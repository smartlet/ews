package kits

import "github.com/smartlet/ews"

type MemoryCredential struct {
	data map[string][2]string
}

func NewMemoryCredential(data map[string][2]string) *MemoryCredential {
	return &MemoryCredential{
		data: data,
	}
}

func (m MemoryCredential) Get(sess ews.Session) (string, string) {
	if vs, ok := m.data[sess.GetId()]; ok {
		return vs[0], vs[1]
	}
	return "", ""
}

var _ ews.Credential = (*MemoryCredential)(nil)