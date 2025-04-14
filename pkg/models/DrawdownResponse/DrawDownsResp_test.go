package DrawdownResponse

import (
	"encoding/xml"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario1_Step2_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000602"
	message.data.CreateDatetime = time.Now()
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
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000601"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000601"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01Step1InstrId001",
		OriginalEndToEndId:    "Scenario1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario1_Step2_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario2_Step2_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000612"
	message.data.CreateDatetime = time.Now()
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
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000611"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000611"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario02Step1InstrId001",
		OriginalEndToEndId:    "Scenario2EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.Rejected,
		StatusReasonInfoCode:  InsufficientFunds,
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario2_Step2_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario3_Step2_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000622"
	message.data.CreateDatetime = time.Now()
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
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000621"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000621"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario03Step1InstrId001",
		OriginalEndToEndId:    "Scenario3EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario3_Step2_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario4_Step2_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000682"
	message.data.CreateDatetime = time.Now()
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
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000681"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000681"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario04Step1InstrId001",
		OriginalEndToEndId:    "Scenario4EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario4_Step2_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario5_Step2_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000632"
	message.data.CreateDatetime = time.Now()
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
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000631"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000631"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario04Step1InstrId001",
		OriginalEndToEndId:    "Scenario4EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.TransPending,
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario5_Step2_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000712"
	message.data.CreateDatetime = time.Now()
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
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000711"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000711"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01InstrId001",
		OriginalEndToEndId:    "Scenario01Step1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step2_pain.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2b_pain_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000712"
	message.data.CreateDatetime = time.Now()
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
	message.data.DebtorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.CreditorAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.OriginalMessageId = "20250310B1QDRCQR000711"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalPaymentInfoId = "20250310B1QDRCQR000711"
	message.data.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01InstrId001",
		OriginalEndToEndId:    "Scenario01Step1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		TransactionStatus:     model.AcceptedTechnicalValidation,
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step2b_pain.xml", xmlData)
	require.NoError(t, err)
}
