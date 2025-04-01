package Pacs008Message

import (
	"reflect"

	"cloud.google.com/go/civil"
	CustomerCreditTransfer "github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer_pacs_008_001_08"
	fedwire "github.com/moov-io/fedwire20022/pkg/fedwire"
)

type PaymentSystemType string
type CodeOrProprietaryType string
type SettlementMethodType string
type CommonClearingSysCodeType string
type InstrumentPropCodeType string
type ChargeBearerType string
type PurposeOfPaymentType string
type RemittanceDeliveryMethod string

const (
	PaymentSysUSABA PaymentSystemType = "USABA" // American Bankers Association (ABA) routing number system
	PaymentSysCHIPS PaymentSystemType = "CHIPS" // Clearing House Interbank Payments System
	PaymentSysSEPA  PaymentSystemType = "SEPA"  // Single Euro Payments Area
	PaymentSysRTGS  PaymentSystemType = "RTGS"  // Real-Time Gross Settlement
	PaymentSysSWIFT PaymentSystemType = "SWIFT" // Society for Worldwide Interbank Financial Telecommunication
	PaymentSysBACS  PaymentSystemType = "BACS"  // Bankers' Automated Clearing Services
)
const (
	CodeCINV CodeOrProprietaryType = "CINV" // Invoice
	CodeCREQ CodeOrProprietaryType = "CREQ" // Credit Request
	CodeCNTR CodeOrProprietaryType = "CNTR" // Credit Note
	CodeDBTR CodeOrProprietaryType = "DBTR" // Debtor
	CodeCRED CodeOrProprietaryType = "CRED" // Credit
	CodeSCT  CodeOrProprietaryType = "SCT"  // SEPA Credit Transfer
	CodePAYM CodeOrProprietaryType = "PAYM" // Payment Message
	CodeRTGS CodeOrProprietaryType = "RTGS" // Real-Time Gross Settlement
	CodeRCLS CodeOrProprietaryType = "RCLS" // Reversal
	CodeRFF  CodeOrProprietaryType = "RFF"  // Reference
	CodeCMCN CodeOrProprietaryType = "CMCN" // Reference
)

const (
	SettlementCLRG SettlementMethodType = "CLRG" // Settlement via Clearing System (e.g., ACH, SEPA, RTGS)
	SettlementINDA SettlementMethodType = "INDA" // In-House Settlement (within the same bank)
	SettlementCOVE SettlementMethodType = "COVE" // Settlement through a Correspondent Bank
	SettlementTDSO SettlementMethodType = "TDSO" // Settlement via Target2 with a Settlement Agent
	SettlementTDSA SettlementMethodType = "TDSA" // Settlement via Target2 with a Direct Account
)

const (
	ClearingSysFDW   CommonClearingSysCodeType = "FDW"   // Fedwire (U.S.)
	ClearingSysCHIPS CommonClearingSysCodeType = "CHIPS" // CHIPS (U.S. Clearing House Interbank Payments System)
	ClearingSysSEPA  CommonClearingSysCodeType = "SEPA"  // SEPA (Single Euro Payments Area)
	ClearingSysRTGS  CommonClearingSysCodeType = "RTGS"  // Real-Time Gross Settlement
	ClearingSysSWIFT CommonClearingSysCodeType = "SWIFT" // SWIFT Network
	ClearingSysBACS  CommonClearingSysCodeType = "BACS"  // BACS (UK Clearing System)
	ClearingSysCNAPS CommonClearingSysCodeType = "CNAPS" // CNAPS (China’s Clearing System)
)

const (
	InstrumentCTRC                      InstrumentPropCodeType = "CTRC" // Credit Transfer (Proprietary Code)
	InstrumentDD                        InstrumentPropCodeType = "DD"   // Direct Debit
	InstrumentStraightThroughProcessing InstrumentPropCodeType = "STP"  // Straight Through Processing
	InstrumentNCT                       InstrumentPropCodeType = "NCT"  // National Credit Transfer
	InstrumentCTRD                      InstrumentPropCodeType = "CTRD" // National Credit Transfer
)

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

