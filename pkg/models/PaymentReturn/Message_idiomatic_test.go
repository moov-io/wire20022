package PaymentReturn

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/civil"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestReadWriteXML tests the XML-first API for reading and writing
func TestReadWriteXML(t *testing.T) {
	// Create a valid message model
	now := time.Now().UTC()
	model := MessageModel{
		PaymentCore: base.PaymentCore{
			MessageHeader: base.MessageHeader{
				MessageId:       "PR20230915001",
				CreatedDateTime: now,
			},
			NumberOfTransactions:  "1",
			SettlementMethod:      models.SettlementCLRG,
			CommonClearingSysCode: models.ClearingSysFDW,
		},
		OriginalMessageId:        "MSG20230914001",
		OriginalMessageNameId:    "pacs.008.001.08",
		OriginalCreationDateTime: now.AddDate(0, 0, -1),
		OriginalInstructionId:    "INST20230914001",
		OriginalEndToEndId:       "E2E20230914001",
		ReturnedInterbankSettlementAmount: models.CurrencyAndAmount{
			Amount:   1000.00,
			Currency: "USD",
		},
		InterbankSettlementDate: fedwire.ISODate(civil.DateOf(now)),
		ReturnedInstructedAmount: models.CurrencyAndAmount{
			Amount:   1000.00,
			Currency: "USD",
		},
		ChargeBearer: models.ChargeBearerSLEV,
		ReturnReasonInformation: models.Reason{
			Reason:         "AC01",
			AdditionalInfo: "Incorrect account number",
		},
		OriginalTransactionRef: models.InstrumentCTRC,
		AgentPair: base.AgentPair{
			InstructingAgent: models.Agent{
				BusinessIdCode:     "BANKUSNY",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "123456789",
				BankName:           "Bank of New York",
			},
			InstructedAgent: models.Agent{
				BusinessIdCode:     "BANKUSLA",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "987654321",
				BankName:           "Bank of Los Angeles",
			},
		},
	}

	// Test WriteXML
	var buf bytes.Buffer
	err := model.WriteXML(&buf, PACS_004_001_08)
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
	require.Contains(t, buf.String(), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	require.Contains(t, buf.String(), "PR20230915001")
	require.Contains(t, buf.String(), "MSG20230914001")

	// Test ReadXML with the generated XML
	var readModel MessageModel
	reader := strings.NewReader(buf.String())
	err = readModel.ReadXML(reader)
	require.NoError(t, err)

	// Verify key fields were preserved
	assert.Equal(t, model.MessageId, readModel.MessageId)
	assert.Equal(t, model.OriginalMessageId, readModel.OriginalMessageId)
	assert.Equal(t, model.OriginalInstructionId, readModel.OriginalInstructionId)
	assert.WithinDuration(t, model.CreatedDateTime, readModel.CreatedDateTime, 5*time.Hour)
}

// TestWriteXMLVersions tests writing XML for different versions
func TestWriteXMLVersions(t *testing.T) {
	versions := []struct {
		name        string
		version     PACS_004_001_VERSION
		hasEnhanced bool
	}{
		{"V2", PACS_004_001_02, false},
		{"V7", PACS_004_001_07, false},
		{"V9", PACS_004_001_09, true},
		{"V13", PACS_004_001_13, true},
	}

	for _, tc := range versions {
		t.Run(tc.name, func(t *testing.T) {
			model := NewMessageForVersion(tc.version)
			model.MessageId = "VERSION_TEST_" + string(tc.version)
			model.CreatedDateTime = time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
			model.NumberOfTransactions = "1"
			model.SettlementMethod = models.SettlementCLRG
			model.CommonClearingSysCode = models.ClearingSysFDW
			model.OriginalMessageId = "ORIG_MSG_" + string(tc.version)
			model.OriginalMessageNameId = "pacs.008.001.08"
			model.OriginalCreationDateTime = time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)
			model.OriginalInstructionId = "INST_" + string(tc.version)
			model.ReturnedInterbankSettlementAmount = models.CurrencyAndAmount{
				Amount:   500.00,
				Currency: "USD",
			}
			model.InterbankSettlementDate = fedwire.ISODate(civil.Date{Year: 2024, Month: 1, Day: 1})
			model.ReturnedInstructedAmount = models.CurrencyAndAmount{
				Amount:   500.00,
				Currency: "USD",
			}
			model.ReturnReasonInformation = models.Reason{Reason: "AC01"}
			model.OriginalTransactionRef = models.InstrumentCTRC
			model.InstructingAgent = models.Agent{
				BusinessIdCode:     "BANKUSNY",
				PaymentSysMemberId: "123456789",
			}
			model.InstructedAgent = models.Agent{
				BusinessIdCode:     "BANKUSLA",
				PaymentSysMemberId: "987654321",
			}

			// Verify version-specific fields are properly initialized
			if tc.hasEnhanced {
				assert.NotNil(t, model.EnhancedTransaction, "EnhancedTransaction should be initialized for %s", tc.version)
				model.EnhancedTransaction.OriginalUETR = "550e8400-e29b-41d4-a716-446655440000"
			} else {
				assert.Nil(t, model.EnhancedTransaction, "EnhancedTransaction should be nil for %s", tc.version)
			}

			var buf bytes.Buffer
			err := model.WriteXML(&buf, tc.version)
			require.NoError(t, err)
			assert.NotEmpty(t, buf.String())
			assert.Contains(t, buf.String(), "<?xml")
			assert.Contains(t, buf.String(), "VERSION_TEST_")

			// Verify namespace is correct for version
			expectedNamespace := VersionNameSpaceMap[tc.version]
			assert.Contains(t, buf.String(), expectedNamespace)
		})
	}
}

