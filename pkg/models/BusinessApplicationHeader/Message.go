package BusinessApplicationHeader

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
	"time"

	head001 "github.com/moov-io/fedwire20022/gen/BusinessApplicationHeader_head_001_001_03"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

type MessageModel struct {
	MessageSenderId   string
	MessageReceiverId string
	//BizMsgIdr (Business Message Identifier) is a unique identifier assigned to an ISO 20022 message to distinguish it from other messages.
	BusinessMessageId string
	//MsgDefIdr (Message Definition Identifier) and BizSvc (Business Service) are part of the Business Application Header (BAH), which helps identify and process financial messages.
	MessageDefinitionId string
	//BizSvc specifies a business service or process related to the message.
	BusinessService string
	//<MktPrctc> (Market Practice) is used to specify market-specific rules and guidelines that apply to the message.
	MarketInfo MarketPractice

	CreateDatetime time.Time
	//BizPrcgDt stands for Business Processing Date. It represents the date when a financial transaction or message is processed by a financial institution.
	BusinessProcessingDate time.Time
	//It refers to a related Business Application Header (BAH) of type BusinessApplicationHeader71
	Relations BusinessApplicationHeader
}
type Message struct {
	data MessageModel
	doc  head001.AppHdr
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}

func (msg *Message) CreateDocument() {
	msg.doc = head001.AppHdr{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:head.001.001.03",
			Local: "AppHdr",
		},
	}
	if msg.data.MessageSenderId != "" {
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(msg.data.MessageSenderId),
				},
			},
		}
		msg.doc.Fr = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if msg.data.MessageReceiverId != "" {
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(msg.data.MessageReceiverId),
				},
			},
		}
		msg.doc.To = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if msg.data.BusinessMessageId != "" {
		msg.doc.BizMsgIdr = head001.Max35Text(msg.data.BusinessMessageId)
	}
	if msg.data.MessageDefinitionId != "" {
		msg.doc.MsgDefIdr = head001.MessageNameIdentificationFRS1(msg.data.MessageDefinitionId)
	}
	if msg.data.BusinessService != "" {
		msg.doc.BizSvc = head001.BusinessServiceFedwireFunds1(msg.data.BusinessService)
	}
	if !isEmpty(msg.data.MarketInfo) {
		MktPrctc := ImplementationSpecification11From(msg.data.MarketInfo)
		if !isEmpty(MktPrctc) {
			msg.doc.MktPrctc = MktPrctc
		}
	}
	if !isEmpty(msg.data.CreateDatetime) {
		msg.doc.CreDt = fedwire.ISODateTime(msg.data.CreateDatetime)
	}
	if !isEmpty(msg.data.BusinessProcessingDate) {
		BizPrcgDt := fedwire.ISODateTime(msg.data.BusinessProcessingDate)
		msg.doc.BizPrcgDt = &BizPrcgDt
	}
	if !isEmpty(msg.data.Relations) {
		Rltd := BusinessApplicationHeader71From(msg.data.Relations)
		if !isEmpty(Rltd) {
			msg.doc.Rltd = &Rltd
		}
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
	return json.MarshalIndent(msg.doc, "", " ")
}

func removeExtraXMLNS(xmlStr string) string {
	// Find the first occurrence of <Document ...> (keep this)
	docStart := strings.Index(xmlStr, "<AppHdr")
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:head.001.001.03"`, "")

	return cleanXML
}
