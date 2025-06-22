#!/bin/bash
# Refactoring script to make wire20022 idiomatic Go before v1.0

set -e

echo "Starting idiomatic Go refactoring for wire20022..."

# Step 1: Move from wrapper to messages package
echo "Step 1: Moving from wrapper to messages package..."
cp -r pkg/wrapper/* pkg/messages/
cd pkg/messages

# Step 2: Update package declaration in all files
echo "Step 2: Updating package declarations..."
for file in *.go; do
    sed -i '' 's/package wrapper/package messages/' "$file"
done

# Step 3: Remove WrapperGeneric suffix from type names
echo "Step 3: Removing WrapperGeneric suffix from type names..."
for file in *.go; do
    # Update type names
    sed -i '' 's/WrapperGeneric//g' "$file"
    
    # Update function names
    sed -i '' 's/NewAccountReportingRequestWrapperGeneric/NewAccountReportingRequest/g' "$file"
    sed -i '' 's/NewActivityReportWrapperGeneric/NewActivityReport/g' "$file"
    sed -i '' 's/NewConnectionCheckWrapperGeneric/NewConnectionCheck/g' "$file"
    sed -i '' 's/NewCustomerCreditTransferWrapperGeneric/NewCustomerCreditTransfer/g' "$file"
    sed -i '' 's/NewDrawdownRequestWrapperGeneric/NewDrawdownRequest/g' "$file"
    sed -i '' 's/NewDrawdownResponseWrapperGeneric/NewDrawdownResponse/g' "$file"
    sed -i '' 's/NewEndpointDetailsReportWrapperGeneric/NewEndpointDetailsReport/g' "$file"
    sed -i '' 's/NewEndpointGapReportWrapperGeneric/NewEndpointGapReport/g' "$file"
    sed -i '' 's/NewEndpointTotalsReportWrapperGeneric/NewEndpointTotalsReport/g' "$file"
    sed -i '' 's/NewFedwireFundsAcknowledgementWrapperGeneric/NewFedwireFundsAcknowledgement/g' "$file"
    sed -i '' 's/NewFedwireFundsPaymentStatusWrapperGeneric/NewFedwireFundsPaymentStatus/g' "$file"
    sed -i '' 's/NewFedwireFundsSystemResponseWrapperGeneric/NewFedwireFundsSystemResponse/g' "$file"
    sed -i '' 's/NewMasterWrapperGeneric/NewMaster/g' "$file"
    sed -i '' 's/NewPaymentReturnWrapperGeneric/NewPaymentReturn/g' "$file"
    sed -i '' 's/NewPaymentStatusRequestWrapperGeneric/NewPaymentStatusRequest/g' "$file"
    sed -i '' 's/NewReturnRequestResponseWrapperGeneric/NewReturnRequestResponse/g' "$file"
done

# Step 4: Rename CheckRequireField to Validate throughout
echo "Step 4: Standardizing validation method names..."
for file in *.go; do
    sed -i '' 's/CheckRequireField/Validate/g' "$file"
done

# Step 5: Update comments to remove references to wrapper/generic
echo "Step 5: Updating comments..."
for file in *.go; do
    sed -i '' 's/generic wrapper implementation/message processor/g' "$file"
    sed -i '' 's/wrapper for/processor for/g' "$file"
    sed -i '' 's/The generic wrapper/The message processor/g' "$file"
done

# Step 6: Rename generic.go to processor.go for clarity
echo "Step 6: Renaming generic.go to processor.go..."
mv generic.go processor.go

# Step 7: Update import statements throughout the codebase
echo "Step 7: Updating imports throughout codebase..."
cd ../..
find . -name "*.go" -type f -exec sed -i '' 's|github.com/moov-io/wire20022/pkg/wrapper|github.com/moov-io/wire20022/pkg/messages|g' {} \;

# Step 8: Update test file references
echo "Step 8: Updating test files..."
cd pkg/messages
for file in *_test.go; do
    sed -i '' 's/WrapperGeneric//g' "$file"
    sed -i '' 's/Wrapper{}/messages.&/g' "$file"
    sed -i '' 's/wrapper\./messages./g' "$file"
done

echo "Refactoring complete!"
echo ""
echo "Summary of changes:"
echo "1. Moved pkg/wrapper → pkg/messages"
echo "2. Removed 'WrapperGeneric' suffix from all type names"
echo "3. Renamed validation method from CheckRequireField → Validate"
echo "4. Updated all imports and references"
echo "5. Made naming more idiomatic throughout"
echo ""
echo "Next steps:"
echo "1. Run 'make check' to ensure everything compiles"
echo "2. Update documentation to reflect new API"
echo "3. Consider additional refactoring for Helper types"