package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestInvestigations_Scenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021040078"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310QMGFT015000901"
	mesage.data.MessageDefinitionId = "camt.110.001.01"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario1_Step2_head.001")
	genterated := filepath.Join("generated", "Investigations_Scenario1_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestInvestigations_Scenario1_Step3_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "011104238"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250310B1QDRCQR000901"
	mesage.data.MessageDefinitionId = "camt.111.001.01"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()

	cErr := mesage.CreateDocument()
	require.NoError(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("Investigations_Scenario1_Step3_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "Investigations_Scenario1_Step3_head.001")
	genterated := filepath.Join("generated", "Investigations_Scenario1_Step3_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
