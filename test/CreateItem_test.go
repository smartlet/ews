package test

import (
	"encoding/base64"
	"github.com/smartlet/ews"
	"testing"
)

var CreateItemId ews.XsString
var CreateItemChangeKey ews.XsString

func TestCreateItem_email(t *testing.T) {
	defer dumpFile.Sync()
	rsp, err := service.CreateItem(ews.MakeContext(testSess), &ews.CreateItemSoapIn{
		CreateItem: &ews.CreateItemType{
			SavedItemFolderId: &ews.TargetFolderIdType{
				DistinguishedFolderId: &ews.DistinguishedFolderIdType{
					Id: ews.DistinguishedFolderIdNameTypeSentitems,
				},
			},
			Items: &ews.NonEmptyArrayOfAllItemsType{
				Message: []*ews.MessageType{
					{
						ItemClass: "IPM.Note",
						ToRecipients: &ews.ArrayOfRecipientsType{
							Mailbox: []*ews.EmailAddressType{
								{EmailAddress: "zhaoshi@wahaha.com"},
							},
						},
						MimeContent: &ews.MimeContentType{
							CharacterSet: "utf8",
							CharData:     ews.XsString(base64.StdEncoding.EncodeToString([]byte("this is a test data"))),
						},
						IsRead: XsBoolean(false),
					},
				},
			},
			MessageDisposition: ews.MessageDispositionTypeSaveOnly, // 测试不要使用Send, 否则会自动发送To/CC相关的人
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	item := rsp.CreateItemResponse.ResponseMessages.CreateItemResponseMessage[0].Items.Message[0].ItemId
	CreateItemId = item.Id
	CreateItemChangeKey = item.ChangeKey
	if CreateItemId == "" {
		panic("create failed")
	}

}
