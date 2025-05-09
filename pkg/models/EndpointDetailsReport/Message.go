package EndpointDetailsReport

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
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
	Data   MessageModel
	Doc    camt052.Document
	Helper MessageHelper
}

func (msg *Message) GetDataModel() interface{} {
	return &msg.Data
}
func (msg *Message) GetDocument() interface{} {
	return &msg.Doc
}
func (msg *Message) GetHelper() interface{} {
	return &msg.Helper
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
  - With XML path: Loads file, parses XML into message.Doc
*/
func NewMessage(filepath string) (*Message, error) {
	msg := Message{Data: MessageModel{}} // Initialize with zero value
	msg.Helper = BuildMessageHelper()

	if filepath == "" {
		return &msg, nil // Return early for empty filepath
	}

	// Read and validate file
	data, err := model.ReadXMLFile(filepath)
	if err != nil {
		return &msg, fmt.Errorf("file read error: %w", err)
	}

	// Handle empty XML data
	if len(data) == 0 {
		return &msg, fmt.Errorf("empty XML file: %s", filepath)
	}

	// Parse XML with structural validation
	if err := xml.Unmarshal(data, &msg.Doc); err != nil {
		return &msg, fmt.Errorf("XML parse error: %w", err)
	}

	return &msg, nil
}
func (msg *Message) ValidateRequiredFields() *model.ValidateError {
	// Initialize the RequireError object
	var ParamNames []string

	// Check required fields and append missing ones to ParamNames
	if msg.Data.MessageId == "" {
		ParamNames = append(ParamNames, "MessageId")
	}
	if msg.Data.CreationDateTime.IsZero() {
		ParamNames = append(ParamNames, "CreationDateTime")
	}
	if isEmpty(msg.Data.MessagePagination) {
		ParamNames = append(ParamNames, "MessagePagination")
	}
	if msg.Data.BussinessQueryMsgId == "" {
		ParamNames = append(ParamNames, "BussinessQueryMsgId")
	}
	if msg.Data.BussinessQueryMsgNameId == "" {
		ParamNames = append(ParamNames, "BussinessQueryMsgNameId")
	}
	if msg.Data.BussinessQueryCreateDatetime.IsZero() {
		ParamNames = append(ParamNames, "BussinessQueryCreateDatetime")
	}
	if msg.Data.ReportId == "" {
		ParamNames = append(ParamNames, "ReportId")
	}
	if isEmpty(msg.Data.ReportingSequence) {
		ParamNames = append(ParamNames, "ReportingSequence")
	}
	if msg.Data.AccountOtherId == "" {
		ParamNames = append(ParamNames, "AccountOtherId")
	}
	// Return nil if no required fields are missing
	if len(ParamNames) == 0 {
		return nil
	}
	return &model.ValidateError{
		ParamName: "RequiredFields",
		Message:   strings.Join(ParamNames, ", "),
	}
}
func (msg *Message) CreateDocument() *model.ValidateError {
	requireErr := msg.ValidateRequiredFields()
	if requireErr != nil {
		return requireErr
	}
	msg.Doc = camt052.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var BkToCstmrAcctRpt camt052.BankToCustomerAccountReportV08
	var GrpHdr camt052.GroupHeader811
	// MessageId string
	if msg.Data.MessageId != "" {
		err := camt052.Max35Text(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = camt052.AccountReportingFedwireFunds1(msg.Data.MessageId)
	}
	// CreationDateTime time.Time
	if !isEmpty(msg.Data.CreationDateTime) {
		err := fedwire.ISODateTime(msg.Data.CreationDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreationDateTime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.Data.CreationDateTime)
	}
	// MessagePagination model.MessagePagenation
	if !isEmpty(msg.Data.MessagePagination) {
		err := camt052.Max5NumericText(msg.Data.MessagePagination.PageNumber).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "PageNumber",
				Message:   err.Error(),
			}
			vErr.InsertPath("MessagePagination")
			return &vErr
		}
		err = camt052.YesNoIndicator(msg.Data.MessagePagination.LastPageIndicator).Validate()
		if err != nil {
			vErr := model.ValidateError{
				ParamName: "LastPageIndicator",
				Message:   err.Error(),
			}
			vErr.InsertPath("MessagePagination")
			return &vErr
		}
		GrpHdr.MsgPgntn = camt052.Pagination1{
			PgNb:      camt052.Max5NumericText(msg.Data.MessagePagination.PageNumber),
			LastPgInd: camt052.YesNoIndicator(msg.Data.MessagePagination.LastPageIndicator),
		}
	}
	var OrgnlBizQry camt052.OriginalBusinessQuery11
	// BussinessQueryMsgId string
	if msg.Data.BussinessQueryMsgId != "" {
		err := camt052.Max35Text(msg.Data.BussinessQueryMsgId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BussinessQueryMsgId",
				Message:   err.Error(),
			}
		}
		OrgnlBizQry.MsgId = camt052.Max35Text(msg.Data.BussinessQueryMsgId)
	}
	// BussinessQueryMsgNameId string
	if msg.Data.BussinessQueryMsgNameId != "" {
		err := camt052.MessageNameIdentificationFRS1(msg.Data.BussinessQueryMsgNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BussinessQueryMsgNameId",
				Message:   err.Error(),
			}
		}
		OrgnlBizQry.MsgNmId = camt052.MessageNameIdentificationFRS1(msg.Data.BussinessQueryMsgNameId)
	}
	// BussinessQueryCreateDatetime time.Time
	if !isEmpty(msg.Data.BussinessQueryCreateDatetime) {
		err := fedwire.ISODateTime(msg.Data.BussinessQueryCreateDatetime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "BussinessQueryCreateDatetime",
				Message:   err.Error(),
			}
		}
		OrgnlBizQry.CreDtTm = fedwire.ISODateTime(msg.Data.BussinessQueryCreateDatetime)
	}
	if !isEmpty(OrgnlBizQry) {
		GrpHdr.OrgnlBizQry = OrgnlBizQry
	}
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}
	var Rpt camt052.AccountReport251
	// ReportId ReportDayType
	if msg.Data.ReportId != "" {
		err := camt052.ReportTimingFRS1(msg.Data.ReportId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportId",
				Message:   err.Error(),
			}
		}
		Rpt.Id = camt052.ReportTimingFRS1(msg.Data.ReportId)
	}
	// ReportingSequence model.SequenceRange
	if !isEmpty(msg.Data.ReportingSequence) {
		FrSeq, err := strconv.ParseFloat(msg.Data.ReportingSequence.FromSeq, 64)
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportingSequence.FromSeq",
				Message:   err.Error(),
			}
		}
		ToSeq, err := strconv.ParseFloat(msg.Data.ReportingSequence.ToSeq, 64)
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
	if !isEmpty(msg.Data.ReportCreateDateTime) {
		err := fedwire.ISODateTime(msg.Data.ReportCreateDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportCreateDateTime",
				Message:   err.Error(),
			}
		}
		CreDtTm := fedwire.ISODateTime(msg.Data.ReportCreateDateTime)
		Rpt.CreDtTm = &CreDtTm
	}
	// AccountOtherId string
	if msg.Data.AccountOtherId != "" {
		err := camt052.EndpointIdentifierFedwireFunds1(msg.Data.AccountOtherId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountOtherId",
				Message:   err.Error(),
			}
		}
		Othr := camt052.GenericAccountIdentification11{
			Id: camt052.EndpointIdentifierFedwireFunds1(msg.Data.AccountOtherId),
		}
		Rpt.Acct = camt052.CashAccount391{
			Id: camt052.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	// TotalDebitEntries model.NumberAndSumOfTransactions
	var TxsSummry camt052.TotalTransactions61
	if !isEmpty(msg.Data.TotalDebitEntries) {
		err := camt052.Max15NumericText(msg.Data.TotalDebitEntries.NumberOfEntries).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalDebitEntries.NumberOfEntries",
				Message:   err.Error(),
			}
		}
		err = camt052.DecimalNumber(msg.Data.TotalDebitEntries.Sum).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalDebitEntries.Sum",
				Message:   err.Error(),
			}
		}
		TtlDbtNtries := camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.Data.TotalDebitEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.Data.TotalDebitEntries.Sum),
		}
		TxsSummry.TtlDbtNtries = &TtlDbtNtries
	}
	if !isEmpty(msg.Data.TotalCreditEntries) {
		err := camt052.Max15NumericText(msg.Data.TotalCreditEntries.NumberOfEntries).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalCreditEntries.NumberOfEntries",
				Message:   err.Error(),
			}
		}
		err = camt052.DecimalNumber(msg.Data.TotalCreditEntries.Sum).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalCreditEntries.Sum",
				Message:   err.Error(),
			}
		}
		TtlCdtNtries := camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.Data.TotalCreditEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.Data.TotalCreditEntries.Sum),
		}
		TxsSummry.TtlCdtNtries = &TtlCdtNtries
	}
	// TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	if !isEmpty(msg.Data.TotalEntriesPerTransactionCode) {
		var TtlNtriesPerBkTxCd []camt052.TotalsPerBankTransactionCode51
		for _, entity := range msg.Data.TotalEntriesPerTransactionCode {
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
	if !isEmpty(msg.Data.EntryDetails) {
		var Ntry []*camt052.ReportEntry101
		for _, entity := range msg.Data.EntryDetails {
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
		msg.Doc.BkToCstmrAcctRpt = BkToCstmrAcctRpt
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc.BkToCstmrAcctRpt) {
		if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr) {
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgId) {
				msg.Data.MessageId = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgId)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.CreDtTm) {
				msg.Data.CreationDateTime = time.Time(msg.Doc.BkToCstmrAcctRpt.GrpHdr.CreDtTm)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb) {
					msg.Data.MessagePagination.PageNumber = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd) {
					msg.Data.MessagePagination.LastPageIndicator = bool(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd)
				}
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId) {
					msg.Data.BussinessQueryMsgId = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId) {
					msg.Data.BussinessQueryMsgNameId = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm) {
					msg.Data.BussinessQueryCreateDatetime = time.Time(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm)
				}
			}
		}
		if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt) {
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Id) {
				msg.Data.ReportId = model.ReportType(msg.Doc.BkToCstmrAcctRpt.Rpt.Id)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.RptgSeq) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.RptgSeq.FrToSeq) {
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.RptgSeq.FrToSeq.FrSeq) {
						msg.Data.ReportingSequence.FromSeq = strconv.FormatFloat(float64(msg.Doc.BkToCstmrAcctRpt.Rpt.RptgSeq.FrToSeq.FrSeq), 'f', -1, 64)
					}
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.RptgSeq.FrToSeq.ToSeq) {
						msg.Data.ReportingSequence.ToSeq = strconv.FormatFloat(float64(msg.Doc.BkToCstmrAcctRpt.Rpt.RptgSeq.FrToSeq.ToSeq), 'f', -1, 64)
					}
				}
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.CreDtTm) {
				msg.Data.ReportCreateDateTime = time.Time(*msg.Doc.BkToCstmrAcctRpt.Rpt.CreDtTm)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr) {
					msg.Data.AccountOtherId = string(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr.Id)
				}
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlDbtNtries) {
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlDbtNtries.NbOfNtries) {
						msg.Data.TotalDebitEntries.NumberOfEntries = string(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlDbtNtries.NbOfNtries)
					}
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlDbtNtries.Sum) {
						msg.Data.TotalDebitEntries.Sum = float64(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlDbtNtries.Sum)
					}
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries) {
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries.NbOfNtries) {
						msg.Data.TotalCreditEntries.NumberOfEntries = string(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries.NbOfNtries)
					}
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries.Sum) {
						msg.Data.TotalCreditEntries.Sum = float64(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries.Sum)
					}
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd) {
					if len(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd) > 0 {
						msg.Data.TotalEntriesPerTransactionCode = make([]model.NumberAndStatusOfTransactions, len(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd))
						for i, entity := range msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd {
							if !isEmpty(entity.NbOfNtries) && !isEmpty(entity.BkTxCd.Prtry) {
								msg.Data.TotalEntriesPerTransactionCode[i] = model.NumberAndStatusOfTransactions{
									NumberOfEntries: string(entity.NbOfNtries),
									Status:          model.TransactionStatusCode(entity.BkTxCd.Prtry.Cd),
								}
							}
						}
					}
				}
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Ntry) {
				if len(msg.Doc.BkToCstmrAcctRpt.Rpt.Ntry) > 0 {
					msg.Data.EntryDetails = make([]model.Entry, len(msg.Doc.BkToCstmrAcctRpt.Rpt.Ntry))
					for i, entity := range msg.Doc.BkToCstmrAcctRpt.Rpt.Ntry {
						msg.Data.EntryDetails[i] = ReportEntry101To(*entity)
					}
				}
			}
		}

	}
	return nil
}
