package ConnectionCheck

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: EventType, EventParam, EventTime")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.EventType == "" {
		msg.Data.EventType = "PING"
	}
	if msg.Data.EventParam == "" {
		msg.Data.EventParam = "BMQFMI01"
	}
	if isEmpty(msg.Data.EventTime) {
		msg.Data.EventTime = time.Now()
	}
	return msg
}
func TestConnectionCheckFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step1_admi.004")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.Doc.SysEvtNtfctn.EvtInf.EvtCd), "PING")
	require.Equal(t, string(message.Doc.SysEvtNtfctn.EvtInf.EvtParam), "BMQFMI01")
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"

func TestAccountBalanceReportValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"EventType",
			Message{Data: MessageModel{EventType: INVALID_COUNT}},
			"error occur at EventType: UNKNOWN fails enumeration validation",
		},
		{
			"EvntParam",
			Message{Data: MessageModel{EventParam: INVALID_COUNT}},
			"error occur at EvntParam: UNKNOWN fails validation with pattern [A-Z0-9]{8,8}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			nMsg := generateRequreFields(tt.msg)
			msgErr := nMsg.CreateDocument()
			require.Equal(t, tt.expectedErr, msgErr.Error())
		})
	}
}
func TestConnectionCheck_Scenario1_Step1_admi(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.EventType = "PING"
	mesage.Data.EventParam = "BMQFMI01"
	mesage.Data.EventTime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("ConnectionCheck_Scenario1_Step1_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step1_admi.004")
	genterated := filepath.Join("generated", "ConnectionCheck_Scenario1_Step1_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
