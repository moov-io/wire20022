package CustomerCreditTransfer_pacs_008_001_08

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/require"
)

func TestCustomerCreditTransfer_Scenario1_Step1_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000001"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			RoomNumber:     "60532",
			PostalCode:     "85268",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = PostalAddress{
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
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario1_Step2_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000001"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = PostalAddress{
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
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario2_Step1_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000002"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario02InstrId001"
	mesage.model.EndToEndId = "Scenario02EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040079",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = PostalAddress{
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
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario3_Step1_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000001"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = PostalAddress{
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
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario4_Step1_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000004"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario04InstrId001"
	mesage.model.EndToEndId = "Scenario04EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 999008.53,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 999008.53,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "091036164",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "091036164",
		BankName:           "Internal Revenue Service",
		PostalAddress: PostalAddress{
			StreetName:     "West Perching Road",
			BuildingNumber: "333",
			PostalCode:     "64108",
			TownName:       "Kansas City",
			Subdivision:    "MO",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation A"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "567876543"
	mesage.model.RemittanceInfor = RemittanceDocument{
		TaxDetail: TaxRecord{
			TaxId:              "123456789",
			TaxTypeCode:        "09455",
			TaxPeriodYear:      civil.DateOf(time.Now()),
			TaxperiodTimeFrame: "MM04",
		},
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario5_Step1_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000005"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario05InstrId001"
	mesage.model.EndToEndId = "Scenario05EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.SericeLevel = "G001"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSHAR
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	mesage.model.DebtorName = "Corporation Z"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName: "Avenue Moliere 70",
		PostalCode: "85268",
		TownName:   "Brussels",
		Country:    "BE",
	}
	mesage.model.DebtorIBAN = "BE34001216371411"
	mesage.model.DebtorAgent = Agent{
		BusinessIdCode: "BANZBEBB",
	}
	mesage.model.CreditorAgent = Agent{
		BusinessIdCode: "BANYBRRJ",
	}
	mesage.model.CreditorName = "Corporation Y"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Av. Lucio Costa",
		BuildingNumber: "5220",
		BuildingName:   "Barra da Tijuca",
		PostalCode:     "22630-012",
		TownName:       "Rio de Janeiro",
		Country:        "US",
	}
	mesage.model.CreditorIBAN = "BR9700360305000010009795493P1"
	mesage.model.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: CodeOrProprietaryType(DocumentType6CodeCmcn),
		Number:            "ABC-987",
		RelatedDate:       civil.DateOf(time.Now()),
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario5_Step2_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000005"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario05InstrId001"
	mesage.model.EndToEndId = "Scenario05EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.SericeLevel = "G001"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSHAR
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	mesage.model.DebtorName = "Corporation Z"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName: "Avenue Moliere 70",
		PostalCode: "85268",
		TownName:   "Brussels",
		Country:    "BE",
	}
	mesage.model.DebtorIBAN = "BE34001216371411"
	mesage.model.DebtorAgent = Agent{
		BusinessIdCode: "BANZBEBB",
	}
	mesage.model.CreditorAgent = Agent{
		BusinessIdCode: "BANYBRRJ",
	}
	mesage.model.CreditorName = "Corporation Y"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Av. Lucio Costa",
		BuildingNumber: "5220",
		BuildingName:   "Barra da Tijuca",
		PostalCode:     "22630-012",
		TownName:       "Rio de Janeiro",
		Country:        "US",
	}
	mesage.model.CreditorIBAN = "BR9700360305000010009795493P1"
	mesage.model.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: CodeOrProprietaryType(DocumentType6CodeCmcn),
		Number:            "ABC-987",
		RelatedDate:       civil.DateOf(time.Now()),
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}