// TestValidateForVersion tests version-specific validation
func TestValidateForVersion(t *testing.T) {
	testCases := []struct {
		name    string
		model   MessageModel
		version PACS_004_001_VERSION
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid model for V7",
			model: MessageModel{
				PaymentCore: base.PaymentCore{
					MessageHeader: base.MessageHeader{
						MessageId:       "VALID007",
						CreatedDateTime: time.Now(),
					},
				},
				OriginalMessageId:     "ORIG007",
				OriginalMessageNameId: "pacs.008.001.08",
				OriginalInstructionId: "INST007",
			},
			version: PACS_004_001_07,
			wantErr: false,
		},
		{
			name: "Valid model for V9 with enhanced fields",
			model: MessageModel{
				PaymentCore: base.PaymentCore{
					MessageHeader: base.MessageHeader{
						MessageId:       "VALID009",
						CreatedDateTime: time.Now(),
					},
				},
				OriginalMessageId:     "ORIG009",
				OriginalMessageNameId: "pacs.008.001.08",
				OriginalInstructionId: "INST009",
				EnhancedTransaction: &EnhancedTransactionFields{
					OriginalUETR: "550e8400-e29b-41d4-a716-446655440000",
				},
			},
			version: PACS_004_001_09,
			wantErr: false,
		},
		{
			name: "Missing MessageId",
			model: MessageModel{
				PaymentCore: base.PaymentCore{
					MessageHeader: base.MessageHeader{
						CreatedDateTime: time.Now(),
					},
				},
				OriginalMessageId:     "ORIG001",
				OriginalInstructionId: "INST001",
			},
			version: PACS_004_001_07,
			wantErr: true,
			errMsg:  "MessageId is required",
		},
		{
			name: "Missing CreatedDateTime",
			model: MessageModel{
				PaymentCore: base.PaymentCore{
					MessageHeader: base.MessageHeader{
						MessageId: "INVALID002",
					},
				},
				OriginalMessageId:     "ORIG002",
				OriginalInstructionId: "INST002",
			},
			version: PACS_004_001_07,
			wantErr: true,
			errMsg:  "CreatedDateTime is required",
		},
		{
			name: "V9 missing EnhancedTransaction",
			model: MessageModel{
				PaymentCore: base.PaymentCore{
					MessageHeader: base.MessageHeader{
						MessageId:       "INVALID009",
						CreatedDateTime: time.Now(),
					},
				},
				OriginalMessageId:     "ORIG009",
				OriginalMessageNameId: "pacs.008.001.08",
				OriginalInstructionId: "INST009",
				// Missing EnhancedTransaction for V9+
			},
			version: PACS_004_001_09,
			wantErr: true,
			errMsg:  "EnhancedTransactionFields required for version",
		},
		{
			name: "V9 with empty OriginalUETR",
			model: MessageModel{
				PaymentCore: base.PaymentCore{
					MessageHeader: base.MessageHeader{
						MessageId:       "INVALID010",
						CreatedDateTime: time.Now(),
					},
				},
				OriginalMessageId:     "ORIG010",
				OriginalMessageNameId: "pacs.008.001.08",
				OriginalInstructionId: "INST010",
				EnhancedTransaction:   &EnhancedTransactionFields{
					// Empty OriginalUETR for V9+
				},
			},
			version: PACS_004_001_09,
			wantErr: true,
			errMsg:  "OriginalUETR is required for versions V9+",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.model.ValidateForVersion(tc.version)
			if tc.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestValidateCoreFields tests core field validation directly
func TestValidateCoreFields(t *testing.T) {
	t.Run("Valid core fields", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					MessageId:       "CORE001",
					CreatedDateTime: time.Now(),
				},
			},
			OriginalMessageId:     "ORIG001",
			OriginalInstructionId: "INST001",
		}
		err := model.validateCoreFields()
		assert.NoError(t, err)
	})

	t.Run("Empty MessageId", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					CreatedDateTime: time.Now(),
				},
			},
			OriginalMessageId:     "ORIG002",
			OriginalInstructionId: "INST002",
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "MessageId is required", err.Error())
	})

	t.Run("Zero CreatedDateTime", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					MessageId: "CORE003",
				},
			},
			OriginalMessageId:     "ORIG003",
			OriginalInstructionId: "INST003",
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "CreatedDateTime is required", err.Error())
	})

	t.Run("Empty OriginalMessageId", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					MessageId:       "CORE004",
					CreatedDateTime: time.Now(),
				},
			},
			OriginalInstructionId: "INST004",
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "OriginalMessageId is required", err.Error())
	})

	t.Run("Empty OriginalInstructionId", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					MessageId:       "CORE005",
					CreatedDateTime: time.Now(),
				},
			},
			OriginalMessageId: "ORIG005",
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "OriginalInstructionId is required", err.Error())
	})
}

