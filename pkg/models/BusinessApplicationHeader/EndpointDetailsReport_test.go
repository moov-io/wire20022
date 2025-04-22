package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestEndpointDetailsReport_Scenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250311143738 DTLS Request"
	mesage.data.MessageDefinitionId = "camt.060.001.05"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointDetailsReport_Scenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137a"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.edr.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario1_Step2_head.001")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointDetailsReport_Scenario2_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250311143738 DTLR Request"
	mesage.data.MessageDefinitionId = "camt.060.001.05"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointDetailsReport_Scenario2_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.edr.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario2_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario2_Step2_head.001")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario2_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointGapReport_Scenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.abs.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointGapReport_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointGapReport_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "EndpointGapReport_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointTotalsReport_Scenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250311143738 ETOT Request"
	mesage.data.MessageDefinitionId = "camt.060.001.05"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointTotalsReport_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointTotalsReport_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "EndpointTotalsReport_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}

func TestEndpointTotalsReport_Scenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137a"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.etr.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointTotalsReport_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointTotalsReport_Scenario1_Step2_head.001")
	genterated := filepath.Join("generated", "EndpointTotalsReport_Scenario1_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointTotalsReport_Scenario2_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137a"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.etr.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointTotalsReport_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointTotalsReport_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "EndpointTotalsReport_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
