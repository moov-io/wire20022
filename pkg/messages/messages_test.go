package messages_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/moov-io/wire20022/pkg/messages"
	AccountReportingRequestModel "github.com/moov-io/wire20022/pkg/models/AccountReportingRequest"
	ActivityReportModel "github.com/moov-io/wire20022/pkg/models/ActivityReport"
	CustomerCreditTransferModel "github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
)

// Test the new idiomatic API for AccountReportingRequest
func TestAccountReportingRequest(t *testing.T) {
	processor := messages.NewAccountReportingRequest()

	t.Run("CreateDocument with valid JSON", func(t *testing.T) {
		validJSON := []byte(`{
			"messageId": "ARR123456789012345678901",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"reportRequestId": "ABAR",
			"requestedMsgNameId": "camt.052.001.08",
			"accountOtherId": "ACC123456789"
		}`)

		xmlData, err := processor.CreateDocument(validJSON, AccountReportingRequestModel.CAMT_060_001_05)
		assert.NoError(t, err)
		assert.NotEmpty(t, xmlData)

		// Verify it's valid XML
		var doc interface{}
		err = xml.Unmarshal(xmlData, &doc)
		assert.NoError(t, err)
	})

	t.Run("GetHelp returns valid JSON", func(t *testing.T) {
		helpJSON, err := processor.GetHelp()
		assert.NoError(t, err)
		assert.NotEmpty(t, helpJSON)

		// Verify it's valid JSON
		var help interface{}
		err = json.Unmarshal([]byte(helpJSON), &help)
		assert.NoError(t, err)
	})

	t.Run("Validate with missing required fields", func(t *testing.T) {
		invalidJSON := `{"messageId": ""}`
		err := processor.ValidateDocument(invalidJSON, AccountReportingRequestModel.CAMT_060_001_05)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "AccountReportingRequest")
	})
}

// Test the new idiomatic API for ActivityReport
func TestActivityReport(t *testing.T) {
	processor := messages.NewActivityReport()

	t.Run("CreateDocument with valid JSON", func(t *testing.T) {
		validJSON := []byte(`{
			"messageId": "AR1234567890123456789012",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"reportCreateDateTime": "2024-01-01T10:00:00Z",
			"reportId": "RPT123",
			"accountId": "ACC123456789"
		}`)

		xmlData, err := processor.CreateDocument(validJSON, ActivityReportModel.CAMT_052_001_05)
		assert.NoError(t, err)
		assert.NotEmpty(t, xmlData)

		// Verify it's valid XML
		var doc interface{}
		err = xml.Unmarshal(xmlData, &doc)
		assert.NoError(t, err)
	})

	t.Run("Enhanced error messages include message type", func(t *testing.T) {
		invalidJSON := `{"messageId": ""}`
		err := processor.ValidateDocument(invalidJSON, ActivityReportModel.CAMT_052_001_05)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ActivityReport")
	})
}

// Test the new idiomatic API for CustomerCreditTransfer
func TestCustomerCreditTransfer(t *testing.T) {
	processor := messages.NewCustomerCreditTransfer()

	t.Run("CreateDocument with valid JSON", func(t *testing.T) {
		validJSON := []byte(`{
			"messageId": "20240101B1QDRCQR000001",
			"createdDateTime": "2024-01-01T10:00:00Z",
			"numberOfTransactions": "1",
			"settlementMethod": "CLRG",
			"commonClearingSysCode": "FDW",
			"instructionId": "INSTR001",
			"endToEndId": "E2E001",
			"uniqueEndToEndTransactionRef": "12345678-1234-1234-1234-123456789012",
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

		xmlData, err := processor.CreateDocument(validJSON, CustomerCreditTransferModel.PACS_008_001_08)
		assert.NoError(t, err)
		assert.NotEmpty(t, xmlData)

		// Verify it's valid XML
		var doc interface{}
		err = xml.Unmarshal(xmlData, &doc)
		assert.NoError(t, err)
	})
}

// Test that all message processors can be created successfully
func TestAllMessageProcessors(t *testing.T) {
	processors := map[string]interface{}{
		"AccountReportingRequest":     messages.NewAccountReportingRequest(),
		"ActivityReport":              messages.NewActivityReport(),
		"ConnectionCheck":             messages.NewConnectionCheck(),
		"CustomerCreditTransfer":      messages.NewCustomerCreditTransfer(),
		"DrawdownRequest":             messages.NewDrawdownRequest(),
		"DrawdownResponse":            messages.NewDrawdownResponse(),
		"EndpointDetailsReport":       messages.NewEndpointDetailsReport(),
		"EndpointGapReport":           messages.NewEndpointGapReport(),
		"EndpointTotalsReport":        messages.NewEndpointTotalsReport(),
		"FedwireFundsAcknowledgement": messages.NewFedwireFundsAcknowledgement(),
		"FedwireFundsPaymentStatus":   messages.NewFedwireFundsPaymentStatus(),
		"FedwireFundsSystemResponse":  messages.NewFedwireFundsSystemResponse(),
		"Master":                      messages.NewMaster(),
		"PaymentReturn":               messages.NewPaymentReturn(),
		"PaymentStatusRequest":        messages.NewPaymentStatusRequest(),
		"ReturnRequestResponse":       messages.NewReturnRequestResponse(),
	}

	for name, processor := range processors {
		t.Run(name, func(t *testing.T) {
			assert.NotNil(t, processor, "Processor %s should be created successfully", name)
		})
	}
}

// Benchmark the new API to ensure performance is acceptable
func BenchmarkCustomerCreditTransfer(b *testing.B) {
	processor := messages.NewCustomerCreditTransfer()
	validJSON := []byte(`{
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := processor.CreateDocument(validJSON, CustomerCreditTransferModel.PACS_008_001_08)
		if err != nil {
			b.Fatal(err)
		}
	}
}
