package Master

import "github.com/moov-io/wire20022/pkg/models"

type CreditLineHelper struct {
	Included models.ElementHelper
	Type     models.ElementHelper
	Amount   models.CurrencyAndAmountHelper
	DateTime models.ElementHelper
}

func BuildCreditLineHelper() CreditLineHelper {
	return CreditLineHelper{
		Included: models.ElementHelper{
			Title:         "Included",
			Rules:         "",
			Type:          `Boolean (based on string)`,
			Documentation: `Indicates whether or not the credit line is included in the balance.`,
		},
		Type: models.ElementHelper{
			Title:         "Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Type of the credit line provided when multiple credit lines may be provided.`,
		},
		Amount: models.BuildCurrencyAndAmountHelper(),
		DateTime: models.ElementHelper{
			Title:         "Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Indicates the date (and time) of the balance.`,
		},
	}
}

type BalanceHelper struct {
	BalanceTypeId        models.ElementHelper
	CdtLines             CreditLineHelper
	Amount               models.CurrencyAndAmountHelper
	CreditDebitIndicator models.ElementHelper
	DateTime             models.ElementHelper
}

func BuildBalanceHelper() BalanceHelper {
	return BalanceHelper{
		BalanceTypeId: models.ElementHelper{
			Title:         "Balance Type Id",
			Rules:         "",
			Type:          `BalanceType(AccountBalance, AvailableBalanceFromAccountBalance ...)`,
			Documentation: `Specifies the nature of a balance.`,
		},
		CdtLines: BuildCreditLineHelper(),
		Amount:   models.BuildCurrencyAndAmountHelper(),
		CreditDebitIndicator: models.ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `CdtDbtInd(Credit, Debit)`,
			Documentation: `Indicates whether the balance is a credit or a debit balance.`,
		},
		DateTime: models.ElementHelper{
			Title:         "Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Indicates the date (and time) of the balance.`,
		},
	}
}

type TotalsPerBankTransactionCodeHelper struct {
	TotalNetEntryAmount  models.ElementHelper
	CreditDebitIndicator models.ElementHelper
	CreditEntries        models.NumberAndSumOfTransactionsHelper
	DebitEntries         models.NumberAndSumOfTransactionsHelper
	BankTransactionCode  models.ElementHelper
	Date                 models.ElementHelper
}

func BuildTotalsPerBankTransactionCodeHelper() TotalsPerBankTransactionCodeHelper {
	return TotalsPerBankTransactionCodeHelper{
		TotalNetEntryAmount: models.ElementHelper{
			Title:         "Total Net Entry Amount",
			Rules:         "",
			Type:          `CurrencyAndAmount (based on string)`,
			Documentation: `Total net entry amount of the transactions included in the report.`,
		},
		CreditDebitIndicator: models.ElementHelper{
			Title:         "Credit Debit Indicator",
			Rules:         "",
			Type:          `ExternalAccountIdentification1Code (based on string)`,
			Documentation: `Indicates whether the balance is a credit or a debit balance.`,
		},
		CreditEntries: models.BuildNumberAndSumOfTransactionsHelper(),
		DebitEntries:  models.BuildNumberAndSumOfTransactionsHelper(),
		BankTransactionCode: models.ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the bank transaction code to which the entry refers.`,
		},
		Date: models.ElementHelper{
			Title:         "Date",
			Rules:         "",
			Type:          `ISODate (based on string)`,
			Documentation: `Date at which the transaction was executed.`,
		},
	}
}

type MessageHelper struct {
	//Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId models.ElementHelper
	//Date and time at which the message was created.
	CreationDateTime models.ElementHelper
	//Provides details on the page number of the message.
	MessagePagination models.MessagePagenationHelper
	//Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.
	OriginalBusinessMsgId models.ElementHelper
	//Specifies the query message name identifier to which the message refers.
	OriginalBusinessMsgNameId models.ElementHelper
	//Date and time at which the message was created.
	OriginalBusinessMsgCreateTime models.ElementHelper
	//Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	ReportTypeId models.ElementHelper
	//Date and time at which the report was created.
	ReportCreatedDate models.ElementHelper
	//Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId models.ElementHelper
	AccountType    models.ElementHelper
	//Identifies the parent account of the account for which the report has been issued.
	RelatedAccountOtherId models.ElementHelper
	//Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balances BalanceHelper
	//Provides summary information on entries.
	TransactionsSummary TotalsPerBankTransactionCodeHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: models.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `CAMTReportType(AccountBalanceReport, ActivityReport ...)`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.`,
		},
		CreationDateTime: models.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the message was created.`,
		},
		MessagePagination: models.BuildMessagePagenationHelper(),
		OriginalBusinessMsgId: models.ElementHelper{
			Title:         "Original Business Msg Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.`,
		},
		OriginalBusinessMsgNameId: models.ElementHelper{
			Title:         "Original Business Msg Name Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the query message name identifier to which the message refers.`,
		},
		OriginalBusinessMsgCreateTime: models.ElementHelper{
			Title:         "Original Business Msg Create Time",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the message was created.`,
		},
		ReportTypeId: models.ElementHelper{
			Title:         "Report Type Id",
			Rules:         "",
			Type:          `AccountReportType(ABMS, FINAL, INTERIM ...)`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
		AccountOtherId: models.ElementHelper{
			Title:         "Account Other Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: "Unambiguous identification of the account to which credit and debit entries are made.",
		},
		AccountType: models.ElementHelper{
			Title:         "Account Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: "Type of the account.",
		},
		RelatedAccountOtherId: models.ElementHelper{
			Title:         "Related Account Other Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Identifies the parent account of the account for which the report has been issued.`,
		},
		ReportCreatedDate: models.ElementHelper{
			Title:         "Report Created Date",
			Rules:         "",
			Type:          `ISODateTime (based on string)`,
			Documentation: `Date and time at which the report was created.`,
		},
		Balances:            BuildBalanceHelper(),
		TransactionsSummary: BuildTotalsPerBankTransactionCodeHelper(),
	}
}
