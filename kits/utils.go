package kits

import "io"

func DiscardAndCloseBody(body io.ReadCloser) {
	io.Copy(io.Discard, body)
	body.Close()
}