// TestGetVersionCapabilities tests version capability detection
func TestGetVersionCapabilities(t *testing.T) {
	testCases := []struct {
		name         string
		version      PACS_004_001_VERSION
		expectedCaps map[string]bool
	}{
		{
			name:    "V7 - no enhanced fields",
			version: PACS_004_001_07,
			expectedCaps: map[string]bool{
				"EnhancedTransaction": false,
			},
		},
		{
			name:    "V9 - enhanced fields",
			version: PACS_004_001_09,
			expectedCaps: map[string]bool{
				"EnhancedTransaction": true,
			},
		},
		{
			name:    "V13 - enhanced fields",
			version: PACS_004_001_13,
			expectedCaps: map[string]bool{
				"EnhancedTransaction": true,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			model := NewMessageForVersion(tc.version)
			caps := model.GetVersionCapabilities()
			assert.Equal(t, tc.expectedCaps, caps)
		})
	}
}

// TestNewMessageForVersion tests version-specific initialization
func TestNewMessageForVersion(t *testing.T) {
	versions := []struct {
		version       PACS_004_001_VERSION
		hasEnhancedTx bool
	}{
		{PACS_004_001_02, false},
		{PACS_004_001_07, false},
		{PACS_004_001_09, true},
		{PACS_004_001_13, true},
	}

	for _, v := range versions {
		t.Run(string(v.version), func(t *testing.T) {
			model := NewMessageForVersion(v.version)
			assert.NotNil(t, model)

			// Check base fields are initialized to zero values
			assert.Empty(t, model.MessageId)
			assert.Empty(t, model.OriginalMessageId)
			assert.True(t, model.CreatedDateTime.IsZero())

			// Check version-specific field initialization
			if v.hasEnhancedTx {
				assert.NotNil(t, model.EnhancedTransaction)
			} else {
				assert.Nil(t, model.EnhancedTransaction)
			}
		})
	}
}

