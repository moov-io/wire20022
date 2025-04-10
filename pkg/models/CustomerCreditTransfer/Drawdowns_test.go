package CustomerCreditTransfer

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario1_Step3CreateXML(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageId = "20250310B1QDRCQR000634"
	mesage.data.CreatedDateTime = time.Now()
	mesage.data.NumberOfTransactions = 1
	mesage.data.SettlementMethod = model.SettlementCLRG
	mesage.data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.data.InstructionId = "Scenario04Step5InstrId001"
	mesage.data.EndToEndId = "Scenario4EndToEndId001"
	mesage.data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f066"
	mesage.data.InstrumentPropCode = model.InstrumentCTRD
	mesage.data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 6000000.00,
	}
	mesage.data.InterBankSettDate = model.FromTime(time.Now())
	mesage.data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 6000000.00,
	}
	mesage.data.ChargeBearer = ChargeBearerSLEV
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
	mesage.data.DebtorOtherTypeId = "92315266453"
	mesage.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "85268",
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
		BuildingNumber: "1167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.data.CreditorOtherTypeId = "5647772655"
	mesage.data.RemittanceInfor = RemittanceDocument{
		UnstructuredRemitInfo: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := xml.MarshalIndent(mesage, "", "\t")
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "Drawdowns_Scenario1_Step3.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario1_Step3.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario1_Step5CreateXML(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageId = "20250310B1QDRCQR000603"
	mesage.data.CreatedDateTime = time.Now()
	mesage.data.NumberOfTransactions = 1
	mesage.data.SettlementMethod = model.SettlementCLRG
	mesage.data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.data.InstructionId = "Scenario01Step3InstrId001"
	mesage.data.EndToEndId = "Scenario1EndToEndId001"
	mesage.data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f066"
	mesage.data.InstrumentPropCode = model.InstrumentCTRD
	mesage.data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 6000000.00,
	}
	mesage.data.InterBankSettDate = model.FromTime(time.Now())
	mesage.data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 6000000.00,
	}
	mesage.data.ChargeBearer = ChargeBearerSLEV
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
	mesage.data.DebtorOtherTypeId = "92315266453"
	mesage.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			RoomNumber:     "60532",
			PostalCode:     "85268",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "167565",
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
		UnstructuredRemitInfo: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := xml.MarshalIndent(mesage, "", "\t")
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "Drawdowns_Scenario1_Step5.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario1_Step5.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
