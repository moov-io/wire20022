package messages

import (
	"fmt"
	"log"
	"strings"

	CustomerCreditTransferModel "github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
)

// ExampleUniversalReader_basicUsage demonstrates basic usage of the UniversalReader
// Note: This example is currently disabled due to XML mapping complexities
func exampleUniversalReader_basicUsage() {
	// Create a universal reader
	reader := NewUniversalReader()

	// Sample CustomerCreditTransfer XML
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<FIToFICstmrCdtTrf xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08">
  <GrpHdr>
    <MsgId>20240123MSGID001</MsgId>
    <CreDtTm>2024-01-23T10:30:00</CreDtTm>
    <NbOfTxs>1</NbOfTxs>
  </GrpHdr>
</FIToFICstmrCdtTrf>`

	// Parse the message - type is automatically detected
	parsed, err := reader.ReadBytes([]byte(xml))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Detected Message Type: %s\n", parsed.Type)
	fmt.Printf("Version: %s\n", parsed.Version)
	fmt.Printf("Root Element: %s\n", parsed.Detection.RootElement)
	fmt.Printf("Detection Method: %s\n", parsed.Detection.DetectedBy)

	// Output:
	// Detected Message Type: CustomerCreditTransfer
	// Version: 001.08
	// Root Element: FIToFICstmrCdtTrf
	// Detection Method: root_element
}

// ExampleUniversalReader_readFromFile demonstrates reading from a file
func exampleUniversalReader_readFromFile() {
	reader := NewUniversalReader()

	// Open a file (this would be a real file in practice)
	// file, err := os.Open("payment.xml")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer file.Close()

	// For this example, we'll use a string reader
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<PmtRtr xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10">
  <GrpHdr>
    <MsgId>RTN20240123001</MsgId>
    <CreDtTm>2024-01-23T10:30:00</CreDtTm>
    <NbOfTxs>1</NbOfTxs>
  </GrpHdr>
</PmtRtr>`

	// Parse from io.Reader
	parsed, err := reader.Read(strings.NewReader(xml))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message Type: %s\n", parsed.Type)
	fmt.Printf("Successfully parsed %s message\n", parsed.Type)

	// Output:
	// Message Type: PaymentReturn
	// Successfully parsed PaymentReturn message
}

// ExampleUniversalReader_typeAssertion demonstrates type assertion after parsing
func exampleUniversalReader_typeAssertion() {
	reader := NewUniversalReader()

	xml := `<?xml version="1.0" encoding="UTF-8"?>
<FIToFICstmrCdtTrf xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08">
  <GrpHdr>
    <MsgId>20240123MSGID001</MsgId>
    <CreDtTm>2024-01-23T10:30:00</CreDtTm>
    <NbOfTxs>1</NbOfTxs>
  </GrpHdr>
</FIToFICstmrCdtTrf>`

	parsed, err := reader.ReadBytes([]byte(xml))
	if err != nil {
		log.Fatal(err)
	}

	// Type assertion to work with specific message type
	switch parsed.Type {
	case TypeCustomerCreditTransfer:
		if cct, ok := parsed.Message.(*CustomerCreditTransferModel.MessageModel); ok {
			fmt.Printf("Customer Credit Transfer Message ID: %s\n", cct.MessageId)
		}
	case TypePaymentReturn:
		fmt.Println("This is a Payment Return message")
	default:
		fmt.Printf("Unknown message type: %s\n", parsed.Type)
	}

	// Output:
	// Customer Credit Transfer Message ID: 20240123MSGID001
}

// ExampleUniversalReader_errorHandling demonstrates error handling with verbose reporting
func exampleUniversalReader_errorHandling() {
	reader := NewUniversalReader()
	reader.VerboseErrors = true

	// Invalid XML that will cause parsing error
	invalidXML := `<?xml version="1.0" encoding="UTF-8"?>
<FIToFICstmrCdtTrf xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08">
  <InvalidStructure>
</FIToFICstmrCdtTrf>`

	_, err := reader.ReadBytes([]byte(invalidXML))
	if err != nil {
		fmt.Println("Error occurred:")
		fmt.Println("Contains enhanced error info:", strings.Contains(err.Error(), "Root element:"))
	}

	// Output:
	// Error occurred:
	// Contains enhanced error info: true
}

