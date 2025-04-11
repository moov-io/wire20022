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

func InvestigationReason21From(p InvestigationReason) camt110.InvestigationReason21 {
	var result camt110.InvestigationReason21
	if p.Reason != "" {
		Cd := camt110.ExternalInvestigationReason1Code(p.Reason)
		result.Rsn = camt110.InvestigationReason1Choice1{
			Cd: &Cd,
		}
	}
	if p.AdditionalRequestData != "" {
		ReqNrrtv := camt110.Max500Text(p.AdditionalRequestData)
		AddtlReqData := camt110.AdditionalRequestData1Choice1{
			ReqNrrtv: &ReqNrrtv,
		}
		result.AddtlReqData = &AddtlReqData
	}
	return result
}
func UnderlyingData2Choice1From(p Underlying) camt110.UnderlyingData2Choice1 {
	var result camt110.UnderlyingData2Choice1
	var IntrBk camt110.UnderlyingPaymentTransaction71
	var OrgnlGrpInf camt110.UnderlyingGroupInformation11
	if p.OriginalMessageId != "" {
		OrgnlGrpInf.OrgnlMsgId = camt110.Max35Text(p.OriginalMessageId)
	}
	if p.OriginalMessageNameId != "" {
		OrgnlGrpInf.OrgnlMsgNmId = camt110.MessageNameIdentificationFRS1(p.OriginalMessageNameId)
	}
	if !isEmpty(p.OriginalCreationDateTime) {
		OrgnlGrpInf.OrgnlCreDtTm = fedwire.ISODateTime(p.OriginalCreationDateTime)
	}
	if !isEmpty(OrgnlGrpInf) {
		IntrBk.OrgnlGrpInf = OrgnlGrpInf
	}
	if p.OriginalInstructionId != "" {
		OrgnlPmtInfId := camt110.Max35Text(p.OriginalInstructionId)
		IntrBk.OrgnlInstrId = &OrgnlPmtInfId
	}
	if p.OriginalEndToEndId != "" {
		OrgnlEndToEndId := camt110.Max35Text(p.OriginalEndToEndId)
		IntrBk.OrgnlEndToEndId = &OrgnlEndToEndId
	}
	if p.OriginalUETR != "" {
		OrgnlUETR := camt110.UUIDv4Identifier(p.OriginalUETR)
		IntrBk.OrgnlUETR = OrgnlUETR
	}
	if !isEmpty(p.OriginalInterbankSettlementAmount) {
		OrgnlIntrBkSttlmAmt := camt110.ActiveOrHistoricCurrencyAndAmount{
			Value: camt110.ActiveOrHistoricCurrencyAndAmountSimpleType(p.OriginalInterbankSettlementAmount.Amount),
			Ccy:   camt110.ActiveOrHistoricCurrencyCode(p.OriginalInterbankSettlementAmount.Currency),
		}
		IntrBk.OrgnlIntrBkSttlmAmt = &OrgnlIntrBkSttlmAmt
	}
	if !isEmpty(p.OriginalInterbankSettlementDate) {
		OrgnlIntrBkSttlmDt := p.OriginalInterbankSettlementDate.Date()
		IntrBk.OrgnlIntrBkSttlmDt = &OrgnlIntrBkSttlmDt
	}
	if !isEmpty(IntrBk) {
		result.IntrBk = &IntrBk
	}
	return result
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
