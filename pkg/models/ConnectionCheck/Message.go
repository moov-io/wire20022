package ConnectionCheck

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	admi004 "github.com/moov-io/fedwire20022/gen/ConnectionCheck_admi_004_001_02"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02"

type MessageModel struct {
	EventType  string
	EventParam string
	EventTime  time.Time
}
type Message struct {
	data MessageModel
	doc  admi004.Document
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
	if msg.data.EventType == "" {
		ParamNames = append(ParamNames, "EventType")
	}
	if msg.data.EventParam == "" {
		ParamNames = append(ParamNames, "EventParam")
	}
	if isEmpty(msg.data.EventTime) {
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
	msg.doc = admi004.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var EvtInf admi004.Event21
	if msg.data.EventType != "" {
		err := admi004.EventFedwireFunds1(msg.data.EventType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventType",
				Message:   err.Error(),
			}
		}
		EvtInf.EvtCd = admi004.EventFedwireFunds1(msg.data.EventType)
	}
	if msg.data.EventParam != "" {
		err := admi004.EndpointIdentifierFedwireFunds1(msg.data.EventParam).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EvntParam",
				Message:   err.Error(),
			}
		}
		EvtInf.EvtParam = admi004.EndpointIdentifierFedwireFunds1(msg.data.EventParam)
	}
	if !isEmpty(msg.data.EventTime) {
		err := fedwire.ISODateTime(msg.data.EventTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventTime",
				Message:   err.Error(),
			}
		}
		EvtInf.EvtTm = fedwire.ISODateTime(msg.data.EventTime)
	}
	if !isEmpty(EvtInf) {
		msg.doc.SysEvtNtfctn = admi004.SystemEventNotificationV02{
			EvtInf: EvtInf,
		}
	}
	return nil
}
