package FedwireFundsSystemResponse

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/FedwireFundsSystemResponse/admi_011_001_01"
	"github.com/moov-io/wire20022/pkg/errors"
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

var NameSpaceModelMap = map[string]models.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:admi.011.001.01": func() models.ISODocument {
		return &admi_011_001_01.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[ADMI_011_001_01], Local: "Document"}}
	},
}

var RequiredFields = []string{
	"MessageId", "EventCode", "EventParam", "EventTime",
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, xmlns, err := models.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, errors.NewParseError("document creation", "XML data", err)
	}
	version := NameSpaceVersonMap[xmlns]

	dataModel := MessageModel{}
	pathMap := VersionPathMap[version]
	rePathMap := models.RemakeMapping(doc, pathMap, true)
	for sourcePath, targetPath := range rePathMap {
		models.CopyDocumentValueToMessage(doc, sourcePath, &dataModel, targetPath)
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel, version ADMI_011_001_VERSION) (models.ISODocument, error) {
	// Check required fields in the model
	if err := CheckRequiredFields(model); err != nil {
		return nil, err
	}

	// Retrieve the path map and document factory for the given version
	pathMap, pathExists := VersionPathMap[version]
	factory, factoryExists := NameSpaceModelMap[VersionNameSpaceMap[version]]
	if !pathExists || !factoryExists {
		return nil, errors.NewInvalidFieldError("version", "unsupported document version")
	}

	// Create the document using the factory
	document := factory()
	// Remap paths and copy values from the model to the document
	rePathMap := models.RemakeMapping(model, pathMap, false)
	for targetPath, sourcePath := range rePathMap {
		if err := models.CopyMessageValueToDocument(model, sourcePath, document, targetPath); err != nil {
			return document, errors.NewFieldError(targetPath, "copy", err)
		}
	}
	return document, nil
}
func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"MessageId":  model.MessageId,
		"EventCode":  model.EventCode,
		"EventParam": model.EventParam,
		"EventTime":  model.EventTime,
	}

	for _, field := range RequiredFields {
		if value, ok := fieldMap[field]; ok {
			if models.IsEmpty(value) {
				return errors.NewRequiredFieldError(field)
			}
		}
	}

	return nil
}
