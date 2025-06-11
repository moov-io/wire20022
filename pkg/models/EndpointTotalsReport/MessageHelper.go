package EndpointTotalsReport

import "github.com/moov-io/wire20022/pkg/models"

type MessageHelper struct {
	MessageId                          models.ElementHelper
	CreatedDateTime                    models.ElementHelper
	MessagePagination                  models.MessagePagenationHelper
	ReportId                           models.ElementHelper
	ReportCreateDateTime               models.ElementHelper
	AccountOtherId                     models.ElementHelper
	TotalCreditEntries                 models.NumberAndSumOfTransactionsHelper
	TotalDebitEntries                  models.NumberAndSumOfTransactionsHelper
	TotalEntriesPerBankTransactionCode models.NumberAndStatusOfTransactionsHelper
	AdditionalReportInfo               models.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: models.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message. Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.`,
		},
		CreatedDateTime: models.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		MessagePagination: models.BuildMessagePagenationHelper(),
		ReportId: models.ElementHelper{
			Title:         "Report Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
		ReportCreateDateTime: models.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the report was created.`,
		},
		AccountOtherId: models.ElementHelper{
			Title:         "Account Other Id",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Unique identification of an account, as assigned by the account servicer, using an identification scheme.`,
		},
		TotalCreditEntries:                 models.BuildNumberAndSumOfTransactionsHelper(),
		TotalDebitEntries:                  models.BuildNumberAndSumOfTransactionsHelper(),
		TotalEntriesPerBankTransactionCode: models.BuildNumberAndStatusOfTransactionsHelper(),
		AdditionalReportInfo: models.ElementHelper{
			Title:         "Additional Report Information",
			Rules:         "",
			Type:          `Max500Text (based on string) minLength: 1 maxLength: 500`,
			Documentation: `Further details of the account report.`,
		},
	}
}
