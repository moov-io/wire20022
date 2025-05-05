package FedwireFundsBroadcast

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageHelper struct {
	EventCode        model.ElementHelper
	EventParam       model.ElementHelper
	EventDescription model.ElementHelper
	EventTime        model.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		EventCode: model.ElementHelper{
			Title:         "Event Code",
			Rules:         "",
			Type:          `FundEventType(AdHoc, ConnectionCheck, SystemClosed ...)`,
			Documentation: `Proprietary code used to specify an event that occurred in a system.`,
		},
		EventParam: model.ElementHelper{
			Title:         "Event Parameter",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Describes the parameters of an event which occurred in a system.`,
		},
		EventDescription: model.ElementHelper{
			Title:         "Event Description",
			Rules:         "",
			Type:          `Max1000Text (based on string) minLength: 1 maxLength: 1000`,
			Documentation: `Free text used to describe an event which occurred in a system.`,
		},
		EventTime: model.ElementHelper{
			Title:         "Event Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the event occurred.`,
		},
	}
}
