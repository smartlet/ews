package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

// 可能阻塞
func TestGetStreamEvents(t *testing.T) {
	TestSubscribe(t)
	rsp, err := service.GetStreamingEvents(ews.MakeContext(testSess), &ews.GetStreamingEventsSoapIn{
		GetStreamingEvents: &ews.GetStreamingEventsType{
			SubscriptionIds: &ews.NonEmptyArrayOfSubscriptionIdsType{
				SubscriptionId: []ews.SubscriptionIdType{
					SubscriptionId,
				},
			},
			ConnectionTimeout: 30,
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.GetStreamingEventsResponse.ResponseMessages.GetStreamingEventsResponseMessage[0].MessageText)
}
