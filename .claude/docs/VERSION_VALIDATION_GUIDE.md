# Version-Specific Validation Guide for wire20022

This guide explains how to implement and validate version-specific fields when adding new ISO 20022 message types to the wire20022 library.

## Overview

ISO 20022 message types evolve over time, with new versions adding, removing, or modifying fields. The wire20022 library uses a sophisticated approach to handle version-specific field availability while maintaining clean, idiomatic Go APIs.

## Architecture Principles

### 1. **Unified Go Struct with Version-Aware Processing**

- **Single MessageModel struct** contains ALL fields from ALL versions
- **Version-specific path mapping** controls which fields are populated from XML
- **JSON omitempty tags** prevent empty version-specific fields in JSON output
- **Field organization** groups fields by version introduction for clarity

### 2. **Version-Specific Validation**

- **RequiredFieldsByVersion map** defines required fields per version
- **CheckRequiredFieldsForVersion** function validates specific versions
- **GetVersionSpecificFields** provides field availability mapping

## Implementation Guide

### Step 1: Organize MessageModel by Version

Structure your MessageModel to clearly show version evolution:

```go
type MessageModel struct {
    // Embed common message fields
    base.MessageHeader `json:",inline"`

    // Core fields present in all versions (V1+)
    CoreField1     string    `json:"coreField1"`
    CoreField2     time.Time `json:"coreField2"`

    // Version 3+ fields (omitempty prevents JSON output when empty)
    V3Field1       string    `json:"v3Field1,omitempty"`
    V3Field2       string    `json:"v3Field2,omitempty"`

    // Version 7+ fields
    V7Field1       models.SomeType `json:"v7Field1,omitempty"`
    V7Field2       string          `json:"v7Field2,omitempty"`
}
```

### Step 2: Define Version-Specific Required Fields

Create a map defining which fields are required for each version:

```go
// Version-specific required fields map
var RequiredFieldsByVersion = map[YOUR_VERSION_TYPE][]string{
    // V1/V2: Basic fields only
    VERSION_01: {"MessageId", "CreatedDateTime", "CoreField1", "CoreField2"},
    VERSION_02: {"MessageId", "CreatedDateTime", "CoreField1", "CoreField2"},
    
    // V3+: Add new required fields
    VERSION_03: {"MessageId", "CreatedDateTime", "CoreField1", "CoreField2", "V3Field1"},
    VERSION_04: {"MessageId", "CreatedDateTime", "CoreField1", "CoreField2", "V3Field1"},
    
    // V7+: Add more required fields
    VERSION_07: {"MessageId", "CreatedDateTime", "CoreField1", "CoreField2", "V3Field1", "V7Field1"},
}

// Maintain backward compatibility
var RequiredFields = RequiredFieldsByVersion[VERSION_LATEST]
```

### Step 3: Implement Version-Specific Validation

Add validation functions for specific versions:

```go
// CheckRequiredFieldsForVersion validates required fields for a specific version
func CheckRequiredFieldsForVersion(model MessageModel, version YOUR_VERSION_TYPE) error {
    requiredFields, exists := RequiredFieldsByVersion[version]
    if !exists {
        requiredFields = RequiredFields
    }
    
    return validateRequiredFieldsReflection(model, requiredFields)
}

// validateRequiredFieldsReflection performs version-aware validation using reflection
func validateRequiredFieldsReflection(model MessageModel, requiredFields []string) error {
    modelValue := reflect.ValueOf(model)
    modelType := reflect.TypeOf(model)
    
    for _, fieldName := range requiredFields {
        var fieldValue reflect.Value
        var found bool
        
        // Check direct fields
        if _, ok := modelType.FieldByName(fieldName); ok {
            fieldValue = modelValue.FieldByName(fieldName)
            found = true
        } else {
            // Check embedded MessageHeader fields
            headerField := modelValue.FieldByName("MessageHeader")
            if headerField.IsValid() {
                if _, ok := headerField.Type().FieldByName(fieldName); ok {
                    fieldValue = headerField.FieldByName(fieldName)
                    found = true
                }
            }
        }
        
        if !found {
            return fmt.Errorf("field %s not found", fieldName)
        }
        
        if isEmpty := isFieldEmpty(fieldValue); isEmpty {
            return fmt.Errorf("required field %s is empty", fieldName)
        }
    }
    
    return nil
}

// isFieldEmpty checks if a field value is considered empty
func isFieldEmpty(value reflect.Value) bool {
    switch value.Kind() {
    case reflect.String:
        return value.String() == ""
    case reflect.Struct:
        if value.Type().String() == "time.Time" {
            return value.Interface().(time.Time).IsZero()
        }
        return value.IsZero()
    case reflect.Slice, reflect.Array:
        return value.Len() == 0
    case reflect.Ptr, reflect.Interface:
        return value.IsNil()
    default:
        return value.IsZero()
    }
}
```

### Step 4: Add Field Availability Documentation

Create a function that documents which fields are available in which versions:

```go
// GetVersionSpecificFields returns a map indicating which fields are available for each version
func GetVersionSpecificFields() map[string][]YOUR_VERSION_TYPE {
    return map[string][]YOUR_VERSION_TYPE{
        "V3Field1": {VERSION_03, VERSION_04, VERSION_05, VERSION_06, VERSION_07},
        "V3Field2": {VERSION_03, VERSION_04, VERSION_05, VERSION_06, VERSION_07},
        "V7Field1": {VERSION_07},
        "V7Field2": {VERSION_07},
    }
}
```

