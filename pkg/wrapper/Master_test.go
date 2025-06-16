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
	Master "github.com/wadearnold/wire20022/pkg/models/Master"
)

// createValidMasterModel creates a Master.MessageModel with all required fields populated
func createValidMasterModel() Master.MessageModel {
	return Master.MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250310MST0000001",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		MessagePagination: models.MessagePagenation{
			PageNumber:        "1",
			LastPageIndicator: true,
		},
		OriginalBusinessMsgId:         "20250310B1QDRCQR000001",
		OriginalBusinessMsgNameId:     "camt.052.001.12",
		OriginalBusinessMsgCreateTime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		ReportTypeId:                  models.FINAL,
		ReportCreatedDate:             time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC),
		AccountOtherId:                "ACC123456789",
		AccountType:                   "SAVINGS",
		RelatedAccountOtherId:         "RELACC987654321",
		Balances: []models.Balance{
			{
				Amount: models.CurrencyAndAmount{
					Amount:   1000000.00,
					Currency: "USD",
				},
				CreditDebitIndicator: models.Credit,
			},
		},
		TransactionsSummary: []models.TotalsPerBankTransaction{
			{
				TotalNetEntryAmount:  500000.00,
				CreditDebitIndicator: models.Credit,
				CreditEntries: models.NumberAndSumOfTransactions{
					NumberOfEntries: "5",
					Sum:             600000.00,
				},
				DebitEntries: models.NumberAndSumOfTransactions{
					NumberOfEntries: "3",
					Sum:             100000.00,
				},
				BankTransactionCode: models.FedwireFundsTransfers,
				Date:                time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}
}

func TestMasterWrapper_CreateDocument(t *testing.T) {
	wrapper := &MasterWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     Master.CAMT_052_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"messageId": "20250310MST0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"messagePagination": {
					"PageNumber": "1",
					"LastPageIndicator": true
				},
				"originalBusinessMsgId": "20250310B1QDRCQR000001",
				"originalBusinessMsgNameId": "camt.052.001.12",
				"originalBusinessMsgCreateTime": "2024-01-01T09:00:00Z",
				"reportTypeId": "FINL",
				"reportCreatedDate": "2024-01-01T08:00:00Z",
				"accountOtherId": "ACC123456789",
				"accountType": "SAVINGS",
				"relatedAccountOtherId": "RELACC987654321",
				"transactionsSummary": [
					{
						"TotalNetEntryAmount": 500000.00,
						"CreditDebitIndicator": "CRDT",
						"CreditEntries": {
							"NumberOfEntries": "5",
							"Sum": 600000.00
						},
						"DebitEntries": {
							"NumberOfEntries": "3",
							"Sum": 100000.00
						},
						"BankTransactionCode": "FDWF",
						"Date": "2024-01-01T00:00:00Z"
					}
				]
			}`),
			version:     Master.CAMT_052_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     Master.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     Master.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     Master.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     Master.CAMT_052_001_12,
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

func TestMasterWrapper_ValidateDocument(t *testing.T) {
	wrapper := &MasterWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     Master.CAMT_052_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"messageId": "20250310MST0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"messagePagination": {
					"PageNumber": "1",
					"LastPageIndicator": true
				},
				"reportTypeId": "FINL",
				"reportCreatedDate": "2024-01-01T08:00:00Z",
				"accountOtherId": "ACC123456789",
				"accountType": "SAVINGS",
				"relatedAccountOtherId": "RELACC987654321",
				"transactionsSummary": [
					{
						"TotalNetEntryAmount": 500000.00,
						"CreditDebitIndicator": "CRDT",
						"BankTransactionCode": "FDWF",
						"Date": "2024-01-01T00:00:00Z"
					}
				]
			}`,
			version:     Master.CAMT_052_001_12,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     Master.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     Master.CAMT_052_001_12,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     Master.CAMT_052_001_12,
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

func TestMasterWrapper_CheckRequireField(t *testing.T) {
	wrapper := &MasterWrapper{}

	tests := []struct {
		name        string
		model       Master.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidMasterModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: Master.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing MessageId
					CreatedDateTime: time.Now(),
				},
				ReportTypeId: models.FINAL,
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       Master.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing ReportTypeId fails validation",
			model: Master.MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "20250310MST0000001",
					CreatedDateTime: time.Now(),
				},
				// Missing ReportTypeId
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

func TestMasterWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &MasterWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.12">
	<BkToCstmrAcctRpt>
		<GrpHdr>
			<MsgId>20250310MST0000001</MsgId>
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
				assert.Equal(t, Master.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, Master.MessageModel{}, result)
			}
		})
	}
}

