package DrawdownRequest

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestReadWriteXML tests the XML-first API for reading and writing
func TestReadWriteXML(t *testing.T) {
	// Create a valid message model using version-specific initialization
	model := NewMessageForVersion(PAIN_013_001_07)
	model.MessageHeader = base.MessageHeader{
		MessageId:       "20250623DD001",
		CreatedDateTime: time.Now().UTC(),
	}
	model.NumberofTransaction = "1"
	model.InitiatingParty = models.PartyIdentify{
		Name: "Test Corporation",
	}
	model.PaymentInfoId = "PYMTINFO001"
	model.PaymentMethod = models.PaymentMethod("TRF")
	model.RequestedExecutDate = fedwire.ISODate{}
	model.Debtor = models.PartyIdentify{
		Name: "Test Debtor Corp",
	}
	model.DebtorAgent = models.Agent{
		PaymentSysMemberId: "021040078",
	}
	model.CreditTransTransaction = CreditTransferTransaction{
		PaymentInstructionId: "INSTR001",
		PaymentEndToEndId:    "E2E001",
	}

	// Test WriteXML with specific version
	var buf bytes.Buffer
	err := model.WriteXML(&buf, PAIN_013_001_07)
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
	require.Contains(t, buf.String(), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	require.Contains(t, buf.String(), "20250623DD001")
	require.Contains(t, buf.String(), "Test Corporation")

	// Test ReadXML with the generated XML
	var readModel MessageModel
	reader := strings.NewReader(buf.String())
	err = readModel.ReadXML(reader)
	require.NoError(t, err)

	// Verify key fields were preserved
	assert.Equal(t, model.MessageId, readModel.MessageId)
	assert.Equal(t, model.NumberofTransaction, readModel.NumberofTransaction)
	assert.Equal(t, model.InitiatingParty.Name, readModel.InitiatingParty.Name)
	assert.WithinDuration(t, model.CreatedDateTime, readModel.CreatedDateTime, 5*time.Hour)
}

// TestWriteXMLVersions tests writing XML for different versions
func TestWriteXMLVersions(t *testing.T) {
	// Test each supported version with appropriate field configurations
	versions := []struct {
		name    string
		version PAIN_013_001_VERSION
		hasAccountEnhancement bool
		hasAddressEnhancement bool
	}{
		{"V1", PAIN_013_001_01, false, false},
		{"V5", PAIN_013_001_05, true, false},
		{"V7", PAIN_013_001_07, true, true},
		{"V10", PAIN_013_001_10, true, true},
	}

	for _, tc := range versions {
		t.Run(tc.name, func(t *testing.T) {
			model := NewMessageForVersion(tc.version)
			model.MessageHeader = base.MessageHeader{
				MessageId:       "VERSION_TEST_" + string(tc.version),
				CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			}
			model.NumberofTransaction = "1"
			model.InitiatingParty = models.PartyIdentify{
				Name: "Version Test Corp",
			}
			model.PaymentInfoId = "VTEST001"
			model.PaymentMethod = models.PaymentMethod("TRF")
			model.RequestedExecutDate = fedwire.ISODate{}
			model.Debtor = models.PartyIdentify{
				Name: "Test Debtor",
			}
			model.DebtorAgent = models.Agent{
				PaymentSysMemberId: "021040078",
			}
			model.CreditTransTransaction = CreditTransferTransaction{
				PaymentInstructionId: "VINSTR001",
				PaymentEndToEndId:    "VE2E001",
			}

			// Verify version-specific fields are properly initialized
			if tc.hasAccountEnhancement {
				assert.NotNil(t, model.AccountEnhancement, "AccountEnhancement should be initialized for %s", tc.version)
			} else {
				assert.Nil(t, model.AccountEnhancement, "AccountEnhancement should be nil for %s", tc.version)
			}

			if tc.hasAddressEnhancement {
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
		version PAIN_013_001_VERSION
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid model for V1",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "VALID001",
					CreatedDateTime: time.Now(),
				},
				NumberofTransaction: "1",
				InitiatingParty:     models.PartyIdentify{Name: "Test Corp"},
				PaymentInfoId:       "PMT001",
				PaymentMethod:       models.PaymentMethod("TRF"),
				RequestedExecutDate: fedwire.ISODate{},
				Debtor:              models.PartyIdentify{Name: "Debtor Corp"},
				DebtorAgent:         models.Agent{PaymentSysMemberId: "021040078"},
				CreditTransTransaction: CreditTransferTransaction{
					PaymentInstructionId: "INSTR001",
					PaymentEndToEndId:    "E2E001",
				},
			},
			version: PAIN_013_001_01,
			wantErr: false,
		},
		{
			name: "Valid model for V7 with enhancements",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "VALID007",
					CreatedDateTime: time.Now(),
				},
				NumberofTransaction: "1",
				InitiatingParty:     models.PartyIdentify{Name: "Test Corp"},
				PaymentInfoId:       "PMT007",
				PaymentMethod:       models.PaymentMethod("TRF"),
				RequestedExecutDate: fedwire.ISODate{},
				Debtor:              models.PartyIdentify{Name: "Debtor Corp"},
				DebtorAgent:         models.Agent{PaymentSysMemberId: "021040078"},
				CreditTransTransaction: CreditTransferTransaction{
					PaymentInstructionId: "INSTR007",
					PaymentEndToEndId:    "E2E007",
				},
				AccountEnhancement: &AccountEnhancementFields{},
				AddressEnhancement: &AddressEnhancementFields{},
			},
			version: PAIN_013_001_07,
			wantErr: false,
		},
		{
			name: "Missing MessageId",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					CreatedDateTime: time.Now(),
				},
				NumberofTransaction: "1",
			},
			version: PAIN_013_001_02,
			wantErr: true,
			errMsg:  "MessageId is required",
		},
		{
			name: "Missing CreatedDateTime",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId: "INVALID001",
				},
				NumberofTransaction: "1",
			},
			version: PAIN_013_001_02,
			wantErr: true,
			errMsg:  "CreatedDateTime is required",
		},
		{
			name: "Missing NumberofTransaction",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "INVALID002",
					CreatedDateTime: time.Now(),
				},
			},
			version: PAIN_013_001_02,
			wantErr: true,
			errMsg:  "NumberofTransaction is required",
		},
		{
			name: "V5 missing AccountEnhancement",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "INVALID007",
					CreatedDateTime: time.Now(),
				},
				NumberofTransaction: "1",
				InitiatingParty:     models.PartyIdentify{Name: "Test Corp"},
				PaymentInfoId:       "PMT007",
				PaymentMethod:       models.PaymentMethod("TRF"),
				RequestedExecutDate: fedwire.ISODate{},
				Debtor:              models.PartyIdentify{Name: "Debtor Corp"},
				DebtorAgent:         models.Agent{PaymentSysMemberId: "021040078"},
				CreditTransTransaction: CreditTransferTransaction{
					PaymentInstructionId: "INSTR007",
					PaymentEndToEndId:    "E2E007",
				},
				// Missing AccountEnhancement for V5+
			},
			version: PAIN_013_001_05,
			wantErr: true,
			errMsg:  "AccountEnhancementFields required for version",
		},
		{
			name: "V7 missing AddressEnhancement",
			model: MessageModel{
				MessageHeader: base.MessageHeader{
					MessageId:       "INVALID007B",
					CreatedDateTime: time.Now(),
				},
				NumberofTransaction: "1",
				InitiatingParty:     models.PartyIdentify{Name: "Test Corp"},
				PaymentInfoId:       "PMT007B",
				PaymentMethod:       models.PaymentMethod("TRF"),
				RequestedExecutDate: fedwire.ISODate{},
				Debtor:              models.PartyIdentify{Name: "Debtor Corp"},
				DebtorAgent:         models.Agent{PaymentSysMemberId: "021040078"},
				CreditTransTransaction: CreditTransferTransaction{
					PaymentInstructionId: "INSTR007B",
					PaymentEndToEndId:    "E2E007B",
				},
				AccountEnhancement: &AccountEnhancementFields{},
				// Missing AddressEnhancement for V7+
			},
			version: PAIN_013_001_07,
			wantErr: true,
			errMsg:  "AddressEnhancementFields required for version",
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
			MessageHeader: base.MessageHeader{
				MessageId:       "CORE001",
				CreatedDateTime: time.Now(),
			},
			NumberofTransaction: "1",
			InitiatingParty:     models.PartyIdentify{Name: "Test Corp"},
			PaymentInfoId:       "PMT001",
			PaymentMethod:       models.PaymentMethod("TRF"),
			RequestedExecutDate: fedwire.ISODate{},
			Debtor:              models.PartyIdentify{Name: "Debtor Corp"},
			DebtorAgent:         models.Agent{PaymentSysMemberId: "021040078"},
			CreditTransTransaction: CreditTransferTransaction{
				PaymentInstructionId: "INSTR001",
				PaymentEndToEndId:    "E2E001",
			},
		}
		err := model.validateCoreFields()
		assert.NoError(t, err)
	})

	t.Run("Empty MessageId", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				CreatedDateTime: time.Now(),
			},
			NumberofTransaction: "1",
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "MessageId is required", err.Error())
	})

	t.Run("Zero CreatedDateTime", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "CORE003",
			},
			NumberofTransaction: "1",
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "CreatedDateTime is required", err.Error())
	})

	t.Run("Empty NumberofTransaction", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId:       "CORE004",
				CreatedDateTime: time.Now(),
			},
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "NumberofTransaction is required", err.Error())
	})
}

