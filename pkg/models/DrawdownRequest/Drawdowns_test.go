package DrawdownRequest

import (
	"encoding/xml"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario1_Step1_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.data.RequestedExecutDate = model.FromTime(time.Time{})
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
		PayCategoryType:      DrawDownRequestCredit,
		PayRequestType:       IntraCompanyPayment,
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
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CrediorAccountOtherId: "5647772655",
		RemittanceInformation: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("Drawdowns_Scenario1_Step1_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario2_Step1_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.data.RequestedExecutDate = model.FromTime(time.Time{})
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
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		PayCategoryType:      DrawDownRequestCredit,
		PayRequestType:       IntraCompanyPayment,
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
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CrediorAccountOtherId: "5647772655",
		RemittanceInformation: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("Drawdowns_Scenario2_Step1_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario3_Step1_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000621"
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
	message.data.PaymentInfoId = "20250310B1QDRCQR000621"
	message.data.PaymentMethod = CreditTransform
	message.data.RequestedExecutDate = model.FromTime(time.Time{})
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
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		PayCategoryType:      DrawDownRequestDebit,
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
		CrediorAccountOtherId: "5647772655",
		RemittanceInformation: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("Drawdowns_Scenario3_Step1_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario4_Step1_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.data.RequestedExecutDate = model.FromTime(time.Time{})
	message.data.Debtor = model.PartyIdentify{
		Name: "Bank Bb",
		Address: model.PostalAddress{
			StreetName:     "Avenue C",
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
		PaymentInstructionId: "Scenario04Step1InstrId001",
		PaymentEndToEndId:    "Scenario4EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		PayCategoryType:      DrawDownRequestCredit,
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
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		RemittanceInformation: "Additional margin call for 03/10/2025 with reference XYZDF22.",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("Drawdowns_Scenario4_Step1_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario5_Step1_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.data.RequestedExecutDate = model.FromTime(time.Time{})
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
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		PayCategoryType:      DrawDownRequestCredit,
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
			Name: "Bank Aa",
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
		RemittanceInformation: "Additional margin call for 03/10/2025 with reference XYZDF22.",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("Drawdowns_Scenario5_Step1_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step1_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.data.RequestedExecutDate = model.FromTime(time.Time{})
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
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		PayCategoryType:      DrawDownRequestCredit,
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
			Name: "Bank Aa",
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
		document: RemittanceDocument{
			CodeOrProprietary: model.CodeCINV,
			Number:            "INV12345",
			RelatedDate:       model.FromTime(time.Time{}),
		},
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step1_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step1b_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.data.RequestedExecutDate = model.FromTime(time.Time{})
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
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		PayCategoryType:      DrawDownRequestCredit,
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
			Name: "Bank A",
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
		document: RemittanceDocument{
			CodeOrProprietary: model.CodeCINV,
			Number:            "INV12345",
			RelatedDate:       model.FromTime(time.Time{}),
		},
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step1b_pain.xml", xmlData)
	require.NoError(t, err)
}
