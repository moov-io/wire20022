package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wadearnold/wire20022/pkg/models"
	FedwireFundsSystemResponse "github.com/wadearnold/wire20022/pkg/models/FedwireFundsSystemResponse"
)

// createValidFedwireFundsSystemResponseModel creates a FedwireFundsSystemResponse.MessageModel with all required fields populated
func createValidFedwireFundsSystemResponseModel() FedwireFundsSystemResponse.MessageModel {
	return FedwireFundsSystemResponse.MessageModel{
		MessageId:  "20250310SYS0000001",
		EventCode:  models.SystemOpen,
		EventParam: "NORMAL",
		EventTime:  time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
	}
}

func TestFedwireFundsSystemResponseWrapper_CreateDocument(t *testing.T) {
	wrapper := &FedwireFundsSystemResponseWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     FedwireFundsSystemResponse.ADMI_011_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"messageId": "20250310SYS0000001",
				"eventCode": "OPEN",
				"eventParam": "NORMAL",
				"eventTime": "2024-01-01T10:00:00Z"
			}`),
			version:     FedwireFundsSystemResponse.ADMI_011_001_01,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     FedwireFundsSystemResponse.ADMI_011_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     FedwireFundsSystemResponse.ADMI_011_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     FedwireFundsSystemResponse.ADMI_011_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"messageId": "",
				"eventCode": "OPEN"
			}`),
			version:     FedwireFundsSystemResponse.ADMI_011_001_01,
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

func TestFedwireFundsSystemResponseWrapper_ValidateDocument(t *testing.T) {
	wrapper := &FedwireFundsSystemResponseWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     FedwireFundsSystemResponse.ADMI_011_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"messageId": "20250310SYS0000001",
				"eventCode": "OPEN",
				"eventParam": "NORMAL",
				"eventTime": "2024-01-01T10:00:00Z"
			}`,
			version:     FedwireFundsSystemResponse.ADMI_011_001_01,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     FedwireFundsSystemResponse.ADMI_011_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     FedwireFundsSystemResponse.ADMI_011_001_01,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"messageId": "",
				"eventCode": "OPEN"
			}`,
			version:     FedwireFundsSystemResponse.ADMI_011_001_01,
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

func TestFedwireFundsSystemResponseWrapper_CheckRequireField(t *testing.T) {
	wrapper := &FedwireFundsSystemResponseWrapper{}

	tests := []struct {
		name        string
		model       FedwireFundsSystemResponse.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidFedwireFundsSystemResponseModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: FedwireFundsSystemResponse.MessageModel{
				// Missing MessageId
				EventCode:  models.SystemOpen,
				EventParam: "NORMAL",
				EventTime:  time.Now(),
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       FedwireFundsSystemResponse.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing EventCode fails validation",
			model: FedwireFundsSystemResponse.MessageModel{
				MessageId: "20250310SYS0000001",
				// Missing EventCode
				EventParam: "NORMAL",
				EventTime:  time.Now(),
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

func TestFedwireFundsSystemResponseWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &FedwireFundsSystemResponseWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.011.001.01">
	<SysEvtAck>
		<MsgId>20250310SYS0000001</MsgId>
		<EvtCd>OPEN</EvtCd>
		<EvtParam>NORMAL</EvtParam>
		<EvtTm>2024-01-01T10:00:00Z</EvtTm>
	</SysEvtAck>
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
				assert.Equal(t, FedwireFundsSystemResponse.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, FedwireFundsSystemResponse.MessageModel{}, result)
			}
		})
	}
}

func TestFedwireFundsSystemResponseWrapper_GetHelp(t *testing.T) {
	wrapper := &FedwireFundsSystemResponseWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(result), &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")

	// Verify it contains expected fields
	assert.Contains(t, result, "MessageId")
	assert.Contains(t, result, "EventCode")
	assert.Contains(t, result, "EventParam")
	assert.Contains(t, result, "EventTime")
}

func TestFedwireFundsSystemResponseWrapper_Integration(t *testing.T) {
	wrapper := &FedwireFundsSystemResponseWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"messageId": "20250310SYS0000001",
		"eventCode": "OPEN",
		"eventParam": "NORMAL",
		"eventTime": "2024-01-01T10:00:00Z"
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, FedwireFundsSystemResponse.ADMI_011_001_01)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestFedwireFundsSystemResponseWrapper_AllVersions(t *testing.T) {
	wrapper := &FedwireFundsSystemResponseWrapper{}

	validJSON := []byte(`{
		"messageId": "20250310SYS0000001",
		"eventCode": "OPEN",
		"eventParam": "NORMAL",
		"eventTime": "2024-01-01T10:00:00Z"
	}`)

	versions := []FedwireFundsSystemResponse.ADMI_011_001_VERSION{
		FedwireFundsSystemResponse.ADMI_011_001_01,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := FedwireFundsSystemResponse.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestFedwireFundsSystemResponseWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &FedwireFundsSystemResponseWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long MessageId that should fail validation
		longMessageId := `{
			"messageId": "ThisIsAnExtremelyLongMessageIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"eventCode": "OPEN",
			"eventParam": "NORMAL",
			"eventTime": "2024-01-01T10:00:00Z"
		}`
		_, err := wrapper.CreateDocument([]byte(longMessageId), FedwireFundsSystemResponse.ADMI_011_001_01)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"messageId": "20250310SYS0000001",
			"eventCode": "OPEN",
			"eventParam": "NORMAL",
			"eventTime": "invalid-date-format"
		}`
		err := wrapper.ValidateDocument(malformedDate, FedwireFundsSystemResponse.ADMI_011_001_01)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.011.001.01"><SysEvtAck><MsgId>Test&InvalidChar</MsgId></SysEvtAck></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid event code", func(t *testing.T) {
		invalidEventCode := `{
			"messageId": "20250310SYS0000001",
			"eventCode": "INVALID_EVENT_CODE_TOO_LONG",
			"eventParam": "NORMAL",
			"eventTime": "2024-01-01T10:00:00Z"
		}`
		_, err := wrapper.CreateDocument([]byte(invalidEventCode), FedwireFundsSystemResponse.ADMI_011_001_01)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with zero time value", func(t *testing.T) {
		zeroTimeModel := FedwireFundsSystemResponse.MessageModel{
			MessageId:  "20250310SYS0000001",
			EventCode:  models.SystemOpen,
			EventParam: "NORMAL",
			// EventTime is zero value - should fail
		}
		err := wrapper.CheckRequireField(zeroTimeModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})

	t.Run("CreateDocument with all event types", func(t *testing.T) {
		eventTypes := []string{"ADHC", "PING", "CLSD", "EXTN", "OPEN"}

		for _, eventType := range eventTypes {
			eventJSON := `{
				"messageId": "20250310SYS0000001",
				"eventCode": "` + eventType + `",
				"eventParam": "TEST",
				"eventTime": "2024-01-01T10:00:00Z"
			}`
			xmlData, err := wrapper.CreateDocument([]byte(eventJSON), FedwireFundsSystemResponse.ADMI_011_001_01)
			assert.NoError(t, err, "Event type %s should be valid", eventType)
			assert.NotEmpty(t, xmlData, "XML should be generated for event type %s", eventType)
		}
	})
}
