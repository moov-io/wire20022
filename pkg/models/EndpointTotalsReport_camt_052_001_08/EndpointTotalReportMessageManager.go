package EndpointTotalsReport_camt_052_001_08

import (
	"reflect"

	EndpointTotalsReport "github.com/moov-io/fedwire20022/gen/EndpointTotalsReport_camt_052_001_08"
	model "github.com/moov-io/wire20022/pkg/models"
)

func TotalsPerBankTransactionCode51From(m model.NumberAndStatusOfTransactions) EndpointTotalsReport.TotalsPerBankTransactionCode51 {
	var result EndpointTotalsReport.TotalsPerBankTransactionCode51
	if m.NumberOfEntries != "" {
		result.NbOfNtries = EndpointTotalsReport.Max15NumericText(m.NumberOfEntries)
	}
	if m.Status != "" {
		result.BkTxCd = EndpointTotalsReport.BankTransactionCodeStructure41{
			Prtry: EndpointTotalsReport.ProprietaryBankTransactionCodeStructure11{
				Cd: EndpointTotalsReport.BankTransactionCodeFedwireFunds1(m.Status),
			},
		}
	}
	return result
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
