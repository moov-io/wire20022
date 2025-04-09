package FedwireFundsPaymentStatus_002_001_10

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestCustomerCreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP31000001"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000001"
	message.model.OriginalMessageNameId = "pacs.008.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.AcceptanceDateTime = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "FDWA1B2C3D4E5F6G7H8I9J10K11L12M0"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000002"
	message.model.OriginalMessageNameId = "pacs.008.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.model.TransactionStatus = model.Rejected
	message.model.StatusReasonInformation = "E433"
	message.model.ReasonAdditionalInfo = "The routing number of the Instructed Agent is not permissible to receive Fedwire Funds transaction."
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario3_Step3_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP31000001"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000001"
	message.model.OriginalMessageNameId = "pacs.008.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.AcceptanceDateTime = time.Now()
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step3_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP31000002"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000004"
	message.model.OriginalMessageNameId = "pacs.008.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP31000003"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000005"
	message.model.OriginalMessageNameId = "pacs.008.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP62000501"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000501"
	message.model.OriginalMessageNameId = "pacs.009.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.AcceptanceDateTime = time.Now()
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario1_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP62000502"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000502"
	message.model.OriginalMessageNameId = "pacs.009.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.AcceptanceDateTime = time.Now()
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario2_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario3_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP62000503"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000503"
	message.model.OriginalMessageNameId = "pacs.009.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.AcceptanceDateTime = time.Now()
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario3_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP62000504"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000504"
	message.model.OriginalMessageNameId = "pacs.009.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.AcceptanceDateTime = time.Now()
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario4_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP62000505"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310QMGFNP62000505"
	message.model.OriginalMessageNameId = "pacs.009.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.AcceptanceDateTime = time.Now()
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario5_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario6_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewPacs002Message()
	message.model.MessageId = "20250310QMGFNP62000506"
	message.model.CreatedDateTime = time.Now()
	message.model.OriginalMessageId = "20250310B1QDRCQR000506"
	message.model.OriginalMessageNameId = "pacs.009.001.08"
	message.model.OriginalMessageCreateTime = time.Now()
	message.model.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.model.TransactionStatus = model.AcceptedSettlementCompleted
	message.model.AcceptanceDateTime = time.Now()
	message.model.EffectiveInterbankSettlementDate = time.Now()
	message.model.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.model.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario6_Step2_pacs.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
