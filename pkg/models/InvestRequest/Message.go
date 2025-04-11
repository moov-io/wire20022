package InvestRequest

import (
	"os"
	"path/filepath"

	camt110 "github.com/moov-io/fedwire20022/gen/InvestigationRequest_camt_110_001_01"
	model "github.com/moov-io/wire20022/pkg/models"
)

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

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = camt110.Document{}
	var InvstgtnReq camt110.InvestigationRequestV01
	var _InvstgtnReq camt110.InvestigationRequest21
	if msg.data.MessageId != "" {
		MsgId := camt110.IMADFedwireFunds1(msg.data.MessageId)
		_InvstgtnReq.MsgId = MsgId
	}
	if msg.data.InvestigationType != "" {
		Cd := camt110.ExternalInvestigationType1Code(msg.data.InvestigationType)
		_InvstgtnReq.InvstgtnTp = camt110.InvestigationType1Choice1{
			Cd: &Cd,
		}
	}
	if !isEmpty(msg.data.UnderlyingData) {
		_InvstgtnReq.Undrlyg = UnderlyingData2Choice1From(msg.data.UnderlyingData)
	}
	if !isEmpty(msg.data.Requestor) {
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
		reason := InvestigationReason21From(msg.data.InvestReason)
		InvstgtnData = append(InvstgtnData, reason)
	}
	if !isEmpty(InvstgtnData) {
		InvstgtnReq.InvstgtnData = InvstgtnData
	}
	if !isEmpty(InvstgtnReq) {
		msg.doc.InvstgtnReq = InvstgtnReq
	}
}
func WriteXMLTo(filePath string, xml []byte) error {
	os.Mkdir("generated", 0755)
	xmlFileName := filepath.Join("generated", filePath)

	return os.WriteFile(xmlFileName, xml, 0644)
}
