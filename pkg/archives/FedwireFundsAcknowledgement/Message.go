package FedwireFundsAcknowledgement

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/moov-io/fedwire20022/gen/FedwireFundsAcknowledgement/admi_007_001_01"
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageModel struct {
	MessageId         string
	CreatedDateTime   time.Time
	RelationReference string
	ReferenceName     string
	RequestHandling   Archive.RelatedStatusCode
}

var NameSpaceModelMap = map[string]Archive.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:admi.007.001.01": func() Archive.ISODocument {
		return &admi_007_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[ADMI_007_001_01], Local: "Document"}}
	},
}
var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "RelationReference", "ReferenceName", "RequestHandling",
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, xmlns, err := Archive.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, fmt.Errorf("failed to create document: %w", err)
	}
	version := NameSpaceVersonMap[xmlns]

	dataModel := MessageModel{}
	pathMap := VersionPathMap[version]
	rePathMap := Archive.RemakeMapping(doc, pathMap, true)
	for sourcePath, targetPath := range rePathMap {
		Archive.CopyDocumentValueToMessage(doc, sourcePath, &dataModel, targetPath)
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel, version ADMI_007_001_VESION) (Archive.ISODocument, error) {
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
	rePathMap := Archive.RemakeMapping(model, pathMap, false)
	for targetPath, sourcePath := range rePathMap {
		if err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath); err != nil {
			return document, err
		}
	}
	return document, nil
}
func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"MessageId":         model.MessageId,
		"CreationDateTime":  model.CreatedDateTime,
		"RelationReference": model.RelationReference,
		"ReferenceName":     model.ReferenceName,
		"RequestHandling":   model.RequestHandling,
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
