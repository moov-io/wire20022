package FedwireFundsPaymentStatus

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageHelper struct {
	MessageId                        model.ElementHelper
	CreatedDateTime                  model.ElementHelper
	OriginalMessageId                model.ElementHelper
	OriginalMessageNameId            model.ElementHelper
	OriginalMessageCreateTime        model.ElementHelper
	OriginalUETR                     model.ElementHelper
	TransactionStatus                model.ElementHelper
	AcceptanceDateTime               model.ElementHelper
	EffectiveInterbankSettlementDate model.ElementHelper
	StatusReasonInformation          model.ElementHelper
	ReasonAdditionalInfo             model.ElementHelper
	InstructingAgent                 model.AgentHelper
	InstructedAgent                  model.AgentHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message. Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.`,
		},
		CreatedDateTime: model.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		OriginalMessageId: model.ElementHelper{
			Title:         "Original Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.`,
		},
		OriginalMessageNameId: model.ElementHelper{
			Title:         "Original Message Name Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the original message name identifier to which the message refers.`,
		},
		OriginalMessageCreateTime: model.ElementHelper{
			Title:         "Original Message Create Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the original message was created.`,
		},
		OriginalUETR: model.ElementHelper{
			Title:         "Original Unique End To End Transaction Reference",
			Rules:         "",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		TransactionStatus: model.ElementHelper{
			Title:         "Transaction Status",
			Rules:         "",
			Type:          `TransactionStatus1Code`,
			Documentation: `Specifies the status of a transaction, in a coded form.`,
		},
		AcceptanceDateTime: model.ElementHelper{
			Title:         "Acceptance Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.`,
		},
		EffectiveInterbankSettlementDate: model.ElementHelper{
			Title:         "Effective Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date and time at which a transaction is completed and cleared, that is, payment is effected.`,
		},
		StatusReasonInformation: model.ElementHelper{
			Title:         "Status Reason Information",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the reason for the status report.`,
		},
		ReasonAdditionalInfo: model.ElementHelper{
			Title:         "Reason Additional Information",
			Rules:         "",
			Type:          `Max105Text (based on string) minLength: 1 maxLength: 105`,
			Documentation: `Further details on the status reason. Usage: Additional information can be used for several purposes such as the reporting of repaired information.`,
		},
		InstructingAgent: model.BuildAgentHelper(),
		InstructedAgent:  model.BuildAgentHelper(),
	}
}
