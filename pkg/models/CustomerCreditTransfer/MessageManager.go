package CustomerCreditTransfer

import (
	"reflect"

	"cloud.google.com/go/civil"
	pacs008 "github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer_pacs_008_001_08"
	fedwire "github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type ChargeBearerType string
type PurposeOfPaymentType string
type RemittanceDeliveryMethod string

const (
	ChargeBearerSLEV ChargeBearerType = "SLEV" // Sender Pays All Charges
	ChargeBearerRECV ChargeBearerType = "RECV" // Receiver Pays All Charges
	ChargeBearerSHAR ChargeBearerType = "SHAR" // Shared Charges
)

const (
	CashWithdrawal    PurposeOfPaymentType = "CASH" // Cash Withdrawal
	GoodsAndServices  PurposeOfPaymentType = "GDSV" // Goods and Services
	LabourInsurance   PurposeOfPaymentType = "CASH" // Labour Insurance
	SupplierPayment   PurposeOfPaymentType = "SUPP" // Supplier Payment
	TradeSettlement   PurposeOfPaymentType = "TRAD" // Trade Settlement
	InvestmentPayment PurposeOfPaymentType = "IVPT" // Investment Payment
	PensionPayment    PurposeOfPaymentType = "PENS" // Pension Payment
	AlimonyPayment    PurposeOfPaymentType = "ALMY" // Alimony Payment
)

const (
	Fax                       RemittanceDeliveryMethod = "FAXI" //Fax
	ElectronicDataInterchange RemittanceDeliveryMethod = "EDIC" //Electronic Data Interchange (EDI)
	UniformResourceIdentifier RemittanceDeliveryMethod = "URID" //Uniform Resource Identifier (URI)
	PostalMail                RemittanceDeliveryMethod = "POST" //Postal Mail
	Email                     RemittanceDeliveryMethod = "EMAL" //Email
)

type TaxRecord struct {
	//is used by governments to track tax obligations and payments.
	TaxId string
	//tax type code
	TaxTypeCode string
	// Tax Period Year
	TaxPeriodYear      civil.Date
	TaxperiodTimeFrame string
}
type RemittanceDetail struct {
	//unique reference number used to identify a remittance transaction.
	RemittanceId string
	//Specifies how the remittance information is delivered.
	Method RemittanceDeliveryMethod
	//Provides the email address where the remittance details should be sent.
	ElectronicAddress string
}
type RemittanceDocument struct {
	//refers to Unstructured Remittance Information in the ISO 20022 payment message standard
	UnstructuredRemitInfo string
	//Code or Proprietary :It is used to specify the method for identifying the type of a document or reference.
	CodeOrProprietary model.CodeOrProprietaryType
	//invoice number
	Number string
	//default value: current date
	RelatedDate civil.Date
	// Tax detail
	TaxDetail TaxRecord
}
type ChargeInfo struct {
	amount         model.CurrencyAndAmount
	BusinessIdCode string
}

/*********************************************************/
/** Internal functions  **/
/*********************************************************/
func isEmptyDate(d civil.Date) bool {
	return d == civil.Date{}
}

func PostalAddress241From(param model.PostalAddress) pacs008.PostalAddress241 {
	var Dbtr_PstlAdr pacs008.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := pacs008.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := pacs008.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := pacs008.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := pacs008.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := pacs008.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := pacs008.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		TwnNm := pacs008.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := pacs008.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Ctry := pacs008.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs008.PostalAddress241{}
	}

	return Dbtr_PstlAdr
}
func isEmptyPostalAddress241(address pacs008.PostalAddress241) bool {
	// Compare the struct with its zero value
	return address.StrtNm == nil &&
		address.BldgNb == nil &&
		address.BldgNm == nil &&
		address.Flr == nil &&
		address.Room == nil &&
		address.PstCd == nil &&
		address.TwnNm == nil &&
		address.CtrySubDvsn == nil &&
		address.Ctry == nil
}
func PostalAddress242From(param model.PostalAddress) pacs008.PostalAddress242 {
	var Dbtr_PstlAdr pacs008.PostalAddress242

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := pacs008.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := pacs008.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.TownName != "" {
		Dbtr_PstlAdr.TwnNm = pacs008.Max35Text(param.TownName)
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := pacs008.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := pacs008.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := pacs008.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := pacs008.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := pacs008.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Dbtr_PstlAdr.Ctry = pacs008.CountryCode(param.Country)
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs008.PostalAddress242{}
	}

	return Dbtr_PstlAdr
}
func isEmptyPostalAddress242(address pacs008.PostalAddress242) bool {
	// Compare the struct with its zero value
	return address.StrtNm == nil &&
		address.BldgNb == nil &&
		address.BldgNm == nil &&
		address.Flr == nil &&
		address.TwnNm == "" &&
		address.Room == nil &&
		address.PstCd == nil &&
		address.Ctry == ""
}
func CashAccount38From(iban string, other string) pacs008.CashAccount38 {
	if iban == "" && other == "" {
		return pacs008.CashAccount38{} // Return empty struct if input is empty
	}
	var account pacs008.AccountIdentification4Choice
	if iban != "" {
		_IBAN := pacs008.IBAN2007Identifier(iban)
		account.IBAN = &_IBAN
	}
	if other != "" {
		account_Othr := pacs008.GenericAccountIdentification1{}
		account.Othr = &account_Othr
		account.Othr.Id = pacs008.Max34Text(other)
	}
	return pacs008.CashAccount38{
		Id: account,
	}
}
func ClearingSystemMemberIdentification21From(param model.PaymentSystemType, paymentSysMemberId string) pacs008.ClearingSystemMemberIdentification21 {
	var result pacs008.ClearingSystemMemberIdentification21
	var hasData bool // Flag to check if there's valid data

	if param != "" {
		Cd := pacs008.ExternalClearingSystemIdentification1Code(param)
		result.ClrSysId = pacs008.ClearingSystemIdentification2Choice1{
			Cd: &Cd,
		}
		hasData = true
	}

	if paymentSysMemberId != "" {
		result.MmbId = pacs008.Max35Text(paymentSysMemberId)
		hasData = true
	}

	// If no valid data, return an empty struct
	if !hasData {
		return pacs008.ClearingSystemMemberIdentification21{}
	}

	return result
}
func RemittanceInformation161From(doc RemittanceDocument) pacs008.RemittanceInformation161 {
	var result pacs008.RemittanceInformation161
	var hasData bool // Flag to check if we have any meaningful data

	// Set UnstructuredRemitInfo if not empty
	if doc.UnstructuredRemitInfo != "" {
		UnstructuredRemitInfo := pacs008.Max140Text(doc.UnstructuredRemitInfo)
		result.Ustrd = &UnstructuredRemitInfo
		hasData = true
	}

	// Prepare referred document information
	var RD_item pacs008.ReferredDocumentInformation71
	var hasRDData bool // Check if RD_item contains meaningful data
	var hasTaxData bool
	var hasTaxPrData bool

	if doc.CodeOrProprietary != "" {
		RD_item_Tp_Cd := pacs008.DocumentType6Code(doc.CodeOrProprietary)
		RD_item.Tp = &pacs008.ReferredDocumentType4{
			CdOrPrtry: pacs008.ReferredDocumentType3Choice{
				Cd: &RD_item_Tp_Cd,
			},
		}
		hasRDData = true
	}

	if doc.Number != "" {
		RD_item_Nb := pacs008.Max35Text(doc.Number)
		RD_item.Nb = &RD_item_Nb
		hasRDData = true
	}

	if !isEmptyDate(doc.RelatedDate) {
		RD_item_RltdDt := fedwire.ISODate(doc.RelatedDate)
		RD_item.RltdDt = &RD_item_RltdDt
		hasRDData = true
	}

	var TaxRmt pacs008.TaxInformation7
	if doc.TaxDetail.TaxId != "" {
		TaxId := pacs008.Max35Text(doc.TaxDetail.TaxId)
		TaxRmt_Cdtr := pacs008.TaxParty1{}
		TaxRmt.Cdtr = &TaxRmt_Cdtr
		TaxRmt.Cdtr.TaxId = &TaxId
		hasTaxData = true
	}
	var TaxRecode pacs008.TaxRecord2
	if doc.TaxDetail.TaxTypeCode != "" {
		TaxRecode_Tp := pacs008.Max35Text(doc.TaxDetail.TaxTypeCode)
		TaxRecode.Tp = &TaxRecode_Tp
		hasTaxPrData = true
	}
	if !isEmptyDate(doc.TaxDetail.TaxPeriodYear) {
		TaxRecode_Prd_Y := fedwire.ISODate(doc.TaxDetail.TaxPeriodYear)
		if TaxRecode.Prd == nil {
			TaxRecode_Prd := pacs008.TaxPeriod2{}
			TaxRecode.Prd = &TaxRecode_Prd
		}
		TaxRecode.Prd.Yr = &TaxRecode_Prd_Y
		hasTaxPrData = true
	}
	if doc.TaxDetail.TaxperiodTimeFrame != "" {
		TaxRecode_Prd_tp := pacs008.TaxRecordPeriod1Code(doc.TaxDetail.TaxperiodTimeFrame)
		if TaxRecode.Prd == nil {
			TaxRecode_Prd := pacs008.TaxPeriod2{}
			TaxRecode.Prd = &TaxRecode_Prd
		}
		TaxRecode.Prd.Tp = &TaxRecode_Prd_tp
		hasTaxPrData = true
	}
	if hasTaxPrData {
		TaxRmt.Rcrd = []*pacs008.TaxRecord2{
			&TaxRecode,
		}
		hasTaxPrData = true
	}

	SR_item := pacs008.StructuredRemittanceInformation161{}
	if hasRDData {
		SR_item.RfrdDocInf = []*pacs008.ReferredDocumentInformation71{
			&RD_item,
		}
	}
	if hasTaxData {
		SR_item.TaxRmt = &TaxRmt
	}

	// If RD_item has data, add it to structured remittance info
	if hasRDData || hasTaxData {
		result.Strd = []*pacs008.StructuredRemittanceInformation161{
			&SR_item,
		}
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return pacs008.RemittanceInformation161{}
	}

	return result
}

