package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
)

// StatusView displays system status and logs
type StatusView struct {
	logs      []string
	cursor    int
	status    string
	provider  string
	model     string
	workspace string
}

// NewStatusView creates a new status view
func NewStatusView() *StatusView {
	return &StatusView{
		logs:      make([]string, 0),
		status:    "Ready",
		provider:  "None",
		model:     "None",
		workspace: "None",
	}
}

// Init initializes the status view
func (sv *StatusView) Init() tea.Model {
	return sv
}

// Update handles messages
func (sv *StatusView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return sv, tea.Quit
		case "up":
			if sv.cursor > 0 {
				sv.cursor--
			}
		case "down":
			if sv.cursor < len(sv.logs)-1 {
				sv.cursor++
			}
		case "c":
			// Clear logs
			sv.logs = make([]string, 0)
			sv.cursor = 0
		}
	}
	
	return sv, nil
}

// View renders the status view
func (sv *StatusView) View() string {
	var sb strings.Builder
	
	// Header
	sb.WriteString("=== System Status ===\n\n")
	
	// Status info
	sb.WriteString(fmt.Sprintf("Status: %s\n", sv.status))
	sb.WriteString(fmt.Sprintf("Provider: %s\n", sv.provider))
	sb.WriteString(fmt.Sprintf("Model: %s\n", sv.model))
	sb.WriteString(fmt.Sprintf("Workspace: %s\n", sv.workspace))
	sb.WriteString("\n")
	
	// Logs
	sb.WriteString("=== Logs ===\n\n")
	
	if len(sv.logs) == 0 {
		sb.WriteString("No logs\n")
	} else {
		start := sv.cursor
		end := sv.cursor + 20
		if end > len(sv.logs) {
			end = len(sv.logs)
		}
		if start < 0 {
			start = 0
		}
		
		for i := start; i < end; i++ {
			cursor := "  "
			if i == sv.cursor {
				cursor = "> "
			}
			sb.WriteString(fmt.Sprintf("%s%s\n", cursor, sv.logs[i]))
		}
	}
	
	// Help
	sb.WriteString("\nPress q to quit, up/down to navigate, c to clear\n")
	
	return sb.String()
}

// Name returns the view name
func (sv *StatusView) Name() string {
	return "status"
}

// AddLog adds a log message
func (sv *StatusView) AddLog(msg string) {
	sv.logs = append(sv.logs, msg)
}

// SetStatus sets the current status
func (sv *StatusView) SetStatus(status string) {
	sv.status = status
}

// SetProvider sets the current provider
func (sv *StatusView) SetProvider(provider string) {
	sv.provider = provider
}

// SetModel sets the current model
func (sv *StatusView) SetModel(model string) {
	sv.model = model
}

// SetWorkspace sets the current workspace
func (sv *StatusView) SetWorkspace(workspace string) {
	sv.workspace = workspace
}