# Type-Safe Migration Framework for wire20022

## Overview

This framework provides a systematic approach to migrate all 16 message types from reflection-based validation to compile-time type-safe validation, based on the proven EndpointDetailsReport implementation.

## Migration Priority Matrix

### **Tier 1: High-Impact Messages (Immediate)**
1. **CustomerCreditTransfer** - 11 versions, complex evolution, high usage
2. **FedwireFundsPaymentStatus** - 12 versions, significant V10+ enhancements
3. **PaymentReturn** - 12 versions, payment processing core

### **Tier 2: Moderate Complexity (Next)**
4. **DrawdownRequest** - 10 versions, moderate field evolution
5. **DrawdownResponse** - 10 versions, paired with DrawdownRequest
6. **ActivityReport** - 12 versions, reporting family
7. **Master** - 11 versions, core message type

### **Tier 3: Simpler/Stable (Final)**
8. **PaymentStatusRequest** - 11 versions, stable evolution
9. **AccountReportingRequest** - 6 versions, manageable scope
10. **EndpointGapReport** - 11 versions, similar to EndpointDetailsReport
11. **EndpointTotalsReport** - 11 versions, similar to EndpointDetailsReport
12. **ReturnRequestResponse** - 9 versions, moderate complexity
13. **ConnectionCheck** - 2 versions, minimal changes
14. **FedwireFundsAcknowledgement** - 1 version, stable
15. **FedwireFundsSystemResponse** - 1 version, event-based pattern

## Migration Patterns by Message Type

### **Pattern A: Payment Messages with Agent Evolution**
**Used by**: CustomerCreditTransfer, PaymentReturn, FedwireFundsPaymentStatus

**Characteristics**:
- Uses `base.PaymentCore` + `base.AgentPair`
- Address field enhancements in V8+ (Floor, RoomNumber, BuildingName)
- UETR field additions in newer versions
- Complex agent relationship evolution

**Migration Strategy**:
```go
// Version-specific field groups
type AddressEnhancementFields struct {
    Floor        string `json:"floor,omitempty"`
    RoomNumber   string `json:"roomNumber,omitempty"`
    BuildingName string `json:"buildingName,omitempty"`
}

type TransactionFields struct {
    UniqueEndToEndTransactionRef string `json:"uniqueEndToEndTransactionRef,omitempty"`
    OriginalUETR                 string `json:"originalUETR,omitempty"`
}

type MessageModel struct {
    base.PaymentCore `json:",inline"`
    base.AgentPair   `json:",inline"`
    
    // Core fields (all versions)
    // ... existing core fields
    
    // V8+ enhancements
    AddressEnhancements *AddressEnhancementFields `json:",inline,omitempty"`
    TransactionRefs     *TransactionFields        `json:",inline,omitempty"`
}
```

### **Pattern B: Reporting Messages with Query Evolution**
**Used by**: ActivityReport, EndpointGapReport, EndpointTotalsReport, Master

**Characteristics**:
- Uses `base.MessageHeader`
- Business query fields in V3+
- Reporting sequence fields in V7+
- Similar evolution to EndpointDetailsReport

**Migration Strategy**: 
Follow the exact EndpointDetailsReport pattern with BusinessQueryFields and ReportingFields.

### **Pattern C: Request/Response Pairs**
**Used by**: DrawdownRequest/DrawdownResponse, PaymentStatusRequest

**Characteristics**:
- Uses `base.MessageHeader`
- Date structure evolution (V6: `Dt` → `Dt.Dt`)
- UETR additions in V7+
- Enhanced address fields in V7+

**Migration Strategy**:
```go
type EnhancedFields struct {
    PaymentUniqueId   string                    `json:"paymentUniqueId,omitempty"`
    EnhancedAddress   *AddressEnhancementFields `json:",inline,omitempty"`
}

type MessageModel struct {
    base.MessageHeader `json:",inline"`
    
    // Core fields (all versions)
    // ... existing core fields
    
    // V7+ enhancements
    Enhanced *EnhancedFields `json:",inline,omitempty"`
}
```

### **Pattern D: Simple/Stable Messages**
**Used by**: ConnectionCheck, FedwireFundsAcknowledgement, FedwireFundsSystemResponse

**Characteristics**:
- Minimal or no version evolution
- Single version or simple V1→V2 changes
- May not need grouped fields

**Migration Strategy**:
Keep flat structure but add type-safe validation methods.

## Automated Migration Framework

### **Script 1: Field Analysis**
```python
# analyze_field_evolution.py
"""
Analyzes map.go files to identify version-specific field patterns
Outputs field grouping recommendations for each message type
"""

def analyze_message_type(message_type):
    # Parse map.go PathMapVX functions
    # Identify field introduction versions
    # Group fields by version ranges
    # Output recommended field groups
```

