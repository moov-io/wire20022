package EndpointDetailsReport

import (
	"encoding/xml"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestEndpointDetailsReport_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
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
				MessageId:                  "20250310B1QDRCQR000003",
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
		{
			Amount: model.CurrencyAndAmount{
				Amount:   5749.56,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Debit,
			Status:               model.Book,
			BankTransactionCode:  model.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: model.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000004",
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
		{
			Amount: model.CurrencyAndAmount{
				Amount:   60000.00,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Debit,
			Status:               model.Book,
			BankTransactionCode:  model.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: model.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000005",
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("EndpointDetailsReport_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestEndpointDetailsReport_Scenario2_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.data.TotalDebitEntries = model.NumberAndSumOfTransactions{
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
		{
			Amount: model.CurrencyAndAmount{
				Amount:   1080.00,
				Currency: "USD",
			},
			CreditDebitIndicator: model.Credit,
			Status:               model.Book,
			BankTransactionCode:  model.TransCredit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: model.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000003",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310ISOTEST100000303100900FT02",
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("EndpointDetailsReport_Scenario2_Step2_camt.xml", xmlData)
	require.NoError(t, err)
}
