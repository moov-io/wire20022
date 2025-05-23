package DrawdownRequest

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	pain013 "github.com/moov-io/fedwire20022/gen/DrawdownRequest_pain_013_001_07"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:pain.013.001.07"

type MessageModel struct {
	//A unique identifier (IMADFedwireFunds1) assigned to the message.
	MessageId string
	//The creation date and time (ISODateTime) of the message.
	CreateDatetime time.Time
	//The total number of transactions (Max15NumericTextFixed) included in the message.
	NumberofTransaction string
	//In the Fedwire Funds Service, this is a person or entity that requests a drawdown.
	InitiatingParty model.PartyIdentify

	//Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInfoId string
	//Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod PaymentMethod
	//Date at which the initiating party requests the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutDate model.Date
	//Party that owes an amount of money to the (ultimate) creditor.
	Debtor model.PartyIdentify
	//This is the account that will be debited by the debtor agent if the drawdown request is honored.
	DebtorAccountOtherId string
	//Financial institution servicing an account for the debtor.
	DebtorAgent model.Agent
	//Payment processes required to transfer cash from the debtor to the creditor.
	CreditTransTransaction CreditTransferTransaction
}
type Message struct {
	Data   MessageModel
	Doc    pain013.Document
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
	if msg.Data.CreateDatetime.IsZero() { // Check if CreatedDateTime is empty
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if msg.Data.NumberofTransaction == "" {
		ParamNames = append(ParamNames, "NumberofTransaction")
	}
	if isEmpty(msg.Data.InitiatingParty) {
		ParamNames = append(ParamNames, "InitiatingParty")
	}
	if msg.Data.PaymentInfoId == "" {
		ParamNames = append(ParamNames, "PaymentInfoId")
	}
	if msg.Data.PaymentMethod == "" {
		ParamNames = append(ParamNames, "PaymentMethod")
	}
	if isEmpty(msg.Data.RequestedExecutDate) {
		ParamNames = append(ParamNames, "RequestedExecutDate")
	}
	if isEmpty(msg.Data.Debtor) {
		ParamNames = append(ParamNames, "Debtor")
	}
	if isEmpty(msg.Data.DebtorAgent) {
		ParamNames = append(ParamNames, "DebtorAgent")
	}
	if isEmpty(msg.Data.CreditTransTransaction) {
		ParamNames = append(ParamNames, "CreditTransTransaction")
	} else if msg.Data.CreditTransTransaction.PaymentEndToEndId == "" {
		ParamNames = append(ParamNames, "CreditTransTransaction.PaymentEndToEndId")
	} else if msg.Data.CreditTransTransaction.PaymentUniqueId == "" {
		ParamNames = append(ParamNames, "CreditTransTransaction.PaymentUniqueId")
	} else if msg.Data.CreditTransTransaction.PayRequestType == "" {
		ParamNames = append(ParamNames, "CreditTransTransaction.PayRequestType")
	} else if isEmpty(msg.Data.CreditTransTransaction.Amount) {
		ParamNames = append(ParamNames, "CreditTransTransaction.Amount")
	} else if isEmpty(msg.Data.CreditTransTransaction.ChargeBearer) {
		ParamNames = append(ParamNames, "CreditTransTransaction.ChargeBearer")
	} else if isEmpty(msg.Data.CreditTransTransaction.CreditorAgent) {
		ParamNames = append(ParamNames, "CreditTransTransaction.CreditorAgent")
	} else if isEmpty(msg.Data.CreditTransTransaction.Creditor) {
		ParamNames = append(ParamNames, "CreditTransTransaction.Creditor")
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
	msg.Doc = pain013.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var CdtrPmtActvtnReq pain013.CreditorPaymentActivationRequestV07
	var GrpHdr pain013.GroupHeader781
	if msg.Data.MessageId != "" {
		err := pain013.IMADFedwireFunds1(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = pain013.IMADFedwireFunds1(msg.Data.MessageId)
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
	if msg.Data.NumberofTransaction != "" {
		err := pain013.Max15NumericTextFixed(msg.Data.NumberofTransaction).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "NumberofTransaction",
				Message:   err.Error(),
			}
		}
		GrpHdr.NbOfTxs = pain013.Max15NumericTextFixed(msg.Data.NumberofTransaction)
	}
	if !isEmpty(msg.Data.InitiatingParty) {
		InitgPty, vErr := PartyIdentification1351From(msg.Data.InitiatingParty)
		if vErr != nil {
			vErr.InsertPath("InitiatingParty")
			return vErr
		}
		GrpHdr.InitgPty = InitgPty
	}
	if !isEmpty(GrpHdr) {
		CdtrPmtActvtnReq.GrpHdr = GrpHdr
	}
	var PmtInf pain013.PaymentInstruction311
	if msg.Data.PaymentInfoId != "" {
		err := pain013.Max35Text(msg.Data.PaymentInfoId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "PaymentInfoId",
				Message:   err.Error(),
			}
		}
		PmtInf.PmtInfId = pain013.Max35Text(msg.Data.PaymentInfoId)
	}
	if msg.Data.PaymentMethod != "" {
		err := pain013.PaymentMethod7Code1(msg.Data.PaymentMethod).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "PaymentMethod",
				Message:   err.Error(),
			}
		}
		PmtInf.PmtMtd = pain013.PaymentMethod7Code1(msg.Data.PaymentMethod)
	}
	if !isEmpty(msg.Data.RequestedExecutDate) {
		err := msg.Data.RequestedExecutDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RequestedExecutDate",
				Message:   err.Error(),
			}
		}
		Dt := msg.Data.RequestedExecutDate.Date()
		PmtInf.ReqdExctnDt = pain013.DateAndDateTime2Choice1{
			Dt: &Dt,
		}
	}
	if !isEmpty(msg.Data.Debtor) {
		Dbtr, vErr := PartyIdentification1352From(msg.Data.Debtor)
		if vErr != nil {
			vErr.InsertPath("Debtor")
			return vErr
		}
		PmtInf.Dbtr = Dbtr
	}
	if msg.Data.DebtorAccountOtherId != "" {
		err := pain013.Max34Text(msg.Data.DebtorAccountOtherId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "DebtorAccountOtherId",
				Message:   err.Error(),
			}
		}
		Othr := pain013.GenericAccountIdentification1{
			Id: pain013.Max34Text(msg.Data.DebtorAccountOtherId),
		}
		DbtrAcct := pain013.CashAccount38{
			Id: pain013.AccountIdentification4Choice{
				Othr: &Othr,
			},
		}
		PmtInf.DbtrAcct = &DbtrAcct
	}
	if !isEmpty(msg.Data.DebtorAgent) {
		err := pain013.ExternalClearingSystemIdentification1CodeFixed(msg.Data.DebtorAgent.PaymentSysCode).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
			vErr.InsertPath("DebtorAgent")
			return &vErr
		}
		err = pain013.RoutingNumberFRS1(msg.Data.DebtorAgent.PaymentSysMemberId).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysMemberId",
				Message:   err.Error(),
			}
			vErr.InsertPath("DebtorAgent")
			return &vErr
		}
		Cd := pain013.ExternalClearingSystemIdentification1CodeFixed(msg.Data.DebtorAgent.PaymentSysCode)
		DbtrAgt := pain013.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: pain013.FinancialInstitutionIdentification181{
				ClrSysMmbId: pain013.ClearingSystemMemberIdentification21{
					ClrSysId: pain013.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: pain013.RoutingNumberFRS1(msg.Data.DebtorAgent.PaymentSysMemberId),
				},
			},
		}
		PmtInf.DbtrAgt = DbtrAgt
	}
	if !isEmpty(msg.Data.CreditTransTransaction) {
		CdtTrfTx, vErr := CreditTransferTransaction351From(msg.Data.CreditTransTransaction)
		if vErr != nil {
			vErr.InsertPath("CreditTransTransaction")
			return vErr
		}
		PmtInf.CdtTrfTx = CdtTrfTx
	}
	if !isEmpty(PmtInf) {
		CdtrPmtActvtnReq.PmtInf = PmtInf
	}
	if !isEmpty(CdtrPmtActvtnReq) {
		msg.Doc.CdtrPmtActvtnReq = CdtrPmtActvtnReq
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.CdtrPmtActvtnReq) {
		if !isEmpty(msg.Doc.CdtrPmtActvtnReq.GrpHdr) {
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.GrpHdr.MsgId) {
				msg.Data.MessageId = string(msg.Doc.CdtrPmtActvtnReq.GrpHdr.MsgId)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.GrpHdr.CreDtTm) {
				msg.Data.CreateDatetime = time.Time(msg.Doc.CdtrPmtActvtnReq.GrpHdr.CreDtTm)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.GrpHdr.NbOfTxs) {
				msg.Data.NumberofTransaction = string(msg.Doc.CdtrPmtActvtnReq.GrpHdr.NbOfTxs)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.GrpHdr.InitgPty) {
				msg.Data.InitiatingParty = PartyIdentification1351To(msg.Doc.CdtrPmtActvtnReq.GrpHdr.InitgPty)
			}
		}
		if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf) {
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.PmtInfId) {
				msg.Data.PaymentInfoId = string(msg.Doc.CdtrPmtActvtnReq.PmtInf.PmtInfId)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.PmtMtd) {
				msg.Data.PaymentMethod = PaymentMethod(msg.Doc.CdtrPmtActvtnReq.PmtInf.PmtMtd)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.ReqdExctnDt) && !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.ReqdExctnDt.Dt) {
				msg.Data.RequestedExecutDate = model.FromDate(*msg.Doc.CdtrPmtActvtnReq.PmtInf.ReqdExctnDt.Dt)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.Dbtr) {
				msg.Data.Debtor = PartyIdentification1352To(msg.Doc.CdtrPmtActvtnReq.PmtInf.Dbtr)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAcct) {
				msg.Data.DebtorAccountOtherId = string(msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAcct.Id.Othr.Id)
			}
			if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAgt) {
				if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAgt.FinInstnId) {
					if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAgt.FinInstnId.ClrSysMmbId) {
						if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId) && !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
							msg.Data.DebtorAgent.PaymentSysCode = model.PaymentSystemType(*msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
						}
						if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId) {
							msg.Data.DebtorAgent.PaymentSysMemberId = string(msg.Doc.CdtrPmtActvtnReq.PmtInf.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId)
						}
					}
				}
				if !isEmpty(msg.Doc.CdtrPmtActvtnReq.PmtInf.CdtTrfTx) {
					msg.Data.CreditTransTransaction = CreditTransferTransaction351To(msg.Doc.CdtrPmtActvtnReq.PmtInf.CdtTrfTx)
				}
			}
		}
	}
	return nil
}
