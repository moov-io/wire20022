# Refactoring Plan for wire20022

This document outlines the refactoring needed to bring this library to Go standard library quality. Each section represents a feature branch, and each TODO is sized for a single commit.

## Branch 1: Fix Critical Bugs and Typos

### Status: ✅ **COMPLETED**

### Completed TODOs:
- [x] Fix typo: Rename `SericeLevel` to `ServiceLevel` in CustomerCreditTransfer MessageModel
- [x] Fix typo: Rename `PACS_008_001_VESION` to `PACS_008_001_VERSION` in all version constants
- [x] Fix missing return statement in CustomerCreditTransferWrapper.ConvertXMLToModel
- [x] Fix Go version in go.mod from 1.24.0 to latest stable version (1.23.x)
- [x] Remove or implement commented-out main.go functionality
- [x] Fix path remapping inconsistencies between message types

## Branch 2: Implement Proper Error Handling (Idiomatic Go)

### Status: ✅ **COMPLETED**

### Completed TODOs:
- [x] Create concrete error types with Error() method (no interfaces): ValidationError, ParseError
- [x] Use error wrapping with fmt.Errorf("operation failed: %w", err) for context
- [x] Return sentinel errors as package variables: var ErrInvalidField = errors.New("invalid field")
- [x] Add Is() and As() methods to custom errors for errors.Is() and errors.As() compatibility
- [x] Replace GetElement nil returns with zero values and explicit error returns
- [x] Use errors.Join() for multiple validation errors (Go 1.20+)
- [x] Create error constructors: NewValidationError(field, reason string) error
- [x] Document error behavior in function comments with "returns ErrX if Y" pattern
- [x] Use panic only for programmer errors, not user input errors
- [x] Avoid error strings starting with capital letters or ending with punctuation
- [x] Return concrete error types, not error interfaces, from constructors

### Additional Fixes Completed:
- [x] Fixed exhaustive linter errors in ValidateNotNil and setValue functions
- [x] Fixed string-to-numeric conversion support in setValue function
- [x] Fixed fedwire.ISODate zero value marshaling issues
- [x] Fixed WriteXMLToGenerate function path handling
- [x] Fixed EndpointTotalsReport field mapping and version compatibility issues
- [x] Created version-specific path mappings for different XML schema versions
- [x] Updated test assertions to match new error message paths after idiomatic error handling refactor
- [x] Fixed FedwireFundsPaymentStatus and FedwireFundsSystemResponse test expectations to align with updated field validation paths

## Branch 3: Create Base Abstractions to Reduce Duplication (Idiomatic Go)

### Status: **READY TO BEGIN** - Critical foundation completed

