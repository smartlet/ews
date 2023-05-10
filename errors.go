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
