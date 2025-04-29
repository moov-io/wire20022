package FedwireFundsSystemResponse

import (
	"encoding/xml"
	"fmt"
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
	if msg.data.EventCode == "" {
		ParamNames = append(ParamNames, "EventCode")
	}
	if msg.data.EventParam == "" {
		ParamNames = append(ParamNames, "EventParam")
	}
	if msg.data.EventTime.IsZero() {
		ParamNames = append(ParamNames, "EventTime")
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
	msg.doc = admi011.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var SysEvtAck admi011.SystemEventAcknowledgementV01
	if msg.data.MessageId != "" {
		err := admi011.Max35Text(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		SysEvtAck.MsgId = admi011.Max35Text(msg.data.MessageId)
	}
	var AckDtls admi011.Event11
	if msg.data.EventCode != "" {
		err := admi011.EventFedwireFunds1(msg.data.EventCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventCode",
				Message:   err.Error(),
			}
		}
		AckDtls.EvtCd = admi011.EventFedwireFunds1(msg.data.EventCode)
	}
	if msg.data.EventParam != "" {
		err := admi011.EndpointIdentifierFedwireFunds1(msg.data.EventParam).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventParam",
				Message:   err.Error(),
			}
		}
		AckDtls.EvtParam = admi011.EndpointIdentifierFedwireFunds1(msg.data.EventParam)
	}
	if !isEmpty(msg.data.EventTime) {
		err := fedwire.ISODateTime(msg.data.EventTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventParam",
				Message:   err.Error(),
			}
		}
		AckDtls.EvtTm = fedwire.ISODateTime(msg.data.EventTime)
	}
	if !isEmpty(AckDtls) {
		SysEvtAck.AckDtls = AckDtls
	}
	if !isEmpty(SysEvtAck) {
		msg.doc.SysEvtAck = SysEvtAck
	}
	return nil
}
