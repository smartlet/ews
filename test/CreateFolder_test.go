package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	defer dumpFile.Sync()

	rsp, err := service.CreateFolder(ews.MakeContext(testSess), &ews.CreateFolderSoapIn{
		CreateFolder: &ews.CreateFolderType{
			ParentFolderId: &ews.TargetFolderIdType{
				DistinguishedFolderId: &ews.DistinguishedFolderIdType{
					Id: ews.DistinguishedFolderIdNameTypeMsgfolderroot,
				},
			},
			Folders: &ews.NonEmptyArrayOfFoldersType{
				Folder: []*ews.FolderType{
					{DisplayName: "test2"},
					{DisplayName: "test3"},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.CreateFolderResponse.ResponseMessages.CreateFolderResponseMessage[0].ResponseClass)
	fmt.Println(rsp.CreateFolderResponse.ResponseMessages.CreateFolderResponseMessage[0].Folders.Folder[0].FolderId.Id)
}