// TestCheckRequiredFields tests the required field validation helper
func TestCheckRequiredFields(t *testing.T) {
	t.Run("All required fields present", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					MessageId:       "REQ001",
					CreatedDateTime: time.Now(),
				},
				NumberOfTransactions:  "1",
				SettlementMethod:      models.SettlementCLRG,
				CommonClearingSysCode: models.ClearingSysFDW,
			},
			OriginalMessageId:                 "ORIG_REQ001",
			OriginalMessageNameId:             "pacs.008.001.08",
			OriginalCreationDateTime:          time.Now().AddDate(0, 0, -1),
			OriginalInstructionId:             "INST_REQ001",
			OriginalEndToEndId:                "E2E_REQ001",
			ReturnedInterbankSettlementAmount: models.CurrencyAndAmount{Amount: 100.00, Currency: "USD"},
			InterbankSettlementDate:           fedwire.ISODate(civil.DateOf(time.Now())),
			ReturnedInstructedAmount:          models.CurrencyAndAmount{Amount: 100.00, Currency: "USD"},
			ChargeBearer:                      models.ChargeBearerSLEV,
			ReturnReasonInformation:           models.Reason{Reason: "AC01"},
			OriginalTransactionRef:            models.InstrumentCTRC,
			AgentPair: base.AgentPair{
				InstructingAgent: models.Agent{
					BusinessIdCode:     "BANKUSNY",
					PaymentSysCode:     "USABA",
					PaymentSysMemberId: "123456789",
				},
				InstructedAgent: models.Agent{
					BusinessIdCode:     "BANKUSLA",
					PaymentSysCode:     "USABA",
					PaymentSysMemberId: "987654321",
				},
			},
		}
		err := CheckRequiredFields(model)
		assert.NoError(t, err)
	})

	t.Run("Missing required fields", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					MessageId: "REQ002",
					// Missing CreatedDateTime and other required fields
				},
			},
		}
		err := CheckRequiredFields(model)
		require.Error(t, err)
		// The error should mention missing required fields
		assert.Contains(t, err.Error(), "required")
	})
}

