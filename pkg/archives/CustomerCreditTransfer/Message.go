package CustomerCreditTransfer

import (
	"encoding/xml"
	"fmt"
	"reflect"
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
	Archive "github.com/moov-io/wire20022/pkg/archives"
)

type MessageModel struct {
	MessageId string
	CreatedDateTime time.Time
	NumberOfTransactions string
	SettlementMethod Archive.SettlementMethodType
	CommonClearingSysCode Archive.CommonClearingSysCodeType
	InstructionId string
	EndToEndId string
	TaxId string
	UniqueEndToEndTransactionRef string
	SericeLevel string
	InstrumentPropCode Archive.InstrumentPropCodeType
	InterBankSettAmount Archive.CurrencyAndAmount
	InterBankSettDate fedwire.ISODate
	InstructedAmount Archive.CurrencyAndAmount
	ExchangeRate float64
	ChargeBearer Archive.ChargeBearerType
	ChargesInfo []ChargeInfo
	InstructingAgents Archive.Agent
	InstructedAgent      Archive.Agent
	IntermediaryAgent1Id string
	UltimateDebtorName    string
	UltimateDebtorAddress Archive.PostalAddress
	DebtorName string
	DebtorAddress Archive.PostalAddress
	DebtorIBAN string
	DebtorOtherTypeId string
	DebtorAgent Archive.Agent
	CreditorAgent Archive.Agent
	CreditorName string
	CreditorPostalAddress Archive.PostalAddress
	UltimateCreditorName    string
	UltimateCreditorAddress Archive.PostalAddress
	CreditorIBAN        string
	CreditorOtherTypeId string
	PurposeOfPayment    Archive.PurposeOfPaymentType
	RelatedRemittanceInfo RemittanceDetail
	RemittanceInfor RemittanceDocument
}
var NameSpaceModelMap = map[string]Archive.DocumentFactory{
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.02": func() Archive.ISODocument {
		return &pacs_008_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_02], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.03": func() Archive.ISODocument {
		return &pacs_008_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_03], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.04": func() Archive.ISODocument {
		return &pacs_008_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_04], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.05": func() Archive.ISODocument {
		return &pacs_008_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_05], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.06": func() Archive.ISODocument {
		return &pacs_008_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_06], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.07": func() Archive.ISODocument {
		return &pacs_008_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_07], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08": func() Archive.ISODocument {
		return &pacs_008_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_08], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09": func() Archive.ISODocument {
		return &pacs_008_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_09], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.10": func() Archive.ISODocument {
		return &pacs_008_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_10], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.11": func() Archive.ISODocument {
		return &pacs_008_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_11], Local: "Document"}}
	},
	"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.12": func() Archive.ISODocument {
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
	doc, xmlns, err := Archive.DocumentFrom(data, NameSpaceModelMap)
	if err != nil {
		return MessageModel{}, fmt.Errorf("failed to create document: %w", err)
	}
	version := NameSpaceVersonMap[xmlns]

	dataModel := MessageModel{}
	pathMap := VersionPathMap[version]
	rePathMap := Archive.RemakeMapping(doc, pathMap, true)
	for sourcePath, targetPath := range rePathMap {
		Archive.CopyDocumentValueToMessage(doc, sourcePath, &dataModel, targetPath)
	}
	return dataModel, nil
}
func DocumentWith(model MessageModel, version PACS_008_001_VESION) (Archive.ISODocument, error) {
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
	rePathMap := Archive.RemakeMapping(model, pathMap, false)
	for targetPath, sourcePath := range rePathMap {
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
				return fmt.Errorf("required field %s is missing", field)
			}
		case "CreatedDateTime":
			if model.CreatedDateTime.IsZero() {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "NumberOfTransactions":
			if model.NumberOfTransactions == "" {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "SettlementMethod":
			if model.SettlementMethod == "" {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "CommonClearingSysCode":
			if model.CommonClearingSysCode == "" {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "InstructionId":
			if model.InstructionId == "" {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "EndToEndId":
			if model.EndToEndId == "" {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "InstrumentPropCode":
			if model.InstrumentPropCode == "" {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "InterBankSettAmount":
			if isEmpty(model.InterBankSettAmount) {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "InterBankSettDate":
			if isEmpty(model.InterBankSettDate) {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "InstructedAmount":
			if isEmpty(model.InstructedAmount) {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "ChargeBearer":
			if model.ChargeBearer == "" {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "InstructingAgents":
			if isEmpty(model.InstructingAgents) {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "InstructedAgent":
			if isEmpty(model.InstructedAgent) {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "DebtorName":
			if model.DebtorName == "" {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "DebtorAddress":
			if isEmpty(model.DebtorAddress) {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "DebtorAgent":
			if isEmpty(model.DebtorAgent) {
				return fmt.Errorf("required field %s is missing", field)
			}
		case "CreditorAgent":
			if isEmpty(model.CreditorAgent) {
				return fmt.Errorf("required field %s is missing", field)
			}
		default:
			return nil
		}
	}
	return nil
}
func isEmpty[T any](s T) bool {
	var zero T // Declare a zero value of type T
	return reflect.DeepEqual(s, zero)
}