package Master

func PathMapV2() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":               "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":       "MessagePagination.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":  "MessagePagination.LastPageIndicator",
		"BkToCstmrAcctRpt.Rpt[0].Id":                  "ReportTypeId",
		"BkToCstmrAcctRpt.Rpt[0].CreDtTm":             "ReportCreatedDate",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":     "AccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Tp.Prtry":       "AccountType",
		"BkToCstmrAcctRpt.Rpt[0].RltdAcct.Id.Othr.Id": "RelatedAccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TransactionsSummary": map[string]string{
			"TtlNetNtryAmt":   "TotalNetEntryAmount",
			"CdtDbtInd":       "CreditDebitIndicator",
			"BkTxCd.Prtry.Cd": "BankTransactionCode",
		},
		"BkToCstmrAcctRpt.Rpt[0].Bal : Balances": map[string]any{
			"Tp.CdOrPrtry.Prtry": "BalanceTypeId",
			"Amt.Value":          "Amount.Amount",
			"Amt.Ccy":            "Amount.Currency",
			"CdtDbtInd":          "CreditDebitIndicator",
			"Dt.DtTm":            "DateTime",
		},
	}
}
func PathMapV3() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":               "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":       "MessagePagination.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":  "MessagePagination.LastPageIndicator",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":   "OriginalBusinessMsgId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId": "OriginalBusinessMsgNameId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm": "OriginalBusinessMsgCreateTime",
		"BkToCstmrAcctRpt.Rpt[0].Id":                  "ReportTypeId",
		"BkToCstmrAcctRpt.Rpt[0].CreDtTm":             "ReportCreatedDate",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":     "AccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Tp.Prtry":       "AccountType",
		"BkToCstmrAcctRpt.Rpt[0].RltdAcct.Id.Othr.Id": "RelatedAccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TransactionsSummary": map[string]string{
			"TtlNetNtryAmt":   "TotalNetEntryAmount",
			"CdtDbtInd":       "CreditDebitIndicator",
			"BkTxCd.Prtry.Cd": "BankTransactionCode",
		},
		"BkToCstmrAcctRpt.Rpt[0].Bal : Balances": map[string]any{
			"Tp.CdOrPrtry.Prtry": "BalanceTypeId",
			"Amt.Value":          "Amount.Amount",
			"Amt.Ccy":            "Amount.Currency",
			"CdtDbtInd":          "CreditDebitIndicator",
			"Dt.DtTm":            "DateTime",
		},
	}
}
func PathMapV4() map[string]any {
	return PathMapV6()
}
func PathMapV5() map[string]any {
	return PathMapV6()
}
func PathMapV6() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":               "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":       "MessagePagination.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":  "MessagePagination.LastPageIndicator",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":   "OriginalBusinessMsgId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId": "OriginalBusinessMsgNameId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm": "OriginalBusinessMsgCreateTime",
		"BkToCstmrAcctRpt.Rpt[0].Id":                  "ReportTypeId",
		"BkToCstmrAcctRpt.Rpt[0].CreDtTm":             "ReportCreatedDate",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":     "AccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Tp.Prtry":       "AccountType",
		"BkToCstmrAcctRpt.Rpt[0].RltdAcct.Id.Othr.Id": "RelatedAccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TransactionsSummary": map[string]string{
			"TtlNetNtry.Amt":       "TotalNetEntryAmount",
			"TtlNetNtry.CdtDbtInd": "CreditDebitIndicator",
			"BkTxCd.Prtry.Cd":      "BankTransactionCode",
		},
		"BkToCstmrAcctRpt.Rpt[0].Bal : Balances": map[string]any{
			"Tp.CdOrPrtry.Prtry": "BalanceTypeId",
			"Amt.Value":          "Amount.Amount",
			"Amt.Ccy":            "Amount.Currency",
			"CdtDbtInd":          "CreditDebitIndicator",
			"Dt.DtTm":            "DateTime",
		},
	}
}
func PathMapV7() map[string]any {
	return PathMapV8()
}
func PathMapV8() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":               "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":       "MessagePagination.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":  "MessagePagination.LastPageIndicator",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":   "OriginalBusinessMsgId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId": "OriginalBusinessMsgNameId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm": "OriginalBusinessMsgCreateTime",
		"BkToCstmrAcctRpt.Rpt[0].Id":                  "ReportTypeId",
		"BkToCstmrAcctRpt.Rpt[0].CreDtTm":             "ReportCreatedDate",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":     "AccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Tp.Prtry":       "AccountType",
		"BkToCstmrAcctRpt.Rpt[0].RltdAcct.Id.Othr.Id": "RelatedAccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TransactionsSummary": map[string]string{
			"TtlNetNtry.Amt":       "TotalNetEntryAmount",
			"TtlNetNtry.CdtDbtInd": "CreditDebitIndicator",
			"CdtNtries.NbOfNtries": "CreditEntries.NumberOfEntries",
			"CdtNtries.Sum":        "CreditEntries.Sum",
			"DbtNtries.NbOfNtries": "DebitEntries.NumberOfEntries",
			"DbtNtries.Sum":        "DebitEntries.Sum",
			"BkTxCd.Prtry.Cd":      "BankTransactionCode",
			"Dt.DtTm":              "Date",
		},
		"BkToCstmrAcctRpt.Rpt[0].Bal : Balances": map[string]any{
			"Tp.CdOrPrtry.Prtry": "BalanceTypeId",
			"Amt.Value":          "Amount.Amount",
			"Amt.Ccy":            "Amount.Currency",
			"CdtDbtInd":          "CreditDebitIndicator",
			"Dt.DtTm":            "DateTime",
			"CdtLine : CdtLines": map[string]string{
				"Incl":      "Included",
				"Tp.Prtry":  "Type",
				"Amt.Value": "Amount.Amount",
				"Amt.Ccy":   "Amount.Currency",
				"Dt.DtTm":   "DateTime",
			},
		},
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
