package BusinessApplicationHeader

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMessageReject_Scenario1_Step2_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137c"
	mesage.data.MessageDefinitionId = "admi.002.001.01"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.acr.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	WriteXMLTo("MessageReject_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)
}
func TestMessageRetrieval_Scenario1_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137c"
	mesage.data.MessageDefinitionId = "admi.002.001.01"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.acr.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	WriteXMLTo("MessageRetrieval_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)
}
