package CamtMessage_060_001_05

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/credit_transfer"
	"github.com/stretchr/testify/require"
)

func TestAccountBalanceReport_Scenario1_Step1_camt_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20250311114001500ABARSrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportDetails
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOtherId = "114001500"
	mesage.model.AccountProperty = AccountTypeMerchant
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "114001500",
		},
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1_MM_camt.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestAccountBalanceReport_Scenario1_Step1_camt_MS_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20230921231981435ABARMSrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportABARequest
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOtherId = "231981435"
	mesage.model.AccountProperty = AccountTypeSavings
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1__MS_camt.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestAccountBalanceReport_Scenario1_Step1_camt_SM_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20230921231981435ABARSMrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportABARequest
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOtherId = "231981435"
	mesage.model.AccountProperty = AccountTypeMerchant
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "114001500",
		},
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1__SM_camt.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestAccountBalanceReport_Scenario1_Step1_camt_SS_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20230921231981435ABARSSrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportABARequest
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOtherId = "114001500"
	mesage.model.AccountProperty = AccountTypeSavings
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "114001500",
		},
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountBalanceReport_Scenario1_Step1__SS_camt.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_M_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20250311231981435ABARMrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportABARequest
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOtherId = "231981435"
	mesage.model.AccountProperty = AccountTypeMerchant
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountReportingRequest_Step1_camt_M.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_S_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20250311231981435ABARMrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportABARequest
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOtherId = "114001500"
	mesage.model.AccountProperty = AccountTypeSavings
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "114001500",
		},
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountReportingRequest_Step1_camt_S.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_DTLR_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20250311231981435DTLRrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportSummary
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "QMGFT001",
	}
	mesage.model.FromToSeuence = SequenceRange{
		FromSeq: 000001,
		ToSeq:   000001,
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountReportingRequest_Step1_camt_DTLR.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_DTLS_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20250311231981435DTLSrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportDetails
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	mesage.model.FromToSeuence = SequenceRange{
		FromSeq: 000100,
		ToSeq:   000200,
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountReportingRequest_Step1_camt_DTLS.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestAccountReportingRequest_Step1_camt_ETOT_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20250311231981435ETOTrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportTotal
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "AccountReportingRequest_Step1_camt_ETOT.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestEndpointDetailsReport_Scenario1_Step1_camt_DTLS_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20250311231981435DTLSrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportDetails
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	mesage.model.FromToSeuence = SequenceRange{
		FromSeq: 000001,
		ToSeq:   000100,
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step1_camt_DTLS.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestEndpointDetailsReport_Scenario1_Step1_camt_DTLR_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20250311231981435DTLRrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportDetails
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	mesage.model.FromToSeuence = SequenceRange{
		FromSeq: 000001,
		ToSeq:   000100,
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step1_camt_DTLR.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
func TestEndpointDetailsReport_Scenario1_Step1_camt_ETOT_CreateXML(t *testing.T) {
	var mesage = NewCamt060MessageMessage()
	mesage.model.MessageId = "20250311231981435ETOTrequest1"
	mesage.model.CreatedDateTime = time.Now()
	mesage.model.ReportRequestId = ReportTotal
	mesage.model.RequestedMsgNameId = "camt.052.001.08"
	mesage.model.AccountOwnerAgent = Camt060Agent{
		agent: credit_transfer.Agent{
			PaymentSysCode:     credit_transfer.PaymentSysUSABA,
			PaymentSysMemberId: "231981435",
		},
		OtherId: "B1QDRCQR",
	}
	mesage.model.FromToSeuence = SequenceRange{
		FromSeq: 000001,
		ToSeq:   000100,
	}
	mesage.CreateDocument()
	// jsonData, err := mesage.GetJson()
	// require.NoError(t, err)
	xmlData, err := mesage.GetXML()
	require.NoError(t, err)
	os.Mkdir("generated", 0755)
	// jsonFileName := filepath.Join("generated", "CustomerCreditTransfer_Scenario1_Step1.json")
	xnlFileName := filepath.Join("generated", "EndpointDetailsReport_Scenario1_Step1_camt_ETOT.xml")
	// err = os.WriteFile(jsonFileName, jsonData, 0644)
	// require.NoError(t, err)
	err = os.WriteFile(xnlFileName, xmlData, 0644)
	require.NoError(t, err)
}
