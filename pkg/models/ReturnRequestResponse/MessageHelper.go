package ReturnRequestResponse

import "github.com/moov-io/wire20022/pkg/models"

type ReasonHelper struct {
	Originator     models.ElementHelper
	Reason         models.ElementHelper
	AdditionalInfo models.ElementHelper
}

func BuildReasonHelper() ReasonHelper {
	return ReasonHelper{
		Originator: models.ElementHelper{
			Title:         "Originator",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Party that issues the cancellation request.`,
		},
		Reason: models.ElementHelper{
			Title:         "Reason",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the reason for the cancellation.`,
		},
		AdditionalInfo: models.ElementHelper{
			Title:         "Additional Info",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Further details on the cancellation request reason.`,
		},
	}
}

type MessageHelper struct {
	AssignmentId                 models.ElementHelper
	Assigner                     models.AgentHelper
	Assignee                     models.AgentHelper
	AssignmentCreateTime         models.ElementHelper
	ResolvedCaseId               models.ElementHelper
	Creator                      models.AgentHelper
	Status                       models.ElementHelper
	OriginalMessageId            models.ElementHelper
	OriginalMessageNameId        models.ElementHelper
	OriginalMessageCreateTime    models.ElementHelper
	OriginalInstructionId        models.ElementHelper
	OriginalEndToEndId           models.ElementHelper
	OriginalUETR                 models.ElementHelper
	CancellationStatusReasonInfo ReasonHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		AssignmentId: models.ElementHelper{
			Title:         "Assignment Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Uniquely identifies the case assignment.`,
		},
		Assigner: models.BuildAgentHelper(),
		Assignee: models.BuildAgentHelper(),
		AssignmentCreateTime: models.ElementHelper{
			Title:         "Assignment Create Time",
			Rules:         "",
			Type:          `ISODateTime`,
			Documentation: `Date and time at which the assignment was created.`,
		},
		ResolvedCaseId: models.ElementHelper{
			Title:         "Resolved Case Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identifies a resolved case.`,
		},
		Creator: models.BuildAgentHelper(),
		Status: models.ElementHelper{
			Title:         "Status",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the status of the investigation, in a coded form.`,
		},
		OriginalMessageId: models.ElementHelper{
			Title:         "Original Message Id",
			Rules:         "",
			Type:          `Status(ReturnRequestAccepted, ReturnRequestRejected ...)`,
			Documentation: `Point to point reference assigned by the original instructing party to unambiguously identify the original message.`,
		},
		OriginalMessageNameId: models.ElementHelper{
			Title:         "Original Message Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.`,
		},
		OriginalMessageCreateTime: models.ElementHelper{
			Title:         "Original Message Create Time",
			Rules:         "",
			Type:          `ISODateTime`,
			Documentation: `Original date and time at which the message was created.`,
		},
		OriginalInstructionId: models.ElementHelper{
			Title:         "Original Instruction Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.`,
		},
		OriginalEndToEndId: models.ElementHelper{
			Title:         "Original End To End Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.`,
		},
		OriginalUETR: models.ElementHelper{
			Title:         "Original UETR",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		CancellationStatusReasonInfo: BuildReasonHelper(),
	}
}