// TestJSONMarshaling tests JSON serialization
func TestJSONMarshaling(t *testing.T) {
	original := MessageModel{
		PaymentCore: base.PaymentCore{
			MessageHeader: base.MessageHeader{
				MessageId:       "JSON001",
				CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			},
			NumberOfTransactions:  "1",
			SettlementMethod:      models.SettlementCLRG,
			CommonClearingSysCode: models.ClearingSysFDW,
		},
		OriginalMessageId:        "JSON_ORIG001",
		OriginalMessageNameId:    "pacs.008.001.08",
		OriginalCreationDateTime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		OriginalInstructionId:    "JSON_INST001",
		ReturnedInterbankSettlementAmount: models.CurrencyAndAmount{
			Amount:   750.50,
			Currency: "USD",
		},
		InterbankSettlementDate: fedwire.ISODate(civil.Date{Year: 2024, Month: 1, Day: 1}),
		ReturnedInstructedAmount: models.CurrencyAndAmount{
			Amount:   750.50,
			Currency: "USD",
		},
		ReturnReasonInformation: models.Reason{
			Reason:         "AC01",
			AdditionalInfo: "Incorrect account number",
		},
		OriginalTransactionRef: models.InstrumentCTRC,
		AgentPair: base.AgentPair{
			InstructingAgent: models.Agent{
				BusinessIdCode:     "BANKUSNY",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "123456789",
				BankName:           "Bank of New York",
			},
		},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	require.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Unmarshal back
	var decoded MessageModel
	err = json.Unmarshal(jsonData, &decoded)
	require.NoError(t, err)

	// Verify fields
	assert.Equal(t, original.MessageId, decoded.MessageId)
	assert.Equal(t, original.OriginalMessageId, decoded.OriginalMessageId)
	assert.Equal(t, original.OriginalInstructionId, decoded.OriginalInstructionId)
	assert.Equal(t, original.CreatedDateTime.UTC(), decoded.CreatedDateTime.UTC())
	assert.Equal(t, original.ReturnedInterbankSettlementAmount.Amount, decoded.ReturnedInterbankSettlementAmount.Amount)
}

// TestJSONUnmarshalWithEnhancedFields tests JSON unmarshaling with enhanced transaction fields
func TestJSONUnmarshalWithEnhancedFields(t *testing.T) {
	jsonData := `{
		"messageId": "JSON002",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"originalMessageId": "ORIG002",
		"originalInstructionId": "INST002",
		"enhancedTransaction": {
			"originalUETR": "550e8400-e29b-41d4-a716-446655440000"
		}
	}`

	var model MessageModel
	err := json.Unmarshal([]byte(jsonData), &model)
	require.NoError(t, err)

	// Verify enhanced transaction fields were properly initialized
	assert.NotNil(t, model.EnhancedTransaction)
	assert.Equal(t, "550e8400-e29b-41d4-a716-446655440000", model.EnhancedTransaction.OriginalUETR)
}

// TestParseXMLWithInvalidData tests ParseXML error handling
func TestParseXMLWithInvalidData(t *testing.T) {
	testCases := []struct {
		name    string
		xmlData string
		wantErr bool
	}{
		{
			name:    "Empty XML",
			xmlData: "",
			wantErr: true,
		},
		{
			name:    "Invalid XML",
			xmlData: "not xml at all",
			wantErr: true,
		},
		{
			name:    "Wrong namespace",
			xmlData: `<?xml version="1.0"?><Document xmlns="wrong:namespace"><PmtRtr></PmtRtr></Document>`,
			wantErr: true,
		},
		{
			name: "Missing required fields",
			xmlData: `<?xml version="1.0"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.004.001.08">
	<PmtRtr>
		<GrpHdr>
			<MsgId>TEST001</MsgId>
		</GrpHdr>
	</PmtRtr>
</Document>`,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseXML([]byte(tc.xmlData))
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestReadXMLWithErrors tests ReadXML error handling
func TestReadXMLWithErrors(t *testing.T) {
	t.Run("Invalid reader", func(t *testing.T) {
		var model MessageModel
		err := model.ReadXML(strings.NewReader(""))
		assert.Error(t, err)
	})

	t.Run("Malformed XML", func(t *testing.T) {
		var model MessageModel
		err := model.ReadXML(strings.NewReader("<invalid>"))
		assert.Error(t, err)
	})
}

// TestWriteXMLWithInvalidModel tests WriteXML with invalid data
func TestWriteXMLWithInvalidModel(t *testing.T) {
	// Model missing required fields
	model := MessageModel{}

	var buf bytes.Buffer
	err := model.WriteXML(&buf)
	assert.Error(t, err)
}

// TestDocumentWithValidation tests DocumentWith with validation
func TestDocumentWithValidation(t *testing.T) {
	t.Run("Valid model creates document", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					MessageId:       "DOC001",
					CreatedDateTime: time.Now(),
				},
				NumberOfTransactions:  "1",
				SettlementMethod:      models.SettlementCLRG,
				CommonClearingSysCode: models.ClearingSysFDW,
			},
			OriginalMessageId:                 "DOC_ORIG001",
			OriginalMessageNameId:             "pacs.008.001.08",
			OriginalCreationDateTime:          time.Now().AddDate(0, 0, -1),
			OriginalInstructionId:             "DOC_INST001",
			OriginalEndToEndId:                "DOC_E2E001",
			ReturnedInterbankSettlementAmount: models.CurrencyAndAmount{Amount: 200.00, Currency: "USD"},
			InterbankSettlementDate:           fedwire.ISODate(civil.DateOf(time.Now())),
			ReturnedInstructedAmount:          models.CurrencyAndAmount{Amount: 200.00, Currency: "USD"},
			ChargeBearer:                      models.ChargeBearerSLEV,
			ReturnReasonInformation:           models.Reason{Reason: "AC01"},
			OriginalTransactionRef:            models.InstrumentCTRC,
			AgentPair: base.AgentPair{
				InstructingAgent: models.Agent{
					BusinessIdCode:     "BANKUSNY",
					PaymentSysCode:     "USABA",
					PaymentSysMemberId: "123456789",
				},
				InstructedAgent: models.Agent{
					BusinessIdCode:     "BANKUSLA",
					PaymentSysCode:     "USABA",
					PaymentSysMemberId: "987654321",
				},
			},
		}

		doc, err := DocumentWith(model, PACS_004_001_07)
		require.NoError(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("Invalid model fails validation", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					MessageId: "DOC002",
					// Missing required fields
				},
			},
		}

		doc, err := DocumentWith(model, PACS_004_001_07)
		require.Error(t, err)
		assert.Nil(t, doc)
		assert.Contains(t, err.Error(), "required")
	})
}

