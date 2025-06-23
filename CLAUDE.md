# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

wire20022 is a Go library for reading, writing, and validating Fedwire ISO20022 messages. **The primary purpose is to read and write Fedwire XML files**, with JSON support as a secondary use case. The library provides idiomatic Go abstractions around generated XSD structs for simplified ISO20022 message processing.

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
  - **All message types use base abstractions and idiomatic XML-first API**
  - **XML-first design**: Primary API uses `ReadXML()`, `WriteXML()`, and `ParseXML()` methods
- `pkg/messages/`: **Type-safe message processors (v1.0 API)**
  - Unified generic interface for all ISO 20022 message types
  - 68% code reduction through generic architecture
  - Enhanced error handling with message type context
  - Compile-time type safety for models and versions
  - Idiomatic Go naming conventions (NewCustomerCreditTransfer, Validate, etc.)
- `pkg/errors/`: Domain-specific error types following Go standard library conventions
- `internal/server/`: Internal HTTP server implementation (planned)
- `cmd/wire20022/`: Command-line application (planned)

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

## Claude Development Files Organization

### `.claude/` Directory Structure

All Claude-generated files are organized in the `.claude/` directory to keep the project root clean:

```
.claude/
├── docs/           # Documentation and architectural guides
├── scripts/        # Automation scripts and build tools
├── coverage/       # Test coverage reports (gitignored)
└── archive/        # Historical/deprecated files
```

#### File Placement Guidelines

**`.claude/docs/`** - Documentation files:
- Architecture guides (BASE_ABSTRACTIONS.md, IMPLEMENTATION_GUIDE.md)
- Design decisions (ERROR_DESIGN_PROPOSAL.md, TYPE_SAFETY_ANALYSIS.md)
- Migration documentation (MIGRATION_STATUS.md, WRAPPER_MIGRATION_PLAN.md)
- Field mapping guides (XML_TO_GO_MAPPING.md)
- Test strategies and coverage analysis

**`.claude/scripts/`** - Automation scripts:
- Python scripts for code generation (`apply_*.py`, `fix_*.py`)
- Shell scripts for batch operations (`*.sh`)
- Refactoring and migration tools
- Build and deployment automation

**`.claude/coverage/`** - Coverage reports (gitignored):
- `cover.out`, `coverage.out`, `coverage.txt`
- Test coverage analysis files
- Performance benchmarks

**`.claude/archive/`** - Deprecated files:
- Old test outputs
- Temporary files kept for reference
- Legacy scripts no longer needed

#### Best Practices for Claude

1. **New Scripts**: Always place automation scripts in `.claude/scripts/`
2. **Documentation**: Create architectural docs in `.claude/docs/`
3. **Temporary Files**: Use `.claude/archive/` for files that may be removed later
4. **Coverage**: Let coverage reports go to `.claude/coverage/` (gitignored)
5. **Large Test Files**: Use modularization script for files >30K bytes

### Test File Modularization

**Problem**: Some `*_version_test.go` files are too large for Claude's context window (>30K bytes).

**Solution**: Split large test files into smaller version groups:
```bash
# Use the modularization script
python3 .claude/scripts/modularize_test_files.py --dry-run  # Preview changes
python3 .claude/scripts/modularize_test_files.py           # Apply changes
```

**Large Files Requiring Modularization**:
- `ActivityReport/Message_version_test.go` (96K, 12 versions)
- `CustomerCreditTransfer/Message_version_test.go` (68K, 11 versions)
- `EndpointDetailsReport/Message_version_test.go` (64K, 11 versions)
- `Master/Message_version_test.go` (60K, 11 versions)
- `PaymentReturn/Message_version_test.go` (55K, 12 versions)

**Result**: Each large file becomes 3-4 smaller files (~25K each), making them manageable for Claude.

**Documentation**: See `.claude/docs/TEST_FILE_MODULARIZATION.md` for detailed strategy.

## Import Path Management

### Local Development vs Upstream Commits

**CRITICAL**: This repository requires different import path strategies for local development versus upstream commits:

#### Local Development (Fork)
When working on a local fork (`github.com/wadearnold/wire20022`), use **relative import paths** in go.mod:

```go
// go.mod for local development
module github.com/wadearnold/wire20022

// Use relative imports for local packages
import (
    "./pkg/messages"
    "./pkg/models/CustomerCreditTransfer"
)
```

#### Upstream Commits (Moov-io)
When preparing commits for the upstream repository (`github.com/moov-io/wire20022`), use **absolute import paths**:

```go
// go.mod for upstream commits
module github.com/moov-io/wire20022

// Use absolute imports for upstream
import (
    "github.com/moov-io/wire20022/pkg/messages"
    "github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
)
```

#### Development Workflow
1. **During development**: Use relative paths for faster local builds and testing
2. **Before commits**: Switch to absolute upstream paths for compatibility
3. **Always verify**: Run `make check` with both import strategies before final commit

#### Why This Matters
- **Local development**: Faster builds, no network dependency, easier testing
- **Upstream commits**: Proper module resolution, CI compatibility, public API consistency
- **Import conflicts**: Different paths prevent accidental cross-contamination

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

