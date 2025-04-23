package FedwireFundsAcknowledgement

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsAcknowledgement_Scenario1_Step1a_admi_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP7500070103101100FT03"
	message.data.CreatedDateTime = time.Now()
	message.data.RelationReference = "20250310B1QDRCQR000711"
	message.data.ReferenceName = "pain.013.001.07"
	message.data.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step1a_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step1a_admi.007")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step1a_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2a_admi_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP7500070203101130FT03"
	message.data.CreatedDateTime = time.Now()
	message.data.RelationReference = "20250310B1QDRCQR000712"
	message.data.ReferenceName = "pain.014.001.07"
	message.data.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step2a_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step2a_admi.007")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step2a_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario2_Step2a_admi_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP6500072203101100FT03"
	message.data.CreatedDateTime = time.Now()
	message.data.RelationReference = "20250310B1QDRCQR000722"
	message.data.ReferenceName = "camt.056.001.08"
	message.data.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step2a_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step2a_admi.007")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step2a_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario2_Step3a_admi_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP6500072303101100FT03"
	message.data.CreatedDateTime = time.Now()
	message.data.RelationReference = "20250310B1QDRCQR000723"
	message.data.ReferenceName = "camt.029.001.09"
	message.data.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step3a_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step3a_admi.007")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step3a_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
