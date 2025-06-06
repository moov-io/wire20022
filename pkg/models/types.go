package model

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	fedwire "github.com/moov-io/fedwire20022/pkg/fedwire"
)

type PaymentSystemType string
type InstrumentPropCodeType string
type TransactionStatusCode string
type CdtDbtInd string
type ReportStatus string
type WorkingDayType string
type CAMTReportType string
type ReportType string
type CodeOrProprietaryType string
type RelatedStatusCode string
type FundEventType string
type SettlementMethodType string
type CommonClearingSysCodeType string
type ChargeBearerType string
type RequestType string

const (
	RequestReceived RequestType = "R"
	RequestSent     RequestType = "S"
)
const (
	EveryDay ReportType = "EDAY"
	Intraday ReportType = "IDAY"
)
const (
	ChargeBearerSLEV   ChargeBearerType = "SLEV" // Sender Pays All Charges
	ChargeBearerRECV   ChargeBearerType = "RECV" // Receiver Pays All Charges
	ChargeBearerSHAR   ChargeBearerType = "SHAR" // Shared Charges
	ChargeBearerDEBT   ChargeBearerType = "DEBT" // Shared Charges
	ChargeBearerCREDIT ChargeBearerType = "CRED" // Shared Charges
)
const (
	//Ad hoc Fedwire Funds Service customized message.
	AdHoc           FundEventType = "ADHC"
	ConnectionCheck FundEventType = "PING"
	SystemClosed    FundEventType = "CLSD"
	SystemExtension FundEventType = "EXTN"
	SystemOpen      FundEventType = "OPEN"
)

const (
	Book     ReportStatus = "BOOK"
	Pending  ReportStatus = "PDNG"
	Received ReportStatus = "RCVD"
	Settled  ReportStatus = "SETT"
)
const (
	Credit CdtDbtInd = "CRDT"
	Debit  CdtDbtInd = "DBIT"
)
const (
	BusinessProcessingDate WorkingDayType = "BPRD"
)
const (
	MessagesInProcess           TransactionStatusCode = "INPR"
	MessagesIntercepted         TransactionStatusCode = "ICPT"
	AcceptedTechnicalValidation TransactionStatusCode = "ACTC"
	AcceptedSettlementInProcess TransactionStatusCode = "ACSP"
	AcceptedWithChange          TransactionStatusCode = "ACWC"
	AcceptedCreditClearing      TransactionStatusCode = "ACCC"
	Sent                        TransactionStatusCode = "SENT"
	TransReceived               TransactionStatusCode = "RCVD"
	Rejected                    TransactionStatusCode = "RJCT"
	TransPending                TransactionStatusCode = "PDNG"
	Cancelled                   TransactionStatusCode = "CANC"
	AcceptedCustomerProfile     TransactionStatusCode = "ACCP"
	PartiallyAccepted           TransactionStatusCode = "PART"
	TransCredit                 TransactionStatusCode = "CRDT"
	TransDebit                  TransactionStatusCode = "DBIT"
	AcceptedSettlementCompleted TransactionStatusCode = "ACSC"
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
	PaymentSysUSABA PaymentSystemType = "USABA" // American Bankers Association (ABA) routing number system
	PaymentSysCHIPS PaymentSystemType = "CHIPS" // Clearing House Interbank Payments System
	PaymentSysSEPA  PaymentSystemType = "SEPA"  // Single Euro Payments Area
	PaymentSysRTGS  PaymentSystemType = "RTGS"  // Real-Time Gross Settlement
	PaymentSysSWIFT PaymentSystemType = "SWIFT" // Society for Worldwide Interbank Financial Telecommunication
	PaymentSysBACS  PaymentSystemType = "BACS"  // Bankers' Automated Clearing Services
)

const (
	AccountBalanceReport          CAMTReportType = "ABAR"
	ActivityReport                CAMTReportType = "ACTR"
	EndpointDetailsReceivedReport CAMTReportType = "DTLR"
	EndpointDetailsSentReport     CAMTReportType = "DTLS"
	EndpointGapReportType         CAMTReportType = "GAPR"
	EndpointTotalsReport          CAMTReportType = "ETOT"
)
const (
	InstrumentCTRC                      InstrumentPropCodeType = "CTRC" // Credit Transfer (Proprietary Code)
	InstrumentDD                        InstrumentPropCodeType = "DD"   // Direct Debit
	InstrumentStraightThroughProcessing InstrumentPropCodeType = "STP"  // Straight Through Processing
	InstrumentNCT                       InstrumentPropCodeType = "NCT"  // National Credit Transfer
	InstrumentCTRD                      InstrumentPropCodeType = "CTRD" // National Credit Transfer
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
	SchemaValidationFailed RelatedStatusCode = "TS01" // Technical Error
	MessageHeaderIssue     RelatedStatusCode = "TS02"
	BusinessRuleViolation  RelatedStatusCode = "NS01"
	UnknownMessageType     RelatedStatusCode = "NS02" // Unknown Message Type
)

type ValidateError struct {
	ParentPath []string
	ParamName  string
	Message    string
}

func (e *ValidateError) Error() string {
	if e == nil {
		return ""
	}
	fullPath := e.ParamName
	if len(e.ParentPath) > 0 {
		fullPath = strings.Join(e.ParentPath, ".") + "." + e.ParamName
	}
	return fmt.Sprintf("error occur at %s: %s", fullPath, e.Message)
}

