# Implementation Guide: Adding New ISO 20022 Message Types

This guide provides step-by-step instructions for implementing new ISO 20022 message types in wire20022 using the standard base abstractions architecture.

## Overview

wire20022 uses a consistent architecture for all message types built on:
- **Base abstractions** (`pkg/base/`) for common functionality
- **Type-safe generics** for XML processing
- **Embedded structs** for field patterns
- **Factory registrations** for version management

## Prerequisites

Before implementing a new message type, ensure you have:
1. **ISO 20022 XSD schemas** for all versions you want to support
2. **Generated Go structs** from the XSD schemas (via code generation tools)
3. **Sample XML messages** for testing and validation
4. **Understanding of the XML structure** and field mappings

## Step-by-Step Implementation

### Step 1: Create Package Structure

Create a new directory under `pkg/models/` with your message type name:

```bash
mkdir pkg/models/YourMessageType
cd pkg/models/YourMessageType
```

Required files:
- `Message.go` - Core message structure and processing
- `MessageHelper.go` - Helper functions and documentation
- `map.go` - XML to Go struct field mappings
- `Message_test.go` - Comprehensive tests
- `Message_version_test.go` - Version-specific tests
- `swiftSample/` - Directory with sample XML files

### Step 2: Define Message Model (`Message.go`)

Create your message structure using base abstractions:

```go
package YourMessageType

import (
    "encoding/xml"
    "github.com/moov-io/fedwire20022/gen/YourMessageType/your_schema_001_01"
    "github.com/moov-io/fedwire20022/gen/YourMessageType/your_schema_001_02"
    // ... additional version imports
    "github.com/moov-io/wire20022/pkg/base"
    "github.com/moov-io/wire20022/pkg/models"
)

// Version-specific fields available in V2+ versions
type EnhancedFields struct {
    NewFieldInV2 string `json:"newFieldInV2"`
    AnotherField string `json:"anotherField"`
}

// Validate checks if enhanced fields meet requirements
func (e *EnhancedFields) Validate() error {
    if e.NewFieldInV2 == "" {
        return fmt.Errorf("NewFieldInV2 is required for versions V2+")
    }
    return nil
}

// NewMessageForVersion creates a MessageModel with appropriate version-specific fields initialized
func NewMessageForVersion(version YOUR_MESSAGE_VERSION) MessageModel {
    model := MessageModel{
        MessageHeader: base.MessageHeader{},
        // Core fields initialized to zero values
    }
    
    // Type-safe version-specific field initialization
    switch {
    case version >= YOUR_MSG_001_02:
        model.Enhanced = &EnhancedFields{}
    }
    
    return model
}

// ValidateForVersion performs type-safe validation for a specific version
func (m MessageModel) ValidateForVersion(version YOUR_MESSAGE_VERSION) error {
    // Base field validation (always required)
    if err := m.validateCoreFields(); err != nil {
        return fmt.Errorf("core field validation failed: %w", err)
    }
    
    // Type-safe version-specific validation
    switch {
    case version >= YOUR_MSG_001_02:
        if m.Enhanced == nil {
            return fmt.Errorf("EnhancedFields required for version %v but not present", version)
        }
        if err := m.Enhanced.Validate(); err != nil {
            return fmt.Errorf("EnhancedFields validation failed: %w", err)
        }
    }
    
    return nil
}

// validateCoreFields checks required core fields present in all versions
func (m MessageModel) validateCoreFields() error {
    // Direct field access - compile-time verified, no reflection
    if m.MessageId == "" {
        return fmt.Errorf("MessageId is required")
    }
    if m.CreatedDateTime.IsZero() {
        return fmt.Errorf("CreatedDateTime is required")
    }
    if m.YourSpecificField1 == "" {
        return fmt.Errorf("YourSpecificField1 is required")
    }
    return nil
}

// GetVersionCapabilities returns which version-specific features are available
func (m MessageModel) GetVersionCapabilities() map[string]bool {
    return map[string]bool{
        "Enhanced": m.Enhanced != nil,
    }
}

// MessageModel uses base abstractions to eliminate duplicate field definitions
type MessageModel struct {
    // Embed common message fields - choose appropriate base types
    base.MessageHeader `json:",inline"`
    // OR for payment messages:
    // base.PaymentCore `json:",inline"`
    
    // Core fields present in all versions (V1+)
    YourSpecificField1 string                   `json:"yourSpecificField1"`
    YourSpecificField2 models.CurrencyAndAmount `json:"yourSpecificField2"`
    
    // Use common agent patterns if applicable
    base.AgentPair          `json:",inline"`
    base.DebtorCreditorPair `json:",inline"`
    
    // Version-specific field groups (type-safe, nil when not applicable)
    Enhanced *EnhancedFields `json:",inline,omitempty"` // V2+ only
}

// Define required fields for validation
var RequiredFields = []string{
    "MessageId", "CreatedDateTime", "YourSpecificField1",
    // Add all required fields based on ISO 20022 specification
}

// Global processor instance using the base abstraction
var processor *base.MessageProcessor[MessageModel, YOUR_MESSAGE_VERSION]

// Version constants
type YOUR_MESSAGE_VERSION string

const (
    YOUR_MSG_001_01 YOUR_MESSAGE_VERSION = "001.01"
    YOUR_MSG_001_02 YOUR_MESSAGE_VERSION = "001.02"
    // Add all supported versions
)

// Version mappings
var VersionNameSpaceMap = map[YOUR_MESSAGE_VERSION]string{
    YOUR_MSG_001_01: "urn:iso:std:iso:20022:tech:xsd:yourmsg.001.001.01",
    YOUR_MSG_001_02: "urn:iso:std:iso:20022:tech:xsd:yourmsg.001.001.02",
    // Add all version namespaces
}

var VersionPathMap = map[YOUR_MESSAGE_VERSION]map[string]any{
    YOUR_MSG_001_01: PathMapV1(),
    YOUR_MSG_001_02: PathMapV2(),
    // Add all version path maps
}

// init sets up the processor using base abstractions
func init() {
    // Register all versions using factory registration pattern
    registrations := []base.FactoryRegistration[models.ISODocument, YOUR_MESSAGE_VERSION]{
        {
            Namespace: "urn:iso:std:iso:20022:tech:xsd:yourmsg.001.001.01",
            Version:   YOUR_MSG_001_01,
            Factory: func() models.ISODocument {
                return &your_schema_001_01.Document{
                    XMLName: xml.Name{
                        Space: VersionNameSpaceMap[YOUR_MSG_001_01], 
                        Local: "Document",
                    },
                }
            },
        },
        {
            Namespace: "urn:iso:std:iso:20022:tech:xsd:yourmsg.001.001.02",
            Version:   YOUR_MSG_001_02,
            Factory: func() models.ISODocument {
                return &your_schema_001_02.Document{
                    XMLName: xml.Name{
                        Space: VersionNameSpaceMap[YOUR_MSG_001_02], 
                        Local: "Document",
                    },
                }
            },
        },
        // Add all versions
    }

    versionedFactory := base.BuildFactoryFromRegistrations(registrations)

    // Create the processor using base abstractions
    processor = base.NewMessageProcessor[MessageModel, YOUR_MESSAGE_VERSION](
        versionedFactory.BuildNameSpaceModelMap(),
        versionedFactory.GetVersionMap(),
        VersionPathMap,
        RequiredFields,
    )
}

// MessageWith uses base abstractions to replace complex XML processing
func MessageWith(data []byte) (MessageModel, error) {
    return processor.ProcessMessage(data)
}

// DocumentWith uses base abstractions to replace complex document creation
func DocumentWith(model MessageModel, version YOUR_MESSAGE_VERSION) (models.ISODocument, error) {
    // Validate required fields before creating document
    if err := processor.ValidateRequiredFields(model); err != nil {
        return nil, err
    }
    return processor.CreateDocument(model, version)
}

// CheckRequiredFields uses base abstractions for validation
func CheckRequiredFields(model MessageModel) error {
    return processor.ValidateRequiredFields(model)
}
```

### Step 3: Create Field Mappings (`map.go`)

Define the XML to Go struct field mappings for each version:

