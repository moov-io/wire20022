package Master

import (
	"encoding/xml"
	"time"

	"fmt"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_02"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_03"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_04"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_05"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_06"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_07"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_08"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_09"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_10"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_11"
	"github.com/moov-io/fedwire20022/gen/Master/camt_052_001_12"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"io"
)

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// Master-specific fields
	MessagePagination             models.MessagePagenation          `json:"messagePagination"`
	OriginalBusinessMsgId         string                            `json:"originalBusinessMsgId"`
	OriginalBusinessMsgNameId     string                            `json:"originalBusinessMsgNameId"`
	OriginalBusinessMsgCreateTime time.Time                         `json:"originalBusinessMsgCreateTime"`
	ReportTypeId                  models.AccountReportType          `json:"reportTypeId"`
	ReportCreatedDate             time.Time                         `json:"reportCreatedDate"`
	AccountOtherId                string                            `json:"accountOtherId"`
	AccountType                   string                            `json:"accountType"`
	RelatedAccountOtherId         string                            `json:"relatedAccountOtherId"`
	Balances                      []models.Balance                  `json:"balances"`
	TransactionsSummary           []models.TotalsPerBankTransaction `json:"transactionsSummary"`
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
//	err := model.WriteXML(file, Master.{VERSION_CONST})
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
	"MessageId", "CreatedDateTime", "MessagePagination", "ReportTypeId", "ReportCreatedDate",
	"AccountOtherId", "AccountType", "RelatedAccountOtherId", "TransactionsSummary",
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, CAMT_052_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, CAMT_052_001_VERSION]{
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
//	doc, err := Master.DocumentWith(model, VERSION_LATEST)
//	if err != nil {
//	    return err
//	}
//	// Now you can inspect or modify doc before serializing
//	xmlBytes, err := xml.Marshal(doc)
func DocumentWith(model MessageModel, version CAMT_052_001_VERSION) (models.ISODocument, error) {
	// Validate required fields before creating document
	if err := processor.ValidateRequiredFields(model); err != nil {
		return nil, err
	}
	return processor.CreateDocument(model, version)
}

// CheckRequiredFields uses base abstractions to replace 20+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}
