package PaymentStatusRequest

import (
	"encoding/xml"
	"fmt"
	"strings"
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
	Data   MessageModel
	Doc    pacs004.Document
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
	if msg.Data.OriginalMessageId == "" {
		ParamNames = append(ParamNames, "OriginalMessageId")
	}
	if msg.Data.OriginalMessageNameId == "" {
		ParamNames = append(ParamNames, "OriginalMessageNameId")
	}
	if msg.Data.OriginalCreationDateTime.IsZero() {
		ParamNames = append(ParamNames, "OriginalCreationDateTime")
	}
	if msg.Data.OriginalUETR == "" {
		ParamNames = append(ParamNames, "OriginalUETR")
	}
	if isEmpty(msg.Data.InstructingAgent) {
		ParamNames = append(ParamNames, "InstructingAgent")
	}
	if isEmpty(msg.Data.InstructedAgent) {
		ParamNames = append(ParamNames, "InstructedAgent")
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
	msg.Doc = pacs004.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var FIToFIPmtStsReq pacs004.FIToFIPaymentStatusRequestV03
	var GrpHdr pacs004.GroupHeader911
	if msg.Data.MessageId != "" {
		err := pacs004.Max35Text(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = pacs004.Max35Text(msg.Data.MessageId)
	}
	if !isEmpty(msg.Data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.Data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.Data.CreatedDateTime)
	}
	if !isEmpty(GrpHdr) {
		FIToFIPmtStsReq.GrpHdr = GrpHdr
	}
	var TxInf pacs004.PaymentTransaction1131
	var OrgnlGrpInf pacs004.OriginalGroupInformation291
	if msg.Data.OriginalMessageId != "" {
		err := pacs004.IMADFedwireFunds1(msg.Data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgId = pacs004.IMADFedwireFunds1(msg.Data.OriginalMessageId)
	}
	if msg.Data.OriginalMessageNameId != "" {
		err := pacs004.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgNmId = pacs004.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId)
	}
	if !isEmpty(msg.Data.OriginalCreationDateTime) {
		err := fedwire.ISODateTime(msg.Data.OriginalCreationDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalCreationDateTime",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(msg.Data.OriginalCreationDateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		TxInf.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.Data.OriginalInstructionId != "" {
		err := pacs004.Max35Text(msg.Data.OriginalInstructionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInstructionId",
				Message:   err.Error(),
			}
		}
		OrgnlInstrId := pacs004.Max35Text(msg.Data.OriginalInstructionId)
		TxInf.OrgnlInstrId = &OrgnlInstrId
	}
	if msg.Data.OriginalEndToEndId != "" {
		err := pacs004.Max35Text(msg.Data.OriginalEndToEndId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalEndToEndId",
				Message:   err.Error(),
			}
		}
		OrgnlEndToEndId := pacs004.Max35Text(msg.Data.OriginalEndToEndId)
		TxInf.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if msg.Data.OriginalUETR != "" {
		err := pacs004.UUIDv4Identifier(msg.Data.OriginalUETR).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalUETR",
				Message:   err.Error(),
			}
		}
		TxInf.OrgnlUETR = pacs004.UUIDv4Identifier(msg.Data.OriginalUETR)
	}
	if !isEmpty(msg.Data.InstructingAgent) {
		InstgAgt, err := BranchAndFinancialInstitutionIdentification61From(msg.Data.InstructingAgent)
		if err != nil {
			err.InsertPath("InstructingAgent")
			return err
		}
		TxInf.InstgAgt = InstgAgt
	}
	if !isEmpty(msg.Data.InstructedAgent) {
		InstdAgt, err := BranchAndFinancialInstitutionIdentification61From(msg.Data.InstructedAgent)
		if err != nil {
			err.InsertPath("InstructingAgent")
			return err
		}
		TxInf.InstdAgt = InstdAgt
	}
	if !isEmpty(TxInf) {
		FIToFIPmtStsReq.TxInf = TxInf
	}
	if !isEmpty(FIToFIPmtStsReq) {
		msg.Doc.FIToFIPmtStsReq = FIToFIPmtStsReq
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.FIToFIPmtStsReq.GrpHdr) {
		if !isEmpty(msg.Doc.FIToFIPmtStsReq.GrpHdr.MsgId) {
			msg.Data.MessageId = string(msg.Doc.FIToFIPmtStsReq.GrpHdr.MsgId)
		}
		if !isEmpty(msg.Doc.FIToFIPmtStsReq.GrpHdr.CreDtTm) {
			msg.Data.CreatedDateTime = time.Time(msg.Doc.FIToFIPmtStsReq.GrpHdr.CreDtTm)
		}
	}
	if !isEmpty(msg.Doc.FIToFIPmtStsReq.TxInf) {
		if !isEmpty(msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf) {
			if !isEmpty(msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgId) {
				msg.Data.OriginalMessageId = string(msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgId)
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgNmId) {
				msg.Data.OriginalMessageNameId = string(msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgNmId)
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlCreDtTm) {
				msg.Data.OriginalCreationDateTime = time.Time(msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlCreDtTm)
			}
		}
		if !isEmpty(msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlInstrId) {
			msg.Data.OriginalInstructionId = string(*msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlInstrId)
		}
		if !isEmpty(msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlEndToEndId) {
			msg.Data.OriginalEndToEndId = string(*msg.Doc.FIToFIPmtStsReq.TxInf.OrgnlEndToEndId)
		}
		if !isEmpty(msg.Doc.FIToFIPmtStsReq.TxInf.InstgAgt) {
			msg.Data.InstructingAgent = BranchAndFinancialInstitutionIdentification61To(msg.Doc.FIToFIPmtStsReq.TxInf.InstgAgt)
		}
		if !isEmpty(msg.Doc.FIToFIPmtStsReq.TxInf.InstdAgt) {
			msg.Data.InstructedAgent = BranchAndFinancialInstitutionIdentification61To(msg.Doc.FIToFIPmtStsReq.TxInf.InstdAgt)
		}
	}
	return nil
}