```go
package YourMessageType

func PathMapV1() map[string]any {
    return map[string]any{
        // XML path -> Go struct field
        "YourRootElement.GrpHdr.MsgId":           "MessageId",
        "YourRootElement.GrpHdr.CreDtTm":         "CreatedDateTime", 
        "YourRootElement.SpecificElement":        "YourSpecificField1",
        "YourRootElement.Amount.Value":           "YourSpecificField2.Amount",
        "YourRootElement.Amount.Ccy":             "YourSpecificField2.Currency",
        
        // Agent mappings using base patterns
        "YourRootElement.InstgAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructingAgent.PaymentSysCode",
        "YourRootElement.InstgAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructingAgent.PaymentSysMemberId",
        "YourRootElement.InstdAgt.FinInstnId.ClrSysMmbId.ClrSysId.Cd": "InstructedAgent.PaymentSysCode",
        "YourRootElement.InstdAgt.FinInstnId.ClrSysMmbId.MmbId":       "InstructedAgent.PaymentSysMemberId",
        
        // Add complex mappings for nested structures
        "YourRootElement.ComplexField : YourComplexStructField": map[string]string{
            "SubField1": "NestedField1",
            "SubField2": "NestedField2",
        },
    }
}

func PathMapV2() map[string]any {
    return map[string]any{
        // Version 2 includes all V1 mappings plus new fields
        "YourRootElement.GrpHdr.MsgId":           "MessageId",
        "YourRootElement.GrpHdr.CreDtTm":         "CreatedDateTime", 
        "YourRootElement.SpecificElement":        "YourSpecificField1",
        "YourRootElement.Amount.Value":           "YourSpecificField2.Amount",
        "YourRootElement.Amount.Ccy":             "YourSpecificField2.Currency",
        
        // New fields in V2+ point to grouped field structure
        "YourRootElement.NewElement":             "Enhanced.NewFieldInV2",
        "YourRootElement.AnotherElement":         "Enhanced.AnotherField",
        // Add version-specific mappings
    }
}
```

**Critical**: Validate XML paths against actual sample XML files in `swiftSample/` directory.

### Step 4: Create Message Helpers (`MessageHelper.go`)

Provide helper functions and documentation:

```go
package YourMessageType

import (
    "time"
    "github.com/moov-io/wire20022/pkg/base"
    "github.com/moov-io/wire20022/pkg/models"
)

// YourMessageHelper provides field documentation and validation helpers
type YourMessageHelper struct {
    base.PaymentMessageHelper `json:",inline"`
    
    YourSpecificField1 models.ElementHelper `json:"yourSpecificField1"`
    YourSpecificField2 models.ElementHelper `json:"yourSpecificField2"`
}

// BuildMessageHelper creates a helper with field documentation
func BuildMessageHelper() YourMessageHelper {
    return YourMessageHelper{
        PaymentMessageHelper: base.BuildPaymentMessageHelper(),
        YourSpecificField1: models.ElementHelper{
            Title:         "Your Specific Field 1",
            Type:          "Max35Text (based on string) minLength: 1 maxLength: 35",
            Documentation: "Description of what this field represents in the message",
        },
        YourSpecificField2: models.ElementHelper{
            Title:         "Your Specific Field 2", 
            Type:          "CurrencyAndAmount",
            Documentation: "Description of the amount field",
        },
    }
}

// YourMessageDataModel creates a model with appropriate zero values
func YourMessageDataModel() MessageModel {
    return MessageModel{
        MessageHeader: base.MessageHeader{
            MessageId:       "",
            CreatedDateTime: time.Time{},
        },
        // OR for payment messages:
        // PaymentCore: base.PaymentCore{
        //     MessageHeader: base.MessageHeader{
        //         MessageId:       "",
        //         CreatedDateTime: time.Time{},
        //     },
        //     NumberOfTransactions:  "",
        //     SettlementMethod:      "",
        //     CommonClearingSysCode: "",
        // },
        YourSpecificField1: "",
        YourSpecificField2: models.CurrencyAndAmount{},
        AgentPair: base.AgentPair{
            InstructingAgent: models.Agent{},
            InstructedAgent:  models.Agent{},
        },
        DebtorCreditorPair: base.DebtorCreditorPair{
            DebtorAgent:   models.Agent{},
            CreditorAgent: models.Agent{},
        },
    }
}
```

### Step 5: Add Sample XML Files

Create `swiftSample/` directory with authoritative XML samples:

```bash
mkdir swiftSample
# Add sample XML files for each version:
# - YourMessage_01.xml
# - YourMessage_02.xml
# etc.
```

**Critical**: These XML files are the source of truth for field mappings and validation.

### Step 6: Write Comprehensive Tests

