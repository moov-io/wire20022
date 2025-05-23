package Archive

import "time"

type AccountTypeFRS string
type PaymentSystemType string
type CAMTReportType string
type ReportType string
type TransactionStatusCode string
type InstrumentPropCodeType string
type CdtDbtInd string
type ReportStatus string
type WorkingDayType string

const (
	BusinessProcessingDate WorkingDayType = "BPRD"
)
const (
	AccountTypeSavings  AccountTypeFRS = "S" // "S" for Savings Account
	AccountTypeMerchant AccountTypeFRS = "M" // "M" for Merchant Account
)
const (
	EveryDay ReportType = "EDAY"
	Intraday ReportType = "IDAY"
)
const (
	Credit CdtDbtInd = "CRDT"
	Debit  CdtDbtInd = "DBIT"
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
	InstrumentCTRC                      InstrumentPropCodeType = "CTRC" // Credit Transfer (Proprietary Code)
	InstrumentDD                        InstrumentPropCodeType = "DD"   // Direct Debit
	InstrumentStraightThroughProcessing InstrumentPropCodeType = "STP"  // Straight Through Processing
	InstrumentNCT                       InstrumentPropCodeType = "NCT"  // National Credit Transfer
	InstrumentCTRD                      InstrumentPropCodeType = "CTRD" // National Credit Transfer
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
	Book     ReportStatus = "BOOK"
	Pending  ReportStatus = "PDNG"
	Received ReportStatus = "RCVD"
	Settled  ReportStatus = "SETT"
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

type SequenceRange struct {
	FromSeq string
	ToSeq   string
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
type Agent struct {
	BusinessIdCode     string
	PaymentSysCode     PaymentSystemType
	PaymentSysMemberId string
	BankName           string
	PostalAddress      PostalAddress
	OtherTypeId        string
}
type CurrencyAndAmount struct {
	//default: USD
	Currency string
	Amount   float64
}
type TotalsPerBankTransactionCode struct {
	// NbOfNtries (Number of Entries) specifies the total number of transactions for a given bank transaction code.
	// This helps in categorizing transactions based on their type.
	NumberOfEntries string
	// It is used when the transaction code follows a bank-specific classification rather than a standard one.
	BankTransactionCode TransactionStatusCode
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
