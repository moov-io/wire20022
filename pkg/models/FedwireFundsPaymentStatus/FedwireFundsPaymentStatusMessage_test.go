package FedwireFundsPaymentStatus

import (
	"encoding/xml"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestCustomerCreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP31000001"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000001"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.AcceptanceDateTime = time.Now()
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("CustomerCreditTransfer_Scenario1_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "FDWA1B2C3D4E5F6G7H8I9J10K11L12M0"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000002"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.Rejected
	message.data.StatusReasonInformation = "E433"
	message.data.ReasonAdditionalInfo = "The routing number of the Instructed Agent is not permissible to receive Fedwire Funds transaction."
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("CustomerCreditTransfer_Scenario2_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario3_Step3_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP31000001"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000001"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("CustomerCreditTransfer_Scenario3_Step3_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP31000002"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000004"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("CustomerCreditTransfer_Scenario4_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP31000003"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000005"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("CustomerCreditTransfer_Scenario5_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP62000501"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000501"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FICreditTransfer_Scenario1_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP62000502"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000502"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FICreditTransfer_Scenario2_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario3_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP62000503"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000503"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FICreditTransfer_Scenario3_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP62000504"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000504"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FICreditTransfer_Scenario4_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP62000505"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310QMGFNP62000505"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FICreditTransfer_Scenario5_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario6_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310QMGFNP62000506"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000506"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FICreditTransfer_Scenario6_Step2_pacs.xml", xmlData)
	require.NoError(t, err)
}
