package ConnectionCheck

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConnectionCheck_Scenario1_Step1_admi(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.EventType = "PING"
	mesage.data.EvntParam = "BMQFMI01"
	mesage.data.EventTime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	WriteXMLTo("ConnectionCheck_Scenario1_Step1_admi.xml", xmlData)
	require.NoError(t, err)
}
