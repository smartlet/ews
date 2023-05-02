package http

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type (
	RoundTripper = http.RoundTripper
	Request      = http.Request
	Response     = http.Response
)

// 默认值与go/pkg/http相同
const (
	defaultDialerTimeout       = 30 * time.Second
	defaultDialerKeepAlive     = 30 * time.Second
	defaultIdleConnTimeout     = 25 * time.Second
	defaultTLSHandshakeTimeout = 10 * time.Second
	defaultMaxIdleConnsPerHost = 64
	defaultMaxConnsPerHost     = 2048
	defaultWriteBufferSize     = 512 * 1024
	defaultReadBufferSize      = 512 * 1024
)

type Config struct {
	DialerTimeout       time.Duration `json:"dialer_timeout"`
	DialerKeepAlive     time.Duration `json:"dialer_keep_alive"`
	TLSClientConfig     *tls.Config   `json:"tls_client_config"`
	TLSHandshakeTimeout time.Duration `json:"tls_handshake_timeout"`
	MaxIdleConnsPerHost int           `json:"max_idle_conns_per_host"`
	MaxConnsPerHost     int           `json:"max_conns_per_host"`
	IdleConnTimeout     time.Duration `json:"idle_conn_timeout"`
	WriteBufferSize     int           `json:"write_buffer_size"`
	ReadBufferSize      int           `json:"read_buffer_size"`
	InsecureSkipVerify  bool          `json:"insecure_skip_verify"`
	Username            string        `json:"username"`
	Password            string        `json:"password"`
}

func NewHTTPRoundTripper(c *Config) http.RoundTripper {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   nvl(c.DialerTimeout, defaultDialerTimeout),
			KeepAlive: nvl(c.DialerKeepAlive, defaultDialerKeepAlive),
		}).DialContext,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: c.InsecureSkipVerify},
		TLSHandshakeTimeout: nvl(c.TLSHandshakeTimeout, defaultTLSHandshakeTimeout),
		MaxIdleConnsPerHost: nvl(c.MaxIdleConnsPerHost, defaultMaxIdleConnsPerHost),
		MaxConnsPerHost:     nvl(c.MaxConnsPerHost, defaultMaxConnsPerHost),
		IdleConnTimeout:     nvl(c.IdleConnTimeout, defaultIdleConnTimeout),
		WriteBufferSize:     nvl(c.WriteBufferSize, defaultWriteBufferSize),
		ReadBufferSize:      nvl(c.ReadBufferSize, defaultReadBufferSize),
	}
}

func nvl[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | time.Duration](vs ...T) T {
	for _, v := range vs {
		if v != 0 {
			return v
		}
	}
	return 0
}
