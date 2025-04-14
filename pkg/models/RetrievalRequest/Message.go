package RetrievalRequest

import (
	"time"

	admi006 "github.com/moov-io/fedwire20022/gen/RetrievalRequest_admi_006_001_01"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageModel struct {
	//Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	MessageId string
	//Date and time at which the message was created.
	CreatedDateTime time.Time
	//Specific actions to be executed through the request.
	RequestType model.RequestType
	//Date of the business day of the requested messages the resend function is used for.
	BusinessDate model.Date
	//Independent counter for a range of message sequences, which are available once per party technical address.
	SequenceRange model.SequenceRange
	//Unambiguously identifies the original bsiness message, which was delivered by the business sender.
	OriginalMessageNameId string
	//String of characters that uniquely identifies the file, which was delivered by the sender.
	FileReference string
	//Unique identification of the party.
	RecipientId string
	//Entity that assigns the identification.
	RecipientIssuer string
}
type Message struct {
	data MessageModel
	doc  admi006.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = admi006.Document{}
	var RsndReq admi006.ResendRequestV01
	var MsgHdr admi006.MessageHeader71
	if msg.data.MessageId != "" {
		MsgHdr.MsgId = admi006.Max35Text(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		MsgHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
	if msg.data.RequestType != "" {
		Prtry := admi006.GenericIdentification11{
			Id: admi006.TrafficTypeFedwireFunds1(msg.data.RequestType),
		}
		MsgHdr.ReqTp = admi006.RequestType4Choice1{
			Prtry: &Prtry,
		}
	}
	if !isEmpty(MsgHdr) {
		RsndReq.MsgHdr = MsgHdr
	}
	var RsndSchCrit admi006.ResendSearchCriteria21
	if !isEmpty(msg.data.BusinessDate) {
		RsndSchCrit.BizDt = msg.data.BusinessDate.Date()
	}
	if !isEmpty(msg.data.SequenceRange) {
		var FrToSeq []admi006.SequenceRange11
		seqrange := admi006.SequenceRange11{
			FrSeq: admi006.XSequenceNumberFedwireFunds1(msg.data.SequenceRange.FromSeq),
			ToSeq: admi006.XSequenceNumberFedwireFunds1(msg.data.SequenceRange.ToSeq),
		}
		FrToSeq = append(FrToSeq, seqrange)
		SeqRg := admi006.SequenceRange1Choice1{
			FrToSeq: FrToSeq,
		}
		RsndSchCrit.SeqRg = &SeqRg
	}
	if msg.data.OriginalMessageNameId != "" {
		OrgnlMsgNmId := admi006.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
		RsndSchCrit.OrgnlMsgNmId = &OrgnlMsgNmId
	}
	if msg.data.FileReference != "" {
		FileRef := admi006.IMADOrOMADFedwireFunds1(msg.data.FileReference)
		RsndSchCrit.FileRef = &FileRef
	}
	var Rcpt admi006.PartyIdentification1361
	var Id admi006.PartyIdentification120Choice1
	var PrtryId admi006.GenericIdentification361
	if msg.data.RecipientId != "" {
		PrtryId.Id = admi006.EndpointIdentifierFedwireFunds1(msg.data.RecipientId)
	}
	if msg.data.RecipientIssuer != "" {
		PrtryId.Issr = admi006.Max35TextFixed(msg.data.RecipientIssuer)
	}
	if !isEmpty(PrtryId) {
		Id.PrtryId = &PrtryId
	}
	if !isEmpty(Id) {
		Rcpt.Id = Id
	}
	if !isEmpty(Rcpt) {
		RsndSchCrit.Rcpt = Rcpt
	}
	if !isEmpty(RsndSchCrit) {
		RsndReq.RsndSchCrit = RsndSchCrit
	}
	if !isEmpty(RsndReq) {
		msg.doc.RsndReq = RsndReq
	}
}
