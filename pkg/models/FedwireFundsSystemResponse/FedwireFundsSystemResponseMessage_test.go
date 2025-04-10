package FedwireFundsSystemResponse

import (
	"encoding/xml"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestConnectionCheck_Scenario1_Step2_admi_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	message.data.EventCode = model.ConnectionCheck
	message.data.EventParam = "BMQFMI01"
	message.data.EventTime = time.Now()
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("ConnectionCheck_Scenario1_Step2_admi.xml", xmlData)
	require.NoError(t, err)
}
