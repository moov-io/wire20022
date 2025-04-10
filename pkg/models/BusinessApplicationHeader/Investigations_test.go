package BusinessApplicationHeader

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestInvestigations_Scenario1_Step2_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310QMGFT015000901"
	mesage.data.MessageDefinitionId = "camt.110.001.01"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.acr.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	WriteXMLTo("Investigations_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)
}
func TestInvestigations_Scenario1_Step3_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000901"
	mesage.data.MessageDefinitionId = "camt.111.001.01"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.acr.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	WriteXMLTo("Investigations_Scenario1_Step3_head.xml", xmlData)
	require.NoError(t, err)
}
