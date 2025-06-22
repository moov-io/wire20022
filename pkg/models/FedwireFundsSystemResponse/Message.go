package FedwireFundsSystemResponse

import (
	"encoding/xml"
	"time"

	"fmt"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsSystemResponse/admi_011_001_01"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"io"
)

// MessageModel represents a non-standard message type (Pattern 3 - Direct Migration)
// Does not use base.MessageHeader as it follows event-based pattern instead of standard message pattern
type MessageModel struct {
	MessageId  string               `json:"messageId"`
	EventCode  models.FundEventType `json:"eventCode"`
	EventParam string               `json:"eventParam"`
	EventTime  time.Time            `json:"eventTime"`
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

// WriteXML writes the MessageModel as XML to an io.Writer
// If no version is specified, uses the latest version (ADMI_011_001_01)
func (m *MessageModel) WriteXML(w io.Writer, version ...ADMI_011_001_VERSION) error {
	// Default to latest version
	ver := ADMI_011_001_01
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
	"MessageId", "EventCode", "EventParam", "EventTime",
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, ADMI_011_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, ADMI_011_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:admi.011.001.01",
			Version:   ADMI_011_001_01,
			Factory: func() models.ISODocument {
				return &admi_011_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[ADMI_011_001_01], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, ADMI_011_001_VERSION](
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

// DocumentWith uses base abstractions to replace 20+ lines with a single call
func DocumentWith(model MessageModel, version ADMI_011_001_VERSION) (models.ISODocument, error) {
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
