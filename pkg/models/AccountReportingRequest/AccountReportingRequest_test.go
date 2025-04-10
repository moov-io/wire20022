package AccountReportingRequest

import (
	"encoding/xml"
	"testing"
	"time"

	model "github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestAccountBalanceReport_Scenario1_Step1_camt_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250311114001500ABARSrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointDetailsReceivedReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOtherId = "114001500"
	message.data.AccountProperty = AccountTypeMerchant
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "114001500",
		},
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("AccountBalanceReport_Scenario1_Step1_MM_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestAccountBalanceReport_Scenario1_Step1_camt_MS_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("AccountBalanceReport_Scenario1_Step1__MS_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestAccountBalanceReport_Scenario1_Step1_camt_SM_CreateXML(t *testing.T) {
	var message = NewMessage()
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

	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("AccountBalanceReport_Scenario1_Step1__SM_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestAccountBalanceReport_Scenario1_Step1_camt_SS_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("AccountBalanceReport_Scenario1_Step1__SS_camt.xml", xmlData)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_M_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("AccountReportingRequest_Step1_camt_M.xml", xmlData)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_S_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250311231981435ABARMrequest1"
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
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("AccountReportingRequest_Step1_camt_S.xml", xmlData)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_DTLR_CreateXML(t *testing.T) {
	var message = NewMessage()
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
		FromSeq: 000001,
		ToSeq:   000001,
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("AccountReportingRequest_Step1_camt_DTLR.xml", xmlData)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_DTLS_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250311231981435DTLSrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointDetailsReceivedReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	message.data.FromToSeuence = model.SequenceRange{
		FromSeq: 000100,
		ToSeq:   000200,
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("AccountReportingRequest_Step1_camt_DTLS.xml", xmlData)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_ETOT_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250311231981435ETOTrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointTotalsReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("AccountReportingRequest_Step1_camt_ETOT.xml", xmlData)
	require.NoError(t, err)
}
func TestEndpointDetailsReport_Scenario1_Step1_camt_DTLS_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250311231981435DTLSrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointDetailsReceivedReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	message.data.FromToSeuence = model.SequenceRange{
		FromSeq: 000001,
		ToSeq:   000100,
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("EndpointDetailsReport_Scenario1_Step1_camt_DTLS.xml", xmlData)
	require.NoError(t, err)
}
func TestEndpointDetailsReport_Scenario1_Step1_camt_DTLR_CreateXML(t *testing.T) {
	var message = NewMessage()
	message.data.MessageId = "20250311231981435DTLRrequest1"
	message.data.CreatedDateTime = time.Now()
	message.data.ReportRequestId = model.EndpointDetailsReceivedReport
	message.data.RequestedMsgNameId = "camt.052.001.08"
	message.data.AccountOwnerAgent = Camt060Agent{
		agent: model.Agent{
			PaymentSysCode:     model.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	message.data.FromToSeuence = model.SequenceRange{
		FromSeq: 000001,
		ToSeq:   000100,
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("EndpointDetailsReport_Scenario1_Step1_camt_DTLR.xml", xmlData)
	require.NoError(t, err)
}
func TestEndpointDetailsReport_Scenario1_Step1_camt_ETOT_CreateXML(t *testing.T) {
	var message = NewMessage()
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
	message.data.FromToSeuence = model.SequenceRange{
		FromSeq: 000001,
		ToSeq:   000100,
	}
	message.CreateDocument()
	xmlData, err := xml.MarshalIndent(&message.doc, "", "\t")
	WriteXMLTo("EndpointDetailsReport_Scenario1_Step1_camt_ETOT.xml", xmlData)
	require.NoError(t, err)
}
