package ReturnRequest

import (
	"encoding/xml"
	"fmt"
	"time"

	camt056 "github.com/moov-io/fedwire20022/gen/ReturnRequest_camt_056_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.056.001.08"

type MessageModel struct {
	//Uniquely identifies the case assignment.
	AssignmentId string
	//Party who assigns the case.
	Assigner model.Agent
	//Party to which the case is assigned.
	Assignee model.Agent
	//Date and time at which the assignment was created.
	AssignmentCreateTime time.Time
	//Identifies the investigation case.
	CaseId string
	//Party that created the investigation case.
	Creator model.Agent
	//Point to point reference assigned by the original instructing party to unambiguously identify the original message.
	OriginalMessageId string
	//Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.
	OriginalMessageNameId string
	//Original date and time at which the message was created.
	OriginalMessageCreateTime time.Time
	//Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId string
	//Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId string
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUETR string
	//Amount of money moved between the instructing agent and the instructed agent, as provided in the original instruction.
	OriginalInterbankSettlementAmount model.CurrencyAndAmount
	//Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate model.Date
	//Provides detailed information on the cancellation reason.
	CancellationReason Reason
}
type Message struct {
	data MessageModel
	doc  camt056.Document
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
	msg.doc = camt056.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var FIToFIPmtCxlReq camt056.FIToFIPaymentCancellationRequestV08
	var Assgnmt camt056.CaseAssignment51
	if msg.data.AssignmentId != "" {
		err := camt056.IMADFedwireFunds1(msg.data.AssignmentId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AssignmentId",
				Message:   err.Error(),
			}
		}
		Assgnmt.Id = camt056.IMADFedwireFunds1(msg.data.AssignmentId)
	}
	if !isEmpty(msg.data.Assigner) {
		Assgnr, err := Party40Choice1From(msg.data.Assigner)
		if err != nil {
			err.InsertPath("Assigner")
			return err
		}
		Assgnmt.Assgnr = Assgnr
	}
	if !isEmpty(msg.data.Assignee) {
		Assgne, err := Party40Choice1From(msg.data.Assignee)
		if err != nil {
			err.InsertPath("Assgne")
			return err
		}
		Assgnmt.Assgne = Assgne
	}
	if !isEmpty(msg.data.AssignmentCreateTime) {
		err := fedwire.ISODateTime(msg.data.AssignmentCreateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AssignmentCreateTime",
				Message:   err.Error(),
			}
		}
		Assgnmt.CreDtTm = fedwire.ISODateTime(msg.data.AssignmentCreateTime)
	}
	if !isEmpty(Assgnmt) {
		FIToFIPmtCxlReq.Assgnmt = Assgnmt
	}
	var Case camt056.Case51
	if msg.data.CaseId != "" {
		err := camt056.Max35Text(msg.data.CaseId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CaseId",
				Message:   err.Error(),
			}
		}
		Case.Id = camt056.Max35Text(msg.data.CaseId)
	}
	if !isEmpty(msg.data.Creator) {
		Cretr, err := Party40Choice2From(msg.data.Creator)
		if err != nil {
			err.InsertPath("Creator")
			return err
		}
		Case.Cretr = Cretr
	}
	if !isEmpty(Case) {
		FIToFIPmtCxlReq.Case = Case
	}
	var Undrlyg camt056.UnderlyingTransaction231
	var TxInf camt056.PaymentTransaction1061
	var OrgnlGrpInf camt056.OriginalGroupInformation291
	if msg.data.OriginalMessageId != "" {
		err := camt056.Max35Text(msg.data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgId = camt056.Max35Text(msg.data.OriginalMessageId)
	}
	if msg.data.OriginalMessageNameId != "" {
		err := camt056.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgNmId = camt056.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
	}
	if !isEmpty(msg.data.OriginalMessageCreateTime) {
		err := fedwire.ISODateTime(msg.data.OriginalMessageCreateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageCreateTime",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(msg.data.OriginalMessageCreateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		TxInf.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.data.OriginalInstructionId != "" {
		err := camt056.Max35Text(msg.data.OriginalInstructionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInstructionId",
				Message:   err.Error(),
			}
		}
		OrgnlInstrId := camt056.Max35Text(msg.data.OriginalInstructionId)
		TxInf.OrgnlInstrId = &OrgnlInstrId
	}
	if msg.data.OriginalEndToEndId != "" {
		err := camt056.Max35Text(msg.data.OriginalEndToEndId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalEndToEndId",
				Message:   err.Error(),
			}
		}
		OrgnlEndToEndId := camt056.Max35Text(msg.data.OriginalEndToEndId)
		TxInf.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if msg.data.OriginalUETR != "" {
		err := camt056.UUIDv4Identifier(msg.data.OriginalUETR).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalUETR",
				Message:   err.Error(),
			}
		}
		TxInf.OrgnlUETR = camt056.UUIDv4Identifier(msg.data.OriginalUETR)
	}
	if !isEmpty(msg.data.OriginalInterbankSettlementAmount) {
		err := fedwire.Amount(msg.data.OriginalInterbankSettlementAmount.Amount).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInterbankSettlementAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = camt056.ActiveOrHistoricCurrencyCode(msg.data.OriginalInterbankSettlementAmount.Currency).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInterbankSettlementAmount.Currency",
				Message:   err.Error(),
			}
		}
		TxInf.OrgnlIntrBkSttlmAmt = camt056.ActiveOrHistoricCurrencyAndAmount{
			Value: camt056.ActiveOrHistoricCurrencyAndAmountSimpleType(msg.data.OriginalInterbankSettlementAmount.Amount),
			Ccy:   camt056.ActiveOrHistoricCurrencyCode(msg.data.OriginalInterbankSettlementAmount.Currency),
		}
	}
	if !isEmpty(msg.data.OriginalInterbankSettlementDate) {
		err := msg.data.OriginalInterbankSettlementDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInterbankSettlementDate",
				Message:   err.Error(),
			}
		}
		TxInf.OrgnlIntrBkSttlmDt = msg.data.OriginalInterbankSettlementDate.Date()
	}
	if !isEmpty(msg.data.CancellationReason) {
		CxlRsnInf, err := PaymentCancellationReason51From(msg.data.CancellationReason)
		if err != nil {
			err.InsertPath("CancellationReason")
			return err
		}
		TxInf.CxlRsnInf = CxlRsnInf
	}
	if !isEmpty(TxInf) {
		Undrlyg.TxInf = TxInf
	}
	if !isEmpty(Undrlyg) {
		FIToFIPmtCxlReq.Undrlyg = Undrlyg
	}
	if !isEmpty(FIToFIPmtCxlReq) {
		msg.doc.FIToFIPmtCxlReq = FIToFIPmtCxlReq
	}
	return nil
}
