package soap

import (
	"context"
	"encoding/base64"
	"github.com/Azure/go-ntlmssp"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/kits"
	"net/http"
	"net/url"
	"strings"
)

func newNTLMClient(tripper http.RoundTripper, auth ews.Authorizer, cred ews.Credential, options *options) HTTPClient {
	return &ntlmClient{
		tripper:           tripper,
		authorizer:        auth,
		credential:        cred,
		withAccountDomain: options.WithAccountDomain,
	}
}

type ntlmClient struct {
	tripper           http.RoundTripper
	authorizer        ews.Authorizer
	credential        ews.Credential
	withAccountDomain bool // account是否带上domain. 默认false
}

func (rt *ntlmClient) Call(ctx context.Context, header http.Header, buffer *kits.Buffer) (rsp *http.Response, err error) {

	req := kits.NewRequestWithContext(ctx)
	req.Method = http.MethodPost
	req.Header = header
	req.Body = buffer
	req.ContentLength = int64(buffer.Len())

	sess := ews.FromContext(ctx)
	if sess == nil {
		return nil, ews.ErrInvalidSession
	}
	authEndpoint, auth, err := rt.authorizer.Get(sess)
	if err != nil {
		return nil, err
	}
	if authEndpoint != "" && auth != "" {
		req.Header.Set("Cookie", auth)
		req.URL, err = url.Parse(authEndpoint)
		if err != nil {
			return nil, err
		}
		req.Host = req.URL.Host

		if rsp, err = rt.tripper.RoundTrip(req); err != nil {
			return nil, err
		}
		if rsp.StatusCode != http.StatusUnauthorized {
			if v := rsp.Header.Get("Set-Cookie"); v != auth {
				if err = rt.authorizer.Set(sess, authEndpoint, v); err != nil {
					kits.DiscardAndCloseBody(rsp.Body)
					return nil, err
				}
			}
			return rsp, nil
		}
		// 必须关闭此前的response后再尝试一次
		kits.DiscardAndCloseBody(rsp.Body)
	}

	credEndpoint, acc, pwd, err := rt.credential.Get(sess)
	if err != nil {
		return nil, err
	}
	if req.URL == nil || authEndpoint != credEndpoint {
		req.URL, err = url.Parse(credEndpoint)
		if err != nil {
			return nil, err
		}
		req.Host = req.URL.Host
	}
	err = rt.authorization(req, acc, pwd)
	if err != nil {
		return nil, err
	}
	rsp, err = rt.tripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusUnauthorized {
		if v := rsp.Header.Get("Set-Cookie"); v != auth {
			if err = rt.authorizer.Set(sess, credEndpoint, v); err != nil {
				kits.DiscardAndCloseBody(rsp.Body)
				return nil, err
			}
		}
	}
	return rsp, nil
}

func (rt *ntlmClient) authorization(req *http.Request, acc, pwd string) error {

	getBody, body, contentLength := req.GetBody, req.Body, req.ContentLength

	req.GetBody = nil
	req.Body = nil
	req.ContentLength = 0

	if !rt.withAccountDomain {
		acc = trimAccountDomain(acc)
	}
	user, domain, domainNeed := ntlmssp.GetDomain(acc)
	negotiateMessage, err := ntlmssp.NewNegotiateMessage(domain, "")
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(negotiateMessage))
	rsp, err := rt.tripper.RoundTrip(req)
	if err != nil {
		return err
	}
	kits.DiscardAndCloseBody(rsp.Body)

	challengeMessage, err := extractNTLMWwwAuthenticate(rsp.Header.Values("Www-Authenticate"))
	if err != nil {
		return err
	}
	authenticateMessage, err := ntlmssp.ProcessChallenge(challengeMessage, user, pwd, domainNeed)
	req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(authenticateMessage))

	req.GetBody, req.Body, req.ContentLength = getBody, body, contentLength

	return nil
}

func extractNTLMWwwAuthenticate(vs []string) ([]byte, error) {
	for _, v := range vs {
		if idx := strings.IndexByte(v, ' '); idx > 0 {
			if v[:idx] == "NTLM" {
				return base64.StdEncoding.DecodeString(v[idx+1:])
			}
		} else {
			if v == "NTLM" {
				return base64.StdEncoding.DecodeString(v[idx+1:])
			}
		}
	}
	return nil, nil
}

func trimAccountDomain(acc string) string {
	idx := strings.IndexByte(acc, '@')
	if idx == -1 {
		return acc
	}
	return acc[:idx]
}
