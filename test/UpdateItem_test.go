package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestUpdateItem(t *testing.T) {
	TestFindItem(t)
	rsp, err := service.UpdateItem(ews.MakeContext(testSess), &ews.UpdateItemSoapIn{
		UpdateItem: &ews.UpdateItemType{
			ItemChanges: &ews.NonEmptyArrayOfItemChangesType{
				ItemChange: []*ews.ItemChangeType{
					{
						ItemId: &ews.ItemIdType{
							Id:        FindItemId,
							ChangeKey: FindItemChangeKey,
						},
						Updates: &ews.NonEmptyArrayOfItemChangeDescriptionsType{
							AppendToItemField: []*ews.AppendToItemFieldType{
								{
									FieldURI: &ews.PathToUnindexedFieldType{
										FieldURI: ews.UnindexedFieldURITypeItemBody,
									},
									Message: &ews.MessageType{
										Body: &ews.BodyType{
											BodyType: ews.BodyTypeTypeText,
											CharData: "Some additional text to adppend",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.UpdateItemResponse.ResponseMessages.UpdateItemResponseMessage[0].Items.Message[0].ItemId)
}
