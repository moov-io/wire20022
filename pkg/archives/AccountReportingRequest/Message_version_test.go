package ArchiveAccountReportingRequest

import (
	"encoding/xml"
	"testing"
	"time"

	Archive "github.com/moov-io/wire20022/pkg/archives"
	"github.com/stretchr/testify/require"
)

var AccountReportingRequestDataModel = MessageModel{
	MessageId:          "20250311231981435ABARMMrequest1",
	CreatedDateTime:    time.Now(),
	ReportRequestId:    EndpointDetailsSentReport,
	RequestedMsgNameId: "camt.052.001.08",
	AccountOtherId:     "231981435",
	AccountProperty:    AccountTypeMerchant,
	AccountOwnerAgent: Agent{
		PaymentSysCode:     PaymentSysUSABA,
		PaymentSysMemberId: "231981435",
		OtherTypeId:        "B1QDRCQR",
	},
	FromToSequence: SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000100",
	},
}

func TestVersion02(t *testing.T) {
	/*Create Document from Model*/
	var doc02, err = DocumentWith(AccountReportingRequestDataModel, "camt.060.001.02")
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc02.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc02, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLTo("AccountReportingRequest_02.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/AccountReportingRequest_02.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
}
func TestVersion03(t *testing.T) {
	/*Create Document from Model*/
	var doc03, err = DocumentWith(AccountReportingRequestDataModel, "camt.060.001.03")
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc03.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc03, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLTo("AccountReportingRequest_03.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/AccountReportingRequest_03.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
}
func TestVersion04(t *testing.T) {
	/*Create Document from Model*/
	var doc04, err = DocumentWith(AccountReportingRequestDataModel, "camt.060.001.04")
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc04.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc04, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLTo("AccountReportingRequest_04.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/AccountReportingRequest_04.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")
}
func TestVersion05(t *testing.T) {
	/*Create Document from Model*/
	var doc05, err = DocumentWith(AccountReportingRequestDataModel, "camt.060.001.05")
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc05.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc05, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLTo("AccountReportingRequest_05.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/AccountReportingRequest_05.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")
}
func TestVersion06(t *testing.T) {
	/*Create Document from Model*/
	var doc06, err = DocumentWith(AccountReportingRequestDataModel, "camt.060.001.06")
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc06.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc06, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLTo("AccountReportingRequest_06.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/AccountReportingRequest_06.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")
}
func TestVersion07(t *testing.T) {
	/*Create Document from Model*/
	var doc07, err = DocumentWith(AccountReportingRequestDataModel, "camt.060.001.07")
	require.NoError(t, err, "Failed to create document")
	/*Validate Check for created Document*/
	vErr := doc07.Validate()
	require.NoError(t, vErr, "Failed to validate document")
	/*Create XML file from Document*/
	xmlData, err := xml.MarshalIndent(doc07, "", "  ")
	require.NoError(t, err)
	err = Archive.WriteXMLTo("AccountReportingRequest_07.xml", xmlData)
	require.NoError(t, err)

	/*Create Date Model from XML (Read XML)*/
	var xmlDoc, xmlErr = Archive.ReadXMLFile("./generated/AccountReportingRequest_07.xml")
	require.NoError(t, xmlErr, "Failed to read XML file")

	/*Compare*/
	model, err := MessageWith(xmlDoc)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, "20250311231981435ABARMMrequest1", model.MessageId, "Failed to get MessageId")
	require.NotNil(t, model.CreatedDateTime, "Failed to get CreatedDateTime")
	require.Equal(t, CAMTReportType("DTLS"), model.ReportRequestId, "Failed to get MessageId")
	require.Equal(t, "camt.052.001.08", model.RequestedMsgNameId, "Failed to get RequestedMsgNameId")
	require.Equal(t, "231981435", model.AccountOtherId, "Failed to get AccountOtherId")
	require.Equal(t, AccountTypeFRS("M"), model.AccountProperty, "Failed to get AccountProperty")
	require.Equal(t, PaymentSystemType("USABA"), model.AccountOwnerAgent.PaymentSysCode, "Failed to get AccountOwnerAgent.PaymentSysCode")
	require.Equal(t, "231981435", model.AccountOwnerAgent.PaymentSysMemberId, "Failed to get AccountOwnerAgent.PaymentSysMemberId")
	require.Equal(t, "B1QDRCQR", model.AccountOwnerAgent.OtherTypeId, "Failed to get AccountOwnerAgent.OtherTypeId")
	require.Equal(t, "000002", model.FromToSequence.FromSeq, "Failed to get FromToSequence.FromSeq")
	require.Equal(t, "000100", model.FromToSequence.ToSeq, "Failed to get FromToSequence.ToSeq")
}
