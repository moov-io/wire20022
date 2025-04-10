package BusinessApplicationHeader

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConnectionCheck_Scenario1_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021052587"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "ConnectionCheck"
	mesage.data.MessageDefinitionId = "admi.004.001.02"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	WriteXMLTo("ConnectionCheck_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)
}

func TestConnectionCheck_Scenario1_Step2_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "021052587"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	mesage.data.MessageDefinitionId = "admi.011.001.01"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	WriteXMLTo("ConnectionCheck_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)
}
