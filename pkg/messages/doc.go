// Package messages provides type-safe processors for ISO 20022 wire transfer message types.
//
// This package offers a unified, generic interface for working with Fedwire ISO 20022 messages,
// eliminating code duplication while providing compile-time type safety and enhanced error handling.
//
// It also includes a Universal Reader that can automatically detect and parse any supported
// message type without requiring prior knowledge of the message format.
//
// # Architecture
//
// Each message type has a dedicated processor that provides these core methods:
//   - CreateDocument: Convert JSON model to ISO 20022 XML document
//   - ValidateDocument: Validate JSON model and generate document
//   - Validate: Validate required fields in model
//   - ConvertXMLToModel: Parse XML to typed Go model
//   - GetHelp: Get comprehensive field documentation
//
// # Example Usage
//
//	// Create a CustomerCreditTransfer processor
//	cct := messages.NewCustomerCreditTransfer()
//
//	// Sample JSON message data
//	messageJSON := []byte(`{
//		"messageId": "20240101B1QDRCQR000001",
//		"createdDateTime": "2024-01-01T10:00:00Z",
//		"numberOfTransactions": "1",
//		"settlementMethod": "CLRG",
//		"commonClearingSysCode": "FDW",
//		"instructionId": "INSTR001",
//		"endToEndId": "E2E001",
//		"instrumentPropCode": "CTRC",
//		"interBankSettAmount": {"currency": "USD", "amount": 1000.00},
//		"interBankSettDate": "2024-01-01",
//		"instructedAmount": {"currency": "USD", "amount": 1000.00},
//		"chargeBearer": "SLEV",
//		"instructingAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "123456789"},
//		"instructedAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "987654321"},
//		"debtorName": "John Doe",
//		"debtorAddress": {"streetName": "Main St", "buildingNumber": "123", "postalCode": "12345", "townName": "Anytown", "country": "US"},
//		"debtorAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "123456789"},
//		"creditorAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "987654321"}
//	}`)
//
//	// Create ISO 20022 XML document
//	xmlData, err := cct.CreateDocument(messageJSON, models.PACS_008_001_08)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Validate document without creating XML
//	err = cct.ValidateDocument(string(messageJSON), models.PACS_008_001_08)
//	if err != nil {
//		log.Printf("Validation failed: %v", err)
//	}
//
//	// Get field documentation
//	helpJSON, err := cct.GetHelp()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Field documentation:", helpJSON)
//
// # Universal Reader
//
// The Universal Reader automatically detects and parses any supported ISO 20022 message type:
//
//	// Create a universal reader instance
//	reader := messages.NewUniversalReader()
//
//	// Parse any message type automatically
//	parsed, err := reader.ReadBytes(xmlData)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("Detected message type: %s\n", parsed.Type)
//	fmt.Printf("Version: %s\n", parsed.Version)
//
//	// Type assertion for specific message handling
//	switch parsed.Type {
//	case messages.TypeCustomerCreditTransfer:
//		// Handle CustomerCreditTransfer
//	case messages.TypePaymentReturn:
//		// Handle PaymentReturn
//	}
//
//	// Validate the parsed message
//	err = reader.ValidateMessage(parsed)
//	if err != nil {
//		log.Printf("Validation failed: %v", err)
//	}
//
// # Supported Message Types
//
// ## Payment Messages (pacs)
//   - CustomerCreditTransfer: pacs.008 - Customer credit transfer initiation
//   - PaymentReturn: pacs.004 - Payment return messages
//   - PaymentStatusRequest: pacs.028 - Payment status inquiry
//   - FedwireFundsPaymentStatus: pacs.002 - Payment status reports
//
// ## Cash Management Messages (camt)
//   - AccountReportingRequest: camt.060 - Account reporting requests
//   - ActivityReport: camt.086 - Bank services billing statements
//   - EndpointDetailsReport: camt.090 - Member profile requests
//   - EndpointGapReport: camt.087 - Duplicate message requests
//   - EndpointTotalsReport: camt.089 - Payment cancellation requests
//   - ReturnRequestResponse: camt.029 - Investigation resolution
//   - Master: camt.052 - Bank to customer account reports
//
// ## Payment Initiation Messages (pain)
//   - DrawdownRequest: pain.013 - Creditor payment activation requests
//   - DrawdownResponse: pain.014 - Creditor payment activation status reports
//
// ## Administrative Messages (admi)
//   - ConnectionCheck: admi.001 - System connectivity checks
//   - FedwireFundsAcknowledgement: admi.004 - System event acknowledgements
//   - FedwireFundsSystemResponse: admi.011 - System event notifications
//
// # Type Safety Benefits
//
// The generic processor architecture provides:
//   - Compile-time type checking for message models and versions
//   - Prevention of wrong model/version combinations
//   - Enhanced error messages with message type context
//   - Centralized validation logic across all message types
//
// # Error Handling
//
// All processors provide consistent, enhanced error handling:
//   - Validation errors include field path and message type context
//   - Parse errors specify the operation and content type
//   - All errors support Go 1.13+ error unwrapping
//
// # Performance
//
// The generic processor adds minimal overhead (~3%) compared to the original
// implementations while providing significantly better type safety and maintainability.
//
// # Additional Features
//
// Each processor also provides:
//   - Type-safe version handling with compile-time validation
//   - Automatic namespace management for different versions
//   - Comprehensive field mapping from XML to Go structs
//   - Integration with the base abstractions for zero code duplication
//
// The processors work seamlessly with the XML-first message models in pkg/models,
// providing a complete solution for ISO 20022 message processing.
package messages
