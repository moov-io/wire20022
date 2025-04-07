package EndpointTotalsReport_camt_052_001_08

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestEndpointTotalsReport_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message = NewCamtTotal0522Message()
	message.model.MessageId = model.EndpointTotalsReport
	message.model.CreatedDateTime = time.Now()
	message.model.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.model.ReportId = model.Intraday
	message.model.ReportCreateDateTime = time.Now()
	message.model.AccountOtherId = "B1QDRCQR"
	message.model.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "1268",
		Sum:             18423923492.15,
	}
	message.model.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "4433",
		Sum:             12378489145.96,
	}
	message.model.TotalEntriesPerTransactionCode = []model.NumberAndStatusOfTransactions{
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
	message.model.AdditionalReportInfo = "Next IMAD sequence number: 4627. Next OMAD sequence number: 1296. Count of missing IMAD sequence numbers: 0."

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "EndpointTotalsReport_Scenario1_Step2_camt.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}

func TestEndpointTotalsReport_Scenario2_Step1_camt_CreateXML(t *testing.T) {
	var message = NewCamtTotal0522Message()
	message.model.MessageId = model.EndpointTotalsReport
	message.model.CreatedDateTime = time.Now()
	message.model.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.model.ReportId = model.Intraday
	message.model.ReportCreateDateTime = time.Now()
	message.model.AccountOtherId = "B1QDRCQR"
	message.model.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "5915",
		Sum:             33992880250.31,
	}
	message.model.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "7070",
		Sum:             35073483328.29,
	}
	message.model.TotalEntriesPerTransactionCode = []model.NumberAndStatusOfTransactions{
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
	message.model.AdditionalReportInfo = "Next IMAD sequence number: 7794. Next OMAD sequence number: 6840. Count of missing IMAD sequence numbers: 0."

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "EndpointTotalsReport_Scenario2_Step1_camt.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
