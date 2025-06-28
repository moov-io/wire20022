package messages

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/moov-io/wire20022/pkg/errors"
	AccountReportingRequestModel "github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
	ActivityReportModel "github.com/moov-io/wire20022/pkg/models/ActivityReport"
	ConnectionCheckModel "github.com/moov-io/wire20022/pkg/models/ConnectionCheck"
	CustomerCreditTransferModel "github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
	DrawdownRequestModel "github.com/moov-io/wire20022/pkg/models/DrawdownRequest"
	DrawdownResponseModel "github.com/moov-io/wire20022/pkg/models/DrawdownResponse"
	EndpointDetailsReportModel "github.com/moov-io/wire20022/pkg/models/EndpointDetailsReport"
	EndpointGapReportModel "github.com/moov-io/wire20022/pkg/models/EndpointGapReport"
	EndpointTotalsReportModel "github.com/moov-io/wire20022/pkg/models/EndpointTotalsReport"
	FedwireFundsAcknowledgementModel "github.com/moov-io/wire20022/pkg/models/FedwireFundsAcknowledgement"
	FedwireFundsPaymentStatusModel "github.com/moov-io/wire20022/pkg/models/FedwireFundsPaymentStatus"
	FedwireFundsSystemResponseModel "github.com/moov-io/wire20022/pkg/models/FedwireFundsSystemResponse"
	MasterModel "github.com/moov-io/wire20022/pkg/models/Master"
	PaymentReturnModel "github.com/moov-io/wire20022/pkg/models/PaymentReturn"
	PaymentStatusRequestModel "github.com/moov-io/wire20022/pkg/models/PaymentStatusRequest"
	ReturnRequestResponseModel "github.com/moov-io/wire20022/pkg/models/ReturnRequestResponse"
)

// MessageType represents the detected ISO 20022 message type
type MessageType string

const (
	TypeCustomerCreditTransfer      MessageType = "CustomerCreditTransfer"
	TypePaymentReturn               MessageType = "PaymentReturn"
	TypePaymentStatusRequest        MessageType = "PaymentStatusRequest"
	TypeFedwireFundsPaymentStatus   MessageType = "FedwireFundsPaymentStatus"
	TypeDrawdownRequest             MessageType = "DrawdownRequest"
	TypeDrawdownResponse            MessageType = "DrawdownResponse"
	TypeAccountReportingRequest     MessageType = "AccountReportingRequest"
	TypeActivityReport              MessageType = "ActivityReport"
	TypeEndpointDetailsReport       MessageType = "EndpointDetailsReport"
	TypeEndpointGapReport           MessageType = "EndpointGapReport"
	TypeEndpointTotalsReport        MessageType = "EndpointTotalsReport"
	TypeReturnRequestResponse       MessageType = "ReturnRequestResponse"
	TypeConnectionCheck             MessageType = "ConnectionCheck"
	TypeFedwireFundsAcknowledgement MessageType = "FedwireFundsAcknowledgement"
	TypeFedwireFundsSystemResponse  MessageType = "FedwireFundsSystemResponse"
	TypeMaster                      MessageType = "Master"
	TypeUnknown                     MessageType = "Unknown"
)

// DetectionInfo contains information about the detected message type
type DetectionInfo struct {
	MessageType     MessageType
	RootElement     string
	Namespace       string
	NamespacePrefix string
	Version         string
	DetectedBy      string // "namespace", "root_element", or "content_analysis"
	AdditionalInfo  map[string]string
}

// ParsedMessage contains the detected message type and parsed content
type ParsedMessage struct {
	Type      MessageType
	Message   interface{} // The actual message struct (e.g., *CustomerCreditTransfer.MessageModel)
	Version   string
	Detection DetectionInfo
}

// UniversalReader reads and automatically detects Fedwire ISO 20022 message types
type UniversalReader struct {
	// Configuration for enhanced error reporting
	VerboseErrors    bool
	TrackLineNumbers bool
}

