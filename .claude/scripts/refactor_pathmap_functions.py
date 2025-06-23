#!/usr/bin/env python3
"""
Script to refactor PathMap functions from exported to unexported (private).
This changes PathMapV1() to pathMapV1() across all message types.
"""

import os
import re
import glob

def refactor_pathmap_functions():
    """Refactor all PathMap functions to be unexported."""
    
    # Find all map.go files
    map_files = glob.glob("pkg/models/*/map.go")
    version_files = glob.glob("pkg/models/*/version.go")
    
    total_changes = 0
    
    print("=== Refactoring PathMap Functions to pathMap ===")
    
    # Step 1: Update function definitions in map.go files
    for file_path in map_files:
        print(f"Processing: {file_path}")
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        # Count existing PathMap functions
        existing_count = len(re.findall(r'^func PathMap', content, re.MULTILINE))
        print(f"  Found {existing_count} PathMap functions")
        
        # Replace function definitions: func PathMapV1() -> func pathMapV1()
        new_content = re.sub(r'^func PathMap', 'func pathMap', content, flags=re.MULTILINE)
        
        # Count changes
        new_count = len(re.findall(r'^func pathMap', new_content, re.MULTILINE))
        changes = existing_count
        total_changes += changes
        
        if changes > 0:
            with open(file_path, 'w') as f:
                f.write(new_content)
            print(f"  ✅ Renamed {changes} functions")
        else:
            print(f"  ⚠️ No changes made")
    
    # Step 2: Update function calls in version.go files
    for file_path in version_files:
        print(f"Processing: {file_path}")
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        # Count existing PathMap calls
        existing_calls = len(re.findall(r'PathMap[A-Z0-9]+\(\)', content))
        print(f"  Found {existing_calls} PathMap function calls")
        
        # Replace function calls: PathMapV1() -> pathMapV1()
        new_content = re.sub(r'PathMap([A-Z0-9]+)\(\)', r'pathMap\1()', content)
        
        # Count changes
        new_calls = len(re.findall(r'pathMap[A-Z0-9]+\(\)', new_content))
        call_changes = existing_calls
        
        if call_changes > 0:
            with open(file_path, 'w') as f:
                f.write(new_content)
            print(f"  ✅ Updated {call_changes} function calls")
        else:
            print(f"  ⚠️ No call updates needed")
    
    print(f"\n=== Summary ===")
    print(f"Total PathMap functions renamed: {total_changes}")
    print(f"Files processed: {len(map_files)} map.go + {len(version_files)} version.go")
    print(f"API surface reduced by {total_changes} exported functions! 🎉")

if __name__ == "__main__":
    refactor_pathmap_functions()