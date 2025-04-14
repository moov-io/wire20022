package ConnectionCheck

import (
	"encoding/xml"
	"time"

	admi004 "github.com/moov-io/fedwire20022/gen/ConnectionCheck_admi_004_001_02"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02"

type MessageModel struct {
	EventType string
	EvntParam string
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
	var EvtInf admi004.Event21
	if msg.data.EventType != "" {
		EvtInf.EvtCd = admi004.EventFedwireFunds1(msg.data.EventType)
	}
	if msg.data.EvntParam != "" {
		EvtInf.EvtParam = admi004.EndpointIdentifierFedwireFunds1(msg.data.EvntParam)
	}
	if !isEmpty(msg.data.EventTime) {
		EvtInf.EvtTm = fedwire.ISODateTime(msg.data.EventTime)
	}
	if !isEmpty(EvtInf) {
		msg.doc.SysEvtNtfctn = admi004.SystemEventNotificationV02{
			EvtInf: EvtInf,
		}
	}
}
