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
	Data   MessageModel
	Doc    admi011.Document
	Helper MessageHelper
}

func (msg *Message) GetDataModel() interface{} {
	return &msg.Data
}
func (msg *Message) GetDocument() interface{} {
	return &msg.Doc
}
func (msg *Message) GetHelper() interface{} {
	return &msg.Helper
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
  - With XML path: Loads file, parses XML into message.Doc
*/
func NewMessage(filepath string) (*Message, error) {
	msg := Message{Data: MessageModel{}} // Initialize with zero value
	msg.Helper = BuildMessageHelper()

	if filepath == "" {
		return &msg, nil // Return early for empty filepath
	}

	// Read and validate file
	data, err := model.ReadXMLFile(filepath)
	if err != nil {
		return &msg, fmt.Errorf("file read error: %w", err)
	}

	// Handle empty XML data
	if len(data) == 0 {
		return &msg, fmt.Errorf("empty XML file: %s", filepath)
	}

	// Parse XML with structural validation
	if err := xml.Unmarshal(data, &msg.Doc); err != nil {
		return &msg, fmt.Errorf("XML parse error: %w", err)
	}

	return &msg, nil
}

func (msg *Message) ValidateRequiredFields() *model.ValidateError {
	// Initialize the RequireError object
	var ParamNames []string

	// Check required fields and append missing ones to ParamNames
	if msg.Data.MessageId == "" {
		ParamNames = append(ParamNames, "MessageId")
	}
	if msg.Data.EventCode == "" {
		ParamNames = append(ParamNames, "EventCode")
	}
	if msg.Data.EventParam == "" {
		ParamNames = append(ParamNames, "EventParam")
	}
	if msg.Data.EventTime.IsZero() {
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
	msg.Doc = admi011.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var SysEvtAck admi011.SystemEventAcknowledgementV01
	if msg.Data.MessageId != "" {
		err := admi011.Max35Text(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		SysEvtAck.MsgId = admi011.Max35Text(msg.Data.MessageId)
	}
	var AckDtls admi011.Event11
	if msg.Data.EventCode != "" {
		err := admi011.EventFedwireFunds1(msg.Data.EventCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventCode",
				Message:   err.Error(),
			}
		}
		AckDtls.EvtCd = admi011.EventFedwireFunds1(msg.Data.EventCode)
	}
	if msg.Data.EventParam != "" {
		err := admi011.EndpointIdentifierFedwireFunds1(msg.Data.EventParam).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventParam",
				Message:   err.Error(),
			}
		}
		AckDtls.EvtParam = admi011.EndpointIdentifierFedwireFunds1(msg.Data.EventParam)
	}
	if !isEmpty(msg.Data.EventTime) {
		err := fedwire.ISODateTime(msg.Data.EventTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventParam",
				Message:   err.Error(),
			}
		}
		AckDtls.EvtTm = fedwire.ISODateTime(msg.Data.EventTime)
	}
	if !isEmpty(AckDtls) {
		SysEvtAck.AckDtls = AckDtls
	}
	if !isEmpty(SysEvtAck) {
		msg.Doc.SysEvtAck = SysEvtAck
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.SysEvtAck) {
		if !isEmpty(msg.Doc.SysEvtAck.MsgId) {
			msg.Data.MessageId = string(msg.Doc.SysEvtAck.MsgId)
		}
		if !isEmpty(msg.Doc.SysEvtAck.AckDtls) {
			if !isEmpty(msg.Doc.SysEvtAck.AckDtls.EvtCd) {
				msg.Data.EventCode = model.FundEventType(msg.Doc.SysEvtAck.AckDtls.EvtCd)
			}
			if !isEmpty(msg.Doc.SysEvtAck.AckDtls.EvtParam) {
				msg.Data.EventParam = string(msg.Doc.SysEvtAck.AckDtls.EvtParam)
			}
			if !isEmpty(msg.Doc.SysEvtAck.AckDtls.EvtTm) {
				msg.Data.EventTime = time.Time(msg.Doc.SysEvtAck.AckDtls.EvtTm)
			}
		}
	}
	return nil
}
