package PaymentReturn

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreatedDateTime, NumberOfTransactions, SettlementMethod, ClearingSystem, OriginalMessageId, OriginalMessageNameId, OriginalCreationDateTime, OriginalUETR, ReturnedInterbankSettlementAmount, InterbankSettlementDate, ReturnedInstructedAmount, InstructingAgent, InstructedAgent, RtrChain, ReturnReasonInformation, OriginalTransactionRef")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.MessageId == "" {
		msg.Data.MessageId = "20250310B1QDRCQR000724"
	}
	if msg.Data.CreatedDateTime.IsZero() {
		msg.Data.CreatedDateTime = time.Now()
	}
	if msg.Data.NumberOfTransactions == 0 {
		msg.Data.NumberOfTransactions = 1
	}
	if msg.Data.SettlementMethod == "" {
		msg.Data.SettlementMethod = model.SettlementCLRG
	}
	if msg.Data.ClearingSystem == "" {
		msg.Data.ClearingSystem = model.ClearingSysFDW
	}
	if msg.Data.OriginalMessageId == "" {
		msg.Data.OriginalMessageId = "20250310B1QDRCQR000721"
	}
	if msg.Data.OriginalMessageNameId == "" {
		msg.Data.OriginalMessageNameId = "pacs.008.001.08"
	}
	if msg.Data.OriginalCreationDateTime.IsZero() {
		msg.Data.OriginalCreationDateTime = time.Now()
	}
	if msg.Data.OriginalUETR == "" {
		msg.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	}
	if isEmpty(msg.Data.ReturnedInterbankSettlementAmount) {
		msg.Data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
			Amount:   151235.88,
			Currency: "USD",
		}
	}
	if isEmpty(msg.Data.InterbankSettlementDate) {
		msg.Data.InterbankSettlementDate = model.FromTime(time.Now())
	}
	if isEmpty(msg.Data.ReturnedInstructedAmount) {
		msg.Data.ReturnedInstructedAmount = model.CurrencyAndAmount{
			Amount:   151235.88,
			Currency: "USD",
		}
	}
	if isEmpty(msg.Data.InstructingAgent) {
		msg.Data.InstructingAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	if isEmpty(msg.Data.InstructedAgent) {
		msg.Data.InstructedAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if isEmpty(msg.Data.RtrChain) {
		msg.Data.RtrChain = ReturnChain{
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
			Creditor: Party{
				Name: "Corporation A",
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
		}
	}
	if isEmpty(msg.Data.ReturnReasonInformation) {
		msg.Data.ReturnReasonInformation = Reason{
			Reason:                "DUPL",
			AdditionalRequestData: "Order cancelled. Ref:20250310B1QDRCQR000721.",
		}
	}
	if msg.Data.OriginalTransactionRef == "" {
		msg.Data.OriginalTransactionRef = model.InstrumentCTRC
	}
	return msg
}
func TestPaymentReturnFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step4_pacs.004")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250310B1QDRCQR000724", string(message.Doc.PmtRtr.GrpHdr.MsgId))
	require.Equal(t, "1", string(message.Doc.PmtRtr.GrpHdr.NbOfTxs))
	require.Equal(t, "CLRG", string(message.Doc.PmtRtr.GrpHdr.SttlmInf.SttlmMtd))
	require.Equal(t, "FDW", string(*message.Doc.PmtRtr.GrpHdr.SttlmInf.ClrSys.Cd))
	require.Equal(t, "20250310B1QDRCQR000721", string(message.Doc.PmtRtr.TxInf.OrgnlGrpInf.OrgnlMsgId))
	require.Equal(t, "pacs.008.001.08", string(message.Doc.PmtRtr.TxInf.OrgnlGrpInf.OrgnlMsgNmId))
	require.Equal(t, "Scenario01InstrId001", string(*message.Doc.PmtRtr.TxInf.OrgnlInstrId))
	require.Equal(t, "Scenario01EtoEId001", string(*message.Doc.PmtRtr.TxInf.OrgnlEndToEndId))
	require.Equal(t, "8a562c67-ca16-48ba-b074-65581be6f011", string(message.Doc.PmtRtr.TxInf.OrgnlUETR))
	require.Equal(t, "SLEV", string(*message.Doc.PmtRtr.TxInf.ChrgBr))
	require.Equal(t, "USABA", string(*message.Doc.PmtRtr.TxInf.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "021040078", string(message.Doc.PmtRtr.TxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "USABA", string(*message.Doc.PmtRtr.TxInf.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "011104238", string(message.Doc.PmtRtr.TxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "Corporation B", string(*message.Doc.PmtRtr.TxInf.RtrChain.Dbtr.Pty.Nm))
	require.Equal(t, "Desert View Street", string(*message.Doc.PmtRtr.TxInf.RtrChain.Dbtr.Pty.PstlAdr.StrtNm))
	require.Equal(t, "567876543", string(message.Doc.PmtRtr.TxInf.RtrChain.DbtrAcct.Id.Othr.Id))
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

func TestPaymentReturnValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"Invalid MessageId",
			Message{Data: MessageModel{MessageId: INVALID_OTHER_ID}},
			"error occur at MessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid SettlementMethod",
			Message{Data: MessageModel{SettlementMethod: model.SettlementMethodType(INVALID_COUNT)}},
			"error occur at SettlementMethod: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid ClearingSystem",
			Message{Data: MessageModel{ClearingSystem: model.CommonClearingSysCodeType(INVALID_COUNT)}},
			"error occur at ClearingSystem: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid OriginalMessageId",
			Message{Data: MessageModel{OriginalMessageId: INVALID_OTHER_ID}},
			"error occur at OriginalMessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid OriginalMessageNameId",
			Message{Data: MessageModel{OriginalMessageNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at OriginalMessageNameId: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"Invalid ChargeBearer",
			Message{Data: MessageModel{ChargeBearer: model.ChargeBearerType(INVALID_COUNT)}},
			"error occur at ChargeBearer: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid InstructingAgent",
			Message{Data: MessageModel{InstructingAgent: model.Agent{
				PaymentSysCode:     INVALID_PAY_SYSCODE,
				PaymentSysMemberId: "021040078",
			}}},
			"error occur at InstructingAgent.PaymentSysCode: UNKNOWN fails enumeration validation",
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
func TestFedwireFundsAcknowledgement_Scenario2_Step4_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000724"
	message.Data.CreatedDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalCreationDateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ReturnedInstructedAmount = model.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.Data.ChargeBearer = model.ChargeBearerSLEV
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.RtrChain = ReturnChain{
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
	message.Data.ReturnReasonInformation = Reason{
		Reason:                "DUPL",
		AdditionalRequestData: "Order cancelled. Ref:20250310B1QDRCQR000721.",
	}
	message.Data.OriginalTransactionRef = model.InstrumentCTRC

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step4_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step4_pacs.004")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step4_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario2_Step5_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310QMGFT015000912"
	message.Data.CreatedDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.OriginalMessageId = "20250310B1QDRCQR000902"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalCreationDateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   510000.74,
		Currency: "USD",
	}
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ReturnedInstructedAmount = model.CurrencyAndAmount{
		Amount:   510000.74,
		Currency: "USD",
	}
	message.Data.ChargeBearer = model.ChargeBearerSLEV
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.RtrChain = ReturnChain{
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
	message.Data.ReturnReasonInformation = Reason{
		Reason:                "DUPL",
		AdditionalRequestData: "Payment returned. Ref:20250310B1QDRCQR000902.",
	}
	message.Data.OriginalTransactionRef = model.InstrumentCTRC

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario2_Step5_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario2_Step5_pacs.004")
	genterated := filepath.Join("generated", "Investigations_Scenario2_Step5_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario1_Step4_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310ISOTEST1000912"
	message.Data.CreatedDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.OriginalMessageId = "20250310B1QDRCQR000902"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalCreationDateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   151000.74,
		Currency: "USD",
	}
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ReturnedInstructedAmount = model.CurrencyAndAmount{
		Amount:   151000.74,
		Currency: "USD",
	}
	message.Data.ChargeBearer = model.ChargeBearerSLEV
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.RtrChain = ReturnChain{
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
	message.Data.ReturnReasonInformation = Reason{
		Reason:                "DUPL",
		AdditionalRequestData: "Payment deiplicate. Ref:20250310B1QDRCQR000902.",
	}
	message.Data.OriginalTransactionRef = model.InstrumentCTRC

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario1_Step4_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario1_Step4_pacs.004")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario1_Step4_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentReturn_Scenario3_Step4_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000433"
	message.Data.CreatedDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalCreationDateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario03InstrId001"
	message.Data.OriginalEndToEndId = "Scenario03EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.ReturnedInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   446915.78,
		Currency: "USD",
	}
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ReturnedInstructedAmount = model.CurrencyAndAmount{
		Amount:   446915.78,
		Currency: "USD",
	}
	message.Data.ChargeBearer = model.ChargeBearerSHAR
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.RtrChain = ReturnChain{
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
	message.Data.ReturnReasonInformation = Reason{
		Reason:                "FOCR",
		AdditionalRequestData: "As per agreed resolution. Ref:20250310B1QDRCQR000400.",
	}
	message.Data.OriginalTransactionRef = model.InstrumentCTRC

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentReturn_Scenario3_Step4_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario3_Step4_pacs.004")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario3_Step4_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
