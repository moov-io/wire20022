#!/usr/bin/env python3

import os
import re

# List of all message types to fix
MESSAGE_TYPES = [
    "AccountReportingRequest", "ActivityReport", "ConnectionCheck", "CustomerCreditTransfer",
    "DrawdownRequest", "DrawdownResponse", "EndpointDetailsReport", "EndpointGapReport", 
    "EndpointTotalsReport", "FedwireFundsAcknowledgement", "FedwireFundsPaymentStatus",
    "FedwireFundsSystemResponse", "Master", "PaymentReturn", "PaymentStatusRequest", 
    "ReturnRequestResponse"
]

def fix_message_type(message_type):
    """Fix MessageWith reference in a message type file"""
    file_path = f"pkg/messages/{message_type}.go"
    
    if not os.path.exists(file_path):
        print(f"Warning: {file_path} not found")
        return False
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Replace the MessageWith reference with ParseXML wrapper
    model_var = f"{message_type}Model"
    
    old_pattern = f"{model_var}.MessageWith,.*// Type-safe XML converter"
    new_replacement = f"""func(data []byte) ({model_var}.MessageModel, error) {{  // XML converter using new API
\t\t\t\tmsg, err := {model_var}.ParseXML(data)
\t\t\t\tif err != nil {{
\t\t\t\t\treturn {model_var}.MessageModel{{}}, err
\t\t\t\t}}
\t\t\t\treturn *msg, nil
\t\t\t}},"""
    
    content = re.sub(old_pattern, new_replacement, content)
    
    with open(file_path, 'w') as f:
        f.write(content)
    
    return True

def main():
    """Main function"""
    print("Fixing MessageWith references in messages package...")
    
    for message_type in MESSAGE_TYPES:
        print(f"Processing {message_type}...")
        fix_message_type(message_type)
    
    print("Done! Testing compilation...")
    
    # Test compilation
    import subprocess
    try:
        result = subprocess.run(['go', 'build', './pkg/messages'], 
                              capture_output=True, text=True)
        if result.returncode == 0:
            print("✅ Messages package compiles successfully!")
        else:
            print(f"❌ Messages package compilation failed: {result.stderr.strip()}")
    except Exception as e:
        print(f"❌ Error testing compilation: {e}")

if __name__ == "__main__":
    main()