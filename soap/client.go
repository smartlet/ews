package soap

import (
	"context"
	"github.com/smartlet/ews/wsdl"
	"net/http"
)

func NewSOAPClient(tripper http.RoundTripper) wsdl.SOAPClient {
	return &client{
		tripper: tripper,
	}
}

type client struct {
	tripper http.RoundTripper
}

func (c client) Call(ctx context.Context, soapAction string, inputHeader any, inputBody any, outputHeader any, outputBody any, faultDetail any) error {
	envelope := new(Envelope)

}

var _ wsdl.SOAPClient = (*client)(nil)
