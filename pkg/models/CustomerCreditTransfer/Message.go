package CustomerCreditTransfer

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	pacs008 "github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer_pacs_008_001_08"
	fedwire "github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"

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
	InterBankSettDate model.Date
	//stands for Instructed Amount, which represents the amount that the sender has instructed to be transferred in a payment transaction.
	InstructedAmount model.CurrencyAndAmount

	exchangeRate float64
	//Charge Bearer. It specifies who is responsible for paying the charges (fees) associated with the transaction.
	//default value: SLEV
	ChargeBearer model.ChargeBearerType

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
	Data   MessageModel
	Doc    pacs008.Document
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
	if isEmpty(msg.Data.CreatedDateTime) {
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if msg.Data.NumberOfTransactions == 0 {
		ParamNames = append(ParamNames, "NumberOfTransactions")
	}
	if msg.Data.SettlementMethod == "" {
		ParamNames = append(ParamNames, "SettlementMethod")
	}
	if msg.Data.CommonClearingSysCode == "" {
		ParamNames = append(ParamNames, "CommonClearingSysCode")
	}
	if msg.Data.InstructionId == "" {
		ParamNames = append(ParamNames, "InstructionId")
	}
	if msg.Data.EndToEndId == "" {
		ParamNames = append(ParamNames, "EndToEndId")
	}
	if msg.Data.UniqueEndToEndTransactionRef == "" {
		ParamNames = append(ParamNames, "UniqueEndToEndTransactionRef")
	}
	if msg.Data.InstrumentPropCode == "" {
		ParamNames = append(ParamNames, "InstrumentPropCode")
	}
	if isEmpty(msg.Data.InterBankSettAmount) {
		ParamNames = append(ParamNames, "InterBankSettAmount")
	} else if msg.Data.InterBankSettAmount.Amount == 0 {
		ParamNames = append(ParamNames, "InterBankSettAmount.Amount")
	} else if msg.Data.InterBankSettAmount.Currency == "" {
		ParamNames = append(ParamNames, "InterBankSettAmount.Currency")
	}
	if isEmpty(msg.Data.InterBankSettDate) {
		ParamNames = append(ParamNames, "InterBankSettDate")
	}
	if isEmpty(msg.Data.InstructedAmount) {
		ParamNames = append(ParamNames, "InstructedAmount")
	} else if msg.Data.InstructedAmount.Amount == 0 {
		ParamNames = append(ParamNames, "InstructedAmount.Amount")
	} else if msg.Data.InstructedAmount.Currency == "" {
		ParamNames = append(ParamNames, "InstructedAmount.Currency")
	}
	if msg.Data.ChargeBearer == "" {
		ParamNames = append(ParamNames, "ChargeBearer")
	}
	if isEmpty(msg.Data.InstructingAgents) {
		ParamNames = append(ParamNames, "InstructingAgents")
	}
	if isEmpty(msg.Data.InstructedAgent) {
		ParamNames = append(ParamNames, "InstructedAgent")
	}
	if msg.Data.DebtorName == "" {
		ParamNames = append(ParamNames, "DebtorName")
	}
	if isEmpty(msg.Data.DebtorAddress) {
		ParamNames = append(ParamNames, "DebtorAddress")
	}
	if isEmpty(msg.Data.DebtorAgent) {
		ParamNames = append(ParamNames, "DebtorAgent")
	}
	if isEmpty(msg.Data.CreditorAgent) {
		ParamNames = append(ParamNames, "CreditorAgent")
	}
	if isEmpty(msg.Data.CreditorAgent) {
		ParamNames = append(ParamNames, "DebtorAgent")
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
	if msg.Data.SettlementMethod != "" {
		err := pacs008.SettlementMethod1Code1(msg.Data.SettlementMethod).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "SettlementMethod",
				Message:   err.Error(),
			}
		}
	}
	if msg.Data.CommonClearingSysCode != "" {
		err := pacs008.ExternalCashClearingSystem1CodeFixed(msg.Data.CommonClearingSysCode).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CommonClearingSysCode",
				Message:   err.Error(),
			}
		}
		SttlmInf_ClrSys_Cd = pacs008.ExternalCashClearingSystem1CodeFixed(msg.Data.CommonClearingSysCode)
	}

	if msg.Data.InstructionId != "" {
		err := pacs008.Max35Text(msg.Data.InstructionId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "InstructionId",
				Message:   err.Error(),
			}
		}
		CdtTrfTxInf_PmtId_InstrId = pacs008.Max35Text(msg.Data.InstructionId)
	}

	for _, charge := range msg.Data.ChargesInfo {
		converted, err := Charges71From(charge)
		if err != nil {
			err.InsertPath("ChargesInfo")
			return err
		}
		if !isEmpty(converted) {
			charges71List = append(charges71List, &converted)
		}
	}

	if msg.Data.InstructingAgents.PaymentSysCode != "" {
		err := pacs008.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructingAgents.PaymentSysCode).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
			vErr.InsertPath("InstructingAgents")
			return &vErr
		}
		InstgAgt_FinInstnId_ClrSysId = pacs008.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructingAgents.PaymentSysCode)
	}

	if msg.Data.InstructedAgent.PaymentSysCode != "" {
		err := pacs008.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructedAgent.PaymentSysCode).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
			vErr.InsertPath("InstructedAgent")
			return &vErr
		}
		InstdAgt_FinInstnId_ClrSysId = pacs008.ExternalClearingSystemIdentification1CodeFixed(msg.Data.InstructedAgent.PaymentSysCode)
	}
	var validateError *model.ValidateError
	if msg.Data.DebtorIBAN != "" || msg.Data.DebtorOtherTypeId != "" {
		DbtrAcct, validateError = CashAccount38From("DebtorIBAN", msg.Data.DebtorIBAN, "DebtorOtherTypeId", msg.Data.DebtorOtherTypeId)
		if validateError != nil {
			return validateError
		}
	}
	if msg.Data.CreditorName != "" {
		err := pacs008.Max140Text(msg.Data.CreditorName).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "CreditorName",
				Message:   err.Error(),
			}
			return &vErr
		}
		Cdtr_Nm = pacs008.Max140Text(msg.Data.CreditorName)
	}
	_Cdtr_PstlAdr, vErr := PostalAddress241From(msg.Data.CreditorPostalAddress)
	if vErr != nil {
		vErr.InsertPath("CreditorPostalAddress")
		return vErr
	}
	if !isEmpty(_Cdtr_PstlAdr) {
		Cdtr_PstlAdr = _Cdtr_PstlAdr
	}

	if msg.Data.CreditorIBAN != "" || msg.Data.CreditorOtherTypeId != "" {
		CdtrAcct, validateError = CashAccount38From("CreditorIBAN", msg.Data.CreditorIBAN, "CreditorOtherTypeId", msg.Data.CreditorOtherTypeId)
		if validateError != nil {
			return validateError
		}
	}

	_RltdRmtInf, vErr := RemittanceLocation71From(msg.Data.RelatedRemittanceInfo)
	if vErr != nil {
		vErr.InsertPath("RelatedRemittanceInfo")
		return vErr
	}
	if !isEmpty(_RltdRmtInf) {
		RltdRmtInf = _RltdRmtInf
	}

	_RmtInf, vErr := RemittanceInformation161From(msg.Data.RemittanceInfor)
	if vErr != nil {
		vErr.InsertPath("RemittanceInfor")
		return vErr
	}
	if !isEmpty(_RmtInf) {
		RmtInf = _RmtInf
	}
	CdtTrfTxInf_UltimateDbtr, vErr := PartyIdentification1351From(msg.Data.UltimateDebtorName, msg.Data.UltimateDebtorAddress)
	if vErr != nil {
		vErr.InsertPath("UltimateDebtor")
		return vErr
	}
	CdtTrfTxInf_Dbtr, vErr := PartyIdentification1352From(msg.Data.DebtorName, msg.Data.DebtorAddress)
	if vErr != nil {
		vErr.InsertPath("DebtorAddress")
		return vErr
	}
	DbtrAgt_FinInstnId, vErr := FinancialInstitutionIdentification181From(msg.Data.DebtorAgent)
	if vErr != nil {
		vErr.InsertPath("DebtorAgent")
		return vErr
	}
	CdtTrfTxInf_UltimateCdtr, vErr := PartyIdentification1351From(msg.Data.UltimateCreditorName, msg.Data.UltimateCreditorAddress)
	if vErr != nil {
		vErr.InsertPath("UltimateCreditor")
		return vErr
	}
	CdtrAgt_FinInstnId, vErr := FinancialInstitutionIdentification181From(msg.Data.CreditorAgent)
	if vErr != nil {
		vErr.InsertPath("CreditorAgent")
		return vErr
	}
	CdtTrfTxInf_PmtTpInf, vErr := PaymentTypeInformation281From(msg.Data.InstrumentPropCode, msg.Data.SericeLevel)
	if vErr != nil {
		vErr.InsertPath("Instrument")
		return vErr
	}
	// Construct the Document structure
	msg.Doc = pacs008.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
		FIToFICstmrCdtTrf: pacs008.FIToFICustomerCreditTransferV08{
			GrpHdr: pacs008.GroupHeader931{
				MsgId:   pacs008.IMADFedwireFunds1(msg.Data.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.Data.CreatedDateTime),
				NbOfTxs: pacs008.Max15NumericTextFixed(strconv.Itoa(msg.Data.NumberOfTransactions)),
				SttlmInf: pacs008.SettlementInstruction71{
					SttlmMtd: pacs008.SettlementMethod1Code1(msg.Data.SettlementMethod),
					ClrSys: pacs008.ClearingSystemIdentification3Choice1{
						Cd: &SttlmInf_ClrSys_Cd,
					},
				},
			},
			CdtTrfTxInf: pacs008.CreditTransferTransaction391{
				PmtId: pacs008.PaymentIdentification71{
					InstrId:    &CdtTrfTxInf_PmtId_InstrId,
					EndToEndId: pacs008.Max35Text(msg.Data.EndToEndId),
					UETR:       pacs008.UUIDv4Identifier(msg.Data.UniqueEndToEndTransactionRef),
				},
				PmtTpInf: CdtTrfTxInf_PmtTpInf,
				IntrBkSttlmAmt: pacs008.ActiveCurrencyAndAmountFedwire1{
					Value: pacs008.ActiveCurrencyAndAmountFedwire1SimpleType(msg.Data.InterBankSettAmount.Amount),
					Ccy:   pacs008.ActiveCurrencyCodeFixed(msg.Data.InterBankSettAmount.Currency),
				},
				IntrBkSttlmDt: msg.Data.InterBankSettDate.Date(),
				InstdAmt: pacs008.ActiveOrHistoricCurrencyAndAmount{
					Value: pacs008.ActiveOrHistoricCurrencyAndAmountSimpleType(msg.Data.InstructedAmount.Amount),
					Ccy:   pacs008.ActiveOrHistoricCurrencyCode(msg.Data.InstructedAmount.Currency),
				},
				ChrgBr: pacs008.ChargeBearerType1Code(msg.Data.ChargeBearer),
				InstgAgt: pacs008.BranchAndFinancialInstitutionIdentification62{
					FinInstnId: pacs008.FinancialInstitutionIdentification182{
						ClrSysMmbId: pacs008.ClearingSystemMemberIdentification22{
							ClrSysId: pacs008.ClearingSystemIdentification2Choice2{
								Cd: &InstgAgt_FinInstnId_ClrSysId,
							},
							MmbId: pacs008.RoutingNumberFRS1(msg.Data.InstructingAgents.PaymentSysMemberId),
						},
					},
				},
				InstdAgt: pacs008.BranchAndFinancialInstitutionIdentification62{
					FinInstnId: pacs008.FinancialInstitutionIdentification182{
						ClrSysMmbId: pacs008.ClearingSystemMemberIdentification22{
							ClrSysId: pacs008.ClearingSystemIdentification2Choice2{
								Cd: &InstdAgt_FinInstnId_ClrSysId,
							},
							MmbId: pacs008.RoutingNumberFRS1(msg.Data.InstructedAgent.PaymentSysMemberId),
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
			},
		},
	}
	if !isEmpty(CdtrAcct) {
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct = &CdtrAcct
	}
	if len(charges71List) > 0 {
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.ChrgsInf = charges71List
	}

	if msg.Data.exchangeRate != 0 {
		_exchangeRate := pacs008.BaseOneRate(msg.Data.exchangeRate)
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.XchgRate = &_exchangeRate
	}
	if msg.Data.IntermediaryAgent1Id != "" {
		_IntrmyAgt1, vErr := BranchAndFinancialInstitutionIdentification61From(msg.Data.IntermediaryAgent1Id)
		if vErr != nil {
			vErr.InsertPath("IntermediaryAgent1Id")
			return vErr
		}
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrmyAgt1 = &_IntrmyAgt1
	}

	if !isEmpty(CdtTrfTxInf_UltimateDbtr) {
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtDbtr = &CdtTrfTxInf_UltimateDbtr
	}
	if !isEmpty(CdtTrfTxInf_Dbtr) {
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr = CdtTrfTxInf_Dbtr
	}
	if !isEmpty(CdtTrfTxInf_UltimateCdtr) {
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr = &CdtTrfTxInf_UltimateCdtr
	}
	if msg.Data.PurposeOfPayment != "" {
		_Cd := pacs008.ExternalPurpose1Code(msg.Data.PurposeOfPayment)
		CdtTrfTxInf_Purp = pacs008.Purpose2Choice{
			Cd: &_Cd,
		}
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Purp = &CdtTrfTxInf_Purp
	}
	if !isEmpty(RltdRmtInf) {
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RltdRmtInf = &RltdRmtInf
	}
	if !isEmpty(RmtInf) {
		msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf = &RmtInf
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc) && !isEmpty(msg.Doc.FIToFICstmrCdtTrf) {
		if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.GrpHdr) {
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.MsgId) {
				msg.Data.MessageId = string(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.MsgId)
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.CreDtTm) {
				msg.Data.CreatedDateTime = time.Time(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.CreDtTm)
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.NbOfTxs) {
				msg.Data.NumberOfTransactions, _ = strconv.Atoi(string(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.NbOfTxs))
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.SttlmInf) {
				msg.Data.SettlementMethod = model.SettlementMethodType(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.SttlmInf.SttlmMtd)
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys) && !isEmpty(msg.Doc.FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd) {
				msg.Data.CommonClearingSysCode = model.CommonClearingSysCodeType(*msg.Doc.FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd)
			}
		}
		if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf) {
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId) {
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.InstrId) {
					msg.Data.InstructionId = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.InstrId)
				}
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.EndToEndId) {
					msg.Data.EndToEndId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.EndToEndId)
				}
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.UETR) {
					msg.Data.UniqueEndToEndTransactionRef = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.UETR)
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf) {
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.SvcLvl) && len(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.SvcLvl) > 0 {
					if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.SvcLvl[0].Cd) {
						msg.Data.SericeLevel = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.SvcLvl[0].Cd)
					}
				}
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm) {
					if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm.Prtry) {
						msg.Data.InstrumentPropCode = model.InstrumentPropCodeType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm.Prtry)
					}
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt) {
				msg.Data.InterBankSettAmount = model.CurrencyAndAmount{
					Amount:   float64(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Value),
					Currency: string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmAmt.Ccy),
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmDt) {
				msg.Data.InterBankSettDate = model.FromDate(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrBkSttlmDt)
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.InstdAmt) {
				msg.Data.InstructedAmount = model.CurrencyAndAmount{
					Amount:   float64(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.InstdAmt.Value),
					Currency: string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.InstdAmt.Ccy),
				}
			}
			if msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.XchgRate != nil {
				msg.Data.exchangeRate = float64(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.XchgRate)
			}
			if msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.ChrgBr != "" {
				msg.Data.ChargeBearer = model.ChargeBearerType(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.ChrgBr)
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.ChrgsInf) {
				for _, charge := range msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.ChrgsInf {
					converted := Charges71To(*charge)
					msg.Data.ChargesInfo = append(msg.Data.ChargesInfo, converted)
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.InstgAgt) {
				msg.Data.InstructingAgents = model.Agent{
					PaymentSysCode:     model.PaymentSystemType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd),
					PaymentSysMemberId: string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId),
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.InstdAgt) {
				msg.Data.InstructedAgent = model.Agent{
					PaymentSysCode:     model.PaymentSystemType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd),
					PaymentSysMemberId: string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId),
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrmyAgt1) && !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrmyAgt1.FinInstnId) && !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrmyAgt1.FinInstnId.ClrSysMmbId) {
				msg.Data.IntermediaryAgent1Id = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId)
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr) {
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.Nm) {
					msg.Data.DebtorName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.Nm)
				}
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr) {
					msg.Data.DebtorAddress = PostalAddress241To(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Dbtr.PstlAdr)
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct) {
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.Id.IBAN) {
					msg.Data.DebtorIBAN = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.Id.IBAN)
				}
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.Id.Othr) && !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.Id.Othr.Id) {
					msg.Data.DebtorOtherTypeId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAcct.Id.Othr.Id)
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAgt) {
				msg.Data.DebtorAgent = model.Agent{
					PaymentSysCode:     model.PaymentSystemType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd),
					PaymentSysMemberId: string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.DbtrAgt.FinInstnId.ClrSysMmbId.MmbId),
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr) {
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.Nm) {
					msg.Data.CreditorName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.Nm)
				}
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr) {
					msg.Data.CreditorPostalAddress = PostalAddress241To(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Cdtr.PstlAdr)
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr) {
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.Nm) {
					msg.Data.UltimateCreditorName = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.Nm)
				}
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr) {
					msg.Data.UltimateCreditorAddress = PostalAddress242To(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.UltmtCdtr.PstlAdr)
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct) {
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Id) && !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Id.IBAN) {
					msg.Data.CreditorIBAN = string(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Id.IBAN)
				}
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Id) && !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Id.Othr) && !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Id.Othr.Id) {
					msg.Data.CreditorOtherTypeId = string(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.CdtrAcct.Id.Othr.Id)
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Purp) {
				if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Purp.Cd) {
					msg.Data.PurposeOfPayment = PurposeOfPaymentType(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.Purp.Cd)
				}
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RltdRmtInf) {
				msg.Data.RelatedRemittanceInfo = RemittanceLocation71To(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RltdRmtInf)
			}
			if !isEmpty(msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf) {
				msg.Data.RemittanceInfor = RemittanceInformation161To(*msg.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.RmtInf)
			}
		}
	}
	return nil
}
