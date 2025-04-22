package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestBusinessApplicationHeaderFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_head.001")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.doc.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId), "231981435")
	require.Equal(t, string(message.doc.To.FIId.FinInstnId.ClrSysMmbId.MmbId), "021151080")
	require.Equal(t, string(message.doc.BizMsgIdr), "20250311143738 ABAR MM Request")
	require.Equal(t, string(message.doc.BizSvc), "TEST")
	require.Equal(t, string(message.doc.MktPrctc.Regy), "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service")
	require.Equal(t, string(message.doc.MktPrctc.Id), "frb.fedwire.01")
}

const INVALID_ACCOUNT_ID string = "123ABC789"
const INVALID_COUNT string = "UNKNOWN"

func TestBusinessApplicationHeaderValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"MessageSenderId",
			Message{data: MessageModel{MessageSenderId: "Unknown data"}},
			"error occur at MessageSenderId: Unknown data fails validation with pattern [A-Z0-9]{9,9}",
		},
		{
			"MessageReceiverId",
			Message{data: MessageModel{MessageReceiverId: "Unknown data"}},
			"error occur at MessageReceiverId: Unknown data fails validation with pattern [A-Z0-9]{9,9}",
		},
		{
			"MessageDefinitionId",
			Message{data: MessageModel{MessageDefinitionId: "-------"}},
			"error occur at MessageDefinitionId: ------- fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"MarketInfo - FrameworkId",
			Message{data: MessageModel{MarketInfo: MarketPractice{
				ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
				FrameworkId:       "wwww",
			}}},
			"error occur at MarketInfo.FrameworkId: wwww fails validation with pattern frb([.]{1,1})fedwire([.]{1,1})(([a-z]{3,3}[.]{1,1})){0,1}01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			msgErr := tt.msg.CreateDocument()
			if msgErr != nil {
				require.Equal(t, tt.expectedErr, msgErr.Error())
			}
		})
	}
}
func TestScenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "231981435"
	mesage.data.MessageReceiverId = "021151080"
	mesage.data.BusinessMessageId = "20250311143738 ABAR MM Request"
	mesage.data.MessageDefinitionId = "camt.060.001.05"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.01",
	}
	mesage.data.CreateDatetime = time.Now()
	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountBalanceReport_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestScenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137a"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.abm.01",
	}
	mesage.data.CreateDatetime = time.Now()
	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountBalanceReport_Scenario1_Step2_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step2_head.001")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step2_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestScenario2_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.data.MessageSenderId = "021151080"
	mesage.data.MessageReceiverId = "231981435"
	mesage.data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137N"
	mesage.data.MessageDefinitionId = "camt.052.001.08"
	mesage.data.BusinessService = "TEST"
	mesage.data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.abs.01",
	}
	mesage.data.CreateDatetime = time.Now()
	cErr := mesage.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&mesage.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountBalanceReport_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
