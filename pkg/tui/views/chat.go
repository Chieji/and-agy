package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
)

// ChatView is the chat interface
type ChatView struct {
	messages []string
	input    string
	cursor   int
}

// NewChatView creates a new chat view
func NewChatView() *ChatView {
	return &ChatView{
		messages: make([]string, 0),
		input:    "",
		cursor:   0,
	}
}

// Init initializes the chat view
func (cv *ChatView) Init() tea.Model {
	return cv
}

// Update handles messages
func (cv *ChatView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return cv, tea.Quit
		case "enter":
			if cv.input != "" {
				cv.messages = append(cv.messages, "You: "+cv.input)
				// TODO: Send to AI and get response
				cv.messages = append(cv.messages, "AI: Response to: "+cv.input)
				cv.input = ""
				cv.cursor = 0
			}
		case "up":
			// TODO: Navigate history
		case "down":
			// TODO: Navigate history
		case "left":
			if cv.cursor > 0 {
				cv.cursor--
			}
		case "right":
			if cv.cursor < len(cv.input) {
				cv.cursor++
			}
		case "backspace":
			if cv.cursor > 0 {
				cv.input = cv.input[:cv.cursor-1] + cv.input[cv.cursor:]
				cv.cursor--
			}
		case "delete":
			if cv.cursor < len(cv.input) {
				cv.input = cv.input[:cv.cursor] + cv.input[cv.cursor+1:]
			}
		default:
			if len(msg.String()) == 1 {
				cv.input = cv.input[:cv.cursor] + msg.String() + cv.input[cv.cursor:]
				cv.cursor++
			}
		}
	}
	
	return cv, nil
}

// View renders the chat view
func (cv *ChatView) View() string {
	var sb strings.Builder
	
	// Header
	sb.WriteString("=== Chat ===\n\n")
	
	// Messages
	for _, msg := range cv.messages {
		sb.WriteString(msg + "\n")
	}
	
	if len(cv.messages) > 0 {
		sb.WriteString("\n")
	}
	
	// Input
	sb.WriteString("Input: ")
	sb.WriteString(cv.input)
	sb.WriteString("\n")
	
	// Help
	sb.WriteString("\nPress q to quit, enter to send\n")
	
	return sb.String()
}

// Name returns the view name
func (cv *ChatView) Name() string {
	return "chat"
}

// AddMessage adds a message to the chat
func (cv *ChatView) AddMessage(msg string) {
	cv.messages = append(cv.messages, msg)
}