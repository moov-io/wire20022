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
	Data   MessageModel
	Doc    admi007.Document
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
	if msg.Data.CreatedDateTime.IsZero() {
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if msg.Data.RelationReference == "" {
		ParamNames = append(ParamNames, "RelationReference")
	}
	if msg.Data.ReferenceName == "" {
		ParamNames = append(ParamNames, "ReferenceName")
	}
	if msg.Data.RequestHandling == "" {
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
	msg.Doc = admi007.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var RctAck admi007.ReceiptAcknowledgementV01
	if msg.Data.MessageId != "" {
		err := admi007.OMADFedwireFunds1(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		RctAck.MsgId.MsgId = admi007.OMADFedwireFunds1(msg.Data.MessageId)
	}
	if !isEmpty(msg.Data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.Data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		RctAck.MsgId.CreDtTm = fedwire.ISODateTime(msg.Data.CreatedDateTime)
	}
	if msg.Data.RelationReference != "" {
		err := admi007.Max35Text(msg.Data.RelationReference).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RelationReference",
				Message:   err.Error(),
			}
		}
		RctAck.Rpt.RltdRef.Ref = admi007.Max35Text(msg.Data.RelationReference)
	}
	if msg.Data.ReferenceName != "" {
		err := admi007.MessageNameIdentificationFRS1(msg.Data.ReferenceName).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReferenceName",
				Message:   err.Error(),
			}
		}
		RctAck.Rpt.RltdRef.MsgNm = admi007.MessageNameIdentificationFRS1(msg.Data.ReferenceName)
	}
	if msg.Data.RequestHandling != "" {
		err := admi007.Max4AlphaNumericTextFixed(msg.Data.RequestHandling).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RequestHandling",
				Message:   err.Error(),
			}
		}
		RctAck.Rpt.ReqHdlg.StsCd = admi007.Max4AlphaNumericTextFixed(msg.Data.RequestHandling)
	}
	if !isEmpty(RctAck) {
		msg.Doc.RctAck = RctAck
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.RctAck) {
		if !isEmpty(msg.Doc.RctAck.MsgId) {
			msg.Data.MessageId = string(msg.Doc.RctAck.MsgId.MsgId)
		}
		if !isEmpty(msg.Doc.RctAck.MsgId.CreDtTm) {
			msg.Data.CreatedDateTime = time.Time(msg.Doc.RctAck.MsgId.CreDtTm)
		}
		if !isEmpty(msg.Doc.RctAck.Rpt) {
			if !isEmpty(msg.Doc.RctAck.Rpt.RltdRef) {
				msg.Data.RelationReference = string(msg.Doc.RctAck.Rpt.RltdRef.Ref)
			}
			if !isEmpty(msg.Doc.RctAck.Rpt.RltdRef.MsgNm) {
				msg.Data.ReferenceName = string(msg.Doc.RctAck.Rpt.RltdRef.MsgNm)
			}
			if !isEmpty(msg.Doc.RctAck.Rpt.ReqHdlg) {
				msg.Data.RequestHandling = model.RelatedStatusCode(msg.Doc.RctAck.Rpt.ReqHdlg.StsCd)
			}
		}
	}
	return nil
}
