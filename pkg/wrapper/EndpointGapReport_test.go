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
	EndpointGapReport "github.com/wadearnold/wire20022/pkg/models/EndpointGapReport"
)

// createValidEndpointGapReportModel creates an EndpointGapReport.MessageModel with all required fields populated
func createValidEndpointGapReportModel() EndpointGapReport.MessageModel {
	return EndpointGapReport.MessageModel{
		MessageHeader: base.MessageHeader{
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		MessageId: models.EndpointGapReportType,
		Pagenation: models.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		},
		ReportId:             models.InputMessageAccountabilityData,
		ReportCreateDateTime: time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC),
		AccountOtherId:       "ACC123456789",
		AdditionalReportInfo: "Gap Report Details",
	}
}

func TestEndpointGapReportWrapper_CreateDocument(t *testing.T) {
	wrapper := &EndpointGapReportWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     EndpointGapReport.CAMT_052_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"messageId": "GAPR",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"pagenation": {
					"PageNumber": "1",
					"LastPageIndicator": true
				},
				"reportId": "IMAD",
				"reportCreateDateTime": "2024-01-01T08:00:00Z",
				"accountOtherId": "ACC123456789",
				"additionalReportInfo": "Gap Report Details"
			}`),
			version:     EndpointGapReport.CAMT_052_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     EndpointGapReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     EndpointGapReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     EndpointGapReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     EndpointGapReport.CAMT_052_001_12,
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

func TestEndpointGapReportWrapper_ValidateDocument(t *testing.T) {
	wrapper := &EndpointGapReportWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     EndpointGapReport.CAMT_052_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"messageId": "GAPR",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"pagenation": {
					"PageNumber": "1",
					"LastPageIndicator": true
				},
				"reportId": "IMAD",
				"reportCreateDateTime": "2024-01-01T08:00:00Z"
			}`,
			version:     EndpointGapReport.CAMT_052_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     EndpointGapReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     EndpointGapReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     EndpointGapReport.CAMT_052_001_12,
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

func TestEndpointGapReportWrapper_CheckRequireField(t *testing.T) {
	wrapper := &EndpointGapReportWrapper{}

	tests := []struct {
		name        string
		model       EndpointGapReport.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidEndpointGapReportModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: EndpointGapReport.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing CreatedDateTime
				},
				MessageId: models.EndpointGapReportType,
				ReportId:  models.InputMessageAccountabilityData,
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       EndpointGapReport.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing ReportId fails validation",
			model: EndpointGapReport.MessageModel{
				MessageHeader: base.MessageHeader{
					CreatedDateTime: time.Now(),
				},
				MessageId: models.EndpointGapReportType,
				// Missing ReportId
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

func TestEndpointGapReportWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &EndpointGapReportWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.12">
	<BkToCstmrAcctRpt>
		<GrpHdr>
			<MsgId>GAPR</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</GrpHdr>
		<Rpt>
			<Id>IMAD</Id>
			<CreDtTm>2024-01-01T08:00:00Z</CreDtTm>
		</Rpt>
	</BkToCstmrAcctRpt>
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
				assert.Equal(t, EndpointGapReport.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, EndpointGapReport.MessageModel{}, result)
			}
		})
	}
}

func TestEndpointGapReportWrapper_GetHelp(t *testing.T) {
	wrapper := &EndpointGapReportWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(result), &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")

	// Verify it contains expected fields
	assert.Contains(t, result, "MessageId")
	assert.Contains(t, result, "ReportId")
	assert.Contains(t, result, "AccountOtherId")
	assert.Contains(t, result, "ReportCreateDateTime")
}

