package ActivityReport

import (
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type TotalsPerBankTransactionCodeHelper struct {
	NumberOfEntries     Archive.ElementHelper
	BankTransactionCode Archive.ElementHelper
}

func BuildTotalsPerBankTransactionCodeHelper() TotalsPerBankTransactionCodeHelper {
	return TotalsPerBankTransactionCodeHelper{
		NumberOfEntries: Archive.ElementHelper{
			Title:         "Number of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) minLength: 1 maxLength: 15`,
			Documentation: `Number of individual entries for the bank transaction code.`,
		},
		BankTransactionCode: Archive.ElementHelper{
			Title:         "Bank Transaction Code",
			Rules:         "",
			Type:          `TransactionStatusCode(MessagesInProcess, MessagesIntercepted ...)`,
			Documentation: `Bank transaction code in a proprietary form, as defined by the issuer.`,
		},
	}
}

type MessageHelper struct {
	MessageId                          Archive.ElementHelper
	CreatedDateTime                    Archive.ElementHelper
	Pagenation                         Archive.MessagePagenationHelper
	ReportId                           Archive.ElementHelper
	ReportCreateDateTime               Archive.ElementHelper
	AccountOtherId                     Archive.ElementHelper
	TotalEntries                       Archive.ElementHelper
	TotalCreditEntries                 Archive.NumberAndSumOfTransactionsHelper
	TotalDebitEntries                  Archive.NumberAndSumOfTransactionsHelper
	TotalEntriesPerBankTransactionCode TotalsPerBankTransactionCodeHelper
	EntryDetails                       Archive.EntryHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: Archive.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message. Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.`,
		},
		CreatedDateTime: Archive.ElementHelper{
			Title:         "Message Identification",
			Rules:         "This is the calendar date and time in New York City (Eastern Time) when the message is created by the Fedwire Funds Service application. Time is in 24-hour clock format and includes the offset against the Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		Pagenation: Archive.BuildMessagePagenationHelper(),
		ReportId: Archive.ElementHelper{
			Title:         "Report Type Id",
			Rules:         "",
			Type:          `ReportType(EveryDay, Intraday)`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
		ReportCreateDateTime: Archive.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "This is the Fedwire Funds Service funds-transfer business day. Note: Time will be defaulted to 00:00:00 in New York City (Eastern Time) with the offset against the Coordinated Universal Time (UTC).",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the report was created.`,
		},
		AccountOtherId: Archive.ElementHelper{
			Title:         "Account Other Id",
			Rules:         "his is the routing number to which the activity report relates.",
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification assigned by an institution.`,
		},
		TotalEntries: Archive.ElementHelper{
			Title:         "Number Of Entries",
			Rules:         "",
			Type:          `Max15NumericText (based on string) pattern: [0-9]{1,15}`,
			Documentation: `Number of individual entries included in the report.`,
		},
		TotalCreditEntries:                 Archive.BuildNumberAndSumOfTransactionsHelper(),
		TotalDebitEntries:                  Archive.BuildNumberAndSumOfTransactionsHelper(),
		TotalEntriesPerBankTransactionCode: BuildTotalsPerBankTransactionCodeHelper(),
		EntryDetails:                       Archive.BuildEntryHelper(),
	}

}
