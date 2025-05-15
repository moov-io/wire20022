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
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000601"
	mesage.Data.MessageDefinitionId = "pain.013.001.07"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021040078"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000602"
	mesage.Data.MessageDefinitionId = "pain.014.001.07"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step2_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario1_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario1_Step3_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021040078"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000603"
	mesage.Data.MessageDefinitionId = "pacs.008.001.08"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario1_Step3_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario1_Step3_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario1_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario2_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000611"
	mesage.Data.MessageDefinitionId = "pain.013.001.07"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario2_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021040078"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000612"
	mesage.Data.MessageDefinitionId = "pain.014.001.07"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario2_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario2_Step2_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario2_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario3_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000621"
	mesage.Data.MessageDefinitionId = "pain.013.001.07"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario3_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step1_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario3_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021040078"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000622"
	mesage.Data.MessageDefinitionId = "pain.014.001.07"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario3_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step2_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario3_Step3_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021040078"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000623"
	mesage.Data.MessageDefinitionId = "pacs.009.001.08"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario3_Step3_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario3_Step3_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario3_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000681"
	mesage.Data.MessageDefinitionId = "pain.013.001.07"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario4_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step1_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021040078"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000682"
	mesage.Data.MessageDefinitionId = "pain.014.001.07"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario4_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step2_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario4_Step3_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021040078"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000683"
	mesage.Data.MessageDefinitionId = "pacs.009.001.08"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario4_Step3_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario4_Step3_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario4_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestDrawdowns_Scenario5_Step3_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310Scenario04Step3MsgId001"
	mesage.Data.MessageDefinitionId = "pacs.028.001.03"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Drawdowns_Scenario5_Step3_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Drawdowns_Scenario5_Step3_head.001")
	genterated := filepath.Join("generated", "Drawdowns_Scenario5_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
