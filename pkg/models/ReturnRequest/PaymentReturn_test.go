package ReturnRequest

import (
	"encoding/xml"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsAcknowledgement_Scenario2_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000722"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.data.Creator = model.Agent{
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
	message.data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.data.CancellationReason = Reason{
		Originator:     "Corporation A",
		Reason:         "DUPL",
		AdditionalInfo: "Order cancelled. Ref:20250310B1QDRCQR000721",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step2_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario2_Step2b_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000722"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.data.Creator = model.Agent{
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
	message.data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.data.CancellationReason = Reason{
		Originator:     "Corporation A",
		Reason:         "DUPL",
		AdditionalInfo: "Order cancelled. Ref:20250310B1QDRCQR000721",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step2b_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestInvestigations_Scenario2_Step4_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000912"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.data.Creator = model.Agent{
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
	message.data.OriginalMessageId = "20250310B1QDRCQR000002"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   510000.74,
		Currency: "USD",
	}
	message.data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.data.CancellationReason = Reason{
		Originator:     "Corporation A",
		Reason:         "DUPL",
		AdditionalInfo: "Payment is a duplicate. Please return payment.",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Investigations_Scenario2_Step4_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario1_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000401"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.data.Creator = model.Agent{
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
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1510000.74,
		Currency: "USD",
	}
	message.data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.data.CancellationReason = Reason{
		Originator:     "Corporation A",
		Reason:         "DUPL",
		AdditionalInfo: "Order cancelled. Ref:20250310B1QDRCQR000400",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("PaymentReturn_Scenario1_Step2_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario2_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000421"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.CaseId = "20250310011104238Sc02Step1MsgIdSVNR"
	message.data.Creator = model.Agent{
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
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario02InstrId001"
	message.data.OriginalEndToEndId = "Scenario02EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1234578.88,
		Currency: "USD",
	}
	message.data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.data.CancellationReason = Reason{
		Originator: "Corporation A",
		Reason:     "SVNR",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("PaymentReturn_Scenario2_Step2_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario3_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000431"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.CaseId = "20250310011104238Sc02Step1MsgIdSVNR"
	message.data.Creator = model.Agent{
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
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario02InstrId001"
	message.data.OriginalEndToEndId = "Scenario02EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1234578.88,
		Currency: "USD",
	}
	message.data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.data.CancellationReason = Reason{
		Originator: "Corporation C",
		Reason:     "SVNR",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("PaymentReturn_Scenario3_Step2_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestPaymentReturn_Scenario5_Step2_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000431"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.CaseId = "20250310011104238Sc02Step1MsgIdSVNR"
	message.data.Creator = model.Agent{
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
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario02InstrId001"
	message.data.OriginalEndToEndId = "Scenario02EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   1234578.88,
		Currency: "USD",
	}
	message.data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.data.CancellationReason = Reason{
		Originator: "Corporation C",
		Reason:     "SVNR",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("PaymentReturn_Scenario5_Step2_camt.xml", xmlData)
	require.NoError(t, err)
}
