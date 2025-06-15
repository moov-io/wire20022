package DrawdownRequest

import (
	"encoding/xml"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion01(t *testing.T) {
	modelName := PAIN_013_001_01
	xmlName := "DrawdownRequest_01.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Contains(t, err.Error(), "failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion02(t *testing.T) {
	modelName := PAIN_013_001_02
	xmlName := "DrawdownRequest_02.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion03(t *testing.T) {
	modelName := PAIN_013_001_03
	xmlName := "DrawdownRequest_03.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion04(t *testing.T) {
	modelName := PAIN_013_001_04
	xmlName := "DrawdownRequest_04.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion05(t *testing.T) {
	modelName := PAIN_013_001_05
	xmlName := "DrawdownRequest_05.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion06(t *testing.T) {
	modelName := PAIN_013_001_06
	xmlName := "DrawdownRequest_06.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion07(t *testing.T) {
	modelName := PAIN_013_001_07
	xmlName := "DrawdownRequest_07.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PaymentUniqueId, "8a562c67-ca16-48ba-b074-65581be6f066")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion08(t *testing.T) {
	modelName := PAIN_013_001_08
	xmlName := "DrawdownRequest_08.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PaymentUniqueId, "8a562c67-ca16-48ba-b074-65581be6f066")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion09(t *testing.T) {
	modelName := PAIN_013_001_09
	xmlName := "DrawdownRequest_09.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PaymentUniqueId, "8a562c67-ca16-48ba-b074-65581be6f066")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}
func TestVersion10(t *testing.T) {
	modelName := PAIN_013_001_10
	xmlName := "DrawdownRequest_10.xml"

	dataModel := DrawdownRequestDataModel()
	/*Create Document from Model*/
	var doc08, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc08.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc08, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)
	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")
	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000601")
	require.NotNil(t, model.CreditTransTransaction)
	require.Equal(t, model.NumberofTransaction, "1")
	require.Equal(t, model.InitiatingParty.Name, "Corporation A")
	require.Equal(t, model.InitiatingParty.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.InitiatingParty.Address.BuildingNumber, "167565")
	require.Equal(t, model.InitiatingParty.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.InitiatingParty.Address.PostalCode, "85268")
	require.Equal(t, model.InitiatingParty.Address.TownName, "Fountain Hills")
	require.Equal(t, model.InitiatingParty.Address.Subdivision, "AZ")
	require.Equal(t, model.InitiatingParty.Address.Country, "US")
	require.Equal(t, model.PaymentInfoId, "20250310B1QDRCQR000601")
	require.Equal(t, model.PaymentMethod, models.CreditTransform)
	require.NotNil(t, model.RequestedExecutDate)
	require.Equal(t, model.Debtor.Name, "Corporation A")
	require.Equal(t, model.Debtor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.Debtor.Address.BuildingNumber, "167565")
	require.Equal(t, model.Debtor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.Debtor.Address.PostalCode, "85268")
	require.Equal(t, model.Debtor.Address.TownName, "Fountain Hills")
	require.Equal(t, model.Debtor.Address.Subdivision, "AZ")
	require.Equal(t, model.Debtor.Address.Country, "US")
	require.Equal(t, model.DebtorAccountOtherId, "92315266453")
	require.NotNil(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditTransTransaction.PaymentInstructionId, "Scenario01Step1InstrId001")
	require.Equal(t, model.CreditTransTransaction.PaymentEndToEndId, "Scenario1EndToEndId001")
	require.Equal(t, model.CreditTransTransaction.PaymentUniqueId, "8a562c67-ca16-48ba-b074-65581be6f066")
	require.Equal(t, model.CreditTransTransaction.PayCategoryType, models.IntraCompanyPayment)
	require.Equal(t, model.CreditTransTransaction.PayRequestType, models.DrawDownRequestCredit)
	require.Equal(t, model.CreditTransTransaction.Amount.Amount, 6000000.00)
	require.Equal(t, model.CreditTransTransaction.Amount.Currency, "USD")
	require.Equal(t, model.CreditTransTransaction.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.CreditTransTransaction.Creditor.Name, "Corporation A")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.BuildingNumber, "167565")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.RoomNumber, "Suite D110")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.PostalCode, "85268")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.TownName, "Fountain HIlls")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Subdivision, "AZ")
	require.Equal(t, model.CreditTransTransaction.Creditor.Address.Country, "US")
	require.Equal(t, model.CreditTransTransaction.CrediorAccountOtherId, "5647772655")
	require.Contains(t, model.CreditTransTransaction.RemittanceInformation, "EDAY ACCT BALANCING")
	require.Equal(t, model.CreditTransTransaction.Document.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.CreditTransTransaction.Document.Number, "INV12345")
	require.NotNil(t, model.CreditTransTransaction.Document.RelatedDate)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000601"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000601"
}

func DrawdownRequestDataModel() MessageModel {
	var message = MessageModel{}
	message.MessageId = "20250310B1QDRCQR000601"
	message.CreatedDateTime = time.Now()
	message.NumberofTransaction = "1"
	message.InitiatingParty = models.PartyIdentify{
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
	}
	message.PaymentInfoId = "20250310B1QDRCQR000601"
	message.PaymentMethod = models.CreditTransform
	message.RequestedExecutDate = fedwire.ISODate(civil.DateOf(time.Now()))
	message.Debtor = models.PartyIdentify{
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
	}
	message.DebtorAccountOtherId = "92315266453"
	message.DebtorAgent = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.CreditTransTransaction = CreditTransferTransaction{
		PaymentInstructionId: "Scenario01Step1InstrId001",
		PaymentEndToEndId:    "Scenario1EndToEndId001",
		PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
		PayCategoryType:      models.IntraCompanyPayment,
		PayRequestType:       models.DrawDownRequestCredit,
		Amount: models.CurrencyAndAmount{
			Amount:   6000000.00,
			Currency: "USD",
		},
		ChargeBearer: models.ChargeBearerSLEV,
		CreditorAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		Creditor: models.PartyIdentify{
			Name: "Corporation A",
			Address: models.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain HIlls",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		CrediorAccountOtherId: "5647772655",
		RemittanceInformation: "EDAY ACCT BALANCING//10 March 2025//$60,000,000.00",
		Document: RemittanceDocument{
			CodeOrProprietary: models.CodeCINV,
			Number:            "INV12345",
			RelatedDate:       fedwire.ISODate(civil.DateOf(time.Now())),
		},
	}
	return message
}
