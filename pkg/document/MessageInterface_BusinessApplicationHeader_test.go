package document

import (
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/moov-io/wire20022/pkg/models/BusinessApplicationHeader"
	"github.com/stretchr/testify/require"
)

func TestBusinessApplicationHeaderParseXMLFile(t *testing.T) {
	xmlFile := "../models/BusinessApplicationHeader/generated/AccountBalanceReport_Scenario1_Step1_head.xml"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	message, error := ParseXML(xmlData, &BusinessApplicationHeader.Message{})
	require.NoError(t, error, "Failed to make XML structure")
	if msgModel, ok := message.GetDataModel().(*BusinessApplicationHeader.MessageModel); ok {
		require.Equal(t, msgModel.MessageSenderId, "231981435")
	}
}

func TestBusinessApplicationHeaderGenerateXML(t *testing.T) {
	dataModel := BusinessApplicationHeaderModel()
	xmlData, err := GenerateXML(&dataModel, &BusinessApplicationHeader.Message{})
	require.NoError(t, err)
	err = model.WriteXMLTo("BusinessApplicationHeader_test.xml", xmlData)
	require.NoError(t, err)
}

func TestBusinessApplicationHeaderRequireFieldCheck(t *testing.T) {
	dataModel := BusinessApplicationHeaderModel()
	dataModel.MessageSenderId = ""
	dataModel.MessageReceiverId = ""
	valid, err := RequireFieldCheck(&dataModel, &BusinessApplicationHeader.Message{})
	require.Equal(t, valid, false)
	require.Equal(t, err.Error(), "error occur at RequiredFields: MessageSenderId, MessageReceiverId")
}

func TestBusinessApplicationHeaderXMLValidation(t *testing.T) {
	xmlFile := "../models/BusinessApplicationHeader/swiftSample/AccountBalanceReport_Scenario1_Step1_head.001"
	var xmlData, err = model.ReadXMLFile(xmlFile)
	require.NoError(t, err, "Failed to read XML file")
	valid, err := Validate(xmlData, &BusinessApplicationHeader.Message{})
	require.NoError(t, err)
	require.Equal(t, valid, true)
}

func TestBusinessApplicationHeaderAccessToHelper(t *testing.T) {
	message, cErr := CreateMessage(&BusinessApplicationHeader.Message{})
	require.NoError(t, cErr)
	if helper, ok := message.GetHelper().(*BusinessApplicationHeader.MessageHelper); ok {
		require.Equal(t, helper.MessageDefinitionId.Title, "Message Definition Identifier")
		require.Equal(t, helper.MessageDefinitionId.Type, "Max35Text (based on string) minLength: 1 maxLength: 35")
		require.Equal(t, helper.MessageDefinitionId.Documentation, "The Message Definition Identifier of the Business Message instance with which this Business Application Header instance is associated.")
	}
}

func BusinessApplicationHeaderModel() BusinessApplicationHeader.MessageModel {
	var mesage, _ = BusinessApplicationHeader.NewMessage("")
	mesage.Data.MessageSenderId = "231981435"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250311143738 ABAR MM Request"
	mesage.Data.MessageDefinitionId = "camt.060.001.05"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = BusinessApplicationHeader.MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	cErr := mesage.CreateDocument()
	if cErr != nil {
		return mesage.Data
	}
	return mesage.Data
}