// NewUniversalReader creates a new universal reader instance
func NewUniversalReader() *UniversalReader {
	return &UniversalReader{
		VerboseErrors:    true,
		TrackLineNumbers: true,
	}
}

// xmlPeeker helps peek at XML structure without full parsing
type xmlPeeker struct {
	RootElement     string
	Namespace       string
	NamespacePrefix string
	Attributes      map[string]string
}

// peekXML examines the XML structure to extract root element and namespace
func (r *UniversalReader) peekXML(data []byte) (*xmlPeeker, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	peek := &xmlPeeker{
		Attributes: make(map[string]string),
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("error reading XML token: %w", err)
		}

		switch t := token.(type) {
		case xml.StartElement:
			// First start element is our root
			peek.RootElement = t.Name.Local

			// Extract namespace info
			for _, attr := range t.Attr {
				peek.Attributes[attr.Name.Local] = attr.Value

				if attr.Name.Local == "xmlns" {
					peek.Namespace = attr.Value
				} else if strings.HasPrefix(attr.Name.Local, "xmlns:") {
					prefix := strings.TrimPrefix(attr.Name.Local, "xmlns:")
					if strings.Contains(attr.Value, "iso:20022") {
						peek.NamespacePrefix = prefix
						peek.Namespace = attr.Value
					}
				}
			}

			// We only need the root element
			return peek, nil
		}
	}

	return peek, fmt.Errorf("no root element found in XML")
}

// extractMessageTypeFromNamespace extracts the message type and version from namespace
func (r *UniversalReader) extractMessageTypeFromNamespace(namespace string) (string, string) {
	// Pattern: urn:iso:std:iso:20022:tech:xsd:pacs.008.001.12
	parts := strings.Split(namespace, ":")
	if len(parts) < 8 {
		return "", ""
	}

	// Get the last part which contains message.type.version
	lastPart := parts[len(parts)-1]
	msgParts := strings.Split(lastPart, ".")

	if len(msgParts) >= 4 {
		msgType := msgParts[0] + "." + msgParts[1]
		version := strings.Join(msgParts[2:], ".")
		return msgType, version
	}

	return "", ""
}

// detectMessageType determines the message type from XML structure
func (r *UniversalReader) detectMessageType(peek *xmlPeeker, data []byte) (*DetectionInfo, error) {
	info := &DetectionInfo{
		RootElement:     peek.RootElement,
		Namespace:       peek.Namespace,
		NamespacePrefix: peek.NamespacePrefix,
		AdditionalInfo:  make(map[string]string),
	}

	// First try namespace detection
	msgType, version := r.extractMessageTypeFromNamespace(peek.Namespace)
	info.Version = version

	if msgType != "" {
		info.DetectedBy = "namespace"
		info.AdditionalInfo["namespace_type"] = msgType
	}

	// Map based on root element and namespace type
	switch peek.RootElement {
	case "FIToFICstmrCdtTrf":
		info.MessageType = TypeCustomerCreditTransfer
	case "PmtRtr":
		info.MessageType = TypePaymentReturn
	case "FIToFIPmtStsReq":
		info.MessageType = TypePaymentStatusRequest
	case "FIToFIPmtStsRpt":
		info.MessageType = TypeFedwireFundsPaymentStatus
	case "CdtrPmtActvtnReq":
		info.MessageType = TypeDrawdownRequest
	case "CdtrPmtActvtnReqStsRpt":
		info.MessageType = TypeDrawdownResponse
	case "AcctRptgReq":
		info.MessageType = TypeAccountReportingRequest
	case "RsltnOfInvstgtn":
		info.MessageType = TypeReturnRequestResponse
	case "SysEvtNtfctn":
		info.MessageType = TypeConnectionCheck
	case "RctAck":
		info.MessageType = TypeFedwireFundsAcknowledgement
	case "SysEvtAck":
		info.MessageType = TypeFedwireFundsSystemResponse
	case "BkToCstmrAcctRpt":
		// This requires content analysis
		info.DetectedBy = "content_analysis"
		return r.analyzeBkToCstmrAcctRpt(info, data)
	case "Document":
		// ISO 20022 messages are often wrapped in a Document element
		// We need to peek deeper to find the actual message element
		return r.analyzeDocumentWrapper(info, data)
	default:
		info.MessageType = TypeUnknown
		info.DetectedBy = "failed"
		return info, fmt.Errorf("unknown root element: %s", peek.RootElement)
	}

	if info.DetectedBy == "" {
		info.DetectedBy = "root_element"
	}

	return info, nil
}

