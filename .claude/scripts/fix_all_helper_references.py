#!/usr/bin/env python3
"""
Fix all helper function references after making the core ones public again.

This script:
1. Updates MessageHelper.go files to use BuildXxxHelper() (uppercase)
2. Updates internal references in typeHelper.go to use BuildXxxHelper() (uppercase)
"""

import os
import re
import glob

def fix_message_helper_references():
    """Fix helper function references in MessageHelper.go files."""
    
    helper_files = glob.glob("pkg/models/*/MessageHelper.go")
    total_fixes = 0
    
    for file_path in helper_files:
        print(f"Processing {file_path}...")
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        original_content = content
        
        # Fix references to build*Helper() functions -> Build*Helper()
        content = re.sub(r'build([A-Z][a-zA-Z0-9_]*Helper)\(\)', r'Build\1()', content)
        
        # Count fixes made in this file
        fixes_in_file = len(re.findall(r'Build[A-Z][a-zA-Z0-9_]*Helper\(\)', content)) - len(re.findall(r'Build[A-Z][a-zA-Z0-9_]*Helper\(\)', original_content))
        if fixes_in_file > 0:
            total_fixes += fixes_in_file
            print(f"  Fixed {fixes_in_file} helper function references")
        
        # Write back if changes were made
        if content != original_content:
            with open(file_path, 'w') as f:
                f.write(content)
    
    return total_fixes

def fix_typehelper_internal_references():
    """Fix internal references in typeHelper.go."""
    
    file_path = "pkg/models/typeHelper.go"
    print(f"Processing internal references in {file_path}...")
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    original_content = content
    
    # Fix internal calls to build*Helper() functions -> Build*Helper()
    content = re.sub(r'build([A-Z][a-zA-Z0-9_]*Helper)\(\)', r'Build\1()', content)
    
    # Count fixes made
    fixes_in_file = len(re.findall(r'Build[A-Z][a-zA-Z0-9_]*Helper\(\)', content)) - len(re.findall(r'Build[A-Z][a-zA-Z0-9_]*Helper\(\)', original_content))
    if fixes_in_file > 0:
        print(f"  Fixed {fixes_in_file} internal helper function references")
    
    # Write back if changes were made
    if content != original_content:
        with open(file_path, 'w') as f:
            f.write(content)
    
    return fixes_in_file

if __name__ == "__main__":
    message_fixes = fix_message_helper_references()
    internal_fixes = fix_typehelper_internal_references()
    print(f"\nTotal helper function references fixed: {message_fixes + internal_fixes}")