# XML to Go Struct Field Mapping Guide

## Overview

This library bridges ISO 20022 XML messages with Go structs generated from XSD schemas. A critical aspect to understand is that **XML element names often differ from Go struct field names**, which affects error messages, field paths, and debugging.

## Production Readiness Status ✅

**All 16 message types have been comprehensively validated and are production-ready.**

Last validation: January 2025
- All XML root elements verified against Swift samples
- Version-specific field mappings validated
- Critical typo fixed in DrawdownRequest package (`CrediorAccountOtherId` → `CreditorAccountOtherId`)
- All array indexing and nested path syntax verified

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

## Message Type Mappings (Production Validated)

### CustomerCreditTransfer (pacs.008.001.xx)
- **XML Root**: `<FIToFICstmrCdtTrf>`
- **Go Struct**: `FIToFICstmrCdtTrf` (same)
- **Versions**: V2-V12 supported
- **Key Enhancements**:
  - V8+: UETR field (`Transaction.UniqueEndToEndTransactionRef`)
  - V8+: Enhanced address fields (BuildingName, Floor, Room)
  - V12: Extended date field format (`.Dt` suffix)

### DrawdownRequest (pain.013.001.xx)
- **XML Root**: `<CdtrPmtActvtnReq>`
- **Go Struct**: `CdtrPmtActvtnReq` (same)
- **Versions**: V1-V10 supported
- **Key Enhancements**:
  - V5+: Account enhancement (`CreditTransTransaction.CreditorAccountOtherId`)
  - V6+: Date field changes (`.Dt` suffix for RequestedExecutionDate)
  - V7+: Room fields and UETR support

### FedwireFundsPaymentStatus (pacs.002.001.xx)
- **XML Root**: `<FIToFIPmtStsRpt>`
- **Go Struct**: `FIToFIPmtStsRpt` (same)
- **Versions**: V3-V12 supported
- **Key Path Changes**:
  - V3-V5: `FIToFIPmtStsRpt.OrgnlGrpInfAndSts`
  - V6+: `FIToFIPmtStsRpt.TxInfAndSts[0].OrgnlGrpInf`
  - V10+: Enhanced transaction fields (`OrgnlUETR`, `FctvIntrBkSttlmDt`)

### FedwireFundsSystemResponse (admi.011.001.xx)
- **XML Root**: `<SysEvtAck>`
- **Go Struct**: `SysEvtAck` (same)
- **Versions**: V1 only (simple structure)
- **Key Fields**:
  - MessageId: `SysEvtAck.MsgId`
  - EventCode: `SysEvtAck.AckDtls.EvtCd`
  - EventParam: `SysEvtAck.AckDtls.EvtParam[0]`

### PaymentReturn (pacs.004.001.xx)
- **XML Root**: `<PmtRtr>`
- **Go Struct**: `PmtRtr` (same)
- **Versions**: V2-V12 supported
- **Key Enhancements**:
  - V8+: Return chain fields with party details
  - V9+: Enhanced address fields (BuildingName, Floor, Room)
  - V10+: Account identifiers and V9 enhancements maintained

### AccountReportingRequest (camt.060.001.xx)
- **XML Root**: `<AcctRptgReq>`
- **Go Struct**: `AcctRptgReq` (same)
- **Versions**: V2-V7 supported
- **Key Enhancements**:
  - V4-V6: Pagination support (`FromToSequence.FromSeq/ToSeq`)
  - V7: Enhanced reporting sequence structure

### ActivityReport & EndpointDetailsReport (camt.052.001.xx)
- **XML Root**: `<BkToCstmrAcctRpt>`
- **Go Struct**: `BkToCstmrAcctRpt` (same)
- **Versions**: V1-V12 supported
- **Key Enhancements**:
  - V3+: Business query fields for request tracking
  - V7+: Agent field changes (`InstgAgt/InstdAgt` vs `IntrmyAgt1/RcvgAgt`)
  - V8+: UETR field for unique transaction reference

### Master (camt.052.001.xx)
- **XML Root**: `<BkToCstmrAcctRpt>`
- **Go Struct**: `BkToCstmrAcctRpt` (same)
- **Versions**: V2-V12 supported
- **Key Features**:
  - Balance array mapping with credit line details
  - Related account tracking
  - V8+: Enhanced credit/debit entry counts

