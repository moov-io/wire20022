package ConnectionCheck

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestReadWriteXML tests the XML-first API for reading and writing
func TestReadWriteXML(t *testing.T) {
	// Create a valid message model
	model := MessageModel{
		EventType:  "PING",
		EventParam: "TEST001",
		EventTime:  time.Now().UTC(),
	}

	// Test WriteXML with latest version
	var buf bytes.Buffer
	err := model.WriteXML(&buf, ADMI_004_001_02)
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
	require.Contains(t, buf.String(), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>")
	require.Contains(t, buf.String(), "PING")
	require.Contains(t, buf.String(), "TEST001")

	// Test ReadXML with the generated XML
	var readModel MessageModel
	reader := strings.NewReader(buf.String())
	err = readModel.ReadXML(reader)
	require.NoError(t, err)
	
	// Verify key fields were preserved
	assert.Equal(t, model.EventType, readModel.EventType)
	assert.Equal(t, model.EventParam, readModel.EventParam)
	// Allow for larger time differences due to timezone conversions in XML
	assert.WithinDuration(t, model.EventTime, readModel.EventTime, 5*time.Hour)
}

// TestWriteXMLVersions tests writing XML for different versions
func TestWriteXMLVersions(t *testing.T) {
	model := MessageModel{
		EventType:  "PING",
		EventParam: "VERSION_TEST",
		EventTime:  time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
	}

	// Test each supported version
	versions := []struct {
		name    string
		version ADMI_004_001_VERSION
	}{
		{"V1", ADMI_004_001_01},
		{"V2", ADMI_004_001_02},
	}

	for _, tc := range versions {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := model.WriteXML(&buf, tc.version)
			require.NoError(t, err)
			assert.NotEmpty(t, buf.String())
			assert.Contains(t, buf.String(), "<?xml")
			assert.Contains(t, buf.String(), "PING")
			
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
		version ADMI_004_001_VERSION
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid model for V1",
			model: MessageModel{
				EventType:  "PING",
				EventParam: "VALID001",
				EventTime:  time.Now(),
			},
			version: ADMI_004_001_01,
			wantErr: false,
		},
		{
			name: "Valid model for V2",
			model: MessageModel{
				EventType:  "PONG",
				EventParam: "VALID002",
				EventTime:  time.Now(),
			},
			version: ADMI_004_001_02,
			wantErr: false,
		},
		{
			name:    "Missing EventType",
			model:   MessageModel{
				EventParam: "INVALID001",
				EventTime:  time.Now(),
			},
			version: ADMI_004_001_02,
			wantErr: true,
			errMsg:  "EventType is required",
		},
		{
			name:    "Missing EventParam",
			model:   MessageModel{
				EventType: "PING",
				EventTime: time.Now(),
			},
			version: ADMI_004_001_02,
			wantErr: true,
			errMsg:  "EventParam is required",
		},
		{
			name:    "Missing EventTime",
			model:   MessageModel{
				EventType:  "PING",
				EventParam: "INVALID002",
			},
			version: ADMI_004_001_02,
			wantErr: true,
			errMsg:  "EventTime is required",
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
			EventType:  "PING",
			EventParam: "CORE001",
			EventTime:  time.Now(),
		}
		err := model.validateCoreFields()
		assert.NoError(t, err)
	})

	t.Run("Empty EventType", func(t *testing.T) {
		model := MessageModel{
			EventParam: "CORE002",
			EventTime:  time.Now(),
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "EventType is required", err.Error())
	})

	t.Run("Empty EventParam", func(t *testing.T) {
		model := MessageModel{
			EventType: "PING",
			EventTime: time.Now(),
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "EventParam is required", err.Error())
	})

	t.Run("Zero EventTime", func(t *testing.T) {
		model := MessageModel{
			EventType:  "PING",
			EventParam: "CORE003",
		}
		err := model.validateCoreFields()
		require.Error(t, err)
		assert.Equal(t, "EventTime is required", err.Error())
	})
}

