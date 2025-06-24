# Test Coverage Analysis and Improvement Plan

## Current Coverage Status
- **Overall Project Coverage**: 44.2%
- **Target Coverage**: >50% (immediate), >70% (ideal for library)

## Coverage Analysis by Package

### High Coverage Packages (Good Examples)
1. **pkg/errors**: 100% - Complete coverage ✅
2. **pkg/base**: 86.6% - Excellent coverage ✅
3. **pkg/models/ActivityReport**: 69.6% - Good coverage

### Critical Low Coverage Areas

#### 1. Main Entry Point (0% Coverage)
- `cmd/wire20022/main.go`: 0% coverage
- **Impact**: Critical - main application entry point untested
- **Priority**: Low (CLI not primary use case)

#### 2. Message Types with Poor Coverage (<35%)
These message types need comprehensive testing:
- **ConnectionCheck**: 27.5%
- **DrawdownRequest**: 29.9%
- **PaymentStatusRequest**: 29.6%
- **Master**: 32.0%
- **PaymentReturn**: 33.7%
- **FedwireFundsPaymentStatus**: 33.7%
- **ReturnRequestResponse**: 33.0%
- **EndpointTotalsReport**: 34.0%

#### 3. Critical Untested Functions (0% Coverage)

##### XML Processing Functions (CRITICAL for XML-first library)
- All `ReadXML()` methods except AccountReportingRequest/ActivityReport
- All `WriteXML()` methods except AccountReportingRequest/ActivityReport/CustomerCreditTransfer
- Most `UnmarshalJSON()` methods

##### Validation Functions (CRITICAL for data integrity)
- `ValidateForVersion()` - Version-specific validation
- `validateCoreFields()` - Core field validation
- `GetVersionCapabilities()` - Version feature detection

##### Helper Functions (Important for usability)
- All `MessageHelper.go` BuildHelper functions
- Type helper functions in `typeHelper.go`

#### 4. Messages Package (38.1% Coverage)
- Missing tests for `Validate()` and `ConvertXMLToModel()` in processor
- Partial coverage for constructor functions

## Test Coverage Improvement Plan

### Phase 1: Critical Path Coverage (Target: 55%)

#### 1.1 Add XML Processing Tests (High Priority)
For each message type missing XML tests:
```go
// Add to each Message_idiomatic_test.go
func TestReadWriteXML(t *testing.T)
func TestParseXML(t *testing.T)
func TestWriteXMLWithVersions(t *testing.T)
```

#### 1.2 Add Validation Tests (High Priority)
```go
func TestValidateForVersion(t *testing.T)
func TestValidateCoreFields(t *testing.T)
func TestGetVersionCapabilities(t *testing.T)
```

#### 1.3 Add UnmarshalJSON Tests
```go
func TestUnmarshalJSON(t *testing.T)
func TestUnmarshalJSONWithVersionFields(t *testing.T)
```

### Phase 2: Helper Function Coverage (Target: 65%)

#### 2.1 MessageHelper Tests
Create helper tests for each message type:
```go
func TestBuildMessageHelper(t *testing.T)
func TestHelperFieldDocumentation(t *testing.T)
```

#### 2.2 TypeHelper Tests
```go
func TestBuildPostalAddressHelper(t *testing.T)
func TestBuildAgentHelper(t *testing.T)
func TestBuildCurrencyAndAmountHelper(t *testing.T)
```

### Phase 3: Comprehensive Coverage (Target: 70%+)

#### 3.1 Edge Cases and Error Scenarios
- Invalid XML parsing
- Missing required fields
- Version mismatch scenarios
- Malformed data handling

#### 3.2 Integration Tests
- Full message lifecycle tests
- Cross-version compatibility tests
- Performance benchmarks

## Implementation Strategy

### 1. Create Test Template Script
```bash
# .claude/scripts/generate_idiomatic_tests.py
# Generates consistent test structure for all message types
```

### 2. Priority Order for Implementation
1. **ConnectionCheck** - Simplest message type, good starting point
2. **DrawdownRequest/DrawdownResponse** - Paired messages
3. **PaymentStatusRequest** - Important for payment tracking
4. **Master** - Complex but critical message type
5. **PaymentReturn** - High-value business logic
6. **FedwireFundsPaymentStatus** - Status reporting
7. **ReturnRequestResponse** - Completes return flow
8. **EndpointTotalsReport** - Reporting functionality

### 3. Test Data Strategy
- Use existing `swiftSample/` directories for test data
- Create minimal valid XML for each version
- Test both valid and invalid scenarios

## Metrics to Track

1. **Coverage Percentage** - Overall and by package
2. **Critical Path Coverage** - XML/Validation/Core functions
3. **Test Execution Time** - Keep under 30 seconds
4. **Test Maintainability** - Use consistent patterns

## Success Criteria

- [ ] Overall coverage > 50% (Phase 1)
- [ ] All XML processing functions tested
- [ ] All validation functions tested
- [ ] No message type below 40% coverage
- [ ] Critical packages (base, messages) > 60%
- [ ] Documentation for test patterns

## HTML Coverage Report

The HTML coverage report has been generated at:
`.claude/coverage/coverage.html`

Open this file in a browser to:
- View line-by-line coverage
- Identify untested code paths
- Navigate through packages interactively
- See coverage heat maps