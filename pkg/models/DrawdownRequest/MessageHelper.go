package DrawdownRequest

import "github.com/moov-io/wire20022/pkg/models"

type RemittanceDocumentHelper struct {
	CodeOrProprietary models.ElementHelper
	Number            models.ElementHelper
	RelatedDate       models.ElementHelper
}

func BuildRemittanceDocumentHelper() RemittanceDocumentHelper {
	return RemittanceDocumentHelper{
		CodeOrProprietary: models.ElementHelper{
			Title:         "Code or Proprietary",
			Rules:         "",
			Type:          `CodeOrProprietaryType(CodeCINV, CodeCREQ, CodeCNTR ...)`,
			Documentation: `Provides the type details of the referred document.`,
		},
		Number: models.ElementHelper{
			Title:         "Number",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique and unambiguous identification of the referred document.`,
		},
		RelatedDate: models.ElementHelper{
			Title:         "Related Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date associated with the referred document.`,
		},
	}
}

type CreditTransferTransactionHelper struct {
	PaymentInstructionId   models.ElementHelper
	PaymentEndToEndId      models.ElementHelper
	PaymentUniqueId        models.ElementHelper
	PayRequestType         models.ElementHelper
	PayCategoryType        models.ElementHelper
	Amount                 models.CurrencyAndAmountHelper
	ChargeBearer           models.ElementHelper
	CreditorAgent          models.AgentHelper
	Creditor               models.PartyIdentifyHelper
	CreditorAccountOtherId models.ElementHelper
	RemittanceInformation  models.ElementHelper
	document               RemittanceDocumentHelper
}

func BuildCreditTransferTransactionHelper() CreditTransferTransactionHelper {
	return CreditTransferTransactionHelper{
		PaymentInstructionId: models.ElementHelper{
			Title:         "Instruction Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35)`,
			Documentation: `Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction. Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.`,
		},
		PaymentEndToEndId: models.ElementHelper{
			Title:         "End To End Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35)`,
			Documentation: `Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain. Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.`,
		},
		PaymentUniqueId: models.ElementHelper{
			Title:         "UETR",
			Rules:         "",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide an end-to-end reference of a payment transaction.`,
		},
		PayRequestType: models.ElementHelper{
			Title:         "Pay Request Type",
			Rules:         "",
			Type:          `PaymentRequestType(DrawDownRequestCredit, DrawDownRequestDebit, IntraCompanyPayment)`,
			Documentation: `Specifies the local instrument, as a proprietary code.`,
		},
		PayCategoryType: models.ElementHelper{
			Title:         "Pay Category Type",
			Rules:         "",
			Type:          `PaymentRequestType(DrawDownRequestCredit, DrawDownRequestDebit, IntraCompanyPayment)`,
			Documentation: `Category purpose, in a proprietary form.`,
		},
		Amount: models.BuildCurrencyAndAmountHelper(),
		ChargeBearer: models.ElementHelper{
			Title:         "Charge Bearer",
			Rules:         "",
			Type:          `ChargeBearerType(ChargeBearerSLEV, ChargeBearerRECV, ChargeBearerSHAR)`,
			Documentation: `Specifies which party/parties will bear the charges associated with the processing of the payment transaction.`,
		},
		CreditorAgent: models.BuildAgentHelper(),
		Creditor:      models.BuildPartyIdentifyHelper(),
		CreditorAccountOtherId: models.ElementHelper{
			Title:         "Credior Account Other Id",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Unique identification of an account, as assigned by the account servicer, using an identification scheme.`,
		},
		RemittanceInformation: models.ElementHelper{
			Title:         "Remittance Information",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.`,
		},
		document: BuildRemittanceDocumentHelper(),
	}
}

type MessageHelper struct {
	MessageId              models.ElementHelper
	CreateDatetime         models.ElementHelper
	NumberofTransaction    models.ElementHelper
	InitiatingParty        models.PartyIdentifyHelper
	PaymentInfoId          models.ElementHelper
	PaymentMethod          models.ElementHelper
	RequestedExecutDate    models.ElementHelper
	Debtor                 models.PartyIdentifyHelper
	DebtorAccountOtherId   models.ElementHelper
	DebtorAgent            models.AgentHelper
	CreditTransTransaction CreditTransferTransactionHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: models.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification of the message as assigned by the message originator.`,
		},
		CreateDatetime: models.ElementHelper{
			Title:         "Created Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		NumberofTransaction: models.ElementHelper{
			Title:         "Number Of Transactions",
			Rules:         "",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `Number of transactions contained in the message.`,
		},
		InitiatingParty: models.BuildPartyIdentifyHelper(),
		PaymentInfoId: models.ElementHelper{
			Title:         "Payment Information Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification as assigned by the instructing party to unambiguously identify the payment information block. Usage: The payment information identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual payment information block. It can be included in several messages related to the payment information block.`,
		},
		PaymentMethod: models.ElementHelper{
			Title:         "Payment Method",
			Rules:         "",
			Type:          `SettlementMethodType(Clearing, Gross, Net, DeferredNet, DeliveryVsPayment, PaymentVsPayment, PaymentVsDelivery, PaymentVsPayment)`,
			Documentation: `Method used to settle a payment transaction.`,
		},
		RequestedExecutDate: models.ElementHelper{
			Title:         "Requested Execution Date",
			Rules:         "",
			Type:          `ISODate (based on date)`,
			Documentation: `Date on which the payment transaction is to be executed.`,
		},
		Debtor: models.BuildPartyIdentifyHelper(),
		DebtorAccountOtherId: models.ElementHelper{
			Title:         "Debtor Account Other Id",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Unique identification of an account, as assigned by the account servicer, using an identification scheme.`,
		},
		DebtorAgent:            models.BuildAgentHelper(),
		CreditTransTransaction: BuildCreditTransferTransactionHelper(),
	}
}
