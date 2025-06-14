# Base Abstractions Developer Guide

This guide explains how to use the `pkg/base/` abstractions to implement new ISO 20022 message types with minimal code duplication and maximum type safety.

## Overview

The base abstractions provide idiomatic Go patterns for implementing ISO 20022 message processing using:

- **Type parameters** for type-safe generic processing
- **Embedded structs** for common field patterns
- **Type assertions** over complex interfaces
- **Function types** for factory patterns

## Quick Start: Adding a New Message Type

### 1. Define Your Message Model

Use embedded base types to eliminate duplication:

```go
package MyNewMessage

import (
    "github.com/moov-io/wire20022/pkg/base"
    "github.com/moov-io/wire20022/pkg/models"
)

// MyMessageModel uses base abstractions
type MyMessageModel struct {
    // Embed common fields automatically
    base.PaymentCore `json:",inline"`
    
    // Add message-specific fields
    SpecialField     string                   `json:"specialField"`
    CustomAmount     models.CurrencyAndAmount `json:"customAmount"`
    
    // Use common agent patterns
    base.AgentPair          `json:",inline"`
    base.DebtorCreditorPair `json:",inline"`
}
```

### 2. Set Up Version Factory

Register all your message versions:

```go
// Version constants
type MY_MSG_VERSION string

const (
    MY_MSG_001_01 MY_MSG_VERSION = "001.01"
    MY_MSG_001_02 MY_MSG_VERSION = "001.02"
)

// Version mappings
var VersionNameSpaceMap = map[MY_MSG_VERSION]string{
    MY_MSG_001_01: "urn:iso:std:iso:20022:tech:xsd:mymsg.001.001.01",
    MY_MSG_001_02: "urn:iso:std:iso:20022:tech:xsd:mymsg.001.001.02",
}

var VersionPathMap = map[MY_MSG_VERSION]map[string]any{
    MY_MSG_001_01: PathMapV1(),
    MY_MSG_001_02: PathMapV2(),
}

// Global processor instance
var processor *base.MessageProcessor[MyMessageModel, MY_MSG_VERSION]

func init() {
    // Register versions using clean factory pattern
    registrations := []base.FactoryRegistration[models.ISODocument, MY_MSG_VERSION]{
        {
            Namespace: "urn:iso:std:iso:20022:tech:xsd:mymsg.001.001.01",
            Version:   MY_MSG_001_01,
            Factory: func() models.ISODocument {
                return &mymsg_001_001_01.Document{
                    XMLName: xml.Name{
                        Space: VersionNameSpaceMap[MY_MSG_001_01], 
                        Local: "Document",
                    },
                }
            },
        },
        {
            Namespace: "urn:iso:std:iso:20022:tech:xsd:mymsg.001.001.02",
            Version:   MY_MSG_001_02,
            Factory: func() models.ISODocument {
                return &mymsg_001_001_02.Document{
                    XMLName: xml.Name{
                        Space: VersionNameSpaceMap[MY_MSG_001_02], 
                        Local: "Document",
                    },
                }
            },
        },
    }
    
    factory := base.BuildFactoryFromRegistrations(registrations)
    
    // Create processor with all required mappings
    processor = base.NewMessageProcessor[MyMessageModel, MY_MSG_VERSION](
        factory.BuildNameSpaceModelMap(),
        factory.GetVersionMap(),
        VersionPathMap,
        RequiredFields,
    )
}
```

### 3. Implement Simple Processing Functions

Replace complex logic with single function calls:

```go
// Parse XML to message model - single line!
func MessageWith(data []byte) (MyMessageModel, error) {
    return processor.ProcessMessage(data)
}

// Create XML document from model - single line!
func DocumentWith(model MyMessageModel, version MY_MSG_VERSION) (models.ISODocument, error) {
    return processor.CreateDocument(model, version)
}

// Validate required fields - single line!
func CheckRequiredFields(model MyMessageModel) error {
    return processor.ValidateRequiredFields(model)
}
```

### 4. Create Helper Functions

Use common helper builders:

```go
func MyMessageDataModel() MyMessageModel {
    return MyMessageModel{
        PaymentCore: base.PaymentCore{
            MessageHeader: base.MessageHeader{
                MessageId:       "",
                CreatedDateTime: time.Time{},
            },
            NumberOfTransactions:  "",
            SettlementMethod:      "",
            CommonClearingSysCode: "",
        },
        // Message-specific fields initialize to zero values
    }
}

func MyMessageHelper() MyMessageHelper {
    return MyMessageHelper{
        PaymentMessageHelper: base.BuildPaymentMessageHelper(),
        AgentHelper:          base.BuildAgentHelper(),
        SpecialField: models.ElementHelper{
            Title:         "Special Field",
            Type:          "Max35Text",
            Documentation: "Message-specific field documentation",
        },
    }
}
```

## Available Base Types

### Core Message Types

