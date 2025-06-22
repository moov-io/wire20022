#!/usr/bin/env python3

import os
import re
import sys

# Message types with their latest versions
MESSAGE_TYPES = {
    "ConnectionCheck": "ADMI_001_001_01",
    "DrawdownRequest": "PAIN_013_001_10", 
    "DrawdownResponse": "PAIN_014_001_10",
    "EndpointDetailsReport": "CAMT_086_001_02",
    "EndpointGapReport": "CAMT_087_001_02", 
    "EndpointTotalsReport": "CAMT_089_001_02",
    "FedwireFundsAcknowledgement": "ADMI_004_001_02",
    "FedwireFundsPaymentStatus": "PACS_002_001_14",
    "FedwireFundsSystemResponse": "ADMI_010_001_01",
    "Master": "CAMT_052_001_12",
    "PaymentReturn": "PACS_004_001_12",
    "PaymentStatusRequest": "PACS_028_001_05",
    "ReturnRequestResponse": "CAMT_029_001_12"
}

def find_latest_version(message_type):
    """Find the latest version from version.go file"""
    version_file = f"pkg/models/{message_type}/version.go"
    
    if not os.path.exists(version_file):
        print(f"Warning: {version_file} not found")
        return None
        
    with open(version_file, 'r') as f:
        content = f.read()
    
    # Find all version constants
    version_pattern = r'([A-Z_]+_\d+_\d+_\d+)\s+[A-Z_]+_VERSION\s*='
    versions = re.findall(version_pattern, content)
    
    if not versions:
        print(f"Warning: No versions found in {version_file}")
        return None
        
    # Return the last version (assuming they're ordered)
    return versions[-1]

def find_version_type(message_type):
    """Find the version type from version.go file"""
    version_file = f"pkg/models/{message_type}/version.go"
    
    if not os.path.exists(version_file):
        return None
        
    with open(version_file, 'r') as f:
        content = f.read()
    
    # Find version type definition
    type_pattern = r'type\s+([A-Z_]+_VERSION)\s+string'
    match = re.search(type_pattern, content)
    
    if match:
        return match.group(1)
    
    return None

def add_imports_to_message(file_path):
    """Add fmt and io imports to Message.go"""
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Check if imports already exist
    if '"fmt"' in content and '"io"' in content:
        return content
    
    # Find import block and add missing imports
    import_pattern = r'import\s*\(\s*\n(\s*"[^"]+"\s*\n)*'
    
    def add_imports(match):
        imports = match.group(0)
        if '"fmt"' not in imports:
            imports = imports.rstrip() + '\n\t"fmt"\n'
        if '"io"' not in imports:
            imports = imports.rstrip() + '\n\t"io"\n'
        return imports
    
    content = re.sub(import_pattern, add_imports, content)
    
    with open(file_path, 'w') as f:
        f.write(content)
    
    return content

