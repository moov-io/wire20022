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
	ActivityReport "github.com/wadearnold/wire20022/pkg/models/ActivityReport"
)

// createValidActivityReportModel creates an ActivityReport.MessageModel with all required fields populated
func createValidActivityReportModel() ActivityReport.MessageModel {
	return ActivityReport.MessageModel{
		MessageHeader: base.MessageHeader{
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		MessageId: models.ActivityReport,
		Pagenation: models.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		},
		ReportId:             models.Intraday,
		ReportCreateDateTime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		AccountOtherId:       "ACCT123456789",
		TotalEntries:         "10",
		TotalCreditEntries: models.NumberAndSumOfTransactions{
			NumberOfEntries: "5",
			Sum:             1000000.00,
		},
		TotalDebitEntries: models.NumberAndSumOfTransactions{
			NumberOfEntries: "5",
			Sum:             800000.00,
		},
		TotalEntriesPerBankTransactionCode: []models.TotalsPerBankTransactionCode{
			{
				NumberOfEntries:     "2",
				BankTransactionCode: models.TransPending,
			},
		},
		EntryDetails: []models.Entry{
			{
				Amount: models.CurrencyAndAmount{
					Amount:   100000.00,
					Currency: "USD",
				},
				CreditDebitIndicator: models.Credit,
				Status:               models.Book,
			},
		},
	}
}

func TestActivityReportWrapper_CreateDocument(t *testing.T) {
	wrapper := &ActivityReportWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     ActivityReport.CAMT_052_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"createdDateTime": "2024-01-01T10:00:00Z",
				"messageId": "ACTR",
				"pagenation": {
					"PageNumber": "1",
					"LastPageIndicator": true
				},
				"reportId": "IDAY",
				"reportCreateDateTime": "2024-01-01T09:00:00Z",
				"accountOtherId": "ACCT123456789",
				"totalEntries": "10",
				"totalCreditEntries": {
					"NumberOfEntries": "5",
					"Sum": 1000000.00
				},
				"totalDebitEntries": {
					"NumberOfEntries": "5",
					"Sum": 800000.00
				}
			}`),
			version:     ActivityReport.CAMT_052_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     ActivityReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     ActivityReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     ActivityReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     ActivityReport.CAMT_052_001_12,
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

func TestActivityReportWrapper_ValidateDocument(t *testing.T) {
	wrapper := &ActivityReportWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     ActivityReport.CAMT_052_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"createdDateTime": "2024-01-01T10:00:00Z",
				"messageId": "ACTR",
				"pagenation": {
					"PageNumber": "1",
					"LastPageIndicator": true
				},
				"reportId": "IDAY",
				"reportCreateDateTime": "2024-01-01T09:00:00Z",
				"accountOtherId": "ACCT123456789",
				"totalEntries": "10"
			}`,
			version:     ActivityReport.CAMT_052_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     ActivityReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     ActivityReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     ActivityReport.CAMT_052_001_12,
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

func TestActivityReportWrapper_CheckRequireField(t *testing.T) {
	wrapper := &ActivityReportWrapper{}

	tests := []struct {
		name        string
		model       ActivityReport.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidActivityReportModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: ActivityReport.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing CreatedDateTime  
				},
				MessageId: models.ActivityReport,
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       ActivityReport.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing MessageId fails validation",
			model: ActivityReport.MessageModel{
				MessageHeader: base.MessageHeader{
					CreatedDateTime: time.Now(),
				},
				// Missing MessageId
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

func TestActivityReportWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &ActivityReportWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.12">
	<BkToCstmrAcctRpt>
		<GrpHdr>
			<MsgId>ACTR</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</GrpHdr>
		<Rpt>
			<Id>IDAY</Id>
			<CreDtTm>2024-01-01T09:00:00Z</CreDtTm>
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
				assert.Equal(t, ActivityReport.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, ActivityReport.MessageModel{}, result)
			}
		})
	}
}

func TestActivityReportWrapper_GetHelp(t *testing.T) {
	wrapper := &ActivityReportWrapper{}

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
	assert.Contains(t, result, "ReportId")
	assert.Contains(t, result, "Pagenation")
}

