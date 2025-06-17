package messages

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/moov-io/wire20022/pkg/models"
)

// MessageWrapper provides a generic, type-safe processor for ISO 20022 message processing
// M: Message model type (e.g., CustomerCreditTransfer.MessageModel)
// V: Version type (e.g., CustomerCreditTransfer.PACS_008_001_VERSION)
type MessageWrapper[M any, V comparable] struct {
	name            string
	documentCreator func(M, V) (models.ISODocument, error)
	fieldValidator  func(M) error
	helpBuilder     func() any
	xmlConverter    func([]byte) (M, error)
}

// NewMessageWrapper creates a new generic message processor with type safety
func NewMessageWrapper[M any, V comparable](
	name string,
	documentCreator func(M, V) (models.ISODocument, error),
	fieldValidator func(M) error,
	helpBuilder func() any,
	xmlConverter func([]byte) (M, error),
) *MessageWrapper[M, V] {
	return &MessageWrapper[M, V]{
		name:            name,
		documentCreator: documentCreator,
		fieldValidator:  fieldValidator,
		helpBuilder:     helpBuilder,
		xmlConverter:    xmlConverter,
	}
}

// CreateDocument converts JSON model to XML document with type safety
func (w *MessageWrapper[M, V]) CreateDocument(modelJson []byte, version V) ([]byte, error) {
	// Unmarshal JSON to typed model
	var model M
	if err := json.Unmarshal(modelJson, &model); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to %s MessageModel: %w", w.name, err)
	}

	// Create document using type-safe document creator
	doc, err := w.documentCreator(model, version)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s document: %w", w.name, err)
	}

	// Marshal to XML
	xmlData, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %s document to XML: %w", w.name, err)
	}

	return xmlData, nil
}

// ValidateDocument validates JSON model without creating XML document
func (w *MessageWrapper[M, V]) ValidateDocument(modelJson string, version V) error {
	// Unmarshal JSON to typed model
	var model M
	if err := json.Unmarshal([]byte(modelJson), &model); err != nil {
		return fmt.Errorf("failed to unmarshal JSON to %s MessageModel: %w", w.name, err)
	}

	// Create document to trigger validation
	_, err := w.documentCreator(model, version)
	if err != nil {
		return fmt.Errorf("failed to create %s document: %w", w.name, err)
	}

	return nil
}

// Validate validates required fields using the configured validator
func (w *MessageWrapper[M, V]) Validate(model M) error {
	if w.fieldValidator == nil {
		return fmt.Errorf("no field validator configured for %s", w.name)
	}

	return w.fieldValidator(model)
}

// ConvertXMLToModel converts XML to typed model with validation
func (w *MessageWrapper[M, V]) ConvertXMLToModel(xmlData []byte) (M, error) {
	var zeroModel M

	if w.xmlConverter == nil {
		return zeroModel, fmt.Errorf("no XML converter configured for %s", w.name)
	}

	// Convert XML to model
	model, err := w.xmlConverter(xmlData)
	if err != nil {
		return zeroModel, fmt.Errorf("failed to convert XML to %s model: %w", w.name, err)
	}

	// Validate the converted model
	if w.fieldValidator != nil {
		if err := w.fieldValidator(model); err != nil {
			return zeroModel, fmt.Errorf("validation failed for converted %s model: %w", w.name, err)
		}
	}

	return model, nil
}

// GetHelp returns help information for the message type
func (w *MessageWrapper[M, V]) GetHelp() (string, error) {
	if w.helpBuilder == nil {
		return "", fmt.Errorf("no help builder configured for %s", w.name)
	}

	helpData := w.helpBuilder()
	helpJson, err := json.MarshalIndent(helpData, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal %s help data to JSON: %w", w.name, err)
	}

	return string(helpJson), nil
}
