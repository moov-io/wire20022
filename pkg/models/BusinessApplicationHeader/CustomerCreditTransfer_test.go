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
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000001"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestACustomerCreditTransfer_Scenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "011104238"
	mesage.Data.BusinessMessageId = "20250310QMGFNP31000001"
	mesage.Data.MessageDefinitionId = "pacs.002.001.10"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	mesage.Data.BusinessProcessingDate = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
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
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "021040078"
	mesage.Data.BusinessMessageId = "20250310QMGFNP31000001"
	mesage.Data.MessageDefinitionId = "pacs.008.001.08"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	mesage.Data.BusinessProcessingDate = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
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
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000002"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario2_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "011104238"
	mesage.Data.BusinessMessageId = "FDWA1B2C3D4E5F6G7H8I9J10K11L12M0"
	mesage.Data.MessageDefinitionId = "pacs.002.001.10"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	mesage.Data.BusinessProcessingDate = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
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
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310Scenario03Step2MsgId001"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step2_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario3_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310Scenario03Step2MsgId001"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario3_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario3_Step2_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario3_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario3_Step3_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "011104238"
	mesage.Data.BusinessMessageId = "A1B2C3D4E5F6G7H8I9J10K11L12M13N1400"
	mesage.Data.MessageDefinitionId = "pacs.002.001.10"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	mesage.Data.BusinessProcessingDate = time.Now()
	mesage.Data.Relations = BusinessApplicationHeader{
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
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
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
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000004"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario4_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario4_Step1_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario4_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario4_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "011104238"
	mesage.Data.BusinessMessageId = "20250310QMGFNP31000002"
	mesage.Data.MessageDefinitionId = "pacs.002.001.10"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	mesage.Data.BusinessProcessingDate = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
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
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "091036164"
	mesage.Data.BusinessMessageId = "20250310QMGFNP31000002"
	mesage.Data.MessageDefinitionId = "pacs.008.001.08"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	mesage.Data.BusinessProcessingDate = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
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
	mesage.Data.MessageSenderId = "021307481"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000005"
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
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario5_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario5_Step1_head.001")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestCustomerCreditTransfer_Scenario5_Step2_head_BankC(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "021307481"
	mesage.Data.BusinessMessageId = "20250310QMGFNP31000003"
	mesage.Data.MessageDefinitionId = "pacs.002.001.10"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	mesage.Data.BusinessProcessingDate = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
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
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "231981435"
	mesage.Data.BusinessMessageId = "20250310QMGFNP31000003"
	mesage.Data.MessageDefinitionId = "pacs.008.001.08"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	mesage.Data.BusinessProcessingDate = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("CustomerCreditTransfer_Scenario5_Step2_head_BankD.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "CustomerCreditTransfer_Scenario5_Step2_head.001_BankD")
	genterated := filepath.Join("generated", "CustomerCreditTransfer_Scenario5_Step2_head_BankD.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
