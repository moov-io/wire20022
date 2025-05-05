package EndpointDetailsReport

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageHelper struct {
	MessageId                      model.ElementHelper
	CreationDateTime               model.ElementHelper
	MessagePagination              model.MessagePagenationHelper
	BussinessQueryMsgId            model.ElementHelper
	BussinessQueryMsgNameId        model.ElementHelper
	BussinessQueryCreateDatetime   model.ElementHelper
	ReportId                       model.ElementHelper
	ReportingSequence              model.SequenceRangeHelper
	ReportCreateDateTime           model.ElementHelper
	AccountOtherId                 model.ElementHelper
	TotalCreditEntries             model.NumberAndSumOfTransactionsHelper
	TotalDebitEntries              model.NumberAndSumOfTransactionsHelper
	TotalEntriesPerTransactionCode model.NumberAndStatusOfTransactionsHelper
	EntryDetails                   model.EntryHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message. Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.`,
		},
		CreationDateTime: model.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		MessagePagination: model.BuildMessagePagenationHelper(),
		BussinessQueryMsgId: model.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.`,
		},
		BussinessQueryMsgNameId: model.ElementHelper{
			Title:         "Message Name Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Specifies the query message name identifier to which the message refers.`,
		},
		BussinessQueryCreateDatetime: model.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		ReportId: model.ElementHelper{
			Title:         "Report Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
		ReportingSequence: model.BuildSequenceRangeHelper(),
		ReportCreateDateTime: model.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the report was created.`,
		},
		AccountOtherId: model.ElementHelper{
			Title:         "Account Other Id",
			Rules:         "",
			Type:          `Max34Text (based on string) minLength: 1 maxLength: 34`,
			Documentation: `Unique identification of an account, as assigned by the account servicer, using an identification scheme.`,
		},
		TotalCreditEntries:             model.BuildNumberAndSumOfTransactionsHelper(),
		TotalDebitEntries:              model.BuildNumberAndSumOfTransactionsHelper(),
		TotalEntriesPerTransactionCode: model.BuildNumberAndStatusOfTransactionsHelper(),
		EntryDetails:                   model.BuildEntryHelper(),
	}
}
