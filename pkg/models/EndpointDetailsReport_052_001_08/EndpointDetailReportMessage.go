package EndpointDetailsReport_052_001_08

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	EndpointDetailsReport "github.com/moov-io/fedwire20022/gen/EndpointDetailsReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Camt052 struct {
	// /Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId string
	//This is the calendar date and time in New York City (Eastern Time) when the message is created by the Fedwire Funds Service application. Time is in 24-hour clock format and includes the offset against the Coordinated Universal Time (UTC).
	CreationDateTime time.Time
	//Provides details on the page number of the message.
	MessagePagination model.MessagePagenation
	// Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.
	BussinessQueryMsgId string
	//Specifies the query message name identifier to which the message refers.
	BussinessQueryMsgNameId string
	//Date and time at which the message was created.
	BussinessQueryCreateDatetime time.Time
	//Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	ReportId model.ReportType
	//Specifies the range of identification sequence numbers, as provided in the request.
	ReportingSequence model.SequenceRange
	//Date and time at which the report was created.
	ReportCreateDateTime time.Time
	//Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId string
	//Specifies the total number and sum of debit entries.
	TotalDebitEntries model.NumberAndSumOfTransactions
	//Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	//Provides details on the entry.
	EntryDetails []model.Entry
}
type Camt052Message struct {
	model Camt052
	doc   EndpointDetailsReport.Document
}

func NewCamt052Message() Camt052Message {
	return Camt052Message{
		model: Camt052{},
	}
}
func (msg *Camt052Message) CreateDocument() {
	msg.doc = EndpointDetailsReport.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08",
			Local: "Document",
		},
	}
	var BkToCstmrAcctRpt EndpointDetailsReport.BankToCustomerAccountReportV08
	var GrpHdr EndpointDetailsReport.GroupHeader811
	// MessageId string
	if msg.model.MessageId != "" {
		GrpHdr.MsgId = EndpointDetailsReport.AccountReportingFedwireFunds1(msg.model.MessageId)
	}
	// CreationDateTime time.Time
	if !isEmpty(msg.model.CreationDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.model.CreationDateTime)
	}
	// MessagePagination model.MessagePagenation
	if !isEmpty(msg.model.MessagePagination) {
		GrpHdr.MsgPgntn = EndpointDetailsReport.Pagination1{
			PgNb:      EndpointDetailsReport.Max5NumericText(msg.model.MessagePagination.PageNumber),
			LastPgInd: EndpointDetailsReport.YesNoIndicator(msg.model.MessagePagination.LastPageIndicator),
		}
	}
	var OrgnlBizQry EndpointDetailsReport.OriginalBusinessQuery11
	// BussinessQueryMsgId string
	if msg.model.BussinessQueryMsgId != "" {
		OrgnlBizQry.MsgId = EndpointDetailsReport.Max35Text(msg.model.BussinessQueryMsgId)
	}
	// BussinessQueryMsgNameId string
	if msg.model.BussinessQueryMsgNameId != "" {
		OrgnlBizQry.MsgNmId = EndpointDetailsReport.MessageNameIdentificationFRS1(msg.model.BussinessQueryMsgNameId)
	}
	// BussinessQueryCreateDatetime time.Time
	if !isEmpty(msg.model.BussinessQueryCreateDatetime) {
		OrgnlBizQry.CreDtTm = fedwire.ISODateTime(msg.model.BussinessQueryCreateDatetime)
	}
	if !isEmpty(OrgnlBizQry) {
		GrpHdr.OrgnlBizQry = OrgnlBizQry
	}
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}
	var Rpt EndpointDetailsReport.AccountReport251
	// ReportId ReportDayType
	if msg.model.ReportId != "" {
		Rpt.Id = EndpointDetailsReport.ReportTimingFRS1(msg.model.ReportId)
	}
	// ReportingSequence model.SequenceRange
	if !isEmpty(msg.model.ReportingSequence) {
		FrToSeq := EndpointDetailsReport.SequenceRange11{
			FrSeq: EndpointDetailsReport.XSequenceNumberFedwireFunds1(msg.model.ReportingSequence.FromSeq),
			ToSeq: EndpointDetailsReport.XSequenceNumberFedwireFunds1(msg.model.ReportingSequence.ToSeq),
		}
		Rpt.RptgSeq = EndpointDetailsReport.SequenceRange1Choice1{
			FrToSeq: &FrToSeq,
		}
	}
	// ReportCreateDateTime time.Time
	if !isEmpty(msg.model.ReportCreateDateTime) {
		CreDtTm := fedwire.ISODateTime(msg.model.ReportCreateDateTime)
		Rpt.CreDtTm = &CreDtTm
	}
	// AccountOtherId string
	if msg.model.AccountOtherId != "" {
		Othr := EndpointDetailsReport.GenericAccountIdentification11{
			Id: EndpointDetailsReport.EndpointIdentifierFedwireFunds1(msg.model.AccountOtherId),
		}
		Rpt.Acct = EndpointDetailsReport.CashAccount391{
			Id: EndpointDetailsReport.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	// TotalDebitEntries model.NumberAndSumOfTransactions
	var TxsSummry EndpointDetailsReport.TotalTransactions61
	if !isEmpty(msg.model.TotalDebitEntries) {
		TtlDbtNtries := EndpointDetailsReport.NumberAndSumOfTransactions11{
			NbOfNtries: EndpointDetailsReport.Max15NumericText(msg.model.TotalDebitEntries.NumberOfEntries),
			Sum:        EndpointDetailsReport.DecimalNumber(msg.model.TotalDebitEntries.Sum),
		}
		TxsSummry.TtlDbtNtries = &TtlDbtNtries
	}
	// TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	if !isEmpty(msg.model.TotalEntriesPerTransactionCode) {
		var TtlNtriesPerBkTxCd []EndpointDetailsReport.TotalsPerBankTransactionCode51
		for _, entity := range msg.model.TotalEntriesPerTransactionCode {
			bankTrans := EndpointDetailsReport.TotalsPerBankTransactionCode51{
				NbOfNtries: EndpointDetailsReport.Max15NumericText(entity.NumberOfEntries),
				BkTxCd: EndpointDetailsReport.BankTransactionCodeStructure41{
					Prtry: EndpointDetailsReport.ProprietaryBankTransactionCodeStructure11{
						Cd: EndpointDetailsReport.BankTransactionCodeFedwireFunds11(entity.Status),
					},
				},
			}
			TtlNtriesPerBkTxCd = append(TtlNtriesPerBkTxCd, bankTrans)
		}
		if !isEmpty(TtlNtriesPerBkTxCd) {
			TxsSummry.TtlNtriesPerBkTxCd = TtlNtriesPerBkTxCd
		}
	}
	if !isEmpty(TxsSummry) {
		Rpt.TxsSummry = &TxsSummry
	}
	// EntryDetails []model.Entry
	if !isEmpty(msg.model.EntryDetails) {
		var Ntry []*EndpointDetailsReport.ReportEntry101
		for _, entity := range msg.model.EntryDetails {
			report := ReportEntry101From(entity)
			Ntry = append(Ntry, &report)
		}
		Rpt.Ntry = Ntry
	}
	if !isEmpty(Rpt) {
		BkToCstmrAcctRpt.Rpt = Rpt
	}
	if !isEmpty(BkToCstmrAcctRpt) {
		msg.doc.BkToCstmrAcctRpt = BkToCstmrAcctRpt
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
