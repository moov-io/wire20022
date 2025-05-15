package ReturnRequestResponse

import (
	"encoding/xml"
	"fmt"
	"strings"
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
	Data   MessageModel
	Doc    camt029.Document
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
	if msg.Data.ResolvedCaseId == "" {
		ParamNames = append(ParamNames, "ResolvedCaseId")
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
	msg.Doc = camt029.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}

	var RsltnOfInvstgtn camt029.ResolutionOfInvestigationV09
	var Assgnmt camt029.CaseAssignment51
	if msg.Data.AssignmentId != "" {
		err := camt029.Max35Text(msg.Data.AssignmentId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AssignmentId",
				Message:   err.Error(),
			}
		}
		Assgnmt.Id = camt029.Max35Text(msg.Data.AssignmentId)
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
			err.InsertPath("Assignee")
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
		RsltnOfInvstgtn.Assgnmt = Assgnmt
	}
	var RslvdCase camt029.Case51
	if msg.Data.ResolvedCaseId != "" {
		err := camt029.Max35Text(msg.Data.ResolvedCaseId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ResolvedCaseId",
				Message:   err.Error(),
			}
		}
		RslvdCase.Id = camt029.Max35Text(msg.Data.ResolvedCaseId)
	}
	if !isEmpty(msg.Data.Creator) {
		Cretr, err := Party40Choice2From(msg.Data.Creator)
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
	if msg.Data.Status != "" {
		err := camt029.ExternalInvestigationExecutionConfirmation1Code(msg.Data.Status).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "Status",
				Message:   err.Error(),
			}
		}
		Conf := camt029.ExternalInvestigationExecutionConfirmation1Code(msg.Data.Status)
		Sts.Conf = &Conf
	}
	if !isEmpty(Sts) {
		RsltnOfInvstgtn.Sts = Sts
	}
	var CxlDtls camt029.UnderlyingTransaction221
	var TxInfAndSts camt029.PaymentTransaction1021
	var OrgnlGrpInf camt029.OriginalGroupInformation291
	if msg.Data.OriginalMessageId != "" {
		err := camt029.Max35Text(msg.Data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgId = camt029.Max35Text(msg.Data.OriginalMessageId)
	}
	if msg.Data.OriginalMessageNameId != "" {
		err := camt029.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgNmId = camt029.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId)
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
		TxInfAndSts.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.Data.OriginalInstructionId != "" {
		err := camt029.Max35Text(msg.Data.OriginalInstructionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInstructionId",
				Message:   err.Error(),
			}
		}
		OrgnlInstrId := camt029.Max35Text(msg.Data.OriginalInstructionId)
		TxInfAndSts.OrgnlInstrId = &OrgnlInstrId
	}
	if msg.Data.OriginalEndToEndId != "" {
		err := camt029.Max35Text(msg.Data.OriginalEndToEndId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalEndToEndId",
				Message:   err.Error(),
			}
		}
		OrgnlEndToEndId := camt029.Max35Text(msg.Data.OriginalEndToEndId)
		TxInfAndSts.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if msg.Data.OriginalUETR != "" {
		err := camt029.UUIDv4Identifier(msg.Data.OriginalUETR).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalUETR",
				Message:   err.Error(),
			}
		}
		TxInfAndSts.OrgnlUETR = camt029.UUIDv4Identifier(msg.Data.OriginalUETR)
	}
	if !isEmpty(msg.Data.CancellationStatusReasonInfo) {
		var CxlStsRsnInf []*camt029.CancellationStatusReason41
		reason, err := CancellationStatusReason41From(msg.Data.CancellationStatusReasonInfo)
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
		msg.Doc.RsltnOfInvstgtn = RsltnOfInvstgtn
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.RsltnOfInvstgtn.Assgnmt) {
		if !isEmpty(msg.Doc.RsltnOfInvstgtn.Assgnmt.Id) {
			msg.Data.AssignmentId = string(msg.Doc.RsltnOfInvstgtn.Assgnmt.Id)
		}
		if !isEmpty(msg.Doc.RsltnOfInvstgtn.Assgnmt.Assgnr) {
			Assgnr := Party40Choice1To(msg.Doc.RsltnOfInvstgtn.Assgnmt.Assgnr)
			msg.Data.Assigner = Assgnr
		}
		if !isEmpty(msg.Doc.RsltnOfInvstgtn.Assgnmt.Assgne) {
			Assgne := Party40Choice1To(msg.Doc.RsltnOfInvstgtn.Assgnmt.Assgne)
			msg.Data.Assignee = Assgne
		}
		if !isEmpty(msg.Doc.RsltnOfInvstgtn.Assgnmt.CreDtTm) {
			msg.Data.AssignmentCreateTime = time.Time(msg.Doc.RsltnOfInvstgtn.Assgnmt.CreDtTm)
		}
	}
	if !isEmpty(msg.Doc.RsltnOfInvstgtn.RslvdCase) {
		if !isEmpty(msg.Doc.RsltnOfInvstgtn.RslvdCase.Id) {
			msg.Data.ResolvedCaseId = string(msg.Doc.RsltnOfInvstgtn.RslvdCase.Id)
		}
		if !isEmpty(msg.Doc.RsltnOfInvstgtn.RslvdCase.Cretr) {
			Cretr := Party40Choice2To(msg.Doc.RsltnOfInvstgtn.RslvdCase.Cretr)
			msg.Data.Creator = Cretr
		}
	}
	if !isEmpty(msg.Doc.RsltnOfInvstgtn.Sts) {
		if !isEmpty(msg.Doc.RsltnOfInvstgtn.Sts.Conf) {
			msg.Data.Status = Status(*msg.Doc.RsltnOfInvstgtn.Sts.Conf)
		}
	}
	if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls) {
		if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts) {
			if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf) {
				if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId) {
					msg.Data.OriginalMessageId = string(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId)
				}
				if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlMsgNmId) {
					msg.Data.OriginalMessageNameId = string(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlMsgNmId)
				}
				if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlCreDtTm) {
					msg.Data.OriginalMessageCreateTime = time.Time(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlCreDtTm)
				}
			}
			if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlInstrId) {
				msg.Data.OriginalInstructionId = string(*msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlInstrId)
			}
			if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlEndToEndId) {
				msg.Data.OriginalEndToEndId = string(*msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlEndToEndId)
			}
			if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlUETR) {
				msg.Data.OriginalUETR = string(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlUETR)
			}
			if !isEmpty(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.CxlStsRsnInf) {
				if len(msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.CxlStsRsnInf) > 0 {
					msg.Data.CancellationStatusReasonInfo = CancellationStatusReason41To(*msg.Doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.CxlStsRsnInf[0])
				}
			}
		}
	}
	return nil
}
