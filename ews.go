package ews

import (
	"context"
)

// Credential 凭证接口
type Credential interface {
	Get(sess Session) (string, string, string, error) //  endpoint, account, password
}

// Authorizer 认证接口
type Authorizer interface {
	Get(sess Session) (string, string, error)
	Set(sess Session, endpoint string, auth string) error
}

const ContextSession = "_session_"

type Session interface {
	GetId() string
}

// AccountSession 企业用户
type AccountSession struct {
	user    string
	corp    string
	account string
	_id     string // 临时数据
}

func NewAccountSession(user, corp, account string) *AccountSession {
	return &AccountSession{
		user:    user,
		corp:    corp,
		account: account,
		_id:     user + "/" + corp + "/" + account,
	}
}

func (as *AccountSession) GetUser() string {
	return as.user
}

func (as *AccountSession) GetCorp() string {
	return as.corp
}

func (as *AccountSession) GetAccount() string {
	return as.account
}

func (as *AccountSession) GetId() string {
	return as._id
}

var _ Session = (*AccountSession)(nil)

func MakeContext(sess Session) context.Context {
	return WithContext(context.Background(), sess)
}

func WithContext(ctx context.Context, sess Session) context.Context {
	return context.WithValue(ctx, ContextSession, sess)
}

func FromContext(ctx context.Context) Session {
	if sess, ok := ctx.Value(ContextSession).(Session); ok {
		return sess
	}
	return nil
}

const (
	CodeInvalidSession    = "InvalidSession"
	CodeInvalidStatus     = "InvalidStatus"
	CodeInvalidAuthorizer = "InvalidAuthorizer"
	CodeInvalidCredential = "InvalidCredential"
)

var (
	ErrInvalidSession    = &Fault{FaultCode: CodeInvalidSession, FaultString: "invalid tokenizer"}
	ErrInvalidAuthorizer = &Fault{FaultCode: CodeInvalidAuthorizer, FaultString: "invalid authorizer"}
	ErrInvalidCredential = &Fault{FaultCode: CodeInvalidCredential, FaultString: "invalid credential"}
)

type FaultDetail struct {
	ResponseCode string `xml:"ResponseCode,omitempty"`
	Message      string `xml:"Message,omitempty"`
}

func (d FaultDetail) Error() string {
	return d.ResponseCode + ":" + d.Message
}
