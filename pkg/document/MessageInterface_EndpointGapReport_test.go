package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/EndpointGapReport"
	"github.com/stretchr/testify/require"
)

var EndpointGapReportxmlFile = "../models/EndpointGapReport/swiftSample/EndpointGapReport_Scenario1_Step1_camt.052_IMAD"

func TestEndpointGapReportParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(EndpointGapReportxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &EndpointGapReport.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*EndpointGapReport.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, model.CAMTReportType("GAPR"))
	}
}

func TestEndpointGapReportGenerateXML(t *testing.T) {
	dataModel := EndpointGapReportDataModel()
	xmlData, err := GenerateXML(&dataModel, &EndpointGapReport.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointGapReport_test.xml", xmlData)
	require.NoError(t, err)
}

func TestEndpointGapReportRequireFieldCheck(t *testing.T) {
	dataModel := EndpointGapReportDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &EndpointGapReport.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestEndpointGapReportXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(EndpointGapReportxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &EndpointGapReport.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestEndpointGapReportAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&EndpointGapReport.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*EndpointGapReport.MessageHelper); ok {
		require.Equal(t, helper.AccountOtherId.Title, "Account Other Id")
		require.Equal(t, helper.AccountOtherId.Type, "Max34Text (based on string) minLength: 1 maxLength: 34")
		require.Equal(t, helper.AccountOtherId.Documentation, "Unique identification of an account, as assigned by the account servicer, using an identification scheme.")
	}
}

func EndpointGapReportDataModel() EndpointGapReport.MessageModel {
	var message, _ = EndpointGapReport.NewMessage("")
	message.Data.MessageId = model.EndpointGapReportType
	message.Data.CreatedDateTime = time.Now()
	message.Data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.Data.ReportId = EndpointGapReport.InputMessageAccountabilityData
	message.Data.ReportCreateDateTime = time.Now()
	message.Data.AccountOtherId = "ISOTEST1"
	message.Data.AdditionalReportInfo = "Next sequence number: 00431. List of missing sequence numbers: 000052 000054 000056 000058 000059 000061 000062 000064-000068 000070 000071 000073 000074 000076 000077 000079 000080 000082 000083 000085 000086 000088 000089 000091 000092 000094 000136 000139 000141 000142 000144 000145 000147 000148 000150 000151 000153 000154 000156 000157 000159 000160 000306 000308 000309 000311 000312 000370 000371 000373 000374 000376 000380 000382 000384 000386 000389 000391 000407 000408 000410 000413"

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