// TestEnhancedFieldValidation tests enhanced field validation
func TestEnhancedFieldValidation(t *testing.T) {
	t.Run("EnhancedTransactionFields validation with OriginalUETR", func(t *testing.T) {
		fields := &EnhancedTransactionFields{
			OriginalUETR: "550e8400-e29b-41d4-a716-446655440000",
		}
		err := fields.Validate()
		assert.NoError(t, err)
	})

	t.Run("EnhancedTransactionFields validation without OriginalUETR", func(t *testing.T) {
		fields := &EnhancedTransactionFields{}
		err := fields.Validate()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "OriginalUETR is required for versions V9+")
	})
}

// TestAgentPairStructures tests the agent pair structures in the message
func TestAgentPairStructures(t *testing.T) {
	t.Run("Valid agent pairs", func(t *testing.T) {
		model := MessageModel{
			PaymentCore: base.PaymentCore{
				MessageHeader: base.MessageHeader{
					MessageId:       "AGENT001",
					CreatedDateTime: time.Now(),
				},
			},
			OriginalMessageId:     "ORIG_AGENT001",
			OriginalInstructionId: "INST_AGENT001",
			AgentPair: base.AgentPair{
				InstructingAgent: models.Agent{
					BusinessIdCode:     "BANKUSNY",
					PaymentSysCode:     "USABA",
					PaymentSysMemberId: "123456789",
					BankName:           "Bank of New York",
				},
				InstructedAgent: models.Agent{
					BusinessIdCode:     "BANKUSLA",
					PaymentSysCode:     "USABA",
					PaymentSysMemberId: "987654321",
					BankName:           "Bank of Los Angeles",
				},
			},
		}

		// Verify agent fields
		assert.Equal(t, "BANKUSNY", model.InstructingAgent.BusinessIdCode)
		assert.Equal(t, "123456789", model.InstructingAgent.PaymentSysMemberId)
		assert.Equal(t, "BANKUSLA", model.InstructedAgent.BusinessIdCode)
		assert.Equal(t, "987654321", model.InstructedAgent.PaymentSysMemberId)
	})
}

// TestCurrencyAndAmountFields tests currency and amount structures
func TestCurrencyAndAmountFields(t *testing.T) {
	t.Run("Valid currency and amounts", func(t *testing.T) {
		model := MessageModel{
			ReturnedInterbankSettlementAmount: models.CurrencyAndAmount{
				Amount:   1250.75,
				Currency: "USD",
			},
			ReturnedInstructedAmount: models.CurrencyAndAmount{
				Amount:   1250.75,
				Currency: "USD",
			},
		}

		assert.Equal(t, 1250.75, model.ReturnedInterbankSettlementAmount.Amount)
		assert.Equal(t, "USD", model.ReturnedInterbankSettlementAmount.Currency)
		assert.Equal(t, 1250.75, model.ReturnedInstructedAmount.Amount)
		assert.Equal(t, "USD", model.ReturnedInstructedAmount.Currency)
	})

	t.Run("Different currencies", func(t *testing.T) {
		currencies := []string{"USD", "EUR", "GBP", "JPY"}

		for _, currency := range currencies {
			model := MessageModel{
				ReturnedInterbankSettlementAmount: models.CurrencyAndAmount{
					Amount:   100.00,
					Currency: currency,
				},
			}
			assert.Equal(t, currency, model.ReturnedInterbankSettlementAmount.Currency)
		}
	})
}

// TestReasonInformation tests reason and additional information structures
func TestReasonInformation(t *testing.T) {
	t.Run("Valid reason information", func(t *testing.T) {
		model := MessageModel{
			ReturnReasonInformation: models.Reason{
				Reason:         "AC01",
				AdditionalInfo: "Incorrect account number specified",
			},
		}

		assert.Equal(t, "AC01", model.ReturnReasonInformation.Reason)
		assert.Equal(t, "Incorrect account number specified", model.ReturnReasonInformation.AdditionalInfo)
	})

	t.Run("Different reason codes", func(t *testing.T) {
		reasonCodes := []string{"AC01", "AC02", "AC03", "AC04", "AC06"}

		for _, reasonCode := range reasonCodes {
			model := MessageModel{
				ReturnReasonInformation: models.Reason{
					Reason: reasonCode,
				},
			}
			assert.Equal(t, reasonCode, model.ReturnReasonInformation.Reason)
		}
	})
}

