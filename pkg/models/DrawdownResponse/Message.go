package DrawdownResponse

import (
	"encoding/xml"
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
	"fmt"
	"io"
)

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// DrawdownResponse-specific fields
	InitiatingParty                 models.PartyIdentify     `json:"initiatingParty"`
	DebtorAgent                     models.Agent             `json:"debtorAgent"`
	CreditorAgent                   models.Agent             `json:"creditorAgent"`
	OriginalMessageId               string                   `json:"originalMessageId"`
	OriginalMessageNameId           string                   `json:"originalMessageNameId"`
	OriginalCreationDateTime        time.Time                `json:"originalCreationDateTime"`
	OriginalPaymentInfoId           string                   `json:"originalPaymentInfoId"`
	TransactionInformationAndStatus TransactionInfoAndStatus `json:"transactionInformationAndStatus"`
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

// WriteXML writes the MessageModel as XML to an io.Writer
// If no version is specified, uses the latest version (PAIN_014_001_10)
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

// MessageWith uses base abstractions to replace 15+ lines with a single call
// ParseXML reads XML data into the MessageModel
// This replaces the non-idiomatic MessageWith function
func ParseXML(data []byte) (*MessageModel, error) {
	model, err := processor.ProcessMessage(data)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

// Deprecated: Use ParseXML instead
func MessageWith(data []byte) (MessageModel, error) {
	return processor.ProcessMessage(data)
}

// DocumentWith uses base abstractions to replace 25+ lines with a single call
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
