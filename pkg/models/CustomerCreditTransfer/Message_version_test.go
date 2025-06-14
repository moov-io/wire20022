package CustomerCreditTransfer

import (
	"encoding/xml"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion02(t *testing.T) {
	modelName := PACS_008_001_02
	xmlName := "CustomerCreditTransfer_02.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.TaxId, "123456789")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion03(t *testing.T) {
	modelName := PACS_008_001_03
	xmlName := "CustomerCreditTransfer_03.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.TaxId, "123456789")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion04(t *testing.T) {
	modelName := PACS_008_001_04
	xmlName := "CustomerCreditTransfer_04.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.TaxId, "123456789")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion05(t *testing.T) {
	modelName := PACS_008_001_05
	xmlName := "CustomerCreditTransfer_05.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.TaxId, "123456789")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxId, "123456789")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxTypeCode, "09455")
	require.NotNil(t, model.RemittanceInfor.TaxDetail.TaxPeriodYear)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxperiodTimeFrame, "MM04")
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion06(t *testing.T) {
	modelName := PACS_008_001_06
	xmlName := "CustomerCreditTransfer_06.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.TaxId, "123456789")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxId, "123456789")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxTypeCode, "09455")
	require.NotNil(t, model.RemittanceInfor.TaxDetail.TaxPeriodYear)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxperiodTimeFrame, "MM04")
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion07(t *testing.T) {
	modelName := PACS_008_001_07
	xmlName := "CustomerCreditTransfer_07.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.TaxId, "123456789")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxId, "123456789")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxTypeCode, "09455")
	require.NotNil(t, model.RemittanceInfor.TaxDetail.TaxPeriodYear)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxperiodTimeFrame, "MM04")
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion08(t *testing.T) {
	modelName := PACS_008_001_08
	xmlName := "CustomerCreditTransfer_08.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.UniqueEndToEndTransactionRef, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.RoomNumber, "Suite D110")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.Floor, "33")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxId, "123456789")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxTypeCode, "09455")
	require.NotNil(t, model.RemittanceInfor.TaxDetail.TaxPeriodYear)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxperiodTimeFrame, "MM04")
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion09(t *testing.T) {
	modelName := PACS_008_001_09
	xmlName := "CustomerCreditTransfer_09.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.UniqueEndToEndTransactionRef, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.RoomNumber, "Suite D110")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.Floor, "33")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxId, "123456789")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxTypeCode, "09455")
	require.NotNil(t, model.RemittanceInfor.TaxDetail.TaxPeriodYear)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxperiodTimeFrame, "MM04")
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion10(t *testing.T) {
	modelName := PACS_008_001_10
	xmlName := "CustomerCreditTransfer_10.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.UniqueEndToEndTransactionRef, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.RoomNumber, "Suite D110")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.Floor, "33")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxId, "123456789")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxTypeCode, "09455")
	require.NotNil(t, model.RemittanceInfor.TaxDetail.TaxPeriodYear)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxperiodTimeFrame, "MM04")
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion11(t *testing.T) {
	modelName := PACS_008_001_11
	xmlName := "CustomerCreditTransfer_11.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.UniqueEndToEndTransactionRef, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.RoomNumber, "Suite D110")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.Floor, "33")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxId, "123456789")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxTypeCode, "09455")
	require.NotNil(t, model.RemittanceInfor.TaxDetail.TaxPeriodYear)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxperiodTimeFrame, "MM04")
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}
func TestVersion12(t *testing.T) {
	modelName := PACS_008_001_12
	xmlName := "CustomerCreditTransfer_12.xml"

	dataModel := CustomerCreditTransferDataModel()
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
	require.Equal(t, model.MessageId, "20250310B1QDRCQR000001")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.NumberOfTransactions, "1")
	require.Equal(t, model.SettlementMethod, models.SettlementMethodType("CLRG"))
	require.Equal(t, model.CommonClearingSysCode, models.CommonClearingSysCodeType("FDW"))
	require.Equal(t, model.InstructionId, "Scenario01InstrId001")
	require.Equal(t, model.EndToEndId, "Scenario01EtoEId001")
	require.Equal(t, model.UniqueEndToEndTransactionRef, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.InstrumentPropCode, models.InstrumentPropCodeType("CTRC"))
	require.Equal(t, model.InterBankSettAmount.Amount, 510000.74)
	require.Equal(t, model.InterBankSettAmount.Currency, "USD")
	require.NotNil(t, model.InterBankSettDate)
	require.Equal(t, model.InstructedAmount.Amount, 510000.74)
	require.Equal(t, model.InstructedAmount.Currency, "USD")
	require.Equal(t, model.ChargeBearer, models.ChargeBearerSLEV)
	require.Equal(t, model.ChargesInfo[0].Amount.Amount, 90.00)
	require.Equal(t, model.ChargesInfo[0].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[0].BusinessIdCode, "BANZBEBB")
	require.Equal(t, model.ChargesInfo[1].Amount.Amount, 40.00)
	require.Equal(t, model.ChargesInfo[1].Amount.Currency, "USD")
	require.Equal(t, model.ChargesInfo[1].BusinessIdCode, "BANCUS33")
	require.Equal(t, model.InstructingAgents.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructingAgents.PaymentSysMemberId, "011104238")
	require.Equal(t, model.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.InstructedAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.DebtorName, "Corporation A")
	require.Equal(t, model.DebtorAddress.StreetName, "Avenue of the Fountains")
	require.Equal(t, model.DebtorAddress.BuildingNumber, "167565")
	require.Equal(t, model.DebtorAddress.RoomNumber, "Suite D110")
	require.Equal(t, model.DebtorAddress.PostalCode, "85268")
	require.Equal(t, model.DebtorAddress.TownName, "Fountain Hills")
	require.Equal(t, model.DebtorAddress.Subdivision, "AZ")
	require.Equal(t, model.DebtorAddress.Country, "US")
	require.Equal(t, model.DebtorOtherTypeId, "5647772655")
	require.Equal(t, model.DebtorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.DebtorAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.DebtorAgent.BankName, "Bank A")
	require.Equal(t, model.DebtorAgent.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.DebtorAgent.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.DebtorAgent.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.DebtorAgent.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.DebtorAgent.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.DebtorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.CreditorAgent.PaymentSysMemberId, "021040078")
	require.Equal(t, model.CreditorAgent.BankName, "Bank B")
	require.Equal(t, model.CreditorAgent.PostalAddress.StreetName, "Avenue B")
	require.Equal(t, model.CreditorAgent.PostalAddress.BuildingNumber, "25")
	require.Equal(t, model.CreditorAgent.PostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorAgent.PostalAddress.TownName, "Yardley")
	require.Equal(t, model.CreditorAgent.PostalAddress.Subdivision, "PA")
	require.Equal(t, model.CreditorAgent.PostalAddress.Country, "US")
	require.Equal(t, model.CreditorName, "Corporation B")
	require.Equal(t, model.CreditorPostalAddress.StreetName, "Desert View Street")
	require.Equal(t, model.CreditorPostalAddress.BuildingNumber, "1")
	require.Equal(t, model.CreditorPostalAddress.Floor, "33")
	require.Equal(t, model.CreditorPostalAddress.PostalCode, "19067")
	require.Equal(t, model.CreditorPostalAddress.TownName, "Palm Springs")
	require.Equal(t, model.CreditorPostalAddress.Subdivision, "CA")
	require.Equal(t, model.CreditorPostalAddress.Country, "US")
	require.Equal(t, model.CreditorOtherTypeId, "567876543")
	require.Equal(t, model.RemittanceInfor.CodeOrProprietary, models.CodeCINV)
	require.Equal(t, model.RemittanceInfor.Number, "INV34563")
	require.NotNil(t, model.RemittanceInfor.RelatedDate)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxId, "123456789")
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxTypeCode, "09455")
	require.NotNil(t, model.RemittanceInfor.TaxDetail.TaxPeriodYear)
	require.Equal(t, model.RemittanceInfor.TaxDetail.TaxperiodTimeFrame, "MM04")
	require.Equal(t, model.RelatedRemittanceInfo.RemittanceId, "Scenario01Var2RemittanceId001")
	require.Equal(t, model.RelatedRemittanceInfo.Method, models.Email)
	require.Equal(t, model.RelatedRemittanceInfo.ElectronicAddress, "CustomerService@CorporationB.com")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250310B1QDRCQR000001"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250310B1QDRCQR000001"
}

