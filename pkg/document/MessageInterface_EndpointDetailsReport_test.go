package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/EndpointDetailsReport"
	"github.com/stretchr/testify/require"
)

func TestEndpointDetailsReportParseXMLFile(t *testing.T) {
	xmlFile := "../models/EndpointDetailsReport/generated/EndpointDetailsReport_Scenario1_Step2_camt.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &EndpointDetailsReport.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*EndpointDetailsReport.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "DTLS")
	}
}

func TestEndpointDetailsReportGenerateXML(t *testing.T) {
	dataModel := EndpointDetailsReportDataModel()
	xmlData, err := GenerateXML(&dataModel, &EndpointDetailsReport.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_test.xml", xmlData)
	require.NoError(t, err)
}

func TestEndpointDetailsReportRequireFieldCheck(t *testing.T) {
	dataModel := EndpointDetailsReportDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &EndpointDetailsReport.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestEndpointDetailsReportXMLValidation(t *testing.T) {
	xmlFile := "../models/EndpointDetailsReport/swiftSample/EndpointDetailsReport_Scenario1_Step2_camt.052_DTLS"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &EndpointDetailsReport.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestEndpointDetailsReportAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&EndpointDetailsReport.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*EndpointDetailsReport.MessageHelper); ok {
		require.Equal(t, helper.AccountOtherId.Title, "Account Other Id")
		require.Equal(t, helper.AccountOtherId.Type, "Max34Text (based on string) minLength: 1 maxLength: 34")
		require.Equal(t, helper.AccountOtherId.Documentation, "Unique identification of an account, as assigned by the account servicer, using an identification scheme.")
	}
}

func EndpointDetailsReportDataModel() EndpointDetailsReport.MessageModel {
	var message, _ = EndpointDetailsReport.NewMessage("")
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
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