// func (r pacs008.RemittanceInformation161) isEmpty() bool {
// 	// Check if Ustrd (Unstructured Remittance Info) is set
// 	if r.Ustrd != nil {
// 		return false
// 	}

// 	// Check if Strd (Structured Remittance Info) contains valid data
// 	if len(r.Strd) > 0 {
// 		for _, srItem := range r.Strd {
// 			if srItem != nil {
// 				// Check if there are valid referred document information items
// 				if len(srItem.RfrdDocInf) > 0 {
// 					for _, refDoc := range srItem.RfrdDocInf {
// 						if refDoc != nil && (refDoc.Tp != nil || refDoc.Nb != nil || refDoc.RltdDt != nil) {
// 							return false
// 						}
// 					}
// 				}

// 				// Check if TaxRmt (Tax Remittance) contains valid data
// 				if srItem.TaxRmt != nil {
// 					if srItem.TaxRmt.Cdtr != nil && srItem.TaxRmt.Cdtr.TaxId != nil {
// 						return false
// 					}
// 					if len(srItem.TaxRmt.Rcrd) > 0 {
// 						for _, taxRecord := range srItem.TaxRmt.Rcrd {
// 							if taxRecord != nil && (taxRecord.Tp != nil || taxRecord.Prd != nil) {
// 								return false
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

