package FinancialInstitutionCreditTransfer_009_001_08

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario3_Step3_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000623"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario03Step3InstrId001"
	message.model.PaymentEndToEndId = "Scenario03EndToEndId001"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f999"
	message.model.LocalInstrument = BankDrawdownTransfer
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   6000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
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
	message.model.Creditor = model.FiniancialInstitutionId{
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
	message.model.RemittanceInfo = "3rd repayment loan with reference ABCD432Z"

	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario3_Step3_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario4_Step3_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000683"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario04Step3InstrId001"
	message.model.PaymentEndToEndId = "Scenario04EndToEndId001"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f999"
	message.model.LocalInstrument = BankDrawdownTransfer
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1500000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
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
	message.model.CreditorAgent = model.FiniancialInstitutionId{
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
	message.model.Creditor = model.FiniancialInstitutionId{
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
	message.model.RemittanceInfo = "Additional margin call for 03/10/2025 with reference XYZDF22."

	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario4_Step3_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario1_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000501"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario01FIInstrId001"
	message.model.PaymentEndToEndId = "BANC338754BAND33"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = BankDrawdownTransfer
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
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
	message.model.Creditor = model.FiniancialInstitutionId{
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

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario1_Step1_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000501"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario01FIInstrId001"
	message.model.PaymentEndToEndId = "BANC338754BAND33"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
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
	message.model.Creditor = model.FiniancialInstitutionId{
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

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario1_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario2_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000502"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario02FIInstrId001"
	message.model.PaymentEndToEndId = "BANC338754BAND33"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BACCUS33",
	}
	message.model.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.model.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BADDUS33",
	}
	message.model.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BADDUS33",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000502"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario02FIInstrId001"
	message.model.PaymentEndToEndId = "BANC338754BAND33"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   500000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BACCUS33",
	}
	message.model.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.model.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BADDUS33",
	}
	message.model.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario2_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario3_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000503"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario03FIInstrId001"
	message.model.PaymentEndToEndId = "BANZBB7854BANYRJ"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1500000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.model.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.model.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.model.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario3_Step1_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario3_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000503"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario03FIInstrId001"
	message.model.PaymentEndToEndId = "BANZBB7854BANYRJ"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1500000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.model.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.model.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.model.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario3_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario4_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000504"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario04FIInstrId001"
	message.model.PaymentEndToEndId = "Scenario04FIEtoEId001"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   200000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
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
	message.model.Creditor = model.FiniancialInstitutionId{
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

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario4_Step1_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000504"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario04FIInstrId001"
	message.model.PaymentEndToEndId = "Scenario04FIEtoEId001"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   200000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
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
	message.model.Creditor = model.FiniancialInstitutionId{
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

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario4_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario5_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000505"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario05FIInstrId001"
	message.model.PaymentEndToEndId = "Scenario05FIEtoEId001"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   350000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
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
	message.model.Creditor = model.FiniancialInstitutionId{
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

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario5_Step1_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000505"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario05FIInstrId001"
	message.model.PaymentEndToEndId = "Scenario05FIEtoEId001"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   350000000.00,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
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
	message.model.Creditor = model.FiniancialInstitutionId{
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

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario5_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario6_Step1_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000506"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario06FIInstrId001"
	message.model.PaymentEndToEndId = "Scenario06FIEtoEId001"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   17985234.25,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.model.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.model.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.model.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.model.UnderlyingCustomerCreditTransfer = CreditTransferTransaction{
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

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario6_Step1_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario6_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs009Message()
	message.model.MessageId = "20250310B1QDRCQR000506"
	message.model.CreateDateTime = time.Now()
	message.model.NumberOfTransactions = 1
	message.model.SettlementMethod = model.SettlementCLRG
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.ClearingSystem = model.ClearingSysFDW
	message.model.PaymentInstructionId = "Scenario06FIInstrId001"
	message.model.PaymentEndToEndId = "Scenario06FIEtoEId001"
	message.model.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.LocalInstrument = CoreBankTransfer
	message.model.InterbankSettlementDate = civil.DateOf(time.Now())
	message.model.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   17985234.25,
		Currency: "USD",
	}
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.model.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.model.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.model.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.model.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.model.UnderlyingCustomerCreditTransfer = CreditTransferTransaction{
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

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario6_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
