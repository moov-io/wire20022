package BusinessApplicationHeader

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	head001 "github.com/moov-io/fedwire20022/gen/BusinessApplicationHeader_head_001_001_03"
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
	data MessageModel
	doc  head001.AppHdr
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
	if msg.data.MessageSenderId == "" {
		ParamNames = append(ParamNames, "MessageSenderId")
	}
	if msg.data.MessageReceiverId == "" {
		ParamNames = append(ParamNames, "MessageReceiverId")
	}
	if msg.data.BusinessMessageId == "" {
		ParamNames = append(ParamNames, "BusinessMessageId")
	}
	if msg.data.MessageDefinitionId == "" {
		ParamNames = append(ParamNames, "MessageDefinitionId")
	}
	if msg.data.BusinessService == "" {
		ParamNames = append(ParamNames, "BusinessService")
	}
	if isEmpty(msg.data.MarketInfo) {
		ParamNames = append(ParamNames, "MarketInfo")
	} else if msg.data.MarketInfo.FrameworkId == "" {
		ParamNames = append(ParamNames, "MarketInfo.FrameworkId")
	} else if msg.data.MarketInfo.ReferenceRegistry == "" {
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
	msg.doc = head001.AppHdr{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "AppHdr",
		},
	}
	if msg.data.MessageSenderId != "" {
		err := head001.ConnectionPartyIdentifierFedwireFunds1(msg.data.MessageSenderId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageSenderId",
				Message:   err.Error(),
			}
		}
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(msg.data.MessageSenderId),
				},
			},
		}
		msg.doc.Fr = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if msg.data.MessageReceiverId != "" {
		err := head001.ConnectionPartyIdentifierFedwireFunds1(msg.data.MessageReceiverId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageReceiverId",
				Message:   err.Error(),
			}
		}
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(msg.data.MessageReceiverId),
				},
			},
		}
		msg.doc.To = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if msg.data.BusinessMessageId != "" {
		err := head001.Max35Text(msg.data.BusinessMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BusinessMessageId",
				Message:   err.Error(),
			}
		}
		msg.doc.BizMsgIdr = head001.Max35Text(msg.data.BusinessMessageId)
	}
	if msg.data.MessageDefinitionId != "" {
		err := head001.MessageNameIdentificationFRS1(msg.data.MessageDefinitionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageDefinitionId",
				Message:   err.Error(),
			}
		}
		msg.doc.MsgDefIdr = head001.MessageNameIdentificationFRS1(msg.data.MessageDefinitionId)
	}
	if msg.data.BusinessService != "" {
		err := head001.BusinessServiceFedwireFunds1(msg.data.BusinessService).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BusinessService",
				Message:   err.Error(),
			}
		}
		msg.doc.BizSvc = head001.BusinessServiceFedwireFunds1(msg.data.BusinessService)
	}
	if !isEmpty(msg.data.MarketInfo) {
		MktPrctc, err := ImplementationSpecification11From(msg.data.MarketInfo)
		if err != nil {
			err.InsertPath("MarketInfo")
			return err
		}
		if !isEmpty(MktPrctc) {
			msg.doc.MktPrctc = MktPrctc
		}
	}
	if !isEmpty(msg.data.CreateDatetime) {
		err := fedwire.ISODateTime(msg.data.CreateDatetime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreateDatetime",
				Message:   err.Error(),
			}
		}
		msg.doc.CreDt = fedwire.ISODateTime(msg.data.CreateDatetime)
	}
	if !isEmpty(msg.data.BusinessProcessingDate) {
		err := fedwire.ISODateTime(msg.data.BusinessProcessingDate).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BusinessProcessingDate",
				Message:   err.Error(),
			}
		}
		BizPrcgDt := fedwire.ISODateTime(msg.data.BusinessProcessingDate)
		msg.doc.BizPrcgDt = &BizPrcgDt
	}
	if !isEmpty(msg.data.Relations) {
		Rltd, err := BusinessApplicationHeader71From(msg.data.Relations)
		if err != nil {
			err.InsertPath("Relations")
			return err
		}
		if !isEmpty(Rltd) {
			msg.doc.Rltd = &Rltd
		}
	}
	return nil
}
