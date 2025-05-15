package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/FinancialInstitutionCreditTransfer"
	"github.com/stretchr/testify/require"
)
var FinancialInstitutionCreditTransferxmlFile = "../models/FinancialInstitutionCreditTransfer/swiftSample/Drawdowns_Scenario3_Step3_pacs.009"
func TestFinancialInstitutionCreditTransferParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(FinancialInstitutionCreditTransferxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &FinancialInstitutionCreditTransfer.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*FinancialInstitutionCreditTransfer.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310B1QDRCQR000623")
	}
}

func TestFinancialInstitutionCreditTransferGenerateXML(t *testing.T) {
	dataModel := FinancialInstitutionCreditTransferDataModel()
	xmlData, err := GenerateXML(&dataModel, &FinancialInstitutionCreditTransfer.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("FinancialInstitutionCreditTransfer_test.xml", xmlData)
	require.NoError(t, err)
}

func TestFinancialInstitutionCreditTransferRequireFieldCheck(t *testing.T) {
	dataModel := FinancialInstitutionCreditTransferDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &FinancialInstitutionCreditTransfer.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestFinancialInstitutionCreditTransferXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(FinancialInstitutionCreditTransferxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &FinancialInstitutionCreditTransfer.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestFinancialInstitutionCreditTransferAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&FinancialInstitutionCreditTransfer.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*FinancialInstitutionCreditTransfer.MessageHelper); ok {
		require.Equal(t, helper.SettlementMethod.Title, "Settlement Method")
		require.Equal(t, helper.SettlementMethod.Type, "SettlementMethodType(SettlementCLRG, SettlementINDA, SettlementCOVE ...)")
		require.Equal(t, helper.SettlementMethod.Documentation, "Method used to settle the (batch of) payment instructions.")
	}
}

func FinancialInstitutionCreditTransferDataModel() FinancialInstitutionCreditTransfer.MessageModel {
	var message, _ = FinancialInstitutionCreditTransfer.NewMessage("")
	message.Data.MessageId = "20250310B1QDRCQR000506"
	message.Data.CreateDateTime = time.Now()
	message.Data.NumberOfTransactions = 1
	message.Data.SettlementMethod = model.SettlementCLRG
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.ClearingSystem = model.ClearingSysFDW
	message.Data.PaymentInstructionId = "Scenario06FIInstrId001"
	message.Data.PaymentEndToEndId = "Scenario06FIEtoEId001"
	message.Data.PaymentUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.Data.LocalInstrument = FinancialInstitutionCreditTransfer.CoreCoverPayment
	message.Data.InterbankSettlementDate = model.FromTime(time.Now())
	message.Data.InterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   179852.25,
		Currency: "USD",
	}
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
	}
	message.Data.Debtor = model.FiniancialInstitutionId{
		BusinessId: "BANZBEBB",
	}
	message.Data.DebtorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANCUS33",
	}
	message.Data.Creditor = model.FiniancialInstitutionId{
		BusinessId: "BANYBRRJ",
	}
	message.Data.CreditorAgent = model.FiniancialInstitutionId{
		BusinessId: "BANDUS33",
	}
	message.Data.UnderlyingCustomerCreditTransfer = FinancialInstitutionCreditTransfer.CreditTransferTransaction{
		Debtor: model.FiniancialInstitutionId{
			Name: "Corporation Z",
			Address: model.PostalAddress{
				StreetName:     "Avenue Moliere",
				BuildingNumber: "70",
				PostalCode:     "1180",
				TownName:       "Brussels",
				Country:        "BE",
			},
		},
		DebtorAccount: "BE34001216371411",
		DebtorAgent: model.FiniancialInstitutionId{
			BusinessId: "BANZBEBB",
		},
		CreditorAgent: model.FiniancialInstitutionId{
			BusinessId: "BANYBRRJ",
		},
		Creditor: model.FiniancialInstitutionId{
			Name: "Corporation Y",
			Address: model.PostalAddress{
				StreetName:     "Av. Lucio Costa",
				BuildingNumber: "5220",
				BuildingName:   "Barra da Tijuca",
				PostalCode:     "22630-012",
				TownName:       "Rio de Janeiro",
				Subdivision:    "RJ",
				Country:        "BR",
			},
		},
		CreditorAccount:       "BR9700360305000010009795493P1",
		RemittanceInformation: "Payment invoice 444563 dated 1st March 2025",
		InstructedAmount: model.CurrencyAndAmount{
			Amount:   17985234.25,
			Currency: "USD",
		},
	}

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
