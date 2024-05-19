package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Background is a UI component for displaying the background. It implements the tea.Model interface.
type Background struct {
	characters [][]rune
}

// NewBackground returns a new Background.
func NewBackground() Background {
	return Background{}
}

// charactersUpdateMsg is a tea.Msg that updates the characters of the background.
type charactersUpdateMsg struct {
	characters [][]rune
}

// Init implements tea.Model by returning a tea.Cmd that initializes the characters of the background.
func (b Background) Init() tea.Cmd {
	return func() tea.Msg { return initializeCharacters(terminalWidth, terminalHeight) }
}

// Update implements tea.Model by updating the model based on the given message.
func (b Background) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case charactersUpdateMsg:
		b.characters = msg.characters
	case tea.WindowSizeMsg:
		cmd = func() tea.Msg { return initializeCharacters(msg.Width, msg.Height) }
	}

	return b, cmd
}

// View implements tea.Model by returning a string that displays the background.
func (b Background) View() string {
	view := ""

	if b.characters != nil {
		for i, line := range b.characters {
			view += string(line)
			if i < len(b.characters)-1 {
				view += "\n"
			}
		}
	}

	return lipgloss.NewStyle().
		Background(backgroundColor).
		Foreground(inactiveTextColor).
		Render(view)
}

// initializeCharacters initializes the characters of the background and returns them in a charactersUpdateMsg.
func initializeCharacters(width int, height int) tea.Msg {
	characters := make([][]rune, height)
	for i := range characters {
		characters[i] = make([]rune, width)
		for j := range characters[i] {
			characters[i][j] = '$'
		}
	}
	return charactersUpdateMsg{characters: characters}
}
