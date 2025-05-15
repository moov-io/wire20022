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
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreationDateTime, MessagePagination, BussinessQueryMsgId, BussinessQueryMsgNameId, BussinessQueryCreateDatetime, ReportId, ReportingSequence, AccountOtherId")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.MessageId == "" {
		msg.Data.MessageId = "DTLS"
	}
	if msg.Data.CreationDateTime.IsZero() {
		msg.Data.CreationDateTime = time.Now()
	}
	if isEmpty(msg.Data.MessagePagination) {
		msg.Data.MessagePagination = model.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		}
	}
	if msg.Data.BussinessQueryMsgId == "" {
		msg.Data.BussinessQueryMsgId = "20250311231981435DTLSrequest1"
	}
	if msg.Data.BussinessQueryMsgNameId == "" {
		msg.Data.BussinessQueryMsgNameId = "camt.060.001.05"
	}
	if msg.Data.BussinessQueryCreateDatetime.IsZero() {
		msg.Data.BussinessQueryCreateDatetime = time.Now()
	}
	if msg.Data.ReportId == "" {
		msg.Data.ReportId = model.Intraday
	}
	if isEmpty(msg.Data.ReportingSequence) {
		msg.Data.ReportingSequence = model.SequenceRange{
			FromSeq: "000001",
			ToSeq:   "000100",
		}
	}
	if msg.Data.AccountOtherId == "" {
		msg.Data.AccountOtherId = "B1QDRCQR"
	}
	return msg
}
func TestEndpointDetailsReportFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario1_Step2_camt.052_DTLS")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.Doc.BkToCstmrAcctRpt.GrpHdr.MsgId), "DTLS")
	require.Equal(t, string(message.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb), "1")
	require.Equal(t, string(message.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId), "20250311231981435DTLSrequest1")
	require.Equal(t, string(message.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId), "camt.060.001.05")
	require.Equal(t, string(message.Doc.BkToCstmrAcctRpt.Rpt.Id), "IDAY")
	require.Equal(t, string(message.Doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr.Id), "B1QDRCQR")
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
			Message{Data: MessageModel{MessageId: INVALID_MESSAGE_ID}},
			"error occur at MessageId: 12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456 fails validation with length 66 <= required maxLength 35",
		},
		{
			"BussinessQueryMsgId",
			Message{Data: MessageModel{BussinessQueryMsgId: INVALID_MESSAGE_ID}},
			"error occur at BussinessQueryMsgId: 12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456 fails validation with length 66 <= required maxLength 35",
		},
		{
			"BussinessQueryMsgNameId",
			Message{Data: MessageModel{BussinessQueryMsgNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at BussinessQueryMsgNameId: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"ReportId",
			Message{Data: MessageModel{ReportId: model.ReportType(INVALID_COUNT)}},
			"error occur at ReportId: UNKNOWN fails enumeration validation",
		},
		{
			"AccountOtherId",
			Message{Data: MessageModel{AccountOtherId: INVALID_OTHER_ID}},
			"error occur at AccountOtherId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [A-Z0-9]{8,8}",
		},
		{
			"TotalEntriesPerTransactionCode",
			Message{Data: MessageModel{TotalEntriesPerTransactionCode: []model.NumberAndStatusOfTransactions{
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
			Message{Data: MessageModel{TotalEntriesPerTransactionCode: []model.NumberAndStatusOfTransactions{
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
//		require.NoError(t, err)
//	}
func TestEndpointDetailsReport_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.Data.MessageId = "DTLS"
	message.Data.CreationDateTime = time.Now()
	message.Data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.Data.BussinessQueryMsgId = "20250311231981435DTLSrequest1"
	message.Data.BussinessQueryMsgNameId = "camt.060.001.05"
	message.Data.BussinessQueryCreateDatetime = time.Now()
	message.Data.ReportId = model.Intraday
	message.Data.ReportingSequence = model.SequenceRange{
		FromSeq: "000001",
		ToSeq:   "000100",
	}
	message.Data.ReportCreateDateTime = time.Now()
	message.Data.AccountOtherId = "B1QDRCQR"
	message.Data.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "100",
		Sum:             8307111.56,
	}
	message.Data.TotalEntriesPerTransactionCode = []model.NumberAndStatusOfTransactions{
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
	message.Data.EntryDetails = []model.Entry{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario1_Step2_camt.052_DTLS")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointDetailsReport_Scenario2_Step2_camt_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.Data.MessageId = "DTLR"
	message.Data.CreationDateTime = time.Now()
	message.Data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.Data.BussinessQueryMsgId = "20250311231981435DTLRrequest1"
	message.Data.BussinessQueryMsgNameId = "camt.060.001.05"
	message.Data.BussinessQueryCreateDatetime = time.Now()
	message.Data.ReportId = model.Intraday
	message.Data.ReportingSequence = model.SequenceRange{
		FromSeq: "000001",
		ToSeq:   "000100",
	}
	message.Data.ReportCreateDateTime = time.Now()
	message.Data.AccountOtherId = "B1QDRCQR"
	message.Data.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "94",
		Sum:             2871734.98,
	}
	message.Data.TotalEntriesPerTransactionCode = []model.NumberAndStatusOfTransactions{
		{
			NumberOfEntries: "6",
			Status:          model.TransReceived,
		},
	}
	message.Data.EntryDetails = []model.Entry{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario2_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario2_Step2_camt.052_DTLR")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario2_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
