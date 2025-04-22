package EndpointGapReport

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestEndpointGapReport_Scenario1_Step1_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointGapReport_Scenario1_Step1_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointGapReport_Scenario1_Step1_camt.052_IMAD")
	genterated := filepath.Join("generated", "EndpointGapReport_Scenario1_Step1_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}

func TestEndpointGapReport_Scenario1_Step1_camt_OMAD_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointGapReport_Scenario1_Step1_camt_OMAD.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointGapReport_Scenario1_Step1_camt.052_OMAD")
	genterated := filepath.Join("generated", "EndpointGapReport_Scenario1_Step1_camt_OMAD.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
