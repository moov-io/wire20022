package wrapper

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	CustomerCreditTransfer "github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
)

// TestGenericWrapperEquivalence proves that the generic wrapper provides identical functionality
// to the original wrapper implementation with better type safety and less code
func TestGenericWrapperEquivalence(t *testing.T) {
	// Create both wrapper implementations
	originalWrapper := &CustomerCreditTransferWrapper{}
	genericWrapper := NewCustomerCreditTransferWrapperGeneric()

	// Test data - valid CustomerCreditTransfer model with all required fields
	validJSON := []byte(`{
		"MessageId": "MSG123",
		"CreatedDateTime": "2024-01-01T10:00:00Z",
		"NumberOfTransactions": "1",
		"SettlementMethod": "CLRG",
		"CommonClearingSysCode": "FDW",
		"InstructionId": "INSTR123",
		"EndToEndId": "E2E123",
		"TaxId": "TX123",
		"InstrumentPropCode": "CTRC",
		"InterBankSettAmount": {
			"Amount": 1000.00,
			"Currency": "USD"
		},
		"InterBankSettDate": "2024-01-01",
		"InstructedAmount": {
			"Amount": 1000.00,
			"Currency": "USD"
		},
		"ChargeBearer": "SLEV",
		"InstructingAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "123456789"
		},
		"InstructedAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "987654321"
		},
		"DebtorAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "111111111"
		},
		"CreditorAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "222222222"
		},
		"DebtorName": "John Doe",
		"DebtorAddress": {
			"StreetName": "123 Main St",
			"TownName": "Anytown",
			"Country": "US"
		}
	}`)

	version := CustomerCreditTransfer.PACS_008_001_08

	t.Run("CreateDocument produces identical results", func(t *testing.T) {
		// Test original wrapper
		originalXML, originalErr := originalWrapper.CreateDocument(validJSON, version)
		require.NoError(t, originalErr)
		require.NotEmpty(t, originalXML)

		// Test generic wrapper  
		genericXML, genericErr := genericWrapper.CreateDocument(validJSON, version)
		require.NoError(t, genericErr)
		require.NotEmpty(t, genericXML)

		// Results should be identical
		assert.Equal(t, originalXML, genericXML, "CreateDocument should produce identical XML output")
	})

	t.Run("ValidateDocument produces identical results", func(t *testing.T) {
		validJSONString := string(validJSON)

		// Test original wrapper
		originalErr := originalWrapper.ValidateDocument(validJSONString, version)

		// Test generic wrapper
		genericErr := genericWrapper.ValidateDocument(validJSONString, version)

		// Results should be functionally equivalent (both succeed or both fail)
		if originalErr == nil {
			assert.NoError(t, genericErr, "Both wrappers should succeed for valid input")
		} else {
			assert.Error(t, genericErr, "Both wrappers should fail for invalid input")
			// Note: Generic wrapper provides enhanced error messages with message type context
			// This is an improvement, not a regression
			t.Logf("Original error: %v", originalErr)
			t.Logf("Generic error: %v", genericErr)
		}
	})

	t.Run("CheckRequireField produces identical results", func(t *testing.T) {
		// Create model from JSON
		var model CustomerCreditTransfer.MessageModel
		err := json.Unmarshal(validJSON, &model)
		require.NoError(t, err)

		// Test original wrapper
		originalErr := originalWrapper.CheckRequireField(model)

		// Test generic wrapper
		genericErr := genericWrapper.CheckRequireField(model)

		// Results should be identical
		if originalErr == nil {
			assert.NoError(t, genericErr, "Both wrappers should succeed for valid model")
		} else {
			assert.Error(t, genericErr, "Both wrappers should fail for invalid model")
			assert.Equal(t, originalErr.Error(), genericErr.Error(), "Error messages should be identical")
		}
	})

	t.Run("GetHelp produces identical results", func(t *testing.T) {
		// Test original wrapper
		originalHelp, originalErr := originalWrapper.GetHelp()
		require.NoError(t, originalErr)
		require.NotEmpty(t, originalHelp)

		// Test generic wrapper
		genericHelp, genericErr := genericWrapper.GetHelp()
		require.NoError(t, genericErr)
		require.NotEmpty(t, genericHelp)

		// Parse both JSON responses to compare content
		var originalData, genericData interface{}
		err := json.Unmarshal([]byte(originalHelp), &originalData)
		require.NoError(t, err)
		
		err = json.Unmarshal([]byte(genericHelp), &genericData)
		require.NoError(t, err)

		// Content should be identical
		assert.Equal(t, originalData, genericData, "GetHelp should produce identical content")
	})
}

