package EndpointDetailsReport

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	camt052 "github.com/moov-io/fedwire20022/gen/EndpointDetailsReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"

type MessageModel struct {
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
			Space: XMLINS,
			Local: "Document",
		},
	}
	var BkToCstmrAcctRpt camt052.BankToCustomerAccountReportV08
	var GrpHdr camt052.GroupHeader811
	// MessageId string
	if msg.data.MessageId != "" {
		GrpHdr.MsgId = camt052.AccountReportingFedwireFunds1(msg.data.MessageId)
	}
	// CreationDateTime time.Time
	if !isEmpty(msg.data.CreationDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreationDateTime)
	}
	// MessagePagination model.MessagePagenation
	if !isEmpty(msg.data.MessagePagination) {
		GrpHdr.MsgPgntn = camt052.Pagination1{
			PgNb:      camt052.Max5NumericText(msg.data.MessagePagination.PageNumber),
			LastPgInd: camt052.YesNoIndicator(msg.data.MessagePagination.LastPageIndicator),
		}
	}
	var OrgnlBizQry camt052.OriginalBusinessQuery11
	// BussinessQueryMsgId string
	if msg.data.BussinessQueryMsgId != "" {
		OrgnlBizQry.MsgId = camt052.Max35Text(msg.data.BussinessQueryMsgId)
	}
	// BussinessQueryMsgNameId string
	if msg.data.BussinessQueryMsgNameId != "" {
		OrgnlBizQry.MsgNmId = camt052.MessageNameIdentificationFRS1(msg.data.BussinessQueryMsgNameId)
	}
	// BussinessQueryCreateDatetime time.Time
	if !isEmpty(msg.data.BussinessQueryCreateDatetime) {
		OrgnlBizQry.CreDtTm = fedwire.ISODateTime(msg.data.BussinessQueryCreateDatetime)
	}
	if !isEmpty(OrgnlBizQry) {
		GrpHdr.OrgnlBizQry = OrgnlBizQry
	}
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}
	var Rpt camt052.AccountReport251
	// ReportId ReportDayType
	if msg.data.ReportId != "" {
		Rpt.Id = camt052.ReportTimingFRS1(msg.data.ReportId)
	}
	// ReportingSequence model.SequenceRange
	if !isEmpty(msg.data.ReportingSequence) {
		FrToSeq := camt052.SequenceRange11{
			FrSeq: camt052.XSequenceNumberFedwireFunds1(msg.data.ReportingSequence.FromSeq),
			ToSeq: camt052.XSequenceNumberFedwireFunds1(msg.data.ReportingSequence.ToSeq),
		}
		Rpt.RptgSeq = camt052.SequenceRange1Choice1{
			FrToSeq: &FrToSeq,
		}
	}
	// ReportCreateDateTime time.Time
	if !isEmpty(msg.data.ReportCreateDateTime) {
		CreDtTm := fedwire.ISODateTime(msg.data.ReportCreateDateTime)
		Rpt.CreDtTm = &CreDtTm
	}
	// AccountOtherId string
	if msg.data.AccountOtherId != "" {
		Othr := camt052.GenericAccountIdentification11{
			Id: camt052.EndpointIdentifierFedwireFunds1(msg.data.AccountOtherId),
		}
		Rpt.Acct = camt052.CashAccount391{
			Id: camt052.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	// TotalDebitEntries model.NumberAndSumOfTransactions
	var TxsSummry camt052.TotalTransactions61
	if !isEmpty(msg.data.TotalDebitEntries) {
		TtlDbtNtries := camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalDebitEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalDebitEntries.Sum),
		}
		TxsSummry.TtlDbtNtries = &TtlDbtNtries
	}
	// TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	if !isEmpty(msg.data.TotalEntriesPerTransactionCode) {
		var TtlNtriesPerBkTxCd []camt052.TotalsPerBankTransactionCode51
		for _, entity := range msg.data.TotalEntriesPerTransactionCode {
			bankTrans := camt052.TotalsPerBankTransactionCode51{
				NbOfNtries: camt052.Max15NumericText(entity.NumberOfEntries),
				BkTxCd: camt052.BankTransactionCodeStructure41{
					Prtry: camt052.ProprietaryBankTransactionCodeStructure11{
						Cd: camt052.BankTransactionCodeFedwireFunds11(entity.Status),
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
	if !isEmpty(msg.data.EntryDetails) {
		var Ntry []*camt052.ReportEntry101
		for _, entity := range msg.data.EntryDetails {
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
func WriteXMLTo(filePath string, xml []byte) error {
	os.Mkdir("generated", 0755)
	xmlFileName := filepath.Join("generated", filePath)

	xmlString := string(xml)
	xmlString = removeExtraXMLNS(xmlString)
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

	return os.WriteFile(xmlFileName, []byte(xmlString), 0644)
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="`+XMLINS+`"`, "")

	return cleanXML
}
