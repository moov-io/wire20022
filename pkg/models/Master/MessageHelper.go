package Master

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type CreditLineHelper struct {
	Included model.ElementHelper
	Type     model.ElementHelper
	Amount   model.CurrencyAndAmountHelper
	DateTime model.ElementHelper
}

func BuildCreditLineHelper() CreditLineHelper {
	return CreditLineHelper{
		Included: model.ElementHelper{
			Title:         "Included",
			Rules:         "",
			Type:          `Boolean (based on string)`,
			Documentation: `Indicates whether or not the credit line is included in the balance.`,
		},
		Type: model.ElementHelper{
			Title:         "Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Type of the credit line provided when multiple credit lines may be provided.`,
		},
		Amount: model.BuildCurrencyAndAmountHelper(),
		DateTime: model.ElementHelper{
			Title:         "Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Indicates the date (and time) of the balance.`,
		},
	}
}

type BalanceHelper struct {
	BalanceTypeId        model.ElementHelper
	CdtLines             CreditLineHelper
	Amount               model.CurrencyAndAmountHelper
	CreditDebitIndicator model.ElementHelper
	DateTime             model.ElementHelper
}

func BuildBalanceHelper() BalanceHelper {
	return BalanceHelper{
		BalanceTypeId: model.ElementHelper{
			Title:         "Balance Type Id",
			Rules:         "",
			Type:          `BalanceType(AccountBalance, AvailableBalanceFromAccountBalance ...)`,
			Documentation: `Specifies the nature of a balance.`,
		},
		CdtLines: BuildCreditLineHelper(),
		Amount:   model.BuildCurrencyAndAmountHelper(),
		CreditDebitIndicator: model.ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `CdtDbtInd(Credit, Debit)`,
			Documentation: `Indicates whether the balance is a credit or a debit balance.`,
		},
		DateTime: model.ElementHelper{
			Title:         "Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Indicates the date (and time) of the balance.`,
		},
	}
}

type TotalsPerBankTransactionCodeHelper struct {
	TotalNetEntryAmount  model.ElementHelper
	CreditDebitIndicator model.ElementHelper
	CreditEntries        model.NumberAndSumOfTransactionsHelper
	DebitEntries         model.NumberAndSumOfTransactionsHelper
	BankTransactionCode  model.ElementHelper
	Date                 model.ElementHelper
}

func BuildTotalsPerBankTransactionCodeHelper() TotalsPerBankTransactionCodeHelper {
	return TotalsPerBankTransactionCodeHelper{
		TotalNetEntryAmount: model.ElementHelper{
			Title:         "Total Net Entry Amount",
			Rules:         "",
			Type:          `CurrencyAndAmount (based on string)`,
			Documentation: `Total net entry amount of the transactions included in the report.`,
		},
		CreditDebitIndicator: model.ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `ExternalAccountIdentification1Code (based on string)`,
			Documentation: `Indicates whether the balance is a credit or a debit balance.`,
		},
		CreditEntries: model.BuildNumberAndSumOfTransactionsHelper(),
		DebitEntries:  model.BuildNumberAndSumOfTransactionsHelper(),
		BankTransactionCode: model.ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the bank transaction code to which the entry refers.`,
		},
		Date: model.ElementHelper{
			Title:         "Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date at which the transaction was executed.`,
		},
	}
}

type MessageHelper struct {
	//Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId model.ElementHelper
	//Date and time at which the message was created.
	CreationDateTime model.ElementHelper
	//Provides details on the page number of the message.
	MessagePagination model.MessagePagenationHelper
	//Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.
	OriginalBusinessMsgId model.ElementHelper
	//Specifies the query message name identifier to which the message refers.
	OriginalBusinessMsgNameId model.ElementHelper
	//Date and time at which the message was created.
	OriginalBusinessMsgCreateTime model.ElementHelper
	//Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	ReportTypeId model.ElementHelper
	//Date and time at which the report was created.
	ReportCreatedDate model.ElementHelper
	//Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId model.ElementHelper
	AccountType    model.ElementHelper
	//Identifies the parent account of the account for which the report has been issued.
	RelatedAccountOtherId model.ElementHelper
	//Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balances BalanceHelper
	//Provides summary information on entries.
	TransactionsSummary TotalsPerBankTransactionCodeHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `CAMTReportType(AccountBalanceReport, ActivityReport ...)`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.`,
		},
		CreationDateTime: model.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the message was created.`,
		},
		MessagePagination: model.BuildMessagePagenationHelper(),
		OriginalBusinessMsgId: model.ElementHelper{
			Title:         "Original Business Msg Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.`,
		},
		OriginalBusinessMsgNameId: model.ElementHelper{
			Title:         "Original Business Msg Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the query message name identifier to which the message refers.`,
		},
		OriginalBusinessMsgCreateTime: model.ElementHelper{
			Title:         "Original Business Msg Create Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the message was created.`,
		},
		ReportTypeId: model.ElementHelper{
			Title:         "Report Type Id",
			Rules:         "",
			Type:          `AccountReportType(ABMS, FINAL, INTERIM ...)`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
		AccountOtherId: model.ElementHelper{
			Title: "Account Other Id",
			Rules: "",
			Type: `Max35Text
			(based on string) minLength: 1 maxLength: 35`,
			Documentation: "Unambiguous identification of the account to which credit and debit entries are made.",
		},
		AccountType: model.ElementHelper{
			Title: "Account Type",
			Rules: "",
			Type: `Max35Text
			(based on string) minLength: 1 maxLength: 35`,
			Documentation: "Type of the account.",
		},
		RelatedAccountOtherId: model.ElementHelper{
			Title:         "Related Account Other Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identifies the parent account of the account for which the report has been issued.`,
		},
		ReportCreatedDate: model.ElementHelper{
			Title:         "Report Created Date",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the report was created.`,
		},
		Balances:            BuildBalanceHelper(),
		TransactionsSummary: BuildTotalsPerBankTransactionCodeHelper(),
	}
}
