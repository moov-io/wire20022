# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

wire20022 is a Go library for reading, writing, and validating Fedwire ISO20022 messages. It provides a wrapper around the generated structs from XSD schemas to simplify working with ISO20022 message types.

## Architecture

### Package Structure
- `pkg/base/`: **Core abstractions for idiomatic Go message processing**
  - Generic message processor using type parameters
  - Common field patterns (MessageHeader, PaymentCore, AgentPair)
  - Versioned document factory patterns
  - Shared ElementHelper definitions
- `pkg/models/`: Contains implementations for each ISO20022 message type
  - Each message type directory contains: `Message.go`, `MessageHelper.go`, `Message_test.go`, `map.go`, and sample SWIFT messages
  - Supports multiple versions of each message type (e.g., pacs.008.001.02 through pacs.008.001.12)
  - **All message types use base abstractions for consistent implementation**
- `pkg/wrapper/`: Simplified wrapper interfaces for each message type
- `pkg/errors/`: Domain-specific error types following Go standard library conventions
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

### **Message Type Architecture**

**All message types use the base abstractions in `pkg/base/` for consistent, type-safe implementation.** This is the standard architecture of the library.

#### Standard Message Type Structure:
1. **Use embedded base types** - `base.PaymentCore`, `base.AgentPair`, etc.
2. **Use generic processor** - Single-line processing functions
3. **Use factory registrations** - Clean version management
4. **Include JSON tags** - API-ready for JSON workflows

```go
type MessageModel struct {
    base.PaymentCore `json:",inline"`        // Embedded common fields
    SpecificField    string `json:"field"`   // Message-specific fields
}

func MessageWith(data []byte) (MessageModel, error) {
    return processor.ProcessMessage(data)  // Single line!
}
```

### Adding New Message Types

**For comprehensive step-by-step instructions on implementing new message types, see [IMPLEMENTATION_GUIDE.md](./IMPLEMENTATION_GUIDE.md).**

This guide covers everything from XSD schemas to complete implementation using base abstractions.

#### Standard File Structure:
1. `Message.go` - Core message structure using base abstractions
2. `MessageHelper.go` - Helper functions for creating and manipulating messages
3. `map.go` - Field mapping logic for XML to Go struct conversion
4. `Message_test.go` - Tests using sample SWIFT messages
5. `swiftSample/` - Authoritative XML samples for validation

When implementing new features or fixing bugs:
- **All message types:** Use base abstractions pattern consistently
- **Maintain compatibility** with all supported message versions
- **Add tests** using sample SWIFT messages from `swiftSample/` directories
- **Follow established patterns** for consistency across the library

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

### **XML Mapping Validation Protocol**

**Before making ANY changes to `map.go` files or test assertions**, you MUST validate against actual XML samples:

#### 1. **Source of Truth**: `swiftSample/` Directories
Each message type contains authoritative XML samples in its `swiftSample/` directory:
- `pkg/models/DrawdownRequest/swiftSample/`
- `pkg/models/FedwireFundsPaymentStatus/swiftSample/`
- `pkg/models/FedwireFundsSystemResponse/swiftSample/`
- `pkg/models/DrawdownResponse/swiftSample/`

#### 2. **Validation Steps** (REQUIRED for any mapping changes):
```bash
# Step 1: Read actual XML structure from Swift samples
cat pkg/models/[MessageType]/swiftSample/[sample-file]

# Step 2: Verify XML path matches map.go mapping
# Look for: Root Element -> Group Header -> Message ID
# Example: <CdtrPmtActvtnReq><GrpHdr><MsgId>

# Step 3: Confirm test assertions use EXACT XML paths
# Test assertions MUST match the actual XML structure, not assumptions

# Step 4: Run make check to verify all mappings work
make check
```

#### 3. **Verified XML Path Structure** (DO NOT CHANGE without validation):
| Message Type | XML Root | Message ID Path | Test Assertion Should Expect |
|-------------|----------|-----------------|------------------------------|
| DrawdownRequest | `<CdtrPmtActvtnReq>` | `CdtrPmtActvtnReq.GrpHdr.MsgId` | `CdtrPmtActvtnReq.GrpHdr.MsgId failed` |
| FedwireFundsPaymentStatus | `<FIToFIPmtStsRpt>` | `FIToFIPmtStsRpt.GrpHdr.MsgId` | `FIToFIPmtStsRpt.GrpHdr.MsgId failed` |
| FedwireFundsSystemResponse | `<SysEvtAck>` | `SysEvtAck.MsgId` | `SysEvtAck.MsgId failed` |
| DrawdownResponse | `<CdtrPmtActvtnReqStsRpt>` | `CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId` | `CdtrPmtActvtnReqStsRpt.GrpHdr.MsgId failed` |

#### 4. **Forbidden Actions** (Will break the library):
- ❌ **Never guess XML paths** - Always verify against actual samples
- ❌ **Never change map.go without checking Swift samples first**
- ❌ **Never update test assertions without verifying runtime behavior**
- ❌ **Never assume XML structure based on Go struct names**

#### 5. **Required Verification After Changes**:
```bash
# MUST pass before any commit involving XML mappings
make check

# Additional verification for mapping changes
go test -v ./pkg/models/[MessageType] -run TestVersion
```

**Breaking this protocol will cause CI failures and incorrect XML processing.**

## Git Workflow and Commit Strategy

### **MANDATORY Summary Commits**

**ALWAYS create a summary commit after successfully completing each task.** This ensures proper version control and makes it easy to track progress:

```bash
# After completing any significant task or set of changes
git add .
git commit -m "Complete [task description]

- Key change 1
- Key change 2  
- Key change 3

🤖 Generated with [Claude Code](https://claude.ai/code)

Co-Authored-By: Claude <noreply@anthropic.com>"
```

**When to create summary commits:**
- After migrating a message type to base abstractions
- After fixing test assertions for multiple message types
- After completing any multi-step refactoring task
- After resolving build or validation issues
- Before switching to work on a different area of the codebase

**Benefits:**
- Provides clear checkpoint for rollback if needed
- Documents progress for team visibility
- Ensures changes are preserved before continuing
- Makes it easier to track what was accomplished

## Development Philosophy

### Idiomatic Go with Base Abstractions
The library follows idiomatic Go programming practices to ensure maintainability:
- **Base abstractions** - All message types use `pkg/base/` to eliminate duplication
- **Type parameters over interfaces** - Generics provide type-safe processing
- **Embedded structs** - Common patterns like `base.PaymentCore`, `base.AgentPair`
- **Type assertions over complex interfaces** - Simple interfaces with fallback logic
- **Go naming conventions** - Standard Go patterns throughout
- **Wrapped error handling** - Domain-specific errors with context
- **Simplicity and clarity** - Prefer readable code over clever code
- **Minimal dependencies** - Clean package structure

#### Architecture Guidelines:
- **All message types** use base abstractions for consistency
- **Embedded structs** for common field patterns (MessageHeader, PaymentCore)
- **Generic processors** for type-safe XML processing
- **Factory registrations** for clean version management
- **JSON tags** on all structs for API compatibility
- **Error types** follow Go standard library conventions

### API Design
- This library is currently pre-1.0 and breaking changes are acceptable
- Prioritize developer experience and idiomatic Go patterns over backwards compatibility
- Once stable, we will adopt semantic versioning for backwards compatibility
- Design APIs that are intuitive and follow Go conventions