// analyzeDocumentWrapper analyzes ISO 20022 messages wrapped in Document element
func (r *UniversalReader) analyzeDocumentWrapper(info *DetectionInfo, data []byte) (*DetectionInfo, error) {
	// Parse the XML to find the child element of Document
	decoder := xml.NewDecoder(bytes.NewReader(data))

	inDocument := false
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return info, fmt.Errorf("error reading XML token: %w", err)
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "Document" {
				inDocument = true
				continue
			}

			if inDocument {
				// This is the actual message element inside the Document
				childInfo := &DetectionInfo{
					RootElement:     t.Name.Local,
					Namespace:       info.Namespace, // Use the namespace from Document
					NamespacePrefix: info.NamespacePrefix,
					AdditionalInfo:  make(map[string]string),
				}

				// Extract version from namespace
				msgType, version := r.extractMessageTypeFromNamespace(info.Namespace)
				childInfo.Version = version

				if msgType != "" {
					childInfo.DetectedBy = "namespace"
					childInfo.AdditionalInfo["namespace_type"] = msgType
				}

				// Map the child element to message type
				switch t.Name.Local {
				case "FIToFICstmrCdtTrf":
					childInfo.MessageType = TypeCustomerCreditTransfer
				case "PmtRtr":
					childInfo.MessageType = TypePaymentReturn
				case "FIToFIPmtStsReq":
					childInfo.MessageType = TypePaymentStatusRequest
				case "FIToFIPmtStsRpt":
					childInfo.MessageType = TypeFedwireFundsPaymentStatus
				case "CdtrPmtActvtnReq":
					childInfo.MessageType = TypeDrawdownRequest
				case "CdtrPmtActvtnReqStsRpt":
					childInfo.MessageType = TypeDrawdownResponse
				case "AcctRptgReq":
					childInfo.MessageType = TypeAccountReportingRequest
				case "RsltnOfInvstgtn":
					childInfo.MessageType = TypeReturnRequestResponse
				case "SysEvtNtfctn":
					childInfo.MessageType = TypeConnectionCheck
				case "RctAck":
					childInfo.MessageType = TypeFedwireFundsAcknowledgement
				case "SysEvtAck":
					childInfo.MessageType = TypeFedwireFundsSystemResponse
				case "BkToCstmrAcctRpt":
					// This requires content analysis
					childInfo.DetectedBy = "content_analysis"
					return r.analyzeBkToCstmrAcctRpt(childInfo, data)
				default:
					childInfo.MessageType = TypeUnknown
					childInfo.DetectedBy = "failed"
					return childInfo, fmt.Errorf("unknown message element inside Document: %s", t.Name.Local)
				}

				if childInfo.DetectedBy == "" {
					childInfo.DetectedBy = "root_element"
				}

				return childInfo, nil
			}
		}
	}

	return info, fmt.Errorf("no message element found inside Document wrapper")
}

// camt052Analyzer helps analyze BkToCstmrAcctRpt messages
type camt052Analyzer struct {
	XMLName xml.Name `xml:"BkToCstmrAcctRpt"`
	GrpHdr  struct {
		MsgId string `xml:"MsgId"`
	} `xml:"GrpHdr"`
	Rpt []struct {
		Id string `xml:"Id"`
	} `xml:"Rpt"`
}

