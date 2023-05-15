package test

import (
	"github.com/smartlet/ews"
	"testing"
)

var MyFirstLevelFolderId ews.XsString

func TestFindFolder(t *testing.T) {
	rsp, err := service.FindFolder(ews.MakeContext(testSess), &ews.FindFolderSoapIn{
		FindFolder: &ews.FindFolderType{
			FolderShape: &ews.FolderResponseShapeType{
				BaseShape: ews.DefaultShapeNamesTypeDefault,
			},
			ParentFolderIds: &ews.NonEmptyArrayOfBaseFolderIdsType{
				DistinguishedFolderId: []*ews.DistinguishedFolderIdType{
					{Id: ews.DistinguishedFolderIdNameTypeDeleteditems},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	folders := rsp.FindFolderResponse.ResponseMessages.FindFolderResponseMessage[0].RootFolder.Folders.Folder
	for _, f := range folders {
		if f.DisplayName == "MyFirstLevelFolder" {
			MyFirstLevelFolderId = f.FolderId.Id
			break
		}
	}
	if MyFirstLevelFolderId == "" {
		panic("not found MyFirstLevelFolder")
	}
}
