#!/bin/bash

# Ultra-Fast IP Generator - Example Usage Scripts
# This file demonstrates various ways to use the IP generators

echo "🚀 Ultra-Fast IP Generator Examples"
echo "===================================="

# Example 1: Generate 10K IPs with ultra-fast algorithm
echo "\n📌 Example 1: Generate 10K IPs (Ultra-Fast)"
echo "Command: go run ultra_fast_ip_generator.go 10000"
go run ultra_fast_ip_generator.go 10000

# Example 2: Generate 5K unique IPs
echo "\n📌 Example 2: Generate 5K Unique IPs"
echo "Command: go run ultra_fast_ip_generator.go 5000 unique"
go run ultra_fast_ip_generator.go 5000 unique

# Example 3: Generate 10K IPs with fast algorithm
echo "\n📌 Example 3: Generate 10K IPs (Fast Algorithm)"
echo "Command: go run fast_ip_generator.go 10000"
go run fast_ip_generator.go 10000

# Example 4: Generate 5K IPs with complete validation
echo "\n📌 Example 4: Generate 5K IPs (Validated)"
echo "Command: go run ip_generator.go 5000"
go run ip_generator.go 5000

# Example 5: Performance comparison
echo "\n📌 Example 5: Performance Comparison (100K IPs each)"
echo "Testing Ultra-Fast Algorithm..."
time go run ultra_fast_ip_generator.go 100000

echo "\nTesting Fast Algorithm..."
time go run fast_ip_generator.go 100000

echo "\nTesting Validated Algorithm..."
time go run ip_generator.go 100000

echo "\n✅ Examples completed! Check the generated .txt files."

# Example 6: File operations
echo "\n📌 Example 6: File Operations"
echo "Latest generated files:"
ls -la *_ips_*.txt | tail -3

echo "\n📋 File sizes:"
du -h *_ips_*.txt | tail -3

echo "\n🎯 All examples completed successfully!"

