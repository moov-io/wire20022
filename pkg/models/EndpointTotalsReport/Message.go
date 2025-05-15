package EndpointTotalsReport

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	camt052 "github.com/moov-io/fedwire20022/gen/EndpointTotalsReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"

type MessageModel struct {
	//Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId model.CAMTReportType
	// Date and time at which the message was created.
	CreatedDateTime time.Time
	//Provides details on the page number of the message.
	MessagePagination model.MessagePagenation
	//Unique identification, as assigned by the account servicer, to unambiguously identify the account report.
	ReportId model.ReportType
	//Date and time at which the report was created.
	ReportCreateDateTime time.Time
	// /Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId string
	//Specifies the total number and sum of credit entries.
	TotalCreditEntries model.NumberAndSumOfTransactions
	TotalDebitEntries  model.NumberAndSumOfTransactions

	//Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerTransactionCode []model.NumberAndStatusOfTransactions
	//Further details of the account report.
	AdditionalReportInfo string
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
	if msg.Data.CreatedDateTime.IsZero() {
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if isEmpty(msg.Data.MessagePagination) {
		ParamNames = append(ParamNames, "MessagePagination")
	}
	if msg.Data.ReportId == "" {
		ParamNames = append(ParamNames, "ReportId")
	}
	if msg.Data.ReportCreateDateTime.IsZero() {
		ParamNames = append(ParamNames, "ReportCreateDateTime")
	}
	if msg.Data.AccountOtherId == "" {
		ParamNames = append(ParamNames, "AccountOtherId")
	}
	if isEmpty(msg.Data.TotalCreditEntries) {
		ParamNames = append(ParamNames, "TotalCreditEntries")
	}
	if isEmpty(msg.Data.TotalDebitEntries) {
		ParamNames = append(ParamNames, "TotalDebitEntries")
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
	if msg.Data.MessageId != "" {
		err := camt052.AccountReportingFedwireFunds1(msg.Data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = camt052.AccountReportingFedwireFunds1(msg.Data.MessageId)
	}
	if !isEmpty(msg.Data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.Data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.Data.CreatedDateTime)
	}
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
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}

	var Rpt camt052.AccountReport251
	if msg.Data.ReportId != "" {
		Rpt.Id = camt052.ReportTimingFRS1(msg.Data.ReportId)
	}
	if !isEmpty(msg.Data.ReportCreateDateTime) {
		Rpt.CreDtTm = fedwire.ISODateTime(msg.Data.ReportCreateDateTime)
	}
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
	var TxsSummry camt052.TotalTransactions61
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
		TxsSummry.TtlCdtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.Data.TotalCreditEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.Data.TotalCreditEntries.Sum),
		}
	}
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
		TxsSummry.TtlDbtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.Data.TotalDebitEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.Data.TotalDebitEntries.Sum),
		}
	}

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
			err = camt052.BankTransactionCodeFedwireFunds1(entity.Status).Validate()
			if err != nil {
				return &model.ValidateError{
					ParamName: "TotalEntriesPerTransactionCode.Status",
					Message:   err.Error(),
				}
			}
			item := camt052.TotalsPerBankTransactionCode51{
				NbOfNtries: camt052.Max15NumericText(entity.NumberOfEntries),
				BkTxCd: camt052.BankTransactionCodeStructure41{
					Prtry: camt052.ProprietaryBankTransactionCodeStructure11{
						Cd: camt052.BankTransactionCodeFedwireFunds1(entity.Status),
					},
				},
			}
			TtlNtriesPerBkTxCd = append(TtlNtriesPerBkTxCd, item)
		}
		if !isEmpty(TtlNtriesPerBkTxCd) {
			TxsSummry.TtlNtriesPerBkTxCd = TtlNtriesPerBkTxCd
		}
	}

	if !isEmpty(TxsSummry) {
		Rpt.TxsSummry = TxsSummry
	}
	if msg.Data.AdditionalReportInfo != "" {
		Rpt.AddtlRptInf = camt052.Max500Text(msg.Data.AdditionalReportInfo)
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
				msg.Data.MessageId = model.CAMTReportType(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgId)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.CreDtTm) {
				msg.Data.CreatedDateTime = time.Time(msg.Doc.BkToCstmrAcctRpt.GrpHdr.CreDtTm)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb) {
					msg.Data.MessagePagination.PageNumber = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd) {
					msg.Data.MessagePagination.LastPageIndicator = bool(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd)
				}
			}
		}
		if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt) {
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Id) {
				msg.Data.ReportId = model.ReportType(msg.Doc.BkToCstmrAcctRpt.Rpt.Id)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.CreDtTm) {
				msg.Data.ReportCreateDateTime = time.Time(msg.Doc.BkToCstmrAcctRpt.Rpt.CreDtTm)
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
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.AddtlRptInf) {
				msg.Data.AdditionalReportInfo = string(msg.Doc.BkToCstmrAcctRpt.Rpt.AddtlRptInf)
			}
		}
	}
	return nil
}
