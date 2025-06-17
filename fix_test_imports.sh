#!/bin/bash
# Fix test file imports to avoid conflicts

set -e

echo "Fixing test file imports..."

cd /Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/messages

# Fix test files to use import aliases
for file in *_test.go; do
    if [[ -f "$file" ]] && [[ ! $file == "example_test.go" ]] && [[ ! $file == "generic_proof_test.go" ]]; then
        # Extract base name
        base_name="${file%_test.go}"
        
        # Add Model suffix to imports in test files
        sed -i '' "s|\"github.com/moov-io/wire20022/pkg/models/${base_name}\"|${base_name}Model \"github.com/moov-io/wire20022/pkg/models/${base_name}\"|" "$file"
        
        # Update all references in test files
        sed -i '' "s|${base_name}\.MessageModel|${base_name}Model.MessageModel|g" "$file"
        sed -i '' "s|${base_name}\.CAMT|${base_name}Model.CAMT|g" "$file"
        sed -i '' "s|${base_name}\.PACS|${base_name}Model.PACS|g" "$file"
        sed -i '' "s|${base_name}\.PAIN|${base_name}Model.PAIN|g" "$file"
        sed -i '' "s|${base_name}\.ADMI|${base_name}Model.ADMI|g" "$file"
        sed -i '' "s|${base_name}\.BuildMessageHelper|${base_name}Model.BuildMessageHelper|g" "$file"
        
        # Fix any version references
        sed -i '' "s|${base_name}_|${base_name}Model.|g" "$file"
    fi
done

# Special handling for generic_proof_test.go
if [[ -f "generic_proof_test.go" ]]; then
    sed -i '' 's|"github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"|CustomerCreditTransferModel "github.com/moov-io/wire20022/pkg/models/CustomerCreditTransfer"|' generic_proof_test.go
    sed -i '' 's|CustomerCreditTransfer\.|CustomerCreditTransferModel.|g' generic_proof_test.go
fi

echo "Test imports fixed!"