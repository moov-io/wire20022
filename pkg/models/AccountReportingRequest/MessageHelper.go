package AccountReportingRequest

import "github.com/wadearnold/wire20022/pkg/models"

type MessageHelper struct {
	MessageId          models.ElementHelper
	CreatedDateTime    models.ElementHelper
	ReportRequestId    models.ElementHelper
	RequestedMsgNameId models.ElementHelper
	AccountOtherId     models.ElementHelper
	AccountProperty    models.ElementHelper
	AccountOwnerAgent  models.AgentHelper
	FromToSeuence      models.SequenceRangeHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: models.ElementHelper{
			Title:         "Message Identification",
			Rules:         "Message Identification is the reference of the account reporting request assigned by the Fedwire Sender.",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the account owner or the party authorised to send the message, and sent to the account servicing institution, to unambiguously identify the message. Usage: The sender has to make sure that 'MessageIdentification' is unique per account servicing institution for a pre-agreed period.`,
		},
		CreatedDateTime: models.ElementHelper{
			Title:         "Creation Date Time",
			Rules:         "Must be date and time when the message is created by the Fedwire Sender. Time must be in 24-hour clock format and either in Coordinated Universal Time (UTC) or in local time with offset against UTC.",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time at which the message was created.`,
		},
		ReportRequestId: models.ElementHelper{
			Title:         "Report Request Identification",
			Rules:         `- EndpointDetailsRule If an endpoint details sent report (DTLS) or an endpoint details received report (DTLR) is requested, then Reporting Sequence and the Fedwire Funds participant's endpoint (Account Owner Other) are mandatory, and Account is not allowed. - AccountBalanceRule If an account balance report (ABAR) is requested, then Account Type is mandatory, and Reporting Sequence and the Fedwire Funds participant's endpoint (Account Owner Other) are not allowed. - EndpointTotalsRule If an endpoint totals report (ETOT) is requested, then the Fedwire Funds participant's endpoint (Account Owner Other) is mandatory, and Account and Reporting Sequence are not allowed. `,
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unique identification, as assigned by the account owner, to unambiguously identify the account reporting request.`,
		},
		RequestedMsgNameId: models.ElementHelper{
			Title:         "Requested Message Name Identification",
			Rules:         "This must contain the requested account report message name identification (i.e., camt.052.001.08 or a subsequent version of that message as it is introduced in a future release of the Fedwire Funds Service).",
			Type:          `MessageNameIdentification_FRS_1 (based on string) exactLength: 15 pattern: [a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`,
			Documentation: `Specifies the type of the requested reporting message.`,
		},
		AccountOtherId: models.ElementHelper{
			Title:         "Account Identification",
			Rules:         `Must be the Inquiry Routing Number. Usage: It may be a subaccount routing number or the master account routing number. This element in conjunction with the balance type code in the Account\Type element will determine the information reported in the Account Balance Report.`,
			Type:          `RoutingNumber_FRS_1 (based on string) exactLength: 9 pattern: [0-9]{9,9}`,
			Documentation: `Identification assigned by an institution.`,
		},
		AccountProperty: models.ElementHelper{
			Title:         "Account Type Proprietary",
			Rules:         "",
			Type:          `AccountTypeFRS(AccountTypeSavings, AccountTypeMerchant)`,
			Documentation: `Nature or use of the account in a proprietary form.`,
		},
		AccountOwnerAgent: models.BuildAgentHelper(),
		FromToSeuence:     models.BuildSequenceRangeHelper(),
	}
}