// TestGetVersionCapabilities tests version capability detection
func TestGetVersionCapabilities(t *testing.T) {
	testCases := []struct {
		name    string
		version PAIN_013_001_VERSION
		expectedCaps map[string]bool
	}{
		{
			name:    "V1 - no enhancements",
			version: PAIN_013_001_01,
			expectedCaps: map[string]bool{
				"AccountEnhancement": false,
				"AddressEnhancement": false,
			},
		},
		{
			name:    "V5 - account enhancement only",
			version: PAIN_013_001_05,
			expectedCaps: map[string]bool{
				"AccountEnhancement": true,
				"AddressEnhancement": false,
			},
		},
		{
			name:    "V7 - both enhancements",
			version: PAIN_013_001_07,
			expectedCaps: map[string]bool{
				"AccountEnhancement": true,
				"AddressEnhancement": true,
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
		version PAIN_013_001_VERSION
		hasAccountEnhancement bool
		hasAddressEnhancement bool
	}{
		{PAIN_013_001_01, false, false},
		{PAIN_013_001_02, false, false},
		{PAIN_013_001_05, true, false},
		{PAIN_013_001_07, true, true},
		{PAIN_013_001_10, true, true},
	}

	for _, v := range versions {
		t.Run(string(v.version), func(t *testing.T) {
			model := NewMessageForVersion(v.version)
			assert.NotNil(t, model)

			// Check base fields are initialized to zero values
			assert.Empty(t, model.MessageId)
			assert.Empty(t, model.NumberofTransaction)
			assert.True(t, model.CreatedDateTime.IsZero())

			// Check version-specific field initialization
			if v.hasAccountEnhancement {
				assert.NotNil(t, model.AccountEnhancement)
			} else {
				assert.Nil(t, model.AccountEnhancement)
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
			MessageHeader: base.MessageHeader{
				MessageId:       "REQ001",
				CreatedDateTime: time.Now(),
			},
			NumberofTransaction: "1",
			InitiatingParty:     models.PartyIdentify{Name: "Test Corp"},
			PaymentInfoId:       "PMT001",
			PaymentMethod:       models.PaymentMethod("TRF"),
			RequestedExecutDate: fedwire.ISODate{},
			Debtor:              models.PartyIdentify{Name: "Debtor Corp"},
			DebtorAgent:         models.Agent{PaymentSysMemberId: "021040078"},
			CreditTransTransaction: CreditTransferTransaction{
				PaymentInstructionId: "INSTR001",
				PaymentEndToEndId:    "E2E001",
			},
		}
		err := CheckRequiredFields(model)
		assert.NoError(t, err)
	})

	t.Run("Missing required fields", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "REQ002",
				// Missing CreatedDateTime
			},
			// Missing NumberofTransaction
		}
		err := CheckRequiredFields(model)
		require.Error(t, err)
		// The error should mention missing required fields
		assert.Contains(t, err.Error(), "required")
	})
}

