package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
)

// ConfigView is the configuration interface
type ConfigView struct {
	config     map[string]string
	keys       []string
	cursor     int
	editingKey string
	editingVal string
	mode       string // "view", "edit"
}

// NewConfigView creates a new config view
func NewConfigView() *ConfigView {
	return &ConfigView{
		config: map[string]string{
			"provider": "gemini",
			"model":    "gemini-1.5-pro",
			"theme":    "default",
		},
		keys: make([]string, 0),
		mode: "view",
	}
}

// Init initializes the config view
func (cv *ConfigView) Init() tea.Model {
	// Populate keys
	for k := range cv.config {
		cv.keys = append(cv.keys, k)
	}
	return cv
}

// Update handles messages
func (cv *ConfigView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return cv, tea.Quit
		case "up":
			if cv.cursor > 0 {
				cv.cursor--
			}
		case "down":
			if cv.cursor < len(cv.keys)-1 {
				cv.cursor++
			}
		case "enter":
			if cv.mode == "view" {
				cv.mode = "edit"
				cv.editingKey = cv.keys[cv.cursor]
				cv.editingVal = cv.config[cv.editingKey]
			} else {
				// Save
				cv.config[cv.editingKey] = cv.editingVal
				cv.mode = "view"
				cv.editingKey = ""
				cv.editingVal = ""
			}
		case "escape":
			if cv.mode == "edit" {
				cv.mode = "view"
				cv.editingKey = ""
				cv.editingVal = ""
			}
		case "backspace":
			if cv.mode == "edit" && len(cv.editingVal) > 0 {
				cv.editingVal = cv.editingVal[:len(cv.editingVal)-1]
			}
		default:
			if cv.mode == "edit" && len(msg.String()) == 1 {
				cv.editingVal += msg.String()
			}
		}
	}
	
	return cv, nil
}

// View renders the config view
func (cv *ConfigView) View() string {
	var sb strings.Builder
	
	// Header
	sb.WriteString("=== Configuration ===\n\n")
	
	// Config items
	for i, key := range cv.keys {
		cursor := "  "
		if i == cv.cursor {
			cursor = "> "
		}
		
		if cv.mode == "edit" && key == cv.editingKey {
			sb.WriteString(fmt.Sprintf("%s%s: %s_\n", cursor, key, cv.editingVal))
		} else {
			sb.WriteString(fmt.Sprintf("%s%s: %s\n", cursor, key, cv.config[key]))
		}
	}
	
	// Help
	sb.WriteString("\n")
	if cv.mode == "view" {
		sb.WriteString("Press q to quit, enter to edit, up/down to navigate\n")
	} else {
		sb.WriteString("Press enter to save, escape to cancel\n")
	}
	
	return sb.String()
}

// Name returns the view name
func (cv *ConfigView) Name() string {
	return "config"
}