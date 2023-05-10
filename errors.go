package ews

import (
	"github.com/smartlet/ews/wsdl"
)

const (
	CodeInvalidSession    = "InvalidSession"
	CodeInvalidStatus     = "InvalidStatus"
	CodeInvalidAuthorizer = "InvalidAuthorizer"
	CodeInvalidCredential = "InvalidCredential"
)

var (
	ErrInvalidSession    = &wsdl.Fault{FaultCode: CodeInvalidSession, FaultString: "invalid tokenizer"}
	ErrInvalidAuthorizer = &wsdl.Fault{FaultCode: CodeInvalidAuthorizer, FaultString: "invalid authorizer"}
	ErrInvalidCredential = &wsdl.Fault{FaultCode: CodeInvalidCredential, FaultString: "invalid credential"}
)

type FaultDetail struct {
	ResponseCode string `xml:"ResponseCode,omitempty"`
	Message      string `xml:"Message,omitempty"`
}

func (d FaultDetail) Error() string {
	return d.ResponseCode + ":" + d.Message
}
