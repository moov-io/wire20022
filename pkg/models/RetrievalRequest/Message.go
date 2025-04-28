package RetrievalRequest

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	admi006 "github.com/moov-io/fedwire20022/gen/RetrievalRequest_admi_006_001_01"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:admi.006.001.01"

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

/*
NewMessage creates a new Message instance with optional XML initialization.

Parameters:
  - filepath: File path to XML (optional)
    If provided, loads and parses XML from specified path

Returns:
  - Message: Initialized message structure
  - error: File read or XML parsing errors (if XML path provided)

Behavior:
  - Without arguments: Returns empty Message with default MessageModel
  - With XML path: Loads file, parses XML into message.doc
*/
func NewMessage(filepath string) (Message, error) {
	msg := Message{data: MessageModel{}} // Initialize with zero value

	if filepath == "" {
		return msg, nil // Return early for empty filepath
	}

	// Read and validate file
	data, err := model.ReadXMLFile(filepath)
	if err != nil {
		return msg, fmt.Errorf("file read error: %w", err)
	}

	// Handle empty XML data
	if len(data) == 0 {
		return msg, fmt.Errorf("empty XML file: %s", filepath)
	}

	// Parse XML with structural validation
	if err := xml.Unmarshal(data, &msg.doc); err != nil {
		return msg, fmt.Errorf("XML parse error: %w", err)
	}

	return msg, nil
}
func (msg *Message) CreateDocument() *model.ValidateError {
	msg.doc = admi006.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var RsndReq admi006.ResendRequestV01
	var MsgHdr admi006.MessageHeader71
	if msg.data.MessageId != "" {
		err := admi006.Max35Text(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		MsgHdr.MsgId = admi006.Max35Text(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		MsgHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
	if msg.data.RequestType != "" {
		err := admi006.TrafficTypeFedwireFunds1(msg.data.RequestType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
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
		err := msg.data.BusinessDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BusinessDate",
				Message:   err.Error(),
			}
		}
		RsndSchCrit.BizDt = msg.data.BusinessDate.Date()
	}
	if !isEmpty(msg.data.SequenceRange) {
		var FrToSeq []admi006.SequenceRange11
		FrSeq, err := strconv.ParseFloat(msg.data.SequenceRange.FromSeq, 64)
		if err != nil {
			return  &model.ValidateError{
				ParamName: "SequenceRange.FromSeq",
				Message:   err.Error(),
			}
		}
		ToSeq, err := strconv.ParseFloat(msg.data.SequenceRange.ToSeq, 64)
		if err != nil {
			return &model.ValidateError{
				ParamName: "SequenceRange.ToSeq",
				Message:   err.Error(),
			}
		}
		seqrange := admi006.SequenceRange11{
			FrSeq: admi006.XSequenceNumberFedwireFunds1(FrSeq),
			ToSeq: admi006.XSequenceNumberFedwireFunds1(ToSeq),
		}
		FrToSeq = append(FrToSeq, seqrange)
		SeqRg := admi006.SequenceRange1Choice1{
			FrToSeq: FrToSeq,
		}
		RsndSchCrit.SeqRg = &SeqRg
	}
	if msg.data.OriginalMessageNameId != "" {
		err := admi006.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlMsgNmId := admi006.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
		RsndSchCrit.OrgnlMsgNmId = &OrgnlMsgNmId
	}
	if msg.data.FileReference != "" {
		err := admi006.IMADOrOMADFedwireFunds1(msg.data.FileReference).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "FileReference",
				Message:   err.Error(),
			}
		}
		FileRef := admi006.IMADOrOMADFedwireFunds1(msg.data.FileReference)
		RsndSchCrit.FileRef = &FileRef
	}
	var Rcpt admi006.PartyIdentification1361
	var Id admi006.PartyIdentification120Choice1
	var PrtryId admi006.GenericIdentification361
	if msg.data.RecipientId != "" {
		err := admi006.EndpointIdentifierFedwireFunds1(msg.data.RecipientId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RecipientId",
				Message:   err.Error(),
			}
		}
		PrtryId.Id = admi006.EndpointIdentifierFedwireFunds1(msg.data.RecipientId)
	}
	if msg.data.RecipientIssuer != "" {
		err := admi006.Max35TextFixed(msg.data.RecipientIssuer).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RecipientIssuer",
				Message:   err.Error(),
			}
		}
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
	return nil
}
