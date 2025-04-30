package RetrievalRequest

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreatedDateTime, RequestType, BusinessDate, RecipientId")
}
func generateRequreFields(msg Message) Message {
	if msg.data.MessageId == "" {
		msg.data.MessageId = "20250310Scenario03Step2MsgId001"
	}
	if msg.data.CreatedDateTime.IsZero() {
		msg.data.CreatedDateTime = time.Now()
	}
	if msg.data.RequestType == "" {
		msg.data.RequestType = model.RequestSent
	}
	if isEmpty(msg.data.BusinessDate) {
		msg.data.BusinessDate = model.FromTime(time.Now())
	}
	if msg.data.RecipientId == "" {
		msg.data.RecipientId = "B1QDRCQR"
	}
	return msg
}
func TestRetrievalRequestFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "MessageRetrieval_Scenario1_Step1_admi.006")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250301011104238MRSc1Step1MsgId", string(message.doc.RsndReq.MsgHdr.MsgId))
	require.Equal(t, "S", string(message.doc.RsndReq.MsgHdr.ReqTp.Prtry.Id))
	require.Equal(t, "pacs.008.001.08", string(*message.doc.RsndReq.RsndSchCrit.OrgnlMsgNmId))
	require.Equal(t, "20250310B1QDRCQR000001", string(*message.doc.RsndReq.RsndSchCrit.FileRef))
	require.Equal(t, "B1QDRCQR", string(message.doc.RsndReq.RsndSchCrit.Rcpt.Id.PrtryId.Id))
	require.Equal(t, "NA", string(message.doc.RsndReq.RsndSchCrit.Rcpt.Id.PrtryId.Issr))
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

func TestRetrievalRequestValidator(t *testing.T) {
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
			"Invalid RequestType",
			Message{data: MessageModel{RequestType: model.RequestType(INVALID_COUNT)}},
			"error occur at CreatedDateTime: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid OriginalMessageNameId",
			Message{data: MessageModel{OriginalMessageNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at OriginalMessageNameId: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"Invalid RecipientId",
			Message{data: MessageModel{RecipientId: INVALID_MESSAGE_NAME_ID}},
			"error occur at RecipientId: sabcd-123-001-12 fails validation with pattern [A-Z0-9]{8,8}",
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
func TestMessageRetrieval_Scenario1_Step1_admi_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250301011104238MRSc1Step1MsgId"
	message.data.CreatedDateTime = time.Now()
	message.data.RequestType = model.RequestSent
	message.data.BusinessDate = model.FromTime(time.Now())
	message.data.OriginalMessageNameId = "pacs.008.001.08"
	message.data.FileReference = "20250310B1QDRCQR000001"
	message.data.RecipientId = "B1QDRCQR"
	message.data.RecipientIssuer = "NA"

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("MessageRetrieval_Scenario1_Step1_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "MessageRetrieval_Scenario1_Step1_admi.006")
	genterated := filepath.Join("generated", "MessageRetrieval_Scenario1_Step1_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestMessageRetrieval_Scenario2_Step1_admi_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.Nil(t, mErr)
	message.data.MessageId = "20250301011104238MRSc2Step1MsgId"
	message.data.CreatedDateTime = time.Now()
	message.data.RequestType = model.RequestSent
	message.data.BusinessDate = model.FromTime(time.Now())
	message.data.SequenceRange = model.SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000003",
	}
	message.data.RecipientId = "B1QDRCQR"
	message.data.RecipientIssuer = "NA"

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("MessageRetrieval_Scenario2_Step1_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "MessageRetrieval_Scenario2_Step1_admi.006")
	genterated := filepath.Join("generated", "MessageRetrieval_Scenario2_Step1_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
