package FedwireFundsPaymentStatus

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_03"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_04"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_05"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_06"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_07"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_08"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_09"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_10"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_11"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_12"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_13"
	"github.com/moov-io/fedwire20022/gen/FedwireFundsPaymentStatus/pacs_002_001_14"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/models"
)

type MessageModel struct {
	MessageId                        string
	CreatedDateTime                  time.Time
	OriginalMessageId                string
	OriginalMessageNameId            string
	OriginalMessageCreateTime        time.Time
	OriginalUETR                     string
	TransactionStatus                models.TransactionStatusCode
	AcceptanceDateTime               time.Time
	EffectiveInterbankSettlementDate fedwire.ISODate
	StatusReasonInformation          string
	ReasonAdditionalInfo             string
	InstructingAgent                 models.Agent
	InstructedAgent                  models.Agent
}

var NameSpaceModelMap = map[string]models.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.03": func() models.ISODocument {
		return &pacs_002_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.04": func() models.ISODocument {
		return &pacs_002_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.05": func() models.ISODocument {
		return &pacs_002_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.06": func() models.ISODocument {
		return &pacs_002_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.07": func() models.ISODocument {
		return &pacs_002_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_07], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08": func() models.ISODocument {
		return &pacs_002_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_08], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.09": func() models.ISODocument {
		return &pacs_002_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_09], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10": func() models.ISODocument {
		return &pacs_002_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_10], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.11": func() models.ISODocument {
		return &pacs_002_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_11], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12": func() models.ISODocument {
		return &pacs_002_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_12], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.13": func() models.ISODocument {
		return &pacs_002_001_13.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_13], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.14": func() models.ISODocument {
		return &pacs_002_001_14.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_002_001_14], Local: "Document"}}
	},
}
var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "TransactionStatus", "InstructingAgent", "InstructedAgent",
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
func DocumentWith(model MessageModel, version PACS_002_001_VESION) (models.ISODocument, error) {
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
		"MessageId":         model.MessageId,
		"CreationDateTime":  model.CreatedDateTime,
		"TransactionStatus": model.TransactionStatus,
		"InstructingAgent":  model.InstructedAgent,
		"InstructedAgent":   model.InstructedAgent,
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
