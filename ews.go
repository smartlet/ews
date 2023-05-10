package ews

import "context"

// Credential 凭证接口
type Credential interface {
	Get(sess Session) (string, string, error)
}

// Authorizer 认证接口
type Authorizer interface {
	Get(sess Session) (string, error)
	Set(sess Session, auth string) error
}

const ContextSession = "_session_"

type Session interface {
	GetId() string
	GetEndpoint() string
}

// AccountSession 企业用户
type AccountSession struct {
	user     string
	corp     string
	account  string
	endpoint string
	_id      string // 临时数据
}

func NewAccountSession(user, corp, account, endpoint string) *AccountSession {
	return &AccountSession{
		user:     user,
		corp:     corp,
		account:  account,
		endpoint: endpoint,
		_id:      user + "/" + corp + "/" + account,
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

func (as *AccountSession) GetEndpoint() string {
	return as.endpoint
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
