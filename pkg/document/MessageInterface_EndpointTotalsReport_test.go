package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/EndpointTotalsReport"
	"github.com/stretchr/testify/require"
)

func TestEndpointTotalsReportParseXMLFile(t *testing.T) {
	xmlFile := "../models/EndpointTotalsReport/generated/EndpointTotalsReport_Scenario1_Step2_camt.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &EndpointTotalsReport.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*EndpointTotalsReport.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, model.EndpointTotalsReport)
	}
}

func TestEndpointTotalsReportGenerateXML(t *testing.T) {
	dataModel := EndpointTotalsReportDataModel()
	xmlData, err := GenerateXML(&dataModel, &EndpointTotalsReport.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointTotalsReport_test.xml", xmlData)
	require.NoError(t, err)
}

func TestEndpointTotalsReportRequireFieldCheck(t *testing.T) {
	dataModel := EndpointTotalsReportDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &EndpointTotalsReport.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestEndpointTotalsReportXMLValidation(t *testing.T) {
	xmlFile := "../models/EndpointTotalsReport/swiftSample/EndpointTotalsReport_Scenario1_Step2_camt.052_ETOT"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &EndpointTotalsReport.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestEndpointTotalsReportAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&EndpointTotalsReport.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*EndpointTotalsReport.MessageHelper); ok {
		require.Equal(t, helper.AccountOtherId.Title, "Account Other Id")
		require.Equal(t, helper.AccountOtherId.Type, "Max34Text (based on string) minLength: 1 maxLength: 34")
		require.Equal(t, helper.AccountOtherId.Documentation, "Unique identification of an account, as assigned by the account servicer, using an identification scheme.")
	}
}

func EndpointTotalsReportDataModel() EndpointTotalsReport.MessageModel {
	var message, _ = EndpointTotalsReport.NewMessage("")
	message.Data.MessageId = model.EndpointTotalsReport
	message.Data.CreatedDateTime = time.Now()
	message.Data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.Data.ReportId = model.Intraday
	message.Data.ReportCreateDateTime = time.Now()
	message.Data.AccountOtherId = "B1QDRCQR"
	message.Data.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "1268",
		Sum:             18423923492.15,
	}
	message.Data.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "4433",
		Sum:             12378489145.96,
	}
	message.Data.TotalEntriesPerTransactionCode = []model.NumberAndStatusOfTransactions{
		{
			NumberOfEntries: "1",
			Status:          model.Rejected,
		},
		{
			NumberOfEntries: "0",
			Status:          model.MessagesIntercepted,
		},
		{
			NumberOfEntries: "0",
			Status:          model.MessagesInProcess,
		},
		{
			NumberOfEntries: "27",
			Status:          model.TransReceived,
		},
		{
			NumberOfEntries: "193",
			Status:          model.Sent,
		},
	}
	message.Data.AdditionalReportInfo = "Next IMAD sequence number: 4627. Next OMAD sequence number: 1296. Count of missing IMAD sequence numbers: 0."

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
