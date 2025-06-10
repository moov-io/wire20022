package PaymentStatusRequest

import (
	"time"

	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_01"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type Message struct {
	Data MessageModel
	Doc  pacs_028_001_01.Document
}

func Convert() {
	msg := Message{}
	msg.Data.MessageId = string(msg.Doc.FIToFIPmtStsReq.GrpHdr.MsgId)
	msg.Data.CreatedDateTime = time.Time(msg.Doc.FIToFIPmtStsReq.GrpHdr.CreDtTm)
	msg.Data.OriginalMessageId = string(msg.Doc.FIToFIPmtStsReq.TxInf[0].OrgnlGrpInf.OrgnlMsgId)
	msg.Data.OriginalMessageNameId = string(msg.Doc.FIToFIPmtStsReq.TxInf[0].OrgnlGrpInf.OrgnlMsgNmId)
	msg.Data.OriginalCreationDateTime = time.Time(*msg.Doc.FIToFIPmtStsReq.TxInf[0].OrgnlGrpInf.OrgnlCreDtTm)
	msg.Data.OriginalInstructionId = string(*msg.Doc.FIToFIPmtStsReq.TxInf[0].OrgnlInstrId)
	msg.Data.OriginalEndToEndId = string(*msg.Doc.FIToFIPmtStsReq.TxInf[0].OrgnlEndToEndId)
	// msg.Data.OriginalUETR = string(*msg.Doc.FIToFIPmtStsReq.TxInf[0].OrgnlUETR)
	msg.Data.InstructingAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.FIToFIPmtStsReq.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.InstructingAgent.PaymentSysMemberId = string(msg.Doc.FIToFIPmtStsReq.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId)
	msg.Data.InstructedAgent.PaymentSysCode = Archive.PaymentSystemType(*msg.Doc.FIToFIPmtStsReq.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd)
	msg.Data.InstructedAgent.PaymentSysMemberId = string(msg.Doc.FIToFIPmtStsReq.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId)
}

func PathMapV1() map[string]string {
	return PathMapV2()
}
func PathMapV2() map[string]string {
	return map[string]string{
		"FIToFIPmtStsReq.GrpHdr.MsgId":                                         "MessageId",
		"FIToFIPmtStsReq.GrpHdr.CreDtTm":                                       "CreatedDateTime",
		"FIToFIPmtStsReq.TxInf[0].OrgnlGrpInf.OrgnlMsgId":                      "OriginalMessageId",
		"FIToFIPmtStsReq.TxInf[0].OrgnlGrpInf.OrgnlMsgNmId":                    "OriginalMessageNameId",
		"FIToFIPmtStsReq.TxInf[0].OrgnlGrpInf.OrgnlCreDtTm":                    "OriginalCreationDateTime",
		"FIToFIPmtStsReq.TxInf[0].OrgnlInstrId":                                "OriginalInstructionId",
		"FIToFIPmtStsReq.TxInf[0].OrgnlEndToEndId":                             "OriginalEndToEndId",
		"FIToFIPmtStsReq.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructingAgent.PaymentSysCode",
		"FIToFIPmtStsReq.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructingAgent.PaymentSysMemberId",
		"FIToFIPmtStsReq.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructedAgent.PaymentSysCode",
		"FIToFIPmtStsReq.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructedAgent.PaymentSysMemberId",
	}
}
func PathMapV3() map[string]string {
	return map[string]string{
		"FIToFIPmtStsReq.GrpHdr.MsgId":                                         "MessageId",
		"FIToFIPmtStsReq.GrpHdr.CreDtTm":                                       "CreatedDateTime",
		"FIToFIPmtStsReq.TxInf[0].OrgnlGrpInf.OrgnlMsgId":                      "OriginalMessageId",
		"FIToFIPmtStsReq.TxInf[0].OrgnlGrpInf.OrgnlMsgNmId":                    "OriginalMessageNameId",
		"FIToFIPmtStsReq.TxInf[0].OrgnlGrpInf.OrgnlCreDtTm":                    "OriginalCreationDateTime",
		"FIToFIPmtStsReq.TxInf[0].OrgnlInstrId":                                "OriginalInstructionId",
		"FIToFIPmtStsReq.TxInf[0].OrgnlEndToEndId":                             "OriginalEndToEndId",
		"FIToFIPmtStsReq.TxInf[0].OrgnlUETR":                                   "OriginalUETR",
		"FIToFIPmtStsReq.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructingAgent.PaymentSysCode",
		"FIToFIPmtStsReq.TxInf[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructingAgent.PaymentSysMemberId",
		"FIToFIPmtStsReq.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructedAgent.PaymentSysCode",
		"FIToFIPmtStsReq.TxInf[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructedAgent.PaymentSysMemberId",
	}
}
func PathMapV4() map[string]string {
	return PathMapV3()
}
func PathMapV5() map[string]string {
	return PathMapV3()
}
func PathMapV6() map[string]string {
	return PathMapV3()
}
