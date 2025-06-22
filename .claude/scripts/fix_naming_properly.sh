#!/bin/bash
# Fix naming conflicts properly by using import aliases

set -e

echo "Fixing naming conflicts with import aliases..."

cd /Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/messages

# For each message file, add an import alias
for file in *.go; do
    if [[ ! $file =~ _test\.go$ ]] && [[ ! $file == "processor.go" ]] && [[ ! $file == "api.go" ]]; then
        # Extract the base name without .go extension
        base_name="${file%.go}"
        
        # Change the import to use an alias with "model" suffix
        sed -i '' "s|\"github.com/moov-io/wire20022/pkg/models/${base_name}\"|${base_name}Model \"github.com/moov-io/wire20022/pkg/models/${base_name}\"|" "$file"
        
        # Update references to use the alias
        sed -i '' "s|${base_name}\.MessageModel|${base_name}Model.MessageModel|g" "$file"
        sed -i '' "s|${base_name}\.CAMT|${base_name}Model.CAMT|g" "$file"
        sed -i '' "s|${base_name}\.PACS|${base_name}Model.PACS|g" "$file"
        sed -i '' "s|${base_name}\.PAIN|${base_name}Model.PAIN|g" "$file"
        sed -i '' "s|${base_name}\.ADMI|${base_name}Model.ADMI|g" "$file"
        sed -i '' "s|${base_name}\.DocumentWith|${base_name}Model.DocumentWith|g" "$file"
        sed -i '' "s|${base_name}\.CheckRequiredFields|${base_name}Model.CheckRequiredFields|g" "$file"
        sed -i '' "s|${base_name}\.BuildMessageHelper|${base_name}Model.BuildMessageHelper|g" "$file"
        sed -i '' "s|${base_name}\.MessageWith|${base_name}Model.MessageWith|g" "$file"
    fi
done

echo "Import aliases added successfully!"