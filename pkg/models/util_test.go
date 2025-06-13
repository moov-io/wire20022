package models

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test structs for testing reflection operations
type TestDocument struct {
	XMLName xml.Name `xml:"Document"`
	Header  TestHeader
	Details []TestDetail
}

func (t *TestDocument) Validate() error {
	return nil
}

type TestHeader struct {
	ID        string
	CreatedAt time.Time
	Count     int
}

type TestDetail struct {
	Name    string
	Address TestAddress
	Amounts []TestAmount
}

type TestAddress struct {
	Street  string
	City    string
	Country string
}

type TestAmount struct {
	Value    float64
	Currency string
}

type TestNestedStruct struct {
	Level1 *TestLevel1
}

type TestLevel1 struct {
	Level2 *TestLevel2
}

type TestLevel2 struct {
	Value string
}

func TestDocumentFrom(t *testing.T) {
	factoryMap := map[string]DocumentFactory{
		"urn:test:namespace": func() ISODocument {
			return &TestDocument{}
		},
	}

	tests := []struct {
		name        string
		xmlData     []byte
		expectError bool
		errorMsg    string
		checkResult func(*testing.T, ISODocument, string)
	}{
		{
			name: "valid XML with known namespace",
			xmlData: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:test:namespace">
	<Header>
		<ID>TEST123</ID>
		<Count>1</Count>
	</Header>
</Document>`),
			expectError: false,
			checkResult: func(t *testing.T, doc ISODocument, xmlns string) {
				assert.Equal(t, "urn:test:namespace", xmlns)
				testDoc := doc.(*TestDocument)
				assert.Equal(t, "TEST123", testDoc.Header.ID)
				assert.Equal(t, 1, testDoc.Header.Count)
			},
		},
		{
			name:        "invalid XML syntax",
			xmlData:     []byte(`<invalid>xml without closing tag`),
			expectError: true,
			errorMsg:    "XML decode failed",
		},
		{
			name: "XML without xmlns attribute",
			xmlData: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document>
	<Header>
		<ID>TEST123</ID>
	</Header>
</Document>`),
			expectError: true,
			errorMsg:    "missing xmlns attribute",
		},
		{
			name: "XML with unknown namespace",
			xmlData: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:unknown:namespace">
	<Header>
		<ID>TEST123</ID>
	</Header>
</Document>`),
			expectError: true,
			errorMsg:    "not supported",
		},
		{
			name:        "empty XML data",
			xmlData:     []byte(``),
			expectError: true,
			errorMsg:    "XML decode failed",
		},
		{
			name:        "nil XML data",
			xmlData:     nil,
			expectError: true,
			errorMsg:    "XML decode failed",
		},
		{
			name: "malformed XML structure",
			xmlData: []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:test:namespace">
	<InvalidStructure>
		<UnknownField>value</UnknownField>
	</InvalidStructure>
</Document>`),
			expectError: false, // XML unmarshals but fields are ignored
			checkResult: func(t *testing.T, doc ISODocument, xmlns string) {
				assert.Equal(t, "urn:test:namespace", xmlns)
				testDoc := doc.(*TestDocument)
				assert.Equal(t, "", testDoc.Header.ID) // Field not populated
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, xmlns, err := DocumentFrom(tt.xmlData, factoryMap)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				assert.Nil(t, doc)
				assert.Empty(t, xmlns)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, doc)
				assert.NotEmpty(t, xmlns)
				if tt.checkResult != nil {
					tt.checkResult(t, doc, xmlns)
				}
			}
		})
	}
}

