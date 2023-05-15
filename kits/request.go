package kits

import (
	"context"
	"net/http"
)

var prototype = new(http.Request)

func NewRequestWithContext(ctx context.Context) *http.Request {
	return prototype.WithContext(ctx)
}
