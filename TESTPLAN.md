# Test Plan for and-agy

## 1. Test Strategy

### 1.1 Test Levels
- Unit tests: Individual functions and methods
- Integration tests: Component interactions
- E2E tests: Complete user workflows
- Termux-specific tests: Platform compatibility

### 1.2 Test Coverage Targets
- Core packages: 90%+
- CLI commands: 100%
- TUI components: 90%+
- Auth providers: 100%

## 2. Test Environment

### 2.1 Platforms
- Termux on Android (primary)
- Linux aarch64 (secondary)
- Linux x86_64 (development)

### 2.2 Dependencies
- Go 1.21+
- Git
- Termux API (for Termux-specific features)

## 3. Unit Tests

### 3.1 internal/agent
- [ ] harness_test.go: Test agent lifecycle
- [ ] planner_test.go: Test task decomposition
- [ ] jobmanager_test.go: Test job queue

### 3.2 internal/auth
- [ ] gemini_test.go: Test Gemini auth
- [ ] provider_test.go: Test provider interface

### 3.3 internal/cli
- [ ] root_test.go: Test root command
- [ ] agent_test.go: Test agent command
- [ ] config_test.go: Test config command
- [ ] chat_test.go: Test chat command

### 3.4 internal/tools
- [ ] filesystem_test.go: Test FS operations
- [ ] shell_test.go: Test shell execution

### 3.5 internal/workspace
- [ ] sandbox_test.go: Test sandbox isolation

### 3.6 pkg/tui
- [ ] app_test.go: Test TUI app
- [ ] theme_test.go: Test styling
- [ ] views/*_test.go: Test individual views

## 4. Integration Tests

### 4.1 Auth System
- [ ] Test provider registration
- [ ] Test authentication flow
- [ ] Test token persistence
- [ ] Test provider switching

### 4.2 Agent System
- [ ] Test agent initialization
- [ ] Test tool execution
- [ ] Test job management
- [ ] Test progress tracking

### 4.3 TUI System
- [ ] Test view transitions
- [ ] Test input handling
- [ ] Test rendering
- [ ] Test keybindings

## 5. E2E Tests

### 5.1 CLI Commands
- [ ] agy version
- [ ] agy help
- [ ] agy chat
- [ ] agy agent
- [ ] agy config get/set/list
- [ ] agy auth login/list/logout/default

### 5.2 User Workflows
- [ ] First-time setup
- [ ] Provider authentication
- [ ] Chat session
- [ ] Agent task execution
- [ ] Configuration management

## 6. Termux-Specific Tests

### 6.1 Platform Compatibility
- [ ] Termux environment detection
- [ ] Termux API integration
- [ ] Storage permissions
- [ ] Network connectivity

### 6.2 Build Verification
- [ ] Pure Go build (CGO_ENABLED=0)
- [ ] aarch64 compatibility
- [ ] Static linking
- [ ] Binary size optimization

## 7. Security Tests

### 7.1 RCE Protection
- [ ] Shell command validation
- [ ] User confirmation prompts
- [ ] Dangerous command blocking
- [ ] Workspace isolation

### 7.2 Auth Security
- [ ] Credential encryption
- [ ] Token storage security
- [ ] Network security (HTTPS)
- [ ] Certificate validation

## 8. Performance Tests

### 8.1 Startup Time
- [ ] CLI startup < 500ms
- [ ] TUI startup < 1s

### 8.2 Memory Usage
- [ ] Idle < 50MB
- [ ] Chat session < 100MB
- [ ] Agent task < 200MB

### 8.3 Response Time
- [ ] Command response < 100ms
- [ ] Chat response < 5s (depends on provider)

## 9. Test Execution

### 9.1 Local Development
```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -run TestHarness ./internal/agent
```

### 9.2 CI/CD
- GitHub Actions for PR testing
- Termux-specific workflows
- Release builds for all platforms

### 9.3 Termux Testing
```bash
# Install dependencies
pkg install git golang

# Clone and test
git clone https://github.com/Chieji/and-agy.git
cd and-agy
go test ./...
```

## 10. Test Reporting

- JUnit XML format for CI
- Coverage reports
- Performance metrics
- Security scan results

## 11. Acceptance Criteria

- [ ] All unit tests pass
- [ ] All integration tests pass
- [ ] All E2E tests pass on Termux
- [ ] No critical security issues
- [ ] Performance targets met
- [ ] Documentation complete