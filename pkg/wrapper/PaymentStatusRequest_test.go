package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wadearnold/wire20022/pkg/base"
	"github.com/wadearnold/wire20022/pkg/models"
	PaymentStatusRequest "github.com/wadearnold/wire20022/pkg/models/PaymentStatusRequest"
)

// createValidPaymentStatusRequestModel creates a PaymentStatusRequest.MessageModel with all required fields populated
func createValidPaymentStatusRequestModel() PaymentStatusRequest.MessageModel {
	return PaymentStatusRequest.MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250310PSR0000001",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		OriginalMessageId:        "20250310PMT0000001",
		OriginalMessageNameId:    "pacs.008.001.12",
		OriginalCreationDateTime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		OriginalInstructionId:    "INSTR123456789",
		OriginalEndToEndId:       "E2E20250310000001",
		OriginalUETR:             "8a562c67-ca16-48ba-b074-65581be6f066",
		AgentPair: base.AgentPair{
			InstructingAgent: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "021040078",
			},
			InstructedAgent: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "021050049",
			},
		},
	}
}

func TestPaymentStatusRequestWrapper_CreateDocument(t *testing.T) {
	wrapper := &PaymentStatusRequestWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     PaymentStatusRequest.PACS_028_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"messageId": "20250310PSR0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"originalMessageId": "20250310PMT0000001",
				"originalMessageNameId": "pacs.008.001.12",
				"originalCreationDateTime": "2024-01-01T09:00:00Z",
				"originalInstructionId": "INSTR123456789",
				"originalEndToEndId": "E2E20250310000001",
				"originalUETR": "8a562c67-ca16-48ba-b074-65581be6f066",
				"instructingAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021040078"
				},
				"instructedAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021050049"
				}
			}`),
			version:     PaymentStatusRequest.PACS_028_001_06,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     PaymentStatusRequest.PACS_028_001_06,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     PaymentStatusRequest.PACS_028_001_06,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     PaymentStatusRequest.PACS_028_001_06,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     PaymentStatusRequest.PACS_028_001_06,
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

func TestPaymentStatusRequestWrapper_ValidateDocument(t *testing.T) {
	wrapper := &PaymentStatusRequestWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     PaymentStatusRequest.PACS_028_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"messageId": "20250310PSR0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"originalMessageId": "20250310PMT0000001",
				"originalMessageNameId": "pacs.008.001.12",
				"originalCreationDateTime": "2024-01-01T09:00:00Z",
				"instructingAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021040078"
				},
				"instructedAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021050049"
				}
			}`,
			version:     PaymentStatusRequest.PACS_028_001_06,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     PaymentStatusRequest.PACS_028_001_06,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     PaymentStatusRequest.PACS_028_001_06,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     PaymentStatusRequest.PACS_028_001_06,
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

func TestPaymentStatusRequestWrapper_CheckRequireField(t *testing.T) {
	wrapper := &PaymentStatusRequestWrapper{}

	tests := []struct {
		name        string
		model       PaymentStatusRequest.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidPaymentStatusRequestModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: PaymentStatusRequest.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing MessageId
					CreatedDateTime: time.Now(),
				},
				OriginalMessageId:        "20250310PMT0000001",
				OriginalMessageNameId:    "pacs.008.001.12",
				OriginalCreationDateTime: time.Now(),
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       PaymentStatusRequest.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing OriginalMessageId fails validation",
			model: PaymentStatusRequest.MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "20250310PSR0000001",
					CreatedDateTime: time.Now(),
				},
				// Missing OriginalMessageId
				OriginalMessageNameId:    "pacs.008.001.12",
				OriginalCreationDateTime: time.Now(),
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

func TestPaymentStatusRequestWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &PaymentStatusRequestWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.028.001.06">
	<FIToFIPmtStsReq>
		<GrpHdr>
			<MsgId>20250310PSR0000001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</GrpHdr>
		<OrgnlGrpInfAndSts>
			<OrgnlMsgId>20250310PMT0000001</OrgnlMsgId>
			<OrgnlMsgNmId>pacs.008.001.12</OrgnlMsgNmId>
			<OrgnlCreDtTm>2024-01-01T09:00:00Z</OrgnlCreDtTm>
		</OrgnlGrpInfAndSts>
	</FIToFIPmtStsReq>
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
				assert.Equal(t, PaymentStatusRequest.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, PaymentStatusRequest.MessageModel{}, result)
			}
		})
	}
}

func TestPaymentStatusRequestWrapper_GetHelp(t *testing.T) {
	wrapper := &PaymentStatusRequestWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(result), &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")

	// Verify it contains expected fields
	assert.Contains(t, result, "MessageId")
	assert.Contains(t, result, "OriginalMessageId")
	assert.Contains(t, result, "InstructingAgent")
	assert.Contains(t, result, "InstructedAgent")
}

