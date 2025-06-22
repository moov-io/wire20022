#!/bin/bash
# Fix naming conflicts in messages package

set -e

echo "Fixing naming conflicts..."

cd /Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/messages

# We need to add a suffix to avoid conflicts with imported package names
# Let's use "Message" as the suffix

for file in *.go; do
    if [[ ! $file =~ _test\.go$ ]] && [[ ! $file == "processor.go" ]] && [[ ! $file == "api.go" ]]; then
        # Extract the base name without .go extension
        base_name="${file%.go}"
        
        # Add Message suffix to type name
        sed -i '' "s/type ${base_name} struct/type ${base_name}Message struct/" "$file"
        
        # Update constructor
        sed -i '' "s/func New${base_name}()/func New${base_name}Message()/" "$file"
        
        # Update return types
        sed -i '' "s/\*${base_name} {/\*${base_name}Message {/" "$file"
        sed -i '' "s/\*${base_name}$/\*${base_name}Message/" "$file"
        
        # Update comments
        sed -i '' "s/${base_name} demonstrates/${base_name}Message demonstrates/" "$file"
        sed -i '' "s/processor for ${base_name} messages/processor for ${base_name} messages/" "$file"
    fi
done

# Update the test files to use the new names
for file in *_test.go; do
    if [[ -f "$file" ]]; then
        # Extract base name
        base_name="${file%_test.go}"
        
        # Update references
        sed -i '' "s/New${base_name}()/New${base_name}Message()/" "$file"
        sed -i '' "s/messages\.${base_name}/messages.${base_name}Message/" "$file"
    fi
done

# Fix the example test
sed -i '' 's/messages\.NewCustomerCreditTransfer()/messages.NewCustomerCreditTransferMessage()/' example_test.go
sed -i '' 's/ExampleCustomerCreditTransfer_/ExampleCustomerCreditTransferMessage_/' example_test.go

echo "Naming conflicts fixed!"