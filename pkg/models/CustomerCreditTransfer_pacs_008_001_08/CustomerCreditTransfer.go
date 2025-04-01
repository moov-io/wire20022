package CustomerCreditTransfer_pacs_008_001_08

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
	fedwire "github.com/moov-io/wire20022/pkg/internal"
)

type CustomerCreditTransfer struct {
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
	SettlementMethod SettlementMethodType
	//CommonClearingSysCode is stands for Code, which represents the Clearing System Code used for settling the payment.
	//default value: FDW
	CommonClearingSysCode CommonClearingSysCodeType
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
	InstrumentPropCode InstrumentPropCodeType
	//Interbank Settlement Amount. It represents the total amount that will be settled between banks as part of a financial transaction.
	InterBankSettAmount CurrencyAndAmount
	//<IntrBkSttlmDt> stands for Interbank Settlement Date. It refers to the date on which the interbank settlement of the payment will occur.
	// default: current date
	InterBankSettDate civil.Date
	//stands for Instructed Amount, which represents the amount that the sender has instructed to be transferred in a payment transaction.
	InstructedAmount CurrencyAndAmount

	exchangeRate float64
	//Charge Bearer. It specifies who is responsible for paying the charges (fees) associated with the transaction.
	//default value: SLEV
	ChargeBearer ChargeBearerType

	ChargesInfo []ChargeInfo
	// Instructing Agent is  This is the financial institution or bank that is instructing the payment transaction to be processed.
	InstructingAgents Agent
	// InstructedAgent is the financial institution or bank that is receiving the payment instruction from the Instructing Agent (the bank sending the payment).
	InstructedAgent      Agent
	IntermediaryAgent1Id string
	//The <UltmtDbtr> (Ultimate Debtor) is an optional element in financial transactions, particularly in ISO 20022 payment messages (such as PACS.008 or PACS.009).
	UltimateDebtorName    string
	UltimateDebtorAddress PostalAddress
	//DebtorName represent the name of the debtor. This could be an individual person, a company, or any other legal entity initiating the payment.
	DebtorName string
	//DebtorAddress is postal address of the debtor (the party making the payment).
	DebtorAddress PostalAddress
	// standardized international format for bank account numbers used to facilitate cross-border payments.
	DebtorIBAN string
	//other types of identification systems for the account, which can vary by region or financial institution.
	DebtorOtherTypeId string
	//refers to the debtor’s agent or the debtor’s bank. This is the financial institution that is responsible for processing the payment on behalf of the debtor (the party making the payment).
	DebtorAgent Agent
	//Represents the creditor's bank or agent that is responsible for receiving the payment on behalf of the creditor.
	CreditorAgent Agent
	//name of the creditor, which is the entity (person or organization) receiving the payment.
	CreditorName string
	//Postal Address of the Creditor
	CreditorPostalAddress PostalAddress

	UltimateCreditorName    string
	UltimateCreditorAddress PostalAddress
	//element holds the actual identifier (e.g., an account number or other form of account ID) for the creditor's account.
	CreditorIBAN        string
	CreditorOtherTypeId string
	PurposeOfPayment    PurposeOfPaymentType
	//(Related Remittance Information) is a field in ISO 20022 payment messages that links a payment to related remittance details.
	RelatedRemittanceInfo RemittanceDetail
	//Remittance Information. It provides detailed information related to a payment, typically describing what the payment is for.
	RemittanceInfor RemittanceDocument
}

type CustomerCreditTransferMessage struct {
	model CustomerCreditTransfer
	doc   Document
}

func NewCustomerCreditTransferMessage() CustomerCreditTransferMessage {
	return CustomerCreditTransferMessage{
		model: CustomerCreditTransfer{},
	}
}

