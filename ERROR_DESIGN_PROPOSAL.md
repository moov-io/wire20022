# Error Design Proposal: Enhanced XML Path Information

## Current State

The library currently uses domain-specific error types that follow Go standard library conventions:
- `ValidationError` - Field validation failures
- `ParseError` - XML/JSON parsing failures  
- `FieldError` - Field access/manipulation failures

Error messages include XML paths in the message text, which is helpful for debugging but not programmatically accessible.

## Proposed Enhancement

Enhance the `FieldError` type to include both Go field paths and XML paths as separate properties, while maintaining idiomatic Go error messages.

### Enhanced FieldError Structure

```go
// FieldError represents an error accessing or setting a field value.
// It provides context about both the Go field path and XML path for debugging.
type FieldError struct {
    Path      string // The Go struct field path (e.g., "MessageId")
    XMLPath   string // The XML element path (e.g., "FIToFICstmrCdtTrf.GrpHdr.MsgId")
    Operation string // The operation being performed (e.g., "get", "set", "copy")
    Err       error  // Underlying error
}

// Error implements the error interface with idiomatic Go messages
func (e *FieldError) Error() string {
    if e.XMLPath != "" {
        return fmt.Sprintf("field %s %s failed: %v", e.Operation, e.Path, e.Err)
    }
    return fmt.Sprintf("field %s %s failed: %v", e.Operation, e.Path, e.Err)
}

// GetXMLPath returns the XML path for programmatic access
func (e *FieldError) GetXMLPath() string {
    return e.XMLPath
}

// GetGoPath returns the Go struct field path
func (e *FieldError) GetGoPath() string {
    return e.Path
}
```

### Usage Examples

```go
// Creating enhanced field errors
err := NewFieldErrorWithXML("MessageId", "FIToFICstmrCdtTrf.GrpHdr.MsgId", "copy", validationErr)

// Error message remains idiomatic
fmt.Println(err.Error())
// Output: "field copy MessageId failed: failed to set MessageId: value too long"

// XML path available programmatically
if fieldErr, ok := err.(*FieldError); ok {
    xmlPath := fieldErr.GetXMLPath()  // "FIToFICstmrCdtTrf.GrpHdr.MsgId"
    goPath := fieldErr.GetGoPath()    // "MessageId"
}
```

## Benefits

1. **Backward Compatibility**: Error messages remain idiomatic and follow Go conventions
2. **Enhanced Debugging**: XML path information available programmatically
3. **Better Tooling**: IDEs and tools can extract structured error information
4. **Maintained Performance**: No additional overhead in error creation
5. **Standard Library Alignment**: Follows patterns from Go's standard library

## Implementation Plan

### Phase 1: Enhance Error Types
- Add `XMLPath` field to `FieldError`
- Add `GetXMLPath()` and `GetGoPath()` methods
- Add `NewFieldErrorWithXML()` constructor
- Maintain backward compatibility with existing constructors

### Phase 2: Update Error Creation
- Modify `HandleFieldCopyError` to include XML path information
- Update base processor to pass XML paths to error constructors
- Ensure all field copy operations include XML context

### Phase 3: Update Tests
- Update test assertions to match new error message format
- Add tests for programmatic access to XML paths
- Verify backward compatibility

## Example Implementation

```go
// Enhanced constructor
func NewFieldErrorWithXML(goPath, xmlPath, operation string, cause error) *FieldError {
    return &FieldError{
        Path:      goPath,
        XMLPath:   xmlPath,
        Operation: operation,
        Err:       cause,
    }
}

// Updated HandleFieldCopyError
func HandleFieldCopyError(goPath, xmlPath string, err error) error {
    return wirerrors.NewFieldErrorWithXML(goPath, xmlPath, "copy", err)
}

// Usage in processor
for sourcePath, targetPath := range rePathMap {
    if err := models.CopyMessageValueToDocument(&message, sourcePath, doc, targetPath); err != nil {
        return nil, HandleFieldCopyError(sourcePath, targetPath, err)
    }
}
```

## Error Message Examples

### Current Format
```
field copy FIToFICstmrCdtTrf.GrpHdr.MsgId failed: failed to set MessageId: value too long
```

### Proposed Format
```
field copy MessageId failed: failed to set MessageId: value too long
```

The XML path (`FIToFICstmrCdtTrf.GrpHdr.MsgId`) becomes available via `GetXMLPath()` method.

## Migration Strategy

1. **Non-Breaking Change**: All existing error handling continues to work
2. **Gradual Adoption**: New XML path features are opt-in
3. **Test Compatibility**: Existing test assertions continue to pass
4. **Documentation**: Update guides to show new programmatic access patterns

## Considerations

### Error Message Format
- Keep Go struct field names in error messages (more meaningful to Go developers)
- XML paths available programmatically for debugging tools
- Maintain consistency with Go standard library error patterns

### Performance Impact
- Minimal: Only adds one string field to error struct
- No allocation overhead in normal (non-error) code paths
- Error creation is already expensive relative to field access

### Developer Experience
- More debugging information available when needed
- Error messages remain readable and Go-idiomatic
- Tooling can extract structured information for better debugging

## Alternative Considered

### XML Path in Error Message
**Pros**: Immediately visible in logs
**Cons**: Error messages become non-idiomatic, harder to read, inconsistent with Go conventions

### Separate XMLError Type
**Pros**: Clear separation of concerns
**Cons**: Complex error handling, multiple error types for similar operations

### Context in Error Wrapping
**Pros**: Uses Go 1.13+ error wrapping
**Cons**: Less structured, harder for tools to extract information

## Recommendation

Implement the enhanced `FieldError` with separate `XMLPath` property. This provides the best balance of:
- Idiomatic Go error messages
- Programmatically accessible XML debugging information
- Backward compatibility
- Performance

The error messages remain focused on Go development while providing XML context for debugging when needed.