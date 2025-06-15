package ReturnRequestResponse

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_03"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_04"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_05"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_06"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_07"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_08"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_09"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_10"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_11"
	"github.com/moov-io/fedwire20022/gen/ReturnRequestResponse/camt_029_001_12"
	"github.com/moov-io/wire20022/pkg/errors"
	"github.com/moov-io/wire20022/pkg/models"
)

// MessageModel represents an investigation resolution message (Pattern 3 - Direct Migration)
// Does not use base.MessageHeader as it has a unique assignment-based structure
type MessageModel struct {
	AssignmentId                 string        `json:"assignmentId"`
	Assigner                     models.Agent  `json:"assigner"`
	Assignee                     models.Agent  `json:"assignee"`
	AssignmentCreateTime         time.Time     `json:"assignmentCreateTime"`
	ResolvedCaseId               string        `json:"resolvedCaseId"`
	Creator                      models.Agent  `json:"creator"`
	Status                       models.Status `json:"status"`
	OriginalMessageId            string        `json:"originalMessageId"`
	OriginalMessageNameId        string        `json:"originalMessageNameId"`
	OriginalMessageCreateTime    time.Time     `json:"originalMessageCreateTime"`
	OriginalInstructionId        string        `json:"originalInstructionId"`
	OriginalEndToEndId           string        `json:"originalEndToEndId"`
	OriginalUETR                 string        `json:"originalUETR"`
	CancellationStatusReasonInfo models.Reason `json:"cancellationStatusReasonInfo"`
}

var NameSpaceModelMap = map[string]models.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.03": func() models.ISODocument {
		return &camt_029_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.04": func() models.ISODocument {
		return &camt_029_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.05": func() models.ISODocument {
		return &camt_029_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.06": func() models.ISODocument {
		return &camt_029_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.07": func() models.ISODocument {
		return &camt_029_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_07], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.08": func() models.ISODocument {
		return &camt_029_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_08], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09": func() models.ISODocument {
		return &camt_029_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_09], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.10": func() models.ISODocument {
		return &camt_029_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_10], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.11": func() models.ISODocument {
		return &camt_029_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_11], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:camt.029.001.12": func() models.ISODocument {
		return &camt_029_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[CAMT_029_001_12], Local: "Document"}}
	},
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
		models.CopyDocumentValueToMessage(doc, sourcePath, &dataModel, targetPath)
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel, version CAMT_029_001_VERSION) (models.ISODocument, error) {
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

	document := factory()
	for targetPath, sourcePath := range pathMap {
		if err := models.CopyMessageValueToDocument(model, sourcePath, document, targetPath); err != nil {
			return document, errors.NewFieldError(targetPath, "copy", err)
		}
	}
	return document, nil
}

var RequiredFields = []string{
	"AssignmentId", "Assigner", "Assignee",
	"AssignmentCreateTime", "ResolvedCaseId", "Creator", "OriginalMessageId",
	"OriginalMessageNameId", "OriginalMessageCreateTime",
}

func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"AssignmentId":              model.AssignmentId,
		"Assigner":                  model.Assigner,
		"Assignee":                  model.Assignee,
		"AssignmentCreateTime":      model.AssignmentCreateTime,
		"ResolvedCaseId":            model.ResolvedCaseId,
		"Creator":                   model.Creator,
		"OriginalMessageId":         model.OriginalMessageId,
		"OriginalMessageNameId":     model.OriginalMessageNameId,
		"OriginalMessageCreateTime": model.OriginalMessageCreateTime,
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
