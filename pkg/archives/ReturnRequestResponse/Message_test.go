package ReturnRequestResponse

import (
	"path/filepath"
	"testing"

	Archive "github.com/moov-io/wire20022/pkg/archives"
	"github.com/stretchr/testify/require"
)

func TestDocumentToMode09(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "Paymentreturn_Scenario2_Step3_camt.029")
	var xmlData, err = Archive.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, Archive.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "NARR")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")
}
