package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestFICreditTransfer_Scenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021307481"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000501"
	mesage.data.MessageDefinitionId = "pacs.009.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step1_head_bankc(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "021307481"
	mesage.data.BusinessMessageId = "20250310QMGFNP62000501"
	mesage.data.MessageDefinitionId = "pacs.002.001.10"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()
	mesage.data.BusinessProcessingDate = time.Now()

	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step1_head_bankc.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step2_head.001_BankC")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_head_bankc.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step1_head_bankd(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "20250310QMGFNP62000501"
	mesage.data.MessageDefinitionId = "pacs.009.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()
	mesage.data.BusinessProcessingDate = time.Now()

	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step1_head_bankd.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario1_Step2_head.001_BankD")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_head_bankd.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestFICreditTransfer_Scenario2_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021307481"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000502"
	mesage.data.MessageDefinitionId = "pacs.009.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("FICreditTransfer_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "FICreditTransfer_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
