package FedwireFundsSystemResponse

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageId, EventCode, EventParam, EventTime")
}
func generateRequreFields(msg Message) Message {
	if msg.data.MessageId == "" {
		msg.data.MessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	}
	if msg.data.EventCode == "" {
		msg.data.EventCode = model.ConnectionCheck
	}
	if msg.data.EventParam == "" {
		msg.data.EventParam = "BMQFMI01"
	}
	if msg.data.EventTime.IsZero() {
		msg.data.EventTime = time.Now()
	}
	return msg
}
func TestFedwireFundsSystemResponseFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step2_admi.011")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	// Validate the parsed message fields
	require.Equal(t, "98z2cb3d0f2f3094f24a16389713541137b", string(message.doc.SysEvtAck.MsgId))
	require.Equal(t, "PING", string(message.doc.SysEvtAck.AckDtls.EvtCd))
	require.Equal(t, "BMQFMI01", string(message.doc.SysEvtAck.AckDtls.EvtParam))
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

func TestFedwireFundsSystemResponseValidator(t *testing.T) {
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
			"Invalid EventCode",
			Message{data: MessageModel{EventCode: model.FundEventType(INVALID_COUNT)}},
			"error occur at EventCode: UNKNOWN fails enumeration validation",
		},
		{
			"Invalid EventParam",
			Message{data: MessageModel{EventParam: INVALID_OTHER_ID}},
			"error occur at EventParam: aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa fails validation with pattern [A-Z0-9]{8,8}",
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
func TestConnectionCheck_Scenario1_Step2_admi_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.data.MessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	message.data.EventCode = model.ConnectionCheck
	message.data.EventParam = "BMQFMI01"
	message.data.EventTime = time.Now()
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("ConnectionCheck_Scenario1_Step2_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step2_admi.011")
	genterated := filepath.Join("generated", "ConnectionCheck_Scenario1_Step2_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
