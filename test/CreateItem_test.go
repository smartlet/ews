package test

import (
	"encoding/base64"
	"fmt"
	"github.com/smartlet/ews"
	"os"
	"testing"
)

func TestCreateItem_email(t *testing.T) {
	rsp, err := service.CreateItem(ews.MakeContext(testSess), &ews.CreateItemSoapIn{
		CreateItem: &ews.CreateItemType{
			SavedItemFolderId: &ews.TargetFolderIdType{
				DistinguishedFolderId: &ews.DistinguishedFolderIdType{
					Id: ews.DistinguishedFolderIdNameTypeSsentitems,
				},
			},
			Items: &ews.NonEmptyArrayOfAllItemsType{
				Message: []*ews.MessageType{
					{
						ItemClass: "IPM.Note",
						ToRecipients: &ews.ArrayOfRecipientsType{
							Mailbox: []*ews.EmailAddressType{
								{EmailAddress: "hezhaowu1@ksomail.com"},
								{EmailAddress: "hezhaowu2@wpsemail.com"},
							},
						},
						MimeContent: &ews.MimeContentType{
							CharacterSet: "utf8",
							CharData:     eml(),
						},
						IsRead: false,
					},
				},
			},
			MessageDisposition: ews.MessageDispositionTypeSaveOnly, // 测试不要使用Send
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.CreateItemResponse.ResponseMessages.CreateItemResponseMessage[0])
}

func eml() ews.XsString {
	const path = `C:\Users\jason\Desktop\exchange\测试完整邮件.eml`
	data, _ := os.ReadFile(path)
	return ews.XsString(base64.StdEncoding.EncodeToString(data))
}
