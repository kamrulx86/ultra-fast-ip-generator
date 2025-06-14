package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Valid public IP ranges (first octet ranges that are definitely public)
var publicRanges = []struct {
	start, end int
}{
	{1, 9},     // 1.x.x.x - 9.x.x.x (excluding 10.x.x.x which is private)
	{11, 126},  // 11.x.x.x - 126.x.x.x (excluding 127.x.x.x which is loopback)
	{128, 168}, // 128.x.x.x - 168.x.x.x (excluding 169.x.x.x which is link-local)
	{170, 171}, // 170.x.x.x - 171.x.x.x (excluding 172.x.x.x which has private ranges)
	{173, 191}, // 173.x.x.x - 191.x.x.x (excluding 192.x.x.x which has private ranges)
	{193, 223}, // 193.x.x.x - 223.x.x.x (before multicast range)
}

// generateFastPublicIP generates a random public IP quickly
func generateFastPublicIP(rng *rand.Rand) string {
	// Pick a random public range
	rangeIdx := rng.Intn(len(publicRanges))
	selectedRange := publicRanges[rangeIdx]
	
	// Generate first octet from the selected range
	firstOctet := selectedRange.start + rng.Intn(selectedRange.end-selectedRange.start+1)
	
	// Handle special cases for ranges that might have private subnets
	var secondOctet int
	if firstOctet == 172 {
		// For 172.x.x.x, avoid 172.16-31.x.x (private range)
		for {
			secondOctet = rng.Intn(256)
			if secondOctet < 16 || secondOctet > 31 {
				break
			}
		}
	} else if firstOctet == 192 {
		// For 192.x.x.x, avoid 192.168.x.x (private range)
		for {
			secondOctet = rng.Intn(256)
			if secondOctet != 168 {
				break
			}
		}
	} else if firstOctet == 169 {
		// For 169.x.x.x, avoid 169.254.x.x (link-local)
		for {
			secondOctet = rng.Intn(256)
			if secondOctet != 254 {
				break
			}
		}
	} else {
		secondOctet = rng.Intn(256)
	}
	
	thirdOctet := rng.Intn(256)
	fourthOctet := rng.Intn(256)
	
	return fmt.Sprintf("%d.%d.%d.%d", firstOctet, secondOctet, thirdOctet, fourthOctet)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run fast_ip_generator.go <number_of_ips>")
		fmt.Println("Example: go run fast_ip_generator.go 500000")
		os.Exit(1)
	}

	// Parse the number of IPs to generate
	numIPs, err := strconv.Atoi(os.Args[1])
	if err != nil || numIPs <= 0 {
		fmt.Println("Error: Please provide a valid positive number")
		os.Exit(1)
	}

	fmt.Printf("Generating %d public IP addresses (fast mode)...\n", numIPs)

	// Create a dedicated random number generator for better performance
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Pre-allocate slice for better performance
	ips := make([]string, 0, numIPs)

	// Progress tracking
	startTime := time.Now()
	progressInterval := numIPs / 20 // Show progress every 5%
	if progressInterval == 0 {
		progressInterval = 1
	}

	// Generate IPs without duplicate checking for maximum speed
	for i := 0; i < numIPs; i++ {
		ip := generateFastPublicIP(rng)
		ips = append(ips, ip)
		
		// Show progress
		if (i+1)%progressInterval == 0 || i == numIPs-1 {
			progress := float64(i+1) / float64(numIPs) * 100
			fmt.Printf("Progress: %.1f%% (%d/%d IPs generated)\n", progress, i+1, numIPs)
		}
	}

	// Create output filename with timestamp
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("fast_generated_ips_%s.txt", timestamp)

	// Write IPs to file efficiently
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Use buffered writer for better performance
	writer := bufio.NewWriterSize(file, 64*1024) // 64KB buffer
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
	fmt.Printf("Generated %d public IP addresses\n", len(ips))
	fmt.Printf("Time taken: %v\n", elapsed)
	fmt.Printf("Rate: %.0f IPs/second\n", float64(numIPs)/elapsed.Seconds())
	fmt.Printf("Output saved to: %s\n", filename)

	// Show some sample IPs
	fmt.Println("\nFirst 10 generated IPs:")
	for i := 0; i < 10 && i < len(ips); i++ {
		fmt.Printf("%d. %s\n", i+1, ips[i])
	}
	
	if len(ips) > 10 {
		fmt.Println("\nLast 5 generated IPs:")
		for i := len(ips) - 5; i < len(ips); i++ {
			fmt.Printf("%d. %s\n", i+1, ips[i])
		}
	}
}

