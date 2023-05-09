package http

import (
	"encoding/base64"
	"github.com/Azure/go-ntlmssp"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/kits"
	"net/http"
	"strings"
)

func NewNTLMRoundTripper(tripper http.RoundTripper, auth ews.Authorizer, cred ews.Credential) http.RoundTripper {
	return &ntlmRoundTripper{
		tripper:    tripper,
		authorizer: auth,
		credential: cred,
	}
}

type ntlmRoundTripper struct {
	tripper    http.RoundTripper
	authorizer ews.Authorizer
	credential ews.Credential
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
	if auth != "" {
		req.Header.Set("Cookie", auth)
	}
	rsp, err := rt.tripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusUnauthorized {
		return rsp, err
	}
	req.Body, _ = req.GetBody()
	kits.DiscardAndCloseBody(rsp.Body)

	_, ntlm, err := extractWwwAuthenticate(rsp.Header.Values("Www-Authenticate"))
	if err != nil {
		return nil, err
	}
	acc, pwd, err := rt.credential.Get(sess)
	if err != nil {
		return nil, err
	}
	user, domain, domainNeed := ntlmssp.GetDomain(acc)
	negotiateMessage, err := ntlmssp.NewNegotiateMessage(domain, "")
	if err != nil {
		return nil, err
	}
	if ntlm {
		req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(negotiateMessage))
	} else {
		req.Header.Set("Authorization", "Negotiate "+base64.StdEncoding.EncodeToString(negotiateMessage))
	}
	rsp, err = rt.tripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	req.Body, _ = req.GetBody()
	kits.DiscardAndCloseBody(rsp.Body)

	challengeMessage, ntlm, err := extractWwwAuthenticate(rsp.Header.Values("Www-Authenticate"))
	if err != nil {
		return nil, err
	}
	authenticateMessage, err := ntlmssp.ProcessChallenge(challengeMessage, user, pwd, domainNeed)
	if ntlm {
		req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(authenticateMessage))
	} else {
		req.Header.Set("Authorization", "Negotiate "+base64.StdEncoding.EncodeToString(authenticateMessage))
	}
	rsp, err = rt.tripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusUnauthorized {
		auth = rsp.Header.Get("Set-Cookie")
		err = rt.authorizer.Set(sess, auth)
		if err != nil {
			kits.DiscardAndCloseBody(rsp.Body)
			return nil, err
		}
	}
	return rsp, err
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
