package PacsMessage_008_001_08

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/moov-io/wire20022/pkg/credit_transfer"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsAcknowledgement_Scenario1_Step3CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000713"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01Step3EndToEndId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f099"
	mesage.model.InstrumentPropCode = InstrumentCTRD
	mesage.model.InterBankSettAmount = credit_transfer.CurrencyAndAmount{
		Currency: "USD", Amount: 60000.00,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = credit_transfer.CurrencyAndAmount{
		Currency: "USD", Amount: 60000.00,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = credit_transfer.Agent{
		PaymentSysCode:     credit_transfer.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.InstructedAgent = credit_transfer.Agent{
		PaymentSysCode:     credit_transfer.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = credit_transfer.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = credit_transfer.Agent{
		PaymentSysCode:     credit_transfer.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: credit_transfer.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19037",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = credit_transfer.Agent{
		PaymentSysCode:     credit_transfer.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: credit_transfer.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation A"
	mesage.model.CreditorPostalAddress = credit_transfer.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.CreditorOtherTypeId = "5647772655"
	mesage.model.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: CodeCINV,
		Number:            "INV12345",
		RelatedDate:       civil.DateOf(time.Now()),
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step3.json")
	xmlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step3.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xmlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario2_Step1CreateXML(t *testing.T) {
	var mesage = NewPacs008Message()
	mesage.model.MessageId = "20250310B1QDRCQR000721"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01InstrId001"
	mesage.model.EndToEndId = "Scenario01EtoEId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f099"
	mesage.model.InstrumentPropCode = InstrumentCTRD
	mesage.model.InterBankSettAmount = credit_transfer.CurrencyAndAmount{
		Currency: "USD", Amount: 151235.88,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = credit_transfer.CurrencyAndAmount{
		Currency: "USD", Amount: 151235.88,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = credit_transfer.Agent{
		PaymentSysCode:     credit_transfer.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.model.InstructedAgent = credit_transfer.Agent{
		PaymentSysCode:     credit_transfer.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.DebtorName = "Corporation A"
	mesage.model.DebtorAddress = credit_transfer.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.model.DebtorOtherTypeId = "5647772655"
	mesage.model.DebtorAgent = credit_transfer.Agent{
		PaymentSysCode:     credit_transfer.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: credit_transfer.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.model.CreditorAgent = credit_transfer.Agent{
		PaymentSysCode:     credit_transfer.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: credit_transfer.PostalAddress{
			StreetName:     "Avenue B<",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.model.CreditorName = "Corporation B"
	mesage.model.CreditorPostalAddress = credit_transfer.PostalAddress{
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
	// jsonFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step3.json")
	xmlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step1.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xmlFileName, xmlData, 0644)
	require.NoError(t, err)
}
