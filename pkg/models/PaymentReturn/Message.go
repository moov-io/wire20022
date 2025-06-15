package PaymentReturn

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_02"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_03"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_04"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_05"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_06"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_07"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_08"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_09"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_10"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_11"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_12"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_13"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/wadearnold/wire20022/pkg/base"
	"github.com/wadearnold/wire20022/pkg/models"
)

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common payment fields instead of duplicating them
	base.PaymentCore `json:",inline"`

	// PaymentReturn-specific fields
	OriginalMessageId                 string                        `json:"originalMessageId"`
	OriginalMessageNameId             string                        `json:"originalMessageNameId"`
	OriginalCreationDateTime          time.Time                     `json:"originalCreationDateTime"`
	OriginalInstructionId             string                        `json:"originalInstructionId"`
	OriginalEndToEndId                string                        `json:"originalEndToEndId"`
	OriginalUETR                      string                        `json:"originalUETR"`
	ReturnedInterbankSettlementAmount models.CurrencyAndAmount      `json:"returnedInterbankSettlementAmount"`
	InterbankSettlementDate           fedwire.ISODate               `json:"interbankSettlementDate"`
	ReturnedInstructedAmount          models.CurrencyAndAmount      `json:"returnedInstructedAmount"`
	ChargeBearer                      models.ChargeBearerType       `json:"chargeBearer"`
	RtrChain                          models.ReturnChain            `json:"rtrChain"`
	ReturnReasonInformation           models.Reason                 `json:"returnReasonInformation"`
	OriginalTransactionRef            models.InstrumentPropCodeType `json:"originalTransactionRef"`

	// Use embedded agent pairs
	base.AgentPair `json:",inline"`
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "NumberOfTransactions", "SettlementMethod", "CommonClearingSysCode",
	"OriginalMessageId", "OriginalMessageNameId", "OriginalCreationDateTime",
	"ReturnedInterbankSettlementAmount", "InterbankSettlementDate", "ReturnedInstructedAmount",
	"InstructingAgent", "InstructedAgent", "ReturnReasonInformation", "OriginalTransactionRef",
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, PACS_004_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, PACS_004_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02",
			Version:   PACS_004_001_02,
			Factory: func() models.ISODocument {
				return &pacs_004_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_02], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.03",
			Version:   PACS_004_001_03,
			Factory: func() models.ISODocument {
				return &pacs_004_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.04",
			Version:   PACS_004_001_04,
			Factory: func() models.ISODocument {
				return &pacs_004_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.05",
			Version:   PACS_004_001_05,
			Factory: func() models.ISODocument {
				return &pacs_004_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.06",
			Version:   PACS_004_001_06,
			Factory: func() models.ISODocument {
				return &pacs_004_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_06], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.07",
			Version:   PACS_004_001_07,
			Factory: func() models.ISODocument {
				return &pacs_004_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_07], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.08",
			Version:   PACS_004_001_08,
			Factory: func() models.ISODocument {
				return &pacs_004_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_08], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.09",
			Version:   PACS_004_001_09,
			Factory: func() models.ISODocument {
				return &pacs_004_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_09], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10",
			Version:   PACS_004_001_10,
			Factory: func() models.ISODocument {
				return &pacs_004_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_10], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.11",
			Version:   PACS_004_001_11,
			Factory: func() models.ISODocument {
				return &pacs_004_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_11], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.12",
			Version:   PACS_004_001_12,
			Factory: func() models.ISODocument {
				return &pacs_004_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_12], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.13",
			Version:   PACS_004_001_13,
			Factory: func() models.ISODocument {
				return &pacs_004_001_13.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_13], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, PACS_004_001_VERSION](
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
func DocumentWith(model MessageModel, version PACS_004_001_VERSION) (models.ISODocument, error) {
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
