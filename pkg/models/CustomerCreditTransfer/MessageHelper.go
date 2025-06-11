package CustomerCreditTransfer

import "github.com/moov-io/wire20022/pkg/models"

type ChargeInfoHelper struct {
	Amount         models.CurrencyAndAmountHelper
	BusinessIdCode models.ElementHelper
}

func BuildChargeInfoHelper() ChargeInfoHelper {
	return ChargeInfoHelper{
		Amount: models.BuildCurrencyAndAmountHelper(),
		BusinessIdCode: models.ElementHelper{
			Title:         "Business Identifier Code",
			Rules:         "Must be the ISO 9362 Business Identifier Code (BIC) of the party to which the charge is to be paid.",
			Type:          `BICIdentifier (based on string) pattern: [A-Z]{6}[A-Z0-9]{2}([A-Z0-9]{3})?`,
			Documentation: `Standard code used to uniquely identify a financial institution or other entity in a financial transaction.`,
		},
	}
}

type RemittanceDetailHelper struct {
	RemittanceId      models.ElementHelper
	Method            models.ElementHelper
	ElectronicAddress models.ElementHelper
}

func BuildRemittanceDetailHelper() RemittanceDetailHelper {
	return RemittanceDetailHelper{
		RemittanceId: models.ElementHelper{
			Title:         "Remittance Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the initiating party, to unambiguously identify the remittance information sent separately from the payment instruction, such as a remittance advice.`,
		},
		Method: models.ElementHelper{
			Title:         "Method",
			Rules:         "",
			Type:          `RemittanceDeliveryMethod(Fax, ElectronicDataInterchange, UniformResourceIdentifier ...)`,
			Documentation: `Method used to deliver the remittance advice information.`,
		},
		ElectronicAddress: models.ElementHelper{
			Title:         "Electronic Address",
			Rules:         "",
			Type:          `Max2048Text (based on string) minLength: 1 maxLength: 2048`,
			Documentation: `Electronic address to which an agent is to send the remittance information.`,
		},
	}
}

type TaxRecordHelper struct {
	TaxId              models.ElementHelper
	TaxTypeCode        models.ElementHelper
	TaxPeriodYear      models.ElementHelper
	TaxperiodTimeFrame models.ElementHelper
}

func BuildTaxRecordHelper() TaxRecordHelper {
	return TaxRecordHelper{
		TaxId: models.ElementHelper{
			Title:         "Tax Identification",
			Rules:         "For IRS tax payments, i.e., if Instructed Agent contains one of the Treasury tax payment RTNs, a Tax Identification Number (TIN) or Employer Identification Number (EIN) of exactly 9 numeric characters (excluding '000000000' and '999999999') must be provided in the tax remittance component (i.e., RemittanceInformation/Structured/TaxRemittance/Creditor/Tax Identification).",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Tax identification number of the creditor.`,
		},
		TaxTypeCode: models.ElementHelper{
			Title:         "Tax Type Code",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Type of tax payer.`,
		},
		TaxPeriodYear: models.ElementHelper{
			Title:         "Tax Period Year",
			Rules:         "For IRS tax payments, i.e., if Instructed Agent contains one of the Treasury tax payment RTNs, a Tax Year of exactly 4 numerical characters must be provided in the tax remittance component (i.e., RemittanceInformation/Structured/TaxRemittance/Record/Period/Year). Note: to obtain a valid Tax Year, the year YYYY should appear as YYYY-12-31.",
			Type:          `ISODate (based on date)`,
			Documentation: `Year related to the tax payment.`,
		},
		TaxperiodTimeFrame: models.ElementHelper{
			Title:         "Tax period Time Frame",
			Rules:         "For IRS tax payments, i.e., if Instructed Agent contains one of the Treasury tax payment RTNs, a Tax Month must be provided in the tax remittance component (i.e.,  RemittanceInformation/Structured/TaxRemittance/Record/Period/Type) and must contain one of the following 4 alphanumeric character codes: MM01 (January), MM02 (February), MM03 (March), MM04 (April), MM05 (May), MM06 (June), MM07 (July), MM08 (August), MM09 (September), MM10 (October), MM11 (November) or MM12 (December). ",
			Type:          `TaxRecordPeriod1Code (based on string)`,
			Documentation: `Range of time between a start date and an end date for which the tax report is provided.`,
		},
	}
}

type RemittanceDocumentHelper struct {
	UnstructuredRemitInfo models.ElementHelper
	CodeOrProprietary     models.ElementHelper
	Number                models.ElementHelper
	RelatedDate           models.ElementHelper
	TaxDetail             TaxRecordHelper
}