Create version-specific tests (`Message_version_test.go`):

```go
package YourMessageType

import (
    "encoding/xml"
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/moov-io/wire20022/pkg/models"
)

func TestVersion01(t *testing.T) {
    modelName := YOUR_MSG_001_01
    xmlName := "YourMessage_01.xml"
    dataModel := YourMessageDataModel()
    
    // Populate with test data
    dataModel.MessageId = "TEST001"
    dataModel.YourSpecificField1 = "test_value"
    
    // Create Document from Model
    doc01, err := DocumentWith(dataModel, modelName)
    require.NoError(t, err, "Failed to create document")
    
    // Validate document
    vErr := doc01.Validate()
    require.NoError(t, vErr, "Failed to validate document")
    
    // Create XML from document
    xmlData, err := xml.MarshalIndent(doc01, "", "  ")
    require.NoError(t, err)
    err = models.WriteXMLToGenerate(xmlName, xmlData)
    require.NoError(t, err)

    // Read back from XML
    xmlDoc, xmlErr := models.ReadXMLFile("./generated/" + xmlName)
    require.NoError(t, xmlErr, "Failed to read XML file")

    // Test round-trip conversion
    model, err := MessageWith(xmlDoc)
    require.NoError(t, err, "Failed to convert XML to model")
    require.Equal(t, "TEST001", model.MessageId, "Failed to get MessageId")
    require.Equal(t, "test_value", model.YourSpecificField1, "Failed to get YourSpecificField1")

    // Test validation errors with proper XML path format
    model.MessageId = "InvalidMessageIdThatIsTooLongForValidation12345"
    _, err = DocumentWith(model, modelName)
    require.NotNil(t, err, "Expected error but got nil")
    require.Equal(t, err.Error(), "field copy YourRootElement.GrpHdr.MsgId failed: failed to set MessageId: InvalidMessageIdThatIsTooLongForValidation12345 fails validation with length 45 <= required maxLength 35")
    model.MessageId = "TEST001"

    // Test required field validation
    model.MessageId = ""
    _, err = DocumentWith(model, modelName)
    require.NotNil(t, err, "Expected error but got nil")
    require.Equal(t, err.Error(), "validation failed for field \"MessageId\": is required: required field missing")
}
```

### Step 7: Validation and Testing

1. **Validate XML Paths**: Ensure all mappings in `map.go` match actual XML structure in samples
2. **Run Tests**: `go test -v ./pkg/models/YourMessageType`
3. **Test All Versions**: Verify each version works correctly
4. **Check Coverage**: Aim for >90% test coverage
5. **Validate Against Samples**: Test with real-world XML samples

### Step 8: Integration

1. **Add to Supported Types**: Update main documentation
2. **Add Wrapper** (if needed): Create wrapper in `pkg/wrapper/`
3. **Update Build**: Ensure `make check` passes
4. **Documentation**: Add usage examples

## Best Practices

### Field Mapping Guidelines

1. **Always validate against XML samples** - Never guess XML paths
2. **Use exact XML element names** in path mappings  
3. **Map to Go struct field names** (not XML names) for values
4. **Handle version differences** with separate PathMap functions
5. **Test error message paths** match XML structure

### Embedded Struct Usage

Choose appropriate base types:
- `base.MessageHeader` - For simple messages with just ID and timestamp
- `base.PaymentCore` - For payment messages with transaction info
- `base.AgentPair` - For messages with instructing/instructed agents
- `base.DebtorCreditorPair` - For messages with debtor/creditor agents

### Error Handling

- Use proper error message formats with XML paths
- Test validation error messages
- Include helpful context in error messages
- Follow Go error handling conventions

### Performance Considerations

- Embedded structs have zero allocation overhead
- Type parameters provide compile-time safety
- Generic processors optimize XML processing
- Factory pattern enables code reuse

## Common Issues and Solutions

### "field not found" errors
- Check embedded field names match path mappings
- Verify `json:",inline"` tags on embedded structs
- Ensure path mappings use correct Go struct field names

### Version registration issues  
- Verify namespace strings exactly match XML samples
- Check all versions are registered in init()
- Ensure VersionPathMap has entries for all versions

### Test assertion failures
- Update error message assertions to match XML path format
- Use actual XML samples for validation
- Run `make check` before committing

