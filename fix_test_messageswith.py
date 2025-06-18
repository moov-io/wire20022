#!/usr/bin/env python3

import os
import re
import glob

def fix_test_files():
    """Fix MessageWith references in test files"""
    
    # Find all test files that reference MessageWith
    test_files = []
    
    # Search in pkg/models directories
    for message_dir in glob.glob("pkg/models/*/"):
        for test_file in glob.glob(f"{message_dir}*_test.go"):
            with open(test_file, 'r') as f:
                content = f.read()
                if "MessageWith" in content:
                    test_files.append(test_file)
    
    print(f"Found {len(test_files)} test files with MessageWith references")
    
    for test_file in test_files:
        print(f"Fixing {test_file}...")
        
        with open(test_file, 'r') as f:
            content = f.read()
        
        # Replace MessageWith(data) with ParseXML pattern
        # Pattern: variable, err := MessageWith(data)
        content = re.sub(
            r'(\w+), err := MessageWith\((\w+)\)',
            r'\1, err := ParseXML(\2)\n\tif err != nil {\n\t\tt.Fatal(err)\n\t}\n\t\1 = *\1',
            content
        )
        
        # Pattern: MessageWith(data) without assignment
        content = re.sub(
            r'MessageWith\((\w+)\)',
            r'func() { msg, err := ParseXML(\1); if err != nil { t.Fatal(err) } }()',
            content
        )
        
        with open(test_file, 'w') as f:
            f.write(content)
    
    print("Done fixing test files!")

def main():
    """Main function"""
    print("Fixing MessageWith references in test files...")
    fix_test_files()
    
    print("Testing specific message type compilation...")
    
    # Test a few message types to see if they compile
    import subprocess
    test_types = ["CustomerCreditTransfer", "PaymentReturn", "AccountReportingRequest"]
    
    for msg_type in test_types:
        try:
            result = subprocess.run(['go', 'build', f'./pkg/models/{msg_type}'], 
                                  capture_output=True, text=True)
            if result.returncode == 0:
                print(f"✅ {msg_type}")
            else:
                print(f"❌ {msg_type}: {result.stderr.strip()}")
        except Exception as e:
            print(f"❌ {msg_type}: {e}")

if __name__ == "__main__":
    main()