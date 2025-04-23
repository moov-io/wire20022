package FinancialInstitutionCreditTransfer

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario3_Step3_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario3_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step3_pacs.009")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step3_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
		Amount:   1500000000.00,
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario4_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step3_pacs.009")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario1_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario1_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario1_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario1_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario1_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario2_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario2_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario3_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
		Amount:   1500000000.00,
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario3_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario3_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario3_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario3_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
		Amount:   1500000000.00,
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario3_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario3_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario3_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario4_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario4_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario4_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario4_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario4_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario4_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario4_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario5_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
		Amount:   350000000.00,
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario5_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario5_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario5_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
		Amount:   350000000.00,
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario5_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario5_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario5_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario6_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
		Amount:   17985234.25,
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario6_Step1_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario6_Step1_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario6_Step1_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario6_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
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
		Amount:   17985234.25,
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario6_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario6_Step2_pacs.009")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario6_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