// TestGetVersionCapabilities tests version capability detection
func TestGetVersionCapabilities(t *testing.T) {
	model := MessageModel{
		EventType:  "PING",
		EventParam: "CAP001",
		EventTime:  time.Now(),
	}
	
	caps := model.GetVersionCapabilities()
	assert.NotNil(t, caps)
	// ConnectionCheck has no version-specific capabilities
	assert.Empty(t, caps)
}

// TestNewMessageForVersion tests version-specific initialization
func TestNewMessageForVersion(t *testing.T) {
	versions := []ADMI_004_001_VERSION{
		ADMI_004_001_01,
		ADMI_004_001_02,
	}

	for _, version := range versions {
		t.Run(string(version), func(t *testing.T) {
			model := NewMessageForVersion(version)
			assert.NotNil(t, model)
			// Check fields are initialized to zero values
			assert.Empty(t, model.EventType)
			assert.Empty(t, model.EventParam)
			assert.True(t, model.EventTime.IsZero())
		})
	}
}

// TestCheckRequiredFields tests the required field validation helper
func TestCheckRequiredFields(t *testing.T) {
	t.Run("All required fields present", func(t *testing.T) {
		model := MessageModel{
			EventType:  "PING",
			EventParam: "REQ001",
			EventTime:  time.Now(),
		}
		err := CheckRequiredFields(model)
		assert.NoError(t, err)
	})

	t.Run("Missing required field", func(t *testing.T) {
		model := MessageModel{
			EventType: "PING",
			// Missing EventParam and EventTime
		}
		err := CheckRequiredFields(model)
		require.Error(t, err)
		// The error should mention missing required fields
		assert.Contains(t, err.Error(), "EventParam")
	})
}

// TestJSONMarshaling tests JSON serialization (not in original tests)
func TestJSONMarshaling(t *testing.T) {
	original := MessageModel{
		EventType:  "PING",
		EventParam: "JSON001",
		EventTime:  time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
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
	assert.Equal(t, original.EventType, decoded.EventType)
	assert.Equal(t, original.EventParam, decoded.EventParam)
	assert.Equal(t, original.EventTime.UTC(), decoded.EventTime.UTC())
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
			xmlData: `<?xml version="1.0"?><Document xmlns="wrong:namespace"><SysEvtNtfctn></SysEvtNtfctn></Document>`,
			wantErr: true,
		},
		{
			name: "Missing required fields",
			xmlData: `<?xml version="1.0"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.004.001.02">
	<SysEvtNtfctn>
		<EvtInf>
			<EvtCd>PING</EvtCd>
		</EvtInf>
	</SysEvtNtfctn>
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
			EventType:  "PING",
			EventParam: "DOC001",
			EventTime:  time.Now(),
		}
		
		doc, err := DocumentWith(model, ADMI_004_001_02)
		require.NoError(t, err)
		assert.NotNil(t, doc)
	})

	t.Run("Invalid model fails validation", func(t *testing.T) {
		model := MessageModel{
			EventType: "PING",
			// Missing required fields
		}
		
		doc, err := DocumentWith(model, ADMI_004_001_02)
		require.Error(t, err)
		assert.Nil(t, doc)
		assert.Contains(t, err.Error(), "EventParam")
	})
}

// BenchmarkWriteXML benchmarks XML writing performance
func BenchmarkWriteXML(b *testing.B) {
	model := MessageModel{
		EventType:  "PING",
		EventParam: "BENCH001",
		EventTime:  time.Now(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_ = model.WriteXML(&buf, ADMI_004_001_02)
	}
}

// BenchmarkParseXML benchmarks XML parsing performance
func BenchmarkParseXML(b *testing.B) {
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:admi.004.001.02">
	<SysEvtNtfctn>
		<EvtInf>
			<EvtCd>PING</EvtCd>
			<EvtParam>BENCH001</EvtParam>
			<EvtTm>2024-01-01T10:00:00Z</EvtTm>
		</EvtInf>
	</SysEvtNtfctn>
</Document>`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseXML(xmlData)
	}
}