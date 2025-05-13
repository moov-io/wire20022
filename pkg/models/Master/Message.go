package Master

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	camt052 "github.com/moov-io/fedwire20022/gen/Master_camt_052_001_08"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08"

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
	AccountType    string
	//Identifies the parent account of the account for which the report has been issued.
	RelatedAccountOtherId string
	//Set of elements used to define the balance as a numerical representation of the net increases and decreases in an account at a specific point in time.
	Balances []Balance
	//Provides summary information on entries.
	TransactionsSummary []TotalsPerBankTransactionCode
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
  - With XML path: Loads file, parses XML into message.doc
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
	if msg.Data.ReportTypeId == "" {
		ParamNames = append(ParamNames, "ReportTypeId")
	}
	if msg.Data.ReportCreatedDate.IsZero() {
		ParamNames = append(ParamNames, "ReportCreatedDate")
	}
	if msg.Data.AccountOtherId == "" {
		ParamNames = append(ParamNames, "AccountOtherId")
	}
	if msg.Data.AccountType == "" {
		ParamNames = append(ParamNames, "AccountType")
	}
	if msg.Data.RelatedAccountOtherId == "" {
		ParamNames = append(ParamNames, "RelatedAccountOtherId")
	}
	if isEmpty(msg.Data.TransactionsSummary) {
		ParamNames = append(ParamNames, "TransactionsSummary")
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
	if !isEmpty(msg.Data.MessagePagination) {
		err := camt052.Max5NumericText(msg.Data.MessagePagination.PageNumber).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessagePagination.PageNumber",
				Message:   err.Error(),
			}
		}
		err = camt052.YesNoIndicator(msg.Data.MessagePagination.LastPageIndicator).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "MessagePagination.LastPageIndicator",
				Message:   err.Error(),
			}
		}
		GrpHdr.MsgPgntn = camt052.Pagination1{
			PgNb:      camt052.Max5NumericText(msg.Data.MessagePagination.PageNumber),
			LastPgInd: camt052.YesNoIndicator(msg.Data.MessagePagination.LastPageIndicator),
		}
	}
	var OrgnlBizQry camt052.OriginalBusinessQuery11
	if msg.Data.OriginalBusinessMsgId != "" {
		err := camt052.Max35Text(msg.Data.OriginalBusinessMsgId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalBusinessMsgId",
				Message:   err.Error(),
			}
		}
		OrgnlBizQry.MsgId = camt052.Max35Text(msg.Data.OriginalBusinessMsgId)
	}
	if msg.Data.OriginalBusinessMsgNameId != "" {
		err := camt052.MessageNameIdentificationFRS1(msg.Data.OriginalBusinessMsgNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalBusinessMsgNameId",
				Message:   err.Error(),
			}
		}
		OrgnlBizQry.MsgNmId = camt052.MessageNameIdentificationFRS1(msg.Data.OriginalBusinessMsgNameId)
	}
	if !isEmpty(msg.Data.OriginalBusinessMsgCreateTime) {
		err := fedwire.ISODateTime(msg.Data.OriginalBusinessMsgCreateTime).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "OriginalBusinessMsgCreateTime",
				Message:   err.Error(),
			}
		}
		OrgnlBizQry.CreDtTm = fedwire.ISODateTime(msg.Data.OriginalBusinessMsgCreateTime)
	}
	if !isEmpty(OrgnlBizQry) {
		GrpHdr.OrgnlBizQry = &OrgnlBizQry
	}
	if !isEmpty(GrpHdr) {
		BkToCstmrAcctRpt.GrpHdr = GrpHdr
	}
	var Rpt camt052.AccountReport251
	if msg.Data.ReportTypeId != "" {
		err := camt052.BalanceReportFRS1(msg.Data.ReportTypeId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportTypeId",
				Message:   err.Error(),
			}
		}
		Rpt.Id = camt052.BalanceReportFRS1(msg.Data.ReportTypeId)
	}
	if !isEmpty(msg.Data.ReportCreatedDate) {
		err := fedwire.ISODateTime(msg.Data.ReportCreatedDate).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportCreatedDate",
				Message:   err.Error(),
			}
		}
		Rpt.CreDtTm = fedwire.ISODateTime(msg.Data.ReportCreatedDate)
	}
	if msg.Data.AccountOtherId != "" {
		err := camt052.RoutingNumberFRS1(msg.Data.AccountOtherId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountOtherId",
				Message:   err.Error(),
			}
		}
		Othr := camt052.GenericAccountIdentification11{
			Id: camt052.RoutingNumberFRS1(msg.Data.AccountOtherId),
		}
		Rpt.Acct = camt052.CashAccount391{
			Id: camt052.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	if msg.Data.AccountType != "" {
		err := camt052.AccountTypeFRS1(msg.Data.AccountType).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountType",
				Message:   err.Error(),
			}
		}
		if !isEmpty(Rpt.Acct) {
			Prtry := camt052.AccountTypeFRS1(msg.Data.AccountType)
			Rpt.Acct.Tp = camt052.CashAccountType2Choice1{
				Prtry: &Prtry,
			}
		}
	}
	if msg.Data.RelatedAccountOtherId != "" {
		err := camt052.RoutingNumberFRS1(msg.Data.RelatedAccountOtherId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RelatedAccountOtherId",
				Message:   err.Error(),
			}
		}
		Othr := camt052.GenericAccountIdentification11{
			Id: camt052.RoutingNumberFRS1(msg.Data.RelatedAccountOtherId),
		}
		Rpt.RltdAcct = camt052.CashAccount381{
			Id: camt052.AccountIdentification4Choice1{
				Othr: &Othr,
			},
		}
	}
	var Bal []camt052.CashBalance81
	if !isEmpty(msg.Data.Balances) {
		for _, item := range msg.Data.Balances {
			line, vErr := CashBalance81From(item)
			if vErr != nil {
				vErr.InsertPath("Bal")
				return vErr
			}
			Bal = append(Bal, line)
		}
	}
	if !isEmpty(Bal) {
		Rpt.Bal = Bal
	}

	var TxsSummry camt052.TotalTransactions61
	var TtlNtriesPerBkTxCd []camt052.TotalsPerBankTransactionCode51
	for _, item := range msg.Data.TransactionsSummary {
		code, vErr := TotalsPerBankTransactionCode51From(item)
		if vErr != nil {
			vErr.InsertPath("TtlNtriesPerBkTxCd")
			return vErr
		}
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
				msg.Data.CreationDateTime = time.Time(msg.Doc.BkToCstmrAcctRpt.GrpHdr.CreDtTm)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn) {
				msg.Data.MessagePagination.PageNumber = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.PgNb)
				msg.Data.MessagePagination.LastPageIndicator = bool(msg.Doc.BkToCstmrAcctRpt.GrpHdr.MsgPgntn.LastPgInd)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId) {
					msg.Data.OriginalBusinessMsgId = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgId)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId) {
					msg.Data.OriginalBusinessMsgNameId = string(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.MsgNmId)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm) {
					msg.Data.OriginalBusinessMsgCreateTime = time.Time(msg.Doc.BkToCstmrAcctRpt.GrpHdr.OrgnlBizQry.CreDtTm)
				}
			}
		}
		if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt) {
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Id) {
				msg.Data.ReportTypeId = AccountReportType(msg.Doc.BkToCstmrAcctRpt.Rpt.Id)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.CreDtTm) {
				msg.Data.ReportCreatedDate = time.Time(msg.Doc.BkToCstmrAcctRpt.Rpt.CreDtTm)
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct.Id) {
					msg.Data.AccountOtherId = string(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct.Id.Othr.Id)
				}
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Acct.Tp) {
					msg.Data.AccountType = string(*msg.Doc.BkToCstmrAcctRpt.Rpt.Acct.Tp.Prtry)
				}
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.RltdAcct) {
				if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.RltdAcct.Id) {
					msg.Data.RelatedAccountOtherId = string(msg.Doc.BkToCstmrAcctRpt.Rpt.RltdAcct.Id.Othr.Id)
				}
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry) {
				for _, item := range msg.Doc.BkToCstmrAcctRpt.Rpt.TxsSummry.TtlNtriesPerBkTxCd {
					line := TotalsPerBankTransactionCode51To(item)
					msg.Data.TransactionsSummary = append(msg.Data.TransactionsSummary, line)
				}
			}
			if !isEmpty(msg.Doc.BkToCstmrAcctRpt.Rpt.Bal) {
				for _, item := range msg.Doc.BkToCstmrAcctRpt.Rpt.Bal {
					line := CashBalance81To(item)
					msg.Data.Balances = append(msg.Data.Balances, line)
				}
			}
		}

	}
	return nil
}
