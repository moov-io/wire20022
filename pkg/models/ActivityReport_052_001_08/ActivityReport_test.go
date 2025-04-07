package ActivityReport_052_001_08

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestActivityReport_Scenario1_Step1_camt_CreateXML(t *testing.T) {
	var mesage = NewCamt052Message()
	mesage.model.MessageId = "20250311114001500ABARSrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.Pagenation = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	mesage.model.ReportType = model.EveryDay
	mesage.model.ReportCreateDateTime = time.Now()
	mesage.model.AccountOtherId = "011104238"
	mesage.model.TotalEntries = "1"
	mesage.model.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "29",
		Sum:             8775299.29,
	}
	mesage.model.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "27",
		Sum:             9932294.43,
	}
	mesage.model.TotalEntriesPerBankTransactionCode = []TotalsPerBankTransactionCode{
		{
			NumberOfEntries:     "0",
			BankTransactionCode: model.Sent,
		},
		{
			NumberOfEntries:     "5",
			BankTransactionCode: model.TransReceived,
		},
	}
	mesage.model.EntryDetails = []model.Entry{
		{
			Amount: model.CurrencyAndAmount{
				Amount:   240.67,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Debit,
			Status:               model.Book,
			BankTransactionCode:  model.TransDebit,
			MessageNameId:        "acs.008.001.08",
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
			MessageNameId:        "acs.008.001.08",
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
				Amount:   1000.00,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Debit,
			Status:               model.Book,
			BankTransactionCode:  model.TransDebit,
			MessageNameId:        "acs.008.001.08",
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
	}

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "ActivityReport_Scenario1_Step1_camt.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
