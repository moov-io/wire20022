package EndpointDetailsReport

import (
	"encoding/xml"
	"strconv"
	"time"

	camt052 "github.com/moov-io/fedwire20022/gen/EndpointDetailsReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"

type MessageModel struct {
	// /Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId string
	//This is the calendar date and time in New York City (Eastern Time) when the message is created by the Fedwire Funds Service application. Time is in 24-hour clock format and includes the offset against the Coordinated Universal Time (UTC).
	CreationDateTime time.Time
	//Provides details on the page number of the message.
	MessagePagination model.MessagePagenation
	// Point to point reference, as assigned by the original initiating party, to unambiguously identify the original query message.
	BussinessQueryMsgId string
	//Specifies the query message name identifier to which the message refers.
	BussinessQueryMsgNameId string
	//Date and time at which the message was created.
	BussinessQueryCreateDatetime time.Time
	//Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	ReportId model.ReportType
	//Specifies the range of identification sequence numbers, as provided in the request.
	ReportingSequence model.SequenceRange
	//Date and time at which the report was created.
	ReportCreateDateTime time.Time
	//Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId string
	//Specifies the total number and sum of debit entries.
	TotalDebitEntries model.NumberAndSumOfTransactions
	//Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	//Provides details on the entry.
	EntryDetails []model.Entry
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
	msg.doc = camt052.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var BkToCstmrAcctRpt camt052.BankToCustomerAccountReportV08
	var GrpHdr camt052.GroupHeader811
	// MessageId string
	if msg.data.MessageId != "" {
		GrpHdr.MsgId = camt052.AccountReportingFedwireFunds1(msg.data.MessageId)
	}
	// CreationDateTime time.Time
	if !isEmpty(msg.data.CreationDateTime) {
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreationDateTime)
	}
	// MessagePagination model.MessagePagenation
	if !isEmpty(msg.data.MessagePagination) {
		GrpHdr.MsgPgntn = camt052.Pagination1{
			PgNb:      camt052.Max5NumericText(msg.data.MessagePagination.PageNumber),
			LastPgInd: camt052.YesNoIndicator(msg.data.MessagePagination.LastPageIndicator),
		}
	}
	var OrgnlBizQry camt052.OriginalBusinessQuery11
	// BussinessQueryMsgId string
	if msg.data.BussinessQueryMsgId != "" {
		OrgnlBizQry.MsgId = camt052.Max35Text(msg.data.BussinessQueryMsgId)
	}
	// BussinessQueryMsgNameId string
	if msg.data.BussinessQueryMsgNameId != "" {
		OrgnlBizQry.MsgNmId = camt052.MessageNameIdentificationFRS1(msg.data.BussinessQueryMsgNameId)
	}
	// BussinessQueryCreateDatetime time.Time
	if !isEmpty(msg.data.BussinessQueryCreateDatetime) {
		OrgnlBizQry.CreDtTm = fedwire.ISODateTime(msg.data.BussinessQueryCreateDatetime)
	}
	if !isEmpty(OrgnlBizQry) {
		GrpHdr.OrgnlBizQry = OrgnlBizQry
	}
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}
	var Rpt camt052.AccountReport251
	// ReportId ReportDayType
	if msg.data.ReportId != "" {
		Rpt.Id = camt052.ReportTimingFRS1(msg.data.ReportId)
	}
	// ReportingSequence model.SequenceRange
	if !isEmpty(msg.data.ReportingSequence) {
		FrSeq, err := strconv.ParseFloat(msg.data.ReportingSequence.FromSeq, 64)
		if err != nil {
			return
		}
		ToSeq, err := strconv.ParseFloat(msg.data.ReportingSequence.ToSeq, 64)
		if err != nil {
			return
		}
		FrToSeq := camt052.SequenceRange11{
			FrSeq: camt052.XSequenceNumberFedwireFunds1(FrSeq),
			ToSeq: camt052.XSequenceNumberFedwireFunds1(ToSeq),
		}
		Rpt.RptgSeq = camt052.SequenceRange1Choice1{
			FrToSeq: &FrToSeq,
		}
	}
	// ReportCreateDateTime time.Time
	if !isEmpty(msg.data.ReportCreateDateTime) {
		CreDtTm := fedwire.ISODateTime(msg.data.ReportCreateDateTime)
		Rpt.CreDtTm = &CreDtTm
	}
	// AccountOtherId string
	if msg.data.AccountOtherId != "" {
		Othr := camt052.GenericAccountIdentification11{
			Id: camt052.EndpointIdentifierFedwireFunds1(msg.data.AccountOtherId),
		}
		Rpt.Acct = camt052.CashAccount391{
			Id: camt052.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	// TotalDebitEntries model.NumberAndSumOfTransactions
	var TxsSummry camt052.TotalTransactions61
	if !isEmpty(msg.data.TotalDebitEntries) {
		TtlDbtNtries := camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalDebitEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalDebitEntries.Sum),
		}
		TxsSummry.TtlDbtNtries = &TtlDbtNtries
	}
	// TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	if !isEmpty(msg.data.TotalEntriesPerTransactionCode) {
		var TtlNtriesPerBkTxCd []camt052.TotalsPerBankTransactionCode51
		for _, entity := range msg.data.TotalEntriesPerTransactionCode {
			bankTrans := camt052.TotalsPerBankTransactionCode51{
				NbOfNtries: camt052.Max15NumericText(entity.NumberOfEntries),
				BkTxCd: camt052.BankTransactionCodeStructure41{
					Prtry: camt052.ProprietaryBankTransactionCodeStructure11{
						Cd: camt052.BankTransactionCodeFedwireFunds11(entity.Status),
					},
				},
			}
			TtlNtriesPerBkTxCd = append(TtlNtriesPerBkTxCd, bankTrans)
		}
		if !isEmpty(TtlNtriesPerBkTxCd) {
			TxsSummry.TtlNtriesPerBkTxCd = TtlNtriesPerBkTxCd
		}
	}
	if !isEmpty(TxsSummry) {
		Rpt.TxsSummry = &TxsSummry
	}
	// EntryDetails []model.Entry
	if !isEmpty(msg.data.EntryDetails) {
		var Ntry []*camt052.ReportEntry101
		for _, entity := range msg.data.EntryDetails {
			report := ReportEntry101From(entity)
			Ntry = append(Ntry, &report)
		}
		Rpt.Ntry = Ntry
	}
	if !isEmpty(Rpt) {
		BkToCstmrAcctRpt.Rpt = Rpt
	}
	if !isEmpty(BkToCstmrAcctRpt) {
		msg.doc.BkToCstmrAcctRpt = BkToCstmrAcctRpt
	}
}
