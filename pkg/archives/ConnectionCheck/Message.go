package ConnectionCheck

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/moov-io/fedwire20022/gen/ConnectionCheck/admi_004_001_01"
	"github.com/moov-io/fedwire20022/gen/ConnectionCheck/admi_004_001_02"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageModel struct {
	EventType  string
	EventParam string
	EventTime  time.Time
}

var NameSpaceModelMap = map[string]Archive.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:admi.004.001.01": func() Archive.ISODocument {
		return &admi_004_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[ADMI_004_001_01], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02": func() Archive.ISODocument {
		return &admi_004_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[ADMI_004_001_02], Local: "Document"}}
	},
}

var RequiredFields = []string{
	"EventType", "EventParam", "EventTime",
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
func DocumentWith(model MessageModel, version ADMI_004_001_VESION) (Archive.ISODocument, error) {
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
	for _, field := range RequiredFields {
		switch field {
		case "EventType":
			if model.EventType == "" {
				return fmt.Errorf("missing required field: %s", field)
			}
		case "EventParam":
			if model.EventParam == "" {
				return fmt.Errorf("missing required field: %s", field)
			}
		case "EventTime":
			if model.EventTime.IsZero() {
				return fmt.Errorf("missing required field: %s", field)
			}
		default:
			return nil
		}
	}
	return nil
}
