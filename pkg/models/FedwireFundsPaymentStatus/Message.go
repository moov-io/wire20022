package FedwireFundsPaymentStatus

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"fmt"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_03"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_04"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_05"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_06"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_07"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_08"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_09"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_10"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_11"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_12"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_13"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_14"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"io"
)

// EnhancedTransactionFields available in V10+ versions
type EnhancedTransactionFields struct {
	OriginalUETR                     string          `json:"originalUETR"`
	EffectiveInterbankSettlementDate fedwire.ISODate `json:"effectiveInterbankSettlementDate"`
}

// Validate checks if enhanced transaction fields meet requirements
func (e *EnhancedTransactionFields) Validate() error {
	if e.OriginalUETR == "" {
		return fmt.Errorf("OriginalUETR is required for versions V10+")
	}
	return nil
}

// NewMessageForVersion creates a MessageModel with appropriate version-specific fields initialized
func NewMessageForVersion(version PACS_002_001_VERSION) MessageModel {
	model := MessageModel{
		MessageHeader: base.MessageHeader{},
		// Core fields initialized to zero values
	}

	// Type-safe version-specific field initialization
	switch {
	case version >= PACS_002_001_10:
		model.EnhancedTransaction = &EnhancedTransactionFields{}
	}

	return model
}

// ValidateForVersion performs type-safe validation for a specific version
func (m MessageModel) ValidateForVersion(version PACS_002_001_VERSION) error {
	// Base field validation (always required for V5+)
	if err := m.validateCoreFields(); err != nil {
		return fmt.Errorf("core field validation failed: %w", err)
	}

	// Type-safe version-specific validation
	switch {
	case version >= PACS_002_001_10:
		if m.EnhancedTransaction == nil {
			return fmt.Errorf("EnhancedTransactionFields required for version %v but not present", version)
		}
		if err := m.EnhancedTransaction.Validate(); err != nil {
			return fmt.Errorf("EnhancedTransactionFields validation failed: %w", err)
		}
	}

	return nil
}

// validateCoreFields checks required core fields present in V5+ versions
func (m MessageModel) validateCoreFields() error {
	// Direct field access - compile-time verified, no reflection
	if m.MessageId == "" {
		return fmt.Errorf("MessageId is required")
	}
	if m.CreatedDateTime.IsZero() {
		return fmt.Errorf("CreatedDateTime is required")
	}
	if m.OriginalMessageId == "" {
		return fmt.Errorf("OriginalMessageId is required")
	}
	if m.TransactionStatus == "" {
		return fmt.Errorf("TransactionStatus is required")
	}
	return nil
}

// GetVersionCapabilities returns which version-specific features are available
func (m MessageModel) GetVersionCapabilities() map[string]bool {
	return map[string]bool{
		"EnhancedTransaction": m.EnhancedTransaction != nil,
	}
}

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// Core fields present in all versions (V5+)
	OriginalMessageId         string                       `json:"originalMessageId"`
	OriginalMessageNameId     string                       `json:"originalMessageNameId"`
	OriginalMessageCreateTime time.Time                    `json:"originalMessageCreateTime"`
	TransactionStatus         models.TransactionStatusCode `json:"transactionStatus"`
	AcceptanceDateTime        time.Time                    `json:"acceptanceDateTime"`
	StatusReasonInformation   string                       `json:"statusReasonInformation"`
	ReasonAdditionalInfo      string                       `json:"reasonAdditionalInfo"`

	// Version-specific field groups (type-safe, nil when not applicable)
	EnhancedTransaction *EnhancedTransactionFields `json:",inline,omitempty"` // V10+ only

	// Use embedded agent pairs
	base.AgentPair `json:",inline"`
}

// UnmarshalJSON implements custom JSON unmarshaling to properly handle grouped fields
func (m *MessageModel) UnmarshalJSON(data []byte) error {
	// Parse into a generic map first to check for inline fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return err
	}

	// Create an alias to avoid recursion
	type Alias MessageModel

	// Unmarshal into the aliased structure normally
	var temp Alias
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Copy all fields
	*m = MessageModel(temp)

	// Post-process: Initialize grouped fields based on presence of inline fields
	if _, hasOriginalUETR := rawMap["originalUETR"]; hasOriginalUETR {
		if m.EnhancedTransaction == nil {
			m.EnhancedTransaction = &EnhancedTransactionFields{}
		}
	}
	if _, hasEffectiveDate := rawMap["effectiveInterbankSettlementDate"]; hasEffectiveDate {
		if m.EnhancedTransaction == nil {
			m.EnhancedTransaction = &EnhancedTransactionFields{}
		}
	}

	return nil
}

// ReadXML reads XML data from an io.Reader into the MessageModel
func (m *MessageModel) ReadXML(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("reading XML: %w", err)
	}

	model, err := processor.ProcessMessage(data)
	if err != nil {
		return err
	}

	*m = model
	return nil
}