// analyzeBkToCstmrAcctRpt performs content analysis for camt.052 messages
func (r *UniversalReader) analyzeBkToCstmrAcctRpt(info *DetectionInfo, data []byte) (*DetectionInfo, error) {
	var analyzer camt052Analyzer
	if err := xml.Unmarshal(data, &analyzer); err != nil {
		return info, fmt.Errorf("failed to analyze BkToCstmrAcctRpt content: %w", err)
	}

	// Store analysis results
	info.AdditionalInfo["GrpHdr.MsgId"] = analyzer.GrpHdr.MsgId
	if len(analyzer.Rpt) > 0 {
		info.AdditionalInfo["Rpt.Id"] = analyzer.Rpt[0].Id
	}

	// Determine specific type based on MsgId and Rpt.Id
	msgId := analyzer.GrpHdr.MsgId
	rptId := ""
	if len(analyzer.Rpt) > 0 {
		rptId = analyzer.Rpt[0].Id
	}

	switch {
	case strings.HasPrefix(msgId, "ACTR"):
		info.MessageType = TypeActivityReport
	case strings.HasPrefix(msgId, "DTLS") || strings.HasPrefix(msgId, "DTLR"):
		info.MessageType = TypeEndpointDetailsReport
	case strings.HasPrefix(msgId, "GAPR"):
		info.MessageType = TypeEndpointGapReport
	case strings.HasPrefix(msgId, "ETOT"):
		info.MessageType = TypeEndpointTotalsReport
	case strings.HasPrefix(msgId, "ABAR"):
		info.MessageType = TypeMaster
	default:
		// Fallback to Rpt.Id analysis
		switch rptId {
		case "EDAY":
			info.MessageType = TypeActivityReport
		case "IMAD", "OMAD":
			info.MessageType = TypeEndpointGapReport
		case "IDAY":
			// Could be EndpointDetailsReport or EndpointTotalsReport
			// Need more context, default to EndpointDetailsReport
			info.MessageType = TypeEndpointDetailsReport
		case "ABMS":
			info.MessageType = TypeMaster
		default:
			info.MessageType = TypeUnknown
			return info, fmt.Errorf("unable to determine BkToCstmrAcctRpt subtype: MsgId=%s, Rpt.Id=%s", msgId, rptId)
		}
	}

	return info, nil
}

// Read reads XML from an io.Reader and returns the parsed message
func (r *UniversalReader) Read(reader io.Reader) (*ParsedMessage, error) {
	// Read all data
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML data: %w", err)
	}

	return r.ReadBytes(data)
}

// ReadBytes reads XML from byte slice and returns the parsed message
func (r *UniversalReader) ReadBytes(data []byte) (*ParsedMessage, error) {
	// Peek at XML structure
	peek, err := r.peekXML(data)
	if err != nil {
		return nil, fmt.Errorf("failed to peek XML structure: %w", err)
	}

	// Detect message type
	detection, err := r.detectMessageType(peek, data)
	if err != nil {
		return nil, fmt.Errorf("failed to detect message type: %w", err)
	}

	// Parse based on detected type
	parsed := &ParsedMessage{
		Type:      detection.MessageType,
		Version:   detection.Version,
		Detection: *detection,
	}

	// Parse the actual message
	switch detection.MessageType {
	case TypeUnknown:
		return nil, fmt.Errorf("unsupported message type: %s", detection.MessageType)
	case TypeCustomerCreditTransfer:
		msg, err := CustomerCreditTransferModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypePaymentReturn:
		msg, err := PaymentReturnModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypePaymentStatusRequest:
		msg, err := PaymentStatusRequestModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeFedwireFundsPaymentStatus:
		msg, err := FedwireFundsPaymentStatusModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeDrawdownRequest:
		msg, err := DrawdownRequestModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeDrawdownResponse:
		msg, err := DrawdownResponseModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeAccountReportingRequest:
		msg, err := AccountReportingRequestModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeActivityReport:
		msg, err := ActivityReportModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeEndpointDetailsReport:
		msg, err := EndpointDetailsReportModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeEndpointGapReport:
		msg, err := EndpointGapReportModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeEndpointTotalsReport:
		msg, err := EndpointTotalsReportModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeReturnRequestResponse:
		msg, err := ReturnRequestResponseModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeConnectionCheck:
		msg, err := ConnectionCheckModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeFedwireFundsAcknowledgement:
		msg, err := FedwireFundsAcknowledgementModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeFedwireFundsSystemResponse:
		msg, err := FedwireFundsSystemResponseModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	case TypeMaster:
		msg, err := MasterModel.ParseXML(data)
		if err != nil {
			return nil, r.enhanceError(err, detection, data)
		}
		parsed.Message = msg

	default:
		return nil, fmt.Errorf("unsupported message type: %s", detection.MessageType)
	}

	return parsed, nil
}

