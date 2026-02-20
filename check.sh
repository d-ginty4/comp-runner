#!/bin/bash

# Candidate names and suffixes
names=("titan" "atlas" "hercules" "barbell" "mammoth" "power")
suffixes=("forge" "hub" "central" "hq" "strength" "arena" "vault" "lift" "works" "x")

check_domain() {
    local domain=$1
    # Run whois quietly
    local result
    result=$(whois "$domain" 2>/dev/null)
    
    # Simple check for availability
    if echo "$result" | grep -qiE "No match|NOT FOUND|No entries"; then
        echo "$domain: AVAILABLE"
    fi
}

# Loop through all combinations
for name in "${names[@]}"; do
    for suffix in "${suffixes[@]}"; do
		domain="$name$suffix.com"
		check_domain "$domain" &
    done
done

wait  # Wait for all background processes to finish
