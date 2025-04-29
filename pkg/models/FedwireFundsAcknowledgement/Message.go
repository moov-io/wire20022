package FedwireFundsAcknowledgement

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	admi007 "github.com/moov-io/fedwire20022/gen/FedwireFundsAcknowledgement_admi_007_001_01"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:admi.007.001.01"

type MessageModel struct {
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
type Message struct {
	data MessageModel
	doc  admi007.Document
}

/*
NewMessage creates a new Message instance with optional XML initialization.

Parameters:
  - filepath: File path to XML (optional)
    If provided, loads and parses XML from specified path

Returns:
  - Message: Initialized message structure
  - error: File read or XML parsing errors (if XML path provided)

Behavior:
  - Without arguments: Returns empty Message with default MessageModel
  - With XML path: Loads file, parses XML into message.doc
*/
func NewMessage(filepath string) (Message, error) {
	msg := Message{data: MessageModel{}} // Initialize with zero value

	if filepath == "" {
		return msg, nil // Return early for empty filepath
	}

	// Read and validate file
	data, err := model.ReadXMLFile(filepath)
	if err != nil {
		return msg, fmt.Errorf("file read error: %w", err)
	}

	// Handle empty XML data
	if len(data) == 0 {
		return msg, fmt.Errorf("empty XML file: %s", filepath)
	}

	// Parse XML with structural validation
	if err := xml.Unmarshal(data, &msg.doc); err != nil {
		return msg, fmt.Errorf("XML parse error: %w", err)
	}

	return msg, nil
}

func (msg *Message) ValidateRequiredFields() *model.ValidateError {
	// Initialize the RequireError object
	var ParamNames []string

	// Check required fields and append missing ones to ParamNames
	if msg.data.MessageId == "" {
		ParamNames = append(ParamNames, "MessageId")
	}
	if msg.data.CreatedDateTime.IsZero() {
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if msg.data.RelationReference == "" {
		ParamNames = append(ParamNames, "RelationReference")
	}
	if msg.data.ReferenceName == "" {
		ParamNames = append(ParamNames, "ReferenceName")
	}
	if msg.data.RequestHandling == "" {
		ParamNames = append(ParamNames, "RequestHandling")
	}
	// Return nil if no required fields are missing
	if len(ParamNames) == 0 {
		return nil
	}
	return &model.ValidateError{
		ParamName: "RequiredFields",
		Message:   strings.Join(ParamNames, ", "),
	}
}

func (msg *Message) CreateDocument() *model.ValidateError {
	requireErr := msg.ValidateRequiredFields()
	if requireErr != nil {
		return requireErr
	}
	msg.doc = admi007.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var RctAck admi007.ReceiptAcknowledgementV01
	if msg.data.MessageId != "" {
		err := admi007.OMADFedwireFunds1(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		RctAck.MsgId.MsgId = admi007.OMADFedwireFunds1(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		RctAck.MsgId.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
	if msg.data.RelationReference != "" {
		err := admi007.Max35Text(msg.data.RelationReference).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RelationReference",
				Message:   err.Error(),
			}
		}
		RctAck.Rpt.RltdRef.Ref = admi007.Max35Text(msg.data.RelationReference)
	}
	if msg.data.ReferenceName != "" {
		err := admi007.MessageNameIdentificationFRS1(msg.data.ReferenceName).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReferenceName",
				Message:   err.Error(),
			}
		}
		RctAck.Rpt.RltdRef.MsgNm = admi007.MessageNameIdentificationFRS1(msg.data.ReferenceName)
	}
	if msg.data.RequestHandling != "" {
		err := admi007.Max4AlphaNumericTextFixed(msg.data.RequestHandling).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RequestHandling",
				Message:   err.Error(),
			}
		}
		RctAck.Rpt.ReqHdlg.StsCd = admi007.Max4AlphaNumericTextFixed(msg.data.RequestHandling)
	}
	if !isEmpty(RctAck) {
		msg.doc.RctAck = RctAck
	}
	return nil
}
