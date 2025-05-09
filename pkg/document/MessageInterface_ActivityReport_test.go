package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/ActivityReport"
	"github.com/stretchr/testify/require"
)

func TestActivityReportParseXMLFile(t *testing.T) {
	xmlFile := "../models/ActivityReport/generated/ActivityReport_Scenario1_Step1_camt.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &ActivityReport.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*ActivityReport.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, model.ActivityReport)
	}
}

func TestActivityReportGenerateXML(t *testing.T) {
	dataModel := ActivityReportDataModel()
	xmlData, err := GenerateXML(&dataModel, &ActivityReport.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("ActivityReport_test.xml", xmlData)
	require.NoError(t, err)
}

func TestActivityReportRequireFieldCheck(t *testing.T) {
	dataModel := ActivityReportDataModel()
	dataModel.MessageId = ""
	dataModel.ReportType = ""
	valid, err := RequireFieldCheck(&dataModel, &ActivityReport.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId, ReportType")
}

func TestActivityReportXMLValidation(t *testing.T) {
	xmlFile := "../models/ActivityReport/swiftSample/ActivityReport_Scenario1_Step1_camt.052_ACTR"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &ActivityReport.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestActivityReportAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&ActivityReport.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*ActivityReport.MessageHelper); ok {
		require.Equal(t, helper.AccountOtherId.Title, "Account Other Id")
		require.Equal(t, helper.AccountOtherId.Type, "RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}")
		require.Equal(t, helper.AccountOtherId.Documentation, "Identification assigned by an institution.")
	}
}

func ActivityReportDataModel() ActivityReport.MessageModel {
	var mesage, _ = ActivityReport.NewMessage("")
	mesage.Data.MessageId = model.ActivityReport
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.Pagenation = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	mesage.Data.ReportType = model.EveryDay
	mesage.Data.ReportCreateDateTime = time.Now()
	mesage.Data.AccountOtherId = "011104238"
	mesage.Data.TotalEntries = "61"
	mesage.Data.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "29",
		Sum:             8775299.29,
	}
	mesage.Data.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "27",
		Sum:             9932294.43,
	}
	mesage.Data.TotalEntriesPerBankTransactionCode = []ActivityReport.TotalsPerBankTransactionCode{
		{
			NumberOfEntries:     "0",
			BankTransactionCode: model.Sent,
		},
		{
			NumberOfEntries:     "5",
			BankTransactionCode: model.TransReceived,
		},
	}
	mesage.Data.EntryDetails = []model.Entry{
		{
			Amount: model.CurrencyAndAmount{
				Amount:   240.67,
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
				ClearingSystemRef:          "20230310QMGFNP6000000103100900FT02",
				InstructingAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice:   model.InstrumentCTRC,
				RelatedDatesProprietary: model.BusinessProcessingDate,
				RelatedDateTime:         time.Now(),
			},
		},

		{
			Amount: model.CurrencyAndAmount{
				Amount:   1000.00,
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
				ClearingSystemRef:          "20230310QMGFNP6000000203100900FT02",
				InstructingAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice:   model.InstrumentCTRC,
				RelatedDatesProprietary: model.BusinessProcessingDate,
				RelatedDateTime:         time.Now(),
			},
		},

		{
			Amount: model.CurrencyAndAmount{
				Amount:   1197.00,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Debit,
			Status:               model.Book,
			BankTransactionCode:  model.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: model.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000003",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310QMGFNP6000000303100900FT02",
				InstructingAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice:   model.InstrumentCTRC,
				RelatedDatesProprietary: model.BusinessProcessingDate,
				RelatedDateTime:         time.Now(),
			},
		},
	}

	cErr := mesage.CreateDocument()
	if cErr != nil {
		return mesage.Data
	}
	return mesage.Data
}
