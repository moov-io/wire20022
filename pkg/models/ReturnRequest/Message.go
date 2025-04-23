package ReturnRequest

import (
	"encoding/xml"
	"time"

	camt056 "github.com/moov-io/fedwire20022/gen/ReturnRequest_camt_056_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.056.001.08"

type MessageModel struct {
	//Uniquely identifies the case assignment.
	AssignmentId string
	//Party who assigns the case.
	Assigner model.Agent
	//Party to which the case is assigned.
	Assignee model.Agent
	//Date and time at which the assignment was created.
	AssignmentCreateTime time.Time
	//Identifies the investigation case.
	CaseId string
	//Party that created the investigation case.
	Creator model.Agent
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
	//Amount of money moved between the instructing agent and the instructed agent, as provided in the original instruction.
	OriginalInterbankSettlementAmount model.CurrencyAndAmount
	//Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate model.Date
	//Provides detailed information on the cancellation reason.
	CancellationReason Reason
}
type Message struct {
	data MessageModel
	doc  camt056.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = camt056.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var FIToFIPmtCxlReq camt056.FIToFIPaymentCancellationRequestV08
	var Assgnmt camt056.CaseAssignment51
	if msg.data.AssignmentId != "" {
		Assgnmt.Id = camt056.IMADFedwireFunds1(msg.data.AssignmentId)
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
		FIToFIPmtCxlReq.Assgnmt = Assgnmt
	}
	var Case camt056.Case51
	if msg.data.CaseId != "" {
		Case.Id = camt056.Max35Text(msg.data.CaseId)
	}
	if !isEmpty(msg.data.Creator) {
		Case.Cretr = Party40Choice2From(msg.data.Creator)
	}
	if !isEmpty(Case) {
		FIToFIPmtCxlReq.Case = Case
	}
	var Undrlyg camt056.UnderlyingTransaction231
	var TxInf camt056.PaymentTransaction1061
	var OrgnlGrpInf camt056.OriginalGroupInformation291
	if msg.data.OriginalMessageId != "" {
		OrgnlGrpInf.OrgnlMsgId = camt056.Max35Text(msg.data.OriginalMessageId)
	}
	if msg.data.OriginalMessageNameId != "" {
		OrgnlGrpInf.OrgnlMsgNmId = camt056.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
	}
	if !isEmpty(msg.data.OriginalMessageCreateTime) {
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(msg.data.OriginalMessageCreateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		TxInf.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.data.OriginalInstructionId != "" {
		OrgnlInstrId := camt056.Max35Text(msg.data.OriginalInstructionId)
		TxInf.OrgnlInstrId = &OrgnlInstrId
	}
	if msg.data.OriginalEndToEndId != "" {
		OrgnlEndToEndId := camt056.Max35Text(msg.data.OriginalEndToEndId)
		TxInf.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if msg.data.OriginalUETR != "" {
		TxInf.OrgnlUETR = camt056.UUIDv4Identifier(msg.data.OriginalUETR)
	}
	if !isEmpty(msg.data.OriginalInterbankSettlementAmount) {
		TxInf.OrgnlIntrBkSttlmAmt = camt056.ActiveOrHistoricCurrencyAndAmount{
			Value: camt056.ActiveOrHistoricCurrencyAndAmountSimpleType(msg.data.OriginalInterbankSettlementAmount.Amount),
			Ccy:   camt056.ActiveOrHistoricCurrencyCode(msg.data.OriginalInterbankSettlementAmount.Currency),
		}
	}
	if !isEmpty(msg.data.OriginalInterbankSettlementDate) {
		TxInf.OrgnlIntrBkSttlmDt = msg.data.OriginalInterbankSettlementDate.Date()
	}
	if !isEmpty(msg.data.CancellationReason) {
		TxInf.CxlRsnInf = PaymentCancellationReason51From(msg.data.CancellationReason)
	}
	if !isEmpty(TxInf) {
		Undrlyg.TxInf = TxInf
	}
	if !isEmpty(Undrlyg) {
		FIToFIPmtCxlReq.Undrlyg = Undrlyg
	}
	if !isEmpty(FIToFIPmtCxlReq) {
		msg.doc.FIToFIPmtCxlReq = FIToFIPmtCxlReq
	}
}
