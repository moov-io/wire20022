package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestConnectionCheck_Scenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021052587"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "ConnectionCheck"
	mesage.Data.MessageDefinitionId = "admi.004.001.02"
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
	err = model.WriteXMLTo("ConnectionCheck_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "ConnectionCheck_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}

func TestConnectionCheck_Scenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "021052587"
	mesage.Data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137b"
	mesage.Data.MessageDefinitionId = "admi.011.001.01"
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
	err = model.WriteXMLTo("ConnectionCheck_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "ConnectionCheck_Scenario1_Step2_head.001")
	genterated := filepath.Join("generated", "ConnectionCheck_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
