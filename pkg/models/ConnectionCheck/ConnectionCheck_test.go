package ConnectionCheck

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestConnectionCheckFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step1_admi.004")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.doc.SysEvtNtfctn.EvtInf.EvtCd), "PING")
	require.Equal(t, string(message.doc.SysEvtNtfctn.EvtInf.EvtParam), "BMQFMI01")
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
			Message{data: MessageModel{EventType: INVALID_COUNT}},
			"error occur at EventType: UNKNOWN fails enumeration validation",
		},
		{
			"EvntParam",
			Message{data: MessageModel{EvntParam: INVALID_COUNT}},
			"error occur at EvntParam: UNKNOWN fails validation with pattern [A-Z0-9]{8,8}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			msgErr := tt.msg.CreateDocument()
			require.Equal(t, tt.expectedErr, msgErr.Error())
		})
	}
}
func TestConnectionCheck_Scenario1_Step1_admi(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.EventType = "PING"
	mesage.data.EvntParam = "BMQFMI01"
	mesage.data.EventTime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("ConnectionCheck_Scenario1_Step1_admi.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step1_admi.004")
	genterated := filepath.Join("generated", "ConnectionCheck_Scenario1_Step1_admi.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
