package ReturnRequestResponse

import (
	"encoding/xml"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsAcknowledgement_Scenario2_Step3_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000723"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc01Step1MsgIdDUPL"
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
	message.data.Status = ReturnRequestAccepted
	message.data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step3_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario2_Step3b_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000723"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc01Step1MsgIdDUPL"
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
	message.data.Status = ReturnRequestAccepted
	message.data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step3b_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestPaymentreturn_Scenario1_Step3_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000402"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc01Step1MsgIdDUPL"
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
	message.data.Status = ReturnRequestAccepted
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Paymentreturn_Scenario1_Step3_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestPaymentreturn_Scenario2_Step3_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000422"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc02Step1MsgIdSVNR"
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
	message.data.Status = ReturnRequestRejected
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.CancellationStatusReasonInfo = Reason{
		Reason:         "NARR",
		AdditionalInfo: "Corporation B delivered goods and services are in-line with client’s order.",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Paymentreturn_Scenario2_Step3_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestPaymentreturn_Scenario3_Step3_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.AssignmentId = "20250310B1QDRCQR000432"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc03Step1MsgIdSVNR"
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
	message.data.Status = ReturnRequestRejected
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.CancellationStatusReasonInfo = Reason{
		Reason:         "NARR",
		AdditionalInfo: "As agreed, partial refund of 20% will be paid for service shortcomings.",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	model.WriteXMLTo("Paymentreturn_Scenario3_Step3_camt.xml", xmlData)
	require.NoError(t, err)
}
