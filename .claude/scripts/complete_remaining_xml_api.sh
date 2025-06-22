#!/bin/bash

# Complete remaining XML API implementations
# This script applies the XML methods to the remaining message types

# Define message types with their latest versions
declare -A MESSAGE_VERSIONS=(
    ["DrawdownResponse"]="PAIN_014_001_10"
    ["EndpointDetailsReport"]="CAMT_086_001_02"
    ["EndpointGapReport"]="CAMT_087_001_02"
    ["EndpointTotalsReport"]="CAMT_089_001_02"
    ["FedwireFundsAcknowledgement"]="ADMI_004_001_02"
    ["FedwireFundsPaymentStatus"]="PACS_002_001_14"
    ["FedwireFundsSystemResponse"]="ADMI_010_001_01"
    ["PaymentStatusRequest"]="PACS_028_001_05"
    ["ReturnRequestResponse"]="CAMT_029_001_12"
)

# Define version types for each message type
declare -A VERSION_TYPES=(
    ["DrawdownResponse"]="PAIN_014_001_VERSION"
    ["EndpointDetailsReport"]="CAMT_086_001_VERSION"
    ["EndpointGapReport"]="CAMT_087_001_VERSION"
    ["EndpointTotalsReport"]="CAMT_089_001_VERSION"
    ["FedwireFundsAcknowledgement"]="ADMI_004_001_VERSION"
    ["FedwireFundsPaymentStatus"]="PACS_002_001_VERSION"
    ["FedwireFundsSystemResponse"]="ADMI_010_001_VERSION"
    ["PaymentStatusRequest"]="PACS_028_001_VERSION"
    ["ReturnRequestResponse"]="CAMT_029_001_VERSION"
)

add_xml_methods() {
    local message_type=$1
    local latest_version=${MESSAGE_VERSIONS[$message_type]}
    local version_type=${VERSION_TYPES[$message_type]}
    local file_path="pkg/models/$message_type/Message.go"
    
    echo "Processing $message_type..."
    
    if [ ! -f "$file_path" ]; then
        echo "Warning: $file_path not found"
        return 1
    fi
    
    # Check if ReadXML already exists
    if grep -q "func (m \*MessageModel) ReadXML" "$file_path"; then
        echo "$message_type already has ReadXML method"
        return 0
    fi
    
    # Create backup
    cp "$file_path" "${file_path}.backup"
    
    # Create temporary file with the XML methods
    cat > "/tmp/xml_methods_${message_type}.go" << EOF

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
// If no version is specified, uses the latest version ($latest_version)
func (m *MessageModel) WriteXML(w io.Writer, version ...$version_type) error {
	// Default to latest version
	ver := $latest_version
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
EOF
    
    # Insert the XML methods before "var RequiredFields"
    awk -v methods_file="/tmp/xml_methods_${message_type}.go" '
    /^var RequiredFields/ && !inserted {
        while ((getline line < methods_file) > 0) {
            print line
        }
        close(methods_file)
        print ""
        inserted = 1
    }
    { print }
    ' "$file_path" > "${file_path}.tmp" && mv "${file_path}.tmp" "$file_path"
    
    # Clean up
    rm "/tmp/xml_methods_${message_type}.go"
    
    # Test compilation
    if go build "./pkg/models/$message_type" 2>/dev/null; then
        echo "✅ $message_type completed successfully"
        return 0
    else
        echo "❌ $message_type compilation failed"
        # Restore backup
        mv "${file_path}.backup" "$file_path"
        return 1
    fi
}

# Process all remaining message types
successful=0
total=0

for message_type in "${!MESSAGE_VERSIONS[@]}"; do
    total=$((total + 1))
    if add_xml_methods "$message_type"; then
        successful=$((successful + 1))
    fi
done

echo ""
echo "Results:"
echo "✅ Successfully completed: $successful/$total message types"

# Test all message types compilation
echo ""
echo "Final compilation test for all message types:"
for dir in pkg/models/*/; do
    if [ -f "$dir/Message.go" ]; then
        package=$(basename "$dir")
        echo -n "Testing $package... "
        if go build "./$dir" 2>/dev/null; then
            echo "✅"
        else
            echo "❌"
        fi
    fi
done