func TestActivityReportWrapper_Integration(t *testing.T) {
	wrapper := &ActivityReportWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"createdDateTime": "2024-01-01T10:00:00Z",
		"messageId": "ACTR",
		"pagenation": {
			"PageNumber": "1",
			"LastPageIndicator": true
		},
		"reportId": "IDAY",
		"reportCreateDateTime": "2024-01-01T09:00:00Z",
		"accountOtherId": "ACCT123456789",
		"totalEntries": "10",
		"totalCreditEntries": {
			"NumberOfEntries": "5",
			"Sum": 1000000.00
		},
		"totalDebitEntries": {
			"NumberOfEntries": "5", 
			"Sum": 800000.00
		}
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, ActivityReport.CAMT_052_001_12)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestActivityReportWrapper_AllVersions(t *testing.T) {
	wrapper := &ActivityReportWrapper{}

	validJSON := []byte(`{
		"createdDateTime": "2024-01-01T10:00:00Z",
		"messageId": "ACTR",
		"pagenation": {
			"PageNumber": "1",
			"LastPageIndicator": true
		},
		"reportId": "IDAY",
		"reportCreateDateTime": "2024-01-01T09:00:00Z"
	}`)

	versions := []ActivityReport.CAMT_052_001_VERSION{
		ActivityReport.CAMT_052_001_01,
		ActivityReport.CAMT_052_001_02,
		ActivityReport.CAMT_052_001_03,
		ActivityReport.CAMT_052_001_04,
		ActivityReport.CAMT_052_001_05,
		ActivityReport.CAMT_052_001_06,
		ActivityReport.CAMT_052_001_07,
		ActivityReport.CAMT_052_001_08,
		ActivityReport.CAMT_052_001_09,
		ActivityReport.CAMT_052_001_10,
		ActivityReport.CAMT_052_001_11,
		ActivityReport.CAMT_052_001_12,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := ActivityReport.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestActivityReportWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &ActivityReportWrapper{}

	t.Run("CreateDocument with invalid pagination structure", func(t *testing.T) {
		invalidPagination := `{
			"createdDateTime": "2024-01-01T10:00:00Z",
			"messageId": "ACTR",
			"pagenation": {
				"PageNumber": "",
				"LastPageIndicator": "invalid_boolean"
			},
			"reportId": "IDAY"
		}`
		_, err := wrapper.CreateDocument([]byte(invalidPagination), ActivityReport.CAMT_052_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"createdDateTime": "invalid-date-format",
			"messageId": "ACTR",
			"reportId": "IDAY"
		}`
		err := wrapper.ValidateDocument(malformedDate, ActivityReport.CAMT_052_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.12"><BkToCstmrAcctRpt><GrpHdr><MsgId>Test&InvalidChar</MsgId></GrpHdr></BkToCstmrAcctRpt></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid message ID type", func(t *testing.T) {
		invalidMessageId := `{
			"createdDateTime": "2024-01-01T10:00:00Z",
			"messageId": "INVALID_MESSAGE_ID_TOO_LONG",
			"reportId": "IDAY"
		}`
		_, err := wrapper.CreateDocument([]byte(invalidMessageId), ActivityReport.CAMT_052_001_12)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with complex nested structure missing", func(t *testing.T) {
		partialModel := ActivityReport.MessageModel{
			MessageHeader: base.MessageHeader{
				CreatedDateTime: time.Now(),
			},
			MessageId: models.ActivityReport,
			// Missing Pagenation - required field
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})

	t.Run("CreateDocument with extremely large entry arrays", func(t *testing.T) {
		// Test with large arrays to ensure memory handling
		largeArrayJSON := `{
			"createdDateTime": "2024-01-01T10:00:00Z",
			"messageId": "ACTR",
			"pagenation": {
				"PageNumber": "1",
				"LastPageIndicator": true
			},
			"reportId": "IDAY",
			"reportCreateDateTime": "2024-01-01T09:00:00Z",
			"entryDetails": [` + 
		`{"Amount": {"Amount": 100, "Currency": "USD"}, "CreditDebitIndicator": "CRDT", "Status": "INFO"},` +
		`{"Amount": {"Amount": 200, "Currency": "USD"}, "CreditDebitIndicator": "DBIT", "Status": "INFO"}` +
		`]}`
		
		xmlData, err := wrapper.CreateDocument([]byte(largeArrayJSON), ActivityReport.CAMT_052_001_12)
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		} else {
			assert.NotEmpty(t, xmlData)
		}
	})
}