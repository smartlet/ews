package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestDeleteItem(t *testing.T) {
	TestFindItem(t)
	rsp, err := service.DeleteItem(ews.MakeContext(testSess), &ews.DeleteItemSoapIn{
		DeleteItem: &ews.DeleteItemType{
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
	fmt.Println(rsp.DeleteItemResponse.ResponseMessages.DeleteItemResponseMessage[0].ResponseClass)
}
