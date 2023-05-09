package ews

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int
	Status  int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

var _ error = (*Error)(nil)

const (
	CodeInvalidSession    = 50000
	CodeInvalidStatus     = 50001
	CodeInvalidAuthorizer = 50002
	CodeInvalidCredential = 50003
)

var (
	ErrInvalidSession    = &Error{Code: CodeInvalidSession, Status: http.StatusForbidden, Message: "invalid tokenizer"}
	ErrInvalidAuthorizer = &Error{Code: CodeInvalidAuthorizer, Status: http.StatusForbidden, Message: "invalid authorizer"}
	ErrInvalidCredential = &Error{Code: CodeInvalidCredential, Status: http.StatusForbidden, Message: "invalid credential"}
)
