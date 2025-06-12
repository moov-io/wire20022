package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	DrawdownRequest "github.com/moov-io/wire20022/pkg/models/DrawdownRequest"
)

type DrawdownRequestWrapper struct{}

// CreateDocument generates a pain.013 XML document based on the provided JSON string representation of the MessageModel and version.
// It uses the DrawdownRequest.DocumentWith function to create the document structure
// and then marshals it into an indented XML format.
//
// Parameters:
// - modelJson: A JSON string representing the MessageModel containing the data for the document.
// - version: The PAIN_013_001_VESION specifying the version of the document.
//
// Returns:
// - []byte: The XML representation of the document.
// - error: An error if the document creation, JSON unmarshaling, or XML marshaling fails.
func (w *DrawdownRequestWrapper) CreateDocument(modelJson []byte, version DrawdownRequest.PAIN_013_001_VESION) ([]byte, error) {
	// Unmarshal the JSON string into the MessageModel
	var model DrawdownRequest.MessageModel
	err := json.Unmarshal(modelJson, &model)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to MessageModel: %w", err)
	}
	// Create the XML document
	doc, err := DrawdownRequest.DocumentWith(model, version)
	if err != nil {
		return nil, fmt.Errorf("failed to create document: %w", err)
	}
	// Convert the document to XML
	xmlData, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal document to XML: %w", err)
	}

	return xmlData, nil
}

// ValidateDocument validates a pain.013 XML document based on the provided JSON string representation of the MessageModel and version.
// It unmarshals the JSON string into a MessageModel, creates the XML document using the DrawdownRequest.DocumentWith function,
// and validates the document structure.
//
// Parameters:
// - modelJson: A JSON string representing the MessageModel containing the data for the document.
// - version: The PAIN_013_001_VESION specifying the version of the document.
//
// Returns:
// - error: An error if the JSON unmarshaling, document creation, or validation fails.
func (w *DrawdownRequestWrapper) ValidateDocument(modelJson string, version DrawdownRequest.PAIN_013_001_VESION) error {
	// Unmarshal the JSON string into the MessageModel
	var model DrawdownRequest.MessageModel
	err := json.Unmarshal([]byte(modelJson), &model)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON to MessageModel: %w", err)
	}
	// Create the XML document
	doc, err := DrawdownRequest.DocumentWith(model, version)
	if err != nil {
		return fmt.Errorf("failed to create document: %w", err)
	}
	if err := doc.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

// CheckRequireField checks if all required fields in the provided MessageModel are populated.
// It uses the DrawdownRequest.CheckRequiredFields function to perform the validation.
//
// Parameters:
// - model: The MessageModel to validate.
//
// Returns:
// - error: An error if any required field is missing; otherwise, nil.
func (w *DrawdownRequestWrapper) CheckRequireField(model DrawdownRequest.MessageModel) error {
	return DrawdownRequest.CheckRequiredFields(model)
}

// ConvertXMLToModel converts a pain.013 XML document into a MessageModel.
// It uses the DrawdownRequest.MessageWith function to parse the XML data and populate the MessageModel.
//
// Parameters:
// - xmlData: A byte slice containing the pain.013 XML document to be converted.
//
// Returns:
// - DrawdownRequest.MessageModel: The parsed MessageModel.
// - error: An error if the XML parsing fails.
func (w *DrawdownRequestWrapper) ConvertXMLToModel(xmlData []byte) (DrawdownRequest.MessageModel, error) {
	model, err := DrawdownRequest.MessageWith(xmlData)
	if err != nil {
		return DrawdownRequest.MessageModel{}, fmt.Errorf("failed to convert XML to model: %w", err)
	}

	return model, nil
}

// GetHelpStringJSON generates a JSON string containing help information for the MessageHelper structure.
//
// Returns:
// - string: A JSON string representation of the MessageHelper structure.
// - error: An error if the JSON marshaling fails.
func (w *DrawdownRequestWrapper) GetHelp() (string, error) {
	// Build the MessageHelper structure
	helper := DrawdownRequest.BuildMessageHelper()

	// Marshal the structure into a JSON string
	jsonData, err := json.MarshalIndent(helper, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal MessageHelper to JSON: %w", err)
	}

	return string(jsonData), nil
}
