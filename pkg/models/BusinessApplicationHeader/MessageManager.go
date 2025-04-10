package BusinessApplicationHeader

import (
	"reflect"
	"time"

	head001 "github.com/moov-io/fedwire20022/gen/BusinessApplicationHeader_head_001_001_03"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

type MarketPractice struct {
	// specifies the URL or reference to a registry where market practice guidelines are maintained.
	ReferenceRegistry string
	//uniquely identifies a specific market practice framework that applies to the message.
	FrameworkId string
}

type BusinessApplicationHeader struct {
	MessageSenderId   string
	MessageReceiverId string

	BusinessMessageId string
	//MsgDefIdr (Message Definition Identifier) and BizSvc (Business Service) are part of the Business Application Header (BAH), which helps identify and process financial messages.
	MessageDefinitionId string
	//BizSvc specifies a business service or process related to the message.
	BusinessService string
	//<MktPrctc> (Market Practice) is used to specify market-specific rules and guidelines that apply to the message.
	MarketInfo MarketPractice

	CreateDatetime time.Time
	//BizPrcgDt stands for Business Processing Date. It represents the date when a financial transaction or message is processed by a financial institution.
	BusinessProcessingDate time.Time
}

func ImplementationSpecification11From(p MarketPractice) head001.ImplementationSpecification11 {
	var result head001.ImplementationSpecification11
	if p.ReferenceRegistry != "" {
		result.Regy = head001.Max350TextFixed(p.ReferenceRegistry)
	}
	if p.FrameworkId != "" {
		result.Id = head001.MarketPracticeIdentificationFedwireFunds1(p.FrameworkId)
	}
	return result
}

func BusinessApplicationHeader71From(p BusinessApplicationHeader) head001.BusinessApplicationHeader71 {
	var result head001.BusinessApplicationHeader71
	if p.MessageSenderId != "" {
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(p.MessageSenderId),
				},
			},
		}
		result.Fr = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if p.MessageReceiverId != "" {
		_FIId := head001.BranchAndFinancialInstitutionIdentification61{
			FinInstnId: head001.FinancialInstitutionIdentification181{
				ClrSysMmbId: head001.ClearingSystemMemberIdentification21{
					MmbId: head001.ConnectionPartyIdentifierFedwireFunds1(p.MessageReceiverId),
				},
			},
		}
		result.To = head001.Party44Choice1{
			FIId: &_FIId,
		}
	}
	if p.BusinessMessageId != "" {
		result.BizMsgIdr = head001.Max35Text(p.BusinessMessageId)
	}
	if p.MessageDefinitionId != "" {
		result.MsgDefIdr = head001.MessageNameIdentificationFRS1(p.MessageDefinitionId)
	}
	if p.BusinessService != "" {
		result.BizSvc = head001.BusinessServiceFedwireFunds1(p.BusinessService)
	}
	if !isEmpty(p.MarketInfo) {
		MktPrctc := head001.ImplementationSpecification12{
			Regy: head001.Max350TextFixed(p.MarketInfo.ReferenceRegistry),
			Id:   head001.Max2048Text(p.MarketInfo.FrameworkId),
		}
		if !isEmpty(MktPrctc) {
			result.MktPrctc = MktPrctc
		}
	}
	if !isEmpty(p.CreateDatetime) {
		result.CreDt = fedwire.ISODateTime(p.CreateDatetime)
	}
	if !isEmpty(p.BusinessProcessingDate) {
		BizPrcgDt := fedwire.ISODateTime(p.BusinessProcessingDate)
		result.BizPrcgDt = &BizPrcgDt
	}

	return result
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
