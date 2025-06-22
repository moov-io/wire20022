package PaymentStatusRequest

import "github.com/moov-io/wire20022/pkg/models"

type MessageHelper struct {
	//Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	MessageId models.ElementHelper
	//Date and time at which the message was created.
	CreatedDateTime models.ElementHelper
	//Point to point reference assigned by the original instructing party to unambiguously identify the original message.
	OriginalMessageId models.ElementHelper
	//Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.
	OriginalMessageNameId models.ElementHelper
	//Original date and time at which the message was created.
	OriginalCreationDateTime models.ElementHelper
	//Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId models.ElementHelper
	//Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId models.ElementHelper
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUETR models.ElementHelper
	//Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent models.AgentHelper
	//Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent models.AgentHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: models.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.`,
		},
		CreatedDateTime: models.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the message was created.`,
		},
		OriginalMessageId: models.ElementHelper{
			Title:         "Original Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference assigned by the original instructing party to unambiguously identify the original message.`,
		},
		OriginalMessageNameId: models.ElementHelper{
			Title:         "Original Message Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.`,
		},
		OriginalCreationDateTime: models.ElementHelper{
			Title:         "Original Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
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
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		InstructingAgent: models.BuildAgentHelper(),
		InstructedAgent:  models.BuildAgentHelper(),
	}
}
