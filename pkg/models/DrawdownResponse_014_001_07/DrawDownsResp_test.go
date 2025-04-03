package DrawdownResponse_014_001_07

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario1_Step2_pain_CreateXML(t *testing.T) {
	var message = NewPain014Message()
	message.model.MessageId = "20250310B1QDRCQR000602"
	message.model.CreateDatetime = time.Now()
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
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.model.OriginalMessageId = "20250310B1QDRCQR000601"
	message.model.OriginalMessageNameId = "pain.013.001.07"
	message.model.OriginalCreationDateTime = time.Now()
	message.model.OriginalPaymentInfoId = "20250310B1QDRCQR000601"
	message.model.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01Step1InstrId001",
		OriginalEndToEndId:    "Scenario1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario1_Step2_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario2_Step2_pain_CreateXML(t *testing.T) {
	var message = NewPain014Message()
	message.model.MessageId = "20250310B1QDRCQR000612"
	message.model.CreateDatetime = time.Now()
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
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.model.OriginalMessageId = "20250310B1QDRCQR000611"
	message.model.OriginalMessageNameId = "pain.013.001.07"
	message.model.OriginalCreationDateTime = time.Now()
	message.model.OriginalPaymentInfoId = "20250310B1QDRCQR000611"
	message.model.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario02Step1InstrId001",
		OriginalEndToEndId:    "Scenario2EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.Rejected,
		StatusReasonInfoCode:  InsufficientFunds,
	}

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario2_Step2_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario3_Step2_pain_CreateXML(t *testing.T) {
	var message = NewPain014Message()
	message.model.MessageId = "20250310B1QDRCQR000622"
	message.model.CreateDatetime = time.Now()
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
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.model.OriginalMessageId = "20250310B1QDRCQR000621"
	message.model.OriginalMessageNameId = "pain.013.001.07"
	message.model.OriginalCreationDateTime = time.Now()
	message.model.OriginalPaymentInfoId = "20250310B1QDRCQR000621"
	message.model.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario03Step1InstrId001",
		OriginalEndToEndId:    "Scenario3EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario3_Step2_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario4_Step2_pain_CreateXML(t *testing.T) {
	var message = NewPain014Message()
	message.model.MessageId = "20250310B1QDRCQR000682"
	message.model.CreateDatetime = time.Now()
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
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.model.OriginalMessageId = "20250310B1QDRCQR000681"
	message.model.OriginalMessageNameId = "pain.013.001.07"
	message.model.OriginalCreationDateTime = time.Now()
	message.model.OriginalPaymentInfoId = "20250310B1QDRCQR000681"
	message.model.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario04Step1InstrId001",
		OriginalEndToEndId:    "Scenario4EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario4_Step2_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario5_Step2_pain_CreateXML(t *testing.T) {
	var message = NewPain014Message()
	message.model.MessageId = "20250310B1QDRCQR000632"
	message.model.CreateDatetime = time.Now()
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
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.model.OriginalMessageId = "20250310B1QDRCQR000631"
	message.model.OriginalMessageNameId = "pain.013.001.07"
	message.model.OriginalCreationDateTime = time.Now()
	message.model.OriginalPaymentInfoId = "20250310B1QDRCQR000631"
	message.model.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario04Step1InstrId001",
		OriginalEndToEndId:    "Scenario4EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.TransPending,
	}

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario5_Step2_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2_pain_CreateXML(t *testing.T) {
	var message = NewPain014Message()
	message.model.MessageId = "20250310B1QDRCQR000712"
	message.model.CreateDatetime = time.Now()
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
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.model.OriginalMessageId = "20250310B1QDRCQR000711"
	message.model.OriginalMessageNameId = "pain.013.001.07"
	message.model.OriginalCreationDateTime = time.Now()
	message.model.OriginalPaymentInfoId = "20250310B1QDRCQR000711"
	message.model.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01InstrId001",
		OriginalEndToEndId:    "Scenario01Step1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step2_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2b_pain_CreateXML(t *testing.T) {
	var message = NewPain014Message()
	message.model.MessageId = "20250310B1QDRCQR000712"
	message.model.CreateDatetime = time.Now()
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
	message.model.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.model.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.model.OriginalMessageId = "20250310B1QDRCQR000711"
	message.model.OriginalMessageNameId = "pain.013.001.07"
	message.model.OriginalCreationDateTime = time.Now()
	message.model.OriginalPaymentInfoId = "20250310B1QDRCQR000711"
	message.model.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01InstrId001",
		OriginalEndToEndId:    "Scenario01Step1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step2b_pain.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
