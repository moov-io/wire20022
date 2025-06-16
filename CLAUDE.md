# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

wire20022 is a Go library for reading, writing, and validating Fedwire ISO20022 messages. It provides a wrapper around the generated structs from XSD schemas to simplify working with ISO20022 message types.

## Architecture

### Package Structure
- `pkg/base/`: **Base abstractions for idiomatic Go message processing**
  - Generic message processor using type parameters
  - Common field patterns (MessageHeader, PaymentCore, AgentPair)
  - Versioned document factory patterns
  - Shared ElementHelper definitions
- `pkg/models/`: Contains implementations for each ISO20022 message type
  - Each message type directory contains: `Message.go`, `MessageHelper.go`, `Message_test.go`, `map.go`, and sample SWIFT messages
  - Supports multiple versions of each message type (e.g., pacs.008.001.02 through pacs.008.001.12)
  - **New message types should use base abstractions to minimize code duplication**
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

### **IMPORTANT: Use Base Abstractions for New Message Types**

**For new message type implementations, always use the base abstractions in `pkg/base/`.** See [BASE_ABSTRACTIONS.md](./BASE_ABSTRACTIONS.md) for complete implementation guide.

#### Quick Start for New Message Types:
1. **Use embedded base types** - `base.PaymentCore`, `base.AgentPair`, etc.
2. **Use generic processor** - Single-line processing functions
3. **Use factory registrations** - Clean version management
4. **Add JSON tags** - Future-ready for JSON workflows

```go
type NewMessageModel struct {
    base.PaymentCore `json:",inline"`        // Embedded common fields
    SpecificField    string `json:"field"`   // Message-specific fields
}

func MessageWith(data []byte) (NewMessageModel, error) {
    return processor.ProcessMessage(data)  // Single line!
}
```

### Legacy Message Type Pattern

Existing message types follow this pattern (new types should use base abstractions):
1. `Message.go` defines the core message structure
2. `MessageHelper.go` provides utility functions for creating and manipulating messages
3. `map.go` contains field mapping logic
4. Tests use sample SWIFT messages from the `swiftSample/` directories

When implementing new features or fixing bugs:
- **For new message types:** Use base abstractions pattern (see [BASE_ABSTRACTIONS.md](./BASE_ABSTRACTIONS.md))
- **For existing message types:** Maintain compatibility with all supported message versions
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

### **MANDATORY XML Mapping Validation Protocol**

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
Always use idiomatic Go programming practices to ensure maintainability:
- **Use base abstractions** - Leverage `pkg/base/` for new message types to eliminate duplication
- **Type parameters over interfaces** - Use generics for type-safe processing
- **Embedded structs over duplication** - Use `base.PaymentCore`, `base.AgentPair` patterns
- **Type assertions over complex interfaces** - Simple interfaces with fallback logic
- Follow Go naming conventions and patterns
- Use proper error handling with wrapped errors
- Prefer simplicity and clarity over cleverness
- Structure packages to minimize dependencies

#### Base Abstractions Guidelines:
- **New message types MUST use base abstractions** to maintain consistency
- **Embedded structs** for common field patterns (MessageHeader, PaymentCore)
- **Generic processors** for type-safe XML processing
- **Factory registrations** for clean version management
- **JSON tags** on all structs for future API compatibility

### API Design
- This library is currently pre-1.0 and breaking changes are acceptable
- Prioritize developer experience and idiomatic Go patterns over backwards compatibility
- Once stable, we will adopt semantic versioning for backwards compatibility
- Design APIs that are intuitive and follow Go conventions