package FedwireFundsAcknowledgement_007_001_01

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsAcknowledgement_Scenario1_Step1a_admi_CreateXML(t *testing.T) {
	var message = NewAdmi007Message()
	message.model.MessageId = "20250310QMGFNP7500070103101100FT03"
	message.model.CreatedDateTime = time.Now()
	message.model.RelationReference = "20250310B1QDRCQR000711"
	message.model.ReferenceName = "pain.013.001.07"
	message.model.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step1a_admi.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2a_admi_CreateXML(t *testing.T) {
	var message = NewAdmi007Message()
	message.model.MessageId = "20250310QMGFNP7500070203101130FT03"
	message.model.CreatedDateTime = time.Now()
	message.model.RelationReference = "20250310B1QDRCQR000712"
	message.model.ReferenceName = "pain.014.001.07"
	message.model.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step2a_admi.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario2_Step2a_admi_CreateXML(t *testing.T) {
	var message = NewAdmi007Message()
	message.model.MessageId = "20250310QMGFNP6500072203101100FT03"
	message.model.CreatedDateTime = time.Now()
	message.model.RelationReference = "20250310B1QDRCQR000722"
	message.model.ReferenceName = "pain.014.001.07"
	message.model.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step2a_admi.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario2_Step3a_admi_CreateXML(t *testing.T) {
	var message = NewAdmi007Message()
	message.model.MessageId = "20250310QMGFNP6500072203101100FT03"
	message.model.CreatedDateTime = time.Now()
	message.model.RelationReference = "20250310B1QDRCQR000723"
	message.model.ReferenceName = "pain.014.001.07"
	message.model.RequestHandling = model.SchemaValidationFailed
	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step3a_admi.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
