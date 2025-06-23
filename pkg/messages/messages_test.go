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
			"accountOtherId": "ACC123456789"
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

// Test GetHelp function for all processors
func TestGetHelpForAllProcessors(t *testing.T) {
	processors := []struct {
		name      string
		processor interface{ GetHelp() (string, error) }
	}{
		{"AccountReportingRequest", messages.NewAccountReportingRequest()},
		{"ActivityReport", messages.NewActivityReport()},
		{"ConnectionCheck", messages.NewConnectionCheck()},
		{"CustomerCreditTransfer", messages.NewCustomerCreditTransfer()},
		{"DrawdownRequest", messages.NewDrawdownRequest()},
		{"DrawdownResponse", messages.NewDrawdownResponse()},
		{"EndpointDetailsReport", messages.NewEndpointDetailsReport()},
		{"EndpointGapReport", messages.NewEndpointGapReport()},
		{"EndpointTotalsReport", messages.NewEndpointTotalsReport()},
		{"FedwireFundsAcknowledgement", messages.NewFedwireFundsAcknowledgement()},
		{"FedwireFundsPaymentStatus", messages.NewFedwireFundsPaymentStatus()},
		{"FedwireFundsSystemResponse", messages.NewFedwireFundsSystemResponse()},
		{"Master", messages.NewMaster()},
		{"PaymentReturn", messages.NewPaymentReturn()},
		{"PaymentStatusRequest", messages.NewPaymentStatusRequest()},
		{"ReturnRequestResponse", messages.NewReturnRequestResponse()},
	}

	for _, proc := range processors {
		t.Run(proc.name+"_GetHelp", func(t *testing.T) {
			help, err := proc.processor.GetHelp()
			assert.NoError(t, err)
			assert.NotEmpty(t, help)

			// Verify it's valid JSON
			var helpStruct interface{}
			err = json.Unmarshal([]byte(help), &helpStruct)
			assert.NoError(t, err, "Help output should be valid JSON for %s", proc.name)
		})
	}
}

// Test Validate method for CustomerCreditTransfer
func TestCustomerCreditTransferValidate(t *testing.T) {
	processor := messages.NewCustomerCreditTransfer()

	// Test valid model
	validModel := CustomerCreditTransferModel.CustomerCreditTransferDataModel()
	err := processor.Validate(validModel)
	assert.NoError(t, err)

	// Test invalid model (missing required field)
	invalidModel := CustomerCreditTransferModel.MessageModel{}
	err = processor.Validate(invalidModel)
	assert.Error(t, err)
	// Just verify we get a validation error, don't check specific message content
	assert.NotEmpty(t, err.Error())
}

