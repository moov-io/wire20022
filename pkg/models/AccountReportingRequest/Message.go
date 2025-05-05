package AccountReportingRequest

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	camt060 "github.com/moov-io/fedwire20022/gen/AccountReportingRequest_camt_060_001_05"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	model "github.com/moov-io/wire20022/pkg/models"
)

const XMLINS string = "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05"

type MessageModel struct {
	//MessageId (Message Identification) is a unique identifier assigned to an entire message.
	MessageId string
	//CreatedDateTime represents the timestamp when a message, instruction, or transaction was created
	//ISO 8601 format
	CreatedDateTime time.Time
	//Unique identification, as assigned by the account owner, to unambiguously identify the account reporting request.
	ReportRequestId model.CAMTReportType
	//Specifies the type of the requested reporting message.
	RequestedMsgNameId string
	//account or entity identifier does not conform to any predefined ISO 20022 standard
	AccountOtherId string
	//AccountProperty defines the properties of a financial account.
	AccountProperty AccountTypeFRS
	// It is defined as a Camt060Agent type which encapsulates the choice of different party identification options for the account owner.
	AccountOwnerAgent Camt060Agent
	//"From-To" sequence within the ISO 20022 camt.060.001.05 message.
	FromToSeuence model.SequenceRange
}

type Message struct {
	Data   MessageModel
	Doc    camt060.Document
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
func (msg *Message) GetXML() ([]byte, error) {
	return xml.MarshalIndent(msg.Doc, "", "\t")
}
func (msg *Message) GetJSON() ([]byte, error) {
	return json.MarshalIndent(msg.Doc, "", "\t")
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
	msg := Message{Data: MessageModel{}} // Initialize with zero value
	msg.Helper = BuildMessageHelper()

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
	if err := xml.Unmarshal(data, &msg.Doc); err != nil {
		return msg, fmt.Errorf("XML parse error: %w", err)
	}

	return msg, nil
}
func (msg *Message) ValidateRequiredFields() *model.ValidateError {
	// Initialize the RequireError object
	var ParamNames []string

	// Check required fields and append missing ones to ParamNames
	if msg.Data.MessageId == "" {
		ParamNames = append(ParamNames, "MessageId")
	}
	if msg.Data.CreatedDateTime.IsZero() { // Check if CreatedDateTime is empty
		ParamNames = append(ParamNames, "CreatedDateTime")
	}
	if msg.Data.ReportRequestId == "" {
		ParamNames = append(ParamNames, "ReportRequestId")
	}
	if msg.Data.RequestedMsgNameId == "" {
		ParamNames = append(ParamNames, "RequestedMsgNameId")
	}
	if isEmpty(msg.Data.AccountOwnerAgent.Agent) {
		ParamNames = append(ParamNames, "AccountOwnerAgent.agent")
	}

	// Return nil if no required fields are missing
	if len(ParamNames) == 0 {
		return nil
	}

	// Return the error with missing fields
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
		err := camt060.Max35Text(msg.Data.MessageId).Validate()
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
	msg.Doc = camt060.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
		AcctRptgReq: camt060.AccountReportingRequestV05{
			GrpHdr: camt060.GroupHeader771{
				MsgId:   camt060.Max35Text(msg.Data.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.Data.CreatedDateTime),
			},
		},
	}
	var RptgReq camt060.ReportingRequest51
	if msg.Data.ReportRequestId != "" {
		err := msg.Data.ReportRequestId.Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportRequestId",
				Message:   err.Error(),
			}
		}
		RptgReq.Id = camt060.AccountReportingFedwireFunds1(msg.Data.ReportRequestId)
	}
	if msg.Data.RequestedMsgNameId != "" {
		err := camt060.MessageNameIdentificationFRS1(msg.Data.RequestedMsgNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RequestedMsgNameId",
				Message:   err.Error(),
			}
		}
		RptgReq.ReqdMsgNmId = camt060.MessageNameIdentificationFRS1(msg.Data.RequestedMsgNameId)
	}
	if msg.Data.AccountOtherId != "" {
		err := camt060.RoutingNumberFRS1(msg.Data.AccountOtherId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountOtherId",
				Message:   err.Error(),
			}
		}
		id_othr := camt060.GenericAccountIdentification11{
			Id: camt060.RoutingNumberFRS1(msg.Data.AccountOtherId),
		}

		_account := camt060.CashAccount381{
			Id: camt060.AccountIdentification4Choice1{
				Othr: &id_othr,
			},
		}
		RptgReq.Acct = &_account
	}
	if msg.Data.AccountProperty != "" {
		err := msg.Data.AccountProperty.Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountProperty",
				Message:   err.Error(),
			}
		}
		_Prtry := camt060.AccountTypeFRS1(msg.Data.AccountProperty)
		RptgReq.Acct.Tp = camt060.CashAccountType2Choice1{
			Prtry: &_Prtry,
		}
	}
	if !isEmpty(msg.Data.AccountOwnerAgent.Agent) {
		_AcctOwnr, err := Party40Choice1From(msg.Data.AccountOwnerAgent.Agent)
		if err != nil {
			err.ParentPath = []string{"AccountOwnerAgent", "agent"}
			return err
		}
		if !isEmpty(_AcctOwnr) {
			RptgReq.AcctOwnr = _AcctOwnr
		}
		if msg.Data.AccountOwnerAgent.OtherId != "" {
			err := camt060.EndpointIdentifierFedwireFunds1(msg.Data.AccountOwnerAgent.OtherId).Validate()
			if err != nil {
				vErr := model.ValidateError{
					ParamName: "OtherId",
					Message:   err.Error(),
				}
				vErr.ParentPath = []string{"AccountOwnerAgent", "agent"}
				return &vErr
			}
			_Other := camt060.GenericFinancialIdentification11{
				Id: camt060.EndpointIdentifierFedwireFunds1(msg.Data.AccountOwnerAgent.OtherId),
			}
			RptgReq.AcctOwnr.Agt.FinInstnId.Othr = &_Other
		}
	}
	if !isEmpty(msg.Data.FromToSeuence) {
		FrSeq, err := strconv.ParseFloat(msg.Data.FromToSeuence.FromSeq, 64)
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"FromToSeuence"},
				ParamName:  "FromSeq",
				Message:    err.Error(),
			}
		}
		ToSeq, err := strconv.ParseFloat(msg.Data.FromToSeuence.ToSeq, 64)
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"FromToSeuence"},
				ParamName:  "ToSeq",
				Message:    err.Error(),
			}
		}
		// err = camt060.XSequenceNumberFedwireFunds1(FrSeq).Validate()
		// if err != nil {
		// 	return &model.ValidateError{
		// 		ParentPath: []string{"FromToSeuence"},
		// 		ParamName:  "FromSeq",
		// 		Message:    err.Error(),
		// 	}
		// }
		// err = camt060.XSequenceNumberFedwireFunds1(ToSeq).Validate()
		// if err != nil {
		// 	return &model.ValidateError{
		// 		ParentPath: []string{"FromToSeuence"},
		// 		ParamName:  "ToSeq",
		// 		Message:    err.Error(),
		// 	}
		// }
		_FrToSeq := camt060.SequenceRange11{
			FrSeq: camt060.XSequenceNumberFedwireFunds1(FrSeq),
			ToSeq: camt060.XSequenceNumberFedwireFunds1(ToSeq),
		}
		_RptgSeq := camt060.SequenceRange1Choice1{
			FrToSeq: &_FrToSeq,
		}
		RptgReq.RptgSeq = &_RptgSeq
	}
	if !isEmpty(RptgReq) {
		msg.Doc.AcctRptgReq.RptgReq = RptgReq
	}
	return nil
}
