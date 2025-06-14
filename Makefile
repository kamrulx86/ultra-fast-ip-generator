# Ultra-Fast IP Generator Makefile

.PHONY: all build clean test examples help

# Default target
all: build

# Build all generators
build:
	@echo "🚀 Building Ultra-Fast IP Generators..."
	go build -ldflags="-s -w" -o bin/ultra_ip_gen ultra_fast_ip_generator.go
	go build -ldflags="-s -w" -o bin/fast_ip_gen fast_ip_generator.go
	go build -ldflags="-s -w" -o bin/ip_gen ip_generator.go
	@echo "✅ Build completed! Binaries are in ./bin/"

# Create bin directory
bin:
	mkdir -p bin

# Build with optimizations
build-optimized: bin
	@echo "⚡ Building optimized binaries..."
	go build -ldflags="-s -w" -gcflags="-l=4" -o bin/ultra_ip_gen ultra_fast_ip_generator.go
	go build -ldflags="-s -w" -gcflags="-l=4" -o bin/fast_ip_gen fast_ip_generator.go
	go build -ldflags="-s -w" -gcflags="-l=4" -o bin/ip_gen ip_generator.go
	@echo "✅ Optimized build completed!"

# Cross-compile for different platforms
build-all: bin
	@echo "🌍 Cross-compiling for all platforms..."
	# Linux
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/ultra_ip_gen_linux_amd64 ultra_fast_ip_generator.go
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/ultra_ip_gen_linux_arm64 ultra_fast_ip_generator.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/ultra_ip_gen_windows_amd64.exe ultra_fast_ip_generator.go
	# macOS
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/ultra_ip_gen_darwin_amd64 ultra_fast_ip_generator.go
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/ultra_ip_gen_darwin_arm64 ultra_fast_ip_generator.go
	@echo "✅ Cross-compilation completed!"

# Run basic tests
test:
	@echo "🧪 Testing IP generators..."
	@echo "Testing ultra-fast generator (1000 IPs)..."
	go run ultra_fast_ip_generator.go 1000
	@echo "Testing fast generator (1000 IPs)..."
	go run fast_ip_generator.go 1000
	@echo "Testing validated generator (1000 IPs)..."
	go run ip_generator.go 1000
	@echo "✅ All tests passed!"

# Run comprehensive tests
test-comprehensive:
	@echo "🧪 Running comprehensive tests..."
	@echo "Testing with 10K IPs each..."
	time go run ultra_fast_ip_generator.go 10000
	time go run fast_ip_generator.go 10000
	time go run ip_generator.go 10000
	@echo "✅ Comprehensive tests completed!"

# Run performance benchmarks
benchmark:
	@echo "📈 Running performance benchmarks..."
	@echo "Ultra-Fast Algorithm (100K IPs):"
	time go run ultra_fast_ip_generator.go 100000
	@echo "\nFast Algorithm (100K IPs):"
	time go run fast_ip_generator.go 100000
	@echo "\nValidated Algorithm (50K IPs):"
	time go run ip_generator.go 50000

# Run examples
examples:
	@echo "🎆 Running examples..."
	./examples.sh

# Format code
format:
	@echo "🎨 Formatting code..."
	go fmt *.go
	@echo "✅ Code formatted!"

# Check code quality
vet:
	@echo "🔍 Checking code quality..."
	go vet *.go
	@echo "✅ Code quality check passed!"

# Clean generated files
clean:
	@echo "🧹 Cleaning up..."
	rm -rf bin/
	rm -f *.txt
	rm -rf ip_chunks_*/
	@echo "✅ Cleanup completed!"

# Clean only IP files (keep binaries)
clean-ips:
	@echo "🧹 Cleaning IP files..."
	rm -f *_ips_*.txt
	rm -f generated_ips_*.txt
	rm -rf ip_chunks_*/
	@echo "✅ IP files cleaned!"

# Install (copy binaries to system PATH)
install: build
	@echo "📦 Installing binaries..."
	cp bin/ultra_ip_gen /usr/local/bin/
	cp bin/fast_ip_gen /usr/local/bin/
	cp bin/ip_gen /usr/local/bin/
	@echo "✅ Installation completed!"

# Uninstall
uninstall:
	@echo "🗂️ Uninstalling binaries..."
	rm -f /usr/local/bin/ultra_ip_gen
	rm -f /usr/local/bin/fast_ip_gen
	rm -f /usr/local/bin/ip_gen
	@echo "✅ Uninstallation completed!"

# Development setup
dev-setup:
	@echo "🛠️ Setting up development environment..."
	go mod init ultra-fast-ip-generator 2>/dev/null || true
	go mod tidy
	mkdir -p bin
	@echo "✅ Development environment ready!"

# Release preparation
release: clean format vet test build-all
	@echo "🎉 Release preparation completed!"
	@echo "Binaries created:"
	@ls -la bin/

# Help
help:
	@echo "🚀 Ultra-Fast IP Generator - Available Commands:"
	@echo ""
	@echo "Build Commands:"
	@echo "  build              - Build all generators"
	@echo "  build-optimized    - Build with optimizations"
	@echo "  build-all          - Cross-compile for all platforms"
	@echo ""
	@echo "Test Commands:"
	@echo "  test               - Run basic tests"
	@echo "  test-comprehensive - Run comprehensive tests"
	@echo "  benchmark          - Run performance benchmarks"
	@echo "  examples           - Run example scripts"
	@echo ""
	@echo "Development Commands:"
	@echo "  format             - Format Go code"
	@echo "  vet                - Check code quality"
	@echo "  dev-setup          - Setup development environment"
	@echo ""
	@echo "Utility Commands:"
	@echo "  clean              - Clean all generated files"
	@echo "  clean-ips          - Clean only IP files"
	@echo "  install            - Install binaries to system"
	@echo "  uninstall          - Remove installed binaries"
	@echo "  release            - Prepare release"
	@echo "  help               - Show this help"
	@echo ""
	@echo "Example Usage:"
	@echo "  make build && ./bin/ultra_ip_gen 100000"
	@echo "  make test"
	@echo "  make benchmark"

