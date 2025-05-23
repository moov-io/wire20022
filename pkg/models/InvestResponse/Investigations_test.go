package InvestResponse

import (
	"encoding/xml"
	"path/filepath"
	"testing"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestRequireField(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	cErr := message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, InvestigationStatus, InvestRequestMessageId, InvestigationType, Requestor, Responder")
}
func generateRequreFields(msg Message) Message {
	// Check required fields and append missing ones to ParamNames
	if msg.Data.MessageId == "" {
		msg.Data.MessageId = "20250310B1QDRCQR000901"
	}
	if msg.Data.InvestigationStatus == "" {
		msg.Data.InvestigationStatus = "CLSD"
	}
	if msg.Data.InvestRequestMessageId == "" {
		msg.Data.InvestRequestMessageId = "20250310QMGFT015000901"
	}
	if msg.Data.InvestigationType == "" {
		msg.Data.InvestigationType = "UTAP"
	}
	if isEmpty(msg.Data.Requestor) {
		msg.Data.Requestor = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if isEmpty(msg.Data.Responder) {
		msg.Data.Responder = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		}
	}
	return msg
}
func TestInvestResponseFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "Investigations_Scenario1_Step3_camt.111")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250310B1QDRCQR000901", string(message.Doc.InvstgtnRspn.InvstgtnRspn.MsgId))
	require.Equal(t, "CLSD", string(message.Doc.InvstgtnRspn.InvstgtnRspn.InvstgtnSts.Sts))
	require.Equal(t, "20250310QMGFT015000901", string(message.Doc.InvstgtnRspn.OrgnlInvstgtnReq.MsgId))
	require.Equal(t, "UTAP", string(*message.Doc.InvstgtnRspn.OrgnlInvstgtnReq.InvstgtnTp.Cd))
	require.Equal(t, "USABA", string(*message.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "011104238", string(message.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rqstr.Agt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "USABA", string(*message.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "021040078", string(message.Doc.InvstgtnRspn.OrgnlInvstgtnReq.Rspndr.Agt.FinInstnId.ClrSysMmbId.MmbId))

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

func TestInvestResponseValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"Invalid MessageId",
			Message{Data: MessageModel{MessageId: INVALID_OTHER_ID}},
			"error occur at MessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid InvestigationStatus",
			Message{Data: MessageModel{InvestigationStatus: INVALID_MESSAGE_NAME_ID}},
			"error occur at InvestigationStatus: sabcd-123-001-12 fails validation with length 16 <= required maxLength 4",
		},
		{
			"Invalid InvestRequestMessageId",
			Message{Data: MessageModel{InvestRequestMessageId: INVALID_OTHER_ID}},
			"error occur at InvestRequestMessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid InvestigationType",
			Message{Data: MessageModel{InvestigationType: INVALID_MESSAGE_NAME_ID}},
			"error occur at InvestigationType: sabcd-123-001-12 fails validation with length 16 <= required maxLength 4",
		},
		{
			"Invalid Requestor",
			Message{Data: MessageModel{Requestor: model.Agent{
				PaymentSysCode:     model.PaymentSysUSABA,
				PaymentSysMemberId: INVALID_MESSAGE_NAME_ID,
			}}},
			"error occur at Requestor.PaymentSysMemberId: sabcd-123-001-12 fails validation with pattern [0-9]{9,9}",
		},
		{
			"Invalid Responder",
			Message{Data: MessageModel{Responder: model.Agent{
				PaymentSysCode:     model.PaymentSysUSABA,
				PaymentSysMemberId: INVALID_MESSAGE_NAME_ID,
			}}},
			"error occur at Responder.PaymentSysMemberId: sabcd-123-001-12 fails validation with pattern [0-9]{9,9}",
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

func TestInvestigations_Scenario1_Step3_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000901"
	message.Data.InvestigationStatus = "CLSD"
	message.Data.InvestigationData = "Please correct Creditor Account number. It should be 567876543."
	message.Data.InvestRequestMessageId = "20250310QMGFT015000901"
	message.Data.InvestigationType = "UTAP"
	message.Data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario1_Step3_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario1_Step3_camt.111")
	genterated := filepath.Join("generated", "Investigations_Scenario1_Step3_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario2_Step3_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000902"
	message.Data.InvestigationStatus = "CLSD"
	message.Data.InvestigationData = "Payment is a duplicate. Please consider VOID. Return request will follow."
	message.Data.InvestRequestMessageId = "20250310QMGFT015000902"
	message.Data.InvestigationType = "OTHR"
	message.Data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario2_Step3_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario2_Step3_camt.111")
	genterated := filepath.Join("generated", "Investigations_Scenario2_Step3_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario3_Step3_camt_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310B1QDRCQR000903"
	message.Data.InvestigationStatus = "CLSD"
	message.Data.InvestigationData = "Remittance information was sent separately. Email: AccountsReceivable@CorporationB.com"
	message.Data.InvestRequestMessageId = "20250310QMGFT015000903"
	message.Data.InvestigationType = "RQFI"
	message.Data.Requestor = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.Responder = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario3_Step3_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario3_Step3_camt.111")
	genterated := filepath.Join("generated", "Investigations_Scenario3_Step3_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
