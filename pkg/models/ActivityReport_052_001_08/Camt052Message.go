package ActivityReport_052_001_08

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
	"time"

	ActivityReport "github.com/moov-io/fedwire20022/gen/ActivityReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Camt052 struct {
	//MessageId (Message Identification) is a unique identifier assigned to an entire message.
	MessageId string
	//CreatedDateTime represents the timestamp when a message, instruction, or transaction was created
	//ISO 8601 format
	CreatedDateTime time.Time
	// MsgPgntn (Message Pagination) provides details about the pagination of the report.
	// It helps in handling reports split across multiple pages.
	Pagenation model.MessagePagenation
	// Id (Report Identification) uniquely identifies the report.
	// It provides a reference to the specific report being generated or requested.
	// Example value: "EDAY" (End-of-Day Report).
	ReportType           ReportType
	ReportCreateDateTime time.Time
	//// Acct (Account Information) contains details about the account being reported on.
	AccountOtherId string
	// TtlNtries (Total Entries) represents the overall count and sum of all transactions in the report.
	// This includes both credit and debit transactions.
	TotalEntries string
	// TtlCdtNtries (Total Credit Entries) specifies the total number and sum of credit transactions.
	// It represents transactions where funds are credited to an account.
	TotalCreditEntries model.NumberAndSumOfTransactions
	// TtlDbtNtries (Total Debit Entries) specifies the total number and sum of debit transactions.
	// It represents transactions where funds are debited from an account.
	TotalDebitEntries model.NumberAndSumOfTransactions
	// TtlNtriesPerBkTxCd (Total Entries Per Bank Transaction Code) provides a breakdown of transactions.
	// It groups transactions based on bank-specific transaction codes.
	TotalEntriesPerBankTransactionCode []TotalsPerBankTransactionCode
	// NtryDtls (Entry Details) contains detailed breakdowns of the transaction.
	// This can include supplementary details such as references, related transactions, and remittance information.
	EntryDetails []model.Entry
}
type Camt052Message struct {
	model Camt052
	doc   ActivityReport.Document
}

