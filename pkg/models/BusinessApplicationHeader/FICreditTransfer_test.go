package BusinessApplicationHeader

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFICreditTransfer_Scenario1_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000501"
	mesage.data.MessageDefinitionId = "pacs.009.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := xml.MarshalIndent(mesage, "", "\t")
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario1_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario2_Step1_head_bankc(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310QMGFNP62000501"
	mesage.data.MessageDefinitionId = "pacs.002.001.10"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()
	mesage.data.BusinessProcessingDate = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := xml.MarshalIndent(mesage, "", "\t")
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_head_bankc.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario2_Step1_head_bankd(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310QMGFNP62000501"
	mesage.data.MessageDefinitionId = "pacs.009.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()
	mesage.data.BusinessProcessingDate = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := xml.MarshalIndent(mesage, "", "\t")
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_head_bankd.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFICreditTransfer_Scenario2_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000502"
	mesage.data.MessageDefinitionId = "pacs.009.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := xml.MarshalIndent(mesage, "", "\t")
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FICreditTransfer_Scenario2_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
