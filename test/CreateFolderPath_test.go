package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestCreatFolderPath(t *testing.T) {
	defer dumpFile.Sync()

	rsp, err := service.CreateFolderPath(ews.MakeContext(testSess), &ews.CreateFolderPathSoapIn{
		RequestServerVersion: &ews.RequestServerVersionType{
			Version: ews.ExchangeVersionTypeExchange2013,
		},
		MailboxCulture: "en-US",
		TimeZoneContext: &ews.TimeZoneContextType{
			TimeZoneDefinition: &ews.TimeZoneDefinitionType{
				Id: "GMT Standard Time",
			},
		},
		CreateFolderPath: &ews.CreateFolderPathType{
			ParentFolderId: &ews.TargetFolderIdType{
				DistinguishedFolderId: &ews.DistinguishedFolderIdType{
					Id: ews.DistinguishedFolderIdNameTypeDeleteditems,
				},
			},
			RelativeFolderPath: &ews.NonEmptyArrayOfFoldersType{
				Folder: []*ews.FolderType{
					{DisplayName: "MyFirstLevelFolder"},
					{DisplayName: "MySecondLevelFolder"},
					{DisplayName: "MyThirdLevelFolder"},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(rsp.CreateFolderPathResponse.ResponseMessages.CreateFolderPathResponseMessage) == 3)
}
