package base

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wadearnold/wire20022/pkg/models"
)

// Mock types for testing
type TestMessage struct {
	MessageHeader
	TestField string `json:"testField"`
}

type TestVersion string

const (
	TestV1 TestVersion = "001.01"
	TestV2 TestVersion = "001.02"
)

type MockDocument struct {
	XMLName xml.Name `xml:"Document"`
	Content string   `xml:"content"`
}

func (m *MockDocument) Validate() error {
	return nil
}

func TestMessageProcessor(t *testing.T) {
	t.Run("NewMessageProcessor creation", func(t *testing.T) {
		namespaceMap := map[string]models.DocumentFactory{
			"test:namespace:v1": func() models.ISODocument {
				return &MockDocument{Content: "test"}
			},
		}

		versionMap := map[string]TestVersion{
			"test:namespace:v1": TestV1,
		}

		pathMaps := map[TestVersion]map[string]any{
			TestV1: {
				"Document.content": "TestField",
			},
		}

		requiredFields := []string{"MessageId", "TestField"}

		processor := NewMessageProcessor[TestMessage, TestVersion](
			namespaceMap,
			versionMap,
			pathMaps,
			requiredFields,
		)

		assert.NotNil(t, processor)
		assert.Equal(t, namespaceMap, processor.namespaceMap)
		assert.Equal(t, versionMap, processor.versionMap)
		assert.Equal(t, pathMaps, processor.pathMaps)
		assert.Equal(t, requiredFields, processor.requiredFields)
	})

	t.Run("ProcessMessage with unsupported namespace", func(t *testing.T) {
		processor := createTestProcessor()
		xmlData := []byte(`<?xml version="1.0"?><Document xmlns="unknown:namespace">content</Document>`)

		_, err := processor.ProcessMessage(xmlData)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "document creation")
	})

	t.Run("ProcessMessage with missing path map", func(t *testing.T) {
		namespaceMap := map[string]models.DocumentFactory{
			"test:namespace:v1": func() models.ISODocument {
				return &MockDocument{Content: "test"}
			},
		}

		versionMap := map[string]TestVersion{
			"test:namespace:v1": TestV1,
		}

		// Empty path maps to trigger missing path map error
		pathMaps := map[TestVersion]map[string]any{}

		processor := NewMessageProcessor[TestMessage, TestVersion](
			namespaceMap,
			versionMap,
			pathMaps,
			[]string{},
		)

		xmlData := []byte(`<?xml version="1.0"?><Document xmlns="test:namespace:v1">content</Document>`)

		_, err := processor.ProcessMessage(xmlData)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "path map lookup")
	})

	t.Run("CreateDocument with unsupported version", func(t *testing.T) {
		processor := createTestProcessor()
		message := TestMessage{
			MessageHeader: MessageHeader{MessageId: "TEST001"},
			TestField:     "test",
		}

		_, err := processor.CreateDocument(message, TestVersion("unsupported"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported version")
	})

	t.Run("CreateDocument with missing factory", func(t *testing.T) {
		// Create processor with version map but no corresponding factory
		namespaceMap := map[string]models.DocumentFactory{} // Empty

		versionMap := map[string]TestVersion{
			"test:namespace:v1": TestV1,
		}

		pathMaps := map[TestVersion]map[string]any{
			TestV1: {
				"Document.content": "TestField",
			},
		}

		processor := NewMessageProcessor[TestMessage, TestVersion](
			namespaceMap,
			versionMap,
			pathMaps,
			[]string{},
		)

		message := TestMessage{
			MessageHeader: MessageHeader{MessageId: "TEST001"},
			TestField:     "test",
		}

		_, err := processor.CreateDocument(message, TestV1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing factory for namespace")
	})

	t.Run("ValidateRequiredFields success", func(t *testing.T) {
		processor := createTestProcessor()
		message := TestMessage{
			MessageHeader: MessageHeader{MessageId: "TEST001"},
			TestField:     "test",
		}

		err := processor.ValidateRequiredFields(message)
		assert.NoError(t, err)
	})

	t.Run("ValidateRequiredFields with missing field", func(t *testing.T) {
		processor := createTestProcessorWithRequiredFields()
		message := TestMessage{
			MessageHeader: MessageHeader{MessageId: ""}, // Missing required field
			TestField:     "test",
		}

		err := processor.ValidateRequiredFields(message)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "MessageId")
	})
}

func TestFieldValidator(t *testing.T) {
	t.Run("ValidateRequired with valid fields", func(t *testing.T) {
		validator := &FieldValidator{
			requiredFields: []string{"MessageId", "TestField"},
		}

		message := TestMessage{
			MessageHeader: MessageHeader{MessageId: "TEST001"},
			TestField:     "test",
		}

		err := validator.ValidateRequired(message)
		assert.NoError(t, err)
	})

	t.Run("ValidateRequired with missing field", func(t *testing.T) {
		validator := &FieldValidator{
			requiredFields: []string{"MessageId", "TestField"},
		}

		message := TestMessage{
			MessageHeader: MessageHeader{MessageId: ""}, // Empty required field
			TestField:     "test",
		}

		err := validator.ValidateRequired(message)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "MessageId")
	})

	t.Run("ValidateRequired with invalid field name", func(t *testing.T) {
		validator := &FieldValidator{
			requiredFields: []string{"NonExistentField"},
		}

		message := TestMessage{
			MessageHeader: MessageHeader{MessageId: "TEST001"},
			TestField:     "test",
		}

		err := validator.ValidateRequired(message)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "NonExistentField")
	})

	t.Run("ValidateRequired with pointer input", func(t *testing.T) {
		validator := &FieldValidator{
			requiredFields: []string{"MessageId"},
		}

		message := &TestMessage{
			MessageHeader: MessageHeader{MessageId: "TEST001"},
			TestField:     "test",
		}

		err := validator.ValidateRequired(message)
		assert.NoError(t, err)
	})
}

