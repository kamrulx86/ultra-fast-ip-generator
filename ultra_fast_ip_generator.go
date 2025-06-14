package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Valid public IP ranges for maximum speed
var fastRanges = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 170, 171, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223}

// generateUltraFastIP generates IPs at maximum speed
func generateUltraFastIP(rng *rand.Rand) string {
	firstOctet := fastRanges[rng.Intn(len(fastRanges))]
	return fmt.Sprintf("%d.%d.%d.%d", firstOctet, rng.Intn(256), rng.Intn(256), rng.Intn(256))
}

// generateBatch generates a batch of IPs using string builder for maximum efficiency
func generateBatch(rng *rand.Rand, batchSize int) []string {
	batch := make([]string, batchSize)
	for i := 0; i < batchSize; i++ {
		batch[i] = generateUltraFastIP(rng)
	}
	return batch
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run ultra_fast_ip_generator.go <number_of_ips> [unique]")
		fmt.Println("Examples:")
		fmt.Println("  go run ultra_fast_ip_generator.go 500000")
		fmt.Println("  go run ultra_fast_ip_generator.go 500000 unique")
		os.Exit(1)
	}

	// Parse the number of IPs to generate
	numIPs, err := strconv.Atoi(os.Args[1])
	if err != nil || numIPs <= 0 {
		fmt.Println("Error: Please provide a valid positive number")
		os.Exit(1)
	}

	// Check if unique mode is requested
	uniqueMode := len(os.Args) == 3 && strings.ToLower(os.Args[2]) == "unique"

	if uniqueMode {
		fmt.Printf("Generating %d UNIQUE public IP addresses (slower but unique)...\n", numIPs)
	} else {
		fmt.Printf("Generating %d public IP addresses (ultra-fast mode, may have duplicates)...\n", numIPs)
	}

	// Create a dedicated random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	startTime := time.Now()
	var ips []string

	if uniqueMode {
		// Unique mode - use map for deduplication
		ipSet := make(map[string]bool)
		ips = make([]string, 0, numIPs)
		
		batchSize := 10000
		for len(ips) < numIPs {
			batch := generateBatch(rng, batchSize)
			for _, ip := range batch {
				if !ipSet[ip] {
					ipSet[ip] = true
					ips = append(ips, ip)
					if len(ips) >= numIPs {
						break
					}
				}
			}
			
			// Show progress for unique mode
			if len(ips)%50000 == 0 || len(ips) >= numIPs {
				progress := float64(len(ips)) / float64(numIPs) * 100
				fmt.Printf("Progress: %.1f%% (%d/%d unique IPs)\n", progress, len(ips), numIPs)
			}
		}
		ips = ips[:numIPs] // Trim to exact count
	} else {
		// Ultra-fast mode - no duplicate checking
		ips = make([]string, numIPs)
		batchSize := 50000
		
		for i := 0; i < numIPs; i += batchSize {
			currentBatchSize := batchSize
			if i+batchSize > numIPs {
				currentBatchSize = numIPs - i
			}
			
			batch := generateBatch(rng, currentBatchSize)
			copy(ips[i:], batch)
			
			// Show progress
			progress := float64(i+currentBatchSize) / float64(numIPs) * 100
			fmt.Printf("Progress: %.1f%% (%d/%d IPs generated)\n", progress, i+currentBatchSize, numIPs)
		}
	}

	// Create output filename
	timestamp := time.Now().Format("20060102_150405")
	mode := "ultra_fast"
	if uniqueMode {
		mode = "unique"
	}
	filename := fmt.Sprintf("%s_ips_%s.txt", mode, timestamp)

	// Write to file with large buffer
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, 1024*1024) // 1MB buffer
	for _, ip := range ips {
		writer.WriteString(ip + "\n")
	}
	writer.Flush()

	elapsed := time.Since(startTime)
	fmt.Printf("\n🚀 COMPLETED SUCCESSFULLY! 🚀\n")
	fmt.Printf("Generated: %d public IP addresses\n", len(ips))
	fmt.Printf("Time taken: %v\n", elapsed)
	fmt.Printf("Rate: %.0f IPs/second\n", float64(len(ips))/elapsed.Seconds())
	fmt.Printf("File size: %.2f MB\n", float64(len(ips)*15)/1024/1024)
	fmt.Printf("Output saved to: %s\n", filename)

	// Show samples
	fmt.Println("\n📋 Sample IPs:")
	for i := 0; i < 5 && i < len(ips); i++ {
		fmt.Printf("  %s\n", ips[i])
	}
	if len(ips) > 5 {
		fmt.Println("  ...")
		for i := len(ips) - 3; i < len(ips); i++ {
			fmt.Printf("  %s\n", ips[i])
		}
	}
}

