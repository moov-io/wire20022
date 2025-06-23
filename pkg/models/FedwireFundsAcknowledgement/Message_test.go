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

	model, err := ParseXML(xmlData)
	if err != nil {
		t.Fatal(err)
	}
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

	model, err := ParseXML(xmlData)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP7500070203101130FT03")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.RelationReference, "20250310B1QDRCQR000712")
	require.Equal(t, model.ReferenceName, "pain.014.001.07")
	require.Equal(t, model.RequestHandling, models.SchemaValidationFailed)
}

// Test helper functions for better coverage
func TestHelperFunctions(t *testing.T) {
	t.Run("NewMessageForVersion", func(t *testing.T) {
		model := NewMessageForVersion(ADMI_007_001_01)
		require.NotNil(t, model)
		// Basic model should have zero values
		require.Empty(t, model.MessageId)
	})

	t.Run("ValidateForVersion", func(t *testing.T) {
		model := MessageModel{}
		err := model.ValidateForVersion(ADMI_007_001_01)
		require.Error(t, err) // Should fail validation with empty model
		require.Contains(t, err.Error(), "MessageId")
	})

	t.Run("GetVersionCapabilities", func(t *testing.T) {
		model := MessageModel{}
		capabilities := model.GetVersionCapabilities()
		require.NotNil(t, capabilities)
		// Should return a map of capabilities
		require.IsType(t, map[string]bool{}, capabilities)
	})

	t.Run("CheckRequiredFields", func(t *testing.T) {
		model := MessageModel{}
		err := CheckRequiredFields(model)
		require.Error(t, err) // Should fail with empty model
	})

	t.Run("BuildMessageHelper", func(t *testing.T) {
		helper := BuildMessageHelper()
		require.NotNil(t, helper)
		// Helper should be a valid object
		require.NotEmpty(t, helper)
	})
}
