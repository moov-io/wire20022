package PaymentReturn

import (
	"encoding/xml"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion2(t *testing.T) {
	modelName := PACS_004_001_02
	xmlName := "PaymentReturn_02.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}
func TestVersion3(t *testing.T) {
	modelName := PACS_004_001_03
	xmlName := "PaymentReturn_03.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}
func TestVersion4(t *testing.T) {
	modelName := PACS_004_001_04
	xmlName := "PaymentReturn_04.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}
func TestVersion5(t *testing.T) {
	modelName := PACS_004_001_05
	xmlName := "PaymentReturn_05.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}
func TestVersion6(t *testing.T) {
	modelName := PACS_004_001_06
	xmlName := "PaymentReturn_06.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}

func TestVersion7(t *testing.T) {
	modelName := PACS_004_001_07
	xmlName := "PaymentReturn_07.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}

func TestVersion8(t *testing.T) {
	modelName := PACS_004_001_08
	xmlName := "PaymentReturn_08.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.Debtor.Name, "Corporation B")
	require.Equal(t, model.RtrChain.Debtor.Address.StreetName, "Desert View Street")
	require.Equal(t, model.RtrChain.Debtor.Address.BuildingNumber, "1")
	require.Equal(t, model.RtrChain.Debtor.Address.PostalCode, "92262")
	require.Equal(t, model.RtrChain.Debtor.Address.TownName, "Palm Springs")
	require.Equal(t, model.RtrChain.Debtor.Address.Subdivision, "CA")
	require.Equal(t, model.RtrChain.Debtor.Address.Country, "US")
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.RtrChain.DebtorAgent.BankName, "BankB")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.CreditorAgent.BankName, "BankA")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.Creditor.Name, "Corporation A")
	require.Equal(t, model.RtrChain.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.RtrChain.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.RtrChain.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.RtrChain.Creditor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.RtrChain.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.RtrChain.Creditor.Address.Country, "US")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}
func TestVersion9(t *testing.T) {
	modelName := PACS_004_001_09
	xmlName := "PaymentReturn_09.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.Debtor.Name, "Corporation B")
	require.Equal(t, model.RtrChain.Debtor.Address.StreetName, "Desert View Street")
	require.Equal(t, model.RtrChain.Debtor.Address.BuildingNumber, "1")
	require.Equal(t, model.RtrChain.Debtor.Address.Floor, "33")
	require.Equal(t, model.RtrChain.Debtor.Address.PostalCode, "92262")
	require.Equal(t, model.RtrChain.Debtor.Address.TownName, "Palm Springs")
	require.Equal(t, model.RtrChain.Debtor.Address.Subdivision, "CA")
	require.Equal(t, model.RtrChain.Debtor.Address.Country, "US")
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.RtrChain.DebtorAgent.BankName, "BankB")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.CreditorAgent.BankName, "BankA")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.Creditor.Name, "Corporation A")
	require.Equal(t, model.RtrChain.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.RtrChain.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.RtrChain.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.RtrChain.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.RtrChain.Creditor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.RtrChain.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.RtrChain.Creditor.Address.Country, "US")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}
func TestVersion10(t *testing.T) {
	modelName := PACS_004_001_10
	xmlName := "PaymentReturn_10.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.Debtor.Name, "Corporation B")
	require.Equal(t, model.RtrChain.Debtor.Address.StreetName, "Desert View Street")
	require.Equal(t, model.RtrChain.Debtor.Address.BuildingNumber, "1")
	require.Equal(t, model.RtrChain.Debtor.Address.Floor, "33")
	require.Equal(t, model.RtrChain.Debtor.Address.PostalCode, "92262")
	require.Equal(t, model.RtrChain.Debtor.Address.TownName, "Palm Springs")
	require.Equal(t, model.RtrChain.Debtor.Address.Subdivision, "CA")
	require.Equal(t, model.RtrChain.Debtor.Address.Country, "US")
	require.Equal(t, model.RtrChain.DebtorOtherTypeId, "567876543")
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.RtrChain.DebtorAgent.BankName, "BankB")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.CreditorAgent.BankName, "BankA")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.Creditor.Name, "Corporation A")
	require.Equal(t, model.RtrChain.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.RtrChain.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.RtrChain.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.RtrChain.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.RtrChain.Creditor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.RtrChain.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.RtrChain.Creditor.Address.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAccountOtherTypeId, "5647772655")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}
