package PaymentStatusRequest

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_01"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_02"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_03"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_04"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_05"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_06"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
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

// MessageWith uses base abstractions to replace 15+ lines with a single call
func MessageWith(data []byte) (MessageModel, error) {
	return processor.ProcessMessage(data)
}

// DocumentWith uses base abstractions to replace 25+ lines with a single call
func DocumentWith(model MessageModel, version PACS_028_001_VERSION) (models.ISODocument, error) {
	return processor.CreateDocument(model, version)
}

// CheckRequiredFields uses base abstractions to replace 30+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}
