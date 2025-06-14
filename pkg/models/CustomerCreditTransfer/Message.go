package CustomerCreditTransfer

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_02"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_03"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_04"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_05"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_06"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_07"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_08"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_09"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_10"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_11"
	"github.com/moov-io/fedwire20022/gen/CustomerCreditTransfer/pacs_008_001_12"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/errors"
	"github.com/moov-io/wire20022/pkg/models"
)

type MessageModel struct {
	MessageId                    string
	CreatedDateTime              time.Time
	NumberOfTransactions         string
	SettlementMethod             models.SettlementMethodType
	CommonClearingSysCode        models.CommonClearingSysCodeType
	InstructionId                string
	EndToEndId                   string
	TaxId                        string
	UniqueEndToEndTransactionRef string
	ServiceLevel                 string
	InstrumentPropCode           models.InstrumentPropCodeType
	InterBankSettAmount          models.CurrencyAndAmount
	InterBankSettDate            fedwire.ISODate
	InstructedAmount             models.CurrencyAndAmount
	ExchangeRate                 float64
	ChargeBearer                 models.ChargeBearerType
	ChargesInfo                  []ChargeInfo
	InstructingAgents            models.Agent
	InstructedAgent              models.Agent
	IntermediaryAgent1Id         string
	UltimateDebtorName           string
	UltimateDebtorAddress        models.PostalAddress
	DebtorName                   string
	DebtorAddress                models.PostalAddress
	DebtorIBAN                   string
	DebtorOtherTypeId            string
	DebtorAgent                  models.Agent
	CreditorAgent                models.Agent
	CreditorName                 string
	CreditorPostalAddress        models.PostalAddress
	UltimateCreditorName         string
	UltimateCreditorAddress      models.PostalAddress
	CreditorIBAN                 string
	CreditorOtherTypeId          string
	PurposeOfPayment             models.PurposeOfPaymentType
	RelatedRemittanceInfo        RemittanceDetail
	RemittanceInfor              RemittanceDocument
}

var NameSpaceModelMap = map[string]models.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.02": func() models.ISODocument {
		return &pacs_008_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.03": func() models.ISODocument {
		return &pacs_008_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.04": func() models.ISODocument {
		return &pacs_008_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.05": func() models.ISODocument {
		return &pacs_008_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.06": func() models.ISODocument {
		return &pacs_008_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.07": func() models.ISODocument {
		return &pacs_008_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_07], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08": func() models.ISODocument {
		return &pacs_008_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_08], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09": func() models.ISODocument {
		return &pacs_008_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_09], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.10": func() models.ISODocument {
		return &pacs_008_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_10], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.11": func() models.ISODocument {
		return &pacs_008_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_11], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.12": func() models.ISODocument {
		return &pacs_008_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_12], Local: "Document"}}
	},
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "NumberOfTransactions",
	"SettlementMethod", "CommonClearingSysCode", "InstructionId",
	"EndToEndId", "InstrumentPropCode",
	"InterBankSettAmount", "InterBankSettDate", "InstructedAmount",
	"ChargeBearer", "InstructingAgents", "InstructedAgent",
	"DebtorName", "DebtorAddress", "DebtorAgent",
	"CreditorAgent",
}

func MessageWith(data []byte) (MessageModel, error) {
	doc, xmlns, err := models.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, errors.NewParseError("document creation", "XML data", err)
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
func DocumentWith(model MessageModel, version PACS_008_001_VERSION) (models.ISODocument, error) {
	if err := CheckRequiredFields(model); err != nil {
		return nil, err
	}

	pathMap, pathExists := VersionPathMap[version]
	factory, factoryExists := NameSpaceModelMap[VersionNameSpaceMap[version]]
	if !pathExists || !factoryExists {
		return nil, errors.NewInvalidFieldError("version", "unsupported document version")
	}

	document := factory()
	rePathMap := models.RemakeMapping(model, pathMap, false)
	for targetPath, sourcePath := range rePathMap {
		if err := models.CopyMessageValueToDocument(model, sourcePath, document, targetPath); err != nil {
			return document, errors.NewFieldError(targetPath, "copy", err)
		}
	}
	return document, nil
}
func CheckRequiredFields(model MessageModel) error {
	fieldMap := map[string]interface{}{
		"MessageId":             model.MessageId,
		"CreatedDateTime":       model.CreatedDateTime,
		"NumberOfTransactions":  model.NumberOfTransactions,
		"SettlementMethod":      model.SettlementMethod,
		"CommonClearingSysCode": model.CommonClearingSysCode,
		"InstructionId":         model.InstructionId,
		"EndToEndId":            model.EndToEndId,
		"InstrumentPropCode":    model.InstrumentPropCode,
		"InterBankSettAmount":   model.InterBankSettAmount,
		"InterBankSettDate":     model.InterBankSettDate,
		"InstructedAmount":      model.InstructedAmount,
		"ChargeBearer":          model.ChargeBearer,
		"InstructingAgents":     model.InstructingAgents,
		"InstructedAgent":       model.InstructedAgent,
		"DebtorName":            model.DebtorName,
		"DebtorAddress":         model.DebtorAddress,
		"DebtorAgent":           model.DebtorAgent,
		"CreditorAgent":         model.CreditorAgent,
	}

	for _, field := range RequiredFields {
		value, ok := fieldMap[field]
		if !ok {
			continue
		}
		if models.IsEmpty(value) {
			return errors.NewRequiredFieldError(field)
		}
	}

	return nil
}
