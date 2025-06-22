package EndpointTotalsReport

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

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func TestVersion03(t *testing.T) {
	modelName := CAMT_052_001_03
	xmlName := "ActivityReport_03.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func TestVersion04(t *testing.T) {
	modelName := CAMT_052_001_04
	xmlName := "ActivityReport_04.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func TestVersion05(t *testing.T) {
	modelName := CAMT_052_001_05
	xmlName := "ActivityReport_05.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func TestVersion06(t *testing.T) {
	modelName := CAMT_052_001_06
	xmlName := "ActivityReport_06.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func TestVersion07(t *testing.T) {
	modelName := CAMT_052_001_07
	xmlName := "ActivityReport_07.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}

func TestVersion08(t *testing.T) {
	modelName := CAMT_052_001_08
	xmlName := "ActivityReport_08.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func TestVersion09(t *testing.T) {
	modelName := CAMT_052_001_09
	xmlName := "ActivityReport_09.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func TestVersion10(t *testing.T) {
	modelName := CAMT_052_001_10
	xmlName := "ActivityReport_10.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func TestVersion11(t *testing.T) {
	modelName := CAMT_052_001_11
	xmlName := "ActivityReport_11.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func TestVersion12(t *testing.T) {
	modelName := CAMT_052_001_12
	xmlName := "ActivityReport_12.xml"

	dataModel := EndpointTotalsReportDataModel()

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
	model, err := ParseXML(xmlDoc)
	if err != nil {
		t.Fatal(err)
	}
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.MessageId, models.EndpointTotalsReport)
	require.NotNil(t, model.CreatedDateTime)
	require.Equal(t, model.Pagenation.PageNumber, "1")
	require.Equal(t, model.Pagenation.LastPageIndicator, true)
	require.Equal(t, model.ReportId, models.Intraday)
	require.NotNil(t, model.ReportCreateDateTime)
	require.Equal(t, model.AccountOtherId, "B1QDRCQR")
	require.Equal(t, model.TotalCreditEntries.NumberOfEntries, "1268")
	require.Equal(t, model.TotalCreditEntries.Sum, 18423923492.15)
	require.Equal(t, model.TotalDebitEntries.NumberOfEntries, "4433")
	require.Equal(t, model.TotalDebitEntries.Sum, 12378489145.96)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].NumberOfEntries, "1")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[0].BankTransactionCode, models.Rejected)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[1].BankTransactionCode, models.MessagesIntercepted)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].NumberOfEntries, "0")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[2].BankTransactionCode, models.MessagesInProcess)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].NumberOfEntries, "27")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[3].BankTransactionCode, models.TransReceived)
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].NumberOfEntries, "193")
	require.Equal(t, model.TotalEntriesPerBankTransactionCode[4].BankTransactionCode, models.Sent)
	require.Contains(t, model.AdditionalReportInfo, "Next IMAD sequence number:")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.EndpointTotalsReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.EndpointTotalsReport
}
func EndpointTotalsReportDataModel() MessageModel {
	message := MessageModel{}
	message.MessageId = models.EndpointTotalsReport
	message.CreatedDateTime = time.Now()
	message.Pagenation = models.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.BussinessQueryMsgId = "BIZ12345"
	message.BussinessQueryMsgNameId = "BIZNAME001"
	message.BussinessQueryCreateDatetime = time.Now()
	message.ReportId = models.Intraday
	message.ReportingSequence = models.SequenceRange{
		FromSeq: "1",
		ToSeq:   "100",
	}
	message.ReportCreateDateTime = time.Now()
	message.AccountOtherId = "B1QDRCQR"
	message.TotalCreditEntries = models.NumberAndSumOfTransactions{
		NumberOfEntries: "1268",
		Sum:             18423923492.15,
	}
	message.TotalDebitEntries = models.NumberAndSumOfTransactions{
		NumberOfEntries: "4433",
		Sum:             12378489145.96,
	}
	message.TotalEntriesPerBankTransactionCode = []models.TotalsPerBankTransactionCode{
		{
			NumberOfEntries:     "1",
			BankTransactionCode: models.Rejected,
		},
		{
			NumberOfEntries:     "0",
			BankTransactionCode: models.MessagesIntercepted,
		},
		{
			NumberOfEntries:     "0",
			BankTransactionCode: models.MessagesInProcess,
		},
		{
			NumberOfEntries:     "27",
			BankTransactionCode: models.TransReceived,
		},
		{
			NumberOfEntries:     "193",
			BankTransactionCode: models.Sent,
		},
	}
	message.AdditionalReportInfo = "Next IMAD sequence number: 4627. Next OMAD sequence number: 1296. Count of missing IMAD sequence numbers: 0."
	return message
}
