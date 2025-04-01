package CustomerCreditTransfer_pacs_008_001_08

import (
	"cloud.google.com/go/civil"
	fedwire "github.com/moov-io/wire20022/pkg/internal"
)

type Pacs008MessageManager interface {
	// Create Business Model from parameters
	CreateDocument()
	GetXML() ([]byte, error)
	GetJson() ([]byte, error)
}

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

func PostalAddress241From(param PostalAddress) PostalAddress241 {
	var Dbtr_PstlAdr PostalAddress241

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.TownName != "" {
		TwnNm := Max35Text(param.TownName)
		Dbtr_PstlAdr.TwnNm = &TwnNm
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Ctry := CountryCode(param.Country)
		Dbtr_PstlAdr.Ctry = &Ctry
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return PostalAddress241{}
	}

	return Dbtr_PstlAdr
}
func isEmptyPostalAddress241(address PostalAddress241) bool {
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
func PostalAddress242From(param PostalAddress) PostalAddress242 {
	var Dbtr_PstlAdr PostalAddress242

	// Flag to track if any field is set
	hasData := false

	// Check and set each field if not empty
	if param.StreetName != "" {
		StrtNm := Max70Text(param.StreetName)
		Dbtr_PstlAdr.StrtNm = &StrtNm
		hasData = true
	}
	if param.BuildingNumber != "" {
		BldgNb := Max16Text(param.BuildingNumber)
		Dbtr_PstlAdr.BldgNb = &BldgNb
		hasData = true
	}
	if param.TownName != "" {
		Dbtr_PstlAdr.TwnNm = Max35Text(param.TownName)
		hasData = true
	}
	if param.BuildingName != "" {
		BldgNm := Max35Text(param.BuildingName)
		Dbtr_PstlAdr.BldgNm = &BldgNm
		hasData = true
	}
	if param.Floor != "" {
		Floor := Max70Text(param.Floor)
		Dbtr_PstlAdr.Flr = &Floor
		hasData = true
	}
	if param.RoomNumber != "" {
		Room := Max70Text(param.RoomNumber)
		Dbtr_PstlAdr.Room = &Room
		hasData = true
	}
	if param.PostalCode != "" {
		PstCd := Max16Text(param.PostalCode)
		Dbtr_PstlAdr.PstCd = &PstCd
		hasData = true
	}
	if param.Subdivision != "" {
		CtrySubDvsn := Max35Text(param.Subdivision)
		Dbtr_PstlAdr.CtrySubDvsn = &CtrySubDvsn
		hasData = true
	}
	if param.Country != "" {
		Dbtr_PstlAdr.Ctry = CountryCode(param.Country)
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return PostalAddress242{}
	}

	return Dbtr_PstlAdr
}
func isEmptyPostalAddress242(address PostalAddress242) bool {
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
func CashAccount38From(iban string, other string) CashAccount38 {
	if iban == "" && other == "" {
		return CashAccount38{} // Return empty struct if input is empty
	}
	var account AccountIdentification4Choice
	if iban != "" {
		_IBAN := IBAN2007Identifier(iban)
		account.IBAN = &_IBAN
	}
	if other != "" {
		account_Othr := GenericAccountIdentification1{}
		account.Othr = &account_Othr
		account.Othr.Id = Max34Text(other)
	}
	return CashAccount38{
		Id: account,
	}
}
func ClearingSystemMemberIdentification21From(param PaymentSystemType, paymentSysMemberId string) ClearingSystemMemberIdentification21 {
	var result ClearingSystemMemberIdentification21
	var hasData bool // Flag to check if there's valid data

	if param != "" {
		Cd := ExternalClearingSystemIdentification1Code(param)
		result.ClrSysId = ClearingSystemIdentification2Choice1{
			Cd: &Cd,
		}
		hasData = true
	}

	if paymentSysMemberId != "" {
		result.MmbId = Max35Text(paymentSysMemberId)
		hasData = true
	}

	// If no valid data, return an empty struct
	if !hasData {
		return ClearingSystemMemberIdentification21{}
	}

	return result
}
func RemittanceInformation161From(doc RemittanceDocument) RemittanceInformation161 {
	var result RemittanceInformation161
	var hasData bool // Flag to check if we have any meaningful data

	// Set UnstructuredRemitInfo if not empty
	if doc.UnstructuredRemitInfo != "" {
		UnstructuredRemitInfo := Max140Text(doc.UnstructuredRemitInfo)
		result.Ustrd = &UnstructuredRemitInfo
		hasData = true
	}

	// Prepare referred document information
	var RD_item ReferredDocumentInformation71
	var hasRDData bool // Check if RD_item contains meaningful data
	var hasTaxData bool
	var hasTaxPrData bool

	if doc.CodeOrProprietary != "" {
		RD_item_Tp_Cd := DocumentType6Code(doc.CodeOrProprietary)
		RD_item.Tp = &ReferredDocumentType4{
			CdOrPrtry: ReferredDocumentType3Choice{
				Cd: &RD_item_Tp_Cd,
			},
		}
		hasRDData = true
	}

	if doc.Number != "" {
		RD_item_Nb := Max35Text(doc.Number)
		RD_item.Nb = &RD_item_Nb
		hasRDData = true
	}

	if !isEmptyDate(doc.RelatedDate) {
		RD_item_RltdDt := fedwire.ISODate(doc.RelatedDate)
		RD_item.RltdDt = &RD_item_RltdDt
		hasRDData = true
	}

	var TaxRmt TaxInformation7
	if doc.TaxDetail.TaxId != "" {
		TaxId := Max35Text(doc.TaxDetail.TaxId)
		TaxRmt_Cdtr := TaxParty1{}
		TaxRmt.Cdtr = &TaxRmt_Cdtr
		TaxRmt.Cdtr.TaxId = &TaxId
		hasTaxData = true
	}
	var TaxRecode TaxRecord2
	if doc.TaxDetail.TaxTypeCode != "" {
		TaxRecode_Tp := Max35Text(doc.TaxDetail.TaxTypeCode)
		TaxRecode.Tp = &TaxRecode_Tp
		hasTaxPrData = true
	}
	if !isEmptyDate(doc.TaxDetail.TaxPeriodYear) {
		TaxRecode_Prd_Y := fedwire.ISODate(doc.TaxDetail.TaxPeriodYear)
		if TaxRecode.Prd == nil {
			TaxRecode_Prd := TaxPeriod2{}
			TaxRecode.Prd = &TaxRecode_Prd
		}
		TaxRecode.Prd.Yr = &TaxRecode_Prd_Y
		hasTaxPrData = true
	}
	if doc.TaxDetail.TaxperiodTimeFrame != "" {
		TaxRecode_Prd_tp := TaxRecordPeriod1Code(doc.TaxDetail.TaxperiodTimeFrame)
		if TaxRecode.Prd == nil {
			TaxRecode_Prd := TaxPeriod2{}
			TaxRecode.Prd = &TaxRecode_Prd
		}
		TaxRecode.Prd.Tp = &TaxRecode_Prd_tp
		hasTaxPrData = true
	}
	if hasTaxPrData {
		TaxRmt.Rcrd = []*TaxRecord2{
			&TaxRecode,
		}
		hasTaxPrData = true
	}

	SR_item := StructuredRemittanceInformation161{}
	if hasRDData {
		SR_item.RfrdDocInf = []*ReferredDocumentInformation71{
			&RD_item,
		}
	}
	if hasTaxData {
		SR_item.TaxRmt = &TaxRmt
	}

	// If RD_item has data, add it to structured remittance info
	if hasRDData || hasTaxData {
		result.Strd = []*StructuredRemittanceInformation161{
			&SR_item,
		}
		hasData = true
	}

	// If no data was set, return an empty struct
	if !hasData {
		return RemittanceInformation161{}
	}

	return result
}
func (r RemittanceInformation161) isEmpty() bool {
	// Check if Ustrd (Unstructured Remittance Info) is set
	if r.Ustrd != nil {
		return false
	}

	// Check if Strd (Structured Remittance Info) contains valid data
	if len(r.Strd) > 0 {
		for _, srItem := range r.Strd {
			if srItem != nil {
				// Check if there are valid referred document information items
				if len(srItem.RfrdDocInf) > 0 {
					for _, refDoc := range srItem.RfrdDocInf {
						if refDoc != nil && (refDoc.Tp != nil || refDoc.Nb != nil || refDoc.RltdDt != nil) {
							return false
						}
					}
				}

				// Check if TaxRmt (Tax Remittance) contains valid data
				if srItem.TaxRmt != nil {
					if srItem.TaxRmt.Cdtr != nil && srItem.TaxRmt.Cdtr.TaxId != nil {
						return false
					}
					if len(srItem.TaxRmt.Rcrd) > 0 {
						for _, taxRecord := range srItem.TaxRmt.Rcrd {
							if taxRecord != nil && (taxRecord.Tp != nil || taxRecord.Prd != nil) {
								return false
							}
						}
					}
				}
			}
		}
	}

	// If none of the fields contain data, return true (struct is empty)
	return true
}
func FinancialInstitutionIdentification181From(agent Agent) FinancialInstitutionIdentification181 {
	var result FinancialInstitutionIdentification181
	if agent.BusinessIdCode != "" {
		_BICFI := BICFIDec2014Identifier(agent.BusinessIdCode)
		result.BICFI = &_BICFI
	}
	if agent.PaymentSysCode != "" || agent.PaymentSysMemberId != "" {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		if agent.PaymentSysCode != "" {
			Cd := ExternalClearingSystemIdentification1Code(agent.PaymentSysCode)
			result.ClrSysMmbId.ClrSysId = ClearingSystemIdentification2Choice1{
				Cd: &Cd,
			}
		}
		if agent.PaymentSysMemberId != "" {
			result.ClrSysMmbId.MmbId = Max35Text(agent.PaymentSysMemberId)
		}
	}
	if agent.BankName != "" {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		_BKNM := Max140Text(agent.BankName)
		result.Nm = &_BKNM
	}
	postalAddress := PostalAddress241From(agent.PostalAddress)
	if !isEmptyPostalAddress241(postalAddress) {
		if result.ClrSysMmbId == nil {
			_resultClrSysMmbId := ClearingSystemMemberIdentification21{}
			result.ClrSysMmbId = &_resultClrSysMmbId
		}
		result.PstlAdr = &postalAddress
	}
	return result
}

func PaymentTypeInformation281From(InstrumentPropCode InstrumentPropCodeType, SericeLevel string) PaymentTypeInformation281 {
	var result PaymentTypeInformation281
	if InstrumentPropCode != "" {
		result.LclInstrm = LocalInstrument2Choice1{}
		CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry := LocalInstrumentFedwireFunds1(InstrumentPropCode)
		result.LclInstrm.Prtry = &CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry
	}
	if SericeLevel != "" {
		svclv := ExternalServiceLevel1Code(SericeLevel)
		CdtTrfTxInf_PmtTpInf_SvcLvl := ServiceLevel8Choice{
			Cd: &svclv,
		}
		result.SvcLvl = []*ServiceLevel8Choice{
			&CdtTrfTxInf_PmtTpInf_SvcLvl,
		}
	}
	return result
}
func RemittanceLocation71From(param RemittanceDetail) RemittanceLocation71 {
	var result RemittanceLocation71
	if param.RemittanceId != "" {
		_RemittanceId := Max35Text(param.RemittanceId)
		result.RmtId = &_RemittanceId
	}
	var locationData RemittanceLocationData11
	var hasLocationData = false
	if param.Method != "" {
		locationData.Mtd = RemittanceLocationMethod2Code(param.Method)
		hasLocationData = true
	}
	if param.ElectronicAddress != "" {
		_ElectronicAddress := Max2048Text(param.ElectronicAddress)
		locationData.ElctrncAdr = &_ElectronicAddress
		hasLocationData = true
	}
	if hasLocationData {
		result.RmtLctnDtls = []*RemittanceLocationData11{
			&locationData,
		}
	}
	return result
}
func (r RemittanceLocation71) isEmpty() bool {
	// Check if RmtId is nil (i.e., no remittance ID is set)
	if r.RmtId != nil {
		return false
	}

	// Check if RmtLctnDtls contains any valid location data
	if len(r.RmtLctnDtls) > 0 {
		for _, loc := range r.RmtLctnDtls {
			if loc != nil && (loc.Mtd != "" || loc.ElctrncAdr != nil) {
				return false
			}
		}
	}

	// If none of the above fields have meaningful data, it's empty
	return true
}
func PartyIdentification1352From(Nm string, PstlAdr PostalAddress) PartyIdentification1352 {
	var result PartyIdentification1352
	if Nm != "" {
		_nm := Max140Text(Nm)
		result.Nm = &_nm
	}
	_PstlAdr := PostalAddress241From(PstlAdr)
	if !isEmptyPostalAddress241(_PstlAdr) {
		result.PstlAdr = &_PstlAdr
	}
	return result
}

func (p PartyIdentification1352) isEmpty() bool {
	if p.Nm == nil {
		return true
	}
	if isEmptyPostalAddress241(*p.PstlAdr) {
		return true
	}
	return false
}
func PartyIdentification1351From(Nm string, PstlAdr PostalAddress) PartyIdentification1351 {
	var result PartyIdentification1351
	if Nm != "" {
		_nm := Max140Text(Nm)
		result.Nm = &_nm
	}
	_PstlAdr := PostalAddress242From(PstlAdr)
	if !isEmptyPostalAddress242(_PstlAdr) {
		result.PstlAdr = &_PstlAdr
	}
	return result
}
func (p PartyIdentification1351) isEmpty() bool {
	if p.Nm == nil {
		return true
	}
	if isEmptyPostalAddress242(*p.PstlAdr) {
		return true
	}
	return false
}
func Charges71From(data ChargeInfo) Charges71 {
	var result Charges71
	if data.amount.Amount != 0 || data.amount.Currency != "" {
		result.Amt = ActiveOrHistoricCurrencyAndAmount{
			Value: ActiveOrHistoricCurrencyAndAmountSimpleType(data.amount.Amount),
			Ccy:   ActiveOrHistoricCurrencyCode(data.amount.Currency),
		}
	}
	if data.BusinessIdCode != "" {
		result.Agt = BranchAndFinancialInstitutionIdentification61{}
		result.Agt.FinInstnId = FinancialInstitutionIdentification181{}
		_BICFI := BICFIDec2014Identifier(data.BusinessIdCode)
		result.Agt.FinInstnId.BICFI = &_BICFI
	}
	return result
}
func (c Charges71) isEmpty() bool {
	if c.Amt.Value != 0 {
		return false
	}
	if c.Agt.FinInstnId.BICFI != nil {
		return false
	}
	return true
}
func BranchAndFinancialInstitutionIdentification61From(BICFI string) BranchAndFinancialInstitutionIdentification61 {
	var result BranchAndFinancialInstitutionIdentification61
	if BICFI != "" {
		result.FinInstnId = FinancialInstitutionIdentification181{}
		_BICFI := BICFIDec2014Identifier(BICFI)
		result.FinInstnId.BICFI = &_BICFI
	}
	return result
}
func (b BranchAndFinancialInstitutionIdentification61) isEmpty() bool {
	return b.FinInstnId.BICFI == nil
}
