package DrawdownRequest

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

var DrawdownRequestsample1XML = filepath.Join("swiftSample", "Drawdowns_Scenario1_Step1_pain.013")
var DrawdownRequestsample2XML = filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step1_pain.013")

func TestDocumentElementToModelOne(t *testing.T) {
	var xmlData, err = models.ReadXMLFile(DrawdownRequestsample1XML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PaymentUniqueId, "8a562c67-ca16-48ba-b074-65581be6f066")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
}

func TestDocumentElementToModelTwo(t *testing.T) {
	var xmlData, err = models.ReadXMLFile(DrawdownRequestsample2XML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)
}