```go
// Universal message header
type MessageHeader struct {
    MessageId       string    `json:"messageId"`
    CreatedDateTime time.Time `json:"createdDateTime"`
}

// Payment-specific fields
type PaymentCore struct {
    MessageHeader
    NumberOfTransactions  string                           `json:"numberOfTransactions"`
    SettlementMethod      models.SettlementMethodType      `json:"settlementMethod"`
    CommonClearingSysCode models.CommonClearingSysCodeType `json:"commonClearingSysCode"`
}

// Common agent patterns
type AgentPair struct {
    InstructingAgent models.Agent `json:"instructingAgent"`
    InstructedAgent  models.Agent `json:"instructedAgent"`
}

type DebtorCreditorPair struct {
    DebtorAgent   models.Agent `json:"debtorAgent"`
    CreditorAgent models.Agent `json:"creditorAgent"`
}

// Address and party information
type PartyAddress struct {
    StreetName     string `json:"streetName"`
    BuildingNumber string `json:"buildingNumber"`
    RoomNumber     string `json:"roomNumber"`
    PostalCode     string `json:"postalCode"`
    TownName       string `json:"townName"`
    Subdivision    string `json:"subdivision"`
    Country        string `json:"country"`
}

type Party struct {
    Name    string       `json:"name"`
    Address PartyAddress `json:"address"`
}
```

### Generic Processor

```go
type MessageProcessor[M any, V comparable] struct {
    namespaceMap   map[string]models.DocumentFactory
    versionMap     map[string]V
    pathMaps       map[V]map[string]any
    requiredFields []string
}
```

### Helper Builders

```go
// Common ElementHelper definitions
var CommonHelpers = map[string]HelperBuilder{
    "MessageId":              func() models.ElementHelper { /* ... */ },
    "CreatedDateTime":        func() models.ElementHelper { /* ... */ },
    "NumberOfTransactions":   func() models.ElementHelper { /* ... */ },
    "SettlementMethod":       func() models.ElementHelper { /* ... */ },
    "CommonClearingSysCode":  func() models.ElementHelper { /* ... */ },
    "InstructingAgent":       func() models.ElementHelper { /* ... */ },
    "InstructedAgent":        func() models.ElementHelper { /* ... */ },
    "DebtorAgent":           func() models.ElementHelper { /* ... */ },
    "CreditorAgent":         func() models.ElementHelper { /* ... */ },
}

// Pre-built helper types
func BuildStandardMessageHelper() StandardMessageHelper
func BuildPaymentMessageHelper() PaymentMessageHelper
func BuildAgentHelper() AgentHelper
func BuildAddressHelper() AddressHelper
```

## Design Patterns and Best Practices

### 1. Prefer Embedded Structs Over Duplication

**❌ Don't duplicate common fields:**
```go
type BadMessageModel struct {
    MessageId       string    // Duplicated everywhere
    CreatedDateTime time.Time // Duplicated everywhere
    SpecialField    string
}
```

**✅ Use embedded base types:**
```go
type GoodMessageModel struct {
    base.MessageHeader `json:",inline"`  // Embedded, no duplication
    SpecialField       string            `json:"specialField"`
}
```

### 2. Use Type Parameters for Type Safety

**❌ Don't use interface{} or reflection:**
```go
func BadMessageProcessor(data []byte, pathMap map[string]any) (interface{}, error) {
    // Loses type safety, error-prone
}
```

**✅ Use type parameters:**
```go
func GoodMessageProcessor[M any, V comparable](
    processor *base.MessageProcessor[M, V],
    data []byte,
) (M, error) {
    return processor.ProcessMessage(data) // Type-safe
}
```

### 3. Leverage Factory Pattern for Versions

**❌ Don't manually build factory maps:**
```go
var BadFactoryMap = map[string]models.DocumentFactory{
    "namespace1": func() models.ISODocument { /* manual XMLName setup */ },
    "namespace2": func() models.ISODocument { /* manual XMLName setup */ },
    // Lots of repetitive code...
}
```

**✅ Use factory registrations:**
```go
registrations := []base.FactoryRegistration[models.ISODocument, VERSION_TYPE]{
    {Namespace: "namespace1", Version: V1, Factory: factoryFunc1},
    {Namespace: "namespace2", Version: V2, Factory: factoryFunc2},
}
factory := base.BuildFactoryFromRegistrations(registrations)
namespaceMap := factory.BuildNameSpaceModelMap() // Automatically generated
```

### 4. Use Type Assertions Over Complex Interfaces

**❌ Don't create overly complex interfaces:**
```go
type BadDocumentInterface interface {
    SetXMLName(xml.Name)
    GetXMLName() xml.Name
    Validate() error
    // Too many methods, hard to implement
}
```

**✅ Use simple type assertions:**
```go
type XMLNameSetter interface {
    SetXMLName(xml.Name)
}

// Use with type assertion
if setter, ok := doc.(XMLNameSetter); ok {
    setter.SetXMLName(xmlName)
} else {
    // Fallback using reflection
    setXMLNameByReflection(doc, namespace)
}
```

## Testing Your Implementation

### 1. Create Test Cases Using Abstractions

