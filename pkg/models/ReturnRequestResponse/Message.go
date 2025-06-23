package ReturnRequestResponse

import (
	"encoding/xml"
	"time"

	"fmt"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_03"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_04"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_05"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_06"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_07"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_08"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_09"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_10"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_11"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_12"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"io"
)

// Enhanced Transaction fields available in V9+ versions
type EnhancedTransactionFields struct {
	OriginalUETR string `json:"originalUETR"`
}

// Validate checks if enhanced transaction fields meet requirements
func (e *EnhancedTransactionFields) Validate() error {
	// OriginalUETR field is optional but should be valid if present
	return nil
}

// Address Enhancement fields available in V9+ versions
type AddressEnhancementFields struct {
	// Enhanced address fields for Creator agent
	// These are handled through the Creator.PostalAddress.BuildingName, Floor, RoomNumber fields
	// This struct serves as a marker for V9+ capabilities
}

// Validate checks if address enhancement fields meet requirements
func (a *AddressEnhancementFields) Validate() error {
	// Address enhancement fields are optional
	return nil
}

// NewMessageForVersion creates a MessageModel with appropriate version-specific fields initialized
func NewMessageForVersion(version CAMT_029_001_VERSION) MessageModel {
	model := MessageModel{
		// Core fields initialized to zero values
	}

	// Type-safe version-specific field initialization
	switch {
	case version >= CAMT_029_001_09:
		model.EnhancedTransaction = &EnhancedTransactionFields{}
		model.AddressEnhancement = &AddressEnhancementFields{}
	}

	return model
}

// ValidateForVersion performs type-safe validation for a specific version
func (m MessageModel) ValidateForVersion(version CAMT_029_001_VERSION) error {
	// Base field validation (always required)
	if err := m.validateCoreFields(); err != nil {
		return fmt.Errorf("core field validation failed: %w", err)
	}

	// Type-safe version-specific validation
	switch {
	case version >= CAMT_029_001_09:
		if m.EnhancedTransaction == nil {
			return fmt.Errorf("EnhancedTransactionFields required for version %v but not present", version)
		}
		if err := m.EnhancedTransaction.Validate(); err != nil {
			return fmt.Errorf("EnhancedTransactionFields validation failed: %w", err)
		}
		if m.AddressEnhancement == nil {
			return fmt.Errorf("AddressEnhancementFields required for version %v but not present", version)
		}
		if err := m.AddressEnhancement.Validate(); err != nil {
			return fmt.Errorf("AddressEnhancementFields validation failed: %w", err)
		}
	}

	return nil
}

// validateCoreFields checks required core fields present in all versions
func (m MessageModel) validateCoreFields() error {
	// Direct field access - compile-time verified, no reflection
	if m.AssignmentId == "" {
		return fmt.Errorf("AssignmentId is required")
	}
	if m.AssignmentCreateTime.IsZero() {
		return fmt.Errorf("AssignmentCreateTime is required")
	}
	if m.ResolvedCaseId == "" {
		return fmt.Errorf("ResolvedCaseId is required")
	}
	if m.OriginalMessageId == "" {
		return fmt.Errorf("OriginalMessageId is required")
	}
	if m.OriginalMessageNameId == "" {
		return fmt.Errorf("OriginalMessageNameId is required")
	}
	if m.OriginalMessageCreateTime.IsZero() {
		return fmt.Errorf("OriginalMessageCreateTime is required")
	}
	return nil
}

// GetVersionCapabilities returns which version-specific features are available
func (m MessageModel) GetVersionCapabilities() map[string]bool {
	return map[string]bool{
		"EnhancedTransaction": m.EnhancedTransaction != nil,
		"AddressEnhancement":  m.AddressEnhancement != nil,
	}
}

// MessageModel uses base abstractions to eliminate duplicate field definitions
// (Pattern 3 - Direct Migration with unique assignment-based structure)
type MessageModel struct {
	// Core fields present in all versions (V3+)
	AssignmentId                 string        `json:"assignmentId"`
	Assigner                     models.Agent  `json:"assigner"`
	Assignee                     models.Agent  `json:"assignee"`
	AssignmentCreateTime         time.Time     `json:"assignmentCreateTime"`
	ResolvedCaseId               string        `json:"resolvedCaseId"`
	Creator                      models.Agent  `json:"creator"`
	Status                       models.Status `json:"status"`
	OriginalMessageId            string        `json:"originalMessageId"`
	OriginalMessageNameId        string        `json:"originalMessageNameId"`
	OriginalMessageCreateTime    time.Time     `json:"originalMessageCreateTime"`
	OriginalInstructionId        string        `json:"originalInstructionId"`
	OriginalEndToEndId           string        `json:"originalEndToEndId"`
	CancellationStatusReasonInfo models.Reason `json:"cancellationStatusReasonInfo"`

	// Version-specific field groups (type-safe, nil when not applicable)
	EnhancedTransaction *EnhancedTransactionFields `json:",inline,omitempty"` // V9+ only
	AddressEnhancement  *AddressEnhancementFields  `json:",inline,omitempty"` // V9+ only
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, CAMT_029_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, CAMT_029_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.03",
			Version:   CAMT_029_001_03,
			Factory: func() models.ISODocument {
				return &camt_029_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.04",
			Version:   CAMT_029_001_04,
			Factory: func() models.ISODocument {
				return &camt_029_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.05",
			Version:   CAMT_029_001_05,
			Factory: func() models.ISODocument {
				return &camt_029_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.06",
			Version:   CAMT_029_001_06,
			Factory: func() models.ISODocument {
				return &camt_029_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_06], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.07",
			Version:   CAMT_029_001_07,
			Factory: func() models.ISODocument {
				return &camt_029_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_07], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.08",
			Version:   CAMT_029_001_08,
			Factory: func() models.ISODocument {
				return &camt_029_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_08], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09",
			Version:   CAMT_029_001_09,
			Factory: func() models.ISODocument {
				return &camt_029_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_09], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.10",
			Version:   CAMT_029_001_10,
			Factory: func() models.ISODocument {
				return &camt_029_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_10], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.11",
			Version:   CAMT_029_001_11,
			Factory: func() models.ISODocument {
				return &camt_029_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_11], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.12",
			Version:   CAMT_029_001_12,
			Factory: func() models.ISODocument {
				return &camt_029_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_12], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, CAMT_029_001_VERSION](
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
//	doc, err := ReturnRequestResponse.DocumentWith(model, VERSION_LATEST)
//	if err != nil {
//	    return err
//	}
//	// Now you can inspect or modify doc before serializing
//	xmlBytes, err := xml.Marshal(doc)
func DocumentWith(model MessageModel, version CAMT_029_001_VERSION) (models.ISODocument, error) {
	// Validate required fields before creating document
	if err := processor.ValidateRequiredFields(model); err != nil {
		return nil, err
	}
	return processor.CreateDocument(model, version)
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
//	err := model.WriteXML(file, ReturnRequestResponse.{VERSION_CONST})
//
//	// Write to buffer
//	var buf bytes.Buffer
//	err := model.WriteXML(&buf)
//
// For advanced use cases requiring document inspection before serialization, see DocumentWith.
func (m *MessageModel) WriteXML(w io.Writer, version ...CAMT_029_001_VERSION) error {
	// Default to latest version
	ver := CAMT_029_001_12
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
	defer encoder.Close()
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
	"AssignmentId", "Assigner", "Assignee",
	"AssignmentCreateTime", "ResolvedCaseId", "Creator", "OriginalMessageId",
	"OriginalMessageNameId", "OriginalMessageCreateTime",
}

// CheckRequiredFields uses base abstractions to replace 20+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}
