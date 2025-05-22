package ArchiveAccountReportingRequest

import (
	"encoding/xml"
	"fmt"
	"time"

	Archive "github.com/moov-io/wire20022/pkg/archives"

	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_02"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_03"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_04"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_05"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_06"
	"github.com/moov-io/fedwire20022/gen/AccountReportingRequest/camt_060_001_07"
)

type MessageModel struct {
	MessageId          string
	CreatedDateTime    time.Time
	ReportRequestId    Archive.CAMTReportType
	RequestedMsgNameId string
	AccountOtherId     string
	AccountProperty    Archive.AccountTypeFRS
	AccountOwnerAgent  Archive.Agent
	FromToSequence     Archive.SequenceRange
}

var NameSpaceModelMap = map[string]Archive.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.02": func() Archive.ISODocument {
		return &camt_060_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03": func() Archive.ISODocument {
		return &camt_060_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04": func() Archive.ISODocument {
		return &camt_060_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.05": func() Archive.ISODocument {
		return &camt_060_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.06": func() Archive.ISODocument {
		return &camt_060_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.07": func() Archive.ISODocument {
		return &camt_060_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_060_001_07], Local: "Document"}}
	},
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "ReportRequestId", "RequestedMsgNameId", "AccountOwnerAgent",
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
func DocumentWith(model MessageModel, version CAMT_060_001_VESION) (Archive.ISODocument, error) {
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
		case "MessageId":
			if model.MessageId == "" {
				return fmt.Errorf("missing required field: %s", field)
			}
		case "CreatedDateTime":
			if model.CreatedDateTime.IsZero() {
				return fmt.Errorf("missing required field: %s", field)
			}
		case "ReportRequestId":
			if model.ReportRequestId == "" {
				return fmt.Errorf("missing required field: %s", field)
			}
		case "RequestedMsgNameId":
			if model.RequestedMsgNameId == "" {
				return fmt.Errorf("missing required field: %s", field)
			}
		case "AccountOwnerAgent":
			if model.AccountOwnerAgent.PaymentSysCode == "" || model.AccountOwnerAgent.PaymentSysMemberId == "" {
				return fmt.Errorf("missing required field: %s", field)
			}
		default:
			return nil
		}
	}
	return nil
}
