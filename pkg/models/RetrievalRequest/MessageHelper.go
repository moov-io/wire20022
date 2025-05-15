package RetrievalRequest

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageHelper struct {
	//Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	MessageId model.ElementHelper
	//Date and time at which the message was created.
	CreatedDateTime model.ElementHelper
	//Specific actions to be executed through the request.
	RequestType model.ElementHelper
	//Date of the business day of the requested messages the resend function is used for.
	BusinessDate model.ElementHelper
	//Independent counter for a range of message sequences, which are available once per party technical address.
	SequenceRange model.SequenceRangeHelper
	//Unambiguously identifies the original bsiness message, which was delivered by the business sender.
	OriginalMessageNameId model.ElementHelper
	//String of characters that uniquely identifies the file, which was delivered by the sender.
	FileReference model.ElementHelper
	//Unique identification of the party.
	RecipientId model.ElementHelper
	//Entity that assigns the identification.
	RecipientIssuer model.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.`,
		},
		CreatedDateTime: model.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the message was created.`,
		},
		RequestType: model.ElementHelper{
			Title:         "Request Type",
			Rules:         "",
			Type:          `RequestType (RequestReceived, RequestSent)`,
			Documentation: `Specific actions to be executed through the request.`,
		},
		BusinessDate: model.ElementHelper{
			Title:         "Business Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date of the business day of the requested messages the resend function is used for.`,
		},
		SequenceRange: model.BuildSequenceRangeHelper(),
		OriginalMessageNameId: model.ElementHelper{
			Title:         "Original Message Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguously identifies the original business message, which was delivered by the business sender.`,
		},
		FileReference: model.ElementHelper{
			Title:         "File Reference",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `String of characters that uniquely identifies the file, which was delivered by the sender.`,
		},
		RecipientId: model.ElementHelper{
			Title:         "Recipient Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification of the party.`,
		},
		RecipientIssuer: model.ElementHelper{
			Title:         "Recipient Issuer",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Entity that assigns the identification.`,
		},
	}
}
