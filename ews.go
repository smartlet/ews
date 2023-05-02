package ews

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const contextMetadata = "_contextMetadata_"

// Metadata 上下文元数据
type Metadata struct {
	User    string
	Corp    string
	Account string
}

func With(meta *Metadata) context.Context {
	return context.WithValue(context.Background(), contextMetadata, meta)
}

func From(ctx context.Context) *Metadata {
	if meta, ok := ctx.Value(contextMetadata).(*Metadata); ok {
		return meta
	}
	return nil
}

// Credential 凭证接口
type Credential interface {
	Get(meta *Metadata) (string, string, error)
}

// Authorization 认证接口
type Authorization interface {
	Get(meta *Metadata) (string, error)
	Set(meta *Metadata, auth string) error
}

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
	CodeInvalidMetadata   = 10000
	CodeInvalidStatusCode = 10001
)

var (
	ErrInvalidMetadata = &Error{Code: CodeInvalidMetadata, Status: http.StatusForbidden, Message: "invalid metadata"}
)

func DiscardAndCloseBody(body io.ReadCloser) {
	io.Copy(io.Discard, body)
	body.Close()
}
