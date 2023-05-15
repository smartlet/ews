package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestSyncFolderHierarchy(t *testing.T) {
	rsp, err := service.SyncFolderHierarchy(ews.MakeContext(testSess), &ews.SyncFolderHierarchySoapIn{
		SyncFolderHierarchy: &ews.SyncFolderHierarchyType{
			FolderShape: &ews.FolderResponseShapeType{
				BaseShape: ews.DefaultShapeNamesTypeAllProperties,
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SyncFolderHierarchyResponse.ResponseMessages.SyncFolderHierarchyResponseMessage[0].SyncState)
	changes := rsp.SyncFolderHierarchyResponse.ResponseMessages.SyncFolderHierarchyResponseMessage[0].Changes
	for _, c := range changes.Create {
		fmt.Println(c.Folder)
	}
	for _, c := range changes.Update {
		fmt.Println(c.Folder)
	}
	for _, c := range changes.Delete {
		fmt.Println(c.FolderId)
	}
}
