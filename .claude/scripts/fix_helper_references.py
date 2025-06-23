#!/usr/bin/env python3
"""
Fix helper function references after making them private.

This script updates references from models.buildXxxHelper() to models.buildXxxHelper()
by removing the package prefix since they're now private within the models package.
"""

import os
import re
import glob

def fix_helper_references():
    """Fix all helper function references to use private versions."""
    
    # Find all MessageHelper.go files in pkg/models
    helper_files = glob.glob("pkg/models/*/MessageHelper.go")
    
    total_fixes = 0
    
    for file_path in helper_files:
        print(f"Processing {file_path}...")
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        original_content = content
        
        # Fix references to models.buildXxxHelper() functions
        # Pattern: models.build(\w+Helper)\(\)
        content = re.sub(r'models\.build(\w+Helper)\(\)', r'build\1()', content)
        
        # Count fixes made in this file
        old_count = len(re.findall(r'models\.build\w+Helper\(\)', original_content))
        new_count = len(re.findall(r'build\w+Helper\(\)', content))
        fixes_in_file = old_count
        
        if fixes_in_file > 0:
            total_fixes += fixes_in_file
            print(f"  Fixed {fixes_in_file} helper function references")
        
        # Write back if changes were made
        if content != original_content:
            with open(file_path, 'w') as f:
                f.write(content)
    
    print(f"\nTotal helper function references fixed: {total_fixes}")

if __name__ == "__main__":
    fix_helper_references()