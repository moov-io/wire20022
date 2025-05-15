package CustomerCreditTransfer

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestPaymentReturn_Scenario1_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000400"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario01InstrId001"
	mesage.Data.EndToEndId = "Scenario01EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1510000.74,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerSLEV
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.Data.DebtorName = "Corporation A"
	mesage.Data.DebtorAddress = model.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.Data.DebtorOtherTypeId = "5647772655"
	mesage.Data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.Data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.Data.CreditorName = "Corporation B"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.Data.CreditorOtherTypeId = "567876543"
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario1_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario1_Step1_pacs.008")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario2_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000400"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario02InstrId001"
	mesage.Data.EndToEndId = "Scenario02EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1234578.88,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1234578.88,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerSLEV
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.Data.DebtorName = "Corporation C"
	mesage.Data.DebtorAddress = model.PostalAddress{
		StreetName:     "40th Street",
		BuildingNumber: "1180",
		PostalCode:     "11218",
		TownName:       "Brooklyn",
		Subdivision:    "NY",
		Country:        "US",
	}
	mesage.Data.DebtorOtherTypeId = "0031234567"
	mesage.Data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.Data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.Data.CreditorName = "Corporation B"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.Data.CreditorOtherTypeId = "567876543"
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario2_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario2_Step1_pacs.008")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario2_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario3_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000400"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario03InstrId001"
	mesage.Data.EndToEndId = "Scenario03EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 2234578.88,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 2234578.88,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerSLEV
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.Data.DebtorName = "Corporation C"
	mesage.Data.DebtorAddress = model.PostalAddress{
		StreetName:     "40th Street",
		BuildingNumber: "1180",
		PostalCode:     "11218",
		TownName:       "Brooklyn",
		Subdivision:    "NY",
		Country:        "US",
	}
	mesage.Data.DebtorOtherTypeId = "0031234567"
	mesage.Data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.Data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.Data.CreditorName = "Corporation B"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.Data.CreditorOtherTypeId = "567876543"
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario3_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario3_Step1_pacs.008")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario3_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario4_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000400"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario04InstrId001"
	mesage.Data.EndToEndId = "Scenario04EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 2234578.88,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 2234578.88,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerSLEV
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.Data.DebtorName = "Corporation C"
	mesage.Data.DebtorAddress = model.PostalAddress{
		StreetName:     "40th Street",
		BuildingNumber: "1180",
		PostalCode:     "11218",
		TownName:       "Brooklyn",
		Subdivision:    "NY",
		Country:        "US",
	}
	mesage.Data.DebtorOtherTypeId = "0031234567"
	mesage.Data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.Data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.Data.CreditorName = "Corporation B"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.Data.CreditorOtherTypeId = "567876543"
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario4_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario4_Step1_pacs.008")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario4_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario5_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000450"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario05InstrId001"
	mesage.Data.EndToEndId = "Scenario05EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.SericeLevel = "G001"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 3234578.89,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 3234578.89,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerSHAR
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	mesage.Data.DebtorName = "Corporation Z"
	mesage.Data.DebtorAddress = model.PostalAddress{
		StreetName: "Avenue Moliere 70",
		PostalCode: "1180",
		TownName:   "Brussels",
		Country:    "BE",
	}
	mesage.Data.DebtorIBAN = "BE34001216371411"
	mesage.Data.DebtorAgent = model.Agent{
		BusinessIdCode: "BANZBEBB",
	}
	mesage.Data.CreditorAgent = model.Agent{
		BusinessIdCode: "BANYBRRJ",
	}
	mesage.Data.CreditorName = "Corporation Y"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Av. Lucio Costa",
		BuildingNumber: "5220",
		BuildingName:   "Barra da Tijuca",
		PostalCode:     "22630-012",
		TownName:       "Rio de Janeiro",
		Country:        "BR",
	}
	mesage.Data.CreditorIBAN = "BR9700360305000010009795493P1"
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCMCN,
		Number:            "ABC-987",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario5_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario5_Step1_pacs.008")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario5_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
