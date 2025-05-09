package ReturnRequestResponse

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestRequireField(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	cErr := message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: AssignmentId, Assigner, Assignee, AssignmentCreateTime, ResolvedCaseId, Creator, OriginalMessageId, OriginalMessageNameId, OriginalMessageCreateTime, OriginalUETR")
}
func generateRequreFields(msg Message) Message {
	if isEmpty(msg.data.AssignmentId) {
		msg.data.AssignmentId = "20250310B1QDRCQR000723"
	}
	if isEmpty(msg.data.Assigner) {
		msg.data.Assigner = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	if isEmpty(msg.data.Assignee) {
		msg.data.Assignee = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if msg.data.AssignmentCreateTime.IsZero() {
		msg.data.AssignmentCreateTime = time.Now()
	}
	if msg.data.ResolvedCaseId == "" {
		msg.data.ResolvedCaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	}
	if isEmpty(msg.data.Creator) {
		msg.data.Creator = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if msg.data.OriginalMessageId == "" {
		msg.data.OriginalMessageId = "20250310B1QDRCQR000721"
	}
	if msg.data.OriginalMessageNameId == "" {
		msg.data.OriginalMessageNameId = "pacs.008.001.08"
	}
	if msg.data.OriginalMessageCreateTime.IsZero() {
		msg.data.OriginalMessageCreateTime = time.Now()
	}
	if msg.data.OriginalUETR == "" {
		msg.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	}
	return msg
}
func TestReturnRequestResponseFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step3_camt.029")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250310B1QDRCQR000723", string(message.doc.RsltnOfInvstgtn.Assgnmt.Id))
	require.Equal(t, "USABA", string(*message.doc.RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "021040078", string(message.doc.RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "USABA", string(*message.doc.RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "011104238", string(message.doc.RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "20250310011104238Sc01Step1MsgIdDUPL", string(message.doc.RsltnOfInvstgtn.RslvdCase.Id))
	require.Equal(t, "Bank A", string(*message.doc.RsltnOfInvstgtn.RslvdCase.Cretr.Agt.FinInstnId.Nm))
	require.Equal(t, "CNCL", string(*message.doc.RsltnOfInvstgtn.Sts.Conf))
	require.Equal(t, "20250310B1QDRCQR000721", string(message.doc.RsltnOfInvstgtn.CxlDtls.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId))
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"
const INVALID_TRCOUNT string = "123456789012345"
const INVALID_MESSAGE_ID string = "12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456"
const INVALID_OTHER_ID string = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const INVALID_BUILD_NUM string = "12345678901234567"
const INVALID_POSTAL_CODE string = "12345678901234567"
const INVALID_COUNTRY_CODE string = "12345678"
const INVALID_MESSAGE_NAME_ID string = "sabcd-123-001-12"
const INVALID_PAY_SYSCODE model.PaymentSystemType = model.PaymentSystemType(INVALID_COUNT)

func TestReturnRequestResponseValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"Invalid AssignmentId",
			Message{data: MessageModel{AssignmentId: INVALID_OTHER_ID}},
			"error occur at AssignmentId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid Assigner",
			Message{data: MessageModel{Assigner: model.Agent{
				PaymentSysCode:     INVALID_PAY_SYSCODE,
				PaymentSysMemberId: "011104238",
			}}},
			"error occur at Assigner.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid Assignee",
			Message{data: MessageModel{Assignee: model.Agent{
				PaymentSysCode:     INVALID_PAY_SYSCODE,
				PaymentSysMemberId: "011104238",
			}}},
			"error occur at Assignee.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid ResolvedCaseId",
			Message{data: MessageModel{ResolvedCaseId: INVALID_OTHER_ID}},
			"error occur at ResolvedCaseId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid Creator",
			Message{data: MessageModel{Creator: model.Agent{
				PaymentSysCode:     model.PaymentSysUSABA,
				PaymentSysMemberId: "011104238",
				BankName:           "Bank A",
				PostalAddress: model.PostalAddress{
					StreetName:     "Avenue A",
					BuildingNumber: "66",
					PostalCode:     INVALID_POSTAL_CODE,
					TownName:       "Lisle",
					Subdivision:    "IL",
					Country:        "US",
				},
			}}},
			"error occur at Creator.PostalAddress.PostalCode: 12345678901234567 fails validation with length 17 <= required maxLength 16",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			nMsg := generateRequreFields(tt.msg)
			msgErr := nMsg.CreateDocument()
			if msgErr != nil {
				require.Equal(t, tt.expectedErr, msgErr.Error())
			}
		})
	}
}
func TestFedwireFundsAcknowledgement_Scenario2_Step3_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.data.AssignmentId = "20250310B1QDRCQR000723"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.Status = ReturnRequestAccepted
	message.data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step3_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step3_camt.029")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step3_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario2_Step3b_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.data.AssignmentId = "20250310B1QDRCQR000723"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.Status = ReturnRequestAccepted
	message.data.OriginalMessageId = "20250310B1QDRCQR000721"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step3b_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step3b_camt.029")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step3b_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentreturn_Scenario1_Step3_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.data.AssignmentId = "20250310B1QDRCQR000402"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc01Step1MsgIdDUPL"
	message.data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.Status = ReturnRequestAccepted
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Paymentreturn_Scenario1_Step3_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Paymentreturn_Scenario1_Step3_camt.029")
	genterated := filepath.Join("generated", "Paymentreturn_Scenario1_Step3_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentreturn_Scenario2_Step3_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.data.AssignmentId = "20250310B1QDRCQR000422"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc02Step1MsgIdSVNR"
	message.data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.Status = ReturnRequestRejected
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario02InstrId001"
	message.data.OriginalEndToEndId = "Scenario02EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.CancellationStatusReasonInfo = Reason{
		Reason:         "NARR",
		AdditionalInfo: "Corporation B delivered goods and services are in-line with client’s order.",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Paymentreturn_Scenario2_Step3_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Paymentreturn_Scenario2_Step3_camt.029")
	genterated := filepath.Join("generated", "Paymentreturn_Scenario2_Step3_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestPaymentreturn_Scenario3_Step3_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.data.AssignmentId = "20250310B1QDRCQR000432"
	message.data.Assigner = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	message.data.Assignee = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.AssignmentCreateTime = time.Now()
	message.data.ResolvedCaseId = "20250310011104238Sc03Step1MsgIdSVNR"
	message.data.Creator = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
		BankName:           "Bank A",
		PostalAddress: model.PostalAddress{
			StreetName:     "Avenue A",
			BuildingNumber: "66",
			PostalCode:     "60532",
			TownName:       "Lisle",
			Subdivision:    "IL",
			Country:        "US",
		},
	}
	message.data.Status = PartiallyExecutedReturn
	message.data.OriginalMessageId = "20250310B1QDRCQR000400"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario03InstrId001"
	message.data.OriginalEndToEndId = "Scenario03EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.CancellationStatusReasonInfo = Reason{
		Reason:         "NARR",
		AdditionalInfo: "As agreed, partial refund of 20% will be paid for service shortcomings.",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Paymentreturn_Scenario3_Step3_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Paymentreturn_Scenario3_Step3_camt.029")
	genterated := filepath.Join("generated", "Paymentreturn_Scenario3_Step3_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
