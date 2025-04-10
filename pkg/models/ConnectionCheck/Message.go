package ConnectionCheck

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
	"time"

	admi004 "github.com/moov-io/fedwire20022/gen/ConnectionCheck_admi_004_001_02"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

type MessageModel struct {
	EventType string
	EvntParam string
	EventTime time.Time
}
type Message struct {
	data MessageModel
	doc  admi004.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = admi004.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02",
			Local: "Document",
		},
	}
	var EvtInf admi004.Event21
	if msg.data.EventType != "" {
		EvtInf.EvtCd = admi004.EventFedwireFunds1(msg.data.EventType)
	}
	if msg.data.EvntParam != "" {
		EvtInf.EvtParam = admi004.EndpointIdentifierFedwireFunds1(msg.data.EvntParam)
	}
	if !isEmpty(msg.data.EventTime) {
		EvtInf.EvtTm = fedwire.ISODateTime(msg.data.EventTime)
	}
	if !isEmpty(EvtInf) {
		msg.doc.SysEvtNtfctn = admi004.SystemEventNotificationV02{
			EvtInf: EvtInf,
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
	return json.MarshalIndent(msg.doc.SysEvtNtfctn, "", " ")
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:admi.004.001.02"`, "")

	return cleanXML
}
