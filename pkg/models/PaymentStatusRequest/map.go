package PaymentStatusRequest

func PathMapV1() map[string]any {
	return PathMapV2()
}
func PathMapV2() map[string]any {
	return map[string]any{
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
func PathMapV3() map[string]any {
	return map[string]any{
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
func PathMapV4() map[string]any {
	return PathMapV3()
}
func PathMapV5() map[string]any {
	return PathMapV3()
}
func PathMapV6() map[string]any {
	return PathMapV3()
}
