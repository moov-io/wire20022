package AccountReportingRequest

import (
	"fmt"
	"reflect"

	camt060 "github.com/moov-io/fedwire20022/gen/AccountReportingRequest_camt_060_001_05"
	model "github.com/moov-io/wire20022/pkg/models"
)

type AccountTypeFRS string

const (
	AccountTypeSavings  AccountTypeFRS = "S" // "S" for Savings Account
	AccountTypeMerchant AccountTypeFRS = "M" // "M" for Merchant Account
)

type Camt060Agent struct {
	Agent   model.Agent
	OtherId string
}

func Party40Choice1From(agent model.Agent) (camt060.Party40Choice1, *model.ValidateError) {
	var result camt060.Party40Choice1
	Agt := camt060.BranchAndFinancialInstitutionIdentification61{}
	if agent.PaymentSysCode != "" || agent.PaymentSysMemberId != "" {
		err := camt060.ExternalClearingSystemIdentification1CodeFixed(agent.PaymentSysCode).Validate()
		if err != nil {
			return camt060.Party40Choice1{}, &model.ValidateError{
				ParamName: "PaymentSysCode",
				Message:   err.Error(),
			}
		}
		Agt.FinInstnId = camt060.FinancialInstitutionIdentification181{}
		_Cd := camt060.ExternalClearingSystemIdentification1CodeFixed(agent.PaymentSysCode)
		Agt.FinInstnId.ClrSysMmbId = camt060.ClearingSystemMemberIdentification21{}
		if agent.PaymentSysCode != "" {
			Agt.FinInstnId.ClrSysMmbId.ClrSysId = camt060.ClearingSystemIdentification2Choice1{
				Cd: &_Cd,
			}
		}
		if agent.PaymentSysMemberId != "" {
			err := camt060.RoutingNumberFRS1(agent.PaymentSysMemberId).Validate()
			if err != nil {
				return camt060.Party40Choice1{}, &model.ValidateError{
					ParamName: "PaymentSysMemberId",
					Message:   err.Error(),
				}
			}
			Agt.FinInstnId.ClrSysMmbId.MmbId = camt060.RoutingNumberFRS1(agent.PaymentSysMemberId)
		}
	}
	if !isEmpty(Agt) {
		result = camt060.Party40Choice1{
			Agt: &Agt,
		}
	}
	return result, nil
}
func Party40Choice1To(agent camt060.Party40Choice1) model.Agent {
	var result model.Agent
	result.PaymentSysCode = model.PaymentSystemType(*agent.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	result.PaymentSysMemberId = string(agent.Agt.FinInstnId.ClrSysMmbId.MmbId)
	return result
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
func (a AccountTypeFRS) Validate() error {
	switch a {
	case AccountTypeSavings, AccountTypeMerchant:
		return nil
	}
	return fmt.Errorf("invalid AccountTypeFRS: %s", a)
}
