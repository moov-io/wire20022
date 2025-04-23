package FedwireFundsSystemResponse

import (
	"encoding/xml"
	"time"

	admi011 "github.com/moov-io/fedwire20022/gen/FedwireFundsSystemResponse_admi_011_001_01"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:admi.011.001.01"

type MessageModel struct {
	//Unique and unambiguous identifier for the message, as assigned by the sender.
	MessageId string
	//Proprietary code used to specify an event that occurred in a system.
	EventCode model.FundEventType
	//Describes the parameters of an event which occurred in a system.
	EventParam string
	//Date and time at which the event occurred.
	EventTime time.Time
}
type Message struct {
	data MessageModel
	doc  admi011.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = admi011.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var SysEvtAck admi011.SystemEventAcknowledgementV01
	if msg.data.MessageId != "" {
		SysEvtAck.MsgId = admi011.Max35Text(msg.data.MessageId)
	}
	var AckDtls admi011.Event11
	if msg.data.EventCode != "" {
		AckDtls.EvtCd = admi011.EventFedwireFunds1(msg.data.EventCode)
	}
	if msg.data.EventParam != "" {
		AckDtls.EvtParam = admi011.EndpointIdentifierFedwireFunds1(msg.data.EventParam)
	}
	if !isEmpty(msg.data.EventTime) {
		AckDtls.EvtTm = fedwire.ISODateTime(msg.data.EventTime)
	}
	if !isEmpty(AckDtls) {
		SysEvtAck.AckDtls = AckDtls
	}
	if !isEmpty(SysEvtAck) {
		msg.doc.SysEvtAck = SysEvtAck
	}
}