### **Script 2: Code Generation**
```python
# generate_type_safe_structure.py
"""
Generates type-safe field structures and validation methods
Based on field analysis output
"""

def generate_field_groups(field_analysis):
    # Create version-specific field group structs
    # Generate validation methods for each group
    # Generate NewMessageForVersion constructor
    # Generate ValidateForVersion method
```

### **Script 3: Path Mapping Update**
```python
# update_path_mappings.py
"""
Updates map.go files to point to new grouped field structure
"""

def update_path_mappings(message_type, field_groups):
    # Update XML paths to point to FieldGroup.FieldName
    # Preserve version-specific mapping logic
    # Validate mapping consistency
```

### **Script 4: Test Migration**
```python
# migrate_tests.py
"""
Updates test files to use new field structure
"""

def migrate_test_assertions(test_file, field_groups):
    # Update field access patterns
    # Add version-specific nil checks
    # Fix test helper functions
```

## Implementation Steps

### **Phase 1: Framework Development**
1. Create field analysis scripts
2. Develop code generation templates
3. Build path mapping updater
4. Create test migration tools

### **Phase 2: Tier 1 Migration**
1. **CustomerCreditTransfer** (Pattern A)
   - Analyze 11 versions for field evolution
   - Create AddressEnhancementFields + TransactionFields groups
   - Update XML path mappings
   - Migrate tests
   - Validate complete functionality

2. **FedwireFundsPaymentStatus** (Pattern A)
   - Focus on V10+ UETR and settlement date additions
   - Similar pattern to CustomerCreditTransfer

3. **PaymentReturn** (Pattern A)
   - 12 versions with V9+ UETR additions
   - Return chain field evolution

### **Phase 3: Tier 2 Migration**
Execute remaining moderate complexity messages using established patterns.

### **Phase 4: Tier 3 Migration**
Handle simple/stable messages with minimal field grouping needs.

## Validation Framework

### **Automated Testing**
```bash
# test_migration.sh
#!/bin/bash

# For each migrated message type:
# 1. Run existing version tests
# 2. Run XML round-trip tests
# 3. Validate type-safe validation
# 4. Check path mapping consistency
# 5. Verify no regression in coverage

for message_type in "${MIGRATED_TYPES[@]}"; do
    echo "Testing $message_type migration..."
    
    # Existing functionality
    go test -v "./pkg/models/$message_type"
    
    # Type-safe validation
    test_type_safe_validation "$message_type"
    
    # XML processing
    test_xml_round_trip "$message_type"
done
```

### **Type Safety Verification**
```go
// verify_type_safety.go
// Compile-time verification that all field access is type-safe
// No reflection usage detected
// All validation uses direct field access

func verifyNoReflection(messageType string) error {
    // Static analysis to ensure no reflect package usage
    // Verify all field access is compile-time checked
}
```

## Documentation Standards

### **Field Evolution Documentation**
Each migrated message type should include:

```go
// CustomerCreditTransfer field evolution:
// V2-V7: Core payment fields
// V8+: Enhanced address fields (Floor, RoomNumber, BuildingName)  
// V9+: UniqueEndToEndTransactionRef support
// V12: Date structure enhancement (RltdDt.Dt)

type MessageModel struct {
    base.PaymentCore `json:",inline"`
    
    // Core fields present in all versions (V2+)
    // ... core fields
    
    // V8+ address enhancements
    AddressEnhancements *AddressEnhancementFields `json:",inline,omitempty"`
    
    // V9+ transaction references  
    TransactionRefs *TransactionFields `json:",inline,omitempty"`
}
```

### **Migration Verification Checklist**

For each message type migration:

- [ ] Field groups identified and documented
- [ ] NewMessageForVersion() constructor implemented
- [ ] ValidateForVersion() method with type-safe validation
- [ ] XML path mappings updated for grouped fields
- [ ] All existing tests pass
- [ ] Version-specific test assertions added
- [ ] No reflection usage in validation code
- [ ] Type-safe field access throughout
- [ ] IDE autocomplete verified
- [ ] Refactoring safety confirmed

## Benefits Tracking

### **Metrics to Track**
- **Lines of reflection code removed**
- **Compile-time errors caught vs runtime**
- **Performance improvement** (validation speed)
- **Developer experience** (IDE support quality)
- **Maintainability** (refactoring safety)

### **Success Criteria**
- All 16 message types use type-safe validation
- Zero reflection usage in validation logic
- All tests pass with improved type safety
- Documentation covers field evolution patterns
- Framework enables easy future message type additions

This framework ensures systematic, reliable migration of all message types to the proven type-safe validation approach while maintaining full functionality and improving developer experience.