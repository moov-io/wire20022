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
	Data   MessageModel
	Doc    camt110.Document
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
	if msg.Data.InvestigationType == "" {
		ParamNames = append(ParamNames, "InvestigationType")
	}
	if isEmpty(msg.Data.UnderlyingData) {
		ParamNames = append(ParamNames, "UnderlyingData")
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
	msg.Doc = camt110.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var InvstgtnReq camt110.InvestigationRequestV01
	var _InvstgtnReq camt110.InvestigationRequest21
	if msg.Data.MessageId != "" {
		err := camt110.IMADFedwireFunds1(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		MsgId := camt110.IMADFedwireFunds1(msg.Data.MessageId)
		_InvstgtnReq.MsgId = MsgId
	}
	if msg.Data.InvestigationType != "" {
		err := camt110.ExternalInvestigationType1Code(msg.Data.InvestigationType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InvestigationType",
				Message:   err.Error(),
			}
		}
		Cd := camt110.ExternalInvestigationType1Code(msg.Data.InvestigationType)
		_InvstgtnReq.InvstgtnTp = camt110.InvestigationType1Choice1{
			Cd: &Cd,
		}
	}
	if !isEmpty(msg.Data.UnderlyingData) {
		Undrlyg, vErr := UnderlyingData2Choice1From(msg.Data.UnderlyingData)
		if vErr != nil {
			vErr.InsertPath("UnderlyingData")
			return vErr
		}
		_InvstgtnReq.Undrlyg = Undrlyg
	}
	if !isEmpty(msg.Data.Requestor) {
		err := camt110.ExternalClearingSystemIdentification1CodeFixed(msg.Data.Requestor.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Requestor.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = camt110.RoutingNumberFRS1(msg.Data.Requestor.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Requestor.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		var Rqstr camt110.Party40Choice1
		Cd := camt110.ExternalClearingSystemIdentification1CodeFixed(msg.Data.Requestor.PaymentSysCode)
		Agt := camt110.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt110.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt110.ClearingSystemMemberIdentification21{
					ClrSysId: camt110.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt110.RoutingNumberFRS1(msg.Data.Requestor.PaymentSysMemberId),
				},
			},
		}
		Rqstr.Agt = &Agt
		if !isEmpty(Rqstr) {
			_InvstgtnReq.Rqstr = Rqstr
		}
	}
	if !isEmpty(msg.Data.Responder) {
		err := camt110.ExternalClearingSystemIdentification1CodeFixed(msg.Data.Responder.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Responder.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = camt110.RoutingNumberFRS1(msg.Data.Responder.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Responder.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		var Rspndr camt110.Party40Choice1
		Cd := camt110.ExternalClearingSystemIdentification1CodeFixed(msg.Data.Responder.PaymentSysCode)
		Agt := camt110.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt110.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt110.ClearingSystemMemberIdentification21{
					ClrSysId: camt110.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt110.RoutingNumberFRS1(msg.Data.Responder.PaymentSysMemberId),
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
	if !isEmpty(msg.Data.InvestReason) {
		reason, vErr := InvestigationReason21From(msg.Data.InvestReason)
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
		msg.Doc.InvstgtnReq = InvstgtnReq
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.InvstgtnReq) {
		if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq) {
			if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.MsgId) {
				msg.Data.MessageId = string(msg.Doc.InvstgtnReq.InvstgtnReq.MsgId)
			}
			if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.InvstgtnTp) {
				if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.InvstgtnTp.Cd) {
					msg.Data.InvestigationType = string(*msg.Doc.InvstgtnReq.InvstgtnReq.InvstgtnTp.Cd)
				}
			}
			if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Undrlyg) {
				undrlyg := UnderlyingData2Choice1To(msg.Doc.InvstgtnReq.InvstgtnReq.Undrlyg)
				msg.Data.UnderlyingData = undrlyg
			}
			if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rqstr) {
				if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rqstr.Agt) {
					if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rqstr.Agt.FinInstnId) {
						if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId) {
							if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.ClrSysId) {
								if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
									msg.Data.Requestor.PaymentSysCode = model.PaymentSystemType(*msg.Doc.InvstgtnReq.InvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
								}
							}
							if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.MmbId) {
								msg.Data.Requestor.PaymentSysMemberId = string(msg.Doc.InvstgtnReq.InvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.MmbId)
							}
						}
					}
				}
			}
			if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rspndr) {
				if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rspndr.Agt) {
					if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rspndr.Agt.FinInstnId) {
						if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId) {
							if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.ClrSysId) {
								if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
									msg.Data.Responder.PaymentSysCode = model.PaymentSystemType(*msg.Doc.InvstgtnReq.InvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
								}
							}
							if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.MmbId) {
								msg.Data.Responder.PaymentSysMemberId = string(msg.Doc.InvstgtnReq.InvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.MmbId)
							}
						}
					}
				}
			}
		}
		if !isEmpty(msg.Doc.InvstgtnReq.InvstgtnData) {
			for _, invstgtnData := range msg.Doc.InvstgtnReq.InvstgtnData {
				investReason := InvestigationReason21To(invstgtnData)
				msg.Data.InvestReason = investReason
			}
		}

	}
	return nil
}
