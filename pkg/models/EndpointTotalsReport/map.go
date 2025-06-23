package EndpointTotalsReport

func pathMapV1() map[string]any {
	return map[string]any{}
}
func pathMapV2() map[string]any {
	return pathMapV2Only()
}
func pathMapV3() map[string]any {
	return pathMapV3ToV6()
}
func pathMapV4() map[string]any {
	return pathMapV3ToV6()
}
func pathMapV5() map[string]any {
	return pathMapV3ToV6()
}
func pathMapV6() map[string]any {
	return pathMapV3ToV6()
}
func pathMapV7() map[string]any {
	return pathMapV7ToV12()
}
func pathMapV8() map[string]any {
	return pathMapV7ToV12()
}
func pathMapV9() map[string]any {
	return pathMapV7ToV12()
}
func pathMapV10() map[string]any {
	return pathMapV7ToV12()
}
func pathMapV11() map[string]any {
	return pathMapV7ToV12()
}
func pathMapV12() map[string]any {
	return pathMapV7ToV12()
}

// pathMapV2Only handles version 02 which doesn't have OrgnlBizQry or RptgSeq fields
func pathMapV2Only() map[string]any {
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

// pathMapV3ToV6 handles versions 03-06 which have OrgnlBizQry but don't have RptgSeq field
func pathMapV3ToV6() map[string]any {
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

// pathMapV7ToV12 handles versions 07-12 which have RptgSeq field
func pathMapV7ToV12() map[string]any {
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
