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
	PaymentReturn "github.com/moov-io/wire20022/pkg/models/PaymentReturn"
)

// createValidPaymentReturnModel creates a PaymentReturn.MessageModel with all required fields populated
func createValidPaymentReturnModel() PaymentReturn.MessageModel {
	return PaymentReturn.MessageModel{
		PaymentCore: base.PaymentCore{
			MessageHeader: base.MessageHeader{
				MessageId:       "20250310B1QDRCQR000401",
				CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			NumberOfTransactions:  "1",
			SettlementMethod:      models.SettlementCLRG,
			CommonClearingSysCode: "FDW",
		},
		AgentPair: base.AgentPair{
			InstructingAgent: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "021040078",
			},
			InstructedAgent: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "011104238",
			},
		},
		OriginalMessageId:        "20250310B1QDRCQR000001",
		OriginalMessageNameId:    "pacs.008.001.08",
		OriginalCreationDateTime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		OriginalInstructionId:    "Scenario01Step1InstrId001",
		OriginalEndToEndId:       "Scenario1EndToEndId001",
		OriginalUETR:             "8a562c67-ca16-48ba-b074-65581be6f011",
		ReturnedInterbankSettlementAmount: models.CurrencyAndAmount{
			Amount:   6000000.00,
			Currency: "USD",
		},
		InterbankSettlementDate: fedwire.ISODate{Year: 2024, Month: 1, Day: 1},
		ReturnedInstructedAmount: models.CurrencyAndAmount{
			Amount:   6000000.00,
			Currency: "USD",
		},
		ChargeBearer: models.ChargeBearerSLEV,
		ReturnReasonInformation: models.Reason{
			Reason: "AC04",
		},
		OriginalTransactionRef: models.InstrumentCTRC,
	}
}

func TestPaymentReturnWrapper_CreateDocument(t *testing.T) {
	wrapper := &PaymentReturnWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     PaymentReturn.PACS_004_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"messageId": "20250310B1QDRCQR000401",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"numberOfTransactions": "1",
				"settlementMethod": "CLRG",
				"commonClearingSysCode": "FDW",
				"instructingAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021040078"
				},
				"instructedAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "011104238"
				},
				"originalMessageId": "20250310B1QDRCQR000001",
				"originalMessageNameId": "pacs.008.001.08",
				"originalCreationDateTime": "2024-01-01T09:00:00Z",
				"originalTransactionRef": "CTRC"
			}`),
			version:     PaymentReturn.PACS_004_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     PaymentReturn.PACS_004_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     PaymentReturn.PACS_004_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     PaymentReturn.PACS_004_001_12,
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

func TestPaymentReturnWrapper_ValidateDocument(t *testing.T) {
	wrapper := &PaymentReturnWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     PaymentReturn.PACS_004_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"messageId": "20250310B1QDRCQR000401",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"numberOfTransactions": "1",
				"settlementMethod": "CLRG",
				"commonClearingSysCode": "FDW",
				"instructingAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021040078"
				},
				"instructedAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "011104238"
				},
				"originalMessageId": "20250310B1QDRCQR000001",
				"originalMessageNameId": "pacs.008.001.08",
				"originalCreationDateTime": "2024-01-01T09:00:00Z",
				"originalTransactionRef": "CTRC",
				"returnedInterbankSettlementAmount": {
					"amount": 6000000.00,
					"currency": "USD"
				},
				"interbankSettlementDate": "2024-01-01",
				"returnedInstructedAmount": {
					"amount": 6000000.00,
					"currency": "USD"
				},
				"chargeBearer": "SLEV",
				"returnReasonInformation": {
					"reason": "AC04"
				}
			}`,
			version:     PaymentReturn.PACS_004_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     PaymentReturn.PACS_004_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     PaymentReturn.PACS_004_001_12,
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

