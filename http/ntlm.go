package http

import (
	"encoding/base64"
	"github.com/Azure/go-ntlmssp"
	"github.com/smartlet/ews"
	"net/http"
	"strings"
)

func NewNTLMRoundTripper(conf *Config, auth ews.Authorizer, cred ews.Credential) http.RoundTripper {
	return &ntlmRoundTripper{
		RoundTripper: NewHTTPRoundTripper(conf),
		Authorizer:   auth,
		Credential:   cred,
	}
}

type ntlmRoundTripper struct {
	http.RoundTripper
	Authorizer ews.Authorizer
	Credential ews.Credential
}

var ng = ntlmssp.Negotiator{}

func (rt *ntlmRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {

	meta := ews.From(req.Context())
	if meta == nil {
		return nil, ews.ErrInvalidMetadata
	}
	auth, err := rt.Authorizer.Get(meta)
	if err != nil {
		return nil, err
	}
	if auth != "" {
		req.Header.Set("Cookie", auth)
	}
	rsp, err := rt.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusUnauthorized {
		return rsp, err
	}
	ews.DiscardAndCloseBody(rsp.Body)
	req.Body, _ = req.GetBody()

	_, ntlm, err := extractWwwAuthenticate(rsp.Header.Values("Www-Authenticate"))
	if err != nil {
		return nil, err
	}
	acc, pwd, err := rt.Credential.Get(meta)
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
	rsp, err = rt.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	ews.DiscardAndCloseBody(rsp.Body)
	req.Body, _ = req.GetBody()

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
	rsp, err = rt.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusUnauthorized {
		auth = rsp.Header.Get("Set-Cookie")
		rt.Authorizer.Set(meta, auth)
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
