package Master

import (
	"reflect"
	"time"

	camt052 "github.com/moov-io/fedwire20022/gen/Master_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type AccountReportType string
type BalanceType string
type CreditLineType string
type TransactionCode string

const (
	ABMS        AccountReportType = "ABMS" //Solicited balance report sent by the Federal Reserve Banks in response to an account balance report request.
	FINAL       AccountReportType = "FINL" //Unsolicited balance report sent by the Federal Reserve Banks as part of the Account Balance Services end-of-day process.
	INTERIM     AccountReportType = "ITRM" //Unsolicited balance report sent by the Federal Reserve Banks when operating in contingency mode.
	OPENING     AccountReportType = "OPEN" //Unsolicited balance report sent by the Federal Reserve Banks when opening balance is loaded.
	PERIODIC    AccountReportType = "PRDC" //Unsolicited balance report sent by the Federal Reserve Banks throughout the day whenever the Federal Reserve Banks' accounting system updates Account Balance Services.
	PROVISIONAL AccountReportType = "PROV" //Unsolicited balance report sent by the Federal Reserve Banks when memo post is used.
)

const (
	AccountBalance                        BalanceType = "ABAL"
	AvailableBalanceFromAccountBalance    BalanceType = "AVAL"
	AvailableBalanceFromDaylightOverdraft BalanceType = "AVLD"
	DaylightOverdraftBalance              BalanceType = "DLOD"
	OpeningBalanceFinalBalanceLoaded      BalanceType = "OBFL"
	OpeningBalanceNotLoaded               BalanceType = "OBNL"
	OpeningBalancePriorDayBalanceLoaded   BalanceType = "OBPL"
)

const (
	CollateralAvailable                CreditLineType = "COLL"
	CollateralizedCapacity             CreditLineType = "CCAP"
	CollateralizedDaylightOverdrafts   CreditLineType = "CLOD"
	NetDebitCap                        CreditLineType = "NCAP"
	UncollateralizedDaylightOverdrafts CreditLineType = "ULOD"
)

const (
	AvailableAllOtherActivity        TransactionCode = "AVOT"
	FedNowFundsTransfers             TransactionCode = "FDNF"
	FedwireFundsTransfers            TransactionCode = "FDWF"
	FedwireSecuritiesTransfers       TransactionCode = "FDWS"
	MemoPostEntries                  TransactionCode = "MEMO"
	NationalSettlementServiceEntries TransactionCode = "NSSE"
	PrefundedACHCreditItems          TransactionCode = "FDAP"
	UnavailableAllOtherActivity      TransactionCode = "UVOT"
)

type CreditLine struct {
	//Indicates whether or not the credit line is included in the balance.
	Included bool
	//Type of the credit line provided when multiple credit lines may be provided.
	Type CreditLineType
	//Amount of money of the cash balance.
	Amount model.CurrencyAndAmount
	//Indicates the date (and time) of the balance.
	DateTime time.Time
}
type Balance struct {
	//Specifies the nature of a balance.
	BalanceTypeId BalanceType

	CdtLines []CreditLine
	//Amount of money of the cash balance.
	Amount model.CurrencyAndAmount
	//Indicates whether the balance is a credit or a debit balance.
	CreditDebitIndicator model.CdtDbtInd
	//Indicates the date (and time) of the balance.
	DateTime time.Time
}
type TotalsPerBankTransactionCode struct {
	TotalNetEntryAmount  float64
	CreditDebitIndicator model.CdtDbtInd
	CreditEntries        model.NumberAndSumOfTransactions
	DebitEntries         model.NumberAndSumOfTransactions
	BankTransactionCode  TransactionCode
	Date                 time.Time
}

func TotalsPerBankTransactionCode51From(p TotalsPerBankTransactionCode) (camt052.TotalsPerBankTransactionCode51, *model.ValidateError) {
	var result camt052.TotalsPerBankTransactionCode51
	var TtlNetNtry camt052.AmountAndDirection35
	if p.TotalNetEntryAmount != 0 {
		err := camt052.NonNegativeDecimalNumber(p.TotalNetEntryAmount).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "TotalNetEntryAmount",
				Message:   err.Error(),
			}
		}
		TtlNetNtry.Amt = camt052.NonNegativeDecimalNumber(p.TotalNetEntryAmount)
	}
	if p.CreditDebitIndicator != "" {
		err := camt052.CreditDebitCode(p.CreditDebitIndicator).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "CreditDebitIndicator",
				Message:   err.Error(),
			}
		}
		TtlNetNtry.CdtDbtInd = camt052.CreditDebitCode(p.CreditDebitIndicator)
	}
	if !isEmpty(TtlNetNtry) {
		result.TtlNetNtry = TtlNetNtry
	}
	if !isEmpty(p.CreditEntries) {
		err := camt052.Max15NumericText(p.CreditEntries.NumberOfEntries).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "CreditEntries.NumberOfEntries",
				Message:   err.Error(),
			}
		}
		err = camt052.DecimalNumber(p.CreditEntries.Sum).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "CreditEntries.Sum",
				Message:   err.Error(),
			}
		}
		NbOfNtries := camt052.Max15NumericText(p.CreditEntries.NumberOfEntries)
		result.CdtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: &NbOfNtries,
			Sum:        camt052.DecimalNumber(p.CreditEntries.Sum),
		}
	}
	if !isEmpty(p.DebitEntries) {
		err := camt052.Max15NumericText(p.DebitEntries.NumberOfEntries).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "DebitEntries.NumberOfEntries",
				Message:   err.Error(),
			}
		}
		err = camt052.DecimalNumber(p.DebitEntries.Sum).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "DebitEntries.Sum",
				Message:   err.Error(),
			}
		}
		NbOfNtries := camt052.Max15NumericText(p.DebitEntries.NumberOfEntries)
		result.DbtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: &NbOfNtries,
			Sum:        camt052.DecimalNumber(p.DebitEntries.Sum),
		}
	}
	if p.BankTransactionCode != "" {
		err := camt052.TransactionsSummaryTypeFRS1(p.BankTransactionCode).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "BankTransactionCode",
				Message:   err.Error(),
			}
		}
		result.BkTxCd = camt052.BankTransactionCodeStructure41{
			Prtry: camt052.ProprietaryBankTransactionCodeStructure11{
				Cd: camt052.TransactionsSummaryTypeFRS1(p.BankTransactionCode),
			},
		}
	}
	if !isEmpty(p.Date) {
		err := fedwire.ISODateTime(p.Date).Validate()
		if err != nil {
			return camt052.TotalsPerBankTransactionCode51{}, &model.ValidateError{
				ParamName: "Date",
				Message:   err.Error(),
			}
		}
		DtTm := fedwire.ISODateTime(p.Date)
		result.Dt = camt052.DateAndDateTime2Choice1{
			DtTm: &DtTm,
		}
	}
	return result, nil
}
func TotalsPerBankTransactionCode51To(p camt052.TotalsPerBankTransactionCode51) TotalsPerBankTransactionCode {
	var result TotalsPerBankTransactionCode
	if !isEmpty(p.TtlNetNtry) {
		if !isEmpty(p.TtlNetNtry.Amt) {
			result.TotalNetEntryAmount = float64(p.TtlNetNtry.Amt)
		}
		if !isEmpty(p.TtlNetNtry.CdtDbtInd) {
			result.CreditDebitIndicator = model.CdtDbtInd(p.TtlNetNtry.CdtDbtInd)
		}
	}
	if !isEmpty(p.CdtNtries) {
		if !isEmpty(p.CdtNtries.NbOfNtries) {
			result.CreditEntries.NumberOfEntries = string(*p.CdtNtries.NbOfNtries)
		}
		if !isEmpty(p.CdtNtries.Sum) {
			result.CreditEntries.Sum = float64(p.CdtNtries.Sum)
		}
	}
	if !isEmpty(p.DbtNtries) {
		if !isEmpty(p.DbtNtries.NbOfNtries) {
			result.DebitEntries.NumberOfEntries = string(*p.DbtNtries.NbOfNtries)
		}
		if !isEmpty(p.DbtNtries.Sum) {
			result.DebitEntries.Sum = float64(p.DbtNtries.Sum)
		}
	}
	if !isEmpty(p.BkTxCd) {
		if !isEmpty(p.BkTxCd.Prtry) {
			result.BankTransactionCode = TransactionCode(p.BkTxCd.Prtry.Cd)
		}
	}
	if !isEmpty(p.Dt) {
		if !isEmpty(p.Dt.DtTm) {
			result.Date = time.Time(*p.Dt.DtTm)
		}
	}
	return result
}
func CreditLine31From(p CreditLine) (camt052.CreditLine31, *model.ValidateError) {
	var result camt052.CreditLine31
	if !isEmpty(p.Included) {
		err := camt052.TrueFalseIndicator(p.Included).Validate()
		if err != nil {
			return camt052.CreditLine31{}, &model.ValidateError{
				ParamName: "Included",
				Message:   err.Error(),
			}
		}
		result.Incl = camt052.TrueFalseIndicator(p.Included)
	}
	if p.Type != "" {
		err := camt052.CreditLineTypeFRS1(p.Type).Validate()
		if err != nil {
			return camt052.CreditLine31{}, &model.ValidateError{
				ParamName: "Type",
				Message:   err.Error(),
			}
		}
		Prtry := camt052.CreditLineTypeFRS1(p.Type)
		result.Tp = camt052.CreditLineType1Choice1{
			Prtry: &Prtry,
		}
	}
	if !isEmpty(p.Amount) {
		err := fedwire.Amount(p.Amount.Amount).Validate()
		if err != nil {
			return camt052.CreditLine31{}, &model.ValidateError{
				ParamName: "Amount.Amount",
				Message:   err.Error(),
			}
		}
		err = camt052.ActiveOrHistoricCurrencyCode(p.Amount.Currency).Validate()
		if err != nil {
			return camt052.CreditLine31{}, &model.ValidateError{
				ParamName: "Amount.Currency",
				Message:   err.Error(),
			}
		}
		result.Amt = camt052.ActiveOrHistoricCurrencyAndAmount{
			Value: camt052.ActiveOrHistoricCurrencyAndAmountSimpleType(p.Amount.Amount),
			Ccy:   camt052.ActiveOrHistoricCurrencyCode(p.Amount.Currency),
		}
	}
	if !isEmpty(p.DateTime) {
		err := fedwire.ISODateTime(p.DateTime).Validate()
		if err != nil {
			return camt052.CreditLine31{}, &model.ValidateError{
				ParamName: "DateTime",
				Message:   err.Error(),
			}
		}
		DtTm := fedwire.ISODateTime(p.DateTime)
		result.Dt = camt052.DateAndDateTime2Choice1{
			DtTm: &DtTm,
		}
	}
	return result, nil
}
func CreditLine31To(p camt052.CreditLine31) CreditLine {
	var result CreditLine
	if !isEmpty(p.Incl) {
		result.Included = bool(p.Incl)
	}
	if !isEmpty(p.Tp) {
		if !isEmpty(p.Tp.Prtry) {
			result.Type = CreditLineType(*p.Tp.Prtry)
		}
	}
	if !isEmpty(p.Amt) {
		result.Amount = model.CurrencyAndAmount{
			Amount:   float64(p.Amt.Value),
			Currency: string(p.Amt.Ccy),
		}
	}
	if !isEmpty(p.Dt) {
		if !isEmpty(p.Dt.DtTm) {
			result.DateTime = time.Time(*p.Dt.DtTm)
		}
	}
	return result
}
func CashBalance81From(p Balance) (camt052.CashBalance81, *model.ValidateError) {
	var result camt052.CashBalance81
	if p.BalanceTypeId != "" {
		err := camt052.BalanceTypeFRS1(p.BalanceTypeId).Validate()
		if err != nil {
			return camt052.CashBalance81{}, &model.ValidateError{
				ParamName: "BalanceTypeId",
				Message:   err.Error(),
			}
		}
		Prtry := camt052.BalanceTypeFRS1(p.BalanceTypeId)
		result.Tp = camt052.BalanceType131{
			CdOrPrtry: camt052.BalanceType10Choice1{
				Prtry: &Prtry,
			},
		}
	}
	var CdtLine []*camt052.CreditLine31
	if !isEmpty(p.CdtLines) {
		for _, item := range p.CdtLines {
			line, vErr := CreditLine31From(item)
			if vErr != nil {
				vErr.InsertPath("CdtLine")
				return camt052.CashBalance81{}, vErr
			}
			CdtLine = append(CdtLine, &line)
		}
	}
	if !isEmpty(CdtLine) {
		result.CdtLine = CdtLine
	}
	if !isEmpty(p.Amount) {
		err := fedwire.Amount(p.Amount.Amount).Validate()
		if err != nil {
			return camt052.CashBalance81{}, &model.ValidateError{
				ParamName: "Amount.Amount",
				Message:   err.Error(),
			}
		}
		err = camt052.ActiveOrHistoricCurrencyCode(p.Amount.Currency).Validate()
		if err != nil {
			return camt052.CashBalance81{}, &model.ValidateError{
				ParamName: "Amount.Currency",
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
			return camt052.CashBalance81{}, &model.ValidateError{
				ParamName: "CreditDebitIndicator",
				Message:   err.Error(),
			}
		}
		result.CdtDbtInd = camt052.CreditDebitCode(p.CreditDebitIndicator)
	}
	if !isEmpty(p.DateTime) {
		err := fedwire.ISODateTime(p.DateTime).Validate()
		if err != nil {
			return camt052.CashBalance81{}, &model.ValidateError{
				ParamName: "DateTime",
				Message:   err.Error(),
			}
		}
		DtTm := fedwire.ISODateTime(p.DateTime)
		result.Dt = camt052.DateAndDateTime2Choice1{
			DtTm: &DtTm,
		}
	}
	return result, nil
}
func CashBalance81To(p camt052.CashBalance81) Balance {
	var result Balance
	if !isEmpty(p.Tp) {
		if !isEmpty(p.Tp.CdOrPrtry) {
			if !isEmpty(p.Tp.CdOrPrtry.Prtry) {
				result.BalanceTypeId = BalanceType(*p.Tp.CdOrPrtry.Prtry)
			}
		}
	}
	if !isEmpty(p.CdtLine) {
		for _, item := range p.CdtLine {
			line := CreditLine31To(*item)
			result.CdtLines = append(result.CdtLines, line)
		}
	}
	if !isEmpty(p.Amt) {
		result.Amount = model.CurrencyAndAmount{
			Amount:   float64(p.Amt.Value),
			Currency: string(p.Amt.Ccy),
		}
	}
	if p.CdtDbtInd != "" {
		result.CreditDebitIndicator = model.CdtDbtInd(p.CdtDbtInd)
	}
	if !isEmpty(p.Dt) {
		if !isEmpty(p.Dt.DtTm) {
			result.DateTime = time.Time(*p.Dt.DtTm)
		}
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
