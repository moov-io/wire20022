package ReturnRequest

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type ReasonHelper struct {
	Originator     model.ElementHelper
	Reason         model.ElementHelper
	AdditionalInfo model.ElementHelper
}

func BuildReasonHelper() ReasonHelper {
	return ReasonHelper{
		Originator: model.ElementHelper{
			Title:         "Originator",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Party that issues the cancellation request.`,
		},
		Reason: model.ElementHelper{
			Title:         "Reason",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the reason for the cancellation.`,
		},
		AdditionalInfo: model.ElementHelper{
			Title:         "Additional Info",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Further details on the cancellation request reason.`,
		},
	}
}

type MessageHelper struct {
	AssignmentId                      model.ElementHelper
	Assigner                          model.AgentHelper
	Assignee                          model.AgentHelper
	AssignmentCreateTime              model.ElementHelper
	CaseId                            model.ElementHelper
	Creator                           model.AgentHelper
	OriginalMessageId                 model.ElementHelper
	OriginalMessageNameId             model.ElementHelper
	OriginalMessageCreateTime         model.ElementHelper
	OriginalInstructionId             model.ElementHelper
	OriginalEndToEndId                model.ElementHelper
	OriginalUETR                      model.ElementHelper
	OriginalInterbankSettlementAmount model.CurrencyAndAmountHelper
	OriginalInterbankSettlementDate   model.ElementHelper
	CancellationReason                ReasonHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		AssignmentId: model.ElementHelper{
			Title:         "Assignment Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Uniquely identifies the case assignment.`,
		},
		Assigner: model.BuildAgentHelper(),
		Assignee: model.BuildAgentHelper(),
		AssignmentCreateTime: model.ElementHelper{
			Title:         "Assignment Create Time",
			Rules:         "",
			Type:          `ISODateTime`,
			Documentation: `Date and time at which the assignment was created.`,
		},
		CaseId: model.ElementHelper{
			Title:         "Case Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identifies the investigation case.`,
		},
		Creator: model.BuildAgentHelper(),
		OriginalMessageId: model.ElementHelper{
			Title:         "Original Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference assigned by the original instructing party to unambiguously identify the original message.`,
		},
		OriginalMessageNameId: model.ElementHelper{
			Title:         "Original Message Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.`,
		},
		OriginalMessageCreateTime: model.ElementHelper{
			Title:         "Original Message Create Time",
			Rules:         "",
			Type:          `ISODateTime`,
			Documentation: `Original date and time at which the message was created.`,
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
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		OriginalInterbankSettlementAmount: model.BuildCurrencyAndAmountHelper(),
		OriginalInterbankSettlementDate: model.ElementHelper{
			Title:         "Original Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate`,
			Documentation: `Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
		CancellationReason: BuildReasonHelper(),
	}
}
