package ntlm

import (
	"context"
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/kits"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type client struct {
	authorizer ews.Authorizer
	credential ews.Credential
	withDomain bool // account验证是否带有domain. 默认false
	retries    int
	httpPool   *httpPool
}

func NewNTLMClient(authorizer ews.Authorizer, credential ews.Credential, withDomain bool, maximum int, retries int, timeout, keepAlive time.Duration) *client {
	return &client{
		authorizer: authorizer,
		credential: credential,
		withDomain: withDomain,
		retries:    retries,
		httpPool:   initHttpPool(maximum, retries, timeout, keepAlive),
	}
}

func (c *client) Call(ctx context.Context, header http.Header, buffer *kits.Buffer) (*http.Response, error) {
	sess := ews.FromContext(ctx)
	if sess == nil {
		return nil, ews.ErrInvalidSession
	}

	_, token, err := c.authorizer.Get(sess)
	if err != nil {
		return nil, err
	}

	conn := c.httpPool.GetConn(token)
	defer c.httpPool.PutConn(conn)

	if conn == nil {
		endpoint, username, password, err := c.credential.Get(sess)
		if err != nil {
			return nil, err
		}
		epurl, err := url.Parse(endpoint)
		if err != nil {
			return nil, err
		}
		token = Token(endpoint, username, password)
		conn, err = c.httpPool.NewConn(token, epurl, c.User(username), password)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := conn.Call(ctx, header, buffer)
	if err != nil || rsp.StatusCode == http.StatusUnauthorized {
		// 请求错误或者授权错误,立即删除连接,等待下次重建!
		c.httpPool.DelConn(token, conn)
	}
	return rsp, err
}

// User 配置登陆用户是否带上domain部分
func (c *client) User(username string) string {
	if c.withDomain {
		return username
	} else {
		idx := strings.IndexByte(username, '@')
		if idx == -1 {
			return username
		}
		return username[:idx]
	}
}
