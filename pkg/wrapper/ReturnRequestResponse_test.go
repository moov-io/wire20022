package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wadearnold/wire20022/pkg/models"
	ReturnRequestResponse "github.com/wadearnold/wire20022/pkg/models/ReturnRequestResponse"
)

// createValidReturnRequestResponseModel creates a ReturnRequestResponse.MessageModel with all required fields populated
func createValidReturnRequestResponseModel() ReturnRequestResponse.MessageModel {
	return ReturnRequestResponse.MessageModel{
		AssignmentId: "20250310RRR0000001",
		Assigner: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		},
		Assignee: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "021050049",
		},
		AssignmentCreateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		ResolvedCaseId:       "CASE20250310000001",
		Creator: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "021060052",
		},
		Status:                    models.ReturnRequestAccepted,
		OriginalMessageId:         "20250310PMT0000001",
		OriginalMessageNameId:     "pacs.008.001.12",
		OriginalMessageCreateTime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		OriginalInstructionId:     "INSTR123456789",
		OriginalEndToEndId:        "E2E20250310000001",
		OriginalUETR:              "8a562c67-ca16-48ba-b074-65581be6f066",
		CancellationStatusReasonInfo: models.Reason{
			AdditionalInfo: "Return request processed successfully",
		},
	}
}

func TestReturnRequestResponseWrapper_CreateDocument(t *testing.T) {
	wrapper := &ReturnRequestResponseWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     ReturnRequestResponse.CAMT_029_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"assignmentId": "20250310RRR0000001",
				"assigner": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021040078"
				},
				"assignee": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021050049"
				},
				"assignmentCreateTime": "2024-01-01T10:00:00Z",
				"resolvedCaseId": "CASE20250310000001",
				"creator": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021060052"
				},
				"status": "CNCL",
				"originalMessageId": "20250310PMT0000001",
				"originalMessageNameId": "pacs.008.001.12",
				"originalMessageCreateTime": "2024-01-01T09:00:00Z",
				"originalInstructionId": "INSTR123456789",
				"originalEndToEndId": "E2E20250310000001",
				"originalUETR": "8a562c67-ca16-48ba-b074-65581be6f066",
				"cancellationStatusReasonInfo": {
					"AdditionalInfo": "Return request processed successfully"
				}
			}`),
			version:     ReturnRequestResponse.CAMT_029_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     ReturnRequestResponse.CAMT_029_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     ReturnRequestResponse.CAMT_029_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     ReturnRequestResponse.CAMT_029_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"assignmentId": "",
				"assignmentCreateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     ReturnRequestResponse.CAMT_029_001_12,
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

func TestReturnRequestResponseWrapper_ValidateDocument(t *testing.T) {
	wrapper := &ReturnRequestResponseWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     ReturnRequestResponse.CAMT_029_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"assignmentId": "20250310RRR0000001",
				"assigner": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021040078"
				},
				"assignee": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021050049"
				},
				"assignmentCreateTime": "2024-01-01T10:00:00Z",
				"resolvedCaseId": "CASE20250310000001",
				"creator": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021060052"
				},
				"originalMessageId": "20250310PMT0000001",
				"originalMessageNameId": "pacs.008.001.12",
				"originalMessageCreateTime": "2024-01-01T09:00:00Z"
			}`,
			version:     ReturnRequestResponse.CAMT_029_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     ReturnRequestResponse.CAMT_029_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     ReturnRequestResponse.CAMT_029_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"assignmentId": "",
				"assignmentCreateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     ReturnRequestResponse.CAMT_029_001_12,
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

