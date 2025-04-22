package Master

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestAccountBalanceReport_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = model.AccountBalanceReport
	message.data.CreationDateTime = time.Now()
	message.data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.data.OriginalBusinessMsgId = "20230921231981435ABARMMrequest1"
	message.data.OriginalBusinessMsgNameId = "camt.060.001.05"
	message.data.OriginalBusinessMsgCreateTime = time.Now()

	message.data.ReportTypeId = ABMS
	message.data.ReportCreatedDate = time.Now()
	message.data.AccountOtherId = "231981435"
	message.data.AccountType = "M"
	message.data.RelatedAccountOtherId = "231981435"

	message.data.Balances = []Balance{
		{
			BalanceTypeId: DaylightOverdraftBalance,
			Amount: model.CurrencyAndAmount{
				Amount:   270458895930.79,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Credit,
			DateTime:             time.Now(),
		},
		{
			BalanceTypeId: AccountBalance,
			CdtLines: []CreditLine{
				{
					Included: true,
					Type:     NetDebitCap,
					Amount: model.CurrencyAndAmount{
						Amount:   23125500000.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     CollateralizedCapacity,
					Amount: model.CurrencyAndAmount{
						Amount:   316874500000.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     CollateralAvailable,
					Amount: model.CurrencyAndAmount{
						Amount:   82598573368.44,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     CollateralizedDaylightOverdrafts,
					Amount: model.CurrencyAndAmount{
						Amount:   0.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     UncollateralizedDaylightOverdrafts,
					Amount: model.CurrencyAndAmount{
						Amount:   0.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
			},
			Amount: model.CurrencyAndAmount{
				Amount:   270594506052.13,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Credit,
			DateTime:             time.Now(),
		},
		{
			BalanceTypeId: AvailableBalanceFromDaylightOverdraft,
			Amount: model.CurrencyAndAmount{
				Amount:   610458895930.79,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Credit,
			DateTime:             time.Now(),
		},
	}
	message.data.TransactionsSummary = []TotalsPerBankTransactionCode{
		{
			TotalNetEntryAmount:  279595877422.72,
			CreditDebitIndicator: model.Credit,
			CreditEntries: model.NumberAndSumOfTransactions{
				NumberOfEntries: "16281",
				Sum:             420780358976.96,
			},
			DebitEntries: model.NumberAndSumOfTransactions{
				NumberOfEntries: "22134",
				Sum:             141184481554.24,
			},
			BankTransactionCode: FedwireFundsTransfers,
			Date:                time.Now(),
		},
		{
			TotalNetEntryAmount:  608598873.60,
			CreditDebitIndicator: model.Credit,
			CreditEntries: model.NumberAndSumOfTransactions{
				NumberOfEntries: "4",
				Sum:             993425694.01,
			},
			DebitEntries: model.NumberAndSumOfTransactions{
				NumberOfEntries: "6",
				Sum:             384826820.41,
			},
			BankTransactionCode: NationalSettlementServiceEntries,
			Date:                time.Now(),
		},
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("AccountBalanceReport_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step2_camt.052_ABAR_MM")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