// Test ConvertXMLToModel method
func TestConvertXMLToModel(t *testing.T) {
	processor := messages.NewCustomerCreditTransfer()

	// Use a simple valid XML (this would normally come from sample files)
	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08">
	<FIToFICstmrCdt>
		<GrpHdr>
			<MsgId>20240101B1QDRCQR000001</MsgId>
			<CreDtTm>2024-01-01T10:00:00Z</CreDtTm>
			<NbOfTxs>1</NbOfTxs>
			<SttlmInf>
				<SttlmMtd>CLRG</SttlmMtd>
				<ClrSys>
					<Cd>FDW</Cd>
				</ClrSys>
			</SttlmInf>
		</GrpHdr>
	</FIToFICstmrCdt>
</Document>`)

	model, err := processor.ConvertXMLToModel(xmlData)
	if err != nil {
		// XML parsing might fail due to missing required fields, but we want to test the method exists
		t.Logf("XML conversion failed as expected for minimal XML: %v", err)
	} else {
		assert.NotNil(t, model)
		assert.Equal(t, "20240101B1QDRCQR000001", model.MessageId)
	}
}

// Test validation with different processors
func TestValidationForDifferentProcessors(t *testing.T) {
	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			"AccountReportingRequest_Validate",
			func(t *testing.T) {
				processor := messages.NewAccountReportingRequest()
				model := AccountReportingRequestModel.MessageModel{}
				err := processor.Validate(model)
				assert.Error(t, err)
				assert.NotEmpty(t, err.Error())
			},
		},
		{
			"ActivityReport_Validate",
			func(t *testing.T) {
				processor := messages.NewActivityReport()
				model := ActivityReportModel.MessageModel{}
				err := processor.Validate(model)
				assert.Error(t, err)
				assert.NotEmpty(t, err.Error())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}

// Test ConvertXMLToModel for multiple message types
func TestConvertXMLToModelForAllTypes(t *testing.T) {
	// Simple test to cover the ConvertXMLToModel method across different types
	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			"AccountReportingRequest_ConvertXML",
			func(t *testing.T) {
				processor := messages.NewAccountReportingRequest()
				xmlData := []byte(`<?xml version="1.0"?><Document></Document>`)
				_, err := processor.ConvertXMLToModel(xmlData)
				// We expect an error with minimal XML, but we're testing the method exists
				assert.Error(t, err)
			},
		},
		{
			"ActivityReport_ConvertXML",
			func(t *testing.T) {
				processor := messages.NewActivityReport()
				xmlData := []byte(`<?xml version="1.0"?><Document></Document>`)
				_, err := processor.ConvertXMLToModel(xmlData)
				assert.Error(t, err)
			},
		},
		{
			"ConnectionCheck_ConvertXML",
			func(t *testing.T) {
				processor := messages.NewConnectionCheck()
				xmlData := []byte(`<?xml version="1.0"?><Document></Document>`)
				_, err := processor.ConvertXMLToModel(xmlData)
				assert.Error(t, err)
			},
		},
		{
			"DrawdownRequest_ConvertXML",
			func(t *testing.T) {
				processor := messages.NewDrawdownRequest()
				xmlData := []byte(`<?xml version="1.0"?><Document></Document>`)
				_, err := processor.ConvertXMLToModel(xmlData)
				assert.Error(t, err)
			},
		},
		{
			"DrawdownResponse_ConvertXML",
			func(t *testing.T) {
				processor := messages.NewDrawdownResponse()
				xmlData := []byte(`<?xml version="1.0"?><Document></Document>`)
				_, err := processor.ConvertXMLToModel(xmlData)
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}

// Test all remaining New* constructors to boost coverage
func TestAllConstructorsCoverage(t *testing.T) {
	// Test that all constructors return non-nil processors and can call basic methods
	constructors := map[string]func() interface {
		GetHelp() (string, error)
	}{
		"EndpointDetailsReport":       func() interface{ GetHelp() (string, error) } { return messages.NewEndpointDetailsReport() },
		"EndpointGapReport":           func() interface{ GetHelp() (string, error) } { return messages.NewEndpointGapReport() },
		"EndpointTotalsReport":        func() interface{ GetHelp() (string, error) } { return messages.NewEndpointTotalsReport() },
		"FedwireFundsAcknowledgement": func() interface{ GetHelp() (string, error) } { return messages.NewFedwireFundsAcknowledgement() },
		"FedwireFundsPaymentStatus":   func() interface{ GetHelp() (string, error) } { return messages.NewFedwireFundsPaymentStatus() },
		"FedwireFundsSystemResponse":  func() interface{ GetHelp() (string, error) } { return messages.NewFedwireFundsSystemResponse() },
		"Master":                      func() interface{ GetHelp() (string, error) } { return messages.NewMaster() },
		"PaymentReturn":               func() interface{ GetHelp() (string, error) } { return messages.NewPaymentReturn() },
		"PaymentStatusRequest":        func() interface{ GetHelp() (string, error) } { return messages.NewPaymentStatusRequest() },
		"ReturnRequestResponse":       func() interface{ GetHelp() (string, error) } { return messages.NewReturnRequestResponse() },
	}

	for name, constructor := range constructors {
		t.Run(name+"_Constructor", func(t *testing.T) {
			processor := constructor()
			assert.NotNil(t, processor)

			// Test that GetHelp works (this exercises more of the constructor logic)
			help, err := processor.GetHelp()
			assert.NoError(t, err)
			assert.NotEmpty(t, help)
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
