package camt_060_001_03

func XMLINS() string {
	return "urn:iso:std:iso:20022:tech:xsd:camt.060.001.03"
}
func PathMap() map[string]string {
	return map[string]string{
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
