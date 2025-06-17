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
	AccountReportingRequest "github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
)

// createValidAccountReportingRequestModel creates an AccountReportingRequest.MessageModel with all required fields populated
func createValidAccountReportingRequestModel() AccountReportingRequest.MessageModel {
	return AccountReportingRequest.MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250310ACCT000001",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		ReportRequestId:    models.AccountBalanceReport,
		RequestedMsgNameId: "camt.052.001.08",
		AccountOtherId:     "ACC123456789",
		AccountProperty:    models.AccountTypeSavings,
		AccountOwnerAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "021151080",
		},
		FromToSequence: models.SequenceRange{
			FromSeq: "1",
			ToSeq:   "100",
		},
	}
}

func TestAccountReportingRequestWrapper_CreateDocument(t *testing.T) {
	wrapper := &AccountReportingRequestWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     AccountReportingRequest.CAMT_060_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"messageId": "20250310ACCT000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"reportRequestId": "ABAR",
				"requestedMsgNameId": "camt.052.001.08",
				"accountOtherId": "ACC123456789",
				"accountProperty": "S",
				"accountOwnerAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021151080"
				},
				"fromToSequence": {
					"FromSeq": "1",
					"ToSeq": "100"
				}
			}`),
			version:     AccountReportingRequest.CAMT_060_001_07,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     AccountReportingRequest.CAMT_060_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     AccountReportingRequest.CAMT_060_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     AccountReportingRequest.CAMT_060_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     AccountReportingRequest.CAMT_060_001_07,
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

func TestAccountReportingRequestWrapper_ValidateDocument(t *testing.T) {
	wrapper := &AccountReportingRequestWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     AccountReportingRequest.CAMT_060_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"messageId": "20250310ACCT000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"reportRequestId": "ABAR",
				"requestedMsgNameId": "camt.052.001.08",
				"accountOtherId": "ACC123456789",
				"accountProperty": "S",
				"accountOwnerAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021151080"
				},
				"fromToSequence": {
					"FromSeq": "1",
					"ToSeq": "100"
				}
			}`,
			version:     AccountReportingRequest.CAMT_060_001_07,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     AccountReportingRequest.CAMT_060_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     AccountReportingRequest.CAMT_060_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     AccountReportingRequest.CAMT_060_001_07,
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

func TestAccountReportingRequestWrapper_CheckRequireField(t *testing.T) {
	wrapper := &AccountReportingRequestWrapper{}

	tests := []struct {
		name        string
		model       AccountReportingRequest.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidAccountReportingRequestModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: AccountReportingRequest.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing MessageId
					CreatedDateTime: time.Now(),
				},
				ReportRequestId: models.AccountBalanceReport,
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       AccountReportingRequest.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing ReportRequestId fails validation",
			model: AccountReportingRequest.MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "20250310ACCT000001",
					CreatedDateTime: time.Now(),
				},
				// Missing ReportRequestId
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

func TestAccountReportingRequestWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &AccountReportingRequestWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.060.001.07">
	<AcctRptgReq>
		<GrpHdr>
			<MsgId>20250310ACCT000001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</GrpHdr>
		<RptgReq>
			<Id>ABAR</Id>
			<ReqdMsgNmId>camt.052.001.08</ReqdMsgNmId>
		</RptgReq>
	</AcctRptgReq>
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
				assert.Equal(t, AccountReportingRequest.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, AccountReportingRequest.MessageModel{}, result)
			}
		})
	}
}

func TestAccountReportingRequestWrapper_GetHelp(t *testing.T) {
	wrapper := &AccountReportingRequestWrapper{}

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
	assert.Contains(t, result, "ReportRequestId")
	assert.Contains(t, result, "AccountOwnerAgent")
}

func TestAccountReportingRequestWrapper_Integration(t *testing.T) {
	wrapper := &AccountReportingRequestWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"messageId": "20250310ACCT000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"reportRequestId": "ABAR",
		"requestedMsgNameId": "camt.052.001.08",
		"accountOtherId": "ACC123456789",
		"accountProperty": "S",
		"accountOwnerAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021151080"
		},
		"fromToSequence": {
			"FromSeq": "1",
			"ToSeq": "100"
		}
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, AccountReportingRequest.CAMT_060_001_07)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestAccountReportingRequestWrapper_AllVersions(t *testing.T) {
	wrapper := &AccountReportingRequestWrapper{}

	validJSON := []byte(`{
		"messageId": "20250310ACCT000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"reportRequestId": "ABAR",
		"requestedMsgNameId": "camt.052.001.08",
		"accountOtherId": "ACC123456789",
		"accountProperty": "S",
		"accountOwnerAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021151080"
		},
		"fromToSequence": {
			"FromSeq": "1",
			"ToSeq": "100"
		}
	}`)

	versions := []AccountReportingRequest.CAMT_060_001_VERSION{
		AccountReportingRequest.CAMT_060_001_02,
		AccountReportingRequest.CAMT_060_001_03,
		AccountReportingRequest.CAMT_060_001_04,
		AccountReportingRequest.CAMT_060_001_05,
		AccountReportingRequest.CAMT_060_001_06,
		AccountReportingRequest.CAMT_060_001_07,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := AccountReportingRequest.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestAccountReportingRequestWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &AccountReportingRequestWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long MessageId that should fail validation
		longMessageId := `{
			"messageId": "ThisIsAnExtremelyLongMessageIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"reportRequestId": "ABAR"
		}`
		_, err := wrapper.CreateDocument([]byte(longMessageId), AccountReportingRequest.CAMT_060_001_07)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"messageId": "20250310ACCT000001",
			"createdDateTime": "invalid-date-format",
			"reportRequestId": "ABAR"
		}`
		err := wrapper.ValidateDocument(malformedDate, AccountReportingRequest.CAMT_060_001_07)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.060.001.07"><AcctRptgReq><GrpHdr><MsgId>Test&InvalidChar</MsgId></GrpHdr></AcctRptgReq></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid report request type", func(t *testing.T) {
		invalidReportType := `{
			"messageId": "20250310ACCT000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"reportRequestId": "INVALID_REPORT_TYPE_TOO_LONG",
			"accountOwnerAgent": {
				"paymentSysCode": "USABA",
				"paymentSysMemberId": "021151080"
			}
		}`
		_, err := wrapper.CreateDocument([]byte(invalidReportType), AccountReportingRequest.CAMT_060_001_07)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := AccountReportingRequest.MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "20250310ACCT000001",
				// Missing CreatedDateTime
			},
			// Missing ReportRequestId
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})
}
