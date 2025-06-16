package ReturnRequestResponse

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_03"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_04"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_05"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_06"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_07"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_08"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_09"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_10"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_11"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_12"
	"github.com/wadearnold/wire20022/pkg/base"
	"github.com/wadearnold/wire20022/pkg/models"
)

// MessageModel uses base abstractions to eliminate duplicate field definitions
// (Pattern 3 - Direct Migration with unique assignment-based structure)
type MessageModel struct {
	// ReturnRequestResponse-specific fields (does not use MessageHeader due to unique structure)
	AssignmentId                 string        `json:"assignmentId"`
	Assigner                     models.Agent  `json:"assigner"`
	Assignee                     models.Agent  `json:"assignee"`
	AssignmentCreateTime         time.Time     `json:"assignmentCreateTime"`
	ResolvedCaseId               string        `json:"resolvedCaseId"`
	Creator                      models.Agent  `json:"creator"`
	Status                       models.Status `json:"status"`
	OriginalMessageId            string        `json:"originalMessageId"`
	OriginalMessageNameId        string        `json:"originalMessageNameId"`
	OriginalMessageCreateTime    time.Time     `json:"originalMessageCreateTime"`
	OriginalInstructionId        string        `json:"originalInstructionId"`
	OriginalEndToEndId           string        `json:"originalEndToEndId"`
	OriginalUETR                 string        `json:"originalUETR"`
	CancellationStatusReasonInfo models.Reason `json:"cancellationStatusReasonInfo"`
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, CAMT_029_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, CAMT_029_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.03",
			Version:   CAMT_029_001_03,
			Factory: func() models.ISODocument {
				return &camt_029_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.04",
			Version:   CAMT_029_001_04,
			Factory: func() models.ISODocument {
				return &camt_029_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.05",
			Version:   CAMT_029_001_05,
			Factory: func() models.ISODocument {
				return &camt_029_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.06",
			Version:   CAMT_029_001_06,
			Factory: func() models.ISODocument {
				return &camt_029_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_06], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.07",
			Version:   CAMT_029_001_07,
			Factory: func() models.ISODocument {
				return &camt_029_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_07], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.08",
			Version:   CAMT_029_001_08,
			Factory: func() models.ISODocument {
				return &camt_029_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_08], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09",
			Version:   CAMT_029_001_09,
			Factory: func() models.ISODocument {
				return &camt_029_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_09], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.10",
			Version:   CAMT_029_001_10,
			Factory: func() models.ISODocument {
				return &camt_029_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_10], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.11",
			Version:   CAMT_029_001_11,
			Factory: func() models.ISODocument {
				return &camt_029_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_11], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:camt.029.001.12",
			Version:   CAMT_029_001_12,
			Factory: func() models.ISODocument {
				return &camt_029_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_12], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, CAMT_029_001_VERSION](
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

// DocumentWith uses base abstractions to replace 15+ lines with a single call
func DocumentWith(model MessageModel, version CAMT_029_001_VERSION) (models.ISODocument, error) {
	// Validate required fields before creating document
	if err := processor.ValidateRequiredFields(model); err != nil {
		return nil, err
	}
	return processor.CreateDocument(model, version)
}

var RequiredFields = []string{
	"AssignmentId", "Assigner", "Assignee",
	"AssignmentCreateTime", "ResolvedCaseId", "Creator", "OriginalMessageId",
	"OriginalMessageNameId", "OriginalMessageCreateTime",
}

// CheckRequiredFields uses base abstractions to replace 20+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}
