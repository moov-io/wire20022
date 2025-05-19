package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/ConnectionCheck"
	"github.com/stretchr/testify/require"
)

var ConnectionCheckxmlFile = "../models/ConnectionCheck/swiftSample/ConnectionCheck_Scenario1_Step1_admi.004"

func TestConnectionCheckParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(ConnectionCheckxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &ConnectionCheck.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*ConnectionCheck.MessageModel); ok {
		require.Equal(t, msgModel.EventType, "PING")
	}
}

func TestConnectionCheckGenerateXML(t *testing.T) {
	dataModel := ConnectionCheckDataModel()
	xmlData, err := GenerateXML(&dataModel, &ConnectionCheck.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("ConnectionCheck_test.xml", xmlData)
	require.NoError(t, err)
}

func TestConnectionCheckRequireFieldCheck(t *testing.T) {
	dataModel := ConnectionCheckDataModel()
	dataModel.EventType = ""
	dataModel.EventParam = ""
	valid, err := RequireFieldCheck(&dataModel, &ConnectionCheck.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: EventType, EventParam")
}

func TestConnectionCheckXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(ConnectionCheckxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &ConnectionCheck.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestConnectionCheckAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&ConnectionCheck.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*ConnectionCheck.MessageHelper); ok {
		require.Equal(t, helper.EventParam.Title, "Event Parameter")
		require.Equal(t, helper.EventParam.Type, "EndpointIdentifier_FedwireFunds_1 (based on string) minLength: 8 maxLength: 8 pattern: [A-Z0-9]{8,8}")
		require.Equal(t, helper.EventParam.Documentation, "Describes the parameters of an event which occurred in a system.")
	}
}

func ConnectionCheckDataModel() ConnectionCheck.MessageModel {
	var mesage, _ = ConnectionCheck.NewMessage("")
	mesage.Data.EventType = "PING"
	mesage.Data.EventParam = "BMQFMI01"
	mesage.Data.EventTime = time.Now()

	cErr := mesage.CreateDocument()
	if cErr != nil {
		return mesage.Data
	}
	return mesage.Data
}
