package ActivityReport

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
	"time"

	camt052 "github.com/moov-io/fedwire20022/gen/ActivityReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageModel struct {
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
	ReportType           model.ReportType
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
type Message struct {
	data MessageModel
	doc  camt052.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = camt052.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08",
			Local: "Document",
		},
		BkToCstmrAcctRpt: camt052.BankToCustomerAccountReportV08{
			GrpHdr: camt052.GroupHeader811{
				MsgId:   camt052.AccountReportingFedwireFunds1(msg.data.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.data.CreatedDateTime),
			},
		},
	}
	if !isEmpty(msg.data.Pagenation) {
		msg.doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn = camt052.Pagination1{
			PgNb:      camt052.Max5NumericText(msg.data.Pagenation.PageNumber),
			LastPgInd: camt052.YesNoIndicator(msg.data.Pagenation.LastPageIndicator),
		}
	}
	var Rpt camt052.AccountReport251
	if msg.data.ReportType != "" {
		Rpt.Id = camt052.ReportTimingFRS1(msg.data.ReportType)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		Rpt.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
	var Acct camt052.CashAccount391
	if msg.data.AccountOtherId != "" {
		_Othr := camt052.GenericAccountIdentification11{
			Id: camt052.RoutingNumberFRS1(msg.data.AccountOtherId),
		}
		Acct.Id = camt052.AccountIdentification4Choice1{
			Othr: &_Othr,
		}
	}
	if !isEmpty(Acct) {
		Rpt.Acct = Acct
	}
	var TxsSummry camt052.TotalTransactions61
	if !isEmpty(msg.data.TotalEntries) {
		TxsSummry.TtlNtries = camt052.NumberAndSumOfTransactions41{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalEntries),
		}
	}
	if !isEmpty(msg.data.TotalCreditEntries) {
		TxsSummry.TtlCdtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalCreditEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalCreditEntries.Sum),
		}
	}
	if !isEmpty(msg.data.TotalDebitEntries) {
		TxsSummry.TtlDbtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalDebitEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalDebitEntries.Sum),
		}
	}
	if !isEmpty(msg.data.TotalEntriesPerBankTransactionCode) {
		var TtlNtriesPerBkTxCd []camt052.TotalsPerBankTransactionCode51
		for _, entity := range msg.data.TotalEntriesPerBankTransactionCode {
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
	var Ntry []*camt052.ReportEntry101
	if !isEmpty(msg.data.EntryDetails) {
		for _, entity := range msg.data.EntryDetails {
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
func (msg *Message) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	xmlData, err := xml.MarshalIndent(msg.doc, "", "\t")
	if err != nil {
		return err
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
	return e.EncodeToken(xml.CharData([]byte(xmlString)))
	// return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *Message) MarshalJSON() ([]byte, error) {
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
