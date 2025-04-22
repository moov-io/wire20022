package ActivityReport

import (
	"encoding/xml"
	"fmt"
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
	if msg.data.MessageId != "" {
		err := model.CAMTReportType(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
	}
	msg.doc = camt052.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
		BkToCstmrAcctRpt: camt052.BankToCustomerAccountReportV08{
			GrpHdr: camt052.GroupHeader811{
				MsgId:   camt052.AccountReportingFedwireFunds1(msg.data.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.data.CreatedDateTime),
			},
		},
	}
	if !isEmpty(msg.data.Pagenation) {
		err := camt052.Max5NumericText(msg.data.Pagenation.PageNumber).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"Pagenation"},
				ParamName:  "PageNumber",
				Message:    err.Error(),
			}
		}
		err = camt052.YesNoIndicator(msg.data.Pagenation.LastPageIndicator).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"Pagenation"},
				ParamName:  "LastPageIndicator",
				Message:    err.Error(),
			}
		}
		msg.doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn = camt052.Pagination1{
			PgNb:      camt052.Max5NumericText(msg.data.Pagenation.PageNumber),
			LastPgInd: camt052.YesNoIndicator(msg.data.Pagenation.LastPageIndicator),
		}
	}
	var Rpt camt052.AccountReport251
	if msg.data.ReportType != "" {
		err := camt052.ReportTimingFRS1(msg.data.ReportType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportType",
				Message:   err.Error(),
			}
		}
		Rpt.Id = camt052.ReportTimingFRS1(msg.data.ReportType)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		Rpt.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
	var Acct camt052.CashAccount391
	if msg.data.AccountOtherId != "" {
		err := camt052.RoutingNumberFRS1(msg.data.AccountOtherId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountOtherId",
				Message:   err.Error(),
			}
		}
		_Othr := camt052.GenericAccountIdentification11{
			Id: camt052.RoutingNumberFRS1(msg.data.AccountOtherId),
		}
		Acct.Id = camt052.AccountIdentification4Choice1{
			Othr: &_Othr,
		}
	}
	if !isEmpty(Acct) {
		Rpt.Acct = Acct
	}
	var TxsSummry camt052.TotalTransactions61
	if !isEmpty(msg.data.TotalEntries) {
		err := camt052.Max15NumericText(msg.data.TotalEntries).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "TotalEntries",
				Message:   err.Error(),
			}
		}
		TxsSummry.TtlNtries = camt052.NumberAndSumOfTransactions41{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalEntries),
		}
	}
	if !isEmpty(msg.data.TotalCreditEntries) {
		err := camt052.Max15NumericText(msg.data.TotalCreditEntries.NumberOfEntries).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"TotalCreditEntries"},
				ParamName:  "NumberOfEntries",
				Message:    err.Error(),
			}
		}
		err = camt052.DecimalNumber(msg.data.TotalCreditEntries.Sum).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"TotalCreditEntries"},
				ParamName:  "Sum",
				Message:    err.Error(),
			}
		}
		TxsSummry.TtlCdtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalCreditEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalCreditEntries.Sum),
		}
	}
	if !isEmpty(msg.data.TotalDebitEntries) {
		err := camt052.Max15NumericText(msg.data.TotalDebitEntries.NumberOfEntries).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"TotalDebitEntries"},
				ParamName:  "NumberOfEntries",
				Message:    err.Error(),
			}
		}
		err = camt052.DecimalNumber(msg.data.TotalDebitEntries.Sum).Validate()
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"TotalDebitEntries"},
				ParamName:  "Sum",
				Message:    err.Error(),
			}
		}
		TxsSummry.TtlDbtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalDebitEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalDebitEntries.Sum),
		}
	}
	if !isEmpty(msg.data.TotalEntriesPerBankTransactionCode) {
		var TtlNtriesPerBkTxCd []camt052.TotalsPerBankTransactionCode51
		for _, entity := range msg.data.TotalEntriesPerBankTransactionCode {
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
	if !isEmpty(msg.data.EntryDetails) {
		for _, entity := range msg.data.EntryDetails {
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
		msg.doc.BkToCstmrAcctRpt.Rpt = Rpt
	}
	return nil
}