// enhanceError adds context to parsing/validation errors for debugging
func (r *UniversalReader) enhanceError(err error, detection *DetectionInfo, data []byte) error {
	if !r.VerboseErrors {
		return err
	}

	// Build enhanced error with context
	var enhanced strings.Builder
	enhanced.WriteString(fmt.Sprintf("Failed to parse %s message:\n", detection.MessageType))
	enhanced.WriteString(fmt.Sprintf("  Original error: %v\n", err))
	enhanced.WriteString(fmt.Sprintf("  Root element: %s\n", detection.RootElement))
	enhanced.WriteString(fmt.Sprintf("  Namespace: %s\n", detection.Namespace))
	enhanced.WriteString(fmt.Sprintf("  Version: %s\n", detection.Version))
	enhanced.WriteString(fmt.Sprintf("  Detection method: %s\n", detection.DetectedBy))

	// Add additional context
	if len(detection.AdditionalInfo) > 0 {
		enhanced.WriteString("  Additional info:\n")
		for k, v := range detection.AdditionalInfo {
			enhanced.WriteString(fmt.Sprintf("    %s: %s\n", k, v))
		}
	}

	// If it's a validation error, try to extract field path
	if validationErr, ok := err.(*errors.ValidationError); ok {
		enhanced.WriteString(fmt.Sprintf("  Validation field: %s\n", validationErr.Field))
		enhanced.WriteString(fmt.Sprintf("  Validation reason: %s\n", validationErr.Reason))
	}

	// Add line number context if available and tracking is enabled
	if r.TrackLineNumbers {
		// This would require more sophisticated XML parsing with line tracking
		// For now, just note that it's available in verbose mode
		enhanced.WriteString("  Enable XML line tracking for detailed position info\n")
	}

	return fmt.Errorf(enhanced.String())
}

// ValidateMessage validates a parsed message (optional since parsing already validates)
// This method is provided for cases where you want to re-validate after modifying a message.
// Note: The ParseXML methods already perform validation during parsing.
func (r *UniversalReader) ValidateMessage(parsed *ParsedMessage) error {
	if parsed == nil || parsed.Message == nil {
		return fmt.Errorf("no message to validate")
	}

	// Messages are already validated during parsing by the ProcessMessage method
	// in the base processor, so this is mainly for re-validation after modifications.
	// Since the models don't expose a simple Validate() method (they use ValidateForVersion),
	// and we don't know the specific version requirements here, we'll just verify
	// the message structure is intact.

	switch parsed.Message.(type) {
	case *CustomerCreditTransferModel.MessageModel,
		*PaymentReturnModel.MessageModel,
		*PaymentStatusRequestModel.MessageModel,
		*FedwireFundsPaymentStatusModel.MessageModel,
		*DrawdownRequestModel.MessageModel,
		*DrawdownResponseModel.MessageModel,
		*AccountReportingRequestModel.MessageModel,
		*ActivityReportModel.MessageModel,
		*EndpointDetailsReportModel.MessageModel,
		*EndpointGapReportModel.MessageModel,
		*EndpointTotalsReportModel.MessageModel,
		*ReturnRequestResponseModel.MessageModel,
		*ConnectionCheckModel.MessageModel,
		*FedwireFundsAcknowledgementModel.MessageModel,
		*FedwireFundsSystemResponseModel.MessageModel,
		*MasterModel.MessageModel:
		// Message is of a known type and was already validated during parsing
		return nil
	default:
		return fmt.Errorf("unknown message type for validation: %T", parsed.Message)
	}
}
