package FedwireFundsBroadcast_004_001_02

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestFedwireFundsBroadcast_admi_ADHC_CreateXML(t *testing.T) {
	var message = NewAdmi004Message()

	message.model.EventCode = model.AdHoc
	message.model.EventParam = civil.DateOf(time.Now())
	message.model.EventDescription = "The Fedwire Funds Service will open the test environment 15 minutes earlier on 03/13/2025"
	message.model.EventTime = time.Now()
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsBroadcast_admi_ADHC.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsBroadcast_admi_CLSD_CreateXML(t *testing.T) {
	var message = NewAdmi004Message()

	message.model.EventCode = model.SystemClosed
	message.model.EventParam = civil.DateOf(time.Now())
	message.model.EventTime = time.Now()
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsBroadcast_admi_CLSD.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsBroadcast_admi_EXTN_CreateXML(t *testing.T) {
	var message = NewAdmi004Message()

	message.model.EventCode = model.SystemExtension
	message.model.EventParam = civil.DateOf(time.Now())
	message.model.EventDescription = "Fedwire Funds Service cutoff times: Customer Transfers is 00:00; Bank Transfers/Other is 00:00; Special Account is 00:00. \n The Fedwire Funds Service has extended Customer Transfers 60 minutes to 19:45 p.m. Eastern Time for Bank ABCD. Bank Transfers/Other cutoff is 8:00 p.m. Eastern Time."
	message.model.EventTime = time.Now()
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsBroadcast_admi_EXTN.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestFedwireFundsBroadcast_admi_OPEN_CreateXML(t *testing.T) {
	var message = NewAdmi004Message()

	message.model.EventCode = model.SystemOpen
	message.model.EventParam = civil.DateOf(time.Now())
	message.model.EventTime = time.Now()
	message.CreateDocument()

	xmlData, err := message.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "PaymentReturn_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "FedwireFundsBroadcast_admi_OPEN.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
