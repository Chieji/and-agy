package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbletea"
)

// App is the main TUI application
type App struct {
	teaProgram *tea.Program
	currentView string
	views map[string]View
}

// View interface for all TUI views
type View interface {
	Init() tea.Model
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
	Name() string
}

// NewApp creates a new TUI application
func NewApp() *App {
	return &App{
		views: make(map[string]View),
	}
}

// AddView adds a view to the application
func (a *App) AddView(view View) {
	a.views[view.Name()] = view
}

// SwitchView switches to a different view
func (a *App) SwitchView(name string) {
	a.currentView = name
}

// Run starts the TUI application
func (a *App) Run() error {
	if len(a.views) == 0 {
		return fmt.Errorf("no views registered")
	}
	
	// Start with the first view
	if a.currentView == "" {
		for name := range a.views {
			a.currentView = name
			break
		}
	}
	
	view, exists := a.views[a.currentView]
	if !exists {
		return fmt.Errorf("view not found: %s", a.currentView)
	}
	
	model := view.Init()
	a.teaProgram = tea.NewProgram(model)
	
	_, err := a.teaProgram.Run()
	return err
}

// Send sends a message to the current view
func (a *App) Send(msg tea.Msg) {
	if a.teaProgram != nil {
		a.teaProgram.Send(msg)
	}
}

// Quit stops the TUI application
func (a *App) Quit() {
	if a.teaProgram != nil {
		a.teaProgram.Quit()
	}
}