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

func (msg *Message) ValidateRequiredFields() *model.ValidateError {
	// Initialize the RequireError object
	var ParamNames []string

	// Check required fields and append missing ones to ParamNames
	if msg.data.MessageId == "" {
		ParamNames = append(ParamNames, "MessageId")
	}
	if msg.data.CreatedDateTime.IsZero() {
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if isEmpty(msg.data.MessagePagination) {
		ParamNames = append(ParamNames, "MessagePagination")
	}
	if msg.data.ReportId == "" {
		ParamNames = append(ParamNames, "ReportId")
	}
	if msg.data.ReportCreateDateTime.IsZero() {
		ParamNames = append(ParamNames, "ReportCreateDateTime")
	}
	if msg.data.AccountOtherId == "" {
		ParamNames = append(ParamNames, "AccountOtherId")
	}
	if isEmpty(msg.data.TotalCreditEntries) {
		ParamNames = append(ParamNames, "TotalCreditEntries")
	}
	if isEmpty(msg.data.TotalDebitEntries) {
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
	msg.doc = camt052.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
	}
	var BkToCstmrAcctRpt camt052.BankToCustomerAccountReportV08
	var GrpHdr camt052.GroupHeader811
	if msg.data.MessageId != "" {
		err := camt052.AccountReportingFedwireFunds1(msg.data.MessageId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessageId",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgId = camt052.AccountReportingFedwireFunds1(msg.data.MessageId)
	}
	if !isEmpty(msg.data.CreatedDateTime) {
		err := fedwire.ISODateTime(msg.data.CreatedDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "CreatedDateTime",
				Message:   err.Error(),
			}
		}
		GrpHdr.CreDtTm = fedwire.ISODateTime(msg.data.CreatedDateTime)
	}
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
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}

	var Rpt camt052.AccountReport251
	if msg.data.ReportId != "" {
		Rpt.Id = camt052.ReportTimingFRS1(msg.data.ReportId)
	}
	if !isEmpty(msg.data.ReportCreateDateTime) {
		Rpt.CreDtTm = fedwire.ISODateTime(msg.data.ReportCreateDateTime)
	}
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
	var TxsSummry camt052.TotalTransactions61
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
		TxsSummry.TtlCdtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalCreditEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalCreditEntries.Sum),
		}
	}
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
		TxsSummry.TtlDbtNtries = camt052.NumberAndSumOfTransactions11{
			NbOfNtries: camt052.Max15NumericText(msg.data.TotalDebitEntries.NumberOfEntries),
			Sum:        camt052.DecimalNumber(msg.data.TotalDebitEntries.Sum),
		}
	}

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
	if msg.data.AdditionalReportInfo != "" {
		Rpt.AddtlRptInf = camt052.Max500Text(msg.data.AdditionalReportInfo)
	}
	if !isEmpty(Rpt) {
		BkToCstmrAcctRpt.Rpt = Rpt
	}
	if !isEmpty(BkToCstmrAcctRpt) {
		msg.doc.BkToCstmrAcctRpt = BkToCstmrAcctRpt
	}
	return nil
}
