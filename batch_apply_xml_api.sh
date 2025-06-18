#!/bin/bash

# Batch apply idiomatic XML API to all remaining message types

MESSAGE_TYPES=(
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

# Function to add XML methods to a message type
add_xml_methods() {
    local message_type=$1
    local file_path="pkg/models/$message_type/Message.go"
    
    echo "Processing $message_type..."
    
    # Check if file exists
    if [ ! -f "$file_path" ]; then
        echo "Warning: $file_path not found"
        return 1
    fi
    
    # Check if ReadXML already exists
    if grep -q "func (m \*MessageModel) ReadXML" "$file_path"; then
        echo "$message_type already has ReadXML method"
        return 0
    fi
    
    # Add imports if not present
    if ! grep -q '"io"' "$file_path"; then
        sed -i.bak '/import (/a\
	"io"
' "$file_path"
    fi
    
    if ! grep -q '"fmt"' "$file_path"; then
        sed -i.bak '/import (/a\
	"fmt"
' "$file_path"
    fi
    
    # Find version type by looking for _VERSION type definition
    local version_type=$(grep -o '[A-Z_]*_VERSION' "$file_path" | head -1)
    if [ -z "$version_type" ]; then
        echo "Warning: Could not find version type for $message_type"
        return 1
    fi
    
    # For now, use placeholder for latest version - we'll fix these manually
    local latest_version="LATEST_VERSION_PLACEHOLDER"
    
    # Create backup
    cp "$file_path" "${file_path}.backup"
    
    # Add XML methods after struct definition, before var RequiredFields
    awk -v version_type="$version_type" -v latest_version="$latest_version" '
    /^}$/ && next_line_is_var {
        print $0
        print ""
        print "// ReadXML reads XML data from an io.Reader into the MessageModel"
        print "func (m *MessageModel) ReadXML(r io.Reader) error {"
        print "\tdata, err := io.ReadAll(r)"
        print "\tif err != nil {"
        print "\t\treturn fmt.Errorf(\"reading XML: %w\", err)"
        print "\t}"
        print "\t"
        print "\tmodel, err := processor.ProcessMessage(data)"
        print "\tif err != nil {"
        print "\t\treturn err"
        print "\t}"
        print "\t"
        print "\t*m = model"
        print "\treturn nil"
        print "}"
        print ""
        print "// WriteXML writes the MessageModel as XML to an io.Writer"
        print "// If no version is specified, uses the latest version"
        print "func (m *MessageModel) WriteXML(w io.Writer, version ..." version_type ") error {"
        print "\t// Default to latest version"
        print "\tver := " latest_version
        print "\tif len(version) > 0 {"
        print "\t\tver = version[0]"
        print "\t}"
        print "\t"
        print "\t// Create versioned document"
        print "\tdoc, err := DocumentWith(*m, ver)"
        print "\tif err != nil {"
        print "\t\treturn fmt.Errorf(\"creating document: %w\", err)"
        print "\t}"
        print "\t"
        print "\t// Write XML with proper formatting"
        print "\tencoder := xml.NewEncoder(w)"
        print "\tencoder.Indent(\"\", \"  \")"
        print "\t"
        print "\t// Write XML declaration"
        print "\tif _, err := w.Write([]byte(xml.Header)); err != nil {"
        print "\t\treturn fmt.Errorf(\"writing XML header: %w\", err)"
        print "\t}"
        print "\t"
        print "\t// Encode document"
        print "\tif err := encoder.Encode(doc); err != nil {"
        print "\t\treturn fmt.Errorf(\"encoding XML: %w\", err)"
        print "\t}"
        print "\t"
        print "\treturn encoder.Flush()"
        print "}"
        print ""
        next_line_is_var = 0
        next
    }
    /^var RequiredFields/ {
        next_line_is_var = 1
    }
    {
        if (next_line_is_var && /^}$/) {
            # This is the closing brace before var RequiredFields
        } else {
            next_line_is_var = 0
        }
        print $0
    }
    ' "$file_path" > "${file_path}.tmp" && mv "${file_path}.tmp" "$file_path"
    
    # Add ParseXML function and deprecate MessageWith
    sed -i.bak2 '/func MessageWith(data \[\]byte) (MessageModel, error) {/,/^}$/{
        N
        N
        s/func MessageWith(data \[\]byte) (MessageModel, error) {\n\treturn processor.ProcessMessage(data)\n}/\/\/ ParseXML reads XML data into the MessageModel\n\/\/ This replaces the non-idiomatic MessageWith function\nfunc ParseXML(data []byte) (*MessageModel, error) {\n\tmodel, err := processor.ProcessMessage(data)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n\treturn \&model, nil\n}\n\n\/\/ Deprecated: Use ParseXML instead\nfunc MessageWith(data []byte) (MessageModel, error) {\n\treturn processor.ProcessMessage(data)\n}/
    }' "$file_path"
    
    echo "✅ $message_type processed (manual version fix needed)"
    return 0
}

# Process all message types
successful=0
total=0

for message_type in "${MESSAGE_TYPES[@]}"; do
    total=$((total + 1))
    if add_xml_methods "$message_type"; then
        successful=$((successful + 1))
    fi
done

echo ""
echo "Results:"
echo "✅ Successfully processed: $successful/$total message types"
echo ""
echo "⚠️  Manual fixes needed:"
echo "1. Replace LATEST_VERSION_PLACEHOLDER with actual latest version constants"
echo "2. Verify import statements are correct"
echo "3. Test that each message type compiles"