package http

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/Azure/go-ntlmssp"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/soap"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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

	rsp, err := rt.roundTrip(req)
	if err != nil {
		return nil, err
	}
	// FOR DEBUG
	buf := new(bytes.Buffer)

	fmt.Fprintln(buf, req.Method, req.URL)
	for k, v := range req.Header {
		fmt.Fprintln(buf, k, ":", v)
	}
	fmt.Fprintln(buf, string(req.Body.(*soap.Buffer).Buff()))
	fmt.Fprintln(buf)
	fmt.Fprintln(buf, rsp.Proto, rsp.Status)
	for k, v := range rsp.Header {
		fmt.Fprintln(buf, k, ":", v)
	}

	rspBuf := new(bytes.Buffer)
	io.Copy(rspBuf, rsp.Body)
	fmt.Fprintln(buf, rspBuf.String())
	rsp.Body = io.NopCloser(rspBuf)

	fmt.Fprintln(buf)
	os.WriteFile(
		filepath.Join("E:\\temp", time.Now().Format("20060102150405")+".txt"),
		buf.Bytes(),
		os.ModePerm)

	return rsp, nil
}

func (rt *ntlmRoundTripper) roundTrip(req *http.Request) (*http.Response, error) {

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
