package soap

import (
	"github.com/smartlet/ews/soap/internal/encoding/xml"
	"io"
)

var xmlHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>`)

const xmlContentType = `text/xml; charset="utf-8"`

type xmlEncoding struct {
}

func (x *xmlEncoding) Encode(w io.Writer, v interface{}) (string, error) {
	err := xml.NewEncoder(w).Encode(v)
	return xmlContentType, err
}

func (x *xmlEncoding) Decode(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}

func (x *xmlEncoding) NewDecoder(r io.Reader) Decoder {
	return xml.NewDecoder(r)
}

var _ MessageEncoding = (*xmlEncoding)(nil)
