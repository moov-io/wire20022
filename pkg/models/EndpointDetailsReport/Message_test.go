package EndpointDetailsReport

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDocumentToModel08(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "EndpointDetailsReport_Scenario1_Step2_camt.052_DTLS")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := ParseXML(xmlData)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.NotNil(t, model.BusinessQuery, "BusinessQuery should be populated for V8 samples")
	require.Equal(t, model.BusinessQuery.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BusinessQuery.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BusinessQuery.BussinessQueryCreateDatetime)
	require.NotNil(t, model.Reporting, "Reporting should be populated for V8 samples")
	require.Equal(t, model.Reporting.ReportingSequence.FromSeq, "000001")
	require.Equal(t, model.Reporting.ReportingSequence.ToSeq, "000100")
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[0].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)
}
