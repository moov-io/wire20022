package CustomerCreditTransfer

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/civil"
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

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
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

// ReadXML reads XML data from an io.Reader into the MessageModel
func (m *MessageModel) ReadXML(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("reading XML: %w", err)
	}

	model, err := processor.ProcessMessage(data)
	if err != nil {
		return err
	}

	*m = model
	return nil
}

// WriteXML writes the MessageModel as XML to an io.Writer
// If no version is specified, uses the latest version (PACS_008_001_12)
func (m *MessageModel) WriteXML(w io.Writer, version ...PACS_008_001_VERSION) error {
	// Default to latest version
	ver := PACS_008_001_12
	if len(version) > 0 {
		ver = version[0]
	}

	// Create versioned document
	doc, err := DocumentWith(*m, ver)
	if err != nil {
		return fmt.Errorf("creating document: %w", err)
	}

	// Write XML with proper formatting
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")

	// Write XML declaration
	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return fmt.Errorf("writing XML header: %w", err)
	}

	// Encode document
	if err := encoder.Encode(doc); err != nil {
		return fmt.Errorf("encoding XML: %w", err)
	}

	return encoder.Flush()
}

var RequiredFields = []string{
	"MessageId", "CreatedDateTime", "NumberOfTransactions",
	"SettlementMethod", "CommonClearingSysCode", "InstructionId",
	"EndToEndId", "InstrumentPropCode",
	"InterBankSettAmount", "InterBankSettDate", "InstructedAmount",
	"ChargeBearer", "InstructingAgent", "InstructedAgent",
	"DebtorName", "DebtorAddress", "DebtorAgent",
	"CreditorAgent",
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, PACS_008_001_VERSION]

// init sets up the processor using base abstractions
func init() {
	// Register all versions using cleaner factory registration pattern
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
	processor = base.NewMessageProcessor[MessageModel, PACS_008_001_VERSION](
		versionedFactory.BuildNameSpaceModelMap(),
		versionedFactory.GetVersionMap(),
		VersionPathMap,
		RequiredFields,
	)
}

// ParseXML reads XML data into the MessageModel
// This replaces the non-idiomatic MessageWith function
func ParseXML(data []byte) (*MessageModel, error) {
	model, err := processor.ProcessMessage(data)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

// DocumentWith uses base abstractions to replace 20+ lines with a single call
func DocumentWith(model MessageModel, version PACS_008_001_VERSION) (models.ISODocument, error) {
	// Validate required fields before creating document
	if err := processor.ValidateRequiredFields(model); err != nil {
		return nil, err
	}
	return processor.CreateDocument(model, version)
}

// CheckRequiredFields uses base abstractions to replace 35+ lines with a single call
func CheckRequiredFields(model MessageModel) error {
	return processor.ValidateRequiredFields(model)
}

// CustomerCreditTransferDataModel creates a new message model with sample data for testing
func CustomerCreditTransferDataModel() MessageModel {
	return MessageModel{
		PaymentCore: base.PaymentCore{
			MessageHeader: base.MessageHeader{
				MessageId:       "20250310B1QDRCQR000001",
				CreatedDateTime: time.Now(),
			},
			NumberOfTransactions:  "1",
			SettlementMethod:      "CLRG",
			CommonClearingSysCode: "FDW",
		},
		InstructionId:                "Scenario01InstrId001",
		EndToEndId:                   "Scenario01EtoEId001",
		UniqueEndToEndTransactionRef: "8a562c67-ca16-48ba-b074-65581be6f011",
		TaxId:                        "123456789",
		InstrumentPropCode:           "CTRC",
		InterBankSettAmount: models.CurrencyAndAmount{
			Currency: "USD", Amount: 510000.74,
		},
		InterBankSettDate: fedwire.ISODate(civil.DateOf(time.Now())),
		InstructedAmount: models.CurrencyAndAmount{
			Currency: "USD", Amount: 510000.74,
		},
		ChargeBearer: "SLEV",
		ChargesInfo: []ChargeInfo{
			{
				Amount:         models.CurrencyAndAmount{Currency: "USD", Amount: 90.00},
				BusinessIdCode: "BANZBEBB",
			},
			{
				Amount:         models.CurrencyAndAmount{Currency: "USD", Amount: 40.00},
				BusinessIdCode: "BANCUS33",
			},
		},
		AgentPair: base.AgentPair{
			InstructingAgent: models.Agent{
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "011104238",
			},
			InstructedAgent: models.Agent{
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "021040078",
			},
		},
		DebtorCreditorPair: base.DebtorCreditorPair{
			DebtorAgent: models.Agent{
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "011104238",
				BankName:           "Bank A",
				PostalAddress: models.PostalAddress{
					StreetName:     "Avenue A",
					BuildingNumber: "66",
					PostalCode:     "60532",
					TownName:       "Lisle",
					Subdivision:    "IL",
					Country:        "US",
				},
			},
			CreditorAgent: models.Agent{
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "021040078",
				BankName:           "Bank B",
				PostalAddress: models.PostalAddress{
					StreetName:     "Avenue B",
					BuildingNumber: "25",
					PostalCode:     "19067",
					TownName:       "Yardley",
					Subdivision:    "PA",
					Country:        "US",
				},
			},
		},
		DebtorName: "Corporation A",
		DebtorAddress: models.PostalAddress{
			StreetName:     "Avenue of the Fountains",
			BuildingNumber: "167565",
			RoomNumber:     "Suite D110",
			PostalCode:     "85268",
			TownName:       "Fountain Hills",
			Subdivision:    "AZ",
			Country:        "US",
		},
		DebtorOtherTypeId: "5647772655",
		CreditorName:      "Corporation B",
		CreditorPostalAddress: models.PostalAddress{
			StreetName:     "Desert View Street",
			BuildingNumber: "1",
			Floor:          "33",
			PostalCode:     "19067",
			TownName:       "Palm Springs",
			Subdivision:    "CA",
			Country:        "US",
		},
		CreditorOtherTypeId: "567876543",
		RemittanceInfor: RemittanceDocument{
			CodeOrProprietary: models.CodeCINV,
			Number:            "INV34563",
			RelatedDate:       fedwire.ISODate(civil.DateOf(time.Now())),
			TaxDetail: TaxRecord{
				TaxId:              "123456789",
				TaxTypeCode:        "09455",
				TaxPeriodYear:      fedwire.ISODate(civil.DateOf(time.Now())),
				TaxperiodTimeFrame: "MM04",
			},
		},
		RelatedRemittanceInfo: RemittanceDetail{
			RemittanceId:      "Scenario01Var2RemittanceId001",
			Method:            models.Email,
			ElectronicAddress: "CustomerService@CorporationB.com",
		},
	}
}
