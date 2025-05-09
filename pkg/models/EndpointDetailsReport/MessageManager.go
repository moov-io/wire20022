package EndpointDetailsReport

import (
	"reflect"

	camt052 "github.com/moov-io/fedwire20022/gen/EndpointDetailsReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

func EntryDetails91From(p model.EntryDetail) camt052.EntryDetails91 {
	var result camt052.EntryDetails91
	var TxDtls camt052.EntryTransaction101
	var Refs camt052.TransactionReferences61
	if p.MessageId != "" {
		Refs.MsgId = camt052.IMADFedwireFunds1(p.MessageId)
	}
	if p.InstructionId != "" {
		InstrId := camt052.Max35Text(p.InstructionId)
		Refs.InstrId = &InstrId
	}
	if p.UniqueTransactionReference != "" {
		UETR := camt052.UUIDv4Identifier(p.UniqueTransactionReference)
		Refs.UETR = &UETR
	}
	if p.ClearingSystemRef != "" {
		ClrSysRef := camt052.OMADFedwireFunds1(p.ClearingSystemRef)
		Refs.ClrSysRef = &ClrSysRef
	}
	if !isEmpty(Refs) {
		TxDtls.Refs = Refs
	}
	var RltdAgts camt052.TransactionAgents51
	if !isEmpty(p.InstructingAgent) {
		Cd := camt052.ExternalClearingSystemIdentification1CodeFixed(p.InstructingAgent.PaymentSysCode)
		RltdAgts.InstgAgt = camt052.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt052.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt052.ClearingSystemMemberIdentification21{
					ClrSysId: camt052.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt052.RoutingNumberFRS1(p.InstructingAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(p.InstructedAgent) {
		Cd := camt052.ExternalClearingSystemIdentification1CodeFixed(p.InstructedAgent.PaymentSysCode)
		RltdAgts.InstdAgt = camt052.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: camt052.FinancialInstitutionIdentification181{
				ClrSysMmbId: camt052.ClearingSystemMemberIdentification21{
					ClrSysId: camt052.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: camt052.RoutingNumberFRS1(p.InstructedAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(RltdAgts) {
		TxDtls.RltdAgts = RltdAgts
	}
	var LclInstrm camt052.LocalInstrument2Choice1
	if p.LocalInstrumentChoice != "" {
		Prtry := camt052.LocalInstrumentFedwireFunds1(p.LocalInstrumentChoice)
		LclInstrm := camt052.LocalInstrument2Choice1{
			Prtry: &Prtry,
		}
		TxDtls.LclInstrm = &LclInstrm
	}
	if !isEmpty(LclInstrm) {
		TxDtls.LclInstrm = &LclInstrm
	}
	if !isEmpty(TxDtls) {
		result.TxDtls = TxDtls
	}
	return result
}
func EntryDetails91To(p camt052.EntryDetails91) model.EntryDetail {
	var result model.EntryDetail
	if !isEmpty(p.TxDtls) {
		if !isEmpty(p.TxDtls.Refs) {
			if !isEmpty(p.TxDtls.Refs.MsgId) {
				result.MessageId = string(p.TxDtls.Refs.MsgId)
			}
			if !isEmpty(p.TxDtls.Refs.InstrId) {
				result.InstructionId = string(*p.TxDtls.Refs.InstrId)
			}
			if !isEmpty(p.TxDtls.Refs.UETR) {
				result.UniqueTransactionReference = string(*p.TxDtls.Refs.UETR)
			}
			if !isEmpty(p.TxDtls.Refs.ClrSysRef) {
				result.ClearingSystemRef = string(*p.TxDtls.Refs.ClrSysRef)
			}
		}
		if !isEmpty(p.TxDtls.RltdAgts) {
			if !isEmpty(p.TxDtls.RltdAgts.InstgAgt) && !isEmpty(p.TxDtls.RltdAgts.InstgAgt.FinInstnId) && !isEmpty(p.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId) {
				if !isEmpty(p.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId) && !isEmpty(p.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
					result.InstructingAgent.PaymentSysCode = model.PaymentSystemType(*p.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
				}
				if !isEmpty(p.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId) {
					result.InstructingAgent.PaymentSysMemberId = string(p.TxDtls.RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId)
				}
			}
			if !isEmpty(p.TxDtls.RltdAgts.InstdAgt) && !isEmpty(p.TxDtls.RltdAgts.InstdAgt.FinInstnId) && !isEmpty(p.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId) {
				if !isEmpty(p.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId) && !isEmpty(p.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd) {
					result.InstructedAgent.PaymentSysCode = model.PaymentSystemType(*p.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
				}
				if !isEmpty(p.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId) {
					result.InstructedAgent.PaymentSysMemberId = string(p.TxDtls.RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId)
				}
			}
		}
		if !isEmpty(p.TxDtls.LclInstrm) && !isEmpty(p.TxDtls.LclInstrm.Prtry) {
			result.LocalInstrumentChoice = model.InstrumentPropCodeType(*p.TxDtls.LclInstrm.Prtry)
		}
	}
	return result
}
func ReportEntry101From(p model.Entry) (camt052.ReportEntry101, *model.ValidateError) {
	var result camt052.ReportEntry101
	if !isEmpty(p.Amount) {
		err := fedwire.Amount(p.Amount.Amount).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "Amount.Amount",
				Message:   err.Error(),
			}
		}
		err = camt052.ActiveOrHistoricCurrencyCode(p.Amount.Currency).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "Amount.Amount",
				Message:   err.Error(),
			}
		}
		result.Amt = camt052.ActiveOrHistoricCurrencyAndAmount{
			Value: camt052.ActiveOrHistoricCurrencyAndAmountSimpleType(p.Amount.Amount),
			Ccy:   camt052.ActiveOrHistoricCurrencyCode(p.Amount.Currency),
		}
	}
	if p.CreditDebitIndicator != "" {
		err := camt052.CreditDebitCode(p.CreditDebitIndicator).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "CreditDebitIndicator",
				Message:   err.Error(),
			}
		}
		result.CdtDbtInd = camt052.CreditDebitCode(p.CreditDebitIndicator)
	}
	if p.Status != "" {
		err := camt052.ExternalEntryStatus1Code(p.Status).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "Status",
				Message:   err.Error(),
			}
		}
		Cd := camt052.ExternalEntryStatus1Code(p.Status)
		result.Sts = camt052.EntryStatus1Choice1{
			Cd: &Cd,
		}
	}
	if p.BankTransactionCode != "" {
		err := camt052.BankTransactionCodeFedwireFunds1(p.BankTransactionCode).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "BankTransactionCode",
				Message:   err.Error(),
			}
		}
		result.BkTxCd = camt052.BankTransactionCodeStructure42{
			Prtry: camt052.ProprietaryBankTransactionCodeStructure12{
				Cd: camt052.BankTransactionCodeFedwireFunds1(p.BankTransactionCode),
			},
		}
	}
	if p.MessageNameId != "" {
		err := camt052.MessageNameIdentificationFRS1(p.MessageNameId).Validate()
		if err != nil {
			return camt052.ReportEntry101{}, &model.ValidateError{
				ParamName: "MessageNameId",
				Message:   err.Error(),
			}
		}
		result.AddtlInfInd = camt052.MessageIdentification21{
			MsgNmId: camt052.MessageNameIdentificationFRS1(p.MessageNameId),
		}
	}
	if !isEmpty(p.EntryDetails) {
		result.NtryDtls = EntryDetails91From(p.EntryDetails)
	}
	return result, nil
}
func ReportEntry101To(p camt052.ReportEntry101) model.Entry {
	var result model.Entry
	if !isEmpty(p.Amt) {
		result.Amount = model.CurrencyAndAmount{
			Amount:   float64(p.Amt.Value),
			Currency: string(p.Amt.Ccy),
		}
	}
	if !isEmpty(p.CdtDbtInd) {
		result.CreditDebitIndicator = model.CdtDbtInd(p.CdtDbtInd)
	}
	if !isEmpty(p.Sts) && !isEmpty(p.Sts.Cd) {
		result.Status = model.ReportStatus(*p.Sts.Cd)
	}
	if !isEmpty(p.BkTxCd) && !isEmpty(p.BkTxCd.Prtry) {
		result.BankTransactionCode = model.TransactionStatusCode(p.BkTxCd.Prtry.Cd)
	}
	if !isEmpty(p.AddtlInfInd) && !isEmpty(p.AddtlInfInd.MsgNmId) {
		result.MessageNameId = string(p.AddtlInfInd.MsgNmId)
	}
	if !isEmpty(p.NtryDtls) {
		result.EntryDetails = EntryDetails91To(p.NtryDtls)
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
