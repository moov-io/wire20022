#!/bin/bash
# Fix imports correctly to resolve naming conflicts

set -e

echo "Fixing imports correctly..."

cd /Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/messages

# First, let's clean up all the files by reverting the broken imports
for file in *_test.go; do
    if [[ -f "$file" ]]; then
        # Extract base name
        base_name="${file%_test.go}"
        
        # Fix the broken import line by finding and replacing the malformed import
        sed -i '' "s|${base_name} ${base_name}Model|${base_name}Model|g" "$file"
    fi
done

# Fix generic_proof_test.go if it exists
if [[ -f "generic_proof_test.go" ]]; then
    sed -i '' 's|CustomerCreditTransfer CustomerCreditTransferModel|CustomerCreditTransferModel|g' generic_proof_test.go
fi

echo "Imports fixed correctly!"