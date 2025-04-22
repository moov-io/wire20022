package ActivityReport

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestActivityReportFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "ActivityReport_Scenario1_Step1_camt.052_ACTR")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.MsgId), "ACTR")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb), "1")
	require.Equal(t, bool(message.doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd), true)
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.Id), "EDAY")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr.Id), "011104238")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtries.NbOfNtries), "61")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries.NbOfNtries), "29")
	require.Equal(t, float64(message.doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries.Sum), 8775299.29)
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd[0].BkTxCd.Prtry.Cd), "SENT")
	require.Equal(t, string(message.doc.BkToCstmrAcctRpt.Rpt.Ntry[0].NtryDtls.TxDtls.Refs.MsgId), "20250310B1QDRCQR000001")
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"

func TestAccountBalanceReportValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"MessageId",
			Message{data: MessageModel{MessageId: "Unknown data"}},
			"error occur at MessageId: invalid CAMT report type: Unknown data",
		},
		{
			"Pagenation - PageNumber",
			Message{data: MessageModel{Pagenation: model.MessagePagenation{
				PageNumber:        "Unknown data",
				LastPageIndicator: true,
			}}},
			"error occur at Pagenation.PageNumber: Unknown data fails validation with pattern [0-9]{1,5}",
		},
		{
			"Pagenation - ReportType",
			Message{data: MessageModel{ReportType: "Unknown data"}},
			"error occur at ReportType: Unknown data fails enumeration validation",
		},
		{
			"AccountOtherId",
			Message{data: MessageModel{AccountOtherId: INVALID_ACCOUNT_ID}},
			"error occur at AccountOtherId: 123ABC789 fails validation with pattern [0-9]{9,9}",
		},
		{
			"TotalEntries",
			Message{data: MessageModel{TotalEntries: INVALID_COUNT}},
			"error occur at TotalEntries: UNKNOWN fails validation with pattern [0-9]{1,15}",
		},
		{
			"TotalCreditEntries - NumberOfEntries",
			Message{data: MessageModel{TotalCreditEntries: model.NumberAndSumOfTransactions{
				NumberOfEntries: INVALID_COUNT,
				Sum:             100.00,
			}}},
			"error occur at TotalCreditEntries.NumberOfEntries: UNKNOWN fails validation with pattern [0-9]{1,15}",
		},
		{
			"TotalDebitEntries - NumberOfEntries",
			Message{data: MessageModel{TotalDebitEntries: model.NumberAndSumOfTransactions{
				NumberOfEntries: INVALID_COUNT,
				Sum:             100.00,
			}}},
			"error occur at TotalDebitEntries.NumberOfEntries: UNKNOWN fails validation with pattern [0-9]{1,15}",
		},
		{
			"TotalEntriesPerBankTransactionCode - NumberOfEntries",
			Message{data: MessageModel{TotalEntriesPerBankTransactionCode: []TotalsPerBankTransactionCode{
				{
					NumberOfEntries:     INVALID_COUNT,
					BankTransactionCode: model.Sent,
				},
			}}},
			"error occur at TotalEntriesPerBankTransactionCode.NumberOfEntries: UNKNOWN fails validation with pattern [0-9]{1,15}",
		},
		{
			"TotalEntriesPerBankTransactionCode - BankTransactionCode",
			Message{data: MessageModel{TotalEntriesPerBankTransactionCode: []TotalsPerBankTransactionCode{
				{
					NumberOfEntries:     "56",
					BankTransactionCode: model.TransactionStatusCode(INVALID_COUNT),
				},
			}}},
			"error occur at TotalEntriesPerBankTransactionCode.BankTransactionCode: UNKNOWN fails enumeration validation",
		},
		{
			"EntryDetails - LocalInstrumentChoice",
			Message{data: MessageModel{EntryDetails: []model.Entry{
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
						LocalInstrumentChoice:   model.InstrumentPropCodeType(INVALID_COUNT),
						RelatedDatesProprietary: model.BusinessProcessingDate,
						RelatedDateTime:         time.Now(),
					},
				},
			}}},
			"error occur at EntryDetails.EntryDetails.LocalInstrumentChoice: UNKNOWN fails enumeration validation",
		},
		{
			"EntryDetails - InstructingAgent",
			Message{data: MessageModel{EntryDetails: []model.Entry{
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
							PaymentSysCode:     model.PaymentSystemType(INVALID_COUNT),
							PaymentSysMemberId: "231981435",
						},
						InstructedAgent: model.Agent{
							PaymentSysCode:     model.PaymentSysUSABA,
							PaymentSysMemberId: "011104238",
						},
						LocalInstrumentChoice:   model.InstrumentPropCodeType(INVALID_COUNT),
						RelatedDatesProprietary: model.BusinessProcessingDate,
						RelatedDateTime:         time.Now(),
					},
				},
			}}},
			"error occur at EntryDetails.EntryDetails.InstructingAgent.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			msgErr := tt.msg.CreateDocument()
			if msgErr != nil {
				require.Equal(t, tt.expectedErr, msgErr.Error())
			}
		})
	}
}
func TestActivityReport_Scenario1_Step1_camt_CreateXML(t *testing.T) {
	var mesage, err = NewMessage("")
	mesage.data.MessageId = model.ActivityReport
	mesage.data.CreatedDateTime = time.Now()
	mesage.data.Pagenation = model.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	mesage.data.ReportType = model.EveryDay
	mesage.data.ReportCreateDateTime = time.Now()
	mesage.data.AccountOtherId = "011104238"
	mesage.data.TotalEntries = "61"
	mesage.data.TotalCreditEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "29",
		Sum:             8775299.29,
	}
	mesage.data.TotalDebitEntries = model.NumberAndSumOfTransactions{
		NumberOfEntries: "27",
		Sum:             9932294.43,
	}
	mesage.data.TotalEntriesPerBankTransactionCode = []TotalsPerBankTransactionCode{
		{
			NumberOfEntries:     "0",
			BankTransactionCode: model.Sent,
		},
		{
			NumberOfEntries:     "5",
			BankTransactionCode: model.TransReceived,
		},
	}
	mesage.data.EntryDetails = []model.Entry{
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

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("ActivityReport_Scenario1_Step1_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "ActivityReport_Scenario1_Step1_camt.052_ACTR")
	genterated := filepath.Join("generated", "ActivityReport_Scenario1_Step1_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
