# Path Remapping Inconsistencies

This document outlines the path remapping inconsistencies found between different message types in the wire20022 library.

## Key Inconsistencies

### 1. Return Type Mismatch
- **CustomerCreditTransfer**: `PathMapV2()` returns `map[string]any`
- **PaymentReturn** and others: `PathMapV2()` returns `map[string]string`

This prevents PaymentReturn from supporting complex nested mappings.

### 2. Field Naming Conventions

#### Agent Fields
- CustomerCreditTransfer: `InstructingAgents` (plural)
- PaymentReturn: `InstructingAgent` (singular)

#### Address Fields
- CustomerCreditTransfer: Mixed use of `Address` and `PostalAddress`
- PaymentReturn: Consistent use of `Address`

#### Clearing System
- CustomerCreditTransfer: `CommonClearingSysCode`
- PaymentReturn: `ClearingSystem`

### 3. Typos and Abbreviations
- `RemittanceInfor` should be `RemittanceInformation`
- `TxId` → `TaxId` mapping only in CustomerCreditTransfer

## Recommendations

1. **Standardize return types**: Use `map[string]any` for all PathMap functions
2. **Consistent naming**: Adopt singular forms for all agent fields
3. **Address fields**: Standardize on either `Address` or `PostalAddress`
4. **Fix typos**: Correct all abbreviated or misspelled field names
5. **Document conventions**: Create a naming convention guide

## Impact

These inconsistencies don't cause runtime errors but make the API harder to use and maintain. Fixing them would be a breaking change but would improve long-term maintainability.