// TestChargeBearerTypes tests different charge bearer types
func TestChargeBearerTypes(t *testing.T) {
	t.Run("Valid charge bearer types", func(t *testing.T) {
		chargeBearers := []models.ChargeBearerType{
			models.ChargeBearerSLEV,
			models.ChargeBearerSHAR,
			models.ChargeBearerCREDIT,
			models.ChargeBearerDEBT,
		}

		for _, chargeBearer := range chargeBearers {
			model := MessageModel{
				ChargeBearer: chargeBearer,
			}
			assert.Equal(t, chargeBearer, model.ChargeBearer)
		}
	})
}

// BenchmarkWriteXML benchmarks XML writing performance
func BenchmarkWriteXML(b *testing.B) {
	model := MessageModel{
		PaymentCore: base.PaymentCore{
			MessageHeader: base.MessageHeader{
				MessageId:       "BENCH001",
				CreatedDateTime: time.Now(),
			},
		},
		OriginalMessageId:                 "BENCH_ORIG001",
		OriginalMessageNameId:             "pacs.008.001.08",
		OriginalCreationDateTime:          time.Now(),
		OriginalInstructionId:             "BENCH_INST001",
		ReturnedInterbankSettlementAmount: models.CurrencyAndAmount{Amount: 500.00, Currency: "USD"},
		ChargeBearer:                      models.ChargeBearerSLEV,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_ = model.WriteXML(&buf, PACS_004_001_08)
	}
}

// BenchmarkParseXML benchmarks XML parsing performance
func BenchmarkParseXML(b *testing.B) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.004.001.08">
	<PmtRtr>
		<GrpHdr>
			<MsgId>BENCH001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
			<InstgAgt>
				<FinInstnId>
					<ClrSysMmbId>
						<MmbId>123456789</MmbId>
					</ClrSysMmbId>
				</FinInstnId>
			</InstgAgt>
			<InstdAgt>
				<FinInstnId>
					<ClrSysMmbId>
						<MmbId>987654321</MmbId>
					</ClrSysMmbId>
				</FinInstnId>
			</InstdAgt>
		</GrpHdr>
		<OrgnlGrpInfAndSts>
			<OrgnlMsgId>BENCH_ORIG001</OrgnlMsgId>
			<OrgnlMsgNmId>pacs.008.001.08</OrgnlMsgNmId>
			<OrgnlCreDtTm>2024-01-01T09:00:00Z</OrgnlCreDtTm>
		</OrgnlGrpInfAndSts>
		<TxInfAndSts>
			<OrgnlInstrId>BENCH_INST001</OrgnlInstrId>
			<OrgnlEndToEndId>BENCH_E2E001</OrgnlEndToEndId>
		</TxInfAndSts>
	</PmtRtr>
