# Contributing to Ultra-Fast IP Generator

First off, thank you for considering contributing to Ultra-Fast IP Generator! 🎉

The following is a set of guidelines for contributing to this project. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

This section guides you through submitting a bug report. Following these guidelines helps maintainers and the community understand your report, reproduce the behavior, and find related reports.

**Before Submitting A Bug Report:**
- Check the [existing issues](https://github.com/yourusername/ultra-fast-ip-generator/issues)
- Perform a cursory search to see if the problem has already been reported

**How Do I Submit A (Good) Bug Report?**

Bugs are tracked as [GitHub issues](https://github.com/yourusername/ultra-fast-ip-generator/issues). Create an issue and provide the following information:

- **Use a clear and descriptive title**
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples** to demonstrate the steps
- **Describe the behavior you observed** and what behavior you expected
- **Include details about your configuration and environment**:
  - Go version (`go version`)
  - Operating system and version
  - Hardware specifications (RAM, CPU)
  - Command used and parameters

### Suggesting Enhancements

This section guides you through submitting an enhancement suggestion, including completely new features and minor improvements to existing functionality.

**Before Submitting An Enhancement Suggestion:**
- Check if the enhancement has already been suggested
- Consider whether your idea fits with the scope and aims of the project

**How Do I Submit A (Good) Enhancement Suggestion?**

- **Use a clear and descriptive title**
- **Provide a step-by-step description** of the suggested enhancement
- **Provide specific examples** to demonstrate how the enhancement would work
- **Describe the current behavior** and **explain which behavior you expected**
- **Explain why this enhancement would be useful**

### Pull Requests

The process described here has several goals:
- Maintain the project's quality
- Fix problems that are important to users
- Engage the community in working toward the best possible implementation
- Enable a sustainable system for maintainers to review contributions

**Before Submitting a Pull Request:**
1. Fork the repository
2. Create a new branch from `main`
3. Make your changes
4. Test your changes thoroughly
5. Update documentation if necessary
6. Ensure your code follows the project's style guidelines

## Development Setup

### Prerequisites
- Go 1.19 or higher
- Git

### Setting Up Your Development Environment

1. **Fork the repository**
   ```bash
   # On GitHub, click the "Fork" button
   ```

2. **Clone your fork**
   ```bash
   git clone https://github.com/your-username/ultra-fast-ip-generator.git
   cd ultra-fast-ip-generator
   ```

3. **Add the upstream remote**
   ```bash
   git remote add upstream https://github.com/original-owner/ultra-fast-ip-generator.git
   ```

4. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

### Running Tests

```bash
# Run basic functionality tests
go run ultra_fast_ip_generator.go 1000
go run fast_ip_generator.go 1000  
go run ip_generator.go 1000

# Test compilation
go build ultra_fast_ip_generator.go
go build fast_ip_generator.go
go build ip_generator.go

# Format code
go fmt ./...

# Check for issues
go vet ./...
```

### Code Style

- Follow standard Go formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions focused and small
- Use error handling appropriately

### Performance Considerations

- Benchmark any performance-critical changes
- Consider memory usage for large datasets
- Test with various input sizes (1K, 100K, 1M, 10M+ IPs)
- Profile memory and CPU usage when applicable

## Git Workflow

### Branch Naming
- `feature/description` - for new features
- `bugfix/description` - for bug fixes
- `docs/description` - for documentation updates
- `performance/description` - for performance improvements

### Commit Messages

Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
type(scope): description

[optional body]

[optional footer]
```

Examples:
```
feat: add IPv6 support to ultra-fast generator
fix: resolve memory leak in batch processing
docs: update performance benchmarks in README
perf: optimize IP validation algorithm
```

### Pull Request Process

1. **Update documentation** if your changes affect user-facing functionality
2. **Add tests** if you're adding new functionality
3. **Ensure all tests pass** and code compiles successfully
4. **Update CHANGELOG.md** with details of changes
5. **Request review** from maintainers
6. **Address feedback** and make necessary changes
7. **Squash commits** if requested before merging

## Performance Benchmarking

When contributing performance improvements:

1. **Benchmark before and after**
   ```bash
   # Example benchmarking
   time go run ultra_fast_ip_generator.go 1000000
   ```

2. **Test with various sizes**
   - Small: 1K - 10K IPs
   - Medium: 100K - 1M IPs  
   - Large: 10M+ IPs

3. **Monitor resource usage**
   ```bash
   # Monitor memory usage
   /usr/bin/time -v go run ultra_fast_ip_generator.go 1000000
   ```

4. **Document improvements** in pull request description

## Documentation Guidelines

- Keep README.md up to date
- Include code examples for new features
- Update performance benchmarks if applicable
- Write clear, concise explanations
- Use proper markdown formatting

## Questions?

Don't hesitate to ask questions:

- Open an [issue](https://github.com/yourusername/ultra-fast-ip-generator/issues)
- Check existing [discussions](https://github.com/yourusername/ultra-fast-ip-generator/discussions)

## Recognition

Contributors will be recognized in:
- README.md contributors section
- CHANGELOG.md for significant contributions
- GitHub contributor graphs

Thank you for contributing! 🚀

