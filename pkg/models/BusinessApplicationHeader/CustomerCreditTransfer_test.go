package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestACustomerCreditTransfer_Scenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000001"
	mesage.data.MessageDefinitionId = "pacs.008.001.08"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestACustomerCreditTransfer_Scenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "011104238"
	mesage.data.BusinessMessageId = "20250310QMGFNP31000001"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step2_head.001_BankA")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario1_Step2_head_BankB(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "021040078"
	mesage.data.BusinessMessageId = "20250310QMGFNP31000001"
	mesage.data.MessageDefinitionId = "pacs.008.001.08"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario1_Step2_head_BankB.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step2_head.001_BankB")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step2_head_BankB.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario2_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000002"
	mesage.data.MessageDefinitionId = "pacs.008.001.08"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario2_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "011104238"
	mesage.data.BusinessMessageId = "FDWA1B2C3D4E5F6G7H8I9J10K11L12M0"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario2_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario2_Step2_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario3_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310Scenario03Step2MsgId001"
	mesage.data.MessageDefinitionId = "pacs.028.001.03"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step2_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario3_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310Scenario03Step2MsgId001"
	mesage.data.MessageDefinitionId = "pacs.028.001.03"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step2_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario3_Step3_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "011104238"
	mesage.data.BusinessMessageId = "A1B2C3D4E5F6G7H8I9J10K11L12M13N1400"
	mesage.data.MessageDefinitionId = "pacs.002.001.10"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()
	mesage.data.BusinessProcessingDate = time.Now()
	mesage.data.Relations = BusinessApplicationHeader{
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

	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step3_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step3_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario4_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000004"
	mesage.data.MessageDefinitionId = "pacs.008.001.08"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario4_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario4_Step1_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario4_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "011104238"
	mesage.data.BusinessMessageId = "20250310QMGFNP31000002"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario4_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario4_Step2_head.001_BankA")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario4_Step2_head_UStreasury(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "091036164"
	mesage.data.BusinessMessageId = "20250310QMGFNP31000002"
	mesage.data.MessageDefinitionId = "pacs.008.001.08"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario4_Step2_head_UStreasury.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario4_Step2_head.001_USTreasury")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step2_head_UStreasury.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario5_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021307481"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000005"
	mesage.data.MessageDefinitionId = "pacs.008.001.08"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario5_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario5_Step1_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario5_Step2_head_BankC(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "021307481"
	mesage.data.BusinessMessageId = "20250310QMGFNP31000003"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario5_Step2_head_BankC.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario5_Step2_head.001_BankC")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2_head_BankC.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario5_Step2_head_BankD(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "20250310QMGFNP31000003"
	mesage.data.MessageDefinitionId = "pacs.008.001.08"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario5_Step2_head_BankD.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario5_Step2_head.001_BankD")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2_head_BankD.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
