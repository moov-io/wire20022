package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestPaymentReturn_Scenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "011104238"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250310B1QDRCQR000400"
	mesage.Data.MessageDefinitionId = "pacs.008.001.08"
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
	err = model.WriteXMLTo("PaymentReturn_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "PaymentReturn_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "PaymentReturn_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
