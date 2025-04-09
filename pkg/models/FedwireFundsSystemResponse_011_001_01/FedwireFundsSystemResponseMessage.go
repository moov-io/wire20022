package FedwireFundsSystemResponse_011_001_01

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	FedwireFundsSystemResponse "github.com/moov-io/fedwire20022/gen/FedwireFundsSystemResponse_admi_011_001_01"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Admi011 struct {
	//Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageId string
	//Proprietary code used to specify an event that occurred in a system.
	EventCode model.FundEventType
	//Describes the parameters of an event which occurred in a system.
	EventParam string
	//Date and time at which the event occurred.
	EventTime time.Time
}
type Admi011Message struct {
	model Admi011
	doc   FedwireFundsSystemResponse.Document
}

func NewAdmi011Message() Admi011Message {
	return Admi011Message{
		model: Admi011{},
	}
}
func (msg *Admi011Message) CreateDocument() {
	msg.doc = FedwireFundsSystemResponse.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:admi.011.001.01",
			Local: "Document",
		},
	}
	var SysEvtAck FedwireFundsSystemResponse.SystemEventAcknowledgementV01
	if msg.model.MessageId != "" {
		SysEvtAck.MsgId = FedwireFundsSystemResponse.Max35Text(msg.model.MessageId)
	}
	var AckDtls FedwireFundsSystemResponse.Event11
	if msg.model.EventCode != "" {
		AckDtls.EvtCd = FedwireFundsSystemResponse.EventFedwireFunds1(msg.model.EventCode)
	}
	if msg.model.EventParam != "" {
		AckDtls.EvtParam = FedwireFundsSystemResponse.EndpointIdentifierFedwireFunds1(msg.model.EventParam)
	}
	if !isEmpty(msg.model.EventTime) {
		AckDtls.EvtTm = fedwire.ISODateTime(msg.model.EventTime)
	}
	if !isEmpty(AckDtls) {
		SysEvtAck.AckDtls = AckDtls
	}
	if !isEmpty(SysEvtAck) {
		msg.doc.SysEvtAck = SysEvtAck
	}
}
func (msg *Admi011Message) GetXML() ([]byte, error) {
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
func (msg *Admi011Message) GetJson() ([]byte, error) {
	return json.MarshalIndent(msg.doc.SysEvtAck, "", " ")
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:admi.011.001.01"`, "")

	return cleanXML
}
