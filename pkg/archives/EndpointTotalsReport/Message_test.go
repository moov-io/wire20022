package EndpointTotalsReport

import (
	"path/filepath"
	"testing"

	Archive "github.com/moov-io/wire20022/pkg/archives"
	"github.com/stretchr/testify/require"
)

func TestDocumentToModel08(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "EndpointTotalsReport_Scenario1_Step2_camt.052_ETOT")
	var xmlData, err = Archive.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, Archive.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, Archive.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, Archive.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, Archive.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, Archive.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, Archive.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, Archive.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")
}
