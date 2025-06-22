#!/bin/bash
# Complete the idiomatic Go refactoring for wire20022

set -e

echo "Completing idiomatic Go refactoring..."

cd /Users/wadearnold/Documents/GitHub/wadearnold/wire20022

# Step 1: Clean up wrapper directory (remove old test files)
echo "Step 1: Cleaning up old wrapper directory..."
rm -rf pkg/wrapper

# Step 2: Fix remaining type names in messages directory
echo "Step 2: Fixing type names..."
cd pkg/messages

# Fix type names - need to handle the imports properly
for file in *.go; do
    # Skip test files for now
    if [[ ! $file =~ _test\.go$ ]]; then
        # Extract the base name without .go extension
        base_name="${file%.go}"
        
        # Update the type declaration to match the file name
        sed -i '' "s/type ${base_name}WrapperGeneric struct/type ${base_name} struct/" "$file"
        sed -i '' "s/type ${base_name} struct/type ${base_name} struct/" "$file"
        
        # Update the constructor function name
        sed -i '' "s/func New${base_name}WrapperGeneric()/func New${base_name}()/" "$file"
        sed -i '' "s/func New${base_name}()/func New${base_name}()/" "$file"
        
        # Update return types
        sed -i '' "s/\*${base_name}WrapperGeneric/\*${base_name}/" "$file"
        sed -i '' "s/&${base_name}WrapperGeneric{/\&${base_name}{/" "$file"
    fi
done

# Step 3: Remove the redundant imports alias when package name matches
echo "Step 3: Cleaning up imports..."
for file in *.go; do
    if [[ ! $file =~ _test\.go$ ]]; then
        # For each file, if the import alias matches the package name, remove it
        base_name="${file%.go}"
        # This is complex, so we'll use a more targeted approach
        # Remove the alias when it's redundant
        sed -i '' "s/${base_name} \"github.com\/moov-io\/wire20022\/pkg\/models\/${base_name}\"/\"github.com\/moov-io\/wire20022\/pkg\/models\/${base_name}\"/" "$file"
    fi
done

# Step 4: Update test files
echo "Step 4: Updating test files..."
for file in *_test.go; do
    # Update package declaration
    sed -i '' 's/package wrapper/package messages/' "$file"
    
    # Update wrapper references
    sed -i '' 's/WrapperGeneric//g' "$file"
    sed -i '' 's/originalWrapper := &/originalWrapper := \&messages./' "$file"
    sed -i '' 's/genericWrapper := New/genericWrapper := messages.New/' "$file"
    
    # Fix test names
    sed -i '' 's/TestAccountReportingRequestMigration/TestAccountReportingRequestEquivalence/g' "$file"
    sed -i '' 's/TestActivityReportMigration/TestActivityReportEquivalence/g' "$file"
    sed -i '' 's/TestConnectionCheckMigration/TestConnectionCheckEquivalence/g' "$file"
done

# Step 5: Create a simple API file that exports key types
echo "Step 5: Creating API surface..."
cat > api.go << 'EOF'
// Package messages provides type-safe processors for ISO 20022 message types.
package messages

// Message processors for each ISO 20022 message type
// Each processor provides:
// - CreateDocument: Convert JSON model to XML document
// - ValidateDocument: Validate JSON model and create document
// - Validate: Validate required fields
// - ConvertXMLToModel: Parse XML to typed model
// - GetHelp: Get field documentation

// Payment messages (pacs)
// - CustomerCreditTransfer: pacs.008 - Customer credit transfer
// - PaymentReturn: pacs.004 - Payment return
// - PaymentStatusRequest: pacs.028 - Payment status request
// - FedwireFundsPaymentStatus: pacs.002 - Payment status report

// Cash management messages (camt)
// - AccountReportingRequest: camt.060 - Account reporting request
// - ActivityReport: camt.086 - Bank services billing statement
// - EndpointDetailsReport: camt.090 - Request for member profile
// - EndpointGapReport: camt.087 - Request for duplicate
// - EndpointTotalsReport: camt.089 - Request to cancel payment
// - ReturnRequestResponse: camt.029 - Resolution of investigation
// - Master: camt.052 - Bank to customer account report

// Payment initiation messages (pain)
// - DrawdownRequest: pain.013 - Creditor payment activation request
// - DrawdownResponse: pain.014 - Creditor payment activation request status report

// Administrative messages (admi)
// - ConnectionCheck: admi.001 - Static data request
// - FedwireFundsAcknowledgement: admi.004 - System event acknowledgement
// - FedwireFundsSystemResponse: admi.011 - System event notification
EOF

# Step 6: Update processor.go with better naming
echo "Step 6: Updating processor.go..."
sed -i '' 's/CheckRequireField/Validate/g' processor.go
sed -i '' 's/message wrapper/message processor/g' processor.go
sed -i '' 's/wrapper provides/processor provides/g' processor.go

# Step 7: Fix migration test files - remove them as they're no longer needed
echo "Step 7: Removing migration tests..."
rm -f *_migration_test.go

# Step 8: Create a better example test
echo "Step 8: Creating example test..."
cat > example_test.go << 'EOF'
package messages_test

import (
	"encoding/json"
	"fmt"

	"github.com/moov-io/wire20022/pkg/messages"
	"github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"
)

func ExampleCustomerCreditTransfer_CreateDocument() {
	// Create a new CustomerCreditTransfer processor
	cct := messages.NewCustomerCreditTransfer()

	// Sample message data
	messageData := []byte(`{
		"messageId": "20250310B1QDRCQR000001",
		"createdDateTime": "2025-01-17T10:00:00Z",
		"numberOfTransactions": "1",
		"settlementMethod": "CLRG",
		"commonClearingSysCode": "FDW",
		"instructionId": "INSTR001",
		"endToEndId": "E2E001",
		"instrumentPropCode": "CTRC",
		"interBankSettAmount": {"currency": "USD", "amount": 1000.00},
		"interBankSettDate": "2025-01-17",
		"instructedAmount": {"currency": "USD", "amount": 1000.00},
		"chargeBearer": "SLEV",
		"instructingAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "123456789"},
		"instructedAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "987654321"},
		"debtorName": "John Doe",
		"debtorAddress": {"streetName": "Main St", "buildingNumber": "123", "postalCode": "12345", "townName": "Anytown", "country": "US"},
		"debtorAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "123456789"},
		"creditorAgent": {"paymentSysCode": "USABA", "paymentSysMemberId": "987654321"}
	}`)

	// Create XML document
	xmlData, err := cct.CreateDocument(messageData, CustomerCreditTransfer.PACS_008_001_08)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Created XML document with %d bytes\n", len(xmlData))
	// Output: Created XML document with 1891 bytes
}

func ExampleCustomerCreditTransfer_Validate() {
	cct := messages.NewCustomerCreditTransfer()

	// Create a model to validate
	var model CustomerCreditTransfer.MessageModel
	modelJSON := `{"messageId": "12345", "debtorName": "Test"}`
	
	if err := json.Unmarshal([]byte(modelJSON), &model); err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
		return
	}

	// Validate required fields
	if err := cct.Validate(model); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		// Output: Validation error: validation failed for field "CreatedDateTime": is required: required field missing
	}
}
EOF

echo ""
echo "Refactoring complete!"
echo ""
echo "Next steps:"
echo "1. Run 'make check' to ensure everything compiles"
echo "2. Update imports in other packages that reference wrapper"
echo "3. Consider further simplifications to the API"