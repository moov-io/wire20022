package AccountReportingRequest

import (
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
	data MessageModel
	doc  camt060.Document
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
    if msg.data.CreatedDateTime.IsZero() { // Check if CreatedDateTime is empty
        ParamNames = append(ParamNames, "CreatedDateTime")
    }
    if msg.data.ReportRequestId == "" {
        ParamNames = append(ParamNames, "ReportRequestId")
    }
    if msg.data.RequestedMsgNameId == "" {
        ParamNames = append(ParamNames, "RequestedMsgNameId")
    }
    if isEmpty(msg.data.AccountOwnerAgent.agent) {
        ParamNames = append(ParamNames, "AccountOwnerAgent.agent")
    }

    // Return nil if no required fields are missing
    if len(ParamNames) == 0 {
        return nil
    }

    // Return the error with missing fields
    return &model.ValidateError{
		ParamName:  "RequiredFields",
		Message:    strings.Join(ParamNames, ", "),
	}
}
func (msg *Message) CreateDocument() *model.ValidateError {
	requireErr := msg.ValidateRequiredFields()
	if requireErr != nil {
		return requireErr
	}
	if msg.data.MessageId != "" {
		err := camt060.Max35Text(msg.data.MessageId).Validate()
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
	msg.doc = camt060.Document{
		XMLName: xml.Name{
			Space: XMLINS,
			Local: "Document",
		},
		AcctRptgReq: camt060.AccountReportingRequestV05{
			GrpHdr: camt060.GroupHeader771{
				MsgId:   camt060.Max35Text(msg.data.MessageId),
				CreDtTm: fedwire.ISODateTime(msg.data.CreatedDateTime),
			},
		},
	}
	var RptgReq camt060.ReportingRequest51
	if msg.data.ReportRequestId != "" {
		err := msg.data.ReportRequestId.Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "ReportRequestId",
				Message:   err.Error(),
			}
		}
		RptgReq.Id = camt060.AccountReportingFedwireFunds1(msg.data.ReportRequestId)
	}
	if msg.data.RequestedMsgNameId != "" {
		err := camt060.MessageNameIdentificationFRS1(msg.data.RequestedMsgNameId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "RequestedMsgNameId",
				Message:   err.Error(),
			}
		}
		RptgReq.ReqdMsgNmId = camt060.MessageNameIdentificationFRS1(msg.data.RequestedMsgNameId)
	}
	if msg.data.AccountOtherId != "" {
		err := camt060.RoutingNumberFRS1(msg.data.AccountOtherId).Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountOtherId",
				Message:   err.Error(),
			}
		}
		id_othr := camt060.GenericAccountIdentification11{
			Id: camt060.RoutingNumberFRS1(msg.data.AccountOtherId),
		}

		_account := camt060.CashAccount381{
			Id: camt060.AccountIdentification4Choice1{
				Othr: &id_othr,
			},
		}
		RptgReq.Acct = &_account
	}
	if msg.data.AccountProperty != "" {
		err := msg.data.AccountProperty.Validate()
		if err != nil {
			return &model.ValidateError{
				ParamName: "AccountProperty",
				Message:   err.Error(),
			}
		}
		_Prtry := camt060.AccountTypeFRS1(msg.data.AccountProperty)
		RptgReq.Acct.Tp = camt060.CashAccountType2Choice1{
			Prtry: &_Prtry,
		}
	}
	if !isEmpty(msg.data.AccountOwnerAgent.agent) {
		_AcctOwnr, err := Party40Choice1From(msg.data.AccountOwnerAgent.agent)
		if err != nil {
			err.ParentPath = []string{"AccountOwnerAgent", "agent"}
			return err
		}
		if !isEmpty(_AcctOwnr) {
			RptgReq.AcctOwnr = _AcctOwnr
		}
		if msg.data.AccountOwnerAgent.OtherId != "" {
			err := camt060.EndpointIdentifierFedwireFunds1(msg.data.AccountOwnerAgent.OtherId).Validate()
			if err != nil {
				vErr := model.ValidateError{
					ParamName: "OtherId",
					Message:   err.Error(),
				}
				vErr.ParentPath = []string{"AccountOwnerAgent", "agent"}
				return &vErr
			}
			_Other := camt060.GenericFinancialIdentification11{
				Id: camt060.EndpointIdentifierFedwireFunds1(msg.data.AccountOwnerAgent.OtherId),
			}
			RptgReq.AcctOwnr.Agt.FinInstnId.Othr = &_Other
		}
	}
	if !isEmpty(msg.data.FromToSeuence) {
		FrSeq, err := strconv.ParseFloat(msg.data.FromToSeuence.FromSeq, 64)
		if err != nil {
			return &model.ValidateError{
				ParentPath: []string{"FromToSeuence"},
				ParamName:  "FromSeq",
				Message:    err.Error(),
			}
		}
		ToSeq, err := strconv.ParseFloat(msg.data.FromToSeuence.ToSeq, 64)
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
		msg.doc.AcctRptgReq.RptgReq = RptgReq
	}
	return nil
}