func TestVersion11(t *testing.T) {
	modelName := PACS_004_001_11
	xmlName := "PaymentReturn_11.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.Debtor.Name, "Corporation B")
	require.Equal(t, model.RtrChain.Debtor.Address.StreetName, "Desert View Street")
	require.Equal(t, model.RtrChain.Debtor.Address.BuildingNumber, "1")
	require.Equal(t, model.RtrChain.Debtor.Address.Floor, "33")
	require.Equal(t, model.RtrChain.Debtor.Address.PostalCode, "92262")
	require.Equal(t, model.RtrChain.Debtor.Address.TownName, "Palm Springs")
	require.Equal(t, model.RtrChain.Debtor.Address.Subdivision, "CA")
	require.Equal(t, model.RtrChain.Debtor.Address.Country, "US")
	require.Equal(t, model.RtrChain.DebtorOtherTypeId, "567876543")
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.RtrChain.DebtorAgent.BankName, "BankB")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.CreditorAgent.BankName, "BankA")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.Creditor.Name, "Corporation A")
	require.Equal(t, model.RtrChain.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.RtrChain.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.RtrChain.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.RtrChain.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.RtrChain.Creditor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.RtrChain.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.RtrChain.Creditor.Address.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAccountOtherTypeId, "5647772655")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}
func TestVersion12(t *testing.T) {
	modelName := PACS_004_001_12
	xmlName := "PaymentReturn_12.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.Debtor.Name, "Corporation B")
	require.Equal(t, model.RtrChain.Debtor.Address.StreetName, "Desert View Street")
	require.Equal(t, model.RtrChain.Debtor.Address.BuildingNumber, "1")
	require.Equal(t, model.RtrChain.Debtor.Address.Floor, "33")
	require.Equal(t, model.RtrChain.Debtor.Address.PostalCode, "92262")
	require.Equal(t, model.RtrChain.Debtor.Address.TownName, "Palm Springs")
	require.Equal(t, model.RtrChain.Debtor.Address.Subdivision, "CA")
	require.Equal(t, model.RtrChain.Debtor.Address.Country, "US")
	require.Equal(t, model.RtrChain.DebtorOtherTypeId, "567876543")
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.RtrChain.DebtorAgent.BankName, "BankB")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.CreditorAgent.BankName, "BankA")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.Creditor.Name, "Corporation A")
	require.Equal(t, model.RtrChain.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.RtrChain.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.RtrChain.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.RtrChain.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.RtrChain.Creditor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.RtrChain.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.RtrChain.Creditor.Address.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAccountOtherTypeId, "5647772655")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}
