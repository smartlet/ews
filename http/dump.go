package http

import (
	"fmt"
	"github.com/smartlet/ews/kits"
	"io"
	"net/http"
)

func NewDumpRoundTripper(tripper http.RoundTripper, writer io.Writer) http.RoundTripper {
	return &dumpRoundTripper{
		tripper: tripper,
		writer:  writer,
	}
}

type dumpRoundTripper struct {
	tripper http.RoundTripper
	writer  io.Writer
}

func (d *dumpRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {

	rsp, err := d.tripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// dump request/response
	fmt.Fprintln(d.writer, req.Method, req.URL)
	for k, v := range req.Header {
		fmt.Fprintln(d.writer, k, ":", v)
	}
	d.writer.Write(req.Body.(*kits.Buffer).Data())
	fmt.Fprintln(d.writer)
	fmt.Fprintln(d.writer)
	fmt.Fprintln(d.writer, rsp.Proto, rsp.Status)
	for k, v := range rsp.Header {
		fmt.Fprintln(d.writer, k, ":", v)
	}
	body := kits.NewBuffer(1024)
	io.Copy(body, rsp.Body)
	d.writer.Write(body.Data())
	rsp.Body = body
	fmt.Fprintln(d.writer)

	return rsp, nil
}

var _ http.RoundTripper = (*dumpRoundTripper)(nil)
