package RetrievalRequest

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
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
	Data   MessageModel
	Doc    admi006.Document
	Helper MessageHelper
}

func (msg *Message) GetDataModel() interface{} {
	return &msg.Data
}
func (msg *Message) GetDocument() interface{} {
	return &msg.Doc
}
func (msg *Message) GetHelper() interface{} {
	return &msg.Helper
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
func NewMessage(filepath string) (*Message, error) {
	msg := Message{Data: MessageModel{}} // Initialize with zero value
	msg.Helper = BuildMessageHelper()

	if filepath == "" {
		return &msg, nil // Return early for empty filepath
	}

	// Read and validate file
	data, err := model.ReadXMLFile(filepath)
	if err != nil {
		return &msg, fmt.Errorf("file read error: %w", err)
	}

	// Handle empty XML data
	if len(data) == 0 {
		return &msg, fmt.Errorf("empty XML file: %s", filepath)
	}

	// Parse XML with structural validation
	if err := xml.Unmarshal(data, &msg.Doc); err != nil {
		return &msg, fmt.Errorf("XML parse error: %w", err)
	}

	return &msg, nil
}

func (msg *Message) ValidateRequiredFields() *model.ValidateError {
	// Initialize the RequireError object
	var ParamNames []string
	// Check required fields and append missing ones to ParamNames
	if msg.Data.MessageId == "" {
		ParamNames = append(ParamNames, "MessageId")
	}
	if msg.Data.CreatedDateTime.IsZero() {
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if msg.Data.RequestType == "" {
		ParamNames = append(ParamNames, "RequestType")
	}
	if isEmpty(msg.Data.BusinessDate) {
		ParamNames = append(ParamNames, "BusinessDate")
	}
	if msg.Data.RecipientId == "" {
		ParamNames = append(ParamNames, "RecipientId")
	}
	// Return nil if no required fields are missing
	if len(ParamNames) == 0 {
		return nil
	}
	return &model.ValidateError{
		ParamName: "RequiredFields",
		Message:   strings.Join(ParamNames, ", "),
	}
}

func (msg *Message) CreateDocument() *model.ValidateError {
	requireErr := msg.ValidateRequiredFields()
	if requireErr != nil {
		return requireErr
	}
	msg.Doc = admi006.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var RsndReq admi006.ResendRequestV01
	var MsgHdr admi006.MessageHeader71
	if msg.Data.MessageId != "" {
		err := admi006.Max35Text(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		MsgHdr.MsgId = admi006.Max35Text(msg.Data.MessageId)
	}
	if !isEmpty(msg.Data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.Data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		MsgHdr.CreDtTm = fedwire.ISODateTime(msg.Data.CreatedDateTime)
	}
	if msg.Data.RequestType != "" {
		err := admi006.TrafficTypeFedwireFunds1(msg.Data.RequestType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		Prtry := admi006.GenericIdentification11{
			Id: admi006.TrafficTypeFedwireFunds1(msg.Data.RequestType),
		}
		MsgHdr.ReqTp = admi006.RequestType4Choice1{
			Prtry: &Prtry,
		}
	}
	if !isEmpty(MsgHdr) {
		RsndReq.MsgHdr = MsgHdr
	}
	var RsndSchCrit admi006.ResendSearchCriteria21
	if !isEmpty(msg.Data.BusinessDate) {
		err := msg.Data.BusinessDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BusinessDate",
				Message:   err.Error(),
			}
		}
		RsndSchCrit.BizDt = msg.Data.BusinessDate.Date()
	}
	if !isEmpty(msg.Data.SequenceRange) {
		var FrToSeq []admi006.SequenceRange11
		FrSeq, err := strconv.ParseFloat(msg.Data.SequenceRange.FromSeq, 64)
		if err != nil {
			return &model.ValidateError{
				ParamName: "SequenceRange.FromSeq",
				Message:   err.Error(),
			}
		}
		ToSeq, err := strconv.ParseFloat(msg.Data.SequenceRange.ToSeq, 64)
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
	if msg.Data.OriginalMessageNameId != "" {
		err := admi006.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlMsgNmId := admi006.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId)
		RsndSchCrit.OrgnlMsgNmId = &OrgnlMsgNmId
	}
	if msg.Data.FileReference != "" {
		err := admi006.IMADOrOMADFedwireFunds1(msg.Data.FileReference).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "FileReference",
				Message:   err.Error(),
			}
		}
		FileRef := admi006.IMADOrOMADFedwireFunds1(msg.Data.FileReference)
		RsndSchCrit.FileRef = &FileRef
	}
	var Rcpt admi006.PartyIdentification1361
	var Id admi006.PartyIdentification120Choice1
	var PrtryId admi006.GenericIdentification361
	if msg.Data.RecipientId != "" {
		err := admi006.EndpointIdentifierFedwireFunds1(msg.Data.RecipientId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RecipientId",
				Message:   err.Error(),
			}
		}
		PrtryId.Id = admi006.EndpointIdentifierFedwireFunds1(msg.Data.RecipientId)
	}
	if msg.Data.RecipientIssuer != "" {
		err := admi006.Max35TextFixed(msg.Data.RecipientIssuer).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RecipientIssuer",
				Message:   err.Error(),
			}
		}
		PrtryId.Issr = admi006.Max35TextFixed(msg.Data.RecipientIssuer)
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
		msg.Doc.RsndReq = RsndReq
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.RsndReq) {
		if !isEmpty(msg.Doc.RsndReq.MsgHdr) {
			if !isEmpty(msg.Doc.RsndReq.MsgHdr.MsgId) {
				msg.Data.MessageId = string(msg.Doc.RsndReq.MsgHdr.MsgId)
			}
			if !isEmpty(msg.Doc.RsndReq.MsgHdr.CreDtTm) {
				msg.Data.CreatedDateTime = time.Time(msg.Doc.RsndReq.MsgHdr.CreDtTm)
			}
			if !isEmpty(msg.Doc.RsndReq.MsgHdr.ReqTp) {
				if !isEmpty(msg.Doc.RsndReq.MsgHdr.ReqTp.Prtry) {
					msg.Data.RequestType = model.RequestType(msg.Doc.RsndReq.MsgHdr.ReqTp.Prtry.Id)
				}
			}
		}
		if !isEmpty(msg.Doc.RsndReq.RsndSchCrit) {
			if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.BizDt) {
				msg.Data.BusinessDate = model.FromDate(msg.Doc.RsndReq.RsndSchCrit.BizDt)
			}
			if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.SeqRg) {
				if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.SeqRg.FrToSeq) {
					if len(msg.Doc.RsndReq.RsndSchCrit.SeqRg.FrToSeq) > 0 {
						msg.Data.SequenceRange.FromSeq = strconv.FormatFloat(float64(msg.Doc.RsndReq.RsndSchCrit.SeqRg.FrToSeq[0].FrSeq), 'f', -1, 64)
						msg.Data.SequenceRange.ToSeq = strconv.FormatFloat(float64(msg.Doc.RsndReq.RsndSchCrit.SeqRg.FrToSeq[0].ToSeq), 'f', -1, 64)
					}
				}
			}
			if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.OrgnlMsgNmId) {
				msg.Data.OriginalMessageNameId = string(*msg.Doc.RsndReq.RsndSchCrit.OrgnlMsgNmId)
			}
			if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.FileRef) {
				msg.Data.FileReference = string(*msg.Doc.RsndReq.RsndSchCrit.FileRef)
			}
			if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.Rcpt) {
				if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.Rcpt.Id) {
					if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.Rcpt.Id.PrtryId) {
						if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.Rcpt.Id.PrtryId.Id) {
							msg.Data.RecipientId = string(msg.Doc.RsndReq.RsndSchCrit.Rcpt.Id.PrtryId.Id)
						}
						if !isEmpty(msg.Doc.RsndReq.RsndSchCrit.Rcpt.Id.PrtryId.Issr) {
							msg.Data.RecipientIssuer = string(msg.Doc.RsndReq.RsndSchCrit.Rcpt.Id.PrtryId.Issr)
						}
					}
				}
			}
		}
	}
	return nil
}
