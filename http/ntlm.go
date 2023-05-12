package http

import (
	"context"
	"encoding/base64"
	"github.com/Azure/go-ntlmssp"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/kits"
	"net/http"
	"strings"
)

func NewNTLMRoundTripper(tripper http.RoundTripper, auth ews.Authorizer, cred ews.Credential, ops ...operation) http.RoundTripper {
	rt := &ntlmRoundTripper{
		tripper:    tripper,
		authorizer: auth,
		credential: cred,
	}
	for _, op := range ops {
		op(rt)
	}
	return rt
}

type operation func(c *ntlmRoundTripper)

func WithAccountDomain() operation {
	return func(c *ntlmRoundTripper) {
		c.withAccountDomain = true
	}
}

type ntlmRoundTripper struct {
	tripper           http.RoundTripper
	authorizer        ews.Authorizer
	credential        ews.Credential
	withAccountDomain bool // account是否带上domain. 默认false
}

func (rt *ntlmRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {

	sess := ews.FromContext(req.Context())
	if sess == nil {
		return nil, ews.ErrInvalidSession
	}
	auth, err := rt.authorizer.Get(sess)
	if err != nil {
		return nil, err
	}
	if auth == "" {
		auth, err = rt.authenticate(req.Context(), sess)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Set("Cookie", auth)
	rsp, err := rt.tripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusUnauthorized {
		if v := rsp.Header.Get("Set-Cookie"); v != auth {
			if err = rt.authorizer.Set(sess, auth); err != nil {
				kits.DiscardAndCloseBody(rsp.Body)
				return nil, err
			}
		}
	}
	return rsp, nil
}

func (rt *ntlmRoundTripper) authenticate(ctx context.Context, sess ews.Session) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", sess.GetEndpoint(), nil)
	if err != nil {
		return "", err
	}
	rsp, err := rt.tripper.RoundTrip(req)
	if err != nil {
		return "", err
	}
	kits.DiscardAndCloseBody(rsp.Body)
	_, ntlm, err := extractWwwAuthenticate(rsp.Header.Values("Www-Authenticate"))
	if err != nil {
		return "", err
	}
	acc, pwd, err := rt.credential.Get(sess)
	if err != nil {
		return "", err
	}
	if !rt.withAccountDomain {
		acc = trimAccountDomain(acc)
	}
	user, domain, domainNeed := ntlmssp.GetDomain(acc)
	negotiateMessage, err := ntlmssp.NewNegotiateMessage(domain, "")
	if err != nil {
		return "", err
	}
	if ntlm {
		req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(negotiateMessage))
	} else {
		req.Header.Set("Authorization", "Negotiate "+base64.StdEncoding.EncodeToString(negotiateMessage))
	}
	rsp, err = rt.tripper.RoundTrip(req)
	if err != nil {
		return "", err
	}
	kits.DiscardAndCloseBody(rsp.Body)

	challengeMessage, ntlm, err := extractWwwAuthenticate(rsp.Header.Values("Www-Authenticate"))
	if err != nil {
		return "", err
	}
	authenticateMessage, err := ntlmssp.ProcessChallenge(challengeMessage, user, pwd, domainNeed)
	if ntlm {
		req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(authenticateMessage))
	} else {
		req.Header.Set("Authorization", "Negotiate "+base64.StdEncoding.EncodeToString(authenticateMessage))
	}
	rsp, err = rt.tripper.RoundTrip(req)
	if err != nil {
		return "", err
	}
	if rsp.StatusCode != http.StatusUnauthorized {
		auth := rsp.Header.Get("Set-Cookie")
		if err = rt.authorizer.Set(sess, auth); err != nil {
			kits.DiscardAndCloseBody(rsp.Body)
			return "", err
		}
		return auth, nil
	}
	return "", nil
}

func extractWwwAuthenticate(vs []string) (data []byte, ntlm bool, err error) {
	// NTLM优先Negotiate
	for _, v := range vs {
		if idx := strings.IndexByte(v, ' '); idx > 0 {
			switch v[:idx] {
			case "Negotiate":
				ntlm = false
				data, err = base64.StdEncoding.DecodeString(v[idx+1:])
			case "NTLM":
				ntlm = true
				data, err = base64.StdEncoding.DecodeString(v[idx+1:])
				return
			}
		} else {
			switch v {
			case "Negotiate":
				ntlm = false
			case "NTLM":
				ntlm = true
				return
			}
		}
	}
	return
}

func trimAccountDomain(acc string) string {
	idx := strings.IndexByte(acc, '@')
	if idx == -1 {
		return acc
	}
	return acc[:idx]
}
