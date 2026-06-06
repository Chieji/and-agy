package tools

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	
	// Security: Prevent directory traversal
	if !filepath.IsAbs(fullPath) {
		fullPath = filepath.Join(fs.workspace, path)
	}
	
	// Ensure path is within workspace
	if !isWithinWorkspace(fs.workspace, fullPath) {
		return nil, fmt.Errorf("path outside workspace: %s", path)
	}
	
	return os.ReadFile(fullPath)
}

// WriteFile writes a file to the workspace
func (fs *Filesystem) WriteFile(path string, content []byte) error {
	fullPath := filepath.Join(fs.workspace, path)
	
	// Security: Prevent directory traversal
	if !isWithinWorkspace(fs.workspace, fullPath) {
		return fmt.Errorf("path outside workspace: %s", path)
	}
	
	// Create directory if it doesn't exist
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	
	return os.WriteFile(fullPath, content, 0644)
}

// ListFiles lists files in a directory
func (fs *Filesystem) ListFiles(path string) ([]string, error) {
	fullPath := filepath.Join(fs.workspace, path)
	
	if !isWithinWorkspace(fs.workspace, fullPath) {
		return nil, fmt.Errorf("path outside workspace: %s", path)
	}
	
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}
	
	files := make([]string, 0, len(entries))
	for _, entry := range entries {
		files = append(files, entry.Name())
	}
	
	return files, nil
}

// DeleteFile deletes a file
func (fs *Filesystem) DeleteFile(path string) error {
	fullPath := filepath.Join(fs.workspace, path)
	
	if !isWithinWorkspace(fs.workspace, fullPath) {
		return fmt.Errorf("path outside workspace: %s", path)
	}
	
	return os.Remove(fullPath)
}

// CopyFile copies a file
func (fs *Filesystem) CopyFile(src, dst string) error {
	srcPath := filepath.Join(fs.workspace, src)
	dstPath := filepath.Join(fs.workspace, dst)
	
	if !isWithinWorkspace(fs.workspace, srcPath) || !isWithinWorkspace(fs.workspace, dstPath) {
		return fmt.Errorf("path outside workspace")
	}
	
	source, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer source.Close()
	
	destination, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer destination.Close()
	
	_, err = io.Copy(destination, source)
	return err
}

// isWithinWorkspace checks if a path is within the workspace
func isWithinWorkspace(workspace, path string) bool {
	absWorkspace, _ := filepath.Abs(workspace)
	absPath, _ := filepath.Abs(path)
	
	// Use filepath.Rel to check if path is within workspace
	rel, err := filepath.Rel(absWorkspace, absPath)
	if err != nil {
		return false
	}
	
	// If rel doesn't start with "..", it's within workspace
	return !filepath.IsAbs(rel) && !containsDotDot(rel)
}

// containsDotDot checks if a path contains ".."
func containsDotDot(path string) bool {
	parts := filepath.SplitList(path)
	for _, part := range parts {
		if part == ".." {
			return true
		}
	}
	return false
}