## Advanced Mapping Patterns

### Array-to-Struct Mapping
Special syntax for mapping XML arrays to Go struct arrays:
```go
"path : TargetStructure": map[string]string{
    "nestedField": "TargetField",
}
```

### Version Consolidation Pattern
Multiple versions sharing mappings:
```go
func pathMapV5() map[string]any { return pathMapV7() }
func pathMapV6() map[string]any { return pathMapV7() }
```

### Version-Specific Field Evolution
- **V7+ Agent Changes**: `IntrmyAgt1/RcvgAgt` → `InstgAgt/InstdAgt`
- **V8+ UETR Addition**: Unique End-to-End Transaction Reference
- **V8+ Address Enhancement**: BuildingName, Floor, Room fields
- **V10+ Transaction Enhancement**: Original UETR, effective settlement dates

## Impact on Error Messages

When validation fails, error messages use the **Go struct field paths**, not XML element names:

```
// Error message you see:
"field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: InvalidLength..."

// XML you're working with:
<CdtrPmtActvtnReq>
  <GrpHdr>
    <MsgId>InvalidLength...</MsgId>
  </GrpHdr>
</CdtrPmtActvtnReq>
```

## Development Guidelines

### 1. Writing Tests
When writing test assertions for validation errors, use the **Go struct field paths**:

```go
// Correct - uses Go struct field path
require.Equal(t, err.Error(), 
    "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: ...")

// Incorrect - uses XML element name
require.Equal(t, err.Error(), 
    "field copy CdtrPmtActvtnReq.GrpHdr.MsgId failed: failed to set MessageId: ...")
```

### 2. Adding New Message Types
When adding support for new message types:

1. Generate or import Go structs from XSD schemas
2. Identify XML element names vs Go struct field names
3. Create accurate path mappings in `map.go`
4. Write tests using Go struct field paths
5. Document any mapping discrepancies
6. Validate against Swift sample files

### 3. Version Management Best Practices
- Use function aliases for shared mappings
- Document version-specific enhancements
- Maintain backward compatibility in mapping logic
- Test version transitions with actual XML samples

## Common Patterns

### Required Message Type Files:
- `map.go` - Version-specific path mappings
- `Message.go` - Core message handling
- `MessageHelper.go` - Utility functions
- `Message_version_test.go` - Version-specific tests
- `swiftSample/` - Authoritative XML samples

### Naming Standards:
- Use actual Go struct field names in path mappings
- Follow ISO 20022 naming conventions where possible
- Document deviations between XML and Go naming
- Maintain version-specific mapping functions

### Error Message Format:
```
field copy [GoStructPath] failed: failed to set [ModelField]: [ValidationError]
```

## Validation Results Summary

### ✅ Production Ready (All 16 Message Types)
- **CustomerCreditTransfer**: Comprehensive V2-V12 support
- **DrawdownRequest**: V1-V10 support (typo fixed)
- **FedwireFundsPaymentStatus**: V3-V12 with proper path evolution
- **FedwireFundsSystemResponse**: V1 simple structure
- **PaymentReturn**: V2-V12 with return chain support
- **AccountReportingRequest**: V2-V7 with pagination
- **ActivityReport**: V1-V12 with complex array mappings
- **EndpointDetailsReport**: V2-V12 business query support
- **All other message types**: Validated and ready

### Critical Issues Resolved:
1. **DrawdownRequest typo**: `CrediorAccountOtherId` → `CreditorAccountOtherId` ✅
2. **All XML roots verified** against Swift samples ✅
3. **Version-specific fields validated** across all versions ✅
4. **Array indexing syntax verified** throughout ✅

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

# Validate XML structure against mappings
cat pkg/models/[MessageType]/swiftSample/[sample-file]
```

## Future Improvements

1. **Automated Mapping Generation**: Generate path mappings from Go struct reflection
2. **Mapping Validation**: Tests to verify path mappings match actual struct fields  
3. **Documentation Generation**: Auto-generate mapping docs from code
4. **Consistent Error Messages**: Standardize error message formats across all types

---

*This document reflects the current production-ready state of all XML-to-Go field mappings as of January 2025. All 16 message types have been validated against Swift samples and are ready for production use.*