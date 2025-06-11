package CustomerCreditTransfer

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDocumentToModel08(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step1_pacs.008")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")

	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.UniqueEndToEndTransactionRef, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.RoomNumber, "Suite D110")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.Floor, "33")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
}

func TestDocumentToModel08ChargeInfo(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "CustomerCreditTransfer_Variation5_pacs.008")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")

	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
}
func TestDocumentToModel08TaxDetail(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario4_Step1_pacs.008")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxId, "123456789")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxTypeCode, "09455")
	require.NotNil(t, model.RemittanceInfor.TaxDetail.TaxPeriodYear)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxperiodTimeFrame, "MM04")
}
func TestDocumentToModel08ElectronicAddress(t *testing.T) {
	var sampleXML = filepath.Join("swiftSample", "CustomerCreditTransfer_Variation2_pacs.008")
	var xmlData, err = models.ReadXMLFile(sampleXML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")
}
