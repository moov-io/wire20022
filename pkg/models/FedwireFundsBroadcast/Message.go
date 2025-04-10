package FedwireFundsBroadcast

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	admi004 "github.com/moov-io/fedwire20022/gen/FedwireFundsBroadcast_admi_004_001_02"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02"

type MessageModel struct {
	//Proprietary code used to specify an event that occurred in a system.
	EventCode model.FundEventType
	//Describes the parameters of an event which occurred in a system.
	EventParam model.Date
	//Free text used to describe an event which occurred in a system.
	EventDescription string
	//Date and time at which the event occurred.
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
			Space: XMLINS,
			Local: "Document",
		},
	}
	var SysEvtNtfctn admi004.SystemEventNotificationV02
	var EvtInf admi004.Event21
	if msg.data.EventCode != "" {
		EvtInf.EvtCd = admi004.EventFedwireFunds1(msg.data.EventCode)
	}
	if !isEmpty(msg.data.EventParam) {
		EvtInf.EvtParam = msg.data.EventParam.ToIosDate()
	}
	if msg.data.EventDescription != "" {
		EvtDesc := admi004.Max1000Text(msg.data.EventDescription)
		EvtInf.EvtDesc = &EvtDesc
	}
	if !isEmpty(msg.data.EventTime) {
		EvtInf.EvtTm = fedwire.ISODateTime(msg.data.EventTime)
	}
	if !isEmpty(EvtInf) {
		SysEvtNtfctn.EvtInf = EvtInf
	}
	if !isEmpty(SysEvtNtfctn) {
		msg.doc.SysEvtNtfctn = SysEvtNtfctn
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
