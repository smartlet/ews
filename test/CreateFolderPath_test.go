package test

import (
	"github.com/smartlet/ews"
	"testing"
)

func TestCreatFolderPath(t *testing.T) {
	defer dumpFile.Sync()

	rsp, err := service.CreateFolderPath(ews.MakeContext(testSess), &ews.CreateFolderPathSoapIn{
		RequestServerVersion: &ews.RequestServerVersionType{
			Version: ews.ExchangeVersionTypeExchange2013,
		},
		MailboxCulture: &ews.MailboxCultureType{
			CharData:
		},
	}, nil)

}
