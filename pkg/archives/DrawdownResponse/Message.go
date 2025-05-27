package DrawdownResponse

import (
	"encoding/xml"
	"fmt"
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
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageModel struct {
	MessageId                       string
	CreateDatetime                  time.Time
	InitiatingParty                 Archive.PartyIdentify
	DebtorAgent                     Archive.Agent
	CreditorAgent                   Archive.Agent
	OriginalMessageId               string
	OriginalMessageNameId           string
	OriginalCreationDateTime        time.Time
	OriginalPaymentInfoId           string
	TransactionInformationAndStatus TransactionInfoAndStatus
}

var NameSpaceModelMap = map[string]Archive.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.01": func() Archive.ISODocument {
		return &pain_014_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_01], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.02": func() Archive.ISODocument {
		return &pain_014_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.03": func() Archive.ISODocument {
		return &pain_014_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.04": func() Archive.ISODocument {
		return &pain_014_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.05": func() Archive.ISODocument {
		return &pain_014_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.06": func() Archive.ISODocument {
		return &pain_014_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.07": func() Archive.ISODocument {
		return &pain_014_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_07], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.08": func() Archive.ISODocument {
		return &pain_014_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_08], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.09": func() Archive.ISODocument {
		return &pain_014_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_09], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pain.014.001.10": func() Archive.ISODocument {
		return &pain_014_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PAIN_014_001_10], Local: "Document"}}
	},
}
var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "InitiatingParty", "DebtorAgent", "CreditorAgent", "OriginalMessageId",
	"OriginalMessageNameId", "OriginalCreationDateTime", "OriginalPaymentInfoId", "TransactionInformationAndStatus",
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
func DocumentWith(model MessageModel, version PAIN_014_001_VESION) (Archive.ISODocument, error) {
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
		"MessageId":                       model.MessageId,
		"CreatedDateTime":                 model.CreateDatetime,
		"InitiatingParty":                 model.InitiatingParty,
		"DebtorAgent":                     model.DebtorAgent,
		"CreditorAgent":                   model.CreditorAgent,
		"OriginalMessageId":               model.OriginalMessageId,
		"OriginalMessageNameId":           model.OriginalMessageNameId,
		"OriginalCreationDateTime":        model.OriginalCreationDateTime,
		"OriginalPaymentInfoId":           model.OriginalPaymentInfoId,
		"TransactionInformationAndStatus": model.TransactionInformationAndStatus,
	}
	for _, field := range RequiredFields {
		if value, exists := fieldMap[field]; exists && Archive.IsEmpty(value) {
			return fmt.Errorf("missing required field: %s", field)
		}
	}

	return nil
}
