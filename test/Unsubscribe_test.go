package test

import (
	"fmt"
	"github.com/smartlet/ews"
	"testing"
)

func TestUnsubscribe(t *testing.T) {
	TestSubscribe(t)
	rsp, err := service.Unsubscribe(ews.MakeContext(testSess), &ews.UnsubscribeSoapIn{
		Unsubscribe: &ews.UnsubscribeType{
			SubscriptionId: SubscriptionId,
		},
	}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.UnsubscribeResponse.ResponseMessages.UnsubscribeResponseMessage[0].ResponseClass)
}
