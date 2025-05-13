package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/ReturnRequest"
	"github.com/stretchr/testify/require"
)

func TestReturnRequestParseXMLFile(t *testing.T) {
	xmlFile := "../models/ReturnRequest/generated/FedwireFundsAcknowledgement_Scenario2_Step2_camt.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &ReturnRequest.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*ReturnRequest.MessageModel); ok {
		require.Equal(t, msgModel.AssignmentId, "20250310B1QDRCQR000722")
	}
}

func TestReturnRequestGenerateXML(t *testing.T) {
	dataModel := ReturnRequestDataModel()
	xmlData, err := GenerateXML(&dataModel, &ReturnRequest.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("ReturnRequest_test.xml", xmlData)
	require.NoError(t, err)
}

func TestReturnRequestRequireFieldCheck(t *testing.T) {
	dataModel := ReturnRequestDataModel()
	dataModel.AssignmentId = ""
	dataModel.CaseId = ""
	valid, err := RequireFieldCheck(&dataModel, &ReturnRequest.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: AssignmentId, CaseId")
}

func TestReturnRequestXMLValidation(t *testing.T) {
	xmlFile := "../models/ReturnRequest/swiftSample/FedwireFundsAcknowledgement_Scenario2_Step2_camt.056"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &ReturnRequest.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestReturnRequestAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&ReturnRequest.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*ReturnRequest.MessageHelper); ok {
		require.Equal(t, helper.OriginalMessageId.Title, "Original Message Id")
		require.Equal(t, helper.OriginalMessageId.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.OriginalMessageId.Documentation, "Point to point reference assigned by the original instructing party to unambiguously identify the original message.")
	}
}

func ReturnRequestDataModel() ReturnRequest.MessageModel {
	var message, _ = ReturnRequest.NewMessage("")
	message.Data.AssignmentId = "20250310B1QDRCQR000722"
	message.Data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.AssignmentCreateTime = time.Now()
	message.Data.CaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.Data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.Data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.Data.OriginalInterbankSettlementDate = model.FromTime(time.Now())
	message.Data.CancellationReason = ReturnRequest.Reason{
		Originator:     "Corporation A",
		Reason:         "DUPL",
		AdditionalInfo: "Order cancelled. Ref:20250310B1QDRCQR000721",
	}

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
