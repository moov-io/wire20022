package EndpointGapReport

import (
	"encoding/xml"
	"fmt"
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
	var Rpt []camt052.AccountReport251
	var report_data camt052.AccountReport251
	if msg.data.ReportId != "" {
		err := camt052.GapTypeFedwireFunds1(msg.data.ReportId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportId",
				Message:   err.Error(),
			}
		}
		report_data.Id = camt052.GapTypeFedwireFunds1(msg.data.ReportId)
	}
	if !isEmpty(msg.data.ReportCreateDateTime) {
		err := fedwire.ISODateTime(msg.data.ReportCreateDateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportCreateDateTime",
				Message:   err.Error(),
			}
		}
		report_data.CreDtTm = fedwire.ISODateTime(msg.data.ReportCreateDateTime)
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
		report_data.Acct = camt052.CashAccount391{
			Id: camt052.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	if msg.data.AdditionalReportInfo != "" {
		err := camt052.Max500Text(msg.data.AdditionalReportInfo).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AdditionalReportInfo",
				Message:   err.Error(),
			}
		}
		report_data.AddtlRptInf = camt052.Max500Text(msg.data.AdditionalReportInfo)
	}
	if !isEmpty(report_data) {
		Rpt = append(Rpt, report_data)
	}
	if !isEmpty(Rpt) {
		BkToCstmrAcctRpt.Rpt = Rpt
	}
	if !isEmpty(BkToCstmrAcctRpt) {
		msg.doc.BkToCstmrAcctRpt = BkToCstmrAcctRpt
	}
	return nil
}
