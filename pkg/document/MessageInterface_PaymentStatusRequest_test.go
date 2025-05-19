package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest"
	"github.com/stretchr/testify/require"
)

var PaymentStatusRequestxmlFile = "../models/PaymentStatusRequest/swiftSample/CustomerCreditTransfer_Scenario3_Step2_pacs.028"

func TestPaymentStatusRequestParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(PaymentStatusRequestxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &PaymentStatusRequest.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*PaymentStatusRequest.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310Scenario03Step2MsgId001")
	}
}

func TestPaymentStatusRequestGenerateXML(t *testing.T) {
	dataModel := PaymentStatusRequestDataModel()
	xmlData, err := GenerateXML(&dataModel, &PaymentStatusRequest.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("PaymentStatusRequest_test.xml", xmlData)
	require.NoError(t, err)
}

func TestPaymentStatusRequestRequireFieldCheck(t *testing.T) {
	dataModel := PaymentStatusRequestDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &PaymentStatusRequest.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestPaymentStatusRequestXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(PaymentStatusRequestxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &PaymentStatusRequest.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestPaymentStatusRequestAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&PaymentStatusRequest.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*PaymentStatusRequest.MessageHelper); ok {
		require.Equal(t, helper.OriginalInstructionId.Title, "Original Instruction Id")
		require.Equal(t, helper.OriginalInstructionId.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.OriginalInstructionId.Documentation, "Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.")
	}
}

func PaymentStatusRequestDataModel() PaymentStatusRequest.MessageModel {
	var message, _ = PaymentStatusRequest.NewMessage("")
	message.Data.MessageId = "20250310Scenario04Step3MsgId001"
	message.Data.CreatedDateTime = time.Now()
	message.Data.OriginalMessageId = "20250310B1QDRCQR000631"
	message.Data.OriginalMessageNameId = "pain.013.001.07"
	message.Data.OriginalCreationDateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario04Step1InstrId001"
	message.Data.OriginalEndToEndId = "Scenario4EndToEndId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f258"
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
