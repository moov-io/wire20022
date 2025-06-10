package PaymentStatusRequest

import (
	"path/filepath"
	"testing"

	Archive "github.com/moov-io/wire20022/pkg/archives"
	"github.com/stretchr/testify/require"
)

func TestDocumentToMode03(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step2_pacs.028")
	var xmlData, err = Archive.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310Scenario03Step2MsgId001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.InstructingAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021151080")
}