type CurrencyAndAmount struct {
	//default: USD
	Currency string
	Amount   float64
}
type Agent struct {
	//BICFI (Business Identifier Code - Financial Institution) is the ISO 9362 standard format used to identify banks and financial institutions globally.
	BusinessIdCode string
	//code that identifies a specific clearing system or a payment system within a financial network.
	//default value: USABA
	PaymentSysCode PaymentSystemType
	// stands for Member ID, which is a unique identifier for a financial institution or bank within the specified clearing system.
	PaymentSysMemberId string
	BankName           string
	PostalAddress      PostalAddress
}

type PostalAddress struct {
	StreetName     string
	BuildingNumber string
	BuildingName   string
	Floor          string
	RoomNumber     string
	PostalCode     string
	TownName       string
	Subdivision    string
	Country        string
}
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
	CodeOrProprietary CodeOrProprietaryType
	//invoice number
	Number string
	//default value: current date
	RelatedDate civil.Date
	// Tax detail
	TaxDetail TaxRecord
}
type ChargeInfo struct {
	amount         CurrencyAndAmount
	BusinessIdCode string
}

/*********************************************************/
/** Internal functions  **/
/*********************************************************/
func isEmptyDate(d civil.Date) bool {
	return d == civil.Date{}
}

func PostalAddress241From(param PostalAddress) CustomerCreditTransfer.PostalAddress241 {
	var Dbtr_PstlAdr CustomerCreditTransfer.PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := CustomerCreditTransfer.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := CustomerCreditTransfer.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := CustomerCreditTransfer.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := CustomerCreditTransfer.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := CustomerCreditTransfer.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := CustomerCreditTransfer.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		TwnNm := CustomerCreditTransfer.Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := CustomerCreditTransfer.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Ctry := CustomerCreditTransfer.CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return CustomerCreditTransfer.PostalAddress241{}
	}

	return Dbtr_PstlAdr
}
func isEmptyPostalAddress241(address CustomerCreditTransfer.PostalAddress241) bool {
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
func PostalAddress242From(param PostalAddress) CustomerCreditTransfer.PostalAddress242 {
	var Dbtr_PstlAdr CustomerCreditTransfer.PostalAddress242

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := CustomerCreditTransfer.Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := CustomerCreditTransfer.Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.TownName != "" {
		Dbtr_PstlAdr.TwnNm = CustomerCreditTransfer.Max35Text(param.TownName)
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := CustomerCreditTransfer.Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := CustomerCreditTransfer.Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := CustomerCreditTransfer.Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := CustomerCreditTransfer.Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := CustomerCreditTransfer.Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Dbtr_PstlAdr.Ctry = CustomerCreditTransfer.CountryCode(param.Country)
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return CustomerCreditTransfer.PostalAddress242{}
	}

	return Dbtr_PstlAdr
}
func isEmptyPostalAddress242(address CustomerCreditTransfer.PostalAddress242) bool {
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
func CashAccount38From(iban string, other string) CustomerCreditTransfer.CashAccount38 {
	if iban == "" && other == "" {
		return CustomerCreditTransfer.CashAccount38{} // Return empty struct if input is empty
	}
	var account CustomerCreditTransfer.AccountIdentification4Choice
	if iban != "" {
		_IBAN := CustomerCreditTransfer.IBAN2007Identifier(iban)
		account.IBAN = &_IBAN
	}
	if other != "" {
		account_Othr := CustomerCreditTransfer.GenericAccountIdentification1{}
		account.Othr = &account_Othr
		account.Othr.Id = CustomerCreditTransfer.Max34Text(other)
	}
	return CustomerCreditTransfer.CashAccount38{
		Id: account,
	}
}
func ClearingSystemMemberIdentification21From(param PaymentSystemType, paymentSysMemberId string) CustomerCreditTransfer.ClearingSystemMemberIdentification21 {
	var result CustomerCreditTransfer.ClearingSystemMemberIdentification21
	var hasData bool // Flag to check if there's valid data

	if param != "" {
		Cd := CustomerCreditTransfer.ExternalClearingSystemIdentification1Code(param)
		result.ClrSysId = CustomerCreditTransfer.ClearingSystemIdentification2Choice1{
			Cd: &Cd,
		}
		hasData = true
	}

	if paymentSysMemberId != "" {
		result.MmbId = CustomerCreditTransfer.Max35Text(paymentSysMemberId)
		hasData = true
	}

	// If no valid data, return an empty struct
	if !hasData {
		return CustomerCreditTransfer.ClearingSystemMemberIdentification21{}
	}

	return result
}
func RemittanceInformation161From(doc RemittanceDocument) CustomerCreditTransfer.RemittanceInformation161 {
	var result CustomerCreditTransfer.RemittanceInformation161
	var hasData bool // Flag to check if we have any meaningful data

	// Set UnstructuredRemitInfo if not empty
	if doc.UnstructuredRemitInfo != "" {
		UnstructuredRemitInfo := CustomerCreditTransfer.Max140Text(doc.UnstructuredRemitInfo)
		result.Ustrd = &UnstructuredRemitInfo
		hasData = true
	}

	// Prepare referred document information
	var RD_item CustomerCreditTransfer.ReferredDocumentInformation71
	var hasRDData bool // Check if RD_item contains meaningful data
	var hasTaxData bool
	var hasTaxPrData bool

	if doc.CodeOrProprietary != "" {
		RD_item_Tp_Cd := CustomerCreditTransfer.DocumentType6Code(doc.CodeOrProprietary)
		RD_item.Tp = &CustomerCreditTransfer.ReferredDocumentType4{
			CdOrPrtry: CustomerCreditTransfer.ReferredDocumentType3Choice{
				Cd: &RD_item_Tp_Cd,
			},
		}
		hasRDData = true
	}

	if doc.Number != "" {
		RD_item_Nb := CustomerCreditTransfer.Max35Text(doc.Number)
		RD_item.Nb = &RD_item_Nb
		hasRDData = true
	}

	if !isEmptyDate(doc.RelatedDate) {
		RD_item_RltdDt := fedwire.ISODate(doc.RelatedDate)
		RD_item.RltdDt = &RD_item_RltdDt
		hasRDData = true
	}

	var TaxRmt CustomerCreditTransfer.TaxInformation7
	if doc.TaxDetail.TaxId != "" {
		TaxId := CustomerCreditTransfer.Max35Text(doc.TaxDetail.TaxId)
		TaxRmt_Cdtr := CustomerCreditTransfer.TaxParty1{}
		TaxRmt.Cdtr = &TaxRmt_Cdtr
		TaxRmt.Cdtr.TaxId = &TaxId
		hasTaxData = true
	}
	var TaxRecode CustomerCreditTransfer.TaxRecord2
	if doc.TaxDetail.TaxTypeCode != "" {
		TaxRecode_Tp := CustomerCreditTransfer.Max35Text(doc.TaxDetail.TaxTypeCode)
		TaxRecode.Tp = &TaxRecode_Tp
		hasTaxPrData = true
	}
	if !isEmptyDate(doc.TaxDetail.TaxPeriodYear) {
		TaxRecode_Prd_Y := fedwire.ISODate(doc.TaxDetail.TaxPeriodYear)
		if TaxRecode.Prd == nil {
			TaxRecode_Prd := CustomerCreditTransfer.TaxPeriod2{}
			TaxRecode.Prd = &TaxRecode_Prd
		}
		TaxRecode.Prd.Yr = &TaxRecode_Prd_Y
		hasTaxPrData = true
	}
	if doc.TaxDetail.TaxperiodTimeFrame != "" {
		TaxRecode_Prd_tp := CustomerCreditTransfer.TaxRecordPeriod1Code(doc.TaxDetail.TaxperiodTimeFrame)
		if TaxRecode.Prd == nil {
			TaxRecode_Prd := CustomerCreditTransfer.TaxPeriod2{}
			TaxRecode.Prd = &TaxRecode_Prd
		}
		TaxRecode.Prd.Tp = &TaxRecode_Prd_tp
		hasTaxPrData = true
	}
	if hasTaxPrData {
		TaxRmt.Rcrd = []*CustomerCreditTransfer.TaxRecord2{
			&TaxRecode,
		}
		hasTaxPrData = true
	}

	SR_item := CustomerCreditTransfer.StructuredRemittanceInformation161{}
	if hasRDData {
		SR_item.RfrdDocInf = []*CustomerCreditTransfer.ReferredDocumentInformation71{
			&RD_item,
		}
	}
	if hasTaxData {
		SR_item.TaxRmt = &TaxRmt
	}

	// If RD_item has data, add it to structured remittance info
	if hasRDData || hasTaxData {
		result.Strd = []*CustomerCreditTransfer.StructuredRemittanceInformation161{
			&SR_item,
		}
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return CustomerCreditTransfer.RemittanceInformation161{}
	}

	return result
}

// func (r CustomerCreditTransfer.RemittanceInformation161) isEmpty() bool {
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
func FinancialInstitutionIdentification181From(agent Agent) CustomerCreditTransfer.FinancialInstitutionIdentification181 {
	var result CustomerCreditTransfer.FinancialInstitutionIdentification181
	if agent.BusinessIdCode != "" {
		_BICFI := CustomerCreditTransfer.BICFIDec2014Identifier(agent.BusinessIdCode)
		result.BICFI = &_BICFI
	}
	if agent.PaymentSysCode != "" || agent.PaymentSysMemberId != "" {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := CustomerCreditTransfer.ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		if agent.PaymentSysCode != "" {
			Cd := CustomerCreditTransfer.ExternalClearingSystemIdentification1Code(agent.PaymentSysCode)
			result.ClrSysMmbId.ClrSysId = CustomerCreditTransfer.ClearingSystemIdentification2Choice1{
				Cd: &Cd,
			}
		}
		if agent.PaymentSysMemberId != "" {
			result.ClrSysMmbId.MmbId = CustomerCreditTransfer.Max35Text(agent.PaymentSysMemberId)
		}
	}
	if agent.BankName != "" {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := CustomerCreditTransfer.ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		_BKNM := CustomerCreditTransfer.Max140Text(agent.BankName)
		result.Nm = &_BKNM
	}
	postalAddress := PostalAddress241From(agent.PostalAddress)
	if !isEmptyPostalAddress241(postalAddress) {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := CustomerCreditTransfer.ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		result.PstlAdr = &postalAddress
	}
	return result
}

func PaymentTypeInformation281From(InstrumentPropCode InstrumentPropCodeType, SericeLevel string) CustomerCreditTransfer.PaymentTypeInformation281 {
	var result CustomerCreditTransfer.PaymentTypeInformation281
	if InstrumentPropCode != "" {
		result.LclInstrm = CustomerCreditTransfer.LocalInstrument2Choice1{}
		CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry := CustomerCreditTransfer.LocalInstrumentFedwireFunds1(InstrumentPropCode)
		result.LclInstrm.Prtry = &CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry
	}
	if SericeLevel != "" {
		svclv := CustomerCreditTransfer.ExternalServiceLevel1Code(SericeLevel)
		CdtTrfTxInf_PmtTpInf_SvcLvl := CustomerCreditTransfer.ServiceLevel8Choice{
			Cd: &svclv,
		}
		result.SvcLvl = []*CustomerCreditTransfer.ServiceLevel8Choice{
			&CdtTrfTxInf_PmtTpInf_SvcLvl,
		}
	}
	return result
}
func RemittanceLocation71From(param RemittanceDetail) CustomerCreditTransfer.RemittanceLocation71 {
	var result CustomerCreditTransfer.RemittanceLocation71
	if param.RemittanceId != "" {
		_RemittanceId := CustomerCreditTransfer.Max35Text(param.RemittanceId)
		result.RmtId = &_RemittanceId
	}
	var locationData CustomerCreditTransfer.RemittanceLocationData11
	var hasLocationData = false
	if param.Method != "" {
		locationData.Mtd = CustomerCreditTransfer.RemittanceLocationMethod2Code(param.Method)
		hasLocationData = true
	}
	if param.ElectronicAddress != "" {
		_ElectronicAddress := CustomerCreditTransfer.Max2048Text(param.ElectronicAddress)
		locationData.ElctrncAdr = &_ElectronicAddress
		hasLocationData = true
	}
	if hasLocationData {
		result.RmtLctnDtls = []*CustomerCreditTransfer.RemittanceLocationData11{
			&locationData,
		}
	}
	return result
}

// func (r CustomerCreditTransfer.RemittanceLocation71) isEmpty() bool {
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
func PartyIdentification1352From(Nm string, PstlAdr PostalAddress) CustomerCreditTransfer.PartyIdentification1352 {
	var result CustomerCreditTransfer.PartyIdentification1352
	if Nm != "" {
		_nm := CustomerCreditTransfer.Max140Text(Nm)
		result.Nm = &_nm
	}
	_PstlAdr := PostalAddress241From(PstlAdr)
	if !isEmptyPostalAddress241(_PstlAdr) {
		result.PstlAdr = &_PstlAdr
	}
	return result
}

//	func (p CustomerCreditTransfer.PartyIdentification1352) isEmpty() bool {
//		if p.Nm == nil {
//			return true
//		}
//		if isEmptyPostalAddress241(*p.PstlAdr) {
//			return true
//		}
//		return false
//	}
func PartyIdentification1351From(Nm string, PstlAdr PostalAddress) CustomerCreditTransfer.PartyIdentification1351 {
	var result CustomerCreditTransfer.PartyIdentification1351
	if Nm != "" {
		_nm := CustomerCreditTransfer.Max140Text(Nm)
		result.Nm = &_nm
	}
	_PstlAdr := PostalAddress242From(PstlAdr)
	if !isEmptyPostalAddress242(_PstlAdr) {
		result.PstlAdr = &_PstlAdr
	}
	return result
}

//	func (p CustomerCreditTransfer.PartyIdentification1351) isEmpty() bool {
//		if p.Nm == nil {
//			return true
//		}
//		if isEmptyPostalAddress242(*p.PstlAdr) {
//			return true
//		}
//		return false
//	}
func Charges71From(data ChargeInfo) CustomerCreditTransfer.Charges71 {
	var result CustomerCreditTransfer.Charges71
	if data.amount.Amount != 0 || data.amount.Currency != "" {
		result.Amt = CustomerCreditTransfer.ActiveOrHistoricCurrencyAndAmount{
			Value: CustomerCreditTransfer.ActiveOrHistoricCurrencyAndAmountSimpleType(data.amount.Amount),
			Ccy:   CustomerCreditTransfer.ActiveOrHistoricCurrencyCode(data.amount.Currency),
		}
	}
	if data.BusinessIdCode != "" {
		result.Agt = CustomerCreditTransfer.BranchAndFinancialInstitutionIdentification61{}
		result.Agt.FinInstnId = CustomerCreditTransfer.FinancialInstitutionIdentification181{}
		_BICFI := CustomerCreditTransfer.BICFIDec2014Identifier(data.BusinessIdCode)
		result.Agt.FinInstnId.BICFI = &_BICFI
	}
	return result
}

//	func (c CustomerCreditTransfer.Charges71) isEmpty() bool {
//		if c.Amt.Value != 0 {
//			return false
//		}
//		if c.Agt.FinInstnId.BICFI != nil {
//			return false
//		}
//		return true
//	}
func BranchAndFinancialInstitutionIdentification61From(BICFI string) CustomerCreditTransfer.BranchAndFinancialInstitutionIdentification61 {
	var result CustomerCreditTransfer.BranchAndFinancialInstitutionIdentification61
	if BICFI != "" {
		result.FinInstnId = CustomerCreditTransfer.FinancialInstitutionIdentification181{}
		_BICFI := CustomerCreditTransfer.BICFIDec2014Identifier(BICFI)
		result.FinInstnId.BICFI = &_BICFI
	}
	return result
}

//	func (b CustomerCreditTransfer.BranchAndFinancialInstitutionIdentification61) isEmpty() bool {
//		return b.FinInstnId.BICFI == nil
//	}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
