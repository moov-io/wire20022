#!/usr/bin/env python3

import os
import re

# Fix version mismatches for message types
FIXES = {
    "EndpointGapReport": ("CAMT_087_001_VERSION", "CAMT_087_001_02", "CAMT_052_001_VERSION", "CAMT_052_001_12"),
    "EndpointTotalsReport": ("CAMT_089_001_VERSION", "CAMT_089_001_02", "CAMT_052_001_VERSION", "CAMT_052_001_12"),
    "FedwireFundsAcknowledgement": ("ADMI_004_001_VERSION", "ADMI_004_001_02", "ADMI_004_001_VERSION", "ADMI_004_001_02"),
    "FedwireFundsSystemResponse": ("ADMI_010_001_VERSION", "ADMI_010_001_01", "ADMI_010_001_VERSION", "ADMI_010_001_01"),
}

def fix_version_mismatch(message_type, wrong_version_type, wrong_version, correct_version_type, correct_version):
    """Fix version mismatches in a message type"""
    file_path = f"pkg/models/{message_type}/Message.go"
    
    if not os.path.exists(file_path):
        print(f"Warning: {file_path} not found")
        return False
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Replace wrong version type and version with correct ones
    content = content.replace(wrong_version_type, correct_version_type)
    content = content.replace(wrong_version, correct_version)
    
    # Update the comment as well
    content = re.sub(
        rf'// If no version is specified, uses the latest version \({re.escape(wrong_version)}\)',
        f'// If no version is specified, uses the latest version ({correct_version})',
        content
    )
    
    with open(file_path, 'w') as f:
        f.write(content)
    
    return True

def main():
    """Main function"""
    print("Fixing version mismatches...")
    
    for message_type, (wrong_version_type, wrong_version, correct_version_type, correct_version) in FIXES.items():
        print(f"Fixing {message_type}...")
        fix_version_mismatch(message_type, wrong_version_type, wrong_version, correct_version_type, correct_version)
    
    print("Done! Testing compilation...")
    
    # Test compilation
    import subprocess
    for message_type in FIXES.keys():
        try:
            result = subprocess.run(['go', 'build', f'./pkg/models/{message_type}'], 
                                  capture_output=True, text=True)
            if result.returncode == 0:
                print(f"✅ {message_type}")
            else:
                print(f"❌ {message_type}: {result.stderr.strip()}")
        except Exception as e:
            print(f"❌ {message_type}: {e}")

if __name__ == "__main__":
    main()