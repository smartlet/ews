package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

var FindItemId ews.XsString
var FindItemChangeKey ews.XsString

func TestFindItem(t *testing.T) {
	rsp, err := service.FindItem(ews.MakeContext(testSess), &ews.FindItemSoapIn{
		FindItem: &ews.FindItemType{
			ItemShape: &ews.ItemResponseShapeType{
				BaseShape: ews.DefaultShapeNamesTypeIdOnly,
			},
			ParentFolderIds: &ews.NonEmptyArrayOfBaseFolderIdsType{
				DistinguishedFolderId: []*ews.DistinguishedFolderIdType{
					{Id: ews.DistinguishedFolderIdNameTypeInbox},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	// 对应xml的结果 FindItemResponse>ResponseMessages>FindItemResponseMessage*>RootFolder>Items>Message*>ItemId>Id
	FindItemId = rsp.FindItemResponse.ResponseMessages.FindItemResponseMessage[0].RootFolder.Items.Message[0].ItemId.Id
	FindItemChangeKey = rsp.FindItemResponse.ResponseMessages.FindItemResponseMessage[0].RootFolder.Items.Message[0].ItemId.ChangeKey

	fmt.Println(FindItemId)
}
