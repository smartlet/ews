package test

import (
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/http"
	"github.com/smartlet/ews/kits"
	"github.com/smartlet/ews/soap"
	"os"
)

// testSess 特定测试会话,真实场景是每个context携带的session.业务可以扩展Session接口实现更多特性.
var testSess = ews.NewAccountSession(
	"test",
	"test",
	os.Getenv("test.user"), // Exchange用户
)

// dumpFile

// httpTripper http客户端(连接池共享)
var httpTripper = http.NewHTTPRoundTripper(new(http.Config))

// dumpTripper dump客户端.打印所有request/response详情.仅仅用于调试!
var dumpFile, _ = os.Create(`e:/tmp/dumpFile.log`)
var dumpTripper = http.NewDumpRoundTripper(httpTripper, dumpFile)

// soapCli soap客户端. 用于所有PortType服务
var soapCli = soap.NewSOAPClient(
	soap.NewNTLMClient(
		dumpTripper,
		kits.NewMemoryAuthorizer(), // 内存实现. 生产环境用redis
		kits.NewMemoryCredential(map[string][3]string{ // 内存实现
			testSess.GetId(): {
				os.Getenv("test.addr"), // Exchange地址. 例如"https://www.office365.com/EWS/Exchange.asmx"
				os.Getenv("test.user"),
				os.Getenv("test.pass"),
			},
		}),
		false,
	),
)

// service portType服务实例
var service = ews.NewExchangeServicePortTypeExt(soapCli)

var (
	XsTrue  = true
	XsFalse = false
)

func XsBoolean(c bool) ews.XsBoolean {
	if c {
		return &XsTrue
	} else {
		return &XsFalse
	}
}
