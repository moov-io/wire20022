package ConnectionCheck

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/fedwire20022/gen/ConnectionCheck/admi_004_001_02"
	"github.com/stretchr/testify/require"
	"github.com/wadearnold/wire20022/pkg/models"
)

var ConnectionChecksample1XML = filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step1_admi.004")

func TestDocumentElementToModelOne(t *testing.T) {
	var xmlData, err = models.ReadXMLFile(ConnectionChecksample1XML)
	require.NoError(t, err, "Failed to read XML file")

	model, err := MessageWith(xmlData)
	require.NoError(t, err, "Failed to make XML structure")
	require.Equal(t, model.EventType, "PING")
	require.Equal(t, model.EventParam, "BMQFMI01")
	require.NotNil(t, model.EventTime, "EventTime should not be nil")
}
func TestModelToDocument02(t *testing.T) {
	dataModel := ConnectionCheckDataModel()
	var doc02, err = DocumentWith(dataModel, ADMI_004_001_02)
	require.NoError(t, err, "Failed to create document")
	if Doc02, ok := doc02.(*admi_004_001_02.Document); ok {
		require.Equal(t, string(Doc02.SysEvtNtfctn.EvtInf.EvtCd), "PING")
		require.Equal(t, string(*Doc02.SysEvtNtfctn.EvtInf.EvtParam[0]), "BMQFMI01")
		require.NotNil(t, Doc02.SysEvtNtfctn.EvtInf.EvtTm, "EventTime should not be nil")
	}
}
func ConnectionCheckDataModel() MessageModel {
	var mesage = MessageModel{}
	mesage.EventType = "PING"
	mesage.EventParam = "BMQFMI01"
	mesage.EventTime = time.Now()
	return mesage
}
