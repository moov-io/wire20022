package wrapper

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	EndpointTotalsReport "github.com/moov-io/wire20022/pkg/models/EndpointTotalsReport"
)

type EndpointTotalsReportWrapper struct{}

// CreateDocument generates a camt.052 XML document based on the provided JSON string representation of the MessageModel and version.
// It uses the EndpointTotalsReport.DocumentWith function to create the document structure
// and then marshals it into an indented XML format.
//
// Parameters:
// - modelJson: A JSON string representing the MessageModel containing the data for the document.
// - version: The CAMT_052_001_VERSION specifying the version of the document.
//
// Returns:
// - []byte: The XML representation of the document.
// - error: An error if the document creation, JSON unmarshaling, or XML marshaling fails.
func (w *EndpointTotalsReportWrapper) CreateDocument(modelJson []byte, version EndpointTotalsReport.CAMT_052_001_VERSION) ([]byte, error) {
	// Unmarshal the JSON string into the MessageModel
	var model EndpointTotalsReport.MessageModel
	err := json.Unmarshal(modelJson, &model)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON to MessageModel: %w", err)
	}
	// Create the XML document
	doc, err := EndpointTotalsReport.DocumentWith(model, version)
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

// ValidateDocument validates a camt.052 XML document based on the provided JSON string representation of the MessageModel and version.
// It unmarshals the JSON string into a MessageModel, creates the XML document using the EndpointTotalsReport.DocumentWith function,
// and validates the document structure.
//
// Parameters:
// - modelJson: A JSON string representing the MessageModel containing the data for the document.
// - version: The CAMT_052_001_VERSION specifying the version of the document.
//
// Returns:
// - error: An error if the JSON unmarshaling, document creation, or validation fails.
func (w *EndpointTotalsReportWrapper) ValidateDocument(modelJson string, version EndpointTotalsReport.CAMT_052_001_VERSION) error {
	// Unmarshal the JSON string into the MessageModel
	var model EndpointTotalsReport.MessageModel
	err := json.Unmarshal([]byte(modelJson), &model)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON to MessageModel: %w", err)
	}
	// Create the XML document
	doc, err := EndpointTotalsReport.DocumentWith(model, version)
	if err != nil {
		return fmt.Errorf("failed to create document: %w", err)
	}
	if err := doc.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

// CheckRequireField checks if all required fields in the provided MessageModel are populated.
// It uses the EndpointTotalsReport.CheckRequiredFields function to perform the validation.
//
// Parameters:
// - model: The MessageModel to validate.
//
// Returns:
// - error: An error if any required field is missing; otherwise, nil.
func (w *EndpointTotalsReportWrapper) CheckRequireField(model EndpointTotalsReport.MessageModel) error {
	return EndpointTotalsReport.CheckRequiredFields(model)
}

// ConvertXMLToModel converts a camt.052 XML document into a MessageModel.
// It uses the EndpointTotalsReport.MessageWith function to parse the XML data and populate the MessageModel.
//
// Parameters:
// - xmlData: A byte slice containing the camt.052 XML document to be converted.
//
// Returns:
// - EndpointTotalsReport.MessageModel: The parsed MessageModel.
// - error: An error if the XML parsing fails.
func (w *EndpointTotalsReportWrapper) ConvertXMLToModel(xmlData []byte) (EndpointTotalsReport.MessageModel, error) {
	model, err := EndpointTotalsReport.MessageWith(xmlData)
	if err != nil {
		return EndpointTotalsReport.MessageModel{}, fmt.Errorf("failed to convert XML to model: %w", err)
	}

	return model, nil
}

// GetHelpStringJSON generates a JSON string containing help information for the MessageHelper structure.
//
// Returns:
// - string: A JSON string representation of the MessageHelper structure.
// - error: An error if the JSON marshaling fails.
func (w *EndpointTotalsReportWrapper) GetHelp() (string, error) {
	// Build the MessageHelper structure
	helper := EndpointTotalsReport.BuildMessageHelper()

	// Marshal the structure into a JSON string
	jsonData, err := json.MarshalIndent(helper, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal MessageHelper to JSON: %w", err)
	}

	return string(jsonData), nil
}
