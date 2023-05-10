package kits

import "io"

var discard = make([]byte, 512)

func DiscardAndCloseBody(body io.ReadCloser) {
	// 不要用io.Copy()浪费buffer
	for {
		if _, err := body.Read(discard); err != nil {
			break
		}
	}
	body.Close()
}
