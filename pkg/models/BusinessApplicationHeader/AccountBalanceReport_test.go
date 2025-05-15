package BusinessApplicationHeader

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestRequireField(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	cErr := message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("require.xml", xmlData)
	require.NoError(t, err)
	require.Equal(t, cErr.Error(), "error occur at RequiredFields: MessageSenderId, MessageReceiverId, BusinessMessageId, MessageDefinitionId, BusinessService, MarketInfo")
}
func generateRequreFields(msg Message) Message {
	if msg.Data.MessageSenderId == "" {
		msg.Data.MessageSenderId = "231981435"
	}
	if msg.Data.MessageReceiverId == "" {
		msg.Data.MessageReceiverId = "021151080"
	}
	if msg.Data.BusinessMessageId == "" {
		msg.Data.BusinessMessageId = "20250311143738 ABAR MM Request"
	}
	if msg.Data.MessageDefinitionId == "" {
		msg.Data.MessageDefinitionId = "camt.060.001.05"
	}
	if msg.Data.BusinessService == "" {
		msg.Data.BusinessService = "TEST"
	}
	if isEmpty(msg.Data.MarketInfo) {
		msg.Data.MarketInfo = MarketPractice{
			ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
			FrameworkId:       "frb.fedwire.01",
		}
	}
	return msg
}
func TestBusinessApplicationHeaderFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_head.001")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.Doc.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId), "231981435")
	require.Equal(t, string(message.Doc.To.FIId.FinInstnId.ClrSysMmbId.MmbId), "021151080")
	require.Equal(t, string(message.Doc.BizMsgIdr), "20250311143738 ABAR MM Request")
	require.Equal(t, string(message.Doc.BizSvc), "TEST")
	require.Equal(t, string(message.Doc.MktPrctc.Regy), "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service")
	require.Equal(t, string(message.Doc.MktPrctc.Id), "frb.fedwire.01")
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
			Message{Data: MessageModel{MessageSenderId: "Unknown data"}},
			"error occur at MessageSenderId: Unknown data fails validation with pattern [A-Z0-9]{9,9}",
		},
		{
			"MessageReceiverId",
			Message{Data: MessageModel{MessageReceiverId: "Unknown data"}},
			"error occur at MessageReceiverId: Unknown data fails validation with pattern [A-Z0-9]{9,9}",
		},
		{
			"MessageDefinitionId",
			Message{Data: MessageModel{MessageDefinitionId: "-------"}},
			"error occur at MessageDefinitionId: ------- fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"MarketInfo - FrameworkId",
			Message{Data: MessageModel{MarketInfo: MarketPractice{
				ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
				FrameworkId:       "wwww",
			}}},
			"error occur at MarketInfo.FrameworkId: wwww fails validation with pattern frb([.]{1,1})fedwire([.]{1,1})(([a-z]{3,3}[.]{1,1})){0,1}01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			nMsg := generateRequreFields(tt.msg)
			msgErr := nMsg.CreateDocument()
			if msgErr != nil {
				require.Equal(t, tt.expectedErr, msgErr.Error())
			}
		})
	}
}
func TestScenario1_Step1_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "231981435"
	mesage.Data.MessageReceiverId = "021151080"
	mesage.Data.BusinessMessageId = "20250311143738 ABAR MM Request"
	mesage.Data.MessageDefinitionId = "camt.060.001.05"
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
	err = model.WriteXMLTo("AccountBalanceReport_Scenario1_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_head.001")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestScenario1_Step2_head(t *testing.T) {
	var mesage, err = NewMessage("")
	require.NoError(t, err)
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "231981435"
	mesage.Data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137a"
	mesage.Data.MessageDefinitionId = "camt.052.001.08"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.abm.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
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
	mesage.Data.MessageSenderId = "021151080"
	mesage.Data.MessageReceiverId = "231981435"
	mesage.Data.BusinessMessageId = "98z2cb3d0f2f3094f24a16389713541137N"
	mesage.Data.MessageDefinitionId = "camt.052.001.08"
	mesage.Data.BusinessService = "TEST"
	mesage.Data.MarketInfo = MarketPractice{
		ReferenceRegistry: "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
		FrameworkId:       "frb.fedwire.abs.01",
	}
	mesage.Data.CreateDatetime = time.Now()
	cErr := mesage.CreateDocument()
	require.NoError(t, cErr.ToError())
	xmlData, err := xml.MarshalIndent(&mesage.Doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountBalanceReport_Scenario2_Step1_head.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario2_Step1_head.001")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario2_Step1_head.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
