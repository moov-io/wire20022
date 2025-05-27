package DrawdownResponse

import (
	"encoding/xml"
	"testing"
	"time"

	Archive "github.com/moov-io/wire20022/pkg/archives"
	"github.com/stretchr/testify/require"
)

func TestVersion01(t *testing.T) {
	modelName := PAIN_014_001_01
	xmlName := "DrawdownRequest_01.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion02(t *testing.T) {
	modelName := PAIN_014_001_02
	xmlName := "DrawdownRequest_02.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion03(t *testing.T) {
	modelName := PAIN_014_001_03
	xmlName := "DrawdownRequest_03.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion04(t *testing.T) {
	modelName := PAIN_014_001_04
	xmlName := "DrawdownRequest_04.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion05(t *testing.T) {
	modelName := PAIN_014_001_05
	xmlName := "DrawdownRequest_05.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}

func TestVersion06(t *testing.T) {
	modelName := PAIN_014_001_06
	xmlName := "DrawdownRequest_06.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion07(t *testing.T) {
	modelName := PAIN_014_001_07
	xmlName := "DrawdownRequest_07.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalUniqueId, "8a562c67-ca16-48ba-b074-65581be6f078")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion08(t *testing.T) {
	modelName := PAIN_014_001_08
	xmlName := "DrawdownRequest_08.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalUniqueId, "8a562c67-ca16-48ba-b074-65581be6f078")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion09(t *testing.T) {
	modelName := PAIN_014_001_09
	xmlName := "DrawdownRequest_09.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalUniqueId, "8a562c67-ca16-48ba-b074-65581be6f078")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion10(t *testing.T) {
	modelName := PAIN_014_001_10
	xmlName := "DrawdownRequest_10.xml"

	dataModel := DrawdownResponseDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000712")
	require.NotNil(t, model.CreateDatetime)
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, Archive.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000711")
	require.Equal(t, model.OriginalMessageNameId, "pain.013.001.07")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalPaymentInfoId, "20250310B1QDRCQR000711")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalEndToEndId, "Scenario01Step1EndToEndId001")
	require.Equal(t, model.TransactionInformationAndStatus.OriginalUniqueId, "8a562c67-ca16-48ba-b074-65581be6f078")
	require.Equal(t, model.TransactionInformationAndStatus.TransactionStatus, Archive.AcceptedTechnicalValidation)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000601"
}

func DrawdownResponseDataModel() MessageModel {
	var message = MessageModel{}
	message.MessageId = "20250310B1QDRCQR000712"
	message.CreateDatetime = time.Now()
	message.InitiatingParty = Archive.PartyIdentify{
		Name: "Corporation A",
		Address: Archive.PostalAddress{
			StreetName:     "Avenue of the Fountains",
			BuildingNumber: "167565",
			RoomNumber:     "Suite D110",
			PostalCode:     "85268",
			TownName:       "Fountain Hills",
			Subdivision:    "AZ",
			Country:        "US",
		},
	}
	message.DebtorAgent = Archive.Agent{
		PaymentSysCode:     Archive.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.CreditorAgent = Archive.Agent{
		PaymentSysCode:     Archive.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.OriginalMessageId = "20250310B1QDRCQR000711"
	message.OriginalMessageNameId = "pain.013.001.07"
	message.OriginalCreationDateTime = time.Now()
	message.OriginalPaymentInfoId = "20250310B1QDRCQR000711"
	message.TransactionInformationAndStatus = TransactionInfoAndStatus{
		OriginalInstructionId: "Scenario01InstrId001",
		OriginalEndToEndId:    "Scenario01Step1EndToEndId001",
		OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f078",
		TransactionStatus:     Archive.AcceptedTechnicalValidation,
	}
	return message
}
