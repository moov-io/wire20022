package FedwireFundsAcknowledgement

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion01(t *testing.T) {
	modelName := ADMI_007_001_01
	xmlName := "FedwireFundsAcknowledgement_01.xml"

	dataModel := FedwireFundsAcknowledgementDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP7500070103101100FT03")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.RelationReference, "20250310B1QDRCQR000711")
	require.Equal(t, model.ReferenceName, "pain.013.001.07")
	require.Equal(t, model.RequestHandling, models.SchemaValidationFailed)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RctAck.MsgId.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}

func FedwireFundsAcknowledgementDataModel() MessageModel {
	message := MessageModel{}
	message.MessageId = "20250310QMGFNP7500070103101100FT03"
	message.CreatedDateTime = time.Now()
	message.RelationReference = "20250310B1QDRCQR000711"
	message.ReferenceName = "pain.013.001.07"
	message.RequestHandling = models.SchemaValidationFailed
	return message
}
