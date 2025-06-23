#!/usr/bin/env python3
"""
Script to refactor BuildXxxHelper functions from exported to unexported (private).
This changes BuildMessageHelper() to buildMessageHelper() across all message types.
"""

import os
import re
import glob

def refactor_helper_functions():
    """Refactor all BuildXxxHelper functions to be unexported."""
    
    # Find all helper files
    helper_files = glob.glob("pkg/models/**/MessageHelper.go", recursive=True)
    helper_files.extend(glob.glob("pkg/models/typeHelper.go"))
    
    total_changes = 0
    
    print("=== Refactoring BuildXxxHelper Functions to buildXxxHelper ===")
    
    # Step 1: Update function definitions in helper files
    for file_path in helper_files:
        print(f"Processing: {file_path}")
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        # Count existing Build functions
        existing_count = len(re.findall(r'^func Build.*Helper', content, re.MULTILINE))
        print(f"  Found {existing_count} Build functions")
        
        # Replace function definitions: func BuildXxxHelper() -> func buildXxxHelper()
        new_content = re.sub(r'^func Build([A-Z][a-zA-Z0-9_]*Helper)', r'func build\1', content, flags=re.MULTILINE)
        
        # Count changes
        new_count = len(re.findall(r'^func build.*Helper', new_content, re.MULTILINE))
        changes = existing_count
        total_changes += changes
        
        if changes > 0:
            with open(file_path, 'w') as f:
                f.write(new_content)
            print(f"  ✅ Renamed {changes} functions")
        else:
            print(f"  ⚠️ No changes made")
    
    # Step 2: Update function calls in files that use these helpers
    # Find files that call Build*Helper functions
    all_go_files = glob.glob("pkg/**/*.go", recursive=True)
    
    for file_path in all_go_files:
        if file_path.endswith("_test.go"):
            continue  # Skip test files for now
            
        print(f"Checking calls in: {file_path}")
        
        with open(file_path, 'r') as f:
            content = f.read()
        
        # Count existing Build function calls
        existing_calls = len(re.findall(r'Build([A-Z][a-zA-Z0-9_]*Helper)\(', content))
        
        if existing_calls > 0:
            print(f"  Found {existing_calls} Build function calls")
            
            # Replace function calls: Build*Helper() -> build*Helper()
            new_content = re.sub(r'Build([A-Z][a-zA-Z0-9_]*Helper)\(', r'build\1(', content)
            
            with open(file_path, 'w') as f:
                f.write(new_content)
            print(f"  ✅ Updated {existing_calls} function calls")
    
    print(f"\n=== Summary ===")
    print(f"Total BuildXxxHelper functions renamed: {total_changes}")
    print(f"Files processed: {len(helper_files)} helper files + call updates")
    print(f"API surface reduced by {total_changes} more exported functions! 🎉")

if __name__ == "__main__":
    refactor_helper_functions()