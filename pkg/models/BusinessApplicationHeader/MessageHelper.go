package BusinessApplicationHeader

import (
	model "github.com/moov-io/wire20022/pkg/models"
)

type MarketPracticeHelper struct {
	ReferenceRegistry model.ElementHelper
	FrameworkId       model.ElementHelper
}

func BuildMarketPracticeHelper() MarketPracticeHelper {
	return MarketPracticeHelper{
		ReferenceRegistry: model.ElementHelper{
			Title:         "Registry",
			Rules:         "www2.swift.com/mystandards/#/group/Federal_Reserve_Financial_Services/Fedwire_Funds_Service",
			Type:          `Max350Text (based on string) minLength: 1 maxLength: 350`,
			Documentation: `Name of the implementation specification registry in which the implementation specification of the ISO 20022 message is maintained. For example, "MyStandards".`,
		},
		FrameworkId: model.ElementHelper{
			Title:         "Framework Id",
			Rules:         "Must contain the unique identifier of the Fedwire Funds Service implementation guideline to which the message identified in Message Definition Identifier conforms. ",
			Type:          `MarketPracticeIdentification_FedwireFunds_1 (based on string) minLength: 14 maxLength: 18 pattern: frb([.]{1,1})fedwire([.]{1,1})(([a-z]{3,3}[.]{1,1})){0,1}01`,
			Documentation: `Identifier which unambiguously identifies, within the implementation specification registry, the implementation specification to which the ISO 20022 message is compliant. This can be done via a URN. It can also contain a version number or date. For instance, "2018-01-01 – Version 2" or "urn:uuid:6e8bc430-9c3a-11d9-9669-0800200c9a66".`,
		},
	}
}

type BusinessApplicationHeaderHelper struct {
	MessageSenderId        model.ElementHelper
	MessageReceiverId      model.ElementHelper
	BusinessMessageId      model.ElementHelper
	MessageDefinitionId    model.ElementHelper
	BusinessService        model.ElementHelper
	MarketInfo             MarketPracticeHelper
	CreateDatetime         model.ElementHelper
	BusinessProcessingDate model.ElementHelper
}

