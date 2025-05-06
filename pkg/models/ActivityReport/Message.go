package ActivityReport

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	camt052 "github.com/moov-io/fedwire20022/gen/ActivityReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"

type MessageModel struct {
	//MessageId (Message Identification) is a unique identifier assigned to an entire message.
	MessageId model.CAMTReportType
	//CreatedDateTime represents the timestamp when a message, instruction, or transaction was created
	//ISO 8601 format
	CreatedDateTime time.Time
	// MsgPgntn (Message Pagination) provides details about the pagination of the report.
	// It helps in handling reports split across multiple pages.
	Pagenation model.MessagePagenation
	// Id (Report Identification) uniquely identifies the report.
	// It provides a reference to the specific report being generated or requested.
	// Example value: "EDAY" (End-of-Day Report).
	ReportType           model.ReportType
	ReportCreateDateTime time.Time
	//// Acct (Account Information) contains details about the account being reported on.
	AccountOtherId string
	// TtlNtries (Total Entries) represents the overall count and sum of all transactions in the report.
	// This includes both credit and debit transactions.
	TotalEntries string
	// TtlCdtNtries (Total Credit Entries) specifies the total number and sum of credit transactions.
	// It represents transactions where funds are credited to an account.
	TotalCreditEntries model.NumberAndSumOfTransactions
	// TtlDbtNtries (Total Debit Entries) specifies the total number and sum of debit transactions.
	// It represents transactions where funds are debited from an account.
	TotalDebitEntries model.NumberAndSumOfTransactions
	// TtlNtriesPerBkTxCd (Total Entries Per Bank Transaction Code) provides a breakdown of transactions.
	// It groups transactions based on bank-specific transaction codes.
	TotalEntriesPerBankTransactionCode []TotalsPerBankTransactionCode
	// NtryDtls (Entry Details) contains detailed breakdowns of the transaction.
	// This can include supplementary details such as references, related transactions, and remittance information.
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
	Data, err := model.ReadXMLFile(filepath)
	if err != nil {
		return &msg, fmt.Errorf("file read error: %w", err)
	}

	// Handle empty XML Data
	if len(Data) == 0 {
		return &msg, fmt.Errorf("empty XML file: %s", filepath)
	}

	// Parse XML with structural validation
	if err := xml.Unmarshal(Data, &msg.Doc); err != nil {
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
	if isEmpty(msg.Data.CreatedDateTime) { // Check if CreatedDateTime is empty
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if isEmpty(msg.Data.Pagenation) {
		ParamNames = append(ParamNames, "Pagenation")
	}
	if msg.Data.ReportType == "" {
		ParamNames = append(ParamNames, "ReportType")
	}
	if isEmpty(msg.Data.ReportCreateDateTime) {
		ParamNames = append(ParamNames, "ReportCreateDateTime")
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
	if msg.Data.MessageId != "" {
		err := msg.Data.MessageId.Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
	}
	if !isEmpty(msg.Data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.Data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
	}
	msg.Doc = camt052.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
		BkToCstmrAcctRpt: camt052.BankToCustomerAccountReportV08{
			GrpHdr: camt052.GroupHeader811{
				MsgId:   camt052.AccountReportingFedwireFunds1(msg.Data.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.Data.CreatedDateTime),
			},
		},
	}
	if !isEmpty(msg.Data.Pagenation) {
		err := camt052.Max5NumericText(msg.Data.Pagenation.PageNumber).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"Pagenation"},
				ParamName:  "PageNumber",
				Message:    err.Error(),
			}
		}
		err = camt052.YesNoIndicator(msg.Data.Pagenation.LastPageIndicator).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"Pagenation"},
				ParamName:  "LastPageIndicator",
				Message:    err.Error(),
			}
		}
		msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn = camt052.Pagination1{
			PgNb:      camt052.Max5NumericText(msg.Data.Pagenation.PageNumber),
			LastPgInd: camt052.YesNoIndicator(msg.Data.Pagenation.LastPageIndicator),
		}
	}
	var Rpt camt052.AccountReport251
	if msg.Data.ReportType != "" {
		err := camt052.ReportTimingFRS1(msg.Data.ReportType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportType",
				Message:   err.Error(),
			}
		}
		Rpt.Id = camt052.ReportTimingFRS1(msg.Data.ReportType)
	}
	if !isEmpty(msg.Data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.Data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		Rpt.CreDtTm = fedwire.ISODateTime(msg.Data.CreatedDateTime)
	}
	var Acct camt052.CashAccount391
	if msg.Data.AccountOtherId != "" {
		err := camt052.RoutingNumberFRS1(msg.Data.AccountOtherId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountOtherId",
				Message:   err.Error(),
			}
		}
		_Othr := camt052.GenericAccountIdentification11{
			Id: camt052.RoutingNumberFRS1(msg.Data.AccountOtherId),
		}
		Acct.Id = camt052.AccountIdentification4Choice1{
			Othr: &_Othr,
		}
	}
	if !isEmpty(Acct) {
		Rpt.Acct = Acct
	}
	var TxsSummry camt052.TotalTransactions61
	if !isEmpty(msg.Data.TotalEntries) {
		err := camt052.Max15NumericText(msg.Data.TotalEntries).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalEntries",
				Message:   err.Error(),
			}
		}
		TxsSummry.TtlNtries = camt052.NumberAndSumOfTransactions41{
			NbOfNtries: camt052.Max15NumericText(msg.Data.TotalEntries),
		}
	}
	if !isEmpty(msg.Data.TotalCreditEntries) {
		err := camt052.Max15NumericText(msg.Data.TotalCreditEntries.NumberOfEntries).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"TotalCreditEntries"},
				ParamName:  "NumberOfEntries",
				Message:    err.Error(),
			}
		}
		err = camt052.DecimalNumber(msg.Data.TotalCreditEntries.Sum).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"TotalCreditEntries"},
				ParamName:  "Sum",
				Message:    err.Error(),
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
				ParentPath: []string{"TotalDebitEntries"},
				ParamName:  "NumberOfEntries",
				Message:    err.Error(),
			}
		}
		err = camt052.DecimalNumber(msg.Data.TotalDebitEntries.Sum).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"TotalDebitEntries"},
				ParamName:  "Sum",
				Message:    err.Error(),
			}
		}
		TxsSummry.TtlDbtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.Data.TotalDebitEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.Data.TotalDebitEntries.Sum),
		}
	}
	if !isEmpty(msg.Data.TotalEntriesPerBankTransactionCode) {
		var TtlNtriesPerBkTxCd []camt052.TotalsPerBankTransactionCode51
		for _, entity := range msg.Data.TotalEntriesPerBankTransactionCode {
			_item, err := TotalsPerBankTransactionCode51From(entity)
			if err != nil {
				err.InsertPath("TotalEntriesPerBankTransactionCode")
				return err
			}
			TtlNtriesPerBkTxCd = append(TtlNtriesPerBkTxCd, _item)
		}
		if !isEmpty(TtlNtriesPerBkTxCd) {
			TxsSummry.TtlNtriesPerBkTxCd = TtlNtriesPerBkTxCd
		}

	}
	if !isEmpty(TxsSummry) {
		Rpt.TxsSummry = &TxsSummry
	}
	var Ntry []*camt052.ReportEntry101
	if !isEmpty(msg.Data.EntryDetails) {
		for _, entity := range msg.Data.EntryDetails {
			_item, err := ReportEntry101From(entity)
			if err != nil {
				err.InsertPath("EntryDetails")
				return err
			}
			Ntry = append(Ntry, &_item)
		}
	}
	if !isEmpty(Ntry) {
		Rpt.Ntry = Ntry
	}
	if !isEmpty(Rpt) {
		msg.Doc.BkToCstmrAcctRpt.Rpt = Rpt
	}
	return nil
}
func (msg *Message) CreateMessageModel() *model.ValidateError {
	msg.Data = MessageModel{}
	if !isEmpty(msg.Doc) {
		if !isEmpty(msg.Doc.BkToCstmrAcctRpt) {
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgId) {
					msg.Data.MessageId = model.CAMTReportType(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgId)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.CreDtTm) {
					msg.Data.CreatedDateTime = time.Time(msg.Doc.BkToCstmrAcctRpt.GrpHdr.CreDtTm)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn) {
					msg.Data.Pagenation = model.MessagePagenation{
						PageNumber:        string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb),
						LastPageIndicator: bool(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd),
					}
				}
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Id) {
					msg.Data.ReportType = model.ReportType(msg.Doc.BkToCstmrAcctRpt.Rpt.Id)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.CreDtTm) {
					msg.Data.ReportCreateDateTime = time.Time(msg.Doc.BkToCstmrAcctRpt.Rpt.CreDtTm)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct) && !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct.Id) && !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr) {
					msg.Data.AccountOtherId = string(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr.Id)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry) {
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtries) {
						msg.Data.TotalEntries = string(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtries.NbOfNtries)
					}
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries) {
						msg.Data.TotalCreditEntries = model.NumberAndSumOfTransactions{
							NumberOfEntries: string(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries.NbOfNtries),
							Sum:             float64(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlCdtNtries.Sum),
						}
					}
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlDbtNtries) {
						msg.Data.TotalDebitEntries = model.NumberAndSumOfTransactions{
							NumberOfEntries: string(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlDbtNtries.NbOfNtries),
							Sum:             float64(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlDbtNtries.Sum),
						}
					}
					if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd) {
						msg.Data.TotalEntriesPerBankTransactionCode = make([]TotalsPerBankTransactionCode, 0)
						for _, entity := range msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd {
							if !isEmpty(entity) {
								_item := TotalsPerBankTransactionCode51To(entity)
								msg.Data.TotalEntriesPerBankTransactionCode = append(msg.Data.TotalEntriesPerBankTransactionCode, _item)
							}
						}
					}
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Ntry) {
					msg.Data.EntryDetails = make([]model.Entry, 0)
					for _, entity := range msg.Doc.BkToCstmrAcctRpt.Rpt.Ntry {
						if !isEmpty(entity) {
							_item := ReportEntry101To(*entity)
							msg.Data.EntryDetails = append(msg.Data.EntryDetails, _item)
						}
					}
				}
			}
		}
	}
	return nil
}
