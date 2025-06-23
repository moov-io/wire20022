#!/bin/bash

# Fix missing defer encoder.Close() calls in WriteXML methods
files=(
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/FedwireFundsPaymentStatus/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/DrawdownRequest/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/DrawdownResponse/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/ReturnRequestResponse/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/EndpointGapReport/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/FedwireFundsAcknowledgement/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/EndpointTotalsReport/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/FedwireFundsSystemResponse/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/PaymentStatusRequest/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/Master/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/PaymentReturn/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/EndpointDetailsReport/Message.go"
    "/Users/wadearnold/Documents/GitHub/wadearnold/wire20022/pkg/models/CustomerCreditTransfer/Message.go"
)

for file in "${files[@]}"; do
    if [[ -f "$file" ]]; then
        echo "Fixing $file"
        sed -i '' 's/\tencoder := xml\.NewEncoder(w)/\tencoder := xml.NewEncoder(w)\n\tdefer encoder.Close()/' "$file"
    fi
done

echo "All files fixed!"