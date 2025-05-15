package InvestRequest

import (
	"reflect"
	"time"

	camt110 "github.com/moov-io/fedwire20022/gen/InvestigationRequest_camt_110_001_01"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type Underlying struct {
	//Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageId string
	//Specifies the original message name identifier to which the message refers.
	OriginalMessageNameId string
	//Date and time at which the original message was created.
	OriginalCreationDateTime time.Time
	//Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OriginalInstructionId string
	//Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OriginalEndToEndId string
	//Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OriginalUETR string
	//Amount of money moved between the instructing agent and the instructed agent, as provided in the original instruction.
	OriginalInterbankSettlementAmount model.CurrencyAndAmount
	//Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OriginalInterbankSettlementDate model.Date
}

type InvestigationReason struct {
	Reason                string
	AdditionalRequestData string
}

func InvestigationReason21From(p InvestigationReason) (camt110.InvestigationReason21, *model.ValidateError) {
	var result camt110.InvestigationReason21
	if p.Reason != "" {
		err := camt110.ExternalInvestigationReason1Code(p.Reason).Validate()
		if err != nil {
			return camt110.InvestigationReason21{}, &model.ValidateError{
				ParamName: "Reason",
				Message:   err.Error(),
			}
		}
		Cd := camt110.ExternalInvestigationReason1Code(p.Reason)
		result.Rsn = camt110.InvestigationReason1Choice1{
			Cd: &Cd,
		}
	}
	if p.AdditionalRequestData != "" {
		err := camt110.Max500Text(p.AdditionalRequestData).Validate()
		if err != nil {
			return camt110.InvestigationReason21{}, &model.ValidateError{
				ParamName: "AdditionalRequestData",
				Message:   err.Error(),
			}
		}
		ReqNrrtv := camt110.Max500Text(p.AdditionalRequestData)
		AddtlReqData := camt110.AdditionalRequestData1Choice1{
			ReqNrrtv: &ReqNrrtv,
		}
		result.AddtlReqData = &AddtlReqData
	}
	return result, nil
}
func InvestigationReason21To(p camt110.InvestigationReason21) InvestigationReason {
	var result InvestigationReason
	if !isEmpty(p.Rsn.Cd) {
		result.Reason = string(*p.Rsn.Cd)
	}
	if !isEmpty(p.AddtlReqData) && !isEmpty(p.AddtlReqData.ReqNrrtv) {
		result.AdditionalRequestData = string(*p.AddtlReqData.ReqNrrtv)
	}
	return result
}
func UnderlyingData2Choice1From(p Underlying) (camt110.UnderlyingData2Choice1, *model.ValidateError) {
	var result camt110.UnderlyingData2Choice1
	var IntrBk camt110.UnderlyingPaymentTransaction71
	var OrgnlGrpInf camt110.UnderlyingGroupInformation11
	if p.OriginalMessageId != "" {
		err := camt110.Max35Text(p.OriginalMessageId).Validate()
		if err != nil {
			return camt110.UnderlyingData2Choice1{}, &model.ValidateError{
				ParamName: "OriginalMessageId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgId = camt110.Max35Text(p.OriginalMessageId)
	}
	if p.OriginalMessageNameId != "" {
		err := camt110.MessageNameIdentificationFRS1(p.OriginalMessageNameId).Validate()
		if err != nil {
			return camt110.UnderlyingData2Choice1{}, &model.ValidateError{
				ParamName: "OriginalMessageNameId",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlMsgNmId = camt110.MessageNameIdentificationFRS1(p.OriginalMessageNameId)
	}
	if !isEmpty(p.OriginalCreationDateTime) {
		err := fedwire.ISODateTime(p.OriginalCreationDateTime).Validate()
		if err != nil {
			return camt110.UnderlyingData2Choice1{}, &model.ValidateError{
				ParamName: "OriginalCreationDateTime",
				Message:   err.Error(),
			}
		}
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(p.OriginalCreationDateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		IntrBk.OrgnlGrpInf = OrgnlGrpInf
	}
	if p.OriginalInstructionId != "" {
		err := camt110.Max35Text(p.OriginalInstructionId).Validate()
		if err != nil {
			return camt110.UnderlyingData2Choice1{}, &model.ValidateError{
				ParamName: "OriginalInstructionId",
				Message:   err.Error(),
			}
		}
		OrgnlPmtInfId := camt110.Max35Text(p.OriginalInstructionId)
		IntrBk.OrgnlInstrId = &OrgnlPmtInfId
	}
	if p.OriginalEndToEndId != "" {
		err := camt110.Max35Text(p.OriginalEndToEndId).Validate()
		if err != nil {
			return camt110.UnderlyingData2Choice1{}, &model.ValidateError{
				ParamName: "OriginalEndToEndId",
				Message:   err.Error(),
			}
		}
		OrgnlEndToEndId := camt110.Max35Text(p.OriginalEndToEndId)
		IntrBk.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if p.OriginalUETR != "" {
		err := camt110.UUIDv4Identifier(p.OriginalUETR).Validate()
		if err != nil {
			return camt110.UnderlyingData2Choice1{}, &model.ValidateError{
				ParamName: "OriginalUETR",
				Message:   err.Error(),
			}
		}
		OrgnlUETR := camt110.UUIDv4Identifier(p.OriginalUETR)
		IntrBk.OrgnlUETR = OrgnlUETR
	}
	if !isEmpty(p.OriginalInterbankSettlementAmount) {
		err := fedwire.Amount(p.OriginalInterbankSettlementAmount.Amount).Validate()
		if err != nil {
			return camt110.UnderlyingData2Choice1{}, &model.ValidateError{
				ParamName: "OriginalInterbankSettlementAmount.Amount",
				Message:   err.Error(),
			}
		}
		err = camt110.ActiveOrHistoricCurrencyCode(p.OriginalInterbankSettlementAmount.Currency).Validate()
		if err != nil {
			return camt110.UnderlyingData2Choice1{}, &model.ValidateError{
				ParamName: "OriginalInterbankSettlementAmount.Currency",
				Message:   err.Error(),
			}
		}
		OrgnlIntrBkSttlmAmt := camt110.ActiveOrHistoricCurrencyAndAmount{
			Value: camt110.ActiveOrHistoricCurrencyAndAmountSimpleType(p.OriginalInterbankSettlementAmount.Amount),
			Ccy:   camt110.ActiveOrHistoricCurrencyCode(p.OriginalInterbankSettlementAmount.Currency),
		}
		IntrBk.OrgnlIntrBkSttlmAmt = &OrgnlIntrBkSttlmAmt
	}
	if !isEmpty(p.OriginalInterbankSettlementDate) {
		err := p.OriginalInterbankSettlementDate.Date().Validate()
		if err != nil {
			return camt110.UnderlyingData2Choice1{}, &model.ValidateError{
				ParamName: "OriginalInterbankSettlementDate.Amount",
				Message:   err.Error(),
			}
		}
		OrgnlIntrBkSttlmDt := p.OriginalInterbankSettlementDate.Date()
		IntrBk.OrgnlIntrBkSttlmDt = &OrgnlIntrBkSttlmDt
	}
	if !isEmpty(IntrBk) {
		result.IntrBk = &IntrBk
	}
	return result, nil
}
func UnderlyingData2Choice1To(p camt110.UnderlyingData2Choice1) Underlying {
	var result Underlying
	if !isEmpty(p.IntrBk) {
		if !isEmpty(p.IntrBk.OrgnlGrpInf) {
			result.OriginalMessageId = string(p.IntrBk.OrgnlGrpInf.OrgnlMsgId)
			result.OriginalMessageNameId = string(p.IntrBk.OrgnlGrpInf.OrgnlMsgNmId)
			result.OriginalCreationDateTime = time.Time(p.IntrBk.OrgnlGrpInf.OrgnlCreDtTm)
		}
		if !isEmpty(p.IntrBk.OrgnlInstrId) {
			result.OriginalInstructionId = string(*p.IntrBk.OrgnlInstrId)
		}
		if !isEmpty(p.IntrBk.OrgnlEndToEndId) {
			result.OriginalEndToEndId = string(*p.IntrBk.OrgnlEndToEndId)
		}
		if !isEmpty(p.IntrBk.OrgnlUETR) {
			result.OriginalUETR = string(p.IntrBk.OrgnlUETR)
		}
		if !isEmpty(p.IntrBk.OrgnlIntrBkSttlmAmt) {
			result.OriginalInterbankSettlementAmount = model.CurrencyAndAmount{
				Currency: string(p.IntrBk.OrgnlIntrBkSttlmAmt.Ccy),
				Amount:   float64(p.IntrBk.OrgnlIntrBkSttlmAmt.Value),
			}
		}
		if !isEmpty(p.IntrBk.OrgnlIntrBkSttlmDt) {
			result.OriginalInterbankSettlementDate = model.FromDate(*p.IntrBk.OrgnlIntrBkSttlmDt)
		}
	}
	return result
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
