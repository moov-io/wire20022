package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestScenario1_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250311143738 ABAR MM Request"
	mesage.data.MessageDefinitionId = "camt.060.001.05"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()
	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("AccountBalanceReport_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestScenario1_Step2_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137a"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.abm.01",
	}
	mesage.data.CreateDatetime = time.Now()
	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("AccountBalanceReport_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step2_head.001")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestScenario2_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137N"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.abs.01",
	}
	mesage.data.CreateDatetime = time.Now()
	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("AccountBalanceReport_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
