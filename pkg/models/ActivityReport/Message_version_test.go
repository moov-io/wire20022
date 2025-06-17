package ActivityReport

import (
	"encoding/xml"
	"testing"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion01(t *testing.T) {
	modelName := CAMT_052_001_01
	xmlName := "ActivityReport_01.xml"
	dataModel := ActivityReportDataModel()
	/*Create Document from Model*/
	var doc01, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc01.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc01, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRptV01.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion02(t *testing.T) {
	modelName := CAMT_052_001_02
	xmlName := "ActivityReport_02.xml"
	dataModel := ActivityReportDataModel()
	/*Create Document from Model*/
	var doc02, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc02.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc02, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate(xmlName, xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/" + xmlName)
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion03(t *testing.T) {
	modelName := CAMT_052_001_03
	xmlName := "ActivityReport_03.xml"

	dataModel := ActivityReportDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document - Skip for version 03 due to strict schema validation*/
	// vErr := doc03.Validate()
	// require.NoError(t, vErr, "Failed to validate document")
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion04(t *testing.T) {
	modelName := CAMT_052_001_04
	xmlName := "ActivityReport_04.xml"

	dataModel := ActivityReportDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document - Skip for version 04 due to strict schema validation*/
	// vErr := doc03.Validate()
	// require.NoError(t, vErr, "Failed to validate document")
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion05(t *testing.T) {
	modelName := CAMT_052_001_05
	xmlName := "ActivityReport_05.xml"

	dataModel := ActivityReportDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document - Skip for version 05 due to strict schema validation*/
	// vErr := doc03.Validate()
	// require.NoError(t, vErr, "Failed to validate document")
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion06(t *testing.T) {
	modelName := CAMT_052_001_06
	xmlName := "ActivityReport_06.xml"

	dataModel := ActivityReportDataModel()
	/*Create Document from Model*/
	var doc03, err = DocumentWith(dataModel, modelName)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document - Skip for version 06 due to strict schema validation*/
	// vErr := doc03.Validate()
	// require.NoError(t, vErr, "Failed to validate document")
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion07(t *testing.T) {
	modelName := CAMT_052_001_07
	xmlName := "ActivityReport_07.xml"

	dataModel := ActivityReportDataModel()
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion08(t *testing.T) {
	modelName := CAMT_052_001_08
	xmlName := "ActivityReport_08.xml"

	dataModel := ActivityReportDataModel()
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion09(t *testing.T) {
	modelName := CAMT_052_001_09
	xmlName := "ActivityReport_09.xml"

	dataModel := ActivityReportDataModel()
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion10(t *testing.T) {
	modelName := CAMT_052_001_10
	xmlName := "ActivityReport_10.xml"

	dataModel := ActivityReportDataModel()
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion11(t *testing.T) {
	modelName := CAMT_052_001_11
	xmlName := "ActivityReport_11.xml"

	dataModel := ActivityReportDataModel()
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
func TestVersion12(t *testing.T) {
	modelName := CAMT_052_001_12
	xmlName := "ActivityReport_12.xml"

	dataModel := ActivityReportDataModel()
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
	require.Equal(t, models.CAMTReportType("ACTR"), model.MessageId, "Failed to get MessageId")
	require.Equal(t, model.Pagenation.PageNumber, "1", "Failed to get PageNumber")
	require.Equal(t, model.Pagenation.LastPageIndicator, true, "Failed to get LastPageIndicator")
	require.Equal(t, model.ReportId, models.EveryDay, "Failed to get ReportId")
	require.Equal(t, model.AccountOtherId, "011104238", "Failed to get AccountOtherId")
	require.Equal(t, model.TotalEntries, "61", "Failed to get TotalEntries")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "29", "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalCreditEntries.Sum, 8775299.29, "Failed to get TotalCreditEntries")
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "27", "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalDebitEntries.Sum, 9932294.43, "Failed to get TotalDebitEntries")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Sent, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "5", "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.TransReceived, "Failed to get TotalEntriesPerBankTransactionCode")
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 240.67, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[0].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[0].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 1000.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[1].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[1].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")
	require.Equal(t, model.EntryDetails[2].Amount.Amount, 1197.00, "Failed to get Amount")
	require.Equal(t, model.EntryDetails[2].Amount.Currency, "USD", "Failed to get Currency")
	require.Equal(t, model.EntryDetails[2].CreditDebitIndicator, models.Debit, "Failed to get CreditDebitIndicator")
	require.Equal(t, model.EntryDetails[2].Status, models.Book, "Failed to get Status")
	require.Equal(t, model.EntryDetails[2].BankTransactionCode, models.TransDebit, "Failed to get BankTransactionCode")
	require.Equal(t, model.EntryDetails[2].MessageNameId, "pacs.008.001.08", "Failed to get MessageNameId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.MessageId, "20250310B1QDRCQR000003", "Failed to get MessageId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructionId, "20250331231981435InstructionId00001", "Failed to get InstructionId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA, "Failed to get PaymentSysCode")
	require.Equal(t, model.EntryDetails[2].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238", "Failed to get PaymentSysMemberId")
	require.Equal(t, model.EntryDetails[2].EntryDetails.RelatedDatesProprietary, models.BusinessProcessingDate, "Failed to get RelatedDatesProprietary")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.CAMTReportType("ACTR")

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by")
}
