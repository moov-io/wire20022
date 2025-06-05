package Master

import (
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type CreditLineHelper struct {
	Included Archive.ElementHelper
	Type     Archive.ElementHelper
	Amount   Archive.CurrencyAndAmountHelper
	DateTime Archive.ElementHelper
}

func BuildCreditLineHelper() CreditLineHelper {
	return CreditLineHelper{
		Included: Archive.ElementHelper{
			Title:         "Included",
			Rules:         "",
			Type:          `Boolean (based on string)`,
			Documentation: `Indicates whether or not the credit line is included in the balance.`,
		},
		Type: Archive.ElementHelper{
			Title:         "Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Type of the credit line provided when multiple credit lines may be provided.`,
		},
		Amount: Archive.BuildCurrencyAndAmountHelper(),
		DateTime: Archive.ElementHelper{
			Title:         "Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Indicates the date (and time) of the balance.`,
		},
	}
}

type BalanceHelper struct {
	BalanceTypeId        Archive.ElementHelper
	CdtLines             CreditLineHelper
	Amount               Archive.CurrencyAndAmountHelper
	CreditDebitIndicator Archive.ElementHelper
	DateTime             Archive.ElementHelper
}

func BuildBalanceHelper() BalanceHelper {
	return BalanceHelper{
		BalanceTypeId: Archive.ElementHelper{
			Title:         "Balance Type Id",
			Rules:         "",
			Type:          `BalanceType(AccountBalance, AvailableBalanceFromAccountBalance ...)`,
			Documentation: `Specifies the nature of a balance.`,
		},
		CdtLines: BuildCreditLineHelper(),
		Amount:   Archive.BuildCurrencyAndAmountHelper(),
		CreditDebitIndicator: Archive.ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `CdtDbtInd(Credit, Debit)`,
			Documentation: `Indicates whether the balance is a credit or a debit balance.`,
		},
		DateTime: Archive.ElementHelper{
			Title:         "Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Indicates the date (and time) of the balance.`,
		},
	}
}

type TotalsPerBankTransactionCodeHelper struct {
	TotalNetEntryAmount  Archive.ElementHelper
	CreditDebitIndicator Archive.ElementHelper
	CreditEntries        Archive.NumberAndSumOfTransactionsHelper
	DebitEntries         Archive.NumberAndSumOfTransactionsHelper
	BankTransactionCode  Archive.ElementHelper
	Date                 Archive.ElementHelper
}

func BuildTotalsPerBankTransactionCodeHelper() TotalsPerBankTransactionCodeHelper {
	return TotalsPerBankTransactionCodeHelper{
		TotalNetEntryAmount: Archive.ElementHelper{
			Title:         "Total Net Entry Amount",
			Rules:         "",
			Type:          `CurrencyAndAmount (based on string)`,
			Documentation: `Total net entry amount of the transactions included in the report.`,
		},
		CreditDebitIndicator: Archive.ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `ExternalAccountIdentification1Code (based on string)`,
			Documentation: `Indicates whether the balance is a credit or a debit balance.`,
		},
		CreditEntries: Archive.BuildNumberAndSumOfTransactionsHelper(),
		DebitEntries:  Archive.BuildNumberAndSumOfTransactionsHelper(),
		BankTransactionCode: Archive.ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the bank transaction code to which the entry refers.`,
		},
		Date: Archive.ElementHelper{
			Title:         "Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date at which the transaction was executed.`,
		},
	}
}

type MessageHelper struct {
	//Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId Archive.ElementHelper
	//Date and time at which the message was created.
	CreationDateTime Archive.ElementHelper
	//Provides details on the page number of the message.
	MessagePagination Archive.MessagePagenationHelper
	//Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.
	OriginalBusinessMsgId Archive.ElementHelper
	//Specifies the query message name identifier to which the message refers.
	OriginalBusinessMsgNameId Archive.ElementHelper
	//Date and time at which the message was created.
	OriginalBusinessMsgCreateTime Archive.ElementHelper
	//Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	ReportTypeId Archive.ElementHelper
	//Date and time at which the report was created.
	ReportCreatedDate Archive.ElementHelper
	//Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId Archive.ElementHelper
	AccountType    Archive.ElementHelper
	//Identifies the parent account of the account for which the report has been issued.
	RelatedAccountOtherId Archive.ElementHelper
	//Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balances BalanceHelper
	//Provides summary information on entries.
	TransactionsSummary TotalsPerBankTransactionCodeHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: Archive.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `CAMTReportType(AccountBalanceReport, ActivityReport ...)`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.`,
		},
		CreationDateTime: Archive.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the message was created.`,
		},
		MessagePagination: Archive.BuildMessagePagenationHelper(),
		OriginalBusinessMsgId: Archive.ElementHelper{
			Title:         "Original Business Msg Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.`,
		},
		OriginalBusinessMsgNameId: Archive.ElementHelper{
			Title:         "Original Business Msg Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the query message name identifier to which the message refers.`,
		},
		OriginalBusinessMsgCreateTime: Archive.ElementHelper{
			Title:         "Original Business Msg Create Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the message was created.`,
		},
		ReportTypeId: Archive.ElementHelper{
			Title:         "Report Type Id",
			Rules:         "",
			Type:          `AccountReportType(ABMS, FINAL, INTERIM ...)`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
		AccountOtherId: Archive.ElementHelper{
			Title:         "Account Other Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: "Unambiguous identification of the account to which credit and debit entries are made.",
		},
		AccountType: Archive.ElementHelper{
			Title:         "Account Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: "Type of the account.",
		},
		RelatedAccountOtherId: Archive.ElementHelper{
			Title:         "Related Account Other Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identifies the parent account of the account for which the report has been issued.`,
		},
		ReportCreatedDate: Archive.ElementHelper{
			Title:         "Report Created Date",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the report was created.`,
		},
		Balances:            BuildBalanceHelper(),
		TransactionsSummary: BuildTotalsPerBankTransactionCodeHelper(),
	}
}
