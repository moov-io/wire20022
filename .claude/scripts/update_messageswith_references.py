#!/usr/bin/env python3
"""
Script to update all references to the old MessageWith function to use the new ParseXML function.
Updates both Go code comments and README documentation.
"""

import os
import re
import glob

def update_go_comments():
    """Update Go code comments that reference MessageWith"""
    
    # Find all Message.go files
    message_files = glob.glob('/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/*/Message.go')
    
    old_comment = "// This replaces the non-idiomatic MessageWith function"
    new_comment = "// This is the primary function for parsing XML from byte data"
    
    updated_files = []
    
    for file_path in message_files:
        try:
            with open(file_path, 'r') as f:
                content = f.read()
            
            if old_comment in content:
                new_content = content.replace(old_comment, new_comment)
                
                with open(file_path, 'w') as f:
                    f.write(new_content)
                
                updated_files.append(file_path)
                print(f"✅ Updated comment in {os.path.basename(os.path.dirname(file_path))}/Message.go")
        
        except Exception as e:
            print(f"❌ Error updating {file_path}: {e}")
    
    return updated_files

def update_readme_files():
    """Update README files that reference MessageWith function"""
    
    # Find all README.md files in model directories
    readme_files = glob.glob('/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/*/README.md')
    
    # Pattern to find MessageWith usage in documentation
    messageswith_pattern = r'model, err := MessageWith\(xmlBytes\)'
    replacement = 'model, err := ParseXML(xmlBytes)'
    
    # Pattern to find other MessageWith references
    messageswith_func_pattern = r'MessageWith\('
    replacement_func = 'ParseXML('
    
    updated_files = []
    
    for file_path in readme_files:
        try:
            with open(file_path, 'r') as f:
                content = f.read()
            
            # Check if file contains MessageWith references
            if 'MessageWith' in content:
                # Replace the specific pattern first
                new_content = re.sub(messageswith_pattern, replacement, content)
                
                # Replace any remaining MessageWith function calls
                new_content = re.sub(messageswith_func_pattern, replacement_func, new_content)
                
                # Also update any remaining references in documentation text
                new_content = new_content.replace('MessageWith', 'ParseXML')
                
                with open(file_path, 'w') as f:
                    f.write(new_content)
                
                updated_files.append(file_path)
                print(f"✅ Updated README in {os.path.basename(os.path.dirname(file_path))}/README.md")
        
        except Exception as e:
            print(f"❌ Error updating {file_path}: {e}")
    
    return updated_files

def main():
    """Main function to run all updates"""
    print("🔄 Updating MessageWith references to ParseXML...")
    print()
    
    print("📝 Updating Go code comments...")
    go_files = update_go_comments()
    print(f"Updated {len(go_files)} Go files")
    print()
    
    print("📚 Updating README files...")
    readme_files = update_readme_files()
    print(f"Updated {len(readme_files)} README files")
    print()
    
    total_files = len(go_files) + len(readme_files)
    print(f"🎉 Successfully updated {total_files} total files!")
    print()
    print("All MessageWith references have been updated to use ParseXML")

if __name__ == "__main__":
    main()