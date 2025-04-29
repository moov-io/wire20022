package EndpointTotalsReport

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestRequireField(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	cErr := message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreatedDateTime, MessagePagination, ReportId, ReportCreateDateTime, AccountOtherId, TotalCreditEntries, TotalDebitEntries")
}
func generateRequreFields(msg Message) Message {
	if msg.data.MessageId == "" {
		msg.data.MessageId = model.EndpointTotalsReport
	}
	if msg.data.CreatedDateTime.IsZero() {
		msg.data.CreatedDateTime = time.Now()
	}
	if isEmpty(msg.data.MessagePagination) {
		msg.data.MessagePagination = model.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		}
	}
	if msg.data.ReportId == "" {
		msg.data.ReportId = model.Intraday
	}
	if msg.data.ReportCreateDateTime.IsZero() {
		msg.data.ReportCreateDateTime = time.Now()
	}
	if msg.data.AccountOtherId == "" {
		msg.data.AccountOtherId = "B1QDRCQR"
	}
	if isEmpty(msg.data.TotalCreditEntries) {
		msg.data.TotalCreditEntries = model.NumberAndSumOfTransactions{
			NumberOfEntries: "1268",
			Sum:             18423923492.15,
		}
	}
	if isEmpty(msg.data.TotalDebitEntries) {
		msg.data.TotalDebitEntries = model.NumberAndSumOfTransactions{
			NumberOfEntries: "4433",
			Sum:             12378489145.96,
		}
	}
	return msg
}
func TestEndpointTotalsReportFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "EndpointTotalsReport_Scenario1_Step2_camt.052_ETOT")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.MsgId), "ETOT")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb), "1")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.Id), "IDAY")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr.Id), "B1QDRCQR")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries.NbOfNtries), "1268")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd[0].BkTxCd.Prtry.Cd), "RJCT")
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

func TestEndpointTotalsReportReportValidator(t *testing.T) {
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
			"AccountOtherId",
			Message{data: MessageModel{AccountOtherId: INVALID_OTHER_ID}},
			"error occur at AccountOtherId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [A-Z0-9]{8,8}",
		},
		{
			"TotalCreditEntries",
			Message{data: MessageModel{TotalCreditEntries: model.NumberAndSumOfTransactions{
				NumberOfEntries: "aaaaa",
				Sum:             18423923492.15,
			}}},
			"error occur at TotalCreditEntries.NumberOfEntries: aaaaa fails validation with pattern [0-9]{1,15}",
		},
		{
			"TotalCreditEntries",
			Message{data: MessageModel{TotalDebitEntries: model.NumberAndSumOfTransactions{
				NumberOfEntries: "aaaaa",
				Sum:             18423923492.15,
			}}},
			"error occur at TotalDebitEntries.NumberOfEntries: aaaaa fails validation with pattern [0-9]{1,15}",
		},
		{
			"TotalEntriesPerTransactionCode",
			Message{data: MessageModel{TotalEntriesPerTransactionCode: []model.NumberAndStatusOfTransactions{
				{
					NumberOfEntries: "0",
					Status:          model.Rejected,
				},
				{
					NumberOfEntries: "0",
					Status:          model.TransactionStatusCode(INVALID_COUNT),
				},
				{
					NumberOfEntries: "0",
					Status:          model.MessagesIntercepted,
				},
				{
					NumberOfEntries: "0",
					Status:          model.Sent,
				},
			}}},
			"error occur at TotalEntriesPerTransactionCode.Status: UNKNOWN fails enumeration validation",
		},
		{
			"TotalEntriesPerTransactionCode",
			Message{data: MessageModel{TotalEntriesPerTransactionCode: []model.NumberAndStatusOfTransactions{
				{
					NumberOfEntries: "bbbbb",
					Status:          model.Rejected,
				},
				{
					NumberOfEntries: "0",
					Status:          model.MessagesInProcess,
				},
				{
					NumberOfEntries: "0",
					Status:          model.MessagesIntercepted,
				},
				{
					NumberOfEntries: "0",
					Status:          model.Sent,
				},
			}}},
			"error occur at TotalEntriesPerTransactionCode.NumberOfEntries: bbbbb fails validation with pattern [0-9]{1,15}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			nMsg := generateRequreFields(tt.msg)
			msgErr := nMsg.CreateDocument()
			if msgErr != nil {
				require.Equal(t, tt.expectedErr, msgErr.Error())
			}
		})
	}
}
func TestEndpointTotalsReport_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
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

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointTotalsReport_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointTotalsReport_Scenario1_Step2_camt.052_ETOT")
	genterated := filepath.Join("generated", "EndpointTotalsReport_Scenario1_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}

func TestEndpointTotalsReport_Scenario2_Step1_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
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

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointTotalsReport_Scenario2_Step1_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointTotalsReport_Scenario2_Step1_camt.052_ETOT")
	genterated := filepath.Join("generated", "EndpointTotalsReport_Scenario2_Step1_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
