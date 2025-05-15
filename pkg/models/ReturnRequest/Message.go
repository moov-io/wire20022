package ReturnRequest

import (
	"encoding/xml"
	"fmt"
	"strings"
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
	Data   MessageModel
	Doc    camt056.Document
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
  - With XML path: Loads file, parses XML into message.Doc
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
	if isEmpty(msg.Data.AssignmentId) {
		ParamNames = append(ParamNames, "AssignmentId")
	}
	if isEmpty(msg.Data.Assigner) {
		ParamNames = append(ParamNames, "Assigner")
	}
	if isEmpty(msg.Data.Assignee) {
		ParamNames = append(ParamNames, "Assignee")
	}
	if msg.Data.AssignmentCreateTime.IsZero() {
		ParamNames = append(ParamNames, "AssignmentCreateTime")
	}
	if msg.Data.CaseId == "" {
		ParamNames = append(ParamNames, "CaseId")
	}
	if isEmpty(msg.Data.Creator) {
		ParamNames = append(ParamNames, "Creator")
	}
	if msg.Data.OriginalMessageId == "" {
		ParamNames = append(ParamNames, "OriginalMessageId")
	}
	if msg.Data.OriginalMessageNameId == "" {
		ParamNames = append(ParamNames, "OriginalMessageNameId")
	}
	if msg.Data.OriginalMessageCreateTime.IsZero() {
		ParamNames = append(ParamNames, "OriginalMessageCreateTime")
	}
	if msg.Data.OriginalUETR == "" {
		ParamNames = append(ParamNames, "OriginalUETR")
	}
	if isEmpty(msg.Data.OriginalInterbankSettlementAmount) {
		ParamNames = append(ParamNames, "OriginalInterbankSettlementAmount")
	}
	if isEmpty(msg.Data.OriginalInterbankSettlementDate) {
		ParamNames = append(ParamNames, "OriginalInterbankSettlementDate")
	}
	if isEmpty(msg.Data.CancellationReason) {
		ParamNames = append(ParamNames, "CancellationReason")
	} else if msg.Data.CancellationReason.Reason == "" {
		ParamNames = append(ParamNames, "CancellationReason.Reason")
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
	msg.Doc = camt056.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var FIToFIPmtCxlReq camt056.FIToFIPaymentCancellationRequestV08
	var Assgnmt camt056.CaseAssignment51
	if msg.Data.AssignmentId != "" {
		err := camt056.IMADFedwireFunds1(msg.Data.AssignmentId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AssignmentId",
				Message:   err.Error(),
			}
		}
		Assgnmt.Id = camt056.IMADFedwireFunds1(msg.Data.AssignmentId)
	}
	if !isEmpty(msg.Data.Assigner) {
		Assgnr, err := Party40Choice1From(msg.Data.Assigner)
		if err != nil {
			err.InsertPath("Assigner")
			return err
		}
		Assgnmt.Assgnr = Assgnr
	}
	if !isEmpty(msg.Data.Assignee) {
		Assgne, err := Party40Choice1From(msg.Data.Assignee)
		if err != nil {
			err.InsertPath("Assgne")
			return err
		}
		Assgnmt.Assgne = Assgne
	}
	if !isEmpty(msg.Data.AssignmentCreateTime) {
		err := fedwire.ISODateTime(msg.Data.AssignmentCreateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AssignmentCreateTime",
				Message:   err.Error(),
			}
		}
		Assgnmt.CreDtTm = fedwire.ISODateTime(msg.Data.AssignmentCreateTime)
	}
	if !isEmpty(Assgnmt) {
		FIToFIPmtCxlReq.Assgnmt = Assgnmt
	}
	var Case camt056.Case51
	if msg.Data.CaseId != "" {
		err := camt056.Max35Text(msg.Data.CaseId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CaseId",
				Message:   err.Error(),
			}
		}
		Case.Id = camt056.Max35Text(msg.Data.CaseId)
	}
	if !isEmpty(msg.Data.Creator) {
		Cretr, err := Party40Choice2From(msg.Data.Creator)
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
	if msg.Data.OriginalMessageId != "" {
		err := camt056.Max35Text(msg.Data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgId = camt056.Max35Text(msg.Data.OriginalMessageId)
	}
	if msg.Data.OriginalMessageNameId != "" {
		err := camt056.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgNmId = camt056.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId)
	}
	if !isEmpty(msg.Data.OriginalMessageCreateTime) {
		err := fedwire.ISODateTime(msg.Data.OriginalMessageCreateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageCreateTime",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(msg.Data.OriginalMessageCreateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		TxInf.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.Data.OriginalInstructionId != "" {
		err := camt056.Max35Text(msg.Data.OriginalInstructionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInstructionId",
				Message:   err.Error(),
			}
		}
		OrgnlInstrId := camt056.Max35Text(msg.Data.OriginalInstructionId)
		TxInf.OrgnlInstrId = &OrgnlInstrId
	}
	if msg.Data.OriginalEndToEndId != "" {
		err := camt056.Max35Text(msg.Data.OriginalEndToEndId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalEndToEndId",
				Message:   err.Error(),
			}
		}
		OrgnlEndToEndId := camt056.Max35Text(msg.Data.OriginalEndToEndId)
		TxInf.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if msg.Data.OriginalUETR != "" {
		err := camt056.UUIDv4Identifier(msg.Data.OriginalUETR).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalUETR",
				Message:   err.Error(),
			}
		}
		TxInf.OrgnlUETR = camt056.UUIDv4Identifier(msg.Data.OriginalUETR)
	}
	if !isEmpty(msg.Data.OriginalInterbankSettlementAmount) {
		err := fedwire.Amount(msg.Data.OriginalInterbankSettlementAmount.Amount).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInterbankSettlementAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = camt056.ActiveOrHistoricCurrencyCode(msg.Data.OriginalInterbankSettlementAmount.Currency).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInterbankSettlementAmount.Currency",
				Message:   err.Error(),
			}
		}
		TxInf.OrgnlIntrBkSttlmAmt = camt056.ActiveOrHistoricCurrencyAndAmount{
			Value: camt056.ActiveOrHistoricCurrencyAndAmountSimpleType(msg.Data.OriginalInterbankSettlementAmount.Amount),
			Ccy:   camt056.ActiveOrHistoricCurrencyCode(msg.Data.OriginalInterbankSettlementAmount.Currency),
		}
	}
	if !isEmpty(msg.Data.OriginalInterbankSettlementDate) {
		err := msg.Data.OriginalInterbankSettlementDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInterbankSettlementDate",
				Message:   err.Error(),
			}
		}
		TxInf.OrgnlIntrBkSttlmDt = msg.Data.OriginalInterbankSettlementDate.Date()
	}
	if !isEmpty(msg.Data.CancellationReason) {
		CxlRsnInf, err := PaymentCancellationReason51From(msg.Data.CancellationReason)
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
		msg.Doc.FIToFIPmtCxlReq = FIToFIPmtCxlReq
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Assgnmt) {
		if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Assgnmt.Id) {
			msg.Data.AssignmentId = string(msg.Doc.FIToFIPmtCxlReq.Assgnmt.Id)
		}
		if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Assgnmt.Assgnr) {
			Assgnr := Party40Choice1To(msg.Doc.FIToFIPmtCxlReq.Assgnmt.Assgnr)
			msg.Data.Assigner = Assgnr
		}
		if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Assgnmt.Assgne) {
			Assgne := Party40Choice1To(msg.Doc.FIToFIPmtCxlReq.Assgnmt.Assgne)
			msg.Data.Assignee = Assgne
		}
		if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Assgnmt.CreDtTm) {
			msg.Data.AssignmentCreateTime = time.Time(msg.Doc.FIToFIPmtCxlReq.Assgnmt.CreDtTm)
		}
	}
	if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Case) {
		if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Case.Id) {
			msg.Data.CaseId = string(msg.Doc.FIToFIPmtCxlReq.Case.Id)
		}
		if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Case.Cretr) {
			Cretr := Party40Choice2To(msg.Doc.FIToFIPmtCxlReq.Case.Cretr)
			msg.Data.Creator = Cretr
		}
	}
	if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg) {
		if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf) {
			if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf) {
				if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlMsgId) {
					msg.Data.OriginalMessageId = string(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlMsgId)
				}
				if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlMsgNmId) {
					msg.Data.OriginalMessageNameId = string(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlMsgNmId)
				}
				if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlCreDtTm) {
					msg.Data.OriginalMessageCreateTime = time.Time(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlGrpInf.OrgnlCreDtTm)
				}
			}
			if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlInstrId) {
				msg.Data.OriginalInstructionId = string(*msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlInstrId)
			}
			if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlEndToEndId) {
				msg.Data.OriginalEndToEndId = string(*msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlEndToEndId)
			}
			if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlUETR) {
				msg.Data.OriginalUETR = string(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlUETR)
			}
			if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlIntrBkSttlmAmt) {
				msg.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
					Amount:   float64(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlIntrBkSttlmAmt.Value),
					Currency: string(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlIntrBkSttlmAmt.Ccy),
				}
			}
			if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlIntrBkSttlmDt) {
				msg.Data.OriginalInterbankSettlementDate = model.FromDate(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.OrgnlIntrBkSttlmDt)
			}
			if !isEmpty(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf) {
				CxlRsnInf := PaymentCancellationReason51To(msg.Doc.FIToFIPmtCxlReq.Undrlyg.TxInf.CxlRsnInf)
				msg.Data.CancellationReason = CxlRsnInf
			}
		}
	}

	return nil
}