### **XML-First API Design**

**This library is designed primarily for reading and writing Fedwire XML files.** All message types implement an idiomatic Go XML-first API using `io.Reader`/`io.Writer` interfaces.

#### Core XML API Methods:
```go
// ReadXML reads XML data from any io.Reader into the MessageModel
func (m *MessageModel) ReadXML(r io.Reader) error

// WriteXML writes the MessageModel as XML to any io.Writer
// If no version is specified, uses the latest version
func (m *MessageModel) WriteXML(w io.Writer, version ...VERSION) error

// ParseXML reads XML data directly from bytes
func ParseXML(data []byte) (*MessageModel, error)
```

#### Usage Examples:
```go
// Reading XML from file
file, err := os.Open("payment.xml")
if err != nil {
    return err
}
defer file.Close()

var msg CustomerCreditTransfer.MessageModel
if err := msg.ReadXML(file); err != nil {
    return err
}

// Writing XML to file with specific version
outFile, err := os.Create("output.xml")
if err != nil {
    return err
}
defer outFile.Close()

if err := msg.WriteXML(outFile, CustomerCreditTransfer.PACS_008_001_10); err != nil {
    return err
}

// Parsing XML from byte data
xmlData := []byte(`<xml>...</xml>`)
msg, err := CustomerCreditTransfer.ParseXML(xmlData)
if err != nil {
    return err
}
```

### **Type-Safe Message Architecture**

**All message types implement type-safe validation with version-specific field grouping for compile-time safety.**

#### Standard Message Type Structure:
1. **Version-specific field groups** - Group fields by when they were introduced
2. **Type-safe validation methods** - Compile-time verified field access
3. **Generic processor** - Single-line processing functions
4. **Factory registrations** - Clean version management
5. **XML-first methods** - `ReadXML()`, `WriteXML()`, `ParseXML()`

```go
// Version-specific fields grouped by introduction version
type TransactionFields struct {
    NewField string `json:"newField"`
}

type MessageModel struct {
    base.PaymentCore `json:",inline"`        // Embedded common fields
    SpecificField    string `json:"field"`   // Core fields (all versions)
    
    // Version-specific field groups (nil when not applicable)
    Transaction *TransactionFields `json:",inline,omitempty"` // V8+ only
}

// Type-safe validation methods
func NewMessageForVersion(version VERSION) MessageModel
func (m MessageModel) ValidateForVersion(version VERSION) error
func (m MessageModel) GetVersionCapabilities() map[string]bool

// XML-first API methods
func (m *MessageModel) ReadXML(r io.Reader) error { /* ... */ }
func (m *MessageModel) WriteXML(w io.Writer, version ...VERSION) error { /* ... */ }
func ParseXML(data []byte) (*MessageModel, error) { /* ... */ }
```

### Adding New Message Types

**For comprehensive step-by-step instructions on implementing new message types, see [IMPLEMENTATION_GUIDE.md](./.claude/docs/IMPLEMENTATION_GUIDE.md).**

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

**Always consult [XML_TO_GO_MAPPING.md](./.claude/docs/XML_TO_GO_MAPPING.md) before:**
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

### XML-First Design with Idiomatic Go
The library prioritizes XML processing as the primary use case while following idiomatic Go patterns:

#### Core Design Principles:
- **XML-first API** - Primary methods use `io.Reader`/`io.Writer` for XML processing
- **Base abstractions** - All message types use `pkg/base/` to eliminate duplication
- **Type parameters over interfaces** - Generics provide type-safe processing
- **Embedded structs** - Common patterns like `base.PaymentCore`, `base.AgentPair`
- **Idiomatic interfaces** - Standard Go patterns with `io.Reader`/`io.Writer`
- **Go naming conventions** - Standard Go patterns throughout
- **Wrapped error handling** - Domain-specific errors with context
- **Simplicity and clarity** - Prefer readable code over clever code
- **Minimal dependencies** - Clean package structure

#### API Design Priorities:
1. **XML processing** - Primary use case for reading/writing Fedwire files
2. **JSON support** - Secondary use case for API integrations
3. **Developer experience** - Intuitive, idiomatic Go APIs
4. **Type safety** - Compile-time guarantees with generics
5. **Performance** - Efficient XML parsing and generation

#### Architecture Guidelines:
- **XML-first methods** - `ReadXML()`, `WriteXML()`, `ParseXML()` on all message types
- **Base abstractions** - Consistent implementation across all message types
- **Embedded structs** - Common field patterns (MessageHeader, PaymentCore)
- **Generic processors** - Type-safe XML processing with version management
- **Factory registrations** - Clean version management and namespace handling
- **JSON tags** - Future-ready for JSON workflows and APIs
- **Error types** - Follow Go standard library conventions

### API Design Philosophy
- **Breaking changes acceptable** - Library is pre-1.0, prioritize correctness
- **Idiomatic Go first** - Follow Go conventions over backwards compatibility
- **XML processing focus** - Optimize for primary Fedwire XML use case
- **Developer experience** - Intuitive APIs that feel natural to Go developers
- **Future semantic versioning** - Once stable, adopt semver for compatibility