package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SummoningCircle is a UI component for displaying the summoning circle. It implements the tea.Model interface.
type SummoningCircle struct {
	summoningMessage string
	terminalWidth    int
	terminalHeight   int
}

// NewSummoningCircle creates a new SummoningCircle.
func NewSummoningCircle(summoningMessage string, terminalWidth int, terminalHeight int) SummoningCircle {
	return SummoningCircle{
		summoningMessage: summoningMessage,
		terminalWidth:    terminalWidth,
		terminalHeight:   terminalHeight,
	}
}

// Init implements the tea.Model interface by returning nil.
func (c *SummoningCircle) Init() tea.Cmd {
	return nil
}

// Update implements the tea.Model interface by returning nil.
func (c *SummoningCircle) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		c.terminalWidth = msg.Width
		c.terminalHeight = msg.Height
	}
	return c, nil
}

// View implements the tea.Model interface by returning a string that displays the summoning circle.
func (c *SummoningCircle) View() string {
	messageLength := len([]rune(c.summoningMessage))
	style := lipgloss.NewStyle().
		Bold(true).
		PaddingTop(c.terminalHeight / 2).
		PaddingLeft((c.terminalWidth / 2) - (messageLength / 2) - 2)
	return style.Render(c.summoningMessage)
}
