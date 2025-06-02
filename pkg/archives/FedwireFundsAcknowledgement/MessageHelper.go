package FedwireFundsAcknowledgement

import (
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageHelper struct {
	MessageId         Archive.ElementHelper
	CreatedDateTime   Archive.ElementHelper
	RelationReference Archive.ElementHelper
	ReferenceName     Archive.ElementHelper
	RequestHandling   Archive.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: Archive.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the sender, to unambiguously identify the message. Usage: The sender has to make sure that MessageIdentification is unique for a pre-agreed period.`,
		},
		CreatedDateTime: Archive.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		RelationReference: Archive.ElementHelper{
			Title:         "Relation Reference",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguous reference to a previous message having a business relevance with this message.`,
		},
		ReferenceName: Archive.ElementHelper{
			Title:         "Reference Name",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Name of the message which contained the given additional reference as its message reference.`,
		},
		RequestHandling: Archive.ElementHelper{
			Title:         "Request Handling",
			Rules:         "",
			Type:          `RelatedStatusCode(SchemaValidationFailed, MessageHeaderIssue, BusinessRuleViolation, UnknownMessageType)`,
			Documentation: `Specifies the status of the request, for example the result of the schema validation or a business processing result/error.`,
		},
	}
}
