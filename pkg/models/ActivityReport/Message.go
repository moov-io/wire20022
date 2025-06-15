package ActivityReport

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_01"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_02"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_03"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_04"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_05"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_06"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_07"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_08"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_09"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_10"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_11"
	"github.com/moov-io/fedwire20022/gen/ActivityReport/camt_052_001_12"
	"github.com/wadearnold/wire20022/pkg/base"
	"github.com/wadearnold/wire20022/pkg/errors"
	"github.com/wadearnold/wire20022/pkg/models"
)

// MessageModel uses base abstractions for common fields but keeps custom logic for complex arrays
type MessageModel struct {
	// Use base abstraction for common header fields
	base.MessageHeader `json:",inline"`

	// Override MessageId with proper type for ActivityReport
	MessageId models.CAMTReportType `json:"messageId"`

	// ActivityReport-specific fields
	Pagenation                         models.MessagePagenation              `json:"pagenation"`
	ReportId                           models.ReportType                     `json:"reportId"`
	ReportCreateDateTime               time.Time                             `json:"reportCreateDateTime"`
	AccountOtherId                     string                                `json:"accountOtherId"`
	TotalEntries                       string                                `json:"totalEntries"`
	TotalCreditEntries                 models.NumberAndSumOfTransactions     `json:"totalCreditEntries"`
	TotalDebitEntries                  models.NumberAndSumOfTransactions     `json:"totalDebitEntries"`
	TotalEntriesPerBankTransactionCode []models.TotalsPerBankTransactionCode `json:"totalEntriesPerBankTransactionCode"`
	EntryDetails                       []models.Entry                        `json:"entryDetails"`
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime",
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, xmlns, err := models.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, err
	}
	version := NameSpaceVersionMap[xmlns]

	dataModel := MessageModel{}
	pathMap := VersionPathMap[version]
	rePathMap := models.RemakeMapping(doc, pathMap, true)
	for sourcePath, targetPath := range rePathMap {
		models.CopyDocumentValueToMessage(doc, sourcePath, &dataModel, targetPath)
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
		return nil, errors.NewInvalidFieldError("version", "unsupported document version: "+string(version))
	}

	// Create the document using the factory
	document := factory()
	// Remap paths and copy values from the model to the document
	rePathMap := models.RemakeMapping(model, pathMap, false)
	for targetPath, sourcePath := range rePathMap {
		if err := models.CopyMessageValueToDocument(model, sourcePath, document, targetPath); err != nil {
			return document, err
		}
	}
	return document, nil
}
func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"MessageId":            model.MessageId,
		"CreatedDateTime":      model.CreatedDateTime,
		"Pagenation":           model.Pagenation.PageNumber,
		"ReportId":             model.ReportId,
		"ReportCreateDateTime": model.ReportCreateDateTime,
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
