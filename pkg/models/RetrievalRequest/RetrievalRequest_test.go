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
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreatedDateTime, RequestType, BusinessDate, RecipientId")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.MessageId == "" {
		msg.Data.MessageId = "20250310Scenario03Step2MsgId001"
	}
	if msg.Data.CreatedDateTime.IsZero() {
		msg.Data.CreatedDateTime = time.Now()
	}
	if msg.Data.RequestType == "" {
		msg.Data.RequestType = model.RequestSent
	}
	if isEmpty(msg.Data.BusinessDate) {
		msg.Data.BusinessDate = model.FromTime(time.Now())
	}
	if msg.Data.RecipientId == "" {
		msg.Data.RecipientId = "B1QDRCQR"
	}
	return msg
}
func TestRetrievalRequestFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "MessageRetrieval_Scenario1_Step1_admi.006")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "20250301011104238MRSc1Step1MsgId", string(message.Doc.RsndReq.MsgHdr.MsgId))
	require.Equal(t, "S", string(message.Doc.RsndReq.MsgHdr.ReqTp.Prtry.Id))
	require.Equal(t, "pacs.008.001.08", string(*message.Doc.RsndReq.RsndSchCrit.OrgnlMsgNmId))
	require.Equal(t, "20250310B1QDRCQR000001", string(*message.Doc.RsndReq.RsndSchCrit.FileRef))
	require.Equal(t, "B1QDRCQR", string(message.Doc.RsndReq.RsndSchCrit.Rcpt.Id.PrtryId.Id))
	require.Equal(t, "NA", string(message.Doc.RsndReq.RsndSchCrit.Rcpt.Id.PrtryId.Issr))
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
			Message{Data: MessageModel{MessageId: INVALID_OTHER_ID}},
			"error occur at MessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"Invalid RequestType",
			Message{Data: MessageModel{RequestType: model.RequestType(INVALID_COUNT)}},
			"error occur at CreatedDateTime: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid OriginalMessageNameId",
			Message{Data: MessageModel{OriginalMessageNameId: INVALID_MESSAGE_NAME_ID}},
			"error occur at OriginalMessageNameId: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"Invalid RecipientId",
			Message{Data: MessageModel{RecipientId: INVALID_MESSAGE_NAME_ID}},
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
	require.NoError(t, mErr)
	message.Data.MessageId = "20250301011104238MRSc1Step1MsgId"
	message.Data.CreatedDateTime = time.Now()
	message.Data.RequestType = model.RequestSent
	message.Data.BusinessDate = model.FromTime(time.Now())
	message.Data.OriginalMessageNameId = "pacs.008.001.08"
	message.Data.FileReference = "20250310B1QDRCQR000001"
	message.Data.RecipientId = "B1QDRCQR"
	message.Data.RecipientIssuer = "NA"

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("MessageRetrieval_Scenario1_Step1_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "MessageRetrieval_Scenario1_Step1_admi.006")
	genterated := filepath.Join("generated", "MessageRetrieval_Scenario1_Step1_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestMessageRetrieval_Scenario2_Step1_admi_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250301011104238MRSc2Step1MsgId"
	message.Data.CreatedDateTime = time.Now()
	message.Data.RequestType = model.RequestSent
	message.Data.BusinessDate = model.FromTime(time.Now())
	message.Data.SequenceRange = model.SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000003",
	}
	message.Data.RecipientId = "B1QDRCQR"
	message.Data.RecipientIssuer = "NA"

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("MessageRetrieval_Scenario2_Step1_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "MessageRetrieval_Scenario2_Step1_admi.006")
	genterated := filepath.Join("generated", "MessageRetrieval_Scenario2_Step1_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
