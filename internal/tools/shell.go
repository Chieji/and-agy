package tools

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Shell provides shell command execution with safety checks
type Shell struct {
	workspace   string
	allowedCmds []string
	blockedCmds []string
	timeout     time.Duration
}

// NewShell creates a new shell tool
func NewShell(workspace string) *Shell {
	return &Shell{
		workspace: workspace,
		allowedCmds: []string{
			"ls", "cd", "pwd", "cat", "echo", "grep", "find", "sed", "awk",
			"git", "go", "make", "mkdir", "rm", "cp", "mv", "chmod", "chown",
			"head", "tail", "wc", "sort", "uniq", "diff", "patch",
		},
		blockedCmds: []string{
			"rm -rf", "chmod +x", "wget", "curl", "ssh", "scp", "rsync",
			"sudo", "su", "apt", "pkg", "yum", "dnf", "pacman",
			"dd", "mkfs", "fdisk", "parted", "mount", "umount",
			"kill", "pkill", "killall", "reboot", "shutdown",
			"nc", "netcat", "socat", "telnet", "ftp", "sftp",
			"python", "python3", "ruby", "perl", "php", "node",
		},
		timeout: 30 * time.Second,
	}
}

// Execute executes a shell command with safety checks
func (s *Shell) Execute(ctx context.Context, command string, args ...string) (string, error) {
	fullCmd := command
	if len(args) > 0 {
		fullCmd = fmt.Sprintf("%s %s", command, strings.Join(args, " "))
	}
	
	// Safety check
	if err := s.isSafe(fullCmd); err != nil {
		return "", fmt.Errorf("command blocked: %w", err)
	}
	
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Dir = s.workspace
	
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	if err := cmd.Run(); err != nil {
		return stdout.String(), fmt.Errorf("%w: %s", err, stderr.String())
	}
	
	return stdout.String(), nil
}

// ExecuteWithTimeout executes a command with a timeout
func (s *Shell) ExecuteWithTimeout(command string, timeout time.Duration, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	
	return s.Execute(ctx, command, args...)
}

// isSafe checks if a command is safe to execute
func (s *Shell) isSafe(command string) error {
	lowerCmd := strings.ToLower(command)
	
	// Check blocked commands
	for _, blocked := range s.blockedCmds {
		if strings.Contains(lowerCmd, strings.ToLower(blocked)) {
			return fmt.Errorf("command contains blocked pattern: %s", blocked)
		}
	}
	
	// Extract command name (first word)
	parts := strings.Fields(lowerCmd)
	if len(parts) == 0 {
		return fmt.Errorf("empty command")
	}
	
	cmdName := parts[0]
	
	// Check if command is allowed
	allowed := false
	for _, allowed := range s.allowedCmds {
		if cmdName == strings.ToLower(allowed) {
			allowed = true
			break
		}
	}
	
	if !allowed {
		return fmt.Errorf("command not in allowed list: %s", cmdName)
	}
	
	// Additional safety checks
	if strings.Contains(lowerCmd, "> /dev/") ||
		strings.Contains(lowerCmd, "> /proc/") ||
		strings.Contains(lowerCmd, "> /sys/") ||
		strings.Contains(lowerCmd, "&>") ||
		strings.Contains(lowerCmd, "; ") ||
		strings.Contains(lowerCmd, " | ") ||
		strings.Contains(lowerCmd, "`") ||
		strings.Contains(lowerCmd, "$") {
		return fmt.Errorf("command contains unsafe patterns")
	}
	
	return nil
}

// AddAllowedCommand adds a command to the allowed list
func (s *Shell) AddAllowedCommand(cmd string) {
	s.allowedCmds = append(s.allowedCmds, cmd)
}

// AddBlockedCommand adds a command to the blocked list
func (s *Shell) AddBlockedCommand(cmd string) {
	s.blockedCmds = append(s.blockedCmds, cmd)
}

// SetTimeout sets the default timeout
func (s *Shell) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

// SetWorkspace sets the working directory
func (s *Shell) SetWorkspace(workspace string) {
	s.workspace = workspace
}