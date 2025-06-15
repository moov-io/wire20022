package EndpointGapReport

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_02"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_03"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_04"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_05"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_06"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_07"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_08"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_09"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_10"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_11"
	"github.com/moov-io/fedwire20022/gen/Endpoint/camt_052_001_12"
	"github.com/wadearnold/wire20022/pkg/base"
	"github.com/wadearnold/wire20022/pkg/errors"
	"github.com/wadearnold/wire20022/pkg/models"
)

// MessageModel uses base abstractions with field override for MessageId type
type MessageModel struct {
	// Embed common message fields but override MessageId for specific type
	base.MessageHeader `json:",inline"`
	MessageId          models.CAMTReportType `json:"messageId"` // Override to use CAMTReportType instead of string

	// EndpointGapReport-specific fields
	Pagenation           models.MessagePagenation `json:"pagenation"`
	ReportId             models.GapType           `json:"reportId"`
	ReportCreateDateTime time.Time                `json:"reportCreateDateTime"`
	AccountOtherId       string                   `json:"accountOtherId"`
	AdditionalReportInfo string                   `json:"additionalReportInfo"`
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime",
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, xmlns, err := models.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, errors.NewParseError("document creation", "XML data", err)
	}
	version := NameSpaceVersionMap[xmlns]

	dataModel := MessageModel{}
	pathMap := VersionPathMap[version]
	for sourcePath, targetPath := range pathMap {
		if targetStr, ok := targetPath.(string); ok {
			models.CopyDocumentValueToMessage(doc, sourcePath, &dataModel, targetStr)
		}
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel, version CAMT_052_001_VERSION) (models.ISODocument, error) {
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
	for targetPath, sourcePath := range pathMap {
		if sourceStr, ok := sourcePath.(string); ok {
			if err := models.CopyMessageValueToDocument(model, sourceStr, document, targetPath); err != nil {
				return document, errors.NewFieldError(targetPath, "copy", err)
			}
		}
	}
	return document, nil
}
func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"MessageId":        model.MessageId,
		"CreationDateTime": model.CreatedDateTime,
		"Pagenation":       model.Pagenation,
		"ReportId":         model.ReportId,
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
