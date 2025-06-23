package FedwireFundsPaymentStatus

func pathMapV3() map[string]any {
	return pathMapV5()
}
func pathMapV4() map[string]any {
	return pathMapV5()
}
func pathMapV5() map[string]any {
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
func pathMapV6() map[string]any {
	return pathMapV9()
}
func pathMapV7() map[string]any {
	return pathMapV9()
}
func pathMapV8() map[string]any {
	return pathMapV9()
}
func pathMapV9() map[string]any {
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
func pathMapV10() map[string]any {
	return map[string]any{
		"FIToFIPmtStsRpt.GrpHdr.MsgId":                                               "MessageId",
		"FIToFIPmtStsRpt.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgId":                      "OriginalMessageId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgNmId":                    "OriginalMessageNameId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf.OrgnlCreDtTm":                    "OriginalMessageCreateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlUETR":                                   "EnhancedTransaction.OriginalUETR",
		"FIToFIPmtStsRpt.TxInfAndSts[0].TxSts":                                       "TransactionStatus",
		"FIToFIPmtStsRpt.TxInfAndSts[0].AccptncDtTm":                                 "AcceptanceDateTime",
		"FIToFIPmtStsRpt.TxInfAndSts[0].FctvIntrBkSttlmDt.Dt":                        "EnhancedTransaction.EffectiveInterbankSettlementDate",
		"FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].Rsn.Prtry":                      "StatusReasonInformation",
		"FIToFIPmtStsRpt.TxInfAndSts[0].StsRsnInf[0].AddtlInf[0]":                    "ReasonAdditionalInfo",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructingAgent.PaymentSysCode",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructingAgent.PaymentSysMemberId",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructedAgent.PaymentSysCode",
		"FIToFIPmtStsRpt.TxInfAndSts[0].InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructedAgent.PaymentSysMemberId",
	}
}
func pathMapV11() map[string]any {
	return pathMapV10()
}
func pathMapV12() map[string]any {
	return pathMapV10()
}
func pathMapV13() map[string]any {
	return pathMapV10()
}
func pathMapV14() map[string]any {
	return pathMapV10()
}
