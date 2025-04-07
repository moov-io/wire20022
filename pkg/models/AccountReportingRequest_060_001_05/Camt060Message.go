package AccountReportingRequest_060_001_05

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
	"time"

	AccountReportingRequest "github.com/moov-io/fedwire20022/gen/AccountReportingRequest_camt_060_001_05"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Camt060 struct {
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

type Camt060Message struct {
	model Camt060
	doc   AccountReportingRequest.Document
}

func NewCamt060MessageMessage() Camt060Message {
	return Camt060Message{
		model: Camt060{},
	}
}

func (msg *Camt060Message) CreateDocument() {
	msg.doc = AccountReportingRequest.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05",
			Local: "Document",
		},
		AcctRptgReq: AccountReportingRequest.AccountReportingRequestV05{
			GrpHdr: AccountReportingRequest.GroupHeader771{
				MsgId:   AccountReportingRequest.Max35Text(msg.model.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.model.CreatedDateTime),
			},
		},
	}
	var RptgReq AccountReportingRequest.ReportingRequest51
	if msg.model.ReportRequestId != "" {
		RptgReq.Id = AccountReportingRequest.AccountReportingFedwireFunds1(msg.model.ReportRequestId)
	}
	if msg.model.RequestedMsgNameId != "" {
		RptgReq.ReqdMsgNmId = AccountReportingRequest.MessageNameIdentificationFRS1(msg.model.RequestedMsgNameId)
	}
	if msg.model.AccountOtherId != "" {
		id_othr := AccountReportingRequest.GenericAccountIdentification11{
			Id: AccountReportingRequest.RoutingNumberFRS1(msg.model.AccountOtherId),
		}

		_account := AccountReportingRequest.CashAccount381{
			Id: AccountReportingRequest.AccountIdentification4Choice1{
				Othr: &id_othr,
			},
		}
		if msg.model.AccountProperty != "" {
			_Prtry := AccountReportingRequest.AccountTypeFRS1(msg.model.AccountProperty)
			_account.Tp = AccountReportingRequest.CashAccountType2Choice1{
				Prtry: &_Prtry,
			}
		}
		RptgReq.Acct = &_account
	}
	if !isEmpty(msg.model.AccountOwnerAgent.agent) {
		_AcctOwnr := Party40Choice1From(msg.model.AccountOwnerAgent.agent)
		if !isEmpty(_AcctOwnr) {
			RptgReq.AcctOwnr = _AcctOwnr
		}
		if msg.model.AccountOwnerAgent.OtherId != "" {
			_Other := AccountReportingRequest.GenericFinancialIdentification11{
				Id: AccountReportingRequest.EndpointIdentifierFedwireFunds1(msg.model.AccountOwnerAgent.OtherId),
			}
			RptgReq.AcctOwnr.Agt.FinInstnId.Othr = &_Other
		}
	}
	if !isEmpty(msg.model.FromToSeuence) {
		_FrToSeq := AccountReportingRequest.SequenceRange11{
			FrSeq: AccountReportingRequest.XSequenceNumberFedwireFunds1(msg.model.FromToSeuence.FromSeq),
			ToSeq: AccountReportingRequest.XSequenceNumberFedwireFunds1(msg.model.FromToSeuence.ToSeq),
		}
		_RptgSeq := AccountReportingRequest.SequenceRange1Choice1{
			FrToSeq: &_FrToSeq,
		}
		RptgReq.RptgSeq = &_RptgSeq
	}
	if !isEmpty(RptgReq) {
		msg.doc.AcctRptgReq.RptgReq = RptgReq
	}
}
func (msg *Camt060Message) GetXML() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(msg.doc, "", "\t")
	if err != nil {
		return nil, err
	}

	// Convert byte slice to string for manipulation
	xmlString := string(xmlData)

	// Keep the xmlns only in the <Document> tag, remove from others
	xmlString = removeExtraXMLNS(xmlString)
	// Regex to find <FromSeq> and <ToSeq> values
	re := regexp.MustCompile(`<(FrSeq|ToSeq)>(\d+)</(FrSeq|ToSeq)>`)

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
func (msg *Camt060Message) GetJson() ([]byte, error) {
	return json.MarshalIndent(msg.doc.AcctRptgReq, "", " ")
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:camt.060.001.05"`, "")

	return cleanXML
}
