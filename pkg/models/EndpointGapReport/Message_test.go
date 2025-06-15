package EndpointGapReport

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wadearnold/wire20022/pkg/models"
)

func TestDocumentToModel08(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "EndpointGapReport_Scenario1_Step1_camt.052_IMAD")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")

	require.Equal(t, model.MessageId, models.EndpointGapReportType)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.InputMessageAccountabilityData)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Contains(t, model.AdditionalReportInfo, "Next sequence number")
}
