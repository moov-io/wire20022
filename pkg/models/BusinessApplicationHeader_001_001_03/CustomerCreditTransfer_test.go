package BusinessApplicationHeader_001_001_03

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestACustomerCreditTransfer_Scenario1_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000001"
	mesage.model.MessageDefinitionId = "pacs.008.001.08"
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
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestACustomerCreditTransfer_Scenario1_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "011104238"
	mesage.model.BusinessMessageId = "20250310QMGFNP31000001"
	mesage.model.MessageDefinitionId = "pacs.002.001.10"
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
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario1_Step2_head_BankB(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "021040078"
	mesage.model.BusinessMessageId = "20250310QMGFNP31000001"
	mesage.model.MessageDefinitionId = "pacs.008.001.08"
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
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2_head_BankB.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario2_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000002"
	mesage.model.MessageDefinitionId = "pacs.008.001.08"
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
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario2_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "011104238"
	mesage.model.BusinessMessageId = "FDWA1B2C3D4E5F6G7H8I9J10K11L12M0"
	mesage.model.MessageDefinitionId = "pacs.002.001.10"
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
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step2_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario3_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000001"
	mesage.model.MessageDefinitionId = "pacs.008.001.08"
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
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario3_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310Scenario03Step2MsgId001"
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
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step2_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario3_Step3_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "011104238"
	mesage.model.BusinessMessageId = "A1B2C3D4E5F6G7H8I9J10K11L12M13N1400"
	mesage.model.MessageDefinitionId = "pacs.002.001.10"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()
	mesage.model.BusinessProcessingDate = time.Now()
	mesage.model.Relations = BusinessApplicationHeader{
		MessageSenderId:     "021151080",
		MessageReceiverId:   "011104238",
		BusinessMessageId:   "20250310QMGFNP31000001",
		MessageDefinitionId: "pacs.002.001.10",
		BusinessService:     "TEST",
		MarketInfo: MarketPractice{
			ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
			FrameworkId:       "frb.fedwire.01",
		},
		CreateDatetime:         time.Now(),
		BusinessProcessingDate: time.Now(),
	}

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step3_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario4_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "011104238"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000004"
	mesage.model.MessageDefinitionId = "pacs.008.001.08"
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
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario4_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "011104238"
	mesage.model.BusinessMessageId = "20250310QMGFNP31000002"
	mesage.model.MessageDefinitionId = "pacs.002.001.10"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()
	mesage.model.BusinessProcessingDate = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario4_Step2_head_UStreasury(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "011104238"
	mesage.model.BusinessMessageId = "20250310QMGFNP31000002"
	mesage.model.MessageDefinitionId = "pacs.002.001.10"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()
	mesage.model.BusinessProcessingDate = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step2_head_UStreasury.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario5_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021307481"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250310B1QDRCQR000005"
	mesage.model.MessageDefinitionId = "pacs.008.001.08"
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
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario5_Step2_head_BankC(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "021307481"
	mesage.model.BusinessMessageId = "20250310QMGFNP31000003"
	mesage.model.MessageDefinitionId = "pacs.008.001.08"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()
	mesage.model.BusinessProcessingDate = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2_head_BankC.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestCustomerCreditTransfer_Scenario5_Step2_head_BankD(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "231981435"
	mesage.model.BusinessMessageId = "20250310QMGFNP31000003"
	mesage.model.MessageDefinitionId = "pacs.008.001.08"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.model.CreateDatetime = time.Now()
	mesage.model.BusinessProcessingDate = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2_head_BankD.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
