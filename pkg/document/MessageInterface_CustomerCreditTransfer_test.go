package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
	"github.com/stretchr/testify/require"
)

var CustomerCreditTransferxmlFile = "../models/CustomerCreditTransfer/swiftSample/CustomerCreditTransfer_Scenario1_Step1_pacs.008"

func TestCustomerCreditTransferParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(CustomerCreditTransferxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &CustomerCreditTransfer.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*CustomerCreditTransfer.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310B1QDRCQR000001")
	}
}

func TestCustomerCreditTransferGenerateXML(t *testing.T) {
	dataModel := CustomerCreditTransferDataModel()
	xmlData, err := GenerateXML(&dataModel, &CustomerCreditTransfer.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_test.xml", xmlData)
	require.NoError(t, err)
}

func TestCustomerCreditTransferRequireFieldCheck(t *testing.T) {
	dataModel := CustomerCreditTransferDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &CustomerCreditTransfer.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestCustomerCreditTransferXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(CustomerCreditTransferxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &CustomerCreditTransfer.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestCustomerCreditTransferAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&CustomerCreditTransfer.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*CustomerCreditTransfer.MessageHelper); ok {
		require.Equal(t, helper.DebtorOtherTypeId.Title, "Debtor Other Type Id")
		require.Equal(t, helper.DebtorOtherTypeId.Type, "Max34Text (based on string) minLength: 1 maxLength: 34")
		require.Equal(t, helper.DebtorOtherTypeId.Documentation, "Unique identification of an account, as assigned by the account servicer, using an identification scheme.")
	}
}

func CustomerCreditTransferDataModel() CustomerCreditTransfer.MessageModel {
	var mesage, _ = CustomerCreditTransfer.NewMessage("")
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
	mesage.Data.CreditorOtherTypeId = "567876543"
	mesage.Data.RemittanceInfor = CustomerCreditTransfer.RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	if cErr != nil {
		return mesage.Data
	}
	return mesage.Data
}
