package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/Master"
	"github.com/stretchr/testify/require"
)

var MasterxmlFile = "../models/Master/swiftSample/AccountBalanceReport_Scenario1_Step2_camt.052_ABAR_MM"

func TestMasterParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(MasterxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &Master.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*Master.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, model.CAMTReportType("ABAR"))
	}
}

func TestMasterGenerateXML(t *testing.T) {
	dataModel := MasterDataModel()
	xmlData, err := GenerateXML(&dataModel, &Master.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("Master_test.xml", xmlData)
	require.NoError(t, err)
}

func TestMasterRequireFieldCheck(t *testing.T) {
	dataModel := MasterDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &Master.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestMasterXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(MasterxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &Master.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestMasterAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&Master.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*Master.MessageHelper); ok {
		require.Equal(t, helper.AccountOtherId.Title, "Account Other Id")
		require.Equal(t, helper.AccountOtherId.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.AccountOtherId.Documentation, "Unambiguous identification of the account to which credit and debit entries are made.")
	}
}

func MasterDataModel() Master.MessageModel {
	var message, _ = Master.NewMessage("")
	message.Data.MessageId = model.AccountBalanceReport
	message.Data.CreationDateTime = time.Now()
	message.Data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.Data.OriginalBusinessMsgId = "20230921231981435ABARMMrequest1"
	message.Data.OriginalBusinessMsgNameId = "camt.060.001.05"
	message.Data.OriginalBusinessMsgCreateTime = time.Now()

	message.Data.ReportTypeId = Master.ABMS
	message.Data.ReportCreatedDate = time.Now()
	message.Data.AccountOtherId = "231981435"
	message.Data.AccountType = "M"
	message.Data.RelatedAccountOtherId = "231981435"

	message.Data.Balances = []Master.Balance{
		{
			BalanceTypeId: Master.DaylightOverdraftBalance,
			Amount: model.CurrencyAndAmount{
				Amount:   270458895930.79,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Credit,
			DateTime:             time.Now(),
		},
		{
			BalanceTypeId: Master.AccountBalance,
			CdtLines: []Master.CreditLine{
				{
					Included: true,
					Type:     Master.NetDebitCap,
					Amount: model.CurrencyAndAmount{
						Amount:   23125500000.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     Master.CollateralizedCapacity,
					Amount: model.CurrencyAndAmount{
						Amount:   316874500000.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     Master.CollateralAvailable,
					Amount: model.CurrencyAndAmount{
						Amount:   82598573368.44,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     Master.CollateralizedDaylightOverdrafts,
					Amount: model.CurrencyAndAmount{
						Amount:   0.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     Master.UncollateralizedDaylightOverdrafts,
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
			BalanceTypeId: Master.AvailableBalanceFromDaylightOverdraft,
			Amount: model.CurrencyAndAmount{
				Amount:   610458895930.79,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Credit,
			DateTime:             time.Now(),
		},
	}
	message.Data.TransactionsSummary = []Master.TotalsPerBankTransactionCode{
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
			BankTransactionCode: Master.FedwireFundsTransfers,
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
			BankTransactionCode: Master.NationalSettlementServiceEntries,
			Date:                time.Now(),
		},
	}

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
