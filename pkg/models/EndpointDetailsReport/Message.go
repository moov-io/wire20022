package EndpointDetailsReport

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"time"

	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_02"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_03"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_04"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_05"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_06"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_07"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_08"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_09"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_10"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_11"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_12"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
)

// MessageModel uses base abstractions to eliminate duplicate field definitions
// Fields are organized by version introduction to clearly show version evolution
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// Core fields present in all versions (V1+)
	Pagenation                         models.MessagePagenation              `json:"pagenation"`
	ReportId                           models.ReportType                     `json:"reportId"`
	ReportCreateDateTime               time.Time                             `json:"reportCreateDateTime"`
	AccountOtherId                     string                                `json:"accountOtherId"`
	TotalCreditEntries                 models.NumberAndSumOfTransactions     `json:"totalCreditEntries"`
	TotalDebitEntries                  models.NumberAndSumOfTransactions     `json:"totalDebitEntries"`
	TotalEntriesPerBankTransactionCode []models.TotalsPerBankTransactionCode `json:"totalEntriesPerBankTransactionCode"`
	EntryDetails                       []models.Entry                        `json:"entryDetails"`

	// Business Query fields (V3+ only - will be empty in V1/V2)
	BussinessQueryMsgId          string    `json:"bussinessQueryMsgId,omitempty"`
	BussinessQueryMsgNameId      string    `json:"bussinessQueryMsgNameId,omitempty"`
	BussinessQueryCreateDatetime time.Time `json:"bussinessQueryCreateDatetime,omitempty"`

	// Reporting Sequence fields (V7+ only - will be empty in V1-V6)
	ReportingSequence models.SequenceRange `json:"reportingSequence,omitempty"`
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
//	err := model.WriteXML(file, EndpointDetailsReport.{VERSION_CONST})
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

// Version-specific required fields map
var RequiredFieldsByVersion = map[CAMT_052_001_VERSION][]string{
	// V1/V2: Basic fields only
	CAMT_052_001_02: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},
	CAMT_052_001_03: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},

	// V3-V6: Add business query fields
	CAMT_052_001_04: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},
	CAMT_052_001_05: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},
	CAMT_052_001_06: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},

	// V7+: Add reporting sequence fields
	CAMT_052_001_07: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},
	CAMT_052_001_08: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},
	CAMT_052_001_09: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},
	CAMT_052_001_10: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},
	CAMT_052_001_11: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},
	CAMT_052_001_12: {"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime"},
}

// RequiredFields maintains backward compatibility with the base processor
var RequiredFields = RequiredFieldsByVersion[CAMT_052_001_12] // Use latest version as default

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
//	doc, err := EndpointDetailsReport.DocumentWith(model, VERSION_LATEST)
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

// CheckRequiredFieldsForVersion validates required fields for a specific version
// This provides version-aware validation that only checks fields that should be present in that version
func CheckRequiredFieldsForVersion(model MessageModel, version CAMT_052_001_VERSION) error {
	requiredFields, exists := RequiredFieldsByVersion[version]
	if !exists {
		// Fall back to latest version validation if version not found
		requiredFields = RequiredFields
	}

	// Use reflection to validate version-specific required fields
	return validateRequiredFieldsReflection(model, requiredFields)
}

// validateRequiredFieldsReflection performs version-aware validation using reflection
func validateRequiredFieldsReflection(model MessageModel, requiredFields []string) error {
	modelValue := reflect.ValueOf(model)
	modelType := reflect.TypeOf(model)

	for _, fieldName := range requiredFields {
		// Handle embedded base.MessageHeader fields
		var fieldValue reflect.Value
		var found bool

		// First check direct fields
		if _, ok := modelType.FieldByName(fieldName); ok {
			fieldValue = modelValue.FieldByName(fieldName)
			found = true
		} else {
			// Check embedded MessageHeader fields
			headerField := modelValue.FieldByName("MessageHeader")
			if headerField.IsValid() {
				if _, ok := headerField.Type().FieldByName(fieldName); ok {
					fieldValue = headerField.FieldByName(fieldName)
					found = true
				}
			}
		}

		if !found {
			return fmt.Errorf("field %s not found", fieldName)
		}

		// Check if field is empty based on type
		if isEmpty := isFieldEmpty(fieldValue); isEmpty {
			return fmt.Errorf("required field %s is empty", fieldName)
		}
	}

	return nil
}

// isFieldEmpty checks if a field value is considered empty
func isFieldEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.String() == ""
	case reflect.Struct:
		if value.Type().String() == "time.Time" {
			timeVal, ok := value.Interface().(time.Time)
			if !ok {
				return value.IsZero()
			}
			return timeVal.IsZero()
		}
		// For other structs, check if it's the zero value
		return value.IsZero()
	case reflect.Slice, reflect.Array:
		return value.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return value.IsNil()
	case reflect.Bool:
		return false // bool is never considered "empty" for validation
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0.0
	case reflect.Complex64, reflect.Complex128:
		return value.Complex() == 0
	case reflect.Chan, reflect.Func, reflect.Map, reflect.UnsafePointer:
		return value.IsNil()
	case reflect.Invalid:
		return true
	default:
		return value.IsZero()
	}
}

// GetVersionSpecificFields returns a map indicating which fields are available for each version
func GetVersionSpecificFields() map[string][]CAMT_052_001_VERSION {
	return map[string][]CAMT_052_001_VERSION{
		"BussinessQueryMsgId":          {CAMT_052_001_04, CAMT_052_001_05, CAMT_052_001_06, CAMT_052_001_07, CAMT_052_001_08, CAMT_052_001_09, CAMT_052_001_10, CAMT_052_001_11, CAMT_052_001_12},
		"BussinessQueryMsgNameId":      {CAMT_052_001_04, CAMT_052_001_05, CAMT_052_001_06, CAMT_052_001_07, CAMT_052_001_08, CAMT_052_001_09, CAMT_052_001_10, CAMT_052_001_11, CAMT_052_001_12},
		"BussinessQueryCreateDatetime": {CAMT_052_001_04, CAMT_052_001_05, CAMT_052_001_06, CAMT_052_001_07, CAMT_052_001_08, CAMT_052_001_09, CAMT_052_001_10, CAMT_052_001_11, CAMT_052_001_12},
		"ReportingSequence":            {CAMT_052_001_07, CAMT_052_001_08, CAMT_052_001_09, CAMT_052_001_10, CAMT_052_001_11, CAMT_052_001_12},
	}
}
