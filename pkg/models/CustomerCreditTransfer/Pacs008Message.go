package CustomerCreditTransfer

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	pacs008 "github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer_pacs_008_001_08"
	fedwire "github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageModel struct {
	//MessageId (Message Identification) is a unique identifier assigned to an entire message.
	MessageId string
	//CreatedDateTime represents the timestamp when a message, instruction, or transaction was created
	//ISO 8601 format
	CreatedDateTime time.Time
	//NbOfTxs (Number of Transactions) represents the total count of individual payment transactions contained within a financial message.
	// default value: 1
	NumberOfTransactions int
	//SttlmMtd (Settlement Method) specifies how the payment settlement is executed between financial institutions.
	//default value: CLRG
	SettlementMethod model.SettlementMethodType
	//CommonClearingSysCode is stands for Code, which represents the Clearing System Code used for settling the payment.
	//default value: FDW
	CommonClearingSysCode model.CommonClearingSysCodeType
	// InstructionId is a unique identifier assigned to a specific payment instruction within a message.
	InstructionId string
	//EndToEndId is Identifies a payment from sender to receiver across the entire payment chain.
	EndToEndId string
	//UniqueETETransactionRef is stands for Unique End-to-End Transaction Reference. It is a unique identifier that is used to track and identify a payment transaction throughout its entire lifecycle, from initiation to completion.
	UniqueEndToEndTransactionRef string
	//service level code
	SericeLevel string
	// A proprietary code for the local instrument.
	//default value: CTRC
	InstrumentPropCode model.InstrumentPropCodeType
	//Interbank Settlement Amount. It represents the total amount that will be settled between banks as part of a financial transaction.
	InterBankSettAmount model.CurrencyAndAmount
	//<IntrBkSttlmDt> stands for Interbank Settlement Date. It refers to the date on which the interbank settlement of the payment will occur.
	// default: current date
	InterBankSettDate civil.Date
	//stands for Instructed Amount, which represents the amount that the sender has instructed to be transferred in a payment transaction.
	InstructedAmount model.CurrencyAndAmount

	exchangeRate float64
	//Charge Bearer. It specifies who is responsible for paying the charges (fees) associated with the transaction.
	//default value: SLEV
	ChargeBearer ChargeBearerType

	ChargesInfo []ChargeInfo
	// Instructing Agent is  This is the financial institution or bank that is instructing the payment transaction to be processed.
	InstructingAgents model.Agent
	// InstructedAgent is the financial institution or bank that is receiving the payment instruction from the Instructing Agent (the bank sending the payment).
	InstructedAgent      model.Agent
	IntermediaryAgent1Id string
	//The <UltmtDbtr> (Ultimate Debtor) is an optional element in financial transactions, particularly in ISO 20022 payment messages (such as PACS.008 or PACS.009).
	UltimateDebtorName    string
	UltimateDebtorAddress model.PostalAddress
	//DebtorName represent the name of the debtor. This could be an individual person, a company, or any other legal entity initiating the payment.
	DebtorName string
	//DebtorAddress is postal address of the debtor (the party making the payment).
	DebtorAddress model.PostalAddress
	// standardized international format for bank account numbers used to facilitate cross-border payments.
	DebtorIBAN string
	//other types of identification systems for the account, which can vary by region or financial institution.
	DebtorOtherTypeId string
	//refers to the debtor’s agent or the debtor’s bank. This is the financial institution that is responsible for processing the payment on behalf of the debtor (the party making the payment).
	DebtorAgent model.Agent
	//Represents the creditor's bank or agent that is responsible for receiving the payment on behalf of the creditor.
	CreditorAgent model.Agent
	//name of the creditor, which is the entity (person or organization) receiving the payment.
	CreditorName string
	//Postal Address of the Creditor
	CreditorPostalAddress model.PostalAddress

	UltimateCreditorName    string
	UltimateCreditorAddress model.PostalAddress
	//element holds the actual identifier (e.g., an account number or other form of account ID) for the creditor's account.
	CreditorIBAN        string
	CreditorOtherTypeId string
	PurposeOfPayment    PurposeOfPaymentType
	//(Related Remittance Information) is a field in ISO 20022 payment messages that links a payment to related remittance details.
	RelatedRemittanceInfo RemittanceDetail
	//Remittance Information. It provides detailed information related to a payment, typically describing what the payment is for.
	RemittanceInfor RemittanceDocument
}

type Message struct {
	data MessageModel
	doc  pacs008.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}

