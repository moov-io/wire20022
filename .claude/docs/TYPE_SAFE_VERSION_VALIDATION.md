# Type-Safe Version Validation Design

## Problem Statement

The current reflection-based validation system introduces type safety regressions by using:
- String-based field references
- Runtime reflection for field validation
- No compile-time verification of version-field mappings

## Type-Safe Solution Design

### 1. **Generic Version-Aware Validation Interface**

```go
// Type-safe version validation interface
type VersionValidator[M any, V comparable] interface {
    ValidateForVersion(model M, version V) error
    GetRequiredFieldsForVersion(version V) []FieldRef[M]
}

// Compile-time field reference
type FieldRef[M any] func(M) any

// Type-safe field reference helpers
func MessageIdField[M any](model M) any {
    // Use type assertion or interface method to access MessageId
    if m, ok := any(model).(interface{ GetMessageId() string }); ok {
        return m.GetMessageId()
    }
    return nil
}
```

### 2. **Interface-Based Field Access**

```go
// Message types implement this interface for type-safe field access
type MessageWithRequiredFields interface {
    GetMessageId() string
    GetCreatedDateTime() time.Time
    ValidateRequiredFields() error
}

// Extend for version-specific fields
type MessageWithVersionFields interface {
    MessageWithRequiredFields
    GetVersionSpecificFields(version any) map[string]any
}
```

### 3. **Type-Safe Version Mapping**

```go
// Compile-time version-field mapping
type VersionFieldMap[M any, V comparable] struct {
    versionFields map[V][]FieldValidator[M]
}

type FieldValidator[M any] struct {
    Name      string
    Validator func(M) error
    Required  bool
}

// Usage with compile-time safety
func NewEndpointDetailsReportValidator() *VersionFieldMap[MessageModel, CAMT_052_001_VERSION] {
    return &VersionFieldMap[MessageModel, CAMT_052_001_VERSION]{
        versionFields: map[CAMT_052_001_VERSION][]FieldValidator[MessageModel]{
            CAMT_052_001_02: {
                {Name: "MessageId", Validator: validateMessageId, Required: true},
                {Name: "CreatedDateTime", Validator: validateCreatedDateTime, Required: true},
                // No business query fields for V2
            },
            CAMT_052_001_07: {
                {Name: "MessageId", Validator: validateMessageId, Required: true},
                {Name: "CreatedDateTime", Validator: validateCreatedDateTime, Required: true},
                {Name: "ReportingSequence", Validator: validateReportingSequence, Required: true},
            },
        },
    }
}
```

### 4. **Compile-Time Field Validators**

```go
// Type-safe field validators
func validateMessageId(model MessageModel) error {
    if model.MessageId == "" {
        return fmt.Errorf("MessageId is required")
    }
    return nil
}

func validateCreatedDateTime(model MessageModel) error {
    if model.CreatedDateTime.IsZero() {
        return fmt.Errorf("CreatedDateTime is required")
    }
    return nil
}

func validateReportingSequence(model MessageModel) error {
    // Type-safe access to ReportingSequence field
    if model.ReportingSequence.FromSeq == "" {
        return fmt.Errorf("ReportingSequence.FromSeq is required")
    }
    return nil
}
```

### 5. **Enhanced MessageModel with Type Safety**

