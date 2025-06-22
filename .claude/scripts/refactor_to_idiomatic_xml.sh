#!/bin/bash

# Script to refactor all message types to use idiomatic XML API

# List of all message types
MESSAGE_TYPES=(
    "AccountReportingRequest"
    "ActivityReport"
    "ConnectionCheck"
    "DrawdownRequest"
    "DrawdownResponse"
    "EndpointDetailsReport"
    "EndpointGapReport"
    "EndpointTotalsReport"
    "FedwireFundsAcknowledgement"
    "FedwireFundsPaymentStatus"
    "FedwireFundsSystemResponse"
    "Master"
    "PaymentReturn"
    "PaymentStatusRequest"
    "ReturnRequestResponse"
)

# Function to add idiomatic XML methods to a message type
add_idiomatic_methods() {
    local message_type=$1
    local file_path="pkg/models/$message_type/Message.go"
    
    echo "Processing $message_type..."
    
    # Check if file exists
    if [ ! -f "$file_path" ]; then
        echo "Warning: $file_path not found"
        return
    fi
    
    # Check if ReadXML already exists
    if grep -q "func (m \*MessageModel) ReadXML" "$file_path"; then
        echo "$message_type already has ReadXML method"
        return
    fi
    
    # Add the import for io and fmt if not present
    if ! grep -q "import.*io" "$file_path"; then
        sed -i '' '/import (/a\
	"fmt"\
	"io"
' "$file_path"
    fi
    
    # Find the latest version constant (assuming pattern like CAMT_060_001_XX or PACS_XXX_XXX_XX)
    local version_prefix=$(grep -o '[A-Z]*_[0-9]*_[0-9]*_[0-9]*' "$file_path" | head -1 | sed 's/_[0-9]*$//')
    local latest_version=$(grep -o "${version_prefix}_[0-9]*" "$file_path" | sort -V | tail -1)
    
    # Add ReadXML and WriteXML methods after the struct definition
    # This is a placeholder - you'll need to add the actual implementation
    echo "Adding ReadXML and WriteXML methods to $message_type with latest version: $latest_version"
    
    # Create a backup
    cp "$file_path" "${file_path}.bak"
    
    # Add the methods (simplified version - manual editing may be needed)
    cat >> "$file_path" << 'EOF'

// ReadXML reads XML data from an io.Reader into the MessageModel
func (m *MessageModel) ReadXML(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("reading XML: %w", err)
	}
	
	model, err := processor.ProcessMessage(data)
	if err != nil {
		return err
	}
	
	*m = model
	return nil
}

// WriteXML writes the MessageModel as XML to an io.Writer
// If no version is specified, uses the latest version
func (m *MessageModel) WriteXML(w io.Writer, version ...VERSION_TYPE) error {
	// Default to latest version
	ver := LATEST_VERSION
	if len(version) > 0 {
		ver = version[0]
	}
	
	// Create versioned document
	doc, err := DocumentWith(*m, ver)
	if err != nil {
		return fmt.Errorf("creating document: %w", err)
	}
	
	// Write XML with proper formatting
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")
	
	// Write XML declaration
	if _, err := w.Write([]byte(xml.Header)); err != nil {
		return fmt.Errorf("writing XML header: %w", err)
	}
	
	// Encode document
	if err := encoder.Encode(doc); err != nil {
		return fmt.Errorf("encoding XML: %w", err)
	}
	
	return encoder.Flush()
}

// ParseXML reads XML data into the MessageModel
// This replaces the non-idiomatic MessageWith function
func ParseXML(data []byte) (*MessageModel, error) {
	model, err := processor.ProcessMessage(data)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
EOF
    
    echo "Added methods to $message_type - manual review required"
}

# Process CustomerCreditTransfer is already done, skip it
echo "Skipping CustomerCreditTransfer (already completed)"

# Process all other message types
for message_type in "${MESSAGE_TYPES[@]}"; do
    add_idiomatic_methods "$message_type"
done

echo ""
echo "Refactoring complete!"
echo "IMPORTANT: Manual review and adjustments are required for:"
echo "1. Version type names (VERSION_TYPE placeholder)"
echo "2. Latest version constant (LATEST_VERSION placeholder)"
echo "3. Import statements"
echo "4. Test file updates"