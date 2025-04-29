package FedwireFundsBroadcast

import (
	"encoding/xml"
	"fmt"
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
	if msg.data.EventCode == "" {
		ParamNames = append(ParamNames, "EventCode")
	}
	if isEmpty(msg.data.EventParam) {
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
	var SysEvtNtfctn admi004.SystemEventNotificationV02
	var EvtInf admi004.Event21
	if msg.data.EventCode != "" {
		err := admi004.EventFedwireFunds1(msg.data.EventCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventCode",
				Message:   err.Error(),
			}
		}
		EvtInf.EvtCd = admi004.EventFedwireFunds1(msg.data.EventCode)
	}
	if !isEmpty(msg.data.EventParam) {
		err := msg.data.EventParam.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventParam",
				Message:   err.Error(),
			}
		}
		EvtInf.EvtParam = msg.data.EventParam.Date()
	}
	if msg.data.EventDescription != "" {
		err := admi004.Max1000Text(msg.data.EventDescription).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventDescription",
				Message:   err.Error(),
			}
		}
		EvtDesc := admi004.Max1000Text(msg.data.EventDescription)
		EvtInf.EvtDesc = &EvtDesc
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
		SysEvtNtfctn.EvtInf = EvtInf
	}
	if !isEmpty(SysEvtNtfctn) {
		msg.doc.SysEvtNtfctn = SysEvtNtfctn
	}
	return nil
}
