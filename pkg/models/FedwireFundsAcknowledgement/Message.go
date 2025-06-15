package FedwireFundsAcknowledgement

import (
	"encoding/xml"

	"github.com/moov-io/fedwire20022/gen/FedwireFundsAcknowledgement/admi_007_001_01"
	"github.com/wadearnold/wire20022/pkg/base"
	"github.com/wadearnold/wire20022/pkg/models"
)

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// FedwireFundsAcknowledgement-specific fields
	RelationReference string                   `json:"relationReference"`
	ReferenceName     string                   `json:"referenceName"`
	RequestHandling   models.RelatedStatusCode `json:"requestHandling"`
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

// MessageWith uses base abstractions to replace 15+ lines with a single call
func MessageWith(data []byte) (MessageModel, error) {
	return processor.ProcessMessage(data)
}

// DocumentWith uses base abstractions to replace 25+ lines with a single call
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
