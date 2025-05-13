package FedwireFundsBroadcast

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
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: EventCode, EventParam, EventTime")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.EventCode == "" {
		msg.Data.EventCode = model.SystemOpen
	}
	if isEmpty(msg.Data.EventParam) {
		msg.Data.EventParam = model.FromTime(time.Now())
	}
	if isEmpty(msg.Data.EventTime) {
		msg.Data.EventTime = time.Now()
	}
	return msg
}
func TestFedwireFundsBroadcastFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "FedwireFundsBroadcast_admi.004_ADHC")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, "ADHC", string(message.Doc.SysEvtNtfctn.EvtInf.EvtCd))
	require.Equal(t, "The Fedwire Funds Service will open the test environment 15 minutes earlier on 03/13/2025", string(*message.Doc.SysEvtNtfctn.EvtInf.EvtDesc))
	require.NotEmpty(t, message.Doc.SysEvtNtfctn.EvtInf.EvtTm)
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

func TestFedwireFundsBroadcastValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"EventCode",
			Message{Data: MessageModel{EventCode: model.FundEventType(INVALID_COUNT)}},
			"error occur at EventCode: UNKNOWN fails enumeration validation",
		},
		{
			"EventDescriptionTooLong",
			Message{Data: MessageModel{EventDescription: string(make([]byte, 501))}}, // Assuming max length is 500
			"error occur at EventDescription: exceeds maximum length",
		},
		{
			"InvalidEventTime",
			Message{Data: MessageModel{EventTime: time.Time{}}}, // Empty time
			"error occur at EventTime: invalid or missing value",
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
func TestFedwireFundsBroadcast_admi_ADHC_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.EventCode = model.AdHoc
	message.Data.EventParam = model.FromTime(time.Now())
	message.Data.EventDescription = "The Fedwire Funds Service will open the test environment 15 minutes earlier on 03/13/2025"
	message.Data.EventTime = time.Now()
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsBroadcast_admi_ADHC.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsBroadcast_admi.004_ADHC")
	genterated := filepath.Join("generated", "FedwireFundsBroadcast_admi_ADHC.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsBroadcast_admi_CLSD_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.EventCode = model.SystemClosed
	message.Data.EventParam = model.FromTime(time.Now())
	message.Data.EventTime = time.Now()
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsBroadcast_admi_CLSD.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsBroadcast_admi.004_CLSD")
	genterated := filepath.Join("generated", "FedwireFundsBroadcast_admi_CLSD.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsBroadcast_admi_EXTN_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.EventCode = model.SystemExtension
	message.Data.EventParam = model.FromTime(time.Now())
	message.Data.EventDescription = `Fedwire Funds Service cutoff times: Customer Transfers is 00:00; Bank Transfers/Other is 00:00; Special Account is 00:00.

The Fedwire Funds Service has extended Customer Transfers 60 minutes to 19:45 p.m. Eastern Time for Bank ABCD. Bank Transfers/Other cutoff is 8:00 p.m. Eastern Time.`
	message.Data.EventTime = time.Now()
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsBroadcast_admi_EXTN.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsBroadcast_admi.004_EXTN")
	genterated := filepath.Join("generated", "FedwireFundsBroadcast_admi_EXTN.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFedwireFundsBroadcast_admi_OPEN_CreateXML(t *testing.T) {
	var message, mErr = NewMessage("")
	require.NoError(t, mErr)
	message.Data.EventCode = model.SystemOpen
	message.Data.EventParam = model.FromTime(time.Now())
	message.Data.EventTime = time.Now()
	cErr := message.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FedwireFundsBroadcast_admi_OPEN.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FedwireFundsBroadcast_admi.004_OPEN")
	genterated := filepath.Join("generated", "FedwireFundsBroadcast_admi_OPEN.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
