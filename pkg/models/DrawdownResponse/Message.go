package DrawdownResponse

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	pain014 "github.com/moov-io/fedwire20022/gen/DrawdownResponse_pain_014_001_07"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:pain.014.001.07"

type MessageModel struct {
	//A unique identifier (IMADFedwireFunds1) assigned to the message.
	MessageId string
	//The creation date and time (ISODateTime) of the message.
	CreateDatetime time.Time
	//In the Fedwire Funds Service, this is a person or entity that requests a drawdown.
	InitiatingParty model.PartyIdentify
	//Financial institution servicing an account for the debtor.
	DebtorAgent model.Agent
	// /Financial institution servicing an account for the creditor.
	CreditorAgent model.Agent
	//This should be the Message Identification of the original drawdown request (pain.013) message to which this drawdown request response message relates.
	OriginalMessageId string
	//Specifies the original message name identifier to which the message refers.
	OriginalMessageNameId string
	//Date and time at which the original message was created.
	OriginalCreationDateTime time.Time
	//Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInfoId string
	//Provides information on the original transactions to which the status report message refers.
	TransactionInformationAndStatus TransactionInfoAndStatus
}
type Message struct {
	Data   MessageModel
	Doc    pain014.Document
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
	if msg.Data.MessageId == "" {
		ParamNames = append(ParamNames, "MessageId")
	}
	if msg.Data.CreateDatetime.IsZero() {
		ParamNames = append(ParamNames, "CreateDatetime")
	}
	if isEmpty(msg.Data.InitiatingParty) {
		ParamNames = append(ParamNames, "InitiatingParty")
	}
	if isEmpty(msg.Data.DebtorAgent) {
		ParamNames = append(ParamNames, "DebtorAgent")
	}
	if isEmpty(msg.Data.CreditorAgent) {
		ParamNames = append(ParamNames, "CreditorAgent")
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
	if msg.Data.OriginalPaymentInfoId == "" {
		ParamNames = append(ParamNames, "OriginalPaymentInfoId")
	}
	if isEmpty(msg.Data.TransactionInformationAndStatus) {
		ParamNames = append(ParamNames, "TransactionInformationAndStatus")
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
	msg.Doc = pain014.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var CdtrPmtActvtnReqStsRpt pain014.CreditorPaymentActivationRequestStatusReportV07

	var GrpHdr pain014.GroupHeader871
	if msg.Data.MessageId != "" {
		err := pain014.Max35Text(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = pain014.Max35Text(msg.Data.MessageId)
	}
	if !isEmpty(msg.Data.CreateDatetime) {
		err := fedwire.ISODateTime(msg.Data.CreateDatetime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreateDatetime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.Data.CreateDatetime)
	}
	if !isEmpty(msg.Data.InitiatingParty) {
		InitgPty, vErr := PartyIdentification1351From(msg.Data.InitiatingParty)
		if vErr != nil {
			vErr.InsertPath("InitiatingParty")
			return vErr
		}
		GrpHdr.InitgPty = InitgPty
	}
	if !isEmpty(msg.Data.DebtorAgent) {
		err := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.Data.DebtorAgent.PaymentSysCode).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
			vErr.InsertPath("DebtorAgent")
			return &vErr
		}
		err = pain014.RoutingNumberFRS1(msg.Data.DebtorAgent.PaymentSysMemberId).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysMemberId",
				Message:   err.Error(),
			}
			vErr.InsertPath("DebtorAgent")
			return &vErr
		}
		Cd := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.Data.DebtorAgent.PaymentSysCode)
		DbtrAgt := pain014.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pain014.FinancialInstitutionIdentification181{
				ClrSysMmbId: pain014.ClearingSystemMemberIdentification21{
					ClrSysId: pain014.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pain014.RoutingNumberFRS1(msg.Data.DebtorAgent.PaymentSysMemberId),
				},
			},
		}
		if !isEmpty(DbtrAgt) {
			GrpHdr.DbtrAgt = DbtrAgt
		}
	}
	if !isEmpty(msg.Data.CreditorAgent) {
		err := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.Data.CreditorAgent.PaymentSysCode).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
			vErr.InsertPath("CreditorAgent")
			return &vErr
		}
		err = pain014.RoutingNumberFRS1(msg.Data.CreditorAgent.PaymentSysMemberId).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysMemberId",
				Message:   err.Error(),
			}
			vErr.InsertPath("CreditorAgent")
			return &vErr
		}
		Cd := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.Data.CreditorAgent.PaymentSysCode)
		CdtrAgt := pain014.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pain014.FinancialInstitutionIdentification181{
				ClrSysMmbId: pain014.ClearingSystemMemberIdentification21{
					ClrSysId: pain014.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pain014.RoutingNumberFRS1(msg.Data.CreditorAgent.PaymentSysMemberId),
				},
			},
		}
		if !isEmpty(CdtrAgt) {
			GrpHdr.CdtrAgt = CdtrAgt
		}
	}
	if !isEmpty(GrpHdr) {
		CdtrPmtActvtnReqStsRpt.GrpHdr = GrpHdr
	}

	var OrgnlGrpInfAndSts pain014.OriginalGroupInformation301
	if msg.Data.OriginalMessageId != "" {
		err := pain014.IMADFedwireFunds1(msg.Data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInfAndSts.OrgnlMsgId = pain014.IMADFedwireFunds1(msg.Data.OriginalMessageId)
	}
	if msg.Data.OriginalMessageNameId != "" {
		err := pain014.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInfAndSts.OrgnlMsgNmId = pain014.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId)
	}
	if !isEmpty(msg.Data.OriginalCreationDateTime) {
		err := fedwire.ISODateTime(msg.Data.OriginalCreationDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalCreationDateTime",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInfAndSts.OrgnlCreDtTm = fedwire.ISODateTime(msg.Data.OriginalCreationDateTime)
	}
	if !isEmpty(OrgnlGrpInfAndSts) {
		CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts = OrgnlGrpInfAndSts
	}

	var OrgnlPmtInfAndSts pain014.OriginalPaymentInstruction311
	if msg.Data.OriginalPaymentInfoId != "" {
		err := pain014.Max35Text(msg.Data.OriginalPaymentInfoId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalPaymentInfoId",
				Message:   err.Error(),
			}
		}
		OrgnlPmtInfAndSts.OrgnlPmtInfId = pain014.Max35Text(msg.Data.OriginalPaymentInfoId)
	}
	if !isEmpty(msg.Data.TransactionInformationAndStatus) {
		TxInfAndSts, vErr := PaymentTransaction1041From(msg.Data.TransactionInformationAndStatus)
		if vErr != nil {
			vErr.InsertPath("TransactionInformationAndStatus")
			return vErr
		}
		OrgnlPmtInfAndSts.TxInfAndSts = TxInfAndSts
	}

	if !isEmpty(OrgnlPmtInfAndSts) {
		CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts = OrgnlPmtInfAndSts
	}
	if !isEmpty(CdtrPmtActvtnReqStsRpt) {
		msg.Doc.CdtrPmtActvtnReqStsRpt = CdtrPmtActvtnReqStsRpt
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt) {
		if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr) {
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId) {
				msg.Data.MessageId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CreDtTm) {
				msg.Data.CreateDatetime = time.Time(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CreDtTm)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty) {
				msg.Data.InitiatingParty = PartyIdentification1351To(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.InitgPty)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt) {
				if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId) {
					if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId) {
						if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId) {
							if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
								msg.Data.DebtorAgent.PaymentSysCode = model.PaymentSystemType(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
							}
							if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId) {
								msg.Data.DebtorAgent.PaymentSysMemberId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId)
							}
						}
					}
				}
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt) {
				if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId) {
					if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId) {
						if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId) {
							if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
								msg.Data.CreditorAgent.PaymentSysCode = model.PaymentSystemType(*msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
							}
							if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId) {
								msg.Data.CreditorAgent.PaymentSysMemberId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.GrpHdr.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId)
							}
						}
					}
				}
			}
		}
		if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts) {
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId) {
				msg.Data.OriginalMessageId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId) {
				msg.Data.OriginalMessageNameId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlCreDtTm) {
				msg.Data.OriginalCreationDateTime = time.Time(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts.OrgnlCreDtTm)
			}
		}
		if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts) {
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.OrgnlPmtInfId) {
				msg.Data.OriginalPaymentInfoId = string(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.OrgnlPmtInfId)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts) {
				msg.Data.TransactionInformationAndStatus = PaymentTransaction1041To(msg.Doc.CdtrPmtActvtnReqStsRpt.OrgnlPmtInfAndSts.TxInfAndSts)
			}
		}
	}
	return nil
}
