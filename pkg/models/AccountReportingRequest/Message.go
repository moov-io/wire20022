package ArchiveAccountReportingRequest

import (
	"encoding/xml"
	"github.com/moov-io/wire20022/pkg/errors"
	"time"

	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_02"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_03"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_04"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_05"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_06"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_07"
	"github.com/moov-io/wire20022/pkg/models"
)

type MessageModel struct {
	MessageId          string
	CreatedDateTime    time.Time
	ReportRequestId    models.CAMTReportType
	RequestedMsgNameId string
	AccountOtherId     string
	AccountProperty    models.AccountTypeFRS
	AccountOwnerAgent  models.Agent
	FromToSequence     models.SequenceRange
}

var NameSpaceModelMap = map[string]models.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.02": func() models.ISODocument {
		return &camt_060_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03": func() models.ISODocument {
		return &camt_060_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04": func() models.ISODocument {
		return &camt_060_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.05": func() models.ISODocument {
		return &camt_060_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.06": func() models.ISODocument {
		return &camt_060_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.07": func() models.ISODocument {
		return &camt_060_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_07], Local: "Document"}}
	},
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "ReportRequestId", "RequestedMsgNameId", "AccountOwnerAgent",
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, xmlns, err := models.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, errors.NewParseError("document creation", "XML data", err)
	}
	version := NameSpaceVersonMap[xmlns]

	dataModel := MessageModel{}
	pathMap := VersionPathMap[version]
	for sourcePath, targetPath := range pathMap {
		models.CopyDocumentValueToMessage(doc, sourcePath, &dataModel, targetPath)
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel, version CAMT_060_001_VERSION) (models.ISODocument, error) {
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
	for targetPath, sourcePath := range pathMap {
		if err := models.CopyMessageValueToDocument(model, sourcePath, document, targetPath); err != nil {
			return document, errors.NewFieldError(targetPath, "copy", err)
		}
	}
	return document, nil
}

func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"MessageId":          model.MessageId,
		"CreatedDateTime":    model.CreatedDateTime,
		"ReportRequestId":    model.ReportRequestId,
		"RequestedMsgNameId": model.RequestedMsgNameId,
		"AccountOwnerAgent":  model.AccountOwnerAgent,
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
