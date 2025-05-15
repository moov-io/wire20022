package PaymentStatusRequest

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageHelper struct {
	//Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	MessageId model.ElementHelper
	//Date and time at which the message was created.
	CreatedDateTime model.ElementHelper
	//Point to point reference assigned by the original instructing party to unambiguously identify the original message.
	OriginalMessageId model.ElementHelper
	//Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.
	OriginalMessageNameId model.ElementHelper
	//Original date and time at which the message was created.
	OriginalCreationDateTime model.ElementHelper
	//Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId model.ElementHelper
	//Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId model.ElementHelper
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUETR model.ElementHelper
	//Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent model.AgentHelper
	//Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent model.AgentHelper
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
		OriginalCreationDateTime: model.ElementHelper{
			Title:         "Original Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
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
		InstructingAgent: model.BuildAgentHelper(),
		InstructedAgent:  model.BuildAgentHelper(),
	}
}
