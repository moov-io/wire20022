package FedwireFundsSystemResponse

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageHelper struct {
	MessageId  model.ElementHelper
	EventCode  model.ElementHelper
	EventParam model.ElementHelper
	EventTime  model.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique and unambiguous identifier for the message, as assigned by the sender.`,
		},
		EventCode: model.ElementHelper{
			Title:         "Event Code",
			Rules:         "",
			Type:          `FundEventType (AdHoc, ConnectionCheck, SystemClosed ...)`,
			Documentation: `Proprietary code used to specify an event that occurred in a system.`,
		},
		EventParam: model.ElementHelper{
			Title:         "Event Parameter",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Describes the parameters of an event which occurred in a system.`,
		},
		EventTime: model.ElementHelper{
			Title:         "Event Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the event occurred.`,
		},
	}
}
