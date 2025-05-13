package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/FedwireFundsSystemResponse"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsSystemResponseParseXMLFile(t *testing.T) {
	xmlFile := "../models/FedwireFundsSystemResponse/generated/ConnectionCheck_Scenario1_Step2_admi.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &FedwireFundsSystemResponse.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*FedwireFundsSystemResponse.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "98z2cb3d0f2f3094f24a16389713541137b")
	}
}

func TestFedwireFundsSystemResponseGenerateXML(t *testing.T) {
	dataModel := FedwireFundsSystemResponseDataModel()
	xmlData, err := GenerateXML(&dataModel, &FedwireFundsSystemResponse.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsSystemResponse_test.xml", xmlData)
	require.NoError(t, err)
}

func TestFedwireFundsSystemResponseRequireFieldCheck(t *testing.T) {
	dataModel := FedwireFundsSystemResponseDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &FedwireFundsSystemResponse.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestFedwireFundsSystemResponseXMLValidation(t *testing.T) {
	xmlFile := "../models/FedwireFundsSystemResponse/swiftSample/ConnectionCheck_Scenario1_Step2_admi.011"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &FedwireFundsSystemResponse.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestFedwireFundsSystemResponseAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&FedwireFundsSystemResponse.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*FedwireFundsSystemResponse.MessageHelper); ok {
		require.Equal(t, helper.EventParam.Title, "Event Parameter")
		require.Equal(t, helper.EventParam.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.EventParam.Documentation, "Describes the parameters of an event which occurred in a system.")
	}
}

func FedwireFundsSystemResponseDataModel() FedwireFundsSystemResponse.MessageModel {
	var message, _ = FedwireFundsSystemResponse.NewMessage("")
	message.Data.MessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	message.Data.EventCode = model.ConnectionCheck
	message.Data.EventParam = "BMQFMI01"
	message.Data.EventTime = time.Now()
	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
