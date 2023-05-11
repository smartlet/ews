package test

import (
	"encoding/base64"
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

// 通过FindItem获取AAAWAGhlemhhb3d1MUBraW5nc29mdC5jb20ARgAAAAAAC9AYe7rimkOhvpHgvNsOzgcAOjhiyJgKLEO13EB8uUUY2gAAAAABDAAAgoNscah2/UO0tMiYJBHiWQAB0psqxAAA
const itemId = "AAAWAGhlemhhb3d1MUBraW5nc29mdC5jb20ARgAAAAAAC9AYe7rimkOhvpHgvNsOzgcAOjhiyJgKLEO13EB8uUUY2gAAAAABDAAAgoNscah2/UO0tMiYJBHiWQAB0psqxAAA"

func TestCopyItem(t *testing.T) {
	defer dumpFile.Sync()
	rsp, err := service.CopyItem(ews.MakeContext(testSess), &ews.CopyItemSoapIn{
		CopyItem: &ews.CopyItemType{
			ToFolderId: &ews.TargetFolderIdType{
				DistinguishedFolderId: &ews.DistinguishedFolderIdType{
					Id: ews.DistinguishedFolderIdNameTypeDdeleteditems,
				},
			},
			ItemIds: &ews.NonEmptyArrayOfBaseItemIdsType{
				ItemId: []*ews.ItemIdType{
					{Id: itemId},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.CopyItemResponse.ResponseMessages.CopyItemResponseMessage[0].ResponseClass)
	fmt.Println(rsp.CopyItemResponse.ResponseMessages.CopyItemResponseMessage[0].Items.Message[0].ItemId.Id)
	fmt.Println(rsp.CopyItemResponse.ResponseMessages.CopyItemResponseMessage[0].Items.Message[0].ItemId.Id == itemId)
}

func TestIdInfo(t *testing.T) {
	data, err := base64.RawStdEncoding.DecodeString(itemId)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
