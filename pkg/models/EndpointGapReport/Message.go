package EndpointGapReport

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	camt052 "github.com/moov-io/fedwire20022/gen/EndpointGapReport_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"

type MessageModel struct {
	//Point to point reference, as assigned by the account servicing institution, and sent to the account owner or the party authorised to receive the message, to unambiguously identify the message.
	MessageId model.CAMTReportType
	//Date and time at which the message was created.
	CreatedDateTime time.Time
	//Provides details on the page number of the message.
	MessagePagination model.MessagePagenation
	//Report id on a cash account.
	ReportId GapType
	//This is the Fedwire Funds Service funds-transfer business day when the gap was identified.
	ReportCreateDateTime time.Time
	//Unambiguous identification of the account to which credit and debit entries are made.
	AccountOtherId string
	//For the Fedwire Funds Service, this provides the missing sequence numbers.
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
	var Rpt []camt052.AccountReport251
	var report_data camt052.AccountReport251
	if msg.Data.ReportId != "" {
		err := camt052.GapTypeFedwireFunds1(msg.Data.ReportId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportId",
				Message:   err.Error(),
			}
		}
		report_data.Id = camt052.GapTypeFedwireFunds1(msg.Data.ReportId)
	}
	if !isEmpty(msg.Data.ReportCreateDateTime) {
		err := fedwire.ISODateTime(msg.Data.ReportCreateDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportCreateDateTime",
				Message:   err.Error(),
			}
		}
		report_data.CreDtTm = fedwire.ISODateTime(msg.Data.ReportCreateDateTime)
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
		report_data.Acct = camt052.CashAccount391{
			Id: camt052.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	if msg.Data.AdditionalReportInfo != "" {
		err := camt052.Max500Text(msg.Data.AdditionalReportInfo).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AdditionalReportInfo",
				Message:   err.Error(),
			}
		}
		report_data.AddtlRptInf = camt052.Max500Text(msg.Data.AdditionalReportInfo)
	}
	if !isEmpty(report_data) {
		Rpt = append(Rpt, report_data)
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
			if len(msg.Doc.BkToCstmrAcctRpt.Rpt) > 0 {
				var report_data camt052.AccountReport251 = msg.Doc.BkToCstmrAcctRpt.Rpt[0]
				if !isEmpty(report_data.Id) {
					msg.Data.ReportId = GapType(report_data.Id)
				}
				if !isEmpty(report_data.CreDtTm) {
					msg.Data.ReportCreateDateTime = time.Time(report_data.CreDtTm)
				}
				if !isEmpty(report_data.Acct) {
					if !isEmpty(report_data.Acct.Id) {
						if !isEmpty(report_data.Acct.Id.Othr) {
							msg.Data.AccountOtherId = string(report_data.Acct.Id.Othr.Id)
						}
					}
				}
				if !isEmpty(report_data.AddtlRptInf) {
					msg.Data.AdditionalReportInfo = string(report_data.AddtlRptInf)
				}
			}
		}
	}
	return nil
}
