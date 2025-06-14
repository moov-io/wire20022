package Master

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion02(t *testing.T) {
	modelName := CAMT_052_001_02
	xmlName := "Master_02.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion03(t *testing.T) {
	modelName := CAMT_052_001_03
	xmlName := "Master_03.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion04(t *testing.T) {
	modelName := CAMT_052_001_04
	xmlName := "Master_04.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion05(t *testing.T) {
	modelName := CAMT_052_001_05
	xmlName := "Master_05.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion06(t *testing.T) {
	modelName := CAMT_052_001_06
	xmlName := "Master_06.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion07(t *testing.T) {
	modelName := CAMT_052_001_07
	xmlName := "Master_07.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].CdtLines[0].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[0].Type, models.NetDebitCap)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Amount, 23125500000.00)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[0].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[1].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[1].Type, models.CollateralizedCapacity)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Amount, 316874500000.00)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[1].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[2].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[2].Type, models.CollateralAvailable)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Amount, 82598573368.44)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[2].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[3].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[3].Type, models.CollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[3].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[4].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[4].Type, models.UncollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[4].DateTime)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.NumberOfEntries, "16281")
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.Sum, 420780358976.96)
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.NumberOfEntries, "22134")
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.Sum, 141184481554.24)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.NumberOfEntries, "4")
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.Sum, 993425694.01)
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.NumberOfEntries, "6")
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.Sum, 384826820.41)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion08(t *testing.T) {
	modelName := CAMT_052_001_08
	xmlName := "Master_08.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].CdtLines[0].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[0].Type, models.NetDebitCap)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Amount, 23125500000.00)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[0].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[1].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[1].Type, models.CollateralizedCapacity)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Amount, 316874500000.00)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[1].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[2].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[2].Type, models.CollateralAvailable)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Amount, 82598573368.44)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[2].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[3].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[3].Type, models.CollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[3].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[4].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[4].Type, models.UncollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[4].DateTime)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.NumberOfEntries, "16281")
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.Sum, 420780358976.96)
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.NumberOfEntries, "22134")
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.Sum, 141184481554.24)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.NumberOfEntries, "4")
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.Sum, 993425694.01)
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.NumberOfEntries, "6")
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.Sum, 384826820.41)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion09(t *testing.T) {
	modelName := CAMT_052_001_09
	xmlName := "Master_09.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].CdtLines[0].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[0].Type, models.NetDebitCap)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Amount, 23125500000.00)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[0].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[1].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[1].Type, models.CollateralizedCapacity)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Amount, 316874500000.00)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[1].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[2].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[2].Type, models.CollateralAvailable)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Amount, 82598573368.44)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[2].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[3].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[3].Type, models.CollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[3].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[4].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[4].Type, models.UncollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[4].DateTime)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.NumberOfEntries, "16281")
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.Sum, 420780358976.96)
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.NumberOfEntries, "22134")
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.Sum, 141184481554.24)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.NumberOfEntries, "4")
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.Sum, 993425694.01)
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.NumberOfEntries, "6")
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.Sum, 384826820.41)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion10(t *testing.T) {
	modelName := CAMT_052_001_10
	xmlName := "Master_10.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].CdtLines[0].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[0].Type, models.NetDebitCap)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Amount, 23125500000.00)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[0].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[1].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[1].Type, models.CollateralizedCapacity)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Amount, 316874500000.00)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[1].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[2].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[2].Type, models.CollateralAvailable)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Amount, 82598573368.44)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[2].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[3].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[3].Type, models.CollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[3].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[4].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[4].Type, models.UncollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[4].DateTime)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.NumberOfEntries, "16281")
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.Sum, 420780358976.96)
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.NumberOfEntries, "22134")
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.Sum, 141184481554.24)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.NumberOfEntries, "4")
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.Sum, 993425694.01)
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.NumberOfEntries, "6")
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.Sum, 384826820.41)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion11(t *testing.T) {
	modelName := CAMT_052_001_11
	xmlName := "Master_11.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].CdtLines[0].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[0].Type, models.NetDebitCap)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Amount, 23125500000.00)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[0].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[1].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[1].Type, models.CollateralizedCapacity)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Amount, 316874500000.00)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[1].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[2].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[2].Type, models.CollateralAvailable)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Amount, 82598573368.44)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[2].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[3].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[3].Type, models.CollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[3].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[4].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[4].Type, models.UncollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[4].DateTime)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.NumberOfEntries, "16281")
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.Sum, 420780358976.96)
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.NumberOfEntries, "22134")
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.Sum, 141184481554.24)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.NumberOfEntries, "4")
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.Sum, 993425694.01)
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.NumberOfEntries, "6")
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.Sum, 384826820.41)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func TestVersion12(t *testing.T) {
	modelName := CAMT_052_001_12
	xmlName := "Master_12.xml"

	dataModel := MasterDataModel()
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
	require.Equal(t, model.MessageId, models.AccountBalanceReport)
	require.NotNil(t, model.CreationDateTime)
	require.Equal(t, model.MessagePagination.PageNumber, "1")
	require.Equal(t, model.MessagePagination.LastPageIndicator, true)
	require.Equal(t, model.OriginalBusinessMsgId, "20230921231981435ABARMMrequest1")
	require.Equal(t, model.OriginalBusinessMsgNameId, "camt.060.001.05")
	require.NotNil(t, model.OriginalBusinessMsgCreateTime)
	require.Equal(t, model.ReportTypeId, models.ABMS)
	require.NotNil(t, model.ReportCreatedDate)
	require.Equal(t, model.AccountOtherId, "231981435")
	require.Equal(t, model.AccountType, "M")
	require.Equal(t, model.RelatedAccountOtherId, "231981435")
	require.Equal(t, model.Balances[0].BalanceTypeId, models.DaylightOverdraftBalance)
	require.Equal(t, model.Balances[0].Amount.Amount, 270458895930.79)
	require.Equal(t, model.Balances[0].Amount.Currency, "USD")
	require.Equal(t, model.Balances[0].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[0].DateTime)
	require.Equal(t, model.Balances[1].BalanceTypeId, models.AccountBalance)
	require.Equal(t, model.Balances[1].CdtLines[0].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[0].Type, models.NetDebitCap)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Amount, 23125500000.00)
	require.Equal(t, model.Balances[1].CdtLines[0].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[0].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[1].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[1].Type, models.CollateralizedCapacity)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Amount, 316874500000.00)
	require.Equal(t, model.Balances[1].CdtLines[1].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[1].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[2].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[2].Type, models.CollateralAvailable)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Amount, 82598573368.44)
	require.Equal(t, model.Balances[1].CdtLines[2].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[2].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[3].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[3].Type, models.CollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[3].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[3].DateTime)
	require.Equal(t, model.Balances[1].CdtLines[4].Included, true)
	require.Equal(t, model.Balances[1].CdtLines[4].Type, models.UncollateralizedDaylightOverdrafts)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Amount, 0.00)
	require.Equal(t, model.Balances[1].CdtLines[4].Amount.Currency, "USD")
	require.NotNil(t, model.Balances[1].CdtLines[4].DateTime)
	require.Equal(t, model.Balances[1].Amount.Amount, 270594506052.13)
	require.Equal(t, model.Balances[1].Amount.Currency, "USD")
	require.Equal(t, model.Balances[1].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[1].DateTime)
	require.Equal(t, model.Balances[2].BalanceTypeId, models.AvailableBalanceFromDaylightOverdraft)
	require.Equal(t, model.Balances[2].Amount.Amount, 610458895930.79)
	require.Equal(t, model.Balances[2].Amount.Currency, "USD")
	require.Equal(t, model.Balances[2].CreditDebitIndicator, models.Credit)
	require.NotNil(t, model.Balances[2].DateTime)
	require.Equal(t, model.TransactionsSummary[0].TotalNetEntryAmount, 279595877422.72)
	require.Equal(t, model.TransactionsSummary[0].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.NumberOfEntries, "16281")
	require.Equal(t, model.TransactionsSummary[0].CreditEntries.Sum, 420780358976.96)
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.NumberOfEntries, "22134")
	require.Equal(t, model.TransactionsSummary[0].DebitEntries.Sum, 141184481554.24)
	require.Equal(t, model.TransactionsSummary[0].BankTransactionCode, models.FedwireFundsTransfers)
	require.NotNil(t, model.TransactionsSummary[0].Date)
	require.Equal(t, model.TransactionsSummary[1].TotalNetEntryAmount, 608598873.60)
	require.Equal(t, model.TransactionsSummary[1].CreditDebitIndicator, models.Credit)
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.NumberOfEntries, "4")
	require.Equal(t, model.TransactionsSummary[1].CreditEntries.Sum, 993425694.01)
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.NumberOfEntries, "6")
	require.Equal(t, model.TransactionsSummary[1].DebitEntries.Sum, 384826820.41)
	require.Equal(t, model.TransactionsSummary[1].BankTransactionCode, models.NationalSettlementServiceEntries)
	require.NotNil(t, model.TransactionsSummary[1].Date)

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy BkToCstmrAcctRpt.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = models.AccountBalanceReport

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = models.AccountBalanceReport
}
func MasterDataModel() MessageModel {
	message := MessageModel{}
	message.MessageId = models.AccountBalanceReport
	message.CreationDateTime = time.Now()
	message.MessagePagination = models.MessagePagenation{
		PageNumber:        "1",
		LastPageIndicator: true,
	}
	message.OriginalBusinessMsgId = "20230921231981435ABARMMrequest1"
	message.OriginalBusinessMsgNameId = "camt.060.001.05"
	message.OriginalBusinessMsgCreateTime = time.Now()

	message.ReportTypeId = models.ABMS
	message.ReportCreatedDate = time.Now()
	message.AccountOtherId = "231981435"
	message.AccountType = "M"
	message.RelatedAccountOtherId = "231981435"

	message.Balances = []models.Balance{
		{
			BalanceTypeId: models.DaylightOverdraftBalance,
			Amount: models.CurrencyAndAmount{
				Amount:   270458895930.79,
				Currency: "USD",
			},
			CreditDebitIndicator: models.Credit,
			DateTime:             time.Now(),
		},
		{
			BalanceTypeId: models.AccountBalance,
			CdtLines: []models.CreditLine{
				{
					Included: true,
					Type:     models.NetDebitCap,
					Amount: models.CurrencyAndAmount{
						Amount:   23125500000.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     models.CollateralizedCapacity,
					Amount: models.CurrencyAndAmount{
						Amount:   316874500000.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     models.CollateralAvailable,
					Amount: models.CurrencyAndAmount{
						Amount:   82598573368.44,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     models.CollateralizedDaylightOverdrafts,
					Amount: models.CurrencyAndAmount{
						Amount:   0.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
				{
					Included: true,
					Type:     models.UncollateralizedDaylightOverdrafts,
					Amount: models.CurrencyAndAmount{
						Amount:   0.00,
						Currency: "USD",
					},
					DateTime: time.Now(),
				},
			},
			Amount: models.CurrencyAndAmount{
				Amount:   270594506052.13,
				Currency: "USD",
			},
			CreditDebitIndicator: models.Credit,
			DateTime:             time.Now(),
		},
		{
			BalanceTypeId: models.AvailableBalanceFromDaylightOverdraft,
			Amount: models.CurrencyAndAmount{
				Amount:   610458895930.79,
				Currency: "USD",
			},
			CreditDebitIndicator: models.Credit,
			DateTime:             time.Now(),
		},
	}
	message.TransactionsSummary = []models.TotalsPerBankTransaction{
		{
			TotalNetEntryAmount:  279595877422.72,
			CreditDebitIndicator: models.Credit,
			CreditEntries: models.NumberAndSumOfTransactions{
				NumberOfEntries: "16281",
				Sum:             420780358976.96,
			},
			DebitEntries: models.NumberAndSumOfTransactions{
				NumberOfEntries: "22134",
				Sum:             141184481554.24,
			},
			BankTransactionCode: models.FedwireFundsTransfers,
			Date:                time.Now(),
		},
		{
			TotalNetEntryAmount:  608598873.60,
			CreditDebitIndicator: models.Credit,
			CreditEntries: models.NumberAndSumOfTransactions{
				NumberOfEntries: "4",
				Sum:             993425694.01,
			},
			DebitEntries: models.NumberAndSumOfTransactions{
				NumberOfEntries: "6",
				Sum:             384826820.41,
			},
			BankTransactionCode: models.NationalSettlementServiceEntries,
			Date:                time.Now(),
		},
	}
	return message
}
