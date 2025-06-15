package AccountReportingRequest

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

var AccountReportingRequestDataModel = MessageModel{
	MessageHeader: base.MessageHeader{
		MessageId:       "20250311231981435ABARMMrequest1",
		CreatedDateTime: time.Now(),
	},
	ReportRequestId:    models.EndpointDetailsSentReport,
	RequestedMsgNameId: "camt.052.001.08",
	AccountOtherId:     "231981435",
	AccountProperty:    models.AccountTypeMerchant,
	AccountOwnerAgent: models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
		OtherTypeId:        "B1QDRCQR",
	},
	FromToSequence: models.SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000100",
	},
}

func TestVersion02(t *testing.T) {
	/*Create Document from Model*/
	var doc02, err = DocumentWith(AccountReportingRequestDataModel, CAMT_060_001_02)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc02.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc02, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate("AccountReportingRequest_02.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/AccountReportingRequest_02.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, models.CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, models.AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, models.PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, CAMT_060_001_02)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy AcctRptgReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, CAMT_060_001_02)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by the account owner or the party authorised", "Failed to get MessageId Documentation")
}
func TestVersion03(t *testing.T) {
	/*Create Document from Model*/
	var doc03, err = DocumentWith(AccountReportingRequestDataModel, CAMT_060_001_03)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate("AccountReportingRequest_03.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/AccountReportingRequest_03.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, models.CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, models.AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, models.PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, CAMT_060_001_03)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy AcctRptgReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, CAMT_060_001_03)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by the account owner or the party authorised", "Failed to get MessageId Documentation")
}
func TestVersion04(t *testing.T) {
	/*Create Document from Model*/
	var doc04, err = DocumentWith(AccountReportingRequestDataModel, CAMT_060_001_04)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc04.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc04, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate("AccountReportingRequest_04.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/AccountReportingRequest_04.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, models.CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, models.AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, models.PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, CAMT_060_001_04)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy AcctRptgReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, CAMT_060_001_04)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by the account owner or the party authorised", "Failed to get MessageId Documentation")
}
func TestVersion05(t *testing.T) {
	/*Create Document from Model*/
	var doc05, err = DocumentWith(AccountReportingRequestDataModel, CAMT_060_001_05)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc05.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc05, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate("AccountReportingRequest_05.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/AccountReportingRequest_05.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, models.CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, models.AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, models.PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, CAMT_060_001_05)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy AcctRptgReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, CAMT_060_001_05)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by the account owner or the party authorised", "Failed to get MessageId Documentation")
}
func TestVersion06(t *testing.T) {
	/*Create Document from Model*/
	var doc06, err = DocumentWith(AccountReportingRequestDataModel, CAMT_060_001_06)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc06.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc06, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate("AccountReportingRequest_06.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/AccountReportingRequest_06.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, models.CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, models.AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, models.PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, CAMT_060_001_06)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy AcctRptgReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, CAMT_060_001_06)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by the account owner or the party authorised", "Failed to get MessageId Documentation")
}
func TestVersion07(t *testing.T) {
	/*Create Document from Model*/
	var doc07, err = DocumentWith(AccountReportingRequestDataModel, CAMT_060_001_07)
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc07.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc07, "", "  ")
	require.NoError(t, err)
	err = models.WriteXMLToGenerate("AccountReportingRequest_07.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = models.ReadXMLFile("./generated/AccountReportingRequest_07.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, models.CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, models.AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, models.PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")

	/*Validation check*/
	model.MessageId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(model, CAMT_060_001_07)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy AcctRptgReq.GrpHdr.MsgId failed: failed to set MessageId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Require field check*/
	model.MessageId = ""
	_, err = DocumentWith(model, CAMT_060_001_07)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
	model.MessageId = "20250311231981435ABARMMrequest1"

	/*Access to Helper*/
	require.Equal(t, "Message Identification", BuildMessageHelper().MessageId.Title, "Failed to get MessageId Title")
	require.Equal(t, "Max35Text (based on string) minLength: 1 maxLength: 35", BuildMessageHelper().MessageId.Type, "Failed to get MessageId Type")
	require.Contains(t, BuildMessageHelper().MessageId.Documentation, "Point to point reference, as assigned by the account owner or the party authorised", "Failed to get MessageId Documentation")
}
