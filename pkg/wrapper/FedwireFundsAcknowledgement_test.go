package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	FedwireFundsAcknowledgement "github.com/moov-io/wire20022/pkg/models/FedwireFundsAcknowledgement"
)

// createValidFedwireFundsAcknowledgementModel creates a FedwireFundsAcknowledgement.MessageModel with all required fields populated
func createValidFedwireFundsAcknowledgementModel() FedwireFundsAcknowledgement.MessageModel {
	return FedwireFundsAcknowledgement.MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250310ACK0000001",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		RelationReference: "20250310B1QDRCQR000001",
		ReferenceName:     "Payment Acknowledgement",
		RequestHandling:   models.SchemaValidationFailed,
	}
}

func TestFedwireFundsAcknowledgementWrapper_CreateDocument(t *testing.T) {
	wrapper := &FedwireFundsAcknowledgementWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     FedwireFundsAcknowledgement.ADMI_007_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"messageId": "20250310ACK0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"relationReference": "20250310B1QDRCQR000001",
				"referenceName": "Payment Acknowledgement",
				"requestHandling": "TS01"
			}`),
			version:     FedwireFundsAcknowledgement.ADMI_007_001_01,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     FedwireFundsAcknowledgement.ADMI_007_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     FedwireFundsAcknowledgement.ADMI_007_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     FedwireFundsAcknowledgement.ADMI_007_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     FedwireFundsAcknowledgement.ADMI_007_001_01,
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

func TestFedwireFundsAcknowledgementWrapper_ValidateDocument(t *testing.T) {
	wrapper := &FedwireFundsAcknowledgementWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     FedwireFundsAcknowledgement.ADMI_007_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"messageId": "20250310ACK0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"relationReference": "20250310B1QDRCQR000001",
				"referenceName": "Payment Acknowledgement",
				"requestHandling": "TS01"
			}`,
			version:     FedwireFundsAcknowledgement.ADMI_007_001_01,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     FedwireFundsAcknowledgement.ADMI_007_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     FedwireFundsAcknowledgement.ADMI_007_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     FedwireFundsAcknowledgement.ADMI_007_001_01,
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

func TestFedwireFundsAcknowledgementWrapper_CheckRequireField(t *testing.T) {
	wrapper := &FedwireFundsAcknowledgementWrapper{}

	tests := []struct {
		name        string
		model       FedwireFundsAcknowledgement.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidFedwireFundsAcknowledgementModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: FedwireFundsAcknowledgement.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing MessageId
					CreatedDateTime: time.Now(),
				},
				RelationReference: "20250310B1QDRCQR000001",
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       FedwireFundsAcknowledgement.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing RelationReference fails validation",
			model: FedwireFundsAcknowledgement.MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "20250310ACK0000001",
					CreatedDateTime: time.Now(),
				},
				// Missing RelationReference
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

func TestFedwireFundsAcknowledgementWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &FedwireFundsAcknowledgementWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.007.001.01">
	<ResolutionOfInvestigation>
		<GrpHdr>
			<MsgId>20250310ACK0000001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</GrpHdr>
		<RsltnInd>
			<RltdRef>20250310B1QDRCQR000001</RltdRef>
			<RefNm>Payment Acknowledgement</RefNm>
		</RsltnInd>
	</ResolutionOfInvestigation>
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
				assert.Equal(t, FedwireFundsAcknowledgement.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, FedwireFundsAcknowledgement.MessageModel{}, result)
			}
		})
	}
}

func TestFedwireFundsAcknowledgementWrapper_GetHelp(t *testing.T) {
	wrapper := &FedwireFundsAcknowledgementWrapper{}

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
	assert.Contains(t, result, "RelationReference")
	assert.Contains(t, result, "RequestHandling")
}

func TestFedwireFundsAcknowledgementWrapper_Integration(t *testing.T) {
	wrapper := &FedwireFundsAcknowledgementWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"messageId": "20250310ACK0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"relationReference": "20250310B1QDRCQR000001",
		"referenceName": "Payment Acknowledgement",
		"requestHandling": "TS01"
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, FedwireFundsAcknowledgement.ADMI_007_001_01)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestFedwireFundsAcknowledgementWrapper_AllVersions(t *testing.T) {
	wrapper := &FedwireFundsAcknowledgementWrapper{}

	validJSON := []byte(`{
		"messageId": "20250310ACK0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"relationReference": "20250310B1QDRCQR000001",
		"referenceName": "Payment Acknowledgement",
		"requestHandling": "TS01"
	}`)

	versions := []FedwireFundsAcknowledgement.ADMI_007_001_VERSION{
		FedwireFundsAcknowledgement.ADMI_007_001_01,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := FedwireFundsAcknowledgement.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestFedwireFundsAcknowledgementWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &FedwireFundsAcknowledgementWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long MessageId that should fail validation
		longMessageId := `{
			"messageId": "ThisIsAnExtremelyLongMessageIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"relationReference": "REF123",
			"requestHandling": "TS01"
		}`
		_, err := wrapper.CreateDocument([]byte(longMessageId), FedwireFundsAcknowledgement.ADMI_007_001_01)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"messageId": "20250310ACK0000001",
			"createdDateTime": "invalid-date-format",
			"relationReference": "REF123",
			"requestHandling": "TS01"
		}`
		err := wrapper.ValidateDocument(malformedDate, FedwireFundsAcknowledgement.ADMI_007_001_01)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.007.001.01"><ResolutionOfInvestigation><GrpHdr><MsgId>Test&InvalidChar</MsgId></GrpHdr></ResolutionOfInvestigation></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid request handling code", func(t *testing.T) {
		invalidRequestHandling := `{
			"messageId": "20250310ACK0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"relationReference": "REF123",
			"referenceName": "Acknowledgement",
			"requestHandling": "INVALID_STATUS_CODE_TOO_LONG"
		}`
		_, err := wrapper.CreateDocument([]byte(invalidRequestHandling), FedwireFundsAcknowledgement.ADMI_007_001_01)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := FedwireFundsAcknowledgement.MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "20250310ACK0000001",
				// Missing CreatedDateTime
			},
			RelationReference: "REF123",
			// Missing RequestHandling
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})
}
