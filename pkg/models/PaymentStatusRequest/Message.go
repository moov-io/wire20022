package PaymentStatusRequest

import (
	"encoding/xml"
	"time"

	"fmt"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_01"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_02"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_03"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_04"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_05"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_06"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"io"
)

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// PaymentStatusRequest-specific fields
	OriginalMessageId        string    `json:"originalMessageId"`
	OriginalMessageNameId    string    `json:"originalMessageNameId"`
	OriginalCreationDateTime time.Time `json:"originalCreationDateTime"`
	OriginalInstructionId    string    `json:"originalInstructionId"`
	OriginalEndToEndId       string    `json:"originalEndToEndId"`
	OriginalUETR             string    `json:"originalUETR"`

	// Use embedded agent pairs
	base.AgentPair `json:",inline"`
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
//	err := model.WriteXML(file, PaymentStatusRequest.{VERSION_CONST})
//
//	// Write to buffer
//	var buf bytes.Buffer
//	err := model.WriteXML(&buf)
//
// For advanced use cases requiring document inspection before serialization, see DocumentWith.
func (m *MessageModel) WriteXML(w io.Writer, version ...PACS_028_001_VERSION) error {
	// Default to latest version
	ver := PACS_028_001_05
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
	"MessageId", "CreatedDateTime", "OriginalMessageId",
	"OriginalMessageNameId", "OriginalCreationDateTime",
	"InstructingAgent", "InstructedAgent",
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, PACS_028_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, PACS_028_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.01",
			Version:   PACS_028_001_01,
			Factory: func() models.ISODocument {
				return &pacs_028_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_01], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.02",
			Version:   PACS_028_001_02,
			Factory: func() models.ISODocument {
				return &pacs_028_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_02], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03",
			Version:   PACS_028_001_03,
			Factory: func() models.ISODocument {
				return &pacs_028_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04",
			Version:   PACS_028_001_04,
			Factory: func() models.ISODocument {
				return &pacs_028_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.05",
			Version:   PACS_028_001_05,
			Factory: func() models.ISODocument {
				return &pacs_028_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.06",
			Version:   PACS_028_001_06,
			Factory: func() models.ISODocument {
				return &pacs_028_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_06], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, PACS_028_001_VERSION](
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
//	doc, err := PaymentStatusRequest.DocumentWith(model, VERSION_LATEST)
//	if err != nil {
//	    return err
//	}
//	// Now you can inspect or modify doc before serializing
//	xmlBytes, err := xml.Marshal(doc)
func DocumentWith(model MessageModel, version PACS_028_001_VERSION) (models.ISODocument, error) {
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