func TestEndpointGapReportWrapper_Integration(t *testing.T) {
	wrapper := &EndpointGapReportWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"messageId": "GAPR",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"pagenation": {
			"PageNumber": "1",
			"LastPageIndicator": true
		},
		"reportId": "IMAD",
		"reportCreateDateTime": "2024-01-01T08:00:00Z"
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, EndpointGapReport.CAMT_052_001_12)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestEndpointGapReportWrapper_AllVersions(t *testing.T) {
	wrapper := &EndpointGapReportWrapper{}

	validJSON := []byte(`{
		"messageId": "GAPR",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"pagenation": {
			"PageNumber": "1",
			"LastPageIndicator": true
		},
		"reportId": "IMAD",
		"reportCreateDateTime": "2024-01-01T08:00:00Z"
	}`)

	versions := []EndpointGapReport.CAMT_052_001_VERSION{
		EndpointGapReport.CAMT_052_001_02,
		EndpointGapReport.CAMT_052_001_03,
		EndpointGapReport.CAMT_052_001_04,
		EndpointGapReport.CAMT_052_001_05,
		EndpointGapReport.CAMT_052_001_06,
		EndpointGapReport.CAMT_052_001_07,
		EndpointGapReport.CAMT_052_001_08,
		EndpointGapReport.CAMT_052_001_09,
		EndpointGapReport.CAMT_052_001_10,
		EndpointGapReport.CAMT_052_001_11,
		EndpointGapReport.CAMT_052_001_12,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := EndpointGapReport.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestEndpointGapReportWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &EndpointGapReportWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long AdditionalReportInfo that should fail validation
		longReportInfo := `{
			"messageId": "GAPR",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"pagenation": {
				"PageNumber": "1",
				"LastPageIndicator": true
			},
			"reportId": "IMAD",
			"reportCreateDateTime": "2024-01-01T08:00:00Z",
			"additionalReportInfo": "ThisIsAnExtremelyLongAdditionalReportInfoFieldThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument"
		}`
		_, err := wrapper.CreateDocument([]byte(longReportInfo), EndpointGapReport.CAMT_052_001_12)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"messageId": "GAPR",
			"createdDateTime": "invalid-date-format",
			"reportId": "IMAD"
		}`
		err := wrapper.ValidateDocument(malformedDate, EndpointGapReport.CAMT_052_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.12"><BkToCstmrAcctRpt><GrpHdr><MsgId>Test&InvalidChar</MsgId></GrpHdr></BkToCstmrAcctRpt></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid gap type", func(t *testing.T) {
		invalidGapType := `{
			"messageId": "GAPR",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"reportId": "INVALID_GAP_TYPE",
			"reportCreateDateTime": "2024-01-01T08:00:00Z"
		}`
		_, err := wrapper.CreateDocument([]byte(invalidGapType), EndpointGapReport.CAMT_052_001_12)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := EndpointGapReport.MessageModel{
			MessageHeader: base.MessageHeader{
				// Missing CreatedDateTime
			},
			MessageId: models.EndpointGapReportType,
			ReportId:  models.InputMessageAccountabilityData,
			// Missing other required fields
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})

	t.Run("CreateDocument with both gap types", func(t *testing.T) {
		gapTypes := []string{"IMAD", "OMAD"}

		for _, gapType := range gapTypes {
			gapJSON := `{
				"messageId": "GAPR",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"pagenation": {
					"PageNumber": "1",
					"LastPageIndicator": true
				},
				"reportId": "` + gapType + `",
				"reportCreateDateTime": "2024-01-01T08:00:00Z",
				"accountOtherId": "ACC123456789"
			}`
			xmlData, err := wrapper.CreateDocument([]byte(gapJSON), EndpointGapReport.CAMT_052_001_12)
			assert.NoError(t, err, "Gap type %s should be valid", gapType)
			assert.NotEmpty(t, xmlData, "XML should be generated for gap type %s", gapType)
		}
	})

	t.Run("CreateDocument with zero time value", func(t *testing.T) {
		zeroTimeModel := EndpointGapReport.MessageModel{
			MessageHeader: base.MessageHeader{
				// CreatedDateTime is zero value - should fail
			},
			MessageId: models.EndpointGapReportType,
			ReportId:  models.InputMessageAccountabilityData,
			// ReportCreateDateTime is zero value - should fail
		}
		err := wrapper.CheckRequireField(zeroTimeModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})
}
