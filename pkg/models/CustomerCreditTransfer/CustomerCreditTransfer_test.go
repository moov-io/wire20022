package CustomerCreditTransfer

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreatedDateTime, NumberOfTransactions, SettlementMethod, CommonClearingSysCode, InstructionId, EndToEndId, UniqueEndToEndTransactionRef, InstrumentPropCode, InterBankSettAmount, InterBankSettDate, InstructedAmount, ChargeBearer, InstructingAgents, InstructedAgent, DebtorName, DebtorAddress, DebtorAgent, CreditorAgent, DebtorAgent")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.MessageId == "" {
		msg.Data.MessageId = "20250310B1QDRCQR000001"
	}
	if isEmpty(msg.Data.CreatedDateTime) {
		msg.Data.CreatedDateTime = time.Now()
	}
	if msg.Data.NumberOfTransactions == 0 {
		msg.Data.NumberOfTransactions = 1
	}
	if msg.Data.SettlementMethod == "" {
		msg.Data.SettlementMethod = model.SettlementCLRG
	}
	if msg.Data.CommonClearingSysCode == "" {
		msg.Data.CommonClearingSysCode = model.ClearingSysFDW
	}
	if msg.Data.InstructionId == "" {
		msg.Data.InstructionId = "DefaultInstrId001"
	}
	if msg.Data.EndToEndId == "" {
		msg.Data.EndToEndId = "DefaultEtoEId001"
	}
	if msg.Data.UniqueEndToEndTransactionRef == "" {
		msg.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	}
	if msg.Data.InstrumentPropCode == "" {
		msg.Data.InstrumentPropCode = model.InstrumentCTRC
	}
	if msg.Data.SericeLevel == "" {
		msg.Data.SericeLevel = "G001"
	}
	if isEmpty(msg.Data.InterBankSettAmount) {
		msg.Data.InterBankSettAmount = model.CurrencyAndAmount{
			Currency: "USD", Amount: 1000.00,
		}
	}
	if isEmpty(msg.Data.InterBankSettDate) {
		msg.Data.InterBankSettDate = model.FromTime(time.Now())
	}
	if isEmpty(msg.Data.InstructedAmount) {
		msg.Data.InstructedAmount = model.CurrencyAndAmount{
			Currency: "USD", Amount: 1000.00,
		}
	}
	if msg.Data.ChargeBearer == "" {
		msg.Data.ChargeBearer = model.ChargeBearerSLEV
	}
	if isEmpty(msg.Data.InstructingAgents) {
		msg.Data.InstructingAgents = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if isEmpty(msg.Data.InstructedAgent) {
		msg.Data.InstructedAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	if msg.Data.DebtorName == "" {
		msg.Data.DebtorName = "Default Debtor"
	}
	if msg.Data.DebtorOtherTypeId == "" {
		msg.Data.DebtorOtherTypeId = "123456789"
	}
	if isEmpty(msg.Data.DebtorAddress) {
		msg.Data.DebtorAddress = model.PostalAddress{
			StreetName: "Default Street", PostalCode: "12345", TownName: "Default Town", Country: "US",
		}
	}
	if isEmpty(msg.Data.DebtorAgent) {
		msg.Data.DebtorAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if isEmpty(msg.Data.CreditorAgent) {
		msg.Data.CreditorAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	return msg
}
func TestCustomerCreditTransferFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step1_pacs.008")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.Doc.FIToFICstmrCdtTrf.GrpHdr.MsgId), "20250310B1QDRCQR000001")
	require.Equal(t, string(message.Doc.FIToFICstmrCdtTrf.GrpHdr.NbOfTxs), "1")
	require.Equal(t, string(message.Doc.FIToFICstmrCdtTrf.GrpHdr.SttlmInf.SttlmMtd), "CLRG")
	require.Equal(t, string(*message.Doc.FIToFICstmrCdtTrf.GrpHdr.SttlmInf.ClrSys.Cd), "FDW")
	require.Equal(t, string(*message.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.InstrId), "Scenario01InstrId001")
	require.Equal(t, string(message.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.EndToEndId), "Scenario01EtoEId001")
	require.Equal(t, string(message.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtId.UETR), "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, string(*message.Doc.FIToFICstmrCdtTrf.CdtTrfTxInf.PmtTpInf.LclInstrm.Prtry), "CTRC")
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"

func TestCustomerCreditTransferValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"MessageId",
			Message{Data: MessageModel{MessageId: "20250310B1QDRCQR000001"}},
			"",
		},
		{
			"SettlementMethod",
			Message{Data: MessageModel{SettlementMethod: model.SettlementMethodType(INVALID_COUNT)}},
			"error occur at SettlementMethod: UNKNOWN fails enumeration validation",
		},
		{
			"CommonClearingSysCode",
			Message{Data: MessageModel{CommonClearingSysCode: model.CommonClearingSysCodeType(INVALID_COUNT)}},
			"error occur at CommonClearingSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"InstructionId",
			Message{Data: MessageModel{InstructionId: "1234567890123456789012345678901234567890"}},
			"error occur at InstructionId: 1234567890123456789012345678901234567890 fails validation with length 40 <= required maxLength 35",
		},
		{
			"InstrumentPropCode",
			Message{Data: MessageModel{InstrumentPropCode: model.InstrumentPropCodeType(INVALID_COUNT)}},
			"error occur at Instrument.InstrumentPropCode: UNKNOWN fails enumeration validation",
		},
		{
			"InstructingAgents - PaymentSysCode",
			Message{Data: MessageModel{InstructingAgents: model.Agent{
				PaymentSysCode:     model.PaymentSystemType(INVALID_COUNT),
				PaymentSysMemberId: "011104238",
			}}},
			"error occur at InstructingAgents.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"InstructingAgents - PaymentSysMemberId",
			Message{Data: MessageModel{InstructingAgents: model.Agent{
				PaymentSysCode:     model.PaymentSysUSABA,
				PaymentSysMemberId: "----.----.---",
			}}},
			"error occur at Instrument.PaymentSysMemberId: UNKNOWN fails enumeration validation",
		},
		{
			"DebtorOtherTypeId",
			Message{Data: MessageModel{DebtorOtherTypeId: "1234567890123456789012345678901234567890"}},
			"error occur at DebtorOtherTypeId: 1234567890123456789012345678901234567890 fails validation with length 40 <= required maxLength 34",
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
func TestCustomerCreditTransfer_Scenario1_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000001"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario01InstrId001"
	mesage.Data.EndToEndId = "Scenario01EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario1_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step1_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario1_Step2_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000001"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario01InstrId001"
	mesage.Data.EndToEndId = "Scenario01EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario1_Step2.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step2_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario2_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000002"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario02InstrId001"
	mesage.Data.EndToEndId = "Scenario02EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerSLEV
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040079",
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario2_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario2_Step1_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario3_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000001"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario01InstrId001"
	mesage.Data.EndToEndId = "Scenario01EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step1_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario4_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000004"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Scenario04InstrId001"
	mesage.Data.EndToEndId = "Scenario04EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 999008.53,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 999008.53,
	}
	mesage.Data.ChargeBearer = model.ChargeBearerDEBT
	mesage.Data.InstructingAgents = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "091036164",
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
		PaymentSysMemberId: "091036164",
		BankName:           "Internal Revenue Service",
		PostalAddress: model.PostalAddress{
			StreetName:     "West Perching Road",
			BuildingNumber: "333",
			PostalCode:     "64108",
			TownName:       "Kansas City",
			Subdivision:    "MO",
			Country:        "US",
		},
	}
	mesage.Data.CreditorName = "Corporation A"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.Data.RemittanceInfor = RemittanceDocument{
		TaxDetail: TaxRecord{
			TaxId:              "123456789",
			TaxTypeCode:        "09455",
			TaxPeriodYear:      model.FromTime(time.Now()),
			TaxperiodTimeFrame: "MM04",
		},
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario4_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario4_Step1_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario5_Step1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000005"
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
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario5_Step1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario5_Step1_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario5_Step2_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000005"
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
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario5_Step2.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario5_Step2_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}

func TestCustomerCreditTransfer_Variantion1_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000006"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Variation1InstrId001"
	mesage.Data.EndToEndId = "Variation1EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	mesage.Data.PurposeOfPayment = InvestmentPayment
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Variantion1.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Variation1_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Variantion1.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Variantion2_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000007"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Variation2InstrId001"
	mesage.Data.EndToEndId = "Variation2EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	// mesage.Data.PurposeOfPayment = InvestmentPayment
	mesage.Data.RelatedRemittanceInfo = RemittanceDetail{
		RemittanceId:      "Scenario01Var2RemittanceId001",
		Method:            Email,
		ElectronicAddress: "CustomerService@CorporationB.com",
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Variantion2.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Variation2_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Variantion2.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Variantion3_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000008"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Variation3InstrId001"
	mesage.Data.EndToEndId = "Variation3EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	mesage.Data.UltimateDebtorName = "Corporation Aa"
	mesage.Data.UltimateDebtorAddress = model.PostalAddress{
		StreetName:     "Ocean Street",
		BuildingNumber: "1",
		PostalCode:     "97035",
		TownName:       "Portland",
		Subdivision:    "OR",
		Country:        "US",
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
	// mesage.Data.PurposeOfPayment = InvestmentPayment
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Variantion3.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Variation3_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Variantion3.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Variantion4_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000009"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Variation4InstrId001"
	mesage.Data.EndToEndId = "Variation4EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
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
	mesage.Data.UltimateCreditorName = "Corporation Bb"
	mesage.Data.UltimateCreditorAddress = model.PostalAddress{
		StreetName:     "9th Avenue",
		BuildingNumber: "66",
		BuildingName:   "The Porter House",
		RoomNumber:     "Unit 6",
		PostalCode:     "10011",
		TownName:       "New York",
		Subdivision:    "NY",
		Country:        "US",
	}
	mesage.Data.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: model.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       model.FromTime(time.Now()),
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Variantion4.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Variation4_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Variantion4.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Variantion5_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000001"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Variation5InstrId001"
	mesage.Data.EndToEndId = "Variation5EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.SericeLevel = "G001"
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 1009858.99,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "EUR", Amount: 1000000.00,
	}
	mesage.Data.exchangeRate = 0.9901
	mesage.Data.ChargeBearer = model.ChargeBearerCREDIT
	mesage.Data.ChargesInfo = []ChargeInfo{
		{
			Amount:         model.CurrencyAndAmount{Currency: "USD", Amount: 90.00},
			BusinessIdCode: "BANZBEBB",
		},
		{
			Amount:         model.CurrencyAndAmount{Currency: "USD", Amount: 40.00},
			BusinessIdCode: "BANCUS33",
		},
	}

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
	err = model.WriteXMLTo("CustomerCreditTransfer_Variantion5.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Variation5_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Variantion5.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Variantion6_CreateXML(t *testing.T) {
	var mesage, vErr = NewMessage("")
	require.NoError(t, vErr)
	mesage.Data.MessageId = "20250310B1QDRCQR000001"
	mesage.Data.CreatedDateTime = time.Now()
	mesage.Data.NumberOfTransactions = 1
	mesage.Data.SettlementMethod = model.SettlementCLRG
	mesage.Data.CommonClearingSysCode = model.ClearingSysFDW
	mesage.Data.InstructionId = "Variation6InstrId001"
	mesage.Data.EndToEndId = "Variation6EtoEId001"
	mesage.Data.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.Data.InstrumentPropCode = model.InstrumentCTRC
	mesage.Data.SericeLevel = "G001"
	mesage.Data.InterBankSettAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 500000.00,
	}
	mesage.Data.InterBankSettDate = model.FromTime(time.Now())
	mesage.Data.InstructedAmount = model.CurrencyAndAmount{
		Currency: "USD", Amount: 500000.00,
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
	mesage.Data.IntermediaryAgent1Id = "BANYBRRJ"
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
		BusinessIdCode: "BANXBRRJ",
	}
	mesage.Data.CreditorName = "Individual X"
	mesage.Data.CreditorPostalAddress = model.PostalAddress{
		StreetName:     "Rua Aprazivel",
		BuildingNumber: "52",
		PostalCode:     "22630-012",
		TownName:       "Rio de Janeiro",
		Country:        "BR",
	}
	mesage.Data.CreditorIBAN = "BR1800360305000010009795493C1"
	mesage.Data.PurposeOfPayment = INSCPayment
	mesage.Data.RemittanceInfor = RemittanceDocument{
		UnstructuredRemitInfo: "Insurance Benefit/Policy XXAB9876/$500000.00",
	}
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Variantion6.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Variation6_pacs.008")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Variantion6.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
