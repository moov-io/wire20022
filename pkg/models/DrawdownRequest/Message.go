package DrawdownRequest

import (
	"encoding/xml"

	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_01"
	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_02"
	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_03"
	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_04"
	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_05"
	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_06"
	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_07"
	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_08"
	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_09"
	"github.com/moov-io/fedwire20022/gen/DrawdownRequest/pain_013_001_10"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
)

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
	// Embed common message fields instead of duplicating them
	base.MessageHeader `json:",inline"`

	// DrawdownRequest-specific fields
	NumberofTransaction    string                    `json:"numberofTransaction"`
	InitiatingParty        models.PartyIdentify      `json:"initiatingParty"`
	PaymentInfoId          string                    `json:"paymentInfoId"`
	PaymentMethod          models.PaymentMethod      `json:"paymentMethod"`
	RequestedExecutDate    fedwire.ISODate           `json:"requestedExecutDate"`
	Debtor                 models.PartyIdentify      `json:"debtor"`
	DebtorAccountOtherId   string                    `json:"debtorAccountOtherId"`
	DebtorAgent            models.Agent              `json:"debtorAgent"`
	CreditTransTransaction CreditTransferTransaction `json:"creditTransTransaction"`
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, PAIN_013_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
	registrations := []base.FactoryRegistration[models.ISODocument, PAIN_013_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.01",
			Version:   PAIN_013_001_01,
			Factory: func() models.ISODocument {
				return &pain_013_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_01], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.02",
			Version:   PAIN_013_001_02,
			Factory: func() models.ISODocument {
				return &pain_013_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_02], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.03",
			Version:   PAIN_013_001_03,
			Factory: func() models.ISODocument {
				return &pain_013_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.04",
			Version:   PAIN_013_001_04,
			Factory: func() models.ISODocument {
				return &pain_013_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.05",
			Version:   PAIN_013_001_05,
			Factory: func() models.ISODocument {
				return &pain_013_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.06",
			Version:   PAIN_013_001_06,
			Factory: func() models.ISODocument {
				return &pain_013_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_06], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.07",
			Version:   PAIN_013_001_07,
			Factory: func() models.ISODocument {
				return &pain_013_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_07], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.08",
			Version:   PAIN_013_001_08,
			Factory: func() models.ISODocument {
				return &pain_013_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_08], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.09",
			Version:   PAIN_013_001_09,
			Factory: func() models.ISODocument {
				return &pain_013_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_09], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pain.013.001.10",
			Version:   PAIN_013_001_10,
			Factory: func() models.ISODocument {
				return &pain_013_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_10], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	processor = base.NewMessageProcessor[MessageModel, PAIN_013_001_VERSION](
		versionedFactory.BuildNameSpaceModelMap(),
		versionedFactory.GetVersionMap(),
		VersionPathMap,
		RequiredFields,
	)
}

var NameSpaceModelMap = map[string]models.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.01": func() models.ISODocument {
		return &pain_013_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_01], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.02": func() models.ISODocument {
		return &pain_013_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.03": func() models.ISODocument {
		return &pain_013_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.04": func() models.ISODocument {
		return &pain_013_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.05": func() models.ISODocument {
		return &pain_013_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.06": func() models.ISODocument {
		return &pain_013_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.07": func() models.ISODocument {
		return &pain_013_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_07], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08": func() models.ISODocument {
		return &pain_013_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_08], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.09": func() models.ISODocument {
		return &pain_013_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_09], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.10": func() models.ISODocument {
		return &pain_013_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_10], Local: "Document"}}
	},
}
var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "NumberofTransaction", "InitiatingParty", "PaymentInfoId", "PaymentMethod",
	"RequestedExecutDate", "Debtor", "DebtorAgent", "CreditTransTransaction",
}

// MessageWith uses base abstractions to replace 15+ lines with a single call
func MessageWith(data []byte) (MessageModel, error) {
	return processor.ProcessMessage(data)
}

// DocumentWith uses base abstractions to replace 25+ lines with a single call
func DocumentWith(model MessageModel, version PAIN_013_001_VERSION) (models.ISODocument, error) {
	return processor.CreateDocument(model, version)
}

// CheckRequiredFields uses base abstractions to replace 30+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}
