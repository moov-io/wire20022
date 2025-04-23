package FedwireFundsAcknowledgement

import (
	"encoding/xml"
	"time"

	admi007 "github.com/moov-io/fedwire20022/gen/FedwireFundsAcknowledgement_admi_007_001_01"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:admi.007.001.01"

type MessageModel struct {
	//Specifies the identification the message.
	MessageId string
	//Date and time at which the message was created.
	CreatedDateTime time.Time
	//Unambiguous reference to a previous message having a business relevance with this message.
	RelationReference string
	//Name of the message which contained the given additional reference as its message reference.
	ReferenceName string
	//Gives the status of the request.
	RequestHandling model.RelatedStatusCode
}
type Message struct {
	data MessageModel
	doc  admi007.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = admi007.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var RctAck admi007.ReceiptAcknowledgementV01
	if msg.data.MessageId != "" {
		RctAck.MsgId.MsgId = admi007.OMADFedwireFunds1(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		RctAck.MsgId.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
	if msg.data.RelationReference != "" {
		RctAck.Rpt.RltdRef.Ref = admi007.Max35Text(msg.data.RelationReference)
	}
	if msg.data.ReferenceName != "" {
		RctAck.Rpt.RltdRef.MsgNm = admi007.MessageNameIdentificationFRS1(msg.data.ReferenceName)
	}
	if msg.data.RequestHandling != "" {
		RctAck.Rpt.ReqHdlg.StsCd = admi007.Max4AlphaNumericTextFixed(msg.data.RequestHandling)
	}
	if !isEmpty(RctAck) {
		msg.doc.RctAck = RctAck
	}
}