</Document>`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseXML(xmlData)
	}
}

// TestBuildPartyHelper tests the party helper builder
func TestBuildPartyHelper(t *testing.T) {
	helper := BuildPartyHelper()

	// Test field initialization
	assert.Equal(t, "Name", helper.Name.Title)
	assert.Contains(t, helper.Name.Type, "Max140Text")
	assert.Contains(t, helper.Name.Documentation, "party is known")

	// Test nested address helper
	assert.Equal(t, "Building Name", helper.Address.BuildingName.Title)
	assert.Equal(t, "Street Name", helper.Address.StreetName.Title)
}

// TestBuildReturnChainHelper tests the return chain helper builder
func TestBuildReturnChainHelper(t *testing.T) {
	helper := BuildReturnChainHelper()

	// Test debtor and creditor helpers
	assert.Equal(t, "Name", helper.Debtor.Name.Title)
	assert.Equal(t, "Name", helper.Creditor.Name.Title)

	// Test ID fields
	assert.Equal(t, "Debtor Other Type Id", helper.DebtorOtherTypeId.Title)
	assert.Equal(t, "Creditor Account Other Type Id", helper.CreditorAccountOtherTypeId.Title)

	// Test nested agent helpers
	assert.Equal(t, "Member Identification", helper.DebtorAgent.PaymentSysMemberId.Title)
	assert.Equal(t, "Member Identification", helper.CreditorAgent.PaymentSysMemberId.Title)
}

// TestBuildMessageHelper tests the message helper builder
func TestBuildMessageHelper(t *testing.T) {
	helper := BuildMessageHelper()

	// Test basic field initialization
	assert.Equal(t, "Message Id", helper.MessageId.Title)
	assert.Equal(t, "Created Date Time", helper.CreatedDateTime.Title)
	assert.Equal(t, "Original Message Id", helper.OriginalMessageId.Title)
	assert.Equal(t, "Original Instruction Id", helper.OriginalInstructionId.Title)
	assert.Equal(t, "Original End To End Id", helper.OriginalEndToEndId.Title)

	// Test currency and amount helpers
	assert.Equal(t, "Amount", helper.ReturnedInterbankSettlementAmount.Amount.Title)
	assert.Equal(t, "Currency", helper.ReturnedInterbankSettlementAmount.Currency.Title)

	// Test agent helpers are properly initialized
	assert.Equal(t, "Member Identification", helper.InstructingAgent.PaymentSysMemberId.Title)
	assert.Equal(t, "Member Identification", helper.InstructedAgent.PaymentSysMemberId.Title)

	// Test return chain helper
	assert.Equal(t, "Name", helper.RtrChain.Debtor.Name.Title)
	assert.Equal(t, "Name", helper.RtrChain.Creditor.Name.Title)

	// Test type information
	assert.Contains(t, helper.MessageId.Type, "Max35Text")
	assert.Contains(t, helper.CreatedDateTime.Type, "ISODateTime")
	assert.Contains(t, helper.OriginalMessageId.Type, "Max35Text")

	// Test documentation
	assert.Contains(t, helper.MessageId.Documentation, "Point to point reference")
	assert.Contains(t, helper.OriginalMessageId.Documentation, "Point to point reference")
	assert.NotEmpty(t, helper.ReturnedInterbankSettlementAmount.Amount.Documentation)
}

// TestMessageHelperCompleteness tests that all fields have documentation
func TestMessageHelperCompleteness(t *testing.T) {
	helper := BuildMessageHelper()

	// Verify all fields have non-empty documentation
	assert.NotEmpty(t, helper.MessageId.Documentation)
	assert.NotEmpty(t, helper.CreatedDateTime.Documentation)
	assert.NotEmpty(t, helper.OriginalMessageId.Documentation)
	assert.NotEmpty(t, helper.OriginalMessageNameId.Documentation)
	assert.NotEmpty(t, helper.OriginalCreationDateTime.Documentation)
	assert.NotEmpty(t, helper.OriginalInstructionId.Documentation)
	assert.NotEmpty(t, helper.OriginalEndToEndId.Documentation)
	assert.NotEmpty(t, helper.ReturnedInterbankSettlementAmount.Amount.Documentation)
	assert.NotEmpty(t, helper.ReturnedInstructedAmount.Amount.Documentation)
	assert.NotEmpty(t, helper.ChargeBearer.Documentation)

	// Verify all fields have non-empty types
	assert.NotEmpty(t, helper.MessageId.Type)
	assert.NotEmpty(t, helper.CreatedDateTime.Type)
	assert.NotEmpty(t, helper.OriginalMessageId.Type)
	assert.NotEmpty(t, helper.OriginalMessageNameId.Type)
	assert.NotEmpty(t, helper.OriginalCreationDateTime.Type)
	assert.NotEmpty(t, helper.OriginalInstructionId.Type)
	assert.NotEmpty(t, helper.OriginalEndToEndId.Type)

	// Verify agent helpers are complete
	assert.NotEmpty(t, helper.InstructingAgent.BusinessIdCode.Documentation)
	assert.NotEmpty(t, helper.InstructedAgent.BusinessIdCode.Documentation)

	// Verify return chain helper is complete
	assert.NotEmpty(t, helper.RtrChain.Debtor.Name.Documentation)
	assert.NotEmpty(t, helper.RtrChain.Creditor.Name.Documentation)
}
