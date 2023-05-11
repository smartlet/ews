package test

import (
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/http"
	"github.com/smartlet/ews/kits"
	"github.com/smartlet/ews/soap"
	"github.com/smartlet/ews/wsdl"
	"os"
)

// acc 特定测试会话,真实场景是每个context携带的session.业务可以扩展Session接口实现更多特性.
var acc = ews.NewAccountSession(
	"test",
	"test",
	os.Getenv("test.user"), // Exchange用户
	os.Getenv("test.addr"), // Exchange地址. 例如"https://www.office365.com/EWS/Exchange.asmx"
)

// dumpFile

// httpCli http客户端(连接池共享)
var httpCli = http.NewHTTPRoundTripper(new(http.Config))

// dumpCli dump客户端.打印所有request/response详情.仅仅用于调试!
var dumpFile, _ = os.Create(`e:/tmp/dumpFile.log`)
var dumpCli = http.NewDumpRoundTripper(httpCli, dumpFile)

// ntlmCli ntlm客户端.其中Authorizer接口实现Cookie缓存,Credential接口实现Password获取.
var ntlmCli = http.NewNTLMRoundTripper(
	dumpCli,
	kits.NewMemoryAuthorizer(), // 内存实现. 生产环境用redis
	kits.NewMemoryCredential(map[string][2]string{ // 内存实现
		acc.GetId(): {
			os.Getenv("test.user"),
			os.Getenv("test.pass"),
		},
	}),
)

// soapCli soap客户端. 用于所有PortType服务
var soapCli = soap.NewSOAPClient(ntlmCli)

// service portType服务实例
var service = wsdl.NewExchangeServicePortType(soapCli)
