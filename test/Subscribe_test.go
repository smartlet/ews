package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

var SubscriptionId ews.SubscriptionIdType

func TestSubscribe(t *testing.T) {
	dumpFile.Sync()

	rsp, err := service.Subscribe(ews.MakeContext(testSess), &ews.SubscribeSoapIn{
		Subscribe: &ews.SubscribeType{
			StreamingSubscriptionRequest: &ews.StreamingSubscriptionRequestType{
				FolderIds: &ews.NonEmptyArrayOfBaseFolderIdsType{
					DistinguishedFolderId: []*ews.DistinguishedFolderIdType{
						{Id: ews.DistinguishedFolderIdNameTypeInbox},
					},
				},
				EventTypes: &ews.NonEmptyArrayOfNotificationEventTypesType{
					EventType: []ews.NotificationEventTypeType{
						ews.NotificationEventTypeTypeNewMailEvent,
						ews.NotificationEventTypeTypeDeletedEvent,
					},
				},
			},
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	SubscriptionId = rsp.SubscribeResponse.ResponseMessages.SubscribeResponseMessage[0].SubscriptionId
	fmt.Println(SubscriptionId)
}
