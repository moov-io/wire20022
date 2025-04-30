package PaymentStatusRequest

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreatedDateTime, OriginalMessageId, OriginalMessageNameId, OriginalCreationDateTime, OriginalUETR, InstructingAgent, InstructedAgent")
}
func generateRequreFields(msg Message) Message {
	if msg.data.MessageId == "" {
		msg.data.MessageId = "20250310Scenario03Step2MsgId001"
	}
	if msg.data.CreatedDateTime.IsZero() {
		msg.data.CreatedDateTime = time.Now()
	}
	if msg.data.OriginalMessageId == "" {
		msg.data.OriginalMessageId = "20250310B1QDRCQR000001"
	}
	if msg.data.OriginalMessageNameId == "" {
		msg.data.OriginalMessageNameId = "pacs.008.001.08"
	}
	if msg.data.OriginalCreationDateTime.IsZero() {
		msg.data.OriginalCreationDateTime = time.Now()
	}
	if msg.data.OriginalUETR == "" {
		msg.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	}
	if isEmpty(msg.data.InstructingAgent) {
		msg.data.InstructingAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if isEmpty(msg.data.InstructedAgent) {
		msg.data.InstructedAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021151080",
		}
	}
	return msg
}
func TestPaymentStatusRequestFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step2_pacs.028")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250310Scenario03Step2MsgId001", string(message.doc.FIToFIPmtStsReq.GrpHdr.MsgId))
	require.Equal(t, "20250310B1QDRCQR000001", string(message.doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgId))
	require.Equal(t, "pacs.008.001.08", string(message.doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgNmId))
	require.Equal(t, "Scenario01InstrId001", string(*message.doc.FIToFIPmtStsReq.TxInf.OrgnlInstrId))
	require.Equal(t, "Scenario01EtoEId001", string(*message.doc.FIToFIPmtStsReq.TxInf.OrgnlEndToEndId))
	require.Equal(t, "8a562c67-ca16-48ba-b074-65581be6f011", string(message.doc.FIToFIPmtStsReq.TxInf.OrgnlUETR))
	require.Equal(t, "011104238", string(message.doc.FIToFIPmtStsReq.TxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "021151080", string(message.doc.FIToFIPmtStsReq.TxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId))
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

func TestPaymentStatusRequestValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"Invalid MessageId",
			Message{data: MessageModel{MessageId: INVALID_OTHER_ID}},
			"error occur at MessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid OriginalMessageId",
			Message{data: MessageModel{OriginalMessageId: INVALID_OTHER_ID}},
			"error occur at OriginalMessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid OriginalMessageNameId",
			Message{data: MessageModel{OriginalMessageNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at OriginalMessageNameId: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"Invalid OriginalInstructionId",
			Message{data: MessageModel{OriginalInstructionId: INVALID_OTHER_ID}},
			"error occur at OriginalInstructionId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid OriginalEndToEndId",
			Message{data: MessageModel{OriginalEndToEndId: INVALID_OTHER_ID}},
			"error occur at OriginalEndToEndId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid InstructingAgent",
			Message{data: MessageModel{InstructingAgent: model.Agent{
				PaymentSysCode:     INVALID_PAY_SYSCODE,
				PaymentSysMemberId: "011104238",
			}}},
			"error occur at InstructingAgent.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid InstructedAgent",
			Message{data: MessageModel{InstructedAgent: model.Agent{
				PaymentSysCode:     INVALID_PAY_SYSCODE,
				PaymentSysMemberId: "011104238",
			}}},
			"error occur at InstructingAgent.PaymentSysCode: UNKNOWN fails enumeration validation",
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
func TestCustomerCreditTransfer_Scenario3_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310Scenario03Step2MsgId001"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000001"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario01InstrId001"
	message.data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step2_pacs.028")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario5_Step3_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310Scenario04Step3MsgId001"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000631"
	message.data.OriginalMessageNameId = "pain.013.001.07"
	message.data.OriginalCreationDateTime = time.Now()
	message.data.OriginalInstructionId = "Scenario04Step1InstrId001"
	message.data.OriginalEndToEndId = "Scenario4EndToEndId001"
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f258"
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario5_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario5_Step3_pacs.028")
	genterated := filepath.Join("generated", "Drawdowns_Scenario5_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
