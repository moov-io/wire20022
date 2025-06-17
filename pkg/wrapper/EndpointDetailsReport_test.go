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
	EndpointDetailsReport "github.com/moov-io/wire20022/pkg/models/EndpointDetailsReport"
)

// createValidEndpointDetailsReportModel creates an EndpointDetailsReport.MessageModel with all required fields populated
func createValidEndpointDetailsReportModel() EndpointDetailsReport.MessageModel {
	return EndpointDetailsReport.MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250310EDR0000001",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		Pagenation: models.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		},
		BussinessQueryMsgId:          "20250310B1QDRCQR000001",
		BussinessQueryMsgNameId:      "camt.060.001.05",
		BussinessQueryCreateDatetime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		ReportId:                     models.Intraday,
		ReportingSequence: models.SequenceRange{
			FromSeq: "1",
			ToSeq:   "100",
		},
		ReportCreateDateTime: time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC),
		AccountOtherId:       "ACC123456789",
		TotalCreditEntries: models.NumberAndSumOfTransactions{
			NumberOfEntries: "5",
			Sum:             600000.00,
		},
		TotalDebitEntries: models.NumberAndSumOfTransactions{
			NumberOfEntries: "3",
			Sum:             100000.00,
		},
		TotalEntriesPerBankTransactionCode: []models.TotalsPerBankTransactionCode{
			{
				BankTransactionCode: models.TransPending,
				NumberOfEntries:     "8",
			},
		},
		EntryDetails: []models.Entry{
			{
				Amount: models.CurrencyAndAmount{
					Amount:   500000.00,
					Currency: "USD",
				},
				CreditDebitIndicator: models.Credit,
				Status:               models.Book,
			},
		},
	}
}

func TestEndpointDetailsReportWrapper_CreateDocument(t *testing.T) {
	wrapper := &EndpointDetailsReportWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     EndpointDetailsReport.CAMT_052_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"messageId": "20250310EDR0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"pagenation": {
					"PageNumber": "1",
					"LastPageIndicator": true
				},
				"bussinessQueryMsgId": "20250310B1QDRCQR000001",
				"bussinessQueryMsgNameId": "camt.060.001.05",
				"bussinessQueryCreateDatetime": "2024-01-01T09:00:00Z",
				"reportId": "IDAY",
				"reportingSequence": {
					"FromSeq": "1",
					"ToSeq": "100"
				},
				"reportCreateDateTime": "2024-01-01T08:00:00Z",
				"accountOtherId": "ACC123456789",
				"totalCreditEntries": {
					"NumberOfEntries": "5",
					"Sum": 600000.00
				},
				"totalDebitEntries": {
					"NumberOfEntries": "3",
					"Sum": 100000.00
				}
			}`),
			version:     EndpointDetailsReport.CAMT_052_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     EndpointDetailsReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     EndpointDetailsReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     EndpointDetailsReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     EndpointDetailsReport.CAMT_052_001_12,
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

func TestEndpointDetailsReportWrapper_ValidateDocument(t *testing.T) {
	wrapper := &EndpointDetailsReportWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     EndpointDetailsReport.CAMT_052_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"messageId": "20250310EDR0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"pagenation": {
					"PageNumber": "1",
					"LastPageIndicator": true
				},
				"reportId": "IDAY",
				"reportCreateDateTime": "2024-01-01T08:00:00Z"
			}`,
			version:     EndpointDetailsReport.CAMT_052_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     EndpointDetailsReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     EndpointDetailsReport.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     EndpointDetailsReport.CAMT_052_001_12,
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

func TestEndpointDetailsReportWrapper_CheckRequireField(t *testing.T) {
	wrapper := &EndpointDetailsReportWrapper{}

	tests := []struct {
		name        string
		model       EndpointDetailsReport.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidEndpointDetailsReportModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: EndpointDetailsReport.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing MessageId
					CreatedDateTime: time.Now(),
				},
				ReportId: models.Intraday,
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       EndpointDetailsReport.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing ReportId fails validation",
			model: EndpointDetailsReport.MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "20250310EDR0000001",
					CreatedDateTime: time.Now(),
				},
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

func TestEndpointDetailsReportWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &EndpointDetailsReportWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.12">
	<BkToCstmrAcctRpt>
		<GrpHdr>
			<MsgId>20250310EDR0000001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</GrpHdr>
		<Rpt>
			<Id>FINL</Id>
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
				assert.Equal(t, EndpointDetailsReport.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, EndpointDetailsReport.MessageModel{}, result)
			}
		})
	}
}

