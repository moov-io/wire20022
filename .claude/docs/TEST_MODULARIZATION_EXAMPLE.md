# Test File Modularization Example

## Current Problem
Large test files like `ActivityReport/Message_version_test.go` (96K, 1207 lines) are too big for Claude's context window.

## Solution Approach
Split large `*_version_test.go` files into smaller, focused files.

## Example: ActivityReport Modularization

### Before (1 large file):
```
pkg/models/ActivityReport/Message_version_test.go  (96K, 12 version tests)
```

### After (4 smaller files):
```
pkg/models/ActivityReport/
├── Message_versions_01_03_test.go     (~25K, 3 version tests)
├── Message_versions_04_06_test.go     (~25K, 3 version tests)  
├── Message_versions_07_09_test.go     (~25K, 3 version tests)
└── Message_versions_10_12_test.go     (~25K, 3 version tests)
```

## Manual Process (for immediate use)

### Step 1: Identify Function Boundaries
Each version test follows this pattern:
```go
func TestVersionXX(t *testing.T) {
    // ... test implementation
}
```

### Step 2: Extract Header
```go
package ActivityReport

import (
    "encoding/xml"
    "testing"

    "github.com/moov-io/wire20022/pkg/models"
    "github.com/stretchr/testify/require"
)
```

### Step 3: Create Grouped Files
Split by version ranges, keeping each file under 30K bytes.

## Benefits of This Approach

1. **Claude Context**: Each file fits comfortably in context window
2. **Focused Testing**: Can test specific version ranges  
3. **Parallel Development**: Multiple developers can work on different versions
4. **Faster Builds**: Selective test execution possible
5. **Easier Debugging**: Isolate issues to specific version ranges

## Naming Convention
- `Message_versions_XX_YY_test.go` - Version range XX to YY
- Group 3-4 versions per file
- Maintain original test function names

## Implementation Priority

### Immediate (>60K files):
1. ActivityReport (96K) → 4 files
2. CustomerCreditTransfer (68K) → 4 files
3. EndpointDetailsReport (64K) → 4 files
4. Master (60K) → 4 files

### Next (40-60K files):
5. PaymentReturn (55K) → 4 files
6. DrawdownRequest (50K) → 3 files

### Optional (30-40K files):
7. EndpointTotalsReport (37K) → 3 files
8. ReturnRequestResponse (34K) → 3 files
9. DrawdownResponse (33K) → 3 files

## Testing Verification
After modularization:
1. Run `make check` to ensure all tests pass
2. Verify coverage remains the same
3. Check that all version tests are still executed

## Example Implementation

See `.claude/scripts/modularize_test_files.py` for automation script that can perform this modularization automatically.

## Alternative: On-Demand Modularization

Instead of modularizing all files upfront, we can:
1. Keep original files intact
2. Modularize only when Claude needs to work on them
3. Use the script to split files temporarily during development
4. Re-combine if needed for CI/CD consistency