// WriteXML writes the MessageModel as XML to an io.Writer.
// This is the primary method for XML serialization and handles the complete XML generation process.
//
// Features:
//   - Writes XML declaration (<?xml version="1.0" encoding="UTF-8"?>)
//   - Properly formatted with indentation
//   - Automatic namespace handling
//   - Validates required fields before writing
//   - Defaults to latest version if not specified
//
// Example:
//
//	// Write to file
//	file, _ := os.Create("payment.xml")
//	defer file.Close()
//	err := model.WriteXML(file, FedwireFundsPaymentStatus.{VERSION_CONST})
//
//	// Write to buffer
//	var buf bytes.Buffer
//	err := model.WriteXML(&buf)
//
// For advanced use cases requiring document inspection before serialization, see DocumentWith.
func (m *MessageModel) WriteXML(w io.Writer, version ...PACS_002_001_VERSION) error {
	// Default to latest version
	ver := PACS_002_001_14
	if len(version) > 0 {
		ver = version[0]
	}

	// Create versioned document
	doc, err := DocumentWith(*m, ver)
	if err != nil {
		return fmt.Errorf("creating document: %w", err)
	}

	// Write XML with proper formatting
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")

	// Write XML declaration
	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return fmt.Errorf("writing XML header: %w", err)
	}

	// Encode document
	if err := encoder.Encode(doc); err != nil {
		return fmt.Errorf("encoding XML: %w", err)
	}

	return encoder.Flush()
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "TransactionStatus", "InstructingAgent", "InstructedAgent",
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, PACS_002_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, PACS_002_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.03",
			Version:   PACS_002_001_03,
			Factory: func() models.ISODocument {
				return &pacs_002_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.04",
			Version:   PACS_002_001_04,
			Factory: func() models.ISODocument {
				return &pacs_002_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.05",
			Version:   PACS_002_001_05,
			Factory: func() models.ISODocument {
				return &pacs_002_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.06",
			Version:   PACS_002_001_06,
			Factory: func() models.ISODocument {
				return &pacs_002_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_06], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.07",
			Version:   PACS_002_001_07,
			Factory: func() models.ISODocument {
				return &pacs_002_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_07], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08",
			Version:   PACS_002_001_08,
			Factory: func() models.ISODocument {
				return &pacs_002_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_08], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.09",
			Version:   PACS_002_001_09,
			Factory: func() models.ISODocument {
				return &pacs_002_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_09], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10",
			Version:   PACS_002_001_10,
			Factory: func() models.ISODocument {
				return &pacs_002_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_10], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.11",
			Version:   PACS_002_001_11,
			Factory: func() models.ISODocument {
				return &pacs_002_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_11], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12",
			Version:   PACS_002_001_12,
			Factory: func() models.ISODocument {
				return &pacs_002_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_12], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.13",
			Version:   PACS_002_001_13,
			Factory: func() models.ISODocument {
				return &pacs_002_001_13.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_13], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.14",
			Version:   PACS_002_001_14,
			Factory: func() models.ISODocument {
				return &pacs_002_001_14.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_14], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, PACS_002_001_VERSION](
		versionedFactory.BuildNameSpaceModelMap(),
		versionedFactory.GetVersionMap(),
		VersionPathMap,
		RequiredFields,
	)
}

// ParseXML reads XML data into the MessageModel
// This is the primary function for parsing XML from byte data
func ParseXML(data []byte) (*MessageModel, error) {
	model, err := processor.ProcessMessage(data)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

// DocumentWith creates a versioned ISO 20022 document from the MessageModel.
// This is a lower-level API that returns the raw document structure for advanced use cases.
//
// When to use DocumentWith vs WriteXML:
//   - Use WriteXML for standard XML output to files, network connections, or buffers
//   - Use DocumentWith when you need to:
//   - Inspect or modify the document structure before serialization
//   - Integrate with other XML processing libraries
//   - Perform custom validation on the document level
//   - Access version-specific document types directly
//
// Example:
//
//	doc, err := FedwireFundsPaymentStatus.DocumentWith(model, VERSION_LATEST)
//	if err != nil {
//	    return err
//	}
//	// Now you can inspect or modify doc before serializing
//	xmlBytes, err := xml.Marshal(doc)
//
// DocumentWith creates a versioned ISO 20022 document from the MessageModel.
// It validates required fields before creating the document and returns an error
// if validation fails or if the specified version is not supported.
func DocumentWith(model MessageModel, version PACS_002_001_VERSION) (models.ISODocument, error) {
	// Validate required fields before creating document
	if err := processor.ValidateRequiredFields(model); err != nil {
		return nil, err
	}
	return processor.CreateDocument(model, version)
}

// CheckRequiredFields uses base abstractions to replace 30+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}
