package EndpointDetailsReport

import (
	"encoding/xml"
	"fmt"
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
	TotalCreditEntries model.NumberAndSumOfTransactions
	TotalDebitEntries  model.NumberAndSumOfTransactions
	//Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	//Provides details on the entry.
	EntryDetails []model.Entry
}
type Message struct {
	data MessageModel
	doc  camt052.Document
}

/*
NewMessage creates a new Message instance with optional XML initialization.

Parameters:
  - filepath: File path to XML (optional)
    If provided, loads and parses XML from specified path

Returns:
  - Message: Initialized message structure
  - error: File read or XML parsing errors (if XML path provided)

Behavior:
  - Without arguments: Returns empty Message with default MessageModel
  - With XML path: Loads file, parses XML into message.doc
*/
func NewMessage(filepath string) (Message, error) {
	msg := Message{data: MessageModel{}} // Initialize with zero value

	if filepath == "" {
		return msg, nil // Return early for empty filepath
	}

	// Read and validate file
	data, err := model.ReadXMLFile(filepath)
	if err != nil {
		return msg, fmt.Errorf("file read error: %w", err)
	}

	// Handle empty XML data
	if len(data) == 0 {
		return msg, fmt.Errorf("empty XML file: %s", filepath)
	}

	// Parse XML with structural validation
	if err := xml.Unmarshal(data, &msg.doc); err != nil {
		return msg, fmt.Errorf("XML parse error: %w", err)
	}

	return msg, nil
}
func (msg *Message) CreateDocument() *model.ValidateError {
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
		err := camt052.Max35Text(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = camt052.AccountReportingFedwireFunds1(msg.data.MessageId)
	}
	// CreationDateTime time.Time
	if !isEmpty(msg.data.CreationDateTime) {
		err := fedwire.ISODateTime(msg.data.CreationDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreationDateTime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreationDateTime)
	}
	// MessagePagination model.MessagePagenation
	if !isEmpty(msg.data.MessagePagination) {
		err := camt052.Max5NumericText(msg.data.MessagePagination.PageNumber).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PageNumber",
				Message:   err.Error(),
			}
			vErr.InsertPath("MessagePagination")
			return &vErr
		}
		err = camt052.YesNoIndicator(msg.data.MessagePagination.LastPageIndicator).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "LastPageIndicator",
				Message:   err.Error(),
			}
			vErr.InsertPath("MessagePagination")
			return &vErr
		}
		GrpHdr.MsgPgntn = camt052.Pagination1{
			PgNb:      camt052.Max5NumericText(msg.data.MessagePagination.PageNumber),
			LastPgInd: camt052.YesNoIndicator(msg.data.MessagePagination.LastPageIndicator),
		}
	}
	var OrgnlBizQry camt052.OriginalBusinessQuery11
	// BussinessQueryMsgId string
	if msg.data.BussinessQueryMsgId != "" {
		err := camt052.Max35Text(msg.data.BussinessQueryMsgId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BussinessQueryMsgId",
				Message:   err.Error(),
			}
		}
		OrgnlBizQry.MsgId = camt052.Max35Text(msg.data.BussinessQueryMsgId)
	}
	// BussinessQueryMsgNameId string
	if msg.data.BussinessQueryMsgNameId != "" {
		err := camt052.MessageNameIdentificationFRS1(msg.data.BussinessQueryMsgNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BussinessQueryMsgNameId",
				Message:   err.Error(),
			}
		}
		OrgnlBizQry.MsgNmId = camt052.MessageNameIdentificationFRS1(msg.data.BussinessQueryMsgNameId)
	}
	// BussinessQueryCreateDatetime time.Time
	if !isEmpty(msg.data.BussinessQueryCreateDatetime) {
		err := fedwire.ISODateTime(msg.data.BussinessQueryCreateDatetime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BussinessQueryCreateDatetime",
				Message:   err.Error(),
			}
		}
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
		err := camt052.ReportTimingFRS1(msg.data.ReportId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportId",
				Message:   err.Error(),
			}
		}
		Rpt.Id = camt052.ReportTimingFRS1(msg.data.ReportId)
	}
	// ReportingSequence model.SequenceRange
	if !isEmpty(msg.data.ReportingSequence) {
		FrSeq, err := strconv.ParseFloat(msg.data.ReportingSequence.FromSeq, 64)
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportingSequence.FromSeq",
				Message:   err.Error(),
			}
		}
		ToSeq, err := strconv.ParseFloat(msg.data.ReportingSequence.ToSeq, 64)
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportingSequence.ToSeq",
				Message:   err.Error(),
			}
		}
		// err = camt052.XSequenceNumberFedwireFunds1(FrSeq).Validate()
		// if err != nil {
		// 	return &model.ValidateError{
		// 		ParamName: "ReportingSequence.FromSeq",
		// 		Message:   err.Error(),
		// 	}
		// }
		// err = camt052.XSequenceNumberFedwireFunds1(ToSeq).Validate()
		// if err != nil {
		// 	return &model.ValidateError{
		// 		ParamName: "ReportingSequence.ToSeq",
		// 		Message:   err.Error(),
		// 	}
		// }
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
		err := fedwire.ISODateTime(msg.data.ReportCreateDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportCreateDateTime",
				Message:   err.Error(),
			}
		}
		CreDtTm := fedwire.ISODateTime(msg.data.ReportCreateDateTime)
		Rpt.CreDtTm = &CreDtTm
	}
	// AccountOtherId string
	if msg.data.AccountOtherId != "" {
		err := camt052.EndpointIdentifierFedwireFunds1(msg.data.AccountOtherId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountOtherId",
				Message:   err.Error(),
			}
		}
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
		err := camt052.Max15NumericText(msg.data.TotalDebitEntries.NumberOfEntries).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalDebitEntries.NumberOfEntries",
				Message:   err.Error(),
			}
		}
		err = camt052.DecimalNumber(msg.data.TotalDebitEntries.Sum).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalDebitEntries.Sum",
				Message:   err.Error(),
			}
		}
		TtlDbtNtries := camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalDebitEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalDebitEntries.Sum),
		}
		TxsSummry.TtlDbtNtries = &TtlDbtNtries
	}
	if !isEmpty(msg.data.TotalCreditEntries) {
		err := camt052.Max15NumericText(msg.data.TotalCreditEntries.NumberOfEntries).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalCreditEntries.NumberOfEntries",
				Message:   err.Error(),
			}
		}
		err = camt052.DecimalNumber(msg.data.TotalCreditEntries.Sum).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalCreditEntries.Sum",
				Message:   err.Error(),
			}
		}
		TtlCdtNtries := camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalCreditEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalCreditEntries.Sum),
		}
		TxsSummry.TtlCdtNtries = &TtlCdtNtries
	}
	// TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	if !isEmpty(msg.data.TotalEntriesPerTransactionCode) {
		var TtlNtriesPerBkTxCd []camt052.TotalsPerBankTransactionCode51
		for _, entity := range msg.data.TotalEntriesPerTransactionCode {
			err := camt052.Max15NumericText(entity.NumberOfEntries).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "TotalEntriesPerTransactionCode.NumberOfEntries",
					Message:   err.Error(),
				}
			}
			err = camt052.BankTransactionCodeFedwireFunds11(entity.Status).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "TotalEntriesPerTransactionCode.Status",
					Message:   err.Error(),
				}
			}
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
			report, vErr := ReportEntry101From(entity)
			if vErr != nil {
				vErr.InsertPath("EntryDetails")
				return vErr
			}
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
	return nil
}
