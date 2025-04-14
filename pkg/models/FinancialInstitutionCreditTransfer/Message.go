package FinancialInstitutionCreditTransfer

import (
	"encoding/xml"
	"strconv"
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

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = pacs009.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var FICdtTrf pacs009.FinancialInstitutionCreditTransferV08
	var GrpHdr pacs009.GroupHeader931
	if msg.data.MessageId != "" {
		GrpHdr.MsgId = pacs009.IMADFedwireFunds1(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreateDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreateDateTime)
	}
	if msg.data.NumberOfTransactions > 0 {
		GrpHdr.NbOfTxs = pacs009.Max15NumericTextFixed(strconv.Itoa(msg.data.NumberOfTransactions))
	}
	var SttlmInf pacs009.SettlementInstruction71
	if msg.data.SettlementMethod != "" {
		SttlmInf.SttlmMtd = pacs009.SettlementMethod1Code1(msg.data.SettlementMethod)
	}
	if msg.data.ClearingSystem != "" {
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
		InstrId := pacs009.Max35Text(msg.data.PaymentInstructionId)
		PmtId.InstrId = &InstrId
	}
	if msg.data.PaymentEndToEndId != "" {
		PmtId.EndToEndId = pacs009.Max35Text(msg.data.PaymentEndToEndId)
	}
	if msg.data.PaymentUETR != "" {
		PmtId.UETR = pacs009.UUIDv4Identifier(msg.data.PaymentUETR)
	}
	if !isEmpty(PmtId) {
		CdtTrfTxInf.PmtId = PmtId
	}
	var PmtTpInf pacs009.PaymentTypeInformation281
	if msg.data.LocalInstrument != "" {
		Prtry := pacs009.LocalInstrumentFedwireFunds1(msg.data.LocalInstrument)
		PmtTpInf.LclInstrm = pacs009.LocalInstrument2Choice1{
			Prtry: &Prtry,
		}
	}
	if !isEmpty(PmtTpInf) {
		CdtTrfTxInf.PmtTpInf = PmtTpInf
	}
	if !isEmpty(msg.data.InterbankSettlementAmount) {
		CdtTrfTxInf.IntrBkSttlmAmt = pacs009.ActiveCurrencyAndAmountFedwire1{
			Value: pacs009.ActiveCurrencyAndAmountFedwire1SimpleType(msg.data.InterbankSettlementAmount.Amount),
			Ccy:   pacs009.ActiveCurrencyCodeFixed(msg.data.InterbankSettlementAmount.Currency),
		}
	}
	if !isEmpty(msg.data.InterbankSettlementDate) {
		CdtTrfTxInf.IntrBkSttlmDt = msg.data.InterbankSettlementDate.Date()
	}
	if !isEmpty(msg.data.InstructingAgent) {
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
			BICFI := pacs009.BICFIDec2014Identifier(msg.data.Debtor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.data.Debtor.ClearingSystemId != "" {
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.data.Debtor.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.data.Debtor.Name != "" {
			Nm := pacs009.Max140Text(msg.data.Debtor.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.data.Debtor.Address) {
			PstlAdr := PostalAddress241From(msg.data.Debtor.Address)
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(msg.data.UnderlyingCustomerCreditTransfer) {
			UndrlygCstmrCdtTrf := CreditTransferTransaction371From(msg.data.UnderlyingCustomerCreditTransfer)
			if !isEmpty(UndrlygCstmrCdtTrf) {
				CdtTrfTxInf.UndrlygCstmrCdtTrf = &UndrlygCstmrCdtTrf

			}
		}
		if !isEmpty(agent) {
			CdtTrfTxInf.Dbtr = agent
		}
	}
	if !isEmpty(msg.data.DebtorAgent) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.data.DebtorAgent.BusinessId != "" {
			BICFI := pacs009.BICFIDec2014Identifier(msg.data.DebtorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.data.DebtorAgent.ClearingSystemId != "" {
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.data.DebtorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.data.DebtorAgent.Name != "" {
			Nm := pacs009.Max140Text(msg.data.DebtorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.data.DebtorAgent.Address) {
			PstlAdr := PostalAddress241From(msg.data.DebtorAgent.Address)
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
			BICFI := pacs009.BICFIDec2014Identifier(msg.data.CreditorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.data.CreditorAgent.ClearingSystemId != "" {
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.data.CreditorAgent.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.data.CreditorAgent.Name != "" {
			Nm := pacs009.Max140Text(msg.data.CreditorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.data.CreditorAgent.Address) {
			PstlAdr := PostalAddress241From(msg.data.CreditorAgent.Address)
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(agent) {
			CdtTrfTxInf.DbtrAgt = &agent
		}
		CdtTrfTxInf.CdtrAgt = &agent
	}
	if !isEmpty(msg.data.Creditor) {
		var agent pacs009.BranchAndFinancialInstitutionIdentification61
		var finialialId pacs009.FinancialInstitutionIdentification181
		if msg.data.Creditor.BusinessId != "" {
			BICFI := pacs009.BICFIDec2014Identifier(msg.data.Creditor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.data.Creditor.ClearingSystemId != "" {
			Cd := pacs009.ExternalClearingSystemIdentification1Code(msg.data.Creditor.ClearingSystemId)
			ClrSysMmbId := pacs009.ClearingSystemMemberIdentification21{
				ClrSysId: pacs009.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.data.Creditor.Name != "" {
			Nm := pacs009.Max140Text(msg.data.Creditor.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.data.Creditor.Address) {
			PstlAdr := PostalAddress241From(msg.data.Creditor.Address)
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(agent) {
			CdtTrfTxInf.DbtrAgt = &agent
		}
		CdtTrfTxInf.Cdtr = agent
	}
	if msg.data.RemittanceInfo != "" {
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
}
