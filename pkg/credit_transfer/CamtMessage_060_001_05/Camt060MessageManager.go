package CamtMessage_060_001_05

import (
	"reflect"

	AccountReportingRequest "github.com/moov-io/fedwire20022/gen/AccountReportingRequest_camt_060_001_05"
	"github.com/moov-io/wire20022/pkg/credit_transfer"
)

type FedwireFundsReportType string
type AccountTypeFRS string

const (
	ReportDetails    FedwireFundsReportType = "DTLS" // Detailed Report
	ReportSummary    FedwireFundsReportType = "DTLR" // Summary Report
	ReportTotal      FedwireFundsReportType = "ETOT" // End-of-Day Total
	ReportABARequest FedwireFundsReportType = "ABAR" // ABA Routing Request
)
const (
	AccountTypeSavings  AccountTypeFRS = "S" // "S" for Savings Account
	AccountTypeMerchant AccountTypeFRS = "M" // "M" for Merchant Account
)

type Camt060Agent struct {
	agent   credit_transfer.Agent
	OtherId string
}
type SequenceRange struct {
	FromSeq float64
	ToSeq   float64
}

func Party40Choice1From(agent credit_transfer.Agent) AccountReportingRequest.Party40Choice1 {
	var result AccountReportingRequest.Party40Choice1
	Agt := AccountReportingRequest.BranchAndFinancialInstitutionIdentification61{}
	if agent.PaymentSysCode != "" || agent.PaymentSysMemberId != "" {
		Agt.FinInstnId = AccountReportingRequest.FinancialInstitutionIdentification181{}
		_Cd := AccountReportingRequest.ExternalClearingSystemIdentification1CodeFixed(agent.PaymentSysCode)
		Agt.FinInstnId.ClrSysMmbId = AccountReportingRequest.ClearingSystemMemberIdentification21{}
		if agent.PaymentSysCode != "" {
			Agt.FinInstnId.ClrSysMmbId.ClrSysId = AccountReportingRequest.ClearingSystemIdentification2Choice1{
				Cd: &_Cd,
			}
		}
		if agent.PaymentSysMemberId != "" {
			Agt.FinInstnId.ClrSysMmbId.MmbId = AccountReportingRequest.RoutingNumberFRS1(agent.PaymentSysMemberId)
		}
	}
	if !isEmpty(Agt) {
		result = AccountReportingRequest.Party40Choice1{
			Agt: &Agt,
		}
	}
	return result
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
