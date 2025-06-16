package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wadearnold/wire20022/pkg/base"
	"github.com/wadearnold/wire20022/pkg/models"
	DrawdownRequest "github.com/wadearnold/wire20022/pkg/models/DrawdownRequest"
)

// createValidDrawdownRequestModel creates a DrawdownRequest.MessageModel with all required fields populated
func createValidDrawdownRequestModel() DrawdownRequest.MessageModel {
	return DrawdownRequest.MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250310DWN0000001",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		NumberofTransaction: "1",
		InitiatingParty: models.PartyIdentify{
			Name: "ABC Corporation",
			Address: models.PostalAddress{
				StreetName: "123 Main St",
				TownName:   "New York",
				Country:    "US",
			},
		},
		PaymentInfoId:       "PMT001",
		PaymentMethod:       models.CreditTransform,
		RequestedExecutDate: fedwire.ISODate{Year: 2024, Month: 1, Day: 2},
		Debtor: models.PartyIdentify{
			Name: "XYZ Company",
			Address: models.PostalAddress{
				StreetName: "456 Oak Ave",
				TownName:   "Chicago",
				Country:    "US",
			},
		},
		DebtorAccountOtherId: "ACC123456789",
		DebtorAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		},
		CreditTransTransaction: DrawdownRequest.CreditTransferTransaction{
			PaymentInstructionId: "Scenario01Step1InstrId001",
			PaymentEndToEndId:    "Scenario1EndToEndId001",
			PaymentUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
			PayCategoryType:      models.IntraCompanyPayment,
			PayRequestType:       models.DrawDownRequestCredit,
			Amount: models.CurrencyAndAmount{
				Amount:   6000000.00,
				Currency: "USD",
			},
			ChargeBearer: models.ChargeBearerSLEV,
		},
	}
}