func TestPaymentReturnWrapper_CheckRequireField(t *testing.T) {
	wrapper := &PaymentReturnWrapper{}

	tests := []struct {
		name        string
		model       PaymentReturn.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidPaymentReturnModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: PaymentReturn.MessageModel{
				PaymentCore: base.PaymentCore{
					MessageHeader: base.MessageHeader{
						// Missing MessageId
						CreatedDateTime: time.Now(),
					},
				},
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       PaymentReturn.MessageModel{},
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

func TestPaymentReturnWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &PaymentReturnWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.004.001.12">
	<PmtRtr>
		<GrpHdr>
			<MsgId>20250310B1QDRCQR000401</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</GrpHdr>
		<OrgnlGrpInfAndSts>
			<OrgnlMsgId>20250310B1QDRCQR000001</OrgnlMsgId>
			<OrgnlMsgNmId>pacs.008.001.08</OrgnlMsgNmId>
		</OrgnlGrpInfAndSts>
	</PmtRtr>
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
			errorMsg:    "required field missing",
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
				assert.Equal(t, PaymentReturn.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, PaymentReturn.MessageModel{}, result)
			}
		})
	}
}

func TestPaymentReturnWrapper_GetHelp(t *testing.T) {
	wrapper := &PaymentReturnWrapper{}

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
	assert.Contains(t, result, "OriginalMessageId")
}

func TestPaymentReturnWrapper_Integration(t *testing.T) {
	wrapper := &PaymentReturnWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"messageId": "20250310B1QDRCQR000401",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"numberOfTransactions": "1",
		"settlementMethod": "CLRG",
		"commonClearingSysCode": "FDW",
		"instructingAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021040078"
		},
		"instructedAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "011104238"
		},
		"originalMessageId": "20250310B1QDRCQR000001",
		"originalMessageNameId": "pacs.008.001.08",
		"originalCreationDateTime": "2024-01-01T09:00:00Z",
		"originalTransactionRef": "CTRC",
		"returnedInterbankSettlementAmount": {
			"amount": 6000000.00,
			"currency": "USD"
		},
		"interbankSettlementDate": "2024-01-01",
		"returnedInstructedAmount": {
			"amount": 6000000.00,
			"currency": "USD"
		},
		"chargeBearer": "SLEV",
		"returnReasonInformation": {
			"reason": "AC04"
		}
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, PaymentReturn.PACS_004_001_12)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestPaymentReturnWrapper_AllVersions(t *testing.T) {
	wrapper := &PaymentReturnWrapper{}

	validJSON := []byte(`{
		"messageId": "20250310B1QDRCQR000401",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"numberOfTransactions": "1",
		"settlementMethod": "CLRG",
		"commonClearingSysCode": "FDW",
		"instructingAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021040078"
		},
		"instructedAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "011104238"
		},
		"originalMessageId": "20250310B1QDRCQR000001",
		"originalMessageNameId": "pacs.008.001.08",
		"originalCreationDateTime": "2024-01-01T09:00:00Z",
		"originalTransactionRef": "CTRC",
		"returnedInterbankSettlementAmount": {
			"amount": 6000000.00,
			"currency": "USD"
		},
		"interbankSettlementDate": "2024-01-01",
		"returnedInstructedAmount": {
			"amount": 6000000.00,
			"currency": "USD"
		},
		"chargeBearer": "SLEV",
		"returnReasonInformation": {
			"reason": "AC04"
		}
	}`)

	versions := []PaymentReturn.PACS_004_001_VERSION{
		PaymentReturn.PACS_004_001_02,
		PaymentReturn.PACS_004_001_03,
		PaymentReturn.PACS_004_001_04,
		PaymentReturn.PACS_004_001_05,
		PaymentReturn.PACS_004_001_06,
		PaymentReturn.PACS_004_001_07,
		PaymentReturn.PACS_004_001_08,
		PaymentReturn.PACS_004_001_09,
		PaymentReturn.PACS_004_001_10,
		PaymentReturn.PACS_004_001_11,
		PaymentReturn.PACS_004_001_12,
		PaymentReturn.PACS_004_001_13,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := PaymentReturn.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}