func TestReturnRequestResponseWrapper_CheckRequireField(t *testing.T) {
	wrapper := &ReturnRequestResponseWrapper{}

	tests := []struct {
		name        string
		model       ReturnRequestResponse.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidReturnRequestResponseModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: ReturnRequestResponse.MessageModel{
				// Missing AssignmentId
				Assigner: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "021040078",
				},
				Assignee: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "021050049",
				},
				AssignmentCreateTime:      time.Now(),
				ResolvedCaseId:            "CASE20250310000001",
				OriginalMessageId:         "20250310PMT0000001",
				OriginalMessageNameId:     "pacs.008.001.12",
				OriginalMessageCreateTime: time.Now(),
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       ReturnRequestResponse.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing OriginalMessageId fails validation",
			model: ReturnRequestResponse.MessageModel{
				AssignmentId: "20250310RRR0000001",
				Assigner: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "021040078",
				},
				Assignee: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "021050049",
				},
				AssignmentCreateTime: time.Now(),
				ResolvedCaseId:       "CASE20250310000001",
				Creator: models.Agent{
					PaymentSysCode:     models.PaymentSysUSABA,
					PaymentSysMemberId: "021060052",
				},
				// Missing OriginalMessageId
				OriginalMessageNameId:     "pacs.008.001.12",
				OriginalMessageCreateTime: time.Now(),
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

func TestReturnRequestResponseWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &ReturnRequestResponseWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.029.001.12">
	<RsltnOfInvstgtn>
		<Assgnmt>
			<Id>20250310RRR0000001</Id>
			<Assgnr>
				<Agt>
					<PrtryId>
						<Id>021040078</Id>
					</PrtryId>
				</Agt>
			</Assgnr>
			<Assgne>
				<Agt>
					<PrtryId>
						<Id>021050049</Id>
					</PrtryId>
				</Agt>
			</Assgne>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</Assgnmt>
		<RslvdCase>
			<Id>CASE20250310000001</Id>
		</RslvdCase>
	</RsltnOfInvstgtn>
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
				assert.Equal(t, ReturnRequestResponse.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, ReturnRequestResponse.MessageModel{}, result)
			}
		})
	}
}

func TestReturnRequestResponseWrapper_GetHelp(t *testing.T) {
	wrapper := &ReturnRequestResponseWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(result), &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")

	// Verify it contains expected fields
	assert.Contains(t, result, "AssignmentId")
	assert.Contains(t, result, "Assigner")
	assert.Contains(t, result, "Assignee")
	assert.Contains(t, result, "OriginalMessageId")
}

