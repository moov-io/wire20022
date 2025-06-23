#!/usr/bin/env python3
"""
Add models. prefix to Build*Helper function calls in MessageHelper.go files.

Since the MessageHelper.go files are in subdirectories and need to access the 
Build*Helper functions from the parent models package, they need the models. prefix.
"""

import os
import re
import glob

def fix_models_prefix():
    """Add models. prefix to Build*Helper function calls."""
    
    helper_files = glob.glob("pkg/models/*/MessageHelper.go")
    total_fixes = 0
    
    for file_path in helper_files:
        print(f"Processing {file_path}...")
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        original_content = content
        
        # Add models. prefix to Build*Helper function calls
        content = re.sub(r'Build([A-Z][a-zA-Z0-9_]*Helper)\(\)', r'models.Build\1()', content)
        
        # Count fixes made in this file
        fixes_in_file = len(re.findall(r'models\.Build[A-Z][a-zA-Z0-9_]*Helper\(\)', content)) - len(re.findall(r'models\.Build[A-Z][a-zA-Z0-9_]*Helper\(\)', original_content))
        if fixes_in_file > 0:
            total_fixes += fixes_in_file
            print(f"  Added models. prefix to {fixes_in_file} helper function calls")
        
        # Write back if changes were made
        if content != original_content:
            with open(file_path, 'w') as f:
                f.write(content)
    
    print(f"\nTotal helper function calls fixed: {total_fixes}")

if __name__ == "__main__":
    fix_models_prefix()