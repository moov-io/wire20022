package PaymentStatusRequest

import (
	"encoding/xml"
	"time"

	pacs004 "github.com/moov-io/fedwire20022/gen/PaymentStatusRequest_pacs_028_001_03"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03"

type MessageModel struct {
	//Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	MessageId string
	//Date and time at which the message was created.
	CreatedDateTime time.Time
	//Point to point reference assigned by the original instructing party to unambiguously identify the original message.
	OriginalMessageId string
	//Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.
	OriginalMessageNameId string
	//Original date and time at which the message was created.
	OriginalCreationDateTime time.Time
	//Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId string
	//Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId string
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUETR string
	//Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent model.Agent
	//Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent model.Agent
}
type Message struct {
	data MessageModel
	doc  pacs004.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = pacs004.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var FIToFIPmtStsReq pacs004.FIToFIPaymentStatusRequestV03
	var GrpHdr pacs004.GroupHeader911
	if msg.data.MessageId != "" {
		GrpHdr.MsgId = pacs004.Max35Text(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
	if !isEmpty(GrpHdr) {
		FIToFIPmtStsReq.GrpHdr = GrpHdr
	}
	var TxInf pacs004.PaymentTransaction1131
	var OrgnlGrpInf pacs004.OriginalGroupInformation291
	if msg.data.OriginalMessageId != "" {
		OrgnlGrpInf.OrgnlMsgId = pacs004.IMADFedwireFunds1(msg.data.OriginalMessageId)
	}
	if msg.data.OriginalMessageNameId != "" {
		OrgnlGrpInf.OrgnlMsgNmId = pacs004.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
	}
	if !isEmpty(msg.data.OriginalCreationDateTime) {
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(msg.data.OriginalCreationDateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		TxInf.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.data.OriginalInstructionId != "" {
		OrgnlInstrId := pacs004.Max35Text(msg.data.OriginalInstructionId)
		TxInf.OrgnlInstrId = &OrgnlInstrId
	}
	if msg.data.OriginalEndToEndId != "" {
		OrgnlEndToEndId := pacs004.Max35Text(msg.data.OriginalEndToEndId)
		TxInf.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if msg.data.OriginalUETR != "" {
		TxInf.OrgnlUETR = pacs004.UUIDv4Identifier(msg.data.OriginalUETR)
	}
	if !isEmpty(msg.data.InstructingAgent) {
		TxInf.InstgAgt = BranchAndFinancialInstitutionIdentification61From(msg.data.InstructingAgent)
	}
	if !isEmpty(msg.data.InstructedAgent) {
		TxInf.InstdAgt = BranchAndFinancialInstitutionIdentification61From(msg.data.InstructedAgent)
	}
	if !isEmpty(TxInf) {
		FIToFIPmtStsReq.TxInf = TxInf
	}
	if !isEmpty(FIToFIPmtStsReq) {
		msg.doc.FIToFIPmtStsReq = FIToFIPmtStsReq
	}
}
