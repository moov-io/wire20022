package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
	"github.com/stretchr/testify/require"

	camt060 "github.com/moov-io/fedwire20022/gen/AccountReportingRequest_camt_060_001_05"
)

func TestParseXMLFile(t *testing.T) {
	xmlFile := "../models/AccountReportingRequest/generated/AccountBalanceReport_Scenario1_Step1__MS_camt.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.Nil(t, err, "Failed to read XML file")
	var Message, error = CreateMessageFrom(xmlData, &AccountReportingRequest.Message{})
	require.Nil(t, error, "Failed to make XML structure")
	if doc, ok := Message.GetDocument().(*camt060.Document); ok {
		require.Equal(t, string(doc.AcctRptgReq.RptgReq.Id), "ABAR")
		require.Equal(t, string(doc.AcctRptgReq.RptgReq.ReqdMsgNmId), "camt.052.001.08")
	}
}
func TestAccessToExtraField(t *testing.T){
	xmlFile := "../models/AccountReportingRequest/generated/AccountBalanceReport_Scenario1_Step1__MS_camt.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.Nil(t, err, "Failed to read XML file")
	var Message, error = CreateMessageFrom(xmlData, &AccountReportingRequest.Message{})
	require.Nil(t, error, "Failed to make XML structure")
	if doc, ok := Message.GetDocument().(*camt060.Document); ok {
		doc.AcctRptgReq.GrpHdr.MsgId = "1234567890"
	}
}

var AccountReportingRequestDataModel = AccountReportingRequest.MessageModel {
	MessageId: "20250311231981435ABARMMrequest1",
	CreatedDateTime: time.Now(),
	ReportRequestId: model.AccountBalanceReport,
	RequestedMsgNameId: "camt.052.001.08",
	AccountOtherId: "231981435",
	AccountProperty: AccountReportingRequest.AccountTypeMerchant,
	AccountOwnerAgent: AccountReportingRequest.Camt060Agent{
		Agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
	},
}
func TestMessageModelToXML(t *testing.T){
	message, cErr := CreateMessageWith(&AccountReportingRequestDataModel, &AccountReportingRequest.Message{})
	require.Nil(t, cErr)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountReportingRequest_test.xml", xmlData)
	require.NoError(t, err)
}
func TestAccessToDataModel(t *testing.T) {
	message, cErr := CreateMessageWith(&AccountReportingRequestDataModel, &AccountReportingRequest.Message{})
	require.Nil(t, cErr)
	if dataModel, ok := message.GetDataModel().(*AccountReportingRequest.MessageModel); ok {
		require.Equal(t, dataModel.MessageId, "20250311231981435ABARMMrequest1")
	}
}
func TestAccessToHelper(t *testing.T){
	message, cErr := CreateMessage(&AccountReportingRequest.Message{})
	require.Nil(t, cErr)
	if helper, ok := message.GetHelper().(*AccountReportingRequest.MessageHelper); ok {
		require.Equal(t, helper.AccountOtherId.Title, "Account Identification")
		require.Equal(t, helper.AccountOtherId.Type, "RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}")
		require.Equal(t, helper.AccountOtherId.Documentation, "Identification assigned by an institution.")
	}
}
func TestUpdateMessageModel(t *testing.T) {
	message, cErr := CreateMessageWith(&AccountReportingRequestDataModel, &AccountReportingRequest.Message{})
	require.Nil(t, cErr)
	if dataModel, ok := message.GetDataModel().(*AccountReportingRequest.MessageModel); ok {
		require.Equal(t, dataModel.MessageId, "20250311231981435ABARMMrequest1")
		dataModel.MessageId = "20250311231Updated"

		vErr := message.CreateDocument()
		require.Nil(t, vErr)
		if doc, ok := message.GetDocument().(*camt060.Document); ok {
			require.Equal(t, string(doc.AcctRptgReq.GrpHdr.MsgId), "20250311231Updated")
		}
	}
}