package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestMessageReject_Scenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137c"
	mesage.data.MessageDefinitionId = "admi.002.001.01"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("MessageReject_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "MessageReject_Scenario1_Step2_head.001")
	genterated := filepath.Join("generated", "MessageReject_Scenario1_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestMessageRetrieval_Scenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021051080"
	mesage.data.BusinessMessageId = "20250301011104238MRSc1Step1MsgId"
	mesage.data.MessageDefinitionId = "admi.006.001.01"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("MessageRetrieval_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "MessageRetrieval_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "MessageRetrieval_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
