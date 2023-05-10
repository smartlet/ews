package wsdl

import "context"

type ExchangeServicePortTypeExt interface {
	ExchangeServicePortType
	GetStreamingEventsExt(ctx context.Context, input *GetStreamingEventsSoapIn, detail error, call func(*GetStreamingEventsSoapOut) error) error
}

type ExchangeServiceBindingExt struct {
	ExchangeServiceBinding
}

func (c *ExchangeServiceBindingExt) GetStreamingEventsExt(ctx context.Context, input *GetStreamingEventsSoapIn, detail error, call func(*GetStreamingEventsSoapOut) error) error {
	var inputHeader any
	if input.ExchangeImpersonation != nil || input.MailboxCulture != nil || input.RequestServerVersion != nil {
		inputHeader = &struct {
			ExchangeImpersonation *ExchangeImpersonationType `xml:"t:ExchangeImpersonation,omitempty"`
			MailboxCulture        *MailboxCultureType        `xml:"t:MailboxCulture,omitempty"`
			RequestServerVersion  *RequestServerVersionType  `xml:"t:RequestServerVersion,omitempty"`
		}{
			ExchangeImpersonation: input.ExchangeImpersonation,
			MailboxCulture:        input.MailboxCulture,
			RequestServerVersion:  input.RequestServerVersion,
		}
	}
	inputBody := &struct {
		GetStreamingEvents *GetStreamingEventsType `xml:"m:GetStreamingEvents,omitempty"`
	}{
		GetStreamingEvents: input.GetStreamingEvents,
	}
	outputHeader := &struct {
		ServerVersionInfo *ServerVersionInfoType `xml:"t:ServerVersionInfo,omitempty"`
	}{}
	outputBody := &struct {
		GetStreamingEventsResponse *GetStreamingEventsResponseType `xml:"m:GetStreamingEventsResponse,omitempty"`
		Fault                      *Fault                          `xml:"s:Fault,omitempty"`
	}{
		Fault: &Fault{Detail: detail},
	}
	action := func() error {
		if outputBody.Fault != nil && outputBody.Fault.FaultCode != "" {
			return outputBody.Fault
		}
		return call(&GetStreamingEventsSoapOut{
			ServerVersionInfo:          outputHeader.ServerVersionInfo,
			GetStreamingEventsResponse: outputBody.GetStreamingEventsResponse,
		})
	}
	return c.client.Stream(ctx, "http://schemas.microsoft.com/exchange/services/2006/messages/GetEvents", inputHeader, inputBody, func() (any, any, func() error) {
		outputHeader.ServerVersionInfo = nil
		outputBody.GetStreamingEventsResponse = nil
		if outputBody.Fault == nil {
			outputBody.Fault = &Fault{Detail: detail}
		}
		return outputHeader, outputBody, action
	})
}

func NewExchangeServicePortTypeExt(client SOAPClient) ExchangeServicePortTypeExt {
	return &ExchangeServiceBindingExt{ExchangeServiceBinding{client: client}}
}