//		// If none of the fields contain data, return true (struct is empty)
//		return true
//	}
func FinancialInstitutionIdentification181From(agent model.Agent) pacs008.FinancialInstitutionIdentification181 {
	var result pacs008.FinancialInstitutionIdentification181
	if agent.BusinessIdCode != "" {
		_BICFI := pacs008.BICFIDec2014Identifier(agent.BusinessIdCode)
		result.BICFI = &_BICFI
	}
	if agent.PaymentSysCode != "" || agent.PaymentSysMemberId != "" {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := pacs008.ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		if agent.PaymentSysCode != "" {
			Cd := pacs008.ExternalClearingSystemIdentification1Code(agent.PaymentSysCode)
			result.ClrSysMmbId.ClrSysId = pacs008.ClearingSystemIdentification2Choice1{
				Cd: &Cd,
			}
		}
		if agent.PaymentSysMemberId != "" {
			result.ClrSysMmbId.MmbId = pacs008.Max35Text(agent.PaymentSysMemberId)
		}
	}
	if agent.BankName != "" {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := pacs008.ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		_BKNM := pacs008.Max140Text(agent.BankName)
		result.Nm = &_BKNM
	}
	postalAddress := PostalAddress241From(agent.PostalAddress)
	if !isEmptyPostalAddress241(postalAddress) {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := pacs008.ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		result.PstlAdr = &postalAddress
	}
	return result
}

func PaymentTypeInformation281From(InstrumentPropCode model.InstrumentPropCodeType, SericeLevel string) pacs008.PaymentTypeInformation281 {
	var result pacs008.PaymentTypeInformation281
	if InstrumentPropCode != "" {
		result.LclInstrm = pacs008.LocalInstrument2Choice1{}
		CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry := pacs008.LocalInstrumentFedwireFunds1(InstrumentPropCode)
		result.LclInstrm.Prtry = &CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry
	}
	if SericeLevel != "" {
		svclv := pacs008.ExternalServiceLevel1Code(SericeLevel)
		CdtTrfTxInf_PmtTpInf_SvcLvl := pacs008.ServiceLevel8Choice{
			Cd: &svclv,
		}
		result.SvcLvl = []*pacs008.ServiceLevel8Choice{
			&CdtTrfTxInf_PmtTpInf_SvcLvl,
		}
	}
	return result
}
func RemittanceLocation71From(param RemittanceDetail) pacs008.RemittanceLocation71 {
	var result pacs008.RemittanceLocation71
	if param.RemittanceId != "" {
		_RemittanceId := pacs008.Max35Text(param.RemittanceId)
		result.RmtId = &_RemittanceId
	}
	var locationData pacs008.RemittanceLocationData11
	var hasLocationData = false
	if param.Method != "" {
		locationData.Mtd = pacs008.RemittanceLocationMethod2Code(param.Method)
		hasLocationData = true
	}
	if param.ElectronicAddress != "" {
		_ElectronicAddress := pacs008.Max2048Text(param.ElectronicAddress)
		locationData.ElctrncAdr = &_ElectronicAddress
		hasLocationData = true
	}
	if hasLocationData {
		result.RmtLctnDtls = []*pacs008.RemittanceLocationData11{
			&locationData,
		}
	}
	return result
}