### Step 5: Update Path Mapping

Ensure your `map.go` file only maps fields that exist in each version:

```go
func PathMapV1() map[string]any {
    return map[string]any{
        "Root.CoreField1": "CoreField1",
        "Root.CoreField2": "CoreField2",
        // No V3+ or V7+ fields
    }
}

func PathMapV3() map[string]any {
    return map[string]any{
        "Root.CoreField1": "CoreField1",
        "Root.CoreField2": "CoreField2",
        "Root.V3Field1":   "V3Field1",    // New in V3
        "Root.V3Field2":   "V3Field2",    // New in V3
        // No V7+ fields yet
    }
}

func PathMapV7() map[string]any {
    return map[string]any{
        "Root.CoreField1": "CoreField1",
        "Root.CoreField2": "CoreField2",
        "Root.V3Field1":   "V3Field1",
        "Root.V3Field2":   "V3Field2",
        "Root.V7Field1":   "V7Field1",    // New in V7
        "Root.V7Field2":   "V7Field2",    // New in V7
    }
}
```

## Usage Examples

### Basic Validation (Latest Version)

```go
model := MessageModel{
    MessageHeader: base.MessageHeader{
        MessageId:       "TEST123",
        CreatedDateTime: time.Now(),
    },
    CoreField1: "required_value",
    CoreField2: time.Now(),
}

// Validate using latest version requirements
if err := CheckRequiredFields(model); err != nil {
    log.Printf("Validation failed: %v", err)
}
```

### Version-Specific Validation

```go
// Validate for a specific version
if err := CheckRequiredFieldsForVersion(model, VERSION_03); err != nil {
    log.Printf("V3 validation failed: %v", err)
}

// Check field availability for a version
fields := GetVersionSpecificFields()
if versions, exists := fields["V7Field1"]; exists {
    fmt.Printf("V7Field1 is available in versions: %v\n", versions)
}
```

### XML Processing with Version Detection

```go
// ReadXML automatically detects version and maps appropriate fields
var model MessageModel
if err := model.ReadXML(xmlReader); err != nil {
    return err
}

// WriteXML can target specific versions
if err := model.WriteXML(xmlWriter, VERSION_07); err != nil {
    return err
}
```

## Testing Version-Specific Functionality

### Unit Test Structure

```go
func TestVersionSpecificValidation(t *testing.T) {
    model := MessageModel{
        MessageHeader: base.MessageHeader{
            MessageId:       "TEST123",
            CreatedDateTime: time.Now(),
        },
        CoreField1: "value",
        CoreField2: time.Now(),
        // No V3+ fields populated
    }

    // V1/V2 should pass
    assert.NoError(t, CheckRequiredFieldsForVersion(model, VERSION_02))

    // V3+ should fail (missing required V3 fields)
    assert.Error(t, CheckRequiredFieldsForVersion(model, VERSION_03))

    // Add V3 fields
    model.V3Field1 = "v3_value"

    // V3+ should now pass
    assert.NoError(t, CheckRequiredFieldsForVersion(model, VERSION_03))
}
```

### XML Round-Trip Testing

```go
func TestXMLRoundTrip(t *testing.T) {
    versions := []VERSION_TYPE{VERSION_02, VERSION_03, VERSION_07}
    
    for _, version := range versions {
        t.Run(fmt.Sprintf("Version_%s", version), func(t *testing.T) {
            // Create model with fields appropriate for version
            model := createModelForVersion(version)
            
            // Write to XML
            var buf bytes.Buffer
            err := model.WriteXML(&buf, version)
            assert.NoError(t, err)
            
            // Read back from XML
            var parsedModel MessageModel
            err = parsedModel.ReadXML(&buf)
            assert.NoError(t, err)
            
            // Validate version-specific requirements
            err = CheckRequiredFieldsForVersion(parsedModel, version)
            assert.NoError(t, err)
        })
    }
}
```

## Best Practices

### 1. **Field Organization**
- Group fields by version introduction in the struct definition
- Use clear comments indicating version requirements
- Add `omitempty` JSON tags for version-specific fields

### 2. **Validation Strategy**
- Always provide version-specific validation functions
- Use reflection for flexible field validation
- Handle embedded struct fields correctly

### 3. **Documentation**
- Document field availability with `GetVersionSpecificFields()`
- Provide clear version evolution comments
- Include usage examples for each version

### 4. **Testing**
- Test all supported versions
- Verify XML round-trip compatibility
- Test validation with missing required fields

### 5. **Error Handling**
- Provide clear error messages with field names
- Handle unknown versions gracefully
- Use structured error types when appropriate

## Common Pitfalls

### 1. **Forgetting omitempty Tags**
Without `omitempty`, version-specific fields will appear in JSON output even when empty.

### 2. **Inconsistent Path Mapping**
Ensure path mapping functions only include fields available in that version.

### 3. **Missing Validation**
Always implement version-specific validation to catch missing required fields.

### 4. **Embedded Field Handling**
Remember to check embedded `MessageHeader` fields during validation.

### 5. **Version Evolution**
When adding new versions, update all related functions: path mapping, validation, and field availability.

## Integration with Base Abstractions

The version-specific validation system integrates seamlessly with the existing base abstractions:

- **Uses base.MessageHeader** for common fields
- **Leverages base.MessageProcessor** for XML processing  
- **Maintains compatibility** with existing validation patterns
- **Extends functionality** without breaking changes

This approach ensures that version-specific handling is consistent across all message types while maintaining the clean, type-safe APIs that wire20022 provides.