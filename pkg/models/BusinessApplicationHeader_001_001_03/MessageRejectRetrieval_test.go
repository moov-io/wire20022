package BusinessApplicationHeader_001_001_03

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMessageReject_Scenario1_Step2_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "231981435"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137c"
	mesage.model.MessageDefinitionId = "admi.002.001.01"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.acr.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "MessageReject_Scenario1_Step2_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestMessageRetrieval_Scenario1_Step1_head(t *testing.T) {
	var mesage = NewHead001Message()
	mesage.model.MessageSenderId = "231981435"
	mesage.model.MessageReceiverId = "021151080"
	mesage.model.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137c"
	mesage.model.MessageDefinitionId = "admi.002.001.01"
	mesage.model.BusinessService = "TEST"
	mesage.model.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.acr.01",
	}
	mesage.model.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "MessageRetrieval_Scenario1_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
