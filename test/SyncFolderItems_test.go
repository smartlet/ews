package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestSyncFolderItems(t *testing.T) {
	rsp, err := service.SyncFolderItems(ews.MakeContext(testSess), &ews.SyncFolderItemsSoapIn{
		SyncFolderItems: &ews.SyncFolderItemsType{
			ItemShape: &ews.ItemResponseShapeType{
				BaseShape: ews.DefaultShapeNamesTypeDefault,
			},
			SyncFolderId: &ews.TargetFolderIdType{
				DistinguishedFolderId: &ews.DistinguishedFolderIdType{
					Id: ews.DistinguishedFolderIdNameTypeSentitems,
				},
			},
		},
	}, nil)

	if err != nil {
		panic(err)
	}
	changes := rsp.SyncFolderItemsResponse.ResponseMessages.SyncFolderItemsResponseMessage[0].Changes
	for _, c := range changes.Create {
		fmt.Println(c.MeetingRequest)
	}
	for _, c := range changes.Update {
		fmt.Println(c.MeetingRequest)
	}
	for _, c := range changes.Delete {
		fmt.Println(c.ItemId)
	}
}
