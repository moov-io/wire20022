package AccountReportingRequest

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_05"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

var sample1XML = filepath.Join("swiftSample", "EndpointDetailsReport_Scenario1_Step1_camt.060_DTLS")
var sample2XML = filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_camt.060_ABAR_MM")

func TestDocumentElementToModelOne(t *testing.T) {
	var xmlData, err = models.ReadXMLFile(sample1XML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := ParseXML(xmlData)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435DTLSrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, models.CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, models.PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")
}
func TestDocumentElementToModelTwo(t *testing.T) {
	var xmlData, err = models.ReadXMLFile(sample2XML)
	require.NoError(t, err, "Failed to read XML file")
	model, err := ParseXML(xmlData)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, models.CAMTReportType("ABAR"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, models.AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, models.PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
}

var AccountReportingRequestDataModel_1 = MessageModel{
	MessageHeader: base.MessageHeader{
		MessageId:       "20250311231981435ABARMMrequest1",
		CreatedDateTime: time.Now(),
	},
	ReportRequestId:    models.AccountBalanceReport,
	RequestedMsgNameId: "camt.052.001.08",
	AccountOtherId:     "231981435",
	AccountProperty:    models.AccountTypeMerchant,
	AccountOwnerAgent: models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	},
}

func TestModelToDocument05_One(t *testing.T) {
	var doc05, err = DocumentWith(AccountReportingRequestDataModel_1, CAMT_060_001_05)
	require.NoError(t, err, "Failed to create document")
	if Doc05, ok := doc05.(*camt_060_001_05.Document); ok {
		require.Equal(t, string(Doc05.AcctRptgReq.GrpHdr.MsgId), "20250311231981435ABARMMrequest1", "Failed to get MessageId")
		require.NotNil(t, Doc05.AcctRptgReq.GrpHdr.CreDtTm, "Failed to get CreatedDateTime")
		require.Equal(t, models.CAMTReportType(*Doc05.AcctRptgReq.RptgReq[0].Id), models.CAMTReportType("ABAR"), "Failed to get MessageId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].ReqdMsgNmId), "camt.052.001.08", "Failed to get RequestedMsgNameId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id), "231981435", "Failed to get AccountOtherId")
		require.Equal(t, string(*Doc05.AcctRptgReq.RptgReq[0].Acct.Tp.Prtry), "M", "Failed to get AccountProperty")
		require.Equal(t, string(*Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get AccountOwnerAgent.PaymentSysCode")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId), "231981435", "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	}
}

var AccountReportingRequestDataModel_2 = MessageModel{
	MessageHeader: base.MessageHeader{
		MessageId:       "20250311231981435ABARMMrequest1",
		CreatedDateTime: time.Now(),
	},
	ReportRequestId:    models.EndpointDetailsSentReport,
	RequestedMsgNameId: "camt.052.001.08",
	AccountOwnerAgent: models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
		OtherTypeId:        "B1QDRCQR",
	},
	FromToSequence: models.SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000100",
	},
}

func TestModelToDocument05_Two(t *testing.T) {
	var doc05, err = DocumentWith(AccountReportingRequestDataModel_2, CAMT_060_001_05)
	require.NoError(t, err, "Failed to create document")
	if Doc05, ok := doc05.(*camt_060_001_05.Document); ok {
		require.Equal(t, string(Doc05.AcctRptgReq.GrpHdr.MsgId), "20250311231981435ABARMMrequest1", "Failed to get MessageId")
		require.NotNil(t, Doc05.AcctRptgReq.GrpHdr.CreDtTm, "Failed to get CreatedDateTime")
		require.Equal(t, models.CAMTReportType(*Doc05.AcctRptgReq.RptgReq[0].Id), models.CAMTReportType("DTLS"), "Failed to get MessageId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].ReqdMsgNmId), "camt.052.001.08", "Failed to get RequestedMsgNameId")
		require.Equal(t, string(*Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA", "Failed to get AccountOwnerAgent.PaymentSysCode")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId), "231981435", "Failed to get AccountOwnerAgent.PaymentSysMemberId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.Othr.Id), "B1QDRCQR", "Failed to get AccountOwnerAgent.OtherTypeId")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].FrSeq), "000002", "Failed to get FromToSequence.FromSeq")
		require.Equal(t, string(Doc05.AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].ToSeq), "000100", "Failed to get FromToSequence.ToSeq")
	}
}

