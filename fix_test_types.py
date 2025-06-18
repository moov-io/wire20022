#!/usr/bin/env python3

import os
import re
import glob

def fix_test_file(file_path):
    """Fix a specific test file's ParseXML usage"""
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    original_content = content
    
    # Fix the pattern: model, err := ParseXML(data) followed by model = *model
    content = re.sub(
        r'model, err := ParseXML\((\w+)\)\n\tif err != nil \{\n\t\tt\.Fatal\(err\)\n\t\}\n\tmodel = \*model',
        r'model, err := ParseXML(\1)\n\tif err != nil {\n\t\tt.Fatal(err)\n\t}',
        content,
        flags=re.MULTILINE
    )
    
    # Fix DocumentWith calls that expect value but get pointer
    content = re.sub(
        r'DocumentWith\(model, ',
        r'DocumentWith(*model, ',
        content
    )
    
    # Check if any changes were made
    if content != original_content:
        with open(file_path, 'w') as f:
            f.write(content)
        return True
    return False

def main():
    """Main function"""
    print("Fixing test file type issues...")
    
    # Find all test files in message directories
    test_files = []
    for message_dir in glob.glob("pkg/models/*/"):
        for test_file in glob.glob(f"{message_dir}*_test.go"):
            test_files.append(test_file)
    
    fixed_count = 0
    for test_file in test_files:
        if fix_test_file(test_file):
            print(f"Fixed {test_file}")
            fixed_count += 1
    
    print(f"Fixed {fixed_count} test files")
    
    # Test compilation
    print("Testing CustomerCreditTransfer compilation...")
    import subprocess
    try:
        result = subprocess.run(['go', 'test', './pkg/models/CustomerCreditTransfer', '-c'], 
                              capture_output=True, text=True)
        if result.returncode == 0:
            print("✅ CustomerCreditTransfer tests compile successfully!")
        else:
            print(f"❌ CustomerCreditTransfer compilation failed:")
            print(result.stderr[:500])
    except Exception as e:
        print(f"❌ Error testing compilation: {e}")

if __name__ == "__main__":
    main()