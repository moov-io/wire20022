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

func TotalsPerBankTransactionCode51From(param TotalsPerBankTransactionCode) (camt052.TotalsPerBankTransactionCode51, *model.ValidateError) {
	var result camt052.TotalsPerBankTransactionCode51
	if param.NumberOfEntries != "" {
		err := camt052.Max15NumericText(param.NumberOfEntries).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "NumberOfEntries",
				Message:   err.Error(),
			}
		}
		result.NbOfNtries = camt052.Max15NumericText(param.NumberOfEntries)
	}
	if param.BankTransactionCode != "" {
		err := camt052.BankTransactionCodeFedwireFunds1(param.BankTransactionCode).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "BankTransactionCode",
				Message:   err.Error(),
			}
		}
		result.BkTxCd = camt052.BankTransactionCodeStructure41{
			Prtry: camt052.ProprietaryBankTransactionCodeStructure11{
				Cd: camt052.BankTransactionCodeFedwireFunds1(param.BankTransactionCode),
			},
		}
	}
	return result, nil
}
func ReportEntry101From(param model.Entry) (camt052.ReportEntry101, *model.ValidateError) {
	var result camt052.ReportEntry101
	if !isEmpty(param.Amount) {
		err := fedwire.Amount(param.Amount.Amount).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "Amount",
				Message:   err.Error(),
			}
			vErr.InsertPath("Amount")
			return camt052.ReportEntry101{}, &vErr
		}
		err = camt052.ActiveOrHistoricCurrencyCode(param.Amount.Currency).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "Currency",
				Message:   err.Error(),
			}
			vErr.InsertPath("Amount")
			return camt052.ReportEntry101{}, &vErr
		}
		result.Amt = camt052.ActiveOrHistoricCurrencyAndAmount{
			Value: camt052.ActiveOrHistoricCurrencyAndAmountSimpleType(param.Amount.Amount),
			Ccy:   camt052.ActiveOrHistoricCurrencyCode(param.Amount.Currency),
		}
	}
	if param.CreditDebitIndicator != "" {
		err := camt052.CreditDebitCode(param.CreditDebitIndicator).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "CreditDebitIndicator",
				Message:   err.Error(),
			}
		}
		result.CdtDbtInd = camt052.CreditDebitCode(param.CreditDebitIndicator)
	}
	if param.Status != "" {
		err := camt052.ExternalEntryStatus1Code(param.Status).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "Status",
				Message:   err.Error(),
			}
		}
		_Cd := camt052.ExternalEntryStatus1Code(param.Status)
		result.Sts = camt052.EntryStatus1Choice1{
			Cd: &_Cd,
		}
	}
	if param.BankTransactionCode != "" {
		err := camt052.BankTransactionCodeFedwireFunds11(param.BankTransactionCode).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "BankTransactionCode",
				Message:   err.Error(),
			}
		}
		result.BkTxCd = camt052.BankTransactionCodeStructure42{
			Prtry: camt052.ProprietaryBankTransactionCodeStructure12{
				Cd: camt052.BankTransactionCodeFedwireFunds11(param.BankTransactionCode),
			},
		}
	}
	if param.MessageNameId != "" {
		err := camt052.MessageNameIdentificationFRS1(param.MessageNameId).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "MessageNameId",
				Message:   err.Error(),
			}
		}
		result.AddtlInfInd = camt052.MessageIdentification21{
			MsgNmId: camt052.MessageNameIdentificationFRS1(param.MessageNameId),
		}
	}
	if !isEmpty(param.EntryDetails) {
		err := camt052.Max35Text(param.EntryDetails.InstructionId).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "InstructionId",
				Message:   err.Error(),
			}
			vErr.InsertPath("EntryDetails")
			return camt052.ReportEntry101{}, &vErr
		}
		_InstrId := camt052.Max35Text(param.EntryDetails.InstructionId)
		err = camt052.UUIDv4Identifier(param.EntryDetails.UniqueTransactionReference).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "UniqueTransactionReference",
				Message:   err.Error(),
			}
			vErr.InsertPath("EntryDetails")
			return camt052.ReportEntry101{}, &vErr
		}
		_UETR := camt052.UUIDv4Identifier(param.EntryDetails.UniqueTransactionReference)
		err = camt052.OMADFedwireFunds1(param.EntryDetails.ClearingSystemRef).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "ClearingSystemRef",
				Message:   err.Error(),
			}
			vErr.InsertPath("EntryDetails")
			return camt052.ReportEntry101{}, &vErr
		}
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
			err := camt052.ExternalClearingSystemIdentification1CodeFixed(param.EntryDetails.InstructingAgent.PaymentSysCode).Validate()
			if err != nil {
				vErr := model.ValidateError{
					ParamName: "PaymentSysCode",
					Message:   err.Error(),
				}
				vErr.InsertPath("InstructingAgent")
				vErr.InsertPath("EntryDetails")
				return camt052.ReportEntry101{}, &vErr
			}
			err = camt052.RoutingNumberFRS1(param.EntryDetails.InstructingAgent.PaymentSysMemberId).Validate()
			if err != nil {
				vErr := model.ValidateError{
					ParamName: "PaymentSysMemberId",
					Message:   err.Error(),
				}
				vErr.InsertPath("InstructingAgent")
				vErr.InsertPath("EntryDetails")
				return camt052.ReportEntry101{}, &vErr
			}
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
			err := camt052.ExternalClearingSystemIdentification1CodeFixed(param.EntryDetails.InstructedAgent.PaymentSysCode).Validate()
			if err != nil {
				vErr := model.ValidateError{
					ParamName: "PaymentSysCode",
					Message:   err.Error(),
				}
				vErr.InsertPath("InstructedAgent")
				vErr.InsertPath("EntryDetails")
				return camt052.ReportEntry101{}, &vErr
			}
			err = camt052.RoutingNumberFRS1(param.EntryDetails.InstructedAgent.PaymentSysMemberId).Validate()
			if err != nil {
				vErr := model.ValidateError{
					ParamName: "PaymentSysMemberId",
					Message:   err.Error(),
				}
				vErr.InsertPath("InstructedAgent")
				vErr.InsertPath("EntryDetails")
				return camt052.ReportEntry101{}, &vErr
			}
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

			err := camt052.LocalInstrumentFedwireFunds1(param.EntryDetails.LocalInstrumentChoice).Validate()
			if err != nil {
				vErr := model.ValidateError{
					ParamName: "LocalInstrumentChoice",
					Message:   err.Error(),
				}
				vErr.InsertPath("EntryDetails")
				return camt052.ReportEntry101{}, &vErr
			}
			_Prtry := camt052.LocalInstrumentFedwireFunds1(param.EntryDetails.LocalInstrumentChoice)
			_LclInstrm := camt052.LocalInstrument2Choice1{
				Prtry: &_Prtry,
			}
			result.NtryDtls.TxDtls.LclInstrm = &_LclInstrm

		}
		if param.EntryDetails.RelatedDatesProprietary != "" {
			err := camt052.ReportDatesFedwireFunds1(param.EntryDetails.RelatedDatesProprietary).Validate()
			if err != nil {
				vErr := model.ValidateError{
					ParamName: "RelatedDatesProprietary",
					Message:   err.Error(),
				}
				vErr.InsertPath("EntryDetails")
				return camt052.ReportEntry101{}, &vErr
			}
			err = fedwire.ISODateTime(param.EntryDetails.RelatedDateTime).Validate()
			if err != nil {
				vErr := model.ValidateError{
					ParamName: "RelatedDateTime",
					Message:   err.Error(),
				}
				vErr.InsertPath("EntryDetails")
				return camt052.ReportEntry101{}, &vErr
			}
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
	return result, nil
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
