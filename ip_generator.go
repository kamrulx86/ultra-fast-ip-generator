package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

// isPrivateIP checks if an IP address is private
func isPrivateIP(ip net.IP) bool {
	// Check for IPv4 private ranges
	if ip.To4() != nil {
		// 10.0.0.0/8
		if ip[0] == 10 {
			return true
		}
		// 172.16.0.0/12
		if ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31 {
			return true
		}
		// 192.168.0.0/16
		if ip[0] == 192 && ip[1] == 168 {
			return true
		}
		// 127.0.0.0/8 (loopback)
		if ip[0] == 127 {
			return true
		}
		// 169.254.0.0/16 (link-local)
		if ip[0] == 169 && ip[1] == 254 {
			return true
		}
		// 224.0.0.0/4 (multicast)
		if ip[0] >= 224 && ip[0] <= 239 {
			return true
		}
		// 240.0.0.0/4 (reserved)
		if ip[0] >= 240 {
			return true
		}
	}
	return false
}

// generateRandomPublicIP generates a random public IPv4 address
func generateRandomPublicIP() net.IP {
	for {
		// Generate random IPv4 address with better distribution
		// Avoid ranges that are likely to be private/invalid
		firstOctet := byte(1 + rand.Intn(223)) // 1-223, avoiding 0, 224-255
		
		// Skip known private ranges for first octet
		if firstOctet == 10 || firstOctet == 127 || firstOctet == 172 ||
		   firstOctet == 192 || firstOctet == 169 {
			continue
		}
		
		ip := net.IPv4(
			firstOctet,
			byte(rand.Intn(256)),
			byte(rand.Intn(256)),
			byte(rand.Intn(256)),
		)
		
		// Double-check it's not private
		if !isPrivateIP(ip) {
			return ip
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run ip_generator.go <number_of_ips>")
		fmt.Println("Example: go run ip_generator.go 500000")
		os.Exit(1)
	}

	// Parse the number of IPs to generate
	numIPs, err := strconv.Atoi(os.Args[1])
	if err != nil || numIPs <= 0 {
		fmt.Println("Error: Please provide a valid positive number")
		os.Exit(1)
	}

	fmt.Printf("Generating %d unique public IP addresses...\n", numIPs)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Use a map to store unique IPs
	ipSet := make(map[string]bool)
	var ips []string

	// Progress tracking
	startTime := time.Now()
	progressInterval := numIPs / 100 // Show progress every 1%
	if progressInterval == 0 {
		progressInterval = 1
	}

	// Generate unique public IPs
	for len(ips) < numIPs {
		ip := generateRandomPublicIP()
		ipStr := ip.String()
		
		// Check if IP is already generated
		if !ipSet[ipStr] {
			ipSet[ipStr] = true
			ips = append(ips, ipStr)
			
			// Show progress
			if len(ips)%progressInterval == 0 {
				progress := float64(len(ips)) / float64(numIPs) * 100
				fmt.Printf("Progress: %.1f%% (%d/%d IPs generated)\n", progress, len(ips), numIPs)
			}
		}
	}

	// Create output filename with timestamp
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("generated_ips_%s.txt", timestamp)

	// Write IPs to file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, ip := range ips {
		_, err := writer.WriteString(ip + "\n")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
	}
	writer.Flush()

	elapsed := time.Since(startTime)
	fmt.Printf("\nCompleted successfully!\n")
	fmt.Printf("Generated %d unique public IP addresses\n", len(ips))
	fmt.Printf("Time taken: %v\n", elapsed)
	fmt.Printf("Output saved to: %s\n", filename)

	// Show some sample IPs
	fmt.Println("\nFirst 10 generated IPs:")
	for i := 0; i < 10 && i < len(ips); i++ {
		fmt.Printf("%d. %s\n", i+1, ips[i])
	}
}

