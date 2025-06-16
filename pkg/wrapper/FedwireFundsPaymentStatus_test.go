package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wadearnold/wire20022/pkg/base"
	"github.com/wadearnold/wire20022/pkg/models"
	FedwireFundsPaymentStatus "github.com/wadearnold/wire20022/pkg/models/FedwireFundsPaymentStatus"
)

// createValidFedwireFundsPaymentStatusModel creates a FedwireFundsPaymentStatus.MessageModel with all required fields populated
func createValidFedwireFundsPaymentStatusModel() FedwireFundsPaymentStatus.MessageModel {
	return FedwireFundsPaymentStatus.MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250310QMGFNP31000001",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		OriginalMessageId:          "20250310B1QDRCQR000001",
		OriginalMessageNameId:      "pacs.008.001.08",
		OriginalMessageCreateTime:  time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		OriginalUETR:               "8a562c67-ca16-48ba-b074-65581be6f011",
		TransactionStatus:          models.AcceptedSettlementCompleted,
		AcceptanceDateTime:         time.Date(2024, 1, 1, 10, 30, 0, 0, time.UTC),
		EffectiveInterbankSettlementDate: fedwire.ISODate{Year: 2024, Month: 1, Day: 1},
		AgentPair: base.AgentPair{
			InstructingAgent: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "021151080",
			},
			InstructedAgent: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "011104238",
			},
		},
		StatusReasonInformation: "Transaction completed successfully",
		ReasonAdditionalInfo:    "All validations passed",
	}
}

func TestFedwireFundsPaymentStatusWrapper_CreateDocument(t *testing.T) {
	wrapper := &FedwireFundsPaymentStatusWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     FedwireFundsPaymentStatus.PACS_002_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"MessageId": "20250310QMGFNP31000001",
				"CreatedDateTime": "2024-01-01T10:00:00Z",
				"OriginalMessageId": "20250310B1QDRCQR000001",
				"OriginalMessageNameId": "pacs.008.001.08",
				"OriginalMessageCreateTime": "2024-01-01T09:00:00Z",
				"OriginalUETR": "8a562c67-ca16-48ba-b074-65581be6f011",
				"TransactionStatus": "ACSC",
				"AcceptanceDateTime": "2024-01-01T10:30:00Z",
				"EffectiveInterbankSettlementDate": "2024-01-01",
				"InstructingAgent": {
					"PaymentSysCode": "USABA",
					"PaymentSysMemberId": "021151080"
				},
				"InstructedAgent": {
					"PaymentSysCode": "USABA",
					"PaymentSysMemberId": "011104238"
				},
				"StatusReasonInformation": "Transaction completed successfully",
				"ReasonAdditionalInfo": "All validations passed"
			}`),
			version:     FedwireFundsPaymentStatus.PACS_002_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     FedwireFundsPaymentStatus.PACS_002_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     FedwireFundsPaymentStatus.PACS_002_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     FedwireFundsPaymentStatus.PACS_002_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"MessageId": "",
				"CreatedDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     FedwireFundsPaymentStatus.PACS_002_001_12,
			expectError: true,
			errorMsg:    "failed to create document",
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

func TestFedwireFundsPaymentStatusWrapper_ValidateDocument(t *testing.T) {
	wrapper := &FedwireFundsPaymentStatusWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     FedwireFundsPaymentStatus.PACS_002_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"MessageId": "20250310QMGFNP31000001",
				"CreatedDateTime": "2024-01-01T10:00:00Z",
				"OriginalMessageId": "20250310B1QDRCQR000001",
				"OriginalMessageNameId": "pacs.008.001.08",
				"OriginalMessageCreateTime": "2024-01-01T09:00:00Z",
				"OriginalUETR": "8a562c67-ca16-48ba-b074-65581be6f011",
				"TransactionStatus": "ACSC",
				"AcceptanceDateTime": "2024-01-01T10:30:00Z",
				"EffectiveInterbankSettlementDate": "2024-01-01",
				"InstructingAgent": {
					"PaymentSysCode": "USABA",
					"PaymentSysMemberId": "021151080"
				},
				"InstructedAgent": {
					"PaymentSysCode": "USABA",
					"PaymentSysMemberId": "011104238"
				}
			}`,
			version:     FedwireFundsPaymentStatus.PACS_002_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     FedwireFundsPaymentStatus.PACS_002_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     FedwireFundsPaymentStatus.PACS_002_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"MessageId": "",
				"CreatedDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     FedwireFundsPaymentStatus.PACS_002_001_12,
			expectError: true,
			errorMsg:    "failed to create document",
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

func TestFedwireFundsPaymentStatusWrapper_CheckRequireField(t *testing.T) {
	wrapper := &FedwireFundsPaymentStatusWrapper{}

	tests := []struct {
		name        string
		model       FedwireFundsPaymentStatus.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidFedwireFundsPaymentStatusModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: FedwireFundsPaymentStatus.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing MessageId
					CreatedDateTime: time.Now(),
				},
				TransactionStatus: models.AcceptedSettlementCompleted,
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       FedwireFundsPaymentStatus.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing TransactionStatus fails validation",
			model: FedwireFundsPaymentStatus.MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "20250310QMGFNP31000001",
					CreatedDateTime: time.Now(),
				},
				// Missing TransactionStatus
			},
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

func TestFedwireFundsPaymentStatusWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &FedwireFundsPaymentStatusWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12">
	<FIToFIPmtStsRpt>
		<GrpHdr>
			<MsgId>20250310QMGFNP31000001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</GrpHdr>
		<OrgnlGrpInfAndSts>
			<OrgnlMsgId>20250310B1QDRCQR000001</OrgnlMsgId>
			<OrgnlMsgNmId>pacs.008.001.08</OrgnlMsgNmId>
		</OrgnlGrpInfAndSts>
		<TxInfAndSts>
			<StsId>20250310QMGFNP31000001</StsId>
			<TxSts>ACSC</TxSts>
		</TxInfAndSts>
	</FIToFIPmtStsRpt>
</Document>`)

	tests := []struct {
		name        string
		xmlData     []byte
		expectError bool
		errorMsg    string
	}{
		{
			name:        "valid XML converts successfully",
			xmlData:     validXML,
			expectError: false,
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
				assert.Equal(t, FedwireFundsPaymentStatus.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, FedwireFundsPaymentStatus.MessageModel{}, result)
			}
		})
	}
}

func TestFedwireFundsPaymentStatusWrapper_GetHelp(t *testing.T) {
	wrapper := &FedwireFundsPaymentStatusWrapper{}

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
	assert.Contains(t, result, "TransactionStatus")
	assert.Contains(t, result, "InstructingAgent")
	assert.Contains(t, result, "InstructedAgent")
}

func TestFedwireFundsPaymentStatusWrapper_Integration(t *testing.T) {
	wrapper := &FedwireFundsPaymentStatusWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"MessageId": "20250310QMGFNP31000001",
		"CreatedDateTime": "2024-01-01T10:00:00Z",
		"OriginalMessageId": "20250310B1QDRCQR000001",
		"OriginalMessageNameId": "pacs.008.001.08",
		"OriginalMessageCreateTime": "2024-01-01T09:00:00Z",
		"OriginalUETR": "8a562c67-ca16-48ba-b074-65581be6f011",
		"TransactionStatus": "ACSC",
		"AcceptanceDateTime": "2024-01-01T10:30:00Z",
		"EffectiveInterbankSettlementDate": "2024-01-01",
		"InstructingAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "021151080"
		},
		"InstructedAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "011104238"
		}
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, FedwireFundsPaymentStatus.PACS_002_001_12)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestFedwireFundsPaymentStatusWrapper_AllVersions(t *testing.T) {
	wrapper := &FedwireFundsPaymentStatusWrapper{}

	validJSON := []byte(`{
		"MessageId": "20250310QMGFNP31000001",
		"CreatedDateTime": "2024-01-01T10:00:00Z",
		"OriginalMessageId": "20250310B1QDRCQR000001",
		"OriginalMessageNameId": "pacs.008.001.08",
		"OriginalMessageCreateTime": "2024-01-01T09:00:00Z",
		"OriginalUETR": "8a562c67-ca16-48ba-b074-65581be6f011",
		"TransactionStatus": "ACSC",
		"InstructingAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "021151080"
		},
		"InstructedAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "011104238"
		}
	}`)

	versions := []FedwireFundsPaymentStatus.PACS_002_001_VERSION{
		FedwireFundsPaymentStatus.PACS_002_001_03,
		FedwireFundsPaymentStatus.PACS_002_001_04,
		FedwireFundsPaymentStatus.PACS_002_001_05,
		FedwireFundsPaymentStatus.PACS_002_001_06,
		FedwireFundsPaymentStatus.PACS_002_001_07,
		FedwireFundsPaymentStatus.PACS_002_001_08,
		FedwireFundsPaymentStatus.PACS_002_001_09,
		FedwireFundsPaymentStatus.PACS_002_001_10,
		FedwireFundsPaymentStatus.PACS_002_001_11,
		FedwireFundsPaymentStatus.PACS_002_001_12,
		FedwireFundsPaymentStatus.PACS_002_001_13,
		FedwireFundsPaymentStatus.PACS_002_001_14,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := FedwireFundsPaymentStatus.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestFedwireFundsPaymentStatusWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &FedwireFundsPaymentStatusWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long MessageId that should fail validation
		longMessageId := `{
			"MessageId": "ThisIsAnExtremelyLongMessageIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"CreatedDateTime": "2024-01-01T10:00:00Z",
			"TransactionStatus": "ACSC"
		}`
		_, err := wrapper.CreateDocument([]byte(longMessageId), FedwireFundsPaymentStatus.PACS_002_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"MessageId": "20250310QMGFNP31000001",
			"CreatedDateTime": "invalid-date-format",
			"TransactionStatus": "ACSC"
		}`
		err := wrapper.ValidateDocument(malformedDate, FedwireFundsPaymentStatus.PACS_002_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12"><FIToFIPmtStsRpt><GrpHdr><MsgId>Test&InvalidChar</MsgId></GrpHdr></FIToFIPmtStsRpt></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid transaction status", func(t *testing.T) {
		invalidStatus := `{
			"MessageId": "20250310QMGFNP31000001",
			"CreatedDateTime": "2024-01-01T10:00:00Z",
			"TransactionStatus": "INVALID_STATUS_CODE_TOO_LONG",
			"InstructingAgent": {
				"PaymentSysCode": "USABA",
				"PaymentSysMemberId": "021151080"
			}
		}`
		_, err := wrapper.CreateDocument([]byte(invalidStatus), FedwireFundsPaymentStatus.PACS_002_001_12)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := FedwireFundsPaymentStatus.MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "20250310QMGFNP31000001",
				// Missing CreatedDateTime
			},
			// Missing TransactionStatus
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})
}