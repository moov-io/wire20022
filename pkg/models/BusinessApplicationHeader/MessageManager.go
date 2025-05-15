package BusinessApplicationHeader

import (
	"reflect"
	"time"

	head001 "github.com/moov-io/fedwire20022/gen/BusinessApplicationHeader_head_001_001_03"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
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

func ImplementationSpecification11From(p MarketPractice) (head001.ImplementationSpecification11, *model.ValidateError) {
	var result head001.ImplementationSpecification11
	if p.ReferenceRegistry != "" {
		err := head001.Max350TextFixed(p.ReferenceRegistry).Validate()
		if err != nil {
			return head001.ImplementationSpecification11{}, &model.ValidateError{
				ParamName: "ReferenceRegistry",
				Message:   err.Error(),
			}
		}
		result.Regy = head001.Max350TextFixed(p.ReferenceRegistry)
	}
	if p.FrameworkId != "" {
		err := head001.MarketPracticeIdentificationFedwireFunds1(p.FrameworkId).Validate()
		if err != nil {
			return head001.ImplementationSpecification11{}, &model.ValidateError{
				ParamName: "FrameworkId",
				Message:   err.Error(),
			}
		}
		result.Id = head001.MarketPracticeIdentificationFedwireFunds1(p.FrameworkId)
	}
	return result, nil
}
func ImplementationSpecification11To(p head001.ImplementationSpecification11) MarketPractice {
	var result MarketPractice
	if !isEmpty(p.Regy) {
		result.ReferenceRegistry = string(p.Regy)
	}
	if !isEmpty(p.Id) {
		result.FrameworkId = string(p.Id)
	}
	return result
}
func BusinessApplicationHeader71From(p BusinessApplicationHeader) (head001.BusinessApplicationHeader71, *model.ValidateError) {
	var result head001.BusinessApplicationHeader71
	if p.MessageSenderId != "" {
		err := head001.ConnectionPartyIdentifierFedwireFunds1(p.MessageSenderId).Validate()
		if err != nil {
			return head001.BusinessApplicationHeader71{}, &model.ValidateError{
				ParamName: "MessageSenderId",
				Message:   err.Error(),
			}
		}
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
		err := head001.ConnectionPartyIdentifierFedwireFunds1(p.MessageReceiverId).Validate()
		if err != nil {
			return head001.BusinessApplicationHeader71{}, &model.ValidateError{
				ParamName: "MessageReceiverId",
				Message:   err.Error(),
			}
		}
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
		err := head001.Max35Text(p.BusinessMessageId).Validate()
		if err != nil {
			return head001.BusinessApplicationHeader71{}, &model.ValidateError{
				ParamName: "BusinessMessageId",
				Message:   err.Error(),
			}
		}
		result.BizMsgIdr = head001.Max35Text(p.BusinessMessageId)
	}
	if p.MessageDefinitionId != "" {
		err := head001.MessageNameIdentificationFRS1(p.MessageDefinitionId).Validate()
		if err != nil {
			return head001.BusinessApplicationHeader71{}, &model.ValidateError{
				ParamName: "MessageDefinitionId",
				Message:   err.Error(),
			}
		}
		result.MsgDefIdr = head001.MessageNameIdentificationFRS1(p.MessageDefinitionId)
	}
	if p.BusinessService != "" {
		err := head001.BusinessServiceFedwireFunds1(p.BusinessService).Validate()
		if err != nil {
			return head001.BusinessApplicationHeader71{}, &model.ValidateError{
				ParamName: "BusinessService",
				Message:   err.Error(),
			}
		}
		result.BizSvc = head001.BusinessServiceFedwireFunds1(p.BusinessService)
	}
	if !isEmpty(p.MarketInfo) {
		err := head001.Max350TextFixed(p.MarketInfo.ReferenceRegistry).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "ReferenceRegistry",
				Message:   err.Error(),
			}
			vErr.InsertPath("MarketInfo")
			return head001.BusinessApplicationHeader71{}, &vErr
		}
		err = head001.Max2048Text(p.MarketInfo.FrameworkId).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "FrameworkId",
				Message:   err.Error(),
			}
			vErr.InsertPath("MarketInfo")
			return head001.BusinessApplicationHeader71{}, &vErr
		}
		MktPrctc := head001.ImplementationSpecification12{
			Regy: head001.Max350TextFixed(p.MarketInfo.ReferenceRegistry),
			Id:   head001.Max2048Text(p.MarketInfo.FrameworkId),
		}
		if !isEmpty(MktPrctc) {
			result.MktPrctc = MktPrctc
		}
	}
	if !isEmpty(p.CreateDatetime) {
		err := fedwire.ISODateTime(p.CreateDatetime).Validate()
		if err != nil {
			return head001.BusinessApplicationHeader71{}, &model.ValidateError{
				ParamName: "CreateDatetime",
				Message:   err.Error(),
			}
		}
		result.CreDt = fedwire.ISODateTime(p.CreateDatetime)
	}
	if !isEmpty(p.BusinessProcessingDate) {
		err := fedwire.ISODateTime(p.BusinessProcessingDate).Validate()
		if err != nil {
			return head001.BusinessApplicationHeader71{}, &model.ValidateError{
				ParamName: "BusinessProcessingDate",
				Message:   err.Error(),
			}
		}
		BizPrcgDt := fedwire.ISODateTime(p.BusinessProcessingDate)
		result.BizPrcgDt = &BizPrcgDt
	}

	return result, nil
}
func BusinessApplicationHeader71To(p head001.BusinessApplicationHeader71) BusinessApplicationHeader {
	var result BusinessApplicationHeader
	if !isEmpty(p.Fr) && !isEmpty(p.Fr.FIId) && !isEmpty(p.Fr.FIId.FinInstnId) && !isEmpty(p.Fr.FIId.FinInstnId.ClrSysMmbId) && !isEmpty(p.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId) {
		result.MessageSenderId = string(p.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId)
	}
	if !isEmpty(p.To) && !isEmpty(p.To.FIId) && !isEmpty(p.To.FIId.FinInstnId) && !isEmpty(p.To.FIId.FinInstnId.ClrSysMmbId) && !isEmpty(p.To.FIId.FinInstnId.ClrSysMmbId.MmbId) {
		result.MessageReceiverId = string(p.To.FIId.FinInstnId.ClrSysMmbId.MmbId)
	}
	if !isEmpty(p.BizMsgIdr) {
		result.BusinessMessageId = string(p.BizMsgIdr)
	}
	if !isEmpty(p.MsgDefIdr) {
		result.MessageDefinitionId = string(p.MsgDefIdr)
	}
	if !isEmpty(p.BizSvc) {
		result.BusinessService = string(p.BizSvc)
	}
	if !isEmpty(p.MktPrctc) {
		if !isEmpty(p.MktPrctc.Regy) {
			result.MarketInfo.ReferenceRegistry = string(p.MktPrctc.Regy)
		}
		if !isEmpty(p.MktPrctc.Id) {
			result.MarketInfo.FrameworkId = string(p.MktPrctc.Id)
		}
	}
	if !isEmpty(p.CreDt) {
		result.CreateDatetime = time.Time(p.CreDt)
	}
	if !isEmpty(p.BizPrcgDt) {
		result.BusinessProcessingDate = time.Time(*p.BizPrcgDt)
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
