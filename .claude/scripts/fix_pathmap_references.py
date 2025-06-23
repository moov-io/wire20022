#!/usr/bin/env python3
"""
Fix PathMap function references after making them private.

This script updates internal function calls from PathMapVX() to pathMapVX()
to fix compilation errors after the public API refactoring.
"""

import os
import re
import glob

def fix_pathmap_references():
    """Fix all PathMap function references to use lowercase (private) versions."""
    
    # Find all map.go files in pkg/models
    map_files = glob.glob("pkg/models/*/map.go")
    
    total_fixes = 0
    
    for file_path in map_files:
        print(f"Processing {file_path}...")
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        original_content = content
        
        # Fix return statements that call PathMapVX() functions
        # Pattern: return PathMapV(\d+)()
        content = re.sub(r'return PathMap(V\d+)\(\)', r'return pathMap\1()', content)
        
        # Fix direct calls to PathMapVX() functions
        # Pattern: PathMapV(\d+)()
        content = re.sub(r'PathMap(V\d+)\(\)', r'pathMap\1()', content)
        
        # Count fixes made in this file
        fixes_in_file = len(re.findall(r'pathMap(V\d+)\(\)', content)) - len(re.findall(r'pathMap(V\d+)\(\)', original_content))
        if fixes_in_file > 0:
            total_fixes += fixes_in_file
            print(f"  Fixed {fixes_in_file} PathMap references")
        
        # Write back if changes were made
        if content != original_content:
            with open(file_path, 'w') as f:
                f.write(content)
    
    print(f"\nTotal PathMap references fixed: {total_fixes}")

if __name__ == "__main__":
    fix_pathmap_references()