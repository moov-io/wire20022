package Master

import (
	"time"

	camt052 "github.com/moov-io/fedwire20022/gen/Master_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

type MessageModel struct {
	//Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId model.CAMTReportType
	//Date and time at which the message was created.
	CreationDateTime time.Time
	//Provides details on the page number of the message.
	MessagePagination model.MessagePagenation
	//Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.
	OriginalBusinessMsgId string
	//Specifies the query message name identifier to which the message refers.
	OriginalBusinessMsgNameId string
	//Date and time at which the message was created.
	OriginalBusinessMsgCreateTime time.Time
	//Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	ReportTypeId AccountReportType
	//Date and time at which the report was created.
	ReportCreatedDate time.Time
	//Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId string
	//Identifies the parent account of the account for which the report has been issued.
	RelatedAccountOtherId string
	//Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balances []Balance
	//Provides summary information on entries.
	TransactionsSummary []TotalsPerBankTransactionCode
}

type Message struct {
	data MessageModel
	doc  camt052.Document
}

func NewMessage() Message {
	return Message{
		data: MessageModel{},
	}
}
func (msg *Message) CreateDocument() {
	msg.doc = camt052.Document{}
	var BkToCstmrAcctRpt camt052.BankToCustomerAccountReportV08
	var GrpHdr camt052.GroupHeader811
	if msg.data.MessageId != "" {
		GrpHdr.MsgId = camt052.AccountReportingFedwireFunds1(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreationDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreationDateTime)
	}
	if !isEmpty(msg.data.MessagePagination) {
		GrpHdr.MsgPgntn = camt052.Pagination1{
			PgNb:      camt052.Max5NumericText(msg.data.MessagePagination.PageNumber),
			LastPgInd: camt052.YesNoIndicator(msg.data.MessagePagination.LastPageIndicator),
		}
	}
	var OrgnlBizQry camt052.OriginalBusinessQuery11
	if msg.data.OriginalBusinessMsgId != "" {
		OrgnlBizQry.MsgId = camt052.Max35Text(msg.data.OriginalBusinessMsgId)
	}
	if msg.data.OriginalBusinessMsgNameId != "" {
		OrgnlBizQry.MsgNmId = camt052.MessageNameIdentificationFRS1(msg.data.OriginalBusinessMsgNameId)
	}
	if !isEmpty(msg.data.OriginalBusinessMsgCreateTime) {
		OrgnlBizQry.CreDtTm = fedwire.ISODateTime(msg.data.OriginalBusinessMsgCreateTime)
	}
	if !isEmpty(OrgnlBizQry) {
		GrpHdr.OrgnlBizQry = &OrgnlBizQry
	}
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}
	var Rpt camt052.AccountReport251
	if msg.data.ReportTypeId != "" {
		Rpt.Id = camt052.BalanceReportFRS1(msg.data.ReportTypeId)
	}
	if !isEmpty(msg.data.ReportCreatedDate) {
		Rpt.CreDtTm = fedwire.ISODateTime(msg.data.ReportCreatedDate)
	}
	if msg.data.AccountOtherId != "" {
		Othr := camt052.GenericAccountIdentification11{
			Id: camt052.RoutingNumberFRS1(msg.data.AccountOtherId),
		}
		Rpt.Acct = camt052.CashAccount391{
			Id: camt052.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	if msg.data.RelatedAccountOtherId != "" {
		Othr := camt052.GenericAccountIdentification11{
			Id: camt052.RoutingNumberFRS1(msg.data.RelatedAccountOtherId),
		}
		Rpt.RltdAcct = camt052.CashAccount381{
			Id: camt052.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	var Bal []camt052.CashBalance81
	if !isEmpty(Bal) {
		for _, item := range msg.data.Balances {
			line := CashBalance81From(item)
			Bal = append(Bal, line)
		}
	}
	if !isEmpty(Bal) {
		Rpt.Bal = Bal
	}

	var TxsSummry camt052.TotalTransactions61
	var TtlNtriesPerBkTxCd []camt052.TotalsPerBankTransactionCode51
	for _, item := range msg.data.TransactionsSummary {
		code := TotalsPerBankTransactionCode51From(item)
		TtlNtriesPerBkTxCd = append(TtlNtriesPerBkTxCd, code)
	}
	if !isEmpty(TtlNtriesPerBkTxCd) {
		TxsSummry.TtlNtriesPerBkTxCd = TtlNtriesPerBkTxCd
	}

	if !isEmpty(TxsSummry) {
		Rpt.TxsSummry = TxsSummry
	}
	if !isEmpty(Rpt) {
		BkToCstmrAcctRpt.Rpt = Rpt
	}
	if !isEmpty(BkToCstmrAcctRpt) {
		msg.doc.BkToCstmrAcctRpt = BkToCstmrAcctRpt
	}
}