func TestModelToDocument05_ValidateError(t *testing.T) {
	var model = MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250311231981435ABARMMrequest1",
			CreatedDateTime: time.Now(),
		},
		ReportRequestId:    models.EndpointDetailsSentReport,
		RequestedMsgNameId: "camt.052.001.08",
		AccountOwnerAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
			OtherTypeId:        "B1QDRCQR",
		},
		FromToSequence: models.SequenceRange{
			FromSeq: "000002",
			ToSeq:   "000100",
		},
	}
	model.MessageHeader.MessageId = "20250311231981435ABARMMrequest120250311231981435ABARMMrequest1"
	var buf strings.Builder
	err := model.WriteXML(&buf, CAMT_060_001_05)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "creating document: field copy AcctRptgReq.GrpHdr.MsgId failed: failed to set MessageId: 20250311231981435ABARMMrequest120250311231981435ABARMMrequest1 fails validation with length 62 <= required maxLength 35")

	model.MessageHeader.MessageId = "20250311231981435ABARMMrequest1"
	model.RequestedMsgNameId = "camt.060.001.05camt.060.001.05camt.060.001.05camt.060.001.05"
	buf.Reset()
	err = model.WriteXML(&buf, CAMT_060_001_05)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "creating document: field copy AcctRptgReq.RptgReq[0].ReqdMsgNmId failed: failed to set RequestedMsgNameId: camt.060.001.05camt.060.001.05camt.060.001.05camt.060.001.05 fails validation with length 60 <= required maxLength 35")

	model.RequestedMsgNameId = "camt.052.001.08"
	model.AccountOtherId = "231981435231981435231981435231981435231981435231981435231981435231981435231981435231981435"
	buf.Reset()
	err = model.WriteXML(&buf, CAMT_060_001_05)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "creating document: field copy AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id failed: failed to set AccountOtherId: 231981435231981435231981435231981435231981435231981435231981435231981435231981435231981435 fails validation with length 90 <= required maxLength 34")
}
func TestModelToDocument05_CheckRequireField(t *testing.T) {
	var model = MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250311231981435ABARMMrequest1",
			CreatedDateTime: time.Now(),
		},
		ReportRequestId:    models.EndpointDetailsSentReport,
		RequestedMsgNameId: "camt.052.001.08",
		AccountOwnerAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
			OtherTypeId:        "B1QDRCQR",
		},
	}
	model.MessageId = ""
	err := CheckRequiredFields(model)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	var buf strings.Builder
	err = model.WriteXML(&buf, CAMT_060_001_05)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "creating document: validation failed for field \"MessageId\": is required: required field missing")

	model.MessageHeader.MessageId = "20250311231981435ABARMMrequest1"
	model.ReportRequestId = ""
	err = CheckRequiredFields(model)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"ReportRequestId\": is required: required field missing")
	buf.Reset()
	err = model.WriteXML(&buf, CAMT_060_001_05)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "creating document: validation failed for field \"ReportRequestId\": is required: required field missing")
}
func TestModelHelper(t *testing.T) {
	require.Equal(t, BuildMessageHelper().MessageId.Title, "Message Identification", "Failed to get MessageId")
	require.Equal(t, BuildMessageHelper().CreatedDateTime.Title, "Creation Date Time", "Failed to get CreatedDateTime")
	require.Equal(t, BuildMessageHelper().ReportRequestId.Title, "Report Request Identification", "Failed to get ReportRequestId")
	require.Equal(t, BuildMessageHelper().RequestedMsgNameId.Title, "Requested Message Name Identification", "Failed to get RequestedMsgNameId")
	require.Equal(t, BuildMessageHelper().AccountOtherId.Title, "Account Identification", "Failed to get AccountOtherId")
	require.Equal(t, BuildMessageHelper().AccountProperty.Title, "Account Type Proprietary", "Failed to get AccountProperty")
	require.Equal(t, BuildMessageHelper().AccountOwnerAgent.PaymentSysCode.Title, "Clearing System Identification Code", "Failed to get AccountOwnerAgent")
}
