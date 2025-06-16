package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ConnectionCheck "github.com/wadearnold/wire20022/pkg/models/ConnectionCheck"
)

// createValidConnectionCheckModel creates a ConnectionCheck.MessageModel with all required fields populated
func createValidConnectionCheckModel() ConnectionCheck.MessageModel {
	return ConnectionCheck.MessageModel{
		EventType:  "CONN",
		EventParam: "CHK",
		EventTime:  time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
	}
}

func TestConnectionCheckWrapper_CreateDocument(t *testing.T) {
	wrapper := &ConnectionCheckWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     ConnectionCheck.ADMI_004_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"eventType": "CONN",
				"eventParam": "CHK",
				"eventTime": "2024-01-01T10:00:00Z"
			}`),
			version:     ConnectionCheck.ADMI_004_001_02,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     ConnectionCheck.ADMI_004_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     ConnectionCheck.ADMI_004_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     ConnectionCheck.ADMI_004_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"eventType": "",
				"eventParam": "CHECK"
			}`),
			version:     ConnectionCheck.ADMI_004_001_02,
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

func TestConnectionCheckWrapper_ValidateDocument(t *testing.T) {
	wrapper := &ConnectionCheckWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     ConnectionCheck.ADMI_004_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"eventType": "CONN",
				"eventParam": "CHK",
				"eventTime": "2024-01-01T10:00:00Z"
			}`,
			version:     ConnectionCheck.ADMI_004_001_02,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     ConnectionCheck.ADMI_004_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     ConnectionCheck.ADMI_004_001_02,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"eventType": "",
				"eventParam": "CHECK"
			}`,
			version:     ConnectionCheck.ADMI_004_001_02,
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

func TestConnectionCheckWrapper_CheckRequireField(t *testing.T) {
	wrapper := &ConnectionCheckWrapper{}

	tests := []struct {
		name        string
		model       ConnectionCheck.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidConnectionCheckModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: ConnectionCheck.MessageModel{
				// Missing EventType
				EventParam: "CHK",
				EventTime:  time.Now(),
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       ConnectionCheck.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing EventParam fails validation",
			model: ConnectionCheck.MessageModel{
				EventType: "CONN",
				// Missing EventParam
				EventTime: time.Now(),
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

func TestConnectionCheckWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &ConnectionCheckWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.004.001.02">
	<SysEvtNtfctn>
		<EvtTp>CONN</EvtTp>
		<EvtParam>CHK</EvtParam>
		<EvtTm>2024-01-01T10:00:00Z</EvtTm>
	</SysEvtNtfctn>
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
				assert.Equal(t, ConnectionCheck.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, ConnectionCheck.MessageModel{}, result)
			}
		})
	}
}

func TestConnectionCheckWrapper_GetHelp(t *testing.T) {
	wrapper := &ConnectionCheckWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(result), &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")

	// Verify it contains expected fields
	assert.Contains(t, result, "EventType")
	assert.Contains(t, result, "EventParam")
	assert.Contains(t, result, "EventTime")
}

func TestConnectionCheckWrapper_Integration(t *testing.T) {
	wrapper := &ConnectionCheckWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"eventType": "CONN",
		"eventParam": "CHK",
		"eventTime": "2024-01-01T10:00:00Z"
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, ConnectionCheck.ADMI_004_001_02)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestConnectionCheckWrapper_AllVersions(t *testing.T) {
	wrapper := &ConnectionCheckWrapper{}

	validJSON := []byte(`{
		"eventType": "CONN",
		"eventParam": "CHK",
		"eventTime": "2024-01-01T10:00:00Z"
	}`)

	versions := []ConnectionCheck.ADMI_004_001_VERSION{
		ConnectionCheck.ADMI_004_001_01,
		ConnectionCheck.ADMI_004_001_02,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := ConnectionCheck.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestConnectionCheckWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &ConnectionCheckWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long EventType that should fail validation
		longEventType := `{
			"eventType": "ThisIsAnExtremelyLongEventTypeThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"eventParam": "CHECK",
			"eventTime": "2024-01-01T10:00:00Z"
		}`
		_, err := wrapper.CreateDocument([]byte(longEventType), ConnectionCheck.ADMI_004_001_02)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"eventType": "CONN",
			"eventParam": "CHK",
			"eventTime": "invalid-date-format"
		}`
		err := wrapper.ValidateDocument(malformedDate, ConnectionCheck.ADMI_004_001_02)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.004.001.02"><SysEvtNtfctn><EvtTp>Test&InvalidChar</EvtTp></SysEvtNtfctn></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CheckRequireField with zero time value", func(t *testing.T) {
		zeroTimeModel := ConnectionCheck.MessageModel{
			EventType:  "CONN",
			EventParam: "CHK",
			// EventTime is zero value - should fail
		}
		err := wrapper.CheckRequireField(zeroTimeModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})
}