func TestErrorHandlingFunctions(t *testing.T) {
	t.Run("HandleDocumentCreationError", func(t *testing.T) {
		originalErr := assert.AnError
		wrappedErr := HandleDocumentCreationError(originalErr)

		assert.Error(t, wrappedErr)
		assert.Contains(t, wrappedErr.Error(), "document creation")
	})

	t.Run("HandleFieldCopyError", func(t *testing.T) {
		originalErr := assert.AnError
		wrappedErr := HandleFieldCopyError("TestField", originalErr)

		assert.Error(t, wrappedErr)
		assert.Contains(t, wrappedErr.Error(), "TestField")
		assert.Contains(t, wrappedErr.Error(), "copy")
	})

	t.Run("HandleVersionLookupError", func(t *testing.T) {
		wrappedErr := HandleVersionLookupError("unknown:namespace")

		assert.Error(t, wrappedErr)
		assert.Contains(t, wrappedErr.Error(), "version lookup")
		assert.Contains(t, wrappedErr.Error(), "unknown:namespace")
	})
}

// Helper functions for tests
func createTestProcessor() *MessageProcessor[TestMessage, TestVersion] {
	return createTestProcessorWithRequiredFields()
}

func createTestProcessorWithRequiredFields() *MessageProcessor[TestMessage, TestVersion] {
	namespaceMap := map[string]models.DocumentFactory{
		"test:namespace:v1": func() models.ISODocument {
			return &MockDocument{Content: "test"}
		},
	}

	versionMap := map[string]TestVersion{
		"test:namespace:v1": TestV1,
	}

	pathMaps := map[TestVersion]map[string]any{
		TestV1: {
			"Document.content": "TestField",
		},
	}

	requiredFields := []string{"MessageId", "TestField"}

	return NewMessageProcessor[TestMessage, TestVersion](
		namespaceMap,
		versionMap,
		pathMaps,
		requiredFields,
	)
}
