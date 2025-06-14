package FedwireFundsPaymentStatus

import (
	"encoding/xml"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion3(t *testing.T) {
	modelName := PACS_002_001_03
	xmlName := "FedwireFundsAcknowledgement_03.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}

func TestVersion4(t *testing.T) {
	modelName := PACS_002_001_04
	xmlName := "FedwireFundsAcknowledgement_04.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion5(t *testing.T) {
	modelName := PACS_002_001_05
	xmlName := "FedwireFundsAcknowledgement_05.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion6(t *testing.T) {
	modelName := PACS_002_001_06
	xmlName := "FedwireFundsAcknowledgement_06.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion7(t *testing.T) {
	modelName := PACS_002_001_07
	xmlName := "FedwireFundsAcknowledgement_08.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion8(t *testing.T) {
	modelName := PACS_002_001_08
	xmlName := "FedwireFundsAcknowledgement_08.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion9(t *testing.T) {
	modelName := PACS_002_001_09
	xmlName := "FedwireFundsAcknowledgement_09.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion10(t *testing.T) {
	modelName := PACS_002_001_10
	xmlName := "FedwireFundsAcknowledgement_10.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion11(t *testing.T) {
	modelName := PACS_002_001_11
	xmlName := "FedwireFundsAcknowledgement_11.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion12(t *testing.T) {
	modelName := PACS_002_001_12
	xmlName := "FedwireFundsAcknowledgement_12.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion13(t *testing.T) {
	modelName := PACS_002_001_13
	xmlName := "FedwireFundsAcknowledgement_13.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}
func TestVersion14(t *testing.T) {
	modelName := PACS_002_001_14
	xmlName := "FedwireFundsAcknowledgement_14.xml"

	dataModel := FedwireFundsPaymentStatusDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310QMGFNP31000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.TransactionStatus, models.AcceptedSettlementCompleted)
	require.NotNil(t, model.AcceptanceDateTime)
	require.NotNil(t, model.EffectiveInterbankSettlementDate)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021151080")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFIPmtStsRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310QMGFNP7500070103101100FT03"
}

func FedwireFundsPaymentStatusDataModel() MessageModel {
	message := MessageModel{}
	message.MessageId = "20250310QMGFNP31000001"
	message.CreatedDateTime = time.Now()
	message.OriginalMessageId = "20250310B1QDRCQR000001"
	message.OriginalMessageNameId = "pacs.008.001.08"
	message.OriginalMessageCreateTime = time.Now()
	message.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.TransactionStatus = models.AcceptedSettlementCompleted
	message.EffectiveInterbankSettlementDate = fedwire.ISODate(civil.DateOf(time.Now()))
	message.AcceptanceDateTime = time.Now()
	message.InstructingAgent = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.InstructedAgent = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	return message
}
