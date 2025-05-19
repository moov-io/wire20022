package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/InvestRequest"
	"github.com/stretchr/testify/require"
)

var InvestRequestxmlFile = "../models/InvestRequest/swiftSample/Investigations_Scenario1_Step2_camt.110"

func TestInvestRequestParseXMLFile(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(InvestRequestxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &InvestRequest.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*InvestRequest.MessageModel); ok {
		require.Equal(t, msgModel.MessageId, "20250310QMGFT015000901")
	}
}

func TestInvestRequestGenerateXML(t *testing.T) {
	dataModel := InvestRequestDataModel()
	xmlData, err := GenerateXML(&dataModel, &InvestRequest.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("InvestRequest_test.xml", xmlData)
	require.NoError(t, err)
}

func TestInvestRequestRequireFieldCheck(t *testing.T) {
	dataModel := InvestRequestDataModel()
	dataModel.MessageId = ""
	valid, err := RequireFieldCheck(&dataModel, &InvestRequest.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageId")
}

func TestInvestRequestXMLValidation(t *testing.T) {
	var xmlData, err = model.ReadXMLFile(InvestRequestxmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &InvestRequest.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestInvestRequestAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&InvestRequest.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*InvestRequest.MessageHelper); ok {
		require.Equal(t, helper.InvestigationType.Title, "Investigation Type")
		require.Equal(t, helper.InvestigationType.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.InvestigationType.Documentation, "Type of investigation.")
	}
}

func InvestRequestDataModel() InvestRequest.MessageModel {
	var message, _ = InvestRequest.NewMessage("")
	message.Data.MessageId = "20250310QMGFT015000903"
	message.Data.InvestigationType = "RQFI"
	message.Data.UnderlyingData = InvestRequest.Underlying{
		OriginalMessageId:        "20250310B1QDRCQR000007",
		OriginalMessageNameId:    "pacs.008.001.08",
		OriginalCreationDateTime: time.Now(),
		OriginalInstructionId:    "Scenario01InstrId001",
		OriginalEndToEndId:       "Scenario01EtoEId001",
		OriginalUETR:             "8a562c67-ca16-48ba-b074-65581be6f011",
		OriginalInterbankSettlementAmount: model.CurrencyAndAmount{
			Amount:   510000.74,
			Currency: "USD",
		},
		OriginalInterbankSettlementDate: model.FromTime(time.Now()),
	}
	message.Data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.InvestReason = InvestRequest.InvestigationReason{
		Reason: "MS01",
	}

	cErr := message.CreateDocument()
	if cErr != nil {
		return message.Data
	}
	return message.Data
}