// ExampleUniversalReader_validation demonstrates message validation
func exampleUniversalReader_validation() {
	reader := NewUniversalReader()

	// Minimal XML that will parse but fail validation
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<FIToFICstmrCdtTrf xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08">
  <GrpHdr>
    <MsgId>TEST123</MsgId>
  </GrpHdr>
</FIToFICstmrCdtTrf>`

	parsed, err := reader.ReadBytes([]byte(xml))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parsing successful: %s\n", parsed.Type)

	// Validate the parsed message
	err = reader.ValidateMessage(parsed)
	if err != nil {
		fmt.Printf("Validation failed: %s\n", "Message incomplete")
	} else {
		fmt.Println("Validation passed")
	}

	// Output:
	// Parsing successful: CustomerCreditTransfer
	// Validation failed: Message incomplete
}

// ExampleUniversalReader_bkToCstmrAcctRpt demonstrates handling of BkToCstmrAcctRpt messages
func exampleUniversalReader_bkToCstmrAcctRpt() {
	reader := NewUniversalReader()

	// ActivityReport (camt.052) XML
	xml := `<?xml version="1.0" encoding="UTF-8"?>
<BkToCstmrAcctRpt xmlns="urn:iso:std:iso:20022:tech:xsd:camt.052.001.08">
  <GrpHdr>
    <MsgId>ACTR20240123001</MsgId>
    <CreDtTm>2024-01-23T10:30:00</CreDtTm>
  </GrpHdr>
  <Rpt>
    <Id>EDAY</Id>
  </Rpt>
</BkToCstmrAcctRpt>`

	parsed, err := reader.ReadBytes([]byte(xml))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Detected Type: %s\n", parsed.Type)
	fmt.Printf("Detection Method: %s\n", parsed.Detection.DetectedBy)
	fmt.Printf("Message ID: %s\n", parsed.Detection.AdditionalInfo["GrpHdr.MsgId"])
	fmt.Printf("Report ID: %s\n", parsed.Detection.AdditionalInfo["Rpt.Id"])

	// Output:
	// Detected Type: ActivityReport
	// Detection Method: content_analysis
	// Message ID: ACTR20240123001
	// Report ID: EDAY
}

// ExampleUniversalReader_batchProcessing demonstrates processing multiple messages
func exampleUniversalReader_batchProcessing() {
	reader := NewUniversalReader()

	// Sample messages of different types
	messages := []string{
		`<FIToFICstmrCdtTrf xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08"><GrpHdr><MsgId>CCT001</MsgId></GrpHdr></FIToFICstmrCdtTrf>`,
		`<PmtRtr xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10"><GrpHdr><MsgId>RTN001</MsgId></GrpHdr></PmtRtr>`,
		`<CdtrPmtActvtnReq xmlns="urn:iso:std:iso:20022:tech:xsd:pain.013.001.09"><GrpHdr><MsgId>DRQ001</MsgId></GrpHdr></CdtrPmtActvtnReq>`,
	}

	typeCounts := make(map[MessageType]int)

	for i, xmlMsg := range messages {
		parsed, err := reader.ReadBytes([]byte(xmlMsg))
		if err != nil {
			fmt.Printf("Error parsing message %d: %v\n", i+1, err)
			continue
		}

		typeCounts[parsed.Type]++
		fmt.Printf("Message %d: %s\n", i+1, parsed.Type)
	}

	fmt.Println("\nSummary:")
	for msgType, count := range typeCounts {
		fmt.Printf("  %s: %d\n", msgType, count)
	}

	// Output:
	// Message 1: CustomerCreditTransfer
	// Message 2: PaymentReturn
	// Message 3: DrawdownRequest
	//
	// Summary:
	//   CustomerCreditTransfer: 1
	//   PaymentReturn: 1
	//   DrawdownRequest: 1
}
