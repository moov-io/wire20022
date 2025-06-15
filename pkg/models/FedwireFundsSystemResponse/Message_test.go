package FedwireFundsSystemResponse

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wadearnold/wire20022/pkg/models"
)

func TestDocumentToModel01(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step2_admi.011")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "98z2cb3d0f2f3094f24a16389713541137b")
	require.Equal(t, model.EventCode, models.ConnectionCheck)
	require.Equal(t, model.EventParam, "BMQFMI01")
	require.NotNil(t, model.EventTime)
}
