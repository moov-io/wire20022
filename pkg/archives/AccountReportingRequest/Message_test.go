package ArchiveAccountReportingRequest

import (
	"testing"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

var sample1XML = "../../models/AccountReportingRequest/swiftSample/EndpointDetailsReport_Scenario1_Step1_camt.060_DTLS"
var sample2XML = "../../models/AccountReportingRequest/swiftSample/AccountBalanceReport_Scenario1_Step1_camt.060_ABAR_MM"

func TestDocumentElementToModelOne(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(sample1XML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435DTLSrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")
}
func TestDocumentElementToModelTwo(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(sample2XML)
	require.NoError(t, err, "Failed to read XML file")
	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, CAMTReportType("ABAR"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
}
