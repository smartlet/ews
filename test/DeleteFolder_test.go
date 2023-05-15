package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

// 先创建相应的folder再删除
func TestDeleteFolder(t *testing.T) {
	TestFindFolder(t)
	rsp, err := service.DeleteFolder(ews.MakeContext(testSess), &ews.DeleteFolderSoapIn{
		DeleteFolder: &ews.DeleteFolderType{
			FolderIds: &ews.NonEmptyArrayOfBaseFolderIdsType{
				FolderId: []*ews.FolderIdType{
					{Id: MyFirstLevelFolderId},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.DeleteFolderResponse.ResponseMessages.DeleteFolderResponseMessage[0].ResponseClass)
}
