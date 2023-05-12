package http

import (
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

func (rt *ntlmRoundTripper) RoundTrip(req *http.Request) (rsp *http.Response, err error) {

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
		if rsp, err = rt.tripper.RoundTrip(req); err != nil {
			return nil, err
		}
		// 如果过期失败再执行一次完整校验(先释放原来的response.body
		if rsp.StatusCode == http.StatusUnauthorized {
			kits.DiscardAndCloseBody(rsp.Body)
			rsp, err = rt.authorization(req, sess)
		}
	} else {
		if rsp, err = rt.authorization(req, sess); err != nil {
			return nil, err
		}
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

func (rt *ntlmRoundTripper) authorization(dstReq *http.Request, sess ews.Session) (*http.Response, error) {
	req, err := http.NewRequestWithContext(dstReq.Context(), "POST", sess.GetEndpoint(), nil)
	if err != nil {
		return nil, err
	}
	/*断言肯定支持NTLM*/
	//rsp, err := rt.tripper.RoundTrip(req)
	//if err != nil {
	//	return "", err
	//}
	//kits.DiscardAndCloseBody(rsp.Body)
	//_, ntlm, err := extractWwwAuthenticate(rsp.Header.Values("Www-Authenticate"))
	//if err != nil {
	//	return "", err
	//}
	acc, pwd, err := rt.credential.Get(sess)
	if err != nil {
		return nil, err
	}
	if !rt.withAccountDomain {
		acc = trimAccountDomain(acc)
	}
	user, domain, domainNeed := ntlmssp.GetDomain(acc)
	negotiateMessage, err := ntlmssp.NewNegotiateMessage(domain, "")
	if err != nil {
		return nil, err
	}
	/*断言肯定支持NTLM*/
	//if ntlm {
	req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(negotiateMessage))
	//} else {
	//	req.Header.Set("requestByAuthorization", "Negotiate "+base64.StdEncoding.EncodeToString(negotiateMessage))
	//}
	rsp, err := rt.tripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	kits.DiscardAndCloseBody(rsp.Body)

	challengeMessage, ntlm, err := extractWwwAuthenticate(rsp.Header.Values("Www-Authenticate"))
	if err != nil {
		return nil, err
	}
	authenticateMessage, err := ntlmssp.ProcessChallenge(challengeMessage, user, pwd, domainNeed)
	if ntlm {
		dstReq.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(authenticateMessage))
	} else {
		dstReq.Header.Set("Authorization", "Negotiate "+base64.StdEncoding.EncodeToString(authenticateMessage))
	}

	return rt.tripper.RoundTrip(dstReq)
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
