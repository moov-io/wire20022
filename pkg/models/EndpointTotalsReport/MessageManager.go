package EndpointTotalsReport

import (
	"reflect"

	camt052 "github.com/moov-io/fedwire20022/gen/EndpointTotalsReport_camt_052_001_08"
	model "github.com/moov-io/wire20022/pkg/models"
)

func TotalsPerBankTransactionCode51From(m model.NumberAndStatusOfTransactions) camt052.TotalsPerBankTransactionCode51 {
	var result camt052.TotalsPerBankTransactionCode51
	if m.NumberOfEntries != "" {
		result.NbOfNtries = camt052.Max15NumericText(m.NumberOfEntries)
	}
	if m.Status != "" {
		result.BkTxCd = camt052.BankTransactionCodeStructure41{
			Prtry: camt052.ProprietaryBankTransactionCodeStructure11{
				Cd: camt052.BankTransactionCodeFedwireFunds1(m.Status),
			},
		}
	}
	return result
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
