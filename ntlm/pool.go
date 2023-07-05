package ntlm

import (
	"math"
	"net/url"
	"sync"
	"time"
)

type httpPool struct {
	lock      *sync.RWMutex
	cond      *sync.Cond
	cons      map[string]*httpConn
	maximum   int
	retries   int           // 未授权重试次数
	timeout   time.Duration // 连接超时
	keepAlive time.Duration // 连接心跳
}

func initHttpPool(maximum int, retries int, timeout time.Duration, keepAlive time.Duration) *httpPool {
	p := new(httpPool)
	p.lock = new(sync.RWMutex)
	p.cond = sync.NewCond(p.lock)
	p.cons = make(map[string]*httpConn, maximum)
	p.maximum = maximum
	p.retries = retries
	p.timeout = timeout
	p.keepAlive = keepAlive
	return p
}

func (p *httpPool) GetConn(token string) *httpConn {

	p.lock.RLock()
	defer p.lock.RUnlock()

	if cc, ok := p.cons[token]; ok {
		return cc.Reference()
	}
	return nil
}

func (p *httpPool) PutConn(conn *httpConn) {
	if conn != nil {
		p.lock.RLock()
		conn.Dereference()
		p.cond.Broadcast() // 通知其它等待的协程
		p.lock.RUnlock()
	}
}

func (p *httpPool) NewConn(token string, endpoint *url.URL, username, password string) (*httpConn, error) {

	var dropConn *httpConn

	p.lock.Lock()
	// 等待或释放最后一个连接
	for len(p.cons) == p.maximum {

		// 二次检测
		if cc, ok := p.cons[token]; ok {
			p.lock.Unlock()
			return cc.Reference(), nil
		}

		dropStamp := int64(math.MaxInt64)
		dropToken := ""
		for k, v := range p.cons {
			// 没有引用的连接才能释放
			if v._reference == 0 {
				if dropStamp > v._timestamp {
					dropStamp = v._timestamp
					dropToken = k
					dropConn = v
				}
			}
		}
		if dropConn != nil {
			delete(p.cons, dropToken)
		} else {
			p.cond.Wait()
		}
	}
	initConn := &httpConn{
		endpoint: endpoint,
		username: username,
		password: password,
		retries:  p.retries,
	}
	p.cons[token] = initConn
	p.lock.Unlock()

	// 最后退出锁再执行Init()费时操作
	if err := initConn.Init(p.timeout, p.keepAlive); err != nil {
		p.DelConn(token, initConn)
		return nil, err
	}

	// 最后退出锁再执行Close()费时操作
	if dropConn != nil {
		dropConn.Close()
	}

	return initConn.Reference(), nil
}

func (p *httpPool) DelConn(token string, conn *httpConn) {
	p.lock.Lock()
	delete(p.cons, token)
	p.cond.Broadcast() // 通知其它等待的协程
	p.lock.Unlock()

	// 最后退出锁再执行Close()费时操作
	conn.Close()
}
