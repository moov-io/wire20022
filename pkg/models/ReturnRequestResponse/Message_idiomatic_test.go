package ReturnRequestResponse

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestReadWriteXML tests the XML-first API for reading and writing
func TestReadWriteXML(t *testing.T) {
	// Create a valid message model
	now := time.Now().UTC()
	model := MessageModel{
		AssignmentId:              "RRR20230915001",
		AssignmentCreateTime:      now,
		ResolvedCaseId:            "CASE20230915001",
		OriginalMessageId:         "MSG20230914001",
		OriginalMessageNameId:     "pacs.008.001.08",
		OriginalMessageCreateTime: now.AddDate(0, 0, -1),
		OriginalInstructionId:     "INST20230914001",
		OriginalEndToEndId:        "E2E20230914001",
		Assigner: models.Agent{
			BusinessIdCode:     "BANKUSNY",
			PaymentSysCode:     "USABA",
			PaymentSysMemberId: "123456789",
			BankName:           "Bank of New York",
		},
		Assignee: models.Agent{
			BusinessIdCode:     "BANKUSLA",
			PaymentSysCode:     "USABA",
			PaymentSysMemberId: "987654321",
			BankName:           "Bank of Los Angeles",
		},
		Creator: models.Agent{
			BusinessIdCode:     "BANKUSNY",
			PaymentSysCode:     "USABA",
			PaymentSysMemberId: "123456789",
			BankName:           "Bank of New York",
		},
		Status: "ACCP",
		CancellationStatusReasonInfo: models.Reason{
			Reason: "AC01",
			AdditionalInfo: "Return accepted",
		},
	}

	// Test WriteXML
	var buf bytes.Buffer
	err := model.WriteXML(&buf, CAMT_029_001_08)
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
	require.Contains(t, buf.String(), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	require.Contains(t, buf.String(), "RRR20230915001")
	require.Contains(t, buf.String(), "CASE20230915001")

	// Test ReadXML with the generated XML
	var readModel MessageModel
	reader := strings.NewReader(buf.String())
	err = readModel.ReadXML(reader)
	require.NoError(t, err)

	// Verify key fields were preserved
	assert.Equal(t, model.AssignmentId, readModel.AssignmentId)
	assert.Equal(t, model.ResolvedCaseId, readModel.ResolvedCaseId)
	assert.Equal(t, model.OriginalMessageId, readModel.OriginalMessageId)
	assert.WithinDuration(t, model.AssignmentCreateTime, readModel.AssignmentCreateTime, 5*time.Hour)
}

// TestWriteXMLVersions tests writing XML for different versions
func TestWriteXMLVersions(t *testing.T) {
	versions := []struct {
		name                string
		version             CAMT_029_001_VERSION
		hasEnhancedTx       bool
		hasAddressEnhance   bool
	}{
		{"V3", CAMT_029_001_03, false, false},
		{"V8", CAMT_029_001_08, false, false},
		{"V9", CAMT_029_001_09, true, true},
		{"V12", CAMT_029_001_12, true, true},
	}

	for _, tc := range versions {
		t.Run(tc.name, func(t *testing.T) {
			model := NewMessageForVersion(tc.version)
			model.AssignmentId = "VERSION_TEST_" + string(tc.version)
			model.AssignmentCreateTime = time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
			model.ResolvedCaseId = "CASE_" + string(tc.version)
			model.OriginalMessageId = "ORIG_MSG_" + string(tc.version)
			model.OriginalMessageNameId = "pacs.008.001.08"
			model.OriginalMessageCreateTime = time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)
			model.Status = "ACCP"

			// Verify version-specific fields are properly initialized
			if tc.hasEnhancedTx {
				assert.NotNil(t, model.EnhancedTransaction, "EnhancedTransaction should be initialized for %s", tc.version)
			} else {
				assert.Nil(t, model.EnhancedTransaction, "EnhancedTransaction should be nil for %s", tc.version)
			}

			if tc.hasAddressEnhance {
				assert.NotNil(t, model.AddressEnhancement, "AddressEnhancement should be initialized for %s", tc.version)
			} else {
				assert.Nil(t, model.AddressEnhancement, "AddressEnhancement should be nil for %s", tc.version)
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
		version CAMT_029_001_VERSION
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid model for V3",
			model: MessageModel{
				AssignmentId:              "VALID003",
				AssignmentCreateTime:      time.Now(),
				ResolvedCaseId:            "CASE003",
				OriginalMessageId:         "ORIG003",
				OriginalMessageNameId:     "pacs.008.001.08",
				OriginalMessageCreateTime: time.Now().AddDate(0, 0, -1),
				Status:                    "ACCP",
			},
			version: CAMT_029_001_03,
			wantErr: false,
		},
		{
			name: "Valid model for V9 with enhanced fields",
			model: MessageModel{
				AssignmentId:              "VALID009",
				AssignmentCreateTime:      time.Now(),
				ResolvedCaseId:            "CASE009",
				OriginalMessageId:         "ORIG009",
				OriginalMessageNameId:     "pacs.008.001.08",
				OriginalMessageCreateTime: time.Now().AddDate(0, 0, -1),
				Status:                    "ACCP",
				EnhancedTransaction: &EnhancedTransactionFields{
					OriginalUETR: "550e8400-e29b-41d4-a716-446655440000",
				},
				AddressEnhancement: &AddressEnhancementFields{},
			},
			version: CAMT_029_001_09,
			wantErr: false,
		},
		{
			name: "Missing AssignmentId",
			model: MessageModel{
				AssignmentCreateTime:      time.Now(),
				ResolvedCaseId:            "CASE001",
				OriginalMessageId:         "ORIG001",
				OriginalMessageNameId:     "pacs.008.001.08",
				OriginalMessageCreateTime: time.Now(),
			},
			version: CAMT_029_001_03,
			wantErr: true,
			errMsg:  "AssignmentId is required",
		},
		{
			name: "Missing AssignmentCreateTime",
			model: MessageModel{
				AssignmentId:              "INVALID002",
				ResolvedCaseId:            "CASE002",
				OriginalMessageId:         "ORIG002",
				OriginalMessageNameId:     "pacs.008.001.08",
				OriginalMessageCreateTime: time.Now(),
			},
			version: CAMT_029_001_03,
			wantErr: true,
			errMsg:  "AssignmentCreateTime is required",
		},
		{
			name: "Missing ResolvedCaseId",
			model: MessageModel{
				AssignmentId:              "INVALID003",
				AssignmentCreateTime:      time.Now(),
				OriginalMessageId:         "ORIG003",
				OriginalMessageNameId:     "pacs.008.001.08",
				OriginalMessageCreateTime: time.Now(),
			},
			version: CAMT_029_001_03,
			wantErr: true,
			errMsg:  "ResolvedCaseId is required",
		},
		{
			name: "V9 missing EnhancedTransaction",
			model: MessageModel{
				AssignmentId:              "INVALID009",
				AssignmentCreateTime:      time.Now(),
				ResolvedCaseId:            "CASE009",
				OriginalMessageId:         "ORIG009",
				OriginalMessageNameId:     "pacs.008.001.08",
				OriginalMessageCreateTime: time.Now(),
				Status:                    "ACCP",
				// Missing EnhancedTransaction for V9+
			},
			version: CAMT_029_001_09,
			wantErr: true,
			errMsg:  "EnhancedTransactionFields required for version",
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
			AssignmentId:              "CORE001",
			AssignmentCreateTime:      time.Now(),
			ResolvedCaseId:            "CASE001",
			OriginalMessageId:         "ORIG001",
			OriginalMessageNameId:     "pacs.008.001.08",
			OriginalMessageCreateTime: time.Now(),
		}
		err := model.validateCoreFields()
		assert.NoError(t, err)
	})

	t.Run("Empty AssignmentId", func(t *testing.T) {
		model := MessageModel{
			AssignmentCreateTime:      time.Now(),
			ResolvedCaseId:            "CASE002",
			OriginalMessageId:         "ORIG002",
			OriginalMessageNameId:     "pacs.008.001.08",
			OriginalMessageCreateTime: time.Now(),
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "AssignmentId is required", err.Error())
	})

	t.Run("Zero AssignmentCreateTime", func(t *testing.T) {
		model := MessageModel{
			AssignmentId:              "CORE003",
			ResolvedCaseId:            "CASE003",
			OriginalMessageId:         "ORIG003",
			OriginalMessageNameId:     "pacs.008.001.08",
			OriginalMessageCreateTime: time.Now(),
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "AssignmentCreateTime is required", err.Error())
	})

	t.Run("Empty OriginalMessageNameId", func(t *testing.T) {
		model := MessageModel{
			AssignmentId:              "CORE004",
			AssignmentCreateTime:      time.Now(),
			ResolvedCaseId:            "CASE004",
			OriginalMessageId:         "ORIG004",
			OriginalMessageCreateTime: time.Now(),
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "OriginalMessageNameId is required", err.Error())
	})
}

// TestGetVersionCapabilities tests version capability detection
func TestGetVersionCapabilities(t *testing.T) {
	testCases := []struct {
		name         string
		version      CAMT_029_001_VERSION
		expectedCaps map[string]bool
	}{
		{
			name:    "V3 - no enhanced fields",
			version: CAMT_029_001_03,
			expectedCaps: map[string]bool{
				"EnhancedTransaction": false,
				"AddressEnhancement":  false,
			},
		},
		{
			name:    "V9 - enhanced fields",
			version: CAMT_029_001_09,
			expectedCaps: map[string]bool{
				"EnhancedTransaction": true,
				"AddressEnhancement":  true,
			},
		},
		{
			name:    "V12 - enhanced fields",
			version: CAMT_029_001_12,
			expectedCaps: map[string]bool{
				"EnhancedTransaction": true,
				"AddressEnhancement":  true,
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
		version              CAMT_029_001_VERSION
		hasEnhancedTx        bool
		hasAddressEnhancement bool
	}{
		{CAMT_029_001_03, false, false},
		{CAMT_029_001_08, false, false},
		{CAMT_029_001_09, true, true},
		{CAMT_029_001_12, true, true},
	}

	for _, v := range versions {
		t.Run(string(v.version), func(t *testing.T) {
			model := NewMessageForVersion(v.version)
			assert.NotNil(t, model)

			// Check base fields are initialized to zero values
			assert.Empty(t, model.AssignmentId)
			assert.Empty(t, model.ResolvedCaseId)
			assert.True(t, model.AssignmentCreateTime.IsZero())

			// Check version-specific field initialization
			if v.hasEnhancedTx {
				assert.NotNil(t, model.EnhancedTransaction)
			} else {
				assert.Nil(t, model.EnhancedTransaction)
			}

			if v.hasAddressEnhancement {
				assert.NotNil(t, model.AddressEnhancement)
			} else {
				assert.Nil(t, model.AddressEnhancement)
			}
		})
	}
}

// TestCheckRequiredFields tests the required field validation helper
func TestCheckRequiredFields(t *testing.T) {
	t.Run("All required fields present", func(t *testing.T) {
		model := MessageModel{
			AssignmentId:              "REQ001",
			AssignmentCreateTime:      time.Now(),
			ResolvedCaseId:            "CASE_REQ001",
			OriginalMessageId:         "ORIG_REQ001",
			OriginalMessageNameId:     "pacs.008.001.08",
			OriginalMessageCreateTime: time.Now().AddDate(0, 0, -1),
			OriginalInstructionId:     "INST_REQ001",
			OriginalEndToEndId:        "E2E_REQ001",
			Assigner: models.Agent{
				BusinessIdCode:     "BANKUSNY",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "123456789",
			},
			Assignee: models.Agent{
				BusinessIdCode:     "BANKUSLA",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "987654321",
			},
			Creator: models.Agent{
				BusinessIdCode:     "BANKUSNY",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "123456789",
			},
			Status: "ACCP",
			CancellationStatusReasonInfo: models.Reason{
				Reason: "AC01",
			},
		}
		err := CheckRequiredFields(model)
		assert.NoError(t, err)
	})

	t.Run("Missing required fields", func(t *testing.T) {
		model := MessageModel{
			AssignmentId: "REQ002",
			// Missing AssignmentCreateTime and other required fields
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
		AssignmentId:              "JSON001",
		AssignmentCreateTime:      time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		ResolvedCaseId:            "JSON_CASE001",
		OriginalMessageId:         "JSON_ORIG001",
		OriginalMessageNameId:     "pacs.008.001.08",
		OriginalMessageCreateTime: time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC),
		Status:                    "ACCP",
		Assigner: models.Agent{
			BusinessIdCode:     "BANKUSNY",
			PaymentSysCode:     "USABA",
			PaymentSysMemberId: "123456789",
			BankName:           "Bank of New York",
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
	assert.Equal(t, original.AssignmentId, decoded.AssignmentId)
	assert.Equal(t, original.ResolvedCaseId, decoded.ResolvedCaseId)
	assert.Equal(t, original.OriginalMessageId, decoded.OriginalMessageId)
	assert.Equal(t, original.AssignmentCreateTime.UTC(), decoded.AssignmentCreateTime.UTC())
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
			xmlData: `<?xml version="1.0"?><Document xmlns="wrong:namespace"><RsltnOfInvstgtn></RsltnOfInvstgtn></Document>`,
			wantErr: true,
		},
		{
			name: "Missing required fields",
			xmlData: `<?xml version="1.0"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.029.001.08">
	<RsltnOfInvstgtn>
		<Assgnmt>
			<Id>TEST001</Id>
		</Assgnmt>
	</RsltnOfInvstgtn>
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
			AssignmentId:              "DOC001",
			AssignmentCreateTime:      time.Now(),
			ResolvedCaseId:            "DOC_CASE001",
			OriginalMessageId:         "DOC_ORIG001",
			OriginalMessageNameId:     "pacs.008.001.08",
			OriginalMessageCreateTime: time.Now().AddDate(0, 0, -1),
			Status:                    "ACCP",
			Assigner: models.Agent{
				BusinessIdCode:     "BANKUSNY",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "123456789",
			},
			Assignee: models.Agent{
				BusinessIdCode:     "BANKUSLA",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "987654321",
			},
			Creator: models.Agent{
				BusinessIdCode:     "BANKUSNY",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "123456789",
			},
			CancellationStatusReasonInfo: models.Reason{
				Reason: "AC01",
			},
		}

		doc, err := DocumentWith(model, CAMT_029_001_03)
		require.NoError(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("Invalid model fails validation", func(t *testing.T) {
		model := MessageModel{
			AssignmentId: "DOC002",
			// Missing required fields
		}

		doc, err := DocumentWith(model, CAMT_029_001_03)
		require.Error(t, err)
		assert.Nil(t, doc)
		assert.Contains(t, err.Error(), "required")
	})
}

// TestEnhancedFieldValidation tests enhanced field validation
func TestEnhancedFieldValidation(t *testing.T) {
	t.Run("EnhancedTransactionFields validation", func(t *testing.T) {
		fields := &EnhancedTransactionFields{
			OriginalUETR: "550e8400-e29b-41d4-a716-446655440000",
		}
		err := fields.Validate()
		assert.NoError(t, err)
	})

	t.Run("EnhancedTransactionFields empty", func(t *testing.T) {
		fields := &EnhancedTransactionFields{}
		err := fields.Validate()
		assert.NoError(t, err)
	})

	t.Run("AddressEnhancementFields validation", func(t *testing.T) {
		fields := &AddressEnhancementFields{}
		err := fields.Validate()
		assert.NoError(t, err)
	})
}

// TestAgentStructures tests the agent structures in the message
func TestAgentStructures(t *testing.T) {
	t.Run("Valid agents", func(t *testing.T) {
		model := MessageModel{
			AssignmentId:              "AGENT001",
			AssignmentCreateTime:      time.Now(),
			ResolvedCaseId:            "CASE_AGENT001",
			OriginalMessageId:         "ORIG_AGENT001",
			OriginalMessageNameId:     "pacs.008.001.08",
			OriginalMessageCreateTime: time.Now(),
			Assigner: models.Agent{
				BusinessIdCode:     "BANKUSNY",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "123456789",
				BankName:           "Bank of New York",
			},
			Assignee: models.Agent{
				BusinessIdCode:     "BANKUSLA",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "987654321",
				BankName:           "Bank of Los Angeles",
			},
			Creator: models.Agent{
				BusinessIdCode:     "BANKUSCH",
				PaymentSysCode:     "USABA",
				PaymentSysMemberId: "111222333",
				BankName:           "Bank of Chicago",
			},
		}

		// Verify agent fields
		assert.Equal(t, "BANKUSNY", model.Assigner.BusinessIdCode)
		assert.Equal(t, "123456789", model.Assigner.PaymentSysMemberId)
		assert.Equal(t, "BANKUSLA", model.Assignee.BusinessIdCode)
		assert.Equal(t, "987654321", model.Assignee.PaymentSysMemberId)
		assert.Equal(t, "BANKUSCH", model.Creator.BusinessIdCode)
		assert.Equal(t, "111222333", model.Creator.PaymentSysMemberId)
	})
}

// TestStatusAndReasonFields tests status and reason structures
func TestStatusAndReasonFields(t *testing.T) {
	t.Run("Valid status and reason", func(t *testing.T) {
		model := MessageModel{
			Status: "ACCP",
			CancellationStatusReasonInfo: models.Reason{
				Reason:         "AC01",
				AdditionalInfo: "Return accepted as requested",
			},
		}

		assert.Equal(t, models.Status("ACCP"), model.Status)
		assert.Equal(t, "AC01", model.CancellationStatusReasonInfo.Reason)
		assert.Equal(t, "Return accepted as requested", model.CancellationStatusReasonInfo.AdditionalInfo)
	})

	t.Run("Different status values", func(t *testing.T) {
		statuses := []models.Status{
			"ACCP",
			"RJCT",
			"PDNG",
		}

		for _, status := range statuses {
			model := MessageModel{
				Status: status,
			}
			assert.Equal(t, status, model.Status)
		}
	})
}

// BenchmarkWriteXML benchmarks XML writing performance
func BenchmarkWriteXML(b *testing.B) {
	model := MessageModel{
		AssignmentId:              "BENCH001",
		AssignmentCreateTime:      time.Now(),
		ResolvedCaseId:            "BENCH_CASE001",
		OriginalMessageId:         "BENCH_ORIG001",
		OriginalMessageNameId:     "pacs.008.001.08",
		OriginalMessageCreateTime: time.Now(),
		Status:                    "ACCP",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_ = model.WriteXML(&buf, CAMT_029_001_08)
	}
}

// BenchmarkParseXML benchmarks XML parsing performance
func BenchmarkParseXML(b *testing.B) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:camt.029.001.08">
	<RsltnOfInvstgtn>
		<Assgnmt>
			<Id>BENCH001</Id>
			<Assgnr>
				<Agt>
					<FinInstnId>
						<ClrSysMmbId>
							<MmbId>123456789</MmbId>
						</ClrSysMmbId>
					</FinInstnId>
				</Agt>
			</Assgnr>
			<Assgne>
				<Agt>
					<FinInstnId>
						<ClrSysMmbId>
							<MmbId>987654321</MmbId>
						</ClrSysMmbId>
					</FinInstnId>
				</Agt>
			</Assgne>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
		</Assgnmt>
		<RslvdCase>
			<Id>BENCH_CASE001</Id>
			<Cretr>
				<Agt>
					<FinInstnId>
						<ClrSysMmbId>
							<MmbId>123456789</MmbId>
						</ClrSysMmbId>
					</FinInstnId>
				</Agt>
			</Cretr>
		</RslvdCase>
		<Sts>
			<Conf>ACCP</Conf>
		</Sts>
		<CxlDtls>
			<OrgnlGrpInfAndSts>
				<OrgnlMsgId>BENCH_ORIG001</OrgnlMsgId>
				<OrgnlMsgNmId>pacs.008.001.08</OrgnlMsgNmId>
				<OrgnlCreDtTm>2024-01-01T09:00:00Z</OrgnlCreDtTm>
			</OrgnlGrpInfAndSts>
		</CxlDtls>
	</RsltnOfInvstgtn>
</Document>`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseXML(xmlData)
	}
}

// TestBuildReasonHelper tests the reason helper builder
func TestBuildReasonHelper(t *testing.T) {
	helper := BuildReasonHelper()

	// Test field initialization
	assert.Equal(t, "Originator", helper.Originator.Title)
	assert.Equal(t, "Reason", helper.Reason.Title)
	assert.Equal(t, "Additional Info", helper.AdditionalInfo.Title)

	// Test type information
	assert.Contains(t, helper.Originator.Type, "Max35Text")
	assert.Contains(t, helper.Reason.Type, "Max35Text")
	assert.Contains(t, helper.AdditionalInfo.Type, "Max140Text")

	// Test documentation
	assert.Contains(t, helper.Originator.Documentation, "cancellation request")
	assert.Contains(t, helper.Reason.Documentation, "reason for the cancellation")
	assert.Contains(t, helper.AdditionalInfo.Documentation, "Further details")
}

// TestBuildMessageHelper tests the message helper builder
func TestBuildMessageHelper(t *testing.T) {
	helper := BuildMessageHelper()

	// Test basic field initialization
	assert.Equal(t, "Assignment Id", helper.AssignmentId.Title)
	assert.Equal(t, "Assignment Create Time", helper.AssignmentCreateTime.Title)
	assert.Equal(t, "Resolved Case Id", helper.ResolvedCaseId.Title)
	assert.Equal(t, "Status", helper.Status.Title)
	assert.Equal(t, "Original Message Id", helper.OriginalMessageId.Title)
	assert.Equal(t, "Original Message Name Id", helper.OriginalMessageNameId.Title)
	assert.Equal(t, "Original Message Create Time", helper.OriginalMessageCreateTime.Title)
	assert.Equal(t, "Original Instruction Id", helper.OriginalInstructionId.Title)
	assert.Equal(t, "Original End To End Id", helper.OriginalEndToEndId.Title)
	assert.Equal(t, "Original UETR", helper.OriginalUETR.Title)

	// Test agent helpers are properly initialized
	assert.Equal(t, "Member Identification", helper.Assigner.PaymentSysMemberId.Title)
	assert.Equal(t, "Member Identification", helper.Assignee.PaymentSysMemberId.Title)
	assert.Equal(t, "Member Identification", helper.Creator.PaymentSysMemberId.Title)

	// Test nested reason helper
	assert.Equal(t, "Originator", helper.CancellationStatusReasonInfo.Originator.Title)
	assert.Equal(t, "Reason", helper.CancellationStatusReasonInfo.Reason.Title)

	// Test type information
	assert.Contains(t, helper.AssignmentId.Type, "Max35Text")
	assert.Contains(t, helper.AssignmentCreateTime.Type, "ISODateTime")
	assert.Contains(t, helper.ResolvedCaseId.Type, "Max35Text")
	assert.Contains(t, helper.OriginalMessageCreateTime.Type, "ISODateTime")

	// Test documentation
	assert.Contains(t, helper.AssignmentId.Documentation, "assignment")
	assert.Contains(t, helper.ResolvedCaseId.Documentation, "resolved case")
	assert.Contains(t, helper.Status.Documentation, "status of the investigation")
	assert.Contains(t, helper.OriginalMessageId.Documentation, "Point to point reference")
	assert.Contains(t, helper.OriginalUETR.Documentation, "Universally unique identifier")
}

// TestMessageHelperCompleteness tests that all fields have documentation
func TestMessageHelperCompleteness(t *testing.T) {
	helper := BuildMessageHelper()

	// Verify all fields have non-empty documentation
	assert.NotEmpty(t, helper.AssignmentId.Documentation)
	assert.NotEmpty(t, helper.AssignmentCreateTime.Documentation)
	assert.NotEmpty(t, helper.ResolvedCaseId.Documentation)
	assert.NotEmpty(t, helper.Status.Documentation)
	assert.NotEmpty(t, helper.OriginalMessageId.Documentation)
	assert.NotEmpty(t, helper.OriginalMessageNameId.Documentation)
	assert.NotEmpty(t, helper.OriginalMessageCreateTime.Documentation)
	assert.NotEmpty(t, helper.OriginalInstructionId.Documentation)
	assert.NotEmpty(t, helper.OriginalEndToEndId.Documentation)
	assert.NotEmpty(t, helper.OriginalUETR.Documentation)

	// Verify all fields have non-empty types
	assert.NotEmpty(t, helper.AssignmentId.Type)
	assert.NotEmpty(t, helper.AssignmentCreateTime.Type)
	assert.NotEmpty(t, helper.ResolvedCaseId.Type)
	assert.NotEmpty(t, helper.Status.Type)
	assert.NotEmpty(t, helper.OriginalMessageId.Type)
	assert.NotEmpty(t, helper.OriginalMessageNameId.Type)
	assert.NotEmpty(t, helper.OriginalMessageCreateTime.Type)
	assert.NotEmpty(t, helper.OriginalInstructionId.Type)
	assert.NotEmpty(t, helper.OriginalEndToEndId.Type)
	assert.NotEmpty(t, helper.OriginalUETR.Type)

	// Verify agent helpers are complete
	assert.NotEmpty(t, helper.Assigner.BusinessIdCode.Documentation)
	assert.NotEmpty(t, helper.Assignee.BusinessIdCode.Documentation)
	assert.NotEmpty(t, helper.Creator.BusinessIdCode.Documentation)

	// Verify cancellation reason helper is complete
	assert.NotEmpty(t, helper.CancellationStatusReasonInfo.Originator.Documentation)
	assert.NotEmpty(t, helper.CancellationStatusReasonInfo.Reason.Documentation)
	assert.NotEmpty(t, helper.CancellationStatusReasonInfo.AdditionalInfo.Documentation)
}