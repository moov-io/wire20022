package EndpointDetailsReport

func pathMapV1() map[string]any {
	return map[string]any{}
}
func pathMapV2() map[string]any {
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
		"BkToCstmrAcctRpt.Rpt[0].Ntry : EntryDetails": map[string]any{
			"Amt.Value":                            "Amount.Amount",
			"Amt.Ccy":                              "Amount.Currency",
			"CdtDbtInd":                            "CreditDebitIndicator",
			"Sts":                                  "Status",
			"BkTxCd.Prtry.Cd":                      "BankTransactionCode",
			"AddtlInfInd.MsgNmId":                  "MessageNameId",
			"NtryDtls[0].TxDtls[0].Refs.MsgId":     "EntryDetails.MessageId",
			"NtryDtls[0].TxDtls[0].Refs.InstrId":   "EntryDetails.InstructionId",
			"NtryDtls[0].TxDtls[0].Refs.ClrSysRef": "EntryDetails.ClearingSystemRef",
			"NtryDtls[0].TxDtls[0].RltdAgts.IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "EntryDetails.InstructingAgent.PaymentSysCode",
			"NtryDtls[0].TxDtls[0].RltdAgts.IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId":       "EntryDetails.InstructingAgent.PaymentSysMemberId",
			"NtryDtls[0].TxDtls[0].RltdAgts.RcvgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":    "EntryDetails.InstructedAgent.PaymentSysCode",
			"NtryDtls[0].TxDtls[0].RltdAgts.RcvgAgt.FinInstnId.ClrSysMmbId.MmbId":          "EntryDetails.InstructedAgent.PaymentSysMemberId",
			"NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Tp":                                    "EntryDetails.RelatedDatesProprietary",
			"NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Dt.DtTm":                               "EntryDetails.RelatedDateTime",
		},
	}
}
func pathMapV3() map[string]any {
	return pathMapV6()
}
func pathMapV4() map[string]any {
	return pathMapV6()
}
func pathMapV5() map[string]any {
	return pathMapV6()
}
func pathMapV6() map[string]any {
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
		"BkToCstmrAcctRpt.Rpt[0].Ntry : EntryDetails": map[string]any{
			"Amt.Value":                            "Amount.Amount",
			"Amt.Ccy":                              "Amount.Currency",
			"CdtDbtInd":                            "CreditDebitIndicator",
			"Sts":                                  "Status",
			"BkTxCd.Prtry.Cd":                      "BankTransactionCode",
			"AddtlInfInd.MsgNmId":                  "MessageNameId",
			"NtryDtls[0].TxDtls[0].Refs.MsgId":     "EntryDetails.MessageId",
			"NtryDtls[0].TxDtls[0].Refs.InstrId":   "EntryDetails.InstructionId",
			"NtryDtls[0].TxDtls[0].Refs.ClrSysRef": "EntryDetails.ClearingSystemRef",
			"NtryDtls[0].TxDtls[0].Amt.Value":      "Amount.Amount",
			"NtryDtls[0].TxDtls[0].Amt.Ccy":        "Amount.Currency",
			"NtryDtls[0].TxDtls[0].CdtDbtInd":      "CreditDebitIndicator",
			"NtryDtls[0].TxDtls[0].RltdAgts.IntrmyAgt1.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "EntryDetails.InstructingAgent.PaymentSysCode",
			"NtryDtls[0].TxDtls[0].RltdAgts.IntrmyAgt1.FinInstnId.ClrSysMmbId.MmbId":       "EntryDetails.InstructingAgent.PaymentSysMemberId",
			"NtryDtls[0].TxDtls[0].RltdAgts.RcvgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd":    "EntryDetails.InstructedAgent.PaymentSysCode",
			"NtryDtls[0].TxDtls[0].RltdAgts.RcvgAgt.FinInstnId.ClrSysMmbId.MmbId":          "EntryDetails.InstructedAgent.PaymentSysMemberId",
			"NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Tp":                                    "EntryDetails.RelatedDatesProprietary",
			"NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Dt.DtTm":                               "EntryDetails.RelatedDateTime",
		},
	}
}
func pathMapV7() map[string]any {
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
		"BkToCstmrAcctRpt.Rpt[0].Ntry : EntryDetails": map[string]any{
			"Amt.Value":                            "Amount.Amount",
			"Amt.Ccy":                              "Amount.Currency",
			"CdtDbtInd":                            "CreditDebitIndicator",
			"Sts.Cd":                               "Status",
			"BkTxCd.Prtry.Cd":                      "BankTransactionCode",
			"AddtlInfInd.MsgNmId":                  "MessageNameId",
			"NtryDtls[0].TxDtls[0].Refs.MsgId":     "EntryDetails.MessageId",
			"NtryDtls[0].TxDtls[0].Refs.InstrId":   "EntryDetails.InstructionId",
			"NtryDtls[0].TxDtls[0].Refs.ClrSysRef": "EntryDetails.ClearingSystemRef",
			"NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "EntryDetails.InstructingAgent.PaymentSysCode",
			"NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "EntryDetails.InstructingAgent.PaymentSysMemberId",
			"NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "EntryDetails.InstructedAgent.PaymentSysCode",
			"NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "EntryDetails.InstructedAgent.PaymentSysMemberId",
			"NtryDtls[0].TxDtls[0].LclInstrm.Prtry":                                      "EntryDetails.LocalInstrumentChoice",
			"NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Tp":                                  "EntryDetails.RelatedDatesProprietary",
			"NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Dt.DtTm":                             "EntryDetails.RelatedDateTime",
		},
	}
}
func pathMapV8() map[string]any {
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
		"BkToCstmrAcctRpt.Rpt[0].Ntry : EntryDetails": map[string]any{
			"Amt.Value":                            "Amount.Amount",
			"Amt.Ccy":                              "Amount.Currency",
			"CdtDbtInd":                            "CreditDebitIndicator",
			"Sts.Cd":                               "Status",
			"BkTxCd.Prtry.Cd":                      "BankTransactionCode",
			"AddtlInfInd.MsgNmId":                  "MessageNameId",
			"NtryDtls[0].TxDtls[0].Refs.MsgId":     "EntryDetails.MessageId",
			"NtryDtls[0].TxDtls[0].Refs.InstrId":   "EntryDetails.InstructionId",
			"NtryDtls[0].TxDtls[0].Refs.UETR":      "EntryDetails.UniqueTransactionReference",
			"NtryDtls[0].TxDtls[0].Refs.ClrSysRef": "EntryDetails.ClearingSystemRef",
			"NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "EntryDetails.InstructingAgent.PaymentSysCode",
			"NtryDtls[0].TxDtls[0].RltdAgts.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "EntryDetails.InstructingAgent.PaymentSysMemberId",
			"NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "EntryDetails.InstructedAgent.PaymentSysCode",
			"NtryDtls[0].TxDtls[0].RltdAgts.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "EntryDetails.InstructedAgent.PaymentSysMemberId",
			"NtryDtls[0].TxDtls[0].LclInstrm.Prtry":                                      "EntryDetails.LocalInstrumentChoice",
			"NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Tp":                                  "EntryDetails.RelatedDatesProprietary",
			"NtryDtls[0].TxDtls[0].RltdDts.Prtry[0].Dt.DtTm":                             "EntryDetails.RelatedDateTime",
		},
	}
}
func pathMapV9() map[string]any {
	return pathMapV8()
}
func pathMapV10() map[string]any {
	return pathMapV8()
}
func pathMapV11() map[string]any {
	return pathMapV8()
}
func pathMapV12() map[string]any {
	return pathMapV8()
}