func NewCamt052Message() Camt052Message {
	return Camt052Message{
		model: Camt052{},
	}
}
func (msg *Camt052Message) CreateDocument() {
	msg.doc = ActivityReport.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08",
			Local: "Document",
		},
		BkToCstmrAcctRpt: ActivityReport.BankToCustomerAccountReportV08{
			GrpHdr: ActivityReport.GroupHeader811{
				MsgId:   ActivityReport.AccountReportingFedwireFunds1(msg.model.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.model.CreatedDateTime),
			},
		},
	}
	if !isEmpty(msg.model.Pagenation) {
		msg.doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn = ActivityReport.Pagination1{
			PgNb:      ActivityReport.Max5NumericText(msg.model.Pagenation.PageNumber),
			LastPgInd: ActivityReport.YesNoIndicator(msg.model.Pagenation.LastPageIndicator),
		}
	}
	var Rpt ActivityReport.AccountReport251
	if msg.model.ReportType != "" {
		Rpt.Id = ActivityReport.ReportTimingFRS1(msg.model.ReportType)
	}
	if !isEmpty(msg.model.CreatedDateTime) {
		Rpt.CreDtTm = fedwire.ISODateTime(msg.model.CreatedDateTime)
	}
	var Acct ActivityReport.CashAccount391
	if msg.model.AccountOtherId != "" {
		_Othr := ActivityReport.GenericAccountIdentification11{
			Id: ActivityReport.RoutingNumberFRS1(msg.model.AccountOtherId),
		}
		Acct.Id = ActivityReport.AccountIdentification4Choice1{
			Othr: &_Othr,
		}
	}
	if !isEmpty(Acct) {
		Rpt.Acct = Acct
	}
	var TxsSummry ActivityReport.TotalTransactions61
	if !isEmpty(msg.model.TotalEntries) {
		TxsSummry.TtlNtries = ActivityReport.NumberAndSumOfTransactions41{
			NbOfNtries: ActivityReport.Max15NumericText(msg.model.TotalEntries),
		}
	}
	if !isEmpty(msg.model.TotalCreditEntries) {
		TxsSummry.TtlCdtNtries = ActivityReport.NumberAndSumOfTransactions11{
			NbOfNtries: ActivityReport.Max15NumericText(msg.model.TotalCreditEntries.NumberOfEntries),
			Sum:        ActivityReport.DecimalNumber(msg.model.TotalCreditEntries.Sum),
		}
	}
	if !isEmpty(msg.model.TotalDebitEntries) {
		TxsSummry.TtlDbtNtries = ActivityReport.NumberAndSumOfTransactions11{
			NbOfNtries: ActivityReport.Max15NumericText(msg.model.TotalDebitEntries.NumberOfEntries),
			Sum:        ActivityReport.DecimalNumber(msg.model.TotalDebitEntries.Sum),
		}
	}
	if !isEmpty(msg.model.TotalEntriesPerBankTransactionCode) {
		var TtlNtriesPerBkTxCd []ActivityReport.TotalsPerBankTransactionCode51
		for _, entity := range msg.model.TotalEntriesPerBankTransactionCode {
			_item := TotalsPerBankTransactionCode51From(entity)
			TtlNtriesPerBkTxCd = append(TtlNtriesPerBkTxCd, _item)
		}
		if !isEmpty(TtlNtriesPerBkTxCd) {
			TxsSummry.TtlNtriesPerBkTxCd = TtlNtriesPerBkTxCd
		}

	}
	if !isEmpty(TxsSummry) {
		Rpt.TxsSummry = &TxsSummry
	}
	var Ntry []*ActivityReport.ReportEntry101
	if !isEmpty(msg.model.EntryDetails) {
		for _, entity := range msg.model.EntryDetails {
			_item := ReportEntry101From(entity)
			Ntry = append(Ntry, &_item)
		}
	}
	if !isEmpty(Ntry) {
		Rpt.Ntry = Ntry
	}
	if !isEmpty(Rpt) {
		msg.doc.BkToCstmrAcctRpt.Rpt = Rpt
	}
}
func (msg *Camt052Message) GetXML() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(msg.doc, "", "\t")
	if err != nil {
		return nil, err
	}

	// Convert byte slice to string for manipulation
	xmlString := string(xmlData)

	// Keep the xmlns only in the <Document> tag, remove from others
	xmlString = removeExtraXMLNS(xmlString)

	// Regular expression to match scientific notation (e.g., 9.93229443e+06)
	re := regexp.MustCompile(`>(\d+\.\d+e[+-]\d+)<`)

	// Replace scientific notation with properly formatted numbers
	xmlString = re.ReplaceAllStringFunc(xmlString, func(match string) string {
		// Extract the number inside the tags
		numberStr := strings.Trim(match, "<>")

		// Convert to float
		var number float64
		fmt.Sscanf(numberStr, "%e", &number)

		// Format it as a standard decimal number with 2 decimal places
		return fmt.Sprintf(">%.2f<", number)
	})

	// Convert back to []byte
	return []byte(xmlString), nil
	// return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *Camt052Message) GetJson() ([]byte, error) {
	return json.MarshalIndent(msg.doc.BkToCstmrAcctRpt, "", " ")
}

func removeExtraXMLNS(xmlStr string) string {
	// Find the first occurrence of <Document ...> (keep this)
	docStart := strings.Index(xmlStr, "<Document")
	if docStart == -1 {
		return xmlStr // Return original if <Document> not found
	}

	// Find the end of the <Document> opening tag
	docEnd := strings.Index(xmlStr[docStart:], ">")
	if docEnd == -1 {
		return xmlStr
	}
	docEnd += docStart // Adjust index

	// Remove all occurrences of xmlns="..." except in <Document>
	cleanXML := xmlStr[:docEnd+1] + // Keep <Document> with its xmlns
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"`, "")

	return cleanXML
}
