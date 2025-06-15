package FedwireFundsPaymentStatus

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wadearnold/wire20022/pkg/models"
)

func TestDocumentToModel0(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step2_pacs.002")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
}
