package kits

import (
	"reflect"
	"unsafe"
)

var discard = make([]byte, 1024)

// UnsafeString bytes到string覆盖data/len即可.
func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// UnsafeBytes string到bytes覆盖data/len/cap.否则可能出现问题!
func UnsafeBytes(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return
}
