package FedwireFundsPaymentStatus

import "github.com/moov-io/wire20022/pkg/models"

type MessageHelper struct {
	MessageId                        models.ElementHelper
	CreatedDateTime                  models.ElementHelper
	OriginalMessageId                models.ElementHelper
	OriginalMessageNameId            models.ElementHelper
	OriginalMessageCreateTime        models.ElementHelper
	OriginalUETR                     models.ElementHelper
	TransactionStatus                models.ElementHelper
	AcceptanceDateTime               models.ElementHelper
	EffectiveInterbankSettlementDate models.ElementHelper
	StatusReasonInformation          models.ElementHelper
	ReasonAdditionalInfo             models.ElementHelper
	InstructingAgent                 models.AgentHelper
	InstructedAgent                  models.AgentHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: models.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message. Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.`,
		},
		CreatedDateTime: models.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		OriginalMessageId: models.ElementHelper{
			Title:         "Original Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.`,
		},
		OriginalMessageNameId: models.ElementHelper{
			Title:         "Original Message Name Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the original message name identifier to which the message refers.`,
		},
		OriginalMessageCreateTime: models.ElementHelper{
			Title:         "Original Message Create Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the original message was created.`,
		},
		OriginalUETR: models.ElementHelper{
			Title:         "Original Unique End To End Transaction Reference",
			Rules:         "",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		TransactionStatus: models.ElementHelper{
			Title:         "Transaction Status",
			Rules:         "",
			Type:          `TransactionStatus1Code`,
			Documentation: `Specifies the status of a transaction, in a coded form.`,
		},
		AcceptanceDateTime: models.ElementHelper{
			Title:         "Acceptance Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.`,
		},
		EffectiveInterbankSettlementDate: models.ElementHelper{
			Title:         "Effective Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date and time at which a transaction is completed and cleared, that is, payment is effected.`,
		},
		StatusReasonInformation: models.ElementHelper{
			Title:         "Status Reason Information",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the reason for the status report.`,
		},
		ReasonAdditionalInfo: models.ElementHelper{
			Title:         "Reason Additional Information",
			Rules:         "",
			Type:          `Max105Text (based on string) minLength: 1 maxLength: 105`,
			Documentation: `Further details on the status reason. Usage: Additional information can be used for several purposes such as the reporting of repaired information.`,
		},
		InstructingAgent: models.BuildAgentHelper(),
		InstructedAgent:  models.BuildAgentHelper(),
	}
}
