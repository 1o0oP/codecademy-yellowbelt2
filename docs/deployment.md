# 🚀 Deploy & Build

Guia completo para build, distribuição e configuração do Todo List CLI em diferentes ambientes.

## 📦 Build Process

### Compilação Local

#### Usando Makefile (Recomendado)
```bash
# Build completo
make build

# Com limpeza prévia
make clean && make build

# Verificar resultado
ls -la bin/
# -rwxr-xr-x 1 user user 8.2M Aug 25 14:30 todo
```

#### Comandos Go Diretos
```bash
# Build simples
go build -o bin/todo .

# Build com otimizações
go build -ldflags="-s -w" -o bin/todo .

# Build com informações de versão
go build -ldflags="-X main.version=1.0.0 -X main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)" -o bin/todo .
```

### Cross-Platform Build

#### Targets Suportados
```bash
# Linux x64
GOOS=linux GOARCH=amd64 go build -o bin/todo-linux-amd64 .

# macOS x64
GOOS=darwin GOARCH=amd64 go build -o bin/todo-darwin-amd64 .

# macOS ARM (M1/M2)
GOOS=darwin GOARCH=arm64 go build -o bin/todo-darwin-arm64 .

# Windows x64
GOOS=windows GOARCH=amd64 go build -o bin/todo-windows-amd64.exe .

# Linux ARM (Raspberry Pi)
GOOS=linux GOARCH=arm64 go build -o bin/todo-linux-arm64 .
```

#### Script de Build Multi-Platform
```bash
#!/bin/bash
# build-all.sh

platforms=(
    "linux/amd64"
    "linux/arm64" 
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
)

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    
    output_name="todo-$GOOS-$GOARCH"
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    
    echo "Building for $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o bin/$output_name .
done
```

## 🔧 Build Otimizações

### Flags de Compilação

#### Redução de Tamanho
```bash
# Flags de otimização
-ldflags="-s -w"
# -s: strip symbol table
# -w: strip DWARF debug info

# Resultado típico:
# Sem flags: 8.2M
# Com flags: 6.1M (~25% menor)
```

#### Build Tags
```bash
# Build de produção (desabilita debug)
go build -tags production -o bin/todo .

# Build de desenvolvimento (habilita debug)
go build -tags debug -o bin/todo .
```

#### Exemplo no Código
```go
//go:build production
// +build production

package main

const DebugMode = false

//go:build debug
// +build debug  

package main

const DebugMode = true
```

### Compressão Adicional

#### UPX (Ultimate Packer for eXecutables)
```bash
# Instalar UPX
sudo apt install upx-ucl  # Ubuntu/Debian
brew install upx         # macOS

# Comprimir binário
upx --best bin/todo

# Resultado típico:
# Original: 6.1M
# UPX:      2.3M (~62% menor)
```

## 🎯 Build para Produção

### Makefile Completo
```makefile
# Variables
APP_NAME=todo
BUILD_DIR=bin
VERSION?=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS=-ldflags="-s -w -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

# Production build
build-prod: clean
	@echo "🏗️  Building production binary..."
	CGO_ENABLED=0 go build $(LDFLAGS) -tags production -o $(BUILD_DIR)/$(APP_NAME) .
	@echo "📦 Binary size: $$(du -h $(BUILD_DIR)/$(APP_NAME) | cut -f1)"

# Multi-platform release
build-release: clean
	@echo "🌍 Building for all platforms..."
	@./scripts/build-all.sh
	@echo "📦 Release binaries:"
	@ls -la $(BUILD_DIR)/

# Docker build
build-docker:
	docker build -t todo-cli:$(VERSION) .

# Clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

.PHONY: build-prod build-release build-docker clean
```

### Informações de Versão

#### Injeção durante Build
```go
// main.go
package main

var (
    version   = "dev"
    buildTime = "unknown"
    gitCommit = "unknown"
)

func init() {
    if version != "dev" {
        fmt.Printf("Todo CLI v%s (built %s)\n", version, buildTime)
    }
}
```

#### Comando Version
```go
// CLI command
func versionCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "version",
        Short: "Show version information",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Todo CLI v%s\n", version)
            fmt.Printf("Build time: %s\n", buildTime)
            fmt.Printf("Git commit: %s\n", gitCommit)
            fmt.Printf("Go version: %s\n", runtime.Version())
            fmt.Printf("Platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
        },
    }
}
```

## 🐳 Containerização

### Dockerfile
```dockerfile
# Build stage
FROM golang:1.23.5-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o todo .

# Production stage  
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/todo .

# Create data directory
RUN mkdir -p .todo-cli

ENTRYPOINT ["./todo"]
```

