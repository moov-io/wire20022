package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse"
	"github.com/stretchr/testify/require"
)
var ReturnRequestResponsexmlFile = "../models/ReturnRequestResponse/swiftSample/FedwireFundsAcknowledgement_Scenario2_Step3_camt.029"
func TestReturnRequestResponseParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(ReturnRequestResponsexmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &ReturnRequestResponse.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*ReturnRequestResponse.MessageModel); ok {
		require.Equal(t, msgModel.AssignmentId, "20250310B1QDRCQR000723")
	}
}

func TestReturnRequestResponseGenerateXML(t *testing.T) {
	dataModel := ReturnRequestResponseDataModel()
	xmlData, err := GenerateXML(&dataModel, &ReturnRequestResponse.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("ReturnRequestResponse_test.xml", xmlData)
	require.NoError(t, err)
}

func TestReturnRequestResponseRequireFieldCheck(t *testing.T) {
	dataModel := ReturnRequestResponseDataModel()
	dataModel.AssignmentId = ""
	dataModel.OriginalMessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &ReturnRequestResponse.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: AssignmentId, OriginalMessageId")
}

func TestReturnRequestResponseXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(ReturnRequestResponsexmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &ReturnRequestResponse.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestReturnRequestResponseAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&ReturnRequestResponse.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*ReturnRequestResponse.MessageHelper); ok {
		require.Equal(t, helper.OriginalInstructionId.Title, "Original Instruction Id")
		require.Equal(t, helper.OriginalInstructionId.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.OriginalInstructionId.Documentation, "Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.")
	}
}

func ReturnRequestResponseDataModel() ReturnRequestResponse.MessageModel {
	var message, _ = ReturnRequestResponse.NewMessage("")
	message.Data.AssignmentId = "20250310B1QDRCQR000422"
	message.Data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.AssignmentCreateTime = time.Now()
	message.Data.ResolvedCaseId = "20250310011104238Sc02Step1MsgIdSVNR"
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
	message.Data.Status = ReturnRequestResponse.ReturnRequestRejected
	message.Data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario02InstrId001"
	message.Data.OriginalEndToEndId = "Scenario02EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.CancellationStatusReasonInfo = ReturnRequestResponse.Reason{
		Reason:         "NARR",
		AdditionalInfo: "Corporation B delivered goods and services are in-line with client’s order.",
	}

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
