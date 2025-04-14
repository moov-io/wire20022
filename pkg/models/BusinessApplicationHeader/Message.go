package BusinessApplicationHeader

import (
	"encoding/xml"
	"time"

	head001 "github.com/moov-io/fedwire20022/gen/BusinessApplicationHeader_head_001_001_03"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:head.001.001.03"

type MessageModel struct {
	MessageSenderId   string
	MessageReceiverId string
	//BizMsgIdr (Business Message Identifier) is a unique identifier assigned to an ISO 20022 message to distinguish it from other messages.
	BusinessMessageId string
	//MsgDefIdr (Message Definition Identifier) and BizSvc (Business Service) are part of the Business Application Header (BAH), which helps identify and process financial messages.
	MessageDefinitionId string
	//BizSvc specifies a business service or process related to the message.
	BusinessService string
	//<MktPrctc> (Market Practice) is used to specify market-specific rules and guidelines that apply to the message.
	MarketInfo MarketPractice

	CreateDatetime time.Time
	//BizPrcgDt stands for Business Processing Date. It represents the date when a financial transaction or message is processed by a financial institution.
	BusinessProcessingDate time.Time
	//It refers to a related Business Application Header (BAH) of type BusinessApplicationHeader71
	Relations BusinessApplicationHeader
}
type Message struct {
	data MessageModel
	doc  head001.AppHdr
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}

func (msg *Message) CreateDocument() {
	msg.doc = head001.AppHdr{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "AppHdr",
		},
	}
	if msg.data.MessageSenderId != "" {
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(msg.data.MessageSenderId),
				},
			},
		}
		msg.doc.Fr = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if msg.data.MessageReceiverId != "" {
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(msg.data.MessageReceiverId),
				},
			},
		}
		msg.doc.To = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if msg.data.BusinessMessageId != "" {
		msg.doc.BizMsgIdr = head001.Max35Text(msg.data.BusinessMessageId)
	}
	if msg.data.MessageDefinitionId != "" {
		msg.doc.MsgDefIdr = head001.MessageNameIdentificationFRS1(msg.data.MessageDefinitionId)
	}
	if msg.data.BusinessService != "" {
		msg.doc.BizSvc = head001.BusinessServiceFedwireFunds1(msg.data.BusinessService)
	}
	if !isEmpty(msg.data.MarketInfo) {
		MktPrctc := ImplementationSpecification11From(msg.data.MarketInfo)
		if !isEmpty(MktPrctc) {
			msg.doc.MktPrctc = MktPrctc
		}
	}
	if !isEmpty(msg.data.CreateDatetime) {
		msg.doc.CreDt = fedwire.ISODateTime(msg.data.CreateDatetime)
	}
	if !isEmpty(msg.data.BusinessProcessingDate) {
		BizPrcgDt := fedwire.ISODateTime(msg.data.BusinessProcessingDate)
		msg.doc.BizPrcgDt = &BizPrcgDt
	}
	if !isEmpty(msg.data.Relations) {
		Rltd := BusinessApplicationHeader71From(msg.data.Relations)
		if !isEmpty(Rltd) {
			msg.doc.Rltd = &Rltd
		}
	}

}
