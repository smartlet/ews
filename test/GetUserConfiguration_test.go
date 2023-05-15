package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestGetuserConfiguration(t *testing.T) {
	rsp, err := service.GetUserConfiguration(ews.MakeContext(testSess), &ews.GetUserConfigurationSoapIn{
		GetUserConfiguration: &ews.GetUserConfigurationType{
			UserConfigurationName: &ews.UserConfigurationNameType{
				Name: "TestConfig",
				DistinguishedFolderId: &ews.DistinguishedFolderIdType{
					Id: ews.DistinguishedFolderIdNameTypeDeleteditems,
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.GetUserConfigurationResponse.ResponseMessages.GetUserConfigurationResponseMessage[0].UserConfiguration)
}
