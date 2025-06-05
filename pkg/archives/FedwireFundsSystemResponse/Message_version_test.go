package FedwireFundsSystemResponse

import (
	"encoding/xml"
	"testing"
	"time"

	Archive "github.com/moov-io/wire20022/pkg/archives"
	"github.com/stretchr/testify/require"
)

func TestVersion01(t *testing.T) {
	modelName := ADMI_011_001_01
	xmlName := "FedwireFundsSystemResponse_01.xml"

	dataModel := FedwireFundsSystemResponseDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "98z2cb3d0f2f3094f24a16389713541137b")
	require.Equal(t, model.EventCode, Archive.ConnectionCheck)
	require.Equal(t, model.EventParam, "BMQFMI01")
	require.NotNil(t, model.EventTime)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "98z2cb3d0f2f3094f24a16389713541137b"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "98z2cb3d0f2f3094f24a16389713541137b"
}

func FedwireFundsSystemResponseDataModel() MessageModel {
	message := MessageModel{}
	message.MessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	message.EventCode = Archive.ConnectionCheck
	message.EventParam = "BMQFMI01"
	message.EventTime = time.Now() // Set to nil for now, as we don't have a specific time
	return message
}