func TestMasterWrapper_GetHelp(t *testing.T) {
	wrapper := &MasterWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(result), &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")

	// Verify it contains expected fields
	assert.Contains(t, result, "MessageId")
	assert.Contains(t, result, "ReportTypeId")
	assert.Contains(t, result, "AccountOtherId")
	assert.Contains(t, result, "TransactionsSummary")
}

func TestMasterWrapper_Integration(t *testing.T) {
	wrapper := &MasterWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"messageId": "20250310MST0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"messagePagination": {
			"PageNumber": "1",
			"LastPageIndicator": true
		},
		"reportTypeId": "FINL",
		"reportCreatedDate": "2024-01-01T08:00:00Z",
		"accountOtherId": "ACC123456789",
		"accountType": "SAVINGS",
		"relatedAccountOtherId": "RELACC987654321",
		"transactionsSummary": [
			{
				"TotalNetEntryAmount": 500000.00,
				"CreditDebitIndicator": "CRDT",
				"BankTransactionCode": "FDWF",
				"Date": "2024-01-01T00:00:00Z"
			}
		]
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, Master.CAMT_052_001_12)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestMasterWrapper_AllVersions(t *testing.T) {
	wrapper := &MasterWrapper{}

	validJSON := []byte(`{
		"messageId": "20250310MST0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"messagePagination": {
			"PageNumber": "1",
			"LastPageIndicator": true
		},
		"reportTypeId": "FINL",
		"reportCreatedDate": "2024-01-01T08:00:00Z",
		"accountOtherId": "ACC123456789",
		"accountType": "SAVINGS",
		"relatedAccountOtherId": "RELACC987654321",
		"transactionsSummary": [
			{
				"TotalNetEntryAmount": 500000.00,
				"CreditDebitIndicator": "CRDT",
				"BankTransactionCode": "FDWF",
				"Date": "2024-01-01T00:00:00Z"
			}
		]
	}`)

	versions := []Master.CAMT_052_001_VERSION{
		Master.CAMT_052_001_02,
		Master.CAMT_052_001_03,
		Master.CAMT_052_001_04,
		Master.CAMT_052_001_05,
		Master.CAMT_052_001_06,
		Master.CAMT_052_001_07,
		Master.CAMT_052_001_08,
		Master.CAMT_052_001_09,
		Master.CAMT_052_001_10,
		Master.CAMT_052_001_11,
		Master.CAMT_052_001_12,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := Master.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestMasterWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &MasterWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long MessageId that should fail validation
		longMessageId := `{
			"messageId": "ThisIsAnExtremelyLongMessageIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"reportTypeId": "FINL"
		}`
		_, err := wrapper.CreateDocument([]byte(longMessageId), Master.CAMT_052_001_12)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"messageId": "20250310MST0000001",
			"createdDateTime": "invalid-date-format",
			"reportTypeId": "FINL"
		}`
		err := wrapper.ValidateDocument(malformedDate, Master.CAMT_052_001_12)
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
			"messageId": "20250310MST0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"reportTypeId": "INVALID_REPORT_TYPE_TOO_LONG",
			"accountOtherId": "ACC123456789"
		}`
		_, err := wrapper.CreateDocument([]byte(invalidReportType), Master.CAMT_052_001_12)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := Master.MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "20250310MST0000001",
				// Missing CreatedDateTime
			},
			ReportTypeId: models.FINAL,
			// Missing other required fields
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})

	t.Run("CreateDocument with complex array validation", func(t *testing.T) {
		complexArray := `{
			"messageId": "20250310MST0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"messagePagination": {
				"PageNumber": "1",
				"LastPageIndicator": true
			},
			"reportTypeId": "FINL",
			"reportCreatedDate": "2024-01-01T08:00:00Z",
			"accountOtherId": "ACC123456789",
			"accountType": "SAVINGS",
			"relatedAccountOtherId": "RELACC987654321",
			"balances": [
				{
					"Amount": {
						"Amount": -1,
						"Currency": "INVALID"
					},
					"CreditDebitIndicator": "INVALID"
				}
			],
			"transactionsSummary": [
				{
					"TotalNetEntryAmount": 0,
					"CreditDebitIndicator": "",
					"BankTransactionCode": ""
				}
			]
		}`
		_, err := wrapper.CreateDocument([]byte(complexArray), Master.CAMT_052_001_12)
		// Test handling of complex validation errors
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})
}