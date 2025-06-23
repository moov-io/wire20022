#!/usr/bin/env python3
"""
Revert build*Helper functions back to public (uppercase) since they're used across the models package.

This script changes buildXxxHelper() back to BuildXxxHelper() in typeHelper.go only,
since these functions need to be accessible across the models package.
"""

import os
import re

def revert_helper_functions():
    """Revert helper functions in typeHelper.go back to public."""
    
    file_path = "pkg/models/typeHelper.go"
    print(f"Processing {file_path}...")
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    original_content = content
    
    # Revert function definitions: func buildXxxHelper -> func BuildXxxHelper
    content = re.sub(r'^func build([A-Z][a-zA-Z0-9_]*Helper)\(', r'func Build\1(', content, flags=re.MULTILINE)
    
    # Count fixes made
    old_count = len(re.findall(r'^func build[A-Z][a-zA-Z0-9_]*Helper\(', original_content, re.MULTILINE))
    new_count = len(re.findall(r'^func Build[A-Z][a-zA-Z0-9_]*Helper\(', content, re.MULTILINE))
    fixes_made = new_count - len(re.findall(r'^func Build[A-Z][a-zA-Z0-9_]*Helper\(', original_content, re.MULTILINE))
    
    if fixes_made > 0:
        print(f"  Reverted {fixes_made} helper function definitions to public")
    
    # Write back if changes were made
    if content != original_content:
        with open(file_path, 'w') as f:
            f.write(content)
    
    print(f"Reverted helper functions in typeHelper.go back to public")

if __name__ == "__main__":
    revert_helper_functions()