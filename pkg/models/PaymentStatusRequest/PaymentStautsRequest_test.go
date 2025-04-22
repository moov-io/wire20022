package PaymentStatusRequest

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestCustomerCreditTransfer_Scenario3_Step2_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310Scenario03Step2MsgId001"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000001"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step2_pacs.028")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario5_Step3_pacs_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250310Scenario04Step3MsgId001"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000631"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario04Step1InstrId001"
	message.data.OriginalEndToEndId = "Scenario4EndToEndId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f258"
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario5_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario5_Step3_pacs.028")
	genterated := filepath.Join("generated", "Drawdowns_Scenario5_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
