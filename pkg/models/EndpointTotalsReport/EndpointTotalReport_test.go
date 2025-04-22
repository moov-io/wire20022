package EndpointTotalsReport

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestEndpointTotalsReport_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = model.EndpointTotalsReport
	message.data.CreatedDateTime = time.Now()
	message.data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.data.ReportId = model.Intraday
	message.data.ReportCreateDateTime = time.Now()
	message.data.AccountOtherId = "B1QDRCQR"
	message.data.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "1268",
		Sum:             18423923492.15,
	}
	message.data.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "4433",
		Sum:             12378489145.96,
	}
	message.data.TotalEntriesPerTransactionCode = []model.NumberAndStatusOfTransactions{
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
	message.data.AdditionalReportInfo = "Next IMAD sequence number: 4627. Next OMAD sequence number: 1296. Count of missing IMAD sequence numbers: 0."

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("EndpointTotalsReport_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointTotalsReport_Scenario1_Step2_camt.052_ETOT")
	genterated := filepath.Join("generated", "EndpointTotalsReport_Scenario1_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}

func TestEndpointTotalsReport_Scenario2_Step1_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = model.EndpointTotalsReport
	message.data.CreatedDateTime = time.Now()
	message.data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.data.ReportId = model.EveryDay
	message.data.ReportCreateDateTime = time.Now()
	message.data.AccountOtherId = "B1QDRCQR"
	message.data.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "5915",
		Sum:             33992880250.31,
	}
	message.data.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "7070",
		Sum:             35073483328.29,
	}
	message.data.TotalEntriesPerTransactionCode = []model.NumberAndStatusOfTransactions{
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
			NumberOfEntries: "924",
			Status:          model.TransReceived,
		},
		{
			NumberOfEntries: "723",
			Status:          model.Sent,
		},
	}
	message.data.AdditionalReportInfo = "Next IMAD sequence number: 7794. Next OMAD sequence number: 6840. Count of missing IMAD sequence numbers: 0."

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("EndpointTotalsReport_Scenario2_Step1_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointTotalsReport_Scenario2_Step1_camt.052_ETOT")
	genterated := filepath.Join("generated", "EndpointTotalsReport_Scenario2_Step1_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
