package EndpointGapReport

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestEndpointGapReportFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "EndpointGapReport_Scenario1_Step1_camt.052_IMAD")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.MsgId), "GAPR")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb), "1")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt[0].Id), "IMAD")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id), "B1QDRCQR")
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"
const INVALID_TRCOUNT string = "123456789012345"
const INVALID_MESSAGE_ID string = "12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456"
const INVALID_OTHER_ID string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const INVALID_BUILD_NUM string = "12345678901234567"
const INVALID_POSTAL_CODE string = "12345678901234567"
const INVALID_COUNTRY_CODE string = "12345678"
const INVALID_MESSAGE_NAME_ID string = "sabcd-123-001-12"
const INVALID_PAY_SYSCODE model.PaymentSystemType = model.PaymentSystemType(INVALID_COUNT)

func TestEndpointDetailsReportValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"MessageId",
			Message{data: MessageModel{MessageId: model.CAMTReportType(INVALID_COUNT)}},
			"error occur at MessageId: UNKNOWN fails enumeration validation",
		},
		{
			"ReportId",
			Message{data: MessageModel{ReportId: GapType(INVALID_COUNT)}},
			"error occur at ReportId: UNKNOWN fails enumeration validation",
		},
		{
			"AccountOtherId",
			Message{data: MessageModel{AccountOtherId: INVALID_OTHER_ID}},
			"error occur at AccountOtherId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [A-Z0-9]{8,8}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			msgErr := tt.msg.CreateDocument()
			if msgErr != nil {
				require.Equal(t, tt.expectedErr, msgErr.Error())
			}
		})
	}
}
func TestEndpointGapReport_Scenario1_Step1_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = model.EndpointGapReportType
	message.data.CreatedDateTime = time.Now()
	message.data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.data.ReportId = InputMessageAccountabilityData
	message.data.ReportCreateDateTime = time.Now()
	message.data.AccountOtherId = "B1QDRCQR"
	message.data.AdditionalReportInfo = "Next sequence number: 011062. List of missing sequence numbers: 000463 000485 000497 000503-000508 000532 000660 000806 000842 000845 000853 000885 001031 001045 001184 001220 001260 001559 001571 001749 005365 005375 005436 005450 005531 005539 005547 005659 006144 006569 006647 006869 006934 007103 007105 007127 007208 007347 007446 007554 007661 007663 007918 008660 008943 009016 009207 009282 010536 010848 011035 011036 011037 011038 011039 011040 011041 011042 011043 011044 011045"

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointGapReport_Scenario1_Step1_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointGapReport_Scenario1_Step1_camt.052_IMAD")
	genterated := filepath.Join("generated", "EndpointGapReport_Scenario1_Step1_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}

func TestEndpointGapReport_Scenario1_Step1_camt_OMAD_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = model.EndpointGapReportType
	message.data.CreatedDateTime = time.Now()
	message.data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.data.ReportId = InputMessageAccountabilityData
	message.data.ReportCreateDateTime = time.Now()
	message.data.AccountOtherId = "ISOTEST1"
	message.data.AdditionalReportInfo = "Next sequence number: 00431. List of missing sequence numbers: 000052 000054 000056 000058 000059 000061 000062 000064-000068 000070 000071 000073 000074 000076 000077 000079 000080 000082 000083 000085 000086 000088 000089 000091 000092 000094 000136 000139 000141 000142 000144 000145 000147 000148 000150 000151 000153 000154 000156 000157 000159 000160 000306 000308 000309 000311 000312 000370 000371 000373 000374 000376 000380 000382 000384 000386 000389 000391 000407 000408 000410 000413"

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointGapReport_Scenario1_Step1_camt_OMAD.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointGapReport_Scenario1_Step1_camt.052_OMAD")
	genterated := filepath.Join("generated", "EndpointGapReport_Scenario1_Step1_camt_OMAD.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
