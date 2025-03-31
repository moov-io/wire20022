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

type CurrencyAndAmount struct {
	//default: USD
	Currency string
	Amount   float64
}
type Agent struct {
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
	Floor          string
	RoomNumber     string
	PostalCode     string
	TownName       string
	Subdivision    string
	Country        string
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
		address.Flr == nil &&
		address.Room == nil &&
		address.PstCd == nil &&
		address.TwnNm == nil &&
		address.CtrySubDvsn == nil &&
		address.Ctry == nil
}
func CashAccount38From(param string) CashAccount38 {
	if param == "" {
		return CashAccount38{} // Return empty struct if input is empty
	}

	return CashAccount38{
		Id: AccountIdentification4Choice{
			Othr: &GenericAccountIdentification1{
				Id: Max34Text(param),
			},
		},
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

	// If RD_item has data, add it to structured remittance info
	if hasRDData {
		SR_item := StructuredRemittanceInformation161{
			RfrdDocInf: []*ReferredDocumentInformation71{
				&RD_item,
			},
		}
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
	// Check if both Unstructured Remit Info and Structured Remittance Info are empty
	if r.Ustrd == nil && len(r.Strd) == 0 {
		return true
	}
	if r.Ustrd != nil {
		return false
	}
	// Check if Structured Remittance Information contains empty or invalid data
	for _, strdItem := range r.Strd {
		if strdItem != nil && len(strdItem.RfrdDocInf) > 0 {
			// Check if each referred document has any value in Tp, Nb, or RltdDt
			for _, refDoc := range strdItem.RfrdDocInf {
				if refDoc != nil && (refDoc.Tp != nil || refDoc.Nb != nil || refDoc.RltdDt != nil) {
					// If any refDoc has a value, it's not empty
					return false
				}
			}
		}
	}

	// Return true if everything is empty
	return true
}
