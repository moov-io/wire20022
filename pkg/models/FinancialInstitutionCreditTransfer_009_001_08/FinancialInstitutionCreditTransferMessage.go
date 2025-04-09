package FinancialInstitutionCreditTransfer_009_001_08

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	FinancialInstitutionCreditTransfer "github.com/moov-io/fedwire20022/gen/FinancialInstitutionCreditTransfer_pacs_009_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Pacs009 struct {
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
	InterbankSettlementDate civil.Date
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
type Pacs009Message struct {
	model Pacs009
	doc   FinancialInstitutionCreditTransfer.Document
}

func NewPacs009Message() Pacs009Message {
	return Pacs009Message{
		model: Pacs009{},
	}
}
func (msg *Pacs009Message) CreateDocument() {
	msg.doc = FinancialInstitutionCreditTransfer.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08",
			Local: "Document",
		},
	}
	var FICdtTrf FinancialInstitutionCreditTransfer.FinancialInstitutionCreditTransferV08
	var GrpHdr FinancialInstitutionCreditTransfer.GroupHeader931
	if msg.model.MessageId != "" {
		GrpHdr.MsgId = FinancialInstitutionCreditTransfer.IMADFedwireFunds1(msg.model.MessageId)
	}
	if !isEmpty(msg.model.CreateDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.model.CreateDateTime)
	}
	if msg.model.NumberOfTransactions > 0 {
		GrpHdr.NbOfTxs = FinancialInstitutionCreditTransfer.Max15NumericTextFixed(strconv.Itoa(msg.model.NumberOfTransactions))
	}
	var SttlmInf FinancialInstitutionCreditTransfer.SettlementInstruction71
	if msg.model.SettlementMethod != "" {
		SttlmInf.SttlmMtd = FinancialInstitutionCreditTransfer.SettlementMethod1Code1(msg.model.SettlementMethod)
	}
	if msg.model.ClearingSystem != "" {
		Cd := FinancialInstitutionCreditTransfer.ExternalCashClearingSystem1CodeFixed(msg.model.ClearingSystem)
		SttlmInf.ClrSys = FinancialInstitutionCreditTransfer.ClearingSystemIdentification3Choice1{
			Cd: &Cd,
		}
	}
	if !isEmpty(SttlmInf) {
		GrpHdr.SttlmInf = SttlmInf
	}
	if !isEmpty(GrpHdr) {
		FICdtTrf.GrpHdr = GrpHdr
	}
	var CdtTrfTxInf FinancialInstitutionCreditTransfer.CreditTransferTransaction361
	var PmtId FinancialInstitutionCreditTransfer.PaymentIdentification71

	if msg.model.PaymentInstructionId != "" {
		InstrId := FinancialInstitutionCreditTransfer.Max35Text(msg.model.PaymentInstructionId)
		PmtId.InstrId = &InstrId
	}
	if msg.model.PaymentEndToEndId != "" {
		PmtId.EndToEndId = FinancialInstitutionCreditTransfer.Max35Text(msg.model.PaymentEndToEndId)
	}
	if msg.model.PaymentUETR != "" {
		PmtId.UETR = FinancialInstitutionCreditTransfer.UUIDv4Identifier(msg.model.PaymentUETR)
	}
	if !isEmpty(PmtId) {
		CdtTrfTxInf.PmtId = PmtId
	}
	var PmtTpInf FinancialInstitutionCreditTransfer.PaymentTypeInformation281
	if msg.model.LocalInstrument != "" {
		Prtry := FinancialInstitutionCreditTransfer.LocalInstrumentFedwireFunds1(msg.model.LocalInstrument)
		PmtTpInf.LclInstrm = FinancialInstitutionCreditTransfer.LocalInstrument2Choice1{
			Prtry: &Prtry,
		}
	}
	if !isEmpty(PmtTpInf) {
		CdtTrfTxInf.PmtTpInf = PmtTpInf
	}
	if !isEmpty(msg.model.InterbankSettlementAmount) {
		CdtTrfTxInf.IntrBkSttlmAmt = FinancialInstitutionCreditTransfer.ActiveCurrencyAndAmountFedwire1{
			Value: FinancialInstitutionCreditTransfer.ActiveCurrencyAndAmountFedwire1SimpleType(msg.model.InterbankSettlementAmount.Amount),
			Ccy:   FinancialInstitutionCreditTransfer.ActiveCurrencyCodeFixed(msg.model.InterbankSettlementAmount.Currency),
		}
	}
	if !isEmpty(msg.model.InterbankSettlementDate) {
		CdtTrfTxInf.IntrBkSttlmDt = fedwire.ISODate(msg.model.InterbankSettlementDate)
	}
	if !isEmpty(msg.model.InstructingAgent) {
		Cd := FinancialInstitutionCreditTransfer.ExternalClearingSystemIdentification1CodeFixed(msg.model.InstructingAgent.PaymentSysCode)
		CdtTrfTxInf.InstgAgt = FinancialInstitutionCreditTransfer.BranchAndFinancialInstitutionIdentification62{
			FinInstnId: FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification182{
				ClrSysMmbId: FinancialInstitutionCreditTransfer.ClearingSystemMemberIdentification22{
					ClrSysId: FinancialInstitutionCreditTransfer.ClearingSystemIdentification2Choice2{
						Cd: &Cd,
					},
					MmbId: FinancialInstitutionCreditTransfer.RoutingNumberFRS1(msg.model.InstructingAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(msg.model.InstructedAgent) {
		Cd := FinancialInstitutionCreditTransfer.ExternalClearingSystemIdentification1CodeFixed(msg.model.InstructedAgent.PaymentSysCode)
		CdtTrfTxInf.InstdAgt = FinancialInstitutionCreditTransfer.BranchAndFinancialInstitutionIdentification62{
			FinInstnId: FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification182{
				ClrSysMmbId: FinancialInstitutionCreditTransfer.ClearingSystemMemberIdentification22{
					ClrSysId: FinancialInstitutionCreditTransfer.ClearingSystemIdentification2Choice2{
						Cd: &Cd,
					},
					MmbId: FinancialInstitutionCreditTransfer.RoutingNumberFRS1(msg.model.InstructedAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(msg.model.Debtor) {
		var agent FinancialInstitutionCreditTransfer.BranchAndFinancialInstitutionIdentification61
		var finialialId FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification181
		if msg.model.Debtor.BusinessId != "" {
			BICFI := FinancialInstitutionCreditTransfer.BICFIDec2014Identifier(msg.model.Debtor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.model.Debtor.ClearingSystemId != "" {
			Cd := FinancialInstitutionCreditTransfer.ExternalClearingSystemIdentification1Code(msg.model.Debtor.ClearingSystemId)
			ClrSysMmbId := FinancialInstitutionCreditTransfer.ClearingSystemMemberIdentification21{
				ClrSysId: FinancialInstitutionCreditTransfer.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.model.Debtor.Name != "" {
			Nm := FinancialInstitutionCreditTransfer.Max140Text(msg.model.Debtor.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.model.Debtor.Address) {
			PstlAdr := PostalAddress241From(msg.model.Debtor.Address)
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(msg.model.UnderlyingCustomerCreditTransfer) {
			UndrlygCstmrCdtTrf := CreditTransferTransaction371From(msg.model.UnderlyingCustomerCreditTransfer)
			if !isEmpty(UndrlygCstmrCdtTrf) {
				CdtTrfTxInf.UndrlygCstmrCdtTrf = &UndrlygCstmrCdtTrf

			}
		}
		if !isEmpty(agent) {
			CdtTrfTxInf.Dbtr = agent
		}
	}
	if !isEmpty(msg.model.DebtorAgent) {
		var agent FinancialInstitutionCreditTransfer.BranchAndFinancialInstitutionIdentification61
		var finialialId FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification181
		if msg.model.DebtorAgent.BusinessId != "" {
			BICFI := FinancialInstitutionCreditTransfer.BICFIDec2014Identifier(msg.model.DebtorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.model.DebtorAgent.ClearingSystemId != "" {
			Cd := FinancialInstitutionCreditTransfer.ExternalClearingSystemIdentification1Code(msg.model.DebtorAgent.ClearingSystemId)
			ClrSysMmbId := FinancialInstitutionCreditTransfer.ClearingSystemMemberIdentification21{
				ClrSysId: FinancialInstitutionCreditTransfer.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.model.DebtorAgent.Name != "" {
			Nm := FinancialInstitutionCreditTransfer.Max140Text(msg.model.DebtorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.model.DebtorAgent.Address) {
			PstlAdr := PostalAddress241From(msg.model.DebtorAgent.Address)
			finialialId.PstlAdr = &PstlAdr
		}
		if !isEmpty(finialialId) {
			agent.FinInstnId = finialialId
		}
		if !isEmpty(agent) {
			CdtTrfTxInf.DbtrAgt = &agent
		}
	}
	if !isEmpty(msg.model.CreditorAgent) {
		var agent FinancialInstitutionCreditTransfer.BranchAndFinancialInstitutionIdentification61
		var finialialId FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification181
		if msg.model.CreditorAgent.BusinessId != "" {
			BICFI := FinancialInstitutionCreditTransfer.BICFIDec2014Identifier(msg.model.CreditorAgent.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.model.CreditorAgent.ClearingSystemId != "" {
			Cd := FinancialInstitutionCreditTransfer.ExternalClearingSystemIdentification1Code(msg.model.CreditorAgent.ClearingSystemId)
			ClrSysMmbId := FinancialInstitutionCreditTransfer.ClearingSystemMemberIdentification21{
				ClrSysId: FinancialInstitutionCreditTransfer.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.model.CreditorAgent.Name != "" {
			Nm := FinancialInstitutionCreditTransfer.Max140Text(msg.model.CreditorAgent.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.model.CreditorAgent.Address) {
			PstlAdr := PostalAddress241From(msg.model.CreditorAgent.Address)
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
	if !isEmpty(msg.model.Creditor) {
		var agent FinancialInstitutionCreditTransfer.BranchAndFinancialInstitutionIdentification61
		var finialialId FinancialInstitutionCreditTransfer.FinancialInstitutionIdentification181
		if msg.model.Creditor.BusinessId != "" {
			BICFI := FinancialInstitutionCreditTransfer.BICFIDec2014Identifier(msg.model.Creditor.BusinessId)
			finialialId.BICFI = &BICFI
		}
		if msg.model.Creditor.ClearingSystemId != "" {
			Cd := FinancialInstitutionCreditTransfer.ExternalClearingSystemIdentification1Code(msg.model.Creditor.ClearingSystemId)
			ClrSysMmbId := FinancialInstitutionCreditTransfer.ClearingSystemMemberIdentification21{
				ClrSysId: FinancialInstitutionCreditTransfer.ClearingSystemIdentification2Choice1{
					Cd: &Cd,
				},
			}
			finialialId.ClrSysMmbId = &ClrSysMmbId
		}
		if msg.model.Creditor.Name != "" {
			Nm := FinancialInstitutionCreditTransfer.Max140Text(msg.model.Creditor.Name)
			finialialId.Nm = &Nm
		}
		if !isEmpty(msg.model.Creditor.Address) {
			PstlAdr := PostalAddress241From(msg.model.Creditor.Address)
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
	if msg.model.RemittanceInfo != "" {
		Ustrd := FinancialInstitutionCreditTransfer.Max140Text(msg.model.RemittanceInfo)
		RmtInf := FinancialInstitutionCreditTransfer.RemittanceInformation21{
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
func (msg *Pacs009Message) GetXML() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(msg.doc, "", "\t")
	if err != nil {
		return nil, err
	}

	// Convert byte slice to string for manipulation
	xmlString := string(xmlData)

	// Keep the xmlns only in the <Document> tag, remove from others
	xmlString = removeExtraXMLNS(xmlString)

	// Regular expression to match scientific notation (e.g., 9.93229443e+06)
	re := regexp.MustCompile(`>(\d+\.\d+(?:e[+-]?\d+)?|\d+e[+-]?\d+)<`)

	// Replace matched numbers with properly formatted ones
	xmlString = re.ReplaceAllStringFunc(xmlString, func(match string) string {
		// Extract the number inside the tags
		numberStr := strings.Trim(match, "<>")

		// Convert to float
		number, err := strconv.ParseFloat(numberStr, 64)
		if err != nil {
			return match // Return the original string if conversion fails
		}

		// Format it as a standard decimal number with 2 decimal places
		return fmt.Sprintf(">%.2f<", number)
	})

	re = regexp.MustCompile(`<(FrSeq|ToSeq)>(\d+)</(FrSeq|ToSeq)>`)

	// Replace numeric values with zero-padded format (6 digits)
	xmlString = re.ReplaceAllStringFunc(xmlString, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if len(parts) == 4 {
			num := parts[2] // Extract number as string
			return fmt.Sprintf("<%s>%06s</%s>", parts[1], num, parts[3])
		}
		return match
	})

	// Convert back to []byte
	return []byte(xmlString), nil
	// return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *Pacs009Message) GetJson() ([]byte, error) {
	return json.MarshalIndent(msg.doc.FICdtTrf, "", " ")
}

func removeExtraXMLNS(xmlStr string) string {
	// Find the first occurrence of <Document ...> (keep this)
	docStart := strings.Index(xmlStr, "<Document")
	if docStart == -1 {
		return xmlStr // Return original if <Document> not found
	}

	// Find the end of the <Document> opening tag
	docEnd := strings.Index(xmlStr[docStart:], ">")
	if docEnd == -1 {
		return xmlStr
	}
	docEnd += docStart // Adjust index

	// Remove all occurrences of xmlns="..." except in <Document>
	cleanXML := xmlStr[:docEnd+1] + // Keep <Document> with its xmlns
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08"`, "")

	return cleanXML
}