func (msg *CustomerCreditTransferMessage) CreateDocument() {
	// Initialize variables
	var SttlmInf_ClrSys_Cd ExternalCashClearingSystem1CodeFixed
	var CdtTrfTxInf_PmtId_InstrId Max35Text
	var InstgAgt_FinInstnId_ClrSysId ExternalClearingSystemIdentification1CodeFixed
	var InstdAgt_FinInstnId_ClrSysId ExternalClearingSystemIdentification1CodeFixed
	var DbtrAcct CashAccount38
	var Cdtr_Nm Max140Text
	var Cdtr_PstlAdr PostalAddress241
	var CdtrAcct CashAccount38
	var RltdRmtInf RemittanceLocation71
	var RmtInf RemittanceInformation161
	var CdtTrfTxInf_Purp Purpose2Choice
	var charges71List []*Charges71

	// Check each field for non-empty values and set accordingly

	if msg.model.CommonClearingSysCode != "" {
		SttlmInf_ClrSys_Cd = ExternalCashClearingSystem1CodeFixed(msg.model.CommonClearingSysCode)
	}

	if msg.model.InstructionId != "" {
		CdtTrfTxInf_PmtId_InstrId = Max35Text(msg.model.InstructionId)
	}

	for _, charge := range msg.model.ChargesInfo {
		converted := Charges71From(charge)
		if !converted.isEmpty() {
			charges71List = append(charges71List, &converted)
		}
	}

	if msg.model.InstructingAgents.PaymentSysCode != "" {
		InstgAgt_FinInstnId_ClrSysId = ExternalClearingSystemIdentification1CodeFixed(msg.model.InstructingAgents.PaymentSysCode)
	}

	if msg.model.InstructedAgent.PaymentSysCode != "" {
		InstdAgt_FinInstnId_ClrSysId = ExternalClearingSystemIdentification1CodeFixed(msg.model.InstructedAgent.PaymentSysCode)
	}

	if msg.model.DebtorIBAN != "" || msg.model.DebtorOtherTypeId != "" {
		DbtrAcct = CashAccount38From(msg.model.DebtorIBAN, msg.model.DebtorOtherTypeId)
	}
	if msg.model.CreditorName != "" {
		Cdtr_Nm = Max140Text(msg.model.CreditorName)
	}
	_Cdtr_PstlAdr := PostalAddress241From(msg.model.CreditorPostalAddress)
	if !isEmptyPostalAddress241(_Cdtr_PstlAdr) {
		Cdtr_PstlAdr = _Cdtr_PstlAdr
	}

	if msg.model.CreditorIBAN != "" || msg.model.CreditorOtherTypeId != "" {
		CdtrAcct = CashAccount38From(msg.model.CreditorIBAN, msg.model.CreditorOtherTypeId)
	}

	_RltdRmtInf := RemittanceLocation71From(msg.model.RelatedRemittanceInfo)
	if !_RltdRmtInf.isEmpty() {
		RltdRmtInf = _RltdRmtInf
	}

	_RmtInf := RemittanceInformation161From(msg.model.RemittanceInfor)
	if !_RmtInf.isEmpty() {
		RmtInf = _RmtInf
	}
	CdtTrfTxInf_UltimateDbtr := PartyIdentification1351From(msg.model.UltimateDebtorName, msg.model.UltimateDebtorAddress)
	CdtTrfTxInf_Dbtr := PartyIdentification1352From(msg.model.DebtorName, msg.model.DebtorAddress)
	DbtrAgt_FinInstnId := FinancialInstitutionIdentification181From(msg.model.DebtorAgent)
	CdtTrfTxInf_UltimateCdtr := PartyIdentification1351From(msg.model.UltimateCreditorName, msg.model.UltimateCreditorAddress)
	CdtrAgt_FinInstnId := FinancialInstitutionIdentification181From(msg.model.CreditorAgent)
	CdtTrfTxInf_PmtTpInf := PaymentTypeInformation281From(msg.model.InstrumentPropCode, msg.model.SericeLevel)
	// Construct the Document structure
	msg.doc = Document{
		XMLName: xml.Name{
			Space: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08",
			Local: "Document",
		},
		FIToFICstmrCdtTrf: FIToFICustomerCreditTransferV08{
			GrpHdr: GroupHeader931{
				MsgId:   IMADFedwireFunds1(msg.model.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.model.CreatedDateTime),
				NbOfTxs: Max15NumericTextFixed(strconv.Itoa(msg.model.NumberOfTransactions)),
				SttlmInf: SettlementInstruction71{
					SttlmMtd: SettlementMethod1Code1(msg.model.SettlementMethod),
					ClrSys: ClearingSystemIdentification3Choice1{
						Cd: &SttlmInf_ClrSys_Cd,
					},
				},
			},
			CdtTrfTxInf: CreditTransferTransaction391{
				PmtId: PaymentIdentification71{
					InstrId:    &CdtTrfTxInf_PmtId_InstrId,
					EndToEndId: Max35Text(msg.model.EndToEndId),
					UETR:       UUIDv4Identifier(msg.model.UniqueEndToEndTransactionRef),
				},
				PmtTpInf: CdtTrfTxInf_PmtTpInf,
				IntrBkSttlmAmt: ActiveCurrencyAndAmountFedwire1{
					Value: ActiveCurrencyAndAmountFedwire1SimpleType(msg.model.InterBankSettAmount.Amount),
					Ccy:   ActiveCurrencyCodeFixed(msg.model.InterBankSettAmount.Currency),
				},
				IntrBkSttlmDt: fedwire.ISODate(msg.model.InterBankSettDate),
				InstdAmt: ActiveOrHistoricCurrencyAndAmount{
					Value: ActiveOrHistoricCurrencyAndAmountSimpleType(msg.model.InstructedAmount.Amount),
					Ccy:   ActiveOrHistoricCurrencyCode(msg.model.InstructedAmount.Currency),
				},
				ChrgBr: ChargeBearerType1Code(msg.model.ChargeBearer),
				InstgAgt: BranchAndFinancialInstitutionIdentification62{
					FinInstnId: FinancialInstitutionIdentification182{
						ClrSysMmbId: ClearingSystemMemberIdentification22{
							ClrSysId: ClearingSystemIdentification2Choice2{
								Cd: &InstgAgt_FinInstnId_ClrSysId,
							},
							MmbId: RoutingNumberFRS1(msg.model.InstructingAgents.PaymentSysMemberId),
						},
					},
				},
				InstdAgt: BranchAndFinancialInstitutionIdentification62{
					FinInstnId: FinancialInstitutionIdentification182{
						ClrSysMmbId: ClearingSystemMemberIdentification22{
							ClrSysId: ClearingSystemIdentification2Choice2{
								Cd: &InstdAgt_FinInstnId_ClrSysId,
							},
							MmbId: RoutingNumberFRS1(msg.model.InstructedAgent.PaymentSysMemberId),
						},
					},
				},
				DbtrAcct: &DbtrAcct,
				DbtrAgt: BranchAndFinancialInstitutionIdentification61{
					FinInstnId: DbtrAgt_FinInstnId,
				},
				CdtrAgt: BranchAndFinancialInstitutionIdentification63{
					FinInstnId: CdtrAgt_FinInstnId,
				},
				Cdtr: PartyIdentification1352{
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

	if msg.model.exchangeRate != 0 {
		_exchangeRate := BaseOneRate(msg.model.exchangeRate)
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.XchgRate = &_exchangeRate
	}
	if msg.model.IntermediaryAgent1Id != "" {
		_IntrmyAgt1 := BranchAndFinancialInstitutionIdentification61From(msg.model.IntermediaryAgent1Id)
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrmyAgt1 = &_IntrmyAgt1
	}

	if !CdtTrfTxInf_UltimateDbtr.isEmpty() {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr = &CdtTrfTxInf_UltimateDbtr
	}
	if !CdtTrfTxInf_Dbtr.isEmpty() {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr = CdtTrfTxInf_Dbtr
	}
	if !CdtTrfTxInf_UltimateCdtr.isEmpty() {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr = &CdtTrfTxInf_UltimateCdtr
	}
	if msg.model.PurposeOfPayment != "" {
		_Cd := ExternalPurpose1Code(InvestmentPayment)
		CdtTrfTxInf_Purp = Purpose2Choice{
			Cd: &_Cd,
		}
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Purp = &CdtTrfTxInf_Purp
	}
	if !RltdRmtInf.isEmpty() {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RltdRmtInf = &RltdRmtInf
	}
	if !RmtInf.isEmpty() {
		msg.doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf = &RmtInf
	}
}
func (msg *CustomerCreditTransferMessage) GetXML() ([]byte, error) {
	xmlData, err := xml.MarshalIndent(msg.doc, "", "\t")
	if err != nil {
		return nil, err
	}

	// Convert byte slice to string for manipulation
	xmlString := string(xmlData)

	// Keep the xmlns only in the <Document> tag, remove from others
	xmlString = removeExtraXMLNS(xmlString)

	// Convert back to []byte
	return []byte(xmlString), nil
	// return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *CustomerCreditTransferMessage) GetJson() ([]byte, error) {
	return json.MarshalIndent(msg.doc.FIToFICstmrCdtTrf, "", "\t")
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
