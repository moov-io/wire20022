package ActivityReport

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_08"
	"github.com/stretchr/testify/require"
	"github.com/wadearnold/wire20022/pkg/models"
)

func TestDocumentToModel08(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "ActivityReport_Scenario1_Step1_camt.052_ACTR")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011", "Failed to get UniqueTransactionReference")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC, "Failed to get LocalInstrumentChoice")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011", "Failed to get UniqueTransactionReference")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC, "Failed to get LocalInstrumentChoice")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011", "Failed to get UniqueTransactionReference")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC, "Failed to get LocalInstrumentChoice")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
}
func TestModelToDocument08_One(t *testing.T) {
	dataModel := ActivityReportDataModel()
	var doc08, err = DocumentWith(dataModel, CAMT_052_001_08)
	require.NoError(t, err, "Failed to create document")
	if Doc08, ok := doc08.(*camt_052_001_08.Document); ok {
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.GrpHdr.MsgId), "ACTR", "Failed to get MessageId")
		require.NotNil(t, Doc08.BkToCstmrAcctRpt.GrpHdr.CreDtTm, "Failed to get CreatedDateTime")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb), "1", "Failed to get PageNumber")
		require.Equal(t, bool(Doc08.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd), true, "Failed to get LastPageIndicator")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Id), "EDAY", "Failed to get ReportId")
		require.NotNil(t, Doc08.BkToCstmrAcctRpt.Rpt[0].CreDtTm, "Failed to get ReportCreateDateTime")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id), "011104238", "Failed to get AccountOtherId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtries.NbOfNtries), "61", "Failed to get TotalEntries")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.NbOfNtries), "29", "Failed to get TotalCreditEntries")
		require.Equal(t, float64(*Doc08.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.Sum), 8775299.29, "Failed to get TotalCreditEntries")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.NbOfNtries), "27", "Failed to get TotalDebitEntries")
		require.Equal(t, float64(*Doc08.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.Sum), 9932294.43, "Failed to get TotalDebitEntries")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].NbOfNtries), "0", "Failed to get TotalEntriesPerBankTransactionCode")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].BkTxCd.Prtry.Cd), "SENT", "Failed to get TotalEntriesPerBankTransactionCode")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[1].NbOfNtries), "5", "Failed to get TotalEntriesPerBankTransactionCode")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[1].BkTxCd.Prtry.Cd), "RCVD", "Failed to get TotalEntriesPerBankTransactionCode")
		require.Equal(t, float64(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].Amt.Value), 240.67, "Failed to get Amount")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].Amt.Ccy), "USD", "Failed to get Currency")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].CdtDbtInd), "DBIT", "Failed to get CreditDebitIndicator")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].Sts.Cd), "BOOK", "Failed to get Status")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].BkTxCd.Prtry.Cd), "DBIT", "Failed to get BankTransactionCode")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].AddtlInfInd.MsgNmId), "pacs.008.001.08", "Failed to get MessageNameId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].NtryDtls[0].TxDtls[0].Refs.InstrId), "20250331231981435InstructionId00001", "Failed to get InstructionId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].NtryDtls[0].TxDtls[0].Refs.UETR), "8a562c67-ca16-48ba-b074-65581be6f011", "Failed to get UniqueTransactionReference")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get PaymentSysCode")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId), "231981435", "Failed to get PaymentSysMemberId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get PaymentSysCode")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId), "011104238", "Failed to get PaymentSysMemberId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].NtryDtls[0].TxDtls[0].LclInstrm.Prtry), "CTRC", "Failed to get LocalInstrumentChoice")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Tp), "BPRD", "Failed to get RelatedDatesProprietary")
		require.NotNil(t, Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[0].NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Dt.DtTm, "Failed to get RelatedDateTime")
		require.Equal(t, float64(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].Amt.Value), 1000.00, "Failed to get Amount")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].Amt.Ccy), "USD", "Failed to get Currency")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].CdtDbtInd), "DBIT", "Failed to get CreditDebitIndicator")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].Sts.Cd), "BOOK", "Failed to get Status")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].BkTxCd.Prtry.Cd), "DBIT", "Failed to get BankTransactionCode")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].AddtlInfInd.MsgNmId), "pacs.008.001.08", "Failed to get MessageNameId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].NtryDtls[0].TxDtls[0].Refs.InstrId), "20250331231981435InstructionId00001", "Failed to get InstructionId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].NtryDtls[0].TxDtls[0].Refs.UETR), "8a562c67-ca16-48ba-b074-65581be6f011", "Failed to get UniqueTransactionReference")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get PaymentSysCode")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId), "231981435", "Failed to get PaymentSysMemberId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get PaymentSysCode")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId), "011104238", "Failed to get PaymentSysMemberId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].NtryDtls[0].TxDtls[0].LclInstrm.Prtry), "CTRC", "Failed to get LocalInstrumentChoice")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Tp), "BPRD", "Failed to get RelatedDatesProprietary")
		require.NotNil(t, Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[1].NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Dt.DtTm, "Failed to get RelatedDateTime")
		require.Equal(t, float64(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].Amt.Value), 1197.00, "Failed to get Amount")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].Amt.Ccy), "USD", "Failed to get Currency")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].CdtDbtInd), "DBIT", "Failed to get CreditDebitIndicator")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].Sts.Cd), "BOOK", "Failed to get Status")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].BkTxCd.Prtry.Cd), "DBIT", "Failed to get BankTransactionCode")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].AddtlInfInd.MsgNmId), "pacs.008.001.08", "Failed to get MessageNameId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].NtryDtls[0].TxDtls[0].Refs.InstrId), "20250331231981435InstructionId00001", "Failed to get InstructionId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].NtryDtls[0].TxDtls[0].Refs.UETR), "8a562c67-ca16-48ba-b074-65581be6f011", "Failed to get UniqueTransactionReference")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get PaymentSysCode")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId), "231981435", "Failed to get PaymentSysMemberId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get PaymentSysCode")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId), "011104238", "Failed to get PaymentSysMemberId")
		require.Equal(t, string(*Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].NtryDtls[0].TxDtls[0].LclInstrm.Prtry), "CTRC", "Failed to get LocalInstrumentChoice")
		require.Equal(t, string(Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Tp), "BPRD", "Failed to get RelatedDatesProprietary")
		require.NotNil(t, Doc08.BkToCstmrAcctRpt.Rpt[0].Ntry[2].NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Dt.DtTm, "Failed to get RelatedDateTime")
	}
}
func ActivityReportDataModel() MessageModel {
	var mesage = MessageModel{}
	mesage.MessageId = models.ActivityReport
	mesage.CreatedDateTime = time.Now()
	mesage.Pagenation = models.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	mesage.ReportId = models.EveryDay
	mesage.ReportCreateDateTime = time.Now()
	mesage.AccountOtherId = "011104238"
	mesage.TotalEntries = "61"
	mesage.TotalCreditEntries = models.NumberAndSumOfTransactions{
		NumberOfEntries: "29",
		Sum:             8775299.29,
	}
	mesage.TotalDebitEntries = models.NumberAndSumOfTransactions{
		NumberOfEntries: "27",
		Sum:             9932294.43,
	}
	mesage.TotalEntriesPerBankTransactionCode = []models.TotalsPerBankTransactionCode{
		{
			NumberOfEntries:     "0",
			BankTransactionCode: models.Sent,
		},
		{
			NumberOfEntries:     "5",
			BankTransactionCode: models.TransReceived,
		},
	}
	mesage.EntryDetails = []models.Entry{
		{
			Amount: models.CurrencyAndAmount{
				Amount:   240.67,
				Currency: "USD",
			},
			CreditDebitIndicator: models.Debit,
			Status:               models.Book,
			BankTransactionCode:  models.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: models.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000001",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310QMGFNP6000000103100900FT02",
				InstructingAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice:   models.InstrumentCTRC,
				RelatedDatesProprietary: models.BusinessProcessingDate,
				RelatedDateTime:         time.Now(),
			},
		},

		{
			Amount: models.CurrencyAndAmount{
				Amount:   1000.00,
				Currency: "USD",
			},
			CreditDebitIndicator: models.Debit,
			Status:               models.Book,
			BankTransactionCode:  models.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: models.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000002",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310QMGFNP6000000203100900FT02",
				InstructingAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice:   models.InstrumentCTRC,
				RelatedDatesProprietary: models.BusinessProcessingDate,
				RelatedDateTime:         time.Now(),
			},
		},

		{
			Amount: models.CurrencyAndAmount{
				Amount:   1197.00,
				Currency: "USD",
			},
			CreditDebitIndicator: models.Debit,
			Status:               models.Book,
			BankTransactionCode:  models.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: models.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000003",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310QMGFNP6000000303100900FT02",
				InstructingAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice:   models.InstrumentCTRC,
				RelatedDatesProprietary: models.BusinessProcessingDate,
				RelatedDateTime:         time.Now(),
			},
		},
	}

	return mesage
}
