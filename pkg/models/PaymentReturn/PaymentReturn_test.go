package PaymentReturn

import (
	"encoding/xml"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsAcknowledgement_Scenario2_Step4_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000724"
	message.data.CreatedDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ReturnedInstructedAmount = model.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.data.ChargeBearer = model.ChargeBearerSLEV
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.RtrChain = ReturnChain{
		Debtor: Party{
			Name: "Corporation B",
			Address: model.PostalAddress{
				StreetName:     "Desert View Street",
				BuildingNumber: "1",
				Floor:          "33",
				PostalCode:     "92262",
				TownName:       "Palm Springs",
				Subdivision:    "CA",
				Country:        "US",
			},
		},
		DebtorOtherTypeId: "567876543",
		DebtorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
			BankName:           "BankB",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue B",
				BuildingNumber: "25",
				PostalCode:     "19067",
				TownName:       "Yardley",
				Subdivision:    "PA",
				Country:        "US",
			},
		},
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
			BankName:           "BankA",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue A",
				BuildingNumber: "66",
				PostalCode:     "60532",
				TownName:       "Lisle",
				Subdivision:    "IL",
				Country:        "US",
			},
		},
		Creditor: Party{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CreditorAccountOtherTypeId: "5647772655",
	}
	message.data.ReturnReasonInformation = Reason{
		Reason:                "DUPL",
		AdditionalRequestData: "Order cancelled. Ref:20250310B1QDRCQR000721.",
	}
	message.data.OriginalTransactionRef = model.InstrumentCTRC

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step4_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestInvestigations_Scenario2_Step5_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFT015000912"
	message.data.CreatedDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.OriginalMessageId = "20250310B1QDRCQR000902"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   510000.74,
		Currency: "USD",
	}
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ReturnedInstructedAmount = model.CurrencyAndAmount{
		Amount:   510000.74,
		Currency: "USD",
	}
	message.data.ChargeBearer = model.ChargeBearerSLEV
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.RtrChain = ReturnChain{
		Debtor: Party{
			Name: "Corporation B",
			Address: model.PostalAddress{
				StreetName:     "Desert View Street",
				BuildingNumber: "1",
				Floor:          "33",
				PostalCode:     "92262",
				TownName:       "Palm Springs",
				Subdivision:    "CA",
				Country:        "US",
			},
		},
		DebtorOtherTypeId: "567876543",
		DebtorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
			BankName:           "BankB",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue B",
				BuildingNumber: "25",
				PostalCode:     "19067",
				TownName:       "Yardley",
				Subdivision:    "PA",
				Country:        "US",
			},
		},
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
			BankName:           "BankA",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue A",
				BuildingNumber: "66",
				PostalCode:     "60532",
				TownName:       "Lisle",
				Subdivision:    "IL",
				Country:        "US",
			},
		},
		Creditor: Party{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CreditorAccountOtherTypeId: "5647772655",
	}
	message.data.ReturnReasonInformation = Reason{
		Reason:                "DUPL",
		AdditionalRequestData: "Payment returned. Ref:20250310B1QDRCQR000902.",
	}
	message.data.OriginalTransactionRef = model.InstrumentCTRC

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Investigations_Scenario2_Step5_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario1_Step4_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310ISOTEST1000912"
	message.data.CreatedDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.OriginalMessageId = "20250310B1QDRCQR000902"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1510000.74,
		Currency: "USD",
	}
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ReturnedInstructedAmount = model.CurrencyAndAmount{
		Amount:   1510000.74,
		Currency: "USD",
	}
	message.data.ChargeBearer = model.ChargeBearerSLEV
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.RtrChain = ReturnChain{
		Debtor: Party{
			Name: "Corporation B",
			Address: model.PostalAddress{
				StreetName:     "Desert View Street",
				BuildingNumber: "1",
				Floor:          "33",
				PostalCode:     "92262",
				TownName:       "Palm Springs",
				Subdivision:    "CA",
				Country:        "US",
			},
		},
		DebtorOtherTypeId: "567876543",
		DebtorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
			BankName:           "BankB",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue B",
				BuildingNumber: "25",
				PostalCode:     "19067",
				TownName:       "Yardley",
				Subdivision:    "PA",
				Country:        "US",
			},
		},
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
			BankName:           "BankA",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue A",
				BuildingNumber: "66",
				PostalCode:     "60532",
				TownName:       "Lisle",
				Subdivision:    "IL",
				Country:        "US",
			},
		},
		Creditor: Party{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CreditorAccountOtherTypeId: "5647772655",
	}
	message.data.ReturnReasonInformation = Reason{
		Reason:                "DUPL",
		AdditionalRequestData: "Payment deiplicate. Ref:20250310B1QDRCQR000902.",
	}
	message.data.OriginalTransactionRef = model.InstrumentCTRC

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("PaymentReturn_Scenario1_Step4_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario3_Step4_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000433"
	message.data.CreatedDateTime = time.Now()
	message.data.NumberOfTransactions = 1
	message.data.SettlementMethod = model.SettlementCLRG
	message.data.ClearingSystem = model.ClearingSysFDW
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario03InstrId001"
	message.data.OriginalEndToEndId = "Scenario03EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   446915.78,
		Currency: "USD",
	}
	message.data.InterbankSettlementDate = model.FromTime(time.Now())
	message.data.ReturnedInstructedAmount = model.CurrencyAndAmount{
		Amount:   446915.78,
		Currency: "USD",
	}
	message.data.ChargeBearer = model.ChargeBearerSLEV
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.RtrChain = ReturnChain{
		Debtor: Party{
			Name: "Corporation B",
			Address: model.PostalAddress{
				StreetName:     "Desert View Street",
				BuildingNumber: "1",
				Floor:          "33",
				PostalCode:     "92262",
				TownName:       "Palm Springs",
				Subdivision:    "CA",
				Country:        "US",
			},
		},
		DebtorOtherTypeId: "567876543",
		DebtorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
			BankName:           "BankB",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue B",
				BuildingNumber: "25",
				PostalCode:     "19067",
				TownName:       "Yardley",
				Subdivision:    "PA",
				Country:        "US",
			},
		},
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
			BankName:           "BankA",
			PostalAddress: model.PostalAddress{
				StreetName:     "Avenue A",
				BuildingNumber: "66",
				PostalCode:     "60532",
				TownName:       "Lisle",
				Subdivision:    "IL",
				Country:        "US",
			},
		},
		Creditor: Party{
			Name: "Corporation C",
			Address: model.PostalAddress{
				StreetName:     "40th Street",
				BuildingNumber: "1180",
				PostalCode:     "11218",
				TownName:       "Brooklyn",
				Subdivision:    "NY",
				Country:        "US",
			},
		},
		CreditorAccountOtherTypeId: "0031234567",
	}
	message.data.ReturnReasonInformation = Reason{
		Reason:                "FOCR",
		AdditionalRequestData: "As per agreed resolution. Ref:20250310B1QDRCQR000400.",
	}
	message.data.OriginalTransactionRef = model.InstrumentCTRC

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("PaymentReturn_Scenario3_Step4_pacs.xml", xmlData)
	require.NoError(t, err)
}
