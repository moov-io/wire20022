package ConnectionCheck

import (
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageHelper struct {
	EventType  Archive.ElementHelper
	EventParam Archive.ElementHelper
	EventTime  Archive.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		EventType: Archive.ElementHelper{
			Title:         "Event Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Proprietary code used to specify an event that occurred in a system.`,
		},
		EventParam: Archive.ElementHelper{
			Title:         "Event Parameter",
			Rules:         "This must contain the Fedwire Funds participant's endpoint (aka logical terminal) used to connect to the Fedwire Funds Service.",
			Type:          `EndpointIdentifier_FedwireFunds_1 (based on string) minLength: 8 maxLength: 8 pattern: [A-Z0-9]{8,8}`,
			Documentation: `Describes the parameters of an event which occurred in a system.`,
		},
		EventTime: Archive.ElementHelper{
			Title:         "Event Time",
			Rules:         "Must be the calendar date and time of the connection check. Time must be in 24-hour clock format and either in Coordinated Universal Time (UTC) or in local time with offset against UTC.",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the event occurred.`,
		},
	}
}
