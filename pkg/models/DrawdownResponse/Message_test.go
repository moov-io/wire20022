package DrawdownResponse

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

var DrawdownResponsesample1XML = filepath.Join("swiftSample", "Drawdowns_Scenario1_Step2_pain.014")

func TestDocumentElementToModelOne(t *testing.T) {
	var xmlData, err = models.ReadXMLFile(DrawdownResponsesample1XML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := ParseXML(xmlData)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000602")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000601")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalUniqueId, "8a562c67-ca16-48ba-b074-65581be6f066")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, models.AcceptedTechnicalValidation)
}
