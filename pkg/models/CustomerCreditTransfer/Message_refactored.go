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
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
)

// RefactoredMessageModel demonstrates using base abstractions
type RefactoredMessageModel struct {
	// Embed common payment fields instead of duplicating them
	base.PaymentCore `json:",inline"`

	// Message-specific fields
	InstructionId                string                        `json:"instructionId"`
	EndToEndId                   string                        `json:"endToEndId"`
	TaxId                        string                        `json:"taxId"`
	UniqueEndToEndTransactionRef string                        `json:"uniqueEndToEndTransactionRef"`
	ServiceLevel                 string                        `json:"serviceLevel"`
	InstrumentPropCode           models.InstrumentPropCodeType `json:"instrumentPropCode"`
	InterBankSettAmount          models.CurrencyAndAmount      `json:"interBankSettAmount"`
	InterBankSettDate            fedwire.ISODate               `json:"interBankSettDate"`
	InstructedAmount             models.CurrencyAndAmount      `json:"instructedAmount"`
	ExchangeRate                 float64                       `json:"exchangeRate"`
	ChargeBearer                 models.ChargeBearerType       `json:"chargeBearer"`
	ChargesInfo                  []ChargeInfo                  `json:"chargesInfo"`

	// Use embedded agent pairs
	base.AgentPair          `json:",inline"`
	base.DebtorCreditorPair `json:",inline"`

	// Party information
	IntermediaryAgent1Id    string                      `json:"intermediaryAgent1Id"`
	UltimateDebtorName      string                      `json:"ultimateDebtorName"`
	UltimateDebtorAddress   models.PostalAddress        `json:"ultimateDebtorAddress"`
	DebtorName              string                      `json:"debtorName"`
	DebtorAddress           models.PostalAddress        `json:"debtorAddress"`
	DebtorIBAN              string                      `json:"debtorIBAN"`
	DebtorOtherTypeId       string                      `json:"debtorOtherTypeId"`
	CreditorName            string                      `json:"creditorName"`
	CreditorPostalAddress   models.PostalAddress        `json:"creditorPostalAddress"`
	UltimateCreditorName    string                      `json:"ultimateCreditorName"`
	UltimateCreditorAddress models.PostalAddress        `json:"ultimateCreditorAddress"`
	CreditorIBAN            string                      `json:"creditorIBAN"`
	CreditorOtherTypeId     string                      `json:"creditorOtherTypeId"`
	PurposeOfPayment        models.PurposeOfPaymentType `json:"purposeOfPayment"`
	RelatedRemittanceInfo   RemittanceDetail            `json:"relatedRemittanceInfo"`
	RemittanceInfor         RemittanceDocument          `json:"remittanceInfor"`
}

// Global processor instance using the base abstraction
var refactoredProcessor *base.MessageProcessor[RefactoredMessageModel, PACS_008_001_VERSION]

// init sets up the refactored processor using the base abstractions
func init() {

	// Register all versions - much cleaner than manual map building
	registrations := []base.FactoryRegistration[models.ISODocument, PACS_008_001_VERSION]{
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.02",
			Version:   PACS_008_001_02,
			Factory: func() models.ISODocument {
				return &pacs_008_001_02.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_02], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.03",
			Version:   PACS_008_001_03,
			Factory: func() models.ISODocument {
				return &pacs_008_001_03.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_03], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.04",
			Version:   PACS_008_001_04,
			Factory: func() models.ISODocument {
				return &pacs_008_001_04.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_04], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.05",
			Version:   PACS_008_001_05,
			Factory: func() models.ISODocument {
				return &pacs_008_001_05.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_05], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.06",
			Version:   PACS_008_001_06,
			Factory: func() models.ISODocument {
				return &pacs_008_001_06.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_06], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.07",
			Version:   PACS_008_001_07,
			Factory: func() models.ISODocument {
				return &pacs_008_001_07.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_07], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08",
			Version:   PACS_008_001_08,
			Factory: func() models.ISODocument {
				return &pacs_008_001_08.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_08], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09",
			Version:   PACS_008_001_09,
			Factory: func() models.ISODocument {
				return &pacs_008_001_09.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_09], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.10",
			Version:   PACS_008_001_10,
			Factory: func() models.ISODocument {
				return &pacs_008_001_10.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_10], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.11",
			Version:   PACS_008_001_11,
			Factory: func() models.ISODocument {
				return &pacs_008_001_11.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_11], Local: "Document"}}
			},
		},
		{
			Namespace: "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.12",
			Version:   PACS_008_001_12,
			Factory: func() models.ISODocument {
				return &pacs_008_001_12.Document{XMLName: xml.Name{Space: VersionNameSpaceMap[PACS_008_001_12], Local: "Document"}}
			},
		},
	}

	versionedFactory := base.BuildFactoryFromRegistrations(registrations)

	// Create the processor using base abstractions
	refactoredProcessor = base.NewMessageProcessor[RefactoredMessageModel, PACS_008_001_VERSION](
		versionedFactory.BuildNameSpaceModelMap(),
		versionedFactory.GetVersionMap(),
		VersionPathMap,
		RequiredFields,
	)
}

// RefactoredMessageWith demonstrates the simplified message processing using base abstractions
// This replaces the entire MessageWith function with a single call
func RefactoredMessageWith(data []byte) (RefactoredMessageModel, error) {
	return refactoredProcessor.ProcessMessage(data)
}

// RefactoredDocumentWith demonstrates the simplified document creation using base abstractions
// This replaces the entire DocumentWith function with a single call
func RefactoredDocumentWith(model RefactoredMessageModel, version PACS_008_001_VERSION) (models.ISODocument, error) {
	return refactoredProcessor.CreateDocument(model, version)
}

// RefactoredCheckRequiredFields demonstrates the simplified validation using base abstractions
// This replaces the entire CheckRequiredFields function with a single call
func RefactoredCheckRequiredFields(model RefactoredMessageModel) error {
	return refactoredProcessor.ValidateRequiredFields(model)
}

// RefactoredCustomerCreditTransferDataModel creates a new message model with embedded base types
func RefactoredCustomerCreditTransferDataModel() RefactoredMessageModel {
	return RefactoredMessageModel{
		PaymentCore: base.PaymentCore{
			MessageHeader: base.MessageHeader{
				MessageId:       "",
				CreatedDateTime: time.Time{},
			},
			NumberOfTransactions:  "",
			SettlementMethod:      "",
			CommonClearingSysCode: "",
		},
		// All other fields initialize to their zero values
	}
}
