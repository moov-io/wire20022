package BusinessApplicationHeader_001_001_03

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDrawdowns_Scenario1_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000601"
	mesage.model.MessageDefinitionId = "pain.013.001.07"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario1_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario1_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021040078"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000602"
	mesage.model.MessageDefinitionId = "pain.013.001.07"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario1_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario1_Step3_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021040078"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000603"
	mesage.model.MessageDefinitionId = "pain.013.001.07"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario1_Step3_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario2_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000611"
	mesage.model.MessageDefinitionId = "pain.013.001.07"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario2_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario2_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021040078"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000612"
	mesage.model.MessageDefinitionId = "pain.014.001.07"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario2_Step2_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario3_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000621"
	mesage.model.MessageDefinitionId = "pain.013.001.07"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario3_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario3_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021040078"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000622"
	mesage.model.MessageDefinitionId = "pain.014.001.07"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario3_Step2_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario3_Step3_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021040078"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000623"
	mesage.model.MessageDefinitionId = "pacs.009.001.08"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario3_Step3_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario4_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000681"
	mesage.model.MessageDefinitionId = "pain.013.001.07"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario4_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario4_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021040078"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000682"
	mesage.model.MessageDefinitionId = "pain.014.001.07"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario4_Step2_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario4_Step3_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021040078"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000683"
	mesage.model.MessageDefinitionId = "pacs.009.001.08"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario4_Step3_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestDrawdowns_Scenario5_Step3_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310Scenario04Step3MsgId001"
	mesage.model.MessageDefinitionId = "pacs.028.001.03"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "Drawdowns_Scenario5_Step3_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
