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
type SettlementMethodType string
type CommonClearingSysCodeType string
type ChargeBearerType string
type PurposeOfPaymentType string
type RemittanceDeliveryMethod string
type CodeOrProprietaryType string
type PaymentMethod string
type PaymentRequestType string
type StatusReasonInformationCode string
type GapType string
type RelatedStatusCode string
type Status string
type FundEventType string
type AccountReportType string
type BalanceType string
type CreditLineType string
type TransactionCode string

const (
	InputMessageAccountabilityData  GapType = "IMAD"
	OutputMessageAccountabilityData GapType = "OMAD"
)
const (
	BusinessProcessingDate WorkingDayType = "BPRD"
)
const (
	CreditTransform PaymentMethod = "TRF"
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
	DrawDownRequestCredit PaymentRequestType = "DRRC"
	DrawDownRequestDebit  PaymentRequestType = "DRRB"
	IntraCompanyPayment   PaymentRequestType = "INTC"
)
const (
	SchemaValidationFailed RelatedStatusCode = "TS01" // Technical Error
	MessageHeaderIssue     RelatedStatusCode = "TS02"
	BusinessRuleViolation  RelatedStatusCode = "NS01"
	UnknownMessageType     RelatedStatusCode = "NS02" // Unknown Message Type
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
	ReturnRequestAccepted   Status = "CNCL"
	ReturnRequestRejected   Status = "RJCR"
	ReturnRequestPending    Status = "PDCR"
	PartiallyExecutedReturn Status = "PECR"
)

const (
	InstrumentCTRC                      InstrumentPropCodeType = "CTRC" // Credit Transfer (Proprietary Code)
	InstrumentDD                        InstrumentPropCodeType = "DD"   // Direct Debit
	InstrumentStraightThroughProcessing InstrumentPropCodeType = "STP"  // Straight Through Processing
	InstrumentNCT                       InstrumentPropCodeType = "NCT"  // National Credit Transfer
	InstrumentCTRD                      InstrumentPropCodeType = "CTRD" // National Credit Transfer
)

