#!/usr/bin/env python3
"""
Fix function definitions that incorrectly have models. prefix.

Remove models. prefix from function definitions like:
func models.BuildMessageHelper() -> func BuildMessageHelper()
"""

import os
import re
import glob

def fix_func_definitions():
    """Fix function definitions that have incorrect models. prefix."""
    
    helper_files = glob.glob("pkg/models/*/MessageHelper.go")
    total_fixes = 0
    
    for file_path in helper_files:
        print(f"Processing {file_path}...")
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        original_content = content
        
        # Fix function definitions: func models.Build* -> func Build*
        content = re.sub(r'^func models\.(Build[A-Z][a-zA-Z0-9_]*Helper)\(', r'func \1(', content, flags=re.MULTILINE)
        
        # Count fixes made in this file
        fixes_in_file = len(re.findall(r'^func Build[A-Z][a-zA-Z0-9_]*Helper\(', content, re.MULTILINE)) - len(re.findall(r'^func Build[A-Z][a-zA-Z0-9_]*Helper\(', original_content, re.MULTILINE))
        if fixes_in_file > 0:
            total_fixes += fixes_in_file
            print(f"  Fixed {fixes_in_file} function definitions")
        
        # Write back if changes were made
        if content != original_content:
            with open(file_path, 'w') as f:
                f.write(content)
    
    print(f"\nTotal function definitions fixed: {total_fixes}")

if __name__ == "__main__":
    fix_func_definitions()