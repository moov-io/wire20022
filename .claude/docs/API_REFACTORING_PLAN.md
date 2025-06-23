# API Refactoring Plan: Reducing Public Surface Area

## Current State
- **302 total exported functions** - Far too many for a library
- Many internal implementation details are publicly exposed
- Users have access to low-level mapping and helper functions they shouldn't need

## Target State
- **~50-60 exported functions** - Core public API only
- Internal implementation details properly encapsulated
- Clean, focused API that guides users to correct usage patterns

## Categorization and Actions

### 🔄 **Keep Exported - Core Public API (52 functions)**

#### Message Constructors (16 functions) ✅ KEEP
```go
func NewCustomerCreditTransfer() *CustomerCreditTransfer
func NewPaymentReturn() *PaymentReturn 
func NewActivityReport() *ActivityReport
// ... 13 more message constructors
```

#### XML Processing API (32 functions) ✅ KEEP  
```go
func ParseXML(data []byte) (*MessageModel, error)      // 16 functions
func DocumentWith(model, version) (ISODocument, error) // 16 functions
```

#### Essential Utilities (4 functions) ✅ KEEP
```go
func DocumentFrom(data []byte, factoryMap map[string]DocumentFactory) (ISODocument, string, error)
func ReadXMLFile(filename string) ([]byte, error)
func WriteXMLTo(filePath string, data []byte) error  
func IsEmpty(value interface{}) bool
```

---

### 📦 **Make Unexported - Internal Implementation (250+ functions)**

#### 1. Path Mapping Functions (180+ functions) ❌ MAKE PRIVATE
```go
// Current: PathMapV1() through PathMapV14() across all message types
// New: pathMapV1() through pathMapV14() 

func PathMapV10() map[string]any  →  func pathMapV10() map[string]any
```
**Reason**: These are internal XML-to-Go field mappings used only by processors

#### 2. Helper Builder Functions (42 functions) ❌ MAKE PRIVATE  
```go
// Current: BuildMessageHelper(), BuildPartyHelper(), etc.
// New: buildMessageHelper(), buildPartyHelper(), etc.

func BuildMessageHelper() MessageHelper  →  func buildMessageHelper() MessageHelper
```
**Reason**: These provide schema documentation but aren't core API. Used internally for validation.

#### 3. Internal Utilities (15+ functions) ❌ MAKE PRIVATE
```go
func SetElementToDocument(item any, path string, value any) error → 
func setElementToDocument(item any, path string, value any) error

func CopyDocumentValueToMessage(from any, fromPah string, to any, toPath string) →
func copyDocumentValueToMessage(from any, fromPah string, to any, toPath string) 

func RemakeMapping(from any, modelMap map[string]any, toModel bool) map[string]string →
func remakeMapping(from any, modelMap map[string]any, toModel bool) map[string]string
```
**Reason**: Internal field manipulation and transformation logic

#### 4. Error Handlers (3 functions) ❌ MAKE PRIVATE
```go
func HandleDocumentCreationError(err error) error → 
func handleDocumentCreationError(err error) error
```
**Reason**: Internal error handling used by processors

#### 5. Factory Functions (4+ functions) ❌ MAKE PRIVATE
```go
func NewMessageProcessor[M any, V comparable](...) →
func newMessageProcessor[M any, V comparable](...)
```
**Reason**: Internal processor construction

---

### 🤔 **Consider Carefully - Error Package**

#### Error Constructors (13 functions) - DECISION NEEDED
```go
func NewValidationError(field, reason string) *ValidationError
func NewParseError(operation, content string, cause error) *ParseError
func NewFieldError(path, operation string, cause error) *FieldError
// ... 10 more error constructors
```

**Option A**: Keep exported - Users may need to handle specific error types
**Option B**: Make unexported - Force users to use error interface checking instead

**Recommendation**: Keep the most common ones exported:
- `NewValidationError` 
- `NewParseError`
- `NewFieldError`

---

## Implementation Plan

### Phase 1: Path Mapping Functions (Immediate)
```bash
# Rename all PathMapVX functions to pathMapVX
find pkg/models -name "map.go" -exec sed -i 's/func PathMap/func pathMap/g' {} \;
```

### Phase 2: Helper Builders (Quick)
```bash  
# Rename all BuildXxxHelper functions to buildXxxHelper
find pkg/models -name "*Helper.go" -exec sed -i 's/func Build/func build/g' {} \;
```

### Phase 3: Internal Utilities (Careful)
- Review each function's usage
- Rename if only used internally
- Keep exported if used by external packages

### Phase 4: Update References
- Update all internal calls to use new lowercase names
- Ensure no external packages break

## Benefits After Refactoring

### Before: 302 exported functions 😰
```
❌ PathMapV1(), PathMapV2(), ... PathMapV14() (×16 message types)
❌ BuildMessageHelper(), BuildPartyHelper(), etc. (×42 functions)  
❌ SetElementToDocument(), CopyDocumentValueToMessage(), etc.
✅ ParseXML(), DocumentWith(), NewCustomerCreditTransfer(), etc.
```

### After: ~50-60 exported functions 🎉
```
✅ Core XML processing: ParseXML(), DocumentWith()
✅ Message constructors: NewCustomerCreditTransfer(), etc.
✅ Essential utilities: DocumentFrom(), ReadXMLFile(), etc.
✅ Key error types: NewValidationError(), etc.
🔒 Internal implementation hidden and properly encapsulated
```

## Documentation Impact
- Remove documentation from newly-private functions 
- Focus documentation efforts on the ~50 core public functions
- Users get a clean, focused API surface that guides correct usage

## Breaking Change Assessment
- This is a **major breaking change** for any external users
- Most "users" of internal functions are likely doing something wrong
- Benefits outweigh the breaking change costs for long-term API design