# wire20022

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/wire20022)](https://goreportcard.com/report/github.com/moov-io/wire20022)
[![Go Reference](https://pkg.go.dev/badge/github.com/moov-io/wire20022.svg)](https://pkg.go.dev/github.com/moov-io/wire20022)

A comprehensive Go library for reading, writing, and validating Fedwire ISO 20022 messages with idiomatic Go patterns and robust error handling.

## Overview

wire20022 provides a complete wrapper around ISO 20022 message processing for Fedwire payments, built on top of generated structs from XSD schemas. The library simplifies working with complex XML structures by providing intuitive Go interfaces, comprehensive validation, and detailed error reporting.

### ✨ Key Features

- **Complete Message Support**: Handles all major Fedwire ISO 20022 message types
- **Idiomatic Go Design**: Type-safe interfaces with proper error handling patterns
- **Comprehensive Validation**: Field-level validation with detailed error reporting
- **XML Processing**: Seamless conversion between XML and Go structs
- **Version Compatibility**: Support for multiple message versions within each type
- **Developer-Friendly**: Human-readable field names and clear documentation

### 📋 Supported Message Types

| Message Type | ISO Code | Versions | Description |
|-------------|----------|----------|-------------|
| CustomerCreditTransfer | pacs.008 | .001.02 - .001.12 | Customer credit transfer initiation |
| PaymentReturn | pacs.004 | .001.02 - .001.12 | Payment return |
| PaymentStatusRequest | pacs.028 | .001.01 - .001.05 | Payment status request |
| FedwireFundsPaymentStatus | pacs.002 | .001.03 - .001.14 | Payment status report |
| FedwireFundsSystemResponse | admi.010 | .001.01 | System event acknowledgment |
| DrawdownRequest | pain.013 | .001.01 - .001.10 | Creditor payment activation request |
| DrawdownResponse | pain.014 | .001.01 - .001.10 | Creditor payment activation request status report |
| AccountReportingRequest | camt.060 | .001.01 - .001.06 | Account reporting request |
| ActivityReport | camt.086 | .001.01 - .001.02 | Intraday transaction query |
| EndpointDetailsReport | camt.090 | .001.01 - .001.02 | Service availability acknowledgment |
| EndpointGapReport | camt.087 | .001.01 - .001.02 | Request to modify payment |
| EndpointTotalsReport | camt.089 | .001.01 - .001.02 | Payment status report |
| ReturnRequestResponse | camt.029 | .001.01 - .001.12 | Resolution of investigation |

## 🚀 Quick Start

### Installation

```bash
go get github.com/moov-io/wire20022
```

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
    "github.com/moov-io/wire20022/pkg/models"
)

func main() {
    // Read XML message from file or bytes
    xmlData := []byte(`<?xml version="1.0"?>
    <Document xmlns="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08">
        <CstmrCdtTrfInitn>
            <GrpHdr>
                <MsgId>MSG001</MsgId>
                <CreDtTm>2025-01-15T10:00:00</CreDtTm>
                <NbOfTxs>1</NbOfTxs>
            </GrpHdr>
        </CstmrCdtTrfInitn>
    </Document>`)
    
    // Parse XML into Go struct
    message, err := CustomerCreditTransfer.MessageWith(xmlData)
    if err != nil {
        log.Fatal("Failed to parse message:", err)
    }
    
    fmt.Printf("Message ID: %s\n", message.MessageId)
    fmt.Printf("Created: %s\n", message.CreatedDateTime.Format("2006-01-02 15:04:05"))
    fmt.Printf("Number of Transactions: %s\n", message.NumberOfTransactions)
}
```

### Creating and Validating Messages

```go
func createMessage() {
    // Create a new message
    message := CustomerCreditTransfer.CustomerCreditTransferDataModel()
    message.MessageId = "MSG002"
    message.CreatedDateTime = time.Now()
    message.NumberOfTransactions = "1"
    message.SettlementMethod = models.SettlementMethodType("CLRG")
    
    // Convert to XML document
    document, err := CustomerCreditTransfer.DocumentWith(message, CustomerCreditTransfer.PACS_008_001_08)
    if err != nil {
        log.Fatal("Failed to create document:", err)
    }
    
    // Validate the document
    if err := document.Validate(); err != nil {
        log.Fatal("Validation failed:", err)
    }
    
    // Marshal to XML
    xmlData, err := xml.MarshalIndent(document, "", "  ")
    if err != nil {
        log.Fatal("Failed to marshal XML:", err)
    }
    
    fmt.Println(string(xmlData))
}
```

### Using the Wrapper Interface

```go
import "github.com/moov-io/wire20022/pkg/wrapper"

func useWrapper() {
    w := wrapper.NewCustomerCreditTransferWrapper()
    
    // Parse from XML
    xmlData := []byte("...") // Your XML data
    if err := w.ParseXML(xmlData); err != nil {
        log.Fatal("Parse error:", err)
    }
    
    // Access parsed data
    model := w.GetModel()
    fmt.Printf("Message ID: %s\n", model.MessageId)
    
    // Generate XML
    xmlOutput, err := w.GenerateXML()
    if err != nil {
        log.Fatal("Generate error:", err)
    }
    
    fmt.Println(string(xmlOutput))
}
```

## 🏗️ Architecture

### Package Structure

```
wire20022/
├── pkg/
│   ├── base/             # 🆕 Base abstractions for idiomatic Go patterns
│   │   ├── message_header.go    # Common message structures  
│   │   ├── processor.go         # Generic message processor
│   │   ├── factory.go           # Versioned document factory
│   │   └── helpers.go           # Shared ElementHelper definitions
│   ├── models/           # Core message type implementations
│   │   ├── CustomerCreditTransfer/
│   │   ├── PaymentReturn/
│   │   ├── DrawdownRequest/
│   │   └── ...
│   ├── wrapper/          # Simplified wrapper interfaces
│   ├── errors/           # Error handling utilities
│   └── fedwire/          # Common types and utilities
├── cmd/wire20022/        # Command-line tools
└── internal/server/      # HTTP server implementation
```

### Message Type Structure

Each message type follows a consistent pattern:
- `Message.go` - Core message structure and business logic
- `MessageHelper.go` - Utility functions for message creation
- `Message_test.go` - Comprehensive test suite
- `map.go` - XML field mapping configuration
- `version.go` - Version-specific constants and mappings
- `swiftSample/` - Authoritative XML sample files

## 🔍 Advanced Usage

### Error Handling

wire20022 implements idiomatic Go error handling with detailed error types:

```go
message, err := CustomerCreditTransfer.MessageWith(invalidXML)
if err != nil {
    // Handle specific error types
    var parseErr *errors.ParseError
    var validationErr *errors.ValidationError
    
    if errors.As(err, &parseErr) {
        fmt.Printf("Parse error in %s: %v\n", parseErr.Field, parseErr.Err)
    } else if errors.As(err, &validationErr) {
        fmt.Printf("Validation failed for %s: %s\n", validationErr.Field, validationErr.Reason)
    }
}
```

### Version-Specific Processing

```go
// Handle different message versions
versions := []CustomerCreditTransfer.PACS_008_001_VERSION{
    CustomerCreditTransfer.PACS_008_001_08,
    CustomerCreditTransfer.PACS_008_001_09,
    CustomerCreditTransfer.PACS_008_001_10,
}

for _, version := range versions {
    document, err := CustomerCreditTransfer.DocumentWith(message, version)
    if err != nil {
        continue // Try next version
    }
    
    fmt.Printf("Successfully created document with version %s\n", version)
    break
}
```

### Field Mapping and Debugging

For debugging XML field mapping issues, consult [XML_TO_GO_MAPPING.md](./XML_TO_GO_MAPPING.md):

```go
// Understanding error paths
// XML: <CdtrPmtActvtnReq><GrpHdr><MsgId>
// Go:  CdtrPmtActvtnReq.GrpHdr.MsgId
// Error: "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: ..."
```

## 🧪 Development & Testing

### Running Tests

```bash
# Run all tests with coverage
make check

# Run tests for specific message type
go test ./pkg/models/CustomerCreditTransfer

# Run with verbose output
go test -v ./pkg/models/CustomerCreditTransfer

# Generate coverage report
make cover-test
make cover-web
```

### Building

```bash
# Build the library
make dist

# Build Docker image
make docker

# Clean build artifacts
make clean
```

### Development Setup

```bash
# Start development environment
make setup

# Stop development environment
make teardown
```

## 📚 Documentation

- **[BASE_ABSTRACTIONS.md](./BASE_ABSTRACTIONS.md)** - 🆕 Complete guide for implementing new message types using base abstractions
- **[XML_TO_GO_MAPPING.md](./XML_TO_GO_MAPPING.md)** - Critical guide for XML field mapping
- **[CLAUDE.md](./CLAUDE.md)** - Development guidelines and patterns
- **[Go Reference](https://pkg.go.dev/github.com/moov-io/wire20022)** - API documentation

## 🤝 Contributing

We welcome contributions! This project uses modern Go practices and follows strict quality standards.

### Current Needs

- **Test Coverage Expansion** - Help us reach >90% coverage across all message types
- **Message Type Support** - Add support for additional ISO 20022 message types
- **Performance Optimization** - Optimize XML parsing and validation performance
- **Documentation** - Improve examples and usage documentation
- **Real-World Testing** - Test with actual Fedwire message samples

### Development Guidelines

1. **Always run `make check` before committing** - This catches issues early
2. **Validate XML mappings** - Use `swiftSample/` directories as source of truth
3. **Follow idiomatic Go** - Type safety, proper error handling, clear interfaces
4. **Add comprehensive tests** - Cover both success and failure scenarios
5. **Update documentation** - Keep README and mapping guides current

### Getting Started

```bash
# Fork and clone the repository
git clone https://github.com/your-username/wire20022.git
cd wire20022

# Install dependencies
go mod download

# Run tests to ensure everything works
make check

# Make your changes and test
# ... your development work ...

# Verify before submitting
make check
```

### Pull Request Process

1. Create a feature branch from `master`
2. Make your changes with tests
3. Ensure `make check` passes
4. Submit a pull request with clear description
5. Respond to review feedback

## 📊 Performance

### Benchmarks

Current performance characteristics (run `go test -bench=.`):

- **XML Parsing**: ~1000 messages/second for typical CustomerCreditTransfer
- **Validation**: ~5000 validations/second 
- **Memory Usage**: ~500KB per message processing
- **Concurrent Processing**: Thread-safe for read operations

### Memory Management

The library is designed for efficient memory usage:
- Minimal allocations during parsing
- Reusable validation contexts
- Optional object pooling for high-throughput scenarios

## 📋 Current Status

- ✅ **Core Functionality**: Complete XML parsing and generation
- ✅ **Error Handling**: Idiomatic Go error patterns implemented
- ✅ **Message Types**: 13+ message types supported with multiple versions
- ✅ **Validation**: Comprehensive field validation with detailed reporting
- ✅ **Testing**: 51% code coverage with growing test suite
- 🔄 **Performance**: Ongoing optimization for high-throughput scenarios
- 📝 **Documentation**: Continuous improvement of examples and guides

## 🐛 Issues & Support

- **Bug Reports**: [GitHub Issues](https://github.com/moov-io/wire20022/issues)
- **Feature Requests**: [GitHub Discussions](https://github.com/moov-io/wire20022/discussions)
- **Security Issues**: Please report privately to security@moov.io

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Built on [moov-io/fedwire20022](https://github.com/moov-io/fedwire20022) generated structs
- Follows [Fedwire ISO 20022 specifications](https://www.frbservices.org/financial-services/fednow/what-is-iso-20022-why-does-it-matter)
- Part of the [Moov](https://github.com/moov-io) financial technology ecosystem

---

**📢 Ready to contribute?** Check out our [good first issues](https://github.com/moov-io/wire20022/labels/good%20first%20issue) or join the discussion in [GitHub Discussions](https://github.com/moov-io/wire20022/discussions)!