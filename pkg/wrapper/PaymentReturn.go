package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	PaymentReturn "github.com/moov-io/wire20022/pkg/models/PaymentReturn"
)

type PaymentReturnWrapper struct{}

// CreateDocument generates a pacs.004 XML document based on the provided JSON string representation of the MessageModel and version.
// It uses the PaymentReturn.DocumentWith function to create the document structure
// and then marshals it into an indented XML format.
//
// Parameters:
// - modelJson: A JSON string representing the MessageModel containing the data for the document.
// - version: The PACS_004_001_VESION specifying the version of the document.
//
// Returns:
// - []byte: The XML representation of the document.
// - error: An error if the document creation, JSON unmarshaling, or XML marshaling fails.
func (w *PaymentReturnWrapper) CreateDocument(modelJson []byte, version PaymentReturn.PACS_004_001_VESION) ([]byte, error) {
	// Unmarshal the JSON string into the MessageModel
	var model PaymentReturn.MessageModel
	err := json.Unmarshal(modelJson, &model)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to MessageModel: %w", err)
	}
	// Create the XML document
	doc, err := PaymentReturn.DocumentWith(model, version)
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

// ValidateDocument validates a pacs.004 XML document based on the provided JSON string representation of the MessageModel and version.
// It unmarshals the JSON string into a MessageModel, creates the XML document using the PaymentReturn.DocumentWith function,
// and validates the document structure.
//
// Parameters:
// - modelJson: A JSON string representing the MessageModel containing the data for the document.
// - version: The PACS_004_001_VESION specifying the version of the document.
//
// Returns:
// - error: An error if the JSON unmarshaling, document creation, or validation fails.
func (w *PaymentReturnWrapper) ValidateDocument(modelJson string, version PaymentReturn.PACS_004_001_VESION) error {
	// Unmarshal the JSON string into the MessageModel
	var model PaymentReturn.MessageModel
	err := json.Unmarshal([]byte(modelJson), &model)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON to MessageModel: %w", err)
	}
	// Create the XML document
	doc, err := PaymentReturn.DocumentWith(model, version)
	if err != nil {
		return fmt.Errorf("failed to create document: %w", err)
	}
	if err := doc.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

// CheckRequireField checks if all required fields in the provided MessageModel are populated.
// It uses the PaymentReturn.CheckRequiredFields function to perform the validation.
//
// Parameters:
// - model: The MessageModel to validate.
//
// Returns:
// - error: An error if any required field is missing; otherwise, nil.
func (w *PaymentReturnWrapper) CheckRequireField(model PaymentReturn.MessageModel) error {
	return PaymentReturn.CheckRequiredFields(model)
}

// ConvertXMLToModel converts a pacs.004 XML document into a MessageModel.
// It uses the PaymentReturn.MessageWith function to parse the XML data and populate the MessageModel.
//
// Parameters:
// - xmlData: A byte slice containing the pacs.004 XML document to be converted.
//
// Returns:
// - PaymentReturn.MessageModel: The parsed MessageModel.
// - error: An error if the XML parsing fails.
func (w *PaymentReturnWrapper) ConvertXMLToModel(xmlData []byte) (PaymentReturn.MessageModel, error) {
	model, err := PaymentReturn.MessageWith(xmlData)
	if err != nil {
		return PaymentReturn.MessageModel{}, fmt.Errorf("failed to convert XML to model: %w", err)
	}

	return model, nil
}

// GetHelpStringJSON generates a JSON string containing help information for the MessageHelper structure.
//
// Returns:
// - string: A JSON string representation of the MessageHelper structure.
// - error: An error if the JSON marshaling fails.
func (w *PaymentReturnWrapper) GetHelp() (string, error) {
	// Build the MessageHelper structure
	helper := PaymentReturn.BuildMessageHelper()

	// Marshal the structure into a JSON string
	jsonData, err := json.MarshalIndent(helper, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal MessageHelper to JSON: %w", err)
	}

	return string(jsonData), nil
}
