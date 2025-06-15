package CustomerCreditTransfer

import (
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/wadearnold/wire20022/pkg/models"
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
	Method models.RemittanceDeliveryMethod
	//Provides the email address where the remittance details should be sent.
	ElectronicAddress string
}
type RemittanceDocument struct {
	//refers to Unstructured Remittance Information in the ISO 20022 payment message standard
	UnstructuredRemitInfo string
	//Code or Proprietary :It is used to specify the method for identifying the type of a document or reference.
	CodeOrProprietary models.CodeOrProprietaryType
	//invoice number
	Number string
	//default value: current date
	RelatedDate fedwire.ISODate
	// Tax detail
	TaxDetail TaxRecord
}
type ChargeInfo struct {
	Amount         models.CurrencyAndAmount
	BusinessIdCode string
}
