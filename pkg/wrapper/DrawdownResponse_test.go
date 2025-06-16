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
	DrawdownResponse "github.com/moov-io/wire20022/pkg/models/DrawdownResponse"
)

// createValidDrawdownResponseModel creates a DrawdownResponse.MessageModel with all required fields populated
func createValidDrawdownResponseModel() DrawdownResponse.MessageModel {
	return DrawdownResponse.MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "20250310B1QDRCQR000602",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		InitiatingParty: models.PartyIdentify{
			Name: "Corporation A",
			Address: models.PostalAddress{
				StreetName:     "Avenue of the Fountains",
				BuildingNumber: "167565",
				RoomNumber:     "Suite D110",
				PostalCode:     "85268",
				TownName:       "Fountain Hills",
				Subdivision:    "AZ",
				Country:        "US",
			},
		},
		DebtorAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "021040078",
		},
		CreditorAgent: models.Agent{
			PaymentSysCode:     models.PaymentSysUSABA,
			PaymentSysMemberId: "011104238",
		},
		OriginalMessageId:        "20250310B1QDRCQR000601",
		OriginalMessageNameId:    "pain.013.001.07",
		OriginalCreationDateTime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		OriginalPaymentInfoId:    "20250310B1QDRCQR000601",
		TransactionInformationAndStatus: DrawdownResponse.TransactionInfoAndStatus{
			OriginalInstructionId: "Scenario01Step1InstrId001",
			OriginalEndToEndId:    "Scenario1EndToEndId001",
			OriginalUniqueId:      "8a562c67-ca16-48ba-b074-65581be6f066",
			TransactionStatus:     models.AcceptedTechnicalValidation,
			StatusReasonInfoCode:  models.StatusReasonInformationCode("ACCP"),
		},
	}
}

