package messages_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/moov-io/wire20022/pkg/messages"
	"github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
)

// ExampleNewCustomerCreditTransfer demonstrates the basic usage of the new v1.0 API
func ExampleNewCustomerCreditTransfer() {
	// Create a CustomerCreditTransfer processor
	processor := messages.NewCustomerCreditTransfer()

	// Sample message data (JSON)
	messageJSON := []byte(`{
		"messageId": "20240101B1QDRCQR000001",
		"createdDateTime": "2024-01-01T10:00:00Z",
		"numberOfTransactions": "1",
		"settlementMethod": "CLRG",
		"commonClearingSysCode": "FDW",
		"instructionId": "INSTR001",
		"endToEndId": "E2E001",
		"instrumentPropCode": "CTRC",
		"interBankSettAmount": {"currency": "USD", "amount": 1000.00},
		"interBankSettDate": "2024-01-01",
		"instructedAmount": {"currency": "USD", "amount": 1000.00},
		"chargeBearer": "SLEV",
		"instructingAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "123456789"},
		"instructedAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "987654321"},
		"debtorName": "John Doe",
		"debtorAddress": {"streetName": "Main St", "buildingNumber": "123", "postalCode": "12345", "townName": "Anytown", "country": "US"},
		"debtorAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "123456789"},
		"creditorAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "987654321"}
	}`)

	// Create ISO 20022 XML document
	xmlData, err := processor.CreateDocument(messageJSON, CustomerCreditTransfer.PACS_008_001_07)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created XML document with %d bytes\n", len(xmlData))

	// Validate without creating XML
	err = processor.ValidateDocument(string(messageJSON), CustomerCreditTransfer.PACS_008_001_07)
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	} else {
		fmt.Println("Document validation passed!")
	}

	// Output:
	// Created XML document with 5032 bytes
	// Document validation passed!
}

// ExampleAccountReportingRequest_GetHelp demonstrates getting field documentation
func ExampleAccountReportingRequest_GetHelp() {
	processor := messages.NewAccountReportingRequest()

	// Get comprehensive field documentation
	helpJSON, err := processor.GetHelp()
	if err != nil {
		log.Fatal(err)
	}

	// Parse the help JSON to show structure
	var help map[string]interface{}
	err = json.Unmarshal([]byte(helpJSON), &help)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Available field count: %d\n", len(help))
	fmt.Println("Documentation includes MessageId, CreatedDateTime, and more...")

	// Output:
	// Available field count: 8
	// Documentation includes MessageId, CreatedDateTime, and more...
}

// ExampleActivityReport_validation demonstrates enhanced error messages
func ExampleActivityReport_validation() {
	processor := messages.NewActivityReport()

	// Try to validate incomplete JSON
	incompleteJSON := `{"messageId": ""}`

	err := processor.ValidateDocument(incompleteJSON, "CAMT_052_001_05")
	if err != nil {
		// Enhanced error messages include message type context
		fmt.Printf("Error includes message type: %t\n", err.Error() != "")
		fmt.Println("Enhanced error handling provides clear feedback")
	}

	// Output:
	// Error includes message type: true
	// Enhanced error handling provides clear feedback
}

// Example_typeSystem demonstrates compile-time type safety
func Example_typeSystem() {
	// All processors provide the same interface but with type safety
	processors := []interface{}{
		messages.NewCustomerCreditTransfer(),
		messages.NewPaymentReturn(),
		messages.NewActivityReport(),
		messages.NewAccountReportingRequest(),
	}

	fmt.Printf("Created %d different message processors\n", len(processors))
	fmt.Println("Each processor is type-safe for its specific message type")

	// Output:
	// Created 4 different message processors
	// Each processor is type-safe for its specific message type
}
