package CustomerCreditTransfer_pacs_008_001_08

import (
	"time"

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
func newCurrencyAndAmount() CurrencyAndAmount {
	return CurrencyAndAmount{
		Currency: "USD",
		Amount:   0,
	}
}
func newPostalAddress() PostalAddress {
	return PostalAddress{
		StreetName:     "",
		BuildingNumber: "",
		RoomNumber:     "",
		PostalCode:     "",
		TownName:       "",
		Subdivision:    "",
		Country:        "",
	}
}
func newAgent() Agent {
	return Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "",
		BankName:           "",
		PostalAddress:      newPostalAddress(),
	}
}
func newRemittanceDocument() RemittanceDocument {
	return RemittanceDocument{
		CodeOrProprietary: CodeCINV,
		Number:            "",
		RelatedDate:       civil.DateOf(time.Now()),
	}
}
func NewCustomerCreditTransfer() CustomerCreditTransfer {
	return CustomerCreditTransfer{
		MessageId:                    "",
		CreatedDateTime:              time.Now(),
		NumberOfTransactions:         1,
		SettlementMethod:             SettlementCLRG,
		CommonClearingSysCode:        ClearingSysFDW,
		InstructionId:                "",
		EndToEndId:                   "",
		UniqueEndToEndTransactionRef: "",
		InstrumentPropCode:           InstrumentCTRC,
		InterBankSettAmount:          newCurrencyAndAmount(),
		InterBankSettDate:            civil.DateOf(time.Now()),
		InstructedAmount:             newCurrencyAndAmount(),
		ChargeBearer:                 ChargeBearerSLEV,
		InstructingAgents:            newAgent(),
		InstructedAgent:              newAgent(),
		DebtorName:                   "",
		DebtorAddress:                newPostalAddress(),
		DebtorOtherTypeId:            "",
		DebtorAgent:                  newAgent(),
		CreditorAgent:                newAgent(),
		CreditorName:                 "",
		CreditorPostalAddress:        newPostalAddress(),
		CreditorOtherTypeId:          "",
		RemittanceInfor:              newRemittanceDocument(),
	}
}
func PostalAddress241From(param PostalAddress) PostalAddress241 {
	StrtNm := Max70Text(param.StreetName)
	BldgNb := Max16Text(param.BuildingNumber)
	Floor := Max70Text(param.Floor)
	Room := Max70Text(param.RoomNumber)
	PstCd := Max16Text(param.PostalCode)
	TwnNm := Max35Text(param.TownName)
	CtrySubDvsn := Max35Text(param.Subdivision)
	Ctry := CountryCode(param.Country)
	Dbtr_PstlAdr := PostalAddress241{
		StrtNm:      &StrtNm,
		BldgNb:      &BldgNb,
		Flr:         &Floor,
		Room:        &Room,
		PstCd:       &PstCd,
		TwnNm:       &TwnNm,
		CtrySubDvsn: &CtrySubDvsn,
		Ctry:        &Ctry,
	}
	return Dbtr_PstlAdr
}
func CashAccount38From(param string) CashAccount38 {
	Othr := GenericAccountIdentification1{
		Id: Max34Text(param),
	}
	return CashAccount38{
		Id: AccountIdentification4Choice{
			Othr: &Othr,
		},
	}
}
func ClearingSystemMemberIdentification21From(param PaymentSystemType, paymentSysMemberId string) ClearingSystemMemberIdentification21 {
	Cd := ExternalClearingSystemIdentification1Code(param)
	return ClearingSystemMemberIdentification21{
		ClrSysId: ClearingSystemIdentification2Choice1{
			Cd: &Cd,
		},
		MmbId: Max35Text(paymentSysMemberId),
	}
}
func RemittanceInformation161From(doc RemittanceDocument) RemittanceInformation161 {
	RD_item_Tp_Cd := DocumentType6Code(doc.CodeOrProprietary)
	RD_item_Tp := ReferredDocumentType4{
		CdOrPrtry: ReferredDocumentType3Choice{
			Cd: &RD_item_Tp_Cd,
		},
	}
	RD_item_Nb := Max35Text(doc.Number)
	RD_item_RltdDt := fedwire.ISODate(doc.RelatedDate)
	RD_item := ReferredDocumentInformation71{
		Tp:     &RD_item_Tp,
		Nb:     &RD_item_Nb,
		RltdDt: &RD_item_RltdDt,
	}
	SR_item := StructuredRemittanceInformation161{
		RfrdDocInf: []*ReferredDocumentInformation71{
			&RD_item,
		},
	}
	return RemittanceInformation161{
		Strd: []*StructuredRemittanceInformation161{
			&SR_item,
		},
	}
}
