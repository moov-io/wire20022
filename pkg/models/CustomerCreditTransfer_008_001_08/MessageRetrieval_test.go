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

func TestMessageRetrieval_Scenario1_Step2_CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000701"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = model.InstrumentCTRC
	mesage.model.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	// jsonFileName := filepath.Join("generated", "MessageRetrieval_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "MessageRetrieval_Scenario1_Step2.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestMessageRetrieval_Scenario2_Step2_1_CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000004"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = model.InstrumentCTRC
	mesage.model.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	// jsonFileName := filepath.Join("generated", "MessageRetrieval_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "MessageRetrieval_Scenario1_Step2_1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestMessageRetrieval_Scenario2_Step2_2_CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000005"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.model.InstrumentPropCode = model.InstrumentCTRC
	mesage.model.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	// jsonFileName := filepath.Join("generated", "MessageRetrieval_Scenario1_Step2.json")
	xnlFileName := filepath.Join("generated", "MessageRetrieval_Scenario1_Step2_2.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
