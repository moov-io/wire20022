package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/DrawdownResponse"
	"github.com/stretchr/testify/require"
)

var DrawdownResponsexmlFile = "../models/DrawdownResponse/swiftSample/Drawdowns_Scenario1_Step2_pain.014"

func TestDrawdownResponseParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(DrawdownResponsexmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &DrawdownResponse.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*DrawdownResponse.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310B1QDRCQR000602")
	}
}

func TestDrawdownResponseGenerateXML(t *testing.T) {
	dataModel := DrawdownResponseDataModel()
	xmlData, err := GenerateXML(&dataModel, &DrawdownResponse.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("DrawdownResponse_test.xml", xmlData)
	require.NoError(t, err)
}

func TestDrawdownResponseRequireFieldCheck(t *testing.T) {
	dataModel := DrawdownResponseDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &DrawdownResponse.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestDrawdownResponseXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(DrawdownResponsexmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &DrawdownResponse.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestDrawdownResponseAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&DrawdownResponse.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*DrawdownResponse.MessageHelper); ok {
		require.Equal(t, helper.OriginalMessageId.Title, "Original Message Identification")
		require.Equal(t, helper.OriginalMessageId.Type, "IMAD_FedwireFunds_1 (based on string) minLength: 22 maxLength: 22 pattern: [0-9]{8}[A-Z0-9]{8}[0-9]{6}")
		require.Equal(t, helper.OriginalMessageId.Documentation, "Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.")
	}
}

func DrawdownResponseDataModel() DrawdownResponse.MessageModel {
	var message, _ = DrawdownResponse.NewMessage("")
	message.Data.MessageId = "20250310B1QDRCQR000712"
	message.Data.CreateDatetime = time.Now()
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
	message.Data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.OriginalMessageId = "20250310B1QDRCQR000711"
	message.Data.OriginalMessageNameId = "pain.013.001.07"
	message.Data.OriginalCreationDateTime = time.Now()
	message.Data.OriginalPaymentInfoId = "20250310B1QDRCQR000711"
	message.Data.TransactionInformationAndStatus = DrawdownResponse.TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01InstrId001",
		OriginalEndToEndId:    "Scenario01Step1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f078",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
