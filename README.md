# and-agy

Antigravity CLI-style agentic coding tool for Android/Termux.

## Overview

and-agy is a Termux-native Go application that provides an Antigravity-style agentic coding experience on Android devices. It features a Bubble Tea TUI, multi-provider AI support, and secure authentication.

## Features

- **Multi-Provider Support**: Google (Gemini), OpenAI, Anthropic, DeepSeek, xAI, Ollama, LM Studio, llama.cpp
- **Bubble Tea TUI**: Interactive terminal interface
- **Pure Go**: No CGO, works on aarch64
- **Secure Auth**: Provider management with remote auth via .well-known/opencode
- **RCE Protection**: Explicit user confirmation for remote code execution
- **Workspace Sandboxing**: Isolated execution environment

## Installation

### Termux

```bash
pkg update && pkg upgrade
pkg install git golang

# Clone and build
git clone https://github.com/Chieji/and-agy.git
cd and-agy
./scripts/install-termux.sh
```

### Configuration

Set your API key:
```bash
export GEMINI_API_KEY="your-api-key"
```

Or use the auth system:
```bash
agy auth login gemini
```

## Usage

```bash
agy chat       # Start interactive chat
agy agent      # Run agentic coding tasks
agy config     # Configure settings
agy auth       # Manage authentication
```

## Architecture

- `cmd/agy`: Main entry point
- `internal/agent`: Agent harness, planner, job manager
- `internal/auth`: Authentication providers
- `internal/cli`: CLI commands
- `internal/tools`: Filesystem and shell tools
- `internal/workspace`: Sandbox environment
- `pkg/tui`: Bubble Tea UI components

## License

MIT
