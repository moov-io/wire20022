# Message Type Migration Patterns to Base Abstractions

This document outlines the different migration patterns discovered during the refactoring of existing message types to use base abstractions in the wire20022 library.

## Overview

During the migration process, we identified three distinct patterns based on message complexity and field compatibility. Each pattern offers different benefits and trade-offs.

## Pattern 1: Full Base Processor Migration (Recommended)

**Best for**: Standard ISO 20022 messages with simple field mappings and common header patterns.

**Characteristics**:
- Messages with `MessageId` and `CreatedDateTime` fields
- Simple path mappings without complex nested arrays
- Standard validation requirements

**Examples**: CustomerCreditTransfer, PaymentReturn, PaymentStatusRequest, FedwireFundsAcknowledgement, AccountReportingRequest

**Implementation**:
```go
type MessageModel struct {
    base.MessageHeader `json:",inline"`
    // Message-specific fields...
}

// Use base processor for all operations
func MessageWith(data []byte) (MessageModel, error) {
    return processor.ProcessMessage(data)
}

func DocumentWith(model MessageModel, version VERSION) (models.ISODocument, error) {
    return processor.CreateDocument(model, version)
}

func CheckRequiredFields(model MessageModel) error {
    return processor.ValidateRequiredFields(model)
}
```

**Benefits**:
- ✅ Maximum code reduction (~50-60 lines eliminated per message type)
- ✅ Consistent error handling and validation
- ✅ Automatic factory management and version handling
- ✅ Type-safe processing with generics

**Code Reduction**: ~50-60 lines per message type

## Pattern 2: Structural Abstraction with Custom Processing (Hybrid)

**Best for**: Complex messages with intricate nested array mappings that don't work well with generic processors.

**Characteristics**:
- Messages with complex array structures and custom mapping logic
- Non-standard field types that require override
- Special processing requirements for array elements

**Examples**: ActivityReport

**Implementation**:
```go
type MessageModel struct {
    base.MessageHeader `json:",inline"`
    // Override field types if needed
    MessageId models.CAMTReportType `json:"messageId"`
    
    // Complex message-specific fields
    TotalEntriesPerBankTransactionCode []models.TotalsPerBankTransactionCode `json:"..."`
    EntryDetails                       []models.Entry                        `json:"..."`
}

// Keep custom processing logic for complex operations
func MessageWith(data []byte) (MessageModel, error) {
    // Custom implementation for complex array handling
}

func DocumentWith(model MessageModel, version VERSION) (models.ISODocument, error) {
    // Custom implementation for complex array handling  
}
```

**Benefits**:
- ✅ Structural consistency with embedded common fields
- ✅ Flexibility for complex processing requirements
- ✅ Type safety for specialized field types
- ⚠️ Moderate code reduction (~15-20 lines eliminated)

**Trade-offs**:
- ❌ Cannot use generic processor for full automation
- ❌ Requires maintaining custom processing logic

**Code Reduction**: ~15-20 lines per message type

## Pattern 3: Non-Standard Message Types (Direct Migration)

**Best for**: Messages that don't follow standard ISO 20022 header patterns but have simple processing needs.

**Characteristics**:
- No standard `MessageId`/`CreatedDateTime` fields
- Simple field mappings without complex arrays
- Event-based or system messages

**Examples**: ConnectionCheck

**Implementation**:
```go
type MessageModel struct {
    // Direct field definitions without embedded headers
    EventType  string    `json:"eventType"`
    EventParam string    `json:"eventParam"`
    EventTime  time.Time `json:"eventTime"`
}

// Full base processor usage despite non-standard fields
func MessageWith(data []byte) (MessageModel, error) {
    return processor.ProcessMessage(data)
}
```

**Benefits**:
- ✅ Full base processor benefits despite non-standard structure
- ✅ Clean migration path for system/event messages
- ✅ Demonstrates base abstraction flexibility

**Code Reduction**: ~50 lines per message type

## Migration Decision Tree

```
Message Type Analysis
        |
        v
Does it have MessageId/CreatedDateTime?
        |                    |
       Yes                  No
        |                    |
        v                    v
Are field mappings simple?   Use Pattern 3
        |                    (Direct Migration)
  Yes   |    No
        |     |
        v     v
   Pattern 1  Pattern 2
   (Full)     (Hybrid)
```

## Implementation Guidelines

### Before Migration

1. **Analyze field structure**: Check for common header patterns
2. **Review path mappings**: Identify complex array or nested mappings
3. **Examine test coverage**: Understand validation requirements
4. **Check field types**: Note any specialized types that need preservation

### During Migration

1. **Start with Pattern 1**: Always try the full migration first
2. **Fallback to Pattern 2**: If complex mappings cause issues
3. **Update tests**: Error message formats change with base processor
4. **Validate functionality**: Ensure all existing behavior is preserved

### After Migration

1. **Measure code reduction**: Document lines eliminated
2. **Update documentation**: Note any pattern-specific considerations
3. **Run full test suite**: Ensure no regressions
4. **Commit with clear messaging**: Document the pattern used and why

## Error Handling Changes

When migrating to base abstractions, error message formats change from:
```
// Old format
"field copy MessageId failed: ..."

// New format  
"field copy FIToFIPmtStsReq.GrpHdr.MsgId failed: ..."
```

Test assertions must be updated to match the new XML path format.

## Summary

The base abstractions framework successfully accommodates diverse message types through flexible migration patterns. The approach provides significant code reduction while maintaining full API compatibility and type safety.

**Total Impact** (8 of 13 message types migrated):
- **~400+ lines of duplicate code eliminated**
- **Consistent architecture across message types** 
- **Improved maintainability and type safety**
- **Preserved 100% API compatibility**

Each pattern serves specific use cases while leveraging the base abstraction benefits appropriate to the message complexity.