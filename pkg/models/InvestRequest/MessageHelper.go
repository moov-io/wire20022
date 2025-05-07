package InvestRequest

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type UnderlyingHelper struct {
	OriginalMessageId                 model.ElementHelper
	OriginalMessageNameId             model.ElementHelper
	OriginalCreationDateTime          model.ElementHelper
	OriginalInstructionId             model.ElementHelper
	OriginalEndToEndId                model.ElementHelper
	OriginalUETR                      model.ElementHelper
	OriginalInterbankSettlementAmount model.CurrencyAndAmountHelper
	OriginalInterbankSettlementDate   model.ElementHelper
}

func BuildUnderlyingHelper() UnderlyingHelper {
	return UnderlyingHelper{
		OriginalMessageId: model.ElementHelper{
			Title:         "Original Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.`,
		},
		OriginalMessageNameId: model.ElementHelper{
			Title:         "Original Message Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the original message name identifier to which the message refers.`,
		},
		OriginalCreationDateTime: model.ElementHelper{
			Title:         "Original Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the original message was created.`,
		},
		OriginalInstructionId: model.ElementHelper{
			Title:         "Original Instruction Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.`,
		},
		OriginalEndToEndId: model.ElementHelper{
			Title:         "Original End To End Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.`,
		},
		OriginalUETR: model.ElementHelper{
			Title:         "Original UETR",
			Rules:         "",
			Type:          `UUIDv4 (based on string)`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		OriginalInterbankSettlementAmount: model.BuildCurrencyAndAmountHelper(),
		OriginalInterbankSettlementDate: model.ElementHelper{
			Title:         "Original Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
	}
}

type InvestigationReasonHelper struct {
	Reason                model.ElementHelper
	AdditionalRequestData model.ElementHelper
}

func BuildInvestigationReasonHelper() InvestigationReasonHelper {
	return InvestigationReasonHelper{
		Reason: model.ElementHelper{
			Title:         "Investigation Reason",
			Rules:         "",
			Type:          `ExternalInvestigationReason1Code (based on string)`,
			Documentation: `Reason for the investigation being opened, in a proprietary form.`,
		},
		AdditionalRequestData: model.ElementHelper{
			Title:         "Additional Request Data",
			Rules:         "",
			Type:          `Max500Text (based on string) minLength: 1 maxLength: 500`,
			Documentation: `Additional request data for the investigation.`,
		},
	}
}

type MessageHelper struct {
	MessageId         model.ElementHelper
	InvestigationType model.ElementHelper
	UnderlyingData    UnderlyingHelper
	Requestor         model.AgentHelper
	Responder         model.AgentHelper
	InvestReason      InvestigationReasonHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the sender, to unambiguously identify the message. Usage: The sender has to make sure that MessageIdentification is unique for a pre-agreed period.`,
		},
		InvestigationType: model.ElementHelper{
			Title:         "Investigation Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Type of investigation.`,
		},
		UnderlyingData: BuildUnderlyingHelper(),
		Requestor:      model.BuildAgentHelper(),
		Responder:      model.BuildAgentHelper(),
		InvestReason:   BuildInvestigationReasonHelper(),
	}
}
