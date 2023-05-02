package soap

import (
	"context"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/wsdl"
	"io"
	"net/http"
	"strings"
	"unsafe"
)

const (
	MTOM = "MTOM"
	MMA  = "MMA"
)

type Config struct {
	Endpoint string `json:"endpoint"`
	Encoding string `json:"encoding"`
}

func NewClient(c *Config, h http.RoundTripper) wsdl.SOAPClient {
	var encoding MessageEncoding
	switch strings.ToUpper(c.Encoding) {
	case MTOM:
	case MMA:
	default:
		encoding = new(xmlEncoding)
	}
	return &client{
		http:     h,
		endpoint: c.Endpoint,
		encoding: encoding,
	}
}

type client struct {
	http     http.RoundTripper
	endpoint string
	encoding MessageEncoding
}

// Call performs HTTP POST request with a context
func (s *client) Call(ctx context.Context, soapAction string, request, response interface{}) error {
	// SOAP envelope capable of namespace prefixes
	envelope := &SOAPEnvelope{
		XmlNS: XmlNsSoapEnv,
	}
	envelope.Body.Content = request

	buffer := borrowBuffer()
	defer returnBuffer(buffer)

	contentType, err := s.encoding.Encode(buffer, envelope)
	if err != nil {
		return err
	}

	// 不指定body
	req, err := http.NewRequestWithContext(ctx, "POST", s.endpoint, nil)
	if err != nil {
		return err
	}
	req.GetBody = func() (io.ReadCloser, error) {
		buffer.Seek(0, io.SeekStart)
		return buffer, nil
	}
	req.Body = buffer
	req.ContentLength = int64(buffer.Len())
	req.Header.Set("SOAPAction", soapAction)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", "ews")

	rsp, err := s.http.RoundTrip(req)
	if err != nil {
		return err
	}
	defer ews.DiscardAndCloseBody(rsp.Body)

	if rsp.StatusCode > 400 {
		data, _ := io.ReadAll(rsp.Body)
		return &ews.Error{
			Code:    ews.CodeInvalidStatusCode,
			Status:  rsp.StatusCode,
			Message: *(*string)(unsafe.Pointer(&data)),
		}
	}

	// xml Decoder (used with and without MTOM) cannot handle namespace prefixes (yet),
	// so we have to use a namespace-less response envelope
	rspEnvelope := new(SOAPEnvelopeResponse)
	rspEnvelope.Body = SOAPBodyResponse{
		Content: response,
		Fault: &SOAPFault{
			Detail: nil, // TODO: support detail
		},
	}

	err = s.encoding.Decode(rsp.Body, rspEnvelope)
	if err != nil {
		return err
	}

	return rspEnvelope.Body.ErrorFromFault()
}
