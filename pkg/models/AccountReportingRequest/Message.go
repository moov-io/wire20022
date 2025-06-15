package AccountReportingRequest

import (
	"encoding/xml"

	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_02"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_03"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_04"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_05"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_06"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_07"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
)

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// AccountReportingRequest-specific fields
	ReportRequestId    models.CAMTReportType `json:"reportRequestId"`
	RequestedMsgNameId string                `json:"requestedMsgNameId"`
	AccountOtherId     string                `json:"accountOtherId"`
	AccountProperty    models.AccountTypeFRS `json:"accountProperty"`
	AccountOwnerAgent  models.Agent          `json:"accountOwnerAgent"`
	FromToSequence     models.SequenceRange  `json:"fromToSequence"`
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "ReportRequestId", "RequestedMsgNameId", "AccountOwnerAgent",
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, CAMT_060_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, CAMT_060_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.02",
			Version:   CAMT_060_001_02,
			Factory: func() models.ISODocument {
				return &camt_060_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_02], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.03",
			Version:   CAMT_060_001_03,
			Factory: func() models.ISODocument {
				return &camt_060_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.04",
			Version:   CAMT_060_001_04,
			Factory: func() models.ISODocument {
				return &camt_060_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05",
			Version:   CAMT_060_001_05,
			Factory: func() models.ISODocument {
				return &camt_060_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.06",
			Version:   CAMT_060_001_06,
			Factory: func() models.ISODocument {
				return &camt_060_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_06], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.060.001.07",
			Version:   CAMT_060_001_07,
			Factory: func() models.ISODocument {
				return &camt_060_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_07], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, CAMT_060_001_VERSION](
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
func DocumentWith(model MessageModel, version CAMT_060_001_VERSION) (models.ISODocument, error) {
	return processor.CreateDocument(model, version)
}

// CheckRequiredFields uses base abstractions to replace 30+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}
