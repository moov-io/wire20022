package CustomerCreditTransfer

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario1_Step3CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
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
	mesage.data.DebtorOtherTypeId = "92315266453"
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
		UnstructuredRemitInfo: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}
	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario1_Step3.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step3_pacs.008")
	genterated := filepath.Join("generated", "Drawdowns_Scenario1_Step3.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario5_Step5_pacs_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
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
	mesage.data.DebtorOtherTypeId = "92315266453"
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
		UnstructuredRemitInfo: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}
	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario5_Step5_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario5_Step5_pacs.008")
	genterated := filepath.Join("generated", "Drawdowns_Scenario5_Step5_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
