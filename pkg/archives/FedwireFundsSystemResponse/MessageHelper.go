package FedwireFundsSystemResponse

import (
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageHelper struct {
	MessageId  Archive.ElementHelper
	EventCode  Archive.ElementHelper
	EventParam Archive.ElementHelper
	EventTime  Archive.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: Archive.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique and unambiguous identifier for the message, as assigned by the sender.`,
		},
		EventCode: Archive.ElementHelper{
			Title:         "Event Code",
			Rules:         "",
			Type:          `FundEventType (AdHoc, ConnectionCheck, SystemClosed ...)`,
			Documentation: `Proprietary code used to specify an event that occurred in a system.`,
		},
		EventParam: Archive.ElementHelper{
			Title:         "Event Parameter",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Describes the parameters of an event which occurred in a system.`,
		},
		EventTime: Archive.ElementHelper{
			Title:         "Event Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the event occurred.`,
		},
	}
}
