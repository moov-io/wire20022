package PaymentReturn

import "github.com/wadearnold/wire20022/pkg/models"

type PartyHelper struct {
	Name    models.ElementHelper
	Address models.PostalAddressHelper
}

func BuildPartyHelper() PartyHelper {
	return PartyHelper{
		Name: models.ElementHelper{
			Title:         "Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		},
		Address: models.BuildPostalAddressHelper(),
	}
}

type ReturnChainHelper struct {
	Debtor                     PartyHelper
	DebtorOtherTypeId          models.ElementHelper
	DebtorAgent                models.AgentHelper
	CreditorAgent              models.AgentHelper
	Creditor                   PartyHelper
	CreditorAccountOtherTypeId models.ElementHelper
}

func BuildReturnChainHelper() ReturnChainHelper {
	return ReturnChainHelper{
		Debtor: BuildPartyHelper(),
		DebtorOtherTypeId: models.ElementHelper{
			Title:         "Debtor Other Type Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.`,
		},
		DebtorAgent:   models.BuildAgentHelper(),
		CreditorAgent: models.BuildAgentHelper(),
		Creditor:      BuildPartyHelper(),
		CreditorAccountOtherTypeId: models.ElementHelper{
			Title:         "Creditor Account Other Type Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.`,
		},
	}
}

type ReasonHelper struct {
	Reason         models.ElementHelper
	AdditionalInfo models.ElementHelper
}

func BuildReasonHelper() ReasonHelper {
	return ReasonHelper{
		Reason: models.ElementHelper{
			Title:         "Reason",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the reason for the return of the payment transaction.`,
		},
		AdditionalInfo: models.ElementHelper{
			Title:         "Additional Request Data",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Additional information related to the return reason.`,
		},
	}
}

type MessageHelper struct {
	MessageId                         models.ElementHelper
	CreatedDateTime                   models.ElementHelper
	NumberOfTransactions              models.ElementHelper
	SettlementMethod                  models.ElementHelper
	ClearingSystem                    models.ElementHelper
	OriginalMessageId                 models.ElementHelper
	OriginalMessageNameId             models.ElementHelper
	OriginalCreationDateTime          models.ElementHelper
	OriginalInstructionId             models.ElementHelper
	OriginalEndToEndId                models.ElementHelper
	OriginalUETR                      models.ElementHelper
	ReturnedInterbankSettlementAmount models.CurrencyAndAmountHelper
	InterbankSettlementDate           models.ElementHelper
	ReturnedInstructedAmount          models.CurrencyAndAmountHelper
	ChargeBearer                      models.ElementHelper
	InstructingAgent                  models.AgentHelper
	InstructedAgent                   models.AgentHelper
	RtrChain                          ReturnChainHelper
	ReturnReasonInformation           ReasonHelper
	OriginalTransactionRef            models.ElementHelper
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
		NumberOfTransactions: models.ElementHelper{
			Title:         "Number Of Transactions",
			Rules:         "",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `Number of individual transactions contained in the message.`,
		},
		SettlementMethod: models.ElementHelper{
			Title:         "Settlement Method",
			Rules:         "",
			Type:          `SettlementMethodType(SettlementCLRG, SettlementINDA...)`,
			Documentation: `Method used to settle the (batch of) payment instructions.`,
		},
		ClearingSystem: models.ElementHelper{
			Title:         "Clearing System",
			Rules:         "",
			Type:          `CommonClearingSysCodeType(ClearingSysFDW, ClearingSysCHIPS ...)`,
			Documentation: `Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.`,
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
			Type:          `UUIDv4 (based on string)`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		ReturnedInterbankSettlementAmount: models.BuildCurrencyAndAmountHelper(),
		InterbankSettlementDate: models.ElementHelper{
			Title:         "Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
		ReturnedInstructedAmount: models.BuildCurrencyAndAmountHelper(),
		ChargeBearer: models.ElementHelper{
			Title:         "Charge Bearer",
			Rules:         "",
			Type:          `ChargeBearerType(ChargeBearerSLEV, ChargeBearerRECV...)`,
			Documentation: `Specifies which party/parties will bear the charges associated with the processing of the payment transaction.`,
		},
		InstructingAgent:        models.BuildAgentHelper(),
		InstructedAgent:         models.BuildAgentHelper(),
		RtrChain:                BuildReturnChainHelper(),
		ReturnReasonInformation: BuildReasonHelper(),
		OriginalTransactionRef: models.ElementHelper{
			Title:         "Original Transaction Ref",
			Rules:         "",
			Type:          `CommonClearingSysCodeType(ClearingSysFDW, ClearingSysCHIPS ...)`,
			Documentation: `Key elements used to identify the original transaction that is being referred to.`,
		},
	}
}
