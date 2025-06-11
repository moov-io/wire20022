package FedwireFundsAcknowledgement

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDocumentToModel01(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step1a_admi.007")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP7500070103101100FT03")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.RelationReference, "20250310B1QDRCQR000711")
	require.Equal(t, model.ReferenceName, "pain.013.001.07")
	require.Equal(t, model.RequestHandling, models.SchemaValidationFailed)
}

func TestDocumentToModel02(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step2a_admi.007")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP7500070203101130FT03")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.RelationReference, "20250310B1QDRCQR000712")
	require.Equal(t, model.ReferenceName, "pain.014.001.07")
	require.Equal(t, model.RequestHandling, models.SchemaValidationFailed)
}
