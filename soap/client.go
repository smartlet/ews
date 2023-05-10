package soap

import (
	"context"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/kits"
	"github.com/smartlet/ews/wsdl"
	"io"
	"net/http"
	"unsafe"
)

type Decoder interface {
	Decode(v any) error
}

type MessageEncoding interface {
	Encode(w io.Writer, v interface{}) (string, error)
	Decode(r io.Reader, v interface{}) error
	NewDecoder(r io.Reader) Decoder
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

	envelope, response, err := c.request(ctx, soapAction, inputHeader, inputBody)
	if err != nil {
		return err
	}
	defer kits.DiscardAndCloseBody(response.Body)

	envelope.Header = outputHeader
	envelope.Body = outputBody
	err = c.enconding.Decode(response.Body, envelope)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) Stream(ctx context.Context, soapAction string, inputHeader, inputBody any, next func() (any, any, func() error)) error {

	envelope, response, err := c.request(ctx, soapAction, inputHeader, inputBody)
	if err != nil {
		return err
	}
	defer kits.DiscardAndCloseBody(response.Body)

	// 链式回调
	dec := c.enconding.NewDecoder(response.Body)
	for {
		outputHeader, outputBody, action := next()
		envelope.Header = outputHeader
		envelope.Body = outputBody
		if err = dec.Decode(envelope); err != nil {
			if err == io.EOF {
				return nil
			} else {
				return err
			}
		}
		if err = action(); err != nil {
			return err
		}
	}
}

func (c *client) request(ctx context.Context, soapAction string, inputHeader, inputBody any) (*wsdl.Envelope, *http.Response, error) {

	sess := ews.FromContext(ctx)
	if sess == nil {
		return nil, nil, ews.ErrInvalidSession
	}

	envelope := &wsdl.Envelope{
		XmlnsS: wsdl.XmlnsS,
		XmlnsM: wsdl.XmlnsM,
		XmlnsT: wsdl.XmlnsT,
		Header: inputHeader,
		Body:   inputBody,
	}

	buffer := kits.BorrowBuffer()
	defer kits.ReturnBuffer(buffer)

	buffer.Write(xmlHeader)
	contentType, err := c.enconding.Encode(buffer, envelope)
	if err != nil {
		return nil, nil, err
	}
	request, err := http.NewRequestWithContext(ctx, "POST", sess.GetEndpoint(), nil)
	if err != nil {
		return nil, nil, err
	}
	request.GetBody = buffer.GetBody
	request.Body = buffer
	request.ContentLength = int64(buffer.Len())
	request.Header.Set("SOAPAction", soapAction)
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("User-Agent", "ews")

	response, err := c.tripper.RoundTrip(request)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode > http.StatusBadRequest {
		data, _ := io.ReadAll(response.Body)
		kits.DiscardAndCloseBody(response.Body)
		return nil, nil, &wsdl.Fault{
			FaultCode:   ews.CodeInvalidStatus,
			FaultString: *(*string)(unsafe.Pointer(&data)),
		}
	}

	return envelope, response, nil
}

var _ wsdl.SOAPClient = (*client)(nil)
