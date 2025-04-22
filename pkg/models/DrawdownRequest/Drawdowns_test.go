package DrawdownRequest

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdownRequestFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step1_pain.013")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.doc.CdtrPmtActvtnReq.GrpHdr.MsgId), "20250310B1QDRCQR000601")
	require.Equal(t, string(message.doc.CdtrPmtActvtnReq.GrpHdr.NbOfTxs), "1")
	require.Equal(t, string(*message.doc.CdtrPmtActvtnReq.GrpHdr.InitgPty.Nm), "Corporation A")
	require.Equal(t, string(*message.doc.CdtrPmtActvtnReq.GrpHdr.InitgPty.PstlAdr.PstCd), "85268")
	require.Equal(t, string(message.doc.CdtrPmtActvtnReq.PmtInf.PmtInfId), "20250310B1QDRCQR000601")
	require.Equal(t, string(message.doc.CdtrPmtActvtnReq.PmtInf.PmtMtd), "TRF")
	require.Equal(t, string(*message.doc.CdtrPmtActvtnReq.PmtInf.Dbtr.PstlAdr.StrtNm), "Avenue of the Fountains")
	require.Equal(t, string(message.doc.CdtrPmtActvtnReq.PmtInf.DbtrAcct.Id.Othr.Id), "92315266453")
	require.Equal(t, string(message.doc.CdtrPmtActvtnReq.PmtInf.CdtTrfTx.PmtId.EndToEndId), "Scenario1EndToEndId001")
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"
const INVALID_TRCOUNT string = "123456789012345"
const INVALID_MESSAGE_ID string = "12345678abcdEFGH123456"
const INVALID_OTHER_ID string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const INVALID_BUILD_NUM string = "12345678901234567"
const INVALID_POSTAL_CODE string = "12345678901234567"
const INVALID_COUNTRY_CODE string = "12345678"

func TestDrawdownRequestValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"MessageId",
			Message{data: MessageModel{MessageId: "Unknown data"}},
			"error occur at MessageId: Unknown data fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"NumberofTransaction",
			Message{data: MessageModel{NumberofTransaction: "Unknown data"}},
			"error occur at NumberofTransaction: Unknown data fails enumeration validation",
		},
		{
			"InitiatingParty - BuildingNumber",
			Message{data: MessageModel{InitiatingParty: model.PartyIdentify{
				Name: "Corporation A",
				Address: model.PostalAddress{
					StreetName:     "Avenue of the Fountains",
					BuildingNumber: "0122345678901234567890",
					RoomNumber:     "Suite D110",
					PostalCode:     "85268",
					TownName:       "Fountain Hills",
					Subdivision:    "AZ",
					Country:        "US",
				},
			}}},
			"error occur at InitiatingParty.Address.BuildingNumber: 0122345678901234567890 fails validation with length 22 <= required maxLength 16",
		},
		{
			"InitiatingParty - Country",
			Message{data: MessageModel{InitiatingParty: model.PartyIdentify{
				Name: "Corporation A",
				Address: model.PostalAddress{
					StreetName:     "Avenue of the Fountains",
					BuildingNumber: "167565",
					RoomNumber:     "Suite D110",
					PostalCode:     "85268",
					TownName:       "Fountain Hills",
					Subdivision:    "AZ",
					Country:        "Space World",
				},
			}}},
			"error occur at InitiatingParty.Address.Country: Space World fails validation with pattern [A-Z]{2,2}",
		},
		{
			"PaymentInfoId",
			Message{data: MessageModel{PaymentInfoId: "01223456789012345678900122345678901234567890"}},
			"error occur at PaymentInfoId: 01223456789012345678900122345678901234567890 fails validation with length 44 <= required maxLength 35",
		},
		{
			"PaymentMethod",
			Message{data: MessageModel{PaymentMethod: PaymentMethod(INVALID_COUNT)}},
			"error occur at PaymentMethod: UNKNOWN fails enumeration validation",
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
func TestDrawdowns_Scenario1_Step1_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000601"
	message.data.CreateDatetime = time.Now()
	message.data.NumberofTransaction = "1"
	message.data.InitiatingParty = model.PartyIdentify{
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
	}
	message.data.PaymentInfoId = "20250310B1QDRCQR000601"
	message.data.PaymentMethod = CreditTransform
	message.data.RequestedExecutDate = model.FromTime(time.Now())
	message.data.Debtor = model.PartyIdentify{
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
	}
	message.data.DebtorAccountOtherId = "92315266453"
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditTransTransaction = CreditTransferTransaction{
		PaymentInstructionId: "Scenario01Step1InstrId001",
		PaymentEndToEndId:    "Scenario1EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		PayCategoryType:      IntraCompanyPayment,
		PayRequestType:       DrawDownRequestCredit,
		Amount: model.CurrencyAndAmount{
			Amount:   6000000.00,
			Currency: "USD",
		},
		ChargeBearer: ChargeBearerSLEV,
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		Creditor: model.PartyIdentify{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain HIlls",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CrediorAccountOtherId: "5647772655",
		RemittanceInformation: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario1_Step1_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step1_pain.013")
	genterated := filepath.Join("generated", "Drawdowns_Scenario1_Step1_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario2_Step1_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000611"
	message.data.CreateDatetime = time.Now()
	message.data.NumberofTransaction = "1"
	message.data.InitiatingParty = model.PartyIdentify{
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
	}
	message.data.PaymentInfoId = "20250310B1QDRCQR000611"
	message.data.PaymentMethod = CreditTransform
	message.data.RequestedExecutDate = model.FromTime(time.Now())
	message.data.Debtor = model.PartyIdentify{
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
	}
	message.data.DebtorAccountOtherId = "92315266453"
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditTransTransaction = CreditTransferTransaction{
		PaymentInstructionId: "Scenario02Step1InstrId001",
		PaymentEndToEndId:    "Scenario2EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f068",
		PayCategoryType:      IntraCompanyPayment,
		PayRequestType:       DrawDownRequestCredit,
		Amount: model.CurrencyAndAmount{
			Amount:   6000000.00,
			Currency: "USD",
		},
		ChargeBearer: ChargeBearerSLEV,
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		Creditor: model.PartyIdentify{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain HIlls",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CrediorAccountOtherId: "5647772655",
		RemittanceInformation: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario2_Step1_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario2_Step1_pain.013")
	genterated := filepath.Join("generated", "Drawdowns_Scenario2_Step1_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario3_Step1_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000621"
	message.data.CreateDatetime = time.Now()
	message.data.NumberofTransaction = "1"
	message.data.InitiatingParty = model.PartyIdentify{
		Name: "Bank A",
		Address: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.PaymentInfoId = "20250310B1QDRCQR000621"
	message.data.PaymentMethod = CreditTransform
	message.data.RequestedExecutDate = model.FromTime(time.Now())
	message.data.Debtor = model.PartyIdentify{
		Name: "Bank Bb",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "52",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditTransTransaction = CreditTransferTransaction{
		PaymentInstructionId: "Scenario03Step1InstrId001",
		PaymentEndToEndId:    "Scenario3EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f070",
		PayRequestType:       DrawDownRequestDebit,
		Amount: model.CurrencyAndAmount{
			Amount:   1000000000.00,
			Currency: "USD",
		},
		ChargeBearer: ChargeBearerSLEV,
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		Creditor: model.PartyIdentify{
			Name: "Bank A",
			Address: model.PostalAddress{
				StreetName:     "Avenue A",
				BuildingNumber: "66",
				PostalCode:     "60532",
				TownName:       "Lisle",
				Subdivision:    "IL",
				Country:        "US",
			},
		},
		RemittanceInformation: "3rd repayment loan with reference ABCD432Z",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario3_Step1_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step1_pain.013")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step1_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step1_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000681"
	message.data.CreateDatetime = time.Now()
	message.data.NumberofTransaction = "1"
	message.data.InitiatingParty = model.PartyIdentify{
		Name: "Bank Aa",
		Address: model.PostalAddress{
			StreetName:     "Main Road",
			BuildingNumber: "3",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.PaymentInfoId = "20250310B1QDRCQR000681"
	message.data.PaymentMethod = CreditTransform
	message.data.RequestedExecutDate = model.FromTime(time.Now())
	message.data.Debtor = model.PartyIdentify{
		Name: "Bank Bb",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
			BuildingNumber: "52",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditTransTransaction = CreditTransferTransaction{
		PaymentInstructionId: "Scenario04Step1InstrId001",
		PaymentEndToEndId:    "Scenario4EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f070",
		PayRequestType:       DrawDownRequestDebit,
		Amount: model.CurrencyAndAmount{
			Amount:   1500000000.00,
			Currency: "USD",
		},
		ChargeBearer: ChargeBearerSLEV,
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		Creditor: model.PartyIdentify{
			Name: "Bank Aa",
			Address: model.PostalAddress{
				StreetName:     "Main Road",
				BuildingNumber: "3",
				PostalCode:     "60532",
				TownName:       "Lisle",
				Subdivision:    "IL",
				Country:        "US",
			},
		},
		RemittanceInformation: "Additional margin call for 03/10/2025 with reference XYZDF22.",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario4_Step1_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step1_pain.013")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step1_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario5_Step1_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000631"
	message.data.CreateDatetime = time.Now()
	message.data.NumberofTransaction = "1"
	message.data.InitiatingParty = model.PartyIdentify{
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
	}
	message.data.PaymentInfoId = "20250310B1QDRCQR000631"
	message.data.PaymentMethod = CreditTransform
	message.data.RequestedExecutDate = model.FromTime(time.Now())
	message.data.Debtor = model.PartyIdentify{
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
	}
	message.data.DebtorAccountOtherId = "9231526645"
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditTransTransaction = CreditTransferTransaction{
		PaymentInstructionId: "Scenario04Step1InstrId001",
		PaymentEndToEndId:    "Scenario4EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f258",
		PayRequestType:       DrawDownRequestCredit,
		PayCategoryType:      IntraCompanyPayment,
		Amount: model.CurrencyAndAmount{
			Amount:   6000000.00,
			Currency: "USD",
		},
		ChargeBearer: ChargeBearerSLEV,
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		Creditor: model.PartyIdentify{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain HIlls",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CrediorAccountOtherId: "5647772655",
		RemittanceInformation: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario5_Step1_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario5_Step1_pain.013")
	genterated := filepath.Join("generated", "Drawdowns_Scenario5_Step1_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario1_Step1_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000711"
	message.data.CreateDatetime = time.Now()
	message.data.NumberofTransaction = "1"
	message.data.InitiatingParty = model.PartyIdentify{
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
	}
	message.data.PaymentInfoId = "20250310B1QDRCQR000711"
	message.data.PaymentMethod = CreditTransform
	message.data.RequestedExecutDate = model.FromTime(time.Now())
	message.data.Debtor = model.PartyIdentify{
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
	}
	message.data.DebtorAccountOtherId = "5647772655"
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditTransTransaction = CreditTransferTransaction{
		PaymentInstructionId: "Scenario01InstrId001",
		PaymentEndToEndId:    "Scenario01Step1EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f078",
		PayRequestType:       DrawDownRequestCredit,
		Amount: model.CurrencyAndAmount{
			Amount:   60000.00,
			Currency: "USD",
		},
		ChargeBearer: ChargeBearerSLEV,
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		Creditor: model.PartyIdentify{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain HIlls",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CrediorAccountOtherId: "5647772655",
		document: RemittanceDocument{
			CodeOrProprietary: model.CodeCINV,
			Number:            "INV12345",
			RelatedDate:       model.FromTime(time.Now()),
		},
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step1_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step1_pain.013")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step1_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario1_Step1b_pain_CreateXML(t *testing.T) {
	var message, vErr = NewMessage("")
	require.NoError(t, vErr)
	message.data.MessageId = "20250310B1QDRCQR000711"
	message.data.CreateDatetime = time.Now()
	message.data.NumberofTransaction = "1"
	message.data.InitiatingParty = model.PartyIdentify{
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
	}
	message.data.PaymentInfoId = "20250310B1QDRCQR000711"
	message.data.PaymentMethod = CreditTransform
	message.data.RequestedExecutDate = model.FromTime(time.Now())
	message.data.Debtor = model.PartyIdentify{
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
	}
	message.data.DebtorAccountOtherId = "5647772655"
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditTransTransaction = CreditTransferTransaction{
		PaymentInstructionId: "Scenario01InstrId001",
		PaymentEndToEndId:    "Scenario01Step1EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f078",
		PayRequestType:       DrawDownRequestCredit,
		Amount: model.CurrencyAndAmount{
			Amount:   60000.00,
			Currency: "USD",
		},
		ChargeBearer: ChargeBearerSLEV,
		CreditorAgent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		Creditor: model.PartyIdentify{
			Name: "Corporation A",
			Address: model.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain HIlls",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CrediorAccountOtherId: "5647772655",
		document: RemittanceDocument{
			CodeOrProprietary: model.CodeCINV,
			Number:            "INV12345",
			RelatedDate:       model.FromTime(time.Now()),
		},
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step1b_pain.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step1b_pain.013")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step1b_pain.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
