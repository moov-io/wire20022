package document

import (
	"testing"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/InvestResponse"
	"github.com/stretchr/testify/require"
)

var InvestResponsexmlFile = "../models/InvestResponse/swiftSample/Investigations_Scenario1_Step3_camt.111"

func TestInvestResponseParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(InvestResponsexmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &InvestResponse.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*InvestResponse.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310B1QDRCQR000901")
	}
}

func TestInvestResponseGenerateXML(t *testing.T) {
	dataModel := InvestResponseDataModel()
	xmlData, err := GenerateXML(&dataModel, &InvestResponse.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("InvestResponse_test.xml", xmlData)
	require.NoError(t, err)
}

func TestInvestResponseRequireFieldCheck(t *testing.T) {
	dataModel := InvestResponseDataModel()
	dataModel.MessageId = ""
	dataModel.InvestRequestMessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &InvestResponse.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId, InvestRequestMessageId")
}

func TestInvestResponseXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(InvestResponsexmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &InvestResponse.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestInvestResponseAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&InvestResponse.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*InvestResponse.MessageHelper); ok {
		require.Equal(t, helper.InvestRequestMessageId.Title, "Invest Request Message Id")
		require.Equal(t, helper.InvestRequestMessageId.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.InvestRequestMessageId.Documentation, "Point to point reference, as assigned by the requestor, and sent to the responder to unambiguously identify the message.")
	}
}

func InvestResponseDataModel() InvestResponse.MessageModel {
	var message, _ = InvestResponse.NewMessage("")
	message.Data.MessageId = "20250310B1QDRCQR000902"
	message.Data.InvestigationStatus = "CLSD"
	message.Data.InvestigationData = "Payment is a duplicate. Please consider VOID. Return request will follow."
	message.Data.InvestRequestMessageId = "20250310QMGFT015000902"
	message.Data.InvestigationType = "OTHR"
	message.Data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
