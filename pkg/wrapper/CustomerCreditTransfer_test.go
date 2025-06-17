package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	CustomerCreditTransfer "github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
)

// createValidModel creates a CustomerCreditTransfer.MessageModel with all required fields populated
func createValidModel() CustomerCreditTransfer.MessageModel {
	// Initialize with valid dates to avoid "0000-00-00" marshaling
	currentDate := fedwire.ISODate{Year: 2024, Month: 1, Day: 1}

	return CustomerCreditTransfer.MessageModel{
		PaymentCore: base.PaymentCore{
			MessageHeader: base.MessageHeader{
				MessageId:       "MSG123",
				CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			NumberOfTransactions:  "1",
			SettlementMethod:      models.SettlementCLRG,
			CommonClearingSysCode: "FDW",
		},
		InstructionId:       "INSTR123",
		EndToEndId:          "E2E123",
		TaxId:               "TX123",
		InstrumentPropCode:  models.InstrumentCTRC,
		InterBankSettAmount: models.CurrencyAndAmount{Amount: 1000.00, Currency: "USD"},
		InterBankSettDate:   currentDate,
		InstructedAmount:    models.CurrencyAndAmount{Amount: 1000.00, Currency: "USD"},
		ChargeBearer:        "SLEV",
		AgentPair: base.AgentPair{
			InstructingAgent: models.Agent{PaymentSysCode: models.PaymentSysUSABA, PaymentSysMemberId: "123456789"},
			InstructedAgent:  models.Agent{PaymentSysCode: models.PaymentSysUSABA, PaymentSysMemberId: "987654321"},
		},
		DebtorCreditorPair: base.DebtorCreditorPair{
			DebtorAgent:   models.Agent{PaymentSysCode: models.PaymentSysUSABA, PaymentSysMemberId: "111111111"},
			CreditorAgent: models.Agent{PaymentSysCode: models.PaymentSysUSABA, PaymentSysMemberId: "222222222"},
		},
		DebtorName:    "John Doe",
		DebtorAddress: models.PostalAddress{StreetName: "123 Main St", TownName: "Anytown", Country: "US"},
		// Initialize RemittanceInfor with valid dates to prevent "0000-00-00" marshaling
		RemittanceInfor: CustomerCreditTransfer.RemittanceDocument{
			RelatedDate: currentDate,
			TaxDetail: CustomerCreditTransfer.TaxRecord{
				TaxPeriodYear: currentDate,
			},
		},
	}
}

func TestCustomerCreditTransferWrapper_CreateDocument(t *testing.T) {
	wrapper := &CustomerCreditTransferWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     CustomerCreditTransfer.PACS_008_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"MessageId": "MSG123",
				"CreatedDateTime": "2024-01-01T10:00:00Z",
				"NumberOfTransactions": "1",
				"SettlementMethod": "CLRG",
				"CommonClearingSysCode": "FDW",
				"InstructionId": "INSTR123",
				"EndToEndId": "E2E123",
				"TaxId": "TX123",
				"InstrumentPropCode": "OTHR",
				"InterBankSettAmount": {"Amount": 1000.00, "Currency": "USD"},
				"InterBankSettDate": "2024-01-01",
				"InstructedAmount": {"Amount": 1000.00, "Currency": "USD"},
				"ChargeBearer": "SLEV",
				"InstructingAgents": {"PaymentSysCode": "USABA", "PaymentSysMemberId": "123456789"},
				"InstructedAgent": {"PaymentSysCode": "USABA", "PaymentSysMemberId": "987654321"},
				"DebtorName": "John Doe",
				"DebtorAddress": {"StreetName": "123 Main St", "TownName": "Anytown", "Country": "US"},
				"DebtorAgent": {"PaymentSysCode": "USABA", "PaymentSysMemberId": "111111111"},
				"CreditorAgent": {"PaymentSysCode": "USABA", "PaymentSysMemberId": "222222222"}
			}`),
			version:     CustomerCreditTransfer.PACS_008_001_02,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     CustomerCreditTransfer.PACS_008_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     CustomerCreditTransfer.PACS_008_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     CustomerCreditTransfer.PACS_008_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := wrapper.CreateDocument(tt.modelJson, tt.version)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				// Verify it's valid XML
				var xmlDoc interface{}
				err = xml.Unmarshal(result, &xmlDoc)
				assert.NoError(t, err, "Generated XML should be valid")
			}
		})
	}
}

func TestCustomerCreditTransferWrapper_ValidateDocument(t *testing.T) {
	wrapper := &CustomerCreditTransferWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     CustomerCreditTransfer.PACS_008_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"MessageId": "MSG123",
				"CreatedDateTime": "2024-01-01T10:00:00Z",
				"NumberOfTransactions": "1",
				"SettlementMethod": "CLRG",
				"CommonClearingSysCode": "FDW",
				"InstructionId": "INSTR123",
				"EndToEndId": "E2E123",
				"TaxId": "TX123",
				"InstrumentPropCode": "OTHR",
				"InterBankSettAmount": {"Amount": 1000.00, "Currency": "USD"},
				"InterBankSettDate": "2024-01-01",
				"InstructedAmount": {"Amount": 1000.00, "Currency": "USD"},
				"ChargeBearer": "SLEV",
				"InstructingAgents": {"PaymentSysCode": "USABA", "PaymentSysMemberId": "123456789"},
				"InstructedAgent": {"PaymentSysCode": "USABA", "PaymentSysMemberId": "987654321"},
				"DebtorName": "John Doe",
				"DebtorAddress": {"StreetName": "123 Main St", "TownName": "Anytown", "Country": "US"},
				"DebtorAgent": {"PaymentSysCode": "USABA", "PaymentSysMemberId": "111111111"},
				"CreditorAgent": {"PaymentSysCode": "USABA", "PaymentSysMemberId": "222222222"}
			}`,
			version:     CustomerCreditTransfer.PACS_008_001_02,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     CustomerCreditTransfer.PACS_008_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     CustomerCreditTransfer.PACS_008_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := wrapper.ValidateDocument(tt.modelJson, tt.version)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCustomerCreditTransferWrapper_CheckRequireField(t *testing.T) {
	wrapper := &CustomerCreditTransferWrapper{}

	tests := []struct {
		name        string
		model       CustomerCreditTransfer.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: CustomerCreditTransfer.MessageModel{
				PaymentCore: base.PaymentCore{
					MessageHeader: base.MessageHeader{
						// Missing MessageId
						CreatedDateTime: time.Now(),
					},
					NumberOfTransactions: "1",
				},
				InstructionId: "INSTR123",
				EndToEndId:    "E2E123",
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       CustomerCreditTransfer.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := wrapper.CheckRequireField(tt.model)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCustomerCreditTransferWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &CustomerCreditTransferWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.02">
	<FIToFICstmrCdtTrf>
		<GrpHdr>
			<MsgId>MSG123</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
			<NbOfTxs>1</NbOfTxs>
			<SttlmInf>
				<SttlmMtd>CLRG</SttlmMtd>
				<ClrSys>
					<Prtry>FDW</Prtry>
				</ClrSys>
			</SttlmInf>
		</GrpHdr>
		<CdtTrfTxInf>
			<InstrId>INSTR123</InstrId>
			<EndToEndId>E2E123</EndToEndId>
		</CdtTrfTxInf>
	</FIToFICstmrCdtTrf>
</Document>`)

	tests := []struct {
		name        string
		xmlData     []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid XML converts with validation error (expected due to minimal XML)",
			xmlData:     validXML,
			expectError: true,
			errorMsg:    "failed to convert XML to model",
		},
		{
			name:        "invalid XML returns error",
			xmlData:     []byte(`<invalid>xml`),
			expectError: true,
			errorMsg:    "failed to convert XML to model",
		},
		{
			name:        "empty XML returns error",
			xmlData:     []byte(``),
			expectError: true,
			errorMsg:    "failed to convert XML to model",
		},
		{
			name:        "nil XML returns error",
			xmlData:     nil,
			expectError: true,
			errorMsg:    "failed to convert XML to model",
		},
		{
			name:        "malformed XML returns error",
			xmlData:     []byte(`<?xml version="1.0"?><Document>missing namespace</Document>`),
			expectError: true,
			errorMsg:    "failed to convert XML to model",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := wrapper.ConvertXMLToModel(tt.xmlData)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				// For error cases, result should be zero value
				assert.Equal(t, CustomerCreditTransfer.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, CustomerCreditTransfer.MessageModel{}, result)
			}
		})
	}
}

func TestCustomerCreditTransferWrapper_GetHelp(t *testing.T) {
	wrapper := &CustomerCreditTransferWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(result), &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")

	// Verify it contains expected fields
	assert.Contains(t, result, "MessageId")
	assert.Contains(t, result, "CreatedDateTime")
	assert.Contains(t, result, "InstructionId")
}

func TestCustomerCreditTransferWrapper_Integration(t *testing.T) {
	wrapper := &CustomerCreditTransferWrapper{}

	// Test complete round-trip: Model -> JSON -> XML -> Model
	originalModel := createValidModel()

	// Model to JSON
	jsonData, err := json.Marshal(originalModel)
	require.NoError(t, err)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(jsonData, CustomerCreditTransfer.PACS_008_001_02)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// XML back to Model
	resultModel, err := wrapper.ConvertXMLToModel(xmlData)
	require.NoError(t, err)

	// Verify key fields are preserved
	assert.Equal(t, originalModel.MessageId, resultModel.MessageId)
	assert.Equal(t, originalModel.InstructionId, resultModel.InstructionId)
	assert.Equal(t, originalModel.EndToEndId, resultModel.EndToEndId)
	assert.Equal(t, originalModel.NumberOfTransactions, resultModel.NumberOfTransactions)
}

func TestCustomerCreditTransferWrapper_AllVersions(t *testing.T) {
	wrapper := &CustomerCreditTransferWrapper{}

	validModel := createValidModel()
	modelJson, err := json.Marshal(validModel)
	require.NoError(t, err)

	versions := []CustomerCreditTransfer.PACS_008_001_VERSION{
		CustomerCreditTransfer.PACS_008_001_02,
		CustomerCreditTransfer.PACS_008_001_03,
		CustomerCreditTransfer.PACS_008_001_04,
		CustomerCreditTransfer.PACS_008_001_05,
		CustomerCreditTransfer.PACS_008_001_06,
		CustomerCreditTransfer.PACS_008_001_07,
		CustomerCreditTransfer.PACS_008_001_08,
		CustomerCreditTransfer.PACS_008_001_09,
		CustomerCreditTransfer.PACS_008_001_10,
		CustomerCreditTransfer.PACS_008_001_11,
		CustomerCreditTransfer.PACS_008_001_12,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(modelJson, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := CustomerCreditTransfer.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}
