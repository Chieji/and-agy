# and-agy Specification

## 1. Overview

Antigravity CLI-style agentic coding tool for Android/Termux, built in pure Go with Bubble Tea TUI and multi-provider AI support.

## 2. Requirements

### 2.1 Platform
- **Target**: Android/Termux
- **Architecture**: aarch64
- **Build**: CGO_ENABLED=0, pure Go
- **Dependencies**: Minimal, statically linked where possible

### 2.2 Core Features
- Multi-provider AI support (Gemini, OpenAI, Anthropic, DeepSeek, xAI, Ollama, LM Studio, llama.cpp)
- Bubble Tea-based TUI
- Agentic coding capabilities
- Secure authentication system
- Workspace sandboxing
- RCE protection

## 3. Architecture

### 3.1 Component Diagram

```
+------------------+
|     CLI          |
+------------------+
        |
        v
+------------------+
|    Root CMD      |
+------------------+
        |
   +----+----+------+
   |         |      |
   v         v      v
+------+ +------+ +------+
| Chat | |Agent | |Config|
+------+ +------+ +------+
        |
        v
+------------------+
|     Agent        |
+------------------+
   +----+----+------+
   |         |      |
   v         v      v
+------+ +------+ +------+
|Harness| |Planner| |JobMgr |
+------+ +------+ +------+
        |
        v
+------------------+
|     Auth         |
+------------------+
   +----+----+------+
   |    |    |      |
   v    v    v      v
+----+----+----+----+
|Gemi|OAI |Anth|... |
+----+----+----+----+
        |
        v
+------------------+
|     Tools        |
+------------------+
   +----+----+
   |    |    |
   v    v    v
+----+----+----+
| FS |Shell|... |
+----+----+----+
        |
        v
+------------------+
|    Workspace     |
+------------------+
        |
        v
+------------------+
|      TUI         |
+------------------+
   +----+----+------+
   |         |      |
   v         v      v
+------+ +------+ +------+
| App  | |Theme | |Views |
+------+ +------+ +------+
```

### 3.2 Provider Interface

```go
type Provider interface {
    Name() string
    Authenticate(ctx context.Context, config map[string]string) error
    Chat(ctx context.Context, messages []Message) (*Message, error)
    StreamChat(ctx context.Context, messages []Message) (<-chan Message, error)
    Models() []Model
}
```

### 3.3 Auth System

- **Local**: GEMINI_API_KEY env var
- **Remote**: .well-known/opencode endpoint
- **Providers**: Google, OpenAI, Anthropic, DeepSeek, xAI
- **Local LLM**: Ollama, LM Studio, llama.cpp
- **Management**: auth login, auth list, auth logout, auth default
- **RCE Protection**: Explicit user confirmation for remote code execution

## 4. CLI Commands

### 4.1 Root
- `agy version`
- `agy help`

### 4.2 Chat
- `agy chat` - Start interactive chat
- `agy chat --provider <name>` - Use specific provider
- `agy chat --model <model>` - Use specific model

### 4.3 Agent
- `agy agent` - Start agent mode
- `agy agent --task <description>` - Run specific task
- `agy agent --workspace <path>` - Use specific workspace

### 4.4 Config
- `agy config get <key>`
- `agy config set <key> <value>`
- `agy config list`
- `agy config edit`

### 4.5 Auth
- `agy auth login <provider>` - Authenticate with provider
- `agy auth list` - List configured providers
- `agy auth logout <provider>` - Remove provider authentication
- `agy auth default <provider>` - Set default provider
- `agy auth connect` - TUI for remote auth

## 5. TUI Views

- **Chat View**: Interactive chat interface
- **Config View**: Configuration editor
- **Input View**: User input prompt
- **Jobs View**: Background job management
- **Status View**: System status and logs

## 6. Agent Components

### 6.1 Harness
- Manages agent lifecycle
- Handles tool execution
- Coordinates with planner and job manager

### 6.2 Planner
- Task decomposition
- Step-by-step execution planning
- Progress tracking

### 6.3 Job Manager
- Background job execution
- Queue management
- Result caching

## 7. Workspace

- Sandboxed execution environment
- File system isolation
- Shell command restrictions
- Resource limits

## 8. Security

- **RCE Protection**: Explicit user confirmation required
- **Auth Isolation**: Provider credentials encrypted at rest
- **Network**: HTTPS only, certificate validation
- **Filesystem**: Workspace sandboxing
- **Shell**: Command whitelist/blacklist

## 9. Build

```bash
# Build for Termux
go build -o agy -ldflags "-s -w" .

# Build for Android aarch64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o agy .
```

## 10. Testing

- Unit tests for all packages
- Integration tests for providers
- E2E tests for CLI commands
- Termux-specific tests