## Example: Complete Implementation

See existing message types like `CustomerCreditTransfer` or `PaymentReturn` for complete examples following this pattern.

## Step 9: Update Universal Reader

When adding a new message type, you **must** update the Universal Reader to support automatic detection of your new message type. This requires updates to several files:

### Universal Reader Updates Required

#### 1. Add Message Type Constant (`pkg/messages/universal_reader.go`)

Add your new message type to the constants section:

```go
const (
    TypeCustomerCreditTransfer      MessageType = "CustomerCreditTransfer"
    TypePaymentReturn               MessageType = "PaymentReturn"
    // ... existing types ...
    TypeYourNewMessageType          MessageType = "YourNewMessageType"  // <- Add this
    TypeUnknown                     MessageType = "Unknown"
)
```

#### 2. Add Import for Your Message Model

Add the import alias for your new message model:

```go
import (
    // ... existing imports ...
    YourNewMessageTypeModel "github.com/moov-io/wire20022/pkg/models/YourNewMessageType"
)
```

#### 3. Update Root Element Detection

Add your XML root element to the detection logic in `detectMessageType()`:

```go
func (r *UniversalReader) detectMessageType(peek *xmlPeeker, data []byte) (*DetectionInfo, error) {
    // ... existing logic ...
    
    // Map based on root element and namespace type
    switch peek.RootElement {
    case "FIToFICstmrCdtTrf":
        info.MessageType = TypeCustomerCreditTransfer
    case "PmtRtr":
        info.MessageType = TypePaymentReturn
    // ... existing cases ...
    case "YourXMLRootElement":  // <- Add your XML root element
        info.MessageType = TypeYourNewMessageType
    case "BkToCstmrAcctRpt":
        // This requires content analysis
        info.DetectedBy = "content_analysis"
        return r.analyzeBkToCstmrAcctRpt(info, data)
    case "Document":
        // ISO 20022 messages are often wrapped in a Document element
        return r.analyzeDocumentWrapper(info, data)
    default:
        info.MessageType = TypeUnknown
        info.DetectedBy = "failed"
        return info, fmt.Errorf("unknown root element: %s", peek.RootElement)
    }
    // ... rest of function
}
```

#### 4. Update Document Wrapper Analysis

If your message can be wrapped in a Document element, add the case to `analyzeDocumentWrapper()`:

```go
func (r *UniversalReader) analyzeDocumentWrapper(info *DetectionInfo, data []byte) (*DetectionInfo, error) {
    // ... existing logic ...
    
    // Map the child element to message type
    switch t.Name.Local {
    case "FIToFICstmrCdtTrf":
        childInfo.MessageType = TypeCustomerCreditTransfer
    // ... existing cases ...
    case "YourXMLRootElement":  // <- Add your XML root element
        childInfo.MessageType = TypeYourNewMessageType
    case "BkToCstmrAcctRpt":
        // This requires content analysis
        childInfo.DetectedBy = "content_analysis"
        return r.analyzeBkToCstmrAcctRpt(childInfo, data)
    default:
        childInfo.MessageType = TypeUnknown
        childInfo.DetectedBy = "failed"
        return childInfo, fmt.Errorf("unknown message element inside Document: %s", t.Name.Local)
    }
    // ... rest of function
}
```

#### 5. Add Parsing Logic

Add your message type to the parsing switch statement in `ReadBytes()`:

```go
func (r *UniversalReader) ReadBytes(data []byte) (*ParsedMessage, error) {
    // ... existing logic ...
    
    // Parse the actual message
    switch detection.MessageType {
    case TypeUnknown:
        return nil, fmt.Errorf("unsupported message type: %s", detection.MessageType)
    case TypeCustomerCreditTransfer:
        msg, err := CustomerCreditTransferModel.ParseXML(data)
        if err != nil {
            return nil, r.enhanceError(err, detection, data)
        }
        parsed.Message = msg
    // ... existing cases ...
    
    case TypeYourNewMessageType:  // <- Add your new case
        msg, err := YourNewMessageTypeModel.ParseXML(data)
        if err != nil {
            return nil, r.enhanceError(err, detection, data)
        }
        parsed.Message = msg
    
    default:
        return nil, fmt.Errorf("unsupported message type: %s", detection.MessageType)
    }
    // ... rest of function
}
```

#### 6. Update ValidateMessage Method

