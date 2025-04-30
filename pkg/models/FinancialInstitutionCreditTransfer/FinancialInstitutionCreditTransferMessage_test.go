package FinancialInstitutionCreditTransfer

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestRequireField(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	cErr := message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreateDateTime, NumberOfTransactions, SettlementMethod, ClearingSystem, PaymentEndToEndId, PaymentUETR, LocalInstrument, InterbankSettlementAmount, InterbankSettlementDate, InstructingAgent, InstructedAgent, Debtor, Creditor")
}
func generateRequreFields(msg Message) Message {
	if msg.data.MessageId == "" {
		msg.data.MessageId = "20250310B1QDRCQR000623"
	}
	if msg.data.CreateDateTime.IsZero() {
		msg.data.CreateDateTime = time.Now()
	}
	if msg.data.NumberOfTransactions <= 0 {
		msg.data.NumberOfTransactions = 1
	}
	if msg.data.SettlementMethod == "" {
		msg.data.SettlementMethod = model.SettlementCLRG
	}
	if msg.data.ClearingSystem == "" {
		msg.data.ClearingSystem = model.ClearingSysFDW
	}
	if msg.data.PaymentEndToEndId == "" {
		msg.data.PaymentEndToEndId = "98z2cb3d0f2f3094f24a16389713541137b"
	}
	if msg.data.PaymentUETR == "" {
		msg.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f999"
	}
	if msg.data.LocalInstrument == "" {
		msg.data.LocalInstrument = BankDrawdownTransfer
	}
	if isEmpty(msg.data.InterbankSettlementAmount) {
		msg.data.InterbankSettlementAmount = model.CurrencyAndAmount{
			Amount:   1000000000.00,
			Currency: "USD",
		}
	}
	if isEmpty(msg.data.InterbankSettlementDate) {
		msg.data.InterbankSettlementDate = model.FromTime(time.Now())
	}
	if isEmpty(msg.data.InstructingAgent) {
		msg.data.InstructingAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	if isEmpty(msg.data.InstructedAgent) {
		msg.data.InstructedAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if isEmpty(msg.data.Debtor) {
		msg.data.Debtor = model.FiniancialInstitutionId{
			ClearingSystemId:       model.PaymentSysUSABA,
			ClearintSystemMemberId: "122240120",
			Name:                   "Bank Bb",
			Address: model.PostalAddress{
				StreetName:     "Avenue C",
				BuildingNumber: "52",
				PostalCode:     "19067",
				TownName:       "Yardley",
				Subdivision:    "PA",
				Country:        "US",
			},
		}
	}
	if isEmpty(msg.data.Creditor) {
		msg.data.Creditor = model.FiniancialInstitutionId{
			ClearingSystemId:       model.PaymentSysUSABA,
			ClearintSystemMemberId: "011104238",
			Name:                   "BANK A",
			Address: model.PostalAddress{
				StreetName:     "Avenue A",
				BuildingNumber: "66",
				PostalCode:     "60532",
				TownName:       "Lisle",
				Subdivision:    "IL",
				Country:        "US",
			},
		}
	}
	return msg
}
func TestFinancialInstitutionCreditTransferFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step3_pacs.009")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250310B1QDRCQR000623", string(message.doc.FICdtTrf.GrpHdr.MsgId))
	require.Equal(t, "1", string(message.doc.FICdtTrf.GrpHdr.NbOfTxs))
	require.Equal(t, "CLRG", string(message.doc.FICdtTrf.GrpHdr.SttlmInf.SttlmMtd))
	require.Equal(t, "FDW", string(*message.doc.FICdtTrf.GrpHdr.SttlmInf.ClrSys.Cd))
	require.Equal(t, "Scenario03Step3InstrId001", string(*message.doc.FICdtTrf.CdtTrfTxInf.PmtId.InstrId))
	require.Equal(t, "Scenario03EndToEndId001", string(message.doc.FICdtTrf.CdtTrfTxInf.PmtId.EndToEndId))
	require.Equal(t, "8a562c67-ca16-48ba-b074-65581be6f999", string(message.doc.FICdtTrf.CdtTrfTxInf.PmtId.UETR))
	require.Equal(t, "BTRD", string(*message.doc.FICdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm.Prtry))
	require.Equal(t, "USABA", string(*message.doc.FICdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "011104238", string(message.doc.FICdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "Bank Bb", string(*message.doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.Nm))
	require.Equal(t, "60532", string(*message.doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.PstlAdr.PstCd))
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"
const INVALID_TRCOUNT string = "123456789012345"
const INVALID_MESSAGE_ID string = "12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456"
const INVALID_OTHER_ID string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const INVALID_BUILD_NUM string = "12345678901234567"
const INVALID_POSTAL_CODE string = "12345678901234567"
const INVALID_COUNTRY_CODE string = "12345678"
const INVALID_MESSAGE_NAME_ID string = "sabcd-123-001-12"
const INVALID_PAY_SYSCODE model.PaymentSystemType = model.PaymentSystemType(INVALID_COUNT)

func TestFinancialInstitutionCreditTransferValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"Invalid MessageId",
			Message{data: MessageModel{MessageId: INVALID_OTHER_ID}},
			"error occur at MessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid SettlementMethod",
			Message{data: MessageModel{SettlementMethod: model.SettlementMethodType(INVALID_COUNT)}},
			"error occur at SettlementMethod: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid ClearingSystem",
			Message{data: MessageModel{ClearingSystem: model.CommonClearingSysCodeType(INVALID_COUNT)}},
			"error occur at ClearingSystem: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid PaymentInstructionId",
			Message{data: MessageModel{PaymentInstructionId: INVALID_OTHER_ID}},
			"error occur at PaymentInstructionId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid PaymentEndToEndId",
			Message{data: MessageModel{PaymentEndToEndId: INVALID_OTHER_ID}},
			"error occur at PaymentEndToEndId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid PaymentUETR",
			Message{data: MessageModel{PaymentUETR: INVALID_OTHER_ID}},
			"error occur at PaymentUETR: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}",
		},
		{
			"Invalid LocalInstrument",
			Message{data: MessageModel{LocalInstrument: InstrumentType(INVALID_COUNT)}},
			"error occur at LocalInstrument: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid InstructingAgent",
			Message{data: MessageModel{InstructingAgent: model.Agent{
				PaymentSysCode:     model.PaymentSystemType(INVALID_COUNT),
				PaymentSysMemberId: "021040078",
			}}},
			"error occur at InstructingAgent.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid InstructedAgent",
			Message{data: MessageModel{InstructedAgent: model.Agent{
				PaymentSysCode:     model.PaymentSysUSABA,
				PaymentSysMemberId: INVALID_ACCOUNT_ID,
			}}},
			"error occur at InstructedAgent.PaymentSysMemberId: 123ABC789 fails validation with pattern [0-9]{9,9}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			nMsg := generateRequreFields(tt.msg)
			msgErr := nMsg.CreateDocument()
			if msgErr != nil {
				require.Equal(t, tt.expectedErr, msgErr.Error())
			}
		})
	}
}
func TestDrawdowns_Scenario3_Step3_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000623"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario03Step3InstrId001"
	message.data.PaymentEndToEndId = "Scenario03EndToEndId001"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f999"
	message.data.LocalInstrument = BankDrawdownTransfer
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1000000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "122240120",
		Name:                   "Bank Bb",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "52",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "011104238",
		Name:                   "BANK A",
		Address: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.RemittanceInfo = "3rd repayment loan with reference ABCD432Z"

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario3_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step3_pacs.009")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step3_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000683"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario04Step3InstrId001"
	message.data.PaymentEndToEndId = "Scenario04EndToEndId001"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f999"
	message.data.LocalInstrument = BankDrawdownTransfer
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "122240120",
		Name:                   "Bank Bb",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "52",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	message.data.CreditorAgent = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "011104238",
		Name:                   "BANK A",
		Address: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "113194159",
		Name:                   "BANK Aa",
		Address: model.PostalAddress{
			StreetName:     "Main Road",
			BuildingNumber: "3",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.RemittanceInfo = "Additional margin call for 03/10/2025 with reference XYZDF22."

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario4_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step3_pacs.009")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario1_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000501"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario01FIInstrId001"
	message.data.PaymentEndToEndId = "BANC338754BAND33"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "231981435",
		Name:                   "BANK D",
		Address: model.PostalAddress{
			StreetName:     "Avenue D",
			BuildingNumber: "35",
			PostalCode:     "60197",
			TownName:       "Carol Stream",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario1_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario1_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000501"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario01FIInstrId001"
	message.data.PaymentEndToEndId = "BANC338754BAND33"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "231981435",
		Name:                   "BANK D",
		Address: model.PostalAddress{
			StreetName:     "Avenue D",
			BuildingNumber: "35",
			PostalCode:     "60197",
			TownName:       "Carol Stream",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario1_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario1_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000502"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario02FIInstrId001"
	message.data.PaymentEndToEndId = "BACc336092BADd33"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BACCUS33",
	}
	message.data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BADDUS33",
	}
	message.data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario2_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000502"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario02FIInstrId001"
	message.data.PaymentEndToEndId = "BACc336092BADd33"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BACCUS33",
	}
	message.data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BADDUS33",
	}
	message.data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario2_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario3_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000503"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario03FIInstrId001"
	message.data.PaymentEndToEndId = "BANZBB7854BANYRJ"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario3_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario3_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario3_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario3_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000503"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario03FIInstrId001"
	message.data.PaymentEndToEndId = "BANZBB7854BANYRJ"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario3_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario3_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario3_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario4_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000504"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario04FIInstrId001"
	message.data.PaymentEndToEndId = "Scenario04FIEtoEId001"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   200000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario4_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario4_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario4_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000504"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario04FIInstrId001"
	message.data.PaymentEndToEndId = "Scenario04FIEtoEId001"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   200000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario4_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario4_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario4_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario5_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000505"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario05FIInstrId001"
	message.data.PaymentEndToEndId = "Scenario05FIEtoEId001"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario5_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario5_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario5_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000505"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario05FIInstrId001"
	message.data.PaymentEndToEndId = "Scenario05FIEtoEId001"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreBankTransfer
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		ClearingSystemId:       model.PaymentSysUSABA,
		ClearintSystemMemberId: "021307481",
		Name:                   "Bank C",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "1099",
			PostalCode:     "92262",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario5_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario5_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario5_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario6_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000506"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario06FIInstrId001"
	message.data.PaymentEndToEndId = "Scenario06FIEtoEId001"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreCoverPayment
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   179852.25,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.data.UnderlyingCustomerCreditTransfer = CreditTransferTransaction{
		Debtor: model.FiniancialInstitutionId{
			Name: "Corporation Z",
			Address: model.PostalAddress{
				StreetName:     "Avenue Moliere",
				BuildingNumber: "70",
				PostalCode:     "1180",
				TownName:       "Brussels",
				Country:        "BE",
			},
		},
		DebtorAccount: "BE34001216371411",
		DebtorAgent: model.FiniancialInstitutionId{
			BusinessId: "BANZBEBB",
		},
		CreditorAgent: model.FiniancialInstitutionId{
			BusinessId: "BANYBRRJ",
		},
		Creditor: model.FiniancialInstitutionId{
			Name: "Corporation Y",
			Address: model.PostalAddress{
				StreetName:     "Av. Lucio Costa",
				BuildingNumber: "5220",
				BuildingName:   "Barra da Tijuca",
				PostalCode:     "22630-012",
				TownName:       "Rio de Janeiro",
				Subdivision:    "RJ",
				Country:        "BR",
			},
		},
		CreditorAccount:       "BR9700360305000010009795493P1",
		RemittanceInformation: "Payment invoice 444563 dated 1st March 2025",
		InstructedAmount: model.CurrencyAndAmount{
			Amount:   17985234.25,
			Currency: "USD",
		},
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario6_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario6_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario6_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario6_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310B1QDRCQR000506"
	message.data.CreateDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.PaymentInstructionId = "Scenario06FIInstrId001"
	message.data.PaymentEndToEndId = "Scenario06FIEtoEId001"
	message.data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.LocalInstrument = CoreCoverPayment
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   179852.25,
		Currency: "USD",
	}
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.data.UnderlyingCustomerCreditTransfer = CreditTransferTransaction{
		Debtor: model.FiniancialInstitutionId{
			Name: "Corporation Z",
			Address: model.PostalAddress{
				StreetName:     "Avenue Moliere",
				BuildingNumber: "70",
				PostalCode:     "1180",
				TownName:       "Brussels",
				Country:        "BE",
			},
		},
		DebtorAccount: "BE34001216371411",
		DebtorAgent: model.FiniancialInstitutionId{
			BusinessId: "BANZBEBB",
		},
		CreditorAgent: model.FiniancialInstitutionId{
			BusinessId: "BANYBRRJ",
		},
		Creditor: model.FiniancialInstitutionId{
			Name: "Corporation Y",
			Address: model.PostalAddress{
				StreetName:     "Av. Lucio Costa",
				BuildingNumber: "5220",
				BuildingName:   "Barra da Tijuca",
				PostalCode:     "22630-012",
				TownName:       "Rio de Janeiro",
				Subdivision:    "RJ",
				Country:        "BR",
			},
		},
		CreditorAccount:       "BR9700360305000010009795493P1",
		RemittanceInformation: "Payment invoice 444563 dated 1st March 2025",
		InstructedAmount: model.CurrencyAndAmount{
			Amount:   17985234.25,
			Currency: "USD",
		},
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario6_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario6_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario6_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
