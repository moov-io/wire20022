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
	Data   MessageModel
	Doc    admi004.Document
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
	if msg.Data.EventType == "" {
		ParamNames = append(ParamNames, "EventType")
	}
	if msg.Data.EventParam == "" {
		ParamNames = append(ParamNames, "EventParam")
	}
	if isEmpty(msg.Data.EventTime) {
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
	msg.Doc = admi004.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var EvtInf admi004.Event21
	if msg.Data.EventType != "" {
		err := admi004.EventFedwireFunds1(msg.Data.EventType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventType",
				Message:   err.Error(),
			}
		}
		EvtInf.EvtCd = admi004.EventFedwireFunds1(msg.Data.EventType)
	}
	if msg.Data.EventParam != "" {
		err := admi004.EndpointIdentifierFedwireFunds1(msg.Data.EventParam).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EvntParam",
				Message:   err.Error(),
			}
		}
		EvtInf.EvtParam = admi004.EndpointIdentifierFedwireFunds1(msg.Data.EventParam)
	}
	if !isEmpty(msg.Data.EventTime) {
		err := fedwire.ISODateTime(msg.Data.EventTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EventTime",
				Message:   err.Error(),
			}
		}
		EvtInf.EvtTm = fedwire.ISODateTime(msg.Data.EventTime)
	}
	if !isEmpty(EvtInf) {
		msg.Doc.SysEvtNtfctn = admi004.SystemEventNotificationV02{
			EvtInf: EvtInf,
		}
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc) && !isEmpty(msg.Doc.SysEvtNtfctn) && !isEmpty(msg.Doc.SysEvtNtfctn.EvtInf) {
		if !isEmpty(msg.Doc.SysEvtNtfctn.EvtInf.EvtCd) {
			msg.Data.EventType = string(msg.Doc.SysEvtNtfctn.EvtInf.EvtCd)
		}
		if !isEmpty(msg.Doc.SysEvtNtfctn.EvtInf.EvtParam) {
			msg.Data.EventParam = string(msg.Doc.SysEvtNtfctn.EvtInf.EvtParam)

		}
		if !isEmpty(msg.Doc.SysEvtNtfctn.EvtInf.EvtTm) {
			msg.Data.EventTime = time.Time(msg.Doc.SysEvtNtfctn.EvtInf.EvtTm)
		}
	}
	return nil
}
