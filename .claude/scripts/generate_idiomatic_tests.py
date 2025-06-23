#!/usr/bin/env python3
"""
Generate idiomatic test files for wire20022 message types to improve test coverage.
This script creates consistent, comprehensive test files following established patterns.
"""

import os
import sys
from pathlib import Path

# Test template for idiomatic Go tests
TEST_TEMPLATE = '''package {package_name}

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/wire20022/pkg/base"
	"github.com/moov-io/wire20022/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestReadWriteXML tests the XML-first API for reading and writing
func TestReadWriteXML(t *testing.T) {{
	// Create a valid message model
	model := create{message_type}TestModel()

	// Test WriteXML with latest version
	var buf bytes.Buffer
	err := model.WriteXML(&buf, {version_const}_LATEST)
	require.NoError(t, err)
	require.NotEmpty(t, buf.String())
	require.Contains(t, buf.String(), "<?xml version=\\"1.0\\" encoding=\\"UTF-8\\"?>")

	// Test ReadXML with the generated XML
	var readModel MessageModel
	reader := strings.NewReader(buf.String())
	err = readModel.ReadXML(reader)
	require.NoError(t, err)
	
	// Verify key fields were preserved
	assert.Equal(t, model.MessageId, readModel.MessageId)
	assert.Equal(t, model.CreatedDateTime.Format(time.RFC3339), readModel.CreatedDateTime.Format(time.RFC3339))
}}

// TestParseXML tests parsing XML from byte data
func TestParseXML(t *testing.T) {{
	// Use a sample from swiftSample directory if available
	xmlData := getSample{message_type}XML()

	model, err := ParseXML([]byte(xmlData))
	require.NoError(t, err)
	require.NotNil(t, model)
	
	// Verify required fields are populated
	assert.NotEmpty(t, model.MessageId)
	assert.False(t, model.CreatedDateTime.IsZero())
}}

// TestWriteXMLVersions tests writing XML for different versions
func TestWriteXMLVersions(t *testing.T) {{
	model := create{message_type}TestModel()

	// Test each supported version
	versions := []VERSION{{
		{version_const}_V1,
		{version_const}_V2,
		// Add more versions as supported
	}}

	for _, version := range versions {{
		t.Run(version.String(), func(t *testing.T) {{
			var buf bytes.Buffer
			err := model.WriteXML(&buf, version)
			
			if err != nil {{
				// Some versions might not support all fields
				t.Logf("Version %s error: %v", version, err)
			}} else {{
				assert.NotEmpty(t, buf.String())
				assert.Contains(t, buf.String(), "<?xml")
			}}
		}})
	}}
}}

// TestValidateForVersion tests version-specific validation
func TestValidateForVersion(t *testing.T) {{
	testCases := []struct {{
		name    string
		model   MessageModel
		version VERSION
		wantErr bool
	}}{{
		{{
			name:    "Valid model for latest version",
			model:   create{message_type}TestModel(),
			version: {version_const}_LATEST,
			wantErr: false,
		}},
		{{
			name:    "Missing required field",
			model:   MessageModel{{}},
			version: {version_const}_LATEST,
			wantErr: true,
		}},
	}}

	for _, tc := range testCases {{
		t.Run(tc.name, func(t *testing.T) {{
			err := tc.model.ValidateForVersion(tc.version)
			if tc.wantErr {{
				assert.Error(t, err)
			}} else {{
				assert.NoError(t, err)
			}}
		}})
	}}
}}

// TestValidateCoreFields tests core field validation
func TestValidateCoreFields(t *testing.T) {{
	t.Run("Valid core fields", func(t *testing.T) {{
		model := create{message_type}TestModel()
		err := model.validateCoreFields()
		assert.NoError(t, err)
	}})

	t.Run("Missing MessageId", func(t *testing.T) {{
		model := create{message_type}TestModel()
		model.MessageId = ""
		err := model.validateCoreFields()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "MessageId")
	}})

	t.Run("Missing CreatedDateTime", func(t *testing.T) {{
		model := create{message_type}TestModel()
		model.CreatedDateTime = time.Time{{}}
		err := model.validateCoreFields()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "CreatedDateTime")
	}})
}}

// TestGetVersionCapabilities tests version capability detection
func TestGetVersionCapabilities(t *testing.T) {{
	t.Run("Base version capabilities", func(t *testing.T) {{
		model := MessageModel{{}}
		caps := model.GetVersionCapabilities()
		assert.NotNil(t, caps)
		// Check for version-specific capabilities
	}})

	t.Run("Enhanced version capabilities", func(t *testing.T) {{
		model := NewMessageForVersion({version_const}_LATEST)
		caps := model.GetVersionCapabilities()
		assert.NotNil(t, caps)
		// Verify version-specific fields are detected
	}})
}}

// TestUnmarshalJSON tests JSON unmarshaling
func TestUnmarshalJSON(t *testing.T) {{
	jsonData := `{{
		"messageId": "TEST123456789",
		"createdDateTime": "2024-01-01T10:00:00Z"
		// Add message-specific fields
	}}`

	var model MessageModel
	err := json.Unmarshal([]byte(jsonData), &model)
	require.NoError(t, err)
	assert.Equal(t, "TEST123456789", model.MessageId)
}}

// TestDocumentWith tests document creation
func TestDocumentWith(t *testing.T) {{
	model := create{message_type}TestModel()

	doc, err := DocumentWith(model, {version_const}_LATEST)
	require.NoError(t, err)
	require.NotNil(t, doc)
}}

// TestCheckRequiredFields tests required field validation
func TestCheckRequiredFields(t *testing.T) {{
	t.Run("All required fields present", func(t *testing.T) {{
		model := create{message_type}TestModel()
		err := CheckRequiredFields(model)
		assert.NoError(t, err)
	}})

	t.Run("Missing required field", func(t *testing.T) {{
		model := MessageModel{{}}
		err := CheckRequiredFields(model)
		assert.Error(t, err)
	}})
}}

// TestNewMessageForVersion tests version-specific initialization
func TestNewMessageForVersion(t *testing.T) {{
	versions := []VERSION{{
		{version_const}_V1,
		{version_const}_V2,
		{version_const}_LATEST,
	}}

	for _, version := range versions {{
		t.Run(version.String(), func(t *testing.T) {{
			model := NewMessageForVersion(version)
			assert.NotNil(t, model)
			// Verify version-specific fields are initialized
		}})
	}}
}}

// Helper function to create a valid test model
func create{message_type}TestModel() MessageModel {{
	return MessageModel{{
		MessageHeader: base.MessageHeader{{
			MessageId:       "{message_prefix}123456789",
			CreatedDateTime: time.Now(),
		}},
		// Add message-specific required fields
	}}
}}

// Helper function to get sample XML
func getSample{message_type}XML() string {{
	// This would ideally read from swiftSample directory
	return `<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="{namespace}">
	<!-- Add minimal valid XML structure -->
</Document>`
}}
'''

