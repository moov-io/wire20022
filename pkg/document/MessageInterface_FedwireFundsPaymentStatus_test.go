package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/FedwireFundsPaymentStatus"
	"github.com/stretchr/testify/require"
)

var FedwireFundsPaymentStatusxmlFile = "../models/FedwireFundsPaymentStatus/swiftSample/CustomerCreditTransfer_Scenario1_Step2_pacs.002"

func TestFedwireFundsPaymentStatusParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(FedwireFundsPaymentStatusxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &FedwireFundsPaymentStatus.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*FedwireFundsPaymentStatus.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310QMGFNP31000001")
	}
}

func TestFedwireFundsPaymentStatusGenerateXML(t *testing.T) {
	dataModel := FedwireFundsPaymentStatusDataModel()
	xmlData, err := GenerateXML(&dataModel, &FedwireFundsPaymentStatus.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsPaymentStatus_test.xml", xmlData)
	require.NoError(t, err)
}

func TestFedwireFundsPaymentStatusRequireFieldCheck(t *testing.T) {
	dataModel := FedwireFundsPaymentStatusDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &FedwireFundsPaymentStatus.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestFedwireFundsPaymentStatusXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(FedwireFundsPaymentStatusxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &FedwireFundsPaymentStatus.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestFedwireFundsPaymentStatusAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&FedwireFundsPaymentStatus.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*FedwireFundsPaymentStatus.MessageHelper); ok {
		require.Equal(t, helper.AcceptanceDateTime.Title, "Acceptance Date Time")
		require.Equal(t, helper.AcceptanceDateTime.Type, "ISODateTime (based on dateTime)")
		require.Equal(t, helper.AcceptanceDateTime.Documentation, "Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.")
	}
}

func FedwireFundsPaymentStatusDataModel() FedwireFundsPaymentStatus.MessageModel {
	var message, _ = FedwireFundsPaymentStatus.NewMessage("")
	message.Data.MessageId = "FDWA1B2C3D4E5F6G7H8I9J10K11L12M0"
	message.Data.CreatedDateTime = time.Now()
	message.Data.OriginalMessageId = "20250310B1QDRCQR000002"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalMessageCreateTime = time.Now()
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.TransactionStatus = model.Rejected
	message.Data.StatusReasonInformation = "E433"
	message.Data.ReasonAdditionalInfo = "The routing number of the Instructed Agent is not permissible to receive Fedwire Funds transaction."
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
