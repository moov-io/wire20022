package DrawdownRequest

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type RemittanceDocumentHelper struct {
	CodeOrProprietary model.ElementHelper
	Number            model.ElementHelper
	RelatedDate       model.ElementHelper
}

func BuildRemittanceDocumentHelper() RemittanceDocumentHelper {
	return RemittanceDocumentHelper{
		CodeOrProprietary: model.ElementHelper{
			Title:         "Code or Proprietary",
			Rules:         "",
			Type:          `CodeOrProprietaryType(CodeCINV, CodeCREQ, CodeCNTR ...)`,
			Documentation: `Provides the type details of the referred document.`,
		},
		Number: model.ElementHelper{
			Title:         "Number",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique and unambiguous identification of the referred document.`,
		},
		RelatedDate: model.ElementHelper{
			Title:         "Related Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date associated with the referred document.`,
		},
	}
}

type CreditTransferTransactionHelper struct {
	PaymentInstructionId  model.ElementHelper
	PaymentEndToEndId     model.ElementHelper
	PaymentUniqueId       model.ElementHelper
	PayRequestType        model.ElementHelper
	PayCategoryType       model.ElementHelper
	Amount                model.CurrencyAndAmountHelper
	ChargeBearer          model.ElementHelper
	CreditorAgent         model.AgentHelper
	Creditor              model.PartyIdentifyHelper
	CrediorAccountOtherId model.ElementHelper
	RemittanceInformation model.ElementHelper
	document              RemittanceDocumentHelper
}

func BuildCreditTransferTransactionHelper() CreditTransferTransactionHelper {
	return CreditTransferTransactionHelper{
		PaymentInstructionId: model.ElementHelper{
			Title:         "Instruction Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35)`,
			Documentation: `Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction. Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.`,
		},
		PaymentEndToEndId: model.ElementHelper{
			Title:         "End To End Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35)`,
			Documentation: `Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain. Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.`,
		},
		PaymentUniqueId: model.ElementHelper{
			Title:         "UETR",
			Rules:         "",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide an end-to-end reference of a payment transaction.`,
		},
		PayRequestType: model.ElementHelper{
			Title:         "Pay Request Type",
			Rules:         "",
			Type:          `PaymentRequestType(DrawDownRequestCredit, DrawDownRequestDebit, IntraCompanyPayment)`,
			Documentation: `Specifies the local instrument, as a proprietary code.`,
		},
		PayCategoryType: model.ElementHelper{
			Title:         "Pay Category Type",
			Rules:         "",
			Type:          `PaymentRequestType(DrawDownRequestCredit, DrawDownRequestDebit, IntraCompanyPayment)`,
			Documentation: `Category purpose, in a proprietary form.`,
		},
		Amount: model.BuildCurrencyAndAmountHelper(),
		ChargeBearer: model.ElementHelper{
			Title:         "Charge Bearer",
			Rules:         "",
			Type:          `ChargeBearerType(ChargeBearerSLEV, ChargeBearerRECV, ChargeBearerSHAR)`,
			Documentation: `Specifies which party/parties will bear the charges associated with the processing of the payment transaction.`,
		},
		CreditorAgent: model.BuildAgentHelper(),
		Creditor:      model.BuildPartyIdentifyHelper(),
		CrediorAccountOtherId: model.ElementHelper{
			Title:         "Credior Account Other Id",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Unique identification of an account, as assigned by the account servicer, using an identification scheme.`,
		},
		RemittanceInformation: model.ElementHelper{
			Title:         "Remittance Information",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.`,
		},
		document: BuildRemittanceDocumentHelper(),
	}
}

type MessageHelper struct {
	MessageId              model.ElementHelper
	CreateDatetime         model.ElementHelper
	NumberofTransaction    model.ElementHelper
	InitiatingParty        model.PartyIdentifyHelper
	PaymentInfoId          model.ElementHelper
	PaymentMethod          model.ElementHelper
	RequestedExecutDate    model.ElementHelper
	Debtor                 model.PartyIdentifyHelper
	DebtorAccountOtherId   model.ElementHelper
	DebtorAgent            model.AgentHelper
	CreditTransTransaction CreditTransferTransactionHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification of the message as assigned by the message originator.`,
		},
		CreateDatetime: model.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		NumberofTransaction: model.ElementHelper{
			Title:         "Number Of Transactions",
			Rules:         "",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `Number of transactions contained in the message.`,
		},
		InitiatingParty: model.BuildPartyIdentifyHelper(),
		PaymentInfoId: model.ElementHelper{
			Title:         "Payment Information Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification as assigned by the instructing party to unambiguously identify the payment information block. Usage: The payment information identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual payment information block. It can be included in several messages related to the payment information block.`,
		},
		PaymentMethod: model.ElementHelper{
			Title:         "Payment Method",
			Rules:         "",
			Type:          `SettlementMethodType(Clearing, Gross, Net, DeferredNet, DeliveryVsPayment, PaymentVsPayment, PaymentVsDelivery, PaymentVsPayment)`,
			Documentation: `Method used to settle a payment transaction.`,
		},
		RequestedExecutDate: model.ElementHelper{
			Title:         "Requested Execution Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date on which the payment transaction is to be executed.`,
		},
		Debtor: model.BuildPartyIdentifyHelper(),
		DebtorAccountOtherId: model.ElementHelper{
			Title:         "Debtor Account Other Id",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Unique identification of an account, as assigned by the account servicer, using an identification scheme.`,
		},
		DebtorAgent:            model.BuildAgentHelper(),
		CreditTransTransaction: BuildCreditTransferTransactionHelper(),
	}
}
