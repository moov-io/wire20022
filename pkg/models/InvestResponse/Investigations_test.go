package InvestigationResponse_camt_111_001_01

import (
	"encoding/xml"
	"testing"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestInvestigations_Scenario1_Step3_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000901"
	message.data.InvestigationStatus = "CLSD"
	message.data.InvestigationData = "Please correct Creditor Account number. It should be 567876543."
	message.data.InvestRequestMessageId = "20250310QMGFT015000901"
	message.data.InvestigationType = "UTAP"
	message.data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("Investigations_Scenario1_Step3_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestInvestigations_Scenario2_Step3_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000902"
	message.data.InvestigationStatus = "CLSD"
	message.data.InvestigationData = "Payment is a duplicate. Please consider VOID. Return request will follow."
	message.data.InvestRequestMessageId = "20250310QMGFT015000902"
	message.data.InvestigationType = "UTAP"
	message.data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("Investigations_Scenario2_Step3_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestInvestigations_Scenario3_Step3_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310B1QDRCQR000903"
	message.data.InvestigationStatus = "CLSD"
	message.data.InvestigationData = "Remittance information was sent separately. Email: AccountsReceivable@CorporationB.com"
	message.data.InvestRequestMessageId = "20250310QMGFT015000903"
	message.data.InvestigationType = "UTAP"
	message.data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("Investigations_Scenario3_Step3_camt.xml", xmlData)
	require.NoError(t, err)
}
