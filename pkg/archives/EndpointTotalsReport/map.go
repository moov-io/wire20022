package EndpointTotalsReport

func PathMapV1() map[string]any {
	return map[string]any{}
}
func PathMapV2() map[string]any {
	return PathMapV8()
}
func PathMapV3() map[string]any {
	return PathMapV8()
}
func PathMapV4() map[string]any {
	return PathMapV8()
}
func PathMapV5() map[string]any {
	return PathMapV8()
}
func PathMapV6() map[string]any {
	return PathMapV8()
}
func PathMapV7() map[string]any {
	return PathMapV8()
}
func PathMapV8() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":                             "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":                           "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":                     "Pagenation.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":                "Pagenation.LastPageIndicator",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":                 "BussinessQueryMsgId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId":               "BussinessQueryMsgNameId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm":               "BussinessQueryCreateDatetime",
		"BkToCstmrAcctRpt.Rpt[0].Id":                                "ReportId",
		"BkToCstmrAcctRpt.Rpt[0].RptgSeq.FrToSeq[0].FrSeq":          "ReportingSequence.FromSeq",
		"BkToCstmrAcctRpt.Rpt[0].RptgSeq.FrToSeq[0].ToSeq":          "ReportingSequence.ToSeq",
		"BkToCstmrAcctRpt.Rpt[0].CreDtTm":                           "ReportCreateDateTime",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":                   "AccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.NbOfNtries": "TotalCreditEntries.NumberOfEntries",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.Sum":        "TotalCreditEntries.Sum",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.NbOfNtries": "TotalDebitEntries.NumberOfEntries",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.Sum":        "TotalDebitEntries.Sum",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TotalEntriesPerBankTransactionCode": map[string]string{
			"NbOfNtries":      "NumberOfEntries",
			"BkTxCd.Prtry.Cd": "BankTransactionCode",
		},
		"BkToCstmrAcctRpt.Rpt[0].AddtlRptInf": "AdditionalReportInfo",
	}
}
func PathMapV9() map[string]any {
	return PathMapV8()
}
func PathMapV10() map[string]any {
	return PathMapV8()
}
func PathMapV11() map[string]any {
	return PathMapV8()
}
func PathMapV12() map[string]any {
	return PathMapV8()
}
