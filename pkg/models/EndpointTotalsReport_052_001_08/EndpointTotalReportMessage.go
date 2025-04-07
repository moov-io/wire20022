package EndpointTotalsReport_052_001_08

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	EndpointTotalsReport "github.com/moov-io/fedwire20022/gen/EndpointTotalsReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type CamtTotal052 struct {
	//Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId model.CAMTReportType
	// Date and time at which the message was created.
	CreatedDateTime time.Time
	//Provides details on the page number of the message.
	MessagePagination model.MessagePagenation
	//Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	ReportId model.ReportType
	//Date and time at which the report was created.
	ReportCreateDateTime time.Time
	// /Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId string
	//Specifies the total number and sum of credit entries.
	TotalCreditEntries model.NumberAndSumOfTransactions
	TotalDebitEntries  model.NumberAndSumOfTransactions

	//Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	//Further details of the account report.
	AdditionalReportInfo string
}
type CamtTotal0522Message struct {
	model CamtTotal052
	doc   EndpointTotalsReport.Document
}

func NewCamtTotal0522Message() CamtTotal0522Message {
	return CamtTotal0522Message{
		model: CamtTotal052{},
	}
}
func (msg *CamtTotal0522Message) CreateDocument() {
	msg.doc = EndpointTotalsReport.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08",
			Local: "Document",
		},
	}
	var BkToCstmrAcctRpt EndpointTotalsReport.BankToCustomerAccountReportV08
	var GrpHdr EndpointTotalsReport.GroupHeader811
	if msg.model.MessageId != "" {
		GrpHdr.MsgId = EndpointTotalsReport.AccountReportingFedwireFunds1(msg.model.MessageId)
	}
	if !isEmpty(msg.model.CreatedDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.model.CreatedDateTime)
	}
	if !isEmpty(msg.model.MessagePagination) {
		GrpHdr.MsgPgntn = EndpointTotalsReport.Pagination1{
			PgNb:      EndpointTotalsReport.Max5NumericText(msg.model.MessagePagination.PageNumber),
			LastPgInd: EndpointTotalsReport.YesNoIndicator(msg.model.MessagePagination.LastPageIndicator),
		}
	}
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}

	var Rpt EndpointTotalsReport.AccountReport251
	if msg.model.ReportId != "" {
		Rpt.Id = EndpointTotalsReport.ReportTimingFRS1(msg.model.ReportId)
	}
	if !isEmpty(msg.model.ReportCreateDateTime) {
		Rpt.CreDtTm = fedwire.ISODateTime(msg.model.ReportCreateDateTime)
	}
	if msg.model.AccountOtherId != "" {
		Othr := EndpointTotalsReport.GenericAccountIdentification11{
			Id: EndpointTotalsReport.EndpointIdentifierFedwireFunds1(msg.model.AccountOtherId),
		}
		Rpt.Acct = EndpointTotalsReport.CashAccount391{
			Id: EndpointTotalsReport.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	var TxsSummry EndpointTotalsReport.TotalTransactions61
	if !isEmpty(msg.model.TotalCreditEntries) {
		TxsSummry.TtlCdtNtries = EndpointTotalsReport.NumberAndSumOfTransactions11{
			NbOfNtries: EndpointTotalsReport.Max15NumericText(msg.model.TotalCreditEntries.NumberOfEntries),
			Sum:        EndpointTotalsReport.DecimalNumber(msg.model.TotalCreditEntries.Sum),
		}
	}
	if !isEmpty(msg.model.TotalDebitEntries) {
		TxsSummry.TtlDbtNtries = EndpointTotalsReport.NumberAndSumOfTransactions11{
			NbOfNtries: EndpointTotalsReport.Max15NumericText(msg.model.TotalDebitEntries.NumberOfEntries),
			Sum:        EndpointTotalsReport.DecimalNumber(msg.model.TotalDebitEntries.Sum),
		}
	}

	if !isEmpty(msg.model.TotalEntriesPerTransactionCode) {
		var TtlNtriesPerBkTxCd []EndpointTotalsReport.TotalsPerBankTransactionCode51
		for _, entity := range msg.model.TotalEntriesPerTransactionCode {
			item := EndpointTotalsReport.TotalsPerBankTransactionCode51{
				NbOfNtries: EndpointTotalsReport.Max15NumericText(entity.NumberOfEntries),
				BkTxCd: EndpointTotalsReport.BankTransactionCodeStructure41{
					Prtry: EndpointTotalsReport.ProprietaryBankTransactionCodeStructure11{
						Cd: EndpointTotalsReport.BankTransactionCodeFedwireFunds1(entity.Status),
					},
				},
			}
			TtlNtriesPerBkTxCd = append(TtlNtriesPerBkTxCd, item)
		}
		if !isEmpty(TtlNtriesPerBkTxCd) {
			TxsSummry.TtlNtriesPerBkTxCd = TtlNtriesPerBkTxCd
		}
	}

	if !isEmpty(TxsSummry) {
		Rpt.TxsSummry = TxsSummry
	}
	if msg.model.AdditionalReportInfo != "" {
		Rpt.AddtlRptInf = EndpointTotalsReport.Max500Text(msg.model.AdditionalReportInfo)
	}
	if !isEmpty(Rpt) {
		BkToCstmrAcctRpt.Rpt = Rpt
	}
	if !isEmpty(BkToCstmrAcctRpt) {
		msg.doc.BkToCstmrAcctRpt = BkToCstmrAcctRpt
	}
}
func (msg *CamtTotal0522Message) GetXML() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(msg.doc, "", "\t")
	if err != nil {
		return nil, err
	}

	// Convert byte slice to string for manipulation
	xmlString := string(xmlData)

	// Keep the xmlns only in the <Document> tag, remove from others
	xmlString = removeExtraXMLNS(xmlString)

	// Regular expression to match scientific notation (e.g., 9.93229443e+06)
	re := regexp.MustCompile(`>(\d+\.\d+(?:e[+-]?\d+)?|\d+e[+-]?\d+)<`)

	// Replace matched numbers with properly formatted ones
	xmlString = re.ReplaceAllStringFunc(xmlString, func(match string) string {
		// Extract the number inside the tags
		numberStr := strings.Trim(match, "<>")

		// Convert to float
		number, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			return match // Return the original string if conversion fails
		}

		// Format it as a standard decimal number with 2 decimal places
		return fmt.Sprintf(">%.2f<", number)
	})

	re = regexp.MustCompile(`<(FrSeq|ToSeq)>(\d+)</(FrSeq|ToSeq)>`)

	// Replace numeric values with zero-padded format (6 digits)
	xmlString = re.ReplaceAllStringFunc(xmlString, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 4 {
			num := parts[2] // Extract number as string
			return fmt.Sprintf("<%s>%06s</%s>", parts[1], num, parts[3])
		}
		return match
	})

	// Convert back to []byte
	return []byte(xmlString), nil
	// return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *CamtTotal0522Message) GetJson() ([]byte, error) {
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
