package FedwireFundsBroadcast

import (
	"encoding/xml"
	"time"

	admi004 "github.com/moov-io/fedwire20022/gen/FedwireFundsBroadcast_admi_004_001_02"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02"

type MessageModel struct {
	//Proprietary code used to specify an event that occurred in a system.
	EventCode model.FundEventType
	//Describes the parameters of an event which occurred in a system.
	EventParam model.Date
	//Free text used to describe an event which occurred in a system.
	EventDescription string
	//Date and time at which the event occurred.
	EventTime time.Time
}
type Message struct {
	data MessageModel
	doc  admi004.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = admi004.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var SysEvtNtfctn admi004.SystemEventNotificationV02
	var EvtInf admi004.Event21
	if msg.data.EventCode != "" {
		EvtInf.EvtCd = admi004.EventFedwireFunds1(msg.data.EventCode)
	}
	if !isEmpty(msg.data.EventParam) {
		EvtInf.EvtParam = msg.data.EventParam.Date()
	}
	if msg.data.EventDescription != "" {
		EvtDesc := admi004.Max1000Text(msg.data.EventDescription)
		EvtInf.EvtDesc = &EvtDesc
	}
	if !isEmpty(msg.data.EventTime) {
		EvtInf.EvtTm = fedwire.ISODateTime(msg.data.EventTime)
	}
	if !isEmpty(EvtInf) {
		SysEvtNtfctn.EvtInf = EvtInf
	}
	if !isEmpty(SysEvtNtfctn) {
		msg.doc.SysEvtNtfctn = SysEvtNtfctn
	}
}