### Docker Compose
```yaml
# docker-compose.yml
version: '3.8'

services:
  todo-cli:
    build: .
    image: todo-cli:latest
    volumes:
      # Mount data directory for persistence
      - ./data:/root/.todo-cli
    environment:
      - TODO_DATA_DIR=/root/.todo-cli
    # Interactive mode for CLI usage
    stdin_open: true
    tty: true
```

### Comandos Docker
```bash
# Build image
docker build -t todo-cli:latest .

# Run interactively
docker run -it --rm -v $(pwd)/data:/root/.todo-cli todo-cli:latest

# Create alias for easy usage
alias dtodo='docker run -it --rm -v $(pwd)/data:/root/.todo-cli todo-cli:latest'

# Usage with alias
dtodo create "Docker task" "Running in container"
dtodo list
```

## 📱 Instalação do Sistema

### Unix/Linux (Método Global)

#### Script de Instalação
```bash
#!/bin/bash
# install.sh

set -e

# Configuration
BINARY_NAME="todo"
INSTALL_DIR="/usr/local/bin"
VERSION="latest"

# Download binary
echo "📥 Downloading Todo CLI..."
if [[ "$OSTYPE" == "darwin"* ]]; then
    PLATFORM="darwin-amd64"
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
    PLATFORM="linux-amd64" 
else
    echo "❌ Unsupported platform: $OSTYPE"
    exit 1
fi

# Create temp directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

# Download (would be from GitHub releases)
# curl -sL "https://github.com/user/repo/releases/download/v${VERSION}/todo-${PLATFORM}" -o todo

# For local installation, copy from build
cp /path/to/project/bin/todo ./

# Make executable
chmod +x todo

# Install to system
echo "🔧 Installing to $INSTALL_DIR..."
sudo mv todo "$INSTALL_DIR/$BINARY_NAME"

# Verify installation
if command -v "$BINARY_NAME" &> /dev/null; then
    echo "✅ Todo CLI installed successfully!"
    echo "📍 Location: $(which $BINARY_NAME)"
    echo "🎯 Try: $BINARY_NAME --help"
else
    echo "❌ Installation failed"
    exit 1
fi

# Cleanup
cd /
rm -rf "$TMP_DIR"
```

#### Manual Installation
```bash
# Build locally
make build

# Install globally
sudo cp bin/todo /usr/local/bin/

# Verify
which todo
todo --help
```

### Windows

#### PowerShell Script
```powershell
# install.ps1

$BinaryName = "todo.exe"
$InstallDir = "$env:USERPROFILE\bin"
$BinaryPath = "$InstallDir\$BinaryName"

# Create install directory
New-Item -ItemType Directory -Force -Path $InstallDir

# Copy binary (assuming built locally)
Copy-Item "bin\todo-windows-amd64.exe" $BinaryPath

# Add to PATH if not already there
$CurrentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
if ($CurrentPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable("PATH", "$CurrentPath;$InstallDir", "User")
    Write-Host "✅ Added $InstallDir to PATH"
    Write-Host "⚠️  Restart your terminal for PATH changes to take effect"
}

Write-Host "✅ Todo CLI installed successfully!"
Write-Host "📍 Location: $BinaryPath"
```

### Package Managers

#### Homebrew (macOS/Linux)
```ruby
# todo.rb (Homebrew formula)
class Todo < Formula
  desc "Todo List CLI - Task management tool"
  homepage "https://github.com/user/codecademy-yellowbelt2"
  url "https://github.com/user/codecademy-yellowbelt2/archive/v1.0.0.tar.gz"
  sha256 "..."
  license "MIT"

  depends_on "go" => :build

  def install
    cd "app" do
      system "go", "build", "-ldflags", "-s -w", "-o", "todo", "."
      bin.install "todo"
    end
  end

  test do
    output = shell_output("#{bin}/todo --help")
    assert_match "Todo List CLI", output
  end
end
```

Installation:
```bash
# Add tap
brew tap user/tap

# Install
brew install todo

# Usage
todo --help
```

#### Debian Package
```bash
# Build .deb package
mkdir -p todo-cli-1.0.0/DEBIAN
mkdir -p todo-cli-1.0.0/usr/local/bin

# Control file
cat > todo-cli-1.0.0/DEBIAN/control << EOF
Package: todo-cli
Version: 1.0.0
Architecture: amd64
Maintainer: Your Name <email@example.com>
Description: Todo List CLI - Task management tool
 A command-line tool for managing personal tasks
 built with Go and Clean Architecture.
EOF

# Copy binary
cp bin/todo todo-cli-1.0.0/usr/local/bin/

# Build package
dpkg-deb --build todo-cli-1.0.0

# Install
sudo dpkg -i todo-cli-1.0.0.deb
```

