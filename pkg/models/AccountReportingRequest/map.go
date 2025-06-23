package AccountReportingRequest

func pathMapV2() map[string]any {
	return map[string]any{
		"AcctRptgReq.GrpHdr.MsgId":                                               "MessageId",
		"AcctRptgReq.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"AcctRptgReq.RptgReq[0].Id":                                              "ReportRequestId",
		"AcctRptgReq.RptgReq[0].ReqdMsgNmId":                                     "RequestedMsgNameId",
		"AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id":                                 "AccountOtherId",
		"AcctRptgReq.RptgReq[0].Acct.Tp.Prtry":                                   "AccountProperty",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "AccountOwnerAgent.PaymentSysCode",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "AccountOwnerAgent.PaymentSysMemberId",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.Othr.Id":                 "AccountOwnerAgent.OtherTypeId",
	}
}

func pathMapV3() map[string]any {
	return map[string]any{
		"AcctRptgReq.GrpHdr.MsgId":                                               "MessageId",
		"AcctRptgReq.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"AcctRptgReq.RptgReq[0].Id":                                              "ReportRequestId",
		"AcctRptgReq.RptgReq[0].ReqdMsgNmId":                                     "RequestedMsgNameId",
		"AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id":                                 "AccountOtherId",
		"AcctRptgReq.RptgReq[0].Acct.Tp.Prtry":                                   "AccountProperty",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "AccountOwnerAgent.PaymentSysCode",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "AccountOwnerAgent.PaymentSysMemberId",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.Othr.Id":                 "AccountOwnerAgent.OtherTypeId",
	}
}

func pathMapV4() map[string]any {
	return map[string]any{
		"AcctRptgReq.GrpHdr.MsgId":                                               "MessageId",
		"AcctRptgReq.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"AcctRptgReq.RptgReq[0].Id":                                              "ReportRequestId",
		"AcctRptgReq.RptgReq[0].ReqdMsgNmId":                                     "RequestedMsgNameId",
		"AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id":                                 "AccountOtherId",
		"AcctRptgReq.RptgReq[0].Acct.Tp.Prtry":                                   "AccountProperty",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "AccountOwnerAgent.PaymentSysCode",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "AccountOwnerAgent.PaymentSysMemberId",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.Othr.Id":                 "AccountOwnerAgent.OtherTypeId",
		"AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].FrSeq":                        "FromToSequence.FromSeq",
		"AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].ToSeq":                        "FromToSequence.ToSeq",
	}
}

func pathMapV5() map[string]any {
	return map[string]any{
		"AcctRptgReq.GrpHdr.MsgId":                                               "MessageId",
		"AcctRptgReq.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"AcctRptgReq.RptgReq[0].Id":                                              "ReportRequestId",
		"AcctRptgReq.RptgReq[0].ReqdMsgNmId":                                     "RequestedMsgNameId",
		"AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id":                                 "AccountOtherId",
		"AcctRptgReq.RptgReq[0].Acct.Tp.Prtry":                                   "AccountProperty",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "AccountOwnerAgent.PaymentSysCode",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "AccountOwnerAgent.PaymentSysMemberId",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.Othr.Id":                 "AccountOwnerAgent.OtherTypeId",
		"AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].FrSeq":                        "FromToSequence.FromSeq",
		"AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].ToSeq":                        "FromToSequence.ToSeq",
	}
}

func pathMapV6() map[string]any {
	return map[string]any{
		"AcctRptgReq.GrpHdr.MsgId":                                               "MessageId",
		"AcctRptgReq.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"AcctRptgReq.RptgReq[0].Id":                                              "ReportRequestId",
		"AcctRptgReq.RptgReq[0].ReqdMsgNmId":                                     "RequestedMsgNameId",
		"AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id":                                 "AccountOtherId",
		"AcctRptgReq.RptgReq[0].Acct.Tp.Prtry":                                   "AccountProperty",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "AccountOwnerAgent.PaymentSysCode",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "AccountOwnerAgent.PaymentSysMemberId",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.Othr.Id":                 "AccountOwnerAgent.OtherTypeId",
		"AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].FrSeq":                        "FromToSequence.FromSeq",
		"AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].ToSeq":                        "FromToSequence.ToSeq",
	}
}

func pathMapV7() map[string]any {
	return map[string]any{
		"AcctRptgReq.GrpHdr.MsgId":                                               "MessageId",
		"AcctRptgReq.GrpHdr.CreDtTm":                                             "CreatedDateTime",
		"AcctRptgReq.RptgReq[0].Id":                                              "ReportRequestId",
		"AcctRptgReq.RptgReq[0].ReqdMsgNmId":                                     "RequestedMsgNameId",
		"AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id":                                 "AccountOtherId",
		"AcctRptgReq.RptgReq[0].Acct.Tp.Prtry":                                   "AccountProperty",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "AccountOwnerAgent.PaymentSysCode",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.ClrSysMmbId.MmbId":       "AccountOwnerAgent.PaymentSysMemberId",
		"AcctRptgReq.RptgReq[0].AcctOwnr.Agt.FinInstnId.Othr.Id":                 "AccountOwnerAgent.OtherTypeId",
		"AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].FrSeq":                        "ReportingSequence.FromToSequence.FromSeq",
		"AcctRptgReq.RptgReq[0].RptgSeq.FrToSeq[0].ToSeq":                        "ReportingSequence.FromToSequence.ToSeq",
	}
}
