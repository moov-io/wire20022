package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario1_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000601"
	mesage.data.MessageDefinitionId = "pain.013.001.07"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario1_Step2_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000602"
	mesage.data.MessageDefinitionId = "pain.014.001.07"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step2_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario1_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario1_Step3_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000603"
	mesage.data.MessageDefinitionId = "pacs.008.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario1_Step3_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step3_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario1_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario2_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000611"
	mesage.data.MessageDefinitionId = "pain.013.001.07"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario2_Step2_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000612"
	mesage.data.MessageDefinitionId = "pain.014.001.07"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario2_Step2_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario2_Step2_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario2_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario3_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000621"
	mesage.data.MessageDefinitionId = "pain.013.001.07"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario3_Step1_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step1_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario3_Step2_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000622"
	mesage.data.MessageDefinitionId = "pain.014.001.07"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario3_Step2_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step2_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario3_Step3_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000623"
	mesage.data.MessageDefinitionId = "pacs.009.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario3_Step3_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step3_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000681"
	mesage.data.MessageDefinitionId = "pain.013.001.07"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario4_Step1_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step1_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step2_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000682"
	mesage.data.MessageDefinitionId = "pain.014.001.07"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario4_Step2_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step2_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step3_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000683"
	mesage.data.MessageDefinitionId = "pacs.009.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario4_Step3_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step3_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario5_Step3_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310Scenario04Step3MsgId001"
	mesage.data.MessageDefinitionId = "pacs.028.001.03"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	model.WriteXMLTo("Drawdowns_Scenario5_Step3_head.xml", xmlData)
	require.NoError(t, err)
	
	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario5_Step3_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario5_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
