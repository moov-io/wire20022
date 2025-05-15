package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/PaymentReturn"
	"github.com/stretchr/testify/require"
)

var PaymentReturnxmlFile = "../models/PaymentReturn/swiftSample/FedwireFundsAcknowledgement_Scenario2_Step4_pacs.004"

func TestPaymentReturnParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(PaymentReturnxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &PaymentReturn.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*PaymentReturn.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310B1QDRCQR000724")
	}
}

func TestPaymentReturnGenerateXML(t *testing.T) {
	dataModel := PaymentReturnDataModel()
	xmlData, err := GenerateXML(&dataModel, &PaymentReturn.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_test.xml", xmlData)
	require.NoError(t, err)
}

func TestPaymentReturnRequireFieldCheck(t *testing.T) {
	dataModel := PaymentReturnDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &PaymentReturn.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestPaymentReturnXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(PaymentReturnxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &PaymentReturn.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestPaymentReturnAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&PaymentReturn.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*PaymentReturn.MessageHelper); ok {
		require.Equal(t, helper.OriginalMessageId.Title, "Original Message Id")
		require.Equal(t, helper.OriginalMessageId.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.OriginalMessageId.Documentation, "Point to point reference assigned by the original instructing party to unambiguously identify the original message.")
	}
}

func PaymentReturnDataModel() PaymentReturn.MessageModel {
	var message, _ = PaymentReturn.NewMessage("")
	message.Data.MessageId = "20250310ISOTEST1000912"
	message.Data.CreatedDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.OriginalMessageId = "20250310B1QDRCQR000902"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalCreationDateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   151000.74,
		Currency: "USD",
	}
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ReturnedInstructedAmount = model.CurrencyAndAmount{
		Amount:   151000.74,
		Currency: "USD",
	}
	message.Data.ChargeBearer = model.ChargeBearerSLEV
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.RtrChain = PaymentReturn.ReturnChain{
		Debtor: PaymentReturn.Party{
			Name: "Corporation B",
			Address: model.PostalAddress{
				StreetName:     "Desert View Street",
				BuildingNumber: "1",
				Floor:          "33",
				PostalCode:     "92262",
				TownName:       "Palm Springs",
				Subdivision:    "CA",
				Country:        "US",
			},
		},
		DebtorOtherTypeId: "567876543",
		DebtorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
			BankName:           "BankB",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue B",
				BuildingNumber: "25",
				PostalCode:     "19067",
				TownName:       "Yardley",
				Subdivision:    "PA",
				Country:        "US",
			},
		},
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
			BankName:           "BankA",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue A",
				BuildingNumber: "66",
				PostalCode:     "60532",
				TownName:       "Lisle",
				Subdivision:    "IL",
				Country:        "US",
			},
		},
		Creditor: PaymentReturn.Party{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CreditorAccountOtherTypeId: "5647772655",
	}
	message.Data.ReturnReasonInformation = PaymentReturn.Reason{
		Reason:                "DUPL",
		AdditionalRequestData: "Payment deiplicate. Ref:20250310B1QDRCQR000902.",
	}
	message.Data.OriginalTransactionRef = model.InstrumentCTRC

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
