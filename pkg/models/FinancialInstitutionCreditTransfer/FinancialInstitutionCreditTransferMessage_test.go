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
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreateDateTime, NumberOfTransactions, SettlementMethod, ClearingSystem, PaymentEndToEndId, PaymentUETR, LocalInstrument, InterbankSettlementAmount, InterbankSettlementDate, InstructingAgent, InstructedAgent, Debtor, Creditor")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.MessageId == "" {
		msg.Data.MessageId = "20250310B1QDRCQR000623"
	}
	if msg.Data.CreateDateTime.IsZero() {
		msg.Data.CreateDateTime = time.Now()
	}
	if msg.Data.NumberOfTransactions <= 0 {
		msg.Data.NumberOfTransactions = 1
	}
	if msg.Data.SettlementMethod == "" {
		msg.Data.SettlementMethod = model.SettlementCLRG
	}
	if msg.Data.ClearingSystem == "" {
		msg.Data.ClearingSystem = model.ClearingSysFDW
	}
	if msg.Data.PaymentEndToEndId == "" {
		msg.Data.PaymentEndToEndId = "98z2cb3d0f2f3094f24a16389713541137b"
	}
	if msg.Data.PaymentUETR == "" {
		msg.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f999"
	}
	if msg.Data.LocalInstrument == "" {
		msg.Data.LocalInstrument = BankDrawdownTransfer
	}
	if isEmpty(msg.Data.InterbankSettlementAmount) {
		msg.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
			Amount:   1000000000.00,
			Currency: "USD",
		}
	}
	if isEmpty(msg.Data.InterbankSettlementDate) {
		msg.Data.InterbankSettlementDate = model.FromTime(time.Now())
	}
	if isEmpty(msg.Data.InstructingAgent) {
		msg.Data.InstructingAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	if isEmpty(msg.Data.InstructedAgent) {
		msg.Data.InstructedAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if isEmpty(msg.Data.Debtor) {
		msg.Data.Debtor = model.FiniancialInstitutionId{
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
	if isEmpty(msg.Data.Creditor) {
		msg.Data.Creditor = model.FiniancialInstitutionId{
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
	require.Equal(t, "20250310B1QDRCQR000623", string(message.Doc.FICdtTrf.GrpHdr.MsgId))
	require.Equal(t, "1", string(message.Doc.FICdtTrf.GrpHdr.NbOfTxs))
	require.Equal(t, "CLRG", string(message.Doc.FICdtTrf.GrpHdr.SttlmInf.SttlmMtd))
	require.Equal(t, "FDW", string(*message.Doc.FICdtTrf.GrpHdr.SttlmInf.ClrSys.Cd))
	require.Equal(t, "Scenario03Step3InstrId001", string(*message.Doc.FICdtTrf.CdtTrfTxInf.PmtId.InstrId))
	require.Equal(t, "Scenario03EndToEndId001", string(message.Doc.FICdtTrf.CdtTrfTxInf.PmtId.EndToEndId))
	require.Equal(t, "8a562c67-ca16-48ba-b074-65581be6f999", string(message.Doc.FICdtTrf.CdtTrfTxInf.PmtId.UETR))
	require.Equal(t, "BTRD", string(*message.Doc.FICdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm.Prtry))
	require.Equal(t, "USABA", string(*message.Doc.FICdtTrf.CdtTrfTxInf.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "011104238", string(message.Doc.FICdtTrf.CdtTrfTxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "Bank Bb", string(*message.Doc.FICdtTrf.CdtTrfTxInf.Dbtr.FinInstnId.Nm))
	require.Equal(t, "60532", string(*message.Doc.FICdtTrf.CdtTrfTxInf.Cdtr.FinInstnId.PstlAdr.PstCd))
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
			Message{Data: MessageModel{MessageId: INVALID_OTHER_ID}},
			"error occur at MessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid SettlementMethod",
			Message{Data: MessageModel{SettlementMethod: model.SettlementMethodType(INVALID_COUNT)}},
			"error occur at SettlementMethod: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid ClearingSystem",
			Message{Data: MessageModel{ClearingSystem: model.CommonClearingSysCodeType(INVALID_COUNT)}},
			"error occur at ClearingSystem: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid PaymentInstructionId",
			Message{Data: MessageModel{PaymentInstructionId: INVALID_OTHER_ID}},
			"error occur at PaymentInstructionId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid PaymentEndToEndId",
			Message{Data: MessageModel{PaymentEndToEndId: INVALID_OTHER_ID}},
			"error occur at PaymentEndToEndId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid PaymentUETR",
			Message{Data: MessageModel{PaymentUETR: INVALID_OTHER_ID}},
			"error occur at PaymentUETR: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}",
		},
		{
			"Invalid LocalInstrument",
			Message{Data: MessageModel{LocalInstrument: InstrumentType(INVALID_COUNT)}},
			"error occur at LocalInstrument: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid InstructingAgent",
			Message{Data: MessageModel{InstructingAgent: model.Agent{
				PaymentSysCode:     model.PaymentSystemType(INVALID_COUNT),
				PaymentSysMemberId: "021040078",
			}}},
			"error occur at InstructingAgent.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid InstructedAgent",
			Message{Data: MessageModel{InstructedAgent: model.Agent{
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
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000623"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario03Step3InstrId001"
	message.Data.PaymentEndToEndId = "Scenario03EndToEndId001"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f999"
	message.Data.LocalInstrument = BankDrawdownTransfer
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1000000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
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
	message.Data.Creditor = model.FiniancialInstitutionId{
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
	message.Data.RemittanceInfo = "3rd repayment loan with reference ABCD432Z"

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario3_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step3_pacs.009")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step3_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000683"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario04Step3InstrId001"
	message.Data.PaymentEndToEndId = "Scenario04EndToEndId001"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f999"
	message.Data.LocalInstrument = BankDrawdownTransfer
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
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
	message.Data.CreditorAgent = model.FiniancialInstitutionId{
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
	message.Data.Creditor = model.FiniancialInstitutionId{
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
	message.Data.RemittanceInfo = "Additional margin call for 03/10/2025 with reference XYZDF22."

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario4_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step3_pacs.009")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario1_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000501"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario01FIInstrId001"
	message.Data.PaymentEndToEndId = "BANC338754BAND33"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
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
	message.Data.Creditor = model.FiniancialInstitutionId{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario1_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario1_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000501"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario01FIInstrId001"
	message.Data.PaymentEndToEndId = "BANC338754BAND33"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
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
	message.Data.Creditor = model.FiniancialInstitutionId{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario1_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario1_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000502"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario02FIInstrId001"
	message.Data.PaymentEndToEndId = "BACc336092BADd33"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BACCUS33",
	}
	message.Data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.Data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BADDUS33",
	}
	message.Data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario2_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000502"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario02FIInstrId001"
	message.Data.PaymentEndToEndId = "BACc336092BADd33"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BACCUS33",
	}
	message.Data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.Data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BADDUS33",
	}
	message.Data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario2_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario3_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000503"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario03FIInstrId001"
	message.Data.PaymentEndToEndId = "BANZBB7854BANYRJ"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.Data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.Data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.Data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario3_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario3_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario3_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario3_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000503"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario03FIInstrId001"
	message.Data.PaymentEndToEndId = "BANZBB7854BANYRJ"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.Data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.Data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.Data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario3_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario3_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario3_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario4_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000504"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario04FIInstrId001"
	message.Data.PaymentEndToEndId = "Scenario04FIEtoEId001"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   200000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
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
	message.Data.Creditor = model.FiniancialInstitutionId{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario4_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario4_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario4_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000504"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario04FIInstrId001"
	message.Data.PaymentEndToEndId = "Scenario04FIEtoEId001"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   200000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
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
	message.Data.Creditor = model.FiniancialInstitutionId{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario4_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario4_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario4_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario5_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000505"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario05FIInstrId001"
	message.Data.PaymentEndToEndId = "Scenario05FIEtoEId001"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
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
	message.Data.Creditor = model.FiniancialInstitutionId{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario5_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario5_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario5_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000505"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario05FIInstrId001"
	message.Data.PaymentEndToEndId = "Scenario05FIEtoEId001"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreBankTransfer
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
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
	message.Data.Creditor = model.FiniancialInstitutionId{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario5_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario5_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario5_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario6_Step1_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000506"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario06FIInstrId001"
	message.Data.PaymentEndToEndId = "Scenario06FIEtoEId001"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreCoverPayment
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   179852.25,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.Data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.Data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.Data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.Data.UnderlyingCustomerCreditTransfer = CreditTransferTransaction{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario6_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario6_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario6_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario6_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000506"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario06FIInstrId001"
	message.Data.PaymentEndToEndId = "Scenario06FIEtoEId001"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = CoreCoverPayment
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   179852.25,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.Data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.Data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.Data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.Data.UnderlyingCustomerCreditTransfer = CreditTransferTransaction{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario6_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario6_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario6_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
