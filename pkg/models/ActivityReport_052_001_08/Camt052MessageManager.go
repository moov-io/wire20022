package ActivityReport_052_001_08

import (
	"reflect"

	ActivityReport "github.com/moov-io/fedwire20022/gen/ActivityReport_camt_052_001_08"
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

func TotalsPerBankTransactionCode51From(param TotalsPerBankTransactionCode) ActivityReport.TotalsPerBankTransactionCode51 {
	var result ActivityReport.TotalsPerBankTransactionCode51
	if param.NumberOfEntries != "" {
		result.NbOfNtries = ActivityReport.Max15NumericText(param.NumberOfEntries)
	}
	if param.BankTransactionCode != "" {
		result.BkTxCd = ActivityReport.BankTransactionCodeStructure41{
			Prtry: ActivityReport.ProprietaryBankTransactionCodeStructure11{
				Cd: ActivityReport.BankTransactionCodeFedwireFunds1(param.BankTransactionCode),
			},
		}
	}
	return result
}
func ReportEntry101From(param model.Entry) ActivityReport.ReportEntry101 {
	var result ActivityReport.ReportEntry101
	if !isEmpty(param.Amount) {
		result.Amt = ActivityReport.ActiveOrHistoricCurrencyAndAmount{
			Value: ActivityReport.ActiveOrHistoricCurrencyAndAmountSimpleType(param.Amount.Amount),
			Ccy:   ActivityReport.ActiveOrHistoricCurrencyCode(param.Amount.Currency),
		}
	}
	if param.CreditDebitIndicator != "" {
		result.CdtDbtInd = ActivityReport.CreditDebitCode(param.CreditDebitIndicator)
	}
	if param.Status != "" {
		_Cd := ActivityReport.ExternalEntryStatus1Code(param.Status)
		result.Sts = ActivityReport.EntryStatus1Choice1{
			Cd: &_Cd,
		}
	}
	if param.BankTransactionCode != "" {
		result.BkTxCd = ActivityReport.BankTransactionCodeStructure42{
			Prtry: ActivityReport.ProprietaryBankTransactionCodeStructure12{
				Cd: ActivityReport.BankTransactionCodeFedwireFunds11(param.BankTransactionCode),
			},
		}
	}
	if param.MessageNameId != "" {
		result.AddtlInfInd = ActivityReport.MessageIdentification21{
			MsgNmId: ActivityReport.MessageNameIdentificationFRS1(param.MessageNameId),
		}
	}
	if !isEmpty(param.EntryDetails) {
		_InstrId := ActivityReport.Max35Text(param.EntryDetails.InstructionId)
		_UETR := ActivityReport.UUIDv4Identifier(param.EntryDetails.UniqueTransactionReference)
		_ClrSysRef := ActivityReport.OMADFedwireFunds1(param.EntryDetails.ClearingSystemRef)
		result.NtryDtls = ActivityReport.EntryDetails91{
			TxDtls: ActivityReport.EntryTransaction101{
				Refs: ActivityReport.TransactionReferences61{
					MsgId:     ActivityReport.IMADFedwireFunds1(param.EntryDetails.MessageId),
					InstrId:   &_InstrId,
					UETR:      &_UETR,
					ClrSysRef: &_ClrSysRef,
				},
			},
		}
		var RltdAgts ActivityReport.TransactionAgents51
		if !isEmpty(param.EntryDetails.InstructingAgent) {
			_Cd := ActivityReport.ExternalClearingSystemIdentification1CodeFixed(param.EntryDetails.InstructingAgent.PaymentSysCode)
			RltdAgts.InstgAgt = ActivityReport.BranchAndFinancialInstitutionIdentification61{
				FinInstnId: ActivityReport.FinancialInstitutionIdentification181{
					ClrSysMmbId: ActivityReport.ClearingSystemMemberIdentification21{
						ClrSysId: ActivityReport.ClearingSystemIdentification2Choice1{
							Cd: &_Cd,
						},
						MmbId: ActivityReport.RoutingNumberFRS1(param.EntryDetails.InstructingAgent.PaymentSysMemberId),
					},
				},
			}
		}
		if !isEmpty(param.EntryDetails.InstructedAgent) {
			_Cd := ActivityReport.ExternalClearingSystemIdentification1CodeFixed(param.EntryDetails.InstructedAgent.PaymentSysCode)
			RltdAgts.InstdAgt = ActivityReport.BranchAndFinancialInstitutionIdentification61{
				FinInstnId: ActivityReport.FinancialInstitutionIdentification181{
					ClrSysMmbId: ActivityReport.ClearingSystemMemberIdentification21{
						ClrSysId: ActivityReport.ClearingSystemIdentification2Choice1{
							Cd: &_Cd,
						},
						MmbId: ActivityReport.RoutingNumberFRS1(param.EntryDetails.InstructedAgent.PaymentSysMemberId),
					},
				},
			}
		}
		if !isEmpty(RltdAgts) {
			result.NtryDtls.TxDtls.RltdAgts = RltdAgts
		}
		if param.EntryDetails.LocalInstrumentChoice != "" {
			_Prtry := ActivityReport.LocalInstrumentFedwireFunds1(param.EntryDetails.LocalInstrumentChoice)
			_LclInstrm := ActivityReport.LocalInstrument2Choice1{
				Prtry: &_Prtry,
			}
			result.NtryDtls.TxDtls.LclInstrm = &_LclInstrm

		}
		if param.EntryDetails.RelatedDatesProprietary != "" {
			_DtTm := fedwire.ISODateTime(param.EntryDetails.RelatedDateTime)
			result.NtryDtls.TxDtls.RltdDts = ActivityReport.TransactionDates31{
				Prtry: ActivityReport.ProprietaryDate31{
					Tp: ActivityReport.ReportDatesFedwireFunds1(param.EntryDetails.RelatedDatesProprietary),
					Dt: ActivityReport.DateAndDateTime2Choice1{
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
