package DrawdownRequest

import (
	"encoding/xml"
	"fmt"
	"time"

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
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageModel struct {
	MessageId              string
	CreateDatetime         time.Time
	NumberofTransaction    string
	InitiatingParty        Archive.PartyIdentify
	PaymentInfoId          string
	PaymentMethod          Archive.PaymentMethod
	RequestedExecutDate    fedwire.ISODate
	Debtor                 Archive.PartyIdentify
	DebtorAccountOtherId   string
	DebtorAgent            Archive.Agent
	CreditTransTransaction CreditTransferTransaction
}

var NameSpaceModelMap = map[string]Archive.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.01": func() Archive.ISODocument {
		return &pain_013_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_01], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.02": func() Archive.ISODocument {
		return &pain_013_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.03": func() Archive.ISODocument {
		return &pain_013_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.04": func() Archive.ISODocument {
		return &pain_013_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.05": func() Archive.ISODocument {
		return &pain_013_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.06": func() Archive.ISODocument {
		return &pain_013_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.07": func() Archive.ISODocument {
		return &pain_013_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_07], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.08": func() Archive.ISODocument {
		return &pain_013_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_08], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.09": func() Archive.ISODocument {
		return &pain_013_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_09], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.013.001.10": func() Archive.ISODocument {
		return &pain_013_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_013_001_10], Local: "Document"}}
	},
}
var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "NumberofTransaction", "InitiatingParty", "PaymentInfoId", "PaymentMethod",
	"RequestedExecutDate", "Debtor", "DebtorAgent", "CreditTransTransaction",
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, xmlns, err := Archive.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, fmt.Errorf("failed to create document: %w", err)
	}
	version := NameSpaceVersonMap[xmlns]

	dataModel := MessageModel{}
	pathMap := VersionPathMap[version]
	for sourcePath, targetPath := range pathMap {
		Archive.CopyDocumentValueToMessage(doc, sourcePath, &dataModel, targetPath)
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel, version PAIN_013_001_VESION) (Archive.ISODocument, error) {
	// Check required fields in the model
	if err := CheckRequiredFields(model); err != nil {
		return nil, err
	}

	// Retrieve the path map and document factory for the given version
	pathMap, pathExists := VersionPathMap[version]
	factory, factoryExists := NameSpaceModelMap[VersionNameSpaceMap[version]]
	if !pathExists || !factoryExists {
		return nil, fmt.Errorf("unsupported document version: %v", version)
	}

	// Create the document using the factory
	document := factory()
	// Remap paths and copy values from the model to the document
	for targetPath, sourcePath := range pathMap {
		if err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath); err != nil {
			return document, err
		}
	}
	return document, nil
}
func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"MessageId":              model.MessageId,
		"CreateDatetime":         model.CreateDatetime,
		"NumberofTransaction":    model.NumberofTransaction,
		"InitiatingParty":        model.InitiatingParty,
		"PaymentInfoId":          model.PaymentInfoId,
		"PaymentMethod":          model.PaymentMethod,
		"RequestedExecutDate":    model.RequestedExecutDate,
		"Debtor":                 model.Debtor,
		"DebtorAgent":            model.DebtorAgent,
		"CreditTransTransaction": model.CreditTransTransaction,
	}

	for _, field := range RequiredFields {
		if value, ok := fieldMap[field]; ok {
			if Archive.IsEmpty(value) {
				return fmt.Errorf("missing required field: %s", field)
			}
		}
	}

	return nil
}