func TestCustomerCreditTransfer_Variantion1_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000006"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Variation1InstrId001"
	mesage.model.EndToEndId = "Variation1EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "567876543"
	mesage.model.PurposeOfPayment = InvestmentPayment
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
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Variantion1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Variantion2_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000007"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Variation2InstrId001"
	mesage.model.EndToEndId = "Variation2EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "567876543"
	mesage.model.PurposeOfPayment = InvestmentPayment
	mesage.model.RelatedRemittanceInfo = RemittanceDetail{
		RemittanceId:      "Scenario01Var2RemittanceId001",
		Method:            Email,
		ElectronicAddress: "CustomerService@CorporationB.com",
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Variantion2.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Variantion3_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000008"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Variation3InstrId001"
	mesage.model.EndToEndId = "Variation3EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.UltimateDebtorName = "Corporation Aa"
	mesage.model.UltimateDebtorAddress = PostalAddress{
		StreetName:     "Ocean Street",
		BuildingNumber: "1",
		PostalCode:     "97035",
		TownName:       "Portland",
		Subdivision:    "OR",
		Country:        "US",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "567876543"
	mesage.model.PurposeOfPayment = InvestmentPayment
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
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Variantion3.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Variantion4_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000009"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Variation4InstrId001"
	mesage.model.EndToEndId = "Variation4EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "567876543"
	mesage.model.UltimateCreditorName = "Corporation Bb"
	mesage.model.UltimateCreditorAddress = PostalAddress{
		StreetName:     "9th Avenue",
		BuildingNumber: "66",
		BuildingName:   "The Porter House",
		RoomNumber:     "Unit 6",
		PostalCode:     "10011",
		TownName:       "New York",
		Subdivision:    "NY",
		Country:        "US",
	}
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
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Variantion4.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Variantion5_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000001"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Variation5InstrId001"
	mesage.model.EndToEndId = "Variation5EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.SericeLevel = "G001"
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 1009858.99,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "EUR", Amount: 1000000.00,
	}
	mesage.model.exchangeRate = 0.9901
	mesage.model.ChargeBearer = ChargeBearerType(ChargeBearerType1CodeCred)
	mesage.model.ChargesInfo = []ChargeInfo{
		{
			amount:         CurrencyAndAmount{Currency: "USD", Amount: 90.00},
			BusinessIdCode: "BANZBEBB",
		},
		{
			amount:         CurrencyAndAmount{Currency: "USD", Amount: 40.00},
			BusinessIdCode: "BANCUS33",
		},
	}

	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	mesage.model.DebtorName = "Corporation Z"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName: "Avenue Moliere 70",
		PostalCode: "1180",
		TownName:   "Brussels",
		Country:    "BE",
	}
	mesage.model.DebtorIBAN = "BE34001216371411"
	mesage.model.DebtorAgent = Agent{
		BusinessIdCode: "BANZBEBB",
	}
	mesage.model.CreditorAgent = Agent{
		BusinessIdCode: "BANYBRRJ",
	}
	mesage.model.CreditorName = "Corporation Y"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Av. Lucio Costa",
		BuildingNumber: "5220",
		BuildingName:   "Barra da Tijuca",
		PostalCode:     "22630-012",
		TownName:       "Rio de Janeiro",
		Country:        "BR",
	}
	mesage.model.CreditorIBAN = "BR9700360305000010009795493P1"
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
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Variantion5.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Variantion6_CreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000001"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Variation6InstrId001"
	mesage.model.EndToEndId = "Variation6EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = InstrumentCTRC
	mesage.model.SericeLevel = "G001"
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 500000.00,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 500000.00,
	}
	mesage.model.ChargeBearer = ChargeBearerType(ChargeBearerType1CodeShar)
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	mesage.model.IntermediaryAgent1Id = "BANYBRRJ"
	mesage.model.DebtorName = "Corporation Z"
	mesage.model.DebtorAddress = PostalAddress{
		StreetName: "Avenue Moliere 70",
		PostalCode: "1180",
		TownName:   "Brussels",
		Country:    "BE",
	}
	mesage.model.DebtorIBAN = "BE34001216371411"
	mesage.model.DebtorAgent = Agent{
		BusinessIdCode: "BANZBEBB",
	}
	mesage.model.CreditorAgent = Agent{
		BusinessIdCode: "BANYBRRJ",
	}
	mesage.model.CreditorName = "Individual X"
	mesage.model.CreditorPostalAddress = PostalAddress{
		StreetName:     "Rua Aprazivel",
		BuildingNumber: "52",
		PostalCode:     "22630-012",
		TownName:       "Rio de Janeiro",
		Country:        "BR",
	}
	mesage.model.CreditorIBAN = "BR1800360305000010009795493C1"
	mesage.model.RemittanceInfor = RemittanceDocument{
		UnstructuredRemitInfo: "Insurance Benefit/Policy XXAB9876/$500000.00",
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Variantion6.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
