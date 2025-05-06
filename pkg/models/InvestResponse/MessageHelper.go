package InvestigationResponse

import model "github.com/moov-io/wire20022/pkg/models"

type MessageHelper struct {
	//Point to point reference, as assigned by the responder, and sent to the requestor to unambiguously identify the message.
	MessageId model.ElementHelper
	//Status of the investigation request.
	InvestigationStatus model.ElementHelper
	//Provides the response to the request.
	InvestigationData model.ElementHelper
	//Point to point reference, as assigned by the requestor, and sent to the responder to unambiguously identify the message.
	InvestRequestMessageId model.ElementHelper
	//Type of investigation.
	InvestigationType model.ElementHelper
	//Identification of the agent or party requesting a new investigation is opened or status update for an existing investigation.
	Requestor model.AgentHelper
	//Identification of the agent or party expected to open a new investigation or provide a status update for an existing investigation.
	Responder model.AgentHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageId: model.ElementHelper{
			Title:         "Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the sender, to unambiguously identify the message. Usage: The sender has to make sure that MessageIdentification is unique for a pre-agreed period.`,
		},
		InvestigationStatus: model.ElementHelper{
			Title:         "Investigation Status",
			Rules:         "",
			Type:          `ExternalInvestigationStatus1Code (based on string)`,
			Documentation: `Status of the investigation request.`,
		},
		InvestigationData: model.ElementHelper{
			Title:         "Investigation Data",
			Rules:         "",
			Type:          `Max500Text (based on string) minLength: 1 maxLength: 500`,
			Documentation: `Provides the response to the request.`,
		},
		InvestRequestMessageId: model.ElementHelper{
			Title:         "Invest Request Message Id",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Point to point reference, as assigned by the requestor, and sent to the responder to unambiguously identify the message.`,
		},
		InvestigationType: model.ElementHelper{
			Title:         "Investigation Type",
			Rules:         "",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Type of investigation.`,
		},
		Requestor: model.BuildAgentHelper(),
		Responder: model.BuildAgentHelper(),
	}
}
