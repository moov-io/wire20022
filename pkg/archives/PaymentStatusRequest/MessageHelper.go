package PaymentStatusRequest

import (
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageHelper struct {
	//Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	MessageId Archive.ElementHelper
	//Date and time at which the message was created.
	CreatedDateTime Archive.ElementHelper
	//Point to point reference assigned by the original instructing party to unambiguously identify the original message.
	OriginalMessageId Archive.ElementHelper
	//Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.
	OriginalMessageNameId Archive.ElementHelper
	//Original date and time at which the message was created.
	OriginalCreationDateTime Archive.ElementHelper
	//Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId Archive.ElementHelper
	//Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId Archive.ElementHelper
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUETR Archive.ElementHelper
	//Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent Archive.AgentHelper
	//Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent Archive.AgentHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: Archive.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.`,
		},
		CreatedDateTime: Archive.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the message was created.`,
		},
		OriginalMessageId: Archive.ElementHelper{
			Title:         "Original Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference assigned by the original instructing party to unambiguously identify the original message.`,
		},
		OriginalMessageNameId: Archive.ElementHelper{
			Title:         "Original Message Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.`,
		},
		OriginalCreationDateTime: Archive.ElementHelper{
			Title:         "Original Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
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
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		InstructingAgent: Archive.BuildAgentHelper(),
		InstructedAgent:  Archive.BuildAgentHelper(),
	}
}
