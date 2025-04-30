package InvestRequest

import (
	"encoding/xml"
	"fmt"
	"strings"

	camt110 "github.com/moov-io/fedwire20022/gen/InvestigationRequest_camt_110_001_01"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.110.001.01"

type MessageModel struct {
	//Point to point reference, as assigned by the requestor, and sent to the responder to unambiguously identify the message.
	MessageId string
	//Type of investigation.
	InvestigationType string
	//Provides details on the subject to which the investigation refers, for example a payment or statement entry.
	UnderlyingData Underlying
	//Identification of the agent or party requesting a new investigation is opened or status update for an existing investigation.
	Requestor model.Agent
	//Identification of the agent or party expected to open a new investigation or provide a status update for an existing investigation.
	Responder model.Agent
	//Reason for the investigation being opened, in a proprietary form.
	InvestReason InvestigationReason
}

type Message struct {
	data MessageModel
	doc  camt110.Document
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
	if msg.data.InvestigationType == "" {
		ParamNames = append(ParamNames, "InvestigationType")
	}
	if isEmpty(msg.data.UnderlyingData) {
		ParamNames = append(ParamNames, "UnderlyingData")
	}
	if isEmpty(msg.data.Requestor) {
		ParamNames = append(ParamNames, "Requestor")
	}
	if isEmpty(msg.data.Responder) {
		ParamNames = append(ParamNames, "Responder")
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
	msg.doc = camt110.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var InvstgtnReq camt110.InvestigationRequestV01
	var _InvstgtnReq camt110.InvestigationRequest21
	if msg.data.MessageId != "" {
		err := camt110.IMADFedwireFunds1(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		MsgId := camt110.IMADFedwireFunds1(msg.data.MessageId)
		_InvstgtnReq.MsgId = MsgId
	}
	if msg.data.InvestigationType != "" {
		err := camt110.ExternalInvestigationType1Code(msg.data.InvestigationType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestigationType",
				Message:   err.Error(),
			}
		}
		Cd := camt110.ExternalInvestigationType1Code(msg.data.InvestigationType)
		_InvstgtnReq.InvstgtnTp = camt110.InvestigationType1Choice1{
			Cd: &Cd,
		}
	}
	if !isEmpty(msg.data.UnderlyingData) {
		Undrlyg, vErr := UnderlyingData2Choice1From(msg.data.UnderlyingData)
		if vErr != nil {
			vErr.InsertPath("UnderlyingData")
			return vErr
		}
		_InvstgtnReq.Undrlyg = Undrlyg
	}
	if !isEmpty(msg.data.Requestor) {
		err := camt110.ExternalClearingSystemIdentification1CodeFixed(msg.data.Requestor.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Requestor.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = camt110.RoutingNumberFRS1(msg.data.Requestor.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Requestor.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		var Rqstr camt110.Party40Choice1
		Cd := camt110.ExternalClearingSystemIdentification1CodeFixed(msg.data.Requestor.PaymentSysCode)
		Agt := camt110.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt110.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt110.ClearingSystemMemberIdentification21{
					ClrSysId: camt110.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt110.RoutingNumberFRS1(msg.data.Requestor.PaymentSysMemberId),
				},
			},
		}
		Rqstr.Agt = &Agt
		if !isEmpty(Rqstr) {
			_InvstgtnReq.Rqstr = Rqstr
		}
	}
	if !isEmpty(msg.data.Responder) {
		err := camt110.ExternalClearingSystemIdentification1CodeFixed(msg.data.Responder.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Responder.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = camt110.RoutingNumberFRS1(msg.data.Responder.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Responder.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		var Rspndr camt110.Party40Choice1
		Cd := camt110.ExternalClearingSystemIdentification1CodeFixed(msg.data.Responder.PaymentSysCode)
		Agt := camt110.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt110.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt110.ClearingSystemMemberIdentification21{
					ClrSysId: camt110.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt110.RoutingNumberFRS1(msg.data.Responder.PaymentSysMemberId),
				},
			},
		}
		Rspndr.Agt = &Agt
		if !isEmpty(Rspndr) {
			_InvstgtnReq.Rspndr = Rspndr
		}
	}
	if !isEmpty(_InvstgtnReq) {
		InvstgtnReq.InvstgtnReq = _InvstgtnReq
	}
	var InvstgtnData []camt110.InvestigationReason21
	if !isEmpty(msg.data.InvestReason) {
		reason, vErr := InvestigationReason21From(msg.data.InvestReason)
		if vErr != nil {
			vErr.InsertPath("InvestReason")
			return vErr
		}
		InvstgtnData = append(InvstgtnData, reason)
	}
	if !isEmpty(InvstgtnData) {
		InvstgtnReq.InvstgtnData = InvstgtnData
	}
	if !isEmpty(InvstgtnReq) {
		msg.doc.InvstgtnReq = InvstgtnReq
	}
	return nil
}
