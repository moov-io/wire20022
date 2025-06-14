# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

wire20022 is a Go library for reading, writing, and validating Fedwire ISO20022 messages. It provides a wrapper around the generated structs from XSD schemas to simplify working with ISO20022 message types.

## Architecture

### Package Structure
- `pkg/models/`: Contains implementations for each ISO20022 message type
  - Each message type directory contains: `Message.go`, `MessageHelper.go`, `Message_test.go`, `map.go`, and sample SWIFT messages
  - Supports multiple versions of each message type (e.g., pacs.008.001.02 through pacs.008.001.12)
- `pkg/wrapper/`: Simplified wrapper interfaces for each message type
- `internal/server/`: Internal HTTP server implementation
- `cmd/wire20022/`: Command-line application (currently in development)

### Supported Message Types
- CustomerCreditTransfer (pacs.008)
- PaymentReturn (pacs.004)
- PaymentStatusRequest (pacs.028)
- FedwireFundsAcknowledgement (admi.004)
- AccountReportingRequest (camt.060)
- ActivityReport (camt.086)
- ConnectionCheck (admi.001)
- DrawdownRequest (pain.013)
- DrawdownResponse (pain.014)
- EndpointDetailsReport (camt.090)
- EndpointGapReport (camt.087)
- EndpointTotalsReport (camt.089)
- FedwireFundsPaymentStatus (pacs.002)
- FedwireFundsSystemResponse (admi.010)
- ReturnRequestResponse (camt.029)

## Development Commands

### Build and Test
```bash
# Run tests and linting
make check

# Run tests with coverage
make cover-test

# View coverage report in browser
make cover-web

# Build the binary
make dist

# Build Docker image
make docker
```

### **MANDATORY Pre-Commit Verification**

**ALWAYS run `make check` locally before making any commits.** This prevents CI build failures and ensures code quality:

```bash
# REQUIRED before every commit
make check
```

**Why this is critical:**
- Catches test failures, linting issues, and build problems early
- Prevents wasted CI build time and multiple round-trips
- Ensures consistent code quality across all contributions
- Validates that test assertions match actual runtime behavior
- Detects XML-to-Go field mapping inconsistencies

**If `make check` fails:**
1. Fix all reported issues
2. Re-run `make check` to verify fixes
3. Only then proceed with `git commit`

**Never commit code that fails `make check` locally** - it will fail in CI and waste development time.

### Development Setup
```bash
# Start Docker compose services
make setup

# Stop Docker compose services
make teardown

# Clean build artifacts
make clean
```

## Working with Message Types

Each message type follows a consistent pattern:
1. `Message.go` defines the core message structure
2. `MessageHelper.go` provides utility functions for creating and manipulating messages
3. `map.go` contains field mapping logic
4. Tests use sample SWIFT messages from the `swiftSample/` directories

When implementing new features or fixing bugs:
- Ensure compatibility with all supported message versions
- Add tests using the existing pattern with sample SWIFT messages
- Follow the established structure for new message types

### XML to Go Struct Field Mapping

**CRITICAL**: This library bridges ISO 20022 XML messages with Go structs, where XML element names often differ from Go struct field names. This affects error messages, field paths, and debugging.

**Always consult [XML_TO_GO_MAPPING.md](./XML_TO_GO_MAPPING.md) before:**
- Writing test assertions for validation errors
- Debugging field mapping issues
- Adding new message types
- Interpreting error messages

Example of the mapping challenge:
- XML: `<CdtrPmtActvtnReq>` (Creditor Payment Activation Request)
- Go Struct: `CstmrDrctDbtInitn` (Customer Direct Debit Initiation)
- Error Path: `CstmrDrctDbtInitn.GrpHdr.MsgId`

**Key Rule**: Always use Go struct field paths in test assertions and path mappings, not XML element names.

## Development Philosophy

### Idiomatic Go
Always use idiomatic Go programming practices to ensure maintainability:
- Follow Go naming conventions and patterns
- Use proper error handling with wrapped errors
- Prefer simplicity and clarity over cleverness
- Use type assertions over interfaces unless interfaces are the only appropriate solution
- Structure packages to minimize dependencies

### API Design
- This library is currently pre-1.0 and breaking changes are acceptable
- Prioritize developer experience and idiomatic Go patterns over backwards compatibility
- Once stable, we will adopt semantic versioning for backwards compatibility
- Design APIs that are intuitive and follow Go conventions