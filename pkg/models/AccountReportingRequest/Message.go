package AccountReportingRequest

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	camt060 "github.com/moov-io/fedwire20022/gen/AccountReportingRequest_camt_060_001_05"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05"

type MessageModel struct {
	//MessageId (Message Identification) is a unique identifier assigned to an entire message.
	MessageId string
	//CreatedDateTime represents the timestamp when a message, instruction, or transaction was created
	//ISO 8601 format
	CreatedDateTime time.Time
	//Unique identification, as assigned by the account owner, to unambiguously identify the account reporting request.
	ReportRequestId model.CAMTReportType
	//Specifies the type of the requested reporting message.
	RequestedMsgNameId string
	//account or entity identifier does not conform to any predefined ISO 20022 standard
	AccountOtherId string
	//AccountProperty defines the properties of a financial account.
	AccountProperty AccountTypeFRS
	//Proprietary account type used when no ISO 20022 standard code applies
	AcountTypeProprietary string
	// It is defined as a Camt060Agent type which encapsulates the choice of different party identification options for the account owner.
	AccountOwnerAgent Camt060Agent
	//"From-To" sequence within the ISO 20022 camt.060.001.05 message.
	FromToSeuence model.SequenceRange
}

type Message struct {
	data MessageModel
	doc  camt060.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}

func (msg *Message) CreateDocument() {
	msg.doc = camt060.Document{
		XMLName: xml.Name{
			Local: "Document",
		},
		AcctRptgReq: camt060.AccountReportingRequestV05{
			GrpHdr: camt060.GroupHeader771{
				MsgId:   camt060.Max35Text(msg.data.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.data.CreatedDateTime),
			},
		},
	}
	var RptgReq camt060.ReportingRequest51
	if msg.data.ReportRequestId != "" {
		RptgReq.Id = camt060.AccountReportingFedwireFunds1(msg.data.ReportRequestId)
	}
	if msg.data.RequestedMsgNameId != "" {
		RptgReq.ReqdMsgNmId = camt060.MessageNameIdentificationFRS1(msg.data.RequestedMsgNameId)
	}
	if msg.data.AccountOtherId != "" {
		id_othr := camt060.GenericAccountIdentification11{
			Id: camt060.RoutingNumberFRS1(msg.data.AccountOtherId),
		}

		_account := camt060.CashAccount381{
			Id: camt060.AccountIdentification4Choice1{
				Othr: &id_othr,
			},
		}
		if msg.data.AccountProperty != "" {
			_Prtry := camt060.AccountTypeFRS1(msg.data.AccountProperty)
			_account.Tp = camt060.CashAccountType2Choice1{
				Prtry: &_Prtry,
			}
		}
		RptgReq.Acct = &_account
	}
	if !isEmpty(msg.data.AccountOwnerAgent.agent) {
		_AcctOwnr := Party40Choice1From(msg.data.AccountOwnerAgent.agent)
		if !isEmpty(_AcctOwnr) {
			RptgReq.AcctOwnr = _AcctOwnr
		}
		if msg.data.AccountOwnerAgent.OtherId != "" {
			_Other := camt060.GenericFinancialIdentification11{
				Id: camt060.EndpointIdentifierFedwireFunds1(msg.data.AccountOwnerAgent.OtherId),
			}
			RptgReq.AcctOwnr.Agt.FinInstnId.Othr = &_Other
		}
	}
	if !isEmpty(msg.data.FromToSeuence) {
		_FrToSeq := camt060.SequenceRange11{
			FrSeq: camt060.XSequenceNumberFedwireFunds1(msg.data.FromToSeuence.FromSeq),
			ToSeq: camt060.XSequenceNumberFedwireFunds1(msg.data.FromToSeuence.ToSeq),
		}
		_RptgSeq := camt060.SequenceRange1Choice1{
			FrToSeq: &_FrToSeq,
		}
		RptgReq.RptgSeq = &_RptgSeq
	}
	if !isEmpty(RptgReq) {
		msg.doc.AcctRptgReq.RptgReq = RptgReq
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
