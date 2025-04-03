package ConnectionCheck_004_001_02

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
	"time"

	ConnectionCheck "github.com/moov-io/fedwire20022/gen/ConnectionCheck_admi_004_001_02"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

type Admi004 struct {
	EventType string
	EvntParam string
	EventTime time.Time
}
type Admi004Message struct {
	model Admi004
	doc   ConnectionCheck.Document
}

func NewAdmi004Message() Admi004Message {
	return Admi004Message{
		model: Admi004{},
	}
}
func (msg *Admi004Message) CreateDocument() {
	msg.doc = ConnectionCheck.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02",
			Local: "Document",
		},
	}
	var EvtInf ConnectionCheck.Event21
	if msg.model.EventType != "" {
		EvtInf.EvtCd = ConnectionCheck.EventFedwireFunds1(msg.model.EventType)
	}
	if msg.model.EvntParam != "" {
		EvtInf.EvtParam = ConnectionCheck.EndpointIdentifierFedwireFunds1(msg.model.EvntParam)
	}
	if !isEmpty(msg.model.EventTime) {
		EvtInf.EvtTm = fedwire.ISODateTime(msg.model.EventTime)
	}
	if !isEmpty(EvtInf) {
		msg.doc.SysEvtNtfctn = ConnectionCheck.SystemEventNotificationV02{
			EvtInf: EvtInf,
		}
	}
}
func (msg *Admi004Message) GetXML() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(msg.doc, "", "\t")
	if err != nil {
		return nil, err
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
	return []byte(xmlString), nil
	// return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *Admi004Message) GetJson() ([]byte, error) {
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
