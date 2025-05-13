package InvestResponse

import (
	"encoding/xml"
	"fmt"
	"strings"

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
	Data   MessageModel
	Doc    camt111.Document
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
  - With XML path: Loads file, parses XML into message.doc
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
	if msg.Data.InvestigationStatus == "" {
		ParamNames = append(ParamNames, "InvestigationStatus")
	}
	if msg.Data.InvestRequestMessageId == "" {
		ParamNames = append(ParamNames, "InvestRequestMessageId")
	}
	if msg.Data.InvestigationType == "" {
		ParamNames = append(ParamNames, "InvestigationType")
	}
	if isEmpty(msg.Data.Requestor) {
		ParamNames = append(ParamNames, "Requestor")
	}
	if isEmpty(msg.Data.Responder) {
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
	msg.Doc = camt111.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var InvstgtnRspn camt111.InvestigationResponseV01
	var _InvstgtnRspn camt111.InvestigationResponse31

	if msg.Data.MessageId != "" {
		err := camt111.IMADFedwireFunds1(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		_InvstgtnRspn.MsgId = camt111.IMADFedwireFunds1(msg.Data.MessageId)
	}
	if msg.Data.InvestigationStatus != "" {
		err := camt111.ExternalInvestigationStatus1Code(msg.Data.InvestigationStatus).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestigationStatus",
				Message:   err.Error(),
			}
		}
		_InvstgtnRspn.InvstgtnSts = camt111.InvestigationStatus21{
			Sts: camt111.ExternalInvestigationStatus1Code(msg.Data.InvestigationStatus),
		}
	}
	var InvstgtnData []*camt111.InvestigationData21
	if msg.Data.InvestigationData != "" {
		err := camt111.Max500Text(msg.Data.InvestigationData).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestigationData",
				Message:   err.Error(),
			}
		}
		RspnNrrtv := camt111.Max500Text(msg.Data.InvestigationData)
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
	if msg.Data.InvestRequestMessageId != "" {
		err := camt111.Max35Text(msg.Data.InvestRequestMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestRequestMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlInvstgtnReq.MsgId = camt111.Max35Text(msg.Data.InvestRequestMessageId)
	}
	if msg.Data.InvestigationType != "" {
		err := camt111.ExternalInvestigationType1Code(msg.Data.InvestigationType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestigationType",
				Message:   err.Error(),
			}
		}
		Cd := camt111.ExternalInvestigationType1Code(msg.Data.InvestigationType)
		OrgnlInvstgtnReq.InvstgtnTp = camt111.InvestigationType1Choice1{
			Cd: &Cd,
		}
	}
	if !isEmpty(msg.Data.Requestor) {
		err := camt111.ExternalClearingSystemIdentification1CodeFixed(msg.Data.Requestor.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Requestor.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = camt111.RoutingNumberFRS1(msg.Data.Requestor.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Requestor.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		var Rqstr camt111.Party40Choice1
		Cd := camt111.ExternalClearingSystemIdentification1CodeFixed(msg.Data.Requestor.PaymentSysCode)
		Agt := camt111.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt111.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt111.ClearingSystemMemberIdentification21{
					ClrSysId: camt111.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt111.RoutingNumberFRS1(msg.Data.Requestor.PaymentSysMemberId),
				},
			},
		}
		Rqstr.Agt = &Agt
		if !isEmpty(Rqstr) {
			OrgnlInvstgtnReq.Rqstr = Rqstr
		}
	}
	if !isEmpty(msg.Data.Responder) {
		err := camt111.ExternalClearingSystemIdentification1CodeFixed(msg.Data.Responder.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Responder.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = camt111.RoutingNumberFRS1(msg.Data.Responder.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Responder.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		var Rspndr camt111.Party40Choice1
		Cd := camt111.ExternalClearingSystemIdentification1CodeFixed(msg.Data.Responder.PaymentSysCode)
		Agt := camt111.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt111.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt111.ClearingSystemMemberIdentification21{
					ClrSysId: camt111.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt111.RoutingNumberFRS1(msg.Data.Responder.PaymentSysMemberId),
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
		msg.Doc.InvstgtnRspn = InvstgtnRspn
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.InvstgtnRspn.InvstgtnRspn) {
		if !isEmpty(msg.Doc.InvstgtnRspn.InvstgtnRspn.MsgId) {
			msg.Data.MessageId = string(msg.Doc.InvstgtnRspn.InvstgtnRspn.MsgId)
		}
		if !isEmpty(msg.Doc.InvstgtnRspn.InvstgtnRspn.InvstgtnSts) {
			msg.Data.InvestigationStatus = string(msg.Doc.InvstgtnRspn.InvstgtnRspn.InvstgtnSts.Sts)
		}
		if !isEmpty(msg.Doc.InvstgtnRspn.InvstgtnRspn.InvstgtnData) {
			if len(msg.Doc.InvstgtnRspn.InvstgtnRspn.InvstgtnData) > 0 {
				if !isEmpty(msg.Doc.InvstgtnRspn.InvstgtnRspn.InvstgtnData[0].RspnData) {
					if !isEmpty(msg.Doc.InvstgtnRspn.InvstgtnRspn.InvstgtnData[0].RspnData.RspnNrrtv) {
						msg.Data.InvestigationData = string(*msg.Doc.InvstgtnRspn.InvstgtnRspn.InvstgtnData[0].RspnData.RspnNrrtv)
					}
				}
			}
		}
	}
	if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq) {
		if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.MsgId) {
			msg.Data.InvestRequestMessageId = string(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.MsgId)
		}
		if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.InvstgtnTp) {
			if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.InvstgtnTp.Cd) {
				msg.Data.InvestigationType = string(*msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.InvstgtnTp.Cd)
			}
		}
		if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr) {
			if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr.Agt) {
				if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr.Agt.FinInstnId) {
					if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId) {
						if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.ClrSysId) {
							msg.Data.Requestor.PaymentSysCode = model.PaymentSystemType(*msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
						}
						if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.MmbId) {
							msg.Data.Requestor.PaymentSysMemberId = string(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.MmbId)
						}
					}
				}
			}
		}
		if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr) {
			if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr.Agt) {
				if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr.Agt.FinInstnId) {
					if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId) {
						if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.ClrSysId) {
							msg.Data.Responder.PaymentSysCode = model.PaymentSystemType(*msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
						}
						if !isEmpty(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.MmbId) {
							msg.Data.Responder.PaymentSysMemberId = string(msg.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.MmbId)
						}
					}
				}
			}
		}
	}
	return nil
}