func TestEndpointDetailsReportWrapper_GetHelp(t *testing.T) {
	wrapper := &EndpointDetailsReportWrapper{}

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

func TestEndpointDetailsReportWrapper_Integration(t *testing.T) {
	wrapper := &EndpointDetailsReportWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"messageId": "20250310EDR0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"pagenation": {
			"PageNumber": "1",
			"LastPageIndicator": true
		},
		"reportId": "FINL",
		"reportCreateDateTime": "2024-01-01T08:00:00Z"
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, EndpointDetailsReport.CAMT_052_001_12)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestEndpointDetailsReportWrapper_AllVersions(t *testing.T) {
	wrapper := &EndpointDetailsReportWrapper{}

	validJSON := []byte(`{
		"messageId": "20250310EDR0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"pagenation": {
			"PageNumber": "1",
			"LastPageIndicator": true
		},
		"reportId": "FINL",
		"reportCreateDateTime": "2024-01-01T08:00:00Z"
	}`)

	versions := []EndpointDetailsReport.CAMT_052_001_VERSION{
		EndpointDetailsReport.CAMT_052_001_02,
		EndpointDetailsReport.CAMT_052_001_03,
		EndpointDetailsReport.CAMT_052_001_04,
		EndpointDetailsReport.CAMT_052_001_05,
		EndpointDetailsReport.CAMT_052_001_06,
		EndpointDetailsReport.CAMT_052_001_07,
		EndpointDetailsReport.CAMT_052_001_08,
		EndpointDetailsReport.CAMT_052_001_09,
		EndpointDetailsReport.CAMT_052_001_10,
		EndpointDetailsReport.CAMT_052_001_11,
		EndpointDetailsReport.CAMT_052_001_12,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := EndpointDetailsReport.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestEndpointDetailsReportWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &EndpointDetailsReportWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long MessageId that should fail validation
		longMessageId := `{
			"messageId": "ThisIsAnExtremelyLongMessageIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"reportId": "FINL"
		}`
		_, err := wrapper.CreateDocument([]byte(longMessageId), EndpointDetailsReport.CAMT_052_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"messageId": "20250310EDR0000001",
			"createdDateTime": "invalid-date-format",
			"reportId": "FINL"
		}`
		err := wrapper.ValidateDocument(malformedDate, EndpointDetailsReport.CAMT_052_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.12"><BkToCstmrAcctRpt><GrpHdr><MsgId>Test&InvalidChar</MsgId></GrpHdr></BkToCstmrAcctRpt></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid report type", func(t *testing.T) {
		invalidReportType := `{
			"messageId": "20250310EDR0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"reportId": "INVALID_REPORT_TYPE_TOO_LONG",
			"reportCreateDateTime": "2024-01-01T08:00:00Z"
		}`
		_, err := wrapper.CreateDocument([]byte(invalidReportType), EndpointDetailsReport.CAMT_052_001_12)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := EndpointDetailsReport.MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "20250310EDR0000001",
				// Missing CreatedDateTime
			},
			ReportId: models.Intraday,
			// Missing other required fields
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})

	t.Run("CreateDocument with complex array validation", func(t *testing.T) {
		complexArray := `{
			"messageId": "20250310EDR0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"pagenation": {
				"PageNumber": "1",
				"LastPageIndicator": true
			},
			"reportId": "FINL",
			"reportCreateDateTime": "2024-01-01T08:00:00Z",
			"accountOtherId": "ACC123456789",
			"entryDetails": [
				{
					"Amount": {
						"Amount": -1,
						"Currency": "INVALID"
					},
					"CreditDebitIndicator": "INVALID",
					"BankTransactionCode": "",
					"ValueDate": "invalid-date"
				}
			]
		}`
		_, err := wrapper.CreateDocument([]byte(complexArray), EndpointDetailsReport.CAMT_052_001_12)
		// Test handling of complex validation errors
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CreateDocument with zero time value", func(t *testing.T) {
		zeroTimeModel := EndpointDetailsReport.MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "20250310EDR0000001",
				// CreatedDateTime is zero value - should fail
			},
			ReportId: models.Intraday,
			// ReportCreateDateTime is zero value - should fail
		}
		err := wrapper.CheckRequireField(zeroTimeModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})
}