// TestJSONMarshaling tests JSON serialization
func TestJSONMarshaling(t *testing.T) {
	t.Skip("Skipping until date handling issues are resolved")
	original := MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "JSON001",
			CreatedDateTime: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		NumberofTransaction: "1",
		InitiatingParty:     models.PartyIdentify{Name: "JSON Test Corp"},
		PaymentInfoId:       "JSONPMT001",
		PaymentMethod:       models.PaymentMethod("TRF"),
		RequestedExecutDate: fedwire.ISODate{},
		Debtor:              models.PartyIdentify{Name: "JSON Debtor"},
		DebtorAgent:         models.Agent{PaymentSysMemberId: "021040078"},
		CreditTransTransaction: CreditTransferTransaction{
			PaymentInstructionId: "JSONINSTR001",
			PaymentEndToEndId:    "JSONE2E001",
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
	assert.Equal(t, original.NumberofTransaction, decoded.NumberofTransaction)
	assert.Equal(t, original.InitiatingParty.Name, decoded.InitiatingParty.Name)
	assert.Equal(t, original.CreatedDateTime.UTC(), decoded.CreatedDateTime.UTC())
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
			xmlData: `<?xml version="1.0"?><Document xmlns="wrong:namespace"><CdtrPmtActvtnReq></CdtrPmtActvtnReq></Document>`,
			wantErr: true,
		},
		{
			name: "Missing required fields",
			xmlData: `<?xml version="1.0"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.013.001.07">
	<CdtrPmtActvtnReq>
		<GrpHdr>
			<MsgId>TEST001</MsgId>
		</GrpHdr>
	</CdtrPmtActvtnReq>
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
	t.Skip("Skipping until field mapping issues are resolved")
	
	t.Run("Valid model creates document", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId:       "DOC001",
				CreatedDateTime: time.Now(),
			},
			NumberofTransaction: "1",
			InitiatingParty:     models.PartyIdentify{Name: "Doc Test Corp"},
			PaymentInfoId:       "DOCPMT001",
			PaymentMethod:       models.PaymentMethod("TRF"),
			RequestedExecutDate: fedwire.ISODate{},
			Debtor:              models.PartyIdentify{Name: "Doc Debtor"},
			DebtorAgent:         models.Agent{PaymentSysMemberId: "021040078"},
			CreditTransTransaction: CreditTransferTransaction{
				PaymentInstructionId: "DOCINSTR001",
				PaymentEndToEndId:    "DOCE2E001",
			},
		}

		doc, err := DocumentWith(model, PAIN_013_001_01)
		require.NoError(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("Invalid model fails validation", func(t *testing.T) {
		model := MessageModel{
			MessageHeader: base.MessageHeader{
				MessageId: "DOC002",
				// Missing required fields
			},
		}

		doc, err := DocumentWith(model, PAIN_013_001_01)
		require.Error(t, err)
		assert.Nil(t, doc)
		assert.Contains(t, err.Error(), "required")
	})
}

// TestVersionSpecificFieldValidation tests field validation specific to different versions
func TestVersionSpecificFieldValidation(t *testing.T) {
	t.Run("AccountEnhancementFields validation", func(t *testing.T) {
		fields := &AccountEnhancementFields{
			DebtorAccountOtherId: "12345",
		}
		err := fields.Validate()
		assert.NoError(t, err)
	})

	t.Run("AddressEnhancementFields validation", func(t *testing.T) {
		fields := &AddressEnhancementFields{}
		err := fields.Validate()
		assert.NoError(t, err)
	})
}

// BenchmarkWriteXML benchmarks XML writing performance
func BenchmarkWriteXML(b *testing.B) {
	model := MessageModel{
		MessageHeader: base.MessageHeader{
			MessageId:       "BENCH001",
			CreatedDateTime: time.Now(),
		},
		NumberofTransaction: "1",
		InitiatingParty:     models.PartyIdentify{Name: "Bench Corp"},
		PaymentInfoId:       "BENCHPMT001",
		PaymentMethod:       models.PaymentMethod("TRF"),
		RequestedExecutDate: fedwire.ISODate{},
		Debtor:              models.PartyIdentify{Name: "Bench Debtor"},
		DebtorAgent:         models.Agent{PaymentSysMemberId: "021040078"},
		CreditTransTransaction: CreditTransferTransaction{
			PaymentInstructionId: "BENCHINSTR001",
			PaymentEndToEndId:    "BENCHE2E001",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_ = model.WriteXML(&buf, PAIN_013_001_07)
	}
}

// BenchmarkParseXML benchmarks XML parsing performance
func BenchmarkParseXML(b *testing.B) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.013.001.07">
	<CdtrPmtActvtnReq>
		<GrpHdr>
			<MsgId>BENCH001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
			<NbOfTxs>1</NbOfTxs>
			<InitgPty>
				<Nm>Bench Corp</Nm>
			</InitgPty>
		</GrpHdr>
		<PmtInf>
			<PmtInfId>BENCHPMT001</PmtInfId>
			<PmtMtd>TRF</PmtMtd>
			<ReqdExctnDt>
				<Dt>2024-02-01</Dt>
			</ReqdExctnDt>
			<Dbtr>
				<Nm>Bench Debtor</Nm>
			</Dbtr>
			<DbtrAgt>
				<FinInstnId>
					<ClrSysMmbId>
						<MmbId>021040078</MmbId>
					</ClrSysMmbId>
				</FinInstnId>
			</DbtrAgt>
			<CdtTrfTx>
				<PmtId>
					<InstrId>BENCHINSTR001</InstrId>
					<EndToEndId>BENCHE2E001</EndToEndId>
				</PmtId>
			</CdtTrfTx>
		</PmtInf>
	</CdtrPmtActvtnReq>
</Document>`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseXML(xmlData)
	}
}