package test

import (
	"github.com/smartlet/ews"
	"github.com/smartlet/ews/kits"
	"github.com/smartlet/ews/ntlm"
	"github.com/smartlet/ews/soap"
	"os"
	"time"
)

// testSess 特定测试会话,真实场景是每个context携带的session.业务可以扩展Session接口实现更多特性.
var testSess = ews.NewAccountSession(
	"test",
	"test",
	os.Getenv("test.user"), // Exchange用户
)

// dumpFile
var dumpFile, _ = os.Create("/tmp/dump.log")

// ntlmCli
var ntlmCli = ntlm.NewNTLMClient(
	kits.NewMemoryAuthorizer(), // 内存实现. 生产环境用redis
	kits.NewMemoryCredential(map[string][3]string{ // 内存实现
		testSess.GetId(): {
			os.Getenv("test.addr"), // Exchange地址. 例如"https://www.office365.com/EWS/Exchange.asmx"
			os.Getenv("test.user"),
			os.Getenv("test.pass"),
		},
	}),
	false,
	3,
	0,
	0,
	30*time.Second,
)

// soapCli soap客户端. 用于所有PortType服务
var soapCli = soap.NewSOAPClient(ntlmCli)

// service portType服务实例
var service = ews.NewExchangeServicePortTypeExt(soapCli)
