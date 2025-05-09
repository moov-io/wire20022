package CustomerCreditTransfer

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestInvestigations_Scenario1_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000001"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario01InstrId001"
	mesage.Data.EndToEndId = "Scenario01EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerSLEV
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.Data.DebtorName = "Corporation A"
	mesage.Data.DebtorAddress = model.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.Data.DebtorOtherTypeId = "5647772655"
	mesage.Data.DebtorAgent = model.Agent{
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
	mesage.Data.CreditorAgent = model.Agent{
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
	mesage.Data.CreditorName = "Corporation B"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.Data.CreditorOtherTypeId = "5678765"
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario1_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario1_Step1_pacs.008")
	genterated := filepath.Join("generated", "Investigations_Scenario1_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario2_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000002"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario01InstrId001"
	mesage.Data.EndToEndId = "Scenario01EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerSLEV
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.Data.DebtorName = "Corporation A"
	mesage.Data.DebtorAddress = model.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.Data.DebtorOtherTypeId = "5647772655"
	mesage.Data.DebtorAgent = model.Agent{
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
	mesage.Data.CreditorAgent = model.Agent{
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
	mesage.Data.CreditorName = "Corporation B"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.Data.CreditorOtherTypeId = "567876543"
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario2_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario2_Step1_pacs.008")
	genterated := filepath.Join("generated", "Investigations_Scenario2_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}

func TestInvestigations_Scenario3_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000007"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Variation2InstrId001"
	mesage.Data.EndToEndId = "Variation2EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerSLEV
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.Data.DebtorName = "Corporation A"
	mesage.Data.DebtorAddress = model.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.Data.DebtorOtherTypeId = "5647772655"
	mesage.Data.DebtorAgent = model.Agent{
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
	mesage.Data.CreditorAgent = model.Agent{
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
	mesage.Data.CreditorName = "Corporation B"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.Data.CreditorOtherTypeId = "567876543"
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario3_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario3_Step1_pacs.008")
	genterated := filepath.Join("generated", "Investigations_Scenario3_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
