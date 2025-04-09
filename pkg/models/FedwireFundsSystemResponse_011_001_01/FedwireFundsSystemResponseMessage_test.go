package FedwireFundsSystemResponse_011_001_01

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestConnectionCheck_Scenario1_Step2_admi_CreateXML(t *testing.T) {
	var message = NewAdmi011Message()
	message.model.MessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	message.model.EventCode = model.ConnectionCheck
	message.model.EventParam = "BMQFMI01"
	message.model.EventTime = time.Now()
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "ConnectionCheck_Scenario1_Step2_admi.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
