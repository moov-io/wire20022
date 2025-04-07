package EndpointGapReport_052_001_08

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	EndpointGapReport "github.com/moov-io/fedwire20022/gen/EndpointGapReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type CamtGap052 struct {
	//Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId model.CAMTReportType
	//Date and time at which the message was created.
	CreatedDateTime time.Time
	//Provides details on the page number of the message.
	MessagePagination model.MessagePagenation
	//Report id on a cash account.
	ReportId GapType
	//This is the Fedwire Funds Service funds-transfer business day when the gap was identified.
	ReportCreateDateTime time.Time
	//Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId string
	//For the Fedwire Funds Service, this provides the missing sequence numbers.
	AdditionalReportInfo string
}

type CamtGap0522Message struct {
	model CamtGap052
	doc   EndpointGapReport.Document
}

func NewCamtGap0522Message() CamtGap0522Message {
	return CamtGap0522Message{
		model: CamtGap052{},
	}
}

func (msg *CamtGap0522Message) CreateDocument() {
	msg.doc = EndpointGapReport.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08",
			Local: "Document",
		},
	}
	var BkToCstmrAcctRpt EndpointGapReport.BankToCustomerAccountReportV08
	var GrpHdr EndpointGapReport.GroupHeader811
	if msg.model.MessageId != "" {
		GrpHdr.MsgId = EndpointGapReport.AccountReportingFedwireFunds1(msg.model.MessageId)
	}
	if !isEmpty(msg.model.CreatedDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.model.CreatedDateTime)
	}
	if !isEmpty(msg.model.MessagePagination) {
		GrpHdr.MsgPgntn = EndpointGapReport.Pagination1{
			PgNb:      EndpointGapReport.Max5NumericText(msg.model.MessagePagination.PageNumber),
			LastPgInd: EndpointGapReport.YesNoIndicator(msg.model.MessagePagination.LastPageIndicator),
		}
	}
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}
	var Rpt []EndpointGapReport.AccountReport251
	var report_data EndpointGapReport.AccountReport251
	if msg.model.ReportId != "" {
		report_data.Id = EndpointGapReport.GapTypeFedwireFunds1(msg.model.ReportId)
	}
	if !isEmpty(msg.model.ReportCreateDateTime) {
		report_data.CreDtTm = fedwire.ISODateTime(msg.model.ReportCreateDateTime)
	}
	if msg.model.AccountOtherId != "" {
		Othr := EndpointGapReport.GenericAccountIdentification11{
			Id: EndpointGapReport.EndpointIdentifierFedwireFunds1(msg.model.AccountOtherId),
		}
		report_data.Acct = EndpointGapReport.CashAccount391{
			Id: EndpointGapReport.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	if msg.model.AdditionalReportInfo != "" {
		report_data.AddtlRptInf = EndpointGapReport.Max500Text(msg.model.AdditionalReportInfo)
	}
	if !isEmpty(report_data) {
		Rpt = append(Rpt, report_data)
	}
	if !isEmpty(Rpt) {
		BkToCstmrAcctRpt.Rpt = Rpt
	}
	if !isEmpty(BkToCstmrAcctRpt) {
		msg.doc.BkToCstmrAcctRpt = BkToCstmrAcctRpt
	}
}
func (msg *CamtGap0522Message) GetXML() ([]byte, error) {
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
func (msg *CamtGap0522Message) GetJson() ([]byte, error) {
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