func CustomerCreditTransferDataModel() MessageModel {
	var mesage = MessageModel{}
	mesage.MessageId = "20250310B1QDRCQR000001"
	mesage.CreatedDateTime = time.Now()
	mesage.NumberOfTransactions = "1"
	mesage.SettlementMethod = models.SettlementCLRG
	mesage.CommonClearingSysCode = models.ClearingSysFDW
	mesage.InstructionId = "Scenario01InstrId001"
	mesage.EndToEndId = "Scenario01EtoEId001"
	mesage.UniqueEndToEndTransactionRef = "8a562c67-ca16-48ba-b074-65581be6f011"
	mesage.TaxId = "123456789"
	mesage.InstrumentPropCode = models.InstrumentCTRC
	mesage.InterBankSettAmount = models.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.InterBankSettDate = fedwire.ISODate(civil.DateOf(time.Now()))
	mesage.InstructedAmount = models.CurrencyAndAmount{
		Currency: "USD", Amount: 510000.74,
	}
	mesage.ChargeBearer = models.ChargeBearerSLEV
	mesage.ChargesInfo = []ChargeInfo{
		{
			Amount:         models.CurrencyAndAmount{Currency: "USD", Amount: 90.00},
			BusinessIdCode: "BANZBEBB",
		},
		{
			Amount:         models.CurrencyAndAmount{Currency: "USD", Amount: 40.00},
			BusinessIdCode: "BANCUS33",
		},
	}
	mesage.InstructingAgents = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	mesage.InstructedAgent = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	mesage.DebtorName = "Corporation A"
	mesage.DebtorAddress = models.PostalAddress{
		StreetName:     "Avenue of the Fountains",
		BuildingNumber: "167565",
		RoomNumber:     "Suite D110",
		PostalCode:     "85268",
		TownName:       "Fountain Hills",
		Subdivision:    "AZ",
		Country:        "US",
	}
	mesage.DebtorOtherTypeId = "5647772655"
	mesage.DebtorAgent = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: models.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	mesage.CreditorAgent = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
		BankName:           "Bank B",
		PostalAddress: models.PostalAddress{
			StreetName:     "Avenue B",
			BuildingNumber: "25",
			PostalCode:     "19067",
			TownName:       "Yardley",
			Subdivision:    "PA",
			Country:        "US",
		},
	}
	mesage.CreditorName = "Corporation B"
	mesage.CreditorPostalAddress = models.PostalAddress{
		StreetName:     "Desert View Street",
		BuildingNumber: "1",
		Floor:          "33",
		PostalCode:     "19067",
		TownName:       "Palm Springs",
		Subdivision:    "CA",
		Country:        "US",
	}
	mesage.CreditorOtherTypeId = "567876543"
	mesage.RemittanceInfor = RemittanceDocument{
		CodeOrProprietary: models.CodeCINV,
		Number:            "INV34563",
		RelatedDate:       fedwire.ISODate(civil.DateOf(time.Now())),
		TaxDetail: TaxRecord{
			TaxId:              "123456789",
			TaxTypeCode:        "09455",
			TaxPeriodYear:      fedwire.ISODate(civil.DateOf(time.Now())),
			TaxperiodTimeFrame: "MM04",
		},
	}
	mesage.RelatedRemittanceInfo = RemittanceDetail{
		RemittanceId:      "Scenario01Var2RemittanceId001",
		Method:            models.Email,
		ElectronicAddress: "CustomerService@CorporationB.com",
	}
	return mesage
}
