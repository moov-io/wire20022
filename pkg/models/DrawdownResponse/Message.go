package DrawdownResponse

import (
	"encoding/xml"
	"fmt"
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
	data MessageModel
	doc  pain014.Document
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
    if err := xml.Unmarshal(data, &msg.data); err != nil {
        return msg, fmt.Errorf("XML parse error: %w", err)
    }

	return msg, nil
}
func (msg *Message) CreateDocument() *model.ValidateError {
	msg.doc = pain014.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var CdtrPmtActvtnReqStsRpt pain014.CreditorPaymentActivationRequestStatusReportV07

	var GrpHdr pain014.GroupHeader871
	if msg.data.MessageId != "" {
		err := pain014.Max35Text(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = pain014.Max35Text(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreateDatetime) {
		err := fedwire.ISODateTime(msg.data.CreateDatetime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreateDatetime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreateDatetime)
	}
	if !isEmpty(msg.data.InitiatingParty) {
		InitgPty, vErr := PartyIdentification1351From(msg.data.InitiatingParty)
		if vErr != nil {
			vErr.InsertPath("InitiatingParty")
			return vErr
		}
		GrpHdr.InitgPty = InitgPty
	}
	if !isEmpty(msg.data.DebtorAgent) {
		err := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.data.DebtorAgent.PaymentSysCode).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
			vErr.InsertPath("DebtorAgent")
			return &vErr
		}
		err = pain014.RoutingNumberFRS1(msg.data.DebtorAgent.PaymentSysMemberId).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysMemberId",
				Message:   err.Error(),
			}
			vErr.InsertPath("DebtorAgent")
			return &vErr
		}
		Cd := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.data.DebtorAgent.PaymentSysCode)
		DbtrAgt := pain014.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pain014.FinancialInstitutionIdentification181{
				ClrSysMmbId: pain014.ClearingSystemMemberIdentification21{
					ClrSysId: pain014.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pain014.RoutingNumberFRS1(msg.data.DebtorAgent.PaymentSysMemberId),
				},
			},
		}
		if !isEmpty(DbtrAgt) {
			GrpHdr.DbtrAgt = DbtrAgt
		}
	}
	if !isEmpty(msg.data.CreditorAgent) {
		err := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.data.CreditorAgent.PaymentSysCode).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
			vErr.InsertPath("CreditorAgent")
			return &vErr
		}
		err = pain014.RoutingNumberFRS1(msg.data.CreditorAgent.PaymentSysMemberId).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysMemberId",
				Message:   err.Error(),
			}
			vErr.InsertPath("CreditorAgent")
			return &vErr
		}
		Cd := pain014.ExternalClearingSystemIdentification1CodeFixed(msg.data.CreditorAgent.PaymentSysCode)
		CdtrAgt := pain014.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pain014.FinancialInstitutionIdentification181{
				ClrSysMmbId: pain014.ClearingSystemMemberIdentification21{
					ClrSysId: pain014.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pain014.RoutingNumberFRS1(msg.data.CreditorAgent.PaymentSysMemberId),
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
	if msg.data.OriginalMessageId != "" {
		err := pain014.IMADFedwireFunds1(msg.data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInfAndSts.OrgnlMsgId = pain014.IMADFedwireFunds1(msg.data.OriginalMessageId)
	}
	if msg.data.OriginalMessageNameId != "" {
		err := pain014.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInfAndSts.OrgnlMsgNmId = pain014.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
	}
	if !isEmpty(msg.data.OriginalCreationDateTime) {
		err := fedwire.ISODateTime(msg.data.OriginalCreationDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalCreationDateTime",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInfAndSts.OrgnlCreDtTm = fedwire.ISODateTime(msg.data.OriginalCreationDateTime)
	}
	if !isEmpty(OrgnlGrpInfAndSts) {
		CdtrPmtActvtnReqStsRpt.OrgnlGrpInfAndSts = OrgnlGrpInfAndSts
	}

	var OrgnlPmtInfAndSts pain014.OriginalPaymentInstruction311
	if msg.data.OriginalPaymentInfoId != "" {
		err := pain014.Max35Text(msg.data.OriginalPaymentInfoId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalPaymentInfoId",
				Message:   err.Error(),
			}
		}
		OrgnlPmtInfAndSts.OrgnlPmtInfId = pain014.Max35Text(msg.data.OriginalPaymentInfoId)
	}
	if !isEmpty(msg.data.TransactionInformationAndStatus) {
		TxInfAndSts, vErr := PaymentTransaction1041From(msg.data.TransactionInformationAndStatus)
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
		msg.doc.CdtrPmtActvtnReqStsRpt = CdtrPmtActvtnReqStsRpt
	}
	return nil
}