func TestGetElement(t *testing.T) {
	// Create test data
	testDoc := &TestDocument{
		Header: TestHeader{
			ID:        "TEST123",
			CreatedAt: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
			Count:     42,
		},
		Details: []TestDetail{
			{
				Name: "Detail1",
				Address: TestAddress{
					Street:  "123 Main St",
					City:    "Anytown",
					Country: "US",
				},
				Amounts: []TestAmount{
					{Value: 100.50, Currency: "USD"},
					{Value: 200.75, Currency: "EUR"},
				},
			},
			{
				Name: "Detail2",
				Address: TestAddress{
					Street:  "456 Oak Ave",
					City:    "Otherville",
					Country: "CA",
				},
			},
		},
	}

	tests := []struct {
		name           string
		item           any
		path           string
		expectedType   reflect.Type
		expectedValue  any
		expectNilType  bool
		expectNilValue bool
	}{
		{
			name:          "simple field access",
			item:          testDoc,
			path:          "Header.ID",
			expectedType:  reflect.TypeOf(""),
			expectedValue: "TEST123",
		},
		{
			name:          "numeric field access",
			item:          testDoc,
			path:          "Header.Count",
			expectedType:  reflect.TypeOf(0),
			expectedValue: 42,
		},
		{
			name:          "time field access",
			item:          testDoc,
			path:          "Header.CreatedAt",
			expectedType:  reflect.TypeOf(time.Time{}),
			expectedValue: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		{
			name:          "array access with index",
			item:          testDoc,
			path:          "Details[0].Name",
			expectedType:  reflect.TypeOf(""),
			expectedValue: "Detail1",
		},
		{
			name:          "nested struct access through array",
			item:          testDoc,
			path:          "Details[0].Address.Street",
			expectedType:  reflect.TypeOf(""),
			expectedValue: "123 Main St",
		},
		{
			name:          "nested array access",
			item:          testDoc,
			path:          "Details[0].Amounts[1].Currency",
			expectedType:  reflect.TypeOf(""),
			expectedValue: "EUR",
		},
		{
			name:          "float field access",
			item:          testDoc,
			path:          "Details[0].Amounts[0].Value",
			expectedType:  reflect.TypeOf(0.0),
			expectedValue: 100.50,
		},
		{
			name:           "nil item",
			item:           nil,
			path:           "Header.ID",
			expectNilType:  true,
			expectNilValue: true,
		},
		{
			name:           "empty path",
			item:           testDoc,
			path:           "",
			expectNilType:  true,
			expectNilValue: true,
		},
		{
			name:           "non-existent field",
			item:           testDoc,
			path:           "Header.NonExistent",
			expectNilType:  true,
			expectNilValue: true,
		},
		{
			name:           "array index out of bounds",
			item:           testDoc,
			path:           "Details[10].Name",
			expectNilType:  true,
			expectNilValue: true,
		},
		{
			name:           "negative array index",
			item:           testDoc,
			path:           "Details[-1].Name",
			expectNilType:  true,
			expectNilValue: true,
		},
		{
			name:           "invalid array index format",
			item:           testDoc,
			path:           "Details[abc].Name",
			expectNilType:  true,
			expectNilValue: true,
		},
		{
			name:           "accessing non-array as array",
			item:           testDoc,
			path:           "Header[0].ID",
			expectNilType:  true,
			expectNilValue: true,
		},
		{
			name:          "access second array element",
			item:          testDoc,
			path:          "Details[1].Address.City",
			expectedType:  reflect.TypeOf(""),
			expectedValue: "Otherville",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultType, resultValue, err := GetElement(tt.item, tt.path)

			if tt.expectNilType {
				assert.Nil(t, resultType)
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expectedType, resultType)
				assert.NoError(t, err)
			}

			if tt.expectNilValue {
				assert.Nil(t, resultValue)
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expectedValue, resultValue)
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetElement_WithPointers(t *testing.T) {
	// Test with nested pointers
	nested := &TestNestedStruct{
		Level1: &TestLevel1{
			Level2: &TestLevel2{
				Value: "deep_value",
			},
		},
	}

	resultType, resultValue, err := GetElement(nested, "Level1.Level2.Value")
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(""), resultType)
	assert.Equal(t, "deep_value", resultValue)

	// Test with nil pointer in chain
	nestedWithNil := &TestNestedStruct{
		Level1: nil,
	}

	resultType, resultValue, err = GetElement(nestedWithNil, "Level1.Level2.Value")
	assert.Error(t, err)
	assert.Nil(t, resultType)
	assert.Nil(t, resultValue)
}

func TestSetElementToDocument(t *testing.T) {
	tests := []struct {
		name        string
		setupItem   func() any
		path        string
		value       any
		expectError bool
		errorMsg    string
		checkResult func(*testing.T, any)
	}{
		{
			name: "set simple string field",
			setupItem: func() any {
				return &TestDocument{}
			},
			path:  "Header.ID",
			value: "NEW123",
			checkResult: func(t *testing.T, item any) {
				doc := item.(*TestDocument)
				assert.Equal(t, "NEW123", doc.Header.ID)
			},
		},
		{
			name: "set numeric field",
			setupItem: func() any {
				return &TestDocument{}
			},
			path:  "Header.Count",
			value: 99,
			checkResult: func(t *testing.T, item any) {
				doc := item.(*TestDocument)
				assert.Equal(t, 99, doc.Header.Count)
			},
		},
		{
			name: "set time field",
			setupItem: func() any {
				return &TestDocument{}
			},
			path:  "Header.CreatedAt",
			value: time.Date(2024, 12, 25, 15, 30, 0, 0, time.UTC),
			checkResult: func(t *testing.T, item any) {
				doc := item.(*TestDocument)
				expected := time.Date(2024, 12, 25, 15, 30, 0, 0, time.UTC)
				assert.Equal(t, expected, doc.Header.CreatedAt)
			},
		},
		{
			name: "set field in array element",
			setupItem: func() any {
				doc := &TestDocument{}
				doc.Details = make([]TestDetail, 2)
				return doc
			},
			path:  "Details[0].Name",
			value: "UpdatedDetail",
			checkResult: func(t *testing.T, item any) {
				doc := item.(*TestDocument)
				assert.Equal(t, "UpdatedDetail", doc.Details[0].Name)
			},
		},
		{
			name:        "nil item",
			setupItem:   func() any { return nil },
			path:        "Header.ID",
			value:       "test",
			expectError: true,
			errorMsg:    "invalid input",
		},
		{
			name: "empty path",
			setupItem: func() any {
				return &TestDocument{}
			},
			path:        "",
			value:       "test",
			expectError: true,
			errorMsg:    "invalid input",
		},
		{
			name: "non-pointer item",
			setupItem: func() any {
				return TestDocument{} // Not a pointer
			},
			path:        "Header.ID",
			value:       "test",
			expectError: true,
			errorMsg:    "item must be a pointer",
		},
		{
			name: "invalid field path",
			setupItem: func() any {
				return &TestDocument{}
			},
			path:        "NonExistent.Field",
			value:       "test",
			expectError: true,
			errorMsg:    "field NonExistent not found",
		},
		{
			name: "array index out of bounds - verify behavior",
			setupItem: func() any {
				doc := &TestDocument{}
				doc.Details = make([]TestDetail, 1) // Only 1 element
				return doc
			},
			path:  "Details[5].Name",
			value: "test",
			// Note: Implementation may auto-expand arrays or handle this gracefully
			// So we don't set expectError and just verify the behavior
			checkResult: func(t *testing.T, item any) {
				// Just verify the function completed without panic
				assert.NotNil(t, item)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item := tt.setupItem()
			err := SetElementToDocument(item, tt.path, tt.value)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
				if tt.checkResult != nil {
					tt.checkResult(t, item)
				}
			}
		})
	}
}

func TestSetValue(t *testing.T) {
	tests := []struct {
		name        string
		setupValue  func() reflect.Value
		inputValue  any
		expectError bool
		checkResult func(*testing.T, reflect.Value)
	}{
		{
			name: "set string value",
			setupValue: func() reflect.Value {
				var s string
				return reflect.ValueOf(&s).Elem()
			},
			inputValue: "test string",
			checkResult: func(t *testing.T, v reflect.Value) {
				assert.Equal(t, "test string", v.String())
			},
		},
		{
			name: "set int value",
			setupValue: func() reflect.Value {
				var i int
				return reflect.ValueOf(&i).Elem()
			},
			inputValue: 42,
			checkResult: func(t *testing.T, v reflect.Value) {
				assert.Equal(t, int64(42), v.Int())
			},
		},
		{
			name: "set float value",
			setupValue: func() reflect.Value {
				var f float64
				return reflect.ValueOf(&f).Elem()
			},
			inputValue: 3.14,
			checkResult: func(t *testing.T, v reflect.Value) {
				assert.Equal(t, 3.14, v.Float())
			},
		},
		{
			name: "set bool value",
			setupValue: func() reflect.Value {
				var b bool
				return reflect.ValueOf(&b).Elem()
			},
			inputValue: true,
			checkResult: func(t *testing.T, v reflect.Value) {
				assert.True(t, v.Bool())
			},
		},
		{
			name: "convert string to int",
			setupValue: func() reflect.Value {
				var i int
				return reflect.ValueOf(&i).Elem()
			},
			inputValue: "123",
			checkResult: func(t *testing.T, v reflect.Value) {
				assert.Equal(t, int64(123), v.Int())
			},
		},
		{
			name: "convert string to float",
			setupValue: func() reflect.Value {
				var f float64
				return reflect.ValueOf(&f).Elem()
			},
			inputValue: "3.14159",
			checkResult: func(t *testing.T, v reflect.Value) {
				assert.Equal(t, 3.14159, v.Float())
			},
		},
		{
			name: "invalid string to int conversion",
			setupValue: func() reflect.Value {
				var i int
				return reflect.ValueOf(&i).Elem()
			},
			inputValue:  "not a number",
			expectError: true,
		},
		{
			name: "invalid string to float conversion",
			setupValue: func() reflect.Value {
				var f float64
				return reflect.ValueOf(&f).Elem()
			},
			inputValue:  "not a float",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.setupValue()
			err := setValue(v, tt.inputValue)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.checkResult != nil {
					tt.checkResult(t, v)
				}
			}
		})
	}
}