const (
	SettlementCLRG SettlementMethodType = "CLRG" // Settlement via Clearing System (e.g., ACH, SEPA, RTGS)
	SettlementINDA SettlementMethodType = "INDA" // In-House Settlement (within the same bank)
	SettlementCOVE SettlementMethodType = "COVE" // Settlement through a Correspondent Bank
	SettlementTDSO SettlementMethodType = "TDSO" // Settlement via Target2 with a Settlement Agent
	SettlementTDSA SettlementMethodType = "TDSA" // Settlement via Target2 with a Direct Account
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
	ClearingSysFDW   CommonClearingSysCodeType = "FDW"   // Fedwire (U.S.)
	ClearingSysCHIPS CommonClearingSysCodeType = "CHIPS" // CHIPS (U.S. Clearing House Interbank Payments System)
	ClearingSysSEPA  CommonClearingSysCodeType = "SEPA"  // SEPA (Single Euro Payments Area)
	ClearingSysRTGS  CommonClearingSysCodeType = "RTGS"  // Real-Time Gross Settlement
	ClearingSysSWIFT CommonClearingSysCodeType = "SWIFT" // SWIFT Network
	ClearingSysBACS  CommonClearingSysCodeType = "BACS"  // BACS (UK Clearing System)
	ClearingSysCNAPS CommonClearingSysCodeType = "CNAPS" // CNAPS (China’s Clearing System)
)
const (
	ChargeBearerSLEV   ChargeBearerType = "SLEV" // Sender Pays All Charges
	ChargeBearerRECV   ChargeBearerType = "RECV" // Receiver Pays All Charges
	ChargeBearerSHAR   ChargeBearerType = "SHAR" // Shared Charges
	ChargeBearerDEBT   ChargeBearerType = "DEBT" // Shared Charges
	ChargeBearerCREDIT ChargeBearerType = "CRED" // Shared Charges
)
const (
	AccountBalance                        BalanceType = "ABAL"
	AvailableBalanceFromAccountBalance    BalanceType = "AVAL"
	AvailableBalanceFromDaylightOverdraft BalanceType = "AVLD"
	DaylightOverdraftBalance              BalanceType = "DLOD"
	OpeningBalanceFinalBalanceLoaded      BalanceType = "OBFL"
	OpeningBalanceNotLoaded               BalanceType = "OBNL"
	OpeningBalancePriorDayBalanceLoaded   BalanceType = "OBPL"
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
	ABMS        AccountReportType = "ABMS" //Solicited balance report sent by the Federal Reserve Banks in response to an account balance report request.
	FINAL       AccountReportType = "FINL" //Unsolicited balance report sent by the Federal Reserve Banks as part of the Account Balance Services end-of-day process.
	INTERIM     AccountReportType = "ITRM" //Unsolicited balance report sent by the Federal Reserve Banks when operating in contingency mode.
	OPENING     AccountReportType = "OPEN" //Unsolicited balance report sent by the Federal Reserve Banks when opening balance is loaded.
	PERIODIC    AccountReportType = "PRDC" //Unsolicited balance report sent by the Federal Reserve Banks throughout the day whenever the Federal Reserve Banks' accounting system updates Account Balance Services.
	PROVISIONAL AccountReportType = "PROV" //Unsolicited balance report sent by the Federal Reserve Banks when memo post is used.
)
const (
	AvailableAllOtherActivity        TransactionCode = "AVOT"
	FedNowFundsTransfers             TransactionCode = "FDNF"
	FedwireFundsTransfers            TransactionCode = "FDWF"
	FedwireSecuritiesTransfers       TransactionCode = "FDWS"
	MemoPostEntries                  TransactionCode = "MEMO"
	NationalSettlementServiceEntries TransactionCode = "NSSE"
	PrefundedACHCreditItems          TransactionCode = "FDAP"
	UnavailableAllOtherActivity      TransactionCode = "UVOT"
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
	INSCPayment       PurposeOfPaymentType = "INSC"
)
const (
	Fax                       RemittanceDeliveryMethod = "FAXI" //Fax
	ElectronicDataInterchange RemittanceDeliveryMethod = "EDIC" //Electronic Data Interchange (EDI)
	UniformResourceIdentifier RemittanceDeliveryMethod = "URID" //Uniform Resource Identifier (URI)
	PostalMail                RemittanceDeliveryMethod = "POST" //Postal Mail
	Email                     RemittanceDeliveryMethod = "EMAL" //Email
)
const (
	Book     ReportStatus = "BOOK"
	Pending  ReportStatus = "PDNG"
	Received ReportStatus = "RCVD"
	Settled  ReportStatus = "SETT"
)
const (
	CollateralAvailable                CreditLineType = "COLL"
	CollateralizedCapacity             CreditLineType = "CCAP"
	CollateralizedDaylightOverdrafts   CreditLineType = "CLOD"
	NetDebitCap                        CreditLineType = "NCAP"
	UncollateralizedDaylightOverdrafts CreditLineType = "ULOD"
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
	InsufficientFunds         StatusReasonInformationCode = "AM04" // The account does not have enough balance to process the transaction.
	DuplicateTransaction      StatusReasonInformationCode = "AM05" // The transaction appears to be a duplicate of another payment.
	WrongAccount              StatusReasonInformationCode = "AM09" // The account number provided is incorrect or does not exist.
	CreditorBankNotRegistered StatusReasonInformationCode = "CNOR" // The creditor’s bank is unknown or not part of the payment system.
	DebtorBankNotRegistered   StatusReasonInformationCode = "DNOR" // The debtor’s bank is unknown or not part of the payment system.
	InvalidFileFormat         StatusReasonInformationCode = "FF01" // The file submitted does not match the expected format.
	InvalidBIC                StatusReasonInformationCode = "RC01" // The Bank Identifier Code (BIC) provided is incorrect.
	MissingDebtorInfo         StatusReasonInformationCode = "RR01" // Mandatory information about the debtor is missing.
	MissingCreditorInfo       StatusReasonInformationCode = "RR02" // Mandatory information about the creditor is missing.
	CutOffTimeExceeded        StatusReasonInformationCode = "SL01" // The transaction was submitted after the cut-off time for processing.
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

type TotalsPerBankTransaction struct {
	TotalNetEntryAmount  float64
	CreditDebitIndicator CdtDbtInd
	CreditEntries        NumberAndSumOfTransactions
	DebitEntries         NumberAndSumOfTransactions
	BankTransactionCode  TransactionCode
	Date                 time.Time
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
type PartyIdentify struct {
	Name    string
	Address PostalAddress
}
type Reason struct {
	//Party that issues the cancellation request.
	Originator string
	//Specifies the reason for the cancellation.
	Reason string
	//Further details on the cancellation request reason.
	AdditionalInfo string
}
type CreditLine struct {
	//Indicates whether or not the credit line is included in the balance.
	Included bool
	//Type of the credit line provided when multiple credit lines may be provided.
	Type CreditLineType
	//Amount of money of the cash balance.
	Amount CurrencyAndAmount
	//Indicates the date (and time) of the balance.
	DateTime time.Time
}
type Balance struct {
	//Specifies the nature of a balance.
	BalanceTypeId BalanceType

	CdtLines []CreditLine
	//Amount of money of the cash balance.
	Amount CurrencyAndAmount
	//Indicates whether the balance is a credit or a debit balance.
	CreditDebitIndicator CdtDbtInd
	//Indicates the date (and time) of the balance.
	DateTime time.Time
}
type Party struct {
	//Name by which a party is known and which is usually used to identify that party.
	Name string
	//Information that locates and identifies a specific address, as defined by postal services.
	Address PostalAddress
}
type ReturnChain struct {
	//Party that owes an amount of money to the (ultimate) creditor.
	Debtor Party
	//Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorOtherTypeId string
	//Financial institution servicing an account for the debtor.
	DebtorAgent Agent
	//Financial institution servicing an account for the creditor.
	CreditorAgent Agent
	//Party to which an amount of money is due.
	Creditor Party
	//Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccountOtherTypeId string
}
