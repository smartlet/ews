package soap

import (
	"github.com/smartlet/ews/soap/internal/encoding/xml"
	"io"
)

var xmlHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>`)

const xmlContentType = `text/xml; charset="utf-8"`

func newXmlMessageEncoding(ops *options) Encoding {
	return new(xmlMessageEncoding)
}

type xmlMessageEncoding struct {
}

func (x *xmlMessageEncoding) Encode(w io.Writer, v interface{}) (string, error) {
	err := xml.NewEncoder(w).Encode(v)
	return xmlContentType, err
}

func (x *xmlMessageEncoding) Decode(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}

func (x *xmlMessageEncoding) NewDecoder(r io.Reader) Decoder {
	return xml.NewDecoder(r)
}

var _ Encoding = (*xmlMessageEncoding)(nil)
