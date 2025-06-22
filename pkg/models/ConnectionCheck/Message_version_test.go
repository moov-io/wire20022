package ConnectionCheck

import (
	"encoding/xml"
	"testing"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion01(t *testing.T) {
	/*Create Document from Model*/
	var doc01, err = DocumentWith(ConnectionCheckDataModel(), ADMI_004_001_01)
	require.NoError(t, err, "Failed to create document")

	/*Validate Check for created Document*/
	vErr := doc01.Validate()
	require.NoError(t, vErr, "Failed to validate document")

	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc01, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate("ConnectionCheck_01.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/ConnectionCheck_01.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.EventType, "PING")
	require.Equal(t, model.EventParam, "BMQFMI01")
	require.NotNil(t, model.EventTime, "EventTime should not be nil")

	/*Validation check*/
	model.EventType = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, ADMI_004_001_01)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy Admi00400101.EvtInf.EvtCd failed: failed to set EventType: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 4")
	model.EventType = "PING"

	/*Require field check*/
	model.EventType = ""
	_, err = DocumentWith(*model, ADMI_004_001_01)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"EventType\": is required: required field missing")
	model.EventType = "PING"

	/*Access to Helper*/
	require.Equal(t, "Event Type", BuildMessageHelper().EventType.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().EventType.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().EventType.Documentation, "Proprietary code used to specify an")
}

func TestVersion02(t *testing.T) {
	/*Create Document from Model*/
	var doc02, err = DocumentWith(ConnectionCheckDataModel(), ADMI_004_001_02)
	require.NoError(t, err, "Failed to create document")

	/*Validate Check for created Document*/
	vErr := doc02.Validate()
	require.NoError(t, vErr, "Failed to validate document")

	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc02, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate("ConnectionCheck_02.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/ConnectionCheck_02.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.EventType, "PING")
	require.Equal(t, model.EventParam, "BMQFMI01")
	require.NotNil(t, model.EventTime, "EventTime should not be nil")

	/*Validation check*/
	model.EventType = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, ADMI_004_001_02)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy SysEvtNtfctn.EvtInf.EvtCd failed: failed to set EventType: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 4")
	model.EventType = "PING"

	/*Require field check*/
	model.EventType = ""
	_, err = DocumentWith(*model, ADMI_004_001_02)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"EventType\": is required: required field missing")
	model.EventType = "PING"

	/*Access to Helper*/
	require.Equal(t, "Event Type", BuildMessageHelper().EventType.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().EventType.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().EventType.Documentation, "Proprietary code used to specify an")
}