Add your message type to the validation type assertion:

```go
func (r *UniversalReader) ValidateMessage(parsed *ParsedMessage) error {
    // ... existing logic ...
    
    switch parsed.Message.(type) {
    case *CustomerCreditTransferModel.MessageModel,
         *PaymentReturnModel.MessageModel,
         // ... existing types ...
         *YourNewMessageTypeModel.MessageModel:  // <- Add your model type
        // Message is of a known type and was already validated during parsing
        return nil
    default:
        return fmt.Errorf("unknown message type for validation: %T", parsed.Message)
    }
}
```

### Testing Universal Reader Integration

#### 1. Add Test Cases (`pkg/messages/universal_reader_test.go`)

Add test cases for your new message type:

```go
func TestUniversalReader_WithSampleFiles(t *testing.T) {
    reader := NewUniversalReader()

    testCases := []struct {
        messageType MessageType
        samplePath  string
        skipFiles   []string
    }{
        // ... existing test cases ...
        {
            messageType: TypeYourNewMessageType,
            samplePath:  "../../pkg/models/YourNewMessageType/swiftSample",
        },
    }
    // ... rest of test
}
```

#### 2. Update Detection Tests

Add detection test cases for your XML structure:

```go
func TestUniversalReader_DetectMessageType(t *testing.T) {
    reader := NewUniversalReader()

    tests := []struct {
        name         string
        rootElement  string
        namespace    string
        expectedType MessageType
        expectedBy   string
    }{
        // ... existing tests ...
        {
            name:         "YourNewMessageType",
            rootElement:  "YourXMLRootElement",
            namespace:    "urn:iso:std:iso:20022:tech:xsd:yourmsg.001.001.01",
            expectedType: TypeYourNewMessageType,
            expectedBy:   "namespace",
        },
    }
    // ... rest of test
}
```

### Special Cases

#### BkToCstmrAcctRpt Messages

If your new message type uses the `BkToCstmrAcctRpt` root element (camt.052 variants), you need to update the content analysis logic in `analyzeBkToCstmrAcctRpt()`:

```go
func (r *UniversalReader) analyzeBkToCstmrAcctRpt(info *DetectionInfo, data []byte) (*DetectionInfo, error) {
    // ... existing analysis logic ...
    
    switch {
    case strings.HasPrefix(msgId, "ACTR"):
        info.MessageType = TypeActivityReport
    // ... existing cases ...
    case strings.HasPrefix(msgId, "YOUR"):  // <- Add your prefix pattern
        info.MessageType = TypeYourNewMessageType
    default:
        // Fallback to Rpt.Id analysis
        switch rptId {
        case "EDAY":
            info.MessageType = TypeActivityReport
        // ... existing cases ...
        case "YOUR":  // <- Add your Rpt.Id pattern
            info.MessageType = TypeYourNewMessageType
        default:
            info.MessageType = TypeUnknown
            return info, fmt.Errorf("unable to determine BkToCstmrAcctRpt subtype: MsgId=%s, Rpt.Id=%s", msgId, rptId)
        }
    }
    // ... rest of function
}
```

### Validation Requirements

Before committing your Universal Reader updates:

1. **Run Universal Reader Tests**: `go test -v ./pkg/messages -run TestUniversal`
2. **Test with Sample Files**: Ensure your swiftSample files are detected correctly
3. **Run Full Test Suite**: `make check` must pass
4. **Test wire20022 Tool**: Verify your new message type works with the validation tool

### Benefits of Universal Reader Integration

Once integrated, your new message type will:
- ✅ **Automatically detected** from XML structure without prior knowledge
- ✅ **Work with wire20022 CLI tool** for batch processing
- ✅ **Support enhanced error reporting** for debugging
- ✅ **Integrate with batch processing** workflows
- ✅ **Provide consistent API** with all other message types

## Summary

This implementation pattern:
- ✅ Eliminates code duplication through base abstractions
- ✅ Provides type safety with generics
- ✅ Ensures consistent error handling
- ✅ Simplifies testing and validation
- ✅ Maintains performance and idiomatic Go style
- ✅ **Integrates with Universal Reader** for automatic detection

Following this guide ensures your new message type integrates seamlessly with the wire20022 architecture, maintains consistency across the library, and works automatically with the Universal Reader for unknown message type handling.