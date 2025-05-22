package ActivityReport

import (
	"encoding/xml"
	"fmt"
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
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageModel struct {
	MessageId                          Archive.CAMTReportType
	CreatedDateTime                    time.Time
	Pagenation                         Archive.MessagePagenation
	ReportId                           Archive.ReportType
	ReportCreateDateTime               time.Time
	AccountOtherId                     string
	TotalEntries                       string
	TotalCreditEntries                 Archive.NumberAndSumOfTransactions
	TotalDebitEntries                  Archive.NumberAndSumOfTransactions
	TotalEntriesPerBankTransactionCode []Archive.TotalsPerBankTransactionCode
	EntryDetails                       []Archive.Entry
}

var NameSpaceModelMap = map[string]Archive.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.01": func() Archive.ISODocument { return &camt_052_001_01.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.02": func() Archive.ISODocument { return &camt_052_001_02.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.03": func() Archive.ISODocument { return &camt_052_001_03.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.04": func() Archive.ISODocument { return &camt_052_001_04.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.05": func() Archive.ISODocument { return &camt_052_001_05.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.06": func() Archive.ISODocument { return &camt_052_001_06.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.07": func() Archive.ISODocument { return &camt_052_001_07.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.08": func() Archive.ISODocument { return &camt_052_001_08.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.09": func() Archive.ISODocument { return &camt_052_001_09.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.10": func() Archive.ISODocument { return &camt_052_001_10.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.11": func() Archive.ISODocument { return &camt_052_001_11.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.052.001.12": func() Archive.ISODocument { return &camt_052_001_12.Document{} },
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "Pagenation", "ReportId", "ReportCreateDateTime",
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, err := Archive.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, fmt.Errorf("failed to create document: %w", err)
	}
	dataModel := MessageModel{}
	if Doc01, ok := doc.(*camt_052_001_01.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_01]
		rePathMap := Archive.RemakeMapping(Doc01, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc01, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc02, ok := doc.(*camt_052_001_02.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_02]
		rePathMap := Archive.RemakeMapping(Doc02, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc02, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc03, ok := doc.(*camt_052_001_03.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_03]
		rePathMap := Archive.RemakeMapping(Doc03, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc03, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc04, ok := doc.(*camt_052_001_04.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_04]
		rePathMap := Archive.RemakeMapping(Doc04, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc04, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc05, ok := doc.(*camt_052_001_05.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_05]
		rePathMap := Archive.RemakeMapping(Doc05, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc05, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc06, ok := doc.(*camt_052_001_06.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_06]
		rePathMap := Archive.RemakeMapping(Doc06, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc06, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc07, ok := doc.(*camt_052_001_07.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_07]
		rePathMap := Archive.RemakeMapping(Doc07, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc07, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc08, ok := doc.(*camt_052_001_08.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_08]
		rePathMap := Archive.RemakeMapping(Doc08, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc08, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc09, ok := doc.(*camt_052_001_09.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_09]
		rePathMap := Archive.RemakeMapping(Doc09, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc09, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc10, ok := doc.(*camt_052_001_10.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_10]
		rePathMap := Archive.RemakeMapping(Doc10, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc10, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc11, ok := doc.(*camt_052_001_11.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_11]
		rePathMap := Archive.RemakeMapping(Doc11, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc11, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc12, ok := doc.(*camt_052_001_12.Document); ok {
		pathMap := VersionPathMap[CAMT_052_001_12]
		rePathMap := Archive.RemakeMapping(Doc12, pathMap, true)
		for sourcePath, targetPath := range rePathMap {
			Archive.CopyDocumentValueToMessage(Doc12, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel, verson CAMT_052_001_VESION) (Archive.ISODocument, error) {
	err := CheckRequiredFields(model)
	if err != nil {
		return nil, err
	}
	var document Archive.ISODocument
	if verson == CAMT_052_001_01 {
		pathMap := VersionPathMap[CAMT_052_001_01]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_01.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_02 {
		pathMap := VersionPathMap[CAMT_052_001_02]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_02.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_03 {
		pathMap := VersionPathMap[CAMT_052_001_03]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_03.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_04 {
		pathMap := VersionPathMap[CAMT_052_001_04]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_04.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_05 {
		pathMap := VersionPathMap[CAMT_052_001_05]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_05.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_06 {
		pathMap := VersionPathMap[CAMT_052_001_06]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_06.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_07 {
		pathMap := VersionPathMap[CAMT_052_001_07]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_07.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_08 {
		pathMap := VersionPathMap[CAMT_052_001_08]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_08.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_09 {
		pathMap := VersionPathMap[CAMT_052_001_09]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_09.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_10 {
		pathMap := VersionPathMap[CAMT_052_001_10]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_10.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_11 {
		pathMap := VersionPathMap[CAMT_052_001_11]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_11.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
		}
	} else if verson == CAMT_052_001_12 {
		pathMap := VersionPathMap[CAMT_052_001_12]
		rePathMap := Archive.RemakeMapping(model, pathMap, false)
		document = &camt_052_001_12.Document{
			XMLName: xml.Name{
				Space: VersionNameSpaceMap[verson],
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range rePathMap {
			err := Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
			if err != nil {
				return document, err
			}
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
		case "Pagenation":
			if model.Pagenation.PageNumber == "" {
				return fmt.Errorf("missing required field: %s", field)
			}
		case "ReportId":
			if model.ReportId == "" {
				return fmt.Errorf("missing required field: %s", field)
			}
		case "ReportCreateDateTime":
			if model.ReportCreateDateTime.IsZero() {
				return fmt.Errorf("missing required field: %s", field)
			}
		default:
			return nil
		}
	}
	return nil
}