```go
type MessageModel struct {
    base.MessageHeader `json:",inline"`
    
    // Core fields present in all versions (V1+)
    Pagenation           models.MessagePagenation              `json:"pagenation"`
    ReportId             models.ReportType                     `json:"reportId"`
    ReportCreateDateTime time.Time                             `json:"reportCreateDateTime"`
    AccountOtherId       string                                `json:"accountOtherId"`
    
    // Version-specific fields with typed accessors
    businessQueryFields  *BusinessQueryFields  `json:",inline,omitempty"`
    reportingFields      *ReportingFields      `json:",inline,omitempty"`
}

// Type-safe version-specific field groups
type BusinessQueryFields struct {
    BussinessQueryMsgId          string    `json:"bussinessQueryMsgId,omitempty"`
    BussinessQueryMsgNameId      string    `json:"bussinessQueryMsgNameId,omitempty"`
    BussinessQueryCreateDatetime time.Time `json:"bussinessQueryCreateDatetime,omitempty"`
}

type ReportingFields struct {
    ReportingSequence models.SequenceRange `json:"reportingSequence,omitempty"`
}

// Type-safe version-aware constructors
func NewMessageForVersion(version CAMT_052_001_VERSION) MessageModel {
    model := MessageModel{
        MessageHeader: base.MessageHeader{},
        // Core fields initialized
    }
    
    // Type-safe version-specific field initialization
    switch {
    case version >= CAMT_052_001_04:
        model.businessQueryFields = &BusinessQueryFields{}
        fallthrough
    case version >= CAMT_052_001_07:
        model.reportingFields = &ReportingFields{}
    }
    
    return model
}
```

### 6. **Type-Safe Validation Methods**

```go
// Implement MessageWithVersionFields interface
func (m MessageModel) ValidateForVersion(version CAMT_052_001_VERSION) error {
    // Base validation always applies
    if err := m.validateBaseFields(); err != nil {
        return err
    }
    
    // Type-safe version-specific validation
    switch {
    case version >= CAMT_052_001_07:
        if m.reportingFields == nil {
            return fmt.Errorf("ReportingFields required for version %v", version)
        }
        if err := m.reportingFields.Validate(); err != nil {
            return fmt.Errorf("ReportingFields validation failed: %w", err)
        }
        fallthrough
    case version >= CAMT_052_001_04:
        if m.businessQueryFields == nil {
            return fmt.Errorf("BusinessQueryFields required for version %v", version)
        }
        if err := m.businessQueryFields.Validate(); err != nil {
            return fmt.Errorf("BusinessQueryFields validation failed: %w", err)
        }
    }
    
    return nil
}

// Type-safe field group validation
func (b *BusinessQueryFields) Validate() error {
    if b.BussinessQueryMsgId == "" {
        return fmt.Errorf("BussinessQueryMsgId is required")
    }
    return nil
}

func (r *ReportingFields) Validate() error {
    if r.ReportingSequence.FromSeq == "" {
        return fmt.Errorf("ReportingSequence.FromSeq is required")
    }
    return nil
}
```

## Benefits of Type-Safe Approach

### ✅ **Compile-Time Safety**
- Field references are checked at compile time
- Version-field relationships are type-verified
- Refactoring automatically updates all references

### ✅ **Clear Version Semantics**
- Version-specific field groups are explicit
- Constructor patterns enforce version compatibility
- Clear upgrade paths between versions

### ✅ **Performance Benefits**
- No reflection overhead during validation
- Direct field access vs string-based lookups
- Inlined validation functions

### ✅ **Developer Experience**
- IDE autocomplete for all field operations
- Compile-time error detection
- Clear version upgrade documentation

## Migration Strategy

### Phase 1: Add Type-Safe Interfaces
1. Define `MessageWithVersionFields` interface
2. Add version-specific field group types
3. Implement type-safe validation methods

### Phase 2: Enhance MessageModel
1. Group version-specific fields into typed structs
2. Add version-aware constructors
3. Implement interface methods

### Phase 3: Replace Reflection-Based Validation
1. Remove `validateRequiredFieldsReflection`
2. Replace with interface-based validation
3. Update all validation calls

### Phase 4: Test & Verify
1. Comprehensive version testing
2. Performance benchmarking
3. Type safety verification

## Implementation Example

```go
// Type-safe usage
model := NewMessageForVersion(CAMT_052_001_07)
model.MessageId = "TEST123"
model.CreatedDateTime = time.Now()

// Compile-time version validation
if err := model.ValidateForVersion(CAMT_052_001_07); err != nil {
    return fmt.Errorf("validation failed: %w", err)
}

// Type-safe XML processing
xmlData, err := model.WriteXMLForVersion(CAMT_052_001_07)
if err != nil {
    return err
}
```

This approach maintains the original type safety goals while providing sophisticated version-specific validation capabilities.