package CustomerCreditTransfer_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
)

// TestReadXML tests the idiomatic ReadXML method
func TestReadXML(t *testing.T) {
	// Load sample XML file
	xmlFile, err := os.Open("./swiftSample/CustomerCreditTransfer_Scenario1_Step1_pacs.008")
	require.NoError(t, err)
	defer xmlFile.Close()

	// Use the new idiomatic API
	var model CustomerCreditTransfer.MessageModel
	err = model.ReadXML(xmlFile)
	assert.NoError(t, err)

	// Verify some key fields were parsed
	assert.NotEmpty(t, model.MessageId)
	assert.NotEmpty(t, model.CreatedDateTime)
	assert.NotEmpty(t, model.NumberOfTransactions)
}

// TestWriteXML tests the idiomatic WriteXML method
func TestWriteXML(t *testing.T) {
	// Create a test model
	model := CustomerCreditTransfer.CustomerCreditTransferDataModel()

	// Write to buffer with default version (latest)
	var buf bytes.Buffer
	err := model.WriteXML(&buf)
	assert.NoError(t, err)

	// Verify XML was written
	xmlOutput := buf.String()
	assert.Contains(t, xmlOutput, "<?xml")
	assert.Contains(t, xmlOutput, "<Document")
	assert.Contains(t, xmlOutput, model.MessageId)

	// Verify we can read it back
	var parsedModel CustomerCreditTransfer.MessageModel
	err = parsedModel.ReadXML(strings.NewReader(xmlOutput))
	assert.NoError(t, err)
	assert.Equal(t, model.MessageId, parsedModel.MessageId)
}

// TestWriteXMLWithVersion tests writing with specific version
func TestWriteXMLWithVersion(t *testing.T) {
	model := CustomerCreditTransfer.CustomerCreditTransferDataModel()

	// Test with specific older version
	var buf bytes.Buffer
	err := model.WriteXML(&buf, CustomerCreditTransfer.PACS_008_001_08)
	assert.NoError(t, err)

	// Verify version namespace in output
	xmlOutput := buf.String()
	assert.Contains(t, xmlOutput, "pacs.008.001.08")
}

// TestParseXML tests the new ParseXML function
func TestParseXML(t *testing.T) {
	// Load sample XML
	xmlData, err := os.ReadFile("./swiftSample/CustomerCreditTransfer_Scenario1_Step1_pacs.008")
	require.NoError(t, err)

	// Use new ParseXML function
	model, err := CustomerCreditTransfer.ParseXML(xmlData)
	assert.NoError(t, err)
	assert.NotNil(t, model)
	assert.NotEmpty(t, model.MessageId)
}

// TestIdiomaticWorkflow tests the complete idiomatic workflow
func TestIdiomaticWorkflow(t *testing.T) {
	// 1. Read from XML file
	xmlFile, err := os.Open("./swiftSample/CustomerCreditTransfer_Scenario1_Step1_pacs.008")
	require.NoError(t, err)
	defer xmlFile.Close()

	var payment CustomerCreditTransfer.MessageModel
	err = payment.ReadXML(xmlFile)
	require.NoError(t, err)

	// 2. Modify the payment
	originalAmount := payment.InterBankSettAmount.Amount
	payment.InterBankSettAmount.Amount = originalAmount + 100

	// 3. Write to new file
	var output bytes.Buffer
	err = payment.WriteXML(&output)
	require.NoError(t, err)

	// 4. Verify the modification persisted
	var verifyPayment CustomerCreditTransfer.MessageModel
	err = verifyPayment.ReadXML(bytes.NewReader(output.Bytes()))
	require.NoError(t, err)
	assert.Equal(t, originalAmount+100, verifyPayment.InterBankSettAmount.Amount)
}
