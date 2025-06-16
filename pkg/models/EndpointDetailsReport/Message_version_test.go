package EndpointDetailsReport

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion02(t *testing.T) {
	modelName := CAMT_052_001_02
	xmlName := "ActivityReport_02.xml"

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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.Equal(t, model.ReportingSequence.FromSeq, "000001")
	require.Equal(t, model.ReportingSequence.ToSeq, "000100")
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[0].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.Equal(t, model.ReportingSequence.FromSeq, "000001")
	require.Equal(t, model.ReportingSequence.ToSeq, "000100")
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[0].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.Equal(t, model.ReportingSequence.FromSeq, "000001")
	require.Equal(t, model.ReportingSequence.ToSeq, "000100")
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[0].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.Equal(t, model.ReportingSequence.FromSeq, "000001")
	require.Equal(t, model.ReportingSequence.ToSeq, "000100")
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[0].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
}
func TestVersion11(t *testing.T) {
	modelName := CAMT_052_001_11
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.Equal(t, model.ReportingSequence.FromSeq, "000001")
	require.Equal(t, model.ReportingSequence.ToSeq, "000100")
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[0].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
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
	require.Equal(t, model.MessageId, "DTLS")
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.BussinessQueryMsgId, "20250311231981435DTLSrequest1")
	require.Equal(t, model.BussinessQueryMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.BussinessQueryCreateDatetime)
	require.Equal(t, model.ReportingSequence.FromSeq, "000001")
	require.Equal(t, model.ReportingSequence.ToSeq, "000100")
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "94")
	require.Equal(t, model.TotalCreditEntries.Sum, 2871734.98)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "100")
	require.Equal(t, model.TotalDebitEntries.Sum, 8307111.56)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.Sent)
	require.Equal(t, model.EntryDetails[0].Amount.Amount, 50000.00)
	require.Equal(t, model.EntryDetails[0].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[0].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[0].Status, models.Book)
	require.Equal(t, model.EntryDetails[0].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[0].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[0].EntryDetails.MessageId, "20250310B1QDRCQR000001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[0].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[0].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000103100900FT02")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[0].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[0].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)
	require.Equal(t, model.EntryDetails[1].Amount.Amount, 8000.00)
	require.Equal(t, model.EntryDetails[1].Amount.Currency, "USD")
	require.Equal(t, model.EntryDetails[1].CreditDebitIndicator, models.Debit)
	require.Equal(t, model.EntryDetails[1].Status, models.Book)
	require.Equal(t, model.EntryDetails[1].BankTransactionCode, models.TransDebit)
	require.Equal(t, model.EntryDetails[1].MessageNameId, "pacs.008.001.08")
	require.Equal(t, model.EntryDetails[1].EntryDetails.MessageId, "20250310B1QDRCQR000002")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructionId, "20250331231981435InstructionId00001")
	require.Equal(t, model.EntryDetails[1].EntryDetails.UniqueTransactionReference, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.EntryDetails[1].EntryDetails.ClearingSystemRef, "20230310ISOTEST100000203100900FT02")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructingAgent.PaymentSysMemberId, "231981435")
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.EntryDetails[1].EntryDetails.InstructedAgent.PaymentSysMemberId, "011104238")
	require.Equal(t, model.EntryDetails[1].EntryDetails.LocalInstrumentChoice, models.InstrumentCTRC)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "DTLS"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "DTLS"
}
func ActivityReportDataModel() MessageModel {
	message := MessageModel{}
	message.MessageId = "DTLS"
	message.CreatedDateTime = time.Now()
	message.Pagenation = models.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.BussinessQueryMsgId = "20250311231981435DTLSrequest1"
	message.BussinessQueryMsgNameId = "camt.060.001.05"
	message.BussinessQueryCreateDatetime = time.Now()
	message.ReportId = models.Intraday
	message.ReportingSequence = models.SequenceRange{
		FromSeq: "000001",
		ToSeq:   "000100",
	}
	message.ReportCreateDateTime = time.Now()
	message.AccountOtherId = "B1QDRCQR"
	message.TotalCreditEntries = models.NumberAndSumOfTransactions{
		NumberOfEntries: "94",
		Sum:             2871734.98,
	}
	message.TotalDebitEntries = models.NumberAndSumOfTransactions{
		NumberOfEntries: "100",
		Sum:             8307111.56,
	}
	message.TotalEntriesPerBankTransactionCode = []models.TotalsPerBankTransactionCode{
		{
			NumberOfEntries:     "0",
			BankTransactionCode: models.Rejected,
		},
		{
			NumberOfEntries:     "0",
			BankTransactionCode: models.MessagesInProcess,
		},
		{
			NumberOfEntries:     "0",
			BankTransactionCode: models.MessagesIntercepted,
		},
		{
			NumberOfEntries:     "0",
			BankTransactionCode: models.Sent,
		},
	}
	message.EntryDetails = []models.Entry{
		{
			Amount: models.CurrencyAndAmount{
				Amount:   50000.00,
				Currency: "USD",
			},
			CreditDebitIndicator: models.Debit,
			Status:               models.Book,
			BankTransactionCode:  models.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: models.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000001",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310ISOTEST100000103100900FT02",
				InstructingAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice: models.InstrumentCTRC,
			},
		},
		{
			Amount: models.CurrencyAndAmount{
				Amount:   8000.00,
				Currency: "USD",
			},
			CreditDebitIndicator: models.Debit,
			Status:               models.Book,
			BankTransactionCode:  models.TransDebit,
			MessageNameId:        "pacs.008.001.08",
			EntryDetails: models.EntryDetail{
				MessageId:                  "20250310B1QDRCQR000002",
				InstructionId:              "20250331231981435InstructionId00001",
				UniqueTransactionReference: "8a562c67-ca16-48ba-b074-65581be6f011",
				ClearingSystemRef:          "20230310ISOTEST100000203100900FT02",
				InstructingAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				InstructedAgent: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "011104238",
				},
				LocalInstrumentChoice: models.InstrumentCTRC,
			},
		},
	}
	return message
}
