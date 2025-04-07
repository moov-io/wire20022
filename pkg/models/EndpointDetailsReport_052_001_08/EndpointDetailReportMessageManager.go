package EndpointDetailsReport_052_001_08

import (
	"reflect"

	EndpointDetailsReport "github.com/moov-io/fedwire20022/gen/EndpointDetailsReport_camt_052_001_08"
	model "github.com/moov-io/wire20022/pkg/models"
)

func EntryDetails91From(p model.EntryDetail) EndpointDetailsReport.EntryDetails91 {
	var result EndpointDetailsReport.EntryDetails91
	var TxDtls EndpointDetailsReport.EntryTransaction101
	var Refs EndpointDetailsReport.TransactionReferences61
	if p.MessageId != "" {
		Refs.MsgId = EndpointDetailsReport.IMADFedwireFunds1(p.MessageId)
	}
	if p.InstructionId != "" {
		InstrId := EndpointDetailsReport.Max35Text(p.InstructionId)
		Refs.InstrId = &InstrId
	}
	if p.UniqueTransactionReference != "" {
		UETR := EndpointDetailsReport.UUIDv4Identifier(p.UniqueTransactionReference)
		Refs.UETR = &UETR
	}
	if !isEmpty(Refs) {
		TxDtls.Refs = Refs
	}
	var RltdAgts EndpointDetailsReport.TransactionAgents51
	if !isEmpty(p.InstructingAgent) {
		Cd := EndpointDetailsReport.ExternalClearingSystemIdentification1CodeFixed(p.InstructingAgent.PaymentSysCode)
		RltdAgts.InstgAgt = EndpointDetailsReport.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: EndpointDetailsReport.FinancialInstitutionIdentification181{
				ClrSysMmbId: EndpointDetailsReport.ClearingSystemMemberIdentification21{
					ClrSysId: EndpointDetailsReport.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: EndpointDetailsReport.RoutingNumberFRS1(p.InstructingAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(p.InstructedAgent) {
		Cd := EndpointDetailsReport.ExternalClearingSystemIdentification1CodeFixed(p.InstructedAgent.PaymentSysCode)
		RltdAgts.InstdAgt = EndpointDetailsReport.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: EndpointDetailsReport.FinancialInstitutionIdentification181{
				ClrSysMmbId: EndpointDetailsReport.ClearingSystemMemberIdentification21{
					ClrSysId: EndpointDetailsReport.ClearingSystemIdentification2Choice1{
						Cd: &Cd,
					},
					MmbId: EndpointDetailsReport.RoutingNumberFRS1(p.InstructedAgent.PaymentSysMemberId),
				},
			},
		}
	}
	if !isEmpty(RltdAgts) {
		TxDtls.RltdAgts = RltdAgts
	}
	var LclInstrm EndpointDetailsReport.LocalInstrument2Choice1
	if p.LocalInstrumentChoice != "" {
		Prtry := EndpointDetailsReport.LocalInstrumentFedwireFunds1(p.LocalInstrumentChoice)
		LclInstrm := EndpointDetailsReport.LocalInstrument2Choice1{
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
func ReportEntry101From(p model.Entry) EndpointDetailsReport.ReportEntry101 {
	var result EndpointDetailsReport.ReportEntry101
	if !isEmpty(p.Amount) {
		result.Amt = EndpointDetailsReport.ActiveOrHistoricCurrencyAndAmount{
			Value: EndpointDetailsReport.ActiveOrHistoricCurrencyAndAmountSimpleType(p.Amount.Amount),
			Ccy:   EndpointDetailsReport.ActiveOrHistoricCurrencyCode(p.Amount.Currency),
		}
	}
	if p.CreditDebitIndicator != "" {
		result.CdtDbtInd = EndpointDetailsReport.CreditDebitCode(p.CreditDebitIndicator)
	}
	if p.Status != "" {
		Cd := EndpointDetailsReport.ExternalEntryStatus1Code(p.Status)
		result.Sts = EndpointDetailsReport.EntryStatus1Choice1{
			Cd: &Cd,
		}
	}
	if p.BankTransactionCode != "" {
		result.BkTxCd = EndpointDetailsReport.BankTransactionCodeStructure42{
			Prtry: EndpointDetailsReport.ProprietaryBankTransactionCodeStructure12{
				Cd: EndpointDetailsReport.BankTransactionCodeFedwireFunds1(p.BankTransactionCode),
			},
		}
	}
	if p.MessageNameId != "" {
		result.AddtlInfInd = EndpointDetailsReport.MessageIdentification21{
			MsgNmId: EndpointDetailsReport.MessageNameIdentificationFRS1(p.MessageNameId),
		}
	}
	if !isEmpty(p.EntryDetails) {
		result.NtryDtls = EntryDetails91From(p.EntryDetails)
	}
	return result
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
