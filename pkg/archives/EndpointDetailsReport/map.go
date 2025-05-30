package EndpointDetailsReport

func PathMapV1() map[string]any {
	return map[string]any{}
}
func PathMapV2() map[string]any {
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
		"BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd : TotalEntriesPerBankTransactionCode": map[string]string{
			"NbOfNtries":      "NumberOfEntries",
			"BkTxCd.Prtry.Cd": "BankTransactionCode",
		},
		"BkToCstmrAcctRpt.Rpt[0].Ntry : EntryDetails": map[string]string{
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
func PathMapV3() map[string]any {
	return PathMapV6()
}
func PathMapV4() map[string]any {
	return PathMapV6()
}
func PathMapV5() map[string]any {
	return PathMapV6()
}
func PathMapV6() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":                             "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":                           "CreatedDateTime",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb":                     "Pagenation.PageNumber",
		"BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd":                "Pagenation.LastPageIndicator",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId":                 "BussinessQueryMsgId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId":               "BussinessQueryMsgNameId",
		"BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm":               "BussinessQueryCreateDatetime",
		"BkToCstmrAcctRpt.Rpt[0].Id":                                "ReportId",
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
		"BkToCstmrAcctRpt.Rpt[0].Ntry : EntryDetails": map[string]string{
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
func PathMapV7() map[string]any {
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
		"BkToCstmrAcctRpt.Rpt[0].Ntry : EntryDetails": map[string]string{
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
		"BkToCstmrAcctRpt.Rpt[0].Ntry : EntryDetails": map[string]string{
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
