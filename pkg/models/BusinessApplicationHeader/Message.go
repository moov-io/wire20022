package BusinessApplicationHeader

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	head001 "github.com/moov-io/wire20022/pkg/models/BusinessApplicationHeader/BusinessApplicationHeader_head_001_001_03"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:head.001.001.03"

type MessageModel struct {
	MessageSenderId   string
	MessageReceiverId string
	//BizMsgIdr (Business Message Identifier) is a unique identifier assigned to an ISO 20022 message to distinguish it from other messages.
	BusinessMessageId string
	//MsgDefIdr (Message Definition Identifier) and BizSvc (Business Service) are part of the Business Application Header (BAH), which helps identify and process financial messages.
	MessageDefinitionId string
	//BizSvc specifies a business service or process related to the message.
	BusinessService string
	//<MktPrctc> (Market Practice) is used to specify market-specific rules and guidelines that apply to the message.
	MarketInfo MarketPractice

	CreateDatetime time.Time
	//BizPrcgDt stands for Business Processing Date. It represents the date when a financial transaction or message is processed by a financial institution.
	BusinessProcessingDate time.Time
	//It refers to a related Business Application Header (BAH) of type BusinessApplicationHeader71
	Relations BusinessApplicationHeader
}
type Message struct {
	Data   MessageModel
	Doc    head001.AppHdr
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
	if msg.Data.MessageSenderId == "" {
		ParamNames = append(ParamNames, "MessageSenderId")
	}
	if msg.Data.MessageReceiverId == "" {
		ParamNames = append(ParamNames, "MessageReceiverId")
	}
	if msg.Data.BusinessMessageId == "" {
		ParamNames = append(ParamNames, "BusinessMessageId")
	}
	if msg.Data.MessageDefinitionId == "" {
		ParamNames = append(ParamNames, "MessageDefinitionId")
	}
	if msg.Data.BusinessService == "" {
		ParamNames = append(ParamNames, "BusinessService")
	}
	if isEmpty(msg.Data.MarketInfo) {
		ParamNames = append(ParamNames, "MarketInfo")
	} else if msg.Data.MarketInfo.FrameworkId == "" {
		ParamNames = append(ParamNames, "MarketInfo.FrameworkId")
	} else if msg.Data.MarketInfo.ReferenceRegistry == "" {
		ParamNames = append(ParamNames, "MarketInfo.ReferenceRegistry")
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
	msg.Doc = head001.AppHdr{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "AppHdr",
		},
	}
	if msg.Data.MessageSenderId != "" {
		err := head001.ConnectionPartyIdentifierFedwireFunds1(msg.Data.MessageSenderId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageSenderId",
				Message:   err.Error(),
			}
		}
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(msg.Data.MessageSenderId),
				},
			},
		}
		msg.Doc.Fr = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if msg.Data.MessageReceiverId != "" {
		err := head001.ConnectionPartyIdentifierFedwireFunds1(msg.Data.MessageReceiverId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageReceiverId",
				Message:   err.Error(),
			}
		}
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(msg.Data.MessageReceiverId),
				},
			},
		}
		msg.Doc.To = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if msg.Data.BusinessMessageId != "" {
		err := head001.Max35Text(msg.Data.BusinessMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BusinessMessageId",
				Message:   err.Error(),
			}
		}
		msg.Doc.BizMsgIdr = head001.Max35Text(msg.Data.BusinessMessageId)
	}
	if msg.Data.MessageDefinitionId != "" {
		err := head001.MessageNameIdentificationFRS1(msg.Data.MessageDefinitionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageDefinitionId",
				Message:   err.Error(),
			}
		}
		msg.Doc.MsgDefIdr = head001.MessageNameIdentificationFRS1(msg.Data.MessageDefinitionId)
	}
	if msg.Data.BusinessService != "" {
		err := head001.BusinessServiceFedwireFunds1(msg.Data.BusinessService).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BusinessService",
				Message:   err.Error(),
			}
		}
		msg.Doc.BizSvc = head001.BusinessServiceFedwireFunds1(msg.Data.BusinessService)
	}
	if !isEmpty(msg.Data.MarketInfo) {
		MktPrctc, err := ImplementationSpecification11From(msg.Data.MarketInfo)
		if err != nil {
			err.InsertPath("MarketInfo")
			return err
		}
		if !isEmpty(MktPrctc) {
			msg.Doc.MktPrctc = MktPrctc
		}
	}
	if !isEmpty(msg.Data.CreateDatetime) {
		err := fedwire.ISODateTime(msg.Data.CreateDatetime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreateDatetime",
				Message:   err.Error(),
			}
		}
		msg.Doc.CreDt = fedwire.ISODateTime(msg.Data.CreateDatetime)
	}
	if !isEmpty(msg.Data.BusinessProcessingDate) {
		err := fedwire.ISODateTime(msg.Data.BusinessProcessingDate).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BusinessProcessingDate",
				Message:   err.Error(),
			}
		}
		BizPrcgDt := fedwire.ISODateTime(msg.Data.BusinessProcessingDate)
		msg.Doc.BizPrcgDt = &BizPrcgDt
	}
	if !isEmpty(msg.Data.Relations) {
		Rltd, err := BusinessApplicationHeader71From(msg.Data.Relations)
		if err != nil {
			err.InsertPath("Relations")
			return err
		}
		if !isEmpty(Rltd) {
			msg.Doc.Rltd = &Rltd
		}
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.Fr) && !isEmpty(msg.Doc.Fr.FIId) && !isEmpty(msg.Doc.Fr.FIId.FinInstnId) && !isEmpty(msg.Doc.Fr.FIId.FinInstnId.ClrSysMmbId) && !isEmpty(msg.Doc.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId) {
		msg.Data.MessageSenderId = string(msg.Doc.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId)
	}
	if !isEmpty(msg.Doc.To) && !isEmpty(msg.Doc.To.FIId) && !isEmpty(msg.Doc.To.FIId.FinInstnId) && !isEmpty(msg.Doc.To.FIId.FinInstnId.ClrSysMmbId) && !isEmpty(msg.Doc.To.FIId.FinInstnId.ClrSysMmbId.MmbId) {
		msg.Data.MessageReceiverId = string(msg.Doc.To.FIId.FinInstnId.ClrSysMmbId.MmbId)
	}
	if !isEmpty(msg.Doc.BizMsgIdr) {
		msg.Data.BusinessMessageId = string(msg.Doc.BizMsgIdr)
	}
	if !isEmpty(msg.Doc.MsgDefIdr) {
		msg.Data.MessageDefinitionId = string(msg.Doc.MsgDefIdr)
	}
	if !isEmpty(msg.Doc.BizSvc) {
		msg.Data.BusinessService = string(msg.Doc.BizSvc)
	}
	if !isEmpty(msg.Doc.MktPrctc) {
		if !isEmpty(msg.Doc.MktPrctc.Regy) {
			msg.Data.MarketInfo.ReferenceRegistry = string(msg.Doc.MktPrctc.Regy)
		}
		if !isEmpty(msg.Doc.MktPrctc.Id) {
			msg.Data.MarketInfo.FrameworkId = string(msg.Doc.MktPrctc.Id)
		}
	}
	if !isEmpty(msg.Doc.CreDt) {
		msg.Data.CreateDatetime = time.Time(msg.Doc.CreDt)
	}
	if !isEmpty(msg.Doc.BizPrcgDt) {
		msg.Data.BusinessProcessingDate = time.Time(*msg.Doc.BizPrcgDt)
	}
	if !isEmpty(msg.Doc.Rltd) {
		msg.Data.Relations = BusinessApplicationHeader71To(*msg.Doc.Rltd)
	}
	return nil
}
