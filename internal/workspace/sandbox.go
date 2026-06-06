package workspace

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Filesystem provides filesystem operations
type Filesystem struct {
	workspace string
}

// NewFilesystem creates a new filesystem tool
func NewFilesystem(workspace string) *Filesystem {
	return &Filesystem{workspace: workspace}
}

// ReadFile reads a file from the workspace
func (fs *Filesystem) ReadFile(path string) ([]byte, error) {
	fullPath := filepath.Join(fs.workspace, path)
	return os.ReadFile(fullPath)
}

// WriteFile writes a file to the workspace
func (fs *Filesystem) WriteFile(path string, content []byte) error {
	fullPath := filepath.Join(fs.workspace, path)
	
	// Create directory if it doesn't exist
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	
	return os.WriteFile(fullPath, content, 0644)
}

// Shell provides shell command execution
type Shell struct {
	workspace string
}

// NewShell creates a new shell tool
func NewShell(workspace string) *Shell {
	return &Shell{workspace: workspace}
}

// Execute executes a shell command
func (s *Shell) Execute(command string, args ...string) (string, error) {
	// TODO: Implement actual shell execution with safety checks
	return fmt.Sprintf("Executed: %s %v", command, args), nil
}

// Sandbox provides an isolated workspace for agent execution
type Sandbox struct {
	rootPath   string
	fs         *Filesystem
	shell      *Shell
	isolation  bool
	mu         sync.RWMutex
	cleanup    []string
}

// NewSandbox creates a new sandbox
func NewSandbox(rootPath string, isolation bool) (*Sandbox, error) {
	if err := os.MkdirAll(rootPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create sandbox root: %w", err)
	}
	
	return &Sandbox{
		rootPath:  rootPath,
		fs:        NewFilesystem(rootPath),
		shell:     NewShell(rootPath),
		isolation: isolation,
		cleanup:   make([]string, 0),
	}, nil
}

// CreateWorkspace creates a new isolated workspace
func (s *Sandbox) CreateWorkspace(name string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	workspacePath := filepath.Join(s.rootPath, "workspaces", name)
	
	if s.isolation {
		// Create temporary directory
		tmpDir, err := os.MkdirTemp(s.rootPath, "workspace-")
		if err != nil {
			return "", fmt.Errorf("failed to create temp workspace: %w", err)
		}
		workspacePath = tmpDir
		s.cleanup = append(s.cleanup, tmpDir)
	} else {
		if err := os.MkdirAll(workspacePath, 0755); err != nil {
			return "", fmt.Errorf("failed to create workspace: %w", err)
		}
	}
	
	return workspacePath, nil
}

// DeleteWorkspace deletes a workspace
func (s *Sandbox) DeleteWorkspace(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Remove from cleanup list
	for i, p := range s.cleanup {
		if p == path {
			s.cleanup = append(s.cleanup[:i], s.cleanup[i+1:]...)
			break
		}
	}
	
	return os.RemoveAll(path)
}

// Cleanup removes all temporary workspaces
func (s *Sandbox) Cleanup() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	var errs []error
	for _, path := range s.cleanup {
		if err := os.RemoveAll(path); err != nil {
			errs = append(errs, err)
		}
	}
	
	if len(errs) > 0 {
		return fmt.Errorf("cleanup errors: %v", errs)
	}
	
	return nil
}

// GetFilesystem returns the filesystem tool
func (s *Sandbox) GetFilesystem() *Filesystem {
	return s.fs
}

// GetShell returns the shell tool
func (s *Sandbox) GetShell() *Shell {
	return s.shell
}

// SetIsolation enables or disables workspace isolation
func (s *Sandbox) SetIsolation(isolation bool) {
	s.isolation = isolation
}

// Filesystem wrapper for sandbox
func (s *Sandbox) ReadFile(path string) ([]byte, error) {
	return s.fs.ReadFile(path)
}

// WriteFile wrapper for sandbox
func (s *Sandbox) WriteFile(path string, content []byte) error {
	return s.fs.WriteFile(path, content)
}

// ExecuteCommand wrapper for sandbox
func (s *Sandbox) ExecuteCommand(command string, args ...string) (string, error) {
	return s.shell.Execute(command, args...)
}