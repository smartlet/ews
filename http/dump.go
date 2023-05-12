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
	fmt.Fprintln(d.writer, "\n-------------------------------------")
	fmt.Fprintln(d.writer, req.Method, req.URL)
	for k, v := range req.Header {
		fmt.Fprintln(d.writer, k, ":", v)
	}
	if req.Body != nil {
		if bf, ok := req.Body.(*kits.Buffer); ok {
			d.writer.Write(bf.Data())
		} else {
			kits.Copy(d.writer, req.Body)
		}
	}
	fmt.Fprint(d.writer, "\n\n")
	fmt.Fprintln(d.writer, rsp.Proto, rsp.Status)
	for k, v := range rsp.Header {
		fmt.Fprintln(d.writer, k, ":", v)
	}
	rsp.Body = &dumpReadCloser{writer: d.writer, reader: rsp.Body}

	return rsp, nil
}

var _ http.RoundTripper = (*dumpRoundTripper)(nil)

type dumpReadCloser struct {
	writer io.Writer
	reader io.Reader
}

func (d *dumpReadCloser) Read(p []byte) (int, error) {
	n, err := d.reader.Read(p)
	if n > 0 {
		d.writer.Write(p[0:n])
	} else {
		fmt.Fprintln(d.writer)
	}
	return n, err
}

func (d *dumpReadCloser) Close() error {
	return nil
}

var _ io.ReadCloser = (*dumpReadCloser)(nil)
