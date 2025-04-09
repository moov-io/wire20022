package FedwireFundsBroadcast_004_001_02

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	FedwireFundsBroadcast "github.com/moov-io/fedwire20022/gen/FedwireFundsBroadcast_admi_004_001_02"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Admi004 struct {
	//Proprietary code used to specify an event that occurred in a system.
	EventCode model.FundEventType
	//Describes the parameters of an event which occurred in a system.
	EventParam civil.Date
	//Free text used to describe an event which occurred in a system.
	EventDescription string
	//Date and time at which the event occurred.
	EventTime time.Time
}
type Admi004Message struct {
	model Admi004
	doc   FedwireFundsBroadcast.Document
}

func NewAdmi004Message() Admi004Message {
	return Admi004Message{
		model: Admi004{},
	}
}
func (msg *Admi004Message) CreateDocument() {
	msg.doc = FedwireFundsBroadcast.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02",
			Local: "Document",
		},
	}
	var SysEvtNtfctn FedwireFundsBroadcast.SystemEventNotificationV02
	var EvtInf FedwireFundsBroadcast.Event21
	if msg.model.EventCode != "" {
		EvtInf.EvtCd = FedwireFundsBroadcast.EventFedwireFunds1(msg.model.EventCode)
	}
	if !isEmpty(msg.model.EventParam) {
		EvtInf.EvtParam = fedwire.ISODate(msg.model.EventParam)
	}
	if msg.model.EventDescription != "" {
		EvtDesc := FedwireFundsBroadcast.Max1000Text(msg.model.EventDescription)
		EvtInf.EvtDesc = &EvtDesc
	}
	if !isEmpty(msg.model.EventTime) {
		EvtInf.EvtTm = fedwire.ISODateTime(msg.model.EventTime)
	}
	if !isEmpty(EvtInf) {
		SysEvtNtfctn.EvtInf = EvtInf
	}
	if !isEmpty(SysEvtNtfctn) {
		msg.doc.SysEvtNtfctn = SysEvtNtfctn
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
