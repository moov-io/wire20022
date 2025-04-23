package AccountReportingRequest

import (
	"encoding/xml"
	"path/filepath"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestAccountBalanceReportFromXMLFile(t *testing.T) {
	xmlFilePath := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_camt.060_ABAR_MM")
	var message, err = NewMessage(xmlFilePath)
	require.NoError(t, err)
	require.Equal(t, string(message.doc.AcctRptgReq.GrpHdr.MsgId), "20250311231981435ABARMMrequest1")
	require.Equal(t, string(message.doc.AcctRptgReq.RptgReq.Id), "ABAR")
	require.Equal(t, string(message.doc.AcctRptgReq.RptgReq.ReqdMsgNmId), "camt.052.001.08")
	require.Equal(t, string(message.doc.AcctRptgReq.RptgReq.Acct.Id.Othr.Id), "231981435")
	require.Equal(t, string(*message.doc.AcctRptgReq.RptgReq.AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd), "USABA")
	require.Equal(t, string(message.doc.AcctRptgReq.RptgReq.AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId), "231981435")
}
func TestAccountBalanceReportValidator(t *testing.T) {
	tests := []struct {
		title       string
		msg         Message
		expectedErr string
	}{
		{
			"MessageId",
			Message{data: MessageModel{MessageId: "20250311231981435ABARMMrequest120250311231981435ABARMMrequest1"}},
			"error occur at MessageId: 20250311231981435ABARMMrequest120250311231981435ABARMMrequest1 fails validation with length 62 <= required maxLength 35",
		},
		{
			"ReportRequestId",
			Message{data: MessageModel{ReportRequestId: "Unknown"}},
			"error occur at ReportRequestId: invalid CAMT report type: Unknown",
		},
		{
			"RequestedMsgNameId",
			Message{data: MessageModel{RequestedMsgNameId: "ABCD12300199"}},
			"error occur at RequestedMsgNameId: ABCD12300199 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"AccountOtherId",
			Message{data: MessageModel{RequestedMsgNameId: "123ABC789"}},
			"error occur at RequestedMsgNameId: 123ABC789 fails validation with pattern [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}",
		},
		{
			"AccountProperty",
			Message{data: MessageModel{AccountProperty: "Unknown"}},
			"error occur at AccountProperty: invalid AccountTypeFRS: Unknown",
		},
		{
			"AccountOwnerAgent - agent - PaymentSysCode",
			Message{data: MessageModel{AccountOwnerAgent: Camt060Agent{
				agent: model.Agent{
					PaymentSysCode:     "unknown",
					PaymentSysMemberId: "231981435",
				},
			}}},
			"error occur at AccountOwnerAgent.agent.PaymentSysCode: unknown fails enumeration validation",
		},
		{
			"AccountOwnerAgent - agent - PaymentSysMemberId",
			Message{data: MessageModel{AccountOwnerAgent: Camt060Agent{
				agent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "unknown",
				},
			}}},
			"error occur at AccountOwnerAgent.agent.PaymentSysMemberId: unknown fails validation with pattern [0-9]{9,9}",
		},
		{
			"AccountOwnerAgent - agent - OtherId",
			Message{data: MessageModel{AccountOwnerAgent: Camt060Agent{
				agent: model.Agent{
					PaymentSysCode:     model.PaymentSysUSABA,
					PaymentSysMemberId: "231981435",
				},
				OtherId: "unknown",
			}}},
			"error occur at AccountOwnerAgent.agent.OtherId: unknown fails validation with pattern [A-Z0-9]{8,8}",
		},
		{
			"FromToSeuence - FromToSeuence - FromSeq",
			Message{data: MessageModel{FromToSeuence: model.SequenceRange{
				FromSeq: "unknown",
				ToSeq:   "000100",
			}}},
			"error occur at FromToSeuence.FromSeq: strconv.ParseFloat: parsing \"unknown\": invalid syntax",
		},
		{
			"FromToSeuence - FromToSeuence - ToSeq",
			Message{data: MessageModel{FromToSeuence: model.SequenceRange{
				FromSeq: "000100",
				ToSeq:   "unknown",
			}}},
			"error occur at FromToSeuence.ToSeq: strconv.ParseFloat: parsing \"unknown\": invalid syntax",
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			msgErr := tt.msg.CreateDocument()
			require.Equal(t, tt.expectedErr, msgErr.Error())
		})
	}
}
func TestAccountBalanceReport_Scenario1_Step1_camt_MM_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20250311231981435ABARMMrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.AccountBalanceReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOtherId = "231981435"
	message.data.AccountProperty = AccountTypeMerchant
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountBalanceReport_Scenario1_Step1_camt_MM.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_camt.060_ABAR_MM")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1_camt_MM.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestAccountBalanceReport_Scenario1_Step1_camt_MS_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20230921231981435ABARMSrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.AccountBalanceReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOtherId = "231981435"
	message.data.AccountProperty = AccountTypeSavings
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountBalanceReport_Scenario1_Step1__MS_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_camt.060_ABAR_MS")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1__MS_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestAccountBalanceReport_Scenario1_Step1_camt_SM_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20230921231981435ABARSMrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.AccountBalanceReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOtherId = "231981435"
	message.data.AccountProperty = AccountTypeMerchant
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "114001500",
		},
	}

	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountBalanceReport_Scenario1_Step1_SM_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_camt.060_ABAR_SM")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1_SM_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestAccountBalanceReport_Scenario1_Step1_camt_SS_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20230921231981435ABARSSrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.AccountBalanceReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOtherId = "114001500"
	message.data.AccountProperty = AccountTypeSavings
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "114001500",
		},
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountBalanceReport_Scenario1_Step1__SS_camt.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountBalanceReport_Scenario1_Step1_camt.060_ABAR_SS")
	genterated := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1__SS_camt.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestAccountReportingRequest_Step1_camt_M_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20250311231981435ABARMrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.AccountBalanceReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOtherId = "231981435"
	message.data.AccountProperty = AccountTypeMerchant
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountReportingRequest_Step1_camt_M.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountReportingRequest_Step1_camt.060_ABAR_M")
	genterated := filepath.Join("generated", "AccountReportingRequest_Step1_camt_M.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestAccountReportingRequest_Step1_camt_S_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20250311114001500ABARSrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.AccountBalanceReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOtherId = "114001500"
	message.data.AccountProperty = AccountTypeSavings
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "114001500",
		},
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountReportingRequest_Step1_camt_S.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountReportingRequest_Step1_camt.060_ABAR_S")
	genterated := filepath.Join("generated", "AccountReportingRequest_Step1_camt_S.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestAccountReportingRequest_Step1_camt_DTLR_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20250311231981435DTLRrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointDetailsReceivedReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "QMGFT001",
	}
	message.data.FromToSeuence = model.SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000003",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountReportingRequest_Step1_camt_DTLR.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountReportingRequest_Step1_camt.060_DTLR")
	genterated := filepath.Join("generated", "AccountReportingRequest_Step1_camt_DTLR.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestAccountReportingRequest_Step1_camt_DTLS_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20250311231981435DTLSrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointDetailsSentReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	message.data.FromToSeuence = model.SequenceRange{
		FromSeq: "000100",
		ToSeq:   "000200",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountReportingRequest_Step1_camt_DTLS.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountReportingRequest_Step1_camt.060_DTLS")
	genterated := filepath.Join("generated", "AccountReportingRequest_Step1_camt_DTLS.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestAccountReportingRequest_Step1_camt_ETOT_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20250311231981435ETOTrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointTotalsReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("AccountReportingRequest_Step1_camt_ETOT.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "AccountReportingRequest_Step1_camt.060_ETOT")
	genterated := filepath.Join("generated", "AccountReportingRequest_Step1_camt_ETOT.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointDetailsReport_Scenario1_Step1_camt_DTLS_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20250311231981435DTLSrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointDetailsSentReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	message.data.FromToSeuence = model.SequenceRange{
		FromSeq: "000002",
		ToSeq:   "000100",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario1_Step1_camt_DTLS.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario1_Step1_camt.060_DTLS")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step1_camt_DTLS.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointDetailsReport_Scenario1_Step1_camt_DTLR_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20250311231981435DTLRrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointDetailsSentReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	message.data.FromToSeuence = model.SequenceRange{
		FromSeq: "2",
		ToSeq:   "000100",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario1_Step1_camt_DTLR.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointDetailsReport_Scenario2_Step1_camt.060_DTLR")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step1_camt_DTLR.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
func TestEndpointDetailsReport_Scenario1_Step1_camt_ETOT_CreateXML(t *testing.T) {
	var message, err = NewMessage("")
	require.NoError(t, err)
	message.data.MessageId = "20250311231981435ETOTrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointTotalsReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	cErr := message.CreateDocument()
	require.Nil(t, cErr)
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	require.NoError(t, err)
	err = model.WriteXMLTo("EndpointDetailsReport_Scenario1_Step1_camt_ETOT.xml", xmlData)
	require.NoError(t, err)

	swiftSample := filepath.Join("swiftSample", "EndpointTotalsReport_Scenario1_Step1_camt.060_ETOT")
	genterated := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step1_camt_ETOT.xml")
	require.True(t, model.CompareXMLs(swiftSample, genterated))
}
