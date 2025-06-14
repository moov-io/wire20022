package PaymentReturn

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_02"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_03"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_04"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_05"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_06"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_07"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_08"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_09"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_10"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_11"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_12"
	"github.com/moov-io/fedwire20022/gen/PaymentReturn/pacs_004_001_13"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/errors"
	"github.com/moov-io/wire20022/pkg/models"
)

type MessageModel struct {
	MessageId                         string
	CreatedDateTime                   time.Time
	NumberOfTransactions              string
	SettlementMethod                  models.SettlementMethodType
	ClearingSystem                    models.CommonClearingSysCodeType
	OriginalMessageId                 string
	OriginalMessageNameId             string
	OriginalCreationDateTime          time.Time
	OriginalInstructionId             string
	OriginalEndToEndId                string
	OriginalUETR                      string
	ReturnedInterbankSettlementAmount models.CurrencyAndAmount
	InterbankSettlementDate           fedwire.ISODate
	ReturnedInstructedAmount          models.CurrencyAndAmount
	ChargeBearer                      models.ChargeBearerType
	InstructingAgent                  models.Agent
	InstructedAgent                   models.Agent
	RtrChain                          models.ReturnChain
	ReturnReasonInformation           models.Reason
	OriginalTransactionRef            models.InstrumentPropCodeType
}

var NameSpaceModelMap = map[string]models.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.02": func() models.ISODocument {
		return &pacs_004_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.03": func() models.ISODocument {
		return &pacs_004_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.04": func() models.ISODocument {
		return &pacs_004_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.05": func() models.ISODocument {
		return &pacs_004_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.06": func() models.ISODocument {
		return &pacs_004_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.07": func() models.ISODocument {
		return &pacs_004_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_07], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.08": func() models.ISODocument {
		return &pacs_004_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_08], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.09": func() models.ISODocument {
		return &pacs_004_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_09], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10": func() models.ISODocument {
		return &pacs_004_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_10], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.11": func() models.ISODocument {
		return &pacs_004_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_11], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.12": func() models.ISODocument {
		return &pacs_004_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_12], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.13": func() models.ISODocument {
		return &pacs_004_001_13.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_004_001_13], Local: "Document"}}
	},
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
func DocumentWith(model MessageModel, version PACS_004_001_VERSION) (models.ISODocument, error) {
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
			return document, err
		}
	}
	return document, nil
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "NumberOfTransactions", "SettlementMethod", "ClearingSystem",
	"OriginalMessageId", "OriginalMessageNameId", "OriginalCreationDateTime",
	"ReturnedInterbankSettlementAmount", "InterbankSettlementDate", "ReturnedInstructedAmount",
	"InstructingAgent", "InstructedAgent", "ReturnReasonInformation", "OriginalTransactionRef",
}

func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"MessageId":                         model.MessageId,
		"CreatedDateTime":                   model.CreatedDateTime,
		"NumberOfTransactions":              model.NumberOfTransactions,
		"SettlementMethod":                  model.SettlementMethod,
		"ClearingSystem":                    model.ClearingSystem,
		"OriginalMessageId":                 model.OriginalMessageId,
		"OriginalMessageNameId":             model.OriginalMessageNameId,
		"OriginalCreationDateTime":          model.OriginalCreationDateTime,
		"ReturnedInterbankSettlementAmount": model.ReturnedInterbankSettlementAmount,
		"InterbankSettlementDate":           model.InterbankSettlementDate,
		"ReturnedInstructedAmount":          model.ReturnedInstructedAmount,
		"InstructingAgent":                  model.InstructingAgent,
		"InstructedAgent":                   model.InstructedAgent,
		"ReturnReasonInformation":           model.ReturnReasonInformation,
		"OriginalTransactionRef":            model.OriginalTransactionRef,
	}

	for _, field := range RequiredFields {
		value, ok := fieldMap[field]
		if !ok {
			continue
		}
		if models.IsEmpty(value) {
			return fmt.Errorf("missing required field: %s", field)
		}
	}

	return nil
}
