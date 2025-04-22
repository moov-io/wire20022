package ReturnRequestResponse

import (
	"encoding/xml"
	"time"

	camt029 "github.com/moov-io/fedwire20022/gen/ReturnRequestResponse_camt_029_001_09"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09"

type MessageModel struct {
	//Uniquely identifies the case assignment.
	AssignmentId string
	//Party who assigns the case.
	Assigner model.Agent
	//Party to which the case is assigned.
	Assignee model.Agent
	//Date and time at which the assignment was created.
	AssignmentCreateTime time.Time
	//Identifies a resolved case.
	ResolvedCaseId string
	//Party that created the investigation case.
	Creator model.Agent
	//Specifies the status of the investigation, in a coded form.
	Status Status
	//Point to point reference assigned by the original instructing party to unambiguously identify the original message.
	OriginalMessageId string
	//Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.
	OriginalMessageNameId string
	//Original date and time at which the message was created.
	OriginalMessageCreateTime time.Time
	//Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId string
	//Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId string
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUETR string
	//Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInfo Reason
}
type Message struct {
	data MessageModel
	doc  camt029.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = camt029.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}

	var RsltnOfInvstgtn camt029.ResolutionOfInvestigationV09
	var Assgnmt camt029.CaseAssignment51
	if msg.data.AssignmentId != "" {
		Assgnmt.Id = camt029.Max35Text(msg.data.AssignmentId)
	}
	if !isEmpty(msg.data.Assigner) {
		Assgnmt.Assgnr = Party40Choice1From(msg.data.Assigner)
	}
	if !isEmpty(msg.data.Assignee) {
		Assgnmt.Assgne = Party40Choice1From(msg.data.Assignee)
	}
	if !isEmpty(msg.data.AssignmentCreateTime) {
		Assgnmt.CreDtTm = fedwire.ISODateTime(msg.data.AssignmentCreateTime)
	}
	if !isEmpty(Assgnmt) {
		RsltnOfInvstgtn.Assgnmt = Assgnmt
	}
	var RslvdCase camt029.Case51
	if msg.data.ResolvedCaseId != "" {
		RslvdCase.Id = camt029.Max35Text(msg.data.ResolvedCaseId)
	}
	if !isEmpty(msg.data.Creator) {
		RslvdCase.Cretr = Party40Choice2From(msg.data.Creator)
	}
	if !isEmpty(RslvdCase) {
		RsltnOfInvstgtn.RslvdCase = RslvdCase
	}
	var Sts camt029.InvestigationStatus5Choice1
	if msg.data.Status != "" {
		Conf := camt029.ExternalInvestigationExecutionConfirmation1Code(msg.data.Status)
		Sts.Conf = &Conf
	}
	if !isEmpty(Sts) {
		RsltnOfInvstgtn.Sts = Sts
	}
	var CxlDtls camt029.UnderlyingTransaction221
	var TxInfAndSts camt029.PaymentTransaction1021
	var OrgnlGrpInf camt029.OriginalGroupInformation291
	if msg.data.OriginalMessageId != "" {
		OrgnlGrpInf.OrgnlMsgId = camt029.Max35Text(msg.data.OriginalMessageId)
	}
	if msg.data.OriginalMessageNameId != "" {
		OrgnlGrpInf.OrgnlMsgNmId = camt029.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
	}
	if !isEmpty(msg.data.OriginalMessageCreateTime) {
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(msg.data.OriginalMessageCreateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		TxInfAndSts.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.data.OriginalInstructionId != "" {
		OrgnlInstrId := camt029.Max35Text(msg.data.OriginalInstructionId)
		TxInfAndSts.OrgnlInstrId = &OrgnlInstrId
	}
	if msg.data.OriginalEndToEndId != "" {
		OrgnlEndToEndId := camt029.Max35Text(msg.data.OriginalEndToEndId)
		TxInfAndSts.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if msg.data.OriginalUETR != "" {
		TxInfAndSts.OrgnlUETR = camt029.UUIDv4Identifier(msg.data.OriginalUETR)
	}
	if !isEmpty(msg.data.CancellationStatusReasonInfo) {
		var CxlStsRsnInf []*camt029.CancellationStatusReason41
		reason := CancellationStatusReason41From(msg.data.CancellationStatusReasonInfo)
		CxlStsRsnInf = append(CxlStsRsnInf, &reason)
		TxInfAndSts.CxlStsRsnInf = CxlStsRsnInf
	}
	if !isEmpty(TxInfAndSts) {
		CxlDtls.TxInfAndSts = TxInfAndSts
	}
	if !isEmpty(CxlDtls) {
		RsltnOfInvstgtn.CxlDtls = CxlDtls
	}
	if !isEmpty(RsltnOfInvstgtn) {
		msg.doc.RsltnOfInvstgtn = RsltnOfInvstgtn
	}
}
