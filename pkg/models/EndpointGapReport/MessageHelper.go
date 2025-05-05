package EndpointGapReport

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageHelper struct {
	MessageId            model.ElementHelper
	CreatedDateTime      model.ElementHelper
	MessagePagination    model.MessagePagenationHelper
	ReportId             model.ElementHelper
	ReportCreateDateTime model.ElementHelper
	AccountOtherId       model.ElementHelper
	AdditionalReportInfo model.ElementHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Identification",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message. Usage: The account servicing institution has to make sure that MessageIdentification is unique per account owner for a pre-agreed period.`,
		},
		CreatedDateTime: model.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		MessagePagination: model.BuildMessagePagenationHelper(),
		ReportId: model.ElementHelper{
			Title:         "Report Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the account servicer, to unambiguously identify the account report.`,
		},
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
		AdditionalReportInfo: model.ElementHelper{
			Title:         "Additional Report Information",
			Rules:         "",
			Type:          `Max500Text (based on string) minLength: 1 maxLength: 500`,
			Documentation: `Further details of the account report.`,
		},
	}
}