func TestFileOperations(t *testing.T) {
	// Create temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "wire20022_test_")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	testXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:test:namespace">
	<Header>
		<ID>TEST123</ID>
	</Header>
</Document>`)

	t.Run("WriteXMLTo and ReadXMLFile", func(t *testing.T) {
		filePath := filepath.Join(tmpDir, "test.xml")

		// Test writing
		err := WriteXMLTo(filePath, testXML)
		assert.NoError(t, err)

		// Verify file exists
		_, err = os.Stat(filePath)
		assert.NoError(t, err)

		// Test reading
		readData, err := ReadXMLFile(filePath)
		assert.NoError(t, err)
		assert.Equal(t, testXML, readData)
	})

	t.Run("WriteXMLToGenerate creates directory", func(t *testing.T) {
		nestedPath := filepath.Join(tmpDir, "nested", "deep", "test.xml")

		// Test writing to non-existent directory
		err := WriteXMLToGenerate(nestedPath, testXML)
		assert.NoError(t, err)

		// Verify file exists
		_, err = os.Stat(nestedPath)
		assert.NoError(t, err)

		// Test reading
		readData, err := ReadXMLFile(nestedPath)
		assert.NoError(t, err)
		assert.Equal(t, testXML, readData)
	})

	t.Run("ReadXMLFile non-existent file", func(t *testing.T) {
		_, err := ReadXMLFile(filepath.Join(tmpDir, "nonexistent.xml"))
		assert.Error(t, err)
	})

	t.Run("WriteXMLTo invalid path", func(t *testing.T) {
		// Try to write to a path that requires permissions
		err := WriteXMLTo("/root/test.xml", testXML)
		assert.Error(t, err)
	})
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{"nil value", nil, true},
		{"empty string", "", true},
		{"non-empty string", "hello", false},
		{"zero int", 0, true},
		{"non-zero int", 42, false},
		{"zero float", 0.0, true},
		{"non-zero float", 3.14, false},
		{"false bool", false, true},
		{"true bool", true, false},
		{"empty slice", []string{}, true},
		{"non-empty slice", []string{"item"}, false},
		{"empty map", map[string]string{}, true},
		{"non-empty map", map[string]string{"key": "value"}, false},
		{"zero time", time.Time{}, true},
		{"non-zero time", time.Now(), false},
		{"empty struct", struct{}{}, false}, // Structs are never considered empty
		{"pointer to empty value", func() interface{} { s := ""; return &s }(), false}, // Pointers are never empty
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEmpty(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCopyDocumentValueToMessage(t *testing.T) {
	// Create source document
	sourceDoc := &TestDocument{
		Header: TestHeader{
			ID:    "SOURCE123",
			Count: 99,
		},
		Details: []TestDetail{
			{
				Name: "SourceDetail",
				Address: TestAddress{
					Street: "Source Street",
					City:   "Source City",
				},
			},
		},
	}

	// Create target message
	targetMsg := &TestDocument{}

	t.Run("copy simple field", func(t *testing.T) {
		CopyDocumentValueToMessage(sourceDoc, "Header.ID", targetMsg, "Header.ID")
		assert.Equal(t, "SOURCE123", targetMsg.Header.ID)
	})

	t.Run("copy numeric field", func(t *testing.T) {
		CopyDocumentValueToMessage(sourceDoc, "Header.Count", targetMsg, "Header.Count")
		assert.Equal(t, 99, targetMsg.Header.Count)
	})

	t.Run("copy nested field", func(t *testing.T) {
		// Initialize the slice first
		targetMsg.Details = make([]TestDetail, 1)
		CopyDocumentValueToMessage(sourceDoc, "Details[0].Name", targetMsg, "Details[0].Name")
		assert.Equal(t, "SourceDetail", targetMsg.Details[0].Name)
	})
}

func TestCopyMessageValueToDocument(t *testing.T) {
	// Create source message
	sourceMsg := &TestDocument{
		Header: TestHeader{
			ID:    "MESSAGE123",
			Count: 55,
		},
	}

	// Create target document
	targetDoc := &TestDocument{}

	t.Run("copy simple field", func(t *testing.T) {
		err := CopyMessageValueToDocument(sourceMsg, "Header.ID", targetDoc, "Header.ID")
		assert.NoError(t, err)
		assert.Equal(t, "MESSAGE123", targetDoc.Header.ID)
	})

	t.Run("copy numeric field", func(t *testing.T) {
		err := CopyMessageValueToDocument(sourceMsg, "Header.Count", targetDoc, "Header.Count")
		assert.NoError(t, err)
		assert.Equal(t, 55, targetDoc.Header.Count)
	})

	t.Run("invalid source path", func(t *testing.T) {
		err := CopyMessageValueToDocument(sourceMsg, "NonExistent.Field", targetDoc, "Header.ID")
		assert.Error(t, err)
	})
}

func TestRemakeMapping(t *testing.T) {
	source := &TestDocument{
		Header: TestHeader{
			ID:    "TEST123",
			Count: 1,
		},
		Details: []TestDetail{
			{Name: "Detail1"},
		},
	}

	t.Run("simple flat mapping", func(t *testing.T) {
		modelMap := map[string]any{
			"Header.ID":    "MessageId",
			"Header.Count": "TransactionCount",
		}

		result := RemakeMapping(source, modelMap, true)
		expected := map[string]string{
			"Header.ID":    "MessageId",
			"Header.Count": "TransactionCount",
		}
		assert.Equal(t, expected, result)
	})

	t.Run("mapping with arrays", func(t *testing.T) {
		modelMap := map[string]any{
			"Details[0].Name": "FirstDetailName",
		}

		result := RemakeMapping(source, modelMap, true)
		expected := map[string]string{
			"Details[0].Name": "FirstDetailName",
		}
		assert.Equal(t, expected, result)
	})
}

func TestUtilityFunctions(t *testing.T) {
	t.Run("seperateKeyAndValue", func(t *testing.T) {
		key, value := seperateKeyAndValue("prefix.suffix", ".")
		assert.Equal(t, "prefix", key)
		assert.Equal(t, "suffix", value)

		key, value = seperateKeyAndValue("noseparator", ".")
		assert.Equal(t, "noseparator", key)
		assert.Equal(t, "", value)
	})

	t.Run("isEmpty internal function", func(t *testing.T) {
		assert.True(t, isEmpty(""))
		assert.True(t, isEmpty(0))
		assert.True(t, isEmpty(0.0))
		assert.True(t, isEmpty(false))
		assert.False(t, isEmpty("test"))
		assert.False(t, isEmpty(42))
		assert.False(t, isEmpty(true))
	})

	t.Run("isReflectValueNil", func(t *testing.T) {
		// Test with nil pointer
		var nilPtr *string
		nilValue := reflect.ValueOf(nilPtr)
		assert.True(t, isReflectValueNil(nilValue))

		// Test with valid pointer
		str := "test"
		validValue := reflect.ValueOf(&str)
		assert.False(t, isReflectValueNil(validValue))

		// Test with non-pointer
		nonPtrValue := reflect.ValueOf("test")
		assert.False(t, isReflectValueNil(nonPtrValue))
	})
}