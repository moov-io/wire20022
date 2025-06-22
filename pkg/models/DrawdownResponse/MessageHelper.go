package DrawdownResponse

import "github.com/moov-io/wire20022/pkg/models"

type TransactionInfoAndStatusHelper struct {
	OriginalInstructionId models.ElementHelper
	OriginalEndToEndId    models.ElementHelper
	OriginalUniqueId      models.ElementHelper
	TransactionStatus     models.ElementHelper
	StatusReasonInfoCode  models.ElementHelper
}

func BuildTransactionInfoAndStatusHelper() TransactionInfoAndStatusHelper {
	return TransactionInfoAndStatusHelper{
		OriginalInstructionId: models.ElementHelper{
			Title:         "Original Instruction Identification",
			Rules:         "If used, this should be the Instruction Identification of the original drawdown request (pain.013) message to which this drawdown request response relates.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.`,
		},
		OriginalEndToEndId: models.ElementHelper{
			Title:         "Original End To End Identification",
			Rules:         "If used, this should be the End To End Identification of the original drawdown request (pain.013) message to which this drawdown request response relates.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.`,
		},
		OriginalUniqueId: models.ElementHelper{
			Title:         "Original Unique Identification",
			Rules:         "This should be the UETR of the original drawdown request (pain.013) message to which this drawdown request response relates.",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		TransactionStatus: models.ElementHelper{
			Title:         "Transaction Status",
			Rules:         "",
			Type:          `TransactionStatusCode(MessagesInProcess, MessagesIntercepted ...)`,
			Documentation: `Specifies the status of a transaction, in a coded form.`,
		},
		StatusReasonInfoCode: models.ElementHelper{
			Title:         "Status Reason Information Code",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Provides detailed information on the status reason.`,
		},
	}
}

type MessageHelper struct {
	MessageId                       models.ElementHelper
	CreateDatetime                  models.ElementHelper
	InitiatingParty                 models.PartyIdentifyHelper
	DebtorAgent                     models.AgentHelper
	CreditorAgent                   models.AgentHelper
	OriginalMessageId               models.ElementHelper
	OriginalMessageNameId           models.ElementHelper
	OriginalCreationDateTime        models.ElementHelper
	OriginalPaymentInfoId           models.ElementHelper
	TransactionInformationAndStatus TransactionInfoAndStatusHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: models.ElementHelper{
			Title:         "Message Identification",
			Rules:         "For a drawdown response sent by a Fedwire Funds participant, this must be an Input Message Accountability Data (IMAD). For a drawdown response sent by the Fedwire Funds Service in response to a drawdown request that has failed the Fedwire Funds Service business validation requirements, this is the Fedwire Funds Service application reference.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `A unique identifier (IMADFedwireFunds1) assigned to the message.`,
		},
		CreateDatetime: models.ElementHelper{
			Title:         "Creation Date and Time",
			Rules:         "For the rejection by the Fedwire Funds Service of a drawdown request this is the calendar date and time in New York City (Eastern Time) when the message was rejected by the Fedwire Funds Service application. Time is in a 24-hour clock format and includes the offset against the Coordinated Universal Time (UTC). ",
			Type:          `ISODateTime (based on time)`,
			Documentation: `Date and time at which the status report was created by the instructing party.`,
		},
		InitiatingParty: models.BuildPartyIdentifyHelper(),
		DebtorAgent:     models.BuildAgentHelper(),
		CreditorAgent:   models.BuildAgentHelper(),
		OriginalMessageId: models.ElementHelper{
			Title:         "Original Message Identification",
			Rules:         "This should be the Message Identification of the original drawdown request (pain.013) message to which this drawdown request response message relates.",
			Type:          `IMAD_FedwireFunds_1 (based on string) minLength: 22 maxLength: 22 pattern: [0-9]{8}[A-Z0-9]{8}[0-9]{6}`,
			Documentation: `Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.`,
		},
		OriginalMessageNameId: models.ElementHelper{
			Title:         "Original Message Name Identifier",
			Rules:         "This must be the Message Name Identification of the original drawdown request message to which this drawdown request response message relates (i.e., pain.013.001.07 or a subsequent version of that message as it is introduced in a future release of the Fedwire Funds Service).",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `Specifies the original message name identifier to which the message refers.`,
		},
		OriginalCreationDateTime: models.ElementHelper{
			Title:         "Original Creation Date Time",
			Rules:         "This should be the Creation Date Time of the original drawdown request (pain.013) message to which this drawdown request response relates.",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the original message was created.`,
		},
		OriginalPaymentInfoId: models.ElementHelper{
			Title:         "Original Payment Information Identification",
			Rules:         "This should be the Payment Information Identification of the original drawdown request (pain.013) message to which this drawdown request response relates.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.`,
		},
		TransactionInformationAndStatus: BuildTransactionInfoAndStatusHelper(),
	}
}
