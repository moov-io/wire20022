#!/usr/bin/env python3

import os
import re

def fix_encoder_close(file_path):
    """Fix XML encoder by adding defer encoder.Close() after xml.NewEncoder"""
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Pattern: encoder := xml.NewEncoder(w)\n\tencoder.Indent
    # Replace with: encoder := xml.NewEncoder(w)\n\tdefer encoder.Close()\n\tencoder.Indent
    pattern = r'(\tencoder := xml\.NewEncoder\(w\)\n)(\tencoder\.Indent)'
    replacement = r'\1\tdefer encoder.Close()\n\2'
    
    new_content = re.sub(pattern, replacement, content)
    
    # Check if any changes were made
    if new_content != content:
        with open(file_path, 'w') as f:
            f.write(new_content)
        print(f"Fixed encoder.Close() in {file_path}")
        return True
    else:
        print(f"No changes needed in {file_path}")
        return False

def main():
    # Find all Message.go files in pkg/models
    models_dir = "pkg/models"
    files_fixed = 0
    
    for root, dirs, files in os.walk(models_dir):
        for file in files:
            if file == "Message.go":
                file_path = os.path.join(root, file)
                if fix_encoder_close(file_path):
                    files_fixed += 1
    
    print(f"\nTotal files fixed: {files_fixed}")

if __name__ == "__main__":
    main()