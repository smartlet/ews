package kits

import (
	"github.com/smartlet/ews"
	"strings"
	"sync"
)

type MemoryAuthorizer struct {
	data sync.Map
}

func NewMemoryAuthorizer() *MemoryAuthorizer {
	return new(MemoryAuthorizer)
}

func (m *MemoryAuthorizer) Get(sess ews.Session) (string, string, error) {
	rt, ok := m.data.Load(sess.GetId())
	if ok {
		val, _ := rt.(string)
		idx := strings.IndexByte(val, '\n')
		if idx != -1 {
			return val[:idx], val[idx+1:], nil
		}
	}
	return "", "", nil
}

func (m *MemoryAuthorizer) Set(sess ews.Session, endpoint, auth string) error {
	m.data.Store(sess.GetId(), endpoint+"\n"+auth)
	return nil
}

var _ ews.Authorizer = (*MemoryAuthorizer)(nil)
