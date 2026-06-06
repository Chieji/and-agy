# Build Instructions for and-agy

## 1. Prerequisites

### 1.1 Termux
```bash
pkg update && pkg upgrade
pkg install git golang
```

### 1.2 Linux
```bash
# Debian/Ubuntu
sudo apt update && sudo apt upgrade
sudo apt install git golang

# Fedora/RHEL
sudo dnf update && sudo dnf upgrade
sudo dnf install git golang

# Arch
sudo pacman -Syu git go
```

## 2. Build Configuration

### 2.1 Environment Variables
```bash
# Required for pure Go build
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=arm64

# Optional
export GOPROXY=https://proxy.golang.org,direct
```

### 2.2 Go Version
- Minimum: Go 1.21
- Recommended: Latest stable

## 3. Build Commands

### 3.1 Development Build
```bash
# Clone repository
git clone https://github.com/Chieji/and-agy.git
cd and-agy

# Build with debugging
go build -o agy .
```

### 3.2 Release Build
```bash
# Optimized build for Termux
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
go build -ldflags "-s -w" -o agy .

# With version information
VERSION=$(git describe --tags --always --dirty)
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
go build -ldflags "-s -w -X main.Version=$VERSION" -o agy .
```

### 3.3 Cross-Platform Builds

#### Android (aarch64)
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o agy-android .
```

#### Linux (x86_64)
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o agy-linux .
```

#### Windows
```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o agy.exe .
```

#### macOS
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o agy-macos .
```

## 4. Installation

### 4.1 Termux
```bash
# Build
./scripts/install-termux.sh

# Or manual installation
cp agy $PREFIX/bin/agy
chmod +x $PREFIX/bin/agy
```

### 4.2 System-wide (Linux)
```bash
sudo cp agy /usr/local/bin/agy
sudo chmod +x /usr/local/bin/agy
```

### 4.3 User-local (Linux)
```bash
mkdir -p ~/.local/bin
cp agy ~/.local/bin/agy
chmod +x ~/.local/bin/agy
export PATH=$PATH:~/.local/bin
```

## 5. Dependencies

### 5.1 Go Modules
```bash
# Download dependencies
go mod download

# Update dependencies
go get -u ./...

# Clean dependencies
go clean -modcache
```

### 5.2 Static Linking
The build is configured for static linking where possible. Some dependencies may require dynamic linking on certain platforms.

## 6. Build Scripts

### 6.1 install-termux.sh
Automated Termux installation:
- Installs dependencies
- Builds the binary
- Installs to $PREFIX/bin
- Creates desktop entry (optional)

### 6.2 setup-proot.sh
For proot-based environments:
- Sets up proot environment
- Installs required packages
- Configures build environment

## 7. Build Troubleshooting

### 7.1 Common Issues

#### CGO_ENABLED=0 errors
```bash
# Ensure pure Go dependencies
go mod tidy
# Check for CGO dependencies
go list -m -json all | grep -i cgo
```

#### Missing dependencies
```bash
# Install missing system packages
# Termux
pkg install <package>

# Debian/Ubuntu
sudo apt install <package>
```

#### Build failures
```bash
# Clean build
go clean
go build -o agy .

# Verbose build
go build -v -o agy .
```

### 7.2 Platform-Specific Issues

#### Termux
- Ensure Termux is up to date
- Check storage permissions
- Verify Go installation: `go version`

#### Linux
- Check architecture: `uname -m`
- Verify Go supports your architecture
- Check for missing system libraries

## 8. Build Verification

### 8.1 Binary Check
```bash
# Check binary type
file agy

# Check architecture
go tool nm agy | grep main
```

### 8.2 Runtime Check
```bash
# Version
agy version

# Help
agy help

# Quick test
agy chat --help
```

## 9. CI/CD Builds

### 9.1 GitHub Actions
```yaml
name: Build

on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    - run: go build -o agy .
```

### 9.2 Termux CI
```yaml
name: Termux Build

on: [push, pull_request]

jobs:
  termux:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    - run: CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o agy .
```

## 10. Release Builds

### 10.1 Version Tagging
```bash
# Create version tag
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

### 10.2 Release Binaries
```bash
# Build all platforms
./scripts/build-release.sh

# Creates:
# - agy-linux-arm64
# - agy-linux-amd64
# - agy-windows-amd64.exe
# - agy-macos-arm64
```

### 10.3 Release Checklist
- [ ] All tests pass
- [ ] Documentation updated
- [ ] Changelog updated
- [ ] Version bumped
- [ ] Binaries built
- [ ] Release notes prepared