package DrawdownRequest

import (
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type RemittanceDocument struct {
	//Code or Proprietary :It is used to specify the method for identifying the type of a document or reference.
	CodeOrProprietary Archive.CodeOrProprietaryType
	//invoice number
	Number string
	//default value: current date
	RelatedDate fedwire.ISODate
}

type CreditTransferTransaction struct {
	//Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	PaymentInstructionId string
	//Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	PaymentEndToEndId string
	//Universally unique identifier to provide an end-to-end reference of a payment transaction.
	PaymentUniqueId string
	//Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	PayRequestType Archive.PaymentRequestType
	//Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	PayCategoryType Archive.PaymentRequestType
	//Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount Archive.CurrencyAndAmount
	//Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer Archive.ChargeBearerType
	// /Financial institution servicing an account for the creditor.
	CreditorAgent Archive.Agent
	//This is the party whose account will be credited by the creditor agent if the drawdown request is honored.
	Creditor Archive.PartyIdentify
	//Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CrediorAccountOtherId string
	//Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation string
	Document              RemittanceDocument
}
