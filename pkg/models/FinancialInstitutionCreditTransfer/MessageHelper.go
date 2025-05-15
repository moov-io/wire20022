package FinancialInstitutionCreditTransfer

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type CreditTransferTransactionHelper struct {
	Debtor                model.FiniancialInstitutionIdHelper
	DebtorAccount         model.ElementHelper
	DebtorAgent           model.FiniancialInstitutionIdHelper
	CreditorAgent         model.FiniancialInstitutionIdHelper
	Creditor              model.FiniancialInstitutionIdHelper
	CreditorAccount       model.ElementHelper
	RemittanceInformation model.ElementHelper
	InstructedAmount      model.CurrencyAndAmountHelper
}

func BuildCreditTransferTransactionHelper() CreditTransferTransactionHelper {
	return CreditTransferTransactionHelper{
		Debtor: model.FiniancialInstitutionIdHelper{},
		DebtorAccount: model.ElementHelper{
			Title:         "Debtor Account",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Identification assigned by an institution.`,
		},
		DebtorAgent:   model.FiniancialInstitutionIdHelper{},
		CreditorAgent: model.FiniancialInstitutionIdHelper{},
		Creditor:      model.FiniancialInstitutionIdHelper{},
		CreditorAccount: model.ElementHelper{
			Title:         "Creditor Account",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Identification assigned by an institution.`,
		},
		RemittanceInformation: model.ElementHelper{
			Title:         "Remittance Information",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.`,
		},
		InstructedAmount: model.CurrencyAndAmountHelper{},
	}
}

type MessageHelper struct {
	MessageId                        model.ElementHelper
	CreateDateTime                   model.ElementHelper
	NumberOfTransactions             model.ElementHelper
	SettlementMethod                 model.ElementHelper
	ClearingSystem                   model.ElementHelper
	PaymentInstructionId             model.ElementHelper
	PaymentEndToEndId                model.ElementHelper
	PaymentUETR                      model.ElementHelper
	LocalInstrument                  model.ElementHelper
	InterbankSettlementAmount        model.CurrencyAndAmountHelper
	InterbankSettlementDate          model.ElementHelper
	InstructingAgent                 model.AgentHelper
	InstructedAgent                  model.AgentHelper
	Debtor                           model.FiniancialInstitutionIdHelper
	DebtorAgent                      model.FiniancialInstitutionIdHelper
	CreditorAgent                    model.FiniancialInstitutionIdHelper
	Creditor                         model.FiniancialInstitutionIdHelper
	RemittanceInfo                   model.ElementHelper
	UnderlyingCustomerCreditTransfer CreditTransferTransactionHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message. Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.`,
		},
		CreateDateTime: model.ElementHelper{
			Title:         "Create Date Time",
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
			Type:          `SettlementMethodType(SettlementCLRG, SettlementINDA, SettlementCOVE ...)`,
			Documentation: `Method used to settle the (batch of) payment instructions.`,
		},
		ClearingSystem: model.ElementHelper{
			Title:         "Clearing System",
			Rules:         "",
			Type:          `CommonClearingSysCodeType(ClearingSysFDW, ClearingSysCHIPS ... )`,
			Documentation: `Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.`,
		},
		PaymentInstructionId: model.ElementHelper{
			Title:         "Payment Instruction Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction. Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.`,
		},
		PaymentEndToEndId: model.ElementHelper{
			Title:         "Payment End To End Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.`,
		},
		PaymentUETR: model.ElementHelper{
			Title:         "Payment UETR",
			Rules:         "",
			Type:          `UUIDv4 (based on string)`,
			Documentation: `Universally unique identifier to provide an end-to-end reference of a payment transaction.`,
		},
		LocalInstrument: model.ElementHelper{
			Title:         "Local Instrument",
			Rules:         "",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Local instrument code used to identify the local payment scheme or service level.`,
		},
		InterbankSettlementAmount: model.CurrencyAndAmountHelper{},
		InterbankSettlementDate: model.ElementHelper{
			Title:         "Interbank Settlement Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
		InstructingAgent: model.BuildAgentHelper(),
		InstructedAgent:  model.BuildAgentHelper(),
		Debtor:           model.FiniancialInstitutionIdHelper{},
		DebtorAgent:      model.FiniancialInstitutionIdHelper{},
		CreditorAgent:    model.FiniancialInstitutionIdHelper{},
		Creditor:         model.FiniancialInstitutionIdHelper{},
		RemittanceInfo: model.ElementHelper{
			Title:         "Remittance Info",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.`,
		},
		UnderlyingCustomerCreditTransfer: BuildCreditTransferTransactionHelper(),
	}
}
