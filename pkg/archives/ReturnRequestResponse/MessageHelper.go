package ReturnRequestResponse

import (
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type ReasonHelper struct {
	Originator     Archive.ElementHelper
	Reason         Archive.ElementHelper
	AdditionalInfo Archive.ElementHelper
}

func BuildReasonHelper() ReasonHelper {
	return ReasonHelper{
		Originator: Archive.ElementHelper{
			Title:         "Originator",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Party that issues the cancellation request.`,
		},
		Reason: Archive.ElementHelper{
			Title:         "Reason",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the reason for the cancellation.`,
		},
		AdditionalInfo: Archive.ElementHelper{
			Title:         "Additional Info",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Further details on the cancellation request reason.`,
		},
	}
}

type MessageHelper struct {
	AssignmentId                 Archive.ElementHelper
	Assigner                     Archive.AgentHelper
	Assignee                     Archive.AgentHelper
	AssignmentCreateTime         Archive.ElementHelper
	ResolvedCaseId               Archive.ElementHelper
	Creator                      Archive.AgentHelper
	Status                       Archive.ElementHelper
	OriginalMessageId            Archive.ElementHelper
	OriginalMessageNameId        Archive.ElementHelper
	OriginalMessageCreateTime    Archive.ElementHelper
	OriginalInstructionId        Archive.ElementHelper
	OriginalEndToEndId           Archive.ElementHelper
	OriginalUETR                 Archive.ElementHelper
	CancellationStatusReasonInfo ReasonHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		AssignmentId: Archive.ElementHelper{
			Title:         "Assignment Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Uniquely identifies the case assignment.`,
		},
		Assigner: Archive.BuildAgentHelper(),
		Assignee: Archive.BuildAgentHelper(),
		AssignmentCreateTime: Archive.ElementHelper{
			Title:         "Assignment Create Time",
			Rules:         "",
			Type:          `ISODateTime`,
			Documentation: `Date and time at which the assignment was created.`,
		},
		ResolvedCaseId: Archive.ElementHelper{
			Title:         "Resolved Case Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identifies a resolved case.`,
		},
		Creator: Archive.BuildAgentHelper(),
		Status: Archive.ElementHelper{
			Title:         "Status",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the status of the investigation, in a coded form.`,
		},
		OriginalMessageId: Archive.ElementHelper{
			Title:         "Original Message Id",
			Rules:         "",
			Type:          `Status(ReturnRequestAccepted, ReturnRequestRejected ...)`,
			Documentation: `Point to point reference assigned by the original instructing party to unambiguously identify the original message.`,
		},
		OriginalMessageNameId: Archive.ElementHelper{
			Title:         "Original Message Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.`,
		},
		OriginalMessageCreateTime: Archive.ElementHelper{
			Title:         "Original Message Create Time",
			Rules:         "",
			Type:          `ISODateTime`,
			Documentation: `Original date and time at which the message was created.`,
		},
		OriginalInstructionId: Archive.ElementHelper{
			Title:         "Original Instruction Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.`,
		},
		OriginalEndToEndId: Archive.ElementHelper{
			Title:         "Original End To End Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.`,
		},
		OriginalUETR: Archive.ElementHelper{
			Title:         "Original UETR",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		CancellationStatusReasonInfo: BuildReasonHelper(),
	}
}
