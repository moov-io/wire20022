package FedwireFundsSystemResponse

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDocumentToModel01(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step2_admi.011")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := ParseXML(xmlData)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "98z2cb3d0f2f3094f24a16389713541137b")
	require.Equal(t, model.EventCode, models.ConnectionCheck)
	require.Equal(t, model.EventParam, "BMQFMI01")
	require.NotNil(t, model.EventTime)
}

// Test helper functions for better coverage
func TestHelperFunctions(t *testing.T) {
	t.Run("NewMessageForVersion", func(t *testing.T) {
		model := NewMessageForVersion(ADMI_011_001_01)
		require.NotNil(t, model)
		require.Empty(t, model.MessageId)
	})

	t.Run("ValidateForVersion", func(t *testing.T) {
		model := MessageModel{}
		err := model.ValidateForVersion(ADMI_011_001_01)
		require.Error(t, err)
		require.Contains(t, err.Error(), "MessageId")
	})

	t.Run("GetVersionCapabilities", func(t *testing.T) {
		model := MessageModel{}
		capabilities := model.GetVersionCapabilities()
		require.NotNil(t, capabilities)
		require.IsType(t, map[string]bool{}, capabilities)
	})

	t.Run("CheckRequiredFields", func(t *testing.T) {
		model := MessageModel{}
		err := CheckRequiredFields(model)
		require.Error(t, err)
	})

	t.Run("BuildMessageHelper", func(t *testing.T) {
		helper := BuildMessageHelper()
		require.NotNil(t, helper)
		require.NotEmpty(t, helper)
	})
}
