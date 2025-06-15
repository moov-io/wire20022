package ConnectionCheck

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/ConnectionCheck/admi_004_001_01"
	"github.com/moov-io/fedwire20022/gen/ConnectionCheck/admi_004_001_02"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
)

// MessageModel uses base abstractions to eliminate duplicate processing logic
type MessageModel struct {
	EventType  string    `json:"eventType"`
	EventParam string    `json:"eventParam"`
	EventTime  time.Time `json:"eventTime"`
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, ADMI_004_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, ADMI_004_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.01",
			Version:   ADMI_004_001_01,
			Factory: func() models.ISODocument {
				return &admi_004_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[ADMI_004_001_01], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02",
			Version:   ADMI_004_001_02,
			Factory: func() models.ISODocument {
				return &admi_004_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[ADMI_004_001_02], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, ADMI_004_001_VERSION](
		versionedFactory.BuildNameSpaceModelMap(),
		versionedFactory.GetVersionMap(),
		VersionPathMap,
		RequiredFields,
	)
}

var RequiredFields = []string{
	"EventType", "EventParam", "EventTime",
}

// MessageWith uses base abstractions to replace 15+ lines with a single call
func MessageWith(data []byte) (MessageModel, error) {
	return processor.ProcessMessage(data)
}

// DocumentWith uses base abstractions to replace 25+ lines with a single call
func DocumentWith(model MessageModel, version ADMI_004_001_VERSION) (models.ISODocument, error) {
	return processor.CreateDocument(model, version)
}

// CheckRequiredFields uses base abstractions to replace 30+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}
