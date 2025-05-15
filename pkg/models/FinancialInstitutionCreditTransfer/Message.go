package FinancialInstitutionCreditTransfer

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	pacs009 "github.com/moov-io/fedwire20022/gen/FinancialInstitutionCreditTransfer_pacs_009_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08"

type MessageModel struct {
	//Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	MessageId string
	//Date and time at which the message was created.
	CreateDateTime time.Time
	//Number of individual transactions contained in the message.
	NumberOfTransactions int
	//Method used to settle the (batch of) payment instructions.
	SettlementMethod model.SettlementMethodType
	//Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClearingSystem model.CommonClearingSysCodeType
	//Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	PaymentInstructionId string
	//Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	PaymentEndToEndId string
	//Universally unique identifier to provide an end-to-end reference of a payment transaction.
	PaymentUETR string
	//User community specific instrument.
	LocalInstrument InstrumentType
	//Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount model.CurrencyAndAmount
	//Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate model.Date
	//Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent model.Agent
	//Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent model.Agent
	//Financial institution that owes an amount of money to the (ultimate) financial institutional creditor.
	Debtor model.FiniancialInstitutionId
	//Financial institution servicing an account for the debtor.
	DebtorAgent model.FiniancialInstitutionId
	//Financial institution servicing an account for the creditor.
	CreditorAgent model.FiniancialInstitutionId
	//Financial institution that receives an amount of money from the financial institutional debtor.
	Creditor model.FiniancialInstitutionId
	//Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInfo string
	//Provides information on the underlying customer credit transfer for which cover is provided.
	UnderlyingCustomerCreditTransfer CreditTransferTransaction
}
type Message struct {
	Data   MessageModel
	Doc    pacs009.Document
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
	if msg.Data.CreateDateTime.IsZero() {
		ParamNames = append(ParamNames, "CreateDateTime")
	}
	if msg.Data.NumberOfTransactions <= 0 {
		ParamNames = append(ParamNames, "NumberOfTransactions")
	}
	if msg.Data.SettlementMethod == "" {
		ParamNames = append(ParamNames, "SettlementMethod")
	}
	if msg.Data.ClearingSystem == "" {
		ParamNames = append(ParamNames, "ClearingSystem")
	}
	if msg.Data.PaymentEndToEndId == "" {
		ParamNames = append(ParamNames, "PaymentEndToEndId")
	}
	if msg.Data.PaymentUETR == "" {
		ParamNames = append(ParamNames, "PaymentUETR")
	}
	if msg.Data.LocalInstrument == "" {
		ParamNames = append(ParamNames, "LocalInstrument")
	}
	if isEmpty(msg.Data.InterbankSettlementAmount) {
		ParamNames = append(ParamNames, "InterbankSettlementAmount")
	}
	if isEmpty(msg.Data.InterbankSettlementDate) {
		ParamNames = append(ParamNames, "InterbankSettlementDate")
	}
	if isEmpty(msg.Data.InstructingAgent) {
		ParamNames = append(ParamNames, "InstructingAgent")
	}
	if isEmpty(msg.Data.InstructedAgent) {
		ParamNames = append(ParamNames, "InstructedAgent")
	}
	if isEmpty(msg.Data.Debtor) {
		ParamNames = append(ParamNames, "Debtor")
	}
	if isEmpty(msg.Data.Creditor) {
		ParamNames = append(ParamNames, "Creditor")
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
	msg.Doc = pacs009.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var FICdtTrf pacs009.FinancialInstitutionCreditTransferV08
	var GrpHdr pacs009.GroupHeader931
	if msg.Data.MessageId != "" {
		err := pacs009.IMADFedwireFunds1(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = pacs009.IMADFedwireFunds1(msg.Data.MessageId)
	}
	if !isEmpty(msg.Data.CreateDateTime) {
		err := fedwire.ISODateTime(msg.Data.CreateDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreateDateTime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.Data.CreateDateTime)
	}
	if msg.Data.NumberOfTransactions > 0 {
		err := pacs009.Max15NumericTextFixed(strconv.Itoa(msg.Data.NumberOfTransactions)).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "NumberOfTransactions",
				Message:   err.Error(),
			}
		}
		GrpHdr.NbOfTxs = pacs009.Max15NumericTextFixed(strconv.Itoa(msg.Data.NumberOfTransactions))
	}
	var SttlmInf pacs009.SettlementInstruction71
	if msg.Data.SettlementMethod != "" {
		err := pacs009.SettlementMethod1Code1(msg.Data.SettlementMethod).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "SettlementMethod",
				Message:   err.Error(),
			}
		}
		SttlmInf.SttlmMtd = pacs009.SettlementMethod1Code1(msg.Data.SettlementMethod)
	}
	if msg.Data.ClearingSystem != "" {
		err := pacs009.ExternalCashClearingSystem1CodeFixed(msg.Data.ClearingSystem).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ClearingSystem",
				Message:   err.Error(),
			}
		}
		Cd := pacs009.ExternalCashClearingSystem1CodeFixed(msg.Data.ClearingSystem)
		SttlmInf.ClrSys = pacs009.ClearingSystemIdentification3Choice1{
			Cd: &Cd,
		}
	}
	if !isEmpty(SttlmInf) {
		GrpHdr.SttlmInf = SttlmInf
	}
	if !isEmpty(GrpHdr) {
		FICdtTrf.GrpHdr = GrpHdr
	}
	var CdtTrfTxInf pacs009.CreditTransferTransaction361
	var PmtId pacs009.PaymentIdentification71

	if msg.Data.PaymentInstructionId != "" {
		err := pacs009.Max35Text(msg.Data.PaymentInstructionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "PaymentInstructionId",
				Message:   err.Error(),
			}
		}
		InstrId := pacs009.Max35Text(msg.Data.PaymentInstructionId)
		PmtId.InstrId = &InstrId
	}
	if msg.Data.PaymentEndToEndId != "" {
		err := pacs009.Max35Text(msg.Data.PaymentEndToEndId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "PaymentEndToEndId",
				Message:   err.Error(),
			}
		}
		PmtId.EndToEndId = pacs009.Max35Text(msg.Data.PaymentEndToEndId)
	}
	if msg.Data.PaymentUETR != "" {
		err := pacs009.UUIDv4Identifier(msg.Data.PaymentUETR).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "PaymentUETR",
				Message:   err.Error(),
			}
		}
		PmtId.UETR = pacs009.UUIDv4Identifier(msg.Data.PaymentUETR)
	}
	if !isEmpty(PmtId) {
		CdtTrfTxInf.PmtId = PmtId
	}
	var PmtTpInf pacs009.PaymentTypeInformation281
	if msg.Data.LocalInstrument != "" {
		err := pacs009.LocalInstrumentFedwireFunds1(msg.Data.LocalInstrument).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "LocalInstrument",
				Message:   err.Error(),
			}
		}
		Prtry := pacs009.LocalInstrumentFedwireFunds1(msg.Data.LocalInstrument)
		PmtTpInf.LclInstrm = pacs009.LocalInstrument2Choice1{
			Prtry: &Prtry,
		}
	}
	if !isEmpty(PmtTpInf) {
		CdtTrfTxInf.PmtTpInf = PmtTpInf
	}
	if !isEmpty(msg.Data.InterbankSettlementAmount) {
		err := pacs009.ActiveCurrencyAndAmountFedwire1SimpleType(msg.Data.InterbankSettlementAmount.Amount).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InterbankSettlementAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = pacs009.ActiveCurrencyCodeFixed(msg.Data.InterbankSettlementAmount.Currency).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InterbankSettlementAmount.Currency",
				Message:   err.Error(),
			}
		}
		CdtTrfTxInf.IntrBkSttlmAmt = pacs009.ActiveCurrencyAndAmountFedwire1{
			Value: pacs009.ActiveCurrencyAndAmountFedwire1SimpleType(msg.Data.InterbankSettlementAmount.Amount),
			Ccy:   pacs009.ActiveCurrencyCodeFixed(msg.Data.InterbankSettlementAmount.Currency),
		}
	}
	if !isEmpty(msg.Data.InterbankSettlementDate) {
		err := msg.Data.InterbankSettlementDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InterbankSettlementDate",
				Message:   err.Error(),
			}
		}
		CdtTrfTxInf.IntrBkSttlmDt = msg.Data.InterbankSettlementDate.Date()
	}
	if !isEmpty(msg.Data.InstructingAgent) {
		err := pacs009.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructingAgent.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructingAgent.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = pacs009.RoutingNumberFRS1(msg.Data.InstructingAgent.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructingAgent.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		Cd := pacs009.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructingAgent.PaymentSysCode)
		CdtTrfTxInf.InstgAgt = pacs009.BranchAndFinancialInstitutionIdentification62{
			FinInstnId: pacs009.FinancialInstitutionIdentification182{
				ClrSysMmbId: pacs009.ClearingSystemMemberIdentification22{
					ClrSysId: pacs009.ClearingSystemIdentification2Choice2{
						Cd: &Cd,
					},
					MmbId: pacs009.RoutingNumberFRS1(msg.Data.InstructingAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(msg.Data.InstructedAgent) {
		err := pacs009.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructedAgent.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructedAgent.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = pacs009.RoutingNumberFRS1(msg.Data.InstructedAgent.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructedAgent.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		Cd := pacs009.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructedAgent.PaymentSysCode)
		CdtTrfTxInf.InstdAgt = pacs009.BranchAndFinancialInstitutionIdentification62{
			FinInstnId: pacs009.FinancialInstitutionIdentification182{
				ClrSysMmbId: pacs009.ClearingSystemMemberIdentification22{
					ClrSysId: pacs009.ClearingSystemIdentification2Choice2{
						Cd: &Cd,
					},
					MmbId: pacs009.RoutingNumberFRS1(msg.Data.InstructedAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(msg.Data.Debtor) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.Data.Debtor.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(msg.Data.Debtor.BusinessId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Debtor.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(msg.Data.Debtor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.Data.Debtor.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(msg.Data.Debtor.ClearingSystemId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Debtor.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.Data.Debtor.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.Data.Debtor.ClearintSystemMemberId != "" {
			err := pacs009.Max35Text(msg.Data.Debtor.ClearintSystemMemberId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Debtor.ClearintSystemMemberId",
					Message:   err.Error(),
				}
			}
			if !isEmpty(finialialId.ClrSysMmbId) {
				finialialId.ClrSysMmbId.MmbId = pacs009.Max35Text(msg.Data.Debtor.ClearintSystemMemberId)
			}
		}
		if msg.Data.Debtor.Name != "" {
			err := pacs009.Max140Text(msg.Data.Debtor.Name).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Debtor.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(msg.Data.Debtor.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.Data.Debtor.Address) {
			PstlAdr, vErr := PostalAddress241From(msg.Data.Debtor.Address)
			if vErr != nil {
				vErr.InsertPath("Debtor.Address")
				return vErr
			}
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(agent) {
			CdtTrfTxInf.Dbtr = agent
		}
	}
	if !isEmpty(msg.Data.DebtorAgent) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.Data.DebtorAgent.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(msg.Data.DebtorAgent.BusinessId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "DebtorAgent.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(msg.Data.DebtorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.Data.DebtorAgent.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(msg.Data.DebtorAgent.ClearingSystemId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "DebtorAgent.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.Data.DebtorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.Data.DebtorAgent.ClearintSystemMemberId != "" {
			err := pacs009.Max35Text(msg.Data.DebtorAgent.ClearintSystemMemberId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "DebtorAgent.ClearintSystemMemberId",
					Message:   err.Error(),
				}
			}
			finialialId.ClrSysMmbId.MmbId = pacs009.Max35Text(msg.Data.DebtorAgent.ClearintSystemMemberId)
		}
		if msg.Data.DebtorAgent.Name != "" {
			err := pacs009.Max140Text(msg.Data.DebtorAgent.Name).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "DebtorAgent.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(msg.Data.DebtorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.Data.DebtorAgent.Address) {
			PstlAdr, vErr := PostalAddress241From(msg.Data.DebtorAgent.Address)
			if vErr != nil {
				vErr.InsertPath("DebtorAgent")
				return vErr
			}
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(agent) {
			CdtTrfTxInf.DbtrAgt = &agent
		}
	}
	if !isEmpty(msg.Data.CreditorAgent) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.Data.CreditorAgent.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(msg.Data.CreditorAgent.BusinessId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "CreditorAgent.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(msg.Data.CreditorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.Data.CreditorAgent.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(msg.Data.CreditorAgent.ClearingSystemId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "CreditorAgent.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.Data.CreditorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.Data.CreditorAgent.ClearintSystemMemberId != "" {
			err := pacs009.Max35Text(msg.Data.CreditorAgent.ClearintSystemMemberId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "CreditorAgent.ClearintSystemMemberId",
					Message:   err.Error(),
				}
			}
			finialialId.ClrSysMmbId.MmbId = pacs009.Max35Text(msg.Data.CreditorAgent.ClearintSystemMemberId)
		}
		if msg.Data.CreditorAgent.Name != "" {
			err := pacs009.Max140Text(msg.Data.CreditorAgent.Name).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "CreditorAgent.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(msg.Data.CreditorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.Data.CreditorAgent.Address) {
			PstlAdr, vErr := PostalAddress241From(msg.Data.CreditorAgent.Address)
			if vErr != nil {
				vErr.InsertPath("CreditorAgent.Address")
				return vErr
			}
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(agent) {
			CdtTrfTxInf.CdtrAgt = &agent
		}
	}
	if !isEmpty(msg.Data.Creditor) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.Data.Creditor.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(msg.Data.Creditor.BusinessId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Creditor.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(msg.Data.Creditor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.Data.Creditor.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(msg.Data.Creditor.ClearingSystemId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Creditor.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.Data.Creditor.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.Data.Creditor.ClearintSystemMemberId != "" {
			err := pacs009.Max35Text(msg.Data.Creditor.ClearintSystemMemberId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Creditor.ClearintSystemMemberId",
					Message:   err.Error(),
				}
			}
			if !isEmpty(finialialId.ClrSysMmbId) {
				finialialId.ClrSysMmbId.MmbId = pacs009.Max35Text(msg.Data.Creditor.ClearintSystemMemberId)
			}
		}
		if msg.Data.Creditor.Name != "" {
			err := pacs009.Max140Text(msg.Data.Creditor.Name).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Creditor.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(msg.Data.Creditor.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.Data.Creditor.Address) {
			PstlAdr, vErr := PostalAddress241From(msg.Data.Creditor.Address)
			if vErr != nil {
				vErr.InsertPath("Creditor.Address")
				return vErr
			}
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(agent) {
			CdtTrfTxInf.Cdtr = agent
		}
	}
	if msg.Data.RemittanceInfo != "" {
		err := pacs009.Max140Text(msg.Data.RemittanceInfo).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RemittanceInfo",
				Message:   err.Error(),
			}
		}
		Ustrd := pacs009.Max140Text(msg.Data.RemittanceInfo)
		RmtInf := pacs009.RemittanceInformation21{
			Ustrd: &Ustrd,
		}
		CdtTrfTxInf.RmtInf = &RmtInf
	}
	if !isEmpty(msg.Data.UnderlyingCustomerCreditTransfer) {
		UndrlygCstmrCdtTrf, vErr := CreditTransferTransaction371From(msg.Data.UnderlyingCustomerCreditTransfer)
		if vErr != nil {
			vErr.InsertPath("UnderlyingCustomerCreditTransfer")
			return vErr
		}
		if !isEmpty(UndrlygCstmrCdtTrf) {
			CdtTrfTxInf.UndrlygCstmrCdtTrf = &UndrlygCstmrCdtTrf

		}
	}

	if !isEmpty(CdtTrfTxInf) {
		FICdtTrf.CdtTrfTxInf = CdtTrfTxInf
	}
	if !isEmpty(FICdtTrf) {
		msg.Doc.FICdtTrf = FICdtTrf
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.FICdtTrf) {
		if !isEmpty(msg.Doc.FICdtTrf.GrpHdr) {
			msg.Data.MessageId = string(msg.Doc.FICdtTrf.GrpHdr.MsgId)
			msg.Data.CreateDateTime = time.Time(msg.Doc.FICdtTrf.GrpHdr.CreDtTm)
			msg.Data.NumberOfTransactions, _ = strconv.Atoi(string(msg.Doc.FICdtTrf.GrpHdr.NbOfTxs))
			msg.Data.SettlementMethod = model.SettlementMethodType(msg.Doc.FICdtTrf.GrpHdr.SttlmInf.SttlmMtd)
			if !isEmpty(msg.Doc.FICdtTrf.GrpHdr.SttlmInf.ClrSys) {
				msg.Data.ClearingSystem = model.CommonClearingSysCodeType(*msg.Doc.FICdtTrf.GrpHdr.SttlmInf.ClrSys.Cd)
			}
		}
		if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf) {
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.PmtId) {
				if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.PmtId.InstrId) {
					msg.Data.PaymentInstructionId = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.PmtId.InstrId)
				}
				if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.PmtId.EndToEndId) {
					msg.Data.PaymentEndToEndId = string(msg.Doc.FICdtTrf.CdtTrfTxInf.PmtId.EndToEndId)
				}
				if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.PmtId.UETR) {
					msg.Data.PaymentUETR = string(msg.Doc.FICdtTrf.CdtTrfTxInf.PmtId.UETR)
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.PmtTpInf) {
				if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm) {
					msg.Data.LocalInstrument = InstrumentType(*msg.Doc.FICdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm.Prtry)
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.IntrBkSttlmAmt) {
				msg.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
					Amount:   float64(msg.Doc.FICdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Value),
					Currency: string(msg.Doc.FICdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Ccy),
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.IntrBkSttlmDt) {
				msg.Data.InterbankSettlementDate = model.FromDate(msg.Doc.FICdtTrf.CdtTrfTxInf.IntrBkSttlmDt)
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.InstgAgt) {
				msg.Data.InstructingAgent = model.Agent{
					PaymentSysCode:     model.PaymentSystemType(*msg.Doc.FICdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd),
					PaymentSysMemberId: string(msg.Doc.FICdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId),
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.InstdAgt) {
				msg.Data.InstructedAgent = model.Agent{
					PaymentSysCode:     model.PaymentSystemType(*msg.Doc.FICdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd),
					PaymentSysMemberId: string(msg.Doc.FICdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId),
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr) {
				if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId) {
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.BICFI) {
						msg.Data.Debtor.BusinessId = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.BICFI)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.ClrSysMmbId) {
						msg.Data.Debtor.ClearingSystemId = model.PaymentSystemType(*msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
						msg.Data.Debtor.ClearintSystemMemberId = string(msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.ClrSysMmbId.MmbId)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.Nm) {
						msg.Data.Debtor.Name = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.Nm)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.PstlAdr) {
						msg.Data.Debtor.Address = PostalAddress241To(*msg.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.PstlAdr)
					}
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt) {
				if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId) {
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.BICFI) {
						msg.Data.DebtorAgent.BusinessId = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.BICFI)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.ClrSysMmbId) {
						msg.Data.DebtorAgent.ClearingSystemId = model.PaymentSystemType(*msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
						msg.Data.DebtorAgent.ClearintSystemMemberId = string(msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.Nm) {
						msg.Data.DebtorAgent.Name = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.Nm)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.PstlAdr) {
						msg.Data.DebtorAgent.Address = PostalAddress241To(*msg.Doc.FICdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.PstlAdr)
					}
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt) {
				if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId) {
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.BICFI) {
						msg.Data.CreditorAgent.BusinessId = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.BICFI)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.ClrSysMmbId) {
						msg.Data.CreditorAgent.ClearingSystemId = model.PaymentSystemType(*msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
						msg.Data.CreditorAgent.ClearintSystemMemberId = string(msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.ClrSysMmbId.MmbId)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.Nm) {
						msg.Data.CreditorAgent.Name = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.Nm)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.PstlAdr) {
						msg.Data.CreditorAgent.Address = PostalAddress241To(*msg.Doc.FICdtTrf.CdtTrfTxInf.CdtrAgt.FinInstnId.PstlAdr)
					}
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr) {
				if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId) {
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.BICFI) {
						msg.Data.Creditor.BusinessId = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.BICFI)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.ClrSysMmbId) {
						msg.Data.Creditor.ClearingSystemId = model.PaymentSystemType(*msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
						msg.Data.Creditor.ClearintSystemMemberId = string(msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.ClrSysMmbId.MmbId)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.Nm) {
						msg.Data.Creditor.Name = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.Nm)
					}
					if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.PstlAdr) {
						msg.Data.Creditor.Address = PostalAddress241To(*msg.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.PstlAdr)
					}
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.RmtInf) {
				if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.RmtInf.Ustrd) {
					msg.Data.RemittanceInfo = string(*msg.Doc.FICdtTrf.CdtTrfTxInf.RmtInf.Ustrd)
				}
			}
			if !isEmpty(msg.Doc.FICdtTrf.CdtTrfTxInf.UndrlygCstmrCdtTrf) {
				UndrlygCstmrCdtTrf := CreditTransferTransaction371To(*msg.Doc.FICdtTrf.CdtTrfTxInf.UndrlygCstmrCdtTrf)
				msg.Data.UnderlyingCustomerCreditTransfer = UndrlygCstmrCdtTrf
			}
		}
	}
	return nil
}
