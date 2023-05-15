package kits

import (
	"github.com/smartlet/ews"
	"sync"
)

type MemoryAuthorizer struct {
	data sync.Map
}

func NewMemoryAuthorizer() *MemoryAuthorizer {
	return new(MemoryAuthorizer)
}

func (m MemoryAuthorizer) Get(sess ews.Session) (string, error) {
	rt, ok := m.data.Load(sess.GetId())
	if ok {
		return rt.(string), nil
	}
	return "", nil
}

func (m MemoryAuthorizer) Set(sess ews.Session, auth string) error {
	m.data.Store(sess.GetId(), auth)
	return nil
}

var _ ews.Authorizer = (*MemoryAuthorizer)(nil)
