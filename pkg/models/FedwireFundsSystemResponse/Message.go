package FedwireFundsSystemResponse

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/FedwireFundsSystemResponse/admi_011_001_01"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
)

// MessageModel represents a non-standard message type (Pattern 3 - Direct Migration)
// Does not use base.MessageHeader as it follows event-based pattern instead of standard message pattern
type MessageModel struct {
	MessageId  string               `json:"messageId"`
	EventCode  models.FundEventType `json:"eventCode"`
	EventParam string               `json:"eventParam"`
	EventTime  time.Time            `json:"eventTime"`
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

// MessageWith uses base abstractions to replace 15+ lines with a single call
func MessageWith(data []byte) (MessageModel, error) {
	return processor.ProcessMessage(data)
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
