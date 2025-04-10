package BusinessApplicationHeader

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestActivityReport_Scenario1_Step1_head(t *testing.T) {
	var mesage = NewMessage()
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.acr.01",
	}
	mesage.data.CreateDatetime = time.Now()

	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := xml.MarshalIndent(mesage, "", "\t")
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "ActivityReport_Scenario1_Step1_head.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
