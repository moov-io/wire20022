#!/usr/bin/env python3

import os
import re

# Mapping of message types to their latest versions
MESSAGE_VERSIONS = {
    "EndpointDetailsReport": ("CAMT_086_001_02", "CAMT_086_001_VERSION"),
    "EndpointGapReport": ("CAMT_087_001_02", "CAMT_087_001_VERSION"),
    "EndpointTotalsReport": ("CAMT_089_001_02", "CAMT_089_001_VERSION"),
    "FedwireFundsAcknowledgement": ("ADMI_004_001_02", "ADMI_004_001_VERSION"),
    "FedwireFundsPaymentStatus": ("PACS_002_001_14", "PACS_002_001_VERSION"),
    "FedwireFundsSystemResponse": ("ADMI_010_001_01", "ADMI_010_001_VERSION"),
    "PaymentStatusRequest": ("PACS_028_001_05", "PACS_028_001_VERSION"),
    "ReturnRequestResponse": ("CAMT_029_001_12", "CAMT_029_001_VERSION"),
}

def add_xml_methods(message_type, latest_version, version_type):
    """Add XML methods to a message type"""
    file_path = f"pkg/models/{message_type}/Message.go"
    
    if not os.path.exists(file_path):
        print(f"Warning: {file_path} not found")
        return False
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Check if methods already exist
    if 'func (m *MessageModel) ReadXML' in content:
        print(f"{message_type}: Already has ReadXML method")
        return True
    
    # Create the XML methods template
    xml_methods = f'''
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
    
    # Insert before "var RequiredFields"
    pattern = r'\n\nvar RequiredFields'
    if re.search(pattern, content):
        content = re.sub(pattern, xml_methods, content)
    else:
        print(f"Warning: Could not find insertion point in {file_path}")
        return False
    
    # Write back
    with open(file_path, 'w') as f:
        f.write(content)
    
    return True

def main():
    """Main function"""
    print("Applying XML methods to remaining message types...")
    
    successful = []
    failed = []
    
    for message_type, (latest_version, version_type) in MESSAGE_VERSIONS.items():
        print(f"Processing {message_type}...")
        if add_xml_methods(message_type, latest_version, version_type):
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
    
    print(f"\nTotal: {len(successful)}/{len(MESSAGE_VERSIONS)} message types processed")

if __name__ == "__main__":
    main()