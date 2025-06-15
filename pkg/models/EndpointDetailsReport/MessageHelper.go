package EndpointDetailsReport

import "github.com/wadearnold/wire20022/pkg/models"

type MessageHelper struct {
	MessageId                      models.ElementHelper
	CreationDateTime               models.ElementHelper
	MessagePagination              models.MessagePagenationHelper
	BussinessQueryMsgId            models.ElementHelper
	BussinessQueryMsgNameId        models.ElementHelper
	BussinessQueryCreateDatetime   models.ElementHelper
	ReportId                       models.ElementHelper
	ReportingSequence              models.SequenceRangeHelper
	ReportCreateDateTime           models.ElementHelper
	AccountOtherId                 models.ElementHelper
	TotalCreditEntries             models.NumberAndSumOfTransactionsHelper
	TotalDebitEntries              models.NumberAndSumOfTransactionsHelper
	TotalEntriesPerTransactionCode models.NumberAndStatusOfTransactionsHelper
	EntryDetails                   models.EntryHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: models.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message. Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.models`,
		},
		CreationDateTime: models.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		MessagePagination: models.BuildMessagePagenationHelper(),
		BussinessQueryMsgId: models.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.`,
		},
		BussinessQueryMsgNameId: models.ElementHelper{
			Title:         "Message Name Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the query message name identifier to which the message refers.`,
		},
		BussinessQueryCreateDatetime: models.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		ReportId: models.ElementHelper{
			Title:         "Report Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
		ReportingSequence: models.BuildSequenceRangeHelper(),
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
		TotalCreditEntries:             models.BuildNumberAndSumOfTransactionsHelper(),
		TotalDebitEntries:              models.BuildNumberAndSumOfTransactionsHelper(),
		TotalEntriesPerTransactionCode: models.BuildNumberAndStatusOfTransactionsHelper(),
		EntryDetails:                   models.BuildEntryHelper(),
	}
}