## ⚙️ Configuração de Runtime

### Variáveis de Ambiente

#### Customização de Dados
```go
// main.go - Enhanced data directory logic
func getDataDirectory() string {
    // 1. Environment variable override
    if dataDir := os.Getenv("TODO_DATA_DIR"); dataDir != "" {
        return dataDir
    }
    
    // 2. XDG Base Directory (Linux)
    if xdgData := os.Getenv("XDG_DATA_HOME"); xdgData != "" {
        return filepath.Join(xdgData, "todo-cli")
    }
    
    // 3. Default: user home
    homeDir, err := os.UserHomeDir()
    if err != nil {
        log.Fatal("Cannot determine home directory:", err)
    }
    return filepath.Join(homeDir, ".todo-cli")
}
```

#### Configuração Avançada
```bash
# Custom data directory
export TODO_DATA_DIR="/custom/path/todo-data"

# JSON formatting (pretty/compact)
export TODO_JSON_FORMAT="pretty"

# Debug mode
export TODO_DEBUG="true"

# Custom date format
export TODO_DATE_FORMAT="2006-01-02 15:04:05"
```

### Arquivo de Configuração

#### YAML Config Support
```yaml
# ~/.config/todo-cli/config.yaml
data:
  directory: "~/.todo-cli"
  filename: "todos.json"
  backup: true
  
display:
  date_format: "02/01/2006 15:04"
  use_emoji: true
  compact_list: false
  
performance:
  cache_enabled: true
  cache_ttl: "5m"
```

#### Carregamento de Config
```go
type Config struct {
    Data struct {
        Directory string `yaml:"directory"`
        Filename  string `yaml:"filename"`
        Backup    bool   `yaml:"backup"`
    } `yaml:"data"`
    Display struct {
        DateFormat  string `yaml:"date_format"`
        UseEmoji    bool   `yaml:"use_emoji"`
        CompactList bool   `yaml:"compact_list"`
    } `yaml:"display"`
}

func LoadConfig() (*Config, error) {
    configPath := filepath.Join(os.Getenv("HOME"), ".config", "todo-cli", "config.yaml")
    
    data, err := os.ReadFile(configPath)
    if err != nil {
        return DefaultConfig(), nil // Fallback to defaults
    }
    
    var config Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        return nil, err
    }
    
    return &config, nil
}
```

## 🔄 CI/CD Pipeline

### GitHub Actions
```yaml
# .github/workflows/build.yml
name: Build and Release

on:
  push:
    tags: [ 'v*' ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.23.5'
    
    - name: Run tests
      run: |
        cd app
        go test -race -coverprofile=coverage.out ./...
        
    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        file: app/coverage.out

  build:
    needs: test
    runs-on: ubuntu-latest
    strategy:
      matrix:
        goos: [linux, darwin, windows]
        goarch: [amd64, arm64]
        exclude:
          - goos: windows
            goarch: arm64
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.23.5'
    
    - name: Build
      env:
        GOOS: ${{ matrix.goos }}
        GOARCH: ${{ matrix.goarch }}
      run: |
        cd app
        output_name="todo-${GOOS}-${GOARCH}"
        if [ "$GOOS" = "windows" ]; then
          output_name="${output_name}.exe"
        fi
        
        go build -ldflags="-s -w -X main.version=${GITHUB_REF#refs/tags/}" \
                 -o "bin/${output_name}" .
    
    - name: Upload artifacts
      uses: actions/upload-artifact@v3
      with:
        name: binaries
        path: app/bin/

  release:
    if: startsWith(github.ref, 'refs/tags/')
    needs: build
    runs-on: ubuntu-latest
    steps:
    - name: Download artifacts
      uses: actions/download-artifact@v3
      with:
        name: binaries
        path: bin/
    
    - name: Create Release
      uses: softprops/action-gh-release@v1
      with:
        files: bin/*
        generate_release_notes: true
```

