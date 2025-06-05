package FedwireFundsPaymentStatus

import (
	"time"

	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_03"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type Message struct {
	Data MessageModel
	Doc  pacs_002_001_03.Document
}

func Convert() {
	msg := Message{}

	msg.Data.MessageId = string(msg.Doc.FIToFIPmtStsRpt.GrpHdr.MsgId)
	msg.Data.CreatedDateTime = time.Time(msg.Doc.FIToFIPmtStsRpt.GrpHdr.CreDtTm)
	msg.Data.OriginalMessageId = string(msg.Doc.FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId)
	msg.Data.OriginalMessageNameId = string(msg.Doc.FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId)
	msg.Data.OriginalMessageCreateTime = time.Time(*msg.Doc.FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlCreDtTm)
	msg.Data.TransactionStatus = Archive.TransactionStatusCode(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts[0].TxSts)
	msg.Data.AcceptanceDateTime = time.Time(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts[0].AccptncDtTm)
	msg.Data.StatusReasonInformation = string(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].Rsn.Prtry)
	msg.Data.ReasonAdditionalInfo = string(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].AddtlInf[0])
	msg.Data.InstructingAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.InstructingAgent.PaymentSysMemberId = string(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.InstructedAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.InstructedAgent.PaymentSysMemberId = string(msg.Doc.FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId)
}

func PathMapV3() map[string]any {
	return PathMapV5()
}
func PathMapV4() map[string]any {
	return PathMapV5()
}
func PathMapV5() map[string]any {
	return map[string]any{
		"FIToFIPmtStsRpt.GrpHdr.MsgId":                                               "MessageId",
		"FIToFIPmtStsRpt.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgId":                               "OriginalMessageId",
		"FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlMsgNmId":                             "OriginalMessageNameId",
		"FIToFIPmtStsRpt.OrgnlGrpInfAndSts.OrgnlCreDtTm":                             "OriginalMessageCreateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].TxSts":                                       "TransactionStatus",
		"FIToFIPmtStsRpt.TxInfAndSts[0].AccptncDtTm":                                 "AcceptanceDateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].Rsn.Prtry":                      "StatusReasonInformation",
		"FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].AddtlInf[0]":                    "ReasonAdditionalInfo",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructingAgent.PaymentSysCode",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructingAgent.PaymentSysMemberId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructedAgent.PaymentSysCode",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructedAgent.PaymentSysMemberId",
	}
}
func PathMapV6() map[string]any {
	return PathMapV9()
}
func PathMapV7() map[string]any {
	return PathMapV9()
}
func PathMapV8() map[string]any {
	return PathMapV9()
}
func PathMapV9() map[string]any {
	return map[string]any{
		"FIToFIPmtStsRpt.GrpHdr.MsgId":                                               "MessageId",
		"FIToFIPmtStsRpt.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgId":                      "OriginalMessageId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgNmId":                    "OriginalMessageNameId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf.OrgnlCreDtTm":                    "OriginalMessageCreateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].TxSts":                                       "TransactionStatus",
		"FIToFIPmtStsRpt.TxInfAndSts[0].AccptncDtTm":                                 "AcceptanceDateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].Rsn.Prtry":                      "StatusReasonInformation",
		"FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].AddtlInf[0]":                    "ReasonAdditionalInfo",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructingAgent.PaymentSysCode",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructingAgent.PaymentSysMemberId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructedAgent.PaymentSysCode",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructedAgent.PaymentSysMemberId",
	}
}
func PathMapV10() map[string]any {
	return map[string]any{
		"FIToFIPmtStsRpt.GrpHdr.MsgId":                                               "MessageId",
		"FIToFIPmtStsRpt.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgId":                      "OriginalMessageId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgNmId":                    "OriginalMessageNameId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf.OrgnlCreDtTm":                    "OriginalMessageCreateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlUETR":                                   "OriginalUETR",
		"FIToFIPmtStsRpt.TxInfAndSts[0].TxSts":                                       "TransactionStatus",
		"FIToFIPmtStsRpt.TxInfAndSts[0].AccptncDtTm":                                 "AcceptanceDateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].FctvIntrBkSttlmDt.Dt":                        "EffectiveInterbankSettlementDate",
		"FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].Rsn.Prtry":                      "StatusReasonInformation",
		"FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].AddtlInf[0]":                    "ReasonAdditionalInfo",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructingAgent.PaymentSysCode",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructingAgent.PaymentSysMemberId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructedAgent.PaymentSysCode",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructedAgent.PaymentSysMemberId",
	}
}
func PathMapV11() map[string]any {
	return PathMapV10()
}
func PathMapV12() map[string]any {
	return PathMapV10()
}
func PathMapV13() map[string]any {
	return PathMapV10()
}
func PathMapV14() map[string]any {
	return PathMapV10()
}