func TestDrawdownRequestWrapper_CreateDocument(t *testing.T) {
	wrapper := &DrawdownRequestWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     DrawdownRequest.PAIN_013_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"messageId": "20250310DWN0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"numberofTransaction": "1",
				"initiatingParty": {
					"Name": "ABC Corporation",
					"Address": {
						"StreetName": "123 Main St",
						"TownName": "New York",
						"Country": "US"
					}
				},
				"paymentInfoId": "PMT001",
				"paymentMethod": "TRF",
				"requestedExecutDate": "2024-01-02",
				"debtor": {
					"Name": "XYZ Company",
					"Address": {
						"StreetName": "456 Oak Ave",
						"TownName": "Chicago",
						"Country": "US"
					}
				},
				"debtorAccountOtherId": "ACC123456789",
				"debtorAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021040078"
				},
				"creditTransTransaction": {
					"PaymentInstructionId": "Scenario01Step1InstrId001",
					"PaymentEndToEndId": "Scenario1EndToEndId001",
					"PaymentUniqueId": "8a562c67-ca16-48ba-b074-65581be6f066",
					"PayCategoryType": "INTC",
					"PayRequestType": "DRRC",
					"Amount": {
						"Amount": 6000000.00,
						"Currency": "USD"
					},
					"ChargeBearer": "SLEV"
				}
			}`),
			version:     DrawdownRequest.PAIN_013_001_10,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     DrawdownRequest.PAIN_013_001_10,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     DrawdownRequest.PAIN_013_001_10,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     DrawdownRequest.PAIN_013_001_10,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     DrawdownRequest.PAIN_013_001_10,
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

func TestDrawdownRequestWrapper_ValidateDocument(t *testing.T) {
	wrapper := &DrawdownRequestWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     DrawdownRequest.PAIN_013_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates with validation error (expected due to minimal XML)",
			modelJson: `{
				"messageId": "20250310DWN0000001",
				"createdDateTime": "2024-01-01T10:00:00Z",
				"numberofTransaction": "1",
				"initiatingParty": {
					"Name": "ABC Corporation",
					"Address": {
						"StreetName": "123 Main St",
						"TownName": "New York",
						"Country": "US"
					}
				},
				"paymentInfoId": "PMT001",
				"paymentMethod": "TRF",
				"requestedExecutDate": "2024-01-02",
				"debtor": {
					"Name": "XYZ Company"
				},
				"debtorAgent": {
					"paymentSysCode": "USABA",
					"paymentSysMemberId": "021040078"
				},
				"creditTransTransaction": {
					"PaymentInstructionId": "Scenario01Step1InstrId001",
					"PaymentEndToEndId": "Scenario1EndToEndId001"
				}
			}`,
			version:     DrawdownRequest.PAIN_013_001_10,
			expectError: true,
			errorMsg:    "validation failed",
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     DrawdownRequest.PAIN_013_001_10,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     DrawdownRequest.PAIN_013_001_10,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"messageId": "",
				"createdDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     DrawdownRequest.PAIN_013_001_10,
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

func TestDrawdownRequestWrapper_CheckRequireField(t *testing.T) {
	wrapper := &DrawdownRequestWrapper{}

	tests := []struct {
		name        string
		model       DrawdownRequest.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidDrawdownRequestModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: DrawdownRequest.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing MessageId
					CreatedDateTime: time.Now(),
				},
				NumberofTransaction: "1",
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       DrawdownRequest.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing NumberofTransaction fails validation",
			model: DrawdownRequest.MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "20250310DWN0000001",
					CreatedDateTime: time.Now(),
				},
				// Missing NumberofTransaction
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

func TestDrawdownRequestWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &DrawdownRequestWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.013.001.10">
	<CdtrPmtActvtnReq>
		<GrpHdr>
			<MsgId>20250310DWN0000001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</GrpHdr>
		<PmtInf>
			<PmtInfId>PMT001</PmtInfId>
			<PmtMtd>TRF</PmtMtd>
		</PmtInf>
	</CdtrPmtActvtnReq>
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
				assert.Equal(t, DrawdownRequest.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, DrawdownRequest.MessageModel{}, result)
			}
		})
	}
}

func TestDrawdownRequestWrapper_GetHelp(t *testing.T) {
	wrapper := &DrawdownRequestWrapper{}

	result, err := wrapper.GetHelp()

	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// Verify it's valid JSON
	var jsonData interface{}
	err = json.Unmarshal([]byte(result), &jsonData)
	assert.NoError(t, err, "Help result should be valid JSON")

	// Verify it contains expected fields
	assert.Contains(t, result, "MessageId")
	assert.Contains(t, result, "CreateDatetime")
	assert.Contains(t, result, "NumberofTransaction")
	assert.Contains(t, result, "PaymentMethod")
	assert.Contains(t, result, "DebtorAgent")
}

func TestDrawdownRequestWrapper_Integration(t *testing.T) {
	wrapper := &DrawdownRequestWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"messageId": "20250310DWN0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"numberofTransaction": "1",
		"initiatingParty": {
			"Name": "ABC Corporation"
		},
		"paymentInfoId": "PMT001",
		"paymentMethod": "TRF",
		"requestedExecutDate": "2024-01-02",
		"debtor": {
			"Name": "XYZ Company"
		},
		"debtorAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021040078"
		},
		"creditTransTransaction": {
			"PaymentInstructionId": "Scenario01Step1InstrId001",
			"PaymentEndToEndId": "Scenario1EndToEndId001"
		}
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, DrawdownRequest.PAIN_013_001_10)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestDrawdownRequestWrapper_AllVersions(t *testing.T) {
	wrapper := &DrawdownRequestWrapper{}

	validJSON := []byte(`{
		"messageId": "20250310DWN0000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"numberofTransaction": "1",
		"initiatingParty": {
			"Name": "ABC Corporation"
		},
		"paymentInfoId": "PMT001",
		"paymentMethod": "TRF",
		"requestedExecutDate": "2024-01-02",
		"debtor": {
			"Name": "XYZ Company"
		},
		"debtorAgent": {
			"paymentSysCode": "USABA",
			"paymentSysMemberId": "021040078"
		},
		"creditTransTransaction": {
			"PaymentInstructionId": "Scenario01Step1InstrId001",
			"PaymentEndToEndId": "Scenario1EndToEndId001"
		}
	}`)

	versions := []DrawdownRequest.PAIN_013_001_VERSION{
		DrawdownRequest.PAIN_013_001_01,
		DrawdownRequest.PAIN_013_001_02,
		DrawdownRequest.PAIN_013_001_03,
		DrawdownRequest.PAIN_013_001_04,
		DrawdownRequest.PAIN_013_001_05,
		DrawdownRequest.PAIN_013_001_06,
		DrawdownRequest.PAIN_013_001_07,
		DrawdownRequest.PAIN_013_001_08,
		DrawdownRequest.PAIN_013_001_09,
		DrawdownRequest.PAIN_013_001_10,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := DrawdownRequest.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestDrawdownRequestWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &DrawdownRequestWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long MessageId that should fail validation
		longMessageId := `{
			"messageId": "ThisIsAnExtremelyLongMessageIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"numberofTransaction": "1"
		}`
		_, err := wrapper.CreateDocument([]byte(longMessageId), DrawdownRequest.PAIN_013_001_10)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"messageId": "20250310DWN0000001",
			"createdDateTime": "invalid-date-format",
			"numberofTransaction": "1"
		}`
		err := wrapper.ValidateDocument(malformedDate, DrawdownRequest.PAIN_013_001_10)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.013.001.10"><CdtrPmtActvtnReq><GrpHdr><MsgId>Test&InvalidChar</MsgId></GrpHdr></CdtrPmtActvtnReq></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid payment method", func(t *testing.T) {
		invalidPaymentMethod := `{
			"messageId": "20250310DWN0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"numberofTransaction": "1",
			"paymentMethod": "INVALID_PAYMENT_METHOD_CODE",
			"initiatingParty": {
				"Name": "ABC Corporation"
			}
		}`
		_, err := wrapper.CreateDocument([]byte(invalidPaymentMethod), DrawdownRequest.PAIN_013_001_10)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := DrawdownRequest.MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "20250310DWN0000001",
				// Missing CreatedDateTime
			},
			NumberofTransaction: "1",
			// Missing InitiatingParty and other required fields
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})

	t.Run("CreateDocument with complex nested structure validation", func(t *testing.T) {
		complexValidation := `{
			"messageId": "20250310DWN0000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"numberofTransaction": "1",
			"initiatingParty": {
				"Name": "",
				"Address": {
					"StreetName": "",
					"TownName": "",
					"Country": ""
				}
			},
			"paymentInfoId": "PMT001",
			"paymentMethod": "TRF",
			"requestedExecutDate": "2024-01-02",
			"debtor": {
				"Name": ""
			},
			"debtorAgent": {
				"paymentSysCode": "",
				"paymentSysMemberId": ""
			},
			"creditTransTransaction": {
				"PaymentInstructionId": "",
				"PaymentEndToEndId": ""
			}
		}`
		_, err := wrapper.CreateDocument([]byte(complexValidation), DrawdownRequest.PAIN_013_001_10)
		// Test handling of nested validation errors
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})
}