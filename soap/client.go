package soap

import (
	"context"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/kits"
	"io"
	"net/http"
	"strconv"
)

type Decoder interface {
	Decode(v any) error
}

type MessageEncoding interface {
	Encode(w io.Writer, v interface{}) (string, error)
	Decode(r io.Reader, v interface{}) error
	NewDecoder(r io.Reader) Decoder
}

func NewSOAPClient(tripper http.RoundTripper) ews.SOAPClient {
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
	if response.StatusCode < http.StatusBadRequest {
		// 成功. 正常解码.
		return c.enconding.Decode(response.Body, envelope)
	} else {
		// 错误. 尝试解码.
		buffer := kits.BorrowBuffer()
		defer kits.ReturnBuffer(buffer)

		_, err = kits.Copy(buffer, response.Body)
		if err != nil {
			return err
		}
		err = c.enconding.Decode(buffer, envelope)
		if err != nil {
			return &ews.Fault{
				FaultCode:   ews.CodeInvalidStatus,
				FaultString: string(buffer.Data()),
				FaultActor:  strconv.Itoa(response.StatusCode),
			}
		}
		return nil
	}
}

func (c *client) Stream(ctx context.Context, soapAction string, inputHeader, inputBody any, next func() (any, any, func() error)) error {

	envelope, response, err := c.request(ctx, soapAction, inputHeader, inputBody)
	if err != nil {
		return err
	}
	defer kits.DiscardAndCloseBody(response.Body)

	if response.StatusCode < http.StatusBadRequest {
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
	} else {
		// 错误. 尝试解码.
		outputHeader, outputBody, action := next()
		envelope.Header = outputHeader
		envelope.Body = outputBody

		buffer := kits.BorrowBuffer()
		defer kits.ReturnBuffer(buffer)

		_, err = kits.Copy(buffer, response.Body)
		if err != nil {
			return err
		}
		err = c.enconding.Decode(buffer, envelope)
		if err != nil {
			return &ews.Fault{
				FaultCode:   ews.CodeInvalidStatus,
				FaultString: string(buffer.Data()),
				FaultActor:  strconv.Itoa(response.StatusCode),
			}
		}
		return action()
	}
}

func (c *client) request(ctx context.Context, soapAction string, inputHeader, inputBody any) (*ews.Envelope, *http.Response, error) {

	sess := ews.FromContext(ctx)
	if sess == nil {
		return nil, nil, ews.ErrInvalidSession
	}

	envelope := &ews.Envelope{
		XmlnsS: ews.XmlnsS,
		XmlnsM: ews.XmlnsM,
		XmlnsT: ews.XmlnsT,
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
	return envelope, response, nil
}

var _ ews.SOAPClient = (*client)(nil)