func TestVersion13(t *testing.T) {
	modelName := PACS_004_001_13
	xmlName := "PaymentReturn_13.xml"

	dataModel := PaymentReturnDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000724")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementCLRG)
	require.Equal(t, model.ClearingSystem, models.ClearingSysFDW)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000721")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalCreationDateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario01InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInterbankSettlementAmount.Currency, "USD")
	require.NotNil(t, model.InterbankSettlementDate)
	require.Equal(t, model.ReturnedInstructedAmount.Amount, 151235.88)
	require.Equal(t, model.ReturnedInstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.Debtor.Name, "Corporation B")
	require.Equal(t, model.RtrChain.Debtor.Address.StreetName, "Desert View Street")
	require.Equal(t, model.RtrChain.Debtor.Address.BuildingNumber, "1")
	require.Equal(t, model.RtrChain.Debtor.Address.Floor, "33")
	require.Equal(t, model.RtrChain.Debtor.Address.PostalCode, "92262")
	require.Equal(t, model.RtrChain.Debtor.Address.TownName, "Palm Springs")
	require.Equal(t, model.RtrChain.Debtor.Address.Subdivision, "CA")
	require.Equal(t, model.RtrChain.Debtor.Address.Country, "US")
	require.Equal(t, model.RtrChain.DebtorOtherTypeId, "567876543")
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.RtrChain.DebtorAgent.BankName, "BankB")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.RtrChain.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.RtrChain.CreditorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.RtrChain.CreditorAgent.BankName, "BankA")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.RtrChain.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.RtrChain.Creditor.Name, "Corporation A")
	require.Equal(t, model.RtrChain.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.RtrChain.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.RtrChain.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.RtrChain.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.RtrChain.Creditor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.RtrChain.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.RtrChain.Creditor.Address.Country, "US")
	require.Equal(t, model.RtrChain.CreditorAccountOtherTypeId, "5647772655")
	require.Equal(t, model.ReturnReasonInformation.Reason, "DUPL")
	require.Contains(t, model.ReturnReasonInformation.AdditionalInfo, "Order cancelled.")
	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000724"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "missing required field: MessageId")
	model.MessageId = "20250310B1QDRCQR000724"
}

func PaymentReturnDataModel() MessageModel {
	message := MessageModel{}
	message.MessageId = "20250310B1QDRCQR000724"
	message.CreatedDateTime = time.Now()
	message.NumberOfTransactions = "1"
	message.SettlementMethod = models.SettlementCLRG
	message.ClearingSystem = models.ClearingSysFDW
	message.OriginalMessageId = "20250310B1QDRCQR000721"
	message.OriginalMessageNameId = "pacs.008.001.08"
	message.OriginalCreationDateTime = time.Now()
	message.OriginalInstructionId = "Scenario01InstrId001"
	message.OriginalEndToEndId = "Scenario01EtoEId001"
	message.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.ReturnedInterbankSettlementAmount = models.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.InterbankSettlementDate = fedwire.ISODate(civil.DateOf(time.Now()))
	message.ReturnedInstructedAmount = models.CurrencyAndAmount{
		Amount:   151235.88,
		Currency: "USD",
	}
	message.ChargeBearer = models.ChargeBearerSLEV
	message.InstructingAgent = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.InstructedAgent = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.RtrChain = models.ReturnChain{
		Debtor: models.Party{
			Name: "Corporation B",
			Address: models.PostalAddress{
				StreetName:     "Desert View Street",
				BuildingNumber: "1",
				Floor:          "33",
				PostalCode:     "92262",
				TownName:       "Palm Springs",
				Subdivision:    "CA",
				Country:        "US",
			},
		},
		DebtorOtherTypeId: "567876543",
		DebtorAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
			BankName:           "BankB",
			PostalAddress: models.PostalAddress{
				StreetName:     "Avenue B",
				BuildingNumber: "25",
				PostalCode:     "19067",
				TownName:       "Yardley",
				Subdivision:    "PA",
				Country:        "US",
			},
		},
		CreditorAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
			BankName:           "BankA",
			PostalAddress: models.PostalAddress{
				StreetName:     "Avenue A",
				BuildingNumber: "66",
				PostalCode:     "60532",
				TownName:       "Lisle",
				Subdivision:    "IL",
				Country:        "US",
			},
		},
		Creditor: models.Party{
			Name: "Corporation A",
			Address: models.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CreditorAccountOtherTypeId: "5647772655",
	}
	message.ReturnReasonInformation = models.Reason{
		Reason:         "DUPL",
		AdditionalInfo: "Order cancelled. Ref:20250310B1QDRCQR000721.",
	}
	message.OriginalTransactionRef = models.InstrumentCTRC
	return message
}
