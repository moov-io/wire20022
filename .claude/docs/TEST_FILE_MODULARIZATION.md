# Test File Modularization Strategy

## Problem
Some test files are too large to fit in Claude's context window, particularly `*_version_test.go` files that contain 10+ version tests. Files like `ActivityReport/Message_version_test.go` (96K) exceed reasonable context limits.

## Current Structure Analysis

### Large Test Files (>30K):
- `ActivityReport/Message_version_test.go` (96K, 12 versions)
- `CustomerCreditTransfer/Message_version_test.go` (68K, 11 versions) 
- `EndpointDetailsReport/Message_version_test.go` (64K, 11 versions)
- `Master/Message_version_test.go` (60K, 11 versions)
- `PaymentReturn/Message_version_test.go` (55K, 12 versions)
- `DrawdownRequest/Message_version_test.go` (50K, 10 versions)
- `EndpointTotalsReport/Message_version_test.go` (37K, 11 versions)
- `ReturnRequestResponse/Message_version_test.go` (34K, 10 versions)
- `DrawdownResponse/Message_version_test.go` (33K, 10 versions)

### Root Cause
Each `*_version_test.go` file contains multiple `TestVersionXX` functions that:
1. Create a document with `DocumentWith()`
2. Validate the document
3. Marshal to XML
4. Parse back from XML
5. Assert on numerous fields (30-60 assertions per version)

## Proposed Modularization Strategy

### Option 1: Split by Version Groups
Break large files into smaller files based on version ranges:

```
pkg/models/ActivityReport/
├── Message_test.go                    # Core functionality tests
├── Message_versions_01_04_test.go    # Versions 01-04
├── Message_versions_05_08_test.go    # Versions 05-08  
├── Message_versions_09_12_test.go    # Versions 09-12
└── Message_version_helpers_test.go   # Shared test helpers
```

### Option 2: Split by Test Type
Separate different types of tests:

```
pkg/models/ActivityReport/
├── Message_test.go                   # Core functionality tests
├── Message_version_creation_test.go # Document creation tests
├── Message_version_parsing_test.go  # XML parsing tests
├── Message_version_fields_test.go   # Field validation tests
└── Message_version_helpers_test.go  # Shared test helpers
```

### Option 3: One File Per Version (Most Granular)
```
pkg/models/ActivityReport/
├── Message_test.go                 # Core functionality tests
├── Message_version_01_test.go     # Version 01 only
├── Message_version_02_test.go     # Version 02 only
├── ...
├── Message_version_12_test.go     # Version 12 only
└── Message_version_helpers_test.go # Shared helpers
```

## Recommended Approach: Option 1 (Version Groups)

**Benefits:**
- Reduces file sizes to manageable chunks (~20K each)
- Maintains logical grouping of related versions
- Easier to work with than dozens of tiny files
- Still allows targeted testing of version ranges

**Implementation:**
1. Create shared test helpers in `Message_version_helpers_test.go`
2. Split large files into 3-4 version groups
3. Maintain consistent naming: `Message_versions_XX_YY_test.go`

### Target File Sizes
- Keep each file under 25K bytes (~400 lines)
- Group 3-4 versions per file for message types with 10+ versions
- Keep 2-6 version message types in single files

### Shared Helpers Structure
```go
// Message_version_helpers_test.go
package MessageType

import (
    "testing"
    "github.com/stretchr/testify/require"
)

// Helper function to reduce duplication
func testVersionDocument(t *testing.T, version VERSION_TYPE, expectedFields map[string]interface{}) {
    // Common test logic for document creation, validation, and parsing
}

// Helper for field assertions
func assertCommonFields(t *testing.T, model MessageModel, expected map[string]interface{}) {
    // Common field assertions
}
```

## Implementation Priority

### Phase 1: Largest Files (>50K)
1. ActivityReport (96K, 12 versions) → 3 files
2. CustomerCreditTransfer (68K, 11 versions) → 3 files  
3. EndpointDetailsReport (64K, 11 versions) → 3 files
4. Master (60K, 11 versions) → 3 files
5. PaymentReturn (55K, 12 versions) → 3 files

### Phase 2: Medium Files (30-50K)
6. DrawdownRequest (50K, 10 versions) → 3 files
7. EndpointTotalsReport (37K, 11 versions) → 3 files
8. ReturnRequestResponse (34K, 10 versions) → 3 files
9. DrawdownResponse (33K, 10 versions) → 3 files

### Phase 3: Smaller Files (<30K)
Keep as-is, they're already manageable.

## Benefits of Modularization

1. **Claude Context**: All test files fit comfortably in context window
2. **Maintainability**: Easier to focus on specific version ranges
3. **Performance**: Faster test execution for targeted versions
4. **Debugging**: Easier to isolate issues to specific versions
5. **Collaboration**: Multiple developers can work on different version ranges

## Testing Considerations

- Ensure all existing test functionality is preserved
- Maintain code coverage at current levels
- Verify `make check` continues to pass
- Test helpers should reduce duplication, not increase complexity