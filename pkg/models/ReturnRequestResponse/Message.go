package ReturnRequestResponse

import (
	"encoding/xml"
	"fmt"
	"time"

	camt029 "github.com/moov-io/fedwire20022/gen/ReturnRequestResponse_camt_029_001_09"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09"

type MessageModel struct {
	//Uniquely identifies the case assignment.
	AssignmentId string
	//Party who assigns the case.
	Assigner model.Agent
	//Party to which the case is assigned.
	Assignee model.Agent
	//Date and time at which the assignment was created.
	AssignmentCreateTime time.Time
	//Identifies a resolved case.
	ResolvedCaseId string
	//Party that created the investigation case.
	Creator model.Agent
	//Specifies the status of the investigation, in a coded form.
	Status Status
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
	//Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInfo Reason
}
type Message struct {
	data MessageModel
	doc  camt029.Document
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
	msg.doc = camt029.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}

	var RsltnOfInvstgtn camt029.ResolutionOfInvestigationV09
	var Assgnmt camt029.CaseAssignment51
	if msg.data.AssignmentId != "" {
		err := camt029.Max35Text(msg.data.AssignmentId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AssignmentId",
				Message:   err.Error(),
			}
		}
		Assgnmt.Id = camt029.Max35Text(msg.data.AssignmentId)
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
			err.InsertPath("Assignee")
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
		RsltnOfInvstgtn.Assgnmt = Assgnmt
	}
	var RslvdCase camt029.Case51
	if msg.data.ResolvedCaseId != "" {
		err := camt029.Max35Text(msg.data.ResolvedCaseId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ResolvedCaseId",
				Message:   err.Error(),
			}
		}
		RslvdCase.Id = camt029.Max35Text(msg.data.ResolvedCaseId)
	}
	if !isEmpty(msg.data.Creator) {
		Cretr, err := Party40Choice2From(msg.data.Creator)
		if err != nil {
			err.InsertPath("Creator")
			return err
		}
		RslvdCase.Cretr = Cretr
	}
	if !isEmpty(RslvdCase) {
		RsltnOfInvstgtn.RslvdCase = RslvdCase
	}
	var Sts camt029.InvestigationStatus5Choice1
	if msg.data.Status != "" {
		err := camt029.ExternalInvestigationExecutionConfirmation1Code(msg.data.Status).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Status",
				Message:   err.Error(),
			}
		}
		Conf := camt029.ExternalInvestigationExecutionConfirmation1Code(msg.data.Status)
		Sts.Conf = &Conf
	}
	if !isEmpty(Sts) {
		RsltnOfInvstgtn.Sts = Sts
	}
	var CxlDtls camt029.UnderlyingTransaction221
	var TxInfAndSts camt029.PaymentTransaction1021
	var OrgnlGrpInf camt029.OriginalGroupInformation291
	if msg.data.OriginalMessageId != "" {
		err := camt029.Max35Text(msg.data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgId = camt029.Max35Text(msg.data.OriginalMessageId)
	}
	if msg.data.OriginalMessageNameId != "" {
		err := camt029.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgNmId = camt029.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
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
		TxInfAndSts.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.data.OriginalInstructionId != "" {
		err := camt029.Max35Text(msg.data.OriginalInstructionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInstructionId",
				Message:   err.Error(),
			}
		}
		OrgnlInstrId := camt029.Max35Text(msg.data.OriginalInstructionId)
		TxInfAndSts.OrgnlInstrId = &OrgnlInstrId
	}
	if msg.data.OriginalEndToEndId != "" {
		err := camt029.Max35Text(msg.data.OriginalEndToEndId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalEndToEndId",
				Message:   err.Error(),
			}
		}
		OrgnlEndToEndId := camt029.Max35Text(msg.data.OriginalEndToEndId)
		TxInfAndSts.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if msg.data.OriginalUETR != "" {
		err := camt029.UUIDv4Identifier(msg.data.OriginalUETR).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalUETR",
				Message:   err.Error(),
			}
		}
		TxInfAndSts.OrgnlUETR = camt029.UUIDv4Identifier(msg.data.OriginalUETR)
	}
	if !isEmpty(msg.data.CancellationStatusReasonInfo) {
		var CxlStsRsnInf []*camt029.CancellationStatusReason41
		reason, err := CancellationStatusReason41From(msg.data.CancellationStatusReasonInfo)
		if err != nil {
			err.InsertPath("CancellationStatusReasonInfo")
			return err
		}
		CxlStsRsnInf = append(CxlStsRsnInf, &reason)
		TxInfAndSts.CxlStsRsnInf = CxlStsRsnInf
	}
	if !isEmpty(TxInfAndSts) {
		CxlDtls.TxInfAndSts = TxInfAndSts
	}
	if !isEmpty(CxlDtls) {
		RsltnOfInvstgtn.CxlDtls = CxlDtls
	}
	if !isEmpty(RsltnOfInvstgtn) {
		msg.doc.RsltnOfInvstgtn = RsltnOfInvstgtn
	}
	return nil
}
