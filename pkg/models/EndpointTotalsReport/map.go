package EndpointTotalsReport

func PathMapV1() map[string]any {
	return map[string]any{}
}
func PathMapV2() map[string]any {
	return PathMapV2Only()
}
func PathMapV3() map[string]any {
	return PathMapV3ToV6()
}
func PathMapV4() map[string]any {
	return PathMapV3ToV6()
}
func PathMapV5() map[string]any {
	return PathMapV3ToV6()
}
func PathMapV6() map[string]any {
	return PathMapV3ToV6()
}
func PathMapV7() map[string]any {
	return PathMapV7ToV12()
}
func PathMapV8() map[string]any {
	return PathMapV7ToV12()
}
func PathMapV9() map[string]any {
	return PathMapV7ToV12()
}
func PathMapV10() map[string]any {
	return PathMapV7ToV12()
}
func PathMapV11() map[string]any {
	return PathMapV7ToV12()
}
func PathMapV12() map[string]any {
	return PathMapV7ToV12()
}

// PathMapV2Only handles version 02 which doesn't have OrgnlBizQry or RptgSeq fields
func PathMapV2Only() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":                             "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":                           "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":                     "Pagenation.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":                "Pagenation.LastPageIndicator",
		"BkToCstmrAcctRpt.Rpt[0].Id":                                "ReportId",
		"BkToCstmrAcctRpt.Rpt[0].CreDtTm":                           "ReportCreateDateTime",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":                   "AccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.NbOfNtries": "TotalCreditEntries.NumberOfEntries",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.Sum":        "TotalCreditEntries.Sum",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.NbOfNtries": "TotalDebitEntries.NumberOfEntries",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.Sum":        "TotalDebitEntries.Sum",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TotalEntriesPerBankTransactionCode": map[string]any{
			"NbOfNtries":      "NumberOfEntries",
			"BkTxCd.Prtry.Cd": "BankTransactionCode",
		},
		"BkToCstmrAcctRpt.Rpt[0].AddtlRptInf": "AdditionalReportInfo",
	}
}

// PathMapV3ToV6 handles versions 03-06 which have OrgnlBizQry but don't have RptgSeq field
func PathMapV3ToV6() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":                             "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":                           "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":                     "Pagenation.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":                "Pagenation.LastPageIndicator",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":                 "BusinessQuery.BussinessQueryMsgId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId":               "BusinessQuery.BussinessQueryMsgNameId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm":               "BusinessQuery.BussinessQueryCreateDatetime",
		"BkToCstmrAcctRpt.Rpt[0].Id":                                "ReportId",
		"BkToCstmrAcctRpt.Rpt[0].CreDtTm":                           "ReportCreateDateTime",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":                   "AccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.NbOfNtries": "TotalCreditEntries.NumberOfEntries",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.Sum":        "TotalCreditEntries.Sum",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.NbOfNtries": "TotalDebitEntries.NumberOfEntries",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.Sum":        "TotalDebitEntries.Sum",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TotalEntriesPerBankTransactionCode": map[string]any{
			"NbOfNtries":      "NumberOfEntries",
			"BkTxCd.Prtry.Cd": "BankTransactionCode",
		},
		"BkToCstmrAcctRpt.Rpt[0].AddtlRptInf": "AdditionalReportInfo",
	}
}

// PathMapV7ToV12 handles versions 07-12 which have RptgSeq field
func PathMapV7ToV12() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":                             "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":                           "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":                     "Pagenation.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":                "Pagenation.LastPageIndicator",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":                 "BusinessQuery.BussinessQueryMsgId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId":               "BusinessQuery.BussinessQueryMsgNameId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm":               "BusinessQuery.BussinessQueryCreateDatetime",
		"BkToCstmrAcctRpt.Rpt[0].Id":                                "ReportId",
		"BkToCstmrAcctRpt.Rpt[0].RptgSeq.FrToSeq[0].FrSeq":          "Reporting.ReportingSequence.FromSeq",
		"BkToCstmrAcctRpt.Rpt[0].RptgSeq.FrToSeq[0].ToSeq":          "Reporting.ReportingSequence.ToSeq",
		"BkToCstmrAcctRpt.Rpt[0].CreDtTm":                           "ReportCreateDateTime",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":                   "AccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.NbOfNtries": "TotalCreditEntries.NumberOfEntries",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlCdtNtries.Sum":        "TotalCreditEntries.Sum",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.NbOfNtries": "TotalDebitEntries.NumberOfEntries",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlDbtNtries.Sum":        "TotalDebitEntries.Sum",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TotalEntriesPerBankTransactionCode": map[string]any{
			"NbOfNtries":      "NumberOfEntries",
			"BkTxCd.Prtry.Cd": "BankTransactionCode",
		},
		"BkToCstmrAcctRpt.Rpt[0].AddtlRptInf": "AdditionalReportInfo",
	}
}
