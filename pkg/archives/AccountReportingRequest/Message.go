package ArchiveAccountReportingRequest

import (
	"fmt"
	"os"
	"reflect"
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
	if Doc05, ok := doc.(*camt_060_001_05.Document); ok {
		pathMap := map[string]string{
			"AcctRptgReq.GrpHdr.MsgId":               "MessageId",
			"AcctRptgReq.GrpHdr.CreDtTm":             "CreatedDateTime",
			"AcctRptgReq.RptgReq[0].Id":              "ReportRequestId",
			"AcctRptgReq.RptgReq[0].ReqdMsgNmId":     "RequestedMsgNameId",
			"AcctRptgReq.RptgReq[0].Acct.Id.Othr.Id": "AccountOtherId",
			"AcctRptgReq.RptgReq[0].Acct.Tp.Prtry":   "AccountProperty",
		}
		// Iterate over the original map and copy values
		for sourcePath, targetPath := range pathMap {
			err = Archive.CopyDocumentValueToMessage(Doc05, sourcePath, &dataModel, targetPath)
			if err != nil {
				return MessageModel{}, fmt.Errorf("failed to copy value from %s to %s: %w", sourcePath, targetPath, err)
			}
		}
		return dataModel, nil
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel) (Archive.IOSDocument, error) {
	return nil, nil
}

func ReadXMLFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", filename, err)
	}
	return data, nil
}

func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}
