package EndpointDetailsReport

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreationDateTime, MessagePagination, BussinessQueryMsgId, BussinessQueryMsgNameId, BussinessQueryCreateDatetime, ReportId, ReportingSequence, AccountOtherId")
}
func generateRequreFields(msg Message) Message {
	if msg.data.MessageId == "" {
		msg.data.MessageId = "DTLS"
	}
	if msg.data.CreationDateTime.IsZero() {
		msg.data.CreationDateTime = time.Now()
	}
	if isEmpty(msg.data.MessagePagination) {
		msg.data.MessagePagination = model.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		}
	}
	if msg.data.BussinessQueryMsgId == "" {
		msg.data.BussinessQueryMsgId = "20250311231981435DTLSrequest1"
	}
	if msg.data.BussinessQueryMsgNameId == "" {
		msg.data.BussinessQueryMsgNameId = "camt.060.001.05"
	}
	if msg.data.BussinessQueryCreateDatetime.IsZero() {
		msg.data.BussinessQueryCreateDatetime = time.Now()
	}
	if msg.data.ReportId == "" {
		msg.data.ReportId = model.Intraday
	}
	if isEmpty(msg.data.ReportingSequence) {
		msg.data.ReportingSequence = model.SequenceRange{
			FromSeq: "000001",
			ToSeq:   "000100",
		}
	}
	if msg.data.AccountOtherId == "" {
		msg.data.AccountOtherId = "B1QDRCQR"
	}
	return msg
}
func TestEndpointDetailsReportFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario1_Step2_camt.052_DTLS")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.MsgId), "DTLS")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb), "1")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId), "20250311231981435DTLSrequest1")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId), "camt.060.001.05")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.Id), "IDAY")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr.Id), "B1QDRCQR")
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
			Message{data: MessageModel{MessageId: INVALID_MESSAGE_ID}},
			"error occur at MessageId: 12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456 fails validation with length 66 <= required maxLength 35",
		},
		{
			"BussinessQueryMsgId",
			Message{data: MessageModel{BussinessQueryMsgId: INVALID_MESSAGE_ID}},
			"error occur at BussinessQueryMsgId: 12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456 fails validation with length 66 <= required maxLength 35",
		},
		{
			"BussinessQueryMsgNameId",
			Message{data: MessageModel{BussinessQueryMsgNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at BussinessQueryMsgNameId: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"ReportId",
			Message{data: MessageModel{ReportId: model.ReportType(INVALID_COUNT)}},
			"error occur at ReportId: UNKNOWN fails enumeration validation",
		},
		{
			"AccountOtherId",
			Message{data: MessageModel{AccountOtherId: INVALID_OTHER_ID}},
			"error occur at AccountOtherId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [A-Z0-9]{8,8}",
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

//	func TestFloat(t *testing.T){
//		err := camt060.XSequenceNumberFedwireFunds1(float64(000001)).Validate()
//		require.Nil(t, err)
//	}
func TestEndpointDetailsReport_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.Nil(t, vErr)
	message.data.MessageId = "DTLS"
	message.data.CreationDateTime = time.Now()
	message.data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.data.BussinessQueryMsgId = "20250311231981435DTLSrequest1"
	message.data.BussinessQueryMsgNameId = "camt.060.001.05"
	message.data.BussinessQueryCreateDatetime = time.Now()
	message.data.ReportId = model.Intraday
	message.data.ReportingSequence = model.SequenceRange{
		FromSeq: "000001",
		ToSeq:   "000100",
	}
	message.data.ReportCreateDateTime = time.Now()
	message.data.AccountOtherId = "B1QDRCQR"
	message.data.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "100",
		Sum:             8307111.56,
	}
	message.data.TotalEntriesPerTransactionCode = []model.NumberAndStatusOfTransactions{
		{
			NumberOfEntries: "0",
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
	}
	message.data.EntryDetails = []model.Entry{
		{
			Amount: model.CurrencyAndAmount{
				Amount:   50000.00,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Debit,
			Status:               model.Book,
			BankTransactionCode:  model.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: model.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000001",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310ISOTEST100000103100900FT02",
				InstructingAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice: model.InstrumentCTRC,
			},
		},
		{
			Amount: model.CurrencyAndAmount{
				Amount:   8000.00,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Debit,
			Status:               model.Book,
			BankTransactionCode:  model.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: model.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000002",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310ISOTEST100000203100900FT02",
				InstructingAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice: model.InstrumentCTRC,
			},
		},
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario1_Step2_camt.052_DTLS")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointDetailsReport_Scenario2_Step2_camt_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.Nil(t, vErr)
	message.data.MessageId = "DTLR"
	message.data.CreationDateTime = time.Now()
	message.data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.data.BussinessQueryMsgId = "20250311231981435DTLRrequest1"
	message.data.BussinessQueryMsgNameId = "camt.060.001.05"
	message.data.BussinessQueryCreateDatetime = time.Now()
	message.data.ReportId = model.Intraday
	message.data.ReportingSequence = model.SequenceRange{
		FromSeq: "000001",
		ToSeq:   "000100",
	}
	message.data.ReportCreateDateTime = time.Now()
	message.data.AccountOtherId = "B1QDRCQR"
	message.data.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "94",
		Sum:             2871734.98,
	}
	message.data.TotalEntriesPerTransactionCode = []model.NumberAndStatusOfTransactions{
		{
			NumberOfEntries: "6",
			Status:          model.TransReceived,
		},
	}
	message.data.EntryDetails = []model.Entry{
		{
			Amount: model.CurrencyAndAmount{
				Amount:   13139.57,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Credit,
			Status:               model.Book,
			BankTransactionCode:  model.TransCredit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: model.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000001",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310ISOTEST100000103100900FT02",
				InstructingAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice: model.InstrumentCTRC,
			},
		},
		{
			Amount: model.CurrencyAndAmount{
				Amount:   278.47,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Credit,
			Status:               model.Book,
			BankTransactionCode:  model.TransCredit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: model.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000002",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310ISOTEST100000203100900FT02",
				InstructingAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice: model.InstrumentCTRC,
			},
		},
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario2_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario2_Step2_camt.052_DTLR")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario2_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
