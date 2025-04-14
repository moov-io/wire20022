package InvestigationResponse_camt_111_001_01

import (
	camt111 "github.com/moov-io/fedwire20022/gen/InvestigationResponse_camt_111_001_01"
	model "github.com/moov-io/wire20022/pkg/models"
)

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

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = camt111.Document{}
	var InvstgtnRspn camt111.InvestigationResponseV01
	var _InvstgtnRspn camt111.InvestigationResponse31

	if msg.data.MessageId != "" {
		_InvstgtnRspn.MsgId = camt111.IMADFedwireFunds1(msg.data.MessageId)
	}
	if msg.data.InvestigationStatus != "" {
		_InvstgtnRspn.InvstgtnSts = camt111.InvestigationStatus21{
			Sts: camt111.ExternalInvestigationStatus1Code(msg.data.InvestigationStatus),
		}
	}
	var InvstgtnData []*camt111.InvestigationData21
	if msg.data.InvestigationData != "" {
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
		OrgnlInvstgtnReq.MsgId = camt111.Max35Text(msg.data.InvestRequestMessageId)
	}
	if msg.data.InvestigationType != "" {
		Cd := camt111.ExternalInvestigationType1Code(msg.data.InvestigationType)
		OrgnlInvstgtnReq.InvstgtnTp = camt111.InvestigationType1Choice1{
			Cd: &Cd,
		}
	}
	if !isEmpty(msg.data.Requestor) {
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
}
