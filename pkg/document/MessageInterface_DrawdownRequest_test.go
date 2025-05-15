package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/DrawdownRequest"
	"github.com/stretchr/testify/require"
)

var DrawdownRequestxmlFile = "../models/DrawdownRequest/swiftSample/Drawdowns_Scenario1_Step1_pain.013"

func TestDrawdownRequestParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(DrawdownRequestxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &DrawdownRequest.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*DrawdownRequest.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310B1QDRCQR000601")
	}
}

func TestDrawdownRequestGenerateXML(t *testing.T) {
	dataModel := DrawdownRequestDataModel()
	xmlData, err := GenerateXML(&dataModel, &DrawdownRequest.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("DrawdownRequest_test.xml", xmlData)
	require.NoError(t, err)
}

func TestDrawdownRequestRequireFieldCheck(t *testing.T) {
	dataModel := DrawdownRequestDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &DrawdownRequest.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestDrawdownRequestXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(DrawdownRequestxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &DrawdownRequest.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestDrawdownRequestAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&DrawdownRequest.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*DrawdownRequest.MessageHelper); ok {
		require.Equal(t, helper.PaymentMethod.Title, "Payment Method")
		require.Equal(t, helper.PaymentMethod.Type, "SettlementMethodType(Clearing, Gross, Net, DeferredNet, DeliveryVsPayment, PaymentVsPayment, PaymentVsDelivery, PaymentVsPayment)")
		require.Equal(t, helper.PaymentMethod.Documentation, "Method used to settle a payment transaction.")
	}
}

func DrawdownRequestDataModel() DrawdownRequest.MessageModel {
	var message, _ = DrawdownRequest.NewMessage("")
	message.Data.MessageId = "20250310B1QDRCQR000601"
	message.Data.CreateDatetime = time.Now()
	message.Data.NumberofTransaction = "1"
	message.Data.InitiatingParty = model.PartyIdentify{
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
	}
	message.Data.PaymentInfoId = "20250310B1QDRCQR000601"
	message.Data.PaymentMethod = DrawdownRequest.CreditTransform
	message.Data.RequestedExecutDate = model.FromTime(time.Now())
	message.Data.Debtor = model.PartyIdentify{
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
	}
	message.Data.DebtorAccountOtherId = "92315266453"
	message.Data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.CreditTransTransaction = DrawdownRequest.CreditTransferTransaction{
		PaymentInstructionId: "Scenario01Step1InstrId001",
		PaymentEndToEndId:    "Scenario1EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		PayCategoryType:      DrawdownRequest.IntraCompanyPayment,
		PayRequestType:       DrawdownRequest.DrawDownRequestCredit,
		Amount: model.CurrencyAndAmount{
			Amount:   6000000.00,
			Currency: "USD",
		},
		ChargeBearer: DrawdownRequest.ChargeBearerSLEV,
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		Creditor: model.PartyIdentify{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain HIlls",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CrediorAccountOtherId: "5647772655",
		RemittanceInformation: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
