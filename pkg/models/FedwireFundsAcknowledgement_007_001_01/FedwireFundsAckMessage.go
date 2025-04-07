package FedwireFundsAcknowledgement_007_001_01

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	FedwireFundsAcknowledgement "github.com/moov-io/fedwire20022/gen/FedwireFundsAcknowledgement_admi_007_001_01"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Admi007 struct {
	//Specifies the identification the message.
	MessageId string
	//Date and time at which the message was created.
	CreatedDateTime time.Time
	//Unambiguous reference to a previous message having a business relevance with this message.
	RelationReference string
	//Name of the message which contained the given additional reference as its message reference.
	ReferenceName string
	//Gives the status of the request.
	RequestHandling model.RelatedStatusCode
}
type Admi007Message struct {
	model Admi007
	doc   FedwireFundsAcknowledgement.Document
}

func NewAdmi007Message() Admi007Message {
	return Admi007Message{
		model: Admi007{},
	}
}
func (msg *Admi007Message) CreateDocument() {
	msg.doc = FedwireFundsAcknowledgement.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:admi.007.001.01",
			Local: "Document",
		},
	}
	var RctAck FedwireFundsAcknowledgement.ReceiptAcknowledgementV01
	if msg.model.MessageId != "" {
		RctAck.MsgId.MsgId = FedwireFundsAcknowledgement.OMADFedwireFunds1(msg.model.MessageId)
	}
	if !isEmpty(msg.model.CreatedDateTime) {
		RctAck.MsgId.CreDtTm = fedwire.ISODateTime(msg.model.CreatedDateTime)
	}
	if msg.model.RelationReference != "" {
		RctAck.Rpt.RltdRef.Ref = FedwireFundsAcknowledgement.Max35Text(msg.model.RelationReference)
	}
	if msg.model.ReferenceName != "" {
		RctAck.Rpt.RltdRef.MsgNm = FedwireFundsAcknowledgement.MessageNameIdentificationFRS1(msg.model.ReferenceName)
	}
	if msg.model.RequestHandling != "" {
		RctAck.Rpt.ReqHdlg.StsCd = FedwireFundsAcknowledgement.Max4AlphaNumericTextFixed(msg.model.RequestHandling)
	}
	if !isEmpty(RctAck) {
		msg.doc.RctAck = RctAck
	}
}
func (msg *Admi007Message) GetXML() ([]byte, error) {
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
func (msg *Admi007Message) GetJson() ([]byte, error) {
	return json.MarshalIndent(msg.doc.RctAck, "", " ")
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:admi.007.001.01"`, "")

	return cleanXML
}
