package ntlm

import (
	"encoding/base64"
	"github.com/smartlet/ews/kits"
	"golang.org/x/crypto/sha3"
)

var separator = []byte{'\n'}

func Token(endpoint, account, password string) string {
	h := sha3.NewShake256()
	h.Write(kits.UnsafeBytes(endpoint))
	h.Write(separator)
	h.Write(kits.UnsafeBytes(account))
	h.Write(separator)
	h.Write(kits.UnsafeBytes(password))
	r := make([]byte, 64)
	h.Read(r)
	return base64.RawStdEncoding.EncodeToString(r)
}
