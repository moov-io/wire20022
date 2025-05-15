package FedwireFundsAcknowledgement

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, CreatedDateTime, RelationReference, ReferenceName, RequestHandling")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.MessageId == "" {
		msg.Data.MessageId = "20250310QMGFNP7500070103101100FT03"
	}
	if msg.Data.CreatedDateTime.IsZero() {
		msg.Data.CreatedDateTime = time.Now()
	}
	if msg.Data.RelationReference == "" {
		msg.Data.RelationReference = "20250310B1QDRCQR000711"
	}
	if msg.Data.ReferenceName == "" {
		msg.Data.ReferenceName = "pain.013.001.07"
	}
	if msg.Data.RequestHandling == "" {
		msg.Data.RequestHandling = model.SchemaValidationFailed
	}
	return msg
}
func TestFedwireFundsAcknowledgementFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step1a_admi.007")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.Doc.RctAck.MsgId.MsgId), "20250310QMGFNP7500070103101100FT03")
	require.Equal(t, string(message.Doc.RctAck.Rpt.RltdRef.Ref), "20250310B1QDRCQR000711")
	require.Equal(t, string(message.Doc.RctAck.Rpt.RltdRef.MsgNm), "pain.013.001.07")
	require.Equal(t, string(message.Doc.RctAck.Rpt.ReqHdlg.StsCd), "TS01")
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

func TestFedwireFundsAcknowledgementValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"MessageId",
			Message{Data: MessageModel{MessageId: INVALID_OTHER_ID}},
			"error occur at MessageId: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [A-Z0-9]{34,34}",
		},
		{
			"RelationReference",
			Message{Data: MessageModel{RelationReference: INVALID_OTHER_ID}},
			"error occur at RelationReference: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with length 50 <= required maxLength 35",
		},
		{
			"ReferenceName",
			Message{Data: MessageModel{ReferenceName: INVALID_MESSAGE_NAME_ID}},
			"error occur at ReferenceName: sabcd-123-001-12 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"RequestHandling",
			Message{Data: MessageModel{RequestHandling: model.RelatedStatusCode(INVALID_COUNT)}},
			"error occur at RequestHandling: UNKNOWN fails enumeration validation",
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
func TestFedwireFundsAcknowledgement_Scenario1_Step1a_admi_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310QMGFNP7500070103101100FT03"
	message.Data.CreatedDateTime = time.Now()
	message.Data.RelationReference = "20250310B1QDRCQR000711"
	message.Data.ReferenceName = "pain.013.001.07"
	message.Data.RequestHandling = model.SchemaValidationFailed

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step1a_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step1a_admi.007")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step1a_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2a_admi_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310QMGFNP7500070203101130FT03"
	message.Data.CreatedDateTime = time.Now()
	message.Data.RelationReference = "20250310B1QDRCQR000712"
	message.Data.ReferenceName = "pain.014.001.07"
	message.Data.RequestHandling = model.SchemaValidationFailed

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario1_Step2a_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario1_Step2a_admi.007")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step2a_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario2_Step2a_admi_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310QMGFNP6500072203101100FT03"
	message.Data.CreatedDateTime = time.Now()
	message.Data.RelationReference = "20250310B1QDRCQR000722"
	message.Data.ReferenceName = "camt.056.001.08"
	message.Data.RequestHandling = model.SchemaValidationFailed

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step2a_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step2a_admi.007")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step2a_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsAcknowledgement_Scenario2_Step3a_admi_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.MessageId = "20250310QMGFNP6500072303101100FT03"
	message.Data.CreatedDateTime = time.Now()
	message.Data.RelationReference = "20250310B1QDRCQR000723"
	message.Data.ReferenceName = "camt.029.001.09"
	message.Data.RequestHandling = model.SchemaValidationFailed

	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsAcknowledgement_Scenario2_Step3a_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsAcknowledgement_Scenario2_Step3a_admi.007")
	genterated := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step3a_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
