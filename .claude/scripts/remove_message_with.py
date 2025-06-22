#!/usr/bin/env python3

import os
import re

# List of all message types to process
MESSAGE_TYPES = [
    "AccountReportingRequest", "ActivityReport", "ConnectionCheck", "CustomerCreditTransfer",
    "DrawdownRequest", "DrawdownResponse", "EndpointDetailsReport", "EndpointGapReport", 
    "EndpointTotalsReport", "FedwireFundsAcknowledgement", "FedwireFundsPaymentStatus",
    "FedwireFundsSystemResponse", "Master", "PaymentReturn", "PaymentStatusRequest", 
    "ReturnRequestResponse"
]

def remove_message_with(message_type):
    """Remove MessageWith function and deprecated comments from a message type"""
    file_path = f"pkg/models/{message_type}/Message.go"
    
    if not os.path.exists(file_path):
        print(f"Warning: {file_path} not found")
        return False
    
    with open(file_path, 'r') as f:
        content = f.read()
    
    # Remove the deprecated MessageWith function and its comment
    content = re.sub(
        r'// Deprecated: Use ParseXML instead\nfunc MessageWith\(data \[\]byte\) \(MessageModel, error\) \{\n\treturn processor\.ProcessMessage\(data\)\n\}\n\n',
        '',
        content
    )
    
    # Remove any remaining MessageWith references in comments
    content = re.sub(
        r'// MessageWith uses base abstractions to replace \d+\+ lines with a single call\n',
        '',
        content
    )
    
    # Clean up any double newlines that might result
    content = re.sub(r'\n\n\n+', '\n\n', content)
    
    with open(file_path, 'w') as f:
        f.write(content)
    
    return True

def main():
    """Main function"""
    print("Removing deprecated MessageWith functions...")
    
    for message_type in MESSAGE_TYPES:
        print(f"Processing {message_type}...")
        remove_message_with(message_type)
    
    print("Done! Testing compilation...")
    
    # Test compilation
    import subprocess
    failed = []
    for message_type in MESSAGE_TYPES:
        try:
            result = subprocess.run(['go', 'build', f'./pkg/models/{message_type}'], 
                                  capture_output=True, text=True)
            if result.returncode == 0:
                print(f"✅ {message_type}")
            else:
                print(f"❌ {message_type}: {result.stderr.strip()}")
                failed.append(message_type)
        except Exception as e:
            print(f"❌ {message_type}: {e}")
            failed.append(message_type)
    
    if failed:
        print(f"\nFailed message types: {failed}")
    else:
        print("\n🎉 All message types compile successfully!")

if __name__ == "__main__":
    main()