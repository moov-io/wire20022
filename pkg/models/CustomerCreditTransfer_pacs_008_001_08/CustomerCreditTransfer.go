package CustomerCreditTransfer_pacs_008_001_08

import (
	"encoding/json"
	"encoding/xml"
	"strconv"
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
	//Charge Bearer. It specifies who is responsible for paying the charges (fees) associated with the transaction.
	//default value: SLEV
	ChargeBearer ChargeBearerType
	// Instructing Agent is  This is the financial institution or bank that is instructing the payment transaction to be processed.
	InstructingAgents Agent
	// InstructedAgent is the financial institution or bank that is receiving the payment instruction from the Instructing Agent (the bank sending the payment).
	InstructedAgent Agent
	//DebtorName represent the name of the debtor. This could be an individual person, a company, or any other legal entity initiating the payment.
	DebtorName string
	//DebtorAddress is postal address of the debtor (the party making the payment).
	DebtorAddress PostalAddress
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
	//element holds the actual identifier (e.g., an account number or other form of account ID) for the creditor's account.
	CreditorOtherTypeId string
	//Remittance Information. It provides detailed information related to a payment, typically describing what the payment is for.
	RemittanceInfor RemittanceDocument
}

type CustomerCreditTransferMessage struct {
	model CustomerCreditTransfer
	doc   Document
}

func NewCustomerCreditTransferMessage() CustomerCreditTransferMessage {
	return CustomerCreditTransferMessage{
		model: NewCustomerCreditTransfer(),
	}
}

func (msg *CustomerCreditTransferMessage) CreateDocument() {
	SttlmInf_ClrSys_Cd := ExternalCashClearingSystem1CodeFixed(msg.model.CommonClearingSysCode)
	CdtTrfTxInf_PmtId_InstrId := Max35Text(msg.model.InstructionId)
	CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry := LocalInstrumentFedwireFunds1(msg.model.InstrumentPropCode)
	InstgAgt_FinInstnId_ClrSysId := ExternalClearingSystemIdentification1CodeFixed(msg.model.InstructingAgents.PaymentSysCode)
	InstdAgt_FinInstnId_ClrSysId := ExternalClearingSystemIdentification1CodeFixed(msg.model.InstructedAgent.PaymentSysCode)
	Dbtr_Nm := Max140Text(msg.model.DebtorName)
	Dbtr_PstlAdr := PostalAddress241From(msg.model.DebtorAddress)
	DbtrAcct := CashAccount38From(msg.model.DebtorOtherTypeId)
	DbtrAgt_FinInstnId_ClrSysMmbId := ClearingSystemMemberIdentification21From(msg.model.DebtorAgent.PaymentSysCode, msg.model.DebtorAgent.PaymentSysMemberId)
	DbtrAgt_FinInstnId_Nm := Max140Text(msg.model.DebtorAgent.BankName)
	DbtrAgt_FinInstnId_PstlAdr := PostalAddress241From(msg.model.DebtorAgent.PostalAddress)
	CdtrAgt_FinInstnId_ClrSysMmbId := ClearingSystemMemberIdentification21From(msg.model.CreditorAgent.PaymentSysCode, msg.model.CreditorAgent.PaymentSysMemberId)
	CdtrAgt_FinInstnId_Nm := Max140Text(msg.model.CreditorAgent.BankName)
	CdtrAgt_FinInstnId_PstlAdr := PostalAddress241From(msg.model.CreditorAgent.PostalAddress)
	Cdtr_Nm := Max140Text(msg.model.CreditorName)
	Cdtr_PstlAdr := PostalAddress241From(msg.model.CreditorPostalAddress)
	CdtrAcct := CashAccount38From(msg.model.CreditorOtherTypeId)
	RmtInf := RemittanceInformation161From(msg.model.RemittanceInfor)
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
				PmtTpInf: PaymentTypeInformation281{
					LclInstrm: LocalInstrument2Choice1{
						Prtry: &CdtTrfTxInf_PmtTpInf_LclInstrm_Prtry,
					},
				},
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
				Dbtr: PartyIdentification1352{
					Nm:      &Dbtr_Nm,
					PstlAdr: &Dbtr_PstlAdr,
				},
				DbtrAcct: &DbtrAcct,
				DbtrAgt: BranchAndFinancialInstitutionIdentification61{
					FinInstnId: FinancialInstitutionIdentification181{
						ClrSysMmbId: &DbtrAgt_FinInstnId_ClrSysMmbId,
						Nm:          &DbtrAgt_FinInstnId_Nm,
						PstlAdr:     &DbtrAgt_FinInstnId_PstlAdr,
					},
				},
				CdtrAgt: BranchAndFinancialInstitutionIdentification63{
					FinInstnId: FinancialInstitutionIdentification181{
						ClrSysMmbId: &CdtrAgt_FinInstnId_ClrSysMmbId,
						Nm:          &CdtrAgt_FinInstnId_Nm,
						PstlAdr:     &CdtrAgt_FinInstnId_PstlAdr,
					},
				},
				Cdtr: PartyIdentification1352{
					Nm:      &Cdtr_Nm,
					PstlAdr: &Cdtr_PstlAdr,
				},
				CdtrAcct: &CdtrAcct,
				RmtInf:   &RmtInf,
			},
		},
	}
}
func (msg *CustomerCreditTransferMessage) GetXML() ([]byte, error) {
	return xml.MarshalIndent(msg.doc, "", "\t")
}
func (msg *CustomerCreditTransferMessage) GetJson() ([]byte, error) {
	return json.MarshalIndent(msg.doc.FIToFICstmrCdtTrf, "", "\t")
}
