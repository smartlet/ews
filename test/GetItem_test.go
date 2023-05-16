package test

import (
	"encoding/base64"
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestGetItem(t *testing.T) {
	TestFindItem(t)
	rsp, err := service.GetItem(ews.MakeContext(testSess), &ews.GetItemSoapIn{
		GetItem: &ews.GetItemType{
			ItemShape: &ews.ItemResponseShapeType{
				BaseShape:          ews.DefaultShapeNamesTypeDefault,
				IncludeMimeContent: XsBoolean(true),
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
	data, _ := base64.StdEncoding.DecodeString(string(rsp.GetItemResponse.ResponseMessages.GetItemResponseMessage[0].Items.Message[0].MimeContent.CharData))
	fmt.Println(string(data))
}
