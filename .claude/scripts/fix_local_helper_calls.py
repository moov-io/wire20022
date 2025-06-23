#!/usr/bin/env python3
"""
Fix local helper function calls that incorrectly have models. prefix.

When calling helper functions defined in the same file, they shouldn't have the models. prefix.
Only calls to functions defined in typeHelper.go should have the models. prefix.
"""

import os
import re
import glob

def get_local_functions(file_path):
    """Get list of Build*Helper functions defined in this file."""
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Find all function definitions in this file
    local_functions = re.findall(r'^func (Build[A-Z][a-zA-Z0-9_]*Helper)\(', content, re.MULTILINE)
    return local_functions

def fix_local_helper_calls():
    """Fix calls to locally defined helper functions."""
    
    helper_files = glob.glob("pkg/models/*/MessageHelper.go")
    total_fixes = 0
    
    for file_path in helper_files:
        print(f"Processing {file_path}...")
        
        # Get local functions defined in this file
        local_functions = get_local_functions(file_path)
        
        if not local_functions:
            continue
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        original_content = content
        
        # Remove models. prefix from calls to locally defined functions
        for func_name in local_functions:
            content = re.sub(rf'models\.{func_name}\(\)', f'{func_name}()', content)
            
        # Count fixes made in this file
        fixes_in_file = 0
        for func_name in local_functions:
            fixes_in_file += len(re.findall(rf'{func_name}\(\)', content)) - len(re.findall(rf'{func_name}\(\)', original_content))
        
        if fixes_in_file > 0:
            total_fixes += fixes_in_file
            print(f"  Fixed {fixes_in_file} local helper function calls")
        
        # Write back if changes were made
        if content != original_content:
            with open(file_path, 'w') as f:
                f.write(content)
    
    print(f"\nTotal local helper function calls fixed: {total_fixes}")

if __name__ == "__main__":
    fix_local_helper_calls()