package ArchiveAccountReportingRequest

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"

	Archive "github.com/moov-io/wire20022/pkg/archives"
	"github.com/moov-io/wire20022/pkg/archives/AccountReportingRequest/camt_060_001_02"
	"github.com/moov-io/wire20022/pkg/archives/AccountReportingRequest/camt_060_001_03"
	"github.com/moov-io/wire20022/pkg/archives/AccountReportingRequest/camt_060_001_04"
	"github.com/moov-io/wire20022/pkg/archives/AccountReportingRequest/camt_060_001_05"
	"github.com/moov-io/wire20022/pkg/archives/AccountReportingRequest/camt_060_001_06"
	"github.com/moov-io/wire20022/pkg/archives/AccountReportingRequest/camt_060_001_07"
)

type MessageModel struct {
	MessageId          string
	CreatedDateTime    time.Time
	ReportRequestId    CAMTReportType
	RequestedMsgNameId string
	AccountOtherId     string
	AccountProperty    AccountTypeFRS
	AccountOwnerAgent  Agent
	FromToSequence     SequenceRange
}

var NameSpaceModelMap = map[string]Archive.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.02": func() Archive.IOSDocument { return &camt_060_001_02.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03": func() Archive.IOSDocument { return &camt_060_001_03.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.04": func() Archive.IOSDocument { return &camt_060_001_04.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.05": func() Archive.IOSDocument { return &camt_060_001_05.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.06": func() Archive.IOSDocument { return &camt_060_001_06.Document{} },
	"urn:iso:std:iso:20022:tech:xsd:camt.060.001.07": func() Archive.IOSDocument { return &camt_060_001_07.Document{} },
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, err := Archive.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, fmt.Errorf("failed to create document: %w", err)
	}
	dataModel := MessageModel{}
	if Doc02, ok := doc.(*camt_060_001_02.Document); ok {
		pathMap := camt_060_001_02.PathMap()
		for sourcePath, targetPath := range pathMap {
			Archive.CopyDocumentValueToMessage(Doc02, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc03, ok := doc.(*camt_060_001_03.Document); ok {
		pathMap := camt_060_001_03.PathMap()
		for sourcePath, targetPath := range pathMap {
			Archive.CopyDocumentValueToMessage(Doc03, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc04, ok := doc.(*camt_060_001_04.Document); ok {
		pathMap := camt_060_001_04.PathMap()
		for sourcePath, targetPath := range pathMap {
			Archive.CopyDocumentValueToMessage(Doc04, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc05, ok := doc.(*camt_060_001_05.Document); ok {
		pathMap := camt_060_001_05.PathMap()
		for sourcePath, targetPath := range pathMap {
			Archive.CopyDocumentValueToMessage(Doc05, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc06, ok := doc.(*camt_060_001_06.Document); ok {
		pathMap := camt_060_001_06.PathMap()
		for sourcePath, targetPath := range pathMap {
			Archive.CopyDocumentValueToMessage(Doc06, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	} else if Doc07, ok := doc.(*camt_060_001_07.Document); ok {
		pathMap := camt_060_001_07.PathMap()
		for sourcePath, targetPath := range pathMap {
			Archive.CopyDocumentValueToMessage(Doc07, sourcePath, &dataModel, targetPath)
		}
		return dataModel, nil
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel, verson string) (Archive.IOSDocument, error) {
	var document Archive.IOSDocument
	var XMLINS = "urn:iso:std:iso:20022:tech:xsd:" + verson
	if verson == "camt.060.001.02" {
		pathMap := camt_060_001_02.PathMap()
		document = &camt_060_001_02.Document{
			XMLName: xml.Name{
				Space: XMLINS,
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range pathMap {
			Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
		}
	} else if verson == "camt.060.001.03" {
		pathMap := camt_060_001_03.PathMap()
		document = &camt_060_001_03.Document{
			XMLName: xml.Name{
				Space: XMLINS,
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range pathMap {
			Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
		}
	} else if verson == "camt.060.001.04" {
		pathMap := camt_060_001_04.PathMap()
		document = &camt_060_001_04.Document{
			XMLName: xml.Name{
				Space: XMLINS,
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range pathMap {
			Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
		}
	} else if verson == "camt.060.001.05" {
		pathMap := camt_060_001_05.PathMap()
		document = &camt_060_001_05.Document{
			XMLName: xml.Name{
				Space: XMLINS,
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range pathMap {
			Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
		}
	} else if verson == "camt.060.001.06" {
		pathMap := camt_060_001_06.PathMap()
		document = &camt_060_001_06.Document{
			XMLName: xml.Name{
				Space: XMLINS,
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range pathMap {
			Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
		}
	} else if verson == "camt.060.001.07" {
		pathMap := camt_060_001_07.PathMap()
		document = &camt_060_001_07.Document{
			XMLName: xml.Name{
				Space: XMLINS,
				Local: "Document",
			},
		}
		for targetPath, sourcePath := range pathMap {
			Archive.CopyMessageValueToDocument(model, sourcePath, document, targetPath)
		}
	}
	return document, nil
}

func ReadXMLFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", filename, err)
	}
	return data, nil
}
