package PaymentStatusRequest

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_01"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_02"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_03"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_04"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_05"
	"github.com/moov-io/fedwire20022/gen/PaymentStatusRequest/pacs_028_001_06"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageModel struct {
	MessageId                string
	CreatedDateTime          time.Time
	OriginalMessageId        string
	OriginalMessageNameId    string
	OriginalCreationDateTime time.Time
	OriginalInstructionId    string
	OriginalEndToEndId       string
	OriginalUETR             string
	InstructingAgent         Archive.Agent
	InstructedAgent          Archive.Agent
}

var NameSpaceModelMap = map[string]Archive.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.01": func() Archive.ISODocument {
		return &pacs_028_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_01], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.02": func() Archive.ISODocument {
		return &pacs_028_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03": func() Archive.ISODocument {
		return &pacs_028_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.04": func() Archive.ISODocument {
		return &pacs_028_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.05": func() Archive.ISODocument {
		return &pacs_028_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.028.001.06": func() Archive.ISODocument {
		return &pacs_028_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_028_001_06], Local: "Document"}}
	},
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
func DocumentWith(model MessageModel, version PACS_028_001_VESION) (Archive.ISODocument, error) {
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

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "OriginalMessageId",
	"OriginalMessageNameId", "OriginalCreationDateTime",
	"InstructingAgent", "InstructedAgent",
}

func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"MessageId":                model.MessageId,
		"CreatedDateTime":          model.CreatedDateTime,
		"OriginalMessageId":        model.OriginalMessageId,
		"OriginalMessageNameId":    model.OriginalMessageNameId,
		"OriginalCreationDateTime": model.OriginalCreationDateTime,
		"InstructingAgent":         model.InstructingAgent,
		"InstructedAgent":          model.InstructedAgent,
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
