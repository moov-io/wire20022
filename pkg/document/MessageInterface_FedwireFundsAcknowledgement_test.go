package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/FedwireFundsAcknowledgement"
	"github.com/stretchr/testify/require"
)
var FedwireFundsAcknowledgementxmlFile = "../models/FedwireFundsAcknowledgement/swiftSample/FedwireFundsAcknowledgement_Scenario1_Step1a_admi.007"
func TestFedwireFundsAcknowledgementParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(FedwireFundsAcknowledgementxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &FedwireFundsAcknowledgement.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*FedwireFundsAcknowledgement.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310QMGFNP7500070103101100FT03")
	}
}

func TestFedwireFundsAcknowledgementGenerateXML(t *testing.T) {
	dataModel := FedwireFundsAcknowledgementDataModel()
	xmlData, err := GenerateXML(&dataModel, &FedwireFundsAcknowledgement.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_test.xml", xmlData)
	require.NoError(t, err)
}

func TestFedwireFundsAcknowledgementRequireFieldCheck(t *testing.T) {
	dataModel := FedwireFundsAcknowledgementDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &FedwireFundsAcknowledgement.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestFedwireFundsAcknowledgementXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(FedwireFundsAcknowledgementxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &FedwireFundsAcknowledgement.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestFedwireFundsAcknowledgementAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&FedwireFundsAcknowledgement.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*FedwireFundsAcknowledgement.MessageHelper); ok {
		require.Equal(t, helper.RelationReference.Title, "Relation Reference")
		require.Equal(t, helper.RelationReference.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.RelationReference.Documentation, "Unambiguous reference to a previous message having a business relevance with this message.")
	}
}

func FedwireFundsAcknowledgementDataModel() FedwireFundsAcknowledgement.MessageModel {
	var message, _ = FedwireFundsAcknowledgement.NewMessage("")
	message.Data.MessageId = "20250310QMGFNP7500070203101130FT03"
	message.Data.CreatedDateTime = time.Now()
	message.Data.RelationReference = "20250310B1QDRCQR000712"
	message.Data.ReferenceName = "pain.014.001.07"
	message.Data.RequestHandling = model.SchemaValidationFailed

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
