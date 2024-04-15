package ui

import tea "github.com/charmbracelet/bubbletea"

// Placeholder is a UI component that displays a placeholder line with text. It implements tea.Model.
type Placeholder struct {
	text string
}

// NewPlaceholder returns a new Placeholder with the given text.
func NewPlaceholder(text string) Placeholder {
	return Placeholder{
		text: text,
	}
}

// Init implements the tea.Model interface by returning nil.
func (p Placeholder) Init() tea.Cmd {
	return nil
}

// Update implements the tea.Model interface by returning nil.
func (p Placeholder) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}

// View implements the tea.Model interface by returning the placeholder text.
func (p Placeholder) View() string {
	return p.text
}
