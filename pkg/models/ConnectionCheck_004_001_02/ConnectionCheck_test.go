package ConnectionCheck_004_001_02

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConnectionCheck_Scenario1_Step1_admi(t *testing.T) {
	var mesage = NewAdmi004Message()
	mesage.model.EventType = "PING"
	mesage.model.EvntParam = "BMQFMI01"
	mesage.model.EventTime = time.Now()
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "ConnectionCheck_Scenario1_Step1_admi.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
