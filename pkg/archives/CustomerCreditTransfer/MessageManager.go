package CustomerCreditTransfer

import (
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type TaxRecord struct {
	//is used by governments to track tax obligations and payments.
	TaxId string
	//tax type code
	TaxTypeCode string
	// Tax Period Year
	TaxPeriodYear      fedwire.ISODate
	TaxperiodTimeFrame string
}
type RemittanceDetail struct {
	//unique reference number used to identify a remittance transaction.
	RemittanceId string
	//Specifies how the remittance information is delivered.
	Method Archive.RemittanceDeliveryMethod
	//Provides the email address where the remittance details should be sent.
	ElectronicAddress string
}
type RemittanceDocument struct {
	//refers to Unstructured Remittance Information in the ISO 20022 payment message standard
	UnstructuredRemitInfo string
	//Code or Proprietary :It is used to specify the method for identifying the type of a document or reference.
	CodeOrProprietary Archive.CodeOrProprietaryType
	//invoice number
	Number string
	//default value: current date
	RelatedDate fedwire.ISODate
	// Tax detail
	TaxDetail TaxRecord
}
type ChargeInfo struct {
	Amount         Archive.CurrencyAndAmount
	BusinessIdCode string
}