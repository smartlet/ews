package kits

import "io"

var discard = make([]byte, 512)

func DiscardAndCloseBody(body io.ReadCloser) {
	// 不要用io.Copy()浪费buffer
	for {
		if n, _ := body.Read(discard); n == 0 {
			break
		}
	}
	body.Close()
}
