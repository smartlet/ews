package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestEmptyFolder(t *testing.T) {
	TestFindItem(t)
	rsp, err := service.EmptyFolder(ews.MakeContext(testSess), &ews.EmptyFolderSoapIn{
		EmptyFolder: &ews.EmptyFolderType{
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
	fmt.Println(rsp.EmptyFolderResponse.ResponseMessages.EmptyFolderResponseMessage[0].ResponseClass)
}
