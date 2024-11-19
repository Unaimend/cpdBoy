#!/bin/bash

# Check if there is at least one argument
if [ $# -eq 0 ]; then
  echo "Usage: cpd <comma-separated list of CPD values>"
  exit 1
fi

# Join arguments into a single string (in case there are multiple parameters)
CPDS="$*"

# Send the request to the server
RESPONSE=$(curl -s -X POST http://127.0.0.1:3000/message \
  -d "{\"text\": \"$CPDS\"}" \
  -H "Content-Type: application/json")

# Print the response
echo "$RESPONSE"

