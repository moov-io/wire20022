#!/usr/bin/env python3
"""
Fix remaining Build*Helper() calls to use build*Helper() 
"""

import os
import re
import glob

def fix_helper_calls():
    """Fix all remaining Build*Helper() calls."""
    
    files = glob.glob("pkg/models/**/*.go", recursive=True)
    files = [f for f in files if not f.endswith("_test.go")]
    
    total_fixes = 0
    
    for file_path in files:
        with open(file_path, 'r') as f:
            content = f.read()
        
        # Count existing Build function calls
        build_calls = re.findall(r'Build([A-Z][a-zA-Z0-9_]*Helper)\(', content)
        
        if build_calls:
            print(f"Fixing {len(build_calls)} calls in {file_path}")
            
            # Replace Build*Helper() calls with build*Helper()
            new_content = re.sub(r'Build([A-Z][a-zA-Z0-9_]*Helper)\(', r'build\1(', content)
            
            with open(file_path, 'w') as f:
                f.write(new_content)
                
            total_fixes += len(build_calls)
    
    print(f"Fixed {total_fixes} Build*Helper() calls")

if __name__ == "__main__":
    fix_helper_calls()