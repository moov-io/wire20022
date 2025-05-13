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
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreatedDateTime, OriginalMessageId, OriginalMessageNameId, OriginalCreationDateTime, OriginalUETR, InstructingAgent, InstructedAgent")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.MessageId == "" {
		msg.Data.MessageId = "20250310Scenario03Step2MsgId001"
	}
	if msg.Data.CreatedDateTime.IsZero() {
		msg.Data.CreatedDateTime = time.Now()
	}
	if msg.Data.OriginalMessageId == "" {
		msg.Data.OriginalMessageId = "20250310B1QDRCQR000001"
	}
	if msg.Data.OriginalMessageNameId == "" {
		msg.Data.OriginalMessageNameId = "pacs.008.001.08"
	}
	if msg.Data.OriginalCreationDateTime.IsZero() {
		msg.Data.OriginalCreationDateTime = time.Now()
	}
	if msg.Data.OriginalUETR == "" {
		msg.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	}
	if isEmpty(msg.Data.InstructingAgent) {
		msg.Data.InstructingAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	if isEmpty(msg.Data.InstructedAgent) {
		msg.Data.InstructedAgent = model.Agent{
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
	require.Equal(t, "20250310Scenario03Step2MsgId001", string(message.Doc.FIToFIPmtStsReq.GrpHdr.MsgId))
	require.Equal(t, "20250310B1QDRCQR000001", string(message.Doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgId))
	require.Equal(t, "pacs.008.001.08", string(message.Doc.FIToFIPmtStsReq.TxInf.OrgnlGrpInf.OrgnlMsgNmId))
	require.Equal(t, "Scenario01InstrId001", string(*message.Doc.FIToFIPmtStsReq.TxInf.OrgnlInstrId))
	require.Equal(t, "Scenario01EtoEId001", string(*message.Doc.FIToFIPmtStsReq.TxInf.OrgnlEndToEndId))
	require.Equal(t, "8a562c67-ca16-48ba-b074-65581be6f011", string(message.Doc.FIToFIPmtStsReq.TxInf.OrgnlUETR))
	require.Equal(t, "011104238", string(message.Doc.FIToFIPmtStsReq.TxInf.InstgAgt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "021151080", string(message.Doc.FIToFIPmtStsReq.TxInf.InstdAgt.FinInstnId.ClrSysMmbId.MmbId))
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
			Message{Data: MessageModel{MessageId: INVALID_OTHER_ID}},
			"error occur at MessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid OriginalMessageId",
			Message{Data: MessageModel{OriginalMessageId: INVALID_OTHER_ID}},
			"error occur at OriginalMessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid OriginalMessageNameId",
			Message{Data: MessageModel{OriginalMessageNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at OriginalMessageNameId: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"Invalid OriginalInstructionId",
			Message{Data: MessageModel{OriginalInstructionId: INVALID_OTHER_ID}},
			"error occur at OriginalInstructionId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid OriginalEndToEndId",
			Message{Data: MessageModel{OriginalEndToEndId: INVALID_OTHER_ID}},
			"error occur at OriginalEndToEndId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid InstructingAgent",
			Message{Data: MessageModel{InstructingAgent: model.Agent{
				PaymentSysCode:     INVALID_PAY_SYSCODE,
				PaymentSysMemberId: "011104238",
			}}},
			"error occur at InstructingAgent.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid InstructedAgent",
			Message{Data: MessageModel{InstructedAgent: model.Agent{
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
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310Scenario03Step2MsgId001"
	message.Data.CreatedDateTime = time.Now()
	message.Data.OriginalMessageId = "20250310B1QDRCQR000001"
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.OriginalCreationDateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario01InstrId001"
	message.Data.OriginalEndToEndId = "Scenario01EtoEId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step2_pacs.028")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario5_Step3_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310Scenario04Step3MsgId001"
	message.Data.CreatedDateTime = time.Now()
	message.Data.OriginalMessageId = "20250310B1QDRCQR000631"
	message.Data.OriginalMessageNameId = "pain.013.001.07"
	message.Data.OriginalCreationDateTime = time.Now()
	message.Data.OriginalInstructionId = "Scenario04Step1InstrId001"
	message.Data.OriginalEndToEndId = "Scenario4EndToEndId001"
	message.Data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f258"
	message.Data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	message.Data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021040078",
	}
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario5_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario5_Step3_pacs.028")
	genterated := filepath.Join("generated", "Drawdowns_Scenario5_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
