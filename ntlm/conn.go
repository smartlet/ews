package ntlm

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/base64"
	"github.com/Azure/go-ntlmssp"
	"github.com/smartlet/ews/kits"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const defaultBufferReaderSize = 8092

type httpConn struct {
	sync.Mutex
	endpoint   *url.URL // 连接端点
	username   string   // 帐号名称
	password   string   // 帐号密码
	retries    int      // 登陆失败重试次数
	_connect   net.Conn
	_cookies   string // 请求的cookie
	_timestamp int64  // 最近时间戳
	_reference int64  // 引用计数(为0时才能释放)
}

// Init 延迟初始化,避免阻塞池子
func (c *httpConn) Init(timeout, keepAlive time.Duration) error {
	if c._connect == nil {
		var err error
		if c.endpoint.Scheme == "https" {
			addr := c.endpoint.Host
			if strings.IndexByte(addr, ':') == -1 {
				addr += ":443"
			}
			c._connect, err = tls.DialWithDialer(&net.Dialer{
				Timeout:   timeout,
				KeepAlive: keepAlive,
			}, "tcp", addr, nil)
			if err != nil {
				return err
			}
		} else {
			addr := c.endpoint.Host
			if strings.IndexByte(addr, ':') == -1 {
				addr += ":80"
			}
			c._connect, err = (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: keepAlive,
			}).Dial("tcp", addr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *httpConn) Reference() *httpConn {
	c.Lock()
	defer c.Unlock()

	c._reference++
	c._timestamp = time.Now().Unix()
	return c
}

func (c *httpConn) Dereference() *httpConn {
	c.Lock()
	defer c.Unlock()

	if c._reference > 0 {
		c._reference--
	}
	c._timestamp = time.Now().Unix()
	return c
}

func (c *httpConn) Close() {
	c.Lock()
	defer c.Unlock()

	if c._connect != nil {
		c._connect.Close()
	}
}

func (c *httpConn) Call(ctx context.Context, header http.Header, buffer *kits.Buffer) (*http.Response, error) {
	c.Lock()
	defer c.Unlock()

	// 创建请求体
	req := kits.NewRequestWithContext(ctx)
	req.URL = c.endpoint
	req.Host = c.endpoint.Host
	req.Method = http.MethodPost
	req.Header = header
	req.Body = buffer
	req.ContentLength = int64(buffer.Len())

	if c._cookies == "" {
		if err := c.sign(req); err != nil {
			return nil, err
		}
	} else {
		req.Header.Set("Cookie", c._cookies)
	}
	rsp, err := c.send(req)
	if err != nil {
		return nil, err
	}
	// 重新登陆
	if rsp.StatusCode == http.StatusUnauthorized {
		for i := 0; i < c.retries; i++ {
			// 清理重置
			rsp.Body.Close()
			buffer.Seek(0, io.SeekStart)
			// 重新登陆
			if err = c.sign(req); err != nil {
				return nil, err
			}
			// 重新请求
			rsp, err = c.send(req)
			if err != nil {
				return nil, err
			}
			if rsp.StatusCode == http.StatusUnauthorized {
				continue
			} else {
				c._cookies = rsp.Header.Get("Set-Cookie")
				break
			}
		}
	} else {
		c._cookies = rsp.Header.Get("Set-Cookie")
	}

	return rsp, nil
}

func (c *httpConn) sign(req *http.Request) (err error) {

	getBody, body, contentLength := req.GetBody, req.Body, req.ContentLength
	req.GetBody = nil
	req.Body = nil
	req.ContentLength = 0

	user, domain, domainNeed := ntlmssp.GetDomain(c.username)
	negotiateMessage, err := ntlmssp.NewNegotiateMessage(domain, "")
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(negotiateMessage))
	err = req.Write(c._connect)
	if err != nil {
		return err
	}
	rsp, err := http.ReadResponse(bufio.NewReaderSize(c._connect, defaultBufferReaderSize), req)
	if err != nil {
		return err
	}
	rsp.Body.Close()

	challengeMessage, err := extractNTLMWwwAuthenticate(rsp.Header.Values("Www-Authenticate"))
	if err != nil {
		return err
	}
	authenticateMessage, err := ntlmssp.ProcessChallenge(challengeMessage, user, c.password, domainNeed)
	req.Header.Set("Authorization", "NTLM "+base64.StdEncoding.EncodeToString(authenticateMessage))

	req.GetBody, req.Body, req.ContentLength = getBody, body, contentLength

	return nil

}

func (c *httpConn) send(req *http.Request) (*http.Response, error) {

	err := req.Write(c._connect)
	if err != nil {
		return nil, err
	}
	rsp, err := http.ReadResponse(bufio.NewReaderSize(c._connect, defaultBufferReaderSize), req)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func extractNTLMWwwAuthenticate(vs []string) ([]byte, error) {
	for _, v := range vs {
		if idx := strings.IndexByte(v, ' '); idx > 0 {
			if v[:idx] == "NTLM" {
				return base64.StdEncoding.DecodeString(v[idx+1:])
			}
		} else {
			if v == "NTLM" {
				return base64.StdEncoding.DecodeString(v[idx+1:])
			}
		}
	}
	return nil, nil
}
