package Master

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreationDateTime, MessagePagination, ReportTypeId, ReportCreatedDate, AccountOtherId, AccountType, RelatedAccountOtherId, TransactionsSummary")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.MessageId == "" {
		msg.Data.MessageId = model.AccountBalanceReport
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
	if msg.Data.ReportTypeId == "" {
		msg.Data.ReportTypeId = ABMS
	}
	if msg.Data.ReportCreatedDate.IsZero() {
		msg.Data.ReportCreatedDate = time.Now()
	}
	if msg.Data.AccountOtherId == "" {
		msg.Data.AccountOtherId = "231981435"
	}
	if msg.Data.AccountType == "" {
		msg.Data.AccountType = "M"
	}
	if msg.Data.RelatedAccountOtherId == "" {
		msg.Data.RelatedAccountOtherId = "231981435"
	}
	if isEmpty(msg.Data.TransactionsSummary) {
		msg.Data.TransactionsSummary = []TotalsPerBankTransactionCode{
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
		}
	}
	return msg
}
func TestMasterFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step2_camt.052_ABAR_MM")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "ABAR", string(message.Doc.BkToCstmrAcctRpt.GrpHdr.MsgId))
	require.Equal(t, "1", string(message.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb))
	require.Equal(t, "20230921231981435ABARMMrequest1", string(message.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId))
	require.Equal(t, "camt.060.001.05", string(message.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId))
	require.Equal(t, "ABMS", string(message.Doc.BkToCstmrAcctRpt.Rpt.Id))
	require.Equal(t, "231981435", string(message.Doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr.Id))
	require.Equal(t, "M", string(*message.Doc.BkToCstmrAcctRpt.Rpt.Acct.Tp.Prtry))
	require.Equal(t, "231981435", string(message.Doc.BkToCstmrAcctRpt.Rpt.RltdAcct.Id.Othr.Id))
	require.Equal(t, "DLOD", string(*message.Doc.BkToCstmrAcctRpt.Rpt.Bal[0].Tp.CdOrPrtry.Prtry))
	require.Equal(t, "CRDT", string(message.Doc.BkToCstmrAcctRpt.Rpt.Bal[0].CdtDbtInd))
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

func TestMasterValidator(t *testing.T) {
	Balances := []Balance{
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
	Balances[0].BalanceTypeId = BalanceType(INVALID_COUNT)
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"Invalid MessageId",
			Message{Data: MessageModel{MessageId: model.CAMTReportType(INVALID_COUNT)}},
			"error occur at MessageId: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid MessagePagination",
			Message{Data: MessageModel{MessagePagination: model.MessagePagenation{
				PageNumber:        "INVALID_COUNT",
				LastPageIndicator: true,
			}}},
			"error occur at MessagePagination.PageNumber: INVALID_COUNT fails validation with pattern [0-9]{1,5}",
		},
		{
			"Invalid OriginalBusinessMsgNameId",
			Message{Data: MessageModel{OriginalBusinessMsgNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at OriginalBusinessMsgNameId: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"Invalid ReportTypeId",
			Message{Data: MessageModel{ReportTypeId: AccountReportType(INVALID_COUNT)}},
			"error occur at ReportTypeId: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid AccountOtherId",
			Message{Data: MessageModel{AccountOtherId: INVALID_ACCOUNT_ID}},
			"error occur at AccountOtherId: 123ABC789 fails validation with pattern [0-9]{9,9}",
		},
		{
			"Invalid Balances",
			Message{Data: MessageModel{Balances: Balances}},
			"error occur at Bal.BalanceTypeId: UNKNOWN fails enumeration validation",
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
func TestAccountBalanceReport_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = model.AccountBalanceReport
	message.Data.CreationDateTime = time.Now()
	message.Data.MessagePagination = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.Data.OriginalBusinessMsgId = "20230921231981435ABARMMrequest1"
	message.Data.OriginalBusinessMsgNameId = "camt.060.001.05"
	message.Data.OriginalBusinessMsgCreateTime = time.Now()

	message.Data.ReportTypeId = ABMS
	message.Data.ReportCreatedDate = time.Now()
	message.Data.AccountOtherId = "231981435"
	message.Data.AccountType = "M"
	message.Data.RelatedAccountOtherId = "231981435"

	message.Data.Balances = []Balance{
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
	message.Data.TransactionsSummary = []TotalsPerBankTransactionCode{
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

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountBalanceReport_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step2_camt.052_ABAR_MM")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step2_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
