package FedwireFundsBroadcast

import (
	"encoding/xml"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsBroadcast_admi_ADHC_CreateXML(t *testing.T) {
	var message = NewMessage()

	message.data.EventCode = model.AdHoc
	message.data.EventParam = civil.DateOf(time.Now())
	message.data.EventDescription = "The Fedwire Funds Service will open the test environment 15 minutes earlier on 03/13/2025"
	message.data.EventTime = time.Now()
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FedwireFundsBroadcast_admi_ADHC.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsBroadcast_admi_CLSD_CreateXML(t *testing.T) {
	var message = NewMessage()

	message.data.EventCode = model.SystemClosed
	message.data.EventParam = civil.DateOf(time.Now())
	message.data.EventTime = time.Now()
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FedwireFundsBroadcast_admi_CLSD.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsBroadcast_admi_EXTN_CreateXML(t *testing.T) {
	var message = NewMessage()

	message.data.EventCode = model.SystemExtension
	message.data.EventParam = civil.DateOf(time.Now())
	message.data.EventDescription = "Fedwire Funds Service cutoff times: Customer Transfers is 00:00; Bank Transfers/Other is 00:00; Special Account is 00:00. \n The Fedwire Funds Service has extended Customer Transfers 60 minutes to 19:45 p.m. Eastern Time for Bank ABCD. Bank Transfers/Other cutoff is 8:00 p.m. Eastern Time."
	message.data.EventTime = time.Now()
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FedwireFundsBroadcast_admi_EXTN.xml", xmlData)
	require.NoError(t, err)
}
func TestFedwireFundsBroadcast_admi_OPEN_CreateXML(t *testing.T) {
	var message = NewMessage()

	message.data.EventCode = model.SystemOpen
	message.data.EventParam = civil.DateOf(time.Now())
	message.data.EventTime = time.Now()
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("FedwireFundsBroadcast_admi_OPEN.xml", xmlData)
	require.NoError(t, err)
}
