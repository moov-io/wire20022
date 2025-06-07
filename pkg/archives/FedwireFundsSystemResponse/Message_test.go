package FedwireFundsSystemResponse

import (
	"path/filepath"
	"testing"

	Archive "github.com/moov-io/wire20022/pkg/archives"
	"github.com/stretchr/testify/require"
)

func TestDocumentToModel01(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step2_admi.011")
	var xmlData, err = Archive.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "98z2cb3d0f2f3094f24a16389713541137b")
	require.Equal(t, model.EventCode, Archive.ConnectionCheck)
	require.Equal(t, model.EventParam, "BMQFMI01")
	require.NotNil(t, model.EventTime)
}
