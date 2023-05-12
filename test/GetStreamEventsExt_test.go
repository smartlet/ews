package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestGetStreamingEvents(t *testing.T) {
	defer dumpFile.Sync()

	TestSubscribe(t)

	call := func(evt *ews.GetStreamingEventsSoapOut) error {
		fmt.Println(evt.GetStreamingEventsResponse.ResponseMessages.GetStreamingEventsResponseMessage[0].ResponseClass)
		return nil
	}
	err := service.GetStreamingEventsExt(ews.MakeContext(testSess), &ews.GetStreamingEventsSoapIn{
		GetStreamingEvents: &ews.GetStreamingEventsType{
			SubscriptionIds: &ews.NonEmptyArrayOfSubscriptionIdsType{
				SubscriptionId: []ews.SubscriptionIdType{
					ews.SubscriptionIdType(SubscriptionId),
				},
			},
			ConnectionTimeout: 30,
		},
	}, nil, call)

	if err != nil {
		panic(err)
	}

}