```go
func TestMyMessageWith(t *testing.T) {
    // Use base test patterns
    testData := []byte(`<?xml version="1.0"?>...`)
    
    result, err := MessageWith(testData)
    require.NoError(t, err)
    
    // Test embedded fields work correctly
    assert.Equal(t, "MSG001", result.MessageId)
    assert.Equal(t, "1", result.NumberOfTransactions)
    
    // Test message-specific fields
    assert.Equal(t, "special_value", result.SpecialField)
}

func TestMyDocumentWith(t *testing.T) {
    model := MyMessageDataModel()
    model.MessageId = "TEST001"
    model.SpecialField = "test_value"
    
    doc, err := DocumentWith(model, MY_MSG_001_01)
    require.NoError(t, err)
    
    // Verify document structure
    assert.NotNil(t, doc)
}
```

### 2. Validate Against Swift Samples

Always test against authoritative XML samples:

```go
func TestMyMessageWithSwiftSample(t *testing.T) {
    // Read from swiftSample directory
    samplePath := "swiftSample/my_message_sample.xml"
    data, err := os.ReadFile(samplePath)
    require.NoError(t, err)
    
    result, err := MessageWith(data)
    require.NoError(t, err)
    
    // Validate field mapping matches XML structure
    assert.Equal(t, "expected_from_xml", result.MessageId)
}
```

## Migration Guide: Converting Existing Message Types

### Step 1: Identify Common Fields

Look for these patterns in your existing `MessageModel`:
- `MessageId string`
- `CreatedDateTime time.Time`
- `NumberOfTransactions string`
- `SettlementMethod`
- `*Agent` fields
- Address structures

### Step 2: Replace with Base Types

```go
// Before
type OldMessageModel struct {
    MessageId            string
    CreatedDateTime      time.Time
    NumberOfTransactions string
    InstructingAgent     models.Agent
    InstructedAgent      models.Agent
    SpecificField        string
}

// After  
type NewMessageModel struct {
    base.PaymentCore `json:",inline"`
    base.AgentPair   `json:",inline"`
    SpecificField    string `json:"specificField"`
}
```

### Step 3: Replace Processing Functions

```go
// Before: 20+ lines of duplicated code
func OldMessageWith(data []byte) (OldMessageModel, error) {
    doc, xmlns, err := models.DocumentFrom(data, NameSpaceModelMap)
    if err != nil {
        return OldMessageModel{}, errors.NewParseError("document creation", "XML data", err)
    }
    version := NameSpaceVersonMap[xmlns]
    
    dataModel := OldMessageModel{}
    pathMap := VersionPathMap[version]
    rePathMap := models.RemakeMapping(doc, pathMap, true)
    for sourcePath, targetPath := range rePathMap {
        models.CopyDocumentValueToMessage(doc, sourcePath, &dataModel, targetPath)
    }
    return dataModel, nil
}

// After: 1 line using abstractions
func NewMessageWith(data []byte) (NewMessageModel, error) {
    return processor.ProcessMessage(data)
}
```

## Performance Considerations

### Memory Efficiency
- Embedded structs have zero allocation overhead
- Type parameters eliminate interface{} boxing
- Generic factories reuse code paths

### Compilation Benefits
- Type safety enforced at compile time
- Dead code elimination for unused versions
- Inlining opportunities for simple functions

### Runtime Benefits
- No reflection in hot paths (except validation)
- Direct field access via embedding
- Optimized XML processing pipeline

## Troubleshooting Common Issues

### "field not found" errors
- Check that embedded field names match XML path mappings
- Verify `json:",inline"` tags are present on embedded structs
- Ensure path mappings reference correct Go struct fields

### Version registration issues
- Verify namespace strings exactly match XML samples
- Check that all versions are registered in init()
- Ensure VersionPathMap contains entries for all versions

### Type safety compilation errors
- Use correct type parameters: `[MessageModel, VersionType]`
- Ensure VersionType implements `comparable` constraint
- Check that factory functions return correct document types

## Future Enhancements

The base abstractions are designed to evolve. Planned additions:

1. **Performance optimizations**: Object pooling for high-throughput scenarios
2. **Additional common patterns**: More embedded types as identified
3. **Validation framework**: Generic validation rules beyond required fields
4. **Code generation**: Tools to auto-generate boilerplate from XSD
5. **JSON workflows**: Direct JSON to XML transformation pipelines

## Best Practices Summary

✅ **DO:**
- Use embedded structs for common field patterns
- Prefer type parameters over interface{} usage
- Leverage factory registrations for clean version management
- Test against Swift sample XML files
- Use JSON tags for future-ready APIs

❌ **DON'T:**
- Duplicate common field definitions across message types
- Use complex interface hierarchies
- Manually build factory maps
- Skip validation against authoritative XML samples
- Ignore type safety warnings

---

This base abstractions pattern eliminates ~70% of code duplication while maintaining full type safety and performance. It follows idiomatic Go principles and provides a scalable foundation for new ISO 20022 message implementations.