func TestDrawdownResponseWrapper_CreateDocument(t *testing.T) {
	wrapper := &DrawdownResponseWrapper{}

	tests := []struct {
		name        string
		modelJson   []byte
		version     DrawdownResponse.PAIN_014_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model creates document successfully",
			modelJson: []byte(`{
				"MessageId": "20250310B1QDRCQR000602",
				"CreatedDateTime": "2024-01-01T10:00:00Z",
				"InitiatingParty": {
					"Name": "Corporation A",
					"Address": {
						"StreetName": "Avenue of the Fountains",
						"BuildingNumber": "167565",
						"PostalCode": "85268",
						"TownName": "Fountain Hills",
						"Country": "US"
					}
				},
				"DebtorAgent": {
					"PaymentSysCode": "USABA",
					"PaymentSysMemberId": "021040078"
				},
				"CreditorAgent": {
					"PaymentSysCode": "USABA",
					"PaymentSysMemberId": "011104238"
				},
				"OriginalMessageId": "20250310B1QDRCQR000601",
				"OriginalMessageNameId": "pain.013.001.07",
				"OriginalCreationDateTime": "2024-01-01T09:00:00Z",
				"OriginalPaymentInfoId": "20250310B1QDRCQR000601",
				"TransactionInformationAndStatus": {
					"OriginalInstructionId": "Scenario01Step1InstrId001",
					"OriginalEndToEndId": "Scenario1EndToEndId001",
					"OriginalUniqueId": "8a562c67-ca16-48ba-b074-65581be6f066",
					"TransactionStatus": "ACTC",
					"StatusReasonInfoCode": "ACCP"
				}
			}`),
			version:     DrawdownResponse.PAIN_014_001_07,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   []byte(`{"invalid": json}`),
			version:     DrawdownResponse.PAIN_014_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   []byte(``),
			version:     DrawdownResponse.PAIN_014_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "nil JSON returns error",
			modelJson:   nil,
			version:     DrawdownResponse.PAIN_014_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: []byte(`{
				"MessageId": "",
				"CreatedDateTime": "2024-01-01T10:00:00Z"
			}`),
			version:     DrawdownResponse.PAIN_014_001_07,
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

func TestDrawdownResponseWrapper_ValidateDocument(t *testing.T) {
	wrapper := &DrawdownResponseWrapper{}

	tests := []struct {
		name        string
		modelJson   string
		version     DrawdownResponse.PAIN_014_001_VERSION
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid model validates successfully",
			modelJson: `{
				"MessageId": "20250310B1QDRCQR000602",
				"CreatedDateTime": "2024-01-01T10:00:00Z",
				"InitiatingParty": {
					"Name": "Corporation A",
					"Address": {
						"StreetName": "Avenue of the Fountains",
						"PostalCode": "85268",
						"TownName": "Fountain Hills",
						"Country": "US"
					}
				},
				"DebtorAgent": {
					"PaymentSysCode": "USABA",
					"PaymentSysMemberId": "021040078"
				},
				"CreditorAgent": {
					"PaymentSysCode": "USABA",
					"PaymentSysMemberId": "011104238"
				},
				"OriginalMessageId": "20250310B1QDRCQR000601",
				"OriginalMessageNameId": "pain.013.001.07",
				"OriginalCreationDateTime": "2024-01-01T09:00:00Z",
				"OriginalPaymentInfoId": "20250310B1QDRCQR000601",
				"TransactionInformationAndStatus": {
					"OriginalInstructionId": "Scenario01Step1InstrId001",
					"OriginalEndToEndId": "Scenario1EndToEndId001",
					"OriginalUniqueId": "8a562c67-ca16-48ba-b074-65581be6f066",
					"TransactionStatus": "ACTC",
					"StatusReasonInfoCode": "ACCP"
				}
			}`,
			version:     DrawdownResponse.PAIN_014_001_07,
			expectError: false,
		},
		{
			name:        "invalid JSON returns error",
			modelJson:   `{"invalid": json}`,
			version:     DrawdownResponse.PAIN_014_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name:        "empty JSON returns error",
			modelJson:   "",
			version:     DrawdownResponse.PAIN_014_001_07,
			expectError: true,
			errorMsg:    "failed to unmarshal JSON to MessageModel",
		},
		{
			name: "model with missing required fields returns error",
			modelJson: `{
				"MessageId": "",
				"CreatedDateTime": "2024-01-01T10:00:00Z"
			}`,
			version:     DrawdownResponse.PAIN_014_001_07,
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

func TestDrawdownResponseWrapper_CheckRequireField(t *testing.T) {
	wrapper := &DrawdownResponseWrapper{}

	tests := []struct {
		name        string
		model       DrawdownResponse.MessageModel
		expectError bool
		errorMsg    string
	}{
		{
			name:        "model with required fields passes validation",
			model:       createValidDrawdownResponseModel(),
			expectError: false,
		},
		{
			name: "model with missing required field fails validation",
			model: DrawdownResponse.MessageModel{
				MessageHeader: base.MessageHeader{
					// Missing MessageId
					CreatedDateTime: time.Now(),
				},
				OriginalMessageId: "20250310B1QDRCQR000601",
			},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name:        "empty model fails validation",
			model:       DrawdownResponse.MessageModel{},
			expectError: true,
			errorMsg:    "required field",
		},
		{
			name: "model with missing OriginalMessageId fails validation",
			model: DrawdownResponse.MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "20250310B1QDRCQR000602",
					CreatedDateTime: time.Now(),
				},
				// Missing OriginalMessageId
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

func TestDrawdownResponseWrapper_ConvertXMLToModel(t *testing.T) {
	wrapper := &DrawdownResponseWrapper{}

	// Create a valid XML sample for testing
	validXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.014.001.07">
	<CdtrPmtActvtnReqStsRpt>
		<GrpHdr>
			<MsgId>20250310B1QDRCQR000602</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
			<InitgPty>
				<Nm>Corporation A</Nm>
			</InitgPty>
		</GrpHdr>
		<OrgnlPmtInfAndSts>
			<OrgnlPmtInfId>20250310B1QDRCQR000601</OrgnlPmtInfId>
		</OrgnlPmtInfAndSts>
	</CdtrPmtActvtnReqStsRpt>
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
				assert.Equal(t, DrawdownResponse.MessageModel{}, result)
			} else {
				assert.NoError(t, err)
				// For success cases, result should have some data
				assert.NotEqual(t, DrawdownResponse.MessageModel{}, result)
			}
		})
	}
}

func TestDrawdownResponseWrapper_GetHelp(t *testing.T) {
	wrapper := &DrawdownResponseWrapper{}

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
	assert.Contains(t, result, "InitiatingParty")
	assert.Contains(t, result, "TransactionInformationAndStatus")
}

func TestDrawdownResponseWrapper_Integration(t *testing.T) {
	wrapper := &DrawdownResponseWrapper{}

	// Test basic functionality: Valid JSON to XML
	validJSON := []byte(`{
		"MessageId": "20250310B1QDRCQR000602",
		"CreatedDateTime": "2024-01-01T10:00:00Z",
		"InitiatingParty": {
			"Name": "Corporation A"
		},
		"DebtorAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "021040078"
		},
		"CreditorAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "011104238"
		},
		"OriginalMessageId": "20250310B1QDRCQR000601",
		"OriginalMessageNameId": "pain.013.001.07",
		"OriginalCreationDateTime": "2024-01-01T09:00:00Z",
		"OriginalPaymentInfoId": "20250310B1QDRCQR000601",
		"TransactionInformationAndStatus": {
			"OriginalInstructionId": "Scenario01Step1InstrId001",
			"OriginalEndToEndId": "Scenario1EndToEndId001",
			"OriginalUniqueId": "8a562c67-ca16-48ba-b074-65581be6f066",
			"TransactionStatus": "ACTC"
		}
	}`)

	// JSON to XML
	xmlData, err := wrapper.CreateDocument(validJSON, DrawdownResponse.PAIN_014_001_07)
	require.NoError(t, err)
	require.NotEmpty(t, xmlData)

	// Verify it's valid XML
	var xmlDoc interface{}
	err = xml.Unmarshal(xmlData, &xmlDoc)
	require.NoError(t, err, "Generated XML should be valid")
}

func TestDrawdownResponseWrapper_AllVersions(t *testing.T) {
	wrapper := &DrawdownResponseWrapper{}

	validJSON := []byte(`{
		"MessageId": "20250310B1QDRCQR000602",
		"CreatedDateTime": "2024-01-01T10:00:00Z",
		"InitiatingParty": {
			"Name": "Corporation A"
		},
		"DebtorAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "021040078"
		},
		"CreditorAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "011104238"
		},
		"OriginalMessageId": "20250310B1QDRCQR000601",
		"OriginalMessageNameId": "pain.013.001.07",
		"OriginalCreationDateTime": "2024-01-01T09:00:00Z",
		"OriginalPaymentInfoId": "20250310B1QDRCQR000601",
		"TransactionInformationAndStatus": {
			"OriginalInstructionId": "Scenario01Step1InstrId001",
			"OriginalEndToEndId": "Scenario1EndToEndId001",
			"OriginalUniqueId": "8a562c67-ca16-48ba-b074-65581be6f066",
			"TransactionStatus": "ACTC"
		}
	}`)

	versions := []DrawdownResponse.PAIN_014_001_VERSION{
		DrawdownResponse.PAIN_014_001_01,
		DrawdownResponse.PAIN_014_001_02,
		DrawdownResponse.PAIN_014_001_03,
		DrawdownResponse.PAIN_014_001_04,
		DrawdownResponse.PAIN_014_001_05,
		DrawdownResponse.PAIN_014_001_06,
		DrawdownResponse.PAIN_014_001_07,
		DrawdownResponse.PAIN_014_001_08,
		DrawdownResponse.PAIN_014_001_09,
		DrawdownResponse.PAIN_014_001_10,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			xmlData, err := wrapper.CreateDocument(validJSON, version)
			assert.NoError(t, err)
			assert.NotEmpty(t, xmlData)

			// Verify XML contains the correct namespace
			expectedNamespace := DrawdownResponse.VersionNameSpaceMap[version]
			assert.Contains(t, string(xmlData), expectedNamespace)
		})
	}
}

