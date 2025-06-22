#!/usr/bin/env python3
"""
Script to modularize large test files by splitting version tests into smaller files.
Implements the version group strategy from TEST_FILE_MODULARIZATION.md
"""

import os
import re
import sys
from pathlib import Path

def extract_test_functions(file_content):
    """Extract individual test functions from a Go test file."""
    # Pattern to match test functions and their complete bodies
    pattern = r'(func Test\w+\(t \*testing\.T\) \{(?:[^{}]*\{[^{}]*\})*[^{}]*\})'
    
    functions = []
    lines = file_content.split('\n')
    current_function = []
    in_function = False
    brace_count = 0
    
    for line in lines:
        if re.match(r'^func Test\w+\(t \*testing\.T\) \{', line):
            if current_function:  # Save previous function
                functions.append('\n'.join(current_function))
            current_function = [line]
            in_function = True
            brace_count = line.count('{') - line.count('}')
        elif in_function:
            current_function.append(line)
            brace_count += line.count('{') - line.count('}')
            if brace_count == 0:  # Function complete
                functions.append('\n'.join(current_function))
                current_function = []
                in_function = False
    
    # Don't forget the last function
    if current_function:
        functions.append('\n'.join(current_function))
    
    return functions

def extract_file_header(file_content):
    """Extract package declaration and imports."""
    lines = file_content.split('\n')
    header_lines = []
    
    for line in lines:
        header_lines.append(line)
        # Stop when we hit the first function or significant non-import content
        if line.startswith('func ') or (line.strip() and not line.startswith('package') and 
                                      not line.startswith('import') and not line.startswith('//') and
                                      not line.startswith('*/') and not line.startswith('/*') and
                                      not line.strip() == ')' and '(' not in line):
            header_lines.pop()  # Remove the last line as it's not part of header
            break
    
    return '\n'.join(header_lines)

def group_version_tests(functions, versions_per_group=4):
    """Group version test functions into smaller groups."""
    version_functions = []
    other_functions = []
    
    for func in functions:
        if 'func TestVersion' in func:
            # Extract version number
            match = re.search(r'func TestVersion(\d+)', func)
            if match:
                version_num = int(match.group(1))
                version_functions.append((version_num, func))
            else:
                other_functions.append(func)
        else:
            other_functions.append(func)
    
    # Sort by version number
    version_functions.sort(key=lambda x: x[0])
    
    # Group into chunks
    groups = []
    for i in range(0, len(version_functions), versions_per_group):
        group = version_functions[i:i + versions_per_group]
        if group:
            start_version = group[0][0]
            end_version = group[-1][0]
            group_functions = [func for _, func in group]
            groups.append((start_version, end_version, group_functions))
    
    return groups, other_functions

def create_helper_file_content(header, other_functions):
    """Create content for the helper file with shared functions."""
    if not other_functions:
        return None
    
    content = header + '\n\n'
    content += '// Shared test helpers and utility functions\n\n'
    content += '\n\n'.join(other_functions)
    return content

def create_version_group_file_content(header, start_version, end_version, functions):
    """Create content for a version group file."""
    content = header + '\n\n'
    content += f'// Version tests {start_version:02d}-{end_version:02d}\n\n'
    content += '\n\n'.join(functions)
    return content

def modularize_test_file(file_path, dry_run=True):
    """Modularize a single test file."""
    print(f"Processing {file_path}...")
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Check if file is large enough to warrant splitting
    if len(content) < 30000:  # Less than 30KB
        print(f"  Skipping - file too small ({len(content)} bytes)")
        return False
    
    header = extract_file_header(content)
    functions = extract_test_functions(content)
    
    # Count version tests
    version_count = len([f for f in functions if 'func TestVersion' in f])
    if version_count < 6:  # Don't split if less than 6 versions
        print(f"  Skipping - only {version_count} version tests")
        return False
    
    print(f"  Found {len(functions)} total functions, {version_count} version tests")
    
    # Determine optimal group size
    versions_per_group = max(3, version_count // 3)  # Aim for 3 groups
    
    groups, other_functions = group_version_tests(functions, versions_per_group)
    
    if dry_run:
        print(f"  Would create {len(groups)} version group files:")
        for start, end, funcs in groups:
            print(f"    Message_versions_{start:02d}_{end:02d}_test.go ({len(funcs)} functions)")
        if other_functions:
            print(f"    Message_version_helpers_test.go ({len(other_functions)} functions)")
        return True
    
    # Create directory path
    dir_path = os.path.dirname(file_path)
    base_name = os.path.basename(file_path).replace('Message_version_test.go', '')
    
    # Create version group files
    for start, end, funcs in groups:
        group_file = os.path.join(dir_path, f'{base_name}Message_versions_{start:02d}_{end:02d}_test.go')
        content = create_version_group_file_content(header, start, end, funcs)
        
        with open(group_file, 'w') as f:
            f.write(content)
        print(f"  Created {group_file}")
    
    # Create helper file if needed
    if other_functions:
        helper_file = os.path.join(dir_path, f'{base_name}Message_version_helpers_test.go')
        content = create_helper_file_content(header, other_functions)
        
        with open(helper_file, 'w') as f:
            f.write(content)
        print(f"  Created {helper_file}")
    
    # Rename original file to backup
    backup_file = file_path + '.backup'
    os.rename(file_path, backup_file)
    print(f"  Backed up original to {backup_file}")
    
    return True

def main():
    """Main function to process large test files."""
    dry_run = '--dry-run' in sys.argv
    
    # Find large version test files
    large_files = [
        'pkg/models/ActivityReport/Message_version_test.go',
        'pkg/models/CustomerCreditTransfer/Message_version_test.go',
        'pkg/models/EndpointDetailsReport/Message_version_test.go',
        'pkg/models/Master/Message_version_test.go',
        'pkg/models/PaymentReturn/Message_version_test.go',
        'pkg/models/DrawdownRequest/Message_version_test.go',
        'pkg/models/EndpointTotalsReport/Message_version_test.go',
        'pkg/models/ReturnRequestResponse/Message_version_test.go',
        'pkg/models/DrawdownResponse/Message_version_test.go',
    ]
    
    if dry_run:
        print("DRY RUN MODE - No files will be modified")
        print("=" * 50)
    
    processed = 0
    for file_path in large_files:
        if os.path.exists(file_path):
            if modularize_test_file(file_path, dry_run):
                processed += 1
        else:
            print(f"File not found: {file_path}")
    
    print(f"\nProcessed {processed} files")
    
    if dry_run:
        print("\nTo actually perform the modularization, run:")
        print("python3 .claude/scripts/modularize_test_files.py")

if __name__ == "__main__":
    main()