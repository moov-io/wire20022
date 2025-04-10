package FedwireFundsSystemResponse

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	admi011 "github.com/moov-io/fedwire20022/gen/FedwireFundsSystemResponse_admi_011_001_01"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:admi.011.001.01"

type MessageModel struct {
	//Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageId string
	//Proprietary code used to specify an event that occurred in a system.
	EventCode model.FundEventType
	//Describes the parameters of an event which occurred in a system.
	EventParam string
	//Date and time at which the event occurred.
	EventTime time.Time
}
type Message struct {
	data MessageModel
	doc  admi011.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = admi011.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var SysEvtAck admi011.SystemEventAcknowledgementV01
	if msg.data.MessageId != "" {
		SysEvtAck.MsgId = admi011.Max35Text(msg.data.MessageId)
	}
	var AckDtls admi011.Event11
	if msg.data.EventCode != "" {
		AckDtls.EvtCd = admi011.EventFedwireFunds1(msg.data.EventCode)
	}
	if msg.data.EventParam != "" {
		AckDtls.EvtParam = admi011.EndpointIdentifierFedwireFunds1(msg.data.EventParam)
	}
	if !isEmpty(msg.data.EventTime) {
		AckDtls.EvtTm = fedwire.ISODateTime(msg.data.EventTime)
	}
	if !isEmpty(AckDtls) {
		SysEvtAck.AckDtls = AckDtls
	}
	if !isEmpty(SysEvtAck) {
		msg.doc.SysEvtAck = SysEvtAck
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
