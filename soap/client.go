package soap

import (
	"context"
	"fmt"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/kits"
	"github.com/smartlet/ews/soap/internal/encoding/xml"
	"github.com/smartlet/ews/wsdl"
	"io"
	"net/http"
	"unsafe"
)

type MessageEncoding interface {
	Encode(w io.Writer, v interface{}) (string, error)
	Decode(r io.Reader, v interface{}) error
}

func NewSOAPClient(tripper http.RoundTripper) wsdl.SOAPClient {
	return &client{
		tripper:   tripper,
		enconding: new(xmlEncoding),
	}
}

type client struct {
	tripper   http.RoundTripper
	enconding MessageEncoding
}

func (c *client) Call(ctx context.Context, soapAction string, inputHeader any, inputBody any, outputHeader any, outputBody any) error {

	sess := ews.FromContext(ctx)
	if sess == nil {
		return ews.ErrInvalidSession
	}

	buffer := kits.BorrowBuffer()
	defer kits.ReturnBuffer(buffer)

	envelope := &wsdl.Envelope{
		XmlnsS: wsdl.XmlnsS,
		XmlnsM: wsdl.XmlnsM,
		XmlnsT: wsdl.XmlnsT,
		Header: inputHeader,
		Body:   inputBody,
	}
	fmt.Fprint(buffer, xml.Header)
	contentType, err := c.enconding.Encode(buffer, envelope)
	if err != nil {
		return err
	}
	request, err := http.NewRequestWithContext(ctx, "POST", sess.GetEndpoint(), nil)
	if err != nil {
		return err
	}
	request.GetBody = buffer.GetBody
	request.Body = buffer
	request.ContentLength = int64(buffer.Len())
	request.Header.Set("SOAPAction", soapAction)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("User-Agent", "ews")

	response, err := c.tripper.RoundTrip(request)
	if err != nil {
		return err
	}
	defer kits.DiscardAndCloseBody(response.Body)

	if response.StatusCode > http.StatusBadRequest {
		data, _ := io.ReadAll(response.Body)
		return &wsdl.Fault{
			FaultCode:   ews.CodeInvalidStatus,
			FaultString: *(*string)(unsafe.Pointer(&data)),
		}
	}

	envelope.Header = outputHeader
	envelope.Body = outputBody
	err = c.enconding.Decode(response.Body, envelope)
	if err != nil {
		return err
	}
	return nil
}

var _ wsdl.SOAPClient = (*client)(nil)
