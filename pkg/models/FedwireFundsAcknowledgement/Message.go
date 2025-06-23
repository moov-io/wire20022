package FedwireFundsAcknowledgement

import (
	"encoding/xml"

	"fmt"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsAcknowledgement/admi_007_001_01"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"io"
)

// NewMessageForVersion creates a MessageModel with appropriate version-specific fields initialized
func NewMessageForVersion(version ADMI_007_001_VERSION) MessageModel {
	model := MessageModel{
		MessageHeader: base.MessageHeader{},
		// Core fields initialized to zero values
	}
	
	// No version-specific fields for FedwireFundsAcknowledgement - single version message
	
	return model
}

// ValidateForVersion performs type-safe validation for a specific version
func (m MessageModel) ValidateForVersion(version ADMI_007_001_VERSION) error {
	// Base field validation (always required)
	if err := m.validateCoreFields(); err != nil {
		return fmt.Errorf("core field validation failed: %w", err)
	}
	
	// No version-specific validation needed - single version message
	
	return nil
}

// validateCoreFields checks required core fields present in all versions
func (m MessageModel) validateCoreFields() error {
	// Direct field access - compile-time verified, no reflection
	if m.MessageId == "" {
		return fmt.Errorf("MessageId is required")
	}
	if m.CreatedDateTime.IsZero() {
		return fmt.Errorf("CreatedDateTime is required")
	}
	if m.RelationReference == "" {
		return fmt.Errorf("RelationReference is required")
	}
	if m.ReferenceName == "" {
		return fmt.Errorf("ReferenceName is required")
	}
	if m.RequestHandling == "" {
		return fmt.Errorf("RequestHandling is required")
	}
	return nil
}

// GetVersionCapabilities returns which version-specific features are available
func (m MessageModel) GetVersionCapabilities() map[string]bool {
	return map[string]bool{
		// No version-specific capabilities for FedwireFundsAcknowledgement
	}
}

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// FedwireFundsAcknowledgement-specific fields
	RelationReference string                   `json:"relationReference"`
	ReferenceName     string                   `json:"referenceName"`
	RequestHandling   models.RelatedStatusCode `json:"requestHandling"`
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
//	err := model.WriteXML(file, FedwireFundsAcknowledgement.{VERSION_CONST})
//
//	// Write to buffer
//	var buf bytes.Buffer
//	err := model.WriteXML(&buf)
//
// For advanced use cases requiring document inspection before serialization, see DocumentWith.
func (m *MessageModel) WriteXML(w io.Writer, version ...ADMI_007_001_VERSION) error {
	// Default to latest version
	ver := ADMI_007_001_01
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
	"MessageId", "CreatedDateTime", "RelationReference", "ReferenceName", "RequestHandling",
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, ADMI_007_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register version using factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, ADMI_007_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:admi.007.001.01",
			Version:   ADMI_007_001_01,
			Factory: func() models.ISODocument {
				return &admi_007_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[ADMI_007_001_01], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, ADMI_007_001_VERSION](
		versionedFactory.BuildNameSpaceModelMap(),
		versionedFactory.GetVersionMap(),
		VersionPathMap,
		RequiredFields,
	)
}

// ParseXML reads XML data into the MessageModel
// This replaces the non-idiomatic MessageWith function
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
//	doc, err := FedwireFundsAcknowledgement.DocumentWith(model, VERSION_LATEST)
//	if err != nil {
//	    return err
//	}
//	// Now you can inspect or modify doc before serializing
//	xmlBytes, err := xml.Marshal(doc)
func DocumentWith(model MessageModel, version ADMI_007_001_VERSION) (models.ISODocument, error) {
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
