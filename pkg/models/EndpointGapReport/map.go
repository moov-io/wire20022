package EndpointGapReport

func PathMapV1() map[string]string {
	return PathMapV8()
}
func PathMapV2() map[string]string {
	return PathMapV8()
}
func PathMapV3() map[string]string {
	return PathMapV8()
}
func PathMapV4() map[string]string {
	return PathMapV8()
}
func PathMapV5() map[string]string {
	return PathMapV8()
}
func PathMapV6() map[string]string {
	return PathMapV8()
}
func PathMapV7() map[string]string {
	return PathMapV8()
}
func PathMapV8() map[string]string {
	return map[string]string{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":              "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":            "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":      "Pagenation.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd": "Pagenation.LastPageIndicator",
		"BkToCstmrAcctRpt.Rpt[0].Id":                 "ReportId",
		"BkToCstmrAcctRpt.Rpt[0].CreDtTm":            "ReportCreateDateTime",
		"BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id":    "AccountOtherId",
		"BkToCstmrAcctRpt.Rpt[0].AddtlRptInf":        "AdditionalReportInfo",
	}
}
func PathMapV9() map[string]string {
	return PathMapV8()
}
func PathMapV10() map[string]string {
	return PathMapV8()
}
func PathMapV11() map[string]string {
	return PathMapV8()
}
func PathMapV12() map[string]string {
	return PathMapV8()
}
