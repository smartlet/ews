package http

import (
	"encoding/base64"
	"github.com/azure/go-ntlmssp"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/soap"
	"net/http"
	"strings"
)

func NewNTLMClient(c *Config, auth ews.Authorization, cred ews.Credential) soap.HTTPClient {
	return &ntlmClient{
		http: newHTTPClient(c),
		auth: auth,
		cred: cred,
	}
}

type ntlmClient struct {
	http *http.Client
	auth ews.Authorization
	cred ews.Credential
}

func (cli *ntlmClient) Do(req *http.Request) (*http.Response, error) {
	meta := ews.From(req.Context())
	if meta == nil {
		return nil, ews.ErrInvalidMetadata
	}
	auth, err := cli.auth.Get(meta)
	if err != nil {
		return nil, err
	}
	if auth != "" {
		req.Header.Set("Cookie", auth)
		rsp, err := cli.http.Do(req)
		if err != nil {
			return nil, err
		}
		if rsp.StatusCode != http.StatusUnauthorized {
			return rsp, nil
		}
		ews.DiscardAndCloseBody(rsp.Body) // 立即舍弃与关闭
	}

	rsp, err := cli.authorizate(meta, req)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusUnauthorized {
		auth = rsp.Header.Get("Cookie")
		cli.auth.Set(meta, auth) // 忽略设置错误
	}
	return rsp, nil
}

func (cli *ntlmClient) authorizate(meta *ews.Metadata, req *http.Request) (*http.Response, error) {

	account, password, err := cli.cred.Get(meta)
	if err != nil {
		return nil, err
	}

	realuser, domain, domainNeeded := ntlmssp.GetDomain(account)
	negotiateMsg, err := ntlmssp.NewNegotiateMessage(domain, "")
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(negotiateMsg))

	rsp, err := cli.http.Do(req)
	if err != nil {
		return nil, err
	}
	ews.DiscardAndCloseBody(rsp.Body) // 立即舍弃与关闭

	challengeMsg, err := getChallengeMessage(rsp)
	if err != nil {
		return nil, err
	} else if len(challengeMsg) == 0 {
		return rsp, nil
	}

	authorization, err := ntlmssp.ProcessChallenge(challengeMsg, realuser, password, domainNeeded)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(authorization))
	rsp, err = cli.http.Do(req)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func getChallengeMessage(rsp *http.Response) ([]byte, error) {
	for _, val := range rsp.Header.Values("WWW-Authenticate") {
		if val[:4] == "NTLM" {
			data, err := base64.StdEncoding.DecodeString(strings.TrimSpace(val[4:]))
			if err != nil {
				return nil, err
			} else {
				return data, nil
			}
		}
	}
	return nil, nil
}