// func (r pacs008.RemittanceLocation71) isEmpty() bool {
// 	// Check if RmtId is nil (i.e., no remittance ID is set)
// 	if r.RmtId != nil {
// 		return false
// 	}

// 	// Check if RmtLctnDtls contains any valid location data
// 	if len(r.RmtLctnDtls) > 0 {
// 		for _, loc := range r.RmtLctnDtls {
// 			if loc != nil && (loc.Mtd != "" || loc.ElctrncAdr != nil) {
// 				return false
// 			}
// 		}
// 	}

//		// If none of the above fields have meaningful data, it's empty
//		return true
//	}
func PartyIdentification1352From(Nm string, PstlAdr model.PostalAddress) pacs008.PartyIdentification1352 {
	var result pacs008.PartyIdentification1352
	if Nm != "" {
		_nm := pacs008.Max140Text(Nm)
		result.Nm = &_nm
	}
	_PstlAdr := PostalAddress241From(PstlAdr)
	if !isEmptyPostalAddress241(_PstlAdr) {
		result.PstlAdr = &_PstlAdr
	}
	return result
}

//	func (p pacs008.PartyIdentification1352) isEmpty() bool {
//		if p.Nm == nil {
//			return true
//		}
//		if isEmptyPostalAddress241(*p.PstlAdr) {
//			return true
//		}
//		return false
//	}
func PartyIdentification1351From(Nm string, PstlAdr model.PostalAddress) pacs008.PartyIdentification1351 {
	var result pacs008.PartyIdentification1351
	if Nm != "" {
		_nm := pacs008.Max140Text(Nm)
		result.Nm = &_nm
	}
	_PstlAdr := PostalAddress242From(PstlAdr)
	if !isEmptyPostalAddress242(_PstlAdr) {
		result.PstlAdr = &_PstlAdr
	}
	return result
}

//	func (p pacs008.PartyIdentification1351) isEmpty() bool {
//		if p.Nm == nil {
//			return true
//		}
//		if isEmptyPostalAddress242(*p.PstlAdr) {
//			return true
//		}
//		return false
//	}
func Charges71From(data ChargeInfo) pacs008.Charges71 {
	var result pacs008.Charges71
	if data.amount.Amount != 0 || data.amount.Currency != "" {
		result.Amt = pacs008.ActiveOrHistoricCurrencyAndAmount{
			Value: pacs008.ActiveOrHistoricCurrencyAndAmountSimpleType(data.amount.Amount),
			Ccy:   pacs008.ActiveOrHistoricCurrencyCode(data.amount.Currency),
		}
	}
	if data.BusinessIdCode != "" {
		result.Agt = pacs008.BranchAndFinancialInstitutionIdentification61{}
		result.Agt.FinInstnId = pacs008.FinancialInstitutionIdentification181{}
		_BICFI := pacs008.BICFIDec2014Identifier(data.BusinessIdCode)
		result.Agt.FinInstnId.BICFI = &_BICFI
	}
	return result
}

//	func (c pacs008.Charges71) isEmpty() bool {
//		if c.Amt.Value != 0 {
//			return false
//		}
//		if c.Agt.FinInstnId.BICFI != nil {
//			return false
//		}
//		return true
//	}
func BranchAndFinancialInstitutionIdentification61From(BICFI string) pacs008.BranchAndFinancialInstitutionIdentification61 {
	var result pacs008.BranchAndFinancialInstitutionIdentification61
	if BICFI != "" {
		result.FinInstnId = pacs008.FinancialInstitutionIdentification181{}
		_BICFI := pacs008.BICFIDec2014Identifier(BICFI)
		result.FinInstnId.BICFI = &_BICFI
	}
	return result
}

//	func (b pacs008.BranchAndFinancialInstitutionIdentification61) isEmpty() bool {
//		return b.FinInstnId.BICFI == nil
//	}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
