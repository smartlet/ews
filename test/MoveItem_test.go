package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestMoveItem(t *testing.T) {
	TestFindItem(t)
	rsp, err := service.MoveItem(ews.MakeContext(testSess), &ews.MoveItemSoapIn{
		MoveItem: &ews.MoveItemType{
			ToFolderId: &ews.TargetFolderIdType{
				DistinguishedFolderId: &ews.DistinguishedFolderIdType{
					Id: ews.DistinguishedFolderIdNameTypeDeleteditems,
				},
			},
			ItemIds: &ews.NonEmptyArrayOfBaseItemIdsType{
				ItemId: []*ews.ItemIdType{
					{Id: FindItemId},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.MoveItemResponse.ResponseMessages.MoveItemResponseMessage[0].Items.Message[0].ItemId)
}
