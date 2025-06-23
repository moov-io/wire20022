package ActivityReport

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"time"

	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_01"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_02"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_03"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_04"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_05"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_06"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_07"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_08"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_09"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_10"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_11"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_12"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
)

// AccountEnhancementFields available in V2+ versions
type AccountEnhancementFields struct {
	AccountOtherId string `json:"accountOtherId"`
}

// Validate checks if account enhancement fields meet requirements
func (a *AccountEnhancementFields) Validate() error {
	// AccountOtherId is optional but should be non-empty if present
	return nil
}

// NewMessageForVersion creates a MessageModel with appropriate version-specific fields initialized
func NewMessageForVersion(version CAMT_052_001_VERSION) MessageModel {
	model := MessageModel{
		MessageHeader: base.MessageHeader{},
		// Core fields initialized to zero values
	}

	// Type-safe version-specific field initialization
	switch {
	case version >= CAMT_052_001_02:
		model.AccountEnhancement = &AccountEnhancementFields{}
	}

	return model
}

// ValidateForVersion performs type-safe validation for a specific version
func (m MessageModel) ValidateForVersion(version CAMT_052_001_VERSION) error {
	// Base field validation (always required)
	if err := m.validateCoreFields(); err != nil {
		return fmt.Errorf("core field validation failed: %w", err)
	}

	// Type-safe version-specific validation
	switch {
	case version >= CAMT_052_001_02:
		if m.AccountEnhancement == nil {
			return fmt.Errorf("AccountEnhancementFields required for version %v but not present", version)
		}
		if err := m.AccountEnhancement.Validate(); err != nil {
			return fmt.Errorf("AccountEnhancementFields validation failed: %w", err)
		}
	}

	return nil
}

// validateCoreFields checks required core fields present in all versions
func (m MessageModel) validateCoreFields() error {
	// Direct field access - compile-time verified, no reflection
	if m.CreatedDateTime.IsZero() {
		return fmt.Errorf("CreatedDateTime is required")
	}
	if m.ReportId == "" {
		return fmt.Errorf("ReportId is required")
	}
	if m.ReportCreateDateTime.IsZero() {
		return fmt.Errorf("ReportCreateDateTime is required")
	}
	return nil
}

// GetVersionCapabilities returns which version-specific features are available
func (m MessageModel) GetVersionCapabilities() map[string]bool {
	return map[string]bool{
		"AccountEnhancement": m.AccountEnhancement != nil,
	}
}

// MessageModel uses base abstractions for common fields but keeps custom logic for complex arrays
type MessageModel struct {
	// Use base abstraction for common header fields
	base.MessageHeader `json:",inline"`

	// Override MessageId with proper type for ActivityReport
	MessageId models.CAMTReportType `json:"messageId"`

	// Core fields present in all versions
	Pagenation                         models.MessagePagenation              `json:"pagenation"`
	ReportId                           models.ReportType                     `json:"reportId"`
	ReportCreateDateTime               time.Time                             `json:"reportCreateDateTime"`
	TotalEntries                       string                                `json:"totalEntries"`
	TotalCreditEntries                 models.NumberAndSumOfTransactions     `json:"totalCreditEntries"`
	TotalDebitEntries                  models.NumberAndSumOfTransactions     `json:"totalDebitEntries"`
	TotalEntriesPerBankTransactionCode []models.TotalsPerBankTransactionCode `json:"totalEntriesPerBankTransactionCode"`
	EntryDetails                       []models.Entry                        `json:"entryDetails"`

	// Version-specific field groups (type-safe, nil when not applicable)
	AccountEnhancement *AccountEnhancementFields `json:",inline,omitempty"` // V2+ only
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
	if _, hasAccountOtherId := rawMap["accountOtherId"]; hasAccountOtherId {
		if m.AccountEnhancement == nil {
			m.AccountEnhancement = &AccountEnhancementFields{}
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
//	err := model.WriteXML(file, ActivityReport.{VERSION_CONST})
//
//	// Write to buffer
//	var buf bytes.Buffer
//	err := model.WriteXML(&buf)
//
// For advanced use cases requiring document inspection before serialization, see DocumentWith.
func (m *MessageModel) WriteXML(w io.Writer, version ...CAMT_052_001_VERSION) error {
	// Default to latest version
	ver := CAMT_052_001_12
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
	"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime",
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, CAMT_052_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, CAMT_052_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.01",
			Version:   CAMT_052_001_01,
			Factory: func() models.ISODocument {
				return &camt_052_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_01], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.02",
			Version:   CAMT_052_001_02,
			Factory: func() models.ISODocument {
				return &camt_052_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_02], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.03",
			Version:   CAMT_052_001_03,
			Factory: func() models.ISODocument {
				return &camt_052_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.04",
			Version:   CAMT_052_001_04,
			Factory: func() models.ISODocument {
				return &camt_052_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.05",
			Version:   CAMT_052_001_05,
			Factory: func() models.ISODocument {
				return &camt_052_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.06",
			Version:   CAMT_052_001_06,
			Factory: func() models.ISODocument {
				return &camt_052_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_06], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.07",
			Version:   CAMT_052_001_07,
			Factory: func() models.ISODocument {
				return &camt_052_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_07], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08",
			Version:   CAMT_052_001_08,
			Factory: func() models.ISODocument {
				return &camt_052_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_08], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.09",
			Version:   CAMT_052_001_09,
			Factory: func() models.ISODocument {
				return &camt_052_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_09], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.10",
			Version:   CAMT_052_001_10,
			Factory: func() models.ISODocument {
				return &camt_052_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_10], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.11",
			Version:   CAMT_052_001_11,
			Factory: func() models.ISODocument {
				return &camt_052_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_11], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.052.001.12",
			Version:   CAMT_052_001_12,
			Factory: func() models.ISODocument {
				return &camt_052_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_12], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, CAMT_052_001_VERSION](
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
//	doc, err := ActivityReport.DocumentWith(model, VERSION_LATEST)
//	if err != nil {
//	    return err
//	}
//	// Now you can inspect or modify doc before serializing
//	xmlBytes, err := xml.Marshal(doc)
//
// DocumentWith creates a versioned ISO 20022 document from the MessageModel.
// It validates required fields before creating the document and returns an error
// if validation fails or if the specified version is not supported.
func DocumentWith(model MessageModel, version CAMT_052_001_VERSION) (models.ISODocument, error) {
	// Validate required fields before creating document
	if err := processor.ValidateRequiredFields(model); err != nil {
		return nil, err
	}
	return processor.CreateDocument(model, version)
}

// CheckRequiredFields uses base abstractions to replace 15+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}
