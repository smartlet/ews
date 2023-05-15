package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestSendItem(t *testing.T) {
	TestCreateItem_email(t)
	rsp, err := service.SendItem(ews.MakeContext(testSess), &ews.SendItemSoapIn{
		SendItem: &ews.SendItemType{
			ItemIds: &ews.NonEmptyArrayOfBaseItemIdsType{
				ItemId: []*ews.ItemIdType{
					{Id: CreateItemId, ChangeKey: CreateItemChangeKey},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SendItemResponse.ResponseMessages.SendItemResponseMessage[0].ResponseClass)
}
