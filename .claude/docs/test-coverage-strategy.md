# Test Coverage Improvement Strategy

## Current State

### Coverage Analysis
- **Message Types**: 86-96% coverage (good)
- **Wrapper Layer**: 0% coverage (critical gap)
- **Base Utilities**: 0% coverage (critical gap)
- **Server**: 0% coverage (needs implementation)

## Critical Gaps Identified

### 1. Wrapper Layer (Priority: CRITICAL)
**Location**: `pkg/wrapper/`
**Current Coverage**: 0%
**Risk**: High - main API interface

**Untested Functions**:
- `CreateDocument()` - JSON to XML conversion
- `ValidateDocument()` - Document validation
- `CheckRequireField()` - Field validation
- `ConvertXMLToModel()` - XML to model parsing
- `GetHelp()` - Help documentation generation

### 2. Base Utilities (Priority: CRITICAL)
**Location**: `pkg/models/util.go`
**Current Coverage**: 0%
**Risk**: High - core parsing logic

**Untested Functions**:
- `DocumentFrom()` - XML document parsing with namespace resolution
- `GetElement()` - Reflection-based field access with array indexing
- `SetElementToDocument()` - Complex nested value setting
- `RemakeMapping()` - Dynamic path mapping
- File I/O operations

### 3. Message Type Gaps (Priority: MEDIUM)
**Location**: Various message type packages
**Current Coverage**: 86-96%
**Risk**: Medium - missing edge cases

**Missing Coverage**:
- `DocumentWith()` error paths
- `CheckRequiredFields()` validation scenarios
- Version-specific path mapping
- Namespace resolution edge cases

## Test Implementation Strategy

### Phase 1: Wrapper Layer Tests (Week 1)
**Goal**: Achieve 90%+ coverage for wrapper layer

#### Test Structure:
```
pkg/wrapper/
├── CustomerCreditTransfer_test.go
├── PaymentReturn_test.go
├── testdata/
│   ├── valid_models.json
│   ├── invalid_models.json
│   ├── valid_documents.xml
│   └── invalid_documents.xml
```

#### Test Categories:
1. **Happy Path Tests**
   - Valid JSON to XML conversion
   - Valid XML to model parsing
   - Successful validation
   - Help generation

2. **Error Handling Tests**
   - Invalid JSON input
   - Malformed XML input
   - Missing required fields
   - Unsupported versions
   - JSON/XML marshaling failures

3. **Edge Case Tests**
   - Empty input data
   - Null/nil values
   - Large documents
   - Special characters in data

### Phase 2: Base Utilities Tests (Week 2)
**Goal**: Achieve 85%+ coverage for util.go

#### Critical Test Areas:
1. **DocumentFrom() Function**
   ```go
   func TestDocumentFrom(t *testing.T) {
       // Test valid XML with known namespace
       // Test XML without xmlns attribute
       // Test unknown namespace
       // Test malformed XML
       // Test factory instantiation errors
   }
   ```

2. **GetElement() Function**
   ```go
   func TestGetElement(t *testing.T) {
       // Test simple field access
       // Test nested struct access
       // Test array indexing (Field[0])
       // Test invalid paths
       // Test nil pointer handling
       // Test index out of bounds
   }
   ```

3. **SetElementToDocument() Function**
   ```go
   func TestSetElementToDocument(t *testing.T) {
       // Test simple value setting
       // Test nested path creation
       // Test array index setting
       // Test type conversions
       // Test validation failures
   }
   ```

### Phase 3: Message Type Improvements (Week 3)
**Goal**: Achieve 95%+ coverage for all message types

#### Focus Areas:
1. **Error Path Testing**
   - Required field validation failures
   - Invalid version handling
   - Path mapping errors

2. **Version Testing**
   - Test all supported versions (v02-v12)
   - Version-specific behavior differences
   - Namespace mapping accuracy

3. **Integration Testing**
   - End-to-end document creation workflows
   - Cross-version compatibility
   - Real-world SWIFT sample processing

## Test Data Strategy

### 1. Golden Files
Create comprehensive test data sets:
```
testdata/
├── samples/
│   ├── pacs008/
│   │   ├── v02_valid.xml
│   │   ├── v12_valid.xml
│   │   └── invalid_missing_fields.xml
│   └── models/
│       ├── complete_model.json
│       ├── minimal_model.json
│       └── invalid_model.json
```

### 2. Generated Test Cases
- Property-based testing for complex scenarios
- Fuzz testing for parsing robustness
- Boundary value testing for numeric fields

### 3. Real-World Samples
- Use actual SWIFT samples from existing test directories
- Add edge cases found in production usage
- Include regulatory compliance examples

## Testing Tools and Patterns

### 1. Test Structure Pattern
```go
func TestWrapper_Method(t *testing.T) {
    tests := []struct {
        name        string
        input       interface{}
        expected    interface{}
        expectError bool
        errorMsg    string
    }{
        // Test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### 2. Mock and Stub Strategy
- Mock external dependencies (file I/O)
- Stub complex fedwire20022 dependencies
- Use interfaces for testability

### 3. Coverage Tools
```bash
# Generate detailed coverage report
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out -o coverage.html

# Check coverage by function
go tool cover -func=cover.out
```

## Success Metrics

### Target Coverage Goals
- **Overall Project**: 90%+
- **Wrapper Layer**: 95%+
- **Base Utilities**: 90%+
- **Message Types**: 95%+

### Quality Metrics
- All error paths tested
- All public APIs have examples
- No untested critical functions
- Integration tests cover main workflows

## Implementation Order

### Week 1: Wrapper Tests
1. Create test directory structure
2. Implement CustomerCreditTransfer wrapper tests
3. Create reusable test data and helpers
4. Achieve 90%+ wrapper coverage

### Week 2: Utility Tests
1. Test DocumentFrom() function thoroughly
2. Test reflection-based utilities
3. Test file I/O operations
4. Achieve 85%+ util.go coverage

### Week 3: Message Type Completion
1. Fill gaps in existing message type tests
2. Add error path testing
3. Add version compatibility tests
4. Achieve 95%+ overall coverage

### Week 4: Integration and Performance
1. Add end-to-end integration tests
2. Add performance benchmarks
3. Add fuzz testing for robustness
4. Document testing patterns for future development

## Risk Mitigation

### High-Risk Areas
1. **Reflection-based operations** - Test thoroughly with various data types
2. **XML parsing edge cases** - Test malformed and edge-case XML
3. **Error propagation** - Ensure errors bubble up correctly
4. **Memory leaks** - Test with large documents and many iterations

### Testing Best Practices
1. Test both success and failure paths
2. Use meaningful test names describing the scenario
3. Include edge cases and boundary conditions
4. Test with real-world data when possible
5. Keep tests fast and independent
6. Use table-driven tests for multiple scenarios