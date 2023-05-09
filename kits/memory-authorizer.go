package kits

import "github.com/smartlet/ews"

type MemoryAuthorizer struct {
	data map[string]string
}

func NewMemoryAuthorizer() *MemoryAuthorizer {
	return &MemoryAuthorizer{
		data: make(map[string]string),
	}
}

func (m MemoryAuthorizer) Get(sess ews.Session) (string, error) {
	return m.data[sess.GetId()], nil
}

func (m MemoryAuthorizer) Set(sess ews.Session, auth string) error {
	m.data[sess.GetId()] = auth
	return nil
}

var _ ews.Authorizer = (*MemoryAuthorizer)(nil)