func BuildRemittanceDocumentHelper() RemittanceDocumentHelper {
	return RemittanceDocumentHelper{
		UnstructuredRemitInfo: models.ElementHelper{
			Title:         "Unstructured",
			Rules:         "Unstructured and Structured remittance information must not be combined.",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.`,
		},
		CodeOrProprietary: models.ElementHelper{
			Title:         "Code Or Proprietary",
			Rules:         "",
			Type:          `CodeOrProprietaryType(CodeCINV, CodeCREQ, CodeCNTR...)`,
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
		TaxDetail: BuildTaxRecordHelper(),
	}
}

type MessageHelper struct {
	MessageId                    models.ElementHelper
	CreatedDateTime              models.ElementHelper
	NumberOfTransactions         models.ElementHelper
	SettlementMethod             models.ElementHelper
	CommonClearingSysCode        models.ElementHelper
	InstructionId                models.ElementHelper
	EndToEndId                   models.ElementHelper
	UniqueEndToEndTransactionRef models.ElementHelper
	SericeLevel                  models.ElementHelper
	InstrumentPropCode           models.ElementHelper
	InterBankSettAmount          models.CurrencyAndAmountHelper
	InterBankSettDate            models.ElementHelper
	InstructedAmount             models.CurrencyAndAmountHelper
	exchangeRate                 models.ElementHelper
	ChargeBearer                 models.ElementHelper
	ChargesInfo                  ChargeInfoHelper
	InstructingAgents            models.AgentHelper
	InstructedAgent              models.AgentHelper
	IntermediaryAgent1Id         models.ElementHelper
	UltimateDebtorName           models.ElementHelper
	UltimateDebtorAddress        models.PostalAddressHelper
	DebtorName                   models.ElementHelper
	DebtorAddress                models.PostalAddressHelper
	DebtorIBAN                   models.ElementHelper
	DebtorOtherTypeId            models.ElementHelper
	DebtorAgent                  models.AgentHelper
	CreditorAgent                models.AgentHelper
	CreditorName                 models.ElementHelper
	CreditorPostalAddress        models.PostalAddressHelper
	UltimateCreditorName         models.ElementHelper
	UltimateCreditorAddress      models.PostalAddressHelper
	CreditorIBAN                 models.ElementHelper
	CreditorOtherTypeId          models.ElementHelper
	PurposeOfPayment             models.ElementHelper
	RelatedRemittanceInfo        RemittanceDetailHelper
	RemittanceInfor              RemittanceDocumentHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: models.ElementHelper{
			Title:         "Message Identification",
			Rules:         "Must be the Fedwire Funds Input Message Accountability Data (IMAD).",
			Type:          `IMAD_FedwireFunds_1 (based on string) minLength: 22 maxLength: 22 pattern: [0-9]{8}[A-Z0-9]{8}[0-9]{6}`,
			Documentation: `Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message. Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.`,
		},
		CreatedDateTime: models.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "Must be date and time when the message is created by the Fedwire Sender. Time must be in 24-hour clock format and either in Coordinated Universal Time (UTC) or in local time with offset against UTC.",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		NumberOfTransactions: models.ElementHelper{
			Title:         "Number Of Transactions",
			Rules:         "",
			Type:          `Max15NumericText (based on string) pattern: [0-9]{1,15}`,
			Documentation: `Number of individual transactions contained in the message.`,
		},
		SettlementMethod: models.ElementHelper{
			Title:         "Settlement Method",
			Rules:         "",
			Type:          `SettlementMethodType(SettlementCLRG, SettlementINDA, SettlementCOVE, SettlementTDSO, SettlementTDSA)`,
			Documentation: `Method used to settle the (batch of) payment instructions.`,
		},
		CommonClearingSysCode: models.ElementHelper{
			Title:         "Common Clearing System Code",
			Rules:         "",
			Type:          `CommonClearingSysCodeType(ClearingSysFDW, ClearingSysCHIPS, ClearingSysSEPA ...)`,
			Documentation: `Infrastructure through which the payment instruction is processed, as published in an external clearing system identification code list.`,
		},
		InstructionId: models.ElementHelper{
			Title:         "Instruction Identification",
			Rules:         "Fedwire Funds Tag {3320} Sender Reference",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction. Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.`,
		},
		EndToEndId: models.ElementHelper{
			Title:         "End To End Identification",
			Rules:         "If no End To End Identification is available, then `NOTPROVIDED` should be used. ",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.`,
		},
		UniqueEndToEndTransactionRef: models.ElementHelper{
			Title:         "UETR",
			Rules:         "If the payment is a customer drawdown transfer, sent as a result of a customer drawdown request message (pain.013) that is being honored, then this should be the UETR of that customer drawdown request.",
			Type:          `UUIDv4Identifier (based on string) pattern: [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12} identificationScheme: RFC4122; UUIDv4`,
			Documentation: `Universally unique identifier to provide an end-to-end reference of a payment transaction.`,
		},
		SericeLevel: models.ElementHelper{
			Title:         "Service Level",
			Rules:         "The service level is chosen by the Fedwire Sender.",
			Type:          `ServiceLevel8Choice is_choice: true`,
			Documentation: `Agreement under which or rules under which the transaction should be processed.`,
		},
		InstrumentPropCode: models.ElementHelper{
			Title:         "Instrument Prop Code",
			Rules:         "",
			Type:          `InstrumentPropCodeType(InstrumentCTRC, InstrumentDD ...)`,
			Documentation: `Specifies the local instrument, as a proprietary code.`,
		},
		InterBankSettAmount: models.BuildCurrencyAndAmountHelper(),
		InterBankSettDate: models.ElementHelper{
			Title:         "Interbank Settlement Date",
			Rules:         "Must be the date of the current Fedwire funds-transfer business day in local date format (YYYY-MM-DD).",
			Type:          `ISODate (based on date)`,
			Documentation: `Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.`,
		},
		InstructedAmount: models.BuildCurrencyAndAmountHelper(),
		exchangeRate: models.ElementHelper{
			Title:         "Exchange Rate",
			Rules:         "",
			Type:          `BaseOneRate (based on decimal) totalDigits: 11 fractionDigits: 10 baseValue: 1`,
			Documentation: `Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.`,
		},
		ChargeBearer: models.ElementHelper{
			Title:         "Charge Bearer",
			Rules:         "",
			Type:          `ChargeBearerType(ChargeBearerSLEV, ChargeBearerRECV, ChargeBearerSHAR ...)`,
			Documentation: `Specifies which party/parties will bear the charges associated with the processing of the payment transaction.`,
		},
		ChargesInfo:       BuildChargeInfoHelper(),
		InstructingAgents: models.BuildAgentHelper(),
		InstructedAgent:   models.BuildAgentHelper(),
		IntermediaryAgent1Id: models.ElementHelper{
			Title:         "Intermediary Agent 1",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Agent between the debtor's agent and the creditor's agent. Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.`,
		},
		UltimateDebtorName: models.ElementHelper{
			Title:         "Ultimate Debtor Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		},
		UltimateDebtorAddress: models.BuildPostalAddressHelper(),
		DebtorName: models.ElementHelper{
			Title:         "Debtor Name",
			Rules:         "",
			Type:          `Max140Text (based on string) minLength: 1 maxLength: 140`,
			Documentation: `Name by which a party is known and which is usually used to identify that party.`,
		},
		DebtorAddress: models.BuildPostalAddressHelper(),
		DebtorIBAN: models.ElementHelper{
			Title:         "Debtor IBAN",
			Rules:         "",
			Type:          `IBAN2007Identifier (based on string) pattern: [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30} identificationScheme: National Banking Association; International Bank Account Number (ISO 13616)`,
			Documentation: `International Bank Account Number (IBAN) - identifier used internationally by financial institutions to uniquely identify the account of a customer. Further specifications of the format and content of the IBAN can be found in the standard ISO 13616 "Banking and related financial services - International Bank Account Number (IBAN)" version 1997-10-01, or later revisions.`,
		},
		DebtorOtherTypeId: models.ElementHelper{
			Title:         "Debtor Other Type Id",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Unique identification of an account, as assigned by the account servicer, using an identification scheme.`,
		},
		DebtorAgent:   models.BuildAgentHelper(),
		CreditorAgent: models.BuildAgentHelper(),
		CreditorName: models.ElementHelper{
			Title:         "Creditor Name",
			Rules:         "",
			Type:          `Max70Text (based on string) minLength: 1 maxLength: 70`,
			Documentation: `Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account. Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.`,
		},
		CreditorPostalAddress: models.BuildPostalAddressHelper(),
		UltimateCreditorName: models.ElementHelper{
			Title:         "Ultimate Creditor Name",
			Rules:         "",
			Type:          `Max70Text (based on string) minLength: 1 maxLength: 70`,
			Documentation: `Ultimate party to which an amount of money is due.`,
		},
		UltimateCreditorAddress: models.BuildPostalAddressHelper(),
		CreditorIBAN: models.ElementHelper{
			Title:         "Creditor IBAN",
			Rules:         "",
			Type:          `IBAN2007Identifier (based on string) pattern: [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30} identificationScheme: National Banking Association; International Bank Account Number (ISO 13616)`,
			Documentation: `International Bank Account Number (IBAN) - identifier used internationally by financial institutions to uniquely identify the account of a customer. Further specifications of the format and content of the IBAN can be found in the standard ISO 13616 "Banking and related financial services - International Bank Account Number (IBAN)" version 1997-10-01, or later revisions.`,
		},
		CreditorOtherTypeId: models.ElementHelper{
			Title:         "Creditor Other Type Id",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Unique identification of an account, as assigned by the account servicer, using an identification scheme.`,
		},
		RelatedRemittanceInfo: BuildRemittanceDetailHelper(),
		RemittanceInfor:       BuildRemittanceDocumentHelper(),
	}
}
