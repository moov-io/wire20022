package FedwireFundsPaymentStatus

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreatedDateTime, OriginalMessageId, OriginalMessageNameId, OriginalMessageCreateTime, OriginalUETR, TransactionStatus, InstructingAgent, InstructedAgent")
}
func generateRequreFields(msg Message) Message {
	if msg.data.MessageId == "" {
		msg.data.MessageId = "20250310QMGFNP31000001"
	}
	if isEmpty(msg.data.CreatedDateTime) {
		msg.data.CreatedDateTime = time.Now()
	}
	if msg.data.OriginalMessageId == "" {
		msg.data.OriginalMessageId = "20250310B1QDRCQR000001"
	}
	if msg.data.OriginalMessageNameId == "" {
		msg.data.OriginalMessageNameId = "pacs.008.001.08"
	}
	if isEmpty(msg.data.OriginalMessageCreateTime) {
		msg.data.OriginalMessageCreateTime = time.Now()
	}
	if msg.data.OriginalUETR == "" {
		msg.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	}
	if msg.data.TransactionStatus == "" {
		msg.data.TransactionStatus = model.AcceptedSettlementCompleted
	}
	if isEmpty(msg.data.InstructingAgent) {
		msg.data.InstructingAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "021151080",
		}
	}
	if isEmpty(msg.data.InstructedAgent) {
		msg.data.InstructedAgent = model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		}
	}
	return msg
}
func TestFedwireFundsPaymentStatusFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step2_pacs.002")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250310QMGFNP31000001", string(message.doc.FIToFIPmtStsRpt.GrpHdr.MsgId))
	require.Equal(t, "20250310B1QDRCQR000001", string(message.doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId))
	require.Equal(t, "pacs.008.001.08", string(message.doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgNmId))
	require.Equal(t, "8a562c67-ca16-48ba-b074-65581be6f011", string(message.doc.FIToFIPmtStsRpt.TxInfAndSts.OrgnlUETR))
	require.Equal(t, "ACSC", string(message.doc.FIToFIPmtStsRpt.TxInfAndSts.TxSts))

	// Validate instructing and instructed agents
	require.Equal(t, "USABA", string(*message.doc.FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "021151080", string(message.doc.FIToFIPmtStsRpt.TxInfAndSts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId))
	require.Equal(t, "USABA", string(*message.doc.FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd))
	require.Equal(t, "011104238", string(message.doc.FIToFIPmtStsRpt.TxInfAndSts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId))
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

func TestFedwireFundsPaymentStatusValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"Invalid MessageId",
			Message{data: MessageModel{MessageId: INVALID_MESSAGE_ID}},
			"error occur at MessageId: 12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456 fails validation with length 66 <= required maxLength 35",
		},
		{
			"Invalid CreatedDateTime",
			Message{data: MessageModel{CreatedDateTime: time.Time{}}},
			"error occur at CreatedDateTime: invalid or empty date-time",
		},
		{
			"Invalid OriginalMessageId",
			Message{data: MessageModel{OriginalMessageId: INVALID_MESSAGE_ID}},
			"error occur at OriginalMessageId: 12345678abcdEFGH12345612345678abcdEFGH12345612345678abcdEFGH123456 fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid OriginalMessageNameId",
			Message{data: MessageModel{OriginalMessageNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at OriginalMessageId:  fails validation with pattern [0-9]{8}[A-Z0-9]{8}[0-9]{6}",
		},
		{
			"Invalid OriginalMessageCreateTime",
			Message{data: MessageModel{OriginalMessageCreateTime: time.Time{}}},
			"error occur at OriginalMessageCreateTime: invalid or empty date-time",
		},
		{
			"Invalid OriginalUETR",
			Message{data: MessageModel{OriginalUETR: "invalid-uetr"}},
			"error occur at OriginalUETR: invalid-uetr fails validation with pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}",
		},
		{
			"Invalid TransactionStatus",
			Message{data: MessageModel{TransactionStatus: "INVALID_STATUS"}},
			"error occur at TransactionStatus: INVALID_STATUS fails validation with length 14 <= required maxLength 4",
		},
		{
			"Invalid AcceptanceDateTime",
			Message{data: MessageModel{AcceptanceDateTime: time.Time{}}},
			"error occur at AcceptanceDateTime: invalid or empty date-time",
		},
		{
			"Invalid EffectiveInterbankSettlementDate",
			Message{data: MessageModel{EffectiveInterbankSettlementDate: model.Date{}}},
			"error occur at EffectiveInterbankSettlementDate: invalid or empty date",
		},
		{
			"Invalid StatusReasonInformation",
			Message{data: MessageModel{StatusReasonInformation: INVALID_OTHER_ID}},
			"error occur at StatusReasonInformation: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid InstructingAgent PaymentSysCode",
			Message{data: MessageModel{InstructingAgent: model.Agent{PaymentSysCode: INVALID_PAY_SYSCODE}}},
			"error occur at InstructingAgent.PaymentSysCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid InstructingAgent PaymentSysMemberId",
			Message{data: MessageModel{InstructingAgent: model.Agent{PaymentSysMemberId: INVALID_OTHER_ID}}},
			"error occur at InstructingAgent.PaymentSysCode:  fails enumeration validation",
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
func TestCustomerCreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP31000001"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000001"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.AcceptanceDateTime = time.Now()
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario1_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step2_pacs.002")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "FDWA1B2C3D4E5F6G7H8I9J10K11L12M0"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000002"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.Rejected
	message.data.StatusReasonInformation = "E433"
	message.data.ReasonAdditionalInfo = "The routing number of the Instructed Agent is not permissible to receive Fedwire Funds transaction."
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario2_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario2_Step2_pacs.002")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario3_Step3_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP31000001"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000001"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step3_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step3_pacs.002")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step3_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP31000002"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000004"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "011104238",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario4_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario4_Step2_pacs.002")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP31000003"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000005"
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f011"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario5_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario5_Step2_pacs.002")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario1_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP62000501"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000501"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario1_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step2_pacs.002")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario1_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP62000502"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000502"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario2_Step2_pacs.002")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario3_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP62000503"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000503"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario3_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario3_Step2_pacs.002")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario3_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario4_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP62000504"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000504"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario4_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario4_Step2_pacs.002")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario4_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario5_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP62000505"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310QMGFNP62000505"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "122240120",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario5_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario5_Step2_pacs.002")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario5_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario6_Step2_pacs_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250310QMGFNP62000506"
	message.data.CreatedDateTime = time.Now()
	message.data.OriginalMessageId = "20250310B1QDRCQR000506"
	message.data.OriginalMessageNameId = "pacs.009.001.08"
	message.data.OriginalMessageCreateTime = time.Now()
	message.data.OriginalUETR = "8a562c67-ca16-48ba-b074-65581be6f055"
	message.data.TransactionStatus = model.AcceptedSettlementCompleted
	message.data.AcceptanceDateTime = time.Now()
	message.data.EffectiveInterbankSettlementDate = model.FromTime(time.Now())
	message.data.InstructingAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021151080",
	}
	message.data.InstructedAgent = model.Agent{
		PaymentSysCode:     model.PaymentSysUSABA,
		PaymentSysMemberId: "021307481",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario6_Step2_pacs.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario6_Step2_pacs.002")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario6_Step2_pacs.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
