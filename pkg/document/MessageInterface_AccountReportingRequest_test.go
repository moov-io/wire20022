package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
	"github.com/stretchr/testify/require"

	camt060 "github.com/moov-io/fedwire20022/gen/AccountReportingRequest_camt_060_001_05"
)
var sampleXML = "../models/AccountReportingRequest/swiftSample/AccountBalanceReport_Scenario1_Step1_camt.060_ABAR_MM"
func TestParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &AccountReportingRequest.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*AccountReportingRequest.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250311231981435ABARMMrequest1")
		require.Equal(t, msgModel.RequestedMsgNameId, "camt.052.001.08")
	}
}

func TestGenerateXML(t *testing.T) {
	xmlData, err := GenerateXML(&AccountReportingRequestDataModel, &AccountReportingRequest.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountReportingRequest_test.xml", xmlData)
	require.NoError(t, err)
}

func TestRequireFieldCheck(t *testing.T) {
	AccountReportingRequestDataModel.MessageId = ""
	AccountReportingRequestDataModel.RequestedMsgNameId = ""
	valid, err := RequireFieldCheck(&AccountReportingRequestDataModel, &AccountReportingRequest.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId, RequestedMsgNameId")
}

func TestXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &AccountReportingRequest.Message{})
	require.Equal(t, valid, true)
	require.NoError(t, err)
}

func TestAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&AccountReportingRequest.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*AccountReportingRequest.MessageHelper); ok {
		require.Equal(t, helper.AccountOtherId.Title, "Account Identification")
		require.Equal(t, helper.AccountOtherId.Type, "RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}")
		require.Equal(t, helper.AccountOtherId.Documentation, "Identification assigned by an institution.")
	}
}

func TestAccessToExtraField(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")
	var Message, error = CreateMessageFrom(xmlData, &AccountReportingRequest.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if doc, ok := Message.GetDocument().(*camt060.Document); ok {
		doc.AcctRptgReq.GrpHdr.MsgId = "1234567890"
	}
}

var AccountReportingRequestDataModel = AccountReportingRequest.MessageModel{
	MessageId:          "20250311231981435ABARMMrequest1",
	CreatedDateTime:    time.Now(),
	ReportRequestId:    model.AccountBalanceReport,
	RequestedMsgNameId: "camt.052.001.08",
	AccountOtherId:     "231981435",
	AccountProperty:    AccountReportingRequest.AccountTypeMerchant,
	AccountOwnerAgent: AccountReportingRequest.Camt060Agent{
		Agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
	},
}
