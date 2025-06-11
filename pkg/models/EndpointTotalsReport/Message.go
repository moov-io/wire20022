package EndpointTotalsReport

import (
	"encoding/xml"
	"fmt"
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
	"github.com/moov-io/wire20022/pkg/models"
)

type MessageModel struct {
	MessageId                          models.CAMTReportType
	CreatedDateTime                    time.Time
	Pagenation                         models.MessagePagenation
	ReportId                           models.ReportType
	ReportCreateDateTime               time.Time
	AccountOtherId                     string
	TotalCreditEntries                 models.NumberAndSumOfTransactions
	TotalDebitEntries                  models.NumberAndSumOfTransactions
	TotalEntriesPerBankTransactionCode []models.TotalsPerBankTransactionCode
	AdditionalReportInfo               string
}

var NameSpaceModelMap = map[string]models.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.02": func() models.ISODocument {
		return &camt_052_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.03": func() models.ISODocument {
		return &camt_052_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.04": func() models.ISODocument {
		return &camt_052_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.05": func() models.ISODocument {
		return &camt_052_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.06": func() models.ISODocument {
		return &camt_052_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.07": func() models.ISODocument {
		return &camt_052_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_07], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08": func() models.ISODocument {
		return &camt_052_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_08], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.09": func() models.ISODocument {
		return &camt_052_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_09], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.10": func() models.ISODocument {
		return &camt_052_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_10], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.11": func() models.ISODocument {
		return &camt_052_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_11], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.12": func() models.ISODocument {
		return &camt_052_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_052_001_12], Local: "Document"}}
	},
}
var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime",
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, xmlns, err := models.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, fmt.Errorf("failed to create document: %w", err)
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
func DocumentWith(model MessageModel, version CAMT_052_001_VESION) (models.ISODocument, error) {
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
		"MessageId":        model.MessageId,
		"CreationDateTime": model.CreatedDateTime,
		"Pagenation":       model.Pagenation,
		"ReportId":         model.ReportId,
	}

	for _, field := range RequiredFields {
		if value, ok := fieldMap[field]; ok {
			if models.IsEmpty(value) {
				return fmt.Errorf("missing required field: %s", field)
			}
		}
	}

	return nil
}