func BuildBusinessApplicationHeaderHelper() BusinessApplicationHeaderHelper {
	return BusinessApplicationHeaderHelper{
		MessageSenderId: model.ElementHelper{
			Title:         "Message Sender Id",
			Rules:         "For messages sent to the Fedwire Funds Service application, this must be the Fedwire Funds connection party identifier from where the message is sent. For messages sent by the Fedwire Funds Service application, this is the Reserve Bank identifier used to identify the Fedwire Funds Service as sender of the message, i.e., 021151080.",
			Type:          `ConnectionPartyIdentifier_FedwireFunds_1 (based on string) pattern: [A-Z0-9]{9,9}`,
			Documentation: `Identification of a member of a clearing system.`,
		},
		MessageReceiverId: model.ElementHelper{
			Title:         "Message Receiver Id",
			Rules:         "For messages sent to the Fedwire Funds Service application, this must be the Fedwire Funds connection party identifier from where the message is sent. For messages sent by the Fedwire Funds Service application, this is the Reserve Bank identifier used to identify the Fedwire Funds Service as sender of the message, i.e., 021151080.",
			Type:          `ConnectionPartyIdentifier_FedwireFunds_1 (based on string) pattern: [A-Z0-9]{9,9}`,
			Documentation: `Identification of a member of a clearing system.`,
		},
		BusinessMessageId: model.ElementHelper{
			Title:         "Business Message Identifier",
			Rules:         "For messages sent by a Fedwire Funds participant this may be the same as the identification of the underlying ISO 20022 business message, e.g., Message Identification for a pacs.008. ",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguously identifies the Business Message to the MessagingEndpoint that has created the Business Message.`,
		},
		MessageDefinitionId: model.ElementHelper{
			Title:         "Message Definition Identifier",
			Rules:         "This must be the ISO 20022 message definition identifier of the underlying ISO 20022 business message (e.g., pacs.008.001.08, pacs.009.001.08, pain.013.001.07 etc. or a subsequent version of the messages as they are introduced in a future release of the Fedwire Funds Service).",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `The Message Definition Identifier of the Business Message instance with which this Business Application Header instance is associated.`,
		},
		BusinessService: model.ElementHelper{
			Title:         "Business Service",
			Rules:         "This indicates the Fedwire Funds Service environment, i.e., Production or Test.",
			Type:          `BusinessService_FedwireFunds_1 (based on string)`,
			Documentation: `Specifies the business service agreed between the two MessagingEndpoints under which rules this Business Message is exchanged. To be used when there is a choice of processing services or processing service levels.`,
		},
		MarketInfo: BuildMarketPracticeHelper(),
		CreateDatetime: model.ElementHelper{
			Title:         "Creation Date",
			Rules:         "For messages sent by a Fedwire Funds participant this must be the calendar date and time when the ISO 20022 business message is sent. Time must be in 24-hour clock format and either in Coordinated Universal Time (UTC) or in local time with offset against UTC. ",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time when this Business Message (header) was created.`,
		},
		BusinessProcessingDate: model.ElementHelper{
			Title:         "Business Processing Date",
			Rules:         "For all end-to-end value and nonvalue messages delivered by the Fedwire Funds Service application, Business Processing Date provides the date and time the message was processed by the Fedwire Funds Service application. ",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Processing date and time indicated by the sender for the receiver of the business message. This date may be different from the date and time provided in the CreationDate. Usage: Market practice or bilateral agreement should specify how this element should be used.`,
		},
	}

}

type MessageHelper struct {
	MessageSenderId        model.ElementHelper
	MessageReceiverId      model.ElementHelper
	BusinessMessageId      model.ElementHelper
	MessageDefinitionId    model.ElementHelper
	BusinessService        model.ElementHelper
	MarketInfo             MarketPracticeHelper
	CreateDatetime         model.ElementHelper
	BusinessProcessingDate model.ElementHelper
	Relations              BusinessApplicationHeaderHelper
}

func BuildMessageHelper() MessageHelper {
	return MessageHelper{
		MessageSenderId: model.ElementHelper{
			Title:         "Message Sender Id",
			Rules:         "For messages sent to the Fedwire Funds Service application, this must be the Fedwire Funds connection party identifier from where the message is sent. For messages sent by the Fedwire Funds Service application, this is the Reserve Bank identifier used to identify the Fedwire Funds Service as sender of the message, i.e., 021151080.",
			Type:          `ConnectionPartyIdentifier_FedwireFunds_1 (based on string) pattern: [A-Z0-9]{9,9}`,
			Documentation: `Identification of a member of a clearing system.`,
		},
		MessageReceiverId: model.ElementHelper{
			Title:         "Message Receiver Id",
			Rules:         "For messages sent to the Fedwire Funds Service application, this must be the Fedwire Funds connection party identifier from where the message is sent. For messages sent by the Fedwire Funds Service application, this is the Reserve Bank identifier used to identify the Fedwire Funds Service as sender of the message, i.e., 021151080.",
			Type:          `ConnectionPartyIdentifier_FedwireFunds_1 (based on string) pattern: [A-Z0-9]{9,9}`,
			Documentation: `Identification of a member of a clearing system.`,
		},
		BusinessMessageId: model.ElementHelper{
			Title:         "Business Message Identifier",
			Rules:         "For messages sent by a Fedwire Funds participant this may be the same as the identification of the underlying ISO 20022 business message, e.g., Message Identification for a pacs.008. ",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `Unambiguously identifies the Business Message to the MessagingEndpoint that has created the Business Message.`,
		},
		MessageDefinitionId: model.ElementHelper{
			Title:         "Message Definition Identifier",
			Rules:         "This must be the ISO 20022 message definition identifier of the underlying ISO 20022 business message (e.g., pacs.008.001.08, pacs.009.001.08, pain.013.001.07 etc. or a subsequent version of the messages as they are introduced in a future release of the Fedwire Funds Service).",
			Type:          `Max35Text (based on string) minLength: 1 maxLength: 35`,
			Documentation: `The Message Definition Identifier of the Business Message instance with which this Business Application Header instance is associated.`,
		},
		BusinessService: model.ElementHelper{
			Title:         "Business Service",
			Rules:         "This indicates the Fedwire Funds Service environment, i.e., Production or Test.",
			Type:          `BusinessService_FedwireFunds_1 (based on string)`,
			Documentation: `Specifies the business service agreed between the two MessagingEndpoints under which rules this Business Message is exchanged. To be used when there is a choice of processing services or processing service levels.`,
		},
		MarketInfo: BuildMarketPracticeHelper(),
		CreateDatetime: model.ElementHelper{
			Title:         "Creation Date",
			Rules:         "For messages sent by a Fedwire Funds participant this must be the calendar date and time when the ISO 20022 business message is sent. Time must be in 24-hour clock format and either in Coordinated Universal Time (UTC) or in local time with offset against UTC. ",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Date and time when this Business Message (header) was created.`,
		},
		BusinessProcessingDate: model.ElementHelper{
			Title:         "Business Processing Date",
			Rules:         "For all end-to-end value and nonvalue messages delivered by the Fedwire Funds Service application, Business Processing Date provides the date and time the message was processed by the Fedwire Funds Service application. ",
			Type:          `ISODateTime (based on dateTime)`,
			Documentation: `Processing date and time indicated by the sender for the receiver of the business message. This date may be different from the date and time provided in the CreationDate. Usage: Market practice or bilateral agreement should specify how this element should be used.`,
		},
		Relations: BuildBusinessApplicationHeaderHelper(),
	}
}
