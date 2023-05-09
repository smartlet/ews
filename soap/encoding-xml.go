package soap

import (
	"github.com/smartlet/ews/soap/internal/encoding/xml"
	"io"
)

type xmlEncoding struct {
}

func (x *xmlEncoding) Encode(w io.Writer, v interface{}) (string, error) {
	err := xml.NewEncoder(w).Encode(v)
	return `text/xml; charset="utf-8"`, err
}

func (x *xmlEncoding) Decode(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}

var _ MessageEncoding = (*xmlEncoding)(nil)
