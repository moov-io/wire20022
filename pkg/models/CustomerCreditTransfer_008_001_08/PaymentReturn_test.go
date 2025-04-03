package CustomerCreditTransfer_008_001_08

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestPaymentReturn_Scenario1_Step1_CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000400"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = model.InstrumentCTRC
	mesage.model.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = model.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "567876543"
	mesage.model.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: CodeCINV,
		Number:            "INV34563",
		RelatedDate:       civil.DateOf(time.Now()),
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario2_Step1_CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000400"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario02InstrId001"
	mesage.model.EndToEndId = "Scenario02EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = model.InstrumentCTRC
	mesage.model.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1234578.88,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1234578.88,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation C"
	mesage.model.DebtorAddress = model.PostalAddress{
		StreetName:     "40th Street",
		BuildingNumber: "1180",
		PostalCode:     "11218",
		TownName:       "Brooklyn",
		Subdivision:    "NY",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "0031234567"
	mesage.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "567876543"
	mesage.model.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: CodeCINV,
		Number:            "INV34563",
		RelatedDate:       civil.DateOf(time.Now()),
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "PaymentReturn_Scenario2_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario3_Step1_CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000400"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario03InstrId001"
	mesage.model.EndToEndId = "Scenario03EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = model.InstrumentCTRC
	mesage.model.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1234578.88,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1234578.88,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation C"
	mesage.model.DebtorAddress = model.PostalAddress{
		StreetName:     "40th Street",
		BuildingNumber: "1180",
		PostalCode:     "11218",
		TownName:       "Brooklyn",
		Subdivision:    "NY",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "0031234567"
	mesage.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "567876543"
	mesage.model.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: CodeCINV,
		Number:            "INV34563",
		RelatedDate:       civil.DateOf(time.Now()),
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "PaymentReturn_Scenario3_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario4_Step1_CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000400"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario04InstrId001"
	mesage.model.EndToEndId = "Scenario04EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = model.InstrumentCTRC
	mesage.model.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1234578.88,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1234578.88,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation C"
	mesage.model.DebtorAddress = model.PostalAddress{
		StreetName:     "40th Street",
		BuildingNumber: "1180",
		PostalCode:     "11218",
		TownName:       "Brooklyn",
		Subdivision:    "NY",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "0031234567"
	mesage.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "567876543"
	mesage.model.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: CodeCINV,
		Number:            "INV34563",
		RelatedDate:       civil.DateOf(time.Now()),
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "PaymentReturn_Scenario4_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario5_Step1_CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000450"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario05InstrId001"
	mesage.model.EndToEndId = "Scenario05EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = model.InstrumentCTRC
	mesage.model.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 3234578.89,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 3234578.89,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	mesage.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	mesage.model.DebtorName = "Corporation Z"
	mesage.model.DebtorAddress = model.PostalAddress{
		StreetName: "Avenue Moliere 70",
		PostalCode: "1180",
		TownName:   "Brussels",
		Country:    "BE",
	}
	mesage.model.DebtorIBAN = "BE34001216371411"
	mesage.model.DebtorAgent = model.Agent{
		BusinessIdCode: "BANZBEBB",
	}
	mesage.model.CreditorAgent = model.Agent{
		BusinessIdCode: "BANYBRRJ",
	}
	mesage.model.CreditorName = "Corporation Y"
	mesage.model.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Av. Lucio Costa",
		BuildingNumber: "15220",
		BuildingName:   "Barra da Tijuca",
		PostalCode:     "22630-012",
		TownName:       "Rio de Janeiro",
		Country:        "BR",
	}
	mesage.model.CreditorIBAN = "BR9700360305000010009795493P1"
	mesage.model.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: CodeCMCN,
		Number:            "ABC-987",
		RelatedDate:       civil.DateOf(time.Now()),
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "PaymentReturn_Scenario5_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
