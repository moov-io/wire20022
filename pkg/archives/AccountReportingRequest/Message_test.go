package ArchiveAccountReportingRequest

import (
	"testing"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

var sampleXML = "../../models/AccountReportingRequest/swiftSample/AccountBalanceReport_Scenario1_Step1_camt.060_ABAR_MM"

func TestDocumentElementToModel(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, CAMTReportType("ABAR"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
}