def add_xml_methods(message_type):
    """Add ReadXML and WriteXML methods to a message type"""
    message_file = f"pkg/models/{message_type}/Message.go"
    
    if not os.path.exists(message_file):
        print(f"Warning: {message_file} not found")
        return False
    
    # Read the file
    with open(message_file, 'r') as f:
        content = f.read()
    
    # Check if methods already exist
    if 'func (m *MessageModel) ReadXML' in content:
        print(f"{message_type}: Already has ReadXML method")
        return True
    
    # Add imports
    content = add_imports_to_message(message_file)
    
    # Find version type and latest version
    version_type = find_version_type(message_type)
    latest_version = find_latest_version(message_type)
    
    if not version_type or not latest_version:
        print(f"Warning: Could not determine version info for {message_type}")
        return False
    
    # Find where to insert methods (after struct definition, before var RequiredFields)
    pattern = r'(}\s*\n\s*var RequiredFields)'
    
    xml_methods = f'''}}

// ReadXML reads XML data from an io.Reader into the MessageModel
func (m *MessageModel) ReadXML(r io.Reader) error {{
\tdata, err := io.ReadAll(r)
\tif err != nil {{
\t\treturn fmt.Errorf("reading XML: %w", err)
\t}}
\t
\tmodel, err := processor.ProcessMessage(data)
\tif err != nil {{
\t\treturn err
\t}}
\t
\t*m = model
\treturn nil
}}

// WriteXML writes the MessageModel as XML to an io.Writer
// If no version is specified, uses the latest version ({latest_version})
func (m *MessageModel) WriteXML(w io.Writer, version ...{version_type}) error {{
\t// Default to latest version
\tver := {latest_version}
\tif len(version) > 0 {{
\t\tver = version[0]
\t}}
\t
\t// Create versioned document
\tdoc, err := DocumentWith(*m, ver)
\tif err != nil {{
\t\treturn fmt.Errorf("creating document: %w", err)
\t}}
\t
\t// Write XML with proper formatting
\tencoder := xml.NewEncoder(w)
\tencoder.Indent("", "  ")
\t
\t// Write XML declaration
\tif _, err := w.Write([]byte(xml.Header)); err != nil {{
\t\treturn fmt.Errorf("writing XML header: %w", err)
\t}}
\t
\t// Encode document
\tif err := encoder.Encode(doc); err != nil {{
\t\treturn fmt.Errorf("encoding XML: %w", err)
\t}}
\t
\treturn encoder.Flush()
}}

var RequiredFields'''
    
    # Replace the pattern
    if re.search(pattern, content):
        content = re.sub(pattern, xml_methods, content)
    else:
        print(f"Warning: Could not find insertion point in {message_file}")
        return False
    
    # Write the file back
    with open(message_file, 'w') as f:
        f.write(content)
    
    return True

def add_parse_xml_function(message_type):
    """Add ParseXML function and deprecate MessageWith"""
    message_file = f"pkg/models/{message_type}/Message.go"
    
    if not os.path.exists(message_file):
        return False
    
    with open(message_file, 'r') as f:
        content = f.read()
    
    # Check if ParseXML already exists
    if 'func ParseXML(' in content:
        print(f"{message_type}: Already has ParseXML function")
        return True
    
    # Find MessageWith function and replace it
    pattern = r'func MessageWith\(data \[\]byte\) \(MessageModel, error\) \{\s*return processor\.ProcessMessage\(data\)\s*\}'
    
    replacement = '''// ParseXML reads XML data into the MessageModel
// This replaces the non-idiomatic MessageWith function
func ParseXML(data []byte) (*MessageModel, error) {
\tmodel, err := processor.ProcessMessage(data)
\tif err != nil {
\t\treturn nil, err
\t}
\treturn &model, nil
}

// Deprecated: Use ParseXML instead
func MessageWith(data []byte) (MessageModel, error) {
\treturn processor.ProcessMessage(data)
}'''
    
    if re.search(pattern, content):
        content = re.sub(pattern, replacement, content)
        
        with open(message_file, 'w') as f:
            f.write(content)
        return True
    else:
        print(f"Warning: Could not find MessageWith function in {message_file}")
        return False

def process_message_type(message_type):
    """Process a single message type"""
    print(f"Processing {message_type}...")
    
    success1 = add_xml_methods(message_type)
    success2 = add_parse_xml_function(message_type)
    
    if success1 and success2:
        print(f"✅ {message_type} processed successfully")
        return True
    else:
        print(f"❌ {message_type} processing failed")
        return False

def main():
    """Main function"""
    print("Applying idiomatic XML API to all message types...")
    
    successful = []
    failed = []
    
    for message_type in MESSAGE_TYPES.keys():
        if process_message_type(message_type):
            successful.append(message_type)
        else:
            failed.append(message_type)
    
    print(f"\n✅ Successfully processed: {len(successful)}")
    for msg in successful:
        print(f"  - {msg}")
    
    if failed:
        print(f"\n❌ Failed to process: {len(failed)}")
        for msg in failed:
            print(f"  - {msg}")
    
    print(f"\nTotal: {len(successful)}/{len(MESSAGE_TYPES)} message types processed")

if __name__ == "__main__":
    main()