package Master

import (
	"time"

	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_02"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type Message struct {
	Data MessageModel
	Doc  camt_052_001_02.Document
}

func Convert() {
	msg := Message{}

	msg.Data.MessageId = Archive.CAMTReportType(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgId)
	msg.Data.CreationDateTime = time.Time(msg.Doc.BkToCstmrAcctRpt.GrpHdr.CreDtTm)
	msg.Data.MessagePagination.PageNumber = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb)
	msg.Data.MessagePagination.LastPageIndicator = bool(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd)
	// msg.Data.OriginalBusinessMsgId = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId)
	// msg.Data.OriginalBusinessMsgNameId = string(*msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId)
	// msg.Data.OriginalBusinessMsgCreateTime = time.Time(*msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm)
	msg.Data.ReportTypeId = Archive.AccountReportType(msg.Doc.BkToCstmrAcctRpt.Rpt[0].Id)
	msg.Data.ReportCreatedDate = time.Time(msg.Doc.BkToCstmrAcctRpt.Rpt[0].CreDtTm)
	msg.Data.AccountOtherId = string(msg.Doc.BkToCstmrAcctRpt.Rpt[0].Acct.Id.Othr.Id)
	msg.Data.AccountType = string(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].Acct.Tp.Prtry)
	msg.Data.RelatedAccountOtherId = string(msg.Doc.BkToCstmrAcctRpt.Rpt[0].RltdAcct.Id.Othr.Id)
	//msg.Data.TransactionsSummary msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd
	msg.Data.TransactionsSummary[0].TotalNetEntryAmount = float64(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].TtlNetNtryAmt)
	msg.Data.TransactionsSummary[0].CreditDebitIndicator = Archive.CdtDbtInd(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].CdtDbtInd)
	// msg.Data.TransactionsSummary[0].CreditEntries.NumberOfEntries = string(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].CdtNtries.NbOfNtries)
	// msg.Data.TransactionsSummary[0].CreditEntries.Sum = float64(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].CdtNtries.Sum)
	// msg.Data.TransactionsSummary[0].DebitEntries.NumberOfEntries = string(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].DbtNtries.NbOfNtries)
	// msg.Data.TransactionsSummary[0].DebitEntries.Sum = float64(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].DbtNtries.Sum)
	msg.Data.TransactionsSummary[0].BankTransactionCode = Archive.TransactionCode(msg.Doc.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].BkTxCd.Prtry.Cd)
	// msg.Data.TransactionsSummary[0].Date = time.Time(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].TxsSummry.TtlNtriesPerBkTxCd[0].Dt.DtTm)
	//msg.Data.Balances msg.Doc.BkToCstmrAcctRpt.Rpt.Bal
	msg.Data.Balances[0].BalanceTypeId = Archive.BalanceType(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].Tp.CdOrPrtry.Prtry)
	msg.Data.Balances[0].Amount.Amount = float64(msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].Amt.Value)
	msg.Data.Balances[0].Amount.Currency = string(msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].Amt.Ccy)
	msg.Data.Balances[0].CreditDebitIndicator = Archive.CdtDbtInd(msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].CdtDbtInd)
	msg.Data.Balances[0].DateTime = time.Time(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].Dt.DtTm)
	////msg.Data.Balances[0].CdtLines msg.Doc.BkToCstmrAcctRpt.Rpt.Bal[0].CdtLine
	// msg.Data.Balances[0].CdtLines[0].Included = bool(msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].CdtLine[0].Incl)
	// msg.Data.Balances[0].CdtLines[0].Type = Archive.CreditLineType(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].CdtLine[0].Tp.Prtry)
	// msg.Data.Balances[0].CdtLines[0].Amount.Amount = float64(msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].CdtLine[0].Amt.Value)
	// msg.Data.Balances[0].CdtLines[0].Amount.Currency = string(msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].CdtLine[0].Amt.Ccy)
	// msg.Data.Balances[0].CdtLines[0].DateTime = time.Time(*msg.Doc.BkToCstmrAcctRpt.Rpt[0].Bal[0].CdtLine[0].Dt.DtTm)
}

func PathMapV2() map[string]any {
	return map[string]any{
		"BkToCstmrAcctRpt.GrpHdr.MsgId":               "MessageId",
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreationDateTime",
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
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreationDateTime",
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
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreationDateTime",
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
		"BkToCstmrAcctRpt.GrpHdr.CreDtTm":             "CreationDateTime",
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
