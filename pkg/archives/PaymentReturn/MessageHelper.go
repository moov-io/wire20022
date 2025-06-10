package PaymentReturn

import (
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type PartyHelper struct {
	Name    Archive.ElementHelper
	Address Archive.PostalAddressHelper
}

func BuildPartyHelper() PartyHelper {
	return PartyHelper{
		Name: Archive.ElementHelper{
			Title:         "Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		},
		Address: Archive.BuildPostalAddressHelper(),
	}
}

type ReturnChainHelper struct {
	Debtor                     PartyHelper
	DebtorOtherTypeId          Archive.ElementHelper
	DebtorAgent                Archive.AgentHelper
	CreditorAgent              Archive.AgentHelper
	Creditor                   PartyHelper
	CreditorAccountOtherTypeId Archive.ElementHelper
}

func BuildReturnChainHelper() ReturnChainHelper {
	return ReturnChainHelper{
		Debtor: BuildPartyHelper(),
		DebtorOtherTypeId: Archive.ElementHelper{
			Title:         "Debtor Other Type Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.`,
		},
		DebtorAgent:   Archive.BuildAgentHelper(),
		CreditorAgent: Archive.BuildAgentHelper(),
		Creditor:      BuildPartyHelper(),
		CreditorAccountOtherTypeId: Archive.ElementHelper{
			Title:         "Creditor Account Other Type Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.`,
		},
	}
}

type ReasonHelper struct {
	Reason         Archive.ElementHelper
	AdditionalInfo Archive.ElementHelper
}

func BuildReasonHelper() ReasonHelper {
	return ReasonHelper{
		Reason: Archive.ElementHelper{
			Title:         "Reason",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the reason for the return of the payment transaction.`,
		},
		AdditionalInfo: Archive.ElementHelper{
			Title:         "Additional Request Data",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Additional information related to the return reason.`,
		},
	}
}

type MessageHelper struct {
	MessageId                         Archive.ElementHelper
	CreatedDateTime                   Archive.ElementHelper
	NumberOfTransactions              Archive.ElementHelper
	SettlementMethod                  Archive.ElementHelper
	ClearingSystem                    Archive.ElementHelper
	OriginalMessageId                 Archive.ElementHelper
	OriginalMessageNameId             Archive.ElementHelper
	OriginalCreationDateTime          Archive.ElementHelper
	OriginalInstructionId             Archive.ElementHelper
	OriginalEndToEndId                Archive.ElementHelper
	OriginalUETR                      Archive.ElementHelper
	ReturnedInterbankSettlementAmount Archive.CurrencyAndAmountHelper
	InterbankSettlementDate           Archive.ElementHelper
	ReturnedInstructedAmount          Archive.CurrencyAndAmountHelper
	ChargeBearer                      Archive.ElementHelper
	InstructingAgent                  Archive.AgentHelper
	InstructedAgent                   Archive.AgentHelper
	RtrChain                          ReturnChainHelper
	ReturnReasonInformation           ReasonHelper
	OriginalTransactionRef            Archive.ElementHelper
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
		NumberOfTransactions: Archive.ElementHelper{
			Title:         "Number Of Transactions",
			Rules:         "",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `Number of individual transactions contained in the message.`,
		},
		SettlementMethod: Archive.ElementHelper{
			Title:         "Settlement Method",
			Rules:         "",
			Type:          `SettlementMethodType(SettlementCLRG, SettlementINDA...)`,
			Documentation: `Method used to settle the (batch of) payment instructions.`,
		},
		ClearingSystem: Archive.ElementHelper{
			Title:         "Clearing System",
			Rules:         "",
			Type:          `CommonClearingSysCodeType(ClearingSysFDW, ClearingSysCHIPS ...)`,
			Documentation: `Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.`,
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
			Type:          `UUIDv4 (based on string)`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		ReturnedInterbankSettlementAmount: Archive.BuildCurrencyAndAmountHelper(),
		InterbankSettlementDate: Archive.ElementHelper{
			Title:         "Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
		ReturnedInstructedAmount: Archive.BuildCurrencyAndAmountHelper(),
		ChargeBearer: Archive.ElementHelper{
			Title:         "Charge Bearer",
			Rules:         "",
			Type:          `ChargeBearerType(ChargeBearerSLEV, ChargeBearerRECV...)`,
			Documentation: `Specifies which party/parties will bear the charges associated with the processing of the payment transaction.`,
		},
		InstructingAgent:        Archive.BuildAgentHelper(),
		InstructedAgent:         Archive.BuildAgentHelper(),
		RtrChain:                BuildReturnChainHelper(),
		ReturnReasonInformation: BuildReasonHelper(),
		OriginalTransactionRef: Archive.ElementHelper{
			Title:         "Original Transaction Ref",
			Rules:         "",
			Type:          `CommonClearingSysCodeType(ClearingSysFDW, ClearingSysCHIPS ...)`,
			Documentation: `Key elements used to identify the original transaction that is being referred to.`,
		},
	}
}
