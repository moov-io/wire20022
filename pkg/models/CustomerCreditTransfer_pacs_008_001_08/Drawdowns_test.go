package CustomerCreditTransfer_pacs_008_001_08

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/stretchr/testify/require"
)

func TestDrawDownsCreateXML(t *testing.T) {
	var mesage = NewCustomerCreditTransferMessage()
	mesage.model.MessageId = "20250310B1QDRCQR000603"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.NumberOfTransactions = 1
	mesage.model.SettlementMethod = SettlementCLRG
	mesage.model.CommonClearingSysCode = ClearingSysFDW
	mesage.model.InstructionId = "Scenario01Step3InstrId001"
	mesage.model.EndToEndId = "Scenario1EndToEndId001"
	mesage.model.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f066"
	mesage.model.InstrumentPropCode = InstrumentCTRD
	mesage.model.InterBankSettAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 6000000.00,
	}
	mesage.model.InterBankSettDate = civil.DateOf(time.Now())
	mesage.model.InstructedAmount = CurrencyAndAmount{
		Currency: "USD", Amount: 6000000.00,
	}
	mesage.model.ChargeBearer = ChargeBearerSLEV
	mesage.model.InstructingAgents = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.model.InstructedAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
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
	mesage.model.DebtorOtherTypeId = "92315266453"
	mesage.model.DebtorAgent = Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
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
		PaymentSysMemberId: "167565",
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
		UnstructuredRemitInfo: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}
	mesage.CreateDocument()
	jsonData, err := mesage.GetJson()
	require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	jsonFileName := filepath.Join("generated", "Drawdowns.json")
	xnlFileName := filepath.Join("generated", "Drawdowns.xml")
	err = os.WriteFile(jsonFileName, jsonData, 0644)
	require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