func (msg *Message) CreateDocument() {
	// Initialize variables
	var SttlmInf_ClrSys_Cd pacs008.ExternalCashClearingSystem1CodeFixed
	var CdtTrfTxInf_PmtId_InstrId pacs008.Max35Text
	var InstgAgt_FinInstnId_ClrSysId pacs008.ExternalClearingSystemIdentification1CodeFixed
	var InstdAgt_FinInstnId_ClrSysId pacs008.ExternalClearingSystemIdentification1CodeFixed
	var DbtrAcct pacs008.CashAccount38
	var Cdtr_Nm pacs008.Max140Text
	var Cdtr_PstlAdr pacs008.PostalAddress241
	var CdtrAcct pacs008.CashAccount38
	var RltdRmtInf pacs008.RemittanceLocation71
	var RmtInf pacs008.RemittanceInformation161
	var CdtTrfTxInf_Purp pacs008.Purpose2Choice
	var charges71List []*pacs008.Charges71

	// Check each field for non-empty values and set accordingly

	if msg.data.CommonClearingSysCode != "" {
		SttlmInf_ClrSys_Cd = pacs008.ExternalCashClearingSystem1CodeFixed(msg.data.CommonClearingSysCode)
	}

	if msg.data.InstructionId != "" {
		CdtTrfTxInf_PmtId_InstrId = pacs008.Max35Text(msg.data.InstructionId)
	}

	for _, charge := range msg.data.ChargesInfo {
		converted := Charges71From(charge)
		if !isEmpty(converted) {
			charges71List = append(charges71List, &converted)
		}
	}

	if msg.data.InstructingAgents.PaymentSysCode != "" {
		InstgAgt_FinInstnId_ClrSysId = pacs008.ExternalClearingSystemIdentification1CodeFixed(msg.data.InstructingAgents.PaymentSysCode)
	}

	if msg.data.InstructedAgent.PaymentSysCode != "" {
		InstdAgt_FinInstnId_ClrSysId = pacs008.ExternalClearingSystemIdentification1CodeFixed(msg.data.InstructedAgent.PaymentSysCode)
	}

	if msg.data.DebtorIBAN != "" || msg.data.DebtorOtherTypeId != "" {
		DbtrAcct = CashAccount38From(msg.data.DebtorIBAN, msg.data.DebtorOtherTypeId)
	}
	if msg.data.CreditorName != "" {
		Cdtr_Nm = pacs008.Max140Text(msg.data.CreditorName)
	}
	_Cdtr_PstlAdr := PostalAddress241From(msg.data.CreditorPostalAddress)
	if !isEmptyPostalAddress241(_Cdtr_PstlAdr) {
		Cdtr_PstlAdr = _Cdtr_PstlAdr
	}

	if msg.data.CreditorIBAN != "" || msg.data.CreditorOtherTypeId != "" {
		CdtrAcct = CashAccount38From(msg.data.CreditorIBAN, msg.data.CreditorOtherTypeId)
	}

	_RltdRmtInf := RemittanceLocation71From(msg.data.RelatedRemittanceInfo)
	if !isEmpty(_RltdRmtInf) {
		RltdRmtInf = _RltdRmtInf
	}

	_RmtInf := RemittanceInformation161From(msg.data.RemittanceInfor)
	if !isEmpty(_RmtInf) {
		RmtInf = _RmtInf
	}
	CdtTrfTxInf_UltimateDbtr := PartyIdentification1351From(msg.data.UltimateDebtorName, msg.data.UltimateDebtorAddress)
	CdtTrfTxInf_Dbtr := PartyIdentification1352From(msg.data.DebtorName, msg.data.DebtorAddress)
	DbtrAgt_FinInstnId := FinancialInstitutionIdentification181From(msg.data.DebtorAgent)
	CdtTrfTxInf_UltimateCdtr := PartyIdentification1351From(msg.data.UltimateCreditorName, msg.data.UltimateCreditorAddress)
	CdtrAgt_FinInstnId := FinancialInstitutionIdentification181From(msg.data.CreditorAgent)
	CdtTrfTxInf_PmtTpInf := PaymentTypeInformation281From(msg.data.InstrumentPropCode, msg.data.SericeLevel)
	// Construct the Document structure
	msg.doc = pacs008.Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08",
			Local: "Document",
		},
		FIToFICstmrCdtTrf: pacs008.FIToFICustomerCreditTransferV08{
			GrpHdr: pacs008.GroupHeader931{
				MsgId:   pacs008.IMADFedwireFunds1(msg.data.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.data.CreatedDateTime),
				NbOfTxs: pacs008.Max15NumericTextFixed(strconv.Itoa(msg.data.NumberOfTransactions)),
				SttlmInf: pacs008.SettlementInstruction71{
					SttlmMtd: pacs008.SettlementMethod1Code1(msg.data.SettlementMethod),
					ClrSys: pacs008.ClearingSystemIdentification3Choice1{
						Cd: &SttlmInf_ClrSys_Cd,
					},
				},
			},
			CdtTrfTxInf: pacs008.CreditTransferTransaction391{
				PmtId: pacs008.PaymentIdentification71{
					InstrId:    &CdtTrfTxInf_PmtId_InstrId,
					EndToEndId: pacs008.Max35Text(msg.data.EndToEndId),
					UETR:       pacs008.UUIDv4Identifier(msg.data.UniqueEndToEndTransactionRef),
				},
				PmtTpInf: CdtTrfTxInf_PmtTpInf,
				IntrBkSttlmAmt: pacs008.ActiveCurrencyAndAmountFedwire1{
					Value: pacs008.ActiveCurrencyAndAmountFedwire1SimpleType(msg.data.InterBankSettAmount.Amount),
					Ccy:   pacs008.ActiveCurrencyCodeFixed(msg.data.InterBankSettAmount.Currency),
				},
				IntrBkSttlmDt: fedwire.ISODate(msg.data.InterBankSettDate),
				InstdAmt: pacs008.ActiveOrHistoricCurrencyAndAmount{
					Value: pacs008.ActiveOrHistoricCurrencyAndAmountSimpleType(msg.data.InstructedAmount.Amount),
					Ccy:   pacs008.ActiveOrHistoricCurrencyCode(msg.data.InstructedAmount.Currency),
				},
				ChrgBr: pacs008.ChargeBearerType1Code(msg.data.ChargeBearer),
				InstgAgt: pacs008.BranchAndFinancialInstitutionIdentification62{
					FinInstnId: pacs008.FinancialInstitutionIdentification182{
						ClrSysMmbId: pacs008.ClearingSystemMemberIdentification22{
							ClrSysId: pacs008.ClearingSystemIdentification2Choice2{
								Cd: &InstgAgt_FinInstnId_ClrSysId,
							},
							MmbId: pacs008.RoutingNumberFRS1(msg.data.InstructingAgents.PaymentSysMemberId),
						},
					},
				},
				InstdAgt: pacs008.BranchAndFinancialInstitutionIdentification62{
					FinInstnId: pacs008.FinancialInstitutionIdentification182{
						ClrSysMmbId: pacs008.ClearingSystemMemberIdentification22{
							ClrSysId: pacs008.ClearingSystemIdentification2Choice2{
								Cd: &InstdAgt_FinInstnId_ClrSysId,
							},
							MmbId: pacs008.RoutingNumberFRS1(msg.data.InstructedAgent.PaymentSysMemberId),
						},
					},
				},
				DbtrAcct: &DbtrAcct,
				DbtrAgt: pacs008.BranchAndFinancialInstitutionIdentification61{
					FinInstnId: DbtrAgt_FinInstnId,
				},
				CdtrAgt: pacs008.BranchAndFinancialInstitutionIdentification63{
					FinInstnId: CdtrAgt_FinInstnId,
				},
				Cdtr: pacs008.PartyIdentification1352{
					Nm:      &Cdtr_Nm,
					PstlAdr: &Cdtr_PstlAdr,
				},
				CdtrAcct: &CdtrAcct,
			},
		},
	}
	if len(charges71List) > 0 {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.ChrgsInf = charges71List
	}

	if msg.data.exchangeRate != 0 {
		_exchangeRate := pacs008.BaseOneRate(msg.data.exchangeRate)
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.XchgRate = &_exchangeRate
	}
	if msg.data.IntermediaryAgent1Id != "" {
		_IntrmyAgt1 := BranchAndFinancialInstitutionIdentification61From(msg.data.IntermediaryAgent1Id)
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrmyAgt1 = &_IntrmyAgt1
	}

	if !isEmpty(CdtTrfTxInf_UltimateDbtr) {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr = &CdtTrfTxInf_UltimateDbtr
	}
	if !isEmpty(CdtTrfTxInf_Dbtr) {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr = CdtTrfTxInf_Dbtr
	}
	if !isEmpty(CdtTrfTxInf_UltimateCdtr) {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr = &CdtTrfTxInf_UltimateCdtr
	}
	if msg.data.PurposeOfPayment != "" {
		_Cd := pacs008.ExternalPurpose1Code(InvestmentPayment)
		CdtTrfTxInf_Purp = pacs008.Purpose2Choice{
			Cd: &_Cd,
		}
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Purp = &CdtTrfTxInf_Purp
	}
	if !isEmpty(RltdRmtInf) {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RltdRmtInf = &RltdRmtInf
	}
	if !isEmpty(RmtInf) {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf = &RmtInf
	}
}
func (msg *Message) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	xmlData, err := xml.MarshalIndent(msg.doc, "", "\t")
	if err != nil {
		return err
	}

	// Convert byte slice to string for manipulation
	xmlString := string(xmlData)

	// Keep the xmlns only in the <Document> tag, remove from others
	xmlString = removeExtraXMLNS(xmlString)

	// Convert back to []byte
	return e.EncodeToken(xml.CharData([]byte(xmlString)))
	// return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *Message) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(msg.doc.FIToFICstmrCdtTrf, "", " ")
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
		strings.ReplaceAll(xmlStr[docEnd+1:], ` xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"`, "")

	return cleanXML
}
