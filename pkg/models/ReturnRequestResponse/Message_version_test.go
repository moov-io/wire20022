package ReturnRequestResponse

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestVersion3(t *testing.T) {
	modelName := CAMT_029_001_03
	xmlName := "PaymentReturn_03.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}
func TestVersion4(t *testing.T) {
	modelName := CAMT_029_001_04
	xmlName := "PaymentReturn_04.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}
func TestVersion5(t *testing.T) {
	modelName := CAMT_029_001_05
	xmlName := "PaymentReturn_05.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}
func TestVersion6(t *testing.T) {
	modelName := CAMT_029_001_06
	xmlName := "PaymentReturn_06.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}
func TestVersion7(t *testing.T) {
	modelName := CAMT_029_001_07
	xmlName := "PaymentReturn_07.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}
func TestVersion8(t *testing.T) {
	modelName := CAMT_029_001_08
	xmlName := "PaymentReturn_08.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}
func TestVersion9(t *testing.T) {
	modelName := CAMT_029_001_09
	xmlName := "PaymentReturn_09.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}

func TestVersion10(t *testing.T) {
	modelName := CAMT_029_001_10
	xmlName := "PaymentReturn_10.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}

func TestVersion11(t *testing.T) {
	modelName := CAMT_029_001_11
	xmlName := "PaymentReturn_11.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}

func TestVersion12(t *testing.T) {
	modelName := CAMT_029_001_12
	xmlName := "PaymentReturn_12.xml"

	dataModel := ReturnRequestResponseDataModel()
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
	require.Equal(t, model.AssignmentId, "20250310B1QDRCQR000422")
	require.Equal(t, model.Assigner.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assigner.PaymentSysMemberId, "021040078")
	require.Equal(t, model.Assignee.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Assignee.PaymentSysMemberId, "011104238")
	require.NotNil(t, model.AssignmentCreateTime)
	require.Equal(t, model.ResolvedCaseId, "20250310011104238Sc02Step1MsgIdSVNR")
	require.Equal(t, model.Creator.PaymentSysCode, models.PaymentSysUSABA)
	require.Equal(t, model.Creator.PaymentSysMemberId, "011104238")
	require.Equal(t, model.Creator.BankName, "Bank A")
	require.Equal(t, model.Creator.PostalAddress.StreetName, "Avenue A")
	require.Equal(t, model.Creator.PostalAddress.BuildingNumber, "66")
	require.Equal(t, model.Creator.PostalAddress.PostalCode, "60532")
	require.Equal(t, model.Creator.PostalAddress.TownName, "Lisle")
	require.Equal(t, model.Creator.PostalAddress.Subdivision, "IL")
	require.Equal(t, model.Creator.PostalAddress.Country, "US")
	require.Equal(t, model.Status, models.ReturnRequestRejected)
	require.Equal(t, model.OriginalMessageId, "20250310B1QDRCQR000400")
	require.Equal(t, model.OriginalMessageNameId, "pacs.008.001.08")
	require.NotNil(t, model.OriginalMessageCreateTime)
	require.Equal(t, model.OriginalInstructionId, "Scenario02InstrId001")
	require.Equal(t, model.OriginalEndToEndId, "Scenario02EtoEId001")
	require.Equal(t, model.OriginalUETR, "8a562c67-ca16-48ba-b074-65581be6f011")
	require.Equal(t, model.CancellationStatusReasonInfo.Reason, "LEGL")
	require.Contains(t, model.CancellationStatusReasonInfo.AdditionalInfo, "Corporation B delivered goods")

	/*Validation check*/
	model.AssignmentId = "InvalideMessageIdLength5012345678901234567890"
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "field copy RsltnOfInvstgtn.Assgnmt.Id failed: failed to set AssignmentId: InvalideMessageIdLength5012345678901234567890 fails validation with length 45 <= required maxLength 35")
	model.AssignmentId = "20250310B1QDRCQR000422"

	/*Require field check*/
	model.AssignmentId = ""
	_, err = DocumentWith(*model, modelName)
	require.NotNil(t, err, "Expected error but got nil")
	require.Equal(t, err.Error(), "validation failed for field \"AssignmentId\": is required: required field missing")
	model.AssignmentId = "20250310B1QDRCQR000422"
}

func ReturnRequestResponseDataModel() MessageModel {
	message := MessageModel{}
	message.AssignmentId = "20250310B1QDRCQR000422"
	message.Assigner = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.Assignee = models.Agent{
		PaymentSysCode:     models.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.AssignmentCreateTime = time.Now()
	message.ResolvedCaseId = "20250310011104238Sc02Step1MsgIdSVNR"
	message.Creator = models.Agent{
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
	message.Status = models.ReturnRequestRejected
	message.OriginalMessageId = "20250310B1QDRCQR000400"
	message.OriginalMessageNameId = "pacs.008.001.08"
	message.OriginalMessageCreateTime = time.Now()
	message.OriginalInstructionId = "Scenario02InstrId001"
	message.OriginalEndToEndId = "Scenario02EtoEId001"
	message.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.CancellationStatusReasonInfo = models.Reason{
		Reason:         "LEGL",
		AdditionalInfo: "Corporation B delivered goods and services are in-line with client’s order.",
	}
	return message
}
