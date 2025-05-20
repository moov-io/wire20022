package ArchiveAccountReportingRequest

import (
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/archives/AccountReportingRequest/camt_060_001_05"
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

var AccountReportingRequestDataModel_1 = MessageModel{
	MessageId:          "20250311231981435ABARMMrequest1",
	CreatedDateTime:    time.Now(),
	ReportRequestId:    AccountBalanceReport,
	RequestedMsgNameId: "camt.052.001.08",
	AccountOtherId:     "231981435",
	AccountProperty:    AccountTypeMerchant,
	AccountOwnerAgent: Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	},
}

func TestModelToDocument05_One(t *testing.T) {
	var doc05, err = DocumentWith(AccountReportingRequestDataModel_1, "camt.060.001.05")
	require.NoError(t, err, "Failed to create document")
	if Doc05, ok := doc05.(*camt_060_001_05.Document); ok {
		require.Equal(t, string(Doc05.AcctRptgReq.GrpHdr.MsgId), "20250311231981435ABARMMrequest1", "Failed to get MessageId")
		require.NotNil(t, Doc05.AcctRptgReq.GrpHdr.CreDtTm, "Failed to get CreatedDateTime")
		require.Equal(t, CAMTReportType(*Doc05.AcctRptgReq.RptgReq[0].Id), CAMTReportType("ABAR"), "Failed to get MessageId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].ReqdMsgNmId), "camt.052.001.08", "Failed to get RequestedMsgNameId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id), "231981435", "Failed to get AccountOtherId")
		require.Equal(t, string(*Doc05.AcctRptgReq.RptgReq[0].Acct.Tp.Prtry), "M", "Failed to get AccountProperty")
		require.Equal(t, string(*Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get AccountOwnerAgent.PaymentSysCode")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId), "231981435", "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	}
}

var AccountReportingRequestDataModel_2 = MessageModel{
	MessageId:          "20250311231981435ABARMMrequest1",
	CreatedDateTime:    time.Now(),
	ReportRequestId:    EndpointDetailsSentReport,
	RequestedMsgNameId: "camt.052.001.08",
	AccountOwnerAgent: Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
		OtherTypeId:        "B1QDRCQR",
	},
	FromToSequence: SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000100",
	},
}

func TestModelToDocument05_Two(t *testing.T) {
	var doc05, err = DocumentWith(AccountReportingRequestDataModel_2, "camt.060.001.05")
	require.NoError(t, err, "Failed to create document")
	if Doc05, ok := doc05.(*camt_060_001_05.Document); ok {
		require.Equal(t, string(Doc05.AcctRptgReq.GrpHdr.MsgId), "20250311231981435ABARMMrequest1", "Failed to get MessageId")
		require.NotNil(t, Doc05.AcctRptgReq.GrpHdr.CreDtTm, "Failed to get CreatedDateTime")
		require.Equal(t, CAMTReportType(*Doc05.AcctRptgReq.RptgReq[0].Id), CAMTReportType("DTLS"), "Failed to get MessageId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].ReqdMsgNmId), "camt.052.001.08", "Failed to get RequestedMsgNameId")
		require.Equal(t, string(*Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get AccountOwnerAgent.PaymentSysCode")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId), "231981435", "Failed to get AccountOwnerAgent.PaymentSysMemberId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.Othr.Id), "B1QDRCQR", "Failed to get AccountOwnerAgent.OtherTypeId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].FrSeq), "000002", "Failed to get FromToSequence.FromSeq")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].ToSeq), "000100", "Failed to get FromToSequence.ToSeq")
	}
}

var AccountReportingRequestDataModel_Empty = MessageModel{}

func TestModelToDocument05_Empty(t *testing.T) {
	var doc05, err = DocumentWith(AccountReportingRequestDataModel_Empty, "camt.060.001.05")
	require.NoError(t, err, "Failed to create document")
	if Doc05, ok := doc05.(*camt_060_001_05.Document); ok {
		require.Equal(t, string(Doc05.AcctRptgReq.GrpHdr.MsgId), "", "Failed to get MessageId")
	}
}
