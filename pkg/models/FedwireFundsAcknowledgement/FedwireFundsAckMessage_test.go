package FedwireFundsAcknowledgement

import (
	"encoding/xml"
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
	WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step1a_admi.xml", xmlData)
	require.NoError(t, err)
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
	WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step2a_admi.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario2_Step2a_admi_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP6500072203101100FT03"
	message.data.CreatedDateTime = time.Now()
	message.data.RelationReference = "20250310B1QDRCQR000722"
	message.data.ReferenceName = "pain.014.001.07"
	message.data.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step2a_admi.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario2_Step3a_admi_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP6500072203101100FT03"
	message.data.CreatedDateTime = time.Now()
	message.data.RelationReference = "20250310B1QDRCQR000723"
	message.data.ReferenceName = "pain.014.001.07"
	message.data.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step3a_admi.xml", xmlData)
	require.NoError(t, err)
}
