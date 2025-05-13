package FedwireFundsPaymentStatus

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	pacs002 "github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus_pacs_002_001_10"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10"

type MessageModel struct {
	//Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	MessageId string
	//Date and time at which the message was created.
	CreatedDateTime time.Time
	//Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageId string
	//Specifies the original message name identifier to which the message refers.
	OriginalMessageNameId string
	//Date and time at which the original message was created.
	OriginalMessageCreateTime time.Time
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUETR string
	//Specifies the status of a transaction, in a coded form.
	TransactionStatus model.TransactionStatusCode
	//Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AcceptanceDateTime time.Time
	//Date and time at which a transaction is completed and cleared, that is, payment is effected.
	EffectiveInterbankSettlementDate model.Date
	//Provides detailed information on the status reason.
	StatusReasonInformation string
	//Further details on the status reason.
	ReasonAdditionalInfo string
	//Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent model.Agent
	//Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent model.Agent
}
type Message struct {
	Data   MessageModel
	Doc    pacs002.Document
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
	if isEmpty(msg.Data.CreatedDateTime) {
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if msg.Data.OriginalMessageId == "" {
		ParamNames = append(ParamNames, "OriginalMessageId")
	}
	if msg.Data.OriginalMessageNameId == "" {
		ParamNames = append(ParamNames, "OriginalMessageNameId")
	}
	if isEmpty(msg.Data.OriginalMessageCreateTime) {
		ParamNames = append(ParamNames, "OriginalMessageCreateTime")
	}
	if msg.Data.OriginalUETR == "" {
		ParamNames = append(ParamNames, "OriginalUETR")
	}
	if msg.Data.TransactionStatus == "" {
		ParamNames = append(ParamNames, "TransactionStatus")
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
	msg.Doc = pacs002.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10",
			Local: "Document",
		},
	}
	var FIToFIPmtStsRpt pacs002.FIToFIPaymentStatusReportV10
	var GrpHdr pacs002.GroupHeader911
	if msg.Data.MessageId != "" {
		err := pacs002.Max35Text(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = pacs002.Max35Text(msg.Data.MessageId)
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
		FIToFIPmtStsRpt.GrpHdr = GrpHdr
	}
	var TxInfAndSts pacs002.PaymentTransaction1101
	var OrgnlGrpInf pacs002.OriginalGroupInformation291
	if msg.Data.OriginalMessageId != "" {
		err := pacs002.IMADFedwireFunds1(msg.Data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgId = pacs002.IMADFedwireFunds1(msg.Data.OriginalMessageId)
	}
	if msg.Data.OriginalMessageNameId != "" {
		err := pacs002.IMADFedwireFunds1(msg.Data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgNmId = pacs002.MessageNameIdentificationFRS1(msg.Data.OriginalMessageNameId)
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
	if msg.Data.OriginalUETR != "" {
		err := pacs002.UUIDv4Identifier(msg.Data.OriginalUETR).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalUETR",
				Message:   err.Error(),
			}
		}
		TxInfAndSts.OrgnlUETR = pacs002.UUIDv4Identifier(msg.Data.OriginalUETR)
	}
	if msg.Data.TransactionStatus != "" {
		err := pacs002.ExternalPaymentTransactionStatus1Code(msg.Data.TransactionStatus).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TransactionStatus",
				Message:   err.Error(),
			}
		}
		TxInfAndSts.TxSts = pacs002.ExternalPaymentTransactionStatus1Code(msg.Data.TransactionStatus)
	}
	if !isEmpty(msg.Data.AcceptanceDateTime) {
		err := fedwire.ISODateTime(msg.Data.AcceptanceDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AcceptanceDateTime",
				Message:   err.Error(),
			}
		}
		AccptncDtTm := fedwire.ISODateTime(msg.Data.AcceptanceDateTime)
		TxInfAndSts.AccptncDtTm = &AccptncDtTm
	}
	if !isEmpty(msg.Data.EffectiveInterbankSettlementDate) {
		err := msg.Data.EffectiveInterbankSettlementDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "EffectiveInterbankSettlementDate",
				Message:   err.Error(),
			}
		}
		Dt := msg.Data.EffectiveInterbankSettlementDate.Date()
		FctvIntrBkSttlmDt := pacs002.DateAndDateTime2Choice1{
			Dt: &Dt,
		}
		TxInfAndSts.FctvIntrBkSttlmDt = &FctvIntrBkSttlmDt
	}
	if msg.Data.StatusReasonInformation != "" {
		err := pacs002.Max35Text(msg.Data.StatusReasonInformation).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "StatusReasonInformation",
				Message:   err.Error(),
			}
		}
		Prtry := pacs002.Max35Text(msg.Data.StatusReasonInformation)
		var StsRsnInf []*pacs002.StatusReasonInformation121
		reson := pacs002.StatusReasonInformation121{
			Rsn: pacs002.StatusReason6Choice1{
				Prtry: &Prtry,
			},
		}
		if msg.Data.ReasonAdditionalInfo != "" {
			err := pacs002.Max105Text(msg.Data.ReasonAdditionalInfo).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "ReasonAdditionalInfo",
					Message:   err.Error(),
				}
			}
			reson.AddtlInf = pacs002.Max105Text(msg.Data.ReasonAdditionalInfo)
		}
		StsRsnInf = append(StsRsnInf, &reson)
		if !isEmpty(StsRsnInf) {
			TxInfAndSts.StsRsnInf = StsRsnInf
		}
	}

	if !isEmpty(msg.Data.InstructingAgent) {
		err := pacs002.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructingAgent.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructingAgent.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = pacs002.Max35TextFixed(msg.Data.InstructingAgent.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructingAgent.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		Cd := pacs002.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructingAgent.PaymentSysCode)
		TxInfAndSts.InstgAgt = pacs002.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pacs002.FinancialInstitutionIdentification181{
				ClrSysMmbId: pacs002.ClearingSystemMemberIdentification21{
					ClrSysId: pacs002.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pacs002.Max35TextFixed(msg.Data.InstructingAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(msg.Data.InstructedAgent) {
		err := pacs002.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructedAgent.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructedAgent.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = pacs002.RoutingNumberFRS1(msg.Data.InstructedAgent.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructedAgent.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		Cd := pacs002.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructedAgent.PaymentSysCode)
		TxInfAndSts.InstdAgt = pacs002.BranchAndFinancialInstitutionIdentification62{
			FinInstnId: pacs002.FinancialInstitutionIdentification182{
				ClrSysMmbId: pacs002.ClearingSystemMemberIdentification22{
					ClrSysId: pacs002.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pacs002.RoutingNumberFRS1(msg.Data.InstructedAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(TxInfAndSts) {
		FIToFIPmtStsRpt.TxInfAndSts = TxInfAndSts
	}
	if !isEmpty(FIToFIPmtStsRpt) {
		msg.Doc.FIToFIPmtStsRpt = FIToFIPmtStsRpt
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.FIToFIPmtStsRpt) {
		if !isEmpty(msg.Doc.FIToFIPmtStsRpt.GrpHdr) {
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.GrpHdr.MsgId) {
				msg.Data.MessageId = string(msg.Doc.FIToFIPmtStsRpt.GrpHdr.MsgId)
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.GrpHdr.CreDtTm) {
				msg.Data.CreatedDateTime = time.Time(msg.Doc.FIToFIPmtStsRpt.GrpHdr.CreDtTm)
			}
		}
		if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts) {
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf) {
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId) {
					msg.Data.OriginalMessageId = string(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId)
				}
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgNmId) {
					msg.Data.OriginalMessageNameId = string(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgNmId)
				}
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlCreDtTm) {
					msg.Data.OriginalMessageCreateTime = time.Time(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlCreDtTm)
				}
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlUETR) {
				msg.Data.OriginalUETR = string(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlUETR)
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.TxSts) {
				msg.Data.TransactionStatus = model.TransactionStatusCode(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.TxSts)
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.AccptncDtTm) {
				msg.Data.AcceptanceDateTime = time.Time(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.AccptncDtTm)
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.FctvIntrBkSttlmDt) {
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.FctvIntrBkSttlmDt.Dt) {
					msg.Data.EffectiveInterbankSettlementDate = model.FromDate(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.FctvIntrBkSttlmDt.Dt)
				}
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf) {
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf[0].Rsn.Prtry) {
					msg.Data.StatusReasonInformation = string(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf[0].Rsn.Prtry)
				}
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf[0].AddtlInf) {
					msg.Data.ReasonAdditionalInfo = string(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.StsRsnInf[0].AddtlInf)
				}
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstgAgt) {
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
					msg.Data.InstructingAgent.PaymentSysCode = model.PaymentSystemType(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
				}
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId) {
					msg.Data.InstructingAgent.PaymentSysMemberId = string(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId)
				}
			}
			if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstdAgt) {
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
					msg.Data.InstructedAgent.PaymentSysCode = model.PaymentSystemType(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
				}
				if !isEmpty(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId) {
					msg.Data.InstructedAgent.PaymentSysMemberId = string(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId)
				}
			}
		}
	}
	return nil
}
