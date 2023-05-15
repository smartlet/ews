package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

var folderId ews.XsString

func TestGetFolder(t *testing.T) {
	defer dumpFile.Sync()
	rsp, err := service.GetFolder(ews.MakeContext(testSess), &ews.GetFolderSoapIn{
		GetFolder: &ews.GetFolderType{
			FolderShape: &ews.FolderResponseShapeType{
				BaseShape: ews.DefaultShapeNamesTypeDefault,
			},
			FolderIds: &ews.NonEmptyArrayOfBaseFolderIdsType{
				DistinguishedFolderId: []*ews.DistinguishedFolderIdType{
					{Id: ews.DistinguishedFolderIdNameTypeInbox},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	folderId = rsp.GetFolderResponse.ResponseMessages.GetFolderResponseMessage[0].Folders.Folder[0].FolderId.Id
	fmt.Println(folderId)
}

func TestGetFolder_error(t *testing.T) {
	defer dumpFile.Sync()

	detail := new(ews.FaultDetail)
	rsp, err := service.GetFolder(ews.MakeContext(testSess), &ews.GetFolderSoapIn{
		GetFolder: &ews.GetFolderType{
			FolderShape: &ews.FolderResponseShapeType{
				BaseShape: ews.DefaultShapeNamesTypeDefault,
			},
			FolderIds: &ews.NonEmptyArrayOfBaseFolderIdsType{
				DistinguishedFolderId: []*ews.DistinguishedFolderIdType{
					{Id: ews.DistinguishedFolderIdNameTypeInbox},
				},
			},
		},
	}, detail)
	if err != nil {
		fmt.Println(detail)
		fmt.Println(err)
	} else {
		fmt.Println(rsp.ServerVersionInfo)
		fmt.Println(rsp.GetFolderResponse)
	}

}