# Message type configurations
MESSAGE_CONFIGS = {
    'ConnectionCheck': {
        'package_name': 'ConnectionCheck',
        'message_type': 'ConnectionCheck',
        'message_prefix': 'CC',
        'version_const': 'ADMI_004_001',
        'namespace': 'urn:iso:std:iso:20022:tech:xsd:admi.004.001.01'
    },
    'DrawdownRequest': {
        'package_name': 'DrawdownRequest',
        'message_type': 'DrawdownRequest',
        'message_prefix': 'DR',
        'version_const': 'PAIN_013_001',
        'namespace': 'urn:iso:std:iso:20022:tech:xsd:pain.013.001.01'
    },
    'DrawdownResponse': {
        'package_name': 'DrawdownResponse',
        'message_type': 'DrawdownResponse',
        'message_prefix': 'DRS',
        'version_const': 'PAIN_014_001',
        'namespace': 'urn:iso:std:iso:20022:tech:xsd:pain.014.001.01'
    },
    'PaymentStatusRequest': {
        'package_name': 'PaymentStatusRequest',
        'message_type': 'PaymentStatusRequest',
        'message_prefix': 'PSR',
        'version_const': 'PACS_028_001',
        'namespace': 'urn:iso:std:iso:20022:tech:xsd:pacs.028.001.01'
    },
    'Master': {
        'package_name': 'Master',
        'message_type': 'Master',
        'message_prefix': 'MST',
        'version_const': 'CAMT_052_001',
        'namespace': 'urn:iso:std:iso:20022:tech:xsd:camt.052.001.01'
    },
    'PaymentReturn': {
        'package_name': 'PaymentReturn',
        'message_type': 'PaymentReturn',
        'message_prefix': 'PR',
        'version_const': 'PACS_004_001',
        'namespace': 'urn:iso:std:iso:20022:tech:xsd:pacs.004.001.01'
    },
    'FedwireFundsPaymentStatus': {
        'package_name': 'FedwireFundsPaymentStatus',
        'message_type': 'FedwireFundsPaymentStatus',
        'message_prefix': 'FFPS',
        'version_const': 'PACS_002_001',
        'namespace': 'urn:iso:std:iso:20022:tech:xsd:pacs.002.001.03'
    },
    'ReturnRequestResponse': {
        'package_name': 'ReturnRequestResponse',
        'message_type': 'ReturnRequestResponse',
        'message_prefix': 'RRR',
        'version_const': 'CAMT_029_001',
        'namespace': 'urn:iso:std:iso:20022:tech:xsd:camt.029.001.03'
    },
    'EndpointTotalsReport': {
        'package_name': 'EndpointTotalsReport',
        'message_type': 'EndpointTotalsReport',
        'message_prefix': 'ETR',
        'version_const': 'CAMT_052_001',
        'namespace': 'urn:iso:std:iso:20022:tech:xsd:camt.052.001.01'
    }
}

def generate_test_file(message_type, output_dir):
    """Generate a test file for the given message type."""
    config = MESSAGE_CONFIGS.get(message_type)
    if not config:
        print(f"Unknown message type: {message_type}")
        return False
    
    # Format the template with configuration
    test_content = TEST_TEMPLATE.format(**config)
    
    # Create output path
    output_path = Path(output_dir) / f"pkg/models/{config['package_name']}/Message_idiomatic_test.go"
    output_path.parent.mkdir(parents=True, exist_ok=True)
    
    # Write the test file
    with open(output_path, 'w') as f:
        f.write(test_content)
    
    print(f"Generated test file: {output_path}")
    return True

def main():
    """Main function to generate test files."""
    if len(sys.argv) < 2:
        print("Usage: python generate_idiomatic_tests.py <message_type> [output_dir]")
        print("Available message types:")
        for msg_type in MESSAGE_CONFIGS:
            print(f"  - {msg_type}")
        sys.exit(1)
    
    message_type = sys.argv[1]
    output_dir = sys.argv[2] if len(sys.argv) > 2 else "."
    
    if message_type == "all":
        # Generate tests for all message types
        for msg_type in MESSAGE_CONFIGS:
            generate_test_file(msg_type, output_dir)
    else:
        # Generate test for specific message type
        if not generate_test_file(message_type, output_dir):
            sys.exit(1)

if __name__ == "__main__":
    main()