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
	data MessageModel
	doc  pacs009.Document
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

func (msg *Message) ValidateRequiredFields() *model.ValidateError {
	// Initialize the RequireError object
	var ParamNames []string

	// Check required fields and append missing ones to ParamNames
	if msg.data.MessageId == "" {
		ParamNames = append(ParamNames, "MessageId")
	}
	if msg.data.CreateDateTime.IsZero() {
		ParamNames = append(ParamNames, "CreateDateTime")
	}
	if msg.data.NumberOfTransactions <= 0 {
		ParamNames = append(ParamNames, "NumberOfTransactions")
	}
	if msg.data.SettlementMethod == "" {
		ParamNames = append(ParamNames, "SettlementMethod")
	}
	if msg.data.ClearingSystem == "" {
		ParamNames = append(ParamNames, "ClearingSystem")
	}
	if msg.data.PaymentEndToEndId == "" {
		ParamNames = append(ParamNames, "PaymentEndToEndId")
	}
	if msg.data.PaymentUETR == "" {
		ParamNames = append(ParamNames, "PaymentUETR")
	}
	if msg.data.LocalInstrument == "" {
		ParamNames = append(ParamNames, "LocalInstrument")
	}
	if isEmpty(msg.data.InterbankSettlementAmount) {
		ParamNames = append(ParamNames, "InterbankSettlementAmount")
	}
	if isEmpty(msg.data.InterbankSettlementDate) {
		ParamNames = append(ParamNames, "InterbankSettlementDate")
	}
	if isEmpty(msg.data.InstructingAgent) {
		ParamNames = append(ParamNames, "InstructingAgent")
	}
	if isEmpty(msg.data.InstructedAgent) {
		ParamNames = append(ParamNames, "InstructedAgent")
	}
	if isEmpty(msg.data.Debtor) {
		ParamNames = append(ParamNames, "Debtor")
	}
	if isEmpty(msg.data.Creditor) {
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
	msg.doc = pacs009.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var FICdtTrf pacs009.FinancialInstitutionCreditTransferV08
	var GrpHdr pacs009.GroupHeader931
	if msg.data.MessageId != "" {
		err := pacs009.IMADFedwireFunds1(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = pacs009.IMADFedwireFunds1(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreateDateTime) {
		err := fedwire.ISODateTime(msg.data.CreateDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreateDateTime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreateDateTime)
	}
	if msg.data.NumberOfTransactions > 0 {
		err := pacs009.Max15NumericTextFixed(strconv.Itoa(msg.data.NumberOfTransactions)).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "NumberOfTransactions",
				Message:   err.Error(),
			}
		}
		GrpHdr.NbOfTxs = pacs009.Max15NumericTextFixed(strconv.Itoa(msg.data.NumberOfTransactions))
	}
	var SttlmInf pacs009.SettlementInstruction71
	if msg.data.SettlementMethod != "" {
		err := pacs009.SettlementMethod1Code1(msg.data.SettlementMethod).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "SettlementMethod",
				Message:   err.Error(),
			}
		}
		SttlmInf.SttlmMtd = pacs009.SettlementMethod1Code1(msg.data.SettlementMethod)
	}
	if msg.data.ClearingSystem != "" {
		err := pacs009.ExternalCashClearingSystem1CodeFixed(msg.data.ClearingSystem).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ClearingSystem",
				Message:   err.Error(),
			}
		}
		Cd := pacs009.ExternalCashClearingSystem1CodeFixed(msg.data.ClearingSystem)
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

	if msg.data.PaymentInstructionId != "" {
		err := pacs009.Max35Text(msg.data.PaymentInstructionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "PaymentInstructionId",
				Message:   err.Error(),
			}
		}
		InstrId := pacs009.Max35Text(msg.data.PaymentInstructionId)
		PmtId.InstrId = &InstrId
	}
	if msg.data.PaymentEndToEndId != "" {
		err := pacs009.Max35Text(msg.data.PaymentEndToEndId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "PaymentEndToEndId",
				Message:   err.Error(),
			}
		}
		PmtId.EndToEndId = pacs009.Max35Text(msg.data.PaymentEndToEndId)
	}
	if msg.data.PaymentUETR != "" {
		err := pacs009.UUIDv4Identifier(msg.data.PaymentUETR).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "PaymentUETR",
				Message:   err.Error(),
			}
		}
		PmtId.UETR = pacs009.UUIDv4Identifier(msg.data.PaymentUETR)
	}
	if !isEmpty(PmtId) {
		CdtTrfTxInf.PmtId = PmtId
	}
	var PmtTpInf pacs009.PaymentTypeInformation281
	if msg.data.LocalInstrument != "" {
		err := pacs009.LocalInstrumentFedwireFunds1(msg.data.LocalInstrument).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "LocalInstrument",
				Message:   err.Error(),
			}
		}
		Prtry := pacs009.LocalInstrumentFedwireFunds1(msg.data.LocalInstrument)
		PmtTpInf.LclInstrm = pacs009.LocalInstrument2Choice1{
			Prtry: &Prtry,
		}
	}
	if !isEmpty(PmtTpInf) {
		CdtTrfTxInf.PmtTpInf = PmtTpInf
	}
	if !isEmpty(msg.data.InterbankSettlementAmount) {
		err := pacs009.ActiveCurrencyAndAmountFedwire1SimpleType(msg.data.InterbankSettlementAmount.Amount).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InterbankSettlementAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = pacs009.ActiveCurrencyCodeFixed(msg.data.InterbankSettlementAmount.Currency).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InterbankSettlementAmount.Currency",
				Message:   err.Error(),
			}
		}
		CdtTrfTxInf.IntrBkSttlmAmt = pacs009.ActiveCurrencyAndAmountFedwire1{
			Value: pacs009.ActiveCurrencyAndAmountFedwire1SimpleType(msg.data.InterbankSettlementAmount.Amount),
			Ccy:   pacs009.ActiveCurrencyCodeFixed(msg.data.InterbankSettlementAmount.Currency),
		}
	}
	if !isEmpty(msg.data.InterbankSettlementDate) {
		err := msg.data.InterbankSettlementDate.Date().Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InterbankSettlementDate",
				Message:   err.Error(),
			}
		}
		CdtTrfTxInf.IntrBkSttlmDt = msg.data.InterbankSettlementDate.Date()
	}
	if !isEmpty(msg.data.InstructingAgent) {
		err := pacs009.ExternalClearingSystemIdentification1CodeFixed(msg.data.InstructingAgent.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructingAgent.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = pacs009.RoutingNumberFRS1(msg.data.InstructingAgent.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructingAgent.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		Cd := pacs009.ExternalClearingSystemIdentification1CodeFixed(msg.data.InstructingAgent.PaymentSysCode)
		CdtTrfTxInf.InstgAgt = pacs009.BranchAndFinancialInstitutionIdentification62{
			FinInstnId: pacs009.FinancialInstitutionIdentification182{
				ClrSysMmbId: pacs009.ClearingSystemMemberIdentification22{
					ClrSysId: pacs009.ClearingSystemIdentification2Choice2{
						Cd: &Cd,
					},
					MmbId: pacs009.RoutingNumberFRS1(msg.data.InstructingAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(msg.data.InstructedAgent) {
		err := pacs009.ExternalClearingSystemIdentification1CodeFixed(msg.data.InstructedAgent.PaymentSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructedAgent.PaymentSysCode",
				Message:   err.Error(),
			}
		}
		err = pacs009.RoutingNumberFRS1(msg.data.InstructedAgent.PaymentSysMemberId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructedAgent.PaymentSysMemberId",
				Message:   err.Error(),
			}
		}
		Cd := pacs009.ExternalClearingSystemIdentification1CodeFixed(msg.data.InstructedAgent.PaymentSysCode)
		CdtTrfTxInf.InstdAgt = pacs009.BranchAndFinancialInstitutionIdentification62{
			FinInstnId: pacs009.FinancialInstitutionIdentification182{
				ClrSysMmbId: pacs009.ClearingSystemMemberIdentification22{
					ClrSysId: pacs009.ClearingSystemIdentification2Choice2{
						Cd: &Cd,
					},
					MmbId: pacs009.RoutingNumberFRS1(msg.data.InstructedAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(msg.data.Debtor) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.data.Debtor.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(msg.data.Debtor.BusinessId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Debtor.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(msg.data.Debtor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.data.Debtor.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(msg.data.Debtor.ClearingSystemId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Debtor.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.data.Debtor.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.data.Debtor.ClearintSystemMemberId != "" {
			err := pacs009.Max35Text(msg.data.Debtor.ClearintSystemMemberId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Debtor.ClearintSystemMemberId",
					Message:   err.Error(),
				}
			}
			if !isEmpty(finialialId.ClrSysMmbId) {
				finialialId.ClrSysMmbId.MmbId = pacs009.Max35Text(msg.data.Debtor.ClearintSystemMemberId)
			}
		}
		if msg.data.Debtor.Name != "" {
			err := pacs009.Max140Text(msg.data.Debtor.Name).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Debtor.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(msg.data.Debtor.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.data.Debtor.Address) {
			PstlAdr, vErr := PostalAddress241From(msg.data.Debtor.Address)
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
		if !isEmpty(msg.data.UnderlyingCustomerCreditTransfer) {
			UndrlygCstmrCdtTrf, vErr := CreditTransferTransaction371From(msg.data.UnderlyingCustomerCreditTransfer)
			if vErr != nil {
				vErr.InsertPath("UnderlyingCustomerCreditTransfer")
				return vErr
			}
			if !isEmpty(UndrlygCstmrCdtTrf) {
				CdtTrfTxInf.UndrlygCstmrCdtTrf = &UndrlygCstmrCdtTrf

			}
		}
	}
	if !isEmpty(msg.data.DebtorAgent) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.data.DebtorAgent.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(msg.data.DebtorAgent.BusinessId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "DebtorAgent.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(msg.data.DebtorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.data.DebtorAgent.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(msg.data.DebtorAgent.ClearingSystemId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "DebtorAgent.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.data.DebtorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.data.DebtorAgent.ClearintSystemMemberId != "" {
			err := pacs009.Max35Text(msg.data.DebtorAgent.ClearintSystemMemberId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "DebtorAgent.ClearintSystemMemberId",
					Message:   err.Error(),
				}
			}
			finialialId.ClrSysMmbId.MmbId = pacs009.Max35Text(msg.data.DebtorAgent.ClearintSystemMemberId)
		}
		if msg.data.DebtorAgent.Name != "" {
			err := pacs009.Max140Text(msg.data.DebtorAgent.Name).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "DebtorAgent.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(msg.data.DebtorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.data.DebtorAgent.Address) {
			PstlAdr, vErr := PostalAddress241From(msg.data.DebtorAgent.Address)
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
	if !isEmpty(msg.data.CreditorAgent) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.data.CreditorAgent.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(msg.data.CreditorAgent.BusinessId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "CreditorAgent.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(msg.data.CreditorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.data.CreditorAgent.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(msg.data.CreditorAgent.ClearingSystemId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "CreditorAgent.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.data.CreditorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.data.CreditorAgent.ClearintSystemMemberId != "" {
			err := pacs009.Max35Text(msg.data.CreditorAgent.ClearintSystemMemberId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "CreditorAgent.ClearintSystemMemberId",
					Message:   err.Error(),
				}
			}
			finialialId.ClrSysMmbId.MmbId = pacs009.Max35Text(msg.data.CreditorAgent.ClearintSystemMemberId)
		}
		if msg.data.CreditorAgent.Name != "" {
			err := pacs009.Max140Text(msg.data.CreditorAgent.Name).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "CreditorAgent.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(msg.data.CreditorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.data.CreditorAgent.Address) {
			PstlAdr, vErr := PostalAddress241From(msg.data.CreditorAgent.Address)
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
	if !isEmpty(msg.data.Creditor) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.data.Creditor.BusinessId != "" {
			err := pacs009.BICFIDec2014Identifier(msg.data.Creditor.BusinessId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Creditor.BusinessId",
					Message:   err.Error(),
				}
			}
			BICFI := pacs009.BICFIDec2014Identifier(msg.data.Creditor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.data.Creditor.ClearingSystemId != "" {
			err := pacs009.ExternalClearingSystemIdentification1Code(msg.data.Creditor.ClearingSystemId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Creditor.ClearingSystemId",
					Message:   err.Error(),
				}
			}
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.data.Creditor.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.data.Creditor.ClearintSystemMemberId != "" {
			err := pacs009.Max35Text(msg.data.Creditor.ClearintSystemMemberId).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Creditor.ClearintSystemMemberId",
					Message:   err.Error(),
				}
			}
			if !isEmpty(finialialId.ClrSysMmbId) {
				finialialId.ClrSysMmbId.MmbId = pacs009.Max35Text(msg.data.Creditor.ClearintSystemMemberId)
			}
		}
		if msg.data.Creditor.Name != "" {
			err := pacs009.Max140Text(msg.data.Creditor.Name).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "Creditor.Name",
					Message:   err.Error(),
				}
			}
			Nm := pacs009.Max140Text(msg.data.Creditor.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.data.Creditor.Address) {
			PstlAdr, vErr := PostalAddress241From(msg.data.Creditor.Address)
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
	if msg.data.RemittanceInfo != "" {
		err := pacs009.Max140Text(msg.data.RemittanceInfo).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RemittanceInfo",
				Message:   err.Error(),
			}
		}
		Ustrd := pacs009.Max140Text(msg.data.RemittanceInfo)
		RmtInf := pacs009.RemittanceInformation21{
			Ustrd: &Ustrd,
		}
		CdtTrfTxInf.RmtInf = &RmtInf
	}

	if !isEmpty(CdtTrfTxInf) {
		FICdtTrf.CdtTrfTxInf = CdtTrfTxInf
	}
	if !isEmpty(FICdtTrf) {
		msg.doc.FICdtTrf = FICdtTrf
	}
	return nil
}
