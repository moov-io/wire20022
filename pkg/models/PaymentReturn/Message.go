package PaymentReturn

import (
	"encoding/xml"
	"fmt"
	"strconv"
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
	data MessageModel
	doc  pacs004.Document
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
	msg.doc = pacs004.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var PmtRtr pacs004.PaymentReturnV10
	var GrpHdr pacs004.GroupHeader901
	if msg.data.MessageId != "" {
		err := pacs004.IMADFedwireFunds1(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = pacs004.IMADFedwireFunds1(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
	if msg.data.NumberOfTransactions != 0 {
		err := pacs004.Max15NumericTextFixed(strconv.Itoa(msg.data.NumberOfTransactions)).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "NumberOfTransactions",
				Message:   err.Error(),
			}
		}
		GrpHdr.NbOfTxs = pacs004.Max15NumericTextFixed(strconv.Itoa(msg.data.NumberOfTransactions))
	}
	var SttlmInf pacs004.SettlementInstruction71
	if msg.data.SettlementMethod != "" {
		err := pacs004.SettlementMethod1Code1(msg.data.SettlementMethod).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "SettlementMethod",
				Message:   err.Error(),
			}
		}
		SttlmInf.SttlmMtd = pacs004.SettlementMethod1Code1(msg.data.SettlementMethod)
	}
	if msg.data.ClearingSystem != "" {
		err := pacs004.ExternalCashClearingSystem1CodeFixed(msg.data.ClearingSystem).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ClearingSystem",
				Message:   err.Error(),
			}
		}
		Cd := pacs004.ExternalCashClearingSystem1CodeFixed(msg.data.ClearingSystem)
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
	if msg.data.OriginalMessageId != "" {
		err := pacs004.Max35Text(msg.data.OriginalMessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgId = pacs004.Max35Text(msg.data.OriginalMessageId)
	}
	if msg.data.OriginalMessageNameId != "" {
		err := pacs004.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgNmId = pacs004.MessageNameIdentificationFRS1(msg.data.OriginalMessageNameId)
	}
	if !isEmpty(msg.data.OriginalCreationDateTime) {
		err := fedwire.ISODateTime(msg.data.OriginalCreationDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalCreationDateTime",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(msg.data.OriginalCreationDateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		TxInf.OrgnlGrpInf = OrgnlGrpInf
	}
	if msg.data.OriginalInstructionId != "" {
		err := pacs004.Max35Text(msg.data.OriginalInstructionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalInstructionId",
				Message:   err.Error(),
			}
		}
		OrgnlInstrId := pacs004.Max35Text(msg.data.OriginalInstructionId)
		TxInf.OrgnlInstrId = &OrgnlInstrId
	}
	if msg.data.OriginalEndToEndId != "" {
		err := pacs004.Max35Text(msg.data.OriginalEndToEndId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalEndToEndId",
				Message:   err.Error(),
			}
		}
		OrgnlEndToEndId := pacs004.Max35Text(msg.data.OriginalEndToEndId)
		TxInf.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if msg.data.OriginalUETR != "" {
		err := pacs004.UUIDv4Identifier(msg.data.OriginalUETR).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalUETR",
				Message:   err.Error(),
			}
		}
		TxInf.OrgnlUETR = pacs004.UUIDv4Identifier(msg.data.OriginalUETR)
	}
	if !isEmpty(msg.data.ReturnedInterbankSettlementAmount) {
		err := pacs004.ActiveCurrencyAndAmountFedwire1SimpleType(msg.data.ReturnedInterbankSettlementAmount.Amount).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReturnedInterbankSettlementAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = pacs004.ActiveCurrencyCodeFixed(msg.data.ReturnedInterbankSettlementAmount.Currency).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReturnedInterbankSettlementAmount.Currency",
				Message:   err.Error(),
			}
		}
		RtrdIntrBkSttlmAmt := pacs004.ActiveCurrencyAndAmountFedwire1{
			Value: pacs004.ActiveCurrencyAndAmountFedwire1SimpleType(msg.data.ReturnedInterbankSettlementAmount.Amount),
			Ccy:   pacs004.ActiveCurrencyCodeFixed(msg.data.ReturnedInterbankSettlementAmount.Currency),
		}
		TxInf.RtrdIntrBkSttlmAmt = RtrdIntrBkSttlmAmt
	}
	if !isEmpty(msg.data.InterbankSettlementDate) {
		err := msg.data.InterbankSettlementDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InterbankSettlementDate",
				Message:   err.Error(),
			}
		}
		IntrBkSttlmDt := msg.data.InterbankSettlementDate.Date()
		TxInf.IntrBkSttlmDt = IntrBkSttlmDt
	}
	if !isEmpty(msg.data.ReturnedInstructedAmount) {
		err := pacs004.ActiveCurrencyAndAmountFedwire1SimpleType(msg.data.ReturnedInstructedAmount.Amount).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReturnedInstructedAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = pacs004.ActiveCurrencyCodeFixed(msg.data.ReturnedInstructedAmount.Currency).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReturnedInstructedAmount.Currency",
				Message:   err.Error(),
			}
		}
		TxInf.RtrdInstdAmt = pacs004.ActiveOrHistoricCurrencyAndAmount{
			Value: pacs004.ActiveOrHistoricCurrencyAndAmountSimpleType(msg.data.ReturnedInstructedAmount.Amount),
			Ccy:   pacs004.ActiveOrHistoricCurrencyCode(msg.data.ReturnedInstructedAmount.Currency),
		}
	}
	if msg.data.ChargeBearer != "" {
		err := pacs004.ChargeBearerType1Code1(msg.data.ChargeBearer).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ChargeBearer",
				Message:   err.Error(),
			}
		}
		ChrgBr := pacs004.ChargeBearerType1Code1(msg.data.ChargeBearer)
		TxInf.ChrgBr = &ChrgBr
	}
	if !isEmpty(msg.data.InstructingAgent) {
		InstgAgt, err := BranchAndFinancialInstitutionIdentification62From(msg.data.InstructingAgent)
		if err != nil {
			err.InsertPath("InstructingAgent")
			return err
		}
		TxInf.InstgAgt = InstgAgt
	}
	if !isEmpty(msg.data.InstructedAgent) {
		InstdAgt, err := BranchAndFinancialInstitutionIdentification62From(msg.data.InstructedAgent)
		if err != nil {
			err.InsertPath("InstructedAgent")
			return err
		}
		TxInf.InstdAgt = InstdAgt
	}
	if !isEmpty(msg.data.RtrChain) {
		RtrChain, err := TransactionParties81From(msg.data.RtrChain)
		if err != nil {
			err.InsertPath("RtrChain")
		}
		TxInf.RtrChain = RtrChain
	}
	if !isEmpty(msg.data.ReturnReasonInformation) {
		RtrRsnInf, err := PaymentReturnReason61From(msg.data.ReturnReasonInformation)
		if err != nil {
			err.InsertPath("ReturnReasonInformation")
			return err
		}
		TxInf.RtrRsnInf = RtrRsnInf
	}
	if msg.data.OriginalTransactionRef != "" {
		err := pacs004.LocalInstrumentFedwireFunds1(msg.data.OriginalTransactionRef).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalTransactionRef",
				Message:   err.Error(),
			}
		}
		Prtry := pacs004.LocalInstrumentFedwireFunds1(msg.data.OriginalTransactionRef)
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
		msg.doc.PmtRtr = PmtRtr
	}
	return nil
}
