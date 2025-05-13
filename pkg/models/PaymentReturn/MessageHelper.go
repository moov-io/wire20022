package PaymentReturn

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type PartyHelper struct {
	Name    model.ElementHelper
	Address model.PostalAddressHelper
}

func BuildPartyHelper() PartyHelper {
	return PartyHelper{
		Name: model.ElementHelper{
			Title:         "Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		},
		Address: model.BuildPostalAddressHelper(),
	}
}

type ReturnChainHelper struct {
	Debtor                     PartyHelper
	DebtorOtherTypeId          model.ElementHelper
	DebtorAgent                model.AgentHelper
	CreditorAgent              model.AgentHelper
	Creditor                   PartyHelper
	CreditorAccountOtherTypeId model.ElementHelper
}

func BuildReturnChainHelper() ReturnChainHelper {
	return ReturnChainHelper{
		Debtor: BuildPartyHelper(),
		DebtorOtherTypeId: model.ElementHelper{
			Title:         "Debtor Other Type Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.`,
		},
		DebtorAgent:   model.BuildAgentHelper(),
		CreditorAgent: model.BuildAgentHelper(),
		Creditor:      BuildPartyHelper(),
		CreditorAccountOtherTypeId: model.ElementHelper{
			Title:         "Creditor Account Other Type Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.`,
		},
	}
}

type ReasonHelper struct {
	Reason                model.ElementHelper
	AdditionalRequestData model.ElementHelper
}

func BuildReasonHelper() ReasonHelper {
	return ReasonHelper{
		Reason: model.ElementHelper{
			Title:         "Reason",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the reason for the return of the payment transaction.`,
		},
		AdditionalRequestData: model.ElementHelper{
			Title:         "Additional Request Data",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Additional information related to the return reason.`,
		},
	}
}

type MessageHelper struct {
	MessageId                         model.ElementHelper
	CreatedDateTime                   model.ElementHelper
	NumberOfTransactions              model.ElementHelper
	SettlementMethod                  model.ElementHelper
	ClearingSystem                    model.ElementHelper
	OriginalMessageId                 model.ElementHelper
	OriginalMessageNameId             model.ElementHelper
	OriginalCreationDateTime          model.ElementHelper
	OriginalInstructionId             model.ElementHelper
	OriginalEndToEndId                model.ElementHelper
	OriginalUETR                      model.ElementHelper
	ReturnedInterbankSettlementAmount model.CurrencyAndAmountHelper
	InterbankSettlementDate           model.ElementHelper
	ReturnedInstructedAmount          model.CurrencyAndAmountHelper
	ChargeBearer                      model.ElementHelper
	InstructingAgent                  model.AgentHelper
	InstructedAgent                   model.AgentHelper
	RtrChain                          ReturnChainHelper
	ReturnReasonInformation           ReasonHelper
	OriginalTransactionRef            model.ElementHelper
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
		NumberOfTransactions: model.ElementHelper{
			Title:         "Number Of Transactions",
			Rules:         "",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `Number of individual transactions contained in the message.`,
		},
		SettlementMethod: model.ElementHelper{
			Title:         "Settlement Method",
			Rules:         "",
			Type:          `SettlementMethodType(SettlementCLRG, SettlementINDA...)`,
			Documentation: `Method used to settle the (batch of) payment instructions.`,
		},
		ClearingSystem: model.ElementHelper{
			Title:         "Clearing System",
			Rules:         "",
			Type:          `CommonClearingSysCodeType(ClearingSysFDW, ClearingSysCHIPS ...)`,
			Documentation: `Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.`,
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
			Type:          `UUIDv4 (based on string)`,
			Documentation: `Universally unique identifier to provide the original end-to-end reference of a payment transaction.`,
		},
		ReturnedInterbankSettlementAmount: model.BuildCurrencyAndAmountHelper(),
		InterbankSettlementDate: model.ElementHelper{
			Title:         "Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
		ReturnedInstructedAmount: model.BuildCurrencyAndAmountHelper(),
		ChargeBearer: model.ElementHelper{
			Title:         "Charge Bearer",
			Rules:         "",
			Type:          `ChargeBearerType(ChargeBearerSLEV, ChargeBearerRECV...)`,
			Documentation: `Specifies which party/parties will bear the charges associated with the processing of the payment transaction.`,
		},
		InstructingAgent:        model.BuildAgentHelper(),
		InstructedAgent:         model.BuildAgentHelper(),
		RtrChain:                BuildReturnChainHelper(),
		ReturnReasonInformation: BuildReasonHelper(),
		OriginalTransactionRef: model.ElementHelper{
			Title:         "Original Transaction Ref",
			Rules:         "",
			Type:          `CommonClearingSysCodeType(ClearingSysFDW, ClearingSysCHIPS ...)`,
			Documentation: `Key elements used to identify the original transaction that is being referred to.`,
		},
	}
}
