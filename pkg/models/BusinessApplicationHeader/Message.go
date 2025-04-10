package BusinessApplicationHeader

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	head001 "github.com/moov-io/fedwire20022/gen/BusinessApplicationHeader_head_001_001_03"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:head.001.001.03"

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
			Space: XMLINS,
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="`+XMLINS+`"`, "")

	return cleanXML
}
