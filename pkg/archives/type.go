package Archive

type AccountTypeFRS string
type PaymentSystemType string
type CAMTReportType string

const (
	AccountTypeSavings  AccountTypeFRS = "S" // "S" for Savings Account
	AccountTypeMerchant AccountTypeFRS = "M" // "M" for Merchant Account
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
