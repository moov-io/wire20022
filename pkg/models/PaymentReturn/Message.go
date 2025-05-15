package PaymentReturn

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	pacs004 "github.com/moov-io/fedwire20022/gen/PaymentReturn_pacs_004_001_10"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10"

type MessageModel struct {
	//Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	MessageId string
	//Date and time at which the message was created.
	CreatedDateTime time.Time
	//Number of individual transactions contained in the message.
	NumberOfTransactions int
	//Method used to settle the (batch of) payment instructions.
	SettlementMethod model.SettlementMethodType
	//Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem model.CommonClearingSysCodeType
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
	//Amount of money to be moved between the instructing agent and the instructed agent in the returned instruction.
	ReturnedInterbankSettlementAmount model.CurrencyAndAmount
	//Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate model.Date
	//Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	ReturnedInstructedAmount model.CurrencyAndAmount
	//Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer model.ChargeBearerType
	//Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent model.Agent
	//Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent model.Agent
	//Provides all parties (agents and non-agents) involved in a return transaction.
	RtrChain ReturnChain
	//Provides detailed information on the return reason.
	ReturnReasonInformation Reason
	//Key elements used to identify the original transaction that is being referred to.
	OriginalTransactionRef model.InstrumentPropCodeType
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
	if msg.Data.NumberOfTransactions == 0 {
		ParamNames = append(ParamNames, "NumberOfTransactions")
	}
	if msg.Data.SettlementMethod == "" {
		ParamNames = append(ParamNames, "SettlementMethod")
	}
	if msg.Data.ClearingSystem == "" {
		ParamNames = append(ParamNames, "ClearingSystem")
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
	if isEmpty(msg.Data.ReturnedInterbankSettlementAmount) {
		ParamNames = append(ParamNames, "ReturnedInterbankSettlementAmount")
	}
	if isEmpty(msg.Data.InterbankSettlementDate) {
		ParamNames = append(ParamNames, "InterbankSettlementDate")
	}
	if isEmpty(msg.Data.ReturnedInstructedAmount) {
		ParamNames = append(ParamNames, "ReturnedInstructedAmount")
	}
	if isEmpty(msg.Data.InstructingAgent) {
		ParamNames = append(ParamNames, "InstructingAgent")
	}
	if isEmpty(msg.Data.InstructedAgent) {
		ParamNames = append(ParamNames, "InstructedAgent")
	}
	if isEmpty(msg.Data.RtrChain) {
		ParamNames = append(ParamNames, "RtrChain")
	} else if isEmpty(msg.Data.RtrChain.Debtor) {
		ParamNames = append(ParamNames, "RtrChain.Debtor")
	} else if isEmpty(msg.Data.RtrChain.Creditor) {
		ParamNames = append(ParamNames, "RtrChain.Creditor")
	}
	if isEmpty(msg.Data.ReturnReasonInformation) {
		ParamNames = append(ParamNames, "ReturnReasonInformation")
	} else if msg.Data.ReturnReasonInformation.Reason == "" {
		ParamNames = append(ParamNames, "ReturnReasonInformation.Reason")
	}
	if msg.Data.OriginalTransactionRef == "" {
		ParamNames = append(ParamNames, "OriginalTransactionRef")
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
	var PmtRtr pacs004.PaymentReturnV10
	var GrpHdr pacs004.GroupHeader901
	if msg.Data.MessageId != "" {
		err := pacs004.IMADFedwireFunds1(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = pacs004.IMADFedwireFunds1(msg.Data.MessageId)
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
	if msg.Data.NumberOfTransactions != 0 {
		err := pacs004.Max15NumericTextFixed(strconv.Itoa(msg.Data.NumberOfTransactions)).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "NumberOfTransactions",
				Message:   err.Error(),
			}
		}
		GrpHdr.NbOfTxs = pacs004.Max15NumericTextFixed(strconv.Itoa(msg.Data.NumberOfTransactions))
	}
	var SttlmInf pacs004.SettlementInstruction71
	if msg.Data.SettlementMethod != "" {
		err := pacs004.SettlementMethod1Code1(msg.Data.SettlementMethod).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "SettlementMethod",
				Message:   err.Error(),
			}
		}
		SttlmInf.SttlmMtd = pacs004.SettlementMethod1Code1(msg.Data.SettlementMethod)
	}
	if msg.Data.ClearingSystem != "" {
		err := pacs004.ExternalCashClearingSystem1CodeFixed(msg.Data.ClearingSystem).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ClearingSystem",
				Message:   err.Error(),
			}
		}
		Cd := pacs004.ExternalCashClearingSystem1CodeFixed(msg.Data.ClearingSystem)
		SttlmInf.ClrSys = pacs004.ClearingSystemIdentification3Choice1{
			Cd: &Cd,
		}
	}
	if !isEmpty(SttlmInf) {
		GrpHdr.SttlmInf = SttlmInf
	}
	if !isEmpty(GrpHdr) {
		PmtRtr.GrpHdr = GrpHdr
	}
	var TxInf pacs004.PaymentTransaction1181
	var OrgnlGrpInf pacs004.OriginalGroupInformation291
	if msg.Data.OriginalMessageId != "" {
		err := pacs004.Max35Text(msg.Data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgId = pacs004.Max35Text(msg.Data.OriginalMessageId)
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
	if !isEmpty(msg.Data.ReturnedInterbankSettlementAmount) {
		err := pacs004.ActiveCurrencyAndAmountFedwire1SimpleType(msg.Data.ReturnedInterbankSettlementAmount.Amount).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReturnedInterbankSettlementAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = pacs004.ActiveCurrencyCodeFixed(msg.Data.ReturnedInterbankSettlementAmount.Currency).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReturnedInterbankSettlementAmount.Currency",
				Message:   err.Error(),
			}
		}
		RtrdIntrBkSttlmAmt := pacs004.ActiveCurrencyAndAmountFedwire1{
			Value: pacs004.ActiveCurrencyAndAmountFedwire1SimpleType(msg.Data.ReturnedInterbankSettlementAmount.Amount),
			Ccy:   pacs004.ActiveCurrencyCodeFixed(msg.Data.ReturnedInterbankSettlementAmount.Currency),
		}
		TxInf.RtrdIntrBkSttlmAmt = RtrdIntrBkSttlmAmt
	}
	if !isEmpty(msg.Data.InterbankSettlementDate) {
		err := msg.Data.InterbankSettlementDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InterbankSettlementDate",
				Message:   err.Error(),
			}
		}
		IntrBkSttlmDt := msg.Data.InterbankSettlementDate.Date()
		TxInf.IntrBkSttlmDt = IntrBkSttlmDt
	}
	if !isEmpty(msg.Data.ReturnedInstructedAmount) {
		err := pacs004.ActiveCurrencyAndAmountFedwire1SimpleType(msg.Data.ReturnedInstructedAmount.Amount).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReturnedInstructedAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = pacs004.ActiveCurrencyCodeFixed(msg.Data.ReturnedInstructedAmount.Currency).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReturnedInstructedAmount.Currency",
				Message:   err.Error(),
			}
		}
		TxInf.RtrdInstdAmt = pacs004.ActiveOrHistoricCurrencyAndAmount{
			Value: pacs004.ActiveOrHistoricCurrencyAndAmountSimpleType(msg.Data.ReturnedInstructedAmount.Amount),
			Ccy:   pacs004.ActiveOrHistoricCurrencyCode(msg.Data.ReturnedInstructedAmount.Currency),
		}
	}
	if msg.Data.ChargeBearer != "" {
		err := pacs004.ChargeBearerType1Code1(msg.Data.ChargeBearer).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ChargeBearer",
				Message:   err.Error(),
			}
		}
		ChrgBr := pacs004.ChargeBearerType1Code1(msg.Data.ChargeBearer)
		TxInf.ChrgBr = &ChrgBr
	}
	if !isEmpty(msg.Data.InstructingAgent) {
		InstgAgt, err := BranchAndFinancialInstitutionIdentification62From(msg.Data.InstructingAgent)
		if err != nil {
			err.InsertPath("InstructingAgent")
			return err
		}
		TxInf.InstgAgt = InstgAgt
	}
	if !isEmpty(msg.Data.InstructedAgent) {
		InstdAgt, err := BranchAndFinancialInstitutionIdentification62From(msg.Data.InstructedAgent)
		if err != nil {
			err.InsertPath("InstructedAgent")
			return err
		}
		TxInf.InstdAgt = InstdAgt
	}
	if !isEmpty(msg.Data.RtrChain) {
		RtrChain, err := TransactionParties81From(msg.Data.RtrChain)
		if err != nil {
			err.InsertPath("RtrChain")
		}
		TxInf.RtrChain = RtrChain
	}
	if !isEmpty(msg.Data.ReturnReasonInformation) {
		RtrRsnInf, err := PaymentReturnReason61From(msg.Data.ReturnReasonInformation)
		if err != nil {
			err.InsertPath("ReturnReasonInformation")
			return err
		}
		TxInf.RtrRsnInf = RtrRsnInf
	}
	if msg.Data.OriginalTransactionRef != "" {
		err := pacs004.LocalInstrumentFedwireFunds1(msg.Data.OriginalTransactionRef).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalTransactionRef",
				Message:   err.Error(),
			}
		}
		Prtry := pacs004.LocalInstrumentFedwireFunds1(msg.Data.OriginalTransactionRef)
		TxInf.OrgnlTxRef = pacs004.OriginalTransactionReference321{
			PmtTpInf: pacs004.PaymentTypeInformation271{
				LclInstrm: pacs004.LocalInstrument2Choice1{
					Prtry: &Prtry,
				},
			},
		}
	}
	if !isEmpty(TxInf) {
		PmtRtr.TxInf = TxInf
	}
	if !isEmpty(PmtRtr) {
		msg.Doc.PmtRtr = PmtRtr
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.PmtRtr) {
		if !isEmpty(msg.Doc.PmtRtr.GrpHdr) {
			if !isEmpty(msg.Doc.PmtRtr.GrpHdr.MsgId) {
				msg.Data.MessageId = string(msg.Doc.PmtRtr.GrpHdr.MsgId)
			}
			if !isEmpty(msg.Doc.PmtRtr.GrpHdr.CreDtTm) {
				msg.Data.CreatedDateTime = time.Time(msg.Doc.PmtRtr.GrpHdr.CreDtTm)
			}
			if !isEmpty(msg.Doc.PmtRtr.GrpHdr.NbOfTxs) {
				msg.Data.NumberOfTransactions, _ = strconv.Atoi(string(msg.Doc.PmtRtr.GrpHdr.NbOfTxs))
			}
			if !isEmpty(msg.Doc.PmtRtr.GrpHdr.SttlmInf) {
				if !isEmpty(msg.Doc.PmtRtr.GrpHdr.SttlmInf.SttlmMtd) {
					msg.Data.SettlementMethod = model.SettlementMethodType(msg.Doc.PmtRtr.GrpHdr.SttlmInf.SttlmMtd)
				}
			}
			if !isEmpty(msg.Doc.PmtRtr.GrpHdr.SttlmInf.ClrSys) {
				msg.Data.ClearingSystem = model.CommonClearingSysCodeType(*msg.Doc.PmtRtr.GrpHdr.SttlmInf.ClrSys.Cd)
			}
		}
		if !isEmpty(msg.Doc.PmtRtr.TxInf) {
			if !isEmpty(msg.Doc.PmtRtr.TxInf.OrgnlGrpInf) {
				if !isEmpty(msg.Doc.PmtRtr.TxInf.OrgnlGrpInf.OrgnlMsgId) {
					msg.Data.OriginalMessageId = string(msg.Doc.PmtRtr.TxInf.OrgnlGrpInf.OrgnlMsgId)
				}
				if !isEmpty(msg.Doc.PmtRtr.TxInf.OrgnlGrpInf.OrgnlMsgNmId) {
					msg.Data.OriginalMessageNameId = string(msg.Doc.PmtRtr.TxInf.OrgnlGrpInf.OrgnlMsgNmId)
				}
				if !isEmpty(msg.Doc.PmtRtr.TxInf.OrgnlGrpInf.OrgnlCreDtTm) {
					msg.Data.OriginalCreationDateTime = time.Time(msg.Doc.PmtRtr.TxInf.OrgnlGrpInf.OrgnlCreDtTm)
				}
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.OrgnlInstrId) {
				msg.Data.OriginalInstructionId = string(*msg.Doc.PmtRtr.TxInf.OrgnlInstrId)
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.OrgnlEndToEndId) {
				msg.Data.OriginalEndToEndId = string(*msg.Doc.PmtRtr.TxInf.OrgnlEndToEndId)
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.OrgnlUETR) {
				msg.Data.OriginalUETR = string(msg.Doc.PmtRtr.TxInf.OrgnlUETR)
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.RtrdIntrBkSttlmAmt) {
				if !isEmpty(msg.Doc.PmtRtr.TxInf.RtrdIntrBkSttlmAmt.Value) {
					msg.Data.ReturnedInterbankSettlementAmount.Amount = float64(msg.Doc.PmtRtr.TxInf.RtrdIntrBkSttlmAmt.Value)
				}
				if !isEmpty(msg.Doc.PmtRtr.TxInf.RtrdIntrBkSttlmAmt.Ccy) {
					msg.Data.ReturnedInterbankSettlementAmount.Currency = string(msg.Doc.PmtRtr.TxInf.RtrdIntrBkSttlmAmt.Ccy)
				}
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.IntrBkSttlmDt) {
				msg.Data.InterbankSettlementDate = model.FromDate(msg.Doc.PmtRtr.TxInf.IntrBkSttlmDt)
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.RtrdInstdAmt) {
				if !isEmpty(msg.Doc.PmtRtr.TxInf.RtrdInstdAmt.Value) {
					msg.Data.ReturnedInstructedAmount.Amount = float64(msg.Doc.PmtRtr.TxInf.RtrdInstdAmt.Value)
				}
				if !isEmpty(msg.Doc.PmtRtr.TxInf.RtrdInstdAmt.Ccy) {
					msg.Data.ReturnedInstructedAmount.Currency = string(msg.Doc.PmtRtr.TxInf.RtrdInstdAmt.Ccy)
				}
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.ChrgBr) {
				msg.Data.ChargeBearer = model.ChargeBearerType(*msg.Doc.PmtRtr.TxInf.ChrgBr)
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.InstgAgt) {
				InstgAgt := BranchAndFinancialInstitutionIdentification62To(msg.Doc.PmtRtr.TxInf.InstgAgt)
				msg.Data.InstructingAgent = InstgAgt
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.InstdAgt) {
				InstdAgt := BranchAndFinancialInstitutionIdentification62To(msg.Doc.PmtRtr.TxInf.InstdAgt)
				msg.Data.InstructedAgent = InstdAgt
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.RtrChain) {
				RtrChain := TransactionParties81To(msg.Doc.PmtRtr.TxInf.RtrChain)
				msg.Data.RtrChain = RtrChain
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.RtrRsnInf) {
				RtrRsnInf := PaymentReturnReason61To(msg.Doc.PmtRtr.TxInf.RtrRsnInf)
				msg.Data.ReturnReasonInformation = RtrRsnInf
			}
			if !isEmpty(msg.Doc.PmtRtr.TxInf.OrgnlTxRef) {
				msg.Data.OriginalTransactionRef = model.InstrumentPropCodeType(*msg.Doc.PmtRtr.TxInf.OrgnlTxRef.PmtTpInf.LclInstrm.Prtry)
			}

		}
	}
	return nil
}
