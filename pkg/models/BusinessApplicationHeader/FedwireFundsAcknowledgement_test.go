package BusinessApplicationHeader

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFedwireFundsAcknowledgement_Scenario1_Step1a_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310QMGFNP6200070203101000FT03"
	mesage.data.MessageDefinitionId = "admi.007.001.01"
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
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step1a_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step2a_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310QMGFNP6200070303101000FT03"
	mesage.data.MessageDefinitionId = "admi.007.001.01"
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
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step2a_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario1_Step3a_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310QMGFNP6200070403101000FT03"
	mesage.data.MessageDefinitionId = "admi.007.001.01"
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
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario1_Step3a_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario2_Step2a_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310QMGFNP6200070503101000FT03"
	mesage.data.MessageDefinitionId = "admi.007.001.01"
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
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step2a_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsAcknowledgement_Scenario2_Step3a_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310QMGFNP6200070603101000FT03"
	mesage.data.MessageDefinitionId = "admi.007.001.01"
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
	xnlFileName := filepath.Join("generated", "FedwireFundsAcknowledgement_Scenario2_Step3a_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
