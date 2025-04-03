package BusinessApplicationHeader_001_001_03

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestScenario1_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "231981435"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "20250311143738 ABAR MM Request"
	mesage.model.MessageDefinitionId = "camt.060.001.05"
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
	xnlFileName := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestScenario1_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "231981435"
	mesage.model.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137a"
	mesage.model.MessageDefinitionId = "camt.052.001.08"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.abm.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step2_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestScenario2_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "021151080"
	mesage.model.MessageReceiverId = "231981435"
	mesage.model.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137a"
	mesage.model.MessageDefinitionId = "camt.052.001.08"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.abm.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountBalanceReport_Scenario2_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
