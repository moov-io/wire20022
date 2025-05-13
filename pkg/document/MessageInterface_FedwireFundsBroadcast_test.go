package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/FedwireFundsBroadcast"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsBroadcastParseXMLFile(t *testing.T) {
	xmlFile := "../models/FedwireFundsBroadcast/generated/FedwireFundsBroadcast_admi_ADHC.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &FedwireFundsBroadcast.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*FedwireFundsBroadcast.MessageModel); ok {
		require.Equal(t, msgModel.EventCode, model.FundEventType("ADHC"))
	}
}

func TestFedwireFundsBroadcastGenerateXML(t *testing.T) {
	dataModel := FedwireFundsBroadcastDataModel()
	xmlData, err := GenerateXML(&dataModel, &FedwireFundsBroadcast.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsBroadcast_test.xml", xmlData)
	require.NoError(t, err)
}

func TestFedwireFundsBroadcastRequireFieldCheck(t *testing.T) {
	dataModel := FedwireFundsBroadcastDataModel()
	dataModel.EventCode = ""
	valid, err := RequireFieldCheck(&dataModel, &FedwireFundsBroadcast.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: EventCode")
}

func TestFedwireFundsBroadcastXMLValidation(t *testing.T) {
	xmlFile := "../models/FedwireFundsBroadcast/swiftSample/FedwireFundsBroadcast_admi.004_ADHC"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &FedwireFundsBroadcast.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestFedwireFundsBroadcastAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&FedwireFundsBroadcast.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*FedwireFundsBroadcast.MessageHelper); ok {
		require.Equal(t, helper.EventDescription.Title, "Event Description")
		require.Equal(t, helper.EventDescription.Type, "Max1000Text (based on string) minLength: 1 maxLength: 1000")
		require.Equal(t, helper.EventDescription.Documentation, "Free text used to describe an event which occurred in a system.")
	}
}

func FedwireFundsBroadcastDataModel() FedwireFundsBroadcast.MessageModel {
	var message, _ = FedwireFundsBroadcast.NewMessage("")
	message.Data.EventCode = model.AdHoc
	message.Data.EventParam = model.FromTime(time.Now())
	message.Data.EventDescription = "The Fedwire Funds Service will open the test environment 15 minutes earlier on 03/13/2025"
	message.Data.EventTime = time.Now()
	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