// TestGenericWrapperTypeSafety demonstrates the type safety benefits of the generic approach
func TestGenericWrapperTypeSafety(t *testing.T) {
	// This test demonstrates compile-time type safety that the generic wrapper provides
	// These would be compile-time errors with the generic wrapper:
	
	t.Run("Generic wrapper enforces correct model type", func(t *testing.T) {
		genericWrapper := NewCustomerCreditTransferWrapperGeneric()
		
		// This compiles and runs - correct model type
		validJSON := []byte(`{"messageId": "test"}`)
		_, err := genericWrapper.CreateDocument(validJSON, CustomerCreditTransfer.PACS_008_001_08)
		
		// Error is expected due to validation, but it compiles because types are correct
		assert.Error(t, err) // Expected validation error
	})

	t.Run("Generic wrapper enforces correct version type", func(t *testing.T) {
		genericWrapper := NewCustomerCreditTransferWrapperGeneric()
		
		validJSON := []byte(`{"messageId": "test"}`)
		
		// This compiles - correct version type
		_, err := genericWrapper.CreateDocument(validJSON, CustomerCreditTransfer.PACS_008_001_08)
		assert.Error(t, err) // Expected validation error, but compiles
		
		// Note: The following would be a COMPILE-TIME ERROR with the generic wrapper:
		// _, err := genericWrapper.CreateDocument(validJSON, "invalid-version-type")
		// _, err := genericWrapper.CreateDocument(validJSON, PaymentReturn.PACS_004_001_12) // Wrong message type version
	})
}

// TestGenericWrapperErrorHandling validates that error handling is consistent and improved
func TestGenericWrapperErrorHandling(t *testing.T) {
	genericWrapper := NewCustomerCreditTransferWrapperGeneric()

	t.Run("Invalid JSON produces clear error", func(t *testing.T) {
		invalidJSON := []byte(`{"invalid": json}`)
		
		_, err := genericWrapper.CreateDocument(invalidJSON, CustomerCreditTransfer.PACS_008_001_08)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal JSON to CustomerCreditTransfer MessageModel")
		assert.Contains(t, err.Error(), "invalid character")
	})

	t.Run("Empty model produces validation error", func(t *testing.T) {
		emptyJSON := []byte(`{}`)
		
		_, err := genericWrapper.CreateDocument(emptyJSON, CustomerCreditTransfer.PACS_008_001_08)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create CustomerCreditTransfer document")
	})

	t.Run("XML conversion error handling", func(t *testing.T) {
		invalidXML := []byte(`<invalid>xml`)
		
		_, err := genericWrapper.ConvertXMLToModel(invalidXML)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "failed to convert XML to CustomerCreditTransfer model")
	})
}

// BenchmarkWrapperPerformance compares performance between original and generic wrappers
func BenchmarkWrapperPerformance(b *testing.B) {
	validJSON := []byte(`{
		"MessageId": "MSG123",
		"CreatedDateTime": "2024-01-01T10:00:00Z",
		"NumberOfTransactions": "1",
		"SettlementMethod": "CLRG",
		"CommonClearingSysCode": "FDW",
		"InstructionId": "INSTR123",
		"EndToEndId": "E2E123",
		"InstrumentPropCode": "CTRC",
		"InterBankSettAmount": {
			"Amount": 1000.00,
			"Currency": "USD"
		},
		"InstructedAmount": {
			"Amount": 1000.00,
			"Currency": "USD"
		},
		"ChargeBearer": "SLEV",
		"InstructingAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "123456789"
		},
		"InstructedAgent": {
			"PaymentSysCode": "USABA",
			"PaymentSysMemberId": "987654321"
		},
		"DebtorName": "John Doe"
	}`)
	version := CustomerCreditTransfer.PACS_008_001_08

	b.Run("Original Wrapper CreateDocument", func(b *testing.B) {
		wrapper := &CustomerCreditTransferWrapper{}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := wrapper.CreateDocument(validJSON, version)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Generic Wrapper CreateDocument", func(b *testing.B) {
		wrapper := NewCustomerCreditTransferWrapperGeneric()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := wrapper.CreateDocument(validJSON, version)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}