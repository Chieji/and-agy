package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme provides styling for the TUI
type Theme struct {
	Primary     lipgloss.Style
	Secondary   lipgloss.Style
	Success     lipgloss.Style
	Warning     lipgloss.Style
	Error       lipgloss.Style
	Title       lipgloss.Style
	Subtitle    lipgloss.Style
	Border      lipgloss.Style
	Input       lipgloss.Style
	Button      lipgloss.Style
	StatusBar   lipgloss.Style
	ChatBubble  lipgloss.Style
	CodeBlock   lipgloss.Style
}

// DefaultTheme returns the default theme
func DefaultTheme() *Theme {
	return &Theme{
		Primary: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Bold(true),
		
		Secondary: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A0A0A0")),
		
		Success: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4CAF50")).
			Bold(true),
		
		Warning: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFC107")).
			Bold(true),
		
		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F44336")).
			Bold(true),
		
		Title: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#2196F3")).
			Bold(true).
			Underline(true),
		
		Subtitle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#64B5F6")),
		
		Border: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#BDBDBD")),
		
		Input: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#424242")),
		
		Button: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#2196F3")).
			Bold(true),
		
		StatusBar: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#616161")),
		
		ChatBubble: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#424242")).
			Padding(1, 2),
			BorderStyle(lipgloss.RoundedBorder()),
		
		CodeBlock: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFC107")).
			Background(lipgloss.Color("#263238")).
			Padding(1, 2),
			BorderStyle(lipgloss.RoundedBorder()),
	}
}

// NewTheme creates a new custom theme
func NewTheme() *Theme {
	return DefaultTheme()
}

// WithPrimary sets the primary color
func (t *Theme) WithPrimary(color string) *Theme {
	t.Primary = t.Primary.Foreground(lipgloss.Color(color))
	return t
}

// WithSecondary sets the secondary color
func (t *Theme) WithSecondary(color string) *Theme {
	t.Secondary = t.Secondary.Foreground(lipgloss.Color(color))
	return t
}