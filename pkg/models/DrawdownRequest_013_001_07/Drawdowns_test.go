package DrawdownRequest_013_001_07

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario1_Step1_pain_CreateXML(t *testing.T) {
	var message = NewPain013Message()
	message.model.MessageId = "20250310B1QDRCQR000601"
	message.model.CreateDatetime = time.Now()
	message.model.NumberofTransaction = "1"
	message.model.InitiatingParty = model.PartyIdentify{
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
	message.model.PaymentInfoId = "20250310B1QDRCQR000601"
	message.model.PaymentMethod = CreditTransform
	message.model.RequestedExecutDate = civil.DateOf(time.Now())
	message.model.Debtor = model.PartyIdentify{
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
	message.model.DebtorAccountOtherId = "92315266453"
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditTransTransaction = CreditTransferTransaction{
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
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario1_Step1_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario2_Step1_pain_CreateXML(t *testing.T) {
	var message = NewPain013Message()
	message.model.MessageId = "20250310B1QDRCQR000611"
	message.model.CreateDatetime = time.Now()
	message.model.NumberofTransaction = "1"
	message.model.InitiatingParty = model.PartyIdentify{
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
	message.model.PaymentInfoId = "20250310B1QDRCQR000611"
	message.model.PaymentMethod = CreditTransform
	message.model.RequestedExecutDate = civil.DateOf(time.Now())
	message.model.Debtor = model.PartyIdentify{
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
	message.model.DebtorAccountOtherId = "92315266453"
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditTransTransaction = CreditTransferTransaction{
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
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario2_Step1_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario3_Step1_pain_CreateXML(t *testing.T) {
	var message = NewPain013Message()
	message.model.MessageId = "20250310B1QDRCQR000621"
	message.model.CreateDatetime = time.Now()
	message.model.NumberofTransaction = "1"
	message.model.InitiatingParty = model.PartyIdentify{
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
	message.model.PaymentInfoId = "20250310B1QDRCQR000621"
	message.model.PaymentMethod = CreditTransform
	message.model.RequestedExecutDate = civil.DateOf(time.Now())
	message.model.Debtor = model.PartyIdentify{
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
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditTransTransaction = CreditTransferTransaction{
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
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario3_Step1_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario4_Step1_pain_CreateXML(t *testing.T) {
	var message = NewPain013Message()
	message.model.MessageId = "20250310B1QDRCQR000681"
	message.model.CreateDatetime = time.Now()
	message.model.NumberofTransaction = "1"
	message.model.InitiatingParty = model.PartyIdentify{
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
	message.model.PaymentInfoId = "20250310B1QDRCQR000681"
	message.model.PaymentMethod = CreditTransform
	message.model.RequestedExecutDate = civil.DateOf(time.Now())
	message.model.Debtor = model.PartyIdentify{
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
	message.model.DebtorAccountOtherId = "92315266453"
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditTransTransaction = CreditTransferTransaction{
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
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario4_Step1_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario5_Step1_pain_CreateXML(t *testing.T) {
	var message = NewPain013Message()
	message.model.MessageId = "20250310B1QDRCQR000631"
	message.model.CreateDatetime = time.Now()
	message.model.NumberofTransaction = "1"
	message.model.InitiatingParty = model.PartyIdentify{
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
	message.model.PaymentInfoId = "20250310B1QDRCQR000631"
	message.model.PaymentMethod = CreditTransform
	message.model.RequestedExecutDate = civil.DateOf(time.Now())
	message.model.Debtor = model.PartyIdentify{
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
	message.model.DebtorAccountOtherId = "9231526645"
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditTransTransaction = CreditTransferTransaction{
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
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario5_Step1_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step1_pain_CreateXML(t *testing.T) {
	var message = NewPain013Message()
	message.model.MessageId = "20250310B1QDRCQR000711"
	message.model.CreateDatetime = time.Now()
	message.model.NumberofTransaction = "1"
	message.model.InitiatingParty = model.PartyIdentify{
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
	message.model.PaymentInfoId = "20250310B1QDRCQR000711"
	message.model.PaymentMethod = CreditTransform
	message.model.RequestedExecutDate = civil.DateOf(time.Now())
	message.model.Debtor = model.PartyIdentify{
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
	message.model.DebtorAccountOtherId = "5647772655"
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditTransTransaction = CreditTransferTransaction{
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
			CodeOrProprietary: CodeCINV,
			Number:            "INV12345",
			RelatedDate:       civil.DateOf(time.Now()),
		},
	}

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step1_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step1b_pain_CreateXML(t *testing.T) {
	var message = NewPain013Message()
	message.model.MessageId = "20250310B1QDRCQR000711"
	message.model.CreateDatetime = time.Now()
	message.model.NumberofTransaction = "1"
	message.model.InitiatingParty = model.PartyIdentify{
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
	message.model.PaymentInfoId = "20250310B1QDRCQR000711"
	message.model.PaymentMethod = CreditTransform
	message.model.RequestedExecutDate = civil.DateOf(time.Now())
	message.model.Debtor = model.PartyIdentify{
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
	message.model.DebtorAccountOtherId = "5647772655"
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditTransTransaction = CreditTransferTransaction{
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
			CodeOrProprietary: CodeCINV,
			Number:            "INV12345",
			RelatedDate:       civil.DateOf(time.Now()),
		},
	}

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step1b_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
