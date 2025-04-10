package ActivityReport

import (
	"reflect"

	camt052 "github.com/moov-io/fedwire20022/gen/ActivityReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type TotalsPerBankTransactionCode struct {
	// NbOfNtries (Number of Entries) specifies the total number of transactions for a given bank transaction code.
	// This helps in categorizing transactions based on their type.
	NumberOfEntries string
	// It is used when the transaction code follows a bank-specific classification rather than a standard one.
	BankTransactionCode model.TransactionStatusCode
}

func TotalsPerBankTransactionCode51From(param TotalsPerBankTransactionCode) camt052.TotalsPerBankTransactionCode51 {
	var result camt052.TotalsPerBankTransactionCode51
	if param.NumberOfEntries != "" {
		result.NbOfNtries = camt052.Max15NumericText(param.NumberOfEntries)
	}
	if param.BankTransactionCode != "" {
		result.BkTxCd = camt052.BankTransactionCodeStructure41{
			Prtry: camt052.ProprietaryBankTransactionCodeStructure11{
				Cd: camt052.BankTransactionCodeFedwireFunds1(param.BankTransactionCode),
			},
		}
	}
	return result
}
func ReportEntry101From(param model.Entry) camt052.ReportEntry101 {
	var result camt052.ReportEntry101
	if !isEmpty(param.Amount) {
		result.Amt = camt052.ActiveOrHistoricCurrencyAndAmount{
			Value: camt052.ActiveOrHistoricCurrencyAndAmountSimpleType(param.Amount.Amount),
			Ccy:   camt052.ActiveOrHistoricCurrencyCode(param.Amount.Currency),
		}
	}
	if param.CreditDebitIndicator != "" {
		result.CdtDbtInd = camt052.CreditDebitCode(param.CreditDebitIndicator)
	}
	if param.Status != "" {
		_Cd := camt052.ExternalEntryStatus1Code(param.Status)
		result.Sts = camt052.EntryStatus1Choice1{
			Cd: &_Cd,
		}
	}
	if param.BankTransactionCode != "" {
		result.BkTxCd = camt052.BankTransactionCodeStructure42{
			Prtry: camt052.ProprietaryBankTransactionCodeStructure12{
				Cd: camt052.BankTransactionCodeFedwireFunds11(param.BankTransactionCode),
			},
		}
	}
	if param.MessageNameId != "" {
		result.AddtlInfInd = camt052.MessageIdentification21{
			MsgNmId: camt052.MessageNameIdentificationFRS1(param.MessageNameId),
		}
	}
	if !isEmpty(param.EntryDetails) {
		_InstrId := camt052.Max35Text(param.EntryDetails.InstructionId)
		_UETR := camt052.UUIDv4Identifier(param.EntryDetails.UniqueTransactionReference)
		_ClrSysRef := camt052.OMADFedwireFunds1(param.EntryDetails.ClearingSystemRef)
		result.NtryDtls = camt052.EntryDetails91{
			TxDtls: camt052.EntryTransaction101{
				Refs: camt052.TransactionReferences61{
					MsgId:     camt052.IMADFedwireFunds1(param.EntryDetails.MessageId),
					InstrId:   &_InstrId,
					UETR:      &_UETR,
					ClrSysRef: &_ClrSysRef,
				},
			},
		}
		var RltdAgts camt052.TransactionAgents51
		if !isEmpty(param.EntryDetails.InstructingAgent) {
			_Cd := camt052.ExternalClearingSystemIdentification1CodeFixed(param.EntryDetails.InstructingAgent.PaymentSysCode)
			RltdAgts.InstgAgt = camt052.BranchAndFinancialInstitutionIdentification61{
				FinInstnId: camt052.FinancialInstitutionIdentification181{
					ClrSysMmbId: camt052.ClearingSystemMemberIdentification21{
						ClrSysId: camt052.ClearingSystemIdentification2Choice1{
							Cd: &_Cd,
						},
						MmbId: camt052.RoutingNumberFRS1(param.EntryDetails.InstructingAgent.PaymentSysMemberId),
					},
				},
			}
		}
		if !isEmpty(param.EntryDetails.InstructedAgent) {
			_Cd := camt052.ExternalClearingSystemIdentification1CodeFixed(param.EntryDetails.InstructedAgent.PaymentSysCode)
			RltdAgts.InstdAgt = camt052.BranchAndFinancialInstitutionIdentification61{
				FinInstnId: camt052.FinancialInstitutionIdentification181{
					ClrSysMmbId: camt052.ClearingSystemMemberIdentification21{
						ClrSysId: camt052.ClearingSystemIdentification2Choice1{
							Cd: &_Cd,
						},
						MmbId: camt052.RoutingNumberFRS1(param.EntryDetails.InstructedAgent.PaymentSysMemberId),
					},
				},
			}
		}
		if !isEmpty(RltdAgts) {
			result.NtryDtls.TxDtls.RltdAgts = RltdAgts
		}
		if param.EntryDetails.LocalInstrumentChoice != "" {
			_Prtry := camt052.LocalInstrumentFedwireFunds1(param.EntryDetails.LocalInstrumentChoice)
			_LclInstrm := camt052.LocalInstrument2Choice1{
				Prtry: &_Prtry,
			}
			result.NtryDtls.TxDtls.LclInstrm = &_LclInstrm

		}
		if param.EntryDetails.RelatedDatesProprietary != "" {
			_DtTm := fedwire.ISODateTime(param.EntryDetails.RelatedDateTime)
			result.NtryDtls.TxDtls.RltdDts = camt052.TransactionDates31{
				Prtry: camt052.ProprietaryDate31{
					Tp: camt052.ReportDatesFedwireFunds1(param.EntryDetails.RelatedDatesProprietary),
					Dt: camt052.DateAndDateTime2Choice1{
						DtTm: &_DtTm,
					},
				},
			}
		}
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