### GitLab CI
```yaml
# .gitlab-ci.yml
stages:
  - test
  - build
  - release

variables:
  GO_VERSION: "1.23.5"

test:
  stage: test
  image: golang:${GO_VERSION}
  script:
    - cd app
    - go test -race -coverprofile=coverage.out ./...
  coverage: '/coverage: \d+\.\d+% of statements/'
  artifacts:
    reports:
      coverage_report:
        coverage_format: cobertura
        path: app/coverage.xml

build:
  stage: build
  image: golang:${GO_VERSION}
  parallel:
    matrix:
      - GOOS: [linux, darwin, windows]
        GOARCH: [amd64, arm64]
  script:
    - cd app
    - |
      output_name="todo-${GOOS}-${GOARCH}"
      if [ "$GOOS" = "windows" ]; then
        output_name="${output_name}.exe"
      fi
      
      CGO_ENABLED=0 go build \
        -ldflags="-s -w -X main.version=${CI_COMMIT_TAG}" \
        -o "bin/${output_name}" .
  artifacts:
    paths:
      - app/bin/
    expire_in: 1 week

release:
  stage: release
  image: registry.gitlab.com/gitlab-org/release-cli:latest
  rules:
    - if: $CI_COMMIT_TAG
  script:
    - echo "Creating release for ${CI_COMMIT_TAG}"
  release:
    tag_name: '$CI_COMMIT_TAG'
    description: 'Release $CI_COMMIT_TAG'
    assets:
      links:
        - name: 'Linux x64'
          url: '${CI_API_V4_URL}/projects/${CI_PROJECT_ID}/jobs/artifacts/${CI_COMMIT_TAG}/raw/app/bin/todo-linux-amd64?job=build'
        - name: 'macOS x64'
          url: '${CI_API_V4_URL}/projects/${CI_PROJECT_ID}/jobs/artifacts/${CI_COMMIT_TAG}/raw/app/bin/todo-darwin-amd64?job=build'
```

## 📊 Monitoring & Observability

### Health Check Endpoint
```go
// Para uso futuro em versão web/API
func healthCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "health",
        Short: "Check system health",
        Run: func(cmd *cobra.Command, args []string) {
            // Check data directory
            dataDir := getDataDirectory()
            if _, err := os.Stat(dataDir); os.IsNotExist(err) {
                fmt.Println("❌ Data directory not accessible")
                os.Exit(1)
            }
            
            // Check data file
            dataFile := filepath.Join(dataDir, "todos.json")
            if _, err := os.Stat(dataFile); err == nil {
                // File exists, try to parse
                repo := NewFileTodoRepository(dataFile)
                if _, err := repo.GetAll(); err != nil {
                    fmt.Println("❌ Data file corrupted")
                    os.Exit(1)
                }
            }
            
            fmt.Println("✅ System healthy")
            fmt.Printf("📁 Data directory: %s\n", dataDir)
            fmt.Printf("📄 Data file: %s\n", dataFile)
        },
    }
}
```

### Metrics Collection
```go
type Metrics struct {
    TotalTodos     int
    CompletedTodos int
    PendingTodos   int
    LastAccess     time.Time
}

func (m *Metrics) ToJSON() string {
    data, _ := json.MarshalIndent(m, "", "  ")
    return string(data)
}

func metricsCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "metrics",
        Short: "Show usage metrics",
        Run: func(cmd *cobra.Command, args []string) {
            // Calculate metrics
            // Output in JSON or human-readable format
        },
    }
}
```

## 🔧 Troubleshooting

### Common Issues

#### Permission Denied
```bash
# Problem: Can't create data directory
# Solution: Check permissions
ls -la ~/.todo-cli
chmod 755 ~/.todo-cli
```

#### Binary Not Found
```bash
# Problem: Command not found after installation
# Solution: Check PATH
echo $PATH
which todo

# Add to PATH if needed
export PATH="$PATH:/usr/local/bin"
```

#### JSON Corruption
```bash
# Problem: Invalid JSON format
# Solution: Backup and reset
cp ~/.todo-cli/todos.json ~/.todo-cli/todos.json.backup
echo '{}' > ~/.todo-cli/todos.json
```

### Debug Mode
```go
// Enable with environment variable
if os.Getenv("TODO_DEBUG") == "true" {
    log.SetLevel(log.DebugLevel)
    log.Debug("Debug mode enabled")
}
```

---

## 🎯 Quick Deploy Checklist

### ✅ Pre-Deploy
- [ ] Tests passing (`make test`)
- [ ] Code coverage >85%
- [ ] Cross-platform build successful
- [ ] Version information injected
- [ ] Documentation updated

### ✅ Deploy Process
- [ ] Tag release (`git tag v1.0.0`)
- [ ] Build artifacts generated
- [ ] Binaries tested on target platforms
- [ ] Release notes prepared
- [ ] Distribution channels updated

### ✅ Post-Deploy
- [ ] Installation scripts tested
- [ ] User documentation verified
- [ ] Health checks passing
- [ ] Monitoring active
- [ ] Feedback mechanism ready

---

**🚀 Com esta configuração, o Todo List CLI está pronto para ser distribuído e usado em produção em qualquer ambiente!**
