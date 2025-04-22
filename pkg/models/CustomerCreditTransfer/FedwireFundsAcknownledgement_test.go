package CustomerCreditTransfer

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsAcknowledgement_Scenario1_Step3CreateXML(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageId = "20250310B1QDRCQR000713"
	mesage.data.CreatedDateTime = time.Now()
	mesage.data.NumberOfTransactions = 1
	mesage.data.SettlementMethod = model.SettlementCLRG
	mesage.data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.data.InstructionId = "Scenario01InstrId001"
	mesage.data.EndToEndId = "Scenario01Step3EndToEndId001"
	mesage.data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f099"
	mesage.data.InstrumentPropCode = model.InstrumentCTRD
	mesage.data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 60000.00,
	}
	mesage.data.InterBankSettDate = model.FromTime(time.Now())
	mesage.data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 60000.00,
	}
	mesage.data.ChargeBearer = model.ChargeBearerSLEV
	mesage.data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.data.DebtorName = "Corporation A"
	mesage.data.DebtorAddress = model.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.data.DebtorOtherTypeId = "5647772655"
	mesage.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19037",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.data.CreditorAgent = model.Agent{
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
	mesage.data.CreditorName = "Corporation A"
	mesage.data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.data.CreditorOtherTypeId = "5647772655"
	mesage.data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV12345",
		RelatedDate:       model.FromTime(time.Now()),
	}
	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step3.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step3_pacs.008")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step3.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario2_Step1CreateXML(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageId = "20250310B1QDRCQR000721"
	mesage.data.CreatedDateTime = time.Now()
	mesage.data.NumberOfTransactions = 1
	mesage.data.SettlementMethod = model.SettlementCLRG
	mesage.data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.data.InstructionId = "Scenario01InstrId001"
	mesage.data.EndToEndId = "Scenario01EtoEId001"
	mesage.data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.data.InstrumentPropCode = model.InstrumentCTRC
	mesage.data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 151235.88,
	}
	mesage.data.InterBankSettDate = model.FromTime(time.Now())
	mesage.data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 151235.88,
	}
	mesage.data.ChargeBearer = model.ChargeBearerSLEV
	mesage.data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.data.DebtorName = "Corporation A"
	mesage.data.DebtorAddress = model.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.data.DebtorOtherTypeId = "5647772655"
	mesage.data.DebtorAgent = model.Agent{
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
	mesage.data.CreditorAgent = model.Agent{
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
	mesage.data.CreditorName = "Corporation B"
	mesage.data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.data.CreditorOtherTypeId = "567876543"
	mesage.data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step1_pacs.008")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