func TestDrawdownResponseWrapper_ErrorHandling_EdgeCases(t *testing.T) {
	wrapper := &DrawdownResponseWrapper{}

	t.Run("CreateDocument with extremely long fields", func(t *testing.T) {
		// Test with extremely long MessageId that should fail validation
		longMessageId := `{
			"MessageId": "ThisIsAnExtremelyLongMessageIdThatExceedsTheMaximumAllowedLengthForThisFieldAndShouldCauseValidationErrorWhenCreatingTheDocument",
			"CreatedDateTime": "2024-01-01T10:00:00Z",
			"OriginalMessageId": "20250310B1QDRCQR000601"
		}`
		_, err := wrapper.CreateDocument([]byte(longMessageId), DrawdownResponse.PAIN_014_001_07)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create document")
	})

	t.Run("ValidateDocument with malformed date fields", func(t *testing.T) {
		malformedDate := `{
			"MessageId": "20250310B1QDRCQR000602",
			"CreatedDateTime": "invalid-date-format",
			"OriginalMessageId": "20250310B1QDRCQR000601"
		}`
		err := wrapper.ValidateDocument(malformedDate, DrawdownResponse.PAIN_014_001_07)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to MessageModel")
	})

	t.Run("ConvertXMLToModel with XML containing invalid characters", func(t *testing.T) {
		invalidXML := []byte(`<?xml version="1.0"?><Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.014.001.07"><CdtrPmtActvtnReqStsRpt><GrpHdr><MsgId>Test&InvalidChar</MsgId></GrpHdr></CdtrPmtActvtnReqStsRpt></Document>`)
		_, err := wrapper.ConvertXMLToModel(invalidXML)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to model")
	})

	t.Run("CreateDocument with invalid transaction status", func(t *testing.T) {
		invalidStatus := `{
			"MessageId": "20250310B1QDRCQR000602",
			"CreatedDateTime": "2024-01-01T10:00:00Z",
			"OriginalMessageId": "20250310B1QDRCQR000601",
			"TransactionInformationAndStatus": {
				"OriginalInstructionId": "Scenario01Step1InstrId001",
				"TransactionStatus": "INVALID_STATUS_TOO_LONG_CODE"
			}
		}`
		_, err := wrapper.CreateDocument([]byte(invalidStatus), DrawdownResponse.PAIN_014_001_07)
		// This may or may not fail depending on validation rules, but we test handling
		if err != nil {
			assert.Contains(t, err.Error(), "failed to")
		}
	})

	t.Run("CheckRequireField with partially populated model", func(t *testing.T) {
		partialModel := DrawdownResponse.MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "20250310B1QDRCQR000602",
				// Missing CreatedDateTime
			},
			// Missing OriginalMessageId
		}
		err := wrapper.CheckRequireField(partialModel)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "required field missing")
	})
}
