# XML to Go Struct Field Mapping Guide

## Overview

This library bridges ISO 20022 XML messages with Go structs generated from XSD schemas. A critical aspect to understand is that **XML element names often differ from Go struct field names**, which affects error messages, field paths, and debugging.

## The Mapping Challenge

When working with this library, you'll encounter three different naming conventions:

### 1. XML Element Names (External View)
These are the standard ISO 20022 XML element names that appear in XML documents:
- `<CdtrPmtActvtnReq>` (Creditor Payment Activation Request)
- `<FIToFIPmtStsRpt>` (Financial Institution to Financial Institution Payment Status Report)
- `<AcctRptgReq>` (Account Reporting Request)

### 2. Go Struct Field Names (Internal Implementation)
These are the actual field names in the generated Go structs from the external `fedwire20022` package:
- `CstmrDrctDbtInitn` (Customer Direct Debit Initiation)
- `FIToFIPmtStsRpt` (same as XML in this case)
- `AcctRptgReq` (same as XML in this case)

### 3. Path Mappings (Error Context)
These combine the Go struct field names with nested field paths:
- `CstmrDrctDbtInitn.GrpHdr.MsgId`
- `FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId`
- `AcctRptgReq.RptgReq[0].ReqdMsgNmId`

## Impact on Error Messages

When validation fails, error messages use the **Go struct field paths**, not XML element names:

```
// Error message you see:
"field copy CstmrDrctDbtInitn.GrpHdr.MsgId failed: failed to set MessageId: InvalidLength..."

// XML you're working with:
<CdtrPmtActvtnReq>
  <GrpHdr>
    <MsgId>InvalidLength...</MsgId>
  </GrpHdr>
</CdtrPmtActvtnReq>
```

## Package-Specific Mappings

### DrawdownRequest (pain.013.001.xx)
- **XML Root**: `<CdtrPmtActvtnReq>`
- **Go Struct**: `CstmrDrctDbtInitn`
- **All Versions**: 01-10 use `CstmrDrctDbtInitn` field paths

### FedwireFundsPaymentStatus (pacs.002.001.xx)
- **XML Root**: `<FIToFIPmtStsRpt>`
- **Go Struct**: `FIToFIPmtStsRpt` (same)
- **Version-Specific Paths**:
  - Early versions: `FIToFIPmtStsRpt.GrpHdr.MsgId`
  - Later versions: `FIToFIPmtStsRpt.TxInfAndSts.OrgnlGrpInf.OrgnlMsgId`

### FedwireFundsSystemResponse (admi.011.001.xx)
- **XML Root**: `<SysEvtAck>`
- **Go Struct**: `SysEvtNtfctn`
- **Path**: `SysEvtNtfctn.EvtNtfctn.TxInfAndSts.OrgnlInstrId`

### AccountReportingRequest (camt.060.001.xx)
- **XML Root**: `<AcctRptgReq>`
- **Go Struct**: `AcctRptgReq` (same)
- **Path**: `AcctRptgReq.RptgReq[0].ReqdMsgNmId`

## Path Mapping Files

Each message type has a `map.go` file defining version-specific path mappings:

```go
func PathMapV1() map[string]string {
    return map[string]string{
        // Go struct field path -> MessageModel field name
        "CstmrDrctDbtInitn.GrpHdr.MsgId": "MessageId",
        "CstmrDrctDbtInitn.GrpHdr.CreDtTm": "CreateDatetime",
        // ...
    }
}
```

## Development Guidelines

### 1. Writing Tests
When writing test assertions for validation errors, use the **Go struct field paths**:

```go
// Correct - uses Go struct field path
require.Equal(t, err.Error(), 
    "field copy CstmrDrctDbtInitn.GrpHdr.MsgId failed: failed to set MessageId: ...")

// Incorrect - uses XML element name
require.Equal(t, err.Error(), 
    "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: ...")
```

### 2. Debugging Field Mapping Issues
If you see "field [FieldName] not found" errors:

1. Check the external package struct definition
2. Compare XML element names with Go field names
3. Update path mappings to use actual Go struct field names
4. Verify version-specific differences

### 3. Adding New Message Types
When adding support for new message types:

1. Generate or import Go structs from XSD schemas
2. Identify XML element names vs Go struct field names
3. Create accurate path mappings in `map.go`
4. Write tests using Go struct field paths
5. Document any mapping discrepancies

## Consistency Across Packages

To ensure consistency:

### Required Files per Message Type:
- `map.go` - Version-specific path mappings
- `Message.go` - Core message handling
- `MessageHelper.go` - Utility functions
- `Message_version_test.go` - Version-specific tests

### Naming Standards:
- Use actual Go struct field names in path mappings
- Follow ISO 20022 naming conventions where possible
- Document deviations between XML and Go naming
- Maintain version-specific mapping functions

### Error Message Format:
```
field copy [GoStructPath] failed: failed to set [ModelField]: [ValidationError]
```

## Common Pitfalls

1. **Assuming XML element names match Go field names** - Always verify actual struct definitions
2. **Not updating all version mappings** - Changes often affect multiple versions
3. **Mixing XML and Go naming in tests** - Use Go struct paths consistently
4. **Forgetting version-specific differences** - Path mappings can vary between versions

## Tools for Investigation

When debugging mapping issues:

```bash
# Find actual struct field names in external package
go list -m github.com/moov-io/fedwire20022
grep -r "type.*struct" $GOPATH/pkg/mod/github.com/moov-io/fedwire20022@*/

# Test specific message type
go test ./pkg/models/[MessageType]/... -v -run [TestName]

# Search for field path patterns
rg "field copy.*failed" pkg/models/
```

## Future Improvements

1. **Automated Mapping Generation**: Generate path mappings from Go struct reflection
2. **Mapping Validation**: Tests to verify path mappings match actual struct fields  
3. **Documentation Generation**: Auto-generate mapping docs from code
4. **Consistent Error Messages**: Standardize error message formats across all types

---

*This document should be updated whenever new message types are added or field mappings change.*