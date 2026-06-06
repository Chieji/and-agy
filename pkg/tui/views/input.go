package views

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
)

// InputView is a generic input prompt
type InputView struct {
	prompt   string
	input    string
	cursor   int
	callback func(string)
}

// NewInputView creates a new input view
func NewInputView(prompt string, callback func(string)) *InputView {
	return &InputView{
		prompt:   prompt,
		input:    "",
		cursor:   0,
		callback: callback,
	}
}

// Init initializes the input view
func (iv *InputView) Init() tea.Model {
	return iv
}

// Update handles messages
func (iv *InputView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return iv, tea.Quit
		case "enter":
			if iv.callback != nil {
				iv.callback(iv.input)
			}
			return iv, tea.Quit
		case "left":
			if iv.cursor > 0 {
				iv.cursor--
			}
		case "right":
			if iv.cursor < len(iv.input) {
				iv.cursor++
			}
		case "backspace":
			if iv.cursor > 0 {
				iv.input = iv.input[:iv.cursor-1] + iv.input[iv.cursor:]
				iv.cursor--
			}
		case "delete":
			if iv.cursor < len(iv.input) {
				iv.input = iv.input[:iv.cursor] + iv.input[iv.cursor+1:]
			}
		default:
			if len(msg.String()) == 1 {
				iv.input = iv.input[:iv.cursor] + msg.String() + iv.input[iv.cursor:]
				iv.cursor++
			}
		}
	}
	
	return iv, nil
}

// View renders the input view
func (iv *InputView) View() string {
	var sb strings.Builder
	
	// Prompt
	sb.WriteString(iv.prompt)
	sb.WriteString(": ")
	sb.WriteString(iv.input)
	sb.WriteString("\n")
	
	// Help
	sb.WriteString("Press enter to submit, ctrl+c to cancel\n")
	
	return sb.String()
}

// Name returns the view name
func (iv *InputView) Name() string {
	return "input"
}

// SetPrompt sets the prompt text
func (iv *InputView) SetPrompt(prompt string) {
	iv.prompt = prompt
}

// SetInput sets the input text
func (iv *InputView) SetInput(text string) {
	iv.input = text
	iv.cursor = len(text)
}