package soap

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/smartlet/ews"
	"github.com/smartlet/ews/kits"
)

type HTTPClient interface {
	Call(ctx context.Context, header http.Header, body *kits.Buffer) (*http.Response, error)
}

type Decoder interface {
	Decode(v any) error
}

type Encoding interface {
	Encode(w io.Writer, v interface{}) (string, error)
	Decode(r io.Reader, v interface{}) error
	NewDecoder(r io.Reader) Decoder
}

func NewSOAPClient(httpClient HTTPClient) ews.SOAPClient {
	return &soapClient{
		client:   httpClient,
		encoding: newXmlMessageEncoding(),
	}
}

type soapClient struct {
	client   HTTPClient
	encoding Encoding
}

func (c *soapClient) Call(ctx context.Context, soapAction string, inputHeader any, inputBody any, outputHeader any, outputBody any) error {

	envelope, response, err := c.request(ctx, soapAction, inputHeader, inputBody)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	envelope.Header = outputHeader
	envelope.Body = outputBody
	if response.StatusCode < http.StatusBadRequest {
		// 成功. 正常解码.
		return c.encoding.Decode(response.Body, envelope)
	} else {
		// 错误. 尝试解码.
		buffer := kits.BorrowBuffer()
		defer kits.ReturnBuffer(buffer)

		_, err = kits.Copy(buffer, response.Body)
		if err != nil {
			return err
		}
		err = c.encoding.Decode(buffer, envelope)
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

func (c *soapClient) Stream(ctx context.Context, soapAction string, inputHeader, inputBody any, next func() (any, any, func() error)) error {

	envelope, response, err := c.request(ctx, soapAction, inputHeader, inputBody)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode < http.StatusBadRequest {
		// 链式回调
		dec := c.encoding.NewDecoder(response.Body)
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
		err = c.encoding.Decode(buffer, envelope)
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

func (c *soapClient) request(ctx context.Context, soapAction string, inputHeader, inputBody any) (*ews.Envelope, *http.Response, error) {

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
	contentType, err := c.encoding.Encode(buffer, envelope)
	if err != nil {
		return nil, nil, err
	}

	header := http.Header{
		"SOAPAction":   []string{soapAction},
		"Content-Type": []string{contentType},
		"User-Agent":   []string{"ews"},
	}

	response, err := c.client.Call(ctx, header, buffer)
	if err != nil {
		return nil, nil, err
	}
	return envelope, response, nil
}

var _ ews.SOAPClient = (*soapClient)(nil)