### Code Duplication Analysis Summary
**Impact**: ~2,400+ lines of duplicated code across 16+ message types
**Files Affected**: All Message.go, version.go files in pkg/models/*/
**Maintenance Burden**: High - any change requires updates to 16+ files
**Priority**: Critical for developer experience and maintainability

### Priority 1: Generic Message Processing Framework (Critical)
**Problem**: Identical MessageWith/DocumentWith logic across all message types
**Impact**: ~1,500 lines of duplicated code, error-prone manual maintenance

- [ ] Create generic MessageProcessor[T, V] struct in pkg/common/processor.go
- [ ] Implement MessageProcessor.MessageWith(data []byte) (T, error) using generics
- [ ] Implement MessageProcessor.DocumentWith(model T, version V) (models.ISODocument, error)
- [ ] Extract common XML parsing and version handling logic
- [ ] Migrate CustomerCreditTransfer to use generic processor (proof of concept)
- [ ] Migrate remaining message types to generic processor
- [ ] Remove duplicated MessageWith/DocumentWith functions from all message types

### Priority 2: Type-Safe Validation Framework (High)
**Problem**: Reflection-based CheckRequiredFields with manual field mapping
**Impact**: ~400 lines of error-prone validation code, runtime errors instead of compile-time

- [ ] Create Validatable interface with Validate() error method
- [ ] Replace CheckRequiredFields with type-specific Validate() methods
- [ ] Implement compile-time type-safe validation for CustomerCreditTransfer
- [ ] Create validation utilities in pkg/common/validation.go
- [ ] Remove reflection-based field mapping from all message types
- [ ] Add validation error aggregation with errors.Join()
- [ ] Migrate all message types to type-safe validation

### Priority 3: Unified Version Management (High)
**Problem**: Inconsistent version handling patterns and naming conventions
**Impact**: ~800 lines of duplicated version management code

- [ ] Create VersionManager[V] generic type in pkg/common/version.go
- [ ] Standardize version string formats (resolve lowercase vs uppercase inconsistencies)
- [ ] Create unified namespace mapping patterns
- [ ] Extract common VersionPathMap handling logic
- [ ] Migrate message types to unified version management
- [ ] Remove duplicated version.go files where possible

### Priority 4: Factory Pattern Abstraction (Medium)
**Problem**: Identical NameSpaceModelMap factory patterns across all types
**Impact**: ~600 lines of repetitive factory code

- [ ] Create generic DocumentFactory[T] type
- [ ] Extract common document factory creation patterns
- [ ] Standardize XMLName construction logic
- [ ] Create factory generator utilities
- [ ] Migrate message types to shared factory patterns

### Additional Improvements:
- [ ] Move XML/JSON conversion to pkg/convert/ with functions like XMLToModel(data []byte, target any)
- [ ] Use type embedding: embed common fields in structs rather than interfaces
- [ ] Replace getter/setter patterns with direct field access on exported structs
- [ ] Create factory functions instead of builder patterns: NewCustomerCreditTransfer() MessageModel
- [ ] Use type switches instead of interface{} where multiple concrete types are handled
- [ ] Eliminate MessageHandler interface - use concrete function types instead
- [ ] Create shared constants and types in pkg/common/types.go for reused structures

### Implementation Strategy:
1. **Start with CustomerCreditTransfer** as proof of concept for generic processor
2. **Measure before/after** - lines of code, test coverage, performance
3. **Incremental migration** - one message type at a time to minimize breakage
4. **Maintain backward compatibility** during transition where possible
5. **Update XML_TO_GO_MAPPING.md** to reflect new abstractions

## Branch 4: Improve Test Coverage and Quality

### TODOs:
- [ ] Create test data factory functions for each message type
- [ ] Add negative test cases for validation failures
- [ ] Replace hardcoded dates with relative date calculations
- [ ] Add benchmarks for critical paths (parsing, validation)
- [ ] Create integration tests for complete message flows
- [ ] Add fuzzing tests for XML/JSON parsing
- [ ] Implement table-driven tests for version compatibility
- [ ] Add test coverage reporting to CI pipeline
- [ ] Create examples directory with usage examples

## Branch 5: Standardize API and Interfaces

### TODOs:
- [ ] Standardize wrapper function signatures (all use []byte or all use string)
- [ ] Create consistent naming conventions for all public APIs
- [ ] Define clear interface boundaries between packages
- [ ] Implement builder pattern for complex message creation
- [ ] Add functional options pattern for configuration
- [ ] Create consistent validation API across all message types
- [ ] Standardize error return patterns across all functions
- [ ] Document all public APIs with examples

## Branch 6: Performance Optimizations

### TODOs:
- [ ] Add object pooling for frequently created objects
- [ ] Implement lazy loading for large message structures
- [ ] Cache compiled XPath expressions in util.go
- [ ] Use sync.Map for concurrent access to version mappings
- [ ] Optimize XML parsing with streaming where possible
- [ ] Add benchmarks before and after each optimization
- [ ] Profile memory allocations and reduce where possible
- [ ] Implement zero-copy techniques for byte slice handling

## Branch 7: Code Generation for Repetitive Patterns

### TODOs:
- [ ] Create code generation tool for version mapping constants
- [ ] Generate MessageWith/DocumentWith functions from templates
- [ ] Generate test boilerplate for new message types
- [ ] Create generator for field mapping functions
- [ ] Generate wrapper implementations from message definitions
- [ ] Add go:generate directives to relevant files
- [ ] Document code generation process in CONTRIBUTING.md
- [ ] Create CI checks to ensure generated code is up-to-date

## Branch 8: Documentation and Examples

### TODOs:
- [ ] Add package-level documentation for all packages
- [ ] Create detailed API documentation with godoc
- [ ] Add inline comments for complex algorithms
- [ ] Create tutorials for common use cases
- [ ] Document version compatibility matrix
- [ ] Add architecture decision records (ADRs)
- [ ] Create migration guide for breaking changes
- [ ] Add performance tuning guide

## Branch 9: Concurrency and Thread Safety

### TODOs:
- [ ] Audit all shared state for thread safety
- [ ] Add mutex protection where needed
- [ ] Document thread safety guarantees in API
- [ ] Create concurrent test scenarios
- [ ] Implement connection pooling for server
- [ ] Add context.Context support for cancellation
- [ ] Implement graceful shutdown for server
- [ ] Add rate limiting capabilities

## Branch 10: Observability and Debugging

### TODOs:
- [ ] Add structured logging with levels
- [ ] Implement metrics collection (processing time, error rates)
- [ ] Add tracing support with OpenTelemetry
- [ ] Create debug mode with verbose output
- [ ] Add validation result details for debugging
- [ ] Implement health check endpoints
- [ ] Add performance profiling endpoints
- [ ] Create diagnostic tools for message analysis

## Priority Order

1. **Critical**: ✅ Branches 1, 2 (Fix bugs and implement proper error handling) - **COMPLETED**
2. **High**: Branches 3, 4, 5 (Reduce duplication, improve tests, standardize APIs) - **NEXT PRIORITY**
3. **Medium**: Branches 6, 7, 8 (Performance, code generation, documentation)
4. **Low**: Branches 9, 10 (Advanced features)

## Success Metrics

### Completed ✅
- ✅ Zero critical bugs - All critical linting and compilation errors resolved
- ✅ Consistent error handling throughout - Idiomatic Go error patterns implemented
- ✅ "Never panic, always return errors" policy enforced across codebase
- ✅ Version-specific field mappings for XML schema compatibility
- ✅ All tests passing with proper CI pipeline validation

### In Progress / Outstanding
- Test coverage > 90% (currently at 51.0%, needs improvement)
- All public APIs documented
- No code duplication across message types
- Benchmark improvements of 20%+ for key operations
- Thread-safe operations verified
- Examples for all major use cases

### Key Achievements from Branch 2 Completion
- Implemented comprehensive error type system (ValidationError, ParseError, FieldError)
- Added proper error wrapping and unwrapping support
- Created error constructors and validation utilities
- Added exhaustive linter compliance
- Fixed all CI build failures
- Updated test assertions to match new error message field paths (FedwireFundsPaymentStatus, FedwireFundsSystemResponse)
- Resolved test failures caused by idiomatic error handling changes to validation field paths
- Established foundation for robust error handling across the library

## Immediate Next Steps (Priority Order)

### 🎯 **MOST IMPACTFUL: Branch 3 - Generic Message Processing Framework**
**Status**: Ready to begin - critical foundation completed
**Impact**: Eliminate ~2,400+ lines of duplicated code across 16+ message types

**Immediate Action Items (First 2 weeks)**:
1. **Create pkg/common/processor.go** - Generic MessageProcessor[T, V] using Go generics
2. **Proof of concept with CustomerCreditTransfer** - Migrate to generic processor  
3. **Measure impact** - Lines of code reduction, performance, maintainability
4. **Incremental migration plan** - One message type at a time

**Expected Outcomes**:
- Reduce maintenance burden from 16+ files to 1 shared implementation
- Eliminate ~1,500 lines of MessageWith/DocumentWith duplication
- Improve developer experience for adding new message types
- Establish pattern for remaining abstractions

### Branch 4: Improve Test Coverage and Quality  
**Status**: Can be worked on in parallel with Branch 3 processor work

**High Priority Items**:
1. Create test data factory functions for each message type
2. Add negative test cases for validation failures  
3. Add table-driven tests for version compatibility
4. Create integration tests for complete message flows

### Branch 5: Type-Safe Validation Framework
**Status**: Depends on Branch 3 processor completion

**High Priority Items**:
1. Replace reflection-based CheckRequiredFields with compile-time validation
2. Create Validatable interface with type-specific implementations
3. Add validation error aggregation with errors.Join()

### Outstanding Issues from Analysis
- [x] Documented XML to Go struct field mapping patterns (XML_TO_GO_MAPPING.md)
- [ ] Similar field mapping issues may exist in other message types (EndpointDetailsReport, etc.)
- [ ] Test coverage needs improvement (currently 51.0%, target >90%)
- [ ] Version string inconsistencies (lowercase vs uppercase) need standardization
- [ ] Factory pattern duplication across all message types (~600 lines)

## Notes for Implementation

- Each branch should have its own PR with focused changes
- Run benchmarks before and after performance branches
- Update CLAUDE.md after major architectural changes
- Consider semantic versioning after Branch 5 completion
- Add deprecation notices for any breaking changes
- Create automated migration tools where possible