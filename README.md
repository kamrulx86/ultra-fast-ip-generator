# 🚀 Ultra-Fast IP Generator

[![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Performance](https://img.shields.io/badge/Performance-3.9M+%20IPs/sec-red.svg)](#performance)

A blazing-fast Go application that generates millions of random public IP addresses with incredible performance. Perfect for network testing, security research, and load testing scenarios.

## ⚡ Key Features

- **Ultra-Fast Generation**: 3.9+ million IPs per second
- **Public IPs Only**: Automatically excludes private, reserved, and multicast ranges
- **Duplicate Handling**: Optional unique mode with efficient deduplication
- **Batch Processing**: Optimized memory usage with batch generation
- **Multiple Algorithms**: Choose between speed and uniqueness
- **Progress Tracking**: Real-time progress updates for large datasets
- **File Splitting**: Built-in functionality to split large files into manageable chunks

## 🏃‍♂️ Quick Start

### Prerequisites
- Go 1.19 or higher
- Linux/macOS/Windows

### Installation

```bash
git clone https://github.com/yourusername/ultra-fast-ip-generator.git
cd ultra-fast-ip-generator
```

### Basic Usage

```bash
# Generate 500K IPs (ultra-fast, may have duplicates)
go run ultra_fast_ip_generator.go 500000

# Generate 500K unique IPs (slightly slower)
go run ultra_fast_ip_generator.go 500000 unique

# Generate 1M IPs with fast algorithm
go run fast_ip_generator.go 1000000

# Generate IPs with complete validation (slowest, guaranteed unique)
go run ip_generator.go 100000
```

## 📊 Performance Benchmarks

| Algorithm | Speed (IPs/sec) | Duplicates | Use Case |
|-----------|----------------|------------|----------|
| Ultra-Fast | 3.9M+ | Possible | Maximum speed, large datasets |
| Fast | 3.4M+ | Possible | Balanced speed/efficiency |
| Validated | 1.5M+ | None | Complete validation, guaranteed unique |
| Unique Mode | 2.5M+ | None | Fast + guaranteed unique |

### Real Performance Results

```
🚀 20 Million IPs Generated in 5.13 seconds!
📈 Rate: 3,896,646 IPs/second
📁 File Size: 286.10 MB
✅ Memory Efficient: Batch processing
```

## 🛠️ Available Generators

### 1. Ultra-Fast Generator (`ultra_fast_ip_generator.go`)
**Best for**: Maximum speed, large datasets (1M+ IPs)

```bash
# Basic usage
go run ultra_fast_ip_generator.go 1000000

# With unique constraint
go run ultra_fast_ip_generator.go 1000000 unique
```

**Features**:
- Pre-selected valid IP ranges
- Batch processing (50K batches)
- 1MB file buffer
- Optional duplicate removal

### 2. Fast Generator (`fast_ip_generator.go`)
**Best for**: Balanced performance and efficiency

```bash
go run fast_ip_generator.go 500000
```

**Features**:
- Smart range selection
- Optimized private IP avoidance
- Progress tracking every 5%

### 3. Validated Generator (`ip_generator.go`)
**Best for**: Complete validation, guaranteed unique IPs

```bash
go run ip_generator.go 100000
```

**Features**:
- Complete RFC compliance
- Guaranteed unique IPs
- Full private IP range validation

## 📁 File Management

### Automatic File Naming
Generated files are automatically named with timestamps:
```
ultra_fast_ips_20250614_085109.txt
fast_generated_ips_20250614_084737.txt
```

### Duplicate Removal
Remove duplicates from existing files:
```bash
sort input_file.txt | uniq > output_unique.txt
```

### File Splitting
Split large IP files into manageable chunks:
```bash
mkdir ip_chunks_100k
split -l 100000 -d input_file.txt ip_chunks_100k/ips_chunk_
```

## 🎯 Use Cases

- **Network Security Testing**: Generate large IP datasets for penetration testing
- **Load Testing**: Create realistic IP distributions for load testing
- **Research**: Academic research requiring large IP datasets
- **Simulation**: Network simulation and modeling
- **Development**: Testing IP-based applications

## 🔧 Advanced Usage

### Environment Variables
```bash
export IP_BATCH_SIZE=50000  # Batch size for generation
export IP_BUFFER_SIZE=1048576  # File buffer size (1MB)
```

### Custom Compilation
```bash
# Optimize for your system
go build -ldflags="-s -w" -o ultra_ip_gen ultra_fast_ip_generator.go

# Cross-compile for different platforms
GOOS=windows GOARCH=amd64 go build ultra_fast_ip_generator.go
GOOS=darwin GOARCH=amd64 go build ultra_fast_ip_generator.go
```

## 📈 Scaling Examples

### Memory Usage
| IPs Generated | Memory Usage | Time |
|---------------|--------------|------|
| 100K | ~15 MB | 27ms |
| 1M | ~150 MB | 296ms |
| 10M | ~1.5 GB | 2.8s |
| 20M | ~3.0 GB | 5.1s |

### Disk Space
- **100K IPs**: ~1.4 MB
- **1M IPs**: ~14 MB  
- **10M IPs**: ~140 MB
- **20M IPs**: ~286 MB

## 🛡️ IP Range Exclusions

The generators automatically exclude:

- **Private Networks**: 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16
- **Loopback**: 127.0.0.0/8
- **Link-Local**: 169.254.0.0/16
- **Multicast**: 224.0.0.0/4
- **Reserved**: 240.0.0.0/4
- **Invalid**: 0.0.0.0/8, 255.255.255.255

## 🐛 Troubleshooting

### Common Issues

**Out of Memory**:
```bash
# Reduce batch size or use streaming
go run ultra_fast_ip_generator.go 100000  # Start smaller
```

**Slow Performance**:
```bash
# Check available RAM
free -h

# Use ultra-fast mode without unique constraint
go run ultra_fast_ip_generator.go 1000000
```

**File Permission Issues**:
```bash
# Ensure write permissions
chmod 755 .
ls -la *.txt
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Setup
```bash
git clone https://github.com/yourusername/ultra-fast-ip-generator.git
cd ultra-fast-ip-generator
go mod init ultra-fast-ip-generator
go mod tidy
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🏆 Achievements

- ⚡ **3.9M+ IPs/second** generation speed
- 🎯 **20M IPs** generated in 5.13 seconds
- 💾 **Memory efficient** batch processing
- 🔧 **Cross-platform** compatibility
- ✅ **Production ready** with comprehensive validation

## 📧 Support

If you encounter any issues or have questions:

1. Check the [Issues](https://github.com/yourusername/ultra-fast-ip-generator/issues) page
2. Create a new issue with detailed description
3. Include your Go version and OS information

## 🌟 Star History

If this project helped you, please consider giving it a star! ⭐

---

**Made with ❤️ for the networking community**