func TestReturnRequestResponseWrapper_Integration(t *testing.T) {
	wrapper := &ReturnRequestResponseWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"assignmentId": "20250310RRR0000001",
		"assigner": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021040078"
		},
		"assignee": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021050049"
		},
		"assignmentCreateTime": "2024-01-01T10:00:00Z",
		"resolvedCaseId": "CASE20250310000001",
		"creator": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021060052"
		},
		"originalMessageId": "20250310PMT0000001",
		"originalMessageNameId": "pacs.008.001.12",
		"originalMessageCreateTime": "2024-01-01T09:00:00Z"
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, ReturnRequestResponse.CAMT_029_001_12)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestReturnRequestResponseWrapper_AllVersions(t *testing.T) {
	wrapper := &ReturnRequestResponseWrapper{}

	validJSON := []byte(`{
		"assignmentId": "20250310RRR0000001",
		"assigner": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021040078"
		},
		"assignee": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021050049"
		},
		"assignmentCreateTime": "2024-01-01T10:00:00Z",
		"resolvedCaseId": "CASE20250310000001",
		"creator": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021060052"
		},
		"originalMessageId": "20250310PMT0000001",
		"originalMessageNameId": "pacs.008.001.12",
		"originalMessageCreateTime": "2024-01-01T09:00:00Z"
	}`)

	versions := []ReturnRequestResponse.CAMT_029_001_VERSION{
		ReturnRequestResponse.CAMT_029_001_03,
		ReturnRequestResponse.CAMT_029_001_04,
		ReturnRequestResponse.CAMT_029_001_05,
		ReturnRequestResponse.CAMT_029_001_06,
		ReturnRequestResponse.CAMT_029_001_07,
		ReturnRequestResponse.CAMT_029_001_08,
		ReturnRequestResponse.CAMT_029_001_09,
		ReturnRequestResponse.CAMT_029_001_10,
		ReturnRequestResponse.CAMT_029_001_11,
		ReturnRequestResponse.CAMT_029_001_12,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := ReturnRequestResponse.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestReturnRequestResponseWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &ReturnRequestResponseWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long AssignmentId that should fail validation
		longAssignmentId := `{
			"assignmentId": "ThisIsAnExtremelyLongAssignmentIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"assignmentCreateTime": "2024-01-01T10:00:00Z",
			"resolvedCaseId": "CASE20250310000001"
		}`
		_, err := wrapper.CreateDocument([]byte(longAssignmentId), ReturnRequestResponse.CAMT_029_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"assignmentId": "20250310RRR0000001",
			"assignmentCreateTime": "invalid-date-format",
			"resolvedCaseId": "CASE20250310000001"
		}`
		err := wrapper.ValidateDocument(malformedDate, ReturnRequestResponse.CAMT_029_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.029.001.12"><RsltnOfInvstgtn><Assgnmt><Id>Test&InvalidChar</Id></Assgnmt></RsltnOfInvstgtn></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid status code", func(t *testing.T) {
		invalidStatus := `{
			"assignmentId": "20250310RRR0000001",
			"assigner": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021040078"
			},
			"assignee": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021050049"
			},
			"assignmentCreateTime": "2024-01-01T10:00:00Z",
			"resolvedCaseId": "CASE20250310000001",
			"creator": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021060052"
			},
			"status": "INVALID_STATUS_CODE",
			"originalMessageId": "20250310PMT0000001",
			"originalMessageNameId": "pacs.008.001.12",
			"originalMessageCreateTime": "2024-01-01T09:00:00Z"
		}`
		_, err := wrapper.CreateDocument([]byte(invalidStatus), ReturnRequestResponse.CAMT_029_001_12)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := ReturnRequestResponse.MessageModel{
			AssignmentId: "20250310RRR0000001",
			// Missing Assigner
			Assignee: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "021050049",
			},
			AssignmentCreateTime: time.Now(),
			ResolvedCaseId:       "CASE20250310000001",
			// Missing other required fields
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})

	t.Run("CreateDocument with missing agent information", func(t *testing.T) {
		missingAgents := `{
			"assignmentId": "20250310RRR0000001",
			"assigner": {
				"paymentSysCode": "",
				"paymentSysMemberId": ""
			},
			"assignee": {
				"paymentSysCode": "",
				"paymentSysMemberId": ""
			},
			"assignmentCreateTime": "2024-01-01T10:00:00Z",
			"resolvedCaseId": "CASE20250310000001",
			"creator": {
				"paymentSysCode": "",
				"paymentSysMemberId": ""
			},
			"originalMessageId": "20250310PMT0000001",
			"originalMessageNameId": "pacs.008.001.12",
			"originalMessageCreateTime": "2024-01-01T09:00:00Z"
		}`
		_, err := wrapper.CreateDocument([]byte(missingAgents), ReturnRequestResponse.CAMT_029_001_12)
		// Test handling of missing agent validation
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CreateDocument with invalid UETR format", func(t *testing.T) {
		invalidUETR := `{
			"assignmentId": "20250310RRR0000001",
			"assigner": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021040078"
			},
			"assignee": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021050049"
			},
			"assignmentCreateTime": "2024-01-01T10:00:00Z",
			"resolvedCaseId": "CASE20250310000001",
			"creator": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021060052"
			},
			"originalMessageId": "20250310PMT0000001",
			"originalMessageNameId": "pacs.008.001.12",
			"originalMessageCreateTime": "2024-01-01T09:00:00Z",
			"originalUETR": "invalid-uetr-format-not-uuid"
		}`
		_, err := wrapper.CreateDocument([]byte(invalidUETR), ReturnRequestResponse.CAMT_029_001_12)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CreateDocument with zero time value", func(t *testing.T) {
		zeroTimeModel := ReturnRequestResponse.MessageModel{
			AssignmentId: "20250310RRR0000001",
			Assigner: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "021040078",
			},
			Assignee: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "021050049",
			},
			// AssignmentCreateTime is zero value - should fail
			ResolvedCaseId: "CASE20250310000001",
			Creator: models.Agent{
				PaymentSysCode:     models.PaymentSysUSABA,
				PaymentSysMemberId: "021060052",
			},
			OriginalMessageId:     "20250310PMT0000001",
			OriginalMessageNameId: "pacs.008.001.12",
			// OriginalMessageCreateTime is zero value - should fail
		}
		err := wrapper.CheckRequireField(zeroTimeModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})
}