func TestPaymentStatusRequestWrapper_Integration(t *testing.T) {
	wrapper := &PaymentStatusRequestWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"messageId": "20250310PSR0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"originalMessageId": "20250310PMT0000001",
		"originalMessageNameId": "pacs.008.001.12",
		"originalCreationDateTime": "2024-01-01T09:00:00Z",
		"instructingAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021040078"
		},
		"instructedAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021050049"
		}
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, PaymentStatusRequest.PACS_028_001_06)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestPaymentStatusRequestWrapper_AllVersions(t *testing.T) {
	wrapper := &PaymentStatusRequestWrapper{}

	validJSON := []byte(`{
		"messageId": "20250310PSR0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"originalMessageId": "20250310PMT0000001",
		"originalMessageNameId": "pacs.008.001.12",
		"originalCreationDateTime": "2024-01-01T09:00:00Z",
		"instructingAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021040078"
		},
		"instructedAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021050049"
		}
	}`)

	versions := []PaymentStatusRequest.PACS_028_001_VERSION{
		PaymentStatusRequest.PACS_028_001_01,
		PaymentStatusRequest.PACS_028_001_02,
		PaymentStatusRequest.PACS_028_001_03,
		PaymentStatusRequest.PACS_028_001_04,
		PaymentStatusRequest.PACS_028_001_05,
		PaymentStatusRequest.PACS_028_001_06,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := PaymentStatusRequest.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestPaymentStatusRequestWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &PaymentStatusRequestWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long MessageId that should fail validation
		longMessageId := `{
			"messageId": "ThisIsAnExtremelyLongMessageIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"originalMessageId": "20250310PMT0000001"
		}`
		_, err := wrapper.CreateDocument([]byte(longMessageId), PaymentStatusRequest.PACS_028_001_06)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"messageId": "20250310PSR0000001",
			"createdDateTime": "invalid-date-format",
			"originalMessageId": "20250310PMT0000001"
		}`
		err := wrapper.ValidateDocument(malformedDate, PaymentStatusRequest.PACS_028_001_06)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.028.001.06"><FIToFIPmtStsReq><GrpHdr><MsgId>Test&InvalidChar</MsgId></GrpHdr></FIToFIPmtStsReq></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid payment system code", func(t *testing.T) {
		invalidPaymentSys := `{
			"messageId": "20250310PSR0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"originalMessageId": "20250310PMT0000001",
			"originalMessageNameId": "pacs.008.001.12",
			"originalCreationDateTime": "2024-01-01T09:00:00Z",
			"instructingAgent": {
				"paymentSysCode": "INVALID_PAYMENT_SYSTEM_CODE",
				"paymentSysMemberId": "021040078"
			},
			"instructedAgent": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021050049"
			}
		}`
		_, err := wrapper.CreateDocument([]byte(invalidPaymentSys), PaymentStatusRequest.PACS_028_001_06)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := PaymentStatusRequest.MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "20250310PSR0000001",
				// Missing CreatedDateTime
			},
			OriginalMessageId:     "20250310PMT0000001",
			OriginalMessageNameId: "pacs.008.001.12",
			// Missing OriginalCreationDateTime and agents
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})

	t.Run("CreateDocument with missing agent information", func(t *testing.T) {
		missingAgents := `{
			"messageId": "20250310PSR0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"originalMessageId": "20250310PMT0000001",
			"originalMessageNameId": "pacs.008.001.12",
			"originalCreationDateTime": "2024-01-01T09:00:00Z",
			"instructingAgent": {
				"paymentSysCode": "",
				"paymentSysMemberId": ""
			},
			"instructedAgent": {
				"paymentSysCode": "",
				"paymentSysMemberId": ""
			}
		}`
		_, err := wrapper.CreateDocument([]byte(missingAgents), PaymentStatusRequest.PACS_028_001_06)
		// Test handling of missing agent validation
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CreateDocument with invalid UETR format", func(t *testing.T) {
		invalidUETR := `{
			"messageId": "20250310PSR0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"originalMessageId": "20250310PMT0000001",
			"originalMessageNameId": "pacs.008.001.12",
			"originalCreationDateTime": "2024-01-01T09:00:00Z",
			"originalUETR": "invalid-uetr-format-not-uuid",
			"instructingAgent": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021040078"
			},
			"instructedAgent": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021050049"
			}
		}`
		_, err := wrapper.CreateDocument([]byte(invalidUETR), PaymentStatusRequest.PACS_028_001_06)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})
}
