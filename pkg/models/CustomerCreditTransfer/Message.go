package CustomerCreditTransfer

import (
	"encoding/xml"
	"fmt"
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
	SericeLevel                  string
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
func DocumentWith(model MessageModel, version PACS_008_001_VESION) (models.ISODocument, error) {
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
			continue // Or handle unknown field name
		}
		if models.IsEmpty(value) {
			return fmt.Errorf("required field %s is missing", field)
		}
	}

	return nil
}
