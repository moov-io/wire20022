package PaymentStatusRequest

import (
	"reflect"

	pacs004 "github.com/moov-io/fedwire20022/gen/PaymentStatusRequest_pacs_028_001_03"
	model "github.com/moov-io/wire20022/pkg/models"
)

func BranchAndFinancialInstitutionIdentification61From(p model.Agent) pacs004.BranchAndFinancialInstitutionIdentification61 {
	var result pacs004.BranchAndFinancialInstitutionIdentification61
	var FinInstnId pacs004.FinancialInstitutionIdentification181
	var ClrSysMmbId pacs004.ClearingSystemMemberIdentification21
	if p.PaymentSysCode != "" {
		Cd := pacs004.ExternalClearingSystemIdentification1CodeFixed(p.PaymentSysCode)
		ClrSysMmbId.ClrSysId = pacs004.ClearingSystemIdentification2Choice1{
			Cd: &Cd,
		}
	}
	if p.PaymentSysMemberId != "" {
		ClrSysMmbId.MmbId = pacs004.RoutingNumberFRS1(p.PaymentSysMemberId)
	}
	if !isEmpty(ClrSysMmbId) {
		FinInstnId.ClrSysMmbId = ClrSysMmbId
	}
	if !isEmpty(FinInstnId) {
		result.FinInstnId = FinInstnId
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
