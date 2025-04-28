package InvestigationResponse_camt_111_001_01

import (
	"encoding/xml"
	"fmt"

	camt111 "github.com/moov-io/fedwire20022/gen/InvestigationResponse_camt_111_001_01"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.111.001.01"

type MessageModel struct {
	//Point to point reference, as assigned by the responder, and sent to the requestor to unambiguously identify the message.
	MessageId string
	//Status of the investigation request.
	InvestigationStatus string
	//Provides the response to the request.
	InvestigationData string
	//Point to point reference, as assigned by the requestor, and sent to the responder to unambiguously identify the message.
	InvestRequestMessageId string
	//Type of investigation.
	InvestigationType string
	//Identification of the agent or party requesting a new investigation is opened or status update for an existing investigation.
	Requestor model.Agent
	//Identification of the agent or party expected to open a new investigation or provide a status update for an existing investigation.
	Responder model.Agent
}

type Message struct {
	data MessageModel
	doc  camt111.Document
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
func (msg *Message) CreateDocument() *model.ValidateError {
	msg.doc = camt111.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var InvstgtnRspn camt111.InvestigationResponseV01
	var _InvstgtnRspn camt111.InvestigationResponse31

	if msg.data.MessageId != "" {
		err := camt111.IMADFedwireFunds1(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		_InvstgtnRspn.MsgId = camt111.IMADFedwireFunds1(msg.data.MessageId)
	}
	if msg.data.InvestigationStatus != "" {
		err := camt111.ExternalInvestigationStatus1Code(msg.data.InvestigationStatus).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestigationStatus",
				Message:   err.Error(),
			}
		}
		_InvstgtnRspn.InvstgtnSts = camt111.InvestigationStatus21{
			Sts: camt111.ExternalInvestigationStatus1Code(msg.data.InvestigationStatus),
		}
	}
	var InvstgtnData []*camt111.InvestigationData21
	if msg.data.InvestigationData != "" {
		err := camt111.Max500Text(msg.data.InvestigationData).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestigationData",
				Message:   err.Error(),
			}
		}
		RspnNrrtv := camt111.Max500Text(msg.data.InvestigationData)
		data := camt111.InvestigationData21{
			RspnData: camt111.InvestigationDataRecord1Choice1{
				RspnNrrtv: &RspnNrrtv,
			},
		}
		InvstgtnData = append(InvstgtnData, &data)
	}
	if !isEmpty(InvstgtnData) {
		_InvstgtnRspn.InvstgtnData = InvstgtnData
	}
	if !isEmpty(_InvstgtnRspn) {
		InvstgtnRspn.InvstgtnRspn = _InvstgtnRspn
	}

	var OrgnlInvstgtnReq camt111.InvestigationRequest31
	if msg.data.InvestRequestMessageId != "" {
		err := camt111.Max35Text(msg.data.InvestRequestMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestRequestMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlInvstgtnReq.MsgId = camt111.Max35Text(msg.data.InvestRequestMessageId)
	}
	if msg.data.InvestigationType != "" {
		err := camt111.ExternalInvestigationType1Code(msg.data.InvestigationType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestigationType",
				Message:   err.Error(),
			}
		}
		Cd := camt111.ExternalInvestigationType1Code(msg.data.InvestigationType)
		OrgnlInvstgtnReq.InvstgtnTp = camt111.InvestigationType1Choice1{
			Cd: &Cd,
		}
	}
	if !isEmpty(msg.data.Requestor) {
		err := camt111.ExternalClearingSystemIdentification1CodeFixed(msg.data.Requestor.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Requestor.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = camt111.RoutingNumberFRS1(msg.data.Requestor.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Requestor.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		var Rqstr camt111.Party40Choice1
		Cd := camt111.ExternalClearingSystemIdentification1CodeFixed(msg.data.Requestor.PaymentSysCode)
		Agt := camt111.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt111.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt111.ClearingSystemMemberIdentification21{
					ClrSysId: camt111.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt111.RoutingNumberFRS1(msg.data.Requestor.PaymentSysMemberId),
				},
			},
		}
		Rqstr.Agt = &Agt
		if !isEmpty(Rqstr) {
			OrgnlInvstgtnReq.Rqstr = Rqstr
		}
	}
	if !isEmpty(msg.data.Responder) {
		err := camt111.ExternalClearingSystemIdentification1CodeFixed(msg.data.Responder.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Responder.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = camt111.RoutingNumberFRS1(msg.data.Responder.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Responder.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		var Rspndr camt111.Party40Choice1
		Cd := camt111.ExternalClearingSystemIdentification1CodeFixed(msg.data.Responder.PaymentSysCode)
		Agt := camt111.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt111.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt111.ClearingSystemMemberIdentification21{
					ClrSysId: camt111.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt111.RoutingNumberFRS1(msg.data.Responder.PaymentSysMemberId),
				},
			},
		}
		Rspndr.Agt = &Agt
		if !isEmpty(Rspndr) {
			OrgnlInvstgtnReq.Rspndr = Rspndr
		}
	}
	if !isEmpty(OrgnlInvstgtnReq) {
		InvstgtnRspn.OrgnlInvstgtnReq = OrgnlInvstgtnReq
	}
	if !isEmpty(InvstgtnRspn) {
		msg.doc.InvstgtnRspn = InvstgtnRspn
	}
	return nil
}
