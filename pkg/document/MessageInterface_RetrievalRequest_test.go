package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/RetrievalRequest"
	"github.com/stretchr/testify/require"
)

func TestRetrievalRequestParseXMLFile(t *testing.T) {
	xmlFile := "../models/RetrievalRequest/generated/MessageRetrieval_Scenario1_Step1_admi.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &RetrievalRequest.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*RetrievalRequest.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250301011104238MRSc1Step1MsgId")
	}
}

func TestRetrievalRequestGenerateXML(t *testing.T) {
	dataModel := RetrievalRequestDataModel()
	xmlData, err := GenerateXML(&dataModel, &RetrievalRequest.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("RetrievalRequest_test.xml", xmlData)
	require.NoError(t, err)
}

func TestRetrievalRequestRequireFieldCheck(t *testing.T) {
	dataModel := RetrievalRequestDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &RetrievalRequest.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestRetrievalRequestXMLValidation(t *testing.T) {
	xmlFile := "../models/RetrievalRequest/swiftSample/MessageRetrieval_Scenario1_Step1_admi.006"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &RetrievalRequest.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestRetrievalRequestAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&RetrievalRequest.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*RetrievalRequest.MessageHelper); ok {
		require.Equal(t, helper.RecipientId.Title, "Recipient Id")
		require.Equal(t, helper.RecipientId.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.RecipientId.Documentation, "Unique identification of the party.")
	}
}

func RetrievalRequestDataModel() RetrievalRequest.MessageModel {
	var message, _ = RetrievalRequest.NewMessage("")
	message.Data.MessageId = "20250301011104238MRSc2Step1MsgId"
	message.Data.CreatedDateTime = time.Now()
	message.Data.RequestType = model.RequestSent
	message.Data.BusinessDate = model.FromTime(time.Now())
	message.Data.SequenceRange = model.SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000003",
	}
	message.Data.RecipientId = "B1QDRCQR"
	message.Data.RecipientIssuer = "NA"

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
