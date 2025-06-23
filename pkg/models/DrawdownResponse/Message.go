package DrawdownResponse

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"time"

	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_01"
	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_02"
	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_03"
	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_04"
	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_05"
	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_06"
	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_07"
	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_08"
	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_09"
	"github.com/moov-io/fedwire20022/gen/DrawdownResponse/pain_014_001_10"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
)

// AddressEnhancementFields available in V7+ versions
type AddressEnhancementFields struct {
	// RoomNumber fields are added to existing address structures
	// These are handled through the PartyIdentify.Address.RoomNumber fields
	// This struct serves as a marker for V7+ capabilities
}

// Validate checks if address enhancement fields meet requirements
func (a *AddressEnhancementFields) Validate() error {
	return nil
}

// NewMessageForVersion creates a MessageModel with appropriate version-specific fields initialized
func NewMessageForVersion(version PAIN_014_001_VERSION) MessageModel {
	model := MessageModel{
		MessageHeader: base.MessageHeader{},
		// Core fields initialized to zero values
	}

	// Type-safe version-specific field initialization
	switch {
	case version >= PAIN_014_001_07:
		model.AddressEnhancement = &AddressEnhancementFields{}
	}

	return model
}

// ValidateForVersion performs type-safe validation for a specific version
func (m MessageModel) ValidateForVersion(version PAIN_014_001_VERSION) error {
	// Base field validation (always required)
	if err := m.validateCoreFields(); err != nil {
		return fmt.Errorf("core field validation failed: %w", err)
	}

	// Type-safe version-specific validation
	switch {
	case version >= PAIN_014_001_07:
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
	if m.MessageId == "" {
		return fmt.Errorf("MessageId is required")
	}
	if m.CreatedDateTime.IsZero() {
		return fmt.Errorf("CreatedDateTime is required")
	}
	if m.OriginalMessageId == "" {
		return fmt.Errorf("OriginalMessageId is required")
	}
	return nil
}

// GetVersionCapabilities returns which version-specific features are available
func (m MessageModel) GetVersionCapabilities() map[string]bool {
	return map[string]bool{
		"AddressEnhancement": m.AddressEnhancement != nil,
	}
}

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// Core fields present in all versions
	InitiatingParty                 models.PartyIdentify     `json:"initiatingParty"`
	DebtorAgent                     models.Agent             `json:"debtorAgent"`
	CreditorAgent                   models.Agent             `json:"creditorAgent"`
	OriginalMessageId               string                   `json:"originalMessageId"`
	OriginalMessageNameId           string                   `json:"originalMessageNameId"`
	OriginalCreationDateTime        time.Time                `json:"originalCreationDateTime"`
	OriginalPaymentInfoId           string                   `json:"originalPaymentInfoId"`
	TransactionInformationAndStatus TransactionInfoAndStatus `json:"transactionInformationAndStatus"`

	// Version-specific field groups (type-safe, nil when not applicable)
	AddressEnhancement *AddressEnhancementFields `json:",inline,omitempty"` // V7+ only
}

// UnmarshalJSON implements custom JSON unmarshaling to properly handle grouped fields
func (m *MessageModel) UnmarshalJSON(data []byte) error {
	// Create an alias to avoid recursion
	type Alias MessageModel

	// Unmarshal into the aliased structure normally
	var temp Alias
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Copy all fields
	*m = MessageModel(temp)

	// For DrawdownResponse, AddressEnhancement is primarily a marker for V7+ capabilities
	// The actual RoomNumber fields are handled through existing Address structures
	// No specific inline field initialization needed for this case

	return nil
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, PAIN_014_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, PAIN_014_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.01",
			Version:   PAIN_014_001_01,
			Factory: func() models.ISODocument {
				return &pain_014_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_01], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.02",
			Version:   PAIN_014_001_02,
			Factory: func() models.ISODocument {
				return &pain_014_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_02], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.03",
			Version:   PAIN_014_001_03,
			Factory: func() models.ISODocument {
				return &pain_014_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.04",
			Version:   PAIN_014_001_04,
			Factory: func() models.ISODocument {
				return &pain_014_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.05",
			Version:   PAIN_014_001_05,
			Factory: func() models.ISODocument {
				return &pain_014_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.06",
			Version:   PAIN_014_001_06,
			Factory: func() models.ISODocument {
				return &pain_014_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_06], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.07",
			Version:   PAIN_014_001_07,
			Factory: func() models.ISODocument {
				return &pain_014_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_07], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.08",
			Version:   PAIN_014_001_08,
			Factory: func() models.ISODocument {
				return &pain_014_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_08], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.09",
			Version:   PAIN_014_001_09,
			Factory: func() models.ISODocument {
				return &pain_014_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_09], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.014.001.10",
			Version:   PAIN_014_001_10,
			Factory: func() models.ISODocument {
				return &pain_014_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_10], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, PAIN_014_001_VERSION](
		versionedFactory.BuildNameSpaceModelMap(),
		versionedFactory.GetVersionMap(),
		VersionPathMap,
		RequiredFields,
	)
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
//	err := model.WriteXML(file, DrawdownResponse.{VERSION_CONST})
//
//	// Write to buffer
//	var buf bytes.Buffer
//	err := model.WriteXML(&buf)
//
// For advanced use cases requiring document inspection before serialization, see DocumentWith.
func (m *MessageModel) WriteXML(w io.Writer, version ...PAIN_014_001_VERSION) error {
	// Default to latest version
	ver := PAIN_014_001_10
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
	"MessageId", "CreatedDateTime", "InitiatingParty", "DebtorAgent", "CreditorAgent", "OriginalMessageId",
	"OriginalMessageNameId", "OriginalCreationDateTime", "OriginalPaymentInfoId", "TransactionInformationAndStatus",
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
//	doc, err := DrawdownResponse.DocumentWith(model, VERSION_LATEST)
//	if err != nil {
//	    return err
//	}
//	// Now you can inspect or modify doc before serializing
//	xmlBytes, err := xml.Marshal(doc)
func DocumentWith(model MessageModel, version PAIN_014_001_VERSION) (models.ISODocument, error) {
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