func (e *ValidateError) ToError() error {
	if e == nil {
		return nil
	}
	return e
}

func (e *ValidateError) InsertPath(path string) {
	if len(e.ParentPath) == 0 {
		e.ParentPath = []string{path}
	} else {
		e.ParentPath = slices.Insert(e.ParentPath, 0, path)
	}
}

type CurrencyAndAmount struct {
	//default: USD
	Currency string
	Amount   float64
}
type SequenceRange struct {
	FromSeq string
	ToSeq   string
}
type MessagePagenation struct {
	// PgNb (Page Number) indicates the current page of the report.
	// It is used for paginated messages where multiple pages exist.
	PageNumber string
	// LastPgInd (Last Page Indicator) specifies whether this is the last page of the report.
	// A value of 'true' means this is the final page, while 'false' means more pages follow.
	LastPageIndicator bool
}
type NumberAndSumOfTransactions struct {
	// NbOfNtries (Number of Entries) specifies the total count of transactions reported.
	// This value represents the total number of individual transactions included in the report.
	NumberOfEntries string
	// Sum represents the total monetary value of all transactions included in the report.
	// It aggregates the values of individual transactions to provide a summary amount.
	Sum float64
}
type NumberAndStatusOfTransactions struct {
	// NbOfNtries (Number of Entries) specifies the total count of transactions reported.
	// This value represents the total number of individual transactions included in the report.
	NumberOfEntries string
	// Proprietary bank transaction code to identify the underlying transaction.
	Status TransactionStatusCode
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

type PartyIdentify struct {
	Name    string
	Address PostalAddress
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

type FiniancialInstitutionId struct {
	BusinessId             string
	ClearingSystemId       PaymentSystemType
	ClearintSystemMemberId string
	Name                   string
	Address                PostalAddress
}
type Entry struct {
	// Amt (Amount) specifies the transaction amount along with the currency.
	// It represents the value of the transaction.
	Amount CurrencyAndAmount
	// CdtDbtInd (Credit or Debit Indicator) specifies whether the transaction is a credit (CRDT) or debit (DBIT).
	CreditDebitIndicator CdtDbtInd
	// Sts (Status) represents the current status of the transaction entry.
	// It may indicate if the transaction is booked, pending, or in another state
	Status ReportStatus
	// BkTxCd (Bank Transaction Code) defines the type of transaction.
	// It provides structured information about the nature of the transaction (e.g., deposit, withdrawal, fee).
	BankTransactionCode TransactionStatusCode
	// <MsgNmId> (Message Name Identification) specifies the identifier for the message type or
	// message version. In this case, "pacs.008.001.08" is the identifier for a specific type of
	// payment message, indicating it is a version 08 of the pacs.008 (Payment Initiation Request) message.
	MessageNameId string
	//Provides details on the entry.
	EntryDetails EntryDetail
}
type EntryDetail struct {
	// MsgId (Message ID) represents the unique identifier for the message.
	// It is used to track the specific transaction message.
	MessageId string
	// InstrId (Instruction ID) is an optional field representing a reference for the instruction associated with the transaction.
	// It can be used to identify a particular instruction within the system.
	InstructionId string
	// UETR (Unique End-to-End Transaction Reference) is an optional field providing a globally unique reference for the transaction.
	// It is typically used to track and identify a specific transaction across different systems.
	UniqueTransactionReference string
	// ClrSysRef (Clearing System Reference) is an optional field used to provide a reference to the clearing system that processes the transaction.
	// It is specific to the system used for clearing and settlement of the transaction.
	ClearingSystemRef string
	// InstgAgt (Instructing Agent) represents the financial institution or branch that is instructing the transaction.
	// This is the party that initiates the transaction and sends the payment instructions.
	InstructingAgent Agent
	// InstdAgt (Instructed Agent) represents the financial institution or branch that is instructed to process the transaction.
	// This is the party that receives the payment instructions and is responsible for executing the transaction.
	InstructedAgent Agent
	// LclInstrm (Local Instrument) is an optional field that refers to a local instrument or payment method for the transaction.
	// It indicates how the transaction is to be processed (e.g., via a local payment system).
	LocalInstrumentChoice InstrumentPropCodeType
	//Tp (Type) indicates the type of the related date. In this case, 'BPRD' could represent a specific type of related date, like business processing date.
	RelatedDatesProprietary WorkingDayType
	RelatedDateTime         time.Time
}
type Date struct {
	Year  int
	Month int
	Day   int
}

func FromTime(t time.Time) Date {
	return Date{
		Year:  t.Year(),
		Month: int(t.Month()),
		Day:   t.Day(),
	}
}
func FromDate(t fedwire.ISODate) Date {
	return Date{
		Year:  t.Year,
		Month: int(t.Month),
		Day:   t.Day,
	}
}
func (d Date) Date() fedwire.ISODate {
	return fedwire.ISODate(civil.Date{
		Year:  d.Year,
		Month: time.Month(d.Month),
		Day:   d.Day,
	})
}
func (r CAMTReportType) Validate() error {
	switch r {
	case AccountBalanceReport,
		ActivityReport,
		EndpointDetailsReceivedReport,
		EndpointDetailsSentReport,
		EndpointGapReportType,
		EndpointTotalsReport:
		return nil
	}
	return fmt.Errorf("invalid CAMT report type: